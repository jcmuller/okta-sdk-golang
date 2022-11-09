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

// GroupType the model 'GroupType'
type GroupType string

// List of GroupType
const (
	GROUPTYPE_APP_GROUP GroupType = "APP_GROUP"
	GROUPTYPE_BUILT_IN GroupType = "BUILT_IN"
	GROUPTYPE_OKTA_GROUP GroupType = "OKTA_GROUP"
)

// All allowed values of GroupType enum
var AllowedGroupTypeEnumValues = []GroupType{
	"APP_GROUP",
	"BUILT_IN",
	"OKTA_GROUP",
}

func (v *GroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GroupType(value)
	for _, existing := range AllowedGroupTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GroupType", value)
}

// NewGroupTypeFromValue returns a pointer to a valid GroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGroupTypeFromValue(v string) (*GroupType, error) {
	ev := GroupType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GroupType: valid values are %v", v, AllowedGroupTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GroupType) IsValid() bool {
	for _, existing := range AllowedGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GroupType value
func (v GroupType) Ptr() *GroupType {
	return &v
}

type NullableGroupType struct {
	value *GroupType
	isSet bool
}

func (v NullableGroupType) Get() *GroupType {
	return v.value
}

func (v *NullableGroupType) Set(val *GroupType) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupType) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupType(val *GroupType) *NullableGroupType {
	return &NullableGroupType{value: val, isSet: true}
}

func (v NullableGroupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

