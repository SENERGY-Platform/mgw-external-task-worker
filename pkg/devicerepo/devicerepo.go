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

package devicerepo

import (
	"github.com/SENERGY-Platform/external-task-worker/lib/devicerepository"
	"github.com/SENERGY-Platform/external-task-worker/lib/devicerepository/model"
	"github.com/SENERGY-Platform/external-task-worker/util"
	"mgw-process-task-worker/pkg/configuration"
)

type Factory struct {
	Config configuration.Config
}

func (this Factory) Get(_ util.Config) devicerepository.RepoInterface {
	return this.new(this.Config)
}

func (this Factory) new(config configuration.Config) devicerepository.RepoInterface {
	return &Iot{config: config}
}

type Iot struct {
	config configuration.Config
}

func (this *Iot) GetDevice(token devicerepository.Impersonate, id string) (model.Device, error) {
	panic("implement me")
}

func (this *Iot) GetService(token devicerepository.Impersonate, device model.Device, serviceId string) (model.Service, error) {
	panic("implement me")
}

func (this *Iot) GetProtocol(token devicerepository.Impersonate, id string) (model.Protocol, error) {
	panic("implement me")
}

func (this *Iot) GetToken(user string) (devicerepository.Impersonate, error) {
	panic("implement me")
}

func (this *Iot) GetDeviceType(token devicerepository.Impersonate, id string) (model.DeviceType, error) {
	panic("implement me")
}

func (this *Iot) GetDeviceGroup(token devicerepository.Impersonate, id string) (model.DeviceGroup, error) {
	panic("implement me")
}
