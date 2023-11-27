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
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/camunda"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/configuration"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/devicerepo"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/marshaller"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/messaging"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/timescale"
	"github.com/SENERGY-Platform/service-commons/pkg/cache"
	fallback "github.com/SENERGY-Platform/service-commons/pkg/cache/fallback"
	"log"
)

func Start(ctx context.Context, config configuration.Config) {
	c, err := cache.New(cache.Config{FallbackProvider: fallback.NewProvider(config.FallbackFile)})
	if err != nil {
		log.Fatal(err)
		return
	}
	iotProvider := &devicerepo.Provider{Config: config, Cache: c}
	scheduler := util.PARALLEL
	if config.SequentialGroups {
		scheduler = util.SEQUENTIAL
	}

	lib.Worker(
		ctx,
		util.Config{
			CompletionStrategy:                config.CompletionStrategy,
			OptimisticTaskCompletionTimeout:   config.OptimisticTaskCompletionTimeout,
			CamundaWorkerTimeout:              config.CamundaWorkerTimeout,
			CamundaWorkerTasks:                config.CamundaWorkerTasks,
			CamundaFetchLockDuration:          config.CamundaFetchLockDuration,
			CamundaTopic:                      config.CamundaTopic,
			CamundaTaskResultName:             config.CamundaTaskResultName,
			AuthExpirationTimeBuffer:          config.AuthExpirationTimeBuffer,
			AuthEndpoint:                      config.AuthEndpoint,
			AuthClientId:                      config.AuthClientId,
			KafkaIncidentTopic:                messaging.IncidentTopic,
			Debug:                             config.Debug,
			SubResultExpirationInSeconds:      config.SubResultExpirationInSeconds,
			GroupScheduler:                    scheduler,
			TimescaleWrapperUrl:               config.TimescaleWrapperUrl,
			HandleMissingLastValueTimeAsError: true, //because this runs in the mgw environment
		},
		messaging.Factory{Config: config, Correlation: messaging.DefaultCorrelation, IdProvider: configuration.DefaultIdProvider},
		iotProvider,
		camunda.Factory{Config: config},
		marshaller.Factory{Config: config, DeviceRepo: iotProvider.GetImpl()},
		timescale.Factory,
	)
}
