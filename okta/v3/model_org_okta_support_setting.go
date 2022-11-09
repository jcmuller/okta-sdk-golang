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

// OrgOktaSupportSetting the model 'OrgOktaSupportSetting'
type OrgOktaSupportSetting string

// List of OrgOktaSupportSetting
const (
	ORGOKTASUPPORTSETTING_DISABLED OrgOktaSupportSetting = "DISABLED"
	ORGOKTASUPPORTSETTING_ENABLED OrgOktaSupportSetting = "ENABLED"
)

// All allowed values of OrgOktaSupportSetting enum
var AllowedOrgOktaSupportSettingEnumValues = []OrgOktaSupportSetting{
	"DISABLED",
	"ENABLED",
}

func (v *OrgOktaSupportSetting) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrgOktaSupportSetting(value)
	for _, existing := range AllowedOrgOktaSupportSettingEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrgOktaSupportSetting", value)
}

// NewOrgOktaSupportSettingFromValue returns a pointer to a valid OrgOktaSupportSetting
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrgOktaSupportSettingFromValue(v string) (*OrgOktaSupportSetting, error) {
	ev := OrgOktaSupportSetting(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrgOktaSupportSetting: valid values are %v", v, AllowedOrgOktaSupportSettingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrgOktaSupportSetting) IsValid() bool {
	for _, existing := range AllowedOrgOktaSupportSettingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrgOktaSupportSetting value
func (v OrgOktaSupportSetting) Ptr() *OrgOktaSupportSetting {
	return &v
}

type NullableOrgOktaSupportSetting struct {
	value *OrgOktaSupportSetting
	isSet bool
}

func (v NullableOrgOktaSupportSetting) Get() *OrgOktaSupportSetting {
	return v.value
}

func (v *NullableOrgOktaSupportSetting) Set(val *OrgOktaSupportSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgOktaSupportSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgOktaSupportSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgOktaSupportSetting(val *OrgOktaSupportSetting) *NullableOrgOktaSupportSetting {
	return &NullableOrgOktaSupportSetting{value: val, isSet: true}
}

func (v NullableOrgOktaSupportSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgOktaSupportSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

