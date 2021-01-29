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
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/hashicorp/go-tfe"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/aid"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/dao"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/helper"
)

var planValidArgs = []string{"read", "logs"}

// PlanCmd command to display tecli current version
func PlanCmd() *cobra.Command {
	man, err := helper.GetManual("plan", planValidArgs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		Example:   man.Example,
		ValidArgs: planValidArgs,
		Args:      cobra.OnlyValidArgs,
		PreRunE:   planPreRun,
		RunE:      planRun,
	}

	aid.SetPlanFlags(cmd)

	return cmd
}

func planPreRun(cmd *cobra.Command, args []string) error {
	if err := helper.ValidateCmdArgs(cmd, args, "plan"); err != nil {
		return err
	}

	fArg := args[0]
	switch fArg {
	case "read", "logs":
		if err := helper.ValidateCmdArgAndFlag(cmd, args, "plan", fArg, "id"); err != nil {
			return err
		}
	}

	return nil
}

func planRun(cmd *cobra.Command, args []string) error {

	token := dao.GetTeamToken(profile)
	client := aid.GetTFEClient(token)

	fArg := args[0]
	switch fArg {
	case "read":
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return fmt.Errorf("unable to get flag id\n%v", err)
		}

		plan, err := planRead(client, id)
		if err == nil {
			fmt.Println(aid.ToJSON(plan))
		} else {
			return fmt.Errorf("plan %s not found\n%v", id, err)
		}
	case "logs":
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return fmt.Errorf("unable to get flag id\n%v", err)
		}

		logs, err := planLogs(client, id)
		if err != nil {
			logrus.Fatalf("unable to read plan logs\n%v", err)
		}
		fmt.Println(StreamToString(logs))
	}

	return nil
}

// Read a plan by its ID.
func planRead(client *tfe.Client, planID string) (*tfe.Plan, error) {
	return client.Plans.Read(context.Background(), planID)

}

// Logs retrieves the logs of a plan.
func planLogs(client *tfe.Client, planID string) (io.Reader, error) {
	return client.Plans.Logs(context.Background(), planID)

}

// StreamToByte converts io.Reader to []byte
func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

// StreamToString convert io.Reader to string
func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}
