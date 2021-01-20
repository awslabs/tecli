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

// Credentials model
type Credentials struct {
	Profiles []CredentialProfile `yaml:"profiles"`
}

// CredentialProfile model
type CredentialProfile struct {
	Name        string       `yaml:"name"`
	Description string       `yaml:"description,omitempty"`
	Enabled     bool         `yaml:"enabled"`
	CreatedAt   string       `yaml:"createdAt"`
	UpdatedAt   string       `yaml:"updatedAt"`
	Credentials []Credential `yaml:"credentials"`
}

// Credential model
type Credential struct {
	Name         string `yaml:"name,omitempty"`
	Description  string `yaml:"description,omitempty"`
	Enabled      bool   `yaml:"enabled"`
	CreatedAt    string `yaml:"createdAt"`
	UpdatedAt    string `yaml:"updatedAt"`
	Provider     string `yaml:"provider"`
	AccessKey    string `yaml:"accessKey"`
	SecretKey    string `yaml:"secretKey"`
	SessionToken string `yaml:"sessionToken"`
}

// Configurations model
type Configurations struct {
	Profiles []ConfigurationProfile `yaml:"profiles"`
}

// ConfigurationProfile model
type ConfigurationProfile struct {
	Name           string          `yaml:"name"`
	Description    string          `yaml:"description,omitempty"`
	Enabled        bool            `yaml:"enabled"`
	CreatedAt      string          `yaml:"createdAt"`
	UpdatedAt      string          `yaml:"updatedAt"`
	Configurations []Configuration `yaml:"configurations"`
}

// Configuration model
type Configuration struct {
	Name           string `yaml:"name,omitempty"`
	Description    string `yaml:"description,omitempty"`
	Enabled        bool   `yaml:"enabled"`
	CreatedAt      string `yaml:"createdAt"`
	UpdatedAt      string `yaml:"updatedAt"`
	Initialization `yaml:"initialization,omitempty"`
	Unsplash       `yaml:"unsplash,omitempty"`
}
