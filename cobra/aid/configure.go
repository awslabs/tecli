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
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.aws.dev/devops-aws/tecli/cobra/model"
	"gitlab.aws.dev/devops-aws/tecli/helper"
	"gopkg.in/yaml.v2"
)

// SetConfigureFlags TODO ...
func SetConfigureFlags(cmd *cobra.Command) {
	usage := `The new name of your profile`
	cmd.Flags().String("new-name", "", usage)

	usage = `A short description`
	cmd.Flags().String("description", "", usage)

	usage = `Enable or disable the entire profile`
	cmd.Flags().Bool("enabled", false, usage)

	usage = `API tokens may belong directly to a user. User tokens are the most flexible token type because they inherit permissions from the user they are associated with.`
	cmd.Flags().String("user-token", "", usage)

	usage = `API tokens may belong to a specific team. Team API tokens allow access to the workspaces that the team has access to, without being tied to any specific user.`
	cmd.Flags().String("team-token", "", usage)

	usage = `API tokens may generated for a specific organization. Organization API tokens allow access to the organization-level settings and resources, without being tied to any specific team or user.`
	cmd.Flags().String("organization-token", "", usage)
}

// GetCredentialProfileFlags TODO ...
func GetCredentialProfileFlags(cmd *cobra.Command) model.CredentialProfile {
	var cp model.CredentialProfile

	profile, err := cmd.Flags().GetString("profile")
	if err != nil {
		logrus.Fatalf("unable to get flag profile\n%v", err)
	}

	if profile != "" {
		cp.Name = profile
	}

	// new profile name replaces current profile name
	newName, err := cmd.Flags().GetString("new-name")
	if err != nil {
		logrus.Fatalf("unable to get flag new-name\n%v", err)
	}

	if newName != "" {
		cp.Name = newName
	}

	description, err := cmd.Flags().GetString("description")
	if err != nil {
		logrus.Fatalf("unable to get flag description\n%v", err)
	}

	if description != "" {
		cp.Description = description
	}

	enabled, err := cmd.Flags().GetBool("enabled")
	if err != nil {
		logrus.Fatalf("unable to get flag enabled\n%v", err)
	}

	cp.Enabled = enabled

	userToken, err := cmd.Flags().GetString("user-token")
	if err != nil {
		logrus.Fatalf("unable to get flag user-token\n%v", err)
	}

	if userToken != "" {
		cp.UserToken = userToken
	}

	teamToken, err := cmd.Flags().GetString("team-token")
	if err != nil {
		logrus.Fatalf("unable to get flag team-token\n%v", err)
	}

	if teamToken != "" {
		cp.TeamToken = teamToken
	}

	organizationToken, err := cmd.Flags().GetString("organization-token")
	if err != nil {
		logrus.Fatalf("unable to get flag organization-token\n%v", err)
	}

	if organizationToken != "" {
		cp.OrganizationToken = organizationToken
	}

	return cp
}

// HasCreatedConfigurationDir return true if configuration directory was created, false if otherwise
func HasCreatedConfigurationDir() (bool, string) {
	if !ConfigurationsDirectoryExist() {
		return CreateConfigurationsDirectory()
	}
	return false, ""
}

// ConfigurationsDirectoryExist returns `true` if the configuration directory exist, `false` otherwise
func ConfigurationsDirectoryExist() bool {
	return helper.DirOrFileExists(GetAppInfo().ConfigurationsDir)
}

// ConfigurationsFileExist returns `true` if the configuration file exist, `false` otherwise
func ConfigurationsFileExist() bool {
	return helper.DirOrFileExists(GetAppInfo().ConfigurationsPath)
}

// CreateConfigurationsDirectory creates the configuration directory, returns `true` if the configuration directory exist, `false` otherwise
func CreateConfigurationsDirectory() (bool, string) {
	dir := GetAppInfo().ConfigurationsDir
	return helper.MkDirsIfNotExist(dir), dir
}

// CredentialsFileExist returns `true` if the credentials file exist, `false` otherwise
func CredentialsFileExist() bool {
	return helper.DirOrFileExists(GetAppInfo().CredentialsPath)
}

// ReadConfig returns the viper instance of the given configuration `name`
func ReadConfig(name string) (*viper.Viper, error) {
	v := viper.New()
	app := GetAppInfo()

	v.SetConfigName(name)
	v.SetConfigType("yaml")
	v.AddConfigPath(app.ConfigurationsDir)

	err := v.ReadInConfig()
	if err != nil {
		return v, fmt.Errorf("unable to read configuration:%s\n%v", name, err)
	}
	return v, err
}

// ReadConfigAsViper returns...
func ReadConfigAsViper(configPath string, configName string, configType string) (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath(configPath)
	v.SetConfigName(configName)
	v.SetConfigType(configType)

	err := v.ReadInConfig()
	if err != nil {
		return v, fmt.Errorf("unable to read configuration as viper\n%v", err)
	}
	return v, err
}

// ReadTemplate read the given template under tecli/*.yaml
func ReadTemplate(fileName string) (*viper.Viper, error) {
	c := viper.New()
	c.AddConfigPath("tecli")
	c.SetConfigName(fileName)
	c.SetConfigType("yaml")
	c.SetConfigPermissions(os.ModePerm)

	err := c.ReadInConfig() // Find and read the c file
	if err != nil {         // Handle errors reading the c file
		return c, fmt.Errorf("Unable to read "+fileName+" via Viper"+"\n%v", err)
	}

	return c, nil
}

// WriteInterfaceToFile write the given interface into a file
func WriteInterfaceToFile(in interface{}, path string) error {
	b, err := yaml.Marshal(&in)
	if err != nil {
		_, ok := err.(*json.UnsupportedTypeError)
		if ok {
			return fmt.Errorf("json unsupported type error")
		}
	}

	err = ioutil.WriteFile(path, b, os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to update:%s\n%v", path, err)
	}

	return err
}

// DeleteCredentialFile delete the credentials file
func DeleteCredentialFile() error {
	return helper.DeleteFile(GetAppInfo().CredentialsPath)
}

// DeleteConfigurationFile delete the credentials file
func DeleteConfigurationFile() error {
	return helper.DeleteFile(GetAppInfo().ConfigurationsPath)
}

// DeleteConfigurationsDirectory delete the configurations directory
func DeleteConfigurationsDirectory() error {
	return os.RemoveAll(GetAppInfo().ConfigurationsDir)
}

// GetSensitiveUserInput get sensitive input as string
func GetSensitiveUserInput(cmd *cobra.Command, text string, info string) (string, error) {
	return getUserInput(cmd, text+" ["+maskString(info, 3)+"]", "")
}

func maskString(s string, showLastChars int) string {
	maskSize := len(s) - showLastChars
	if maskSize <= 0 {
		return s
	}

	return strings.Repeat("*", maskSize) + s[maskSize:]
}

// GetSensitiveUserInputAsString get sensitive input as string
func GetSensitiveUserInputAsString(cmd *cobra.Command, text string, info string) string {
	answer, err := GetSensitiveUserInput(cmd, text, info)
	if err != nil {
		log.Fatalf("unable to get user input about profile's name\n%v\n", err)
	}

	// if user typed ENTER, keep the current value
	if answer != "" {
		return answer
	}

	return info
}

func getUserInput(cmd *cobra.Command, text string, info string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	if info == "" {
		fmt.Print(text + ": ")
	} else {
		fmt.Print(text + " [" + info + "]: ")
	}

	input, err := reader.ReadString('\n')
	// convert CRLF to LF
	input = strings.Replace(input, "\n", "", -1)
	if err != nil {
		return input, fmt.Errorf("unable to read user input\n%v", err)
	}

	return input, err
}

// GetUserInputAsBool prints `text` on console and return answer as `boolean`
func GetUserInputAsBool(cmd *cobra.Command, text string, info bool) bool {
	answer, err := getUserInput(cmd, text, strconv.FormatBool(info))
	if err != nil {
		log.Fatalf("unable to get user input as boolean\n%v", err)
	}

	if answer == "true" {
		return true
	} else if answer == "false" {
		return false
	}

	return info
}

// GetUserInputAsString prints `text` on console and return answer as `string`
func GetUserInputAsString(cmd *cobra.Command, text string, info string) string {
	answer, err := getUserInput(cmd, text, info)
	if err != nil {
		log.Fatalf("unable to get user input about profile's name\n%v\n", err)
	}

	// if user typed ENTER, keep the current value
	if answer != "" {
		return answer
	}

	return info
}
