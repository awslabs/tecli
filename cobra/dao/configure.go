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

// Package dao represents the Data Access Object Pattern or DAO pattern is used to separate low level data accessing API or operations from high level business services
package dao

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.aws.dev/devops-aws/tecli/cobra/aid"
	"gitlab.aws.dev/devops-aws/tecli/cobra/model"
)

// GetCredentials read the current credentials file and return its model
func GetCredentials() (model.Credentials, error) {
	var creds model.Credentials
	err := viper.ReadInConfig()
	if err != nil {
		return creds, fmt.Errorf("unable to read credentials\n%v", err)
	}

	err = viper.Unmarshal(&creds)
	if err != nil {
		return creds, fmt.Errorf("unable to unmarshall credentials\n%v", err)
	}

	return creds, err
}

// GetCredentialProfile returns credentials of a profile
func GetCredentialProfile(name string) (model.CredentialProfile, error) {
	credentials, err := GetCredentials()

	if err != nil {
		return (model.CredentialProfile{}), err
	}

	for _, profile := range credentials.Profiles {
		if profile.Name == name && profile.Enabled {
			return profile, err
		}
	}

	return (model.CredentialProfile{}), err
}

// GetTeamToken return the team token from credentials file
func GetTeamToken(name string) string {
	cp, err := GetCredentialProfile(name)
	if err != nil {
		logrus.Fatalf("unable to read team token from credentials\n%v\n", err)
	}

	return cp.TeamToken
}

// GetOrganizationToken return the organization token from credentials file
func GetOrganizationToken(name string) string {
	cp, err := GetCredentialProfile(name)
	if err != nil {
		logrus.Fatalf("unable to read organization token from configurations\n%v\n", err)
	}
	return cp.OrganizationToken
}

// SaveCredentials saves the given credential onto the credentials file
func SaveCredentials(credentials model.Credentials) error {
	return aid.WriteInterfaceToFile(credentials, viper.ConfigFileUsed())
}

// SaveCredentialsV2 saves the given credential onto the credentials file
func SaveCredentialsV2(credentials model.Credentials, path string) error {
	return aid.WriteInterfaceToFile(credentials, path)
}
