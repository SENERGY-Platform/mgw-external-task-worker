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

package marshaller

import (
	"errors"
	"reflect"
	"testing"

	"github.com/SENERGY-Platform/external-task-worker/lib/marshaller"
	"github.com/SENERGY-Platform/marshaller/lib/config"
	marshaller_service_model "github.com/SENERGY-Platform/marshaller/lib/marshaller/model"
	marshaller_service_v2 "github.com/SENERGY-Platform/marshaller/lib/marshaller/v2"
)

type aspectRepoMock map[string]marshaller_service_model.AspectNode

func (this aspectRepoMock) GetAspectNode(id string) (marshaller_service_model.AspectNode, error) {
	node, ok := this[id]
	if !ok {
		return node, errors.New("unknown aspect node " + id)
	}
	return node, nil
}

const measureTemperature = marshaller_service_model.MEASURING_FUNCTION_PREFIX + "getTemperature"

var aspectTestProtocol = marshaller_service_model.Protocol{
	Id:               "p1",
	Handler:          "protocol1",
	ProtocolSegments: []marshaller_service_model.ProtocolSegment{{Id: "ms1", Name: "body"}},
}

// both temperatures are air temperatures, they differ only in their second aspect.
// "inside" is the first path, the one found without any aspect, so the aspect cases ask for
// "outside" to fail if the aspects are ignored.
var aspectTestService = marshaller_service_model.Service{
	Id:         "service_1",
	LocalId:    "s1u",
	ProtocolId: "p1",
	Outputs: []marshaller_service_model.Content{
		{
			Id: "metrics",
			ContentVariable: marshaller_service_model.ContentVariable{
				Id:   "metrics",
				Name: "metrics",
				Type: marshaller_service_model.Structure,
				SubContentVariables: []marshaller_service_model.ContentVariable{
					{
						Id:         "inside",
						Name:       "inside",
						Type:       marshaller_service_model.Integer,
						AspectIds:  []string{"air", "inside"},
						FunctionId: measureTemperature,
					},
					{
						Id:         "outside",
						Name:       "outside",
						Type:       marshaller_service_model.Integer,
						AspectIds:  []string{"air", "outside"},
						FunctionId: measureTemperature,
					},
				},
			},
			Serialization:     "json",
			ProtocolSegmentId: "ms1",
		},
	},
}

func newAspectTestMarshaller() *Marshaller {
	return &Marshaller{
		v2: marshaller_service_v2.New(config.Config{}, nil, nil),
		aspects: aspectRepoMock{
			"air":     {Id: "air"},
			"inside":  {Id: "inside"},
			"outside": {Id: "outside"},
		},
	}
}

func aspectTestRequest() marshaller.UnmarshallingV2Request {
	return marshaller.UnmarshallingV2Request{
		Service:    aspectTestService,
		Protocol:   aspectTestProtocol,
		Message:    map[string]string{"body": `{"inside":21,"outside":7}`},
		FunctionId: measureTemperature,
	}
}

func TestUnmarshalV2AspectLists(t *testing.T) {
	cases := []struct {
		name     string
		modify   func(request *marshaller.UnmarshallingV2Request)
		expected interface{}
	}{
		{
			name: "aspect nodes",
			modify: func(request *marshaller.UnmarshallingV2Request) {
				request.AspectNodes = []marshaller_service_model.AspectNode{{Id: "air"}, {Id: "outside"}}
			},
			expected: float64(7),
		},
		{
			name: "aspect node ids",
			modify: func(request *marshaller.UnmarshallingV2Request) {
				request.AspectNodeIds = []string{"air", "outside"}
			},
			expected: float64(7),
		},
		{
			name: "deprecated aspect node is a list with one element",
			modify: func(request *marshaller.UnmarshallingV2Request) {
				request.AspectNode = marshaller_service_model.AspectNode{Id: "outside"}
			},
			expected: float64(7),
		},
		{
			name: "deprecated aspect node id is a list with one element",
			modify: func(request *marshaller.UnmarshallingV2Request) {
				request.AspectNodeId = "outside"
			},
			expected: float64(7),
		},
		{
			name: "deprecated aspect node joins the aspect node ids",
			modify: func(request *marshaller.UnmarshallingV2Request) {
				request.AspectNode = marshaller_service_model.AspectNode{Id: "air"}
				request.AspectNodeIds = []string{"outside"}
			},
			expected: float64(7),
		},
		{
			name: "explicit path wins over aspects",
			modify: func(request *marshaller.UnmarshallingV2Request) {
				request.Path = "metrics.inside"
				request.AspectNodes = []marshaller_service_model.AspectNode{{Id: "outside"}}
			},
			expected: float64(21),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			request := aspectTestRequest()
			c.modify(&request)
			result, err := newAspectTestMarshaller().UnmarshalV2(request)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(result, c.expected) {
				t.Errorf("expected %#v, got %#v", c.expected, result)
			}
		})
	}
}

func TestUnmarshalV2AspectsAreAnAnd(t *testing.T) {
	request := aspectTestRequest()
	request.AspectNodes = []marshaller_service_model.AspectNode{{Id: "inside"}, {Id: "outside"}}
	_, err := newAspectTestMarshaller().UnmarshalV2(request)
	if err == nil {
		t.Fatal("expected an error: no variable is inside and outside at once")
	}
}

func TestUnmarshalV2UnknownAspectNodeId(t *testing.T) {
	request := aspectTestRequest()
	request.AspectNodeIds = []string{"air", "unknown"}
	_, err := newAspectTestMarshaller().UnmarshalV2(request)
	if err == nil {
		t.Fatal("expected an error for an aspect id the device-repository does not know")
	}
}

func TestUnmarshalV2PrefersClosestAspect(t *testing.T) {
	// both variables match "air", "metrics.inside" only through its child "inside_air"
	service := aspectTestService
	service.Outputs = []marshaller_service_model.Content{aspectTestService.Outputs[0]}
	service.Outputs[0].ContentVariable.SubContentVariables = []marshaller_service_model.ContentVariable{
		{Id: "inside", Name: "inside", Type: marshaller_service_model.Integer, AspectIds: []string{"inside_air"}, FunctionId: measureTemperature},
		{Id: "outside", Name: "outside", Type: marshaller_service_model.Integer, AspectIds: []string{"air"}, FunctionId: measureTemperature},
	}
	request := aspectTestRequest()
	request.Service = service
	request.AspectNodeIds = []string{"air"}

	m := newAspectTestMarshaller()
	m.aspects = aspectRepoMock{
		"air":        {Id: "air", ChildIds: []string{"inside_air"}, DescendentIds: []string{"inside_air"}},
		"inside_air": {Id: "inside_air", ParentId: "air", AncestorIds: []string{"air"}},
	}
	result, err := m.UnmarshalV2(request)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result, float64(7)) {
		t.Errorf("expected the value of the variable with the aspect itself, got %#v", result)
	}
}
