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
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var sshKeyValidArgs = []string{"list", "create", "read", "update", "delete"}

// SSHKeyCmd command to display tecli current version
func SSHKeyCmd() *cobra.Command {
	man, err := helper.GetManual("ssh-key", sshKeyValidArgs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:          man.Use,
		Short:        man.Short,
		Long:         man.Long,
		Example:      man.Example,
		ValidArgs:    sshKeyValidArgs,
		Args:         cobra.OnlyValidArgs,
		PreRunE:      sshKeyPreRun,
		RunE:         sshKeyRun,
		SilenceUsage: true,
	}

	usage := `SSH key ID. Required for read and delete.`
	cmd.Flags().String("id", "", usage)

	usage = `A name to identify the SSH key.`
	cmd.Flags().String("name", "", usage)

	usage = `The content of the SSH private key.`
	cmd.Flags().String("value", "", usage)

	return cmd
}

func sshKeyPreRun(cmd *cobra.Command, args []string) error {
	if err := helper.ValidateCmdArgs(cmd, args, "ssh-key"); err != nil {
		return err
	}

	fArg := args[0]
	switch fArg {
	case "create":
		if err := helper.ValidateCmdArgAndFlag(cmd, args, "ssh-key", fArg, "name"); err != nil {
			return err
		}

		if err := helper.ValidateCmdArgAndFlag(cmd, args, "ssh-key", fArg, "value"); err != nil {
			return err
		}
	case "read":
		if err := helper.ValidateCmdArgAndFlag(cmd, args, "ssh-key", fArg, "id"); err != nil {
			return err
		}
	case "update":
		if err := helper.ValidateCmdArgAndFlag(cmd, args, "ssh-key", fArg, "id"); err != nil {
			return err
		}

		if err := helper.ValidateCmdArgAndFlag(cmd, args, "ssh-key", fArg, "name"); err != nil {
			return err
		}

		if err := helper.ValidateCmdArgAndFlag(cmd, args, "ssh-key", fArg, "value"); err != nil {
			return err
		}
	case "delete":
		if err := helper.ValidateCmdArgAndFlag(cmd, args, "ssh-key", fArg, "id"); err != nil {
			return err
		}
	}

	return nil
}

func sshKeyRun(cmd *cobra.Command, args []string) error {
	// config, err := cmd.Flags().GetString("config")
	// if err != nil {
	// 	return fmt.Errorf("unable to get flag config\n%v", err)
	// }

	// aid.LoadViper(config)

	token := dao.GetTeamToken(profile)
	client := aid.GetTFEClient(token)

	var sshKey *tfe.SSHKey
	var err error

	fArg := args[0]
	switch fArg {
	case "list":
		organization := dao.GetOrganization(profile)
		list, err := sshKeyList(client, organization)
		if err == nil {
			fmt.Println(aid.ToJSON(list))
		} else {
			return fmt.Errorf("no ssh key was found")
		}

	case "create":
		organization := dao.GetOrganization(profile)
		options := aid.GetSSHKeysCreateOptions(cmd)
		sshKey, err = sshKeyCreate(client, organization, options)
		if err != nil {
			logrus.Errorln("unable to create ssh key")
			return err
		}

		if err == nil && sshKey.ID != "" {
			fmt.Println(aid.ToJSON(sshKey))
		}
	case "read":
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return fmt.Errorf("unable to get flag id\n%v", err)
		}

		sshKey, err := sshKeyRead(client, id)
		if err == nil {
			fmt.Println(aid.ToJSON(sshKey))
		} else {
			return fmt.Errorf("ssh key %s not found\n%v", id, err)
		}

	case "update":
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return fmt.Errorf("unable to get flag id\n%v", err)
		}

		options := aid.GetSSHKeysUpdateOptions(cmd)
		sshKey, err = sshKeyUpdate(client, id, options)
		if err == nil && sshKey.ID != "" {
			fmt.Println(aid.ToJSON(sshKey))
		} else {
			return fmt.Errorf("unable to update ssh key\n%v", err)
		}
	case "delete":
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return fmt.Errorf("unable to get flag id\n%v", err)
		}

		err = sshKeyDelete(client, id)
		if err == nil {
			cmd.Printf("ssh key %s deleted successfully\n", id)
		} else {
			return fmt.Errorf("unable to delete ssh key %s\n%v", id, err)
		}
	}

	return nil
}

func sshKeyList(client *tfe.Client, organization string) (*tfe.SSHKeyList, error) {
	return client.SSHKeys.List(context.Background(), organization, tfe.SSHKeyListOptions{})
}

// Create is used to create a new sshKey.
func sshKeyCreate(client *tfe.Client, organization string, options tfe.SSHKeyCreateOptions) (*tfe.SSHKey, error) {
	return client.SSHKeys.Create(context.Background(), organization, options)
}

// Read a sshKey by its name.
func sshKeyRead(client *tfe.Client, sshKeyID string) (*tfe.SSHKey, error) {
	return client.SSHKeys.Read(context.Background(), sshKeyID)
}

// Update settings of an existing sshKey.
func sshKeyUpdate(client *tfe.Client, sshKeyID string, options tfe.SSHKeyUpdateOptions) (*tfe.SSHKey, error) {
	return client.SSHKeys.Update(context.Background(), sshKeyID, options)
}

// // Delete a sshKey by its name.
func sshKeyDelete(client *tfe.Client, sshKeyID string) error {
	return client.SSHKeys.Delete(context.Background(), sshKeyID)
}
