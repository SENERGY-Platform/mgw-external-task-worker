/*
 * Copyright 2022 InfAI (CC SES)
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
	"errors"
	"github.com/SENERGY-Platform/marshaller/lib/marshaller/model"
	"mgw-external-task-worker/pkg/configuration"
	"mgw-external-task-worker/pkg/devicerepo"
	"net/http"
)

func NewMarshallerIot(ctx context.Context, conf configuration.Config, iot *devicerepo.Iot) (result *MarshallerIot, err error) {
	return &MarshallerIot{iot: iot}, nil
}

type MarshallerIot struct {
	iot *devicerepo.Iot
}

func (this *MarshallerIot) GetAspectNode(id string) (result model.AspectNode, err error) {
	temp, err := this.iot.GetAspectNode(id)
	if err != nil {
		return result, err
	}
	err = jsonCast(temp, &result)
	return result, err
}

//this method should only be needed for old marshal/unmarshal requests
func (this MarshallerIot) GetDeviceType(id string) (result model.DeviceType, err error, code int) {
	return result, errors.New("not implemented"), http.StatusInternalServerError
}
