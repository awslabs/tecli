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

// Package tests contains command-level tests for tecli's Cobra commands.
package tests

import (
	"bytes"
	"os"
	"testing"

	"github.com/awslabs/tecli/cobra/controller"
	"github.com/spf13/cobra"
)

/* COBRA */

// executeCommand runs cmd inside a fresh temporary working directory and
// returns its combined stdout/stderr output.
func executeCommand(t *testing.T, cmd *cobra.Command, args []string) (stdout string, err error) {
	wd := t.TempDir()
	// execute the command within the new working directory
	if err := os.Chdir(wd); err != nil {
		t.Fatalf("unable to change to temp working directory %q: %v", wd, err)
	}
	_, stdout, err = executeCommandC(cmd, args)

	return stdout, err
}

// executeCommandOnly only executes the given command without changing the
// current directory, useful for combined tests.
func executeCommandOnly(t *testing.T, cmd *cobra.Command, args []string) (stdout string, err error) {
	_, stdout, err = executeCommandC(cmd, args)
	return stdout, err
}

func executeCommandC(cmd *cobra.Command, args []string) (command *cobra.Command, stdout string, err error) {
	buf := new(bytes.Buffer)

	rootCmd := controller.RootCmd()
	rootCmd.AddCommand(cmd)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs(args)

	command, err = rootCmd.ExecuteC()
	stdout = buf.String()
	return command, stdout, err
}
