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
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/aid"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/dao"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/helper"
)

var workspaceValidArgs = []string{
	"list",
	"create",
	"read",
	"update",
	"delete",
	"remove-vcs-connection",
	"lock",
	"unlock",
	"force-unlock",
	"assign-ssh-key",
	"unassign-ssh-key"}

// WorkspaceCmd command to display tecli current version
func WorkspaceCmd() *cobra.Command {
	man, err := helper.GetManual("workspace")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		Example:   man.Example,
		ValidArgs: workspaceValidArgs,
		Args:      cobra.OnlyValidArgs,
		PreRunE:   workspacePreRun,
		RunE:      workspaceRun,
	}

	usage := `Required when execution-mode is set to agent. The ID of the agent pool
	belonging to the workspace's organization. This value must not be specified
	if execution-mode is set to remote or local or if operations is set to true.`
	cmd.Flags().String("agent-pool-id", "", usage)

	usage = `Whether destroy plans can be queued on the workspace.`
	cmd.Flags().Bool("allow-destroy-plan", false, usage)

	usage = `Whether to automatically apply changes when a Terraform plan is successful.`
	cmd.Flags().Bool("auto-apply", false, usage)

	usage = `Which execution mode to use. Valid values are remote, local, and agent.
	When set to local, the workspace will be used for state storage only.
	This value must not be specified if operations is specified.
	'agent' execution mode is not available in Terraform Enterprise.`
	cmd.Flags().String("execution-mode", "", usage)

	usage = `Whether to filter runs based on the changed files in a VCS push. If
	enabled, the working directory and trigger prefixes describe a set of
	paths which must contain changes for a VCS push to trigger a run. If
	disabled, any push will trigger a run.`
	cmd.Flags().Bool("file-triggers-enabled", false, usage)

	usage = `The legacy TFE environment to use as the source of the migration, in the
	form organization/environment. Omit this unless you are migrating a legacy
	environment.`
	cmd.Flags().String("migration-environment", "", usage)

	usage = `The name of the workspace, which can only include letters, numbers, -,
	and _. This will be used as an identifier and must be unique in the
	organization.`
	cmd.Flags().String("name", "", usage)

	usage = `A new name for the workspace, which can only include letters, numbers, -,
	and _. This will be used as an identifier and must be unique in the
	organization. Warning: Changing a workspace's name changes its URL in the
	API and UI.`
	cmd.Flags().String("new-name", "", usage)

	usage = `Whether to queue all runs. Unless this is set to true, runs triggered by
	a webhook will not be queued until at least one run is manually queued.`
	cmd.Flags().Bool("queue-all-runs", false, usage)

	usage = `Whether this workspace allows speculative plans. Setting this to false
	prevents Terraform Cloud or the Terraform Enterprise instance from
	running plans on pull requests, which can improve security if the VCS
	repository is public or includes untrusted contributors.`
	cmd.Flags().Bool("speculative-enabled", false, usage)

	usage = `The version of Terraform to use for this workspace. Upon creating a
	workspace, the latest version is selected unless otherwise specified.`
	cmd.Flags().String("terraform-version", "", usage)

	usage = `List of repository-root-relative paths which list all locations to be
	tracked for changes. See FileTriggersEnabled above for more details.`
	var emptyArray []string
	cmd.Flags().StringArray("trigger-prefixes", emptyArray, usage)

	// Settings for the workspace's VCS repository. If omitted, the workspace is
	// created without a VCS repo. If included, you must specify at least the
	// oauth-token-id and identifier keys below.`

	usage = `The repository branch that Terraform will execute from. If omitted or submitted as an empty string, this defaults to the repository's default branch (e.g. master).`
	cmd.Flags().String("vcs-repo-branch", "", usage)

	usage = `A reference to your VCS repository in the format :org/:repo where :org and :repo refer to the organization and repository in your VCS provider. The format for Azure DevOps is :org/:project/_git/:repo.`
	cmd.Flags().String("vcs-repo-identifier", "", usage)

	usage = `Whether submodules should be fetched when cloning the VCS repository.`
	cmd.Flags().Bool("vcs-repo-ingress-submodules", false, usage)

	usage = `The VCS Connection (OAuth Connection + Token) to use. This ID can be obtained from the oauth-tokens endpoint.`
	cmd.Flags().String("vcs-repo-oauth-token-id", "", usage)

	usage = `A relative path that Terraform will execute within. This defaults to the
	root of your repository and is typically set to a subdirectory matching the
	environment when multiple environments exist within the same repository.`
	cmd.Flags().String("working-directory", "", usage)

	usage = `The SSH key ID to assign to a workspace. Must be created on the organization.`
	cmd.Flags().String("ssh-key-id", "", usage)

	return cmd
}

func workspacePreRun(cmd *cobra.Command, args []string) error {
	if err := helper.ValidateCmdArgs(cmd, args, "workspace"); err != nil {
		return err
	}

	fArg := args[0]
	switch fArg {
	case "create":
		if err := helper.ValidateCmdArgAndFlag(cmd, args, "workspace", fArg, "organization"); err != nil {
			return err
		}

		if err := helper.ValidateCmdArgAndFlag(cmd, args, "workspace", fArg, "name"); err != nil {
			return err
		}
	}

	return nil
}

func workspaceRun(cmd *cobra.Command, args []string) error {

	token := dao.GetTeamToken(profile)
	client := aid.GetTFEClient(token)

	var workspace *tfe.Workspace
	var err error

	fArg := args[0]
	switch fArg {
	case "list":
		list, err := workspaceList(client)
		if err == nil {
			if list.TotalCount > 0 {
				for _, item := range list.Items {
					fmt.Printf("%v\n", aid.ToJSON(item))
				}
			} else {
				return fmt.Errorf("no workspace was found")
			}
		}
	case "create":
		options := aid.GetWorkspaceCreateOptions(cmd)
		workspace, err = workspaceCreate(client, options)

		if err == nil && workspace.ID != "" {
			fmt.Println(aid.ToJSON(workspace))
		}
	case "read":
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		workspace, err := workspaceRead(client, name)
		if err == nil {
			fmt.Println(aid.ToJSON(workspace))
		} else {
			return fmt.Errorf("workspace %s not found\n%v", name, err)
		}
	case "update":
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		options := aid.GetWorkspaceUpdateOptions(cmd)
		workspace, err = workspaceUpdate(client, name, options)
		if err == nil && workspace.ID != "" {
			fmt.Println(aid.ToJSON(workspace))
		} else {
			return fmt.Errorf("unable to update workspace\n%v", err)
		}
	case "delete":
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		err = workspaceDelete(client, name)
		if err == nil {
			fmt.Printf("workspace %s deleted successfully\n", name)
		} else {
			return fmt.Errorf("unable to delete workspace %s\n%v", name, err)
		}
	case "remove-vcs-connection":
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		workspace, err = workspaceRemoveVCSConnection(client, name)
		if err == nil {
			fmt.Println(aid.ToJSON(workspace))
		} else {
			return fmt.Errorf("unable to remove vcs connection\n%v", err)
		}
	case "lock":
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		w, err := workspaceRead(client, name)
		if err != nil {
			return err
		}

		workspace, err := workspaceLock(client, w.ID)
		if err != nil {
			return err
		}

		if workspace.Locked {
			fmt.Printf("workspace %s locked successfully\n", name)
		}

	case "unlock":
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		w, err := workspaceRead(client, name)
		if err != nil {
			return err
		}

		workspace, err := workspaceUnlock(client, w.ID)
		if err != nil {
			return err
		}

		if !workspace.Locked {
			fmt.Printf("workspace %s unlocked successfully\n", name)
		}
	case "force-unlock":
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		w, err := workspaceRead(client, name)
		if err != nil {
			return err
		}

		workspace, err := workspaceForceUnlock(client, w.ID)
		if err != nil {
			return err
		}

		if !workspace.Locked {
			fmt.Printf("workspace %s unlocked successfully\n", name)
		}
	case "assign-ssh-key":
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		w, err := workspaceRead(client, name)
		if err != nil {
			return err
		}

		// TODO: need to fetch the SSH keys via the client.SSHKeys interface
		options := aid.GetWorkspaceAssignSSHKeyOptions(cmd)
		workspace, err := workspaceAssignSSHKey(client, w.ID, options)
		if err != nil {
			return err
		}

		if workspace.ID != "" && workspace.SSHKey.ID != "" {
			fmt.Println(aid.ToJSON(workspace))
			fmt.Println("SSH key assigned  successfully")
		}

	case "unassign-ssh-key":
		fmt.Println("unassign-ssh-key")
	default:
		return fmt.Errorf("unknown argument provided")
	}

	if err != nil {
		logrus.Fatalf("unable to %s workspace\n%v\n", fArg, err)
	}

	return nil
}

func workspaceList(client *tfe.Client) (*tfe.WorkspaceList, error) {
	return client.Workspaces.List(context.Background(), organization, tfe.WorkspaceListOptions{})
}

// Create is used to create a new workspace.
func workspaceCreate(client *tfe.Client, options tfe.WorkspaceCreateOptions) (*tfe.Workspace, error) {
	return client.Workspaces.Create(context.Background(), organization, options)
}

// Read a workspace by its name.
func workspaceRead(client *tfe.Client, workspace string) (*tfe.Workspace, error) {
	return client.Workspaces.Read(context.Background(), organization, workspace)
}

// Update settings of an existing workspace.
func workspaceUpdate(client *tfe.Client, workspace string, options tfe.WorkspaceUpdateOptions) (*tfe.Workspace, error) {
	return client.Workspaces.Update(context.Background(), organization, workspace, options)
}

// // Delete a workspace by its name.
func workspaceDelete(client *tfe.Client, workspace string) error {
	return client.Workspaces.Delete(context.Background(), organization, workspace)
}

// RemoveVCSConnection from a workspace.
func workspaceRemoveVCSConnection(client *tfe.Client, workspace string) (*tfe.Workspace, error) {
	return client.Workspaces.RemoveVCSConnection(context.Background(), organization, workspace)
}

// Lock a workspace by its ID.
func workspaceLock(client *tfe.Client, workspaceID string) (*tfe.Workspace, error) {
	return client.Workspaces.Lock(context.Background(), workspaceID, tfe.WorkspaceLockOptions{})
}

// Unlock a workspace by its ID.
func workspaceUnlock(client *tfe.Client, workspaceID string) (*tfe.Workspace, error) {
	return client.Workspaces.Unlock(context.Background(), workspaceID)
}

// ForceUnlock a workspace by its ID.
func workspaceForceUnlock(client *tfe.Client, workspaceID string) (*tfe.Workspace, error) {
	return client.Workspaces.ForceUnlock(context.Background(), workspaceID)
}

// AssignSSHKey to a workspace.
func workspaceAssignSSHKey(client *tfe.Client, workspaceID string, options tfe.WorkspaceAssignSSHKeyOptions) (*tfe.Workspace, error) {
	return client.Workspaces.AssignSSHKey(context.Background(), workspaceID, options)
}

// UnassignSSHKey from a workspace.
func workspaceUnassignSSHKey(client *tfe.Client, workspaceID string) (*tfe.Workspace, error) {
	return client.Workspaces.UnassignSSHKey(context.Background(), workspaceID)
}
