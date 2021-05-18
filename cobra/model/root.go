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

// App represents the all the necessary information about tecli
type App struct {
	// Name of file to look for inside the path
	Name                       string
	ConfigurationsDir          string
	CredentialsFileName        string
	CredentialsFileType        string
	CredentialsFilePath        string
	CredentialsFilePermissions os.FileMode
	LogsDir                    string
	LogsFileName               string
	LogsFileType               string
	LogsFilePath               string
	LogsFilePermissions        os.FileMode

	WorkingDir string
}
