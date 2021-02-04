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
	"gitlab.aws.dev/devops-aws/tecli/helper"
)

var profile string

var organization string

// RootCmd represents the base command when called without any subcommands
func RootCmd() *cobra.Command {
	man, err := helper.GetManual("root", []string{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Short: man.Short,
		Long:  man.Long,
	}

	// Cobra supports persistent flags, which, if defined here will be global for your application.
	cmd.PersistentFlags().StringVarP(&profile, "profile", "p", "default", "Use a specific profile from your credentials and configurations file.")
	cmd.PersistentFlags().StringVarP(&organization, "organization", "o", "", "Terraform Cloud Organization name")

	return cmd
}
