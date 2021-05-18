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

// Package view represents the visualization of the data that model contains.
package view

import (
	"fmt"

	"github.com/awslabs/tecli/cobra/aid"
	"github.com/awslabs/tecli/cobra/dao"
	"github.com/awslabs/tecli/cobra/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CREDENTIALS

// CreateCredentials ask user for input and create the credentials
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
	return AskAboutCredentialProfile(cmd, cp)
}

// AskAboutCredentialProfile TODO ...
func AskAboutCredentialProfile(cmd *cobra.Command, cp model.CredentialProfile) model.CredentialProfile {
	fmt.Println(">> Profile: " + cp.Name)
	cp.Name = aid.GetUserInputAsString(cmd, ">> Name", cp.Name)
	cp.Description = aid.GetUserInputAsString(cmd, ">> Description", cp.Description)
	cp.Organization = aid.GetUserInputAsString(cmd, ">> Organization", cp.Organization)
	cp.UserToken = aid.GetUserInputAsString(cmd, ">> User Token", cp.UserToken)
	cp.TeamToken = aid.GetUserInputAsString(cmd, ">> Team Token", cp.TeamToken)
	cp.OrganizationToken = aid.GetUserInputAsString(cmd, ">> Organization Token", cp.OrganizationToken)

	return cp
}

// UpdateCredentials update the given credentials
func UpdateCredentials(cmd *cobra.Command, name string) (model.Credentials, error) {
	fmt.Println("> Credentials")

	credentials, err := dao.GetCredentials()
	if err != nil {
		logrus.Fatalf("unable to update credentials\n%v\n", err)
	}

	found := false
	for i, profile := range credentials.Profiles {
		if profile.Name == name {
			found = true
			credentials.Profiles[i] = AskAboutCredentialProfile(cmd, profile)
		}
	}

	if !found {
		return model.Credentials{}, fmt.Errorf("profile %s not found", name)
	}

	return credentials, nil
}
