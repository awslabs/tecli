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

// GetSSHKeysCreateOptions  return options based on the flags values
func GetSSHKeysCreateOptions(cmd *cobra.Command) tfe.SSHKeyCreateOptions {
	var options tfe.SSHKeyCreateOptions
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		logrus.Fatalf("unable to get flag name\n%v\n", err)
	}

	if name != "" {
		options.Name = &name
	}

	value, err := cmd.Flags().GetString("value")
	if err != nil {
		logrus.Fatalf("unable to get flag value\n%v\n", err)
	}

	if value != "" {
		options.Value = &value
	}

	return options
}

// GetSSHKeyByName TODO ...
func GetSSHKeyByName(list *tfe.SSHKeyList, name string) tfe.SSHKey {
	for _, item := range list.Items {
		if item.Name == name {
			return *item
		}
	}

	return tfe.SSHKey{}
}

// GetSSHKeysUpdateOptions return options based on the flag values
func GetSSHKeysUpdateOptions(cmd *cobra.Command) tfe.SSHKeyUpdateOptions {
	var options tfe.SSHKeyUpdateOptions

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		logrus.Fatalf("unable to get flag name\n%v\n", err)
	}

	if name != "" {
		options.Name = &name
	}

	value, err := cmd.Flags().GetString("value")
	if err != nil {
		logrus.Fatalf("unable to get flag value\n%v\n", err)
	}

	if value != "" {
		options.Value = &value
	}

	return options

}
