package aid

import (
	"context"

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
		logrus.Fatalf("unable to get organization by name\n%v", err)
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
		logrus.Fatalf("unable to get flag agent-pool-id\n%v", err)
	}

	allowDestroyPlan, err := cmd.Flags().GetBool("allow-destroy-plan")
	if err != nil {
		logrus.Fatalf("unable to get flag allow-destroy-plan\n%v", err)
	}

	autoApply, err := cmd.Flags().GetBool("auto-apply")
	if err != nil {
		logrus.Fatalf("unable to get flag auto-apply\n%v", err)
	}

	executionMode, err := cmd.Flags().GetString("execution-mode")
	if err != nil {
		logrus.Fatalf("unable to get flag execution-mode\n%v", err)
	}

	fileTriggersEnabled, err := cmd.Flags().GetBool("file-triggers-enabled")
	if err != nil {
		logrus.Fatalf("unable to get flag file-triggers-enabled\n%v", err)
	}

	migrationEnvironment, err := cmd.Flags().GetString("migration-environment")
	if err != nil {
		logrus.Fatalf("unable to get flag migration-environment\n%v", err)
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		logrus.Fatalf("unable to get flag name\n%v", err)
	}

	queueAllRuns, err := cmd.Flags().GetBool("queue-all-runs")
	if err != nil {
		logrus.Fatalf("unable to get flag queue-all-runs\n%v", err)
	}

	speculativeEnabled, err := cmd.Flags().GetBool("speculative-enabled")
	if err != nil {
		logrus.Fatalf("unable to get flag speculative-enabled\n%v", err)
	}

	terraformVersion, err := cmd.Flags().GetString("terraform-version")
	if err != nil {
		logrus.Fatalf("unable to get flag terraform-version\n%v", err)
	}

	workingDirectory, err := cmd.Flags().GetString("working-directory")
	if err != nil {
		logrus.Fatalf("unable to get flag working-directory\n%v", err)
	}

	// vcs

	vcsRepoBranch, err := cmd.Flags().GetString("vcs-repo-branch")
	if err != nil {
		logrus.Fatalf("unable to get flag vcsRepoBranch\n%v", err)
	}

	vcsRepoIdentifier, err := cmd.Flags().GetString("vcs-repo-identifier")
	if err != nil {
		logrus.Fatalf("unable to get flag vcsRepoIdentifier\n%v", err)
	}

	vcsRepoIngressSubmodules, err := cmd.Flags().GetBool("vcs-repo-ingress-submodules")
	if err != nil {
		logrus.Fatalf("unable to get flag vcsRepoIngressSubmodules\n%v", err)
	}

	vcsRepoOauthTokenID, err := cmd.Flags().GetString("vcs-repo-oauth-token-id")
	if err != nil {
		logrus.Fatalf("unable to get flag vcsRepoOauthTokenId\n%v", err)
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

	// if *triggerPrefixes != "" {
	// options.TriggerPrefixes = &triggerPrefixes
	// }

	if repoOptions != (tfe.VCSRepoOptions{}) {
		options.VCSRepo = &repoOptions
	}

	if workingDirectory != "" {
		options.WorkingDirectory = &workingDirectory
	}

	return options

}
