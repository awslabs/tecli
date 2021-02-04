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
	"gitlab.aws.dev/devops-aws/tecli/cobra/controller"
)

func TestRunList(t *testing.T) {
	args := []string{"run", "list", "--workspace-id", "ws-LhxL4zC6zhZAUT9i"}
	out, err := executeCommandOnly(t, controller.RunCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "")
}

func TestRunCreate(t *testing.T) {
	args := []string{"run", "create", "--workspace-id", "ws-LhxL4zC6zhZAUT9i"}
	out, err := executeCommandOnly(t, controller.RunCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "")
}

func TestRunRead(t *testing.T) {
	args := []string{"run", "read", "--organization", "terraform-cloud-pipeline"}
	out, err := executeCommandOnly(t, controller.RunCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "")
}

func TestRunDelete(t *testing.T) {
	args := []string{"run", "delete", "--organization", "terraform-cloud-pipeline"}
	out, err := executeCommandOnly(t, controller.RunCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "")
}
