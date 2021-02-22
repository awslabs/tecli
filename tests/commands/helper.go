// Package commands contains tests for Cobra commands
package commands

import (
	"bytes"
	"os"
	"testing"

	"github.com/awslabs/tecli/cobra/controller"
	tfe "github.com/hashicorp/go-tfe"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

/* COBRA */

func executeCommand(t *testing.T, cmd *cobra.Command, args []string) (stdout string, err error) {
	wd := t.TempDir()
	// execute the command within the new working directory
	os.Chdir(wd)
	_, stdout, err = executeCommandC(cmd, args)

	return stdout, err
}

// executeCommandOnly only executes the given command without changing the current directory, useful for combined tests
func executeCommandOnly(t *testing.T, cmd *cobra.Command, args []string) (stdout string, err error) {
	_, stdout, err = executeCommandC(cmd, args)
	return stdout, err
}

func executeCommandC(cmd *cobra.Command, args []string) (command *cobra.Command, stdout string, err error) {
	buf := new(bytes.Buffer)

	rootCmd := controller.RootCmd()
	rootCmd.AddCommand(cmd)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	// args = append(args, "--verbosity", "trace")
	rootCmd.SetArgs(args)

	command, err = rootCmd.ExecuteC()
	stdout = buf.String()
	return command, stdout, err
}

/* TFE */

// GetTFEClient returns a new terraform api client given a token
func GetTFEClient() *tfe.Client {
	token := os.Getenv("TFC_TEAM_TOKEN")

	config := &tfe.Config{
		Token: token,
	}

	client, err := tfe.NewClient(config)
	if err != nil {
		logrus.Errorln("unable to get terraform cloud api client")
		logrus.Fatalln(err)
	}

	return client
}

