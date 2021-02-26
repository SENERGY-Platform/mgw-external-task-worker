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

package tests

import (
	"github.com/SENERGY-Platform/external-task-worker/util"
	"log"
	"mgw-external-task-worker/pkg/configuration"
	"mgw-external-task-worker/pkg/devicerepo"
	"testing"
)

func TestDeviceRepo(t *testing.T) {
	config, err := configuration.Load("resources/config.json")
	if err != nil {
		t.Error(err)
		return
	}

	fallback, err := devicerepo.NewFallback(config.FallbackFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	cache := devicerepo.NewCache(fallback)
	drf := devicerepo.Factory{Config: config, Cache: cache}
	repo := drf.Get(util.Config{})

	token, err := repo.GetToken(config.AuthUserName) //is not needed but may be useful if implementation changes
	if err != nil {
		log.Fatal(err)
		return
	}

	dt, err := repo.GetDeviceType(token, "urn:infai:ses:device-type:1be88694-2981-416a-8a58-ad7fb2c63a37")
	if err != nil {
		t.Error(err)
		return
	}

	if dt.Id == "" {
		t.Error(dt)
		return
	}
}
