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
	"context"
	"fmt"
	"os"

	"github.com/awslabs/tecli/cobra/aid"
	"github.com/awslabs/tecli/cobra/dao"
	"github.com/awslabs/tecli/helper"
	"github.com/hashicorp/go-tfe"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var variableValidArgs = []string{
	"list",
	"create",
	"read",
	"update",
	"delete",
	"delete-all",
}

// VariableCmd command to display tecli current version
func VariableCmd() *cobra.Command {
	man, err := helper.GetManual("variable", variableValidArgs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		Example:   man.Example,
		ValidArgs: variableValidArgs,
		Args:      cobra.OnlyValidArgs,
		PreRunE:   variablePreRun,
		RunE:      variableRun,
	}

	aid.SetVariableFlags(cmd)

	return cmd
}

func variablePreRun(cmd *cobra.Command, args []string) error {
	logrus.Tracef("start: variablePreRun")

	if err := helper.ValidateCmdArgsV2(cmd, args); err != nil {
		return fmt.Errorf("unexpected error\n%v", err)
	}

	switch args[0] {
	case "list", "create", "delete-all":
		if err := helper.ValidateCmdFlagString(cmd, "workspace-id"); err != nil {
			return err
		}

	case "read", "update", "delete":
		if err := helper.ValidateCmdFlagString(cmd, "id"); err != nil {
			return err
		}

		if err := helper.ValidateCmdFlagString(cmd, "workspace-id"); err != nil {
			return err
		}

	default:
		return fmt.Errorf("unknown argument: %s", args[0])
	}

	return nil
}

func variableRun(cmd *cobra.Command, args []string) error {
	logrus.Tracef("start: variableRun")

	token := dao.GetTeamToken(profile)
	client := aid.GetTFEClient(token)

	fArg := args[0]
	switch fArg {
	case "list":
		workspaceID := helper.GetCmdFlagString(cmd, "workspace-id")

		list, err := variableList(client, workspaceID, tfe.VariableListOptions{})
		if err == nil && len(list.Items) > 0 {
			aid.PrintVariableList(list)
		} else {
			return fmt.Errorf("no variable was found")
		}

	case "create":
		workspaceID := helper.GetCmdFlagString(cmd, "workspace-id")
		options := aid.GetVariableCreateOptions(cmd)

		variable, err := variableCreate(client, workspaceID, options)
		if err == nil && variable.ID != "" {
			fmt.Println(aid.ToJSON(variable))
		} else {
			return fmt.Errorf("unable to create variable\n%v", err)
		}

	case "read":
		workspaceID := helper.GetCmdFlagString(cmd, "workspace-id")
		id := helper.GetCmdFlagString(cmd, "id")

		variable, err := variableRead(client, workspaceID, id)
		if err == nil {
			fmt.Println(aid.ToJSON(variable))
		} else {
			return fmt.Errorf("variable %s not found\n%v", id, err)
		}

	case "update":
		workspaceID := helper.GetCmdFlagString(cmd, "workspace-id")
		id := helper.GetCmdFlagString(cmd, "id")
		options := aid.GetVariableUpdateOptions(cmd)

		variable, err := variableUpdate(client, workspaceID, id, options)
		if err == nil && variable.ID != "" {
			fmt.Println(aid.ToJSON(variable))
		} else {
			return fmt.Errorf("unable to update variable\n%v", err)
		}

	case "delete":
		workspaceID := helper.GetCmdFlagString(cmd, "workspace-id")
		id := helper.GetCmdFlagString(cmd, "id")

		err := variableDelete(client, workspaceID, id)
		if err == nil {
			fmt.Printf("variable %s deleted successfully\n", id)
		} else {
			return fmt.Errorf("unable to delete variable %s\n%v", id, err)
		}

	case "delete-all":
		workspaceID := helper.GetCmdFlagString(cmd, "workspace-id")
		list, err := variableList(client, workspaceID, tfe.VariableListOptions{})
		if err != nil {
			return fmt.Errorf("no variable was found\n%v", err)
		}

		for _, v := range list.Items {
			fmt.Printf("attempting to delete variable %s (%s)\n", v.Key, v.ID)
			err := variableDelete(client, workspaceID, v.ID)
			if err != nil {
				return fmt.Errorf("unable to delete variable %s (%s)\n%v", v.Key, v.ID, err)
			}

			fmt.Printf("variable %s (%s) deleted successfully\n", v.Key, v.ID)
		}
	default:
		return fmt.Errorf("unknown argument provided")
	}

	return nil
}

func variableList(client *tfe.Client, workspaceID string, options tfe.VariableListOptions) (*tfe.VariableList, error) {
	return client.Variables.List(context.Background(), workspaceID, options)
}

func variableCreate(client *tfe.Client, workspaceID string, options tfe.VariableCreateOptions) (*tfe.Variable, error) {
	return client.Variables.Create(context.Background(), workspaceID, options)
}

func variableRead(client *tfe.Client, workspaceID string, variableID string) (*tfe.Variable, error) {
	return client.Variables.Read(context.Background(), workspaceID, variableID)
}

func variableUpdate(client *tfe.Client, workspaceID string, variableID string, options tfe.VariableUpdateOptions) (*tfe.Variable, error) {
	return client.Variables.Update(context.Background(), workspaceID, variableID, options)
}

func variableDelete(client *tfe.Client, workspaceID string, variableID string) error {
	return client.Variables.Delete(context.Background(), workspaceID, variableID)
}
