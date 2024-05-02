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

var artifactpublishedschema = `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"https://cdevents.dev/0.3.0/schema/artifact-published-event","properties":{"context":{"properties":{"version":{"type":"string","minLength":1},"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","enum":["dev.cdevents.artifact.published.0.1.1"],"default":"dev.cdevents.artifact.published.0.1.1"},"timestamp":{"type":"string","format":"date-time"}},"additionalProperties":false,"type":"object","required":["version","id","source","type","timestamp"]},"subject":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","minLength":1,"enum":["artifact"],"default":"artifact"},"content":{"properties":{},"additionalProperties":false,"type":"object"}},"additionalProperties":false,"type":"object","required":["id","type","content"]},"customData":{"oneOf":[{"type":"object"},{"type":"string","contentEncoding":"base64"}]},"customDataContentType":{"type":"string"}},"additionalProperties":false,"type":"object","required":["context","subject"]}`

var (
	// ArtifactPublished event type v0.1.1
	ArtifactPublishedEventTypeV0_1_1 CDEventType = CDEventType{
		Subject:   "artifact",
		Predicate: "published",
		Version:   "0.1.1",
	}
)

type ArtifactPublishedSubjectContent struct {
}

type ArtifactPublishedSubject struct {
	SubjectBase
	Content ArtifactPublishedSubjectContent `json:"content"`
}

func (sc ArtifactPublishedSubject) GetSubjectType() SubjectType {
	return "artifact"
}

type ArtifactPublishedEventV0_1_1 struct {
	Context Context                  `json:"context"`
	Subject ArtifactPublishedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e ArtifactPublishedEventV0_1_1) GetType() CDEventType {
	return ArtifactPublishedEventTypeV0_1_1
}

func (e ArtifactPublishedEventV0_1_1) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ArtifactPublishedEventV0_1_1) GetId() string {
	return e.Context.Id
}

func (e ArtifactPublishedEventV0_1_1) GetSource() string {
	return e.Context.Source
}

func (e ArtifactPublishedEventV0_1_1) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ArtifactPublishedEventV0_1_1) GetSubjectId() string {
	return e.Subject.Id
}

func (e ArtifactPublishedEventV0_1_1) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ArtifactPublishedEventV0_1_1) GetSubject() Subject {
	return e.Subject
}

func (e ArtifactPublishedEventV0_1_1) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e ArtifactPublishedEventV0_1_1) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e ArtifactPublishedEventV0_1_1) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e ArtifactPublishedEventV0_1_1) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *ArtifactPublishedEventV0_1_1) SetId(id string) {
	e.Context.Id = id
}

func (e *ArtifactPublishedEventV0_1_1) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ArtifactPublishedEventV0_1_1) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ArtifactPublishedEventV0_1_1) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ArtifactPublishedEventV0_1_1) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *ArtifactPublishedEventV0_1_1) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e ArtifactPublishedEventV0_1_1) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), artifactpublishedschema
}

// Set subject custom fields

// New creates a new ArtifactPublishedEventV0_1_1
func NewArtifactPublishedEventV0_1_1() (*ArtifactPublishedEventV0_1_1, error) {
	e := &ArtifactPublishedEventV0_1_1{
		Context: Context{
			Type:    ArtifactPublishedEventTypeV0_1_1.String(),
			Version: CDEventsSpecVersion,
		},
		Subject: ArtifactPublishedSubject{
			SubjectBase: SubjectBase{
				Type: "artifact",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}