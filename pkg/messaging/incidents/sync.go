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
	"encoding/json"
	"github.com/SENERGY-Platform/external-task-worker/lib/messages"
)

func (this *Incidents) sendIncidentToSync(command messages.KafkaIncidentsCommand, deploymentName string) error {
	incident := command.Incident
	incident.DeploymentName = deploymentName
	return this.send(this.getStateTopic(incidentTopic), incident)
}

const incidentTopic = "incident"

func (this *Incidents) getStateTopic(entity string, substate ...string) (topic string) {
	topic = this.getBaseTopic() + "/" + entity
	for _, sub := range substate {
		topic = topic + "/" + sub
	}
	return
}

func (this *Incidents) getBaseTopic() string {
	if this.config.SyncNetworkId != "" {
		return "processes/" + this.config.SyncNetworkId
	} else {
		return "processes"
	}
}

func (this *Incidents) send(topic string, message interface{}) error {
	msg, err := json.Marshal(message)
	if err != nil {
		return err
	}
	token := this.mqtt.Publish(topic, 2, false, msg)
	token.Wait()
	return token.Error()
}
