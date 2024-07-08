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
	// BuildFinished event type v0.1.1
	BuildFinishedEventTypeV0_1_1 CDEventType = CDEventType{
		Subject:   "build",
		Predicate: "finished",
		Version:   "0.1.1",
	}
)

type BuildFinishedSubjectContentV0_1_1 struct {
	ArtifactId string `json:"artifactId,omitempty" validate:"purl"`
}

type BuildFinishedSubjectV0_1_1 struct {
	SubjectBase
	Content BuildFinishedSubjectContentV0_1_1 `json:"content"`
}

func (sc BuildFinishedSubjectV0_1_1) GetSubjectType() SubjectType {
	return "build"
}

type BuildFinishedEventV0_1_1 struct {
	Context Context                    `json:"context"`
	Subject BuildFinishedSubjectV0_1_1 `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e BuildFinishedEventV0_1_1) GetType() CDEventType {
	return BuildFinishedEventTypeV0_1_1
}

func (e BuildFinishedEventV0_1_1) GetVersion() string {
	return e.Context.GetVersion()
}

func (e BuildFinishedEventV0_1_1) GetId() string {
	return e.Context.Id
}

func (e BuildFinishedEventV0_1_1) GetSource() string {
	return e.Context.Source
}

func (e BuildFinishedEventV0_1_1) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e BuildFinishedEventV0_1_1) GetSubjectId() string {
	return e.Subject.Id
}

func (e BuildFinishedEventV0_1_1) GetSubjectSource() string {
	return e.Subject.Source
}

func (e BuildFinishedEventV0_1_1) GetSubject() Subject {
	return e.Subject
}

func (e BuildFinishedEventV0_1_1) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e BuildFinishedEventV0_1_1) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e BuildFinishedEventV0_1_1) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e BuildFinishedEventV0_1_1) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *BuildFinishedEventV0_1_1) SetId(id string) {
	e.Context.Id = id
}

func (e *BuildFinishedEventV0_1_1) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *BuildFinishedEventV0_1_1) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *BuildFinishedEventV0_1_1) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *BuildFinishedEventV0_1_1) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *BuildFinishedEventV0_1_1) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e BuildFinishedEventV0_1_1) GetSchema() (string, *jsonschema.Schema, error) {
	eType := e.GetType()
	return CompiledSchemas.GetBySpecSubjectPredicate("0.3.0", eType.Subject, eType.Predicate, eType.Custom)
}

func (e BuildFinishedEventV0_1_1) GetSubjectContent() interface{} {
	return e.Subject.Content
}

// Set subject custom fields

func (e *BuildFinishedEventV0_1_1) SetSubjectArtifactId(artifactId string) {
	e.Subject.Content.ArtifactId = artifactId
}

// New creates a new BuildFinishedEventV0_1_1
func NewBuildFinishedEventV0_1_1(specVersion string) (*BuildFinishedEventV0_1_1, error) {
	e := &BuildFinishedEventV0_1_1{
		Context: Context{
			Type:    BuildFinishedEventTypeV0_1_1,
			Version: specVersion,
		},
		Subject: BuildFinishedSubjectV0_1_1{
			SubjectBase: SubjectBase{
				Type: "build",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
