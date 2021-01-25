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
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/aid"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/dao"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/helper"
)

var plansValidArgs = []string{"delete"}

// PlansCmd command to display tecli current version
func PlansCmd() *cobra.Command {
	man, err := helper.GetManual("plans")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		Example:   man.Example,
		ValidArgs: plansValidArgs,
		Args:      cobra.OnlyValidArgs,
		RunE:      plansRun,
	}

	return cmd
}

func plansRun(cmd *cobra.Command, args []string) error {
	token := dao.GetTeamToken(profile)
	client := aid.GetTFEClient(token)

	orgs, err := client.Organizations.List(context.Background(), tfe.OrganizationListOptions{})
	if err != nil {
		logrus.Fatal(err)
	}

	for i, item := range orgs.Items {
		fmt.Printf("%d, %s", i, item.Email)
	}

	return nil
}
