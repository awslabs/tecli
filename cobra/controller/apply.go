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

// Package controller acts on both model and view. It controls the data flow into model object and updates the view whenever data changes. It keeps view and model separate.
package controller

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/hashicorp/go-tfe"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.aws.dev/devops-aws/tecli/cobra/aid"
	"gitlab.aws.dev/devops-aws/tecli/cobra/dao"
	"gitlab.aws.dev/devops-aws/tecli/helper"
)

var applyValidArgs = []string{"read", "logs"}

// ApplyCmd command to display tecli current version
func ApplyCmd() *cobra.Command {
	man, err := helper.GetManual("apply", applyValidArgs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		Example:   man.Example,
		ValidArgs: applyValidArgs,
		Args:      cobra.OnlyValidArgs,
		PreRunE:   applyPreRun,
		RunE:      applyRun,
	}

	aid.SetApplyFlags(cmd)

	return cmd
}

func applyPreRun(cmd *cobra.Command, args []string) error {
	if err := helper.ValidateCmdArgs(cmd, args, "plan"); err != nil {
		return err
	}

	fArg := args[0]
	switch fArg {
	case "read", "logs":
		if err := helper.ValidateCmdArgAndFlag(cmd, args, "plan", fArg, "id"); err != nil {
			return err
		}
	}

	return nil
}

func applyRun(cmd *cobra.Command, args []string) error {

	token := dao.GetTeamToken(profile)
	client := aid.GetTFEClient(token)

	fArg := args[0]
	switch fArg {
	case "read":
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return fmt.Errorf("unable to get flag id\n%v", err)
		}

		apply, err := applyRead(client, id)
		if err == nil {
			fmt.Println(aid.ToJSON(apply))
		} else {
			return fmt.Errorf("apply %s not found\n%v", id, err)
		}
	case "logs":
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return fmt.Errorf("unable to get flag id\n%v", err)
		}

		logs, err := applyLogs(client, id)
		if err != nil {
			logrus.Fatalf("unable to read apply logs\n%v", err)
		}
		fmt.Println(StreamToString(logs))
	}

	return nil
}

// Read an apply by its ID.
func applyRead(client *tfe.Client, applyID string) (*tfe.Apply, error) {
	return client.Applies.Read(context.Background(), applyID)
}

// Logs retrieves the logs of an apply.
func applyLogs(client *tfe.Client, applyID string) (io.Reader, error) {
	return client.Applies.Logs(context.Background(), applyID)
}
