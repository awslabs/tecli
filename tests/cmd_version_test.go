package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.aws.dev/devops-aws/tecli/cobra/controller"
)

func TestVersionCmd(t *testing.T) {
	args := []string{"version"}
	cmd := controller.VersionCmd()
	out, err := executeCommand(t, cmd, args)
	assert.Nil(t, err)
	assert.Contains(t, out, "tecli v")
}
