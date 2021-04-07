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

package commands

import (
	"context"
	"log"
	"testing"

	"github.com/awslabs/tecli/cobra/controller"
	"github.com/awslabs/tecli/helper"
	tfe "github.com/hashicorp/go-tfe"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSSHKeyList(t *testing.T) {
	args := []string{"ssh-key", "list", "--organization", "tecli-test-org"}
	out, err := executeCommandOnly(t, controller.SSHKeyCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "\"ID\": \"sshkey-")
}

func getSSHKeyID() string {
	client := GetTFEClient()
	list, err := client.SSHKeys.List(context.Background(), "tecli-test-org", tfe.SSHKeyListOptions{})
	if err != nil {
		logrus.Fatalln(err)
	}

	for _, item := range list.Items {
		if item.Name == "foo" {
			return item.ID
		}
	}

	logrus.Fatalln("unable to find ssh key id")

	return ""
}

func TestSSHKeyCreate(t *testing.T) {
	privateKey, err := helper.GeneratePrivateSSHKey(4096)
	if err != nil {
		log.Fatal(err.Error())
	}

	privateKeyBytes := helper.EncodePrivateSSHKeyToPEM(privateKey)

	args := []string{"ssh-key", "create", "--organization", "tecli-test-org", "--name", "foo", "--value", string(privateKeyBytes)}
	out, err := executeCommandOnly(t, controller.SSHKeyCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "\"ID\": \"sshkey-")
	assert.Contains(t, out, "\"Name\": \"foo\"")
}

func TestSSHKeyRead(t *testing.T) {
	args := []string{"ssh-key", "read", "--organization", "tecli-test-org", "--id", getSSHKeyID()}
	out, err := executeCommandOnly(t, controller.SSHKeyCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "\"ID\": \"sshkey-")
}

// func TestSSHKeyUpdate(t *testing.T) {
// 	args := []string{"ssh-key", "update", "--name", "bar", "--value", tempSSHPrivateKey, "--id", getSSHKeyID()}
// 	out, err := executeCommandOnly(t, controller.SSHKeyCmd(), args)
// 	assert.Nil(t, err)
// 	assert.Contains(t, out, "")
// }

func TestSSHKeyDelete(t *testing.T) {
	args := []string{"ssh-key", "delete", "--id", getSSHKeyID()}
	out, err := executeCommandOnly(t, controller.SSHKeyCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "")
}
