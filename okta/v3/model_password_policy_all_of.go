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
)

// PasswordPolicyAllOf struct for PasswordPolicyAllOf
type PasswordPolicyAllOf struct {
	Conditions *PasswordPolicyConditions `json:"conditions,omitempty"`
	Settings *PasswordPolicySettings `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyAllOf PasswordPolicyAllOf

// NewPasswordPolicyAllOf instantiates a new PasswordPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyAllOf() *PasswordPolicyAllOf {
	this := PasswordPolicyAllOf{}
	return &this
}

// NewPasswordPolicyAllOfWithDefaults instantiates a new PasswordPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyAllOfWithDefaults() *PasswordPolicyAllOf {
	this := PasswordPolicyAllOf{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *PasswordPolicyAllOf) GetConditions() PasswordPolicyConditions {
	if o == nil || o.Conditions == nil {
		var ret PasswordPolicyConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyAllOf) GetConditionsOk() (*PasswordPolicyConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *PasswordPolicyAllOf) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given PasswordPolicyConditions and assigns it to the Conditions field.
func (o *PasswordPolicyAllOf) SetConditions(v PasswordPolicyConditions) {
	o.Conditions = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *PasswordPolicyAllOf) GetSettings() PasswordPolicySettings {
	if o == nil || o.Settings == nil {
		var ret PasswordPolicySettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyAllOf) GetSettingsOk() (*PasswordPolicySettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *PasswordPolicyAllOf) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given PasswordPolicySettings and assigns it to the Settings field.
func (o *PasswordPolicyAllOf) SetSettings(v PasswordPolicySettings) {
	o.Settings = &v
}

func (o PasswordPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordPolicyAllOf := _PasswordPolicyAllOf{}

	err = json.Unmarshal(bytes, &varPasswordPolicyAllOf)
	if err == nil {
		*o = PasswordPolicyAllOf(varPasswordPolicyAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "settings")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordPolicyAllOf struct {
	value *PasswordPolicyAllOf
	isSet bool
}

func (v NullablePasswordPolicyAllOf) Get() *PasswordPolicyAllOf {
	return v.value
}

func (v *NullablePasswordPolicyAllOf) Set(val *PasswordPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyAllOf(val *PasswordPolicyAllOf) *NullablePasswordPolicyAllOf {
	return &NullablePasswordPolicyAllOf{value: val, isSet: true}
}

func (v NullablePasswordPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

