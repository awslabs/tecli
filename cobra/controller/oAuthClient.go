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
	"gitlab.aws.dev/devops-aws/tecli/cobra/aid"
	"gitlab.aws.dev/devops-aws/tecli/cobra/dao"
	"gitlab.aws.dev/devops-aws/tecli/helper"
)

var oAuthClientValidArgs = []string{
	"list",
	"create",
	"read",
	"delete",
}

// OAuthClientCmd command to display tecli current version
func OAuthClientCmd() *cobra.Command {
	man, err := helper.GetManual("o-auth-client", oAuthClientValidArgs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		Example:   man.Example,
		ValidArgs: oAuthClientValidArgs,
		Args:      cobra.OnlyValidArgs,
		PreRunE:   oAuthClientPreRun,
		RunE:      oAuthClientRun,
	}

	aid.SetOAuthClientFlags(cmd)

	return cmd
}

func oAuthClientPreRun(cmd *cobra.Command, args []string) error {
	if err := helper.ValidateCmdArgs(cmd, args, "o-auth-client"); err != nil {
		return err
	}

	fArg := args[0]
	switch fArg {
	case "list":
		if err := helper.ValidateCmdArgAndFlag(cmd, args, "o-auth-client", fArg, "organization"); err != nil {
			return err
		}
	case "create":
		if err := helper.ValidateCmdArgAndFlag(cmd, args, "o-auth-client", fArg, "organization"); err != nil {
			return err
		}

		if err := helper.ValidateCmdArgAndFlag(cmd, args, "o-auth-client", fArg, "api-url"); err != nil {
			return err
		}

		if err := helper.ValidateCmdArgAndFlag(cmd, args, "o-auth-client", fArg, "http-url"); err != nil {
			return err
		}

		if err := helper.ValidateCmdArgAndFlag(cmd, args, "o-auth-client", fArg, "o-auth-token"); err != nil {
			return err
		}

		if err := helper.ValidateCmdArgAndFlag(cmd, args, "o-auth-client", fArg, "service-provider"); err != nil {
			return err
		}
	case "read", "delete":
		if err := helper.ValidateCmdArgAndFlag(cmd, args, "o-auth-client", fArg, "id"); err != nil {
			return err
		}
	}

	return nil
}

func oAuthClientRun(cmd *cobra.Command, args []string) error {

	token := dao.GetTeamToken(profile)
	client := aid.GetTFEClient(token)

	fArg := args[0]
	switch fArg {
	case "list":
		list, err := oAuthClientList(client)
		if err == nil {
			aid.PrintOAuthClientList(list)
		} else {
			return fmt.Errorf("no o-auth-clients was found")
		}

	case "create":
		options := aid.GetOAuthClientCreateOptions(cmd)
		oAuthClient, err := oAuthClientCreate(client, options)

		if err == nil && oAuthClient.ID != "" {
			fmt.Println(aid.ToJSON(oAuthClient))
		} else {
			return fmt.Errorf("unable to create o-auth-client\n%v", err)
		}
	case "read":
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return fmt.Errorf("unable to get flag id\n%v", err)
		}

		oAuthClient, err := oAuthClientRead(client, id)
		if err == nil {
			fmt.Println(aid.ToJSON(oAuthClient))
		} else {
			return fmt.Errorf("o-auth-client %s not found\n%v", id, err)
		}
	case "delete":
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return fmt.Errorf("unable to get flag id\n%v", err)
		}

		err = oAuthClientDelete(client, id)
		if err == nil {
			fmt.Printf("o-auth-client %s deleted successfully\n", id)
		} else {
			return fmt.Errorf("unable to delete o-auth-client %s\n%v", id, err)
		}
	}

	return nil
}

func oAuthClientList(client *tfe.Client) (*tfe.OAuthClientList, error) {
	return client.OAuthClients.List(context.Background(), organization, tfe.OAuthClientListOptions{})
}

// Create is used to create a new oAuthClient.
func oAuthClientCreate(client *tfe.Client, options tfe.OAuthClientCreateOptions) (*tfe.OAuthClient, error) {
	return client.OAuthClients.Create(context.Background(), organization, options)
}

// Read an OAuth client by its ID.

// Read a oAuthClient by its name.
func oAuthClientRead(client *tfe.Client, oAuthClientID string) (*tfe.OAuthClient, error) {
	return client.OAuthClients.Read(context.Background(), oAuthClientID)
}

// // Delete a oAuthClient by its name.
func oAuthClientDelete(client *tfe.Client, oAuthClientID string) error {
	return client.OAuthClients.Delete(context.Background(), oAuthClientID)
}
