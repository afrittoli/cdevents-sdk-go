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
	_ "embed"
	"fmt"
	"time"
)

//go:embed spec/schemas/changecreated.json
var changecreatedschema string

var (
	// ChangeCreated event v0.1.2
	ChangeCreatedEventV1 CDEventType = CDEventType{
		Subject:   "change",
		Predicate: "created",
		Version:   "0.1.2",
	}
)

type ChangeCreatedSubjectContent struct {
	Repository Reference `json:"repository"`
}

type ChangeCreatedSubject struct {
	SubjectBase
	Content ChangeCreatedSubjectContent `json:"content"`
}

func (sc ChangeCreatedSubject) GetSubjectType() SubjectType {
	return "change"
}

type ChangeCreatedEvent struct {
	Context Context              `json:"context"`
	Subject ChangeCreatedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e ChangeCreatedEvent) GetType() CDEventType {
	return ChangeCreatedEventV1
}

func (e ChangeCreatedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ChangeCreatedEvent) GetId() string {
	return e.Context.Id
}

func (e ChangeCreatedEvent) GetSource() string {
	return e.Context.Source
}

func (e ChangeCreatedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ChangeCreatedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e ChangeCreatedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ChangeCreatedEvent) GetSubject() Subject {
	return e.Subject
}

func (e ChangeCreatedEvent) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e ChangeCreatedEvent) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e ChangeCreatedEvent) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e ChangeCreatedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *ChangeCreatedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *ChangeCreatedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ChangeCreatedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ChangeCreatedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ChangeCreatedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *ChangeCreatedEvent) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e ChangeCreatedEvent) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), changecreatedschema
}

// Set subject custom fields

func (e *ChangeCreatedEvent) SetSubjectRepository(repository Reference) {
	e.Subject.Content.Repository = repository
}

// New creates a new ChangeCreatedEvent
func NewChangeCreatedEvent() (*ChangeCreatedEvent, error) {
	e := &ChangeCreatedEvent{
		Context: Context{
			Type:    ChangeCreatedEventV1.String(),
			Version: CDEventsSpecVersion,
		},
		Subject: ChangeCreatedSubject{
			SubjectBase: SubjectBase{
				Type: "change",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
