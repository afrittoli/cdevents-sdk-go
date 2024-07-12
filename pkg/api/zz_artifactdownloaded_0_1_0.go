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

	jsonschema "github.com/santhosh-tekuri/jsonschema/v6"
)

var (
	// ArtifactDownloaded event type v0.1.0
	ArtifactDownloadedEventTypeV0_1_0 CDEventType = CDEventType{
		Subject:   "artifact",
		Predicate: "downloaded",
		Version:   "0.1.0",
	}
)

type ArtifactDownloadedSubjectContentV0_1_0 struct {
	User string `json:"user,omitempty"`
}

type ArtifactDownloadedSubjectV0_1_0 struct {
	SubjectBase
	Content ArtifactDownloadedSubjectContentV0_1_0 `json:"content"`
}

func (sc ArtifactDownloadedSubjectV0_1_0) GetSubjectType() SubjectType {
	return "artifact"
}

type ArtifactDownloadedEventV0_1_0 struct {
	Context ContextV04                      `json:"context"`
	Subject ArtifactDownloadedSubjectV0_1_0 `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e ArtifactDownloadedEventV0_1_0) GetType() CDEventType {
	return ArtifactDownloadedEventTypeV0_1_0
}

func (e ArtifactDownloadedEventV0_1_0) GetVersion() string {
	return e.Context.GetVersion()
}

func (e ArtifactDownloadedEventV0_1_0) GetId() string {
	return e.Context.Id
}

func (e ArtifactDownloadedEventV0_1_0) GetSource() string {
	return e.Context.Source
}

func (e ArtifactDownloadedEventV0_1_0) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ArtifactDownloadedEventV0_1_0) GetSubjectId() string {
	return e.Subject.Id
}

func (e ArtifactDownloadedEventV0_1_0) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ArtifactDownloadedEventV0_1_0) GetSubject() Subject {
	return e.Subject
}

func (e ArtifactDownloadedEventV0_1_0) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e ArtifactDownloadedEventV0_1_0) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e ArtifactDownloadedEventV0_1_0) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e ArtifactDownloadedEventV0_1_0) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsReaderV04 implementation

func (e ArtifactDownloadedEventV0_1_0) GetChainId() string {
	return e.Context.ChainId
}

func (e ArtifactDownloadedEventV0_1_0) GetLinks() EmbeddedLinksArray {
	return e.Context.Links
}

func (e ArtifactDownloadedEventV0_1_0) GetSchemaUri() string {
	return e.Context.SchemaUri
}

// GetCustomSchema looks up the SchemaUri, if any is defined. If none is defined, it returns nil.
// If it's defined and cannot be found, it returns an error.
func (e ArtifactDownloadedEventV0_1_0) GetCustomSchema() (*jsonschema.Schema, error) {
	schemaUri := e.GetSchemaUri()
	if schemaUri == "" {
		return nil, nil
	}
	schema, found := CompiledCustomSchemas[schemaUri]
	if !found {
		return nil, fmt.Errorf("schema with id %s could not be found in the local registry", schemaUri)
	}
	return schema, nil
}

// CDEventsWriter implementation

func (e *ArtifactDownloadedEventV0_1_0) SetId(id string) {
	e.Context.Id = id
}

func (e *ArtifactDownloadedEventV0_1_0) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ArtifactDownloadedEventV0_1_0) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ArtifactDownloadedEventV0_1_0) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ArtifactDownloadedEventV0_1_0) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *ArtifactDownloadedEventV0_1_0) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e ArtifactDownloadedEventV0_1_0) GetSchema() (string, *jsonschema.Schema, error) {
	eType := e.GetType()
	return CompiledSchemas.GetBySpecSubjectPredicate("0.4.1", eType.Subject, eType.Predicate, eType.Custom)
}

// CDEventsWriterV04 implementation

func (e *ArtifactDownloadedEventV0_1_0) SetChainId(chainId string) {
	e.Context.ChainId = chainId
}

func (e *ArtifactDownloadedEventV0_1_0) SetLinks(links EmbeddedLinksArray) {
	e.Context.Links = links
}

func (e *ArtifactDownloadedEventV0_1_0) SetSchemaUri(schema string) {
	e.Context.SchemaUri = schema
}

func (e ArtifactDownloadedEventV0_1_0) GetSubjectContent() interface{} {
	return e.Subject.Content
}

// Set subject custom fields

func (e *ArtifactDownloadedEventV0_1_0) SetSubjectUser(user string) {
	e.Subject.Content.User = user
}

// New creates a new ArtifactDownloadedEventV0_1_0
func NewArtifactDownloadedEventV0_1_0(specVersion string) (*ArtifactDownloadedEventV0_1_0, error) {
	e := &ArtifactDownloadedEventV0_1_0{
		Context: ContextV04{
			Context{
				Type:    ArtifactDownloadedEventTypeV0_1_0,
				Version: specVersion,
			},
			ContextLinks{},
			ContextCustom{},
		},
		Subject: ArtifactDownloadedSubjectV0_1_0{
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
