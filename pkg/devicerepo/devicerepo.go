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
	"encoding/json"
	"errors"
	"github.com/SENERGY-Platform/external-task-worker/lib/devicerepository"
	"github.com/SENERGY-Platform/external-task-worker/lib/devicerepository/model"
	"github.com/SENERGY-Platform/external-task-worker/util"
	"log"
	"mgw-process-task-worker/pkg/configuration"
	"net/http"
	"net/url"
	"time"
)

type Factory struct {
	Config configuration.Config
	Cache  *Cache
}

func (this Factory) Get(_ util.Config) devicerepository.RepoInterface {
	return New(this.Config, this.Cache)
}

func New(config configuration.Config, cache *Cache) devicerepository.RepoInterface {
	return &Iot{config: config, cache: cache}
}

type Iot struct {
	config      configuration.Config
	openIdToken *OpenidToken
	cache       *Cache
}

func (this *Iot) getToken() (string, error) {
	if this.openIdToken == nil {
		this.openIdToken = &OpenidToken{}
	}
	return this.openIdToken.EnsureAccess(this.config)
}

// implement GetToken()
// and have token devicerepository.Impersonate parameters
// to implement the devicerepository.RepoInterface (github.com/SENERGY-Platform/external-task-worker/lib/devicerepository)
// but use Iot.getToken() in the submethod like getDevice()
// this allows the usage of the fallback storage if the platform (repository + auth) is not reachable

func (this *Iot) GetToken(user string) (devicerepository.Impersonate, error) {
	return devicerepository.Impersonate(""), nil
}

func (this *Iot) GetDevice(_ devicerepository.Impersonate, id string) (result model.Device, err error) {
	err = this.cache.Use("device."+id, func() (interface{}, error) {
		return this.getDevice(id)
	}, &result)
	return
}

func (this *Iot) GetService(token devicerepository.Impersonate, device model.Device, serviceId string) (result model.Service, err error) {
	dt, err := this.GetDeviceType(token, device.DeviceTypeId)
	if err != nil {
		log.Println("ERROR: unable to load device-type", device.DeviceTypeId)
		return result, err
	}
	for _, service := range dt.Services {
		if service.Id == serviceId {
			return service, nil
		}
	}
	log.Println("ERROR: unable to find service in device-type", device.DeviceTypeId, serviceId)
	return result, errors.New("service not found")
}

func (this *Iot) GetProtocol(_ devicerepository.Impersonate, id string) (result model.Protocol, err error) {
	err = this.cache.Use("protocol."+id, func() (interface{}, error) {
		return this.getProtocol(id)
	}, &result)
	return
}

func (this *Iot) GetDeviceType(_ devicerepository.Impersonate, id string) (result model.DeviceType, err error) {
	err = this.cache.Use("deviceType."+id, func() (interface{}, error) {
		return this.getDeviceType(id)
	}, &result)
	return
}

func (this *Iot) GetDeviceGroup(_ devicerepository.Impersonate, id string) (result model.DeviceGroup, err error) {
	err = this.cache.Use("deviceGroup."+id, func() (interface{}, error) {
		return this.getDeviceGroup(id)
	}, &result)
	return
}

func (this *Iot) getDevice(id string) (result model.Device, err error) {
	token, err := this.getToken()
	if err != nil {
		return result, err
	}
	req, err := http.NewRequest("GET", this.config.DeviceRepoUrl+"/devices/"+url.PathEscape(id), nil)
	if err != nil {
		return result, err
	}
	req.Header.Set("Authorization", token)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return result, errors.New("device-type not found")
	}
	if resp.StatusCode >= 300 {
		return result, errors.New("unexpected status code")
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, err
	}
	return
}

func (this *Iot) getProtocol(id string) (result model.Protocol, err error) {
	token, err := this.getToken()
	if err != nil {
		return result, err
	}
	req, err := http.NewRequest("GET", this.config.DeviceRepoUrl+"/protocols/"+url.PathEscape(id), nil)
	if err != nil {
		return result, err
	}
	req.Header.Set("Authorization", token)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return result, errors.New("device-type not found")
	}
	if resp.StatusCode >= 300 {
		return result, errors.New("unexpected status code")
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, err
	}
	return
}

func (this *Iot) getDeviceType(id string) (result model.DeviceType, err error) {
	token, err := this.getToken()
	if err != nil {
		return result, err
	}
	req, err := http.NewRequest("GET", this.config.DeviceRepoUrl+"/device-types/"+url.PathEscape(id), nil)
	if err != nil {
		return result, err
	}
	req.Header.Set("Authorization", token)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return result, errors.New("device-type not found")
	}
	if resp.StatusCode >= 300 {
		return result, errors.New("unexpected status code")
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, err
	}
	return
}

func (this *Iot) getDeviceGroup(id string) (result model.DeviceGroup, err error) {
	token, err := this.getToken()
	if err != nil {
		return result, err
	}
	req, err := http.NewRequest("GET", this.config.DeviceRepoUrl+"/device-groups/"+url.PathEscape(id), nil)
	if err != nil {
		return result, err
	}
	req.Header.Set("Authorization", token)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return result, errors.New("device-type not found")
	}
	if resp.StatusCode >= 300 {
		return result, errors.New("unexpected status code")
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, err
	}
	return
}
