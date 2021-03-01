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
	"github.com/SENERGY-Platform/converter/lib/converter/example"
	"github.com/SENERGY-Platform/external-task-worker/lib/devicerepository/model"
	"github.com/SENERGY-Platform/external-task-worker/lib/messages"
	"github.com/SENERGY-Platform/external-task-worker/util"
	paho "github.com/eclipse/paho.mqtt.golang"
	"log"
	"mgw-external-task-worker/pkg"
	"mgw-external-task-worker/pkg/configuration"
	"mgw-external-task-worker/pkg/messaging"
	"mgw-external-task-worker/pkg/tests/docker"
	"mgw-external-task-worker/pkg/tests/mocks"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	t.Log("error messages like \n\tERROR: getOpenidToken::PostForm() Post \"/auth/realms/master/protocol/openid-connect/token\": unsupported protocol scheme \"\" \nare expected in this test")
	wg := sync.WaitGroup{}
	defer wg.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	util.GetId = func() string {
		return "worker-id"
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
		Id: "dt1",
		Services: []model.Service{
			{
				Id:         "service_1",
				Name:       "s1",
				LocalId:    "s1u",
				ProtocolId: "p1",
				Outputs: []model.Content{
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
									CharacteristicId: example.Hex,
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
		Version:          2,
		Function:         model.Function{RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION},
		CharacteristicId: example.Rgb,
		DeviceId:         "device_1",
		ServiceId:        "service_1",
		ProtocolId:       "p1",
	}

	cmdMsg1, err := json.Marshal(cmd1)
	if err != nil {
		t.Error(err)
		return
	}

	cmd2 := messages.Command{
		Version:          2,
		Function:         model.Function{RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION},
		CharacteristicId: example.Hex,
		DeviceId:         "device_1",
		ServiceId:        "service_1",
		ProtocolId:       "p1",
	}

	cmdMsg2, err := json.Marshal(cmd2)
	if err != nil {
		t.Error(err)
		return
	}

	fetchRespnses := []interface{}{
		[]messages.CamundaExternalTask{
			{
				Id: "test-task-id-1",
				Variables: map[string]messages.CamundaVariable{
					util.CAMUNDA_VARIABLES_PAYLOAD: {
						Value: string(cmdMsg1),
					},
				},
			},
			{
				Id: "test-task-id-2",
				Variables: map[string]messages.CamundaVariable{
					util.CAMUNDA_VARIABLES_PAYLOAD: {
						Value: string(cmdMsg2),
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
		mgwmux.Lock()
		defer mgwmux.Unlock()
		mgwMessages[message.Topic()] = append(mgwMessages[message.Topic()], string(message.Payload()))
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

	time.Sleep(1 * time.Second)

	mgwmux.Lock()
	for topic, messages := range mgwMessages {
		if strings.HasPrefix(topic, "command") {
			for _, msgStr := range messages {
				msg := messaging.Command{}
				err = json.Unmarshal([]byte(msgStr), &msg)
				if err != nil {
					log.Fatal(err)
				}
				msg.Data = "{\"level\":\"#c83200\"}"
				resp, err := json.Marshal(msg)
				if err != nil {
					log.Fatal(err)
				}
				respTopic := strings.Replace(topic, "command", "response", 1)
				mgwMqtt.Publish(respTopic, 2, false, resp)
			}
		}
	}
	mgwmux.Unlock()

	time.Sleep(1 * time.Second)

	t.Log(mgwMessages)
	t.Log(syncMessages)
	t.Log(camundamock.UnexpectedRequests)
	t.Log(camundamock.CompleteRequests)
	t.Log(camundamock.StopRequests)

	if !reflect.DeepEqual(mgwMessages, map[string][]string{
		"command/d1u/s1u": {
			`{"command_id":"pessimistic-local-task-uuid-0","data":""}`,
			`{"command_id":"pessimistic-local-task-uuid-1","data":""}`,
		},
		"response/d1u/s1u": {
			`{"command_id":"pessimistic-local-task-uuid-0","data":"{\"level\":\"#c83200\"}"}`,
			`{"command_id":"pessimistic-local-task-uuid-1","data":"{\"level\":\"#c83200\"}"}`,
		},
	}) {
		t.Error(mgwMessages)
	}

	if !reflect.DeepEqual(syncMessages, map[string][]string{}) {
		t.Error(syncMessages)
	}

	if !reflect.DeepEqual(camundamock.CompleteRequests, map[string][]interface{}{
		"/engine-rest/external-task/test-task-id-1/complete": {
			map[string]interface{}{
				"workerId": "worker-id",
				"localVariables": map[string]interface{}{
					"result": map[string]interface{}{
						"value": []interface{}{
							map[string]interface{}{
								"r": float64(200),
								"g": float64(50),
								"b": float64(0),
							},
						},
					},
				},
			},
		},
		"/engine-rest/external-task/test-task-id-2/complete": {
			map[string]interface{}{
				"workerId": "worker-id",
				"localVariables": map[string]interface{}{
					"result": map[string]interface{}{
						"value": []interface{}{"#c83200"},
					},
				},
			},
		},
	}) {
		t.Error(camundamock.CompleteRequests)
	}

	if !reflect.DeepEqual(camundamock.StopRequests, map[string][]interface{}{}) {
		t.Error(camundamock.StopRequests)
	}

	if !reflect.DeepEqual(camundamock.UnexpectedRequests, map[string][]interface{}{}) {
		t.Error(camundamock.UnexpectedRequests)
	}

}
