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

package aid

import (
	tfe "github.com/hashicorp/go-tfe"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// GetRunCreateOptions TODO ..
func GetRunCreateOptions(cmd *cobra.Command) tfe.RunCreateOptions {
	var options tfe.RunCreateOptions

	// Specifies if this plan is a destroy plan, which will destroy all
	// provisioned resources.
	isDestroy, err := cmd.Flags().GetBool("is-destroy")
	if err != nil {
		logrus.Fatalf("unable to get flag is-destroy\n%v", err)
	} else {
		options.IsDestroy = &isDestroy
	}

	// Specifies the message to be associated with this run.
	message := cmd.Flags().GetString("message")

	// Specifies the configuration version to use for this run. If the
	// configuration version object is omitted, the run will be created using the
	// workspace's latest configuration version.
	configurationVersionId := cmd.Flags().GetString("configuration-version-id")

	configurationVersionAutoQueueRuns, err := cmd.Flags().GetBool("configuration-version-auto-queue-runs")
	if err != nil {
		logrus.Fatalf("unable to get flag configuration-version-auto-queue-runs\n%v", err)
	} else {
		options.ConfigurationVersion.AutoQueueRuns = configurationVersionAutoQueueRuns
	}

	configurationVersionError := cmd.Flags().GetString("configuration-version-error")
	configurationVersionErrorMessage := cmd.Flags().GetString("configuration-version-error-message")
	configurationVersionSource := cmd.Flags().GetString("configuration-version-source")

	configurationVersionSpeculative, err := cmd.Flags().GetBool("configuration-version-speculative")
	if err != nil {
		logrus.Fatalf("unable to get flag configuration-version-speculative\n%v", err)
	} else {
		options.ConfigurationVersion.Speculative = configurationVersionSpeculative
	}

	configurationVersionStatus := cmd.Flags().GetSTring("configuration-version-status")
	configurationVersionUploadUrl := cmd.Flags().GetString("configuration-version-upload-url")

	// Specifies the workspace where the run will be executed.
	workspaceId := cmd.Flags().GetString("workspace-id")

	workspaceActionsIsDestroyable, err := cmd.Flags().GetBool("workspace-actions-is-destroyable")
	if err != nil {
		logrus.Fatalf("unable to get flag workspace-actions-is-destroyable\n%v", err)
	} else {
		options.Workspace.Actions.IsDestroyable = workspaceActionsIsDestroyable
	}

	workspaceAgentPoolId := cmd.Flags().GetString("workspace-agent-pool-id")

	workspaceAllowDestroyPlan, err := cmd.Flags().GetBool("workspace-allow-destroy-plan")
	if err != nil {
		logrus.Fatalf("unable to get flag workspace-allow-destroy-plan\n%v", err)
	} else {
		options.Workspace.AllowDestroyPlan = workspaceAllowDestroyPlan
	}

	workspaceAutoApply, err := cmd.Flags().GetBool("workspace-auto-apply")
	if err != nil {
		logrus.Fatalf("unable to get flag workspace-auto-apply\n%v", err)
	} else {
		options.Workspace.AutoApply = workspaceAutoApply
	}

	workspaceCanQueueDestroyPlan, err := cmd.Flags().GetBool("workspace-can-queue-destroy-plan")
	if err != nil {
		logrus.Fatalf("unable to get flag workspace-can-queue-destroy-plan\n%v", err)
	} else {
		options.Workspace.CanQueueDestroyPlan = workspaceCanQueueDestroyPlan
	}

	workspaceEnvironment := cmd.Flags().GetString("workspace-environment")
	workspaceExecutionMode := cmd.Flags().GetString("workspace-execution-mode")

	workspaceFileTriggersEnabled, err := cmd.Flags().GetBool("workspace-file-triggers-enabled")
	if err != nil {
		logrus.Fatalf("unable to get flag workspace-file-triggers-enabled\n%v", err)
	} else {
		options.Workspace.FileTriggersEnabled = workspaceFileTriggersEnabled
	}

	workspaceLocked, err := cmd.Flags().GetBool("workspace-locked")
	if err != nil {
		logrus.Fatalf("unable to get flag workspace-locked\n%v", err)
	} else {
		options.Workspace.Locked = workspaceLocked
	}

	workspaceMigrationEnvironment := cmd.Flags().GetString("workspace-migration-environment")
	workspaceName := cmd.Flags().GetString("workspace-name")

	workspacePermissionCanDestroy, err := cmd.Flags().GetBool("workspace-permission-can-destroy")
	if err != nil {
		logrus.Fatalf("unable to get flag workspace-permission-can-destroy\n%v", err)
	} else {
		options.Workspace.Permissions.CanDestroy = workspacePermissionCanDestroy
	}

	workspacePermissionCanForceUnlock, err := cmd.Flags().GetBool("workspace-permission-can-force-unlock")
	if err != nil {
		logrus.Fatalf("unable to get flag workspace-permission-can-force-unlock\n%v", err)
	} else {
		options.Workspace.Permissions.CanForceUnlock = workspacePermissionCanForceUnlock
	}

	workspacePermissionCanLock, err := cmd.Flags().GetBool("workspace-permission-can-lock")
	if err != nil {
		logrus.Fatalf("unable to get flag workspace-permission-can-lock\n%v", err)
	} else {
		options.Workspace.Permissions.CanLock = workspacePermissionCanLock
	}

	workspacePermissionCanQueueApply, err := cmd.Flags().GetBool("workspace-permission-can-queue-apply")
	if err != nil {
		logrus.Fatalf("unable to get flag workspace-permission-can-queue-apply\n%v", err)
	} else {
		options.Workspace.Permissions.CanQueueApply = workspacePermissionCanQueueApply
	}

	workspacePermissionCanQueueDestroy, err := cmd.Flags().GetBool("workspace-permission-can-queue-destroy")
	if err != nil {
		logrus.Fatalf("unable to get flag workspace-permission-can-queue-destroy\n%v", err)
	} else {
		options.Workspace.Permissions.CanQueueDestroy = workspacePermissionCanQueueDestroy
	}

	workspacePermissionCanQueueRun, err := cmd.Flags().GetBool("workspace-permission-can-queue-run")
	if err != nil {
		logrus.Fatalf("unable to get flag workspace-permission-can-queue-run\n%v", err)
	} else {
		options.Workspace.Permissions.CanQueueRun = workspacePermissionCanQueueRun
	}

	workspacePermissionCanReadSettings, err := cmd.Flags().GetBool("workspace-permission-can-read-settings")
	if err != nil {
		logrus.Fatalf("unable to get flag workspace-permission-can-read-settings\n%v", err)
	} else {
		options.Workspace.Permissions.CanReadSettings = workspacePermissionCanReadSettings
	}

	workspacePermissionCanUnlock, err := cmd.Flags().GetBool("workspace-permission-can-unlock")
	if err != nil {
		logrus.Fatalf("unable to get flag workspace-permission-can-unlock\n%v", err)
	} else {
		options.Workspace.Permissions.CanUnlock = workspacePermissionCanUnlock
	}

	workspacePermissionCanUpdate, err := cmd.Flags().GetBool("workspace-permission-can-update")
	if err != nil {
		logrus.Fatalf("unable to get flag workspace-permission-can-update\n%v", err)
	} else {
		options.Workspace.Permissions.CanUpdate = workspacePermissionCanUpdate
	}

	workspacePermissionCanUpdateVariable, err := cmd.Flags().GetBool("workspace-permission-can-update-variable")
	if err != nil {
		logrus.Fatalf("unable to get flag workspace-permission-can-update-variable\n%v", err)
	} else {
		options.Workspace.Permissions.CanUpdateVariable = workspacePermissionCanUpdateVariable
	}

	workspaceQueueAllRuns, err := cmd.Flags().GetBool("workspace-queue-all-runs")
	if err != nil {
		logrus.Fatalf("unable to get flag workspace-queue-all-runs\n%v", err)
	} else {
		options.Workspace.QueueAllRuns = workspaceQueueAllRuns
	}

	workspaceSpeculativeEnabled, err := cmd.Flags().GetBool("workspace-speculative-enabled")
	if err != nil {
		logrus.Fatalf("unable to get flag workspace-speculative-enabled\n%v", err)
	} else {
		options.Workspace.SpeculativeEnabled = workspaceSpeculativeEnabled
	}

	workspaceTerraformVersion := cmd.Flags().GetString("workspace-terraform-version")
	workspaceTriggerPrefixes := cmd.Flags().GetStringArray("workspace-trigger-prefixes")

	workspaceVcsRepoBranch := cmd.Flags().GetString("workspace-vcs-repo-branch")
	workspaceVcsRepoDisplayIdentifier := cmd.Flags().GetString("workspace-vcs-repo-display-identifier")
	workspaceVcsRepoIdentifier := cmd.Flags().GetString("workspace-vcs-repo-identifier")
	workspaceVcsRepoIngressSubmodules, err := cmd.Flags().GetBool("workspace-vcs-repo-ingress-submodules")
	if err != nil {
		logrus.Fatalf("unable to get flag workspace-vcs-repo-ingress-submodules\n%v", err)
	} else {
		options.Workspace.VcsRepoIngressSubmodules = &workspaceVcsRepoIngressSubmodules
	}

	workspaceVcsRepoOAuthTokenId := cmd.Flags().GetString("workspace-vcs-repo-o-auth-token-id")

	workspaceWorkingDirectory := cmd.Flags().GetString("workspace-working-directory")

	// If non-empty, requests that Terraform should create a plan including
	// actions only for the given objects (specified using resource address
	// syntax) and the objects they depend on.
	//
	// This capability is provided for exceptional circumstances only, such as
	// recovering from mistakes or working around existing Terraform
	// limitations. Terraform will generally mention the -target command line
	// option in its error messages describing situations where setting this
	// argument may be appropriate. This argument should not be used as part
	// of routine workflow and Terraform will emit warnings reminding about
	// this whenever this property is set.
	targetAddrs := cmd.Flags().GetStringArray("target-addrs")

	return options

}

// // GetRunUpdateOptions TODO ...
// func GetRunUpdateOptions(cmd *cobra.Command) tfe.RunUpdateOptions {
// 	var options tfe.RunUpdateOptions

// 	return options

// }
