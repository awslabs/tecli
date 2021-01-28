/*
Copyright © 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/controller"
)

func TestPlanList(t *testing.T) {
	args := []string{"plan", "list", "--organization", "terraform-cloud-pipeline"}
	out, err := executeCommandOnly(t, controller.PlanCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "")
}

func TestPlanCreate(t *testing.T) {
	args := []string{"plan", "create", "--organization", "terraform-cloud-pipeline"}
	out, err := executeCommandOnly(t, controller.PlanCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "")
}

func TestPlanRead(t *testing.T) {
	args := []string{"plan", "read", "--organization", "terraform-cloud-pipeline"}
	out, err := executeCommandOnly(t, controller.PlanCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "")
}

func TestPlanDelete(t *testing.T) {
	args := []string{"plan", "delete", "--organization", "terraform-cloud-pipeline"}
	out, err := executeCommandOnly(t, controller.PlanCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "")
}
