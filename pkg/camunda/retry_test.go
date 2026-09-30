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
	"io"
	"log/slog"
	"reflect"
	"testing"
	"time"

	"github.com/SENERGY-Platform/external-task-worker/lib/camunda/interfaces"
	"github.com/SENERGY-Platform/external-task-worker/lib/messages"
)

// recordingCamunda stands in for the camunda http client and records reported incidents
type recordingCamunda struct {
	interfaces.CamundaInterface
	incidents []string
}

func (this *recordingCamunda) Error(externalTaskId string, _ string, _ string, _ string, msg string, _ string) {
	this.incidents = append(this.incidents, externalTaskId+": "+msg)
}

func (this *recordingCamunda) CompleteTask(_ messages.TaskInfo, _ string, _ interface{}) error {
	return nil
}

func newTestRetry(retries int64) (*RetryBeforeIncident, *recordingCamunda) {
	inner := &recordingCamunda{}
	return NewRetryBeforeIncident(inner, retries, 10*time.Second, slog.New(slog.NewTextHandler(io.Discard, nil))), inner
}

func fail(retry *RetryBeforeIncident, taskId string, msg string, times int) {
	for range times {
		retry.Error(taskId, "instance", "definition", "bk", msg, "")
	}
}

func TestRetryReportsIncidentOnlyAfterRetriesAreExhausted(t *testing.T) {
	retry, inner := newTestRetry(3)
	fail(retry, "task-1", "device not found", 3)
	if len(inner.incidents) != 0 {
		t.Fatalf("expected no incident within the retries, got %#v", inner.incidents)
	}
	fail(retry, "task-1", "device not found", 1)
	if !reflect.DeepEqual(inner.incidents, []string{"task-1: device not found"}) {
		t.Fatalf("expected one incident after the last retry, got %#v", inner.incidents)
	}
}

func TestRetryCountsTasksIndependently(t *testing.T) {
	retry, inner := newTestRetry(1)
	fail(retry, "task-1", "err", 1)
	fail(retry, "task-2", "err", 2)
	if !reflect.DeepEqual(inner.incidents, []string{"task-2: err"}) {
		t.Fatalf("expected only task-2 to be reported, got %#v", inner.incidents)
	}
}

func TestRetryStartsOverAfterCompletedTask(t *testing.T) {
	retry, inner := newTestRetry(1)
	fail(retry, "task-1", "err", 1)
	_ = retry.CompleteTask(messages.TaskInfo{TaskId: "task-1"}, "result", nil)
	fail(retry, "task-1", "err", 1)
	if len(inner.incidents) != 0 {
		t.Fatalf("expected the completed task to start with fresh retries, got %#v", inner.incidents)
	}
}

func TestRetryReportsNotRetryableIncidentsImmediately(t *testing.T) {
	retry, inner := newTestRetry(3)
	fail(retry, "task-1", "user triggered incident: stop", 1)
	fail(retry, "task-2", "invalid task format (json)", 1)
	if !reflect.DeepEqual(inner.incidents, []string{"task-1: user triggered incident: stop", "task-2: invalid task format (json)"}) {
		t.Fatalf("expected both incidents without retry, got %#v", inner.incidents)
	}
}

func TestRetryForgetsAttemptsAfterExpiration(t *testing.T) {
	retry, inner := newTestRetry(1)
	now := time.Now()
	retry.now = func() time.Time { return now }
	fail(retry, "task-1", "err", 1)
	now = now.Add(retry.expiration + time.Second)
	fail(retry, "task-2", "err", 1) // any failure prunes expired entries
	if _, ok := retry.attempts["task-1"]; ok {
		t.Fatal("expected expired attempts of task-1 to be removed")
	}
	if len(inner.incidents) != 0 {
		t.Fatalf("expected no incident, got %#v", inner.incidents)
	}
}
