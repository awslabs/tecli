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
	cmd.Println(">> Profile: " + name)
	var cProfile model.CredentialProfile
	cProfile.Name = name
	cProfile.Description = "managed by tecli"
	cProfile.Enabled = true // enabling profile by default
	cProfile.CreatedAt = time.Now().String()
	cProfile.UpdatedAt = time.Now().String()

	var cred model.Credential
	cred.Enabled = true
	cred.CreatedAt = time.Now().String()
	cred = askAboutCredential(cmd, cred)

	cProfile.Credentials = append(cProfile.Credentials, cred)

	for {
		answer := aid.GetUserInputAsBool(cmd, "Would you like to setup another credential?", false)
		if answer {
			var newCred model.Credential
			newCred.Enabled = true
			newCred.CreatedAt = time.Now().String()
			newCred = askAboutCredential(cmd, newCred)
			cProfile.Credentials = append(cProfile.Credentials, newCred)
		} else {
			break
		}
	}

	return cProfile
}

func askAboutCredential(cmd *cobra.Command, credential model.Credential) model.Credential {
	cmd.Println(">>>> Credential")
	credential.Name = aid.GetUserInputAsString(cmd, ">>>>> Name", credential.Name)
	credential.Description = aid.GetUserInputAsString(cmd, ">>>>> Description", credential.Description)
	credential.Enabled = aid.GetUserInputAsBool(cmd, ">>>>> Enabled", credential.Enabled)

	if credential.CreatedAt == "" {
		credential.CreatedAt = time.Now().String()
	}

	credential.UpdatedAt = time.Now().String()
	credential.Provider = aid.GetUserInputAsString(cmd, ">>>>> Provider", credential.Provider)
	credential.AccessKey = aid.GetSensitiveUserInputAsString(cmd, ">>>>> Access Key", credential.AccessKey)
	credential.SecretKey = aid.GetSensitiveUserInputAsString(cmd, ">>>>> Secret Key", credential.SecretKey)
	credential.SessionToken = aid.GetSensitiveUserInputAsString(cmd, ">>>>> Session Token", credential.SessionToken)
	return credential
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

func askAboutCredentialProfile(cmd *cobra.Command, profile model.CredentialProfile) model.CredentialProfile {
	cmd.Println(">> Profile: " + profile.Name)
	profile.Name = aid.GetUserInputAsString(cmd, ">>> Name", profile.Name)
	profile.Description = aid.GetUserInputAsString(cmd, ">>> Description", profile.Description)
	profile.Enabled = aid.GetUserInputAsBool(cmd, ">>> Enabled", profile.Enabled)

	if profile.CreatedAt == "" {
		profile.CreatedAt = time.Now().String()
	}

	profile.UpdatedAt = time.Now().String()

	for i, credential := range profile.Credentials {
		profile.Credentials[i] = askAboutCredential(cmd, credential)
	}

	return profile
}

// CONFIGURATIONS

// CreateConfigurations create the configuration file with the given profile name
func CreateConfigurations(cmd *cobra.Command, name string) model.Configurations {
	cmd.Println("> Configurations")
	var configurations model.Configurations
	cProfile := createConfigurationProfile(cmd, name)
	configurations.Profiles = append(configurations.Profiles, cProfile)
	return configurations
}

// createConfigurationProfile create the given profile name into the configurations file, return the profile created
func createConfigurationProfile(cmd *cobra.Command, name string) model.ConfigurationProfile {
	cmd.Println(">> Profile: " + name)
	var cProfile model.ConfigurationProfile
	cProfile.Name = name
	cProfile.Description = "managed by tecli"
	cProfile.Enabled = true // enabling profile by default
	cProfile.CreatedAt = time.Now().String()
	cProfile.UpdatedAt = time.Now().String()

	var conf model.Configuration
	conf.Enabled = true
	conf.CreatedAt = time.Now().String()
	conf = askAboutConfiguration(cmd, conf)
	cProfile.Configurations = append(cProfile.Configurations, conf)

	for {
		answer := aid.GetUserInputAsBool(cmd, "Would you like to setup another configuration?", false)
		if answer {
			var newConf model.Configuration
			newConf.Enabled = true
			newConf = askAboutConfiguration(cmd, newConf)
			cProfile.Configurations = append(cProfile.Configurations, newConf)
		} else {
			break
		}
	}

	return cProfile
}

func askAboutConfiguration(cmd *cobra.Command, conf model.Configuration) model.Configuration {
	cmd.Println(">>> Configuration")
	conf.Name = aid.GetUserInputAsString(cmd, ">>>> Name", conf.Name)
	conf.Description = aid.GetUserInputAsString(cmd, ">>>> Description", conf.Description)
	conf.Enabled = aid.GetUserInputAsBool(cmd, ">>>> Enabled", conf.Enabled)

	if conf.CreatedAt == "" {
		conf.CreatedAt = time.Now().String()
	}
	return conf
}

// UpdateConfigurations update the given configurations
func UpdateConfigurations(cmd *cobra.Command, name string) model.Configurations {
	cmd.Println("> Configurations")
	configurations, err := dao.GetConfigurations()
	if err != nil {
		logrus.Fatalf("unable to update configurations\n%v", err)
	}

	found := false
	for i, profile := range configurations.Profiles {
		if profile.Name == name {
			found = true
			configurations.Profiles[i] = askAboutConfigurationProfile(cmd, profile)
		}
	}

	if !found {
		cmd.Printf("No configurations not found for profile %s\n", name)
	}

	return configurations
}

func askAboutConfigurationProfile(cmd *cobra.Command, profile model.ConfigurationProfile) model.ConfigurationProfile {
	cmd.Println(">> Profile: " + profile.Name)
	profile.Name = aid.GetUserInputAsString(cmd, ">>> Name", profile.Name)
	profile.Description = aid.GetUserInputAsString(cmd, ">>> Description", profile.Description)
	profile.Enabled = aid.GetUserInputAsBool(cmd, ">>> Enabled", profile.Enabled)

	if profile.CreatedAt == "" {
		profile.CreatedAt = time.Now().String()
	}

	profile.UpdatedAt = time.Now().String()

	for i, configuration := range profile.Configurations {
		profile.Configurations[i] = askAboutConfiguration(cmd, configuration)
	}

	return profile
}
