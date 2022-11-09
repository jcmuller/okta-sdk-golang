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

// FipsEnum the model 'FipsEnum'
type FipsEnum string

// List of FipsEnum
const (
	FIPSENUM_OPTIONAL FipsEnum = "OPTIONAL"
	FIPSENUM_REQUIRED FipsEnum = "REQUIRED"
)

// All allowed values of FipsEnum enum
var AllowedFipsEnumEnumValues = []FipsEnum{
	"OPTIONAL",
	"REQUIRED",
}

func (v *FipsEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FipsEnum(value)
	for _, existing := range AllowedFipsEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FipsEnum", value)
}

// NewFipsEnumFromValue returns a pointer to a valid FipsEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFipsEnumFromValue(v string) (*FipsEnum, error) {
	ev := FipsEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FipsEnum: valid values are %v", v, AllowedFipsEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FipsEnum) IsValid() bool {
	for _, existing := range AllowedFipsEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FipsEnum value
func (v FipsEnum) Ptr() *FipsEnum {
	return &v
}

type NullableFipsEnum struct {
	value *FipsEnum
	isSet bool
}

func (v NullableFipsEnum) Get() *FipsEnum {
	return v.value
}

func (v *NullableFipsEnum) Set(val *FipsEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableFipsEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableFipsEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFipsEnum(val *FipsEnum) *NullableFipsEnum {
	return &NullableFipsEnum{value: val, isSet: true}
}

func (v NullableFipsEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFipsEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

