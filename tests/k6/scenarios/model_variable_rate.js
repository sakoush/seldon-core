import {inferHttp, inferGrpc, connectV2Grpc, disconnectV2Grpc} from '../components/v2.js'
import {generateModel} from '../components/model.js'
import {getConfig} from '../components/settings.js'
import {randomIntBetween} from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';
import { sleep } from 'k6';

export const options = {
    noVUConnectionReuse: true,
    noConnectionReuse: true,
    scenarios: {
        variable_request_rate: {
            executor: 'ramping-arrival-rate',
            startRate: getConfig().requestRate,
            timeUnit: '1s',
            //duration: getConfig().constantRateDurationSeconds.toString()+'s',
            stages: [
                { target: 200, duration: '10s' },
        
                { target: 10, duration: '10s' },
        
                { target: 200, duration: '10s' },

                { target: 10, duration: '10s' },

                { target: 200, duration: '10s' },
        
                { target: 10, duration: '10s' },
        
                { target: 300, duration: '10s' },

                { target: 10, duration: '10s' },

                { target: 200, duration: '10s' },
        
                { target: 10, duration: '10s' },

                { target: 200, duration: '10s' },
        
                { target: 10, duration: '10s' },
        
                { target: 300, duration: '10s' },

                { target: 10, duration: '10s' },

                { target: 300, duration: '10s' },
        
                { target: 10, duration: '10s' },
        
                { target: 200, duration: '10s' },

                { target: 10, duration: '10s' },

                { target: 200, duration: '10s' },
        
                { target: 10, duration: '10s' },
        
            ],
            preAllocatedVUs: 10, // how large the initial pool of VUs would be
            maxVUs: 2000, // if the preAllocatedVUs are not enough, we can initialize more
        },
    },
    setupTimeout: '6000s',
    teardownTimeout: '6000s',
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
    sleep(randomIntBetween(1, 5));
}
