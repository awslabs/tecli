package aid

import (
	tfe "github.com/hashicorp/go-tfe"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// SetWorkspaceFlagsV1 TODO ...
func SetWorkspaceFlagsV1(cmd *cobra.Command) {

	// List
	usage := `A search string (partial workspace name) used to filter the results.`
	cmd.Flags().String("search", "", usage)

	usage = `A list of relations to include. See available resources https://www.terraform.io/docs/cloud/api/workspaces.html#available-related-resources`
	cmd.Flags().String("include", "", usage)

	// Create, Update

	usage = `The workspace ID`
	cmd.Flags().String("id", "", usage)

	usage = `Required when execution-mode is set to agent. The ID of the agent pool belonging to the workspace's organization. This value must not be specified if execution-mode is set to remote or local or if operations is set to true.`
	cmd.Flags().String("agent-pool-id", "", usage)

	usage = `Whether destroy plans can be queued on the workspace.`
	cmd.Flags().Bool("allow-destroy-plan", false, usage)

	usage = `Whether to automatically apply changes when a Terraform plan is successful.`
	cmd.Flags().Bool("auto-apply", false, usage)

	usage = `Which execution mode to use. Valid values are remote, local, and agent. When set to local, the workspace will be used for state storage only. This value must not be specified if operations is specified. 'agent' execution mode is not available in Terraform Enterprise.`
	cmd.Flags().String("execution-mode", "", usage)

	usage = `Whether to filter runs based on the changed files in a VCS push. If enabled, the working directory and trigger prefixes describe a set of paths which must contain changes for a VCS push to trigger a run. If disabled, any push will trigger a run.`
	cmd.Flags().Bool("file-triggers-enabled", false, usage)

	usage = `The legacy TFE environment to use as the source of the migration, in the form organization/environment. Omit this unless you are migrating a legacy environment.`
	cmd.Flags().String("migration-environment", "", usage)

	usage = `The name of the workspace, which can only include letters, numbers, -, and _. This will be used as an identifier and must be unique in the organization.`
	cmd.Flags().String("name", "", usage)

	usage = `A new name for the workspace, which can only include letters, numbers, -, and _. This will be used as an identifier and must be unique in the organization. Warning: Changing a workspace's name changes its URL in the API and UI.`
	cmd.Flags().String("new-name", "", usage)

	usage = `Whether to queue all runs. Unless this is set to true, runs triggered by
	a webhook will not be queued until at least one run is manually queued.`
	cmd.Flags().Bool("queue-all-runs", false, usage)

	usage = `Whether this workspace allows speculative plans. Setting this to false prevents Terraform Cloud or the Terraform Enterprise instance from running plans on pull requests, which can improve security if the VCS repository is public or includes untrusted contributors.`
	cmd.Flags().Bool("speculative-enabled", false, usage)

	usage = `The version of Terraform to use for this workspace. Upon creating a workspace, the latest version is selected unless otherwise specified.`
	cmd.Flags().String("terraform-version", "", usage)

	usage = `List of repository-root-relative paths which list all locations to be tracked for changes. See FileTriggersEnabled above for more details.`
	var emptyArray []string
	cmd.Flags().StringArray("trigger-prefixes", emptyArray, usage)

	// Create/Update/RemoveVCSConnectionOptions
	SetVCSRepoFlags(cmd)

	usage = `A relative path that Terraform will execute within. This defaults to the root of your repository and is typically set to a subdirectory matching the environment when multiple environments exist within the same repository.`
	cmd.Flags().String("working-directory", "", usage)

	// Lock
	usage = `Specifies the reason for locking the workspace.`
	cmd.Flags().String("reason", "", usage)

	// AssignSSHKey / UnassignSSHKey
	usage = `The SSH key ID to assign to a workspace. Must be created on the organization.`
	cmd.Flags().String("ssh-key-id", "", usage)
}

// SetVCSRepoFlags TODO ..
func SetVCSRepoFlags(cmd *cobra.Command) {
	// Settings for the workspace's VCS repository. If omitted, the workspace is
	// created without a VCS repo. If included, you must specify at least the
	// oauth-token-id and identifier keys below.`

	usage := `The repository branch that Terraform will execute from. If omitted or submitted as an empty string, this defaults to the repository's default branch (e.g. master).`
	cmd.Flags().String("vcs-repo-branch", "", usage)

	usage = `A reference to your VCS repository in the format :org/:repo where :org and :repo refer to the organization and repository in your VCS provider. The format for Azure DevOps is :org/:project/_git/:repo.`
	cmd.Flags().String("vcs-repo-identifier", "", usage)

	usage = `Whether submodules should be fetched when cloning the VCS repository.`
	cmd.Flags().Bool("vcs-repo-ingress-submodules", false, usage)

	usage = `The VCS Connection (OAuth Connection + Token) to use. This ID can be obtained from the oauth-tokens endpoint.`
	cmd.Flags().String("vcs-repo-oauth-token-id", "", usage)
}

// GetWorkspaceCreateOptions TODO ..
func GetWorkspaceCreateOptions(cmd *cobra.Command) tfe.WorkspaceCreateOptions {
	var options tfe.WorkspaceCreateOptions

	agentPoolID, err := cmd.Flags().GetString("agent-pool-id")
	if err != nil {
		logrus.Fatalf("unable to get flag agent-pool-id\n%v\n", err)
	}
	if agentPoolID != "" {
		options.AgentPoolID = &agentPoolID
	}

	allowDestroyPlan, err := cmd.Flags().GetBool("allow-destroy-plan")
	if err != nil {
		logrus.Fatalf("unable to get flag allow-destroy-plan\n%v\n", err)
	}

	options.AllowDestroyPlan = &allowDestroyPlan

	autoApply, err := cmd.Flags().GetBool("auto-apply")
	if err != nil {
		logrus.Fatalf("unable to get flag auto-apply\n%v\n", err)
	}

	options.AutoApply = &autoApply

	executionMode, err := cmd.Flags().GetString("execution-mode")
	if err != nil {
		logrus.Fatalf("unable to get flag execution-mode\n%v\n", err)
	}
	if executionMode != "" {
		options.ExecutionMode = &executionMode
	}

	fileTriggersEnabled, err := cmd.Flags().GetBool("file-triggers-enabled")
	if err != nil {
		logrus.Fatalf("unable to get flag file-triggers-enabled\n%v\n", err)
	}

	options.FileTriggersEnabled = &fileTriggersEnabled

	migrationEnvironment, err := cmd.Flags().GetString("migration-environment")
	if err != nil {
		logrus.Fatalf("unable to get flag migration-environment\n%v\n", err)
	}
	if migrationEnvironment != "" {
		options.MigrationEnvironment = &migrationEnvironment
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		logrus.Fatalf("unable to get flag name\n%v\n", err)
	}
	if name != "" {
		options.Name = &name
	}

	queueAllRuns, err := cmd.Flags().GetBool("queue-all-runs")
	if err != nil {
		logrus.Fatalf("unable to get flag queue-all-runs\n%v\n", err)
	}

	options.QueueAllRuns = &queueAllRuns

	speculativeEnabled, err := cmd.Flags().GetBool("speculative-enabled")
	if err != nil {
		logrus.Fatalf("unable to get flag speculative-enabled\n%v\n", err)
	}
	if speculativeEnabled {
		options.SpeculativeEnabled = &speculativeEnabled
	}

	terraformVersion, err := cmd.Flags().GetString("terraform-version")
	if err != nil {
		logrus.Fatalf("unable to get flag terraform-version\n%v\n", err)
	}
	if terraformVersion != "" {
		options.TerraformVersion = &terraformVersion
	}

	triggerPrefixes, err := cmd.Flags().GetStringArray("trigger-prefixes")
	if err != nil {
		logrus.Fatalf("unable to get flag trigger-prefixes\n%v\n", err)
	}
	if len(triggerPrefixes) > 0 {
		options.TriggerPrefixes = triggerPrefixes
	}

	repoOptions := GetVCSRepoFlags(cmd)
	if repoOptions != (tfe.VCSRepoOptions{}) {
		options.VCSRepo = &repoOptions
	}

	workingDirectory, err := cmd.Flags().GetString("working-directory")
	if err != nil {
		logrus.Fatalf("unable to get flag working-directory\n%v\n", err)
	}
	if workingDirectory != "" {
		options.WorkingDirectory = &workingDirectory
	}

	return options

}

// GetVCSRepoFlags TODO ...
func GetVCSRepoFlags(cmd *cobra.Command) tfe.VCSRepoOptions {
	var options tfe.VCSRepoOptions

	vcsRepoBranch, err := cmd.Flags().GetString("vcs-repo-branch")
	if err != nil {
		logrus.Fatalf("unable to get flag vcsRepoBranch\n%v\n", err)
	}
	if vcsRepoBranch != "" {
		options.Branch = &vcsRepoBranch
	}

	vcsRepoIdentifier, err := cmd.Flags().GetString("vcs-repo-identifier")
	if err != nil {
		logrus.Fatalf("unable to get flag vcsRepoIdentifier\n%v\n", err)
	}
	if vcsRepoIdentifier != "" {
		options.Identifier = &vcsRepoIdentifier
	}

	vcsRepoIngressSubmodules, err := cmd.Flags().GetBool("vcs-repo-ingress-submodules")
	if err != nil {
		logrus.Fatalf("unable to get flag vcsRepoIngressSubmodules\n%v\n", err)
	}

	options.IngressSubmodules = &vcsRepoIngressSubmodules

	vcsRepoOauthTokenID, err := cmd.Flags().GetString("vcs-repo-oauth-token-id")
	if err != nil {
		logrus.Fatalf("unable to get flag vcsRepoOauthTokenId\n%v\n", err)
	}
	if vcsRepoOauthTokenID != "" {
		options.OAuthTokenID = &vcsRepoOauthTokenID
	}

	return options
}

// GetWorkspaceUpdateOptions TODO ...
func GetWorkspaceUpdateOptions(cmd *cobra.Command) tfe.WorkspaceUpdateOptions {
	var options tfe.WorkspaceUpdateOptions

	// Required when execution-mode is set to agent. The ID of the agent pool
	// belonging to the workspace's organization. This value must not be specified
	// if execution-mode is set to remote or local or if operations is set to true.
	agentPoolID, err := cmd.Flags().GetString("agent-pool-id")
	if err != nil {
		logrus.Fatalf("unable to get flag agent-pool-id\n%v\n", err)
	}
	if agentPoolID != "" {
		options.AgentPoolID = &agentPoolID
	}

	// Whether destroy plans can be queued on the workspace.
	allowDestroyPlan, err := cmd.Flags().GetBool("allow-destroy-plan")
	if err != nil {
		logrus.Fatalf("unable to get flag allow-destroy-plan\n%v\n", err)
	}

	options.AllowDestroyPlan = &allowDestroyPlan

	// Whether to automatically apply changes when a Terraform plan is successful.
	autoApply, err := cmd.Flags().GetBool("auto-apply")
	if err != nil {
		logrus.Fatalf("unable to get flag auto-apply\n%v\n", err)
	}

	options.AutoApply = &autoApply

	// A new name for the workspace, which can only include letters, numbers, -,
	// and _. This will be used as an identifier and must be unique in the
	// organization. Warning: Changing a workspace's name changes its URL in the
	// API and UI.
	newName, err := cmd.Flags().GetString("new-name")
	if err != nil {
		logrus.Fatalf("unable to get flag new-name\n%v\n", err)
	}

	if newName != "" {
		options.Name = &newName
	}

	// Which execution mode to use. Valid values are remote, local, and agent.
	// When set to local, the workspace will be used for state storage only.
	// This value must not be specified if operations is specified.
	// 'agent' execution mode is not available in Terraform Enterprise.
	executionMode, err := cmd.Flags().GetString("execution-mode")
	if err != nil {
		logrus.Fatalf("unable to get flag execution-mode\n%v\n", err)
	}

	if executionMode != "" {
		options.ExecutionMode = &executionMode
	}

	// Whether to filter runs based on the changed files in a VCS push. If
	// enabled, the working directory and trigger prefixes describe a set of
	// paths which must contain changes for a VCS push to trigger a run. If
	// disabled, any push will trigger a run.
	fileTriggersEnabled, err := cmd.Flags().GetBool("file-triggers-enabled")
	if err != nil {
		logrus.Fatalf("unable to get flag file-triggers-enabled\n%v\n", err)
	}

	options.FileTriggersEnabled = &fileTriggersEnabled

	// Whether to queue all runs. Unless this is set to true, runs triggered by
	// a webhook will not be queued until at least one run is manually queued.
	queueAllRuns, err := cmd.Flags().GetBool("queue-all-runs")
	if err != nil {
		logrus.Fatalf("unable to get flag queue-all-runs\n%v\n", err)
	}

	options.QueueAllRuns = &queueAllRuns

	// Whether this workspace allows speculative plans. Setting this to false
	// prevents Terraform Cloud or the Terraform Enterprise instance from
	// running plans on pull requests, which can improve security if the VCS
	// repository is public or includes untrusted contributors.
	speculativeEnabled, err := cmd.Flags().GetBool("speculative-enabled")
	if err != nil {
		logrus.Fatalf("unable to get flag speculative-enabled\n%v\n", err)
	}

	options.SpeculativeEnabled = &speculativeEnabled

	// The version of Terraform to use for this workspace.
	terraformVersion, err := cmd.Flags().GetString("terraform-version")
	if err != nil {
		logrus.Fatalf("unable to get flag terraform-version\n%v\n", err)
	}

	if terraformVersion != "" {
		options.TerraformVersion = &terraformVersion
	}

	// List of repository-root-relative paths which list all locations to be
	// tracked for changes. See FileTriggersEnabled above for more details.
	triggerPrefixes, err := cmd.Flags().GetStringArray("trigger-prefixes")
	if err != nil {
		logrus.Fatalf("unable to get flag trigger-prefixes\n%v\n", err)
	}

	if len(triggerPrefixes) > 0 {
		options.TriggerPrefixes = triggerPrefixes
	}

	// To delete a workspace's existing VCS repo, specify null instead of an
	// object. To modify a workspace's existing VCS repo, include whichever of
	// the keys below you wish to modify. To add a new VCS repo to a workspace
	// that didn't previously have one, include at least the oauth-token-id and
	// identifier keys.

	repoOptions := GetVCSRepoFlags(cmd)
	if repoOptions != (tfe.VCSRepoOptions{}) {
		options.VCSRepo = &repoOptions
	}

	// A relative path that Terraform will execute within. This defaults to the
	// root of your repository and is typically set to a subdirectory matching
	// the environment when multiple environments exist within the same
	// repository.
	workingDirectory, err := cmd.Flags().GetString("working-directory")
	if err != nil {
		logrus.Fatalf("unable to get flag working-directory\n%v\n", err)
	}

	if workingDirectory != "" {
		options.WorkingDirectory = &workingDirectory
	}

	return options
}

// GetWorkspaceAssignSSHKeyOptions TODO ...
func GetWorkspaceAssignSSHKeyOptions(cmd *cobra.Command) tfe.WorkspaceAssignSSHKeyOptions {
	var options tfe.WorkspaceAssignSSHKeyOptions

	sshKeyID, err := cmd.Flags().GetString("ssh-key-id")
	if err != nil {
		logrus.Fatalf("unable to get flag ssh-key-id\n%v\n", err)
	}

	if sshKeyID != "" {
		options.SSHKeyID = &sshKeyID
	}

	return options
}

// // GetWorspaceFlags TODO ...
// func GetWorspaceFlags(cmd *cobra.Command) tfe.Workspace {

// 	var workspace tfe.Workspace

// 	// Specifies the workspace where the run will be executed.
// 	workspaceID, err := cmd.Flags().GetString("workspace-id")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-id\n%v", err)
// 	}
// 	if workspaceID != "" {
// 		workspace.ID = workspaceID
// 	}

// 	workspaceActionsIsDestroyable, err := cmd.Flags().GetBool("workspace-actions-is-destroyable")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-actions-is-destroyable\n%v", err)
// 	} else {
// 		workspace.Actions.IsDestroyable = workspaceActionsIsDestroyable
// 	}

// 	workspaceAgentPoolID, err := cmd.Flags().GetString("workspace-agent-pool-id")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-agent-pool-id\n%v", err)
// 	}
// 	if workspaceAgentPoolID != "" {
// 		workspace.AgentPool.ID = workspaceAgentPoolID
// 	}

// 	workspaceAllowDestroyPlan, err := cmd.Flags().GetBool("workspace-allow-destroy-plan")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-allow-destroy-plan\n%v", err)
// 	} else {
// 		workspace.AllowDestroyPlan = workspaceAllowDestroyPlan
// 	}

// 	workspaceAutoApply, err := cmd.Flags().GetBool("workspace-auto-apply")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-auto-apply\n%v", err)
// 	} else {
// 		workspace.AutoApply = workspaceAutoApply
// 	}

// 	workspaceCanQueueDestroyPlan, err := cmd.Flags().GetBool("workspace-can-queue-destroy-plan")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-can-queue-destroy-plan\n%v", err)
// 	} else {
// 		workspace.CanQueueDestroyPlan = workspaceCanQueueDestroyPlan
// 	}

// 	workspaceEnvironment, err := cmd.Flags().GetString("workspace-environment")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-environment\n%v", err)
// 	}
// 	if workspaceEnvironment != "" {
// 		workspace.Environment = workspaceEnvironment
// 	}

// 	workspaceExecutionMode, err := cmd.Flags().GetString("workspace-execution-mode")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-execution-mode\n%v", err)
// 	}
// 	if workspaceExecutionMode != "" {
// 		workspace.ExecutionMode = workspaceExecutionMode
// 	}

// 	workspaceFileTriggersEnabled, err := cmd.Flags().GetBool("workspace-file-triggers-enabled")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-file-triggers-enabled\n%v", err)
// 	} else {
// 		workspace.FileTriggersEnabled = workspaceFileTriggersEnabled
// 	}

// 	workspaceLocked, err := cmd.Flags().GetBool("workspace-locked")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-locked\n%v", err)
// 	} else {
// 		workspace.Locked = workspaceLocked
// 	}

// 	workspaceMigrationEnvironment, err := cmd.Flags().GetString("workspace-migration-environment")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-migration-environment\n%v", err)
// 	}
// 	if workspaceMigrationEnvironment != "" {
// 		workspace.MigrationEnvironment = workspaceMigrationEnvironment
// 	}

// 	workspaceName, err := cmd.Flags().GetString("workspace-name")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-name\n%v", err)
// 	}
// 	if workspaceName != "" {
// 		workspace.Name = workspaceName
// 	}

// 	workspacePermissionCanDestroy, err := cmd.Flags().GetBool("workspace-permission-can-destroy")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-permission-can-destroy\n%v", err)
// 	} else {
// 		workspace.Permissions.CanDestroy = workspacePermissionCanDestroy
// 	}

// 	workspacePermissionCanForceUnlock, err := cmd.Flags().GetBool("workspace-permission-can-force-unlock")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-permission-can-force-unlock\n%v", err)
// 	} else {
// 		workspace.Permissions.CanForceUnlock = workspacePermissionCanForceUnlock
// 	}

// 	workspacePermissionCanLock, err := cmd.Flags().GetBool("workspace-permission-can-lock")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-permission-can-lock\n%v", err)
// 	} else {
// 		workspace.Permissions.CanLock = workspacePermissionCanLock
// 	}

// 	workspacePermissionCanQueueApply, err := cmd.Flags().GetBool("workspace-permission-can-queue-apply")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-permission-can-queue-apply\n%v", err)
// 	} else {
// 		workspace.Permissions.CanQueueApply = workspacePermissionCanQueueApply
// 	}

// 	workspacePermissionCanQueueDestroy, err := cmd.Flags().GetBool("workspace-permission-can-queue-destroy")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-permission-can-queue-destroy\n%v", err)
// 	} else {
// 		workspace.Permissions.CanQueueDestroy = workspacePermissionCanQueueDestroy
// 	}

// 	workspacePermissionCanQueueRun, err := cmd.Flags().GetBool("workspace-permission-can-queue-run")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-permission-can-queue-run\n%v", err)
// 	} else {
// 		workspace.Permissions.CanQueueRun = workspacePermissionCanQueueRun
// 	}

// 	workspacePermissionCanReadSettings, err := cmd.Flags().GetBool("workspace-permission-can-read-settings")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-permission-can-read-settings\n%v", err)
// 	} else {
// 		workspace.Permissions.CanReadSettings = workspacePermissionCanReadSettings
// 	}

// 	workspacePermissionCanUnlock, err := cmd.Flags().GetBool("workspace-permission-can-unlock")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-permission-can-unlock\n%v", err)
// 	} else {
// 		workspace.Permissions.CanUnlock = workspacePermissionCanUnlock
// 	}

// 	workspacePermissionCanUpdate, err := cmd.Flags().GetBool("workspace-permission-can-update")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-permission-can-update\n%v", err)
// 	} else {
// 		workspace.Permissions.CanUpdate = workspacePermissionCanUpdate
// 	}

// 	workspacePermissionCanUpdateVariable, err := cmd.Flags().GetBool("workspace-permission-can-update-variable")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-permission-can-update-variable\n%v", err)
// 	} else {
// 		workspace.Permissions.CanUpdateVariable = workspacePermissionCanUpdateVariable
// 	}

// 	workspaceQueueAllRuns, err := cmd.Flags().GetBool("workspace-queue-all-runs")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-queue-all-runs\n%v", err)
// 	} else {
// 		workspace.QueueAllRuns = workspaceQueueAllRuns
// 	}

// 	workspaceSpeculativeEnabled, err := cmd.Flags().GetBool("workspace-speculative-enabled")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-speculative-enabled\n%v", err)
// 	} else {
// 		workspace.SpeculativeEnabled = workspaceSpeculativeEnabled
// 	}

// 	workspaceTerraformVersion, err := cmd.Flags().GetString("workspace-terraform-version")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-terraform-version\n%v", err)
// 	}
// 	if workspaceTerraformVersion != "" {
// 		workspace.TerraformVersion = workspaceTerraformVersion
// 	}

// 	workspaceTriggerPrefixes, err := cmd.Flags().GetStringArray("workspace-trigger-prefixes")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-trigger-prefixes\n%v", err)
// 	}
// 	if len(workspaceTriggerPrefixes) > 0 {
// 		workspace.TriggerPrefixes = workspaceTriggerPrefixes
// 	}

// 	workspaceVcsRepoBranch, err := cmd.Flags().GetString("workspace-vcs-repo-branch")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-vcs-repo-branch\n%v", err)
// 	}
// 	if workspaceVcsRepoBranch != "" {
// 		workspace.VCSRepo.Branch = workspaceVcsRepoBranch
// 	}

// 	workspaceVcsRepoDisplayIdentifier, err := cmd.Flags().GetString("workspace-vcs-repo-display-identifier")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-vcs-repo-display-identifier\n%v", err)
// 	}
// 	if workspaceVcsRepoDisplayIdentifier != "" {
// 		workspace.VCSRepo.DisplayIdentifier = workspaceVcsRepoDisplayIdentifier
// 	}

// 	workspaceVcsRepoIdentifier, err := cmd.Flags().GetString("workspace-vcs-repo-identifier")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-vcs-repo-identifier\n%v", err)
// 	}
// 	if workspaceVcsRepoIdentifier != "" {
// 		workspace.VCSRepo.Identifier = workspaceVcsRepoIdentifier
// 	}

// 	workspaceVcsRepoIngressSubmodules, err := cmd.Flags().GetBool("workspace-vcs-repo-ingress-submodules")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-vcs-repo-ingress-submodules\n%v", err)
// 	} else {
// 		workspace.VCSRepo.IngressSubmodules = workspaceVcsRepoIngressSubmodules
// 	}

// 	workspaceVcsRepoOAuthTokenID, err := cmd.Flags().GetString("workspace-vcs-repo-o-auth-token-id")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-vcs-repo-o-auth-token-id\n%v", err)
// 	}
// 	if workspaceVcsRepoOAuthTokenID != "" {
// 		workspace.VCSRepo.OAuthTokenID = workspaceVcsRepoOAuthTokenID
// 	}

// 	workspaceWorkingDirectory, err := cmd.Flags().GetString("workspace-working-directory")
// 	if err != nil {
// 		logrus.Fatalf("unable to get flag workspace-working-directory\n%v", err)
// 	}
// 	if workspaceWorkingDirectory != "" {
// 		workspace.WorkingDirectory = workspaceWorkingDirectory
// 	}

// 	return workspace
// }

// // SetWorkspaceFlags TODO ...
// func SetWorkspaceFlags(cmd *cobra.Command) {

// 	usage := `TODO`
// 	cmd.Flags().String("workspace-id", "", usage)

// 	usage = `TODO`
// 	cmd.Flags().Bool("workspace-actions-is-destroyable", false, usage)

// 	usage = `TODO`
// 	cmd.Flags().String("workspace-agent-pool-id", "", usage)

// 	usage = `TODO`
// 	cmd.Flags().Bool("workspace-allow-destroy-plan", false, usage)

// 	usage = `TODO`
// 	cmd.Flags().Bool("workspace-auto-apply", false, usage)

// 	usage = `TODO`
// 	cmd.Flags().Bool("workspace-can-queue-destroy-plan", false, usage)

// 	usage = `TODO`
// 	cmd.Flags().String("workspace-environment", "", usage)

// 	usage = `TODO`
// 	cmd.Flags().String("workspace-execution-mode", "", usage)

// 	usage = `TODO`
// 	cmd.Flags().Bool("workspace-file-triggers-enabled", false, usage)

// 	usage = `TODO`
// 	cmd.Flags().Bool("workspace-locked", false, usage)

// 	usage = `TODO`
// 	cmd.Flags().String("workspace-migration-environment", "", usage)

// 	usage = `TODO`
// 	cmd.Flags().String("workspace-name", "", usage)

// 	usage = `TODO`
// 	cmd.Flags().Bool("workspace-permission-can-destroy", false, usage)

// 	usage = `TODO`
// 	cmd.Flags().Bool("workspace-permission-can-force-unlock", false, usage)

// 	usage = `TODO`
// 	cmd.Flags().Bool("workspace-permission-can-lock", false, usage)

// 	usage = `TODO`
// 	cmd.Flags().Bool("workspace-permission-can-queue-apply", false, usage)

// 	usage = `TODO`
// 	cmd.Flags().Bool("workspace-permission-can-queue-destroy", false, usage)

// 	usage = `TODO`
// 	cmd.Flags().Bool("workspace-permission-can-queue-run", false, usage)

// 	usage = `TODO`
// 	cmd.Flags().Bool("workspace-permission-can-read-settings", false, usage)

// 	usage = `TODO`
// 	cmd.Flags().Bool("workspace-permission-can-unlock", false, usage)

// 	usage = `TODO`
// 	cmd.Flags().Bool("workspace-permission-can-update", false, usage)

// 	usage = `TODO`
// 	cmd.Flags().Bool("workspace-permission-can-update-variable", false, usage)

// 	usage = `TODO`
// 	cmd.Flags().Bool("workspace-queue-all-runs", false, usage)

// 	usage = `TODO`
// 	cmd.Flags().Bool("workspace-speculative-enabled", false, usage)

// 	usage = `TODO`
// 	cmd.Flags().String("workspace-terraform-version", "", usage)

// 	usage = `TODO`
// 	var emptyArray []string
// 	cmd.Flags().StringArray("workspace-trigger-prefixes", emptyArray, usage)

// 	usage = `TODO`
// 	cmd.Flags().String("workspace-vcs-repo-branch", "", usage)

// 	usage = `TODO`
// 	cmd.Flags().String("workspace-vcs-repo-display-identifier", "", usage)

// 	usage = `TODO`
// 	cmd.Flags().String("workspace-vcs-repo-identifier", "", usage)

// 	usage = `TODO`
// 	cmd.Flags().Bool("workspace-vcs-repo-ingress-submodules", false, usage)

// 	usage = `TODO`
// 	cmd.Flags().String("workspace-vcs-repo-o-auth-token-id", "", usage)

// 	usage = `TODO`
// 	cmd.Flags().String("workspace-working-directory", "", usage)
// }