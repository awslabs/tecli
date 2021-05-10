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

	"github.com/awslabs/tecli/cobra/aid"
	"github.com/awslabs/tecli/cobra/dao"
	"github.com/awslabs/tecli/helper"
	"github.com/hashicorp/go-tfe"
	"github.com/spf13/cobra"
)

var configurationVersionValidArgs = []string{"list", "create", "read", "upload"}

// ConfigurationVersionCmd command to display tecli current version
func ConfigurationVersionCmd() *cobra.Command {
	man, err := helper.GetManual("configuration-version", configurationVersionValidArgs)
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

	aid.SetConfigurationVersionFlags(cmd)

	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)

	return cmd
}

func configurationVersionPreRun(cmd *cobra.Command, args []string) error {
	if err := helper.ValidateCmdArgs(cmd, args, "configuration-version"); err != nil {
		return err
	}

	fArg := args[0]
	switch fArg {
	case "list", "create":
		if err := helper.ValidateCmdArgAndFlag(cmd, args, "configuration-version", fArg, "workspace-id"); err != nil {
			return err
		}

	case "read":
		if err := helper.ValidateCmdArgAndFlag(cmd, args, "configuration-version", fArg, "id"); err != nil {
			return err
		}
	case "upload":
		if err := helper.ValidateCmdArgAndFlag(cmd, args, "configuration-version", fArg, "url"); err != nil {
			return err
		}

		if err := helper.ValidateCmdArgAndFlag(cmd, args, "configuration-version", fArg, "path"); err != nil {
			return err
		}
	}

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
			aid.PrintConfigurationVersionList(list)
		} else {
			return fmt.Errorf("no configurationVersion was found")
		}

	case "create":
		workspaceID, err := cmd.Flags().GetString("workspace-id")
		if err != nil {
			return fmt.Errorf("unable to get flag workspace-id\n%v", err)
		}

		options := aid.GetConfigurationVersionCreateOptions(cmd)
		cv, err := configurationVersionCreate(client, workspaceID, options)

		if err == nil && cv.ID != "" {
			cmd.Println(aid.ToJSON(cv))
		} else {
			return fmt.Errorf("unable to create configuration version\n%v", err)
		}
	case "read":
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return fmt.Errorf("unable to get flag id\n%v", err)
		}

		cv, err := configurationVersionRead(client, id)
		if err == nil {
			cmd.Println(aid.ToJSON(cv))
		} else {
			return fmt.Errorf("configuration version %s not found\n%v", id, err)
		}

	case "upload":
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			return fmt.Errorf("unable to get flag url\n%v", err)
		}

		path, err := cmd.Flags().GetString("path")
		if err != nil {
			return fmt.Errorf("unable to get flag path\n%v", err)
		}

		err = configurationVersionUpload(client, url, path)
		if err == nil {
			cmd.Println("upload completed successfully")
		} else {
			return fmt.Errorf("unable to upload to configuration version\n%v", err)
		}
	}

	return nil
}

// List returns all configuration versions of a workspace.
func configurationVersionList(client *tfe.Client, workspaceID string, options tfe.ConfigurationVersionListOptions) (*tfe.ConfigurationVersionList, error) {
	return client.ConfigurationVersions.List(context.Background(), workspaceID, options)
}

// Create is used to create a new configuration version. The created
// configuration version will be usable once data is uploaded to it.
func configurationVersionCreate(client *tfe.Client, workspaceID string, options tfe.ConfigurationVersionCreateOptions) (*tfe.ConfigurationVersion, error) {
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
