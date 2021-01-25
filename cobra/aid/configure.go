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
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.aws.dev/devops-aws/terraform-ce-cli/helper"
	"gopkg.in/yaml.v2"
)

// ConfigurationsDirectoryExist returns `true` if the configuration directory exist, `false` otherwise
func ConfigurationsDirectoryExist() bool {
	return helper.DirOrFileExists(GetAppInfo().ConfigurationsDir)
}

// ConfigurationsFileExist returns `true` if the configuration file exist, `false` otherwise
func ConfigurationsFileExist() bool {
	return helper.DirOrFileExists(GetAppInfo().ConfigurationsPath)
}

// CreateConfigurationsDirectory creates the configuration directory, returns `true` if the configuration directory exist, `false` otherwise
func CreateConfigurationsDirectory() (bool, string) {
	dir := GetAppInfo().ConfigurationsDir
	return helper.MkDirsIfNotExist(dir), dir
}

// CredentialsFileExist returns `true` if the credentials file exist, `false` otherwise
func CredentialsFileExist() bool {
	return helper.DirOrFileExists(GetAppInfo().CredentialsPath)
}

// ReadConfig returns the viper instance of the given configuration `name`
func ReadConfig(name string) (*viper.Viper, error) {
	v := viper.New()
	app := GetAppInfo()

	v.SetConfigName(name)
	v.SetConfigType("yaml")
	v.AddConfigPath(app.ConfigurationsDir)

	err := v.ReadInConfig()
	if err != nil {
		return v, fmt.Errorf("unable to read configuration:%s\n%v\n", name, err)
	}
	return v, err
}

// ReadConfigAsViper returns...
func ReadConfigAsViper(configPath string, configName string, configType string) (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath(configPath)
	v.SetConfigName(configName)
	v.SetConfigType(configType)

	err := v.ReadInConfig()
	if err != nil {
		return v, fmt.Errorf("unable to read configuration as viper\n%v\n", err)
	}
	return v, err
}

// ReadTemplate read the given template under tecli/*.yaml
func ReadTemplate(fileName string) (*viper.Viper, error) {
	c := viper.New()
	c.AddConfigPath("tecli")
	c.SetConfigName(fileName)
	c.SetConfigType("yaml")
	c.SetConfigPermissions(os.ModePerm)

	err := c.ReadInConfig() // Find and read the c file
	if err != nil {         // Handle errors reading the c file
		return c, fmt.Errorf("Unable to read "+fileName+" via Viper"+"\n%v\n", err)
	}

	return c, nil
}

// WriteInterfaceToFile write the given interface into a file
func WriteInterfaceToFile(in interface{}, path string) error {
	b, err := yaml.Marshal(&in)
	if err != nil {
		_, ok := err.(*json.UnsupportedTypeError)
		if ok {
			return fmt.Errorf("json unsupported type error")
		}
	}

	err = ioutil.WriteFile(path, b, os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to update:%s\n%v\n", path, err)
	}

	return err
}

// DeleteCredentialFile delete the credentials file
func DeleteCredentialFile() error {
	return helper.DeleteFile(GetAppInfo().CredentialsPath)
}

// DeleteConfigurationFile delete the credentials file
func DeleteConfigurationFile() error {
	return helper.DeleteFile(GetAppInfo().ConfigurationsPath)
}

// DeleteConfigurationsDirectory delete the configurations directory
func DeleteConfigurationsDirectory() error {
	return os.RemoveAll(GetAppInfo().ConfigurationsDir)
}

// GetSensitiveUserInput get sensitive input as string
func GetSensitiveUserInput(cmd *cobra.Command, text string, info string) (string, error) {
	return getUserInput(cmd, text+" ["+maskString(info, 3)+"]", "")
}

func maskString(s string, showLastChars int) string {
	maskSize := len(s) - showLastChars
	if maskSize <= 0 {
		return s
	}

	return strings.Repeat("*", maskSize) + s[maskSize:]
}

// GetSensitiveUserInputAsString get sensitive input as string
func GetSensitiveUserInputAsString(cmd *cobra.Command, text string, info string) string {
	answer, err := GetSensitiveUserInput(cmd, text, info)
	if err != nil {
		log.Fatalf("unable to get user input about profile's name\n%v\n", err)
	}

	// if user typed ENTER, keep the current value
	if answer != "" {
		return answer
	}

	return info
}

func getUserInput(cmd *cobra.Command, text string, info string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	if info == "" {
		fmt.Print(text + ": ")
	} else {
		fmt.Print(text + " [" + info + "]: ")
	}

	input, err := reader.ReadString('\n')
	// convert CRLF to LF
	input = strings.Replace(input, "\n", "", -1)
	if err != nil {
		return input, fmt.Errorf("unable to read user input\n%v\n", err)
	}

	return input, err
}

// GetUserInputAsBool prints `text` on console and return answer as `boolean`
func GetUserInputAsBool(cmd *cobra.Command, text string, info bool) bool {
	answer, err := getUserInput(cmd, text, strconv.FormatBool(info))
	if err != nil {
		log.Fatalf("unable to get user input as boolean\n%s", err)
	}

	if answer == "true" {
		return true
	} else if answer == "false" {
		return false
	}

	return info
}

// GetUserInputAsString prints `text` on console and return answer as `string`
func GetUserInputAsString(cmd *cobra.Command, text string, info string) string {
	answer, err := getUserInput(cmd, text, info)
	if err != nil {
		log.Fatalf("unable to get user input about profile's name\n%v\n", err)
	}

	// if user typed ENTER, keep the current value
	if answer != "" {
		return answer
	}

	return info
}
