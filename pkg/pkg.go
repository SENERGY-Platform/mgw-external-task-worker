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
	"mgw-process-task-worker/pkg/camunda"
	"mgw-process-task-worker/pkg/configuration"
	"mgw-process-task-worker/pkg/devicerepo"
	"mgw-process-task-worker/pkg/marshaller"
	"mgw-process-task-worker/pkg/messaging"
)

func Start(ctx context.Context, config configuration.Config) {
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
			SequentialGroups:                config.SequentialGroups,
		},
		messaging.Factory{Config: config, Correlation: messaging.DefaultCorrelation},
		devicerepo.Factory{Config: config},
		camunda.Factory{Config: config},
		marshaller.Factory{Config: config},
	)
}
