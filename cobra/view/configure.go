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

	"github.com/awslabs/tfe-cli/cobra/aid"
	"github.com/awslabs/tfe-cli/cobra/dao"
	"github.com/awslabs/tfe-cli/cobra/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
	cProfile.Description = "managed by tfe-cli"
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
	cProfile.Description = "managed by tfe-cli"
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

	conf.UpdatedAt = time.Now().String()
	conf.Unsplash = askAboutUnsplashConfiguration(cmd, conf.Unsplash)
	conf.Initialization = askAboutInitialization(cmd, conf.Initialization)
	return conf
}

// askAboutUnsplashConfiguration ask user about configuration
func askAboutUnsplashConfiguration(cmd *cobra.Command, unsplash model.Unsplash) model.Unsplash {
	// configuration can have many types: Unsplash, AWS, etc
	answer := aid.GetUserInputAsBool(cmd, "Would you like to setup Unsplash configuration?", false)
	if answer {
		cmd.Println(">>>> Unsplash")
		unsplash.Name = aid.GetUserInputAsString(cmd, ">>>>> Name", unsplash.Name)
		unsplash.Description = aid.GetUserInputAsString(cmd, ">>>>> Description", unsplash.Description)
		unsplash.Enabled = aid.GetUserInputAsBool(cmd, ">>>>> Enabled", unsplash.Enabled)

		if unsplash.CreatedAt == "" {
			unsplash.CreatedAt = time.Now().String()
		}

		unsplash.UpdatedAt = time.Now().String()
		unsplash.RandomPhoto = askAboutUnsplashRandomPhoto(cmd, unsplash.RandomPhoto)

	} else {
		cmd.Println("Skipping Unplash configuration ...")
	}
	return unsplash
}

func askAboutUnsplashRandomPhoto(cmd *cobra.Command, randomPhoto model.UnsplashRandomPhoto) model.UnsplashRandomPhoto {
	// unsplash configuration may have multiple nested configuration, such as random photo, etc...
	answer := aid.GetUserInputAsBool(cmd, "Would you like to setup Unsplash Random Photo Parameters?", false)

	if answer {
		cmd.Println(">>>>>> Random Photo")
		randomPhoto.Name = aid.GetUserInputAsString(cmd, ">>>>>>> Name", randomPhoto.Name)
		randomPhoto.Description = aid.GetUserInputAsString(cmd, ">>>>>>> Description", randomPhoto.Description)
		randomPhoto.Enabled = aid.GetUserInputAsBool(cmd, ">>>>>>> Enabled", randomPhoto.Enabled)

		if randomPhoto.CreatedAt == "" {
			randomPhoto.CreatedAt = time.Now().String()
		}

		randomPhoto.UpdatedAt = time.Now().String()
		randomPhoto.Parameters = askAboutUnsplashRandomPhotoParameters(cmd, randomPhoto.Parameters)

	} else {
		cmd.Println("Skipping Unplash Random Photo configuration ...")
	}

	return randomPhoto
}

func askAboutUnsplashRandomPhotoParameters(cmd *cobra.Command, params model.UnsplashRandomPhotoParameters) model.UnsplashRandomPhotoParameters {
	cmd.Println(">>>>>>> Parameters")
	params.Collections = aid.GetUserInputAsString(cmd, ">>>>>>>> Collections (Public collection ID to filter selection. If multiple, comma-separated) ", params.Collections)
	params.Featured = aid.GetUserInputAsBool(cmd, ">>>>>>>> Featured (Limit selection to featured photos. Valid values: false and true)", params.Featured)
	params.Filter = aid.GetUserInputAsString(cmd, ">>>>>>>> Filter (Limit results by content safety. Valid values are low and high)", params.Filter)
	params.Orientation = aid.GetUserInputAsString(cmd, ">>>>>>>> Orientation (Filter by photo orientation. Valid values: landscape, portrait, squarish)", params.Orientation)
	params.Query = aid.GetUserInputAsString(cmd, ">>>>>>>> Query (Limit selection to photos matching a search term)", params.Query)
	params.Size = aid.GetUserInputAsString(cmd, ">>>>>>>> Size (Photos size. Valid values: all, thumb, small, regular, full, raw)", params.Size)
	params.Username = aid.GetUserInputAsString(cmd, ">>>>>>>> Username (Limit selection to a single user)", params.Username)

	return params
}

func askAboutInitialization(cmd *cobra.Command, init model.Initialization) model.Initialization {
	// configuration can have many types: Unsplash, AWS, etc
	answer := aid.GetUserInputAsBool(cmd, "Would you like to setup a customized initialization?", false)
	if answer {

		cmd.Println(">>>> Initialization")
		init.Name = aid.GetUserInputAsString(cmd, ">>>> Name", init.Name)
		init.Description = aid.GetUserInputAsString(cmd, ">>>> Description", init.Description)
		init.Enabled = aid.GetUserInputAsBool(cmd, ">>>> Enabled", init.Enabled)
		init.Type = aid.GetUserInputAsString(cmd, ">>>> Type", init.Type)

		if init.CreatedAt == "" {
			init.CreatedAt = time.Now().String()
		}
		init.UpdatedAt = time.Now().String()

		if len(init.Files) == 0 {
			var file model.File
			file = askAboutInitializationFile(cmd, file)
			init.Files = append(init.Files, file)
		} else {
			for i, f := range init.Files {
				init.Files[i] = askAboutInitializationFile(cmd, f)
			}
		}

		for {
			answer = aid.GetUserInputAsBool(cmd, "Would you like to setup a new file?", false)
			if answer {
				var file model.File
				file = askAboutInitializationFile(cmd, file)
				init.Files = append(init.Files, file)
			} else {
				break
			}
		}

	}

	return init
}

func askAboutInitializationFile(cmd *cobra.Command, file model.File) model.File {
	cmd.Println(">>>>> File")
	file.Path = aid.GetUserInputAsString(cmd, ">>>>>> Path", file.Path)
	file.Src = aid.GetUserInputAsString(cmd, ">>>>>> Source", file.Src)
	file.Dest = aid.GetUserInputAsString(cmd, ">>>>>> Destination", file.Dest)
	file.State = aid.GetUserInputAsString(cmd, ">>>>>> State", file.State)
	return file
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
