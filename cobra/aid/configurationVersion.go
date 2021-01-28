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

// SetConfigurationVersionFlags TODO ...
func SetConfigurationVersionFlags(cmd *cobra.Command, withPrefix bool) {
	var prefix string
	if withPrefix {
		prefix = "configuration-version-"
	}
	// Specifies the configuration version to use for this run. If the
	// configuration version object is omitted, the run will be created using the
	// workspace's latest configuration version.

	usage := `The Configuration Version ID.`
	cmd.Flags().String(prefix+"id", "", usage)

	usage = `When true, runs are queued automatically when the configuration version is uploaded.`
	cmd.Flags().Bool(prefix+"auto-queue-runs", false, usage)

	usage = `When true, this configuration version can only be used for planning.`
	cmd.Flags().Bool(prefix+"speculative", false, usage)

	usage = `The Workspace ID`
	cmd.Flags().String("workspace-id", "", usage)

	// Upload packages and uploads Terraform configuration files. It requires
	// the upload URL from a configuration version and the full path to the
	// configuration files on disk.
	usage = `The upload url`
	cmd.Flags().String("url", "", usage)

	usage = `The upload path`
	cmd.Flags().String("path", "", usage)
}

// GetConfigurationVersionCreateOptions TODO ..
func GetConfigurationVersionCreateOptions(cmd *cobra.Command, withPrefix bool) tfe.ConfigurationVersionCreateOptions {
	var options tfe.ConfigurationVersionCreateOptions
	var prefix string

	if withPrefix {
		prefix = "configuration-version-"
	}

	autoQueueRuns, err := cmd.Flags().GetBool(prefix + "auto-queue-runs")
	if err != nil {
		logrus.Fatalf("unable to get flag configuration-version-auto-queue-runs\n%v", err)
	}

	options.AutoQueueRuns = &autoQueueRuns

	speculative, err := cmd.Flags().GetBool(prefix + "speculative")
	if err != nil {
		logrus.Fatalf("unable to get flag configuration-version-speculative\n%v", err)
	}

	options.Speculative = &speculative

	return options

}
