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
	// ArtifactPackaged event type v0.2.0
	ArtifactPackagedEventTypeV0_2_0 CDEventType = CDEventType{
		Subject:   "artifact",
		Predicate: "packaged",
		Version:   "0.2.0",
	}
)

type ArtifactPackagedSubjectContentV0_2_0 struct {
	Change *Reference `json:"change"`

	Sbom *ArtifactPackagedSubjectContentSbomV0_2_0 `json:"sbom,omitempty"`
}

type ArtifactPackagedSubjectV0_2_0 struct {
	SubjectBase
	Content ArtifactPackagedSubjectContentV0_2_0 `json:"content"`
}

func (sc ArtifactPackagedSubjectV0_2_0) GetSubjectType() SubjectType {
	return "artifact"
}

type ArtifactPackagedEventV0_2_0 struct {
	Context ContextV04                    `json:"context"`
	Subject ArtifactPackagedSubjectV0_2_0 `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e ArtifactPackagedEventV0_2_0) GetType() CDEventType {
	return ArtifactPackagedEventTypeV0_2_0
}

func (e ArtifactPackagedEventV0_2_0) GetVersion() string {
	return e.Context.GetVersion()
}

func (e ArtifactPackagedEventV0_2_0) GetId() string {
	return e.Context.Id
}

func (e ArtifactPackagedEventV0_2_0) GetSource() string {
	return e.Context.Source
}

func (e ArtifactPackagedEventV0_2_0) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ArtifactPackagedEventV0_2_0) GetSubjectId() string {
	return e.Subject.Id
}

func (e ArtifactPackagedEventV0_2_0) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ArtifactPackagedEventV0_2_0) GetSubject() Subject {
	return e.Subject
}

func (e ArtifactPackagedEventV0_2_0) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e ArtifactPackagedEventV0_2_0) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e ArtifactPackagedEventV0_2_0) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e ArtifactPackagedEventV0_2_0) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsReaderV04 implementation

func (e ArtifactPackagedEventV0_2_0) GetChainId() string {
	return e.Context.ChainId
}

func (e ArtifactPackagedEventV0_2_0) GetLinks() EmbeddedLinksArray {
	return e.Context.Links
}

func (e ArtifactPackagedEventV0_2_0) GetSchemaUri() string {
	return e.Context.SchemaUri
}

// CDEventsWriter implementation

func (e *ArtifactPackagedEventV0_2_0) SetId(id string) {
	e.Context.Id = id
}

func (e *ArtifactPackagedEventV0_2_0) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ArtifactPackagedEventV0_2_0) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ArtifactPackagedEventV0_2_0) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ArtifactPackagedEventV0_2_0) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *ArtifactPackagedEventV0_2_0) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e ArtifactPackagedEventV0_2_0) GetSchema() (string, *jsonschema.Schema, error) {
	eType := e.GetType()
	return CompiledSchemas.GetBySpecSubjectPredicate("0.4.1", eType.Subject, eType.Predicate, eType.Custom)
}

// CDEventsWriterV04 implementation

func (e *ArtifactPackagedEventV0_2_0) SetChainId(chainId string) {
	e.Context.ChainId = chainId
}

func (e *ArtifactPackagedEventV0_2_0) SetLinks(links EmbeddedLinksArray) {
	e.Context.Links = links
}

func (e *ArtifactPackagedEventV0_2_0) SetSchemaUri(schema string) {
	e.Context.SchemaUri = schema
}

func (e ArtifactPackagedEventV0_2_0) GetSubjectContent() interface{} {
	return e.Subject.Content
}

// Set subject custom fields

func (e *ArtifactPackagedEventV0_2_0) SetSubjectChange(change *Reference) {
	e.Subject.Content.Change = change
}

func (e *ArtifactPackagedEventV0_2_0) SetSubjectSbom(sbom *ArtifactPackagedSubjectContentSbomV0_2_0) {
	e.Subject.Content.Sbom = sbom
}

// New creates a new ArtifactPackagedEventV0_2_0
func NewArtifactPackagedEventV0_2_0(specVersion string) (*ArtifactPackagedEventV0_2_0, error) {
	e := &ArtifactPackagedEventV0_2_0{
		Context: ContextV04{
			Context{
				Type:    ArtifactPackagedEventTypeV0_2_0,
				Version: specVersion,
			},
			ContextLinks{},
			ContextCustom{},
		},
		Subject: ArtifactPackagedSubjectV0_2_0{
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

// ArtifactPackagedSubjectContentSbomV0_2_0 holds the content of a Sbom field in the content
type ArtifactPackagedSubjectContentSbomV0_2_0 struct {
	Uri string `json:"uri"`
}
