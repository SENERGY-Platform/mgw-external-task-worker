/*
 * Copyright 2021 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pkg

import (
	"context"
	"github.com/SENERGY-Platform/external-task-worker/lib"
	"github.com/SENERGY-Platform/external-task-worker/util"
	"log"
	"mgw-external-task-worker/pkg/camunda"
	"mgw-external-task-worker/pkg/configuration"
	"mgw-external-task-worker/pkg/devicerepo"
	"mgw-external-task-worker/pkg/marshaller"
	"mgw-external-task-worker/pkg/messaging"
	"mgw-external-task-worker/pkg/timescale"
)

func Start(ctx context.Context, config configuration.Config) {
	fallback, err := devicerepo.NewFallback(config.FallbackFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	cache := devicerepo.NewCache(fallback)

	iotProvider := &devicerepo.Provider{Config: config, Cache: cache}
	scheduler := util.PARALLEL
	if config.SequentialGroups {
		scheduler = util.SEQUENTIAL
	}

	lib.Worker(
		ctx,
		util.Config{
			CompletionStrategy:              config.CompletionStrategy,
			OptimisticTaskCompletionTimeout: config.OptimisticTaskCompletionTimeout,
			CamundaWorkerTimeout:            config.CamundaWorkerTimeout,
			CamundaWorkerTasks:              config.CamundaWorkerTasks,
			CamundaFetchLockDuration:        config.CamundaFetchLockDuration,
			CamundaTopic:                    config.CamundaTopic,
			CamundaTaskResultName:           config.CamundaTaskResultName,
			AuthExpirationTimeBuffer:        config.AuthExpirationTimeBuffer,
			AuthEndpoint:                    config.AuthEndpoint,
			AuthClientId:                    config.AuthClientId,
			KafkaIncidentTopic:              messaging.IncidentTopic,
			Debug:                           config.Debug,
			SubResultExpirationInSeconds:    config.SubResultExpirationInSeconds,
			GroupScheduler:                  scheduler,
			TimescaleWrapperUrl:             config.TimescaleWrapperUrl,
		},
		messaging.Factory{Config: config, Correlation: messaging.DefaultCorrelation, IdProvider: configuration.DefaultIdProvider},
		iotProvider,
		camunda.Factory{Config: config},
		marshaller.Factory{Config: config, DeviceRepo: iotProvider.GetImpl()},
		timescale.Factory,
	)
}
