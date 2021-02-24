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

package messaging

import (
	"github.com/SENERGY-Platform/external-task-worker/lib/kafka"
	"github.com/SENERGY-Platform/external-task-worker/util"
	"mgw-process-task-worker/pkg/configuration"
)

type Factory struct {
	Config      configuration.Config
	Correlation Correlation
}

func (this Factory) NewConsumer(_ util.Config, listener func(msg string) error) (consumer kafka.ConsumerInterface, err error) {
	return this.newConsumer(this.Config, listener)
}

func (this Factory) NewProducer(_ util.Config) (kafka.ProducerInterface, error) {
	return this.newProducer(this.Config)
}

func (this Factory) newConsumer(config configuration.Config, listener func(msg string) error) (kafka.ConsumerInterface, error) {
	result := &Consumer{config: config, listener: listener, correlation: this.Correlation}
	err := result.start()
	return result, err
}

func (this Factory) newProducer(config configuration.Config) (kafka.ProducerInterface, error) {
	result := &Producer{config: config, correlation: this.Correlation}
	err := result.start()
	return result, err
}
