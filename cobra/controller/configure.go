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

	"github.com/spf13/cobra"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/aid"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/dao"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/view"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/helper"
)

var configureValidArgs = []string{"create", "delete"}

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

	return cmd
}

func configurePreRun(cmd *cobra.Command, args []string) error {
	if err := helper.ValidateCmdArgs(cmd, args, "configure"); err != nil {
		return err
	}

	fArg := args[0]
	switch fArg {
	case "create", "delete":
		if err := helper.ValidateCmdArgAndFlag(cmd, args, "configure", fArg, "profile"); err != nil {
			return err
		}

	default:
		return fmt.Errorf("unknown argument provided")
	}
	return nil
}

func configureRun(cmd *cobra.Command, args []string) error {
	if !aid.ConfigurationsDirectoryExist() {
		if created, dir := aid.CreateConfigurationsDirectory(); created {
			fmt.Printf("tecli configuration directory created at %s\n", dir)
			createCredentials(cmd)
			// createConfigurations(cmd)
		}
	} else {
		// configurations directory exist
		if !aid.CredentialsFileExist() {
			createCredentials(cmd)
		} else {
			updateCredentials(cmd)
		}

		// if !aid.ConfigurationsFileExist() {
		// 	createConfigurations(cmd)
		// } else {
		// 	updateConfigurations(cmd)
		// }
	}

	return nil
}

func createCredentials(cmd *cobra.Command) {
	answer := aid.GetUserInputAsBool(cmd, "Would you like to setup credentials?", false)
	if answer {
		credentials := view.CreateCredentials(cmd, profile)
		dao.SaveCredentials(credentials)
	}
}

func updateCredentials(cmd *cobra.Command) {
	answer := aid.GetUserInputAsBool(cmd, "Would you like to update credentials?", false)
	if answer {
		credentials := view.UpdateCredentials(cmd, profile)
		dao.SaveCredentials(credentials)
	}
}
