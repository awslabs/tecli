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
	"encoding/json"
	"fmt"
	"os"

	"github.com/hashicorp/go-tfe"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/aid"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/dao"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/helper"
)

var oAuthValidArgs = []string{"list"}

// OAuthCmd command to display tecli current version
func OAuthCmd() *cobra.Command {
	man, err := helper.GetManual("oauth")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		Example:   man.Example,
		ValidArgs: oAuthValidArgs,
		Args:      cobra.OnlyValidArgs,
		PreRunE:   oAuthPreRun,
		RunE:      oAuthRun,
	}

	return cmd
}

func oAuthPreRun(cmd *cobra.Command, args []string) error {
	if err := helper.ValidateCmdArgs(cmd, args, "oauth"); err != nil {
		return err
	}
	return nil
}

func oAuthRun(cmd *cobra.Command, args []string) error {
	organization, _ := cmd.Flags().GetString("organization")
	token := dao.GetOrganizationToken(profile)
	client := aid.GetTFEClient(token)

	listOAuthClients(cmd, client, organization)
	// listOAuthToken(client, organization)

	return nil
}

func listOAuthClients(cmd *cobra.Command, client *tfe.Client, organization string) {
	clients, err := client.OAuthClients.List(context.Background(), organization, tfe.OAuthClientListOptions{})
	if err != nil {
		logrus.Fatalf("unable to list oauth clients\n%v", err)
	}

	if len(clients.Items) == 0 {
		cmd.Println("no oautch client found")
	} else {
		for _, item := range clients.Items {
			j, err := json.MarshalIndent(item, "", "  ")
			if err != nil {
				logrus.Fatalf(err.Error())
			}

			cmd.Printf("\n %s\n", string(j))
		}
	}
}

// func listOAuthToken(client *tfe.Client, organization string) {
// 	tokens, err := client.OAuthTokens.List(context.Background(), organization, tfe.OAuthTokenListOptions{})
// 	if err != nil {
// 		logrus.Fatal("unable to list oauth tokens\n%v", err)
// 	}

// 	for i, item := range tokens.Items {
// 		fmt.Printf("%d, %s", i, item.ID)
// 	}
// }
