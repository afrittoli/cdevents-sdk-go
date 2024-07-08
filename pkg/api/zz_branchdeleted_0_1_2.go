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
	"time"

	jsonschema "github.com/santhosh-tekuri/jsonschema/v6"
)

var (
	// BranchDeleted event type v0.1.2
	BranchDeletedEventTypeV0_1_2 CDEventType = CDEventType{
		Subject:   "branch",
		Predicate: "deleted",
		Version:   "0.1.2",
	}
)

type BranchDeletedSubjectContentV0_1_2 struct {
	Repository *Reference `json:"repository,omitempty"`
}

type BranchDeletedSubjectV0_1_2 struct {
	SubjectBase
	Content BranchDeletedSubjectContentV0_1_2 `json:"content"`
}

func (sc BranchDeletedSubjectV0_1_2) GetSubjectType() SubjectType {
	return "branch"
}

type BranchDeletedEventV0_1_2 struct {
	Context Context                    `json:"context"`
	Subject BranchDeletedSubjectV0_1_2 `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e BranchDeletedEventV0_1_2) GetType() CDEventType {
	return BranchDeletedEventTypeV0_1_2
}

func (e BranchDeletedEventV0_1_2) GetVersion() string {
	return e.Context.GetVersion()
}

func (e BranchDeletedEventV0_1_2) GetId() string {
	return e.Context.Id
}

func (e BranchDeletedEventV0_1_2) GetSource() string {
	return e.Context.Source
}

func (e BranchDeletedEventV0_1_2) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e BranchDeletedEventV0_1_2) GetSubjectId() string {
	return e.Subject.Id
}

func (e BranchDeletedEventV0_1_2) GetSubjectSource() string {
	return e.Subject.Source
}

func (e BranchDeletedEventV0_1_2) GetSubject() Subject {
	return e.Subject
}

func (e BranchDeletedEventV0_1_2) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e BranchDeletedEventV0_1_2) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e BranchDeletedEventV0_1_2) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e BranchDeletedEventV0_1_2) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *BranchDeletedEventV0_1_2) SetId(id string) {
	e.Context.Id = id
}

func (e *BranchDeletedEventV0_1_2) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *BranchDeletedEventV0_1_2) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *BranchDeletedEventV0_1_2) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *BranchDeletedEventV0_1_2) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *BranchDeletedEventV0_1_2) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e BranchDeletedEventV0_1_2) GetSchema() (string, *jsonschema.Schema, error) {
	eType := e.GetType()
	return CompiledSchemas.GetBySpecSubjectPredicate("0.3.0", eType.Subject, eType.Predicate, eType.Custom)
}

func (e BranchDeletedEventV0_1_2) GetSubjectContent() interface{} {
	return e.Subject.Content
}

// Set subject custom fields

func (e *BranchDeletedEventV0_1_2) SetSubjectRepository(repository *Reference) {
	e.Subject.Content.Repository = repository
}

// New creates a new BranchDeletedEventV0_1_2
func NewBranchDeletedEventV0_1_2(specVersion string) (*BranchDeletedEventV0_1_2, error) {
	e := &BranchDeletedEventV0_1_2{
		Context: Context{
			Type:    BranchDeletedEventTypeV0_1_2,
			Version: specVersion,
		},
		Subject: BranchDeletedSubjectV0_1_2{
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
