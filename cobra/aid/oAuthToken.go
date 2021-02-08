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
	"fmt"

	tfe "github.com/hashicorp/go-tfe"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// SetOAuthTokenFlags define flags for the cobra command
func SetOAuthTokenFlags(cmd *cobra.Command) {
	usage := `The OAuth Token ID.`
	cmd.Flags().String("id", "", usage)

	usage = `A private SSH key to be used for git clone operations.`
	cmd.Flags().String("private-ssh-key", "", usage)
}

// GetOAuthTokenUpdateOptions return options based on the flag values
func GetOAuthTokenUpdateOptions(cmd *cobra.Command) tfe.OAuthTokenUpdateOptions {
	var options tfe.OAuthTokenUpdateOptions

	// A private SSH key to be used for git clone operations.
	privateSSHKey, err := cmd.Flags().GetString("private-ssh-key")
	if err != nil {
		logrus.Fatalf("unable to get flag private-ssh-key")
	}

	if privateSSHKey != "" {
		options.PrivateSSHKey = &privateSSHKey
	}

	return options

}

// PrintOAuthTokenList TODO ...
func PrintOAuthTokenList(list *tfe.OAuthTokenList) {
	if len(list.Items) > 0 {
		for i, item := range list.Items {
			if i < len(list.Items)-1 {
				fmt.Printf("%v,\n", ToJSON(item))
			} else {
				fmt.Printf("%v\n", ToJSON(item))
			}
		}
	}
}
