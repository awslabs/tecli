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
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/aid"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/dao"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/model"
)

// CREDENTIALS

// CreateCredentials create the credentials
func CreateCredentials(cmd *cobra.Command, name string) model.Credentials {
	cmd.Println("> Credentials")
	var credentials model.Credentials
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
	cmd.Println(">> Profile: " + cp.Name)
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
	cmd.Println("> Credentials")

	credentials, err := dao.GetCredentials()
	if err != nil {
		logrus.Fatalf("unable to update credentials\n%v", err)
	}

	found := false
	for i, profile := range credentials.Profiles {
		if profile.Name == name {
			found = true
			credentials.Profiles[i] = askAboutCredentialProfile(cmd, profile)
		}
	}

	if !found {
		cmd.Printf("No credentials not found for profile %s\n", name)
	}

	return credentials
}

// CONFIGURATIONS

// // CreateConfigurations create the configuration file with the given profile name
// func CreateConfigurations(cmd *cobra.Command, name string) model.Configurations {
// 	cmd.Println("> Configurations")
// 	var configurations model.Configurations
// 	cProfile := createConfigurationProfile(cmd, name)
// 	configurations.Profiles = append(configurations.Profiles, cProfile)
// 	return configurations
// }

// // createConfigurationProfile create the given profile name into the configurations file, return the profile created
// func createConfigurationProfile(cmd *cobra.Command, name string) model.ConfigurationProfile {
// 	cmd.Println(">> Profile: " + name)
// 	var cProfile model.ConfigurationProfile
// 	cProfile.Name = name
// 	cProfile.Description = "managed by tecli"
// 	cProfile.Enabled = true // enabling profile by default
// 	cProfile.CreatedAt = time.Now().String()
// 	cProfile.UpdatedAt = time.Now().String()

// 	var conf model.Configuration
// 	conf.Enabled = true
// 	conf.CreatedAt = time.Now().String()
// 	conf = askAboutConfiguration(cmd, conf)
// 	cProfile.Configurations = append(cProfile.Configurations, conf)

// 	for {
// 		answer := aid.GetUserInputAsBool(cmd, "Would you like to setup another configuration?", false)
// 		if answer {
// 			var newConf model.Configuration
// 			newConf.Enabled = true
// 			newConf = askAboutConfiguration(cmd, newConf)
// 			cProfile.Configurations = append(cProfile.Configurations, newConf)
// 		} else {
// 			break
// 		}
// 	}

// 	return cProfile
// }

// func askAboutConfiguration(cmd *cobra.Command, conf model.Configuration) model.Configuration {
// 	cmd.Println(">>> Configuration")
// 	conf.Name = aid.GetUserInputAsString(cmd, ">>>> Name", conf.Name)
// 	conf.Description = aid.GetUserInputAsString(cmd, ">>>> Description", conf.Description)
// 	conf.Enabled = aid.GetUserInputAsBool(cmd, ">>>> Enabled", conf.Enabled)

// 	if conf.CreatedAt == "" {
// 		conf.CreatedAt = time.Now().String()
// 	}
// 	return conf
// }

// // UpdateConfigurations update the given configurations
// func UpdateConfigurations(cmd *cobra.Command, name string) model.Configurations {
// 	cmd.Println("> Configurations")
// 	configurations, err := dao.GetConfigurations()
// 	if err != nil {
// 		logrus.Fatalf("unable to update configurations\n%v", err)
// 	}

// 	found := false
// 	for i, profile := range configurations.Profiles {
// 		if profile.Name == name {
// 			found = true
// 			configurations.Profiles[i] = askAboutConfigurationProfile(cmd, profile)
// 		}
// 	}

// 	if !found {
// 		cmd.Printf("No configurations not found for profile %s\n", name)
// 	}

// 	return configurations
// }

// func askAboutConfigurationProfile(cmd *cobra.Command, profile model.ConfigurationProfile) model.ConfigurationProfile {
// 	cmd.Println(">> Profile: " + profile.Name)
// 	profile.Name = aid.GetUserInputAsString(cmd, ">>> Name", profile.Name)
// 	profile.Description = aid.GetUserInputAsString(cmd, ">>> Description", profile.Description)
// 	profile.Enabled = aid.GetUserInputAsBool(cmd, ">>> Enabled", profile.Enabled)

// 	if profile.CreatedAt == "" {
// 		profile.CreatedAt = time.Now().String()
// 	}

// 	profile.UpdatedAt = time.Now().String()

// 	for i, configuration := range profile.Configurations {
// 		profile.Configurations[i] = askAboutConfiguration(cmd, configuration)
// 	}

// 	return profile
// }
