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

package camunda

import (
	"log/slog"
	"strings"
	"sync"
	"time"

	"github.com/SENERGY-Platform/external-task-worker/lib/camunda/interfaces"
	"github.com/SENERGY-Platform/external-task-worker/lib/messages"
)

// messages of incidents that fail the same way on every attempt (see external-task-worker lib/worker.go ExecuteTask())
var notRetryableIncidentPrefixes = []string{
	"user triggered incident:",
	"invalid task format (json)",
}

// minimal time an attempt count is kept for a task that is neither completed nor reported as incident
const minAttemptExpiration = time.Hour

// RetryBeforeIncident delays the incident of a failed task until it failed retries+1 times.
// A withheld incident leaves the task locked; camunda hands it out again when the lock expires,
// so the camunda fetch lock duration is the retry interval.
type RetryBeforeIncident struct {
	interfaces.CamundaInterface
	retries    int64
	expiration time.Duration
	logger     *slog.Logger
	now        func() time.Time
	mux        sync.Mutex
	attempts   map[string]taskAttempts
}

type taskAttempts struct {
	failed int64
	last   time.Time
}

func NewRetryBeforeIncident(camunda interfaces.CamundaInterface, retries int64, lockDuration time.Duration, logger *slog.Logger) *RetryBeforeIncident {
	expiration := 10 * lockDuration * time.Duration(retries+1)
	if expiration < minAttemptExpiration {
		expiration = minAttemptExpiration
	}
	return &RetryBeforeIncident{
		CamundaInterface: camunda,
		retries:          retries,
		expiration:       expiration,
		logger:           logger,
		now:              time.Now,
		attempts:         map[string]taskAttempts{},
	}
}

func (this *RetryBeforeIncident) Error(externalTaskId string, processInstanceId string, processDefinitionId string, businessKey string, msg string, tenantId string) {
	if this.retries > 0 && isRetryable(msg) && this.withhold(externalTaskId) {
		this.logger.Warn("task failed, incident withheld until retries are exhausted", "taskId", externalTaskId, "processInstanceId", processInstanceId, "businessKey", businessKey, "retries", this.retries, "error", msg)
		return
	}
	this.CamundaInterface.Error(externalTaskId, processInstanceId, processDefinitionId, businessKey, msg, tenantId)
}

func (this *RetryBeforeIncident) CompleteTask(taskInfo messages.TaskInfo, outputName string, output interface{}) (err error) {
	this.forget(taskInfo.TaskId)
	return this.CamundaInterface.CompleteTask(taskInfo, outputName, output)
}

// withhold counts the failed attempt and reports whether the incident has to be withheld
func (this *RetryBeforeIncident) withhold(taskId string) bool {
	this.mux.Lock()
	defer this.mux.Unlock()
	now := this.now()
	for id, attempts := range this.attempts {
		if now.Sub(attempts.last) > this.expiration {
			delete(this.attempts, id)
		}
	}
	attempts := this.attempts[taskId]
	if attempts.failed >= this.retries {
		delete(this.attempts, taskId)
		return false
	}
	this.attempts[taskId] = taskAttempts{failed: attempts.failed + 1, last: now}
	return true
}

func (this *RetryBeforeIncident) forget(taskId string) {
	this.mux.Lock()
	defer this.mux.Unlock()
	delete(this.attempts, taskId)
}

func isRetryable(msg string) bool {
	for _, prefix := range notRetryableIncidentPrefixes {
		if strings.HasPrefix(msg, prefix) {
			return false
		}
	}
	return true
}
