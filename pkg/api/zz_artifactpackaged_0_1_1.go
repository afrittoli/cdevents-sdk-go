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

var artifactpackagedschema0_1_1 = `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"https://cdevents.dev/0.3.0/schema/artifact-packaged-event","properties":{"context":{"properties":{"version":{"type":"string","minLength":1},"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","enum":["dev.cdevents.artifact.packaged.0.1.1"],"default":"dev.cdevents.artifact.packaged.0.1.1"},"timestamp":{"type":"string","format":"date-time"}},"additionalProperties":false,"type":"object","required":["version","id","source","type","timestamp"]},"subject":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","minLength":1,"enum":["artifact"],"default":"artifact"},"content":{"properties":{"change":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"}},"additionalProperties":false,"type":"object","required":["id"]}},"additionalProperties":false,"type":"object","required":["change"]}},"additionalProperties":false,"type":"object","required":["id","type","content"]},"customData":{"oneOf":[{"type":"object"},{"type":"string","contentEncoding":"base64"}]},"customDataContentType":{"type":"string"}},"additionalProperties":false,"type":"object","required":["context","subject"]}`

var (
	// ArtifactPackaged event type v0.1.1
	ArtifactPackagedEventTypeV0_1_1 CDEventType = CDEventType{
		Subject:   "artifact",
		Predicate: "packaged",
		Version:   "0.1.1",
	}
)

type ArtifactPackagedSubjectContent struct {
	Change *Reference `json:"change"`
}

type ArtifactPackagedSubject struct {
	SubjectBase
	Content ArtifactPackagedSubjectContent `json:"content"`
}

func (sc ArtifactPackagedSubject) GetSubjectType() SubjectType {
	return "artifact"
}

type ArtifactPackagedEventV0_1_1 struct {
	Context Context                 `json:"context"`
	Subject ArtifactPackagedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e ArtifactPackagedEventV0_1_1) GetType() CDEventType {
	return ArtifactPackagedEventTypeV0_1_1
}

func (e ArtifactPackagedEventV0_1_1) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ArtifactPackagedEventV0_1_1) GetId() string {
	return e.Context.Id
}

func (e ArtifactPackagedEventV0_1_1) GetSource() string {
	return e.Context.Source
}

func (e ArtifactPackagedEventV0_1_1) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ArtifactPackagedEventV0_1_1) GetSubjectId() string {
	return e.Subject.Id
}

func (e ArtifactPackagedEventV0_1_1) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ArtifactPackagedEventV0_1_1) GetSubject() Subject {
	return e.Subject
}

func (e ArtifactPackagedEventV0_1_1) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e ArtifactPackagedEventV0_1_1) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e ArtifactPackagedEventV0_1_1) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e ArtifactPackagedEventV0_1_1) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *ArtifactPackagedEventV0_1_1) SetId(id string) {
	e.Context.Id = id
}

func (e *ArtifactPackagedEventV0_1_1) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ArtifactPackagedEventV0_1_1) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ArtifactPackagedEventV0_1_1) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ArtifactPackagedEventV0_1_1) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *ArtifactPackagedEventV0_1_1) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e ArtifactPackagedEventV0_1_1) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), artifactpackagedschema0_1_1
}

// Set subject custom fields

func (e *ArtifactPackagedEventV0_1_1) SetSubjectChange(change *Reference) {
	e.Subject.Content.Change = change
}

// New creates a new ArtifactPackagedEventV0_1_1
func NewArtifactPackagedEventV0_1_1(specVersion string) (*ArtifactPackagedEventV0_1_1, error) {
	e := &ArtifactPackagedEventV0_1_1{
		Context: Context{
			Type:    ArtifactPackagedEventTypeV0_1_1,
			Version: specVersion,
		},
		Subject: ArtifactPackagedSubject{
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
