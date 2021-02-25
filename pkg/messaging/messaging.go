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
	"github.com/SENERGY-Platform/external-task-worker/lib/messages"
	"github.com/SENERGY-Platform/external-task-worker/util"
	"mgw-process-task-worker/pkg/configuration"
)

type Factory struct {
	Config      configuration.Config
	Correlation Correlation
	Incidents   Incidents
}

type Incidents interface {
	Handle(command messages.KafkaIncidentsCommand) error
}

func (this Factory) NewConsumer(_ util.Config, listener func(msg string) error) (consumer kafka.ConsumerInterface, err error) {
	return this.newConsumer(listener)
}

func (this Factory) NewProducer(_ util.Config) (kafka.ProducerInterface, error) {
	return this.newProducer()
}

func (this Factory) newConsumer(listener func(msg string) error) (kafka.ConsumerInterface, error) {
	result := &Consumer{config: this.Config, listener: listener, correlation: this.Correlation}
	err := result.start()
	return result, err
}

func (this Factory) newProducer() (kafka.ProducerInterface, error) {
	result := &Producer{config: this.Config, correlation: this.Correlation, incidents: this.Incidents}
	err := result.start()
	return result, err
}
