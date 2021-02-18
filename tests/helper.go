package tests

import (
	"bytes"
	"os"
	"testing"

	"github.com/awslabs/tecli/cobra/controller"
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
