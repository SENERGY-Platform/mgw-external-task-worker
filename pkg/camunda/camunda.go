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

package camunda

import (
	"github.com/SENERGY-Platform/external-task-worker/lib/camunda"
	"github.com/SENERGY-Platform/external-task-worker/lib/camunda/interfaces"
	"github.com/SENERGY-Platform/external-task-worker/lib/com"
	"github.com/SENERGY-Platform/external-task-worker/util"
	"mgw-external-task-worker/pkg/configuration"
)

type Factory struct {
	Config configuration.Config
}

type Shards string

func (this Shards) GetShards() (result []string, err error) {
	return []string{string(this)}, nil
}

func (this Shards) GetShardForUser(_ string) (shardUrl string, err error) {
	return string(this), nil
}

func (this Factory) Get(config util.Config, producer com.ProducerInterface) (interfaces.CamundaInterface, error) {
	return camunda.NewCamundaWithShards(config, producer, Shards(this.Config.CamundaUrl)), nil
}
