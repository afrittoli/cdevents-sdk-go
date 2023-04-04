// Code generated by tools/generator. DO NOT EDIT.

/*
Copyright 2023 The CDEvents Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

SPDX-License-Identifier: Apache-2.0
*/

package api

import (
	"fmt"
	"time"
)

var taskrunfinishedschema = `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"https://cdevents.dev/0.2.0/schema/task-run-finished-event","properties":{"context":{"properties":{"version":{"type":"string","minLength":1},"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","enum":["dev.cdevents.taskrun.finished.0.1.1"],"default":"dev.cdevents.taskrun.finished.0.1.1"},"timestamp":{"type":"string","format":"date-time"}},"additionalProperties":false,"type":"object","required":["version","id","source","type","timestamp"]},"subject":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","minLength":1,"enum":["taskRun"],"default":"taskRun"},"content":{"properties":{"taskName":{"type":"string"},"url":{"type":"string"},"pipelineRun":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"}},"additionalProperties":false,"type":"object","required":["id"]},"outcome":{"type":"string"},"errors":{"type":"string"}},"additionalProperties":false,"type":"object"}},"additionalProperties":false,"type":"object","required":["id","type","content"]},"customData":{"oneOf":[{"type":"object"},{"type":"string","contentEncoding":"base64"}]},"customDataContentType":{"type":"string"}},"additionalProperties":false,"type":"object","required":["context","subject"]}`

var (
	// TaskRunFinished event v0.1.1
	TaskRunFinishedEventV1 CDEventType = CDEventType{
		Subject:   "taskrun",
		Predicate: "finished",
		Version:   "0.1.1",
	}
)

type TaskRunFinishedSubjectContent struct {
	Errors string `json:"errors"`

	Outcome string `json:"outcome"`

	PipelineRun Reference `json:"pipelineRun"`

	TaskName string `json:"taskName"`

	Url string `json:"url"`
}

type TaskRunFinishedSubject struct {
	SubjectBase
	Content TaskRunFinishedSubjectContent `json:"content"`
}

func (sc TaskRunFinishedSubject) GetSubjectType() SubjectType {
	return "taskRun"
}

type TaskRunFinishedEvent struct {
	Context Context                `json:"context"`
	Subject TaskRunFinishedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e TaskRunFinishedEvent) GetType() CDEventType {
	return TaskRunFinishedEventV1
}

func (e TaskRunFinishedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e TaskRunFinishedEvent) GetId() string {
	return e.Context.Id
}

func (e TaskRunFinishedEvent) GetSource() string {
	return e.Context.Source
}

func (e TaskRunFinishedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e TaskRunFinishedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e TaskRunFinishedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e TaskRunFinishedEvent) GetSubject() Subject {
	return e.Subject
}

func (e TaskRunFinishedEvent) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e TaskRunFinishedEvent) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e TaskRunFinishedEvent) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e TaskRunFinishedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *TaskRunFinishedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *TaskRunFinishedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *TaskRunFinishedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *TaskRunFinishedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *TaskRunFinishedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *TaskRunFinishedEvent) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e TaskRunFinishedEvent) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), taskrunfinishedschema
}

// Set subject custom fields

func (e *TaskRunFinishedEvent) SetSubjectErrors(errors string) {
	e.Subject.Content.Errors = errors
}

func (e *TaskRunFinishedEvent) SetSubjectOutcome(outcome string) {
	e.Subject.Content.Outcome = outcome
}

func (e *TaskRunFinishedEvent) SetSubjectPipelineRun(pipelineRun Reference) {
	e.Subject.Content.PipelineRun = pipelineRun
}

func (e *TaskRunFinishedEvent) SetSubjectTaskName(taskName string) {
	e.Subject.Content.TaskName = taskName
}

func (e *TaskRunFinishedEvent) SetSubjectUrl(url string) {
	e.Subject.Content.Url = url
}

// New creates a new TaskRunFinishedEvent
func NewTaskRunFinishedEvent() (*TaskRunFinishedEvent, error) {
	e := &TaskRunFinishedEvent{
		Context: Context{
			Type:    TaskRunFinishedEventV1.String(),
			Version: CDEventsSpecVersion,
		},
		Subject: TaskRunFinishedSubject{
			SubjectBase: SubjectBase{
				Type: "taskRun",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
