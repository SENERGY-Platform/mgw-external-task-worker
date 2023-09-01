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

package marshaller

import (
	"context"
	"encoding/json"
	"errors"
	converterService "github.com/SENERGY-Platform/converter/lib/converter"
	"github.com/SENERGY-Platform/external-task-worker/lib/devicerepository/model"
	"github.com/SENERGY-Platform/external-task-worker/lib/marshaller"
	"github.com/SENERGY-Platform/marshaller/lib/config"
	marshaller_service_configurables "github.com/SENERGY-Platform/marshaller/lib/configurables"
	marshaller_service "github.com/SENERGY-Platform/marshaller/lib/marshaller"
	marshaller_service_model "github.com/SENERGY-Platform/marshaller/lib/marshaller/model"
	marshaller_service_v2 "github.com/SENERGY-Platform/marshaller/lib/marshaller/v2"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/configuration"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/devicerepo"
	"github.com/SENERGY-Platform/models/go/models"
	"log"
	"runtime/debug"
)

type Factory struct {
	Config     configuration.Config
	DeviceRepo *devicerepo.Iot
}

func (this Factory) New(ctx context.Context, _ string) marshaller.Interface {
	result, err := NewMarshaller(ctx, this.Config, this.DeviceRepo)
	if err != nil {
		panic(err)
	}
	return result
}

func NewMarshaller(ctx context.Context, conf configuration.Config, iot *devicerepo.Iot) (*Marshaller, error) {
	marshallerIot, err := NewMarshallerIot(ctx, conf, iot)
	if err != nil {
		return nil, err
	}

	conceptrepo, err := NewConceptRepo(ctx, conf, iot)
	if err != nil {
		return nil, err
	}

	converter, err := converterService.New()
	if err != nil {
		return nil, err
	}

	return &Marshaller{
		marshaller: marshaller_service.New(converter, conceptrepo, marshallerIot),
		v2:         marshaller_service_v2.New(config.Config{}, converter, conceptrepo),
	}, nil
}

type Marshaller struct {
	marshaller *marshaller_service.Marshaller
	v2         *marshaller_service_v2.Marshaller
}

func (this *Marshaller) MarshalFromServiceAndProtocol(characteristicId string, service model.Service, protocol model.Protocol, characteristicData interface{}, configurables []marshaller.Configurable) (result map[string]string, err error) {
	mockService := marshaller_service_model.Service{}
	mockProtocol := marshaller_service_model.Protocol{}
	mockConfigurables := []marshaller_service_configurables.Configurable{}
	err = jsonCast(service, &mockService)
	if err != nil {
		return result, err
	}
	err = jsonCast(protocol, &mockProtocol)
	if err != nil {
		return result, err
	}
	err = jsonCast(configurables, &mockConfigurables)
	if err != nil {
		return result, err
	}
	return this.marshaller.MarshalInputs(mockProtocol, mockService, characteristicData, characteristicId, nil, mockConfigurables...)
}

func (this *Marshaller) UnmarshalFromServiceAndProtocol(characteristicId string, service model.Service, protocol model.Protocol, message map[string]string, hints []string) (characteristicData interface{}, err error) {
	mockService := marshaller_service_model.Service{}
	mockProtocol := marshaller_service_model.Protocol{}
	err = jsonCast(service, &mockService)
	if err != nil {
		return characteristicData, err
	}
	err = jsonCast(protocol, &mockProtocol)
	if err != nil {
		return characteristicData, err
	}
	return this.marshaller.UnmarshalOutputs(mockProtocol, mockService, message, characteristicId, nil, hints...)
}

func jsonCast(in interface{}, out interface{}) (err error) {
	temp, err := json.Marshal(in)
	if err != nil {
		debug.PrintStack()
		return err
	}
	err = json.Unmarshal(temp, out)
	if err != nil {
		debug.PrintStack()
		return err
	}
	return nil
}

func (this *Marshaller) MarshalV2(service model.Service, protocol model.Protocol, data []marshaller.MarshallingV2RequestData) (result map[string]string, err error) {
	mockData := []marshaller_service_model.MarshallingV2RequestData{}
	err = jsonCast(data, &mockData)
	if err != nil {
		return result, err
	}
	return this.v2.Marshal(protocol, service, mockData)
}

func (this *Marshaller) UnmarshalV2(request marshaller.UnmarshallingV2Request) (result interface{}, err error) {
	var aspect *models.AspectNode
	if request.AspectNode.Id != "" {
		aspect = &request.AspectNode
	}
	if request.Path == "" {
		paths := this.v2.GetOutputPaths(request.Service, request.FunctionId, aspect)
		if len(paths) > 1 {
			log.Println("WARNING: only first path found by FunctionId and AspectNode is used for Unmarshal:", paths)
		}
		if len(paths) == 0 {
			return result, errors.New("no output path found for criteria")
		}
		request.Path = paths[0]
	}
	return this.v2.Unmarshal(request.Protocol, request.Service, request.CharacteristicId, request.Path, request.Message, nil)
}
