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

package configuration

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Config struct {
	CompletionStrategy              string `json:"completion_strategy"`
	OptimisticTaskCompletionTimeout int64  `json:"optimistic_task_completion_timeout"`
	CamundaWorkerTasks              int64  `json:"camunda_worker_tasks"`
	CamundaFetchLockDuration        int64  `json:"camunda_fetch_lock_duration"`
	CamundaTopic                    string `json:"camunda_topic"`
	CamundaTaskResultName           string `json:"camunda_task_result_name"`
	CamundaLongPollTimeout          int64  `json:"camunda_long_poll_timeout"`

	TimescaleWrapperUrl string `json:"timescale_wrapper_url"`

	AuthExpirationTimeBuffer     float64 `json:"auth_expiration_time_buffer"`
	AuthEndpoint                 string  `json:"auth_endpoint"`
	AuthClientId                 string  `json:"auth_client_id" config:"secret"`
	AuthUserName                 string  `json:"auth_user_name" config:"secret"`
	AuthPassword                 string  `json:"auth_password" config:"secret"`
	Debug                        bool    `json:"debug"`
	SubResultExpirationInSeconds int32   `json:"sub_result_expiration_in_seconds"`
	SequentialGroups             bool    `json:"sequential_groups"`
	CamundaUrl                   string  `json:"camunda_url"`

	MqttBroker          string `json:"mqtt_broker"`
	MqttResponseTopic   string `json:"mqtt_response_topic"`
	CorrelationIdPrefix string `json:"correlation_id_prefix"`
	DeviceRepoUrl       string `json:"device_repo_url"`
	FallbackFile        string `json:"fallback_file"`

	ProtocolHandler string `json:"protocol_handler"`
	ProtocolSegment string `json:"protocol_segment"`

	MgwConceptRepoRefreshInterval int64 `json:"mgw_concept_repo_refresh_interval"`

	//optional; use if running as standalone; dont use if running with https://github.com/SENERGY-Platform/senergy-connector
	SyncMqttBroker           string `json:"sync_mqtt_broker"`
	SyncNetworkId            string `json:"sync_network_id"`
	SyncConnectRetryInterval string `json:"sync_connect_retry_interval"`
}

// loads config from json in location and used environment variables (e.g ZookeeperUrl --> ZOOKEEPER_URL)
func Load(location string) (config Config, err error) {
	file, err := os.Open(location)
	if err != nil {
		return config, err
	}
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return config, err
	}
	handleEnvironmentVars(&config)
	return config, nil
}

func (this Config) AuthEnabled() bool {
	return this.AuthEndpoint != "" && this.AuthEndpoint != "-"
}

var camel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")

func fieldNameToEnvName(s string) string {
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	return strings.ToUpper(strings.Join(a, "_"))
}

// preparations for docker
func handleEnvironmentVars(config *Config) {
	configValue := reflect.Indirect(reflect.ValueOf(config))
	configType := configValue.Type()
	for index := 0; index < configType.NumField(); index++ {
		fieldName := configType.Field(index).Name
		fieldConfig := configType.Field(index).Tag.Get("config")
		envName := fieldNameToEnvName(fieldName)
		envValue := os.Getenv(envName)
		if envValue != "" {
			loggedEnvValue := envValue
			if strings.Contains(fieldConfig, "secret") {
				loggedEnvValue = "***"
			}
			fmt.Println("use environment variable: ", envName, " = ", loggedEnvValue)
			if configValue.FieldByName(fieldName).Kind() == reflect.Int64 {
				i, _ := strconv.ParseInt(envValue, 10, 64)
				configValue.FieldByName(fieldName).SetInt(i)
			}
			if configValue.FieldByName(fieldName).Kind() == reflect.String {
				configValue.FieldByName(fieldName).SetString(envValue)
			}
			if configValue.FieldByName(fieldName).Kind() == reflect.Bool {
				b, _ := strconv.ParseBool(envValue)
				configValue.FieldByName(fieldName).SetBool(b)
			}
			if configValue.FieldByName(fieldName).Kind() == reflect.Float64 {
				f, _ := strconv.ParseFloat(envValue, 64)
				configValue.FieldByName(fieldName).SetFloat(f)
			}
			if configValue.FieldByName(fieldName).Kind() == reflect.Slice {
				val := []string{}
				for _, element := range strings.Split(envValue, ",") {
					val = append(val, strings.TrimSpace(element))
				}
				configValue.FieldByName(fieldName).Set(reflect.ValueOf(val))
			}
			if configValue.FieldByName(fieldName).Kind() == reflect.Map {
				value := map[string]string{}
				for _, element := range strings.Split(envValue, ",") {
					keyVal := strings.Split(element, ":")
					key := strings.TrimSpace(keyVal[0])
					val := strings.TrimSpace(keyVal[1])
					value[key] = val
				}
				configValue.FieldByName(fieldName).Set(reflect.ValueOf(value))
			}
		}
	}
}
