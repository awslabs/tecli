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

func TestWorkspaceCmd(t *testing.T) {
	tests := map[string]struct {
		args []string
		out  string
		err  string
	}{
		// argument
		"empty":     {args: []string{"workspace"}, out: "", err: "this command requires one argument"},
		"empty arg": {args: []string{"workspace", ""}, out: "", err: "invalid argument"},
		"wrong arg": {args: []string{"workspace", "foo"}, out: "", err: "invalid argument"},

		// flags
		"wrong flag": {args: []string{"workspace", "--foo"}, out: "", err: "unknown flag"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			out, err := executeCommand(t, controller.WorkspaceCmd(), tc.args)
			assert.Contains(t, out, tc.out)
			assert.Contains(t, err.Error(), tc.err)
		})
	}
}

func TestWorkspaceCreateWithNoArgAndNoFlags(t *testing.T) {
	args := []string{"workspace", "create", "--organization", "terraform-cloud-pipeline", "--name", "my-workspace"}
	out, err := executeCommandOnly(t, controller.WorkspaceCmd(), args)
	assert.NotNil(t, err)
	assert.Equal(t, "this command requires one argument", err.Error())
	assert.Contains(t, out, "")
}
