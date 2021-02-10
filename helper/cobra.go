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

// Package helper contains collection of common functions
package helper

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ValidateCmdArgs basic validation of a command's arguments, return error if fails
func ValidateCmdArgs(cmd *cobra.Command, args []string, cmdName string) error {
	logrus.Tracef("start: validate command %s arguments", cmdName)

	if len(args) == 0 {
		logrus.Error("no arguments passed")
		return fmt.Errorf("this command requires one argument")
	}

	if len(args) > 1 {
		logrus.Errorf("more than one argument passed: %v\n", args)
		return fmt.Errorf("this command accepts only one argument at a time")
	}

	if !ContainsString(cmd.ValidArgs, args[0]) {
		logrus.Errorf("unknow argument passed: %v\n", args)
		logrus.Errorf("command %s only accepts the following arguments: %v\n", cmdName, args)
		return fmt.Errorf("unknown argument provided: %s", args[0])
	}

	logrus.Tracef("end: validate command %s arguments", cmdName)
	return nil
}

// ValidateCmdArgAndFlag basic validation for the given arg on args, and if flag is empty, return error if fails
func ValidateCmdArgAndFlag(cmd *cobra.Command, args []string, cmdName string, arg string, flag string) error {
	logrus.Tracef("start: validate command %s --%s", cmdName, flag)
	if args[0] == arg {
		pName, err := cmd.Flags().GetString(flag)
		if err != nil {
			logrus.Errorf("unable to access --%s: %s", flag, err.Error())
			return err
		}

		if pName == "" {
			logrus.Errorf("empty value passed to --%s", flag)
			return fmt.Errorf("--%s must be defined", flag)
		}
	} else {
		logrus.Errorf("unknow argument passed: %v\n", args)
		return fmt.Errorf("unknown argument provided: %s", args[0])
	}
	logrus.Tracef("end: validate command %s --%s", cmdName, flag)
	return nil
}
