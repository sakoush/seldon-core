/*
Copyright (c) 2024 Seldon Technologies Ltd.

Use of this software is governed BY
(1) the license included in the LICENSE file or
(2) if the license included in the LICENSE file is the Business Source License 1.1,
the Change License after the Change Date as each is defined in accordance with the LICENSE file.
*/

package io.seldon.dataflow

import com.github.michaelbull.retry.policy.binaryExponentialBackoff
import com.github.michaelbull.retry.retry
import io.grpc.ManagedChannelBuilder
import io.grpc.StatusException
import io.grpc.StatusRuntimeException
import io.seldon.dataflow.kafka.*
import io.seldon.mlops.chainer.ChainerGrpcKt
import io.seldon.mlops.chainer.ChainerOuterClass.*
import io.seldon.mlops.chainer.ChainerOuterClass.PipelineUpdateMessage.PipelineOperation
import kotlinx.coroutines.FlowPreview
import kotlinx.coroutines.async
import kotlinx.coroutines.awaitAll
import kotlinx.coroutines.flow.collect
import kotlinx.coroutines.flow.onCompletion
import kotlinx.coroutines.flow.onEach
import kotlinx.coroutines.runBlocking
import java.util.concurrent.ConcurrentHashMap
import io.klogging.logger as coLogger

@OptIn(FlowPreview::class)
class PipelineSubscriber(
    private val name: String,
    private val kafkaProperties: KafkaProperties,
    kafkaAdminProperties: KafkaAdminProperties,
    kafkaStreamsParams: KafkaStreamsParams,
    private val kafkaDomainParams: KafkaDomainParams,
    private val upstreamHost: String,
    private val upstreamPort: Int,
    grpcServiceConfig: Map<String, Any>,
    private val kafkaConsumerGroupIdPrefix: String,
    private val namespace: String,
) {
    private val kafkaAdmin = KafkaAdmin(kafkaAdminProperties, kafkaStreamsParams)
    private val channel = ManagedChannelBuilder
        .forAddress(upstreamHost, upstreamPort)
        .defaultServiceConfig(grpcServiceConfig)
        .usePlaintext() // Use TLS
        .enableRetry()
        .build()
    private val client = ChainerGrpcKt.ChainerCoroutineStub(channel)

    private val pipelines = ConcurrentHashMap<PipelineId, Pipeline>()

    suspend fun subscribe() {
        logger.info("will connect to ${upstreamHost}:${upstreamPort}")
        while (true) {
            retry(binaryExponentialBackoff(50..5_000L)) {
                subscribePipelines(kafkaConsumerGroupIdPrefix, namespace)
            }
        }
    }

    // TODO
    //  - If a topology encounters an error, we should signal back to the scheduler about this.
    //  - If the scheduler updates/removes a topology, we need to cancel the corresponding coroutine.
    //  ...
    //  Pipeline UID should be enough to uniquely key it, even across versions?
    //  ...
    //  - Add map of model name -> (weak) referrents/reference count to avoid recreation of streams
    private suspend fun subscribePipelines(kafkaConsumerGroupIdPrefix: String, namespace: String) {
        logger.info("Subscribing to pipeline updates")
        client
            .subscribePipelineUpdates(request = makeSubscriptionRequest())
            .onEach { update ->
                logger.info("received request for ${update.pipeline}:${update.version} Id:${update.uid}")

                val metadata = PipelineMetadata(
                    id = update.uid,
                    name = update.pipeline,
                    version = update.version,
                )

                when (update.op) {
                    PipelineOperation.Create -> handleCreate(metadata, update.updatesList, kafkaConsumerGroupIdPrefix, namespace)
                    PipelineOperation.Delete -> handleDelete(metadata)
                    else -> logger.warn("unrecognised pipeline operation (${update.op})")
                }
            }
            .onCompletion { cause ->
                if (cause == null) {
                    logger.info("pipeline subscription completed successfully")
                } else {
                    logger.error("pipeline subscription terminated with error ${cause}")
                }
            }
            .collect()
        // TODO - error handling?
        // TODO - use supervisor job(s) for spawning coroutines?
    }

    private fun makeSubscriptionRequest() =
        PipelineSubscriptionRequest
            .newBuilder()
            .setName(name)
            .build()

    private suspend fun handleCreate(
        metadata: PipelineMetadata,
        steps: List<PipelineStepUpdate>,
        kafkaConsumerGroupIdPrefix: String,
        namespace: String,
    ) {
        logger.info("Create pipeline ${metadata.name} version: ${metadata.version} id: ${metadata.id}")
        val pipeline = Pipeline.forSteps(metadata, steps, kafkaProperties, kafkaDomainParams, kafkaConsumerGroupIdPrefix, namespace)
        if (pipeline.size != steps.size) {
            client.pipelineUpdateEvent(
                makePipelineUpdateEvent(
                    metadata = metadata,
                    operation = PipelineOperation.Create,
                    success = false,
                    reason = "failed to create all pipeline steps"
                )
            )

            return
        }

        val previous = pipelines.putIfAbsent(metadata.id, pipeline)
        var pipelineStarted = false
        var updateEventReason = "created pipeline"
        if (previous == null) {
            kafkaAdmin.ensureTopicsExist(steps)
            pipelineStarted = pipeline.start()
            if (pipelineStarted == false) {
                updateEventReason = "kafka topic error"
            }
        } else {
            logger.warn("pipeline ${metadata.id} already exists")
        }

        client.pipelineUpdateEvent(
            makePipelineUpdateEvent(
                metadata = metadata,
                operation = PipelineOperation.Create,
                success = pipelineStarted,
                reason = updateEventReason
            )
        )
    }

    private suspend fun handleDelete(metadata: PipelineMetadata) {
        logger.info("Delete pipeline ${metadata.name} version: ${metadata.version} id: ${metadata.id}")
        pipelines
            .remove(metadata.id)
            ?.also { pipeline ->
                runBlocking {
                    pipeline.stop()
                }
            }
        client.pipelineUpdateEvent(
            makePipelineUpdateEvent(
                metadata = metadata,
                operation = PipelineOperation.Delete,
                success = true,
                reason = "pipeline removed",
            )
        )
    }

    fun cancelPipelines(reason: String) {
        runBlocking {
            logger.info("cancelling pipelines due to: $reason")
            pipelines.values
                .map { pipeline ->
                    async { pipeline.stop() }
                }
                .awaitAll()
        }
    }

    private fun makePipelineUpdateEvent(
        metadata: PipelineMetadata,
        operation: PipelineOperation,
        success: Boolean,
        reason: String = "",
    ): PipelineUpdateStatusMessage {
        return PipelineUpdateStatusMessage
            .newBuilder()
            .setSuccess(success)
            .setReason(reason)
            .setUpdate(
                PipelineUpdateMessage
                    .newBuilder()
                    .setOp(operation)
                    .setPipeline(metadata.name)
                    .setVersion(metadata.version)
                    .setUid(metadata.id)
                    .build()
            )
            .build()
    }

    companion object {
        private val logger = coLogger(PipelineSubscriber::class)
    }
}