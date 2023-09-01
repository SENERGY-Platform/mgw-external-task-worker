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
	"encoding/json"
	"github.com/SENERGY-Platform/external-task-worker/lib/messages"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/configuration"
	paho "github.com/eclipse/paho.mqtt.golang"
	"log"
	"strings"
)

type Consumer struct {
	config      configuration.Config
	correlation Correlation
	listener    func(msg string) error
	mqtt        paho.Client
}

func (this *Consumer) Stop() {
	this.mqtt.Disconnect(0)
}

func (this *Consumer) start() error {
	options := paho.NewClientOptions().
		SetAutoReconnect(true).
		SetCleanSession(true).
		AddBroker(this.config.MqttBroker).
		SetResumeSubs(true).
		SetConnectionLostHandler(func(_ paho.Client, err error) {
			log.Println("consumer connection to mqtt broker lost")
		}).
		SetOnConnectHandler(func(m paho.Client) {
			log.Println("consumer connected to mqtt broker")
			this.subscribe()
		})

	this.mqtt = paho.NewClient(options)
	if token := this.mqtt.Connect(); token.Wait() && token.Error() != nil {
		log.Println("Error on MqttStart.Connect(): ", token.Error())
		return token.Error()
	}

	return nil
}

func (this *Consumer) subscribe() {
	this.mqtt.Subscribe(this.getResponseTopic(), 2, func(client paho.Client, message paho.Message) {
		msg := Command{}
		err := json.Unmarshal(message.Payload(), &msg)
		if err != nil {
			log.Println("ERROR: unable to unmarshal response to mgw command wrapper", err)
			return
		}
		if strings.HasPrefix(msg.CommandId, this.config.CorrelationIdPrefix) {
			convertedMsg, err := this.convertMessage(msg)
			if err != nil {
				log.Println("ERROR: unable to convert response", err)
				return
			}
			go func() {
				err = this.listener(convertedMsg)
				if err != nil {
					log.Println("ERROR: unable to handle response", err)
					return
				}
			}()
		}
	})
}

func (this *Consumer) getResponseTopic() string {
	return this.config.MqttResponseTopic
}

func (this *Consumer) convertMessage(msg Command) (string, error) {
	target, err := this.getCorrelatedTask(msg.CommandId)
	if err != nil {
		return "", err
	}
	if target.Response.Output == nil {
		target.Response.Output = map[string]string{}
	}
	target.Response.Output[this.config.ProtocolSegment] = msg.Data
	temp, err := json.Marshal(target)
	return string(temp), err
}

func (this *Consumer) getCorrelatedTask(id string) (messages.ProtocolMsg, error) {
	return this.correlation.Get(id)
}
