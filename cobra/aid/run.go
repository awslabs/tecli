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

// GetRunCreateOptions TODO ..
func GetRunCreateOptions(cmd *cobra.Command) tfe.RunCreateOptions {
	var options tfe.RunCreateOptions

	// Specifies if this plan is a destroy plan, which will destroy all
	// provisioned resources.
	isDestroy, err := cmd.Flags().GetBool("is-destroy")
	if err != nil {
		logrus.Fatalf("unable to get flag is-destroy\n%v", err)
	} else {
		options.IsDestroy = &isDestroy
	}

	// Specifies the message to be associated with this run.
	message, err := cmd.Flags().GetString("message")
	if err != nil {
		logrus.Fatalf("unable to get flag message\n%v", err)
	}
	if message != "" {
		options.Message = &message
	}

	configurationVersion := GetConfigurationVersionFlags(cmd)
	options.ConfigurationVersion = &configurationVersion

	workspace := GetWorspaceFlags(cmd)
	options.Workspace = &workspace

	// If non-empty, requests that Terraform should create a plan including
	// actions only for the given objects (specified using resource address
	// syntax) and the objects they depend on.
	//
	// This capability is provided for exceptional circumstances only, such as
	// recovering from mistakes or working around existing Terraform
	// limitations. Terraform will generally mention the -target command line
	// option in its error messages describing situations where setting this
	// argument may be appropriate. This argument should not be used as part
	// of routine workflow and Terraform will emit warnings reminding about
	// this whenever this property is set.
	targetAddrs, err := cmd.Flags().GetStringArray("target-addrs")
	if err != nil {
		logrus.Fatalf("unable to get flag target-addrs\n%v", err)
	}
	if len(targetAddrs) > 0 {
		options.TargetAddrs = targetAddrs
	}

	return options

}

// GetConfigurationVersionFlags TODO ...
func GetConfigurationVersionFlags(cmd *cobra.Command) tfe.ConfigurationVersion {
	var configurationVersion tfe.ConfigurationVersion

	// Specifies the configuration version to use for this run. If the
	// configuration version object is omitted, the run will be created using the
	// workspace's latest configuration version.
	configurationVersionID, err := cmd.Flags().GetString("configuration-version-id")
	if err != nil {
		logrus.Fatalf("unable to get flag configuration-version-id\n%v", err)
	}
	if configurationVersionID != "" {
		configurationVersion.ID = configurationVersionID
	}

	configurationVersionAutoQueueRuns, err := cmd.Flags().GetBool("configuration-version-auto-queue-runs")
	if err != nil {
		logrus.Fatalf("unable to get flag configuration-version-auto-queue-runs\n%v", err)
	} else {
		configurationVersion.AutoQueueRuns = configurationVersionAutoQueueRuns
	}

	configurationVersionError, err := cmd.Flags().GetString("configuration-version-error")
	if err != nil {
		logrus.Fatalf("unable to get flag configuration-version-error\n%v", err)
	}
	if configurationVersionError != "" {
		configurationVersion.Error = configurationVersionError
	}

	configurationVersionErrorMessage, err := cmd.Flags().GetString("configuration-version-error-message")
	if err != nil {
		logrus.Fatalf("unable to get flag configuration-version-error-message\n%v", err)
	}
	if configurationVersionErrorMessage != "" {
		configurationVersion.ErrorMessage = configurationVersionErrorMessage
	}

	configurationVersionSpeculative, err := cmd.Flags().GetBool("configuration-version-speculative")
	if err != nil {
		logrus.Fatalf("unable to get flag configuration-version-speculative\n%v", err)
	} else {
		configurationVersion.Speculative = configurationVersionSpeculative
	}

	configurationVersionUploadURL, err := cmd.Flags().GetString("configuration-version-upload-url")
	if err != nil {
		logrus.Fatalf("unable to get flag configuration-version-upload-url\n%v", err)
	}
	if configurationVersionUploadURL != "" {
		configurationVersion.UploadURL = configurationVersionUploadURL
	}

	return configurationVersion
}

// // GetRunUpdateOptions TODO ...
// func GetRunUpdateOptions(cmd *cobra.Command) tfe.RunUpdateOptions {
// 	var options tfe.RunUpdateOptions

// 	return options

// }
