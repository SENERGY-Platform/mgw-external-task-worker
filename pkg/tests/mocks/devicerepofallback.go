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
	"github.com/SENERGY-Platform/external-task-worker/lib/devicerepository/model"
	"mgw-external-task-worker/pkg/devicerepo"
	"os"
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
	fallback, err := devicerepo.NewFallback(filename)
	if err != nil {
		return nil, "", err
	}
	repo = &Repo{fallback: fallback}
	return
}

type Repo struct {
	fallback devicerepo.Fallback
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
