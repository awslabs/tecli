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

package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/cobra/controller"
)

var tempSSHPrivateKey = `
-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----
`

func TestSSHKeyList(t *testing.T) {
	args := []string{"sshKey", "list", "--organization", "terraform-cloud-pipeline"}
	out, err := executeCommandOnly(t, controller.SSHKeyCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "")
}

func TestSSHKeyCreate(t *testing.T) {
	args := []string{"sshKey", "create", "--organization", "terraform-cloud-pipeline", "--name", "foo", "--value", tempSSHPrivateKey}
	out, err := executeCommandOnly(t, controller.SSHKeyCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "")
}

func TestSSHKeyRead(t *testing.T) {
	args := []string{"sshKey", "read", "--organization", "terraform-cloud-pipeline", "--id", "sshkey-BmNjgGuyA8sP7NUK"}
	out, err := executeCommandOnly(t, controller.SSHKeyCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "")
}

func TestSSHKeyDelete(t *testing.T) {
	args := []string{"sshKey", "delete", "--organization", "terraform-cloud-pipeline"}
	out, err := executeCommandOnly(t, controller.SSHKeyCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "")
}
