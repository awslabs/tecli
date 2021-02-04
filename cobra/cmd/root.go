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

package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"gitlab.aws.dev/devops-aws/tecli/cobra/aid"
	"gitlab.aws.dev/devops-aws/tecli/cobra/controller"

	"github.com/spf13/viper"
)

var rootCmd = controller.RootCmd()

//The verbose flag value
var verbosity string
var log string
var logFilePath string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&verbosity, "verbosity", "v", logrus.ErrorLevel.String(), "Valid log level:panic,fatal,error,warn,info,debug,trace).")
	rootCmd.PersistentFlags().StringVarP(&log, "log", "l", "disable", "Enable or disable logs (found at $HOME/.tecli/logs.json). Log outputs will be shown on default output.")
	rootCmd.PersistentFlags().StringVar(&logFilePath, "log-file-path", aid.GetAppInfo().LogsPath, "Log file path.")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	app := aid.GetAppInfo()
	viper.AddConfigPath(app.ConfigurationsDir) // global directory
	viper.SetConfigName(app.ConfigurationsName)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("using config file:", viper.ConfigFileUsed())
	}

	if log == "enable" && logFilePath != "" {
		if err := aid.SetupLoggingLevel(verbosity); err == nil {
			fmt.Printf("logging level: %s\n", verbosity)
		}

		if err := aid.SetupLoggingOutput(logFilePath); err == nil {
			fmt.Printf("logging path: %s\n", logFilePath)
		}
	}

}
