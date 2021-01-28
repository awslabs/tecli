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

	"github.com/hashicorp/go-tfe"
	"github.com/spf13/cobra"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/aid"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/dao"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/helper"
)

var runValidArgs = []string{
	"list",
	"create",
	"read",
	"read-with-options",
	"apply",
	"cancel",
	"force-cancel",
	"discard"}

// RunCmd command to display tecli current version
func RunCmd() *cobra.Command {
	man, err := helper.GetManualV2("run", runValidArgs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		Example:   man.Example,
		ValidArgs: runValidArgs,
		Args:      cobra.OnlyValidArgs,
		PreRunE:   runPreRun,
		RunE:      runRun,
	}

	aid.SetRunFlags(cmd)

	return cmd
}

func runPreRun(cmd *cobra.Command, args []string) error {
	if err := helper.ValidateCmdArgs(cmd, args, "run"); err != nil {
		return err
	}

	fArg := args[0]
	switch fArg {
	case "list", "create":
		if err := helper.ValidateCmdArgAndFlag(cmd, args, "run", fArg, "workspace-id"); err != nil {
			return err
		}

	}

	return nil
}

func runRun(cmd *cobra.Command, args []string) error {

	token := dao.GetTeamToken(profile)
	client := aid.GetTFEClient(token)

	fArg := args[0]
	switch fArg {
	case "list":
		workspaceID, err := cmd.Flags().GetString("workspace-id")
		if err != nil {
			return fmt.Errorf("unable to get flag workspace-id\n%v", err)
		}

		list, err := runList(client, workspaceID, tfe.RunListOptions{})
		if err == nil {
			if len(list.Items) > 0 {
				for _, item := range list.Items {
					fmt.Printf("%v,\n", aid.ToJSON(item))
				}
			} else {
				return fmt.Errorf("no run was found")
			}
		}
	case "create":
		options := aid.GetRunCreateOptions(cmd)

		workspaceID, err := cmd.Flags().GetString("workspace-id")
		if err != nil {
			return fmt.Errorf("unable to get flag workspace-id\n%v", err)
		}

		if workspaceID != "" {
			workspace, err := workspaceReadByID(client, workspaceID)
			if err != nil {
				return fmt.Errorf("unable to find workspace %s\n%v", workspaceID, err)
			}
			options.Workspace = workspace
		}

		cvID, err := cmd.Flags().GetString("configuration-version-id")
		if err != nil {
			return fmt.Errorf("unable to get flag configuration-version-id\n%v", err)
		}

		if cvID != "" {
			cv, err := configurationVersionRead(client, cvID)
			if err != nil {
				return fmt.Errorf("unable to find configuration version %s\n%v", cvID, err)
			}

			if cv.ID != "" {
				options.ConfigurationVersion = cv
			}
		}

		run, err := runCreate(client, options)

		if err == nil && run.ID != "" {
			fmt.Println(aid.ToJSON(run))
		} else {
			return fmt.Errorf("unable to create run\n%v", err)
		}

	case "read":
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return fmt.Errorf("unable to get flag id\n%v", err)
		}

		run, err := runRead(client, id)
		if err == nil {
			fmt.Println(aid.ToJSON(run))
		} else {
			return fmt.Errorf("run %s not found\n%v", id, err)
		}
	case "read-with-options":
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return fmt.Errorf("unable to get flag id\n%v", err)
		}

		options := aid.GetRunReadOptions(cmd)
		run, err := runReadWithOptions(client, id, &options)
		if err == nil {
			fmt.Println(aid.ToJSON(run))
		} else {
			return fmt.Errorf("run %s not found\n%v", id, err)
		}
	case "apply":
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return fmt.Errorf("unable to get flag id\n%v", err)
		}

		options := aid.GetRunApplyOptions(cmd)
		err = runApply(client, id, options)
		if err != nil {
			return fmt.Errorf("unable to apply run\n%v", err)
		}

		fmt.Println("run applied successfully")
	case "cancel":
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return fmt.Errorf("unable to get flag id\n%v", err)
		}

		options := aid.GetRunCancelOptions(cmd)
		err = runCancel(client, id, options)
		if err != nil {
			return fmt.Errorf("unable to cancel run\n%v", err)
		}

		fmt.Println("run cancelled successfully")

	case "force-cancel":
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return fmt.Errorf("unable to get flag id\n%v", err)
		}

		options := aid.GetRunForceCancelOptions(cmd)
		err = runForceCancel(client, id, options)
		if err != nil {
			return fmt.Errorf("unable to force ancel run\n%v", err)
		}

		fmt.Println("run cancelled successfully")
	case "discard":
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return fmt.Errorf("unable to get flag id\n%v", err)
		}

		options := aid.GetRunDiscardOptions(cmd)
		err = runDiscard(client, id, options)
		if err != nil {
			return fmt.Errorf("unable to discard run\n%v", err)
		}

	}

	return nil
}

// List all the runs of the given workspace.
func runList(client *tfe.Client, workspaceID string, options tfe.RunListOptions) (*tfe.RunList, error) {
	return client.Runs.List(context.Background(), workspaceID, options)
}

// Create a new run with the given options.
func runCreate(client *tfe.Client, options tfe.RunCreateOptions) (*tfe.Run, error) {
	return client.Runs.Create(context.Background(), options)
}

// Read a run by its ID.
func runRead(client *tfe.Client, runID string) (*tfe.Run, error) {
	return client.Runs.Read(context.Background(), runID)
}

// ReadWithOptions reads a run by its ID using the options supplied
func runReadWithOptions(client *tfe.Client, runID string, options *tfe.RunReadOptions) (*tfe.Run, error) {
	return client.Runs.ReadWithOptions(context.Background(), runID, options)
}

// Apply a run by its ID.
func runApply(client *tfe.Client, runID string, options tfe.RunApplyOptions) error {
	return client.Runs.Apply(context.Background(), runID, options)
}

// Cancel a run by its ID.
func runCancel(client *tfe.Client, runID string, options tfe.RunCancelOptions) error {
	return client.Runs.Cancel(context.Background(), runID, options)
}

// Force-cancel a run by its ID.
func runForceCancel(client *tfe.Client, runID string, options tfe.RunForceCancelOptions) error {
	return client.Runs.ForceCancel(context.Background(), runID, options)
}

func runDiscard(client *tfe.Client, runID string, options tfe.RunDiscardOptions) error {
	return client.Runs.Discard(context.Background(), runID, options)
}
