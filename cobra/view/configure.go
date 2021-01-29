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

package view

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/aid"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/dao"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/model"
)

// CREDENTIALS

// CreateCredentials create the credentials
func CreateCredentials(cmd *cobra.Command, name string, credentials model.Credentials) model.Credentials {
	fmt.Println("> Credentials")
	cProfile := createCredentialProfile(cmd, name)
	credentials.Profiles = append(credentials.Profiles, cProfile)
	return credentials
}

func createCredentialProfile(cmd *cobra.Command, name string) model.CredentialProfile {
	var cp model.CredentialProfile
	cp.Name = name
	cp.Description = "managed by tecli"
	cp.Enabled = true // enabling profile by default
	cp.CreatedAt = time.Now().String()
	cp.UpdatedAt = time.Now().String()

	return askAboutCredentialProfile(cmd, cp)
}

func askAboutCredentialProfile(cmd *cobra.Command, cp model.CredentialProfile) model.CredentialProfile {
	fmt.Println(">> Profile: " + cp.Name)
	cp.Name = aid.GetUserInputAsString(cmd, ">> Name", cp.Name)
	cp.Description = aid.GetUserInputAsString(cmd, ">> Description", cp.Description)
	cp.Enabled = aid.GetUserInputAsBool(cmd, ">> Enabled", cp.Enabled)
	cp.UserToken = aid.GetUserInputAsString(cmd, ">> User Token", cp.UserToken)
	cp.TeamToken = aid.GetUserInputAsString(cmd, ">> Team Token", cp.TeamToken)
	cp.OrganizationToken = aid.GetUserInputAsString(cmd, ">> Organization Token", cp.OrganizationToken)

	return cp
}

// UpdateCredentials update the given credentials
func UpdateCredentials(cmd *cobra.Command, name string) model.Credentials {
	fmt.Println("> Credentials")

	credentials, err := dao.GetCredentials()
	if err != nil {
		logrus.Fatalf("unable to update credentials\n%v\n", err)
	}

	found := false
	for i, profile := range credentials.Profiles {
		if profile.Name == name {
			found = true
			credentials.Profiles[i] = askAboutCredentialProfile(cmd, profile)
		}
	}

	if !found {
		fmt.Printf("No credentials not found for profile %s\n", name)
	}

	return credentials
}
