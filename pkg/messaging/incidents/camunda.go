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
	"errors"
	"github.com/SENERGY-Platform/external-task-worker/lib/messages"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"runtime/debug"
	"time"
)

func (this *Incidents) stopProcessInCamunda(command messages.KafkaIncidentsCommand) (deploymentName string, err error) {
	incident := command.Incident
	err = this.stopProcessInstance(incident.ProcessInstanceId)
	if err != nil {
		return deploymentName, err
	}
	name, err := this.getProcessName(incident.ProcessDefinitionId)
	if err != nil {
		log.Println("WARNING: unable to get process name", err)
		return incident.ProcessDefinitionId, nil
	} else {
		return name, nil
	}
}

func (this *Incidents) stopProcessInstance(id string) (err error) {
	shard := this.config.CamundaUrl
	client := &http.Client{Timeout: 5 * time.Second}
	request, err := http.NewRequest("DELETE", shard+"/engine-rest/process-instance/"+url.PathEscape(id)+"?skipIoMappings=true", nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil
	}
	if resp.StatusCode == 200 || resp.StatusCode == 204 {
		return nil
	}
	msg, _ := ioutil.ReadAll(resp.Body)
	err = errors.New("error on delete in engine for " + shard + "/engine-rest/process-instance/" + url.PathEscape(id) + ": " + resp.Status + " " + string(msg))
	return err
}

func (this *Incidents) getProcessName(id string) (name string, err error) {
	shard := this.config.CamundaUrl
	client := &http.Client{Timeout: 5 * time.Second}
	request, err := http.NewRequest("GET", shard+"/engine-rest/process-definition/"+url.PathEscape(id), nil)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		temp, _ := ioutil.ReadAll(resp.Body)
		log.Println("ERROR:", resp.Status, string(temp))
		debug.PrintStack()
		return "", errors.New("unexpected response")
	}
	result := NameWrapper{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result.Name, err
}

type NameWrapper struct {
	Name string `json:"name"`
}
