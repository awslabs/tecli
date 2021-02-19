package aid

import (
	"fmt"

	tfe "github.com/hashicorp/go-tfe"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// SetVariableFlags define flags for the cobra command
func SetVariableFlags(cmd *cobra.Command) {
	usage := "The variable ID"
	cmd.Flags().String("id", "", usage)

	usage = "The workspace ID."
	cmd.Flags().String("workspace-id", "", usage)

	usage = "The name of the variable."
	cmd.Flags().String("key", "", usage)

	usage = "The value of the variable."
	cmd.Flags().String("value", "", usage)

	usage = "The description of the variable."
	cmd.Flags().String("description", "", usage)

	usage = "Whether this is a Terraform or environment variable. Valid values: env, policy-set or terraform. Once defined, cannot be modified."
	cmd.Flags().String("category", "", usage)

	usage = "Whether to evaluate the value of the variable as a string of HCL code."
	cmd.Flags().Bool("hcl", false, usage)

	usage = "Whether the value is sensitive."
	cmd.Flags().Bool("sensitive", false, usage)

}

// GetVariableCreateOptions return tfe.VariableCreateOptions with correpondent values given by the flags
func GetVariableCreateOptions(cmd *cobra.Command) tfe.VariableCreateOptions {
	var options tfe.VariableCreateOptions

	if cmd.Flags().Changed("key") {
		// The name of the variable.
		key, err := cmd.Flags().GetString("key")
		if err != nil {
			logrus.Fatalf("unable to get flag key\n%v", err)
		}

		options.Key = &key
	}

	if cmd.Flags().Changed("value") {
		// The value of the variable.
		value, err := cmd.Flags().GetString("value")
		if err != nil {
			logrus.Fatalf("unable to get flag value\n%v", err)
		}

		options.Value = &value
	}

	if cmd.Flags().Changed("description") {
		// The description of the variable.
		description, err := cmd.Flags().GetString("description")
		if err != nil {
			logrus.Fatalf("unable to get flag description\n%v", err)
		}

		options.Description = &description
	}

	if cmd.Flags().Changed("category") {
		category, err := cmd.Flags().GetString("category")
		if err != nil {
			logrus.Fatalf("unable to get flag category\n%v", err)
		}

		switch category {
		case "env":
			options.Category = tfe.Category(tfe.CategoryEnv)
		case "policy-set":
			options.Category = tfe.Category(tfe.CategoryPolicySet)
		case "terraform":
			options.Category = tfe.Category(tfe.CategoryTerraform)
		}
	}

	if cmd.Flags().Changed("hcl") {
		// Whether to evaluate the value of the variable as a string of HCL code.
		hcl, err := cmd.Flags().GetBool("hcl")
		if err != nil {
			logrus.Fatalf("unable to get flag hcl\n%v", err)
		}

		options.HCL = &hcl
	}

	if cmd.Flags().Changed("sensitive") {
		// Whether the value is sensitive.
		sensitive, err := cmd.Flags().GetBool("sensitive")
		if err != nil {
			logrus.Fatalf("unable to get flag sensitive\n%v", err)
		}

		options.Sensitive = &sensitive
	}

	return options
}

// GetVariableUpdateOptions return tfe.VariableUpdateOptions with correpondent values given by the flags
func GetVariableUpdateOptions(cmd *cobra.Command) tfe.VariableUpdateOptions {
	var options tfe.VariableUpdateOptions

	if cmd.Flags().Changed("id") {
		// The name of the variable.
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			logrus.Fatalf("unable to get flag id\n%v", err)
		}

		options.Key = &id
	}

	if cmd.Flags().Changed("key") {
		// The name of the variable.
		key, err := cmd.Flags().GetString("key")
		if err != nil {
			logrus.Fatalf("unable to get flag key\n%v", err)
		}

		options.Key = &key
	}

	if cmd.Flags().Changed("value") {
		// The value of the variable.
		value, err := cmd.Flags().GetString("value")
		if err != nil {
			logrus.Fatalf("unable to get flag value\n%v", err)
		}

		options.Value = &value
	}

	if cmd.Flags().Changed("description") {
		// The description of the variable.
		description, err := cmd.Flags().GetString("description")
		if err != nil {
			logrus.Fatalf("unable to get flag description\n%v", err)
		}

		options.Description = &description
	}

	if cmd.Flags().Changed("hcl") {
		// Whether to evaluate the value of the variable as a string of HCL code.
		hcl, err := cmd.Flags().GetBool("hcl")
		if err != nil {
			logrus.Fatalf("unable to get flag hcl\n%v", err)
		}

		options.HCL = &hcl
	}

	if cmd.Flags().Changed("sensitive") {
		// Whether the value is sensitive.
		sensitive, err := cmd.Flags().GetBool("sensitive")
		if err != nil {
			logrus.Fatalf("unable to get flag sensitive\n%v", err)
		}

		options.Sensitive = &sensitive
	}

	return options
}

// PrintVariableList convert struct to JSON and displays to user
func PrintVariableList(list *tfe.VariableList) {
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
