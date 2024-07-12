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
	// TestCaseRunStarted event type v0.2.0
	TestCaseRunStartedEventTypeV0_2_0 CDEventType = CDEventType{
		Subject:   "testcaserun",
		Predicate: "started",
		Version:   "0.2.0",
	}
)

type TestCaseRunStartedSubjectContentV0_2_0 struct {
	Environment *Reference `json:"environment"`

	TestCase *TestCaseRunStartedSubjectContentTestCaseV0_2_0 `json:"testCase,omitempty"`

	TestSuiteRun *Reference `json:"testSuiteRun,omitempty"`

	Trigger *TestCaseRunStartedSubjectContentTriggerV0_2_0 `json:"trigger,omitempty"`
}

type TestCaseRunStartedSubjectV0_2_0 struct {
	SubjectBase
	Content TestCaseRunStartedSubjectContentV0_2_0 `json:"content"`
}

func (sc TestCaseRunStartedSubjectV0_2_0) GetSubjectType() SubjectType {
	return "testCaseRun"
}

type TestCaseRunStartedEventV0_2_0 struct {
	Context ContextV04                      `json:"context"`
	Subject TestCaseRunStartedSubjectV0_2_0 `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e TestCaseRunStartedEventV0_2_0) GetType() CDEventType {
	return TestCaseRunStartedEventTypeV0_2_0
}

func (e TestCaseRunStartedEventV0_2_0) GetVersion() string {
	return e.Context.GetVersion()
}

func (e TestCaseRunStartedEventV0_2_0) GetId() string {
	return e.Context.Id
}

func (e TestCaseRunStartedEventV0_2_0) GetSource() string {
	return e.Context.Source
}

func (e TestCaseRunStartedEventV0_2_0) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e TestCaseRunStartedEventV0_2_0) GetSubjectId() string {
	return e.Subject.Id
}

func (e TestCaseRunStartedEventV0_2_0) GetSubjectSource() string {
	return e.Subject.Source
}

func (e TestCaseRunStartedEventV0_2_0) GetSubject() Subject {
	return e.Subject
}

func (e TestCaseRunStartedEventV0_2_0) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e TestCaseRunStartedEventV0_2_0) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e TestCaseRunStartedEventV0_2_0) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e TestCaseRunStartedEventV0_2_0) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsReaderV04 implementation

func (e TestCaseRunStartedEventV0_2_0) GetChainId() string {
	return e.Context.ChainId
}

func (e TestCaseRunStartedEventV0_2_0) GetLinks() EmbeddedLinksArray {
	return e.Context.Links
}

func (e TestCaseRunStartedEventV0_2_0) GetSchemaUri() string {
	return e.Context.SchemaUri
}

// GetCustomSchema looks up the SchemaUri, if any is defined. If none is defined, it returns nil.
// If it's defined and cannot be found, it returns an error.
func (e TestCaseRunStartedEventV0_2_0) GetCustomSchema() (*jsonschema.Schema, error) {
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

func (e *TestCaseRunStartedEventV0_2_0) SetId(id string) {
	e.Context.Id = id
}

func (e *TestCaseRunStartedEventV0_2_0) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *TestCaseRunStartedEventV0_2_0) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *TestCaseRunStartedEventV0_2_0) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *TestCaseRunStartedEventV0_2_0) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *TestCaseRunStartedEventV0_2_0) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e TestCaseRunStartedEventV0_2_0) GetSchema() (string, *jsonschema.Schema, error) {
	eType := e.GetType()
	return CompiledSchemas.GetBySpecSubjectPredicate("0.4.1", eType.Subject, eType.Predicate, eType.Custom)
}

// CDEventsWriterV04 implementation

func (e *TestCaseRunStartedEventV0_2_0) SetChainId(chainId string) {
	e.Context.ChainId = chainId
}

func (e *TestCaseRunStartedEventV0_2_0) SetLinks(links EmbeddedLinksArray) {
	e.Context.Links = links
}

func (e *TestCaseRunStartedEventV0_2_0) SetSchemaUri(schema string) {
	e.Context.SchemaUri = schema
}

func (e TestCaseRunStartedEventV0_2_0) GetSubjectContent() interface{} {
	return e.Subject.Content
}

// Set subject custom fields

func (e *TestCaseRunStartedEventV0_2_0) SetSubjectEnvironment(environment *Reference) {
	e.Subject.Content.Environment = environment
}

func (e *TestCaseRunStartedEventV0_2_0) SetSubjectTestCase(testCase *TestCaseRunStartedSubjectContentTestCaseV0_2_0) {
	e.Subject.Content.TestCase = testCase
}

func (e *TestCaseRunStartedEventV0_2_0) SetSubjectTestSuiteRun(testSuiteRun *Reference) {
	e.Subject.Content.TestSuiteRun = testSuiteRun
}

func (e *TestCaseRunStartedEventV0_2_0) SetSubjectTrigger(trigger *TestCaseRunStartedSubjectContentTriggerV0_2_0) {
	e.Subject.Content.Trigger = trigger
}

// New creates a new TestCaseRunStartedEventV0_2_0
func NewTestCaseRunStartedEventV0_2_0(specVersion string) (*TestCaseRunStartedEventV0_2_0, error) {
	e := &TestCaseRunStartedEventV0_2_0{
		Context: ContextV04{
			Context{
				Type:    TestCaseRunStartedEventTypeV0_2_0,
				Version: specVersion,
			},
			ContextLinks{},
			ContextCustom{},
		},
		Subject: TestCaseRunStartedSubjectV0_2_0{
			SubjectBase: SubjectBase{
				Type: "testCaseRun",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// TestCaseRunStartedSubjectContentTestCaseV0_2_0 holds the content of a TestCase field in the content
type TestCaseRunStartedSubjectContentTestCaseV0_2_0 struct {
	Id string `json:"id"`

	Name string `json:"name,omitempty"`

	Type string `json:"type,omitempty"`

	Uri string `json:"uri,omitempty"`

	Version string `json:"version,omitempty"`
}

// TestCaseRunStartedSubjectContentTriggerV0_2_0 holds the content of a Trigger field in the content
type TestCaseRunStartedSubjectContentTriggerV0_2_0 struct {
	Type string `json:"type,omitempty"`

	Uri string `json:"uri,omitempty"`
}
