package tests

import (
	"testing"

	"github.com/awslabs/tfe-cli/cobra/controller"
	"github.com/stretchr/testify/assert"
)

func TestGitIgnoreCmd(t *testing.T) {
	tests := map[string]struct {
		args []string
		out  string
		err  string
	}{
		// argument
		"empty":     {args: []string{"gitignore"}, out: "", err: "no flag or argument provided"},
		"empty arg": {args: []string{"gitignore", ""}, out: "", err: "unknown argument provided"},
		"wrong arg": {args: []string{"gitignore", "foo"}, out: "", err: "unknown argument provided"},

		// // flags
		"wrong flag": {args: []string{"gitignore", "--input"}, out: "", err: "flag needs an argument: --input"},

		// // # projects
		"no project name": {args: []string{"gitignore", "--input", ""}, out: "", err: "no flag or argument provided"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			out, err := executeCommand(t, controller.GitIgnoreCmd(), tc.args)
			assert.Contains(t, out, tc.out)
			assert.Contains(t, err.Error(), tc.err)
		})
	}
}

func TestGitIgnoreList(t *testing.T) {
	args := []string{"gitignore", "list"}
	out, err := executeCommand(t, controller.GitIgnoreCmd(), args)
	assert.Nil(t, err)
	assert.NotEmpty(t, out)
}
