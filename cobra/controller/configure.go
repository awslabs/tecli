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
	fArg := args[0]
	switch fArg {
	case "list":
		creds, err := configureListCredentials()
		if err != nil {
			logrus.Fatalf("unable to list credentials\n%v", err)
		}

		fmt.Println(aid.ToJSON(creds))
	case "create":
		configureCreateCredentials(cmd)
	case "read":
		c, err := configureReadCredentials(cmd)
		if err != nil {
			return fmt.Errorf("unable to read credential")
		}
		fmt.Println(aid.ToJSON(c))
	case "update":
		configureUpdateCredentials(cmd)
	case "delete":
		err := configureDeleteCredential()
		if err != nil {
			return err
		}
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

func configureCreateCredentials(cmd *cobra.Command) {
	c := aid.GetCredentialProfileFlags(cmd)
	var ft bool
	if ft, dir := aid.HasCreatedConfigurationDir(); ft {
		fmt.Printf("tecli configuration directory created at %s\n", dir)
	}

	if (c != model.CredentialProfile{}) {
		// non-interactive
		if ft {
			var creds model.Credentials
			creds.Profiles = append(creds.Profiles, c)
			dao.SaveCredentials(creds)
		} else {
			updateOrAppendCredential(c)
		}
	} else {
		// interactive
		configureCreateCredentialsInteractive(cmd)
	}
}

func updateOrAppendCredential(c model.CredentialProfile) {
	if aid.CredentialsFileExist() {
		creds, err := dao.GetCredentials()
		if err != nil {
			logrus.Fatalf("unable to get credentials")
		}

		found := false
		for i, p := range creds.Profiles {
			if p.Name == profile || p.Name == c.Name {
				creds.Profiles[i] = c
				found = true
			}
		}

		if !found {
			creds.Profiles = append(creds.Profiles, c)
		}

		dao.SaveCredentials(creds)
	} else {
		logrus.Fatalln("tecli credentials file not found")
	}
}

func configureCreateCredentialsInteractive(cmd *cobra.Command) {
	if !aid.CredentialsFileExist() {
		creds := view.CreateCredentialsInteractive(cmd, profile, model.Credentials{})
		dao.SaveCredentials(creds)
	} else {
		if aid.CredentialsFileExist() {
			creds, err := dao.GetCredentials()
			if err != nil {
				logrus.Fatalf("unable to get credentials")
			}

			found := false
			for _, p := range creds.Profiles {
				if p.Name == profile {
					configureUpdateCredentials(cmd)
					found = true
				}
			}

			if !found {
				creds = view.CreateCredentialsInteractive(cmd, profile, creds)
				dao.SaveCredentials(creds)
			}
		}

	}
}

func configureReadCredentials(cmd *cobra.Command) (model.CredentialProfile, error) {
	return dao.GetCredentialProfile(profile)
}

func configureUpdateCredentials(cmd *cobra.Command) {
	credentials := view.UpdateCredentials(cmd, profile)
	dao.SaveCredentials(credentials)
}

func removeCredential(s []model.CredentialProfile, index int) []model.CredentialProfile {
	return append(s[:index], s[index+1:]...)
}

func configureDeleteCredential() error {
	creds, err := dao.GetCredentials()
	if err != nil {
		return fmt.Errorf("unable to delete credential\n%v", err)
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
		return fmt.Errorf("unable to find the profile %s and delete its credentials", profile)
	}

	return nil
}
