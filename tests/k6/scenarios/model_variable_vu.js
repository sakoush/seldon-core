import {inferHttp, inferGrpc, connectV2Grpc, disconnectV2Grpc} from '../components/v2.js'
import {generateModel} from '../components/model.js'
import {getConfig} from '../components/settings.js'
import {randomIntBetween} from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

export const options = {
    discardResponseBodies: true,
    scenarios: {
      contacts: {
        executor: 'ramping-vus',
        startVUs: 0,
        stages: [
          { duration: '2s', target: 10 },
          { duration: '2s', target: 15 },
          { duration: '2s', target: 22 },
          { duration: '2s', target: 10 },
          { duration: '2s', target: 15 },
          { duration: '2s', target: 10 },
          { duration: '2s', target: 20 },
          { duration: '2s', target: 15 },
          { duration: '2s', target: 12 },
          { duration: '2s', target: 20 },
          { duration: '2s', target: 15 },
          { duration: '2s', target: 10 },
        ],
        gracefulRampDown: '0s',
      },
    },
  };

export function setup() {
    return getConfig()
}

export default function (config) {
    // only assume one model type in this scenario
    const idx = 0
    const endIdx = (config.modelEndIdx > 0) ? config.modelEndIdx : config.maxNumModels[idx]  
    const modelIdx = randomIntBetween(config.modelStartIdx, endIdx)
    const modelName = config.modelNamePrefix[idx] + modelIdx.toString()
    const model = generateModel(config.modelType[idx], modelName, 0, 1,
        config.isSchedulerProxy, config.modelMemoryBytes[idx], config.inferBatchSize[idx])
    const httpEndpoint = config.inferHttpEndpoint
    const grpcEndpoint = config.inferGrpcEndpoint

    if (config.inferType === "REST") {
        if (config.modelName !== "") {
            inferHttp(httpEndpoint, config.modelName, model.inference.http, config.isEnvoy, "")
        } else {
            inferHttp(httpEndpoint, modelName, model.inference.http, config.isEnvoy, "")
        }
    } else {
        connectV2Grpc(grpcEndpoint)
        if (config.modelName !== "") {
            inferGrpc(config.modelName, model.inference.grpc, config.isEnvoy, "")
        } else {
            inferGrpc(modelName, model.inference.grpc, config.isEnvoy, "")
        }
        disconnectV2Grpc()
    }
}
