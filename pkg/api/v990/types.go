// Code generated by tools/generator. DO NOT EDIT.

//go:build testonly

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

package v990

import (
	"fmt"

	"github.com/cdevents/sdk-go/pkg/api"
)

var CDEventsTypes = []api.CDEvent{
	&FooSubjectBarPredicateEvent{},
}

var CDEventsByUnversionedTypes map[string]api.CDEvent

func init() {
	// Set up CDEventsByUnversionedTypes for convenience
	CDEventsByUnversionedTypes = make(map[string]api.CDEvent)
	for _, event := range CDEventsTypes {
		CDEventsByUnversionedTypes[event.GetType().UnversionedString()] = event
	}
}

// NewCDEvent produces a CDEvent by type
// This function can be used by users but it's meant mainly for testing purposes
func NewCDEvent(eventType, specVersion string) (api.CDEvent, error) {
	switch eventType {
	case api.FooSubjectBarPredicateEventTypeV1_2_3.String():
		return NewFooSubjectBarPredicateEvent()
	default:
		return nil, fmt.Errorf("event %v not supported", eventType)
	}
}