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
	"errors"
	"github.com/SENERGY-Platform/external-task-worker/lib/messages"
	paho "github.com/eclipse/paho.mqtt.golang"
	uuid "github.com/satori/go.uuid"
	"log"
	"mgw-process-task-worker/pkg/configuration"
	"time"
)

type Producer struct {
	config      configuration.Config
	correlation Correlation
	mqtt        paho.Client
}

func (this *Producer) start() error {
	options := paho.NewClientOptions().
		SetAutoReconnect(true).
		SetCleanSession(true).
		AddBroker(this.config.MqttBroker).
		SetResumeSubs(true).
		SetConnectionLostHandler(func(_ paho.Client, err error) {
			log.Println("producer to mqtt broker lost")
		}).
		SetOnConnectHandler(func(m paho.Client) {
			log.Println("producer connected to mqtt broker")
		})

	this.mqtt = paho.NewClient(options)
	if token := this.mqtt.Connect(); token.Wait() && token.Error() != nil {
		log.Println("Error on MqttStart.Connect(): ", token.Error())
		return token.Error()
	}
	return nil
}

func (this *Producer) Produce(topic string, message string) (err error) {
	convertetTopic, convertedMessage, err := this.convert(topic, message)
	if err != nil {
		return err
	}
	if convertetTopic != "" {
		token := this.mqtt.Publish(convertetTopic, 2, false, convertedMessage)
		if !token.WaitTimeout(10 * time.Second) {
			return errors.New("publish timeout")
		}
		return token.Error()
	}
	return nil
}

func (this *Producer) ProduceWithKey(topic string, _ string, message string) (err error) {
	return this.Produce(topic, message)
}

func (this *Producer) Close() {
	this.mqtt.Disconnect(0)
}

func (this *Producer) Log(logger *log.Logger) {
	//ignore
}

const ProtocolTopic = "protocol"
const IncidentTopic = "incidents"

func (this *Producer) convert(topic string, message string) (resultTopic string, resultMessage []byte, err error) {
	switch topic {
	case IncidentTopic:
		return this.handleIncident(message)
	case ProtocolTopic:
		return this.convertProtocolMessage(message)
	default:
		log.Println("WARNING: usage of unsupported topic/protocol", topic)
		return //scip mqtt produce because resultTopic remains ""
	}
}

func (this *Producer) handleIncident(message string) (mqttTopic string, mqttMessage []byte, err error) {
	//TODO
	panic("not implemented")
}

func (this *Producer) convertProtocolMessage(message string) (mqttTopic string, mqttMessage []byte, err error) {
	source := messages.ProtocolMsg{}
	err = json.Unmarshal([]byte(message), &source)
	if err != nil {
		return
	}
	if source.Request.Input == nil {
		source.Request.Input = map[string]string{}
	}
	correlationId := this.config.CorrelationIdPrefix + uuid.NewV4().String()
	err = this.correlation.Set(correlationId, source)
	if err != nil {
		return
	}

	mqttTopic = "command/" + source.Metadata.Device.LocalId + "/" + source.Metadata.Service.LocalId

	data := source.Request.Input[StandardProtocolPayloadMessageSegment]
	target := Command{
		CommandId: correlationId,
		Data:      data,
	}
	mqttMessage, err = json.Marshal(target)

	return
}
