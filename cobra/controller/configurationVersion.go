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
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/go-tfe"
	"github.com/spf13/cobra"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/aid"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/dao"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/helper"
)

var configurationVersionValidArgs = []string{"list", "create", "read", "update", "delete"}

// ConfigurationVersionCmd command to display tecli current version
func ConfigurationVersionCmd() *cobra.Command {
	man, err := helper.GetManualV2("configurationVersion", configurationVersionValidArgs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		Example:   man.Example,
		ValidArgs: configurationVersionValidArgs,
		Args:      cobra.OnlyValidArgs,
		PreRunE:   configurationVersionPreRun,
		RunE:      configurationVersionRun,
	}

	aid.SetConfigurationVersionFlags(cmd, false)

	return cmd
}

func configurationVersionPreRun(cmd *cobra.Command, args []string) error {
	if err := helper.ValidateCmdArgs(cmd, args, "configurationVersion"); err != nil {
		return err
	}

	// fArg := args[0]
	// switch fArg {
	// case "create", "read", "list", "delete":
	// 	if err := helper.ValidateCmdArgAndFlag(cmd, args, "configurationVersion", fArg, "organization"); err != nil {
	// 		return err
	// 	}

	// 	if err := helper.ValidateCmdArgAndFlag(cmd, args, "configurationVersion", fArg, "name"); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

func configurationVersionRun(cmd *cobra.Command, args []string) error {

	token := dao.GetTeamToken(profile)
	client := aid.GetTFEClient(token)

	fArg := args[0]
	switch fArg {
	case "list":
		workspaceID, err := cmd.Flags().GetString("workspace-id")
		if err != nil {
			return fmt.Errorf("unable to get flag workspace-id\n%v", err)
		}

		list, err := configurationVersionList(client, workspaceID, tfe.ConfigurationVersionListOptions{})
		if err == nil {
			if len(list.Items) > 0 {
				for _, item := range list.Items {
					fmt.Printf("%v,\n", aid.ToJSON(item))
				}
			} else {
				return fmt.Errorf("no configurationVersion was found")
			}
		}
		// case "create":
		// 	options := aid.GetConfigurationVersionCreateOptions(cmd)
		// 	configurationVersion, err = configurationVersionCreate(client, options)

		// 	if err == nil && configurationVersion.ID != "" {
		// 		fmt.Println(aid.ToJSON(configurationVersion))
		// 	}
		// case "read":
		// 	name, err := cmd.Flags().GetString("name")
		// 	if err != nil {
		// 		return err
		// 	}

		// 	configurationVersion, err := configurationVersionRead(client, name)
		// 	if err == nil {
		// 		fmt.Println(aid.ToJSON(configurationVersion))
		// 	} else {
		// 		return fmt.Errorf("configurationVersion %s not found\n%v", name, err)
		// 	}
		// case "update":
		// 	name, err := cmd.Flags().GetString("name")
		// 	if err != nil {
		// 		return err
		// 	}

		// 	options := aid.GetConfigurationVersionUpdateOptions(cmd)
		// 	configurationVersion, err = configurationVersionUpdate(client, name, options)
		// 	if err == nil && configurationVersion.ID != "" {
		// 		fmt.Println(aid.ToJSON(configurationVersion))
		// 	} else {
		// 		return fmt.Errorf("unable to update configurationVersion\n%v", err)
		// 	}
		// case "delete":
		// 	name, err := cmd.Flags().GetString("name")
		// 	if err != nil {
		// 		return err
		// 	}

		// 	err = configurationVersionDelete(client, name)
		// 	if err == nil {
		// 		fmt.Printf("configurationVersion %s deleted successfully\n", name)
		// 	} else {
		// 		return fmt.Errorf("unable to delete configurationVersion %s\n%v", name, err)
		// 	}
	}

	return nil
}

// List returns all configuration versions of a workspace.
func configurationVersionList(client *tfe.Client, workspaceID string, options tfe.ConfigurationVersionListOptions) (*tfe.ConfigurationVersionList, error) {
	return client.ConfigurationVersions.List(context.Background(), workspaceID, options)
}

// Create is used to create a new configuration version. The created
// configuration version will be usable once data is uploaded to it.
func configurationVersionCreate(client *tfe.Client, workspaceID string, string, options tfe.ConfigurationVersionCreateOptions) (*tfe.ConfigurationVersion, error) {
	return client.ConfigurationVersions.Create(context.Background(), workspaceID, options)
}

// Read a configuration version by its ID.
func configurationVersionRead(client *tfe.Client, cvID string) (*tfe.ConfigurationVersion, error) {
	return client.ConfigurationVersions.Read(context.Background(), cvID)
}

// Upload packages and uploads Terraform configuration files. It requires
// the upload URL from a configuration version and the full path to the
// configuration files on disk.
func configurationVersionUpload(client *tfe.Client, url string, path string) error {
	return client.ConfigurationVersions.Upload(context.Background(), url, path)
}
