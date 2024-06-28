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

package tests

import (
	"context"
	"encoding/json"
	"github.com/SENERGY-Platform/external-task-worker/lib/devicerepository/model"
	"github.com/SENERGY-Platform/external-task-worker/lib/messages"
	"github.com/SENERGY-Platform/external-task-worker/util"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/configuration"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/messaging"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/tests/docker"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/tests/mocks"
	paho "github.com/eclipse/paho.mqtt.golang"
	"log"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestIncident(t *testing.T) {
	t.Log("error messages like \n\tERROR: getOpenidToken::PostForm() Post \"/auth/realms/master/protocol/openid-connect/token\": unsupported protocol scheme \"\" \nare expected in this test")
	wg := sync.WaitGroup{}
	defer wg.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	util.GetId = func() string {
		return "worker-id"
	}

	util.TimeNow = func() time.Time {
		return time.Time{}
	}

	count := 0
	configuration.DefaultIdProvider = func() string {
		defer func() { count = count + 1 }()
		return "uuid-" + strconv.Itoa(count)
	}

	config, err := configuration.Load("../../config.json")
	if err != nil {
		t.Error(err)
		return
	}

	config.CompletionStrategy = util.PESSIMISTIC
	config.CamundaWorkerTimeout = 100

	repo, fallbackfile, err := mocks.NewFallbackFile(ctx, &wg)
	if err != nil {
		t.Error(err)
		return
	}
	config.FallbackFile = fallbackfile

	err = repo.RegisterDefaults()
	if err != nil {
		t.Error(err)
		return
	}

	err = repo.RegisterDevice(model.Device{
		Id:           "device_1",
		Name:         "d1",
		DeviceTypeId: "dt1",
		LocalId:      "d1u",
	})
	if err != nil {
		t.Error(err)
		return
	}

	err = repo.RegisterDefaults()
	if err != nil {
		t.Error(err)
		return
	}

	err = repo.RegisterDevice(model.Device{
		Id:           "device_2",
		Name:         "d2",
		DeviceTypeId: "dt1",
		LocalId:      "d1u",
	})
	if err != nil {
		t.Error(err)
		return
	}

	err = repo.RegisterDeviceGroup(model.DeviceGroup{
		Id:        "dg_1",
		Name:      "dg1",
		DeviceIds: []string{"device_1", "device_2"},
	})
	if err != nil {
		t.Error(err)
		return
	}

	err = repo.RegisterProtocol(model.Protocol{
		Id:               "p1",
		Name:             "protocol1",
		Handler:          "protocol1",
		ProtocolSegments: []model.ProtocolSegment{{Id: "ms1", Name: "body"}},
	})
	if err != nil {
		t.Error(err)
		return
	}
	config.ProtocolHandler = "protocol1"
	config.ProtocolSegment = "body"

	err = repo.RegisterDeviceType(model.DeviceType{
		Id:            "dt1",
		DeviceClassId: "dc-test-1",
		Services: []model.Service{
			{
				Id:          "service_1",
				Name:        "s1",
				LocalId:     "s1u",
				ProtocolId:  "p1",
				Interaction: model.REQUEST,
				Inputs: []model.Content{
					{
						Id: "metrics",
						ContentVariable: model.ContentVariable{
							Id:   "metrics",
							Name: "metrics",
							Type: model.Structure,
							SubContentVariables: []model.ContentVariable{
								{
									Id:               "level",
									Name:             "level",
									Type:             model.Integer,
									CharacteristicId: example.Rgb,
									AspectId:         "color",
									FunctionId:       model.CONTROLLING_FUNCTION_PREFIX + ":test-function-2",
								},
							},
						},
						Serialization:     "json",
						ProtocolSegmentId: "ms1",
					},
				},
			},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	cmd1 := messages.Command{
		Version:          3,
		Function:         model.Function{RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION, Id: model.CONTROLLING_FUNCTION_PREFIX + ":test-function-2"},
		CharacteristicId: example.Hex,
		DeviceGroupId:    "dg_1",
		Aspect: &model.AspectNode{
			Id: "color",
		},
		DeviceClass: &model.DeviceClass{
			Id: "dc-test-1",
		},
		Input: "foobar",
	}

	cmdMsg1, err := json.Marshal(cmd1)
	if err != nil {
		t.Error(err)
		return
	}

	fetchRespnses := []interface{}{
		[]messages.CamundaExternalTask{
			{
				Id:                  "test-task-id-1",
				ProcessDefinitionId: "test-definition-id",
				ProcessInstanceId:   "test-instance-id",
				Variables: map[string]messages.CamundaVariable{
					util.CAMUNDA_VARIABLES_PAYLOAD: {
						Value: string(cmdMsg1),
					},
				},
			},
		},
	}

	camundamock := mocks.NewCamundaMock(ctx, fetchRespnses)

	config.CamundaUrl = camundamock.Server.URL

	// mgw mqtt
	mgwPort, _, err := docker.Mqtt(ctx, &wg)
	if err != nil {
		t.Error(err)
		return
	}
	config.MqttBroker = "tcp://localhost:" + mgwPort

	mgwMqtt := paho.NewClient(paho.NewClientOptions().
		SetAutoReconnect(true).
		SetCleanSession(true).
		AddBroker(config.MqttBroker))
	if token := mgwMqtt.Connect(); token.Wait() && token.Error() != nil {
		t.Error(token.Error())
		return
	}

	mgwMessages := map[string][]string{}
	mgwmux := sync.Mutex{}
	token := mgwMqtt.Subscribe("#", 2, func(client paho.Client, message paho.Message) {
		topic := message.Topic()
		payload := message.Payload()
		go func() {
			mgwmux.Lock()
			defer mgwmux.Unlock()
			mgwMessages[topic] = append(mgwMessages[topic], string(payload))
			if strings.HasPrefix(topic, "command") {
				msg := messaging.Command{}
				err = json.Unmarshal([]byte(payload), &msg)
				if err != nil {
					log.Fatal(err)
				}
				msg.Data = ""
				resp, err := json.Marshal(msg)
				if err != nil {
					log.Fatal(err)
				}
				respTopic := strings.Replace(topic, "command", "response", 1)
				mgwMqtt.Publish(respTopic, 2, false, resp)
			}
		}()
	})
	if token.Wait() && token.Error() != nil {
		t.Error(token.Error())
		return
	}

	//sync mqtt
	syncPort, _, err := docker.Mqtt(ctx, &wg)
	if err != nil {
		t.Error(err)
		return
	}
	config.SyncMqttBroker = "tcp://localhost:" + syncPort

	syncMqtt := paho.NewClient(paho.NewClientOptions().
		SetAutoReconnect(true).
		SetCleanSession(true).
		AddBroker(config.SyncMqttBroker))
	if token := syncMqtt.Connect(); token.Wait() && token.Error() != nil {
		t.Error(token.Error())
		return
	}

	syncMessages := map[string][]string{}
	syncmux := sync.Mutex{}
	token = syncMqtt.Subscribe("#", 2, func(client paho.Client, message paho.Message) {
		syncmux.Lock()
		defer syncmux.Unlock()
		syncMessages[message.Topic()] = append(syncMessages[message.Topic()], string(message.Payload()))
	})
	if token.Wait() && token.Error() != nil {
		t.Error(token.Error())
		return
	}

	config.SyncNetworkId = "test-network-id"

	go pkg.Start(ctx, config)

	time.Sleep(2 * time.Second)

	/*
		t.Log(mgwMessages)
		t.Log(syncMessages)
		t.Log(camundamock.UnexpectedRequests)
		t.Log(camundamock.CompleteRequests)
		t.Log(camundamock.StopRequests)
	*/
	if !reflect.DeepEqual(mgwMessages, map[string][]string{}) {
		t.Error(mgwMessages)
		return
	}

	expectedSyncMessages := map[string][]string{
		"processes/test-network-id/state/incident": {
			`{"id":"worker-id","external_task_id":"test-task-id-1","process_instance_id":"test-instance-id","process_definition_id":"test-definition-id","worker_id":"worker-id","error_message":"internal format error: Parsing of color failed, Bad Color","time":"0001-01-01T00:00:00Z","tenant_id":"","deployment_name":"test-definition-id"}`,
		},
	}
	if !reflect.DeepEqual(syncMessages, expectedSyncMessages) {
		t.Errorf("\n%#v\n%#v\n", expectedSyncMessages, syncMessages)
		return
	}

	if !reflect.DeepEqual(camundamock.CompleteRequests, map[string][]interface{}{}) {
		t.Error(camundamock.CompleteRequests)
		return
	}

	if !reflect.DeepEqual(camundamock.StopRequests, map[string][]interface{}{
		"/engine-rest/process-instance/test-instance-id": {nil},
	}) {
		t.Error(camundamock.StopRequests)
		return
	}

	if !reflect.DeepEqual(camundamock.UnexpectedRequests, map[string][]interface{}{
		"/engine-rest/process-definition/test-definition-id": {nil},
	}) {
		t.Error(camundamock.UnexpectedRequests)
		return
	}

}
