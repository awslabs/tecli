package aid

import (
	"fmt"
	"os"
	"strings"

	"github.com/awslabs/tfe-cli/cobra/model"
	"github.com/awslabs/tfe-cli/helper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

/* BASIC PROJECT */

// CreateBasicProject creates a basic project
func CreateBasicProject(cmd *cobra.Command, name string) error {
	err := createAndEnterProjectDir(name)
	if err != nil {
		return err
	}

	if initalized := initProject(); !initalized {
		logrus.Errorf("unable to initialize basic project")
		return fmt.Errorf("unable to initalize project %s", name)
	}

	return nil
}

func createAndEnterProjectDir(name string) error {

	if !helper.MkDirsIfNotExist(name) {
		return fmt.Errorf("unable to create directory %s", name)
	}

	err := os.Chdir(name)
	if err != nil {
		return fmt.Errorf("unable to enter directory %s", name)
	}

	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to returns a rooted path name corresponding to the current directory:\n%v", err)
	}
	logrus.Infof("current working directory changed to %s", wd)

	return nil
}

// create the basic configuration files
func initProject() bool {

	// Create a directory for tfe-cli
	a := helper.MkDirsIfNotExist("tfe-cli")
	b := helper.WriteFileFromBox("/init/tfe-cli/readme.yaml", "tfe-cli/readme.yaml")
	c := helper.WriteFileFromBox("/init/tfe-cli/readme.tmpl", "tfe-cli/readme.tmpl")
	d := helper.WriteFileFromBox("/init/.gitignore", ".gitignore")

	return (a && b && c && d)
}

// InitCustomized TODO...
func InitCustomized(profile string, config model.Configurations) bool {

	for _, p := range config.Profiles {
		if p.Name == profile && p.Enabled {
			for _, c := range p.Configurations {
				if c.Enabled && c.Initialization.Enabled {
					for _, f := range c.Initialization.Files {
						if f.State == "directory" {
							if !helper.MkDirsIfNotExist(f.Path) {
								logrus.Errorf("unable to create directory based on configuration")
								return false
							}
						} else if f.State == "file" {
							if strings.Contains(f.Src, "http") {
								if err := helper.DownloadFileTo(f.Src, f.Dest); err != nil {
									logrus.Errorf("unable to download file based on configuration")
									return false
								}
							} else {
								if err := helper.CopyFileTo(f.Src, f.Dest); err != nil {
									logrus.Errorf("unable to copy file based on configuration")
									return false
								}
							}
						}
					}
				}
			}
		}
	}

	return true
}

/* CLOUD PROJECT */

// CreateCloudProject copies the necessary templates for cloud projects
func CreateCloudProject(cmd *cobra.Command, name string) error {
	if err := CreateBasicProject(cmd, name); err != nil {
		return nil
	}

	if initialized := initCloudProject(); !initialized {
		logrus.Errorf("unable to initialize cloud project")
		return fmt.Errorf("unable to initalize project %s", name)
	}

	return nil
}

// copies the High Level Design template file
func initCloudProject() bool {
	a := helper.WriteFileFromBox("/init/tfe-cli/hld.yaml", "tfe-cli/hld.yaml")
	b := helper.WriteFileFromBox("/init/tfe-cli/hld.tmpl", "tfe-cli/hld.tmpl")

	return (a && b)
}

/* CLOUDFORMATION PROJECT */

// CreateCloudFormationProject creates an AWS CloudFormation project
func CreateCloudFormationProject(cmd *cobra.Command, name string) error {
	if err := CreateBasicProject(cmd, name); err != nil {
		return nil
	}

	if initialized := initCloudProject(); !initialized {
		logrus.Errorf("unable to initialize cloud project")
		return fmt.Errorf("unable to initalize project %s", name)
	}

	if initialized := initCloudFormationProject(); !initialized {
		logrus.Errorf("unable to initialize cloudformation project")
		return fmt.Errorf("unable to initalize project %s", name)
	}

	return nil
}

// initialize a project with CloudFormation structure and copies template files
func initCloudFormationProject() bool {

	a := helper.MkDirsIfNotExist("environments")
	b := helper.MkDirsIfNotExist("environments/dev")
	c := helper.MkDirsIfNotExist("environments/prod")
	d := helper.WriteFileFromBox("/init/project/type/clouformation/skeleton.yaml", "skeleton.yaml")
	e := helper.WriteFileFromBox("/init/project/type/clouformation/skeleton.json", "skeleton.json")

	/* TODO: copy a template to create standard tags for the entire stack easily
	https://docs.aws.amazon.com/cli/latest/reference/cloudformation/create-stack.html
	example aws cloudformation create-stack ... --tags */

	/* TODO: copy Makefile */
	/* TODO: copy LICENSE */

	return (a && b && c && d && e)
}

/* TERRAFORM PROJECT */

// CreateTerraformProject creates a HashiCorp Terraform project
func CreateTerraformProject(cmd *cobra.Command, name string) error {
	if err := CreateBasicProject(cmd, name); err != nil {
		return nil
	}

	if initialized := initCloudProject(); !initialized {
		logrus.Errorf("unable to initialize terraform project")
		return fmt.Errorf("unable to initalize project %s", name)
	}

	if initialized := initTerraformProject(); !initialized {
		logrus.Errorf("unable to initialize cloud project")
		return fmt.Errorf("unable to initalize project %s", name)
	}

	return nil
}

// InitTerraform initialize a project with Terraform structure
func initTerraformProject() bool {
	a := helper.WriteFileFromBox("/init/project/type/terraform/Makefile", "Makefile")
	b := helper.WriteFileFromBox("/init/project/type/terraform/LICENSE", "LICENSE")

	c := helper.MkDirsIfNotExist("environments")
	d := helper.WriteFileFromBox("/init/project/type/terraform/environments/dev.tf", "environments/dev.tf")
	e := helper.WriteFileFromBox("/init/project/type/terraform/environments/prod.tf", "environments/prod.tf")

	f := helper.WriteFileFromBox("/init/project/type/terraform/main.tf", "main.tf")
	g := helper.WriteFileFromBox("/init/project/type/terraform/variables.tf", "variables.tf")
	h := helper.WriteFileFromBox("/init/project/type/terraform/outputs.tf", "outputs.tf")

	return (a && b && c && d && e && f && g && h)

}

// TODO: allow users to inform additional files to be added to their project initialization
