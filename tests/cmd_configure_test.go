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

// Package test contains unit and integration tests
package tests

import (
	"os"
	"testing"

	"github.com/awslabs/tecli/cobra/aid"
	"github.com/awslabs/tecli/cobra/controller"
	"github.com/stretchr/testify/assert"
)

func TestConfigureCmdFlags(t *testing.T) {
	tests := map[string]struct {
		args []string
		out  string
		err  string
	}{
		// argument
		"empty":     {args: []string{"configure"}, out: "", err: "this command requires one argument"},
		"empty arg": {args: []string{"configure", ""}, out: "", err: "invalid argument"},
		"wrong arg": {args: []string{"configure", "foo"}, out: "", err: "invalid argument"},

		// flags
		"wrong flag": {args: []string{"configure", "--foo"}, out: "", err: "unknown flag"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			out, err := executeCommand(t, controller.ConfigureCmd(), tc.args)
			assert.Contains(t, out, tc.out)
			assert.Contains(t, err.Error(), tc.err)
		})
	}
}

func TestConfigureCreateWithNoArgAndNoFlags(t *testing.T) {
	// need to remove global tecli configuration directory
	err := os.RemoveAll(aid.GetAppInfo().AppDir)
	assert.Nil(t, err)

	// need to find a way to run/debug test in interactive mode
	args := []string{"configure", "create", "--mode", "non-interactive"}
	out, err := executeCommandOnly(t, controller.ConfigureCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "created successfully")
}
