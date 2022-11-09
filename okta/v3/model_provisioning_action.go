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

// ProvisioningAction the model 'ProvisioningAction'
type ProvisioningAction string

// List of ProvisioningAction
const (
	PROVISIONINGACTION_AUTO ProvisioningAction = "AUTO"
	PROVISIONINGACTION_CALLOUT ProvisioningAction = "CALLOUT"
	PROVISIONINGACTION_DISABLED ProvisioningAction = "DISABLED"
)

// All allowed values of ProvisioningAction enum
var AllowedProvisioningActionEnumValues = []ProvisioningAction{
	"AUTO",
	"CALLOUT",
	"DISABLED",
}

func (v *ProvisioningAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProvisioningAction(value)
	for _, existing := range AllowedProvisioningActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProvisioningAction", value)
}

// NewProvisioningActionFromValue returns a pointer to a valid ProvisioningAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProvisioningActionFromValue(v string) (*ProvisioningAction, error) {
	ev := ProvisioningAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProvisioningAction: valid values are %v", v, AllowedProvisioningActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProvisioningAction) IsValid() bool {
	for _, existing := range AllowedProvisioningActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProvisioningAction value
func (v ProvisioningAction) Ptr() *ProvisioningAction {
	return &v
}

type NullableProvisioningAction struct {
	value *ProvisioningAction
	isSet bool
}

func (v NullableProvisioningAction) Get() *ProvisioningAction {
	return v.value
}

func (v *NullableProvisioningAction) Set(val *ProvisioningAction) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningAction) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningAction(val *ProvisioningAction) *NullableProvisioningAction {
	return &NullableProvisioningAction{value: val, isSet: true}
}

func (v NullableProvisioningAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

