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

var oAuthTokenValidArgs = []string{
	"list",
	"read",
	"update",
	"delete",
}

// OAuthTokenCmd command to display tecli current version
func OAuthTokenCmd() *cobra.Command {
	man, err := helper.GetManual("o-auth-token", oAuthTokenValidArgs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:          man.Use,
		Short:        man.Short,
		Long:         man.Long,
		Example:      man.Example,
		ValidArgs:    oAuthTokenValidArgs,
		Args:         cobra.OnlyValidArgs,
		PreRunE:      oAuthTokenPreRun,
		RunE:         oAuthTokenRun,
		SilenceUsage: true,
	}

	aid.SetOAuthTokenFlags(cmd)

	return cmd
}

func oAuthTokenPreRun(cmd *cobra.Command, args []string) error {
	if err := helper.ValidateCmdArgs(cmd, args, "o-auth-token"); err != nil {
		return err
	}

	fArg := args[0]
	switch fArg {
	case "read", "update", "delete":
		if err := helper.ValidateCmdArgAndFlag(cmd, args, "o-auth-token", fArg, "id"); err != nil {
			return err
		}
	}

	return nil
}

func oAuthTokenRun(cmd *cobra.Command, args []string) error {

	token := dao.GetOrganizationToken(profile)
	client := aid.GetTFEClient(token)

	fArg := args[0]
	switch fArg {
	case "list":
		organization := dao.GetOrganization(profile)
		list, err := oAuthTokenList(client, organization)
		if err == nil {
			aid.PrintOAuthTokenList(list)
		} else {
			return fmt.Errorf("no o-auth-tokens was found")
		}
	case "read":
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return fmt.Errorf("unable to get flag id\n%v", err)
		}

		oAuthToken, err := oAuthTokenRead(client, id)
		if err == nil {
			fmt.Println(aid.ToJSON(oAuthToken))
		} else {
			return fmt.Errorf("o-auth-token %s not found\n%v", id, err)
		}
	case "update":
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return fmt.Errorf("unable to get flag id\n%v", err)
		}

		options := aid.GetOAuthTokenUpdateOptions(cmd)
		oAuthToken, err := oAuthTokenUpdate(client, id, options)

		if err == nil && oAuthToken.ID != "" {
			fmt.Println(aid.ToJSON(oAuthToken))
		} else {
			return fmt.Errorf("unable to create o-auth-token\n%v", err)
		}
	case "delete":
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return fmt.Errorf("unable to get flag id\n%v", err)
		}

		err = oAuthTokenDelete(client, id)
		if err == nil {
			fmt.Printf("o-auth-token %s deleted successfully\n", id)
		} else {
			return fmt.Errorf("unable to delete o-auth-token %s\n%v", id, err)
		}
	}

	return nil
}

func oAuthTokenList(client *tfe.Client, organization string) (*tfe.OAuthTokenList, error) {
	return client.OAuthTokens.List(context.Background(), organization, tfe.OAuthTokenListOptions{})
}

// Read an OAuth client by its ID.
func oAuthTokenRead(client *tfe.Client, oAuthTokenID string) (*tfe.OAuthToken, error) {
	return client.OAuthTokens.Read(context.Background(), oAuthTokenID)
}

// Create is used to create a new oAuthToken.
func oAuthTokenUpdate(client *tfe.Client, oAuthTokenID string, options tfe.OAuthTokenUpdateOptions) (*tfe.OAuthToken, error) {
	return client.OAuthTokens.Update(context.Background(), oAuthTokenID, options)
}

// // Delete a oAuthToken by its name.
func oAuthTokenDelete(client *tfe.Client, oAuthTokenID string) error {
	return client.OAuthTokens.Delete(context.Background(), oAuthTokenID)
}
