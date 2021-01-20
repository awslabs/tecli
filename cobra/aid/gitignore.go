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

package aid

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/awslabs/tfe-cli/helper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// DownloadGitIgnore ..
func DownloadGitIgnore(cmd *cobra.Command, input string) (bool, error) {
	bytes, err := requestGitIgnore(cmd, input)
	if err != nil {
		logrus.Errorf("unable to download gitignore file\n%v", err)
		return false, err
	}

	if !saveGitIgnoreAsFile(bytes) {
		logrus.Errorln("unable to save gitignore API response as file")
		return false, fmt.Errorf("unable to create .gitignore file")
	}

	return true, nil
}

// GetGitIgnoreList ...
func GetGitIgnoreList() string {
	bytes, err := requestGitIgnoreList()
	if err != nil {
		logrus.Errorf("unable to fetch the gitignore list")
	}
	return string(bytes)
}

func requestGitIgnore(cmd *cobra.Command, input string) ([]byte, error) {
	url := fmt.Sprintf("https://www.toptal.com/developers/gitignore/api/%s", input)
	var response []byte

	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		logrus.Errorf("unexpected error while performing GET on Toptal API\n%v", err)
		return response, fmt.Errorf("unable to fetch gitignore API\n%v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		response, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return response, fmt.Errorf("unexpected error while reading Unsplash response \n%v", err)
		}

		return response, nil
	}

	return response, err
}

func requestGitIgnoreList() ([]byte, error) {
	url := fmt.Sprintf("https://www.toptal.com/developers/gitignore/api/list")
	var response []byte

	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		logrus.Errorf("unexpected error while performing GET on Toptal API\n%v", err)
		return response, fmt.Errorf("unable to fetch gitignore API\n%v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		response, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return response, fmt.Errorf("unexpected error while reading Unsplash response \n%v", err)
		}

		return response, nil
	}

	return response, err
}

func saveGitIgnoreAsFile(bytes []byte) bool {
	return helper.WriteFile(".gitignore", bytes)
}
