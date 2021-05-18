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
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/awslabs/tecli/cobra/model"
	"github.com/awslabs/tecli/helper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// SetConfigureFlags TODO ...
func SetConfigureFlags(cmd *cobra.Command) {
	usage := `Valid values: interactive or non-interactive. Interactive mode ask inputs to user. Non-interactive assumes all information is passed via flags`
	cmd.Flags().String("mode", "interactive", usage)

	usage = `The new name of your profile`
	cmd.Flags().String("new-name", "", usage)

	usage = `A short description`
	cmd.Flags().String("description", "", usage)

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

// UpdateCredentialProfile update credential profile based on input flags
func UpdateCredentialProfile(cmd *cobra.Command, cp model.CredentialProfile) model.CredentialProfile {
	f := GetCredentialProfileFlags(cmd)

	if f.Name != "" && f.Name != cp.Name {
		cp.Name = f.Name
	}

	if f.Description != "" && f.Description != cp.Description {
		cp.Description = f.Description
	}

	if f.Organization != "" && f.Organization != cp.Organization {
		cp.Organization = f.Organization
	}

	if f.UserToken != "" && f.UserToken != cp.UserToken {
		cp.UserToken = f.UserToken
	}

	if f.TeamToken != "" && f.TeamToken != cp.TeamToken {
		cp.TeamToken = f.TeamToken
	}

	if f.OrganizationToken != "" && f.OrganizationToken != cp.OrganizationToken {
		cp.OrganizationToken = f.OrganizationToken
	}

	return cp
}

// CheckAppDirAndFile checks if configuration directory and file exist
func CheckAppDirAndFile() error {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			return fmt.Errorf("configuration file not found\n%v", err)
		}

		// Config file was found but another error was produced
		return fmt.Errorf("unable to read config file\n%v", err)
	}

	return nil
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

func getInput() (string, error) {

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	if runtime.GOOS == "windows" {
		// convert LF to CRLF
		text = strings.Replace(text, "\r\n", "", -1)
	} else {
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
	}

	return text, nil
}

func getUserInput(cmd *cobra.Command, text string, info string) (string, error) {
	if info == "" {
		cmd.Print(text + ": ")
	} else {
		cmd.Print(text + " [" + info + "]: ")
	}

	input, err := getInput()

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

// RemoveCredential TODO ...
func RemoveCredential(s []model.CredentialProfile, index int) []model.CredentialProfile {
	return append(s[:index], s[index+1:]...)
}

// HasCreatedAppDir TODO ...
func HasCreatedAppDir(cmd *cobra.Command) (bool, error) {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// not found anywhere, crea dir

			dirPath := GetAppInfo().ConfigurationsDir
			if helper.MkDirsIfNotExist(dirPath) {
				fmt.Printf("tecli configuration directory created at %s\n", dirPath)

				// necessary to create an empty file for Viper
				f, err := os.Create(GetAppInfo().CredentialsFilePath)
				if err != nil {
					return false, fmt.Errorf("unable to create empty credentials file\n%v", err)
				}
				f.Close()

				// reload viper
				LoadViper()

				return true, nil
			}
		}
	}

	return false, nil
}
