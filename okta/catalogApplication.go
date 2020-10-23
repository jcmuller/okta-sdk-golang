/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// AUTO-GENERATED!  DO NOT EDIT FILE DIRECTLY

package okta

import (
	"time"
)

type CatalogApplication struct {
	Links              interface{} `json:"_links,omitempty"`
	Category           string      `json:"category,omitempty"`
	Description        string      `json:"description,omitempty"`
	DisplayName        string      `json:"displayName,omitempty"`
	Features           []string    `json:"features,omitempty"`
	Id                 string      `json:"id,omitempty"`
	LastUpdated        *time.Time  `json:"lastUpdated,omitempty"`
	Name               string      `json:"name,omitempty"`
	SignOnModes        []string    `json:"signOnModes,omitempty"`
	Status             string      `json:"status,omitempty"`
	VerificationStatus string      `json:"verificationStatus,omitempty"`
	Website            string      `json:"website,omitempty"`
}