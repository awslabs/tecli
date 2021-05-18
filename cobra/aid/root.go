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
	"fmt"
	"os"

	"github.com/awslabs/tecli/cobra/model"
	tfe "github.com/hashicorp/go-tfe"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// GetAppInfo return information about tecli settings
func GetAppInfo() model.App {
	var err error

	configDir, err := os.UserConfigDir()
	if err != nil {
		logrus.Fatalf("unable to read user configuration directory\n%v", err)
	}

	sep := string(os.PathSeparator)

	var app model.App
	app.Name = "tecli"
	app.ConfigurationsDir = configDir + sep + app.Name
	app.CredentialsFileName = "credentials"
	app.CredentialsFileType = "yaml"
	app.CredentialsFilePath = app.ConfigurationsDir + sep + app.CredentialsFileName + "." + app.CredentialsFileType

	// ~/.tecli/logs.json
	app.LogsFileName = "logs"
	app.LogsFileType = "json"
	app.LogsFilePath = app.ConfigurationsDir + sep + app.LogsFileName + "." + app.LogsFileType
	app.LogsFilePermissions = os.ModePerm

	app.WorkingDir, err = os.Getwd()
	if err != nil {
		fmt.Printf("Unable to detect the current directory\n%v", err)
		os.Exit(1)
	}

	return app
}

// SetupLoggingOutput set logrun output file
func SetupLoggingOutput(path string) error {
	// define default path
	if path == "" {
		path = GetAppInfo().LogsFilePath
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to open log file\n%v", err)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(file)

	return nil
}

// LoadViper setup project setting
func LoadViper() {
	// environment variables
	viper.SetEnvPrefix("TFC") // will be uppercased automatically
	viper.BindEnv("ORGANIZATION")
	viper.BindEnv("USER_TOKEN")
	viper.BindEnv("TEAM_TOKEN")
	viper.BindEnv("ORGANIZATION_TOKEN")

	app := GetAppInfo()

	viper.SetConfigName(app.CredentialsFileName)
	viper.SetConfigType(app.CredentialsFileType) // REQUIRED if the config file does not have the extension in the name

	viper.AddConfigPath(app.ConfigurationsDir)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("using config file:", viper.ConfigFileUsed())
	}
	// if config is not found, that's okay, as the user might use env vars
}

// SetupLoggingLevel set logrus level
func SetupLoggingLevel(level string) error {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return fmt.Errorf("unable to set log level\n%v", err)
	}

	logrus.SetLevel(lvl)
	return nil
}

// Returns struct from Terraform Enterprise Cloud API response
func getTFEConfig(token string) *tfe.Config {
	config := &tfe.Config{
		Token: token,
	}
	return config
}

// Returns a new terraform api client
func getTFENewClient(config *tfe.Config) (*tfe.Client, error) {
	client, err := tfe.NewClient(config)
	if err != nil {
		logrus.Errorf("unable to get new terraform enterprise api client\n%v\n", err)
		return client, err
	}

	return client, err
}

// GetTFEClient returns a new terraform api client given a token
func GetTFEClient(token string) *tfe.Client {
	config := getTFEConfig(token)
	client, err := getTFENewClient(config)
	if err != nil {
		logrus.Fatalf("unable to get terraform cloud api client\n%v\n", err)
	}

	return client
}
