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

var testcaserunqueuedschema = `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"https://cdevents.dev/0.3.0/schema/test-case-run-queued-event","properties":{"context":{"properties":{"version":{"type":"string","minLength":1},"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1},"type":{"type":"string","enum":["dev.cdevents.testcaserun.queued.0.1.0"],"default":"dev.cdevents.testcaserun.queued.0.1.0"},"timestamp":{"type":"string","format":"date-time"}},"additionalProperties":false,"type":"object","required":["version","id","source","type","timestamp"]},"subject":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string"},"type":{"type":"string","minLength":1,"enum":["testCaseRun"],"default":"testCaseRun"},"content":{"properties":{"trigger":{"type":"object","properties":{"type":{"type":"string","enum":["manual","pipeline","event","schedule","other"]},"uri":{"type":"string","format":"uri"}}},"environment":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"}},"additionalProperties":false,"type":"object","required":["id"]},"testSuiteRun":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string"}},"additionalProperties":false,"type":"object","required":["id"]},"testCase":{"type":"object","additionalProperties":false,"required":["id"],"properties":{"id":{"type":"string","minLength":1},"version":{"type":"string"},"name":{"type":"string"},"type":{"type":"string","enum":["performance","functional","unit","security","compliance","integration","e2e","other"]},"uri":{"type":"string","format":"uri"}}}},"additionalProperties":false,"type":"object","required":["environment"]}},"additionalProperties":false,"type":"object","required":["id","type","content"]},"customData":{"oneOf":[{"type":"object"},{"type":"string","contentEncoding":"base64"}]},"customDataContentType":{"type":"string"}},"additionalProperties":false,"type":"object","required":["context","subject"]}`

var (
	// TestCaseRunQueued event type v0.1.0
	TestCaseRunQueuedEventTypeV0_1_0 CDEventType = CDEventType{
		Subject:   "testcaserun",
		Predicate: "queued",
		Version:   "0.1.0",
	}
)

type TestCaseRunQueuedSubjectContent struct {
	Environment *Reference `json:"environment"`

	TestCase *TestCaseRunQueuedSubjectContentTestCase `json:"testCase,omitempty"`

	TestSuiteRun *Reference `json:"testSuiteRun,omitempty"`

	Trigger *TestCaseRunQueuedSubjectContentTrigger `json:"trigger,omitempty"`
}

type TestCaseRunQueuedSubject struct {
	SubjectBase
	Content TestCaseRunQueuedSubjectContent `json:"content"`
}

func (sc TestCaseRunQueuedSubject) GetSubjectType() SubjectType {
	return "testCaseRun"
}

type TestCaseRunQueuedEventV0_1_0 struct {
	Context Context                  `json:"context"`
	Subject TestCaseRunQueuedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e TestCaseRunQueuedEventV0_1_0) GetType() CDEventType {
	return TestCaseRunQueuedEventTypeV0_1_0
}

func (e TestCaseRunQueuedEventV0_1_0) GetVersion() string {
	return CDEventsSpecVersion
}

func (e TestCaseRunQueuedEventV0_1_0) GetId() string {
	return e.Context.Id
}

func (e TestCaseRunQueuedEventV0_1_0) GetSource() string {
	return e.Context.Source
}

func (e TestCaseRunQueuedEventV0_1_0) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e TestCaseRunQueuedEventV0_1_0) GetSubjectId() string {
	return e.Subject.Id
}

func (e TestCaseRunQueuedEventV0_1_0) GetSubjectSource() string {
	return e.Subject.Source
}

func (e TestCaseRunQueuedEventV0_1_0) GetSubject() Subject {
	return e.Subject
}

func (e TestCaseRunQueuedEventV0_1_0) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e TestCaseRunQueuedEventV0_1_0) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e TestCaseRunQueuedEventV0_1_0) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e TestCaseRunQueuedEventV0_1_0) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *TestCaseRunQueuedEventV0_1_0) SetId(id string) {
	e.Context.Id = id
}

func (e *TestCaseRunQueuedEventV0_1_0) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *TestCaseRunQueuedEventV0_1_0) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *TestCaseRunQueuedEventV0_1_0) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *TestCaseRunQueuedEventV0_1_0) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *TestCaseRunQueuedEventV0_1_0) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e TestCaseRunQueuedEventV0_1_0) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), testcaserunqueuedschema
}

// Set subject custom fields

func (e *TestCaseRunQueuedEventV0_1_0) SetSubjectEnvironment(environment *Reference) {
	e.Subject.Content.Environment = environment
}

func (e *TestCaseRunQueuedEventV0_1_0) SetSubjectTestCase(testCase *TestCaseRunQueuedSubjectContentTestCase) {
	e.Subject.Content.TestCase = testCase
}

func (e *TestCaseRunQueuedEventV0_1_0) SetSubjectTestSuiteRun(testSuiteRun *Reference) {
	e.Subject.Content.TestSuiteRun = testSuiteRun
}

func (e *TestCaseRunQueuedEventV0_1_0) SetSubjectTrigger(trigger *TestCaseRunQueuedSubjectContentTrigger) {
	e.Subject.Content.Trigger = trigger
}

// New creates a new TestCaseRunQueuedEventV0_1_0
func NewTestCaseRunQueuedEventV0_1_0(specVersion string) (*TestCaseRunQueuedEventV0_1_0, error) {
	e := &TestCaseRunQueuedEventV0_1_0{
		Context: Context{
			Type:    TestCaseRunQueuedEventTypeV0_1_0,
			Version: specVersion,
		},
		Subject: TestCaseRunQueuedSubject{
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

// TestCaseRunQueuedSubjectContentTestCase holds the content of a TestCase field in the content
type TestCaseRunQueuedSubjectContentTestCase struct {
	Id string `json:"id"`

	Name string `json:"name,omitempty"`

	Type string `json:"type,omitempty"`

	Uri string `json:"uri,omitempty"`

	Version string `json:"version,omitempty"`
}

// TestCaseRunQueuedSubjectContentTrigger holds the content of a Trigger field in the content
type TestCaseRunQueuedSubjectContentTrigger struct {
	Type string `json:"type,omitempty"`

	Uri string `json:"uri,omitempty"`
}
