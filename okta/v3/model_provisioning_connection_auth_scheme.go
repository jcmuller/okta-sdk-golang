/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 4.0.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
	"fmt"
)

// ProvisioningConnectionAuthScheme the model 'ProvisioningConnectionAuthScheme'
type ProvisioningConnectionAuthScheme string

// List of ProvisioningConnectionAuthScheme
const (
	PROVISIONINGCONNECTIONAUTHSCHEME_TOKEN ProvisioningConnectionAuthScheme = "TOKEN"
	PROVISIONINGCONNECTIONAUTHSCHEME_UNKNOWN ProvisioningConnectionAuthScheme = "UNKNOWN"
)

// All allowed values of ProvisioningConnectionAuthScheme enum
var AllowedProvisioningConnectionAuthSchemeEnumValues = []ProvisioningConnectionAuthScheme{
	"TOKEN",
	"UNKNOWN",
}

func (v *ProvisioningConnectionAuthScheme) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProvisioningConnectionAuthScheme(value)
	for _, existing := range AllowedProvisioningConnectionAuthSchemeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProvisioningConnectionAuthScheme", value)
}

// NewProvisioningConnectionAuthSchemeFromValue returns a pointer to a valid ProvisioningConnectionAuthScheme
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProvisioningConnectionAuthSchemeFromValue(v string) (*ProvisioningConnectionAuthScheme, error) {
	ev := ProvisioningConnectionAuthScheme(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProvisioningConnectionAuthScheme: valid values are %v", v, AllowedProvisioningConnectionAuthSchemeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProvisioningConnectionAuthScheme) IsValid() bool {
	for _, existing := range AllowedProvisioningConnectionAuthSchemeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProvisioningConnectionAuthScheme value
func (v ProvisioningConnectionAuthScheme) Ptr() *ProvisioningConnectionAuthScheme {
	return &v
}

type NullableProvisioningConnectionAuthScheme struct {
	value *ProvisioningConnectionAuthScheme
	isSet bool
}

func (v NullableProvisioningConnectionAuthScheme) Get() *ProvisioningConnectionAuthScheme {
	return v.value
}

func (v *NullableProvisioningConnectionAuthScheme) Set(val *ProvisioningConnectionAuthScheme) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningConnectionAuthScheme) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningConnectionAuthScheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningConnectionAuthScheme(val *ProvisioningConnectionAuthScheme) *NullableProvisioningConnectionAuthScheme {
	return &NullableProvisioningConnectionAuthScheme{value: val, isSet: true}
}

func (v NullableProvisioningConnectionAuthScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningConnectionAuthScheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

