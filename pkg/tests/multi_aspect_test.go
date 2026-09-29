/*
 * Copyright 2026 InfAI (CC SES)
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
	"log"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/SENERGY-Platform/external-task-worker/lib/devicerepository/model"
	"github.com/SENERGY-Platform/external-task-worker/lib/messages"
	"github.com/SENERGY-Platform/external-task-worker/util"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/configuration"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/messaging"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/tests/docker"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/tests/mocks"
	paho "github.com/eclipse/paho.mqtt.golang"
)

// a device measures two air temperatures, which differ only in their second aspect,
// and can heat each of them with its own service
const multiAspectMeasureTemperature = model.MEASURING_FUNCTION_PREFIX + "getTemperature"
const multiAspectSetTemperature = model.CONTROLLING_FUNCTION_PREFIX + "setTemperature"

func multiAspectTemperatureVariable(name string, functionId string, aspectIds ...string) model.Content {
	return model.Content{
		Id: name,
		ContentVariable: model.ContentVariable{
			Id:   name,
			Name: "metrics",
			Type: model.Structure,
			SubContentVariables: []model.ContentVariable{
				{
					Id:               name + "_level",
					Name:             "level",
					Type:             model.Integer,
					CharacteristicId: example.Lux,
					AspectIds:        aspectIds,
					FunctionId:       functionId,
				},
			},
		},
		Serialization:     "json",
		ProtocolSegmentId: "ms1",
	}
}

var multiAspectDeviceType = model.DeviceType{
	Id:            "dt1",
	DeviceClassId: "dc-test-1",
	Services: []model.Service{
		{
			Id:          "service_measure",
			Name:        "measure",
			LocalId:     "measure",
			ProtocolId:  "p1",
			Interaction: model.REQUEST,
			Outputs: []model.Content{
				{
					Id: "metrics",
					ContentVariable: model.ContentVariable{
						Id:   "metrics",
						Name: "metrics",
						Type: model.Structure,
						SubContentVariables: []model.ContentVariable{
							{
								Id:               "inside",
								Name:             "inside",
								Type:             model.Integer,
								CharacteristicId: example.Lux,
								AspectIds:        []string{"air", "inside"},
								FunctionId:       multiAspectMeasureTemperature,
							},
							{
								Id:               "outside",
								Name:             "outside",
								Type:             model.Integer,
								CharacteristicId: example.Lux,
								AspectIds:        []string{"air", "outside"},
								FunctionId:       multiAspectMeasureTemperature,
							},
						},
					},
					Serialization:     "json",
					ProtocolSegmentId: "ms1",
				},
			},
		},
		{
			Id:          "service_heat_inside",
			Name:        "heat_inside",
			LocalId:     "heat_inside",
			ProtocolId:  "p1",
			Interaction: model.REQUEST,
			Inputs:      []model.Content{multiAspectTemperatureVariable("heat_inside", multiAspectSetTemperature, "air", "inside")},
		},
		{
			Id:          "service_heat_outside",
			Name:        "heat_outside",
			LocalId:     "heat_outside",
			ProtocolId:  "p1",
			Interaction: model.REQUEST,
			Inputs:      []model.Content{multiAspectTemperatureVariable("heat_outside", multiAspectSetTemperature, "air", "outside")},
		},
	},
}

type multiAspectEnv struct {
	camunda  *mocks.CamundaMock
	mux      sync.Mutex
	messages map[string][]string
}

func (this *multiAspectEnv) mgwMessages() map[string][]string {
	this.mux.Lock()
	defer this.mux.Unlock()
	return this.messages
}

// startMultiAspectWorker starts the worker against a device of multiAspectDeviceType, with the
// given commands as camunda tasks. Every command on the mgw mqtt is answered with data.
func startMultiAspectWorker(t *testing.T, ctx context.Context, wg *sync.WaitGroup, data string, commands ...messages.Command) *multiAspectEnv {
	t.Helper()
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
		t.Fatal(err)
	}
	config.CompletionStrategy = util.PESSIMISTIC

	repo, fallbackfile, err := mocks.NewFallbackFile(ctx, wg)
	if err != nil {
		t.Fatal(err)
	}
	config.FallbackFile = fallbackfile
	err = repo.RegisterDefaults()
	if err != nil {
		t.Fatal(err)
	}
	err = repo.RegisterDevice(model.Device{Id: "device_1", Name: "d1", DeviceTypeId: "dt1", LocalId: "d1u"})
	if err != nil {
		t.Fatal(err)
	}
	err = repo.RegisterDeviceGroup(model.DeviceGroup{Id: "dg_1", Name: "dg1", DeviceIds: []string{"device_1"}})
	if err != nil {
		t.Fatal(err)
	}
	err = repo.RegisterProtocol(model.Protocol{
		Id:               "p1",
		Name:             "protocol1",
		Handler:          "protocol1",
		ProtocolSegments: []model.ProtocolSegment{{Id: "ms1", Name: "body"}},
	})
	if err != nil {
		t.Fatal(err)
	}
	config.ProtocolSegment = "body"
	err = repo.RegisterDeviceType(multiAspectDeviceType)
	if err != nil {
		t.Fatal(err)
	}

	tasks := []messages.CamundaExternalTask{}
	for i, command := range commands {
		payload, err := json.Marshal(command)
		if err != nil {
			t.Fatal(err)
		}
		tasks = append(tasks, messages.CamundaExternalTask{
			Id: "test-task-id-" + strconv.Itoa(i+1),
			Variables: map[string]messages.CamundaVariable{
				util.CAMUNDA_VARIABLES_PAYLOAD: {Value: string(payload)},
			},
		})
	}
	env := &multiAspectEnv{messages: map[string][]string{}}
	env.camunda = mocks.NewCamundaMock(ctx, []interface{}{tasks})
	config.CamundaUrl = env.camunda.Server.URL

	mgwPort, _, err := docker.Mqtt(ctx, wg)
	if err != nil {
		t.Fatal(err)
	}
	config.MqttBroker = "tcp://localhost:" + mgwPort
	mgwMqtt := paho.NewClient(paho.NewClientOptions().
		SetAutoReconnect(true).
		SetCleanSession(true).
		AddBroker(config.MqttBroker))
	if token := mgwMqtt.Connect(); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}
	token := mgwMqtt.Subscribe("#", 2, func(client paho.Client, message paho.Message) {
		topic := message.Topic()
		payload := message.Payload()
		go func() {
			env.mux.Lock()
			defer env.mux.Unlock()
			env.messages[topic] = append(env.messages[topic], string(payload))
			if strings.HasPrefix(topic, "command") {
				msg := messaging.Command{}
				err := json.Unmarshal(payload, &msg)
				if err != nil {
					log.Fatal(err)
				}
				msg.Data = data
				resp, err := json.Marshal(msg)
				if err != nil {
					log.Fatal(err)
				}
				mgwMqtt.Publish(strings.Replace(topic, "command", "response", 1), 2, false, resp)
			}
		}()
	})
	if token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}

	syncPort, _, err := docker.Mqtt(ctx, wg)
	if err != nil {
		t.Fatal(err)
	}
	config.SyncMqttBroker = "tcp://localhost:" + syncPort
	config.SyncNetworkId = "test-network-id"

	go pkg.Start(ctx, config)
	time.Sleep(2 * time.Second)
	return env
}

func TestMultiAspectResponse(t *testing.T) {
	wg := sync.WaitGroup{}
	defer wg.Wait()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	measure := messages.Command{
		Version:          3,
		Function:         model.Function{Id: multiAspectMeasureTemperature},
		CharacteristicId: example.Lux,
		DeviceId:         "device_1",
		ServiceId:        "service_measure",
		ProtocolId:       "p1",
	}
	outside := measure
	outside.Aspects = []model.AspectNode{{Id: "air"}, {Id: "outside"}}
	deprecatedInside := measure
	deprecatedInside.Aspect = &model.AspectNode{Id: "inside"}

	env := startMultiAspectWorker(t, ctx, &wg, `{"inside":21,"outside":7}`, outside, deprecatedInside)

	result := func(value float64) []interface{} {
		return []interface{}{map[string]interface{}{
			"workerId": "worker-id",
			"localVariables": map[string]interface{}{
				"result": map[string]interface{}{"value": []interface{}{value}},
			},
		}}
	}
	if !reflect.DeepEqual(env.camunda.CompleteRequests, map[string][]interface{}{
		"/engine-rest/external-task/test-task-id-1/complete": result(7),
		"/engine-rest/external-task/test-task-id-2/complete": result(21),
	}) {
		t.Error(env.camunda.CompleteRequests)
	}
	if !reflect.DeepEqual(env.camunda.StopRequests, map[string][]interface{}{}) {
		t.Error(env.camunda.StopRequests)
	}
	if !reflect.DeepEqual(env.camunda.UnexpectedRequests, map[string][]interface{}{}) {
		t.Error(env.camunda.UnexpectedRequests)
	}
}

func TestMultiAspectDeviceGroup(t *testing.T) {
	wg := sync.WaitGroup{}
	defer wg.Wait()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	env := startMultiAspectWorker(t, ctx, &wg, "", messages.Command{
		Version:          3,
		Function:         model.Function{RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION, Id: multiAspectSetTemperature},
		CharacteristicId: example.Lux,
		DeviceGroupId:    "dg_1",
		Aspects:          []model.AspectNode{{Id: "air"}, {Id: "outside"}},
		DeviceClass:      &model.DeviceClass{Id: "dc-test-1"},
		Input:            float64(20),
	})

	commands := map[string][]string{}
	for topic, msgs := range env.mgwMessages() {
		if strings.HasPrefix(topic, "command") {
			commands[topic] = msgs
		}
	}
	if !reflect.DeepEqual(commands, map[string][]string{
		"command/d1u/heat_outside": {`{"command_id":"pessimistic-local-task-uuid-0","data":"{\"level\":20}"}`},
	}) {
		t.Error(commands)
	}
	if !reflect.DeepEqual(env.camunda.StopRequests, map[string][]interface{}{}) {
		t.Error(env.camunda.StopRequests)
	}
	if !reflect.DeepEqual(env.camunda.UnexpectedRequests, map[string][]interface{}{}) {
		t.Error(env.camunda.UnexpectedRequests)
	}
}
