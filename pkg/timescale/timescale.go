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

package timescale

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/SENERGY-Platform/external-task-worker/lib/messages"
	"github.com/SENERGY-Platform/external-task-worker/lib/timescale"
	"github.com/SENERGY-Platform/external-task-worker/util"
	"io"
	"net/http"
	"strings"
	"time"
)

type Timescale struct {
	TimescaleWrapperUrl string
}

func Factory(ctx context.Context, config util.Config) (timescale.Interface, error) {
	return NewTimescale(config.TimescaleWrapperUrl), nil
}

func NewTimescale(timescaleWrapperUrl string) *Timescale {
	return &Timescale{TimescaleWrapperUrl: timescaleWrapperUrl}
}

func (this *Timescale) Query(token string, request []messages.TimescaleRequest, timeout time.Duration) (result []messages.TimescaleResponse, err error) {
	body := &bytes.Buffer{}
	err = json.NewEncoder(body).Encode(this.castRequest(request))
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", this.TimescaleWrapperUrl+"/last-values", body)
	if err != nil {
		return result, err
	}
	req.Header.Set("Authorization", token)
	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		temp, _ := io.ReadAll(resp.Body)
		return result, errors.New(strings.TrimSpace(string(temp)))
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	return
}

func (this *Timescale) castRequest(request []messages.TimescaleRequest) (result []TimescaleRequest) {
	for _, r := range request {
		result = append(result, TimescaleRequest{
			DeviceId:   r.Device.LocalId,
			ServiceId:  r.Service.LocalId,
			ColumnName: this.castPath(r.ColumnName),
		})
	}
	return
}

func (this *Timescale) castPath(path string) string {
	if path == "" {
		return path
	}
	parts := strings.Split(path, ".")
	if len(parts) == 0 {
		return path
	}
	return strings.Join(parts[1:], ".")
}

type TimescaleRequest struct {
	DeviceId   string
	ServiceId  string
	ColumnName string
}
