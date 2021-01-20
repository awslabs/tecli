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
	"strings"

	"github.com/awslabs/tfe-cli/cobra/aid"
	"github.com/awslabs/tfe-cli/cobra/dao"
	"github.com/awslabs/tfe-cli/cobra/model"
	"github.com/awslabs/tfe-cli/helper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var renderValidArgs = []string{"template"}

// RenderCmd command to render templates
func RenderCmd() *cobra.Command {
	man, err := helper.GetManual("render")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:       man.Use,
		Short:     man.Short,
		Long:      man.Long,
		ValidArgs: renderValidArgs,
		Args:      cobra.OnlyValidArgs,
		PreRunE:   renderPreRun,
		RunE:      renderRun,
	}

	cmd.Flags().StringP("name", "n", "readme", "Database file name of the template to be rendered (it must be under tfe-cli/ directory.")

	return cmd
}

func renderPreRun(cmd *cobra.Command, args []string) error {
	logrus.Traceln("start: command render pre-run")

	if err := helper.ValidateCmdArgs(cmd, args, "render"); err != nil {
		return err
	}

	if err := helper.ValidateCmdArgAndFlag(cmd, args, "render", "template", "name"); err != nil {
		return err
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		logrus.Errorf("error: unable to access flag name\n%v", err)
		return fmt.Errorf("unable to access flag name\n%v", err)
	}

	path := helper.BuildPath("tfe-cli/" + name + ".yaml")
	if !helper.FileExists(path) {
		logrus.Errorf("missing database " + path)
		return errors.New("missing database " + path)
	}

	path = helper.BuildPath("tfe-cli/" + name + ".tmpl")
	if !helper.FileExists(path) {
		logrus.Errorf("missing template " + path)
		return errors.New("missing template " + path)
	}

	logrus.Traceln("end: command render pre-run")
	return nil
}

func renderRun(cmd *cobra.Command, args []string) error {
	logrus.Traceln("start: command render run")

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		logrus.Errorf("error: unable to render template "+name+"\n%v", err)
		return fmt.Errorf("unable to render template "+name+"\n%v", err)
	}

	if err := updateLogo(profile); err != nil {
		logrus.Errorf("Unexpected error: %v", err)
		return fmt.Errorf("unable to update logo url\n%v", err)
	}

	if err := aid.BuildTemplate(name); err != nil {
		logrus.Errorf("Unexpected error: %v", err)
		return fmt.Errorf("unable to render template "+name+"\n%v", err)
	}

	cmd.Println("Template " + name + ".tmpl rendered as " + strings.ToUpper(name) + ".md.")

	logrus.Traceln("end: command render run")
	return nil
}

func updateLogo(profile string) error {

	if !updateLogoFromUnsplashFile() {
		return updateLogoFromConfigurations(profile)
	}

	return nil
}

func updateLogoFromUnsplashFile() bool {
	if helper.FileExists("unsplash.yaml") {
		configPath, _ := os.Getwd()
		configName := "unsplash"
		configType := "yaml"

		var response model.UnsplashRandomPhotoResponse

		v, err := aid.ReadConfigAsViper(configPath, configName, configType)
		if err != nil {
			logrus.Errorf("unable to read unsplash.yaml as viper object\n%v", err)
			return false
		}

		err = v.Unmarshal(&response)
		if err != nil {
			logrus.Errorf("unable to unmarshall unsplash.yaml as unsplash response\n%v", err)
			return false
		}

		readMe, err := dao.GetReadMe()
		if err != nil {
			logrus.Errorf("Unable to get local readme config\n%v", err)
			return false
		}

		err = aid.UpdateReadMeLogoURL(readMe, response)
		if err != nil {
			logrus.Errorf("unable to update logo URL\n%s", err)
			return false
		}

		return true
	}

	return false
}

func updateLogoFromConfigurations(profile string) error {
	if aid.ConfigurationsDirectoryExist() {
		if aid.CredentialsFileExist() && aid.ConfigurationsFileExist() {

			// ignore error, as credentials doesn't exist
			cred, err := dao.GetCredentialByProvider(profile, "unsplash")
			if err != nil {
				logrus.Warnf("no unsplash credential found\n%v", err)
				return nil
			}

			if cred.AccessKey != "" && cred.SecretKey != "" {
				readMe, err := dao.GetReadMe()
				if err != nil {
					return fmt.Errorf("Unable to get local readme config\n%v", err)
				}

				params := dao.GetUnsplashRandomPhotoParameters(profile)
				if (model.UnsplashRandomPhotoParameters{}) == params {
					logrus.Warnf("no unsplash random photo parameters configuration found or enabled\n%v", err)
					return nil
				}

				response, err := aid.RequestRandomPhoto(params, cred)
				if err != nil {
					logrus.Warnf("unable to fetch response from unsplash during render command\n%v", err)
					return err
				}

				err = aid.UpdateReadMeLogoURL(readMe, response)
				if err != nil {
					return fmt.Errorf("unable to update logo URL\n%s", err)
				}
			}
		}
	}

	return nil
}
