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

package controller

import (
	"errors"
	"fmt"
	"os"

	"github.com/awslabs/tfe-cli/cobra/aid"
	"github.com/awslabs/tfe-cli/cobra/dao"
	"github.com/awslabs/tfe-cli/helper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var initValidArgs = []string{"project"}
var initValidProjectTypes = []string{"basic", "cloud", "cloudformation", "terraform"}

// InitCmd command to initialize projects
func InitCmd() *cobra.Command {
	man, err := helper.GetManual("init")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		Example:   man.Example,
		ValidArgs: initValidArgs,
		Args:      cobra.OnlyValidArgs,
		PreRunE:   initPreRun,
		RunE:      initRun,
	}

	cmd.Flags().String("project-name", "", "The project name.")
	cmd.Flags().String("project-type", "basic", "The project type.")
	cmd.MarkFlagRequired("name")

	return cmd
}

func initPreRun(cmd *cobra.Command, args []string) error {
	logrus.Traceln("start: command init pre-run")

	if err := helper.ValidateCmdArgs(cmd, args, "init"); err != nil {
		return err
	}

	if err := helper.ValidateCmdArgAndFlag(cmd, args, "init", "project", "project-name"); err != nil {
		return err
	}

	if err := helper.ValidateCmdArgAndFlag(cmd, args, "init", "project", "project-type"); err != nil {
		return err
	}

	logrus.Traceln("end: command init pre-run")
	return nil
}

func initRun(cmd *cobra.Command, args []string) error {
	logrus.Traceln("start: command init run")

	cmd.SilenceUsage = true

	if args[0] == "project" {

		pName, err := cmd.Flags().GetString("project-name")
		if err != nil {
			logrus.Errorf("unable to access --project-name\n%v", err)
			return fmt.Errorf("unable to access --project-name\n%v", err)
		}

		pType, err := cmd.Flags().GetString("project-type")
		if err != nil {
			logrus.Errorf("unable to access --project-type\n%v", err)
			return fmt.Errorf("unable to access --project-type\n%v", err)
		}

		switch pType {
		case "basic":
			err = aid.CreateBasicProject(cmd, pName)
		case "cloud":
			err = aid.CreateCloudProject(cmd, pName)
		case "cloudformation":
			err = aid.CreateCloudFormationProject(cmd, pName)
		case "terraform":
			err = aid.CreateTerraformProject(cmd, pName)
		default:
			return errors.New("unknow project type")
		}

		if aid.ConfigurationsDirectoryExist() && aid.ConfigurationsFileExist() {
			config, err := dao.GetConfigurations()
			if err != nil {
				logrus.Errorf("unable to get configuration during initialization\n%v", err)
				return fmt.Errorf("unable to initialize project based on configurations\n%v", err)
			}

			created := aid.InitCustomized(profile, config)
			if !created {
				return fmt.Errorf("unable to initialize project based on configurations\n%s", err)
			}

		}

		if err != nil {
			return fmt.Errorf("unable to initialize project sucessfully \n%s", err)
		}

		cmd.Printf("%s was successfully initialized as a %s project\n", pName, pType)

	}

	logrus.Traceln("end: command init run")
	return nil
}
