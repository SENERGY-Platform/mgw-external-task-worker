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
	"context"
	"github.com/SENERGY-Platform/external-task-worker/lib/com"
	"github.com/SENERGY-Platform/external-task-worker/lib/messages"
	"github.com/SENERGY-Platform/external-task-worker/util"
	"mgw-external-task-worker/pkg/configuration"
)

type Factory struct {
	Config      configuration.Config
	Correlation Correlation
	IdProvider  func() string
}

type Incidents interface {
	Handle(command messages.KafkaIncidentsCommand) error
}

func (this Factory) NewConsumer(ctx context.Context, config util.Config, respoinseListener func(msg string) error, errorListener func(msg string) error) (err error) {
	result := &Consumer{config: this.Config, listener: respoinseListener, correlation: this.Correlation}
	err = result.start()
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		result.Stop()
	}()
	return err
}

func (this Factory) NewProducer(ctx context.Context, config util.Config) (com.ProducerInterface, error) {
	result := &Producer{config: this.Config, correlation: this.Correlation, idProvider: this.IdProvider}
	err := result.start()
	if err != nil {
		return result, err
	}
	go func() {
		<-ctx.Done()
		result.Close()
	}()
	return result, err
}
