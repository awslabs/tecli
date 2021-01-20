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

package model

import "os"

// App represents the all the necessary information about tfe-cli
type App struct {
	// Name of file to look for inside the path
	Name                      string
	HomeDir                   string
	ConfigurationsDir         string
	ConfigurationsName        string
	ConfigurationsType        string
	ConfigurationsPath        string
	ConfigurationsPermissions os.FileMode
	CredentialsName           string
	CredentialsType           string
	CredentialsPath           string
	CredentialsPermissions    os.FileMode
	LogsDir                   string
	LogsName                  string
	LogsType                  string
	LogsPath                  string
	LogsPermissions           os.FileMode

	WorkingDir string
}

// ReadMe struct of the readme.yaml
type ReadMe struct {
	Logo struct {
		URL   string `yaml:"url"`
		Label string `yaml:"label"`
	} `yaml:"logo,omitempty"`
	Shields struct {
		Badges []struct {
			Description string `yaml:"description"`
			Image       string `yaml:"image"`
			URL         string `yaml:"url"`
		} `yaml:"badges"`
	} `yaml:"shields,omitempty"`
	App struct {
		Name     string `yaml:"name"`
		Function string `yaml:"function"`
		ID       string `yaml:"id"`
	} `yaml:"app,omitempty"`
	Screenshots []struct {
		Caption string `yaml:"caption"`
		Label   string `yaml:"label"`
		URL     string `yaml:"url"`
	} `yaml:"screenshots,omitempty"`
	Usage         string `yaml:"usage"`
	Prerequisites []struct {
		Description string `yaml:"description"`
		Name        string `yaml:"name"`
		URL         string `yaml:"url"`
	} `yaml:"prerequisites,omitempty"`
	Installing   string   `yaml:"installing,omitempty"`
	Testing      string   `yaml:"testing,omitempty"`
	Deployment   string   `yaml:"deployment,omitempty"`
	Include      []string `yaml:"include,omitempty"`
	Contributors []struct {
		Name  string `yaml:"name"`
		Role  string `yaml:"role"`
		Email string `yaml:"email"`
	} `yaml:"contributors,omitempty"`
	Acknowledgments []struct {
		Name string `yaml:"name"`
		Role string `yaml:"role"`
	} `yaml:"acknowledgments,omitempty"`
	References []struct {
		Description string `yaml:"description"`
		Name        string `yaml:"name"`
		URL         string `yaml:"url"`
	} `yaml:"references,omitempty"`
	License   string `yaml:"license,omitempty"`
	Copyright string `yaml:"copyright,omitempty"`
}
