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

// ProvisioningGroupsAction the model 'ProvisioningGroupsAction'
type ProvisioningGroupsAction string

// List of ProvisioningGroupsAction
const (
	PROVISIONINGGROUPSACTION_APPEND ProvisioningGroupsAction = "APPEND"
	PROVISIONINGGROUPSACTION_ASSIGN ProvisioningGroupsAction = "ASSIGN"
	PROVISIONINGGROUPSACTION_NONE ProvisioningGroupsAction = "NONE"
	PROVISIONINGGROUPSACTION_SYNC ProvisioningGroupsAction = "SYNC"
)

// All allowed values of ProvisioningGroupsAction enum
var AllowedProvisioningGroupsActionEnumValues = []ProvisioningGroupsAction{
	"APPEND",
	"ASSIGN",
	"NONE",
	"SYNC",
}

func (v *ProvisioningGroupsAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProvisioningGroupsAction(value)
	for _, existing := range AllowedProvisioningGroupsActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProvisioningGroupsAction", value)
}

// NewProvisioningGroupsActionFromValue returns a pointer to a valid ProvisioningGroupsAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProvisioningGroupsActionFromValue(v string) (*ProvisioningGroupsAction, error) {
	ev := ProvisioningGroupsAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProvisioningGroupsAction: valid values are %v", v, AllowedProvisioningGroupsActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProvisioningGroupsAction) IsValid() bool {
	for _, existing := range AllowedProvisioningGroupsActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProvisioningGroupsAction value
func (v ProvisioningGroupsAction) Ptr() *ProvisioningGroupsAction {
	return &v
}

type NullableProvisioningGroupsAction struct {
	value *ProvisioningGroupsAction
	isSet bool
}

func (v NullableProvisioningGroupsAction) Get() *ProvisioningGroupsAction {
	return v.value
}

func (v *NullableProvisioningGroupsAction) Set(val *ProvisioningGroupsAction) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningGroupsAction) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningGroupsAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningGroupsAction(val *ProvisioningGroupsAction) *NullableProvisioningGroupsAction {
	return &NullableProvisioningGroupsAction{value: val, isSet: true}
}

func (v NullableProvisioningGroupsAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningGroupsAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

