package tests

import (
	"testing"

	"github.com/awslabs/tfe-cli/cobra/controller"
	"github.com/stretchr/testify/assert"
)

func TestVersionCmd(t *testing.T) {
	args := []string{"version"}
	cmd := controller.VersionCmd()
	out, err := executeCommand(t, cmd, args)
	assert.Nil(t, err)
	assert.Contains(t, out, "tfe-cli v")
}
