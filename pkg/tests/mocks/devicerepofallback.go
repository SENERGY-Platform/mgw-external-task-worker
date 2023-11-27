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
	"github.com/SENERGY-Platform/external-task-worker/lib/devicerepository/model"
	"github.com/SENERGY-Platform/service-commons/pkg/cache/fallback"
	"os"
	"runtime/debug"
	"strconv"
	"sync"
	"time"
)

func NewFallbackFile(ctx context.Context, wg *sync.WaitGroup) (repo *Repo, filename string, err error) {
	filename = "/tmp/test" + strconv.FormatInt(time.Now().Unix(), 10) + ".json"
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		os.Remove(filename)
	}()
	fb, err := fallback.New(filename)
	if err != nil {
		return nil, "", err
	}
	repo = &Repo{fallback: fb}
	return
}

type Repo struct {
	fallback *fallback.Fallback
}

func (this *Repo) RegisterDevice(device model.Device) (err error) {
	return this.fallback.Set("device."+device.Id, device)
}

func (this *Repo) RegisterProtocol(protocol model.Protocol) (err error) {
	return this.fallback.Set("protocol."+protocol.Id, protocol)
}

func (this *Repo) RegisterDeviceType(deviceType model.DeviceType) (err error) {
	return this.fallback.Set("deviceType."+deviceType.Id, deviceType)
}

func (this *Repo) RegisterDeviceGroup(group model.DeviceGroup) error {
	return this.fallback.Set("deviceGroup."+group.Id, group)
}

func (this *Repo) RegisterAspectNode(aspect model.AspectNode) (err error) {
	return this.fallback.Set("aspect-nodes."+aspect.Id, aspect)
}

func (this *Repo) RegisterConceptIds(ids []string) (err error) {
	return this.fallback.Set("concept-ids", ids)
}

func (this *Repo) RegisterListFunctions(functionInfos []model.Function) (err error) {
	return this.fallback.Set("list-functions", functionInfos)
}

func (this *Repo) RegisterCharacteristic(characteristic model.Characteristic) (err error) {
	return this.fallback.Set("characteristics."+characteristic.Id, characteristic)
}

func (this *Repo) RegisterConcept(concept model.Concept) (err error) {
	return this.fallback.Set("concept."+concept.Id, concept)
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
