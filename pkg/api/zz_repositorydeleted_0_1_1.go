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

var repositorydeletedschema = `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"https://cdevents.dev/0.3.0/schema/repository-deleted-event","properties":{"context":{"properties":{"version":{"type":"string","minLength":1},"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","enum":["dev.cdevents.repository.deleted.0.1.1"],"default":"dev.cdevents.repository.deleted.0.1.1"},"timestamp":{"type":"string","format":"date-time"}},"additionalProperties":false,"type":"object","required":["version","id","source","type","timestamp"]},"subject":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","minLength":1,"enum":["repository"],"default":"repository"},"content":{"properties":{"name":{"type":"string"},"owner":{"type":"string"},"url":{"type":"string"},"viewUrl":{"type":"string"}},"additionalProperties":false,"type":"object"}},"additionalProperties":false,"type":"object","required":["id","type","content"]},"customData":{"oneOf":[{"type":"object"},{"type":"string","contentEncoding":"base64"}]},"customDataContentType":{"type":"string"}},"additionalProperties":false,"type":"object","required":["context","subject"]}`

var (
	// RepositoryDeleted event type v0.1.1
	RepositoryDeletedEventTypeV0_1_1 CDEventType = CDEventType{
		Subject:   "repository",
		Predicate: "deleted",
		Version:   "0.1.1",
	}
)

type RepositoryDeletedSubjectContent struct {
	Name string `json:"name,omitempty"`

	Owner string `json:"owner,omitempty"`

	Url string `json:"url,omitempty"`

	ViewUrl string `json:"viewUrl,omitempty"`
}

type RepositoryDeletedSubject struct {
	SubjectBase
	Content RepositoryDeletedSubjectContent `json:"content"`
}

func (sc RepositoryDeletedSubject) GetSubjectType() SubjectType {
	return "repository"
}

type RepositoryDeletedEventV0_1_1 struct {
	Context Context                  `json:"context"`
	Subject RepositoryDeletedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e RepositoryDeletedEventV0_1_1) GetType() CDEventType {
	return RepositoryDeletedEventTypeV0_1_1
}

func (e RepositoryDeletedEventV0_1_1) GetVersion() string {
	return CDEventsSpecVersion
}

func (e RepositoryDeletedEventV0_1_1) GetId() string {
	return e.Context.Id
}

func (e RepositoryDeletedEventV0_1_1) GetSource() string {
	return e.Context.Source
}

func (e RepositoryDeletedEventV0_1_1) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e RepositoryDeletedEventV0_1_1) GetSubjectId() string {
	return e.Subject.Id
}

func (e RepositoryDeletedEventV0_1_1) GetSubjectSource() string {
	return e.Subject.Source
}

func (e RepositoryDeletedEventV0_1_1) GetSubject() Subject {
	return e.Subject
}

func (e RepositoryDeletedEventV0_1_1) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e RepositoryDeletedEventV0_1_1) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e RepositoryDeletedEventV0_1_1) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e RepositoryDeletedEventV0_1_1) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *RepositoryDeletedEventV0_1_1) SetId(id string) {
	e.Context.Id = id
}

func (e *RepositoryDeletedEventV0_1_1) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *RepositoryDeletedEventV0_1_1) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *RepositoryDeletedEventV0_1_1) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *RepositoryDeletedEventV0_1_1) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *RepositoryDeletedEventV0_1_1) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e RepositoryDeletedEventV0_1_1) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), repositorydeletedschema
}

// Set subject custom fields

func (e *RepositoryDeletedEventV0_1_1) SetSubjectName(name string) {
	e.Subject.Content.Name = name
}

func (e *RepositoryDeletedEventV0_1_1) SetSubjectOwner(owner string) {
	e.Subject.Content.Owner = owner
}

func (e *RepositoryDeletedEventV0_1_1) SetSubjectUrl(url string) {
	e.Subject.Content.Url = url
}

func (e *RepositoryDeletedEventV0_1_1) SetSubjectViewUrl(viewUrl string) {
	e.Subject.Content.ViewUrl = viewUrl
}

// New creates a new RepositoryDeletedEventV0_1_1
func NewRepositoryDeletedEventV0_1_1(specVersion string) (*RepositoryDeletedEventV0_1_1, error) {
	e := &RepositoryDeletedEventV0_1_1{
		Context: Context{
			Type:    RepositoryDeletedEventTypeV0_1_1,
			Version: specVersion,
		},
		Subject: RepositoryDeletedSubject{
			SubjectBase: SubjectBase{
				Type: "repository",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
