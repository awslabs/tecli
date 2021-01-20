package tests

import (
	"os"
	"testing"

	"github.com/awslabs/tfe-cli/cobra/aid"
	"github.com/awslabs/tfe-cli/cobra/controller"
	"github.com/awslabs/tfe-cli/helper"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestRenderCmd(t *testing.T) {
	tests := map[string]struct {
		args []string
		out  string
		err  string
	}{
		// argument
		"no arg":                      {args: []string{"render"}, out: "", err: "this command requires one argument"},
		"empty arg":                   {args: []string{"render", ""}, out: "", err: "invalid argument"},
		"wrong arg":                   {args: []string{"render", "foo"}, out: "", err: "invalid argument"},
		"unknown flag":                {args: []string{"render", "--foo"}, out: "", err: "unknown flag: --foo"},
		"missing database":            {args: []string{"render", "template"}, out: "", err: "missing database"},
		"flag needs an argument name": {args: []string{"render", "template", "--name"}, out: "", err: "flag needs an argument: --name"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			out, err := executeCommand(t, controller.RenderCmd(), tc.args)
			assert.Contains(t, out, tc.out)
			assert.Contains(t, err.Error(), tc.err)
		})
	}
}

// /* README */

func initProject(t *testing.T, pType string) {
	aid.DeleteConfigurationsDirectory()
	args := []string{"init", "project", "--project-name", "foo", "--project-type", pType}
	wd, out, err := executeCommandGetWorkingDirectory(t, controller.InitCmd(), args)

	assert.Nil(t, err)
	assert.Contains(t, out, "was successfully initialized")

	if err := os.Chdir(helper.BuildPath(wd + "/" + t.Name() + "/" + "foo")); err != nil {
		logrus.Fatal("unable to change current working directory")
	}

}

func TestRenderDefault(t *testing.T) {
	initProject(t, "basic")

	args := []string{"render", "template"}
	out, err := executeCommandOnly(t, controller.RenderCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "Template readme.tmpl rendered as README.md")
}

func TestRenderReadme(t *testing.T) {
	initProject(t, "basic")

	args := []string{"render", "template", "--name", "readme"}
	out, err := executeCommandOnly(t, controller.RenderCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "Template readme.tmpl rendered as README.md")
}

func TestRenderHLD(t *testing.T) {
	initProject(t, "cloud")

	args := []string{"render", "template", "--name", "hld"}
	out, err := executeCommandOnly(t, controller.RenderCmd(), args)
	assert.Nil(t, err)
	assert.Contains(t, out, "Template hld.tmpl rendered as HLD.md")
}
