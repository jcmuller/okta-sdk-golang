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

// EventHookChannelConfigAuthSchemeType the model 'EventHookChannelConfigAuthSchemeType'
type EventHookChannelConfigAuthSchemeType string

// List of EventHookChannelConfigAuthSchemeType
const (
	EVENTHOOKCHANNELCONFIGAUTHSCHEMETYPE_HEADER EventHookChannelConfigAuthSchemeType = "HEADER"
)

// All allowed values of EventHookChannelConfigAuthSchemeType enum
var AllowedEventHookChannelConfigAuthSchemeTypeEnumValues = []EventHookChannelConfigAuthSchemeType{
	"HEADER",
}

func (v *EventHookChannelConfigAuthSchemeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventHookChannelConfigAuthSchemeType(value)
	for _, existing := range AllowedEventHookChannelConfigAuthSchemeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventHookChannelConfigAuthSchemeType", value)
}

// NewEventHookChannelConfigAuthSchemeTypeFromValue returns a pointer to a valid EventHookChannelConfigAuthSchemeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventHookChannelConfigAuthSchemeTypeFromValue(v string) (*EventHookChannelConfigAuthSchemeType, error) {
	ev := EventHookChannelConfigAuthSchemeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventHookChannelConfigAuthSchemeType: valid values are %v", v, AllowedEventHookChannelConfigAuthSchemeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventHookChannelConfigAuthSchemeType) IsValid() bool {
	for _, existing := range AllowedEventHookChannelConfigAuthSchemeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventHookChannelConfigAuthSchemeType value
func (v EventHookChannelConfigAuthSchemeType) Ptr() *EventHookChannelConfigAuthSchemeType {
	return &v
}

type NullableEventHookChannelConfigAuthSchemeType struct {
	value *EventHookChannelConfigAuthSchemeType
	isSet bool
}

func (v NullableEventHookChannelConfigAuthSchemeType) Get() *EventHookChannelConfigAuthSchemeType {
	return v.value
}

func (v *NullableEventHookChannelConfigAuthSchemeType) Set(val *EventHookChannelConfigAuthSchemeType) {
	v.value = val
	v.isSet = true
}

func (v NullableEventHookChannelConfigAuthSchemeType) IsSet() bool {
	return v.isSet
}

func (v *NullableEventHookChannelConfigAuthSchemeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventHookChannelConfigAuthSchemeType(val *EventHookChannelConfigAuthSchemeType) *NullableEventHookChannelConfigAuthSchemeType {
	return &NullableEventHookChannelConfigAuthSchemeType{value: val, isSet: true}
}

func (v NullableEventHookChannelConfigAuthSchemeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventHookChannelConfigAuthSchemeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

