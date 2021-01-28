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

package aid

import (
	tfe "github.com/hashicorp/go-tfe"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// SetRunFlags TODO ...
func SetRunFlags(cmd *cobra.Command) {
	usage := `A list of relations to include. See available resources: https://www.terraform.io/docs/cloud/api/run.html#available-related-resources`
	cmd.Flags().String("include", "", usage)

	usage = `The Run ID`
	cmd.Flags().String("id", "", usage)

	usage = `Specifies if this plan is a destroy plan, which will destroy all provisioned resources.`
	cmd.Flags().Bool("is-destroy", false, usage)

	usage = `Specifies the message to be associated with this run.`
	cmd.Flags().String("message", "", usage)

	usage = `Specifies the configuration version to use for this run. If the configuration version object is omitted, the run will be created using the workspace's latest configuration version.`
	cmd.Flags().String("configuration-version-id", "", usage)

	usage = `Specifies the workspace where the run will be executed.`
	cmd.Flags().String("workspace-id", "", usage)

	usage = `If non-empty, requests that Terraform should create a plan including actions only for the given objects (specified using resource address syntax) and the objects they depend on. This capability is provided for exceptional circumstances only, such as recovering from mistakes or working around existing Terraform limitations. Terraform will generally mention the -target command line option in its error messages describing situations where setting this argument may be appropriate. This argument should not be used as part of routine workflow and Terraform will emit warnings reminding about this whenever this property is set.`
	cmd.Flags().StringArray("target-addrs", []string{}, usage)

	usage = `An optional comment about the run.`
	cmd.Flags().String("comment", "", usage)

}

// GetRunCreateOptions TODO ..
func GetRunCreateOptions(cmd *cobra.Command) tfe.RunCreateOptions {
	var options tfe.RunCreateOptions

	// Specifies if this plan is a destroy plan, which will destroy all
	// provisioned resources.
	isDestroy, err := cmd.Flags().GetBool("is-destroy")
	if err != nil {
		logrus.Fatalf("unable to get flag is-destroy\n%v", err)
	}

	options.IsDestroy = &isDestroy

	// Specifies the message to be associated with this run.
	message, err := cmd.Flags().GetString("message")
	if err != nil {
		logrus.Fatalf("unable to get flag message\n%v", err)
	}
	if message != "" {
		options.Message = &message
	}

	targetAddrs, err := cmd.Flags().GetStringArray("target-addrs")
	if err != nil {
		logrus.Fatalf("unable to get flag target-addrs\n%v", err)
	}
	if len(targetAddrs) > 0 {
		options.TargetAddrs = targetAddrs
	}

	return options
}

// GetRunReadOptions TODO ...
func GetRunReadOptions(cmd *cobra.Command) tfe.RunReadOptions {
	var options tfe.RunReadOptions
	include, err := cmd.Flags().GetString("include")
	if err != nil {
		logrus.Fatalf("unable to get flag include\n%v", err)
	}

	if include != "" {
		options.Include = include
	}

	return options
}

// // GetRunUpdateOptions TODO ...
// func GetRunUpdateOptions(cmd *cobra.Command) tfe.RunUpdateOptions {
// 	var options tfe.RunUpdateOptions

// 	return options

// }
