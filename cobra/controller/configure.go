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

	"github.com/awslabs/tecli/cobra/aid"
	"github.com/awslabs/tecli/cobra/dao"
	"github.com/awslabs/tecli/cobra/model"
	"github.com/awslabs/tecli/cobra/view"
	"github.com/awslabs/tecli/helper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
		Use:          man.Use,
		Short:        man.Short,
		Long:         man.Long,
		Example:      man.Example,
		ValidArgs:    configureValidArgs,
		Args:         cobra.OnlyValidArgs,
		PreRunE:      configurePreRun,
		RunE:         configureRun,
		SilenceUsage: true,
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

	if cmd.Flags().Changed("mode") {
		mode, err := cmd.Flags().GetString("mode")
		if err != nil {
			return fmt.Errorf("unable to get flag mode\n%v", err)
		}

		if mode != "interactive" && mode != "non-interactive" {
			return fmt.Errorf("invalid mode provided, mode can be only: interactive or non-interactive")
		}
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
		err = configureCreateCredentials(cmd, mode)
		if err != nil {
			return fmt.Errorf("unable to create profile\n%v", err)
		}
		cmd.Printf("profile %s created successfully\n", profile)

	case "read":
		c, err := configureReadCredentials(cmd)
		if err != nil {
			return fmt.Errorf("unable to read credential")
		}
		fmt.Println(aid.ToJSON(c))

	case "update":
		err = configureUpdateCredentials(cmd, mode)
		if err != nil {
			return fmt.Errorf("unable to update profile\n%v", err)
		}
		cmd.Printf("profile %s updated successfully\n", profile)

	case "delete":
		err := configureDeleteCredential()
		if err != nil {
			return fmt.Errorf("unable to delete profile\n%v", err)
		}
		cmd.Printf("profile %s deleted successfully\n", profile)

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

func configureCreateCredentials(cmd *cobra.Command, mode string) error {
	created, err := aid.HasCreatedAppDir(cmd)
	if err != nil {
		return err
	}

	if created {
		var creds model.Credentials
		if mode == "interactive" {
			creds = view.CreateCredentials(cmd, profile, model.Credentials{})
		} else if mode == "non-interactive" {
			c := aid.GetCredentialProfileFlags(cmd)
			creds.Profiles = append(creds.Profiles, c)
		}

		return dao.SaveCredentials(creds)
	}

	creds, err := dao.GetCredentials()
	if err != nil {
		logrus.Fatalf("unable to get credentials\n%v", err)
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
	if mode == "interactive" {
		creds = view.CreateCredentials(cmd, profile, creds)
	} else if mode == "non-interactive" {
		c := aid.GetCredentialProfileFlags(cmd)
		creds.Profiles = append(creds.Profiles, c)
	}

	return dao.SaveCredentials(creds)
}

func configureReadCredentials(cmd *cobra.Command) (model.CredentialProfile, error) {
	return dao.GetCredentialProfile(profile)
}

func configureUpdateCredentials(cmd *cobra.Command, mode string) error {
	var err error
	if err = aid.CheckAppDirAndFile(); err == nil {
		creds, err := dao.GetCredentials()
		if err != nil {
			logrus.Fatalf("unable to update credentials\n%v\n", err)
		}

		found := false
		for i, p := range creds.Profiles {
			if p.Name == profile {
				found = true
				if mode == "interactive" {
					creds.Profiles[i] = view.AskAboutCredentialProfile(cmd, p)
				} else if mode == "non-interactive" {
					creds.Profiles[i] = aid.UpdateCredentialProfile(cmd, p)
				}
			}
		}

		if !found {
			return fmt.Errorf("profile %s not found", profile)
		}

		return dao.SaveCredentials(creds)
	}

	return err
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
			newCreds.Profiles = aid.RemoveCredential(creds.Profiles, i)
			dao.SaveCredentials(newCreds)
			break
		}
	}

	if !found {
		return fmt.Errorf("profile %s not found", profile)
	}

	return nil
}
