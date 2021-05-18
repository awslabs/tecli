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

// Package model represents an object carrying data. It can also have logic to update controller if its data changes.
package model

// Credentials model
type Credentials struct {
	Profiles []CredentialProfile `yaml:"profiles"`
}

// CredentialProfile model
type CredentialProfile struct {
	Name              string `yaml:"name"`
	Description       string `yaml:"description,omitempty"`
	Organization      string `yaml:"organization"`
	UserToken         string `yaml:"userToken"`
	TeamToken         string `yaml:"teamToken"`
	OrganizationToken string `yaml:"organizationToken"`
}
