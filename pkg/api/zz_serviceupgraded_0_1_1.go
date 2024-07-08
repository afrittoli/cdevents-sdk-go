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
	// ServiceUpgraded event type v0.1.1
	ServiceUpgradedEventTypeV0_1_1 CDEventType = CDEventType{
		Subject:   "service",
		Predicate: "upgraded",
		Version:   "0.1.1",
	}
)

type ServiceUpgradedSubjectContentV0_1_1 struct {
	ArtifactId string `json:"artifactId" validate:"purl"`

	Environment *Reference `json:"environment"`
}

type ServiceUpgradedSubjectV0_1_1 struct {
	SubjectBase
	Content ServiceUpgradedSubjectContentV0_1_1 `json:"content"`
}

func (sc ServiceUpgradedSubjectV0_1_1) GetSubjectType() SubjectType {
	return "service"
}

type ServiceUpgradedEventV0_1_1 struct {
	Context Context                      `json:"context"`
	Subject ServiceUpgradedSubjectV0_1_1 `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e ServiceUpgradedEventV0_1_1) GetType() CDEventType {
	return ServiceUpgradedEventTypeV0_1_1
}

func (e ServiceUpgradedEventV0_1_1) GetVersion() string {
	return e.Context.GetVersion()
}

func (e ServiceUpgradedEventV0_1_1) GetId() string {
	return e.Context.Id
}

func (e ServiceUpgradedEventV0_1_1) GetSource() string {
	return e.Context.Source
}

func (e ServiceUpgradedEventV0_1_1) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ServiceUpgradedEventV0_1_1) GetSubjectId() string {
	return e.Subject.Id
}

func (e ServiceUpgradedEventV0_1_1) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ServiceUpgradedEventV0_1_1) GetSubject() Subject {
	return e.Subject
}

func (e ServiceUpgradedEventV0_1_1) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e ServiceUpgradedEventV0_1_1) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e ServiceUpgradedEventV0_1_1) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e ServiceUpgradedEventV0_1_1) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *ServiceUpgradedEventV0_1_1) SetId(id string) {
	e.Context.Id = id
}

func (e *ServiceUpgradedEventV0_1_1) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ServiceUpgradedEventV0_1_1) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ServiceUpgradedEventV0_1_1) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ServiceUpgradedEventV0_1_1) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *ServiceUpgradedEventV0_1_1) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e ServiceUpgradedEventV0_1_1) GetSchema() (string, *jsonschema.Schema, error) {
	eType := e.GetType()
	return CompiledSchemas.GetBySpecSubjectPredicate("0.3.0", eType.Subject, eType.Predicate, eType.Custom)
}

func (e ServiceUpgradedEventV0_1_1) GetSubjectContent() interface{} {
	return e.Subject.Content
}

// Set subject custom fields

func (e *ServiceUpgradedEventV0_1_1) SetSubjectArtifactId(artifactId string) {
	e.Subject.Content.ArtifactId = artifactId
}

func (e *ServiceUpgradedEventV0_1_1) SetSubjectEnvironment(environment *Reference) {
	e.Subject.Content.Environment = environment
}

// New creates a new ServiceUpgradedEventV0_1_1
func NewServiceUpgradedEventV0_1_1(specVersion string) (*ServiceUpgradedEventV0_1_1, error) {
	e := &ServiceUpgradedEventV0_1_1{
		Context: Context{
			Type:    ServiceUpgradedEventTypeV0_1_1,
			Version: specVersion,
		},
		Subject: ServiceUpgradedSubjectV0_1_1{
			SubjectBase: SubjectBase{
				Type: "service",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
