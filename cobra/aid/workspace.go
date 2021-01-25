package aid

import (
	"context"
	"encoding/json"

	tfe "github.com/hashicorp/go-tfe"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// GetOrganizations TODO ...
func GetOrganizations(client *tfe.Client) (*tfe.OrganizationList, error) {
	orgs, err := client.Organizations.List(context.Background(), tfe.OrganizationListOptions{})
	if err != nil {
		logrus.Fatal(err)
	}

	return orgs, err
}

// GetOrganizationByName TODO ...
func GetOrganizationByName(client *tfe.Client, name string) tfe.Organization {
	orgs, err := GetOrganizations(client)
	if err != nil {
		logrus.Fatalf("unable to get organization by name\n%v\n", err)
	}

	if len(orgs.Items) == 0 {
		logrus.Fatalf("no organization found")
	}

	var org tfe.Organization
	for _, item := range orgs.Items {
		if item.Name == name {
			org = *item
		}
	}

	return org
}

// GetWorkspaceCreateOptions TODO ..
func GetWorkspaceCreateOptions(cmd *cobra.Command) tfe.WorkspaceCreateOptions {

	agentPoolID, err := cmd.Flags().GetString("agent-pool-id")
	if err != nil {
		logrus.Fatalf("unable to get flag agent-pool-id\n%v\n", err)
	}

	allowDestroyPlan, err := cmd.Flags().GetBool("allow-destroy-plan")
	if err != nil {
		logrus.Fatalf("unable to get flag allow-destroy-plan\n%v\n", err)
	}

	autoApply, err := cmd.Flags().GetBool("auto-apply")
	if err != nil {
		logrus.Fatalf("unable to get flag auto-apply\n%v\n", err)
	}

	executionMode, err := cmd.Flags().GetString("execution-mode")
	if err != nil {
		logrus.Fatalf("unable to get flag execution-mode\n%v\n", err)
	}

	fileTriggersEnabled, err := cmd.Flags().GetBool("file-triggers-enabled")
	if err != nil {
		logrus.Fatalf("unable to get flag file-triggers-enabled\n%v\n", err)
	}

	migrationEnvironment, err := cmd.Flags().GetString("migration-environment")
	if err != nil {
		logrus.Fatalf("unable to get flag migration-environment\n%v\n", err)
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		logrus.Fatalf("unable to get flag name\n%v\n", err)
	}

	queueAllRuns, err := cmd.Flags().GetBool("queue-all-runs")
	if err != nil {
		logrus.Fatalf("unable to get flag queue-all-runs\n%v\n", err)
	}

	speculativeEnabled, err := cmd.Flags().GetBool("speculative-enabled")
	if err != nil {
		logrus.Fatalf("unable to get flag speculative-enabled\n%v\n", err)
	}

	terraformVersion, err := cmd.Flags().GetString("terraform-version")
	if err != nil {
		logrus.Fatalf("unable to get flag terraform-version\n%v\n", err)
	}

	triggerPrefixes, err := cmd.Flags().GetStringArray("trigger-prefixes")
	if err != nil {
		logrus.Fatalf("unable to get flag trigger-prefixes\n%v\n", err)
	}

	workingDirectory, err := cmd.Flags().GetString("working-directory")
	if err != nil {
		logrus.Fatalf("unable to get flag working-directory\n%v\n", err)
	}

	// vcs

	vcsRepoBranch, err := cmd.Flags().GetString("vcs-repo-branch")
	if err != nil {
		logrus.Fatalf("unable to get flag vcsRepoBranch\n%v\n", err)
	}

	vcsRepoIdentifier, err := cmd.Flags().GetString("vcs-repo-identifier")
	if err != nil {
		logrus.Fatalf("unable to get flag vcsRepoIdentifier\n%v\n", err)
	}

	vcsRepoIngressSubmodules, err := cmd.Flags().GetBool("vcs-repo-ingress-submodules")
	if err != nil {
		logrus.Fatalf("unable to get flag vcsRepoIngressSubmodules\n%v\n", err)
	}

	vcsRepoOauthTokenID, err := cmd.Flags().GetString("vcs-repo-oauth-token-id")
	if err != nil {
		logrus.Fatalf("unable to get flag vcsRepoOauthTokenId\n%v\n", err)
	}

	var repoOptions tfe.VCSRepoOptions

	if vcsRepoBranch != "" {
		repoOptions.Branch = &vcsRepoBranch
	}

	if vcsRepoIdentifier != "" {
		repoOptions.Identifier = &vcsRepoIdentifier
	}

	if vcsRepoIngressSubmodules {
		repoOptions.IngressSubmodules = &vcsRepoIngressSubmodules
	}

	if vcsRepoOauthTokenID != "" {
		repoOptions.OAuthTokenID = &vcsRepoOauthTokenID
	}

	// workspace

	var options tfe.WorkspaceCreateOptions
	if agentPoolID != "" {
		options.AgentPoolID = &agentPoolID
	}

	if allowDestroyPlan {
		options.AllowDestroyPlan = &allowDestroyPlan
	}

	if autoApply {
		options.AutoApply = &autoApply
	}

	if executionMode != "" {
		options.ExecutionMode = &executionMode
	}

	if fileTriggersEnabled {
		options.FileTriggersEnabled = &fileTriggersEnabled
	}

	if migrationEnvironment != "" {
		options.MigrationEnvironment = &migrationEnvironment
	}

	if name != "" {
		options.Name = &name
	}

	// if *operations != "" {
	// options.Operations = &operations
	// }

	if queueAllRuns {
		options.QueueAllRuns = &queueAllRuns
	}

	if speculativeEnabled {
		options.SpeculativeEnabled = &speculativeEnabled
	}

	if terraformVersion != "" {
		options.TerraformVersion = &terraformVersion
	}

	if len(triggerPrefixes) > 0 {
		options.TriggerPrefixes = triggerPrefixes
	}

	if repoOptions != (tfe.VCSRepoOptions{}) {
		options.VCSRepo = &repoOptions
	}

	if workingDirectory != "" {
		options.WorkingDirectory = &workingDirectory
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

	if allowDestroyPlan {
		options.AllowDestroyPlan = &allowDestroyPlan
	}

	// Whether to automatically apply changes when a Terraform plan is successful.
	autoApply, err := cmd.Flags().GetBool("auto-apply")
	if err != nil {
		logrus.Fatalf("unable to get flag auto-apply\n%v\n", err)
	}

	if allowDestroyPlan {
		options.AutoApply = &autoApply
	}

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

	if allowDestroyPlan {
		options.FileTriggersEnabled = &fileTriggersEnabled
	}

	// Whether to queue all runs. Unless this is set to true, runs triggered by
	// a webhook will not be queued until at least one run is manually queued.
	queueAllRuns, err := cmd.Flags().GetBool("queue-all-runs")
	if err != nil {
		logrus.Fatalf("unable to get flag queue-all-runs\n%v\n", err)
	}

	if allowDestroyPlan {
		options.QueueAllRuns = &queueAllRuns
	}

	// Whether this workspace allows speculative plans. Setting this to false
	// prevents Terraform Cloud or the Terraform Enterprise instance from
	// running plans on pull requests, which can improve security if the VCS
	// repository is public or includes untrusted contributors.
	speculativeEnabled, err := cmd.Flags().GetBool("speculative-enabled")
	if err != nil {
		logrus.Fatalf("unable to get flag speculative-enabled\n%v\n", err)
	}

	if allowDestroyPlan {
		options.SpeculativeEnabled = &speculativeEnabled
	}

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

	var repoOptions tfe.VCSRepoOptions
	vcsRepoBranch, err := cmd.Flags().GetString("vcs-repo-branch")
	if err != nil {
		logrus.Fatalf("unable to get flag vcsRepoBranch\n%v\n", err)
	}

	if vcsRepoBranch != "" {
		repoOptions.Branch = &vcsRepoBranch
	}

	vcsRepoIdentifier, err := cmd.Flags().GetString("vcs-repo-identifier")
	if err != nil {
		logrus.Fatalf("unable to get flag vcsRepoIdentifier\n%v\n", err)
	}

	if vcsRepoIdentifier != "" {
		repoOptions.Identifier = &vcsRepoIdentifier
	}

	vcsRepoIngressSubmodules, err := cmd.Flags().GetBool("vcs-repo-ingress-submodules")
	if err != nil {
		logrus.Fatalf("unable to get flag vcsRepoIngressSubmodules\n%v\n", err)
	}

	if vcsRepoIngressSubmodules {
		repoOptions.IngressSubmodules = &vcsRepoIngressSubmodules
	}

	vcsRepoOauthTokenID, err := cmd.Flags().GetString("vcs-repo-oauth-token-id")
	if err != nil {
		logrus.Fatalf("unable to get flag vcsRepoOauthTokenId\n%v\n", err)
	}

	if vcsRepoOauthTokenID != "" {
		repoOptions.OAuthTokenID = &vcsRepoOauthTokenID
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

// ToJSON converts a given struct to json
func ToJSON(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		logrus.Fatalf("unable to convert struct to json\n%v\n", b)
	}

	return string(b)
}

// GetWorkspaceAssignSSHKeyOptions TODO ...
func GetWorkspaceAssignSSHKeyOptions(cmd *cobra.Command) tfe.WorkspaceAssignSSHKeyOptions {
	var options tfe.WorkspaceAssignSSHKeyOptions

	sshKeyID, err := cmd.Flags().GetString("ssh-key-id")
	if err != nil {
		logrus.Fatalf("unable to get flag ssh-key-id\n%v\n", err)
	}

	options.SSHKeyID = &sshKeyID

	return options
}
