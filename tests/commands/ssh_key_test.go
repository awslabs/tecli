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
	"log"
	"testing"

	"github.com/awslabs/tecli/cobra/controller"
	"github.com/awslabs/tecli/helper"
	"github.com/stretchr/testify/assert"
)

// func TestSSHKeyList(t *testing.T) {
// 	args := []string{"ssh-key", "list", "--organization", "tecli-test-org"}
// 	out, err := executeCommandOnly(t, controller.SSHKeyCmd(), args)
// 	assert.Nil(t, err)
// 	assert.Contains(t, out, "\"ID\": \"sshkey-")
// }

func TestSSHKeyCreate(t *testing.T) {
	privateKey, err := helper.GeneratePrivateSSHKey(4096)
	if err != nil {
		log.Fatal(err.Error())
	}

	privateKeyBytes := helper.EncodePrivateSSHKeyToPEM(privateKey)

	args := []string{"ssh-key", "create", "--organization", "tecli-test-org", "--name", "foo", "--value", string(privateKeyBytes)}
	out, err := executeCommandOnly(t, controller.SSHKeyCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "")
}

// func TestSSHKeyRead(t *testing.T) {
// 	args := []string{"ssh-key", "read", "--organization", "tecli-test-org", "--id", "sshkey-BmNjgGuyA8sP7NUK"}
// 	out, err := executeCommandOnly(t, controller.SSHKeyCmd(), args)
// 	assert.Nil(t, err)
// 	assert.Contains(t, out, "")
// }

// func TestSSHKeyUpdate(t *testing.T) {
// 	args := []string{"ssh-key", "update", "--name", "bar", "--value", tempSSHPrivateKey, "--id", "sshkey-BmNjgGuyA8sP7NUK"}
// 	out, err := executeCommandOnly(t, controller.SSHKeyCmd(), args)
// 	assert.Nil(t, err)
// 	assert.Contains(t, out, "")
// }

// func TestSSHKeyDelete(t *testing.T) {
// 	args := []string{"ssh-key", "delete", "--id", "sshkey-BmNjgGuyA8sP7NUK"}
// 	out, err := executeCommandOnly(t, controller.SSHKeyCmd(), args)
// 	assert.Nil(t, err)
// 	assert.Contains(t, out, "")
// }
