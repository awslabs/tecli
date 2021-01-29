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

// SetOAuthClientFlags define flags for the cobra command
func SetOAuthClientFlags(cmd *cobra.Command) {
	usage := `The OAuth Client ID.`
	cmd.Flags().String("id", "", usage)
	usage = `The base URL of your VCS provider's API.`
	cmd.Flags().String("api-url", "", usage)
	usage = `The homepage of your VCS provider.`
	cmd.Flags().String("http-url", "", usage)
	usage = `The token string you were given by your VCS provider.`
	cmd.Flags().String("o-auth-token", "", usage)
	usage = `Private key associated with this vcs provider - only available for azure-devops-server`
	cmd.Flags().String("private-key", "", usage)
	usage = `The VCS provider being connected with. Valid values azure-devops-server, azure-devops-services, bitbucket-hosted, bitbucket-server, bitbucket-server-legacy, github, github-enterprise, gitlab-hosted, gitlab-community-edition, gitlab-enterprise-edition.`
	cmd.Flags().String("service-provider", "", usage)
}

// GetOAuthClientCreateOptions return options based on the flags values
func GetOAuthClientCreateOptions(cmd *cobra.Command) tfe.OAuthClientCreateOptions {
	var options tfe.OAuthClientCreateOptions

	// The base URL of your VCS provider's API.
	apiURL, err := cmd.Flags().GetString("api-url")
	if err != nil {
		logrus.Fatalf("unable to get flag api-url")
	}

	if apiURL != "" {
		options.APIURL = &apiURL
	}

	// The homepage of your VCS provider.
	httpURL, err := cmd.Flags().GetString("http-url")
	if err != nil {
		logrus.Fatalf("unable to get flag http-url")
	}

	if httpURL != "" {
		options.HTTPURL = &httpURL
	}

	// The token string you were given by your VCS provider.
	oAuthToken, err := cmd.Flags().GetString("o-auth-token")
	if err != nil {
		logrus.Fatalf("unable to get flag o-auth-token")
	}

	if oAuthToken != "" {
		options.OAuthToken = &oAuthToken
	}

	// Private key associated with this vcs provider - only available for azure-devops-server
	privateKey, err := cmd.Flags().GetString("private-key")
	if err != nil {
		logrus.Fatalf("unable to get flag private-key")
	}

	if privateKey != "" {
		options.PrivateKey = &privateKey
	}

	// The VCS provider being connected with.
	serviceProvider, err := cmd.Flags().GetString("service-provider")
	if err != nil {
		logrus.Fatalf("unable to get flag service-provider")
	}

	if serviceProvider != "" {
		var sp tfe.ServiceProviderType

		if serviceProvider == "azure-devops-server" {
			sp = tfe.ServiceProviderAzureDevOpsServer
		}
		if serviceProvider == "azure-devops-services" {
			sp = tfe.ServiceProviderAzureDevOpsServices
		}
		if serviceProvider == "bitbucket-hosted" {
			sp = tfe.ServiceProviderBitbucket
		}
		if serviceProvider == "bitbucket-server" {
			sp = tfe.ServiceProviderBitbucketServer
		}
		if serviceProvider == "bitbucket-server-legacy" {
			sp = tfe.ServiceProviderBitbucketServerLegacy
		}
		if serviceProvider == "github" {
			sp = tfe.ServiceProviderGithub
		}
		if serviceProvider == "github-enterprise" {
			sp = tfe.ServiceProviderGithubEE
		}
		if serviceProvider == "gitlab-hosted" {
			sp = tfe.ServiceProviderGitlab
		}
		if serviceProvider == "gitlab-community-edition" {
			sp = tfe.ServiceProviderGitlabCE
		}
		if serviceProvider == "gitlab-enterprise-edition" {
			sp = tfe.ServiceProviderGitlabEE
		}

		options.ServiceProvider = &sp
	}

	return options
}
