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
	"io"
	"log"
	"net/http"
	"net/url"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/SENERGY-Platform/device-repository/lib/client"
	"github.com/SENERGY-Platform/external-task-worker/lib/devicerepository"
	"github.com/SENERGY-Platform/external-task-worker/lib/devicerepository/model"
	"github.com/SENERGY-Platform/external-task-worker/util"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/configuration"
	"github.com/SENERGY-Platform/service-commons/pkg/cache"
)

const CacheExpiration = 60 * time.Second

type Provider struct {
	Config   configuration.Config
	Cache    *cache.Cache
	instance *Iot
	once     sync.Once
}

func (this *Provider) Get(_ util.Config) (devicerepository.RepoInterface, error) {
	return this.GetImpl(), nil
}

func (this *Provider) GetImpl() *Iot {
	this.once.Do(func() {
		this.instance = New(this.Config, this.Cache)
	})
	return this.instance
}

type Factory struct {
	Config configuration.Config
	Cache  *cache.Cache
}

func (this Factory) Get(_ util.Config) devicerepository.RepoInterface {
	return New(this.Config, this.Cache)
}

func New(config configuration.Config, cache *cache.Cache) *Iot {
	return &Iot{config: config, cache: cache}
}

type Iot struct {
	config      configuration.Config
	openIdToken *OpenidToken
	cache       *cache.Cache
}

func (this *Iot) getToken() (string, error) {
	if !this.config.AuthEnabled() {
		return "", nil
	}
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
	use := cache.Use[model.Device]
	if this.config.AsyncCacheRefresh {
		use = cache.UseWithAsyncRefresh[model.Device]
	}
	return use(this.cache, "device."+id, func() (model.Device, error) {
		return this.getDevice(id)
	}, func(device model.Device) error {
		if device.Id == "" {
			return errors.New("invalid device loaded from cache")
		}
		return nil
	}, CacheExpiration)
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
	use := cache.Use[model.Protocol]
	if this.config.AsyncCacheRefresh {
		use = cache.UseWithAsyncRefresh[model.Protocol]
	}
	return use(this.cache, "protocol."+id, func() (model.Protocol, error) {
		return this.getProtocol(id)
	}, func(protocol model.Protocol) error {
		if protocol.Id == "" {
			return errors.New("invalid protocol loaded from cache")
		}
		return nil
	}, CacheExpiration)
}

func (this *Iot) GetDeviceType(_ devicerepository.Impersonate, id string) (result model.DeviceType, err error) {
	use := cache.Use[model.DeviceType]
	if this.config.AsyncCacheRefresh {
		use = cache.UseWithAsyncRefresh[model.DeviceType]
	}
	return use(this.cache, "deviceType."+id, func() (model.DeviceType, error) {
		return this.getDeviceType(id)
	}, func(deviceType model.DeviceType) error {
		if deviceType.Id == "" {
			return errors.New("invalid device-type loaded from cache")
		}
		return nil
	}, CacheExpiration)
}

func (this *Iot) GetDeviceGroup(_ devicerepository.Impersonate, id string) (result model.DeviceGroup, err error) {
	use := cache.Use[model.DeviceGroup]
	if this.config.AsyncCacheRefresh {
		use = cache.UseWithAsyncRefresh[model.DeviceGroup]
	}
	return use(this.cache, "deviceGroup."+id, func() (model.DeviceGroup, error) {
		return this.getDeviceGroup(id)
	}, func(group model.DeviceGroup) error {
		if group.Id == "" {
			return errors.New("invalid device-group loaded from cache")
		}
		return nil
	}, CacheExpiration)
}

func (this *Iot) getDevice(id string) (result model.Device, err error) {
	req, err := http.NewRequest("GET", this.config.DeviceRepoUrl+"/devices/"+url.PathEscape(id), nil)
	if err != nil {
		return result, err
	}
	if this.config.AuthEnabled() {
		token, err := this.getToken()
		if err != nil {
			return result, err
		}
		req.Header.Set("Authorization", token)
	}
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := c.Do(req)
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
	req, err := http.NewRequest("GET", this.config.DeviceRepoUrl+"/protocols/"+url.PathEscape(id), nil)
	if err != nil {
		return result, err
	}
	if this.config.AuthEnabled() {
		token, err := this.getToken()
		if err != nil {
			return result, err
		}
		req.Header.Set("Authorization", token)
	}
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := c.Do(req)
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
	req, err := http.NewRequest("GET", this.config.DeviceRepoUrl+"/device-types/"+url.PathEscape(id), nil)
	if err != nil {
		return result, err
	}
	if this.config.AuthEnabled() {
		token, err := this.getToken()
		if err != nil {
			return result, err
		}
		req.Header.Set("Authorization", token)
	}
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := c.Do(req)
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
	req, err := http.NewRequest("GET", this.config.DeviceRepoUrl+"/device-groups/"+url.PathEscape(id), nil)
	if err != nil {
		return result, err
	}
	if this.config.AuthEnabled() {
		token, err := this.getToken()
		if err != nil {
			return result, err
		}
		req.Header.Set("Authorization", token)
	}
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := c.Do(req)
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

func (this *Iot) GetAspectNode(id string) (result model.AspectNode, err error) {
	use := cache.Use[model.AspectNode]
	if this.config.AsyncCacheRefresh {
		use = cache.UseWithAsyncRefresh[model.AspectNode]
	}
	return use(this.cache, "aspect-nodes."+id, func() (model.AspectNode, error) {
		return this.getAspectNode(id)
	}, func(aspectNode model.AspectNode) error {
		if aspectNode.Id == "" {
			return errors.New("invalid aspect-node loaded from cache")
		}
		return nil
	}, CacheExpiration)
}

func (this *Iot) getAspectNode(id string) (result model.AspectNode, err error) {
	token, err := this.getToken()
	if err != nil {
		return result, err
	}
	err = this.GetJson(token, this.config.DeviceRepoUrl+"/aspect-nodes/"+url.QueryEscape(id), &result)
	return
}

type IdWrapper struct {
	Id string `json:"id"`
}

func (this *Iot) GetConceptIds() (ids []string, err error) {
	use := cache.Use[[]string]
	if this.config.AsyncCacheRefresh {
		use = cache.UseWithAsyncRefresh[[]string]
	}
	return use(this.cache, "concept-ids", func() ([]string, error) {
		return this.getConceptIds()
	}, func(i []string) error {
		return nil
	}, CacheExpiration)
}

func (this *Iot) getConceptIds() (ids []string, err error) {
	limit := 1000
	offset := 0
	temp := []model.Concept{}
	c := client.NewClient(this.config.DeviceRepoUrl, this.getToken)
	for len(temp) == limit || offset == 0 {
		temp, _, err, _ = c.ListConcepts(client.ConceptListOptions{
			Limit:  int64(limit),
			Offset: int64(offset),
			SortBy: "name.asc",
		})
		if err != nil {
			log.Println("ERROR: getConceptIds()", err)
			return ids, err
		}
		for _, wrapper := range temp {
			ids = append(ids, wrapper.Id)
		}
		offset = offset + limit
	}
	return ids, err
}

func (this *Iot) ListFunctions() (functionInfos []model.Function, err error) {
	use := cache.Use[[]model.Function]
	if this.config.AsyncCacheRefresh {
		use = cache.UseWithAsyncRefresh[[]model.Function]
	}
	return use(this.cache, "list-functions", func() ([]model.Function, error) {
		return this.listFunctions()
	}, func(functions []model.Function) error {
		return nil
	}, CacheExpiration)
}

func (this *Iot) listFunctions() (functionInfos []model.Function, err error) {
	limit := 100
	offset := 0
	temp := []model.Function{}
	c := client.NewClient(this.config.DeviceRepoUrl, this.getToken)
	for len(temp) == limit || offset == 0 {
		temp, _, err, _ = c.ListFunctions(client.FunctionListOptions{
			Limit:  int64(limit),
			Offset: int64(offset),
			SortBy: "name.asc",
		})
		if err != nil {
			return functionInfos, err
		}
		for _, f := range temp {
			functionInfos = append(functionInfos, f)
		}
		offset = offset + limit
	}
	return functionInfos, err
}

func (this *Iot) GetCharacteristic(id string) (result model.Characteristic, err error) {
	use := cache.Use[model.Characteristic]
	if this.config.AsyncCacheRefresh {
		use = cache.UseWithAsyncRefresh[model.Characteristic]
	}
	return use(this.cache, "characteristics."+id, func() (model.Characteristic, error) {
		return this.getCharacteristic(id)
	}, func(characteristic model.Characteristic) error {
		if characteristic.Id == "" {
			return errors.New("invalid characteristic loaded from cache")
		}
		return nil
	}, CacheExpiration)
}

func (this *Iot) getCharacteristic(id string) (result model.Characteristic, err error) {
	token, err := this.getToken()
	if err != nil {
		return result, err
	}
	err = this.GetJson(token, this.config.DeviceRepoUrl+"/characteristics/"+url.PathEscape(id), &result)
	return
}

func (this *Iot) GetConcept(id string) (result model.Concept, err error) {
	use := cache.Use[model.Concept]
	if this.config.AsyncCacheRefresh {
		use = cache.UseWithAsyncRefresh[model.Concept]
	}
	return use(this.cache, "concept."+id, func() (model.Concept, error) {
		return this.getConcept(id)
	}, func(concept model.Concept) error {
		if concept.Id == "" {
			return errors.New("invalid concept loaded from cache")
		}
		return nil
	}, CacheExpiration)
}

func (this *Iot) getConcept(id string) (result model.Concept, err error) {
	token, err := this.getToken()
	if err != nil {
		return result, err
	}
	err = this.GetJson(token, this.config.DeviceRepoUrl+"/concepts/"+url.PathEscape(id), &result)
	return
}

func (this *Iot) GetJson(token string, endpoint string, result interface{}) (err error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}
	if token != "" {
		req.Header.Set("Authorization", token)
	}
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		temp, _ := io.ReadAll(resp.Body)
		return errors.New(strings.TrimSpace(string(temp)))
	}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		log.Println("ERROR:", err.Error())
		debug.PrintStack()
	}
	return
}
