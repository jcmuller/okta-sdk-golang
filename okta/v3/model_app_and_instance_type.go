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

// AppAndInstanceType the model 'AppAndInstanceType'
type AppAndInstanceType string

// List of AppAndInstanceType
const (
	APPANDINSTANCETYPE_APP AppAndInstanceType = "APP"
	APPANDINSTANCETYPE_APP_TYPE AppAndInstanceType = "APP_TYPE"
)

// All allowed values of AppAndInstanceType enum
var AllowedAppAndInstanceTypeEnumValues = []AppAndInstanceType{
	"APP",
	"APP_TYPE",
}

func (v *AppAndInstanceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppAndInstanceType(value)
	for _, existing := range AllowedAppAndInstanceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppAndInstanceType", value)
}

// NewAppAndInstanceTypeFromValue returns a pointer to a valid AppAndInstanceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppAndInstanceTypeFromValue(v string) (*AppAndInstanceType, error) {
	ev := AppAndInstanceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppAndInstanceType: valid values are %v", v, AllowedAppAndInstanceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppAndInstanceType) IsValid() bool {
	for _, existing := range AllowedAppAndInstanceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppAndInstanceType value
func (v AppAndInstanceType) Ptr() *AppAndInstanceType {
	return &v
}

type NullableAppAndInstanceType struct {
	value *AppAndInstanceType
	isSet bool
}

func (v NullableAppAndInstanceType) Get() *AppAndInstanceType {
	return v.value
}

func (v *NullableAppAndInstanceType) Set(val *AppAndInstanceType) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAndInstanceType) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAndInstanceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAndInstanceType(val *AppAndInstanceType) *NullableAppAndInstanceType {
	return &NullableAppAndInstanceType{value: val, isSet: true}
}

func (v NullableAppAndInstanceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAndInstanceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

