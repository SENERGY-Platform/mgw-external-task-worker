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

package devicerepo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/configuration"
)

func TestNotFoundErrorNamesTheMissingEntity(t *testing.T) {
	server := httptest.NewServer(http.NotFoundHandler())
	defer server.Close()
	iot := New(configuration.Config{DeviceRepoUrl: server.URL}, nil)

	cases := map[string]func() error{
		"device not found":       func() error { _, err := iot.getDevice("id"); return err },
		"device-type not found":  func() error { _, err := iot.getDeviceType("id"); return err },
		"protocol not found":     func() error { _, err := iot.getProtocol("id"); return err },
		"device-group not found": func() error { _, err := iot.getDeviceGroup("id"); return err },
	}
	for expected, get := range cases {
		err := get()
		if err == nil || err.Error() != expected {
			t.Errorf("expected error %q, got %v", expected, err)
		}
	}
}
