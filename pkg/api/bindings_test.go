/*
Copyright 2022 The CDEvents Authors

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
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	jsonschema "github.com/santhosh-tekuri/jsonschema/v5"
)

const testsFolder = "tests"

type testData struct {
	TestValues []map[string]string `json:"testValues"`
}

var (
	testSource               = "TestAsCloudEvent"
	testSubjectId            = "mySubject123"
	testPipeline             = "myPipeline"
	testSubjecturl           = "https://www.example.com/mySubject123"
	testPipelineOutcome      = PipelineRunOutcomeFailed
	testPipelineErrors       = "Something went wrong\nWith some more details"
	testTaskName             = "myTask"
	testTaskOutcome          = TaskRunOutcomeFailed
	testTaskRunErrors        = "Something went wrong\nWith some more details"
	testRepo                 = "TestRepo"
	testOwner                = "TestOrg"
	testUrl                  = "https://example.org/TestOrg/TestRepo"
	testViewUrl              = "https://example.org/view/TestOrg/TestRepo"
	testArtifactId           = "pkg:oci/myapp@sha256%3A0b31b1c02ff458ad9b7b81cbdf8f028bd54699fa151f221d1e8de6817db93427"
	testEnvironmentId        = "test123"
	testEnvironmentName      = "testEnv"
	testEnvironmentUrl       = "https://example.org/testEnv"
	testDataJson             = testData{TestValues: []map[string]string{{"k1": "v1"}, {"k2": "v2"}}}
	testDataJsonUnmarshalled = map[string]any{
		"testValues": []any{map[string]any{"k1": string("v1")}, map[string]any{"k2": string("v2")}},
	}
	testDataXml       = []byte("<xml>testData</xml>")
	testChangeId      = "myChange123"
	testRepoReference = Reference{
		Id:     "TestRepo/TestOrg",
		Source: "https://example.org",
	}

	pipelineRunQueuedEvent   *PipelineRunQueuedEvent
	pipelineRunStartedEvent  *PipelineRunStartedEvent
	pipelineRunFinishedEvent *PipelineRunFinishedEvent
	taskRunStartedEvent      *TaskRunStartedEvent
	taskRunFinishedEvent     *TaskRunFinishedEvent
	changeCreatedEvent       *ChangeCreatedEvent
	changeUpdatedEvent       *ChangeUpdatedEvent
	changeReviewedEvent      *ChangeReviewedEvent
	changeMergedEvent        *ChangeMergedEvent
	changeAbandonedEvent     *ChangeAbandonedEvent
	repositoryCreatedEvent   *RepositoryCreatedEvent
	repositoryModifiedEvent  *RepositoryModifiedEvent
	repositoryDeletedEvent   *RepositoryDeletedEvent
	branchCreatedEvent       *BranchCreatedEvent
	branchDeletedEvent       *BranchDeletedEvent
	testCaseQueuedEvent      *TestCaseQueuedEvent
	testCaseStartedEvent     *TestCaseStartedEvent
	testCaseFinishedEvent    *TestCaseFinishedEvent
	testSuiteStartedEvent    *TestSuiteStartedEvent
	testSuiteFinishedEvent   *TestSuiteFinishedEvent
	buildQueuedEvent         *BuildQueuedEvent
	buildStartedEvent        *BuildStartedEvent
	buildFinishedEvent       *BuildFinishedEvent
	artifactPackagedEvent    *ArtifactPackagedEvent
	artifactPublishedEvent   *ArtifactPublishedEvent
	environmentCreatedEvent  *EnvironmentCreatedEvent
	environmentModifiedEvent *EnvironmentModifiedEvent
	environmentDeletedEvent  *EnvironmentDeletedEvent
	serviceDeployedEvent     *ServiceDeployedEvent
	serviceUpgradedEvent     *ServiceUpgradedEvent
	serviceRolledBackEvent   *ServiceRolledbackEvent
	serviceRemovedEvent      *ServiceRemovedEvent
	servicePublishedEvent    *ServicePublishedEvent

	eventJsonCustomData             *ArtifactPackagedEvent
	eventNonJsonCustomData          *ArtifactPackagedEvent
	eventJsonCustomDataUnmarshalled *ArtifactPackagedEvent

	pipelineRunQueuedEventJsonTemplate = `
{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.pipelinerun.queued.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "pipelineRun",
		"content": {
			"pipelineName": "myPipeline",
			"url": "https://www.example.com/mySubject123"
		}
	}
}`

	pipelineRunStartedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.pipelinerun.started.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "pipelineRun",
		"content": {
			"pipelineName": "myPipeline",
			"url": "https://www.example.com/mySubject123"
		}
	}
}`

	pipelineRunFinishedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.pipelinerun.finished.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "pipelineRun",
		"content": {
			"pipelineName": "myPipeline",
			"url": "https://www.example.com/mySubject123",
			"outcome": "failure",
			"errors": "Something went wrong\nWith some more details"
		}
	}
}`

	taskRunStartedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.taskrun.started.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "taskRun",
		"content": {
			"taskName": "myTask",
			"url": "https://www.example.com/mySubject123",
			"pipelineRun": {
				"id": "mySubject123"
			}
		}
	}
}`

	taskRunFinishedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.taskrun.finished.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "taskRun",
		"content": {
			"taskName": "myTask",
			"url": "https://www.example.com/mySubject123",
			"pipelineRun": {
				"id": "mySubject123"
			},
			"outcome": "failure",
			"errors": "Something went wrong\nWith some more details"
		}
	}
}`

	changeCreatedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.change.created.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "change",
		"content": {
			"repository": {
				"id": "TestRepo/TestOrg",
				"source": "https://example.org"
			}
		}
	}
}`

	changeUpdatedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.change.updated.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "change",
		"content": {
			"repository": {
				"id": "TestRepo/TestOrg",
				"source": "https://example.org"
			}
		}
	}
}`

	changeReviewedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.change.reviewed.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "change",
		"content": {
			"repository": {
				"id": "TestRepo/TestOrg",
				"source": "https://example.org"
			}
		}
	}
}`

	changeMergedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.change.merged.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "change",
		"content": {
			"repository": {
				"id": "TestRepo/TestOrg",
				"source": "https://example.org"
			}
		}
	}
}`

	changeAbandonedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.change.abandoned.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "change",
		"content": {
			"repository": {
				"id": "TestRepo/TestOrg",
				"source": "https://example.org"
			}
		}
	}
}`

	repositoryCreatedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.repository.created.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "repository",
		"content": {
			"name": "TestRepo",
			"owner": "TestOrg",
			"url": "https://example.org/TestOrg/TestRepo",
			"viewUrl": "https://example.org/view/TestOrg/TestRepo"
		}
	}
}`

	repositoryModifiedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.repository.modified.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "repository",
		"content": {
			"name": "TestRepo",
			"owner": "TestOrg",
			"url": "https://example.org/TestOrg/TestRepo",
			"viewUrl": "https://example.org/view/TestOrg/TestRepo"
		}
	}
}`

	repositoryDeletedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.repository.deleted.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "repository",
		"content": {
			"name": "TestRepo",
			"owner": "TestOrg",
			"url": "https://example.org/TestOrg/TestRepo",
			"viewUrl": "https://example.org/view/TestOrg/TestRepo"
		}
	}
}`

	branchCreatedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.branch.created.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "branch",
		"content": {
			"repository": {
				"id": "TestRepo/TestOrg",
				"source": "https://example.org"
			}
		}
	}
}`

	branchDeletedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.branch.deleted.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "branch",
		"content": {
			"repository": {
				"id": "TestRepo/TestOrg",
				"source": "https://example.org"
			}
		}
	}
}`

	testCaseQueuedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.testcase.queued.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "testCase",
		"content": {}
	}
}`

	testCaseStartedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.testcase.started.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "testCase",
		"content": {}
	}
}`

	testCaseFinishedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.testcase.finished.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "testCase",
		"content": {}
	}
}`

	testSuiteStartedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.testsuite.started.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "testSuite",
		"content": {}
	}
}`

	testSuiteFinishedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.testsuite.finished.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "testSuite",
		"content": {}
	}
}`

	buildQueuedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.build.queued.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "build",
		"content": {}
	}
}`

	buildStartedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.build.started.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "build",
		"content": {}
	}
}`

	buildFinishedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.build.finished.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "build",
		"content": {
			"artifactId": "pkg:oci/myapp@sha256%%3A0b31b1c02ff458ad9b7b81cbdf8f028bd54699fa151f221d1e8de6817db93427"
		}
	}
}`

	artifactPackagedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.artifact.packaged.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "artifact",
		"content": {
			"change": {
				"id": "myChange123"
			}
		}
	}
}`

	artifactPublishedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.artifact.published.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "artifact",
		"content": {}
	}
}`

	environmentCreatedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.environment.created.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "environment",
		"content": {
			"name": "testEnv",
			"url": "https://example.org/testEnv"
		}
	}
}`

	environmentModifiedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.environment.modified.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "environment",
		"content": {
			"name": "testEnv",
			"url": "https://example.org/testEnv"
		}
	}
}`

	environmentDeletedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.environment.deleted.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "environment",
		"content": {
			"name": "testEnv"
		}
	}
}`

	serviceDeployedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.service.deployed.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "service",
		"content": {
			"environment": {
				"id": "test123"
			},
			"artifactId": "pkg:oci/myapp@sha256%%3A0b31b1c02ff458ad9b7b81cbdf8f028bd54699fa151f221d1e8de6817db93427"
		}
	}
}`

	serviceUpgradedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.service.upgraded.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "service",
		"content": {
			"environment": {
				"id": "test123"
			},
			"artifactId": "pkg:oci/myapp@sha256%%3A0b31b1c02ff458ad9b7b81cbdf8f028bd54699fa151f221d1e8de6817db93427"
		}
	}
}`

	serviceRolledBackEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.service.rolledback.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "service",
		"content": {
			"environment": {
				"id": "test123"
			},
			"artifactId": "pkg:oci/myapp@sha256%%3A0b31b1c02ff458ad9b7b81cbdf8f028bd54699fa151f221d1e8de6817db93427"
		}
	}
}`

	serviceRemovedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.service.removed.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "service",
		"content": {
			"environment": {
				"id": "test123"
			}
		}
	}
}`

	servicePublishedEventJsonTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.service.published.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "service",
		"content": {
			"environment": {
				"id": "test123"
			}
		}
	}
}`

	eventJsonCustomDataTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.artifact.packaged.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "artifact",
		"content": {
			"change": {
				"id": "myChange123"
			}
		}
	},
	"customData": {
		"testValues": [
			{"k1": "v1"},
			{"k2": "v2"}
		]
	},
	"customDataContentType": "application/json"
}`

	eventImplicitJsonCustomDataTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.artifact.packaged.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "artifact",
		"content": {
			"change": {
				"id": "myChange123"
			}
		}
	},
	"customData": {
		"testValues": [
			{"k1": "v1"},
			{"k2": "v2"}
		]
	},
	"customDataContentType": "application/json"
}`

	eventNonJsonCustomDataTemplate = `{
	"context": {
		"version": "0.1.0",
		"id": "%s",
		"source": "TestAsCloudEvent",
		"type": "dev.cdevents.artifact.packaged.0.1.0",
		"timestamp": "%s"
	},
	"subject": {
		"id": "mySubject123",
		"source": "TestAsCloudEvent",
		"type": "artifact",
		"content": {
			"change": {
				"id": "myChange123"
			}
		}
	},
	"customData": "PHhtbD50ZXN0RGF0YTwveG1sPg==",
	"customDataContentType": "application/xml"
}`

	pipelineRunQueuedEventJson      string
	pipelineRunStartedEventJson     string
	pipelineRunFinishedEventJson    string
	taskRunStartedEventJson         string
	taskRunFinishedEventJson        string
	changeCreateEventJson           string
	changeUpdatedEventJson          string
	changeReviewedEventJson         string
	changeMergedEventJson           string
	changeAbandonedEventJson        string
	repositoryCreatedEventJson      string
	repositoryModifiedEventJson     string
	repositoryDeletedEventJson      string
	branchCreatedEventJson          string
	branchDeletedEventJson          string
	testCaseQueuedEventJson         string
	testCaseStartedEventJson        string
	testCaseFinishedEventJson       string
	testSuiteStartedEventJson       string
	testSuiteFinishedEventJson      string
	buildQueuedEventJson            string
	buildStartedEventJson           string
	buildFinishedEventJson          string
	artifactPackagedEventJson       string
	artifactPublishedEventJson      string
	environmentCreatedEventJson     string
	environmentModifiedEventJson    string
	environmentDeletedEventJson     string
	serviceDeployedEventJson        string
	serviceUpgradedEventJson        string
	serviceRolledBackEventJson      string
	serviceRemovedEventJson         string
	servicePublishedEventJson       string
	eventJsonCustomDataJson         string
	eventImplicitJsonCustomDataJson string
	eventNonJsonCustomDataJson      string
)

func init() {

	// Get the time once
	t := time.Now().Round(0)
	timeNow = func() time.Time {
		return t
	}

	// Get the UUID once
	u, _ := uuid.NewRandom()
	uuidNewRandom = func() (uuid.UUID, error) {
		return u, nil
	}
}

func setContext(event CDEventWriter) {
	event.SetSource(testSource)
	event.SetSubjectId(testSubjectId)
}

func init() {
	pipelineRunQueuedEvent, _ = NewPipelineRunQueuedEvent()
	setContext(pipelineRunQueuedEvent)
	pipelineRunQueuedEvent.SetSubjectPipelineName(testPipeline)
	pipelineRunQueuedEvent.SetSubjectUrl(testSubjecturl)

	pipelineRunStartedEvent, _ = NewPipelineRunStartedEvent()
	setContext(pipelineRunStartedEvent)
	pipelineRunStartedEvent.SetSubjectPipelineName(testPipeline)
	pipelineRunStartedEvent.SetSubjectUrl(testSubjecturl)

	pipelineRunFinishedEvent, _ = NewPipelineRunFinishedEvent()
	setContext(pipelineRunFinishedEvent)
	pipelineRunFinishedEvent.SetSubjectPipelineName(testPipeline)
	pipelineRunFinishedEvent.SetSubjectUrl(testSubjecturl)
	pipelineRunFinishedEvent.SetSubjectOutcome(testPipelineOutcome)
	pipelineRunFinishedEvent.SetSubjectErrors(testPipelineErrors)

	taskRunStartedEvent, _ = NewTaskRunStartedEvent()
	setContext(taskRunStartedEvent)
	taskRunStartedEvent.SetSubjectTaskName(testTaskName)
	taskRunStartedEvent.SetSubjectUrl(testSubjecturl)
	taskRunStartedEvent.SetSubjectPipelineRun(Reference{Id: testSubjectId})

	taskRunFinishedEvent, _ = NewTaskRunFinishedEvent()
	setContext(taskRunFinishedEvent)
	taskRunFinishedEvent.SetSubjectTaskName(testTaskName)
	taskRunFinishedEvent.SetSubjectUrl(testSubjecturl)
	taskRunFinishedEvent.SetSubjectPipelineRun(Reference{Id: testSubjectId})
	taskRunFinishedEvent.SetSubjectOutcome(testTaskOutcome)
	taskRunFinishedEvent.SetSubjectErrors(testTaskRunErrors)

	changeCreatedEvent, _ = NewChangeCreatedEvent()
	setContext(changeCreatedEvent)
	changeCreatedEvent.SetSubjectRepository(testRepoReference)

	changeUpdatedEvent, _ = NewChangeUpdatedEvent()
	setContext(changeUpdatedEvent)
	changeUpdatedEvent.SetSubjectRepository(testRepoReference)

	changeReviewedEvent, _ = NewChangeReviewedEvent()
	setContext(changeReviewedEvent)
	changeReviewedEvent.SetSubjectRepository(testRepoReference)

	changeMergedEvent, _ = NewChangeMergedEvent()
	setContext(changeMergedEvent)
	changeMergedEvent.SetSubjectRepository(testRepoReference)

	changeAbandonedEvent, _ = NewChangeAbandonedEvent()
	setContext(changeAbandonedEvent)
	changeAbandonedEvent.SetSubjectRepository(testRepoReference)

	repositoryCreatedEvent, _ = NewRepositoryCreatedEvent()
	setContext(repositoryCreatedEvent)
	repositoryCreatedEvent.SetSubjectName(testRepo)
	repositoryCreatedEvent.SetSubjectOwner(testOwner)
	repositoryCreatedEvent.SetSubjectUrl(testUrl)
	repositoryCreatedEvent.SetSubjectViewUrl(testViewUrl)

	repositoryModifiedEvent, _ = NewRepositoryModifiedEvent()
	setContext(repositoryModifiedEvent)
	repositoryModifiedEvent.SetSubjectName(testRepo)
	repositoryModifiedEvent.SetSubjectOwner(testOwner)
	repositoryModifiedEvent.SetSubjectUrl(testUrl)
	repositoryModifiedEvent.SetSubjectViewUrl(testViewUrl)

	repositoryDeletedEvent, _ = NewRepositoryDeletedEvent()
	setContext(repositoryDeletedEvent)
	repositoryDeletedEvent.SetSubjectName(testRepo)
	repositoryDeletedEvent.SetSubjectOwner(testOwner)
	repositoryDeletedEvent.SetSubjectUrl(testUrl)
	repositoryDeletedEvent.SetSubjectViewUrl(testViewUrl)

	branchCreatedEvent, _ = NewBranchCreatedEvent()
	setContext(branchCreatedEvent)
	branchCreatedEvent.SetSubjectRepository(testRepoReference)

	branchDeletedEvent, _ = NewBranchDeletedEvent()
	setContext(branchDeletedEvent)
	branchDeletedEvent.SetSubjectRepository(testRepoReference)

	testCaseQueuedEvent, _ = NewTestCaseQueuedEvent()
	setContext(testCaseQueuedEvent)

	testCaseStartedEvent, _ = NewTestCaseStartedEvent()
	setContext(testCaseStartedEvent)

	testCaseFinishedEvent, _ = NewTestCaseFinishedEvent()
	setContext(testCaseFinishedEvent)

	testSuiteStartedEvent, _ = NewTestSuiteStartedEvent()
	setContext(testSuiteStartedEvent)

	testSuiteFinishedEvent, _ = NewTestSuiteFinishedEvent()
	setContext(testSuiteFinishedEvent)

	buildQueuedEvent, _ = NewBuildQueuedEvent()
	setContext(buildQueuedEvent)

	buildStartedEvent, _ = NewBuildStartedEvent()
	setContext(buildStartedEvent)

	buildFinishedEvent, _ = NewBuildFinishedEvent()
	setContext(buildFinishedEvent)
	buildFinishedEvent.SetSubjectArtifactId(testArtifactId)

	artifactPackagedEvent, _ = NewArtifactPackagedEvent()
	setContext(artifactPackagedEvent)
	artifactPackagedEvent.SetSubjectChange(Reference{Id: testChangeId})

	artifactPublishedEvent, _ = NewArtifactPublishedEvent()
	setContext(artifactPublishedEvent)

	environmentCreatedEvent, _ = NewEnvironmentCreatedEvent()
	setContext(environmentCreatedEvent)
	environmentCreatedEvent.SetSubjectName(testEnvironmentName)
	environmentCreatedEvent.SetSubjectUrl(testEnvironmentUrl)

	environmentModifiedEvent, _ = NewEnvironmentModifiedEvent()
	setContext(environmentModifiedEvent)
	environmentModifiedEvent.SetSubjectName(testEnvironmentName)
	environmentModifiedEvent.SetSubjectUrl(testEnvironmentUrl)

	environmentDeletedEvent, _ = NewEnvironmentDeletedEvent()
	setContext(environmentDeletedEvent)
	environmentDeletedEvent.SetSubjectName(testEnvironmentName)

	serviceDeployedEvent, _ = NewServiceDeployedEvent()
	setContext(serviceDeployedEvent)
	serviceDeployedEvent.SetSubjectEnvironment(Reference{Id: testEnvironmentId})
	serviceDeployedEvent.SetSubjectArtifactId(testArtifactId)

	serviceUpgradedEvent, _ = NewServiceUpgradedEvent()
	setContext(serviceUpgradedEvent)
	serviceUpgradedEvent.SetSubjectEnvironment(Reference{Id: testEnvironmentId})
	serviceUpgradedEvent.SetSubjectArtifactId(testArtifactId)

	serviceRolledBackEvent, _ = NewServiceRolledbackEvent()
	setContext(serviceRolledBackEvent)
	serviceRolledBackEvent.SetSubjectEnvironment(Reference{Id: testEnvironmentId})
	serviceRolledBackEvent.SetSubjectArtifactId(testArtifactId)

	serviceRemovedEvent, _ = NewServiceRemovedEvent()
	setContext(serviceRemovedEvent)
	serviceRemovedEvent.SetSubjectEnvironment(Reference{Id: testEnvironmentId})

	servicePublishedEvent, _ = NewServicePublishedEvent()
	setContext(servicePublishedEvent)
	servicePublishedEvent.SetSubjectEnvironment(Reference{Id: testEnvironmentId})

	eventJsonCustomData, _ = NewArtifactPackagedEvent()
	setContext(eventJsonCustomData)
	eventJsonCustomData.SetSubjectChange(Reference{Id: testChangeId})
	err := eventJsonCustomData.SetCustomData("application/json", testDataJson)
	panicOnError(err)

	eventJsonCustomDataUnmarshalled, _ = NewArtifactPackagedEvent()
	setContext(eventJsonCustomDataUnmarshalled)
	eventJsonCustomDataUnmarshalled.SetSubjectChange(Reference{Id: testChangeId})
	err = eventJsonCustomDataUnmarshalled.SetCustomData("application/json", testDataJsonUnmarshalled)
	panicOnError(err)

	eventNonJsonCustomData, _ = NewArtifactPackagedEvent()
	setContext(eventNonJsonCustomData)
	eventNonJsonCustomData.SetSubjectChange(Reference{Id: testChangeId})
	err = eventNonJsonCustomData.SetCustomData("application/xml", testDataXml)
	panicOnError(err)

	newUUID, _ := uuidNewRandom()
	newTime := timeNow().Format(time.RFC3339Nano)
	pipelineRunQueuedEventJson = fmt.Sprintf(pipelineRunQueuedEventJsonTemplate, newUUID, newTime)
	pipelineRunStartedEventJson = fmt.Sprintf(pipelineRunStartedEventJsonTemplate, newUUID, newTime)
	pipelineRunFinishedEventJson = fmt.Sprintf(pipelineRunFinishedEventJsonTemplate, newUUID, newTime)
	taskRunStartedEventJson = fmt.Sprintf(taskRunStartedEventJsonTemplate, newUUID, newTime)
	taskRunFinishedEventJson = fmt.Sprintf(taskRunFinishedEventJsonTemplate, newUUID, newTime)
	changeCreateEventJson = fmt.Sprintf(changeCreatedEventJsonTemplate, newUUID, newTime)
	changeUpdatedEventJson = fmt.Sprintf(changeUpdatedEventJsonTemplate, newUUID, newTime)
	changeReviewedEventJson = fmt.Sprintf(changeReviewedEventJsonTemplate, newUUID, newTime)
	changeMergedEventJson = fmt.Sprintf(changeMergedEventJsonTemplate, newUUID, newTime)
	changeAbandonedEventJson = fmt.Sprintf(changeAbandonedEventJsonTemplate, newUUID, newTime)
	repositoryCreatedEventJson = fmt.Sprintf(repositoryCreatedEventJsonTemplate, newUUID, newTime)
	repositoryModifiedEventJson = fmt.Sprintf(repositoryModifiedEventJsonTemplate, newUUID, newTime)
	repositoryDeletedEventJson = fmt.Sprintf(repositoryDeletedEventJsonTemplate, newUUID, newTime)
	branchCreatedEventJson = fmt.Sprintf(branchCreatedEventJsonTemplate, newUUID, newTime)
	branchDeletedEventJson = fmt.Sprintf(branchDeletedEventJsonTemplate, newUUID, newTime)
	testCaseQueuedEventJson = fmt.Sprintf(testCaseQueuedEventJsonTemplate, newUUID, newTime)
	testCaseStartedEventJson = fmt.Sprintf(testCaseStartedEventJsonTemplate, newUUID, newTime)
	testCaseFinishedEventJson = fmt.Sprintf(testCaseFinishedEventJsonTemplate, newUUID, newTime)
	testSuiteStartedEventJson = fmt.Sprintf(testSuiteStartedEventJsonTemplate, newUUID, newTime)
	testSuiteFinishedEventJson = fmt.Sprintf(testSuiteFinishedEventJsonTemplate, newUUID, newTime)
	buildQueuedEventJson = fmt.Sprintf(buildQueuedEventJsonTemplate, newUUID, newTime)
	buildStartedEventJson = fmt.Sprintf(buildStartedEventJsonTemplate, newUUID, newTime)
	buildFinishedEventJson = fmt.Sprintf(buildFinishedEventJsonTemplate, newUUID, newTime)
	artifactPackagedEventJson = fmt.Sprintf(artifactPackagedEventJsonTemplate, newUUID, newTime)
	artifactPublishedEventJson = fmt.Sprintf(artifactPublishedEventJsonTemplate, newUUID, newTime)
	environmentCreatedEventJson = fmt.Sprintf(environmentCreatedEventJsonTemplate, newUUID, newTime)
	environmentModifiedEventJson = fmt.Sprintf(environmentModifiedEventJsonTemplate, newUUID, newTime)
	environmentDeletedEventJson = fmt.Sprintf(environmentDeletedEventJsonTemplate, newUUID, newTime)
	serviceDeployedEventJson = fmt.Sprintf(serviceDeployedEventJsonTemplate, newUUID, newTime)
	serviceUpgradedEventJson = fmt.Sprintf(serviceUpgradedEventJsonTemplate, newUUID, newTime)
	serviceRolledBackEventJson = fmt.Sprintf(serviceRolledBackEventJsonTemplate, newUUID, newTime)
	serviceRemovedEventJson = fmt.Sprintf(serviceRemovedEventJsonTemplate, newUUID, newTime)
	servicePublishedEventJson = fmt.Sprintf(servicePublishedEventJsonTemplate, newUUID, newTime)
	eventJsonCustomDataJson = fmt.Sprintf(eventJsonCustomDataTemplate, newUUID, newTime)
	eventImplicitJsonCustomDataJson = fmt.Sprintf(eventImplicitJsonCustomDataTemplate, newUUID, newTime)
	eventNonJsonCustomDataJson = fmt.Sprintf(eventNonJsonCustomDataTemplate, newUUID, newTime)
}

func TestAsCloudEvent(t *testing.T) {

	tests := []struct {
		name            string
		event           CDEventReader
		payloadReceiver interface{}
	}{{
		name:            "pipelinerun queued",
		event:           pipelineRunQueuedEvent,
		payloadReceiver: &PipelineRunQueuedEvent{},
	}, {
		name:            "pipelinerun started",
		event:           pipelineRunStartedEvent,
		payloadReceiver: &PipelineRunStartedEvent{},
	}, {
		name:            "pipelinerun finished",
		event:           pipelineRunFinishedEvent,
		payloadReceiver: &PipelineRunFinishedEvent{},
	}, {
		name:            "taskrun started",
		event:           taskRunStartedEvent,
		payloadReceiver: &TaskRunStartedEvent{},
	}, {
		name:            "taskrun finished",
		event:           taskRunFinishedEvent,
		payloadReceiver: &TaskRunFinishedEvent{},
	}, {
		name:            "change created",
		event:           changeCreatedEvent,
		payloadReceiver: &ChangeCreatedEvent{},
	}, {
		name:            "change updated",
		event:           changeUpdatedEvent,
		payloadReceiver: &ChangeUpdatedEvent{},
	}, {
		name:            "change reviewed",
		event:           changeReviewedEvent,
		payloadReceiver: &ChangeReviewedEvent{},
	}, {
		name:            "change merged",
		event:           changeMergedEvent,
		payloadReceiver: &ChangeMergedEvent{},
	}, {
		name:            "change abandoned",
		event:           changeAbandonedEvent,
		payloadReceiver: &ChangeAbandonedEvent{},
	}, {
		name:            "repository created",
		event:           repositoryCreatedEvent,
		payloadReceiver: &RepositoryCreatedEvent{},
	}, {
		name:            "repository modified",
		event:           repositoryModifiedEvent,
		payloadReceiver: &RepositoryModifiedEvent{},
	}, {
		name:            "repository deleted",
		event:           repositoryDeletedEvent,
		payloadReceiver: &RepositoryDeletedEvent{},
	}, {
		name:            "branch created",
		event:           branchCreatedEvent,
		payloadReceiver: &BranchCreatedEvent{},
	}, {
		name:            "branch deleted",
		event:           branchDeletedEvent,
		payloadReceiver: &BranchDeletedEvent{},
	}, {
		name:            "testcase queued",
		event:           testCaseQueuedEvent,
		payloadReceiver: &TestCaseQueuedEvent{},
	}, {
		name:            "testcase started",
		event:           testCaseStartedEvent,
		payloadReceiver: &TestCaseStartedEvent{},
	}, {
		name:            "testcase finished",
		event:           testCaseFinishedEvent,
		payloadReceiver: &TestCaseFinishedEvent{},
	}, {
		name:            "testsuite started",
		event:           testSuiteStartedEvent,
		payloadReceiver: &TestSuiteStartedEvent{},
	}, {
		name:            "testsuite finished",
		event:           testSuiteFinishedEvent,
		payloadReceiver: &TestSuiteFinishedEvent{},
	}, {
		name:            "build queued",
		event:           buildQueuedEvent,
		payloadReceiver: &BuildQueuedEvent{},
	}, {
		name:            "build started",
		event:           buildStartedEvent,
		payloadReceiver: &BuildStartedEvent{},
	}, {
		name:            "build finished",
		event:           buildFinishedEvent,
		payloadReceiver: &BuildFinishedEvent{},
	}, {
		name:            "artifact packaged",
		event:           artifactPackagedEvent,
		payloadReceiver: &ArtifactPackagedEvent{},
	}, {
		name:            "artifact published",
		event:           artifactPublishedEvent,
		payloadReceiver: &ArtifactPublishedEvent{},
	}, {
		name:            "environment created",
		event:           environmentCreatedEvent,
		payloadReceiver: &EnvironmentCreatedEvent{},
	}, {
		name:            "environment modified",
		event:           environmentModifiedEvent,
		payloadReceiver: &EnvironmentModifiedEvent{},
	}, {
		name:            "environment deleted",
		event:           environmentDeletedEvent,
		payloadReceiver: &EnvironmentDeletedEvent{},
	}, {
		name:            "service deployed",
		event:           serviceDeployedEvent,
		payloadReceiver: &ServiceDeployedEvent{},
	}, {
		name:            "service upgraded",
		event:           serviceUpgradedEvent,
		payloadReceiver: &ServiceUpgradedEvent{},
	}, {
		name:            "service rolledback",
		event:           serviceRolledBackEvent,
		payloadReceiver: &ServiceRolledbackEvent{},
	}, {
		name:            "service removed",
		event:           serviceRemovedEvent,
		payloadReceiver: &ServiceRemovedEvent{},
	}, {
		name:            "service published",
		event:           servicePublishedEvent,
		payloadReceiver: &ServicePublishedEvent{},
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ce, err := AsCloudEvent(tc.event)
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			if d := cmp.Diff(testSubjectId, ce.Context.GetSubject()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(testSource, ce.Context.GetSource()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(tc.event.GetType().String(), ce.Context.GetType()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			err = ce.DataAs(tc.payloadReceiver)
			if err != nil {
				t.Fatalf("somehow cannot unmarshal test event %v, %v", ce, err)
			}
			if d := cmp.Diff(tc.event, tc.payloadReceiver); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
		})
	}
}

func TestAsCloudEventInvalid(t *testing.T) {
	_, err := AsCloudEvent(nil)
	if err == nil {
		t.Fatalf("expected it to fail, but it didn't")
	}
}

func TestAsJsonString(t *testing.T) {

	tests := []struct {
		name       string
		event      CDEvent
		jsonString string
	}{{
		name:       "pipelinerun queued",
		event:      pipelineRunQueuedEvent,
		jsonString: pipelineRunQueuedEventJson,
	}, {
		name:       "pipelinerun started",
		event:      pipelineRunStartedEvent,
		jsonString: pipelineRunStartedEventJson,
	}, {
		name:       "pipelinerun finished",
		event:      pipelineRunFinishedEvent,
		jsonString: pipelineRunFinishedEventJson,
	}, {
		name:       "taskrun started",
		event:      taskRunStartedEvent,
		jsonString: taskRunStartedEventJson,
	}, {
		name:       "taskrun finished",
		event:      taskRunFinishedEvent,
		jsonString: taskRunFinishedEventJson,
	}, {
		name:       "change created",
		event:      changeCreatedEvent,
		jsonString: changeCreateEventJson,
	}, {
		name:       "change updated",
		event:      changeUpdatedEvent,
		jsonString: changeUpdatedEventJson,
	}, {
		name:       "change reviewed",
		event:      changeReviewedEvent,
		jsonString: changeReviewedEventJson,
	}, {
		name:       "change merged",
		event:      changeMergedEvent,
		jsonString: changeMergedEventJson,
	}, {
		name:       "change abandoned",
		event:      changeAbandonedEvent,
		jsonString: changeAbandonedEventJson,
	}, {
		name:       "repository created",
		event:      repositoryCreatedEvent,
		jsonString: repositoryCreatedEventJson,
	}, {
		name:       "repository modified",
		event:      repositoryModifiedEvent,
		jsonString: repositoryModifiedEventJson,
	}, {
		name:       "repository deleted",
		event:      repositoryDeletedEvent,
		jsonString: repositoryDeletedEventJson,
	}, {
		name:       "branch created",
		event:      branchCreatedEvent,
		jsonString: branchCreatedEventJson,
	}, {
		name:       "branch deleted",
		event:      branchDeletedEvent,
		jsonString: branchDeletedEventJson,
	}, {
		name:       "testcase queued",
		event:      testCaseQueuedEvent,
		jsonString: testCaseQueuedEventJson,
	}, {
		name:       "testcase started",
		event:      testCaseStartedEvent,
		jsonString: testCaseStartedEventJson,
	}, {
		name:       "testcase finished",
		event:      testCaseFinishedEvent,
		jsonString: testCaseFinishedEventJson,
	}, {
		name:       "testsuite started",
		event:      testSuiteStartedEvent,
		jsonString: testSuiteStartedEventJson,
	}, {
		name:       "testsuite finished",
		event:      testSuiteFinishedEvent,
		jsonString: testSuiteFinishedEventJson,
	}, {
		name:       "build queued",
		event:      buildQueuedEvent,
		jsonString: buildQueuedEventJson,
	}, {
		name:       "build started",
		event:      buildStartedEvent,
		jsonString: buildStartedEventJson,
	}, {
		name:       "build finished",
		event:      buildFinishedEvent,
		jsonString: buildFinishedEventJson,
	}, {
		name:       "artifact packaged",
		event:      artifactPackagedEvent,
		jsonString: artifactPackagedEventJson,
	}, {
		name:       "artifact published",
		event:      artifactPublishedEvent,
		jsonString: artifactPublishedEventJson,
	}, {
		name:       "environment created",
		event:      environmentCreatedEvent,
		jsonString: environmentCreatedEventJson,
	}, {
		name:       "environment modified",
		event:      environmentModifiedEvent,
		jsonString: environmentModifiedEventJson,
	}, {
		name:       "environment deleted",
		event:      environmentDeletedEvent,
		jsonString: environmentDeletedEventJson,
	}, {
		name:       "service deployed",
		event:      serviceDeployedEvent,
		jsonString: serviceDeployedEventJson,
	}, {
		name:       "service upgraded",
		event:      serviceUpgradedEvent,
		jsonString: serviceUpgradedEventJson,
	}, {
		name:       "service rolledback",
		event:      serviceRolledBackEvent,
		jsonString: serviceRolledBackEventJson,
	}, {
		name:       "service removed",
		event:      serviceRemovedEvent,
		jsonString: serviceRemovedEventJson,
	}, {
		name:       "service published",
		event:      servicePublishedEvent,
		jsonString: servicePublishedEventJson,
	}, {
		name:       "json custom data",
		event:      eventJsonCustomData,
		jsonString: eventJsonCustomDataJson,
	}, {
		name:       "json custom data implicit",
		event:      eventJsonCustomData,
		jsonString: eventImplicitJsonCustomDataJson,
	}, {
		name:       "xml custom data",
		event:      eventNonJsonCustomData,
		jsonString: eventNonJsonCustomDataJson,
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// First validate that the test JSON compiles against the schema
			schema, url := tc.event.GetSchema()
			sch, err := jsonschema.CompileString(schema, url)
			if err != nil {
				t.Fatalf("Cannot compile jsonschema %s: %v", url, err)
			}
			var v interface{}
			if err := json.Unmarshal([]byte(tc.jsonString), &v); err != nil {
				t.Fatalf("Cannot unmarshal test json: %v", err)
			}
			err = sch.Validate(v)
			if err != nil {
				t.Fatalf("Failed to validate events %s", err)
			}
			// Then test that AsJsonString produces a matching JSON from the event
			obtainedJsonString, err := AsJsonString(tc.event)
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			expectedJsonString := &bytes.Buffer{}
			if err := json.Compact(expectedJsonString, []byte(tc.jsonString)); err != nil {
				t.Fatalf("somehow cannot compact test json %s", tc.jsonString)
			}
			if d := cmp.Diff(expectedJsonString.String(), obtainedJsonString); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
		})
	}
}

func TestInvalidEvent(t *testing.T) {

	// mandatory source missing
	eventNoSource, _ := NewCDEvent(ChangeAbandonedEventV1)
	eventNoSource.SetSubjectId(testSubjectId)

	// mandatory subject id missing
	eventNoSubjectId, _ := NewCDEvent(ChangeAbandonedEventV1)
	eventNoSubjectId.SetSource(testSource)

	// forced invalid version
	eventBadVersion, _ := NewChangeAbandonedEvent()
	eventBadVersion.Context.Version = "invalid"

	// mandatory subject url missing
	eventIncompleteSubject, _ := NewRepositoryCreatedEvent()
	eventIncompleteSubject.SetSource(testSource)
	eventIncompleteSubject.SetSubjectId(testSubjectId)
	eventIncompleteSubject.SetSubjectName(testRepo)

	// invalid source format in context
	eventInvalidSource, _ := NewChangeAbandonedEvent()
	eventInvalidSource.SetSource("\\--##@@")

	// invalid source format in reference
	eventInvalidSourceReference, _ := NewServiceDeployedEvent()
	eventInvalidSourceReference.SetSubjectEnvironment(
		Reference{Id: "1234", Source: "\\--##@@"})

	// invalid format of purl
	eventInvalidPurl, _ := NewBuildFinishedEvent()
	setContext(eventInvalidPurl)
	eventInvalidPurl.SetSubjectArtifactId("not-a-valid-purl")

	// invalid event type
	eventInvalidType := &ServicePublishedEvent{
		Context: Context{
			Type:    "not-a-valid-type",
			Version: CDEventsSpecVersion,
		},
		Subject: ServicePublishedSubject{
			SubjectBase: SubjectBase{
				Type: ServiceSubjectType,
			},
		},
	}

	tests := []struct {
		name  string
		event CDEvent
	}{{
		name:  "missing source",
		event: eventNoSource,
	}, {
		name:  "missing subject id",
		event: eventNoSubjectId,
	}, {
		name:  "invalid version",
		event: eventBadVersion,
	}, {
		name:  "missing subject url",
		event: eventIncompleteSubject,
	}, {
		name:  "invalid source in context",
		event: eventInvalidSource,
	}, {
		name:  "invalid source in reference",
		event: eventInvalidSourceReference,
	}, {
		name:  "invalid purl in build finished",
		event: eventInvalidPurl,
	}, {
		name:  "invalid event type",
		event: eventInvalidType,
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// First validate that the test JSON compiles against the schema
			err := Validate(tc.event)
			if err == nil {
				t.Fatalf("Expected validation to fail, but it succeeded instead")
			}
		})
	}
}

func TestAsJsonStringEmpty(t *testing.T) {
	obtainedJsonString, err := AsJsonString(nil)
	if err != nil {
		t.Fatalf("didn't expected it to fail, but it did: %v", err)
	}
	if d := cmp.Diff("", obtainedJsonString); d != "" {
		t.Errorf("args: diff(-want,+got):\n%s", d)
	}
}

func TestNewFromJsonString(t *testing.T) {

	tests := []struct {
		name       string
		event      CDEvent
		jsonString string
	}{{
		name:       "pipelinerun queued",
		event:      pipelineRunQueuedEvent,
		jsonString: pipelineRunQueuedEventJson,
	}, {
		name:       "pipelinerun started",
		event:      pipelineRunStartedEvent,
		jsonString: pipelineRunStartedEventJson,
	}, {
		name:       "pipelinerun finished",
		event:      pipelineRunFinishedEvent,
		jsonString: pipelineRunFinishedEventJson,
	}, {
		name:       "taskrun started",
		event:      taskRunStartedEvent,
		jsonString: taskRunStartedEventJson,
	}, {
		name:       "taskrun finished",
		event:      taskRunFinishedEvent,
		jsonString: taskRunFinishedEventJson,
	}, {
		name:       "change created",
		event:      changeCreatedEvent,
		jsonString: changeCreateEventJson,
	}, {
		name:       "change updated",
		event:      changeUpdatedEvent,
		jsonString: changeUpdatedEventJson,
	}, {
		name:       "change reviewed",
		event:      changeReviewedEvent,
		jsonString: changeReviewedEventJson,
	}, {
		name:       "change merged",
		event:      changeMergedEvent,
		jsonString: changeMergedEventJson,
	}, {
		name:       "change abandoned",
		event:      changeAbandonedEvent,
		jsonString: changeAbandonedEventJson,
	}, {
		name:       "repository created",
		event:      repositoryCreatedEvent,
		jsonString: repositoryCreatedEventJson,
	}, {
		name:       "repository modified",
		event:      repositoryModifiedEvent,
		jsonString: repositoryModifiedEventJson,
	}, {
		name:       "repository deleted",
		event:      repositoryDeletedEvent,
		jsonString: repositoryDeletedEventJson,
	}, {
		name:       "branch created",
		event:      branchCreatedEvent,
		jsonString: branchCreatedEventJson,
	}, {
		name:       "branch deleted",
		event:      branchDeletedEvent,
		jsonString: branchDeletedEventJson,
	}, {
		name:       "testcase queued",
		event:      testCaseQueuedEvent,
		jsonString: testCaseQueuedEventJson,
	}, {
		name:       "testcase started",
		event:      testCaseStartedEvent,
		jsonString: testCaseStartedEventJson,
	}, {
		name:       "testcase finished",
		event:      testCaseFinishedEvent,
		jsonString: testCaseFinishedEventJson,
	}, {
		name:       "testsuite started",
		event:      testSuiteStartedEvent,
		jsonString: testSuiteStartedEventJson,
	}, {
		name:       "testsuite finished",
		event:      testSuiteFinishedEvent,
		jsonString: testSuiteFinishedEventJson,
	}, {
		name:       "build queued",
		event:      buildQueuedEvent,
		jsonString: buildQueuedEventJson,
	}, {
		name:       "build started",
		event:      buildStartedEvent,
		jsonString: buildStartedEventJson,
	}, {
		name:       "build finished",
		event:      buildFinishedEvent,
		jsonString: buildFinishedEventJson,
	}, {
		name:       "artifact packaged",
		event:      artifactPackagedEvent,
		jsonString: artifactPackagedEventJson,
	}, {
		name:       "artifact published",
		event:      artifactPublishedEvent,
		jsonString: artifactPublishedEventJson,
	}, {
		name:       "environment created",
		event:      environmentCreatedEvent,
		jsonString: environmentCreatedEventJson,
	}, {
		name:       "environment modified",
		event:      environmentModifiedEvent,
		jsonString: environmentModifiedEventJson,
	}, {
		name:       "environment deleted",
		event:      environmentDeletedEvent,
		jsonString: environmentDeletedEventJson,
	}, {
		name:       "service deployed",
		event:      serviceDeployedEvent,
		jsonString: serviceDeployedEventJson,
	}, {
		name:       "service upgraded",
		event:      serviceUpgradedEvent,
		jsonString: serviceUpgradedEventJson,
	}, {
		name:       "service rolledback",
		event:      serviceRolledBackEvent,
		jsonString: serviceRolledBackEventJson,
	}, {
		name:       "service removed",
		event:      serviceRemovedEvent,
		jsonString: serviceRemovedEventJson,
	}, {
		name:       "service published",
		event:      servicePublishedEvent,
		jsonString: servicePublishedEventJson,
	}, {
		name:       "json custom data",
		event:      eventJsonCustomDataUnmarshalled,
		jsonString: eventJsonCustomDataJson,
	}, {
		name:       "json custom data implicit",
		event:      eventJsonCustomDataUnmarshalled,
		jsonString: eventImplicitJsonCustomDataJson,
	}, {
		name:       "xml custom data",
		event:      eventNonJsonCustomData,
		jsonString: eventNonJsonCustomDataJson,
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			obtainedEvent, err := NewFromJsonString(tc.jsonString)
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			// Check the context
			if d := cmp.Diff(tc.event.GetId(), obtainedEvent.GetId()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(tc.event.GetVersion(), obtainedEvent.GetVersion()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(tc.event.GetSource(), obtainedEvent.GetSource()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(tc.event.GetTimestamp(), obtainedEvent.GetTimestamp()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			if d := cmp.Diff(tc.event.GetType(), obtainedEvent.GetType()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			// Check the subject
			if d := cmp.Diff(tc.event.GetSubject(), obtainedEvent.GetSubject()); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
			// Check the data
			expectedData, err := tc.event.GetCustomData()
			if err != nil {
				t.Fatalf("cannot get data from test event %s", err)
			}
			obtainedData, err := obtainedEvent.GetCustomData()
			if err != nil {
				t.Fatalf("cannot get data from new event %s", err)
			}
			if d := cmp.Diff(expectedData, obtainedData); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
		})
	}
}

func TestParseType(t *testing.T) {

	tests := []struct {
		name      string
		eventType string
		want      *CDEventType
		wantError string
	}{{
		name:      "valid",
		eventType: "dev.cdevents.artifact.packaged.0.1.2-draft",
		want: &CDEventType{
			Subject:   "artifact",
			Predicate: "packaged",
			Version:   "0.1.2-draft",
		},
		wantError: "",
	}, {
		name:      "invalid root",
		eventType: "foo.bar.subject.predicate.0.1.2-draft",
		want:      nil,
		wantError: "cannot parse event type foo.bar.subject.predicate.0.1.2-draft",
	}, {
		name:      "invalid format",
		eventType: "dev.cdevents.artifact_packaged_0.1.2-draft",
		want:      nil,
		wantError: "cannot parse event type dev.cdevents.artifact_packaged_0.1.2-draft",
	}, {
		name:      "unknown subject",
		eventType: "dev.cdevents.subject.packaged.0.1.2-draft",
		want:      nil,
		wantError: "unknown event type dev.cdevents.subject.packaged",
	}, {
		name:      "unknown predicate",
		eventType: "dev.cdevents.artifact.predicate.0.1.2-draft",
		want:      nil,
		wantError: "unknown event type dev.cdevents.artifact.predicate",
	}, {
		name:      "invalid version",
		eventType: "dev.cdevents.artifact.packaged.0.1-draft",
		want:      nil,
		wantError: "invalid version format 0.1-draft",
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			obtained, err := ParseType(tc.eventType)
			if err != nil {
				if tc.wantError == "" {
					t.Fatalf("didn't expected it to fail, but it did: %v", err)
				} else {
					if d := cmp.Diff(tc.wantError, err.Error()); d != "" {
						t.Errorf("args: diff(-want,+got):\n%s", d)
					}
				}
			}

			// Check the subject
			if d := cmp.Diff(tc.want, obtained); d != "" {
				t.Errorf("args: diff(-want,+got):\n%s", d)
			}
		})
	}
}

func TestNewFromJsonBytes(t *testing.T) {

	minorVersion, _ := NewArtifactPackagedEvent()
	setContext(minorVersion)
	minorVersion.SetSubjectChange(Reference{Id: testChangeId})
	err := minorVersion.SetCustomData("application/json", testDataJsonUnmarshalled)
	panicOnError(err)
	minorType, err := ParseType(minorVersion.Context.Type)
	panicOnError(err)
	minorType.Version = "0.999.0"
	minorVersion.Context.Type = minorType.String()

	patchVersion, _ := NewArtifactPackagedEvent()
	setContext(patchVersion)
	patchVersion.SetSubjectChange(Reference{Id: testChangeId})
	err = patchVersion.SetCustomData("application/json", testDataJsonUnmarshalled)
	panicOnError(err)
	patchType, err := ParseType(patchVersion.Context.Type)
	panicOnError(err)
	patchType.Version = "0.1.999"
	patchVersion.Context.Type = patchType.String()

	tests := []struct {
		testFile    string
		description string
		wantError   string
		wantEvent   CDEvent
	}{{
		testFile:    "future_event_major_version",
		description: "A newer major version in the event is backward incompatible and cannot be parsed",
		wantError:   "sdk event version 0.1.0 not compatible with 999.0.0",
	}, {
		testFile:    "future_event_minor_version",
		description: "A newer minor version in the event is compatible and can be parsed, data is lost",
		wantEvent:   minorVersion,
	}, {
		testFile:    "future_event_patch_version",
		description: "A newer patch version in the event is compatible and can be parsed",
		wantEvent:   patchVersion,
	}, {
		testFile:    "non_unmarshable",
		description: "The event has a valid context but fails to unmarshal",
		wantError:   `invalid character '@' after object key:value pair`,
	}, {
		testFile:    "unknown_type",
		description: "The event has a valid structure but unknown type",
		wantError:   "unknown event type dev.cdevents.artifact.gazumped",
	}, {
		testFile:    "unparsable_context",
		description: "The context cannot be parsed, mandatory field is missing",
		wantError:   `invalid character '&' after object key:value pair`,
	}, {
		testFile:    "unparsable_type",
		description: "The context can be parsed, but the type is invalid",
		wantError:   "cannot parse event type dev.cdevents.artifact_packaged_0.1.0",
	}}
	for _, tc := range tests {
		t.Run(tc.testFile, func(t *testing.T) {
			eventBytes, err := os.ReadFile(testsFolder + string(os.PathSeparator) + tc.testFile + ".json")
			if err != nil {
				t.Fatalf("didn't expected it to fail, but it did: %v", err)
			}
			obtained, err := NewFromJsonBytes(eventBytes)
			if err != nil {
				if tc.wantError == "" {
					t.Fatalf("didn't expected it to fail, but it did: %v", err)
				} else {
					// Check the error is what is expected
					if d := cmp.Diff(tc.wantError, err.Error()); d != "" {
						t.Errorf("args: diff(-want,+got):\n%s", d)
					}
				}
			}
			if err == nil {
				if tc.wantError != "" {
					t.Fatalf("expected an error, but go none")
				} else {
					// Check the event is what is expected
					if d := cmp.Diff(tc.wantEvent, obtained, cmpopts.IgnoreFields(Context{}, "Id", "Timestamp")); d != "" {
						t.Errorf("args: diff(-want,+got):\n%s", d)
					}
				}
			}
		})
	}
}
