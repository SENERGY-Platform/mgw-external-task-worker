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

package incidents

import (
	"github.com/SENERGY-Platform/external-task-worker/lib/messages"
	paho "github.com/eclipse/paho.mqtt.golang"
	"log"
	"mgw-external-task-worker/pkg/configuration"
	"time"
)

func NewWithMqttClient(config configuration.Config, client paho.Client) (result *Incidents) {
	result = &Incidents{
		config: config,
	}
	result.mqtt = client
	return
}

func New(config configuration.Config) (result *Incidents) {
	retryInterval, err := time.ParseDuration(config.SyncConnectRetryInterval)
	if err != nil {
		log.Println("WARNING: unable to parse SyncConnectRetryInterval; use 1m instead")
		retryInterval = time.Minute
	}
	options := paho.NewClientOptions().
		SetAutoReconnect(true).
		SetCleanSession(true).
		AddBroker(config.SyncMqttBroker).
		SetPassword(config.AuthPassword).
		SetUsername(config.AuthUserName).
		SetConnectRetry(true).
		SetConnectRetryInterval(retryInterval).
		SetResumeSubs(true).
		SetConnectionLostHandler(func(_ paho.Client, err error) {
			log.Println("connection to sync mqtt broker lost")
		}).
		SetOnConnectHandler(func(m paho.Client) {
			log.Println("connected to sync mqtt broker")
		})

	client := paho.NewClient(options)
	client.Connect() //dont wait for token because broker mai not be reachable and the ConnectRetry option is set to true
	return NewWithMqttClient(config, client)
}

type Incidents struct {
	config configuration.Config
	mqtt   paho.Client
}

func (this *Incidents) Handle(command messages.KafkaIncidentsCommand) error {
	deploymentName, err := this.stopProcessInCamunda(command)
	if err != nil {
		return err
	}
	return this.sendIncidentToSync(command, deploymentName)
}
