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

var branchcreatedschema = `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"https://cdevents.dev/0.2.0/schema/branch-created-event","properties":{"context":{"properties":{"version":{"type":"string","minLength":1},"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","enum":["dev.cdevents.branch.created.0.1.2"],"default":"dev.cdevents.branch.created.0.1.2"},"timestamp":{"type":"string","format":"date-time"}},"additionalProperties":false,"type":"object","required":["version","id","source","type","timestamp"]},"subject":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","minLength":1,"enum":["branch"],"default":"branch"},"content":{"properties":{"repository":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"}},"additionalProperties":false,"type":"object","required":["id"]}},"additionalProperties":false,"type":"object"}},"additionalProperties":false,"type":"object","required":["id","type","content"]},"customData":{"oneOf":[{"type":"object"},{"type":"string","contentEncoding":"base64"}]},"customDataContentType":{"type":"string"}},"additionalProperties":false,"type":"object","required":["context","subject"]}`

var (
	// BranchCreated event v0.1.2
	BranchCreatedEventV1 CDEventType = CDEventType{
		Subject:   "branch",
		Predicate: "created",
		Version:   "0.1.2",
	}
)

type BranchCreatedSubjectContent struct {
	Repository Reference `json:"repository"`
}

type BranchCreatedSubject struct {
	SubjectBase
	Content BranchCreatedSubjectContent `json:"content"`
}

func (sc BranchCreatedSubject) GetSubjectType() SubjectType {
	return "branch"
}

type BranchCreatedEvent struct {
	Context Context              `json:"context"`
	Subject BranchCreatedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e BranchCreatedEvent) GetType() CDEventType {
	return BranchCreatedEventV1
}

func (e BranchCreatedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e BranchCreatedEvent) GetId() string {
	return e.Context.Id
}

func (e BranchCreatedEvent) GetSource() string {
	return e.Context.Source
}

func (e BranchCreatedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e BranchCreatedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e BranchCreatedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e BranchCreatedEvent) GetSubject() Subject {
	return e.Subject
}

func (e BranchCreatedEvent) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e BranchCreatedEvent) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e BranchCreatedEvent) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e BranchCreatedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *BranchCreatedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *BranchCreatedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *BranchCreatedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *BranchCreatedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *BranchCreatedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *BranchCreatedEvent) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e BranchCreatedEvent) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), branchcreatedschema
}

// Set subject custom fields

func (e *BranchCreatedEvent) SetSubjectRepository(repository Reference) {
	e.Subject.Content.Repository = repository
}

// New creates a new BranchCreatedEvent
func NewBranchCreatedEvent() (*BranchCreatedEvent, error) {
	e := &BranchCreatedEvent{
		Context: Context{
			Type:    BranchCreatedEventV1.String(),
			Version: CDEventsSpecVersion,
		},
		Subject: BranchCreatedSubject{
			SubjectBase: SubjectBase{
				Type: "branch",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
