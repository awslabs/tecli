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

package controller

import (
	"fmt"
	"os"

	"github.com/awslabs/tfe-cli/cobra/aid"
	"github.com/awslabs/tfe-cli/helper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var gitIgnoreArgs = []string{"list"}

// GitIgnoreCmd ....
func GitIgnoreCmd() *cobra.Command {
	man, err := helper.GetManual("gitignore")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:     man.Use,
		Short:   man.Short,
		Long:    man.Long,
		Example: man.Example,
		PreRunE: gitIgnorePreRun,
		RunE:    gitIgnoreRun,
	}

	cmd.Flags().StringP("input", "i", "", "Gitignore input. If multiple, comma-separated")

	return cmd
}

func gitIgnorePreRun(cmd *cobra.Command, args []string) error {
	logrus.Traceln("start: command gitignore pre-run")

	if len(args) == 0 {
		input, err := cmd.Flags().GetString("input")

		if err != nil {
			logrus.Errorf("unable to access flag input\n%v", err)
			return err
		}

		if input == "" {
			logrus.Errorln("no flag or argument provided")
			return fmt.Errorf("no flag or argument provided")
		}
	} else if len(args) == 1 && args[0] != "list" {
		logrus.Errorf("unknow argument passed: %v", args)
		return fmt.Errorf("unknown argument provided: %s", args[0])
	}

	logrus.Traceln("end: command gitignore pre-run")

	return nil
}

func gitIgnoreRun(cmd *cobra.Command, args []string) error {
	logrus.Traceln("start: command gitignore run")

	if len(args) > 0 && args[0] == "list" {
		list := aid.GetGitIgnoreList()
		if list == "" {
			return fmt.Errorf("unable to get gitignore list")
		}

		cmd.Println(list)
	} else {
		input, err := cmd.Flags().GetString("input")
		if err != nil {
			return err
		}

		downloaded, err := aid.DownloadGitIgnore(cmd, input)
		if err != nil {
			logrus.Errorf("unable to download gitignore\n%v", err)
			return err
		}

		if downloaded {
			cmd.Println(".gitignore created successfully")
		}
	}

	logrus.Traceln("end: command gitignore run")
	return nil
}
