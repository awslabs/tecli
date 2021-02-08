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

package controller

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.aws.dev/devops-aws/tecli/cobra/aid"
	"gitlab.aws.dev/devops-aws/tecli/cobra/dao"
	"gitlab.aws.dev/devops-aws/tecli/cobra/model"
	"gitlab.aws.dev/devops-aws/tecli/cobra/view"
	"gitlab.aws.dev/devops-aws/tecli/helper"
)

var configureValidArgs = []string{"list", "create", "read", "update", "delete"}

// ConfigureCmd command to display tecli current version
func ConfigureCmd() *cobra.Command {
	man, err := helper.GetManual("configure", configureValidArgs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		Example:   man.Example,
		ValidArgs: configureValidArgs,
		Args:      cobra.OnlyValidArgs,
		PreRunE:   configurePreRun,
		RunE:      configureRun,
	}

	aid.SetConfigureFlags(cmd)

	return cmd
}

func configurePreRun(cmd *cobra.Command, args []string) error {
	if err := helper.ValidateCmdArgs(cmd, args, "configure"); err != nil {
		return err
	}

	fArg := args[0]
	switch fArg {
	case "list", "create", "read", "update", "delete":
		if err := helper.ValidateCmdArgAndFlag(cmd, args, "configure", fArg, "profile"); err != nil {
			return err
		}

	default:
		return fmt.Errorf("unknown argument provided")
	}
	return nil
}

func configureRun(cmd *cobra.Command, args []string) error {
	mode, err := cmd.Flags().GetString("mode")
	if err != nil {
		return fmt.Errorf("unable to get flag mode\n%v", err)
	}

	fArg := args[0]
	switch fArg {
	case "list":
		creds, err := configureListCredentials()
		if err != nil {
			logrus.Fatalf("unable to list credentials\n%v", err)
		}

		fmt.Println(aid.ToJSON(creds))
	case "create":
		if mode == "interactive" {
			err = configureCreateCredentialsInteractive(cmd)
		} else if mode == "non-interactive" {
			err = configureCreateCredentialsNonInteractive(cmd)
		}

		if err != nil {
			return fmt.Errorf("unable to create profile\n%v", err)
		}

		fmt.Printf("profile %s created successfully\n", profile)
	case "read":
		c, err := configureReadCredentials(cmd)
		if err != nil {
			return fmt.Errorf("unable to read credential")
		}
		fmt.Println(aid.ToJSON(c))
	case "update":
		if mode == "interactive" {
			err = configureUpdateCredentialsInteractive(cmd)
		} else if mode == "non-interactive" {
			err = configureUpdateCredentialsNonInteractive(cmd)
		}

		if err != nil {
			return fmt.Errorf("unable to update profile\n%v", err)
		}

		fmt.Printf("profile %s updated successfully\n", profile)
	case "delete":
		err := configureDeleteCredential()
		if err != nil {
			return fmt.Errorf("unable to delete profile\n%v", err)
		}

		fmt.Printf("profile %s delete successfully\n", profile)
	default:
		return fmt.Errorf("unknown argument provided")
	}

	return nil
}

func configureListCredentials() (model.Credentials, error) {
	creds, err := dao.GetCredentials()
	if err != nil {
		return creds, err
	}

	return creds, nil
}

func configureCreateCredentialsInteractive(cmd *cobra.Command) error {
	ft, dir := aid.HasCreatedConfigurationDir()

	if ft {
		fmt.Printf("tecli configuration directory created at %s\n", dir)
	}

	if !aid.CredentialsFileExist() {
		creds := view.CreateCredentials(cmd, profile, model.Credentials{})
		return dao.SaveCredentials(creds)
	}

	creds, err := dao.GetCredentials()
	if err != nil {
		logrus.Fatalf("unable to get credentials")
	}

	found := false
	for _, p := range creds.Profiles {
		if p.Name == profile {
			found = true
		}
	}

	if found {
		// don't add duplicates
		return fmt.Errorf("profile %s already exist\nprofile names must be unique", profile)
	}

	// append new cred to credentials file
	creds = view.CreateCredentials(cmd, profile, creds)
	return dao.SaveCredentials(creds)
}

func configureCreateCredentialsNonInteractive(cmd *cobra.Command) error {
	c := aid.GetCredentialProfileFlags(cmd)

	// enable profile by default
	if !cmd.Flags().Changed("enabled") {
		c.Enabled = true
	}

	ft, dir := aid.HasCreatedConfigurationDir()

	if ft {
		fmt.Printf("tecli configuration directory created at %s\n", dir)
	}

	if !aid.CredentialsFileExist() {
		var creds model.Credentials
		creds.Profiles = append(creds.Profiles, c)
		return dao.SaveCredentials(creds)
	}

	creds, err := dao.GetCredentials()
	if err != nil {
		logrus.Fatalf("unable to get credentials")
	}

	found := false
	for _, p := range creds.Profiles {
		if p.Name == profile {
			found = true
		}
	}

	if found {
		// don't add duplicates
		return fmt.Errorf("profile %s already exist, profile names must be unique", profile)
	}

	creds.Profiles = append(creds.Profiles, c)
	return dao.SaveCredentials(creds)
}

func configureReadCredentials(cmd *cobra.Command) (model.CredentialProfile, error) {
	return dao.GetCredentialProfile(profile)
}

func configureUpdateCredentialsInteractive(cmd *cobra.Command) error {
	if !aid.ConfigurationsDirectoryExist() {
		return fmt.Errorf("tecli configuration directory not found\nplease run configure create")
	}

	if !aid.CredentialsFileExist() {
		return fmt.Errorf("credentials file not found\nplease run configure create")
	}

	creds, err := view.UpdateCredentials(cmd, profile)
	if err != nil {
		return err
	}

	return dao.SaveCredentials(creds)
}

func configureUpdateCredentialsNonInteractive(cmd *cobra.Command) error {
	if !aid.ConfigurationsDirectoryExist() {
		return fmt.Errorf("configuration directory not found\nplease run configure create")
	}

	if !aid.CredentialsFileExist() {
		return fmt.Errorf("credentials file not found\nplease run configure create")
	}

	creds, err := dao.GetCredentials()
	if err != nil {
		logrus.Fatalf("unable to update credentials\n%v\n", err)
	}

	found := false
	for i, p := range creds.Profiles {
		if p.Name == profile {
			found = true
			creds.Profiles[i] = aid.GetCredentialProfileFlags(cmd)
		}
	}

	if !found {
		return fmt.Errorf("profile %s not found", profile)
	}

	return dao.SaveCredentials(creds)
}

func removeCredential(s []model.CredentialProfile, index int) []model.CredentialProfile {
	return append(s[:index], s[index+1:]...)
}

func configureDeleteCredential() error {
	creds, err := dao.GetCredentials()
	if err != nil {
		return err
	}

	found := false
	for i, p := range creds.Profiles {
		if p.Name == profile {
			found = true

			var newCreds model.Credentials
			newCreds.Profiles = removeCredential(creds.Profiles, i)
			dao.SaveCredentials(newCreds)
			break
		}
	}

	if !found {
		return fmt.Errorf("profile %s not found", profile)
	}

	return nil
}
