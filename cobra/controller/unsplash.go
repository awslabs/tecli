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
	"fmt"
	"os"

	"github.com/awslabs/tfe-cli/cobra/aid"
	"github.com/awslabs/tfe-cli/cobra/dao"
	"github.com/awslabs/tfe-cli/cobra/model"
	"github.com/awslabs/tfe-cli/helper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var unsplashPhotoSizes = []string{"all", "thumb", "small", "regular", "full", "raw"}

// UnsplashCmd command to download photos from Unsplash.com
func UnsplashCmd() *cobra.Command {
	man, err := helper.GetManual("unsplash")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:     man.Use,
		Short:   man.Short,
		Long:    man.Long,
		Example: man.Example,
		PreRunE: unsplashPreRun,
		RunE:    unsplashRun,
	}

	cmd.Flags().String("collections", "", "Public collection ID(‘s) to filter selection. If multiple, comma-separated")
	cmd.Flags().Bool("featured", false, "Limit selection to featured photos. Valid values: false, true.")
	cmd.Flags().String("filter", "low", "Limit results by content safety. Default: low. Valid values are low and high.")
	cmd.Flags().String("orientation", "landscape", "Filter by photo orientation. Valid values: landscape, portrait, squarish.")
	cmd.Flags().String("query", "mountains", "Limit selection to photos matching a search term.")
	cmd.Flags().String("size", "all", "Photos size. Valid values: all, thumb, small, regular, full, raw. Default: all")
	cmd.Flags().String("username", "", "Limit selection to a single user.")

	return cmd
}

func unsplashPreRun(cmd *cobra.Command, args []string) error {
	logrus.Traceln("start: command unsplash pre-run")

	params := aid.GetModelFromFlags(cmd)
	if !helper.ContainsString(unsplashPhotoSizes, params.Size) {
		return fmt.Errorf("unknown photo size provided: %s", params.Size)
	}

	logrus.Traceln("end: command unsplash pre-run")

	return nil
}

func unsplashRun(cmd *cobra.Command, args []string) error {
	logrus.Traceln("start: command unsplash run")

	params := aid.GetModelFromFlags(cmd)
	cred, err := dao.GetCredentialByProvider(profile, "unsplash")
	if err != nil {
		logrus.Errorf("Unexpected error: %v", err)
		return err
	}

	if (model.Credential{}) == cred {
		return fmt.Errorf("no unsplash credential found or no profile enabled")
	}

	err = aid.DownloadPhoto(params, cred, unsplashPhotoSizes)
	if err != nil {
		logrus.Errorf("unable to download photo\n%v", err)
		return err
	}

	logrus.Traceln("end: command unsplash run")
	return err
}
