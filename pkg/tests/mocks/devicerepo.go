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

package mocks

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"runtime/debug"
	"strconv"
	"sync"

	"github.com/SENERGY-Platform/device-repository/v2/lib/client"
	"github.com/SENERGY-Platform/external-task-worker/lib/devicerepository/model"
)

// NewDeviceRepo starts a device-repository mock that serves the registered entities
// on the endpoints pkg/devicerepo reads; it stops when ctx is done
func NewDeviceRepo(ctx context.Context, wg *sync.WaitGroup) (repo *Repo) {
	repo = &Repo{entities: map[string]interface{}{}}
	repo.server = httptest.NewServer(http.HandlerFunc(repo.serve))
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		repo.server.Close()
	}()
	return repo
}

type Repo struct {
	server     *httptest.Server
	mux        sync.Mutex
	entities   map[string]interface{} //by request path
	conceptIds []string
	functions  []model.Function
}

func (this *Repo) Url() string {
	return this.server.URL
}

func (this *Repo) serve(w http.ResponseWriter, r *http.Request) {
	this.mux.Lock()
	defer this.mux.Unlock()
	var result interface{}
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/v2/concepts":
		concepts := []model.Concept{}
		for _, id := range this.conceptIds {
			concepts = append(concepts, model.Concept{Id: id})
		}
		limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
		offset, _ := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
		w.Header().Set("X-Total-Count", strconv.Itoa(len(concepts)))
		result = page(concepts, limit, offset)
	case r.Method == http.MethodPost && r.URL.Path == "/query/functions":
		options := client.FunctionListOptions{}
		err := json.NewDecoder(r.Body).Decode(&options)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("X-Total-Count", strconv.Itoa(len(this.functions)))
		result = page(this.functions, options.Limit, options.Offset)
	case r.Method == http.MethodGet:
		value, ok := this.entities[r.URL.Path]
		if !ok {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		result = value
	default:
		http.Error(w, "unexpected request", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(result)
}

// page returns the elements of a list request; a limit of 0 means no limit, like the device-repository client omits it
func page[T any](list []T, limit int64, offset int64) []T {
	if offset >= int64(len(list)) {
		return []T{}
	}
	list = list[offset:]
	if limit > 0 && limit < int64(len(list)) {
		list = list[:limit]
	}
	return list
}

func (this *Repo) set(path string, value interface{}) error {
	this.mux.Lock()
	defer this.mux.Unlock()
	this.entities[path] = value
	return nil
}

func (this *Repo) RegisterDevice(device model.Device) (err error) {
	return this.set("/devices/"+device.Id, device)
}

func (this *Repo) RegisterProtocol(protocol model.Protocol) (err error) {
	return this.set("/protocols/"+protocol.Id, protocol)
}

func (this *Repo) RegisterDeviceType(deviceType model.DeviceType) (err error) {
	return this.set("/device-types/"+deviceType.Id, deviceType)
}

func (this *Repo) RegisterDeviceGroup(group model.DeviceGroup) error {
	return this.set("/device-groups/"+group.Id, group)
}

func (this *Repo) RegisterAspectNode(aspect model.AspectNode) (err error) {
	return this.set("/aspect-nodes/"+aspect.Id, aspect)
}

func (this *Repo) RegisterConceptIds(ids []string) (err error) {
	this.mux.Lock()
	defer this.mux.Unlock()
	this.conceptIds = ids
	return nil
}

func (this *Repo) RegisterListFunctions(functionInfos []model.Function) (err error) {
	this.mux.Lock()
	defer this.mux.Unlock()
	this.functions = functionInfos
	return nil
}

func (this *Repo) RegisterCharacteristic(characteristic model.Characteristic) (err error) {
	return this.set("/characteristics/"+characteristic.Id, characteristic)
}

func (this *Repo) RegisterConcept(concept model.Concept) (err error) {
	return this.set("/concepts/"+concept.Id, concept)
}

func (this *Repo) RegisterDefaults() (err error) {
	conceptList := []model.Concept{}
	err = json.Unmarshal([]byte(conceptListStr), &conceptList)
	if err != nil {
		debug.PrintStack()
		return err
	}
	conceptIds := []string{}
	for _, e := range conceptList {
		err = this.RegisterConcept(e)
		if err != nil {
			debug.PrintStack()
			return err
		}
		conceptIds = append(conceptIds, e.Id)
	}
	err = this.RegisterConceptIds(conceptIds)
	if err != nil {
		debug.PrintStack()
		return err
	}

	functions := []model.Function{}
	err = json.Unmarshal([]byte(functionsListStr), &functions)
	if err != nil {
		debug.PrintStack()
		return err
	}
	err = this.RegisterListFunctions(functions)
	if err != nil {
		debug.PrintStack()
		return err
	}
	conceptMap := map[string]model.Concept{}
	err = json.Unmarshal([]byte(conceptPathMapStr), &conceptMap)
	if err != nil {
		debug.PrintStack()
		return err
	}
	for _, e := range conceptMap {
		for _, id := range e.CharacteristicIds {
			err = this.RegisterCharacteristic(model.Characteristic{Id: id})
			if err != nil {
				debug.PrintStack()
				return err
			}
		}
	}

	characteristicsMap := map[string]model.Characteristic{}
	err = json.Unmarshal([]byte(characteristicsPathMapStr), &characteristicsMap)
	if err != nil {
		debug.PrintStack()
		return err
	}
	for _, e := range characteristicsMap {
		err = this.RegisterCharacteristic(e)
		if err != nil {
			debug.PrintStack()
			return err
		}
	}

	return nil
}
