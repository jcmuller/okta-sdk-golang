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

// MultifactorEnrollmentPolicyAllOf struct for MultifactorEnrollmentPolicyAllOf
type MultifactorEnrollmentPolicyAllOf struct {
	Conditions *PolicyRuleConditions `json:"conditions,omitempty"`
	Settings *MultifactorEnrollmentPolicySettings `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MultifactorEnrollmentPolicyAllOf MultifactorEnrollmentPolicyAllOf

// NewMultifactorEnrollmentPolicyAllOf instantiates a new MultifactorEnrollmentPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultifactorEnrollmentPolicyAllOf() *MultifactorEnrollmentPolicyAllOf {
	this := MultifactorEnrollmentPolicyAllOf{}
	return &this
}

// NewMultifactorEnrollmentPolicyAllOfWithDefaults instantiates a new MultifactorEnrollmentPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultifactorEnrollmentPolicyAllOfWithDefaults() *MultifactorEnrollmentPolicyAllOf {
	this := MultifactorEnrollmentPolicyAllOf{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *MultifactorEnrollmentPolicyAllOf) GetConditions() PolicyRuleConditions {
	if o == nil || o.Conditions == nil {
		var ret PolicyRuleConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultifactorEnrollmentPolicyAllOf) GetConditionsOk() (*PolicyRuleConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *MultifactorEnrollmentPolicyAllOf) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given PolicyRuleConditions and assigns it to the Conditions field.
func (o *MultifactorEnrollmentPolicyAllOf) SetConditions(v PolicyRuleConditions) {
	o.Conditions = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *MultifactorEnrollmentPolicyAllOf) GetSettings() MultifactorEnrollmentPolicySettings {
	if o == nil || o.Settings == nil {
		var ret MultifactorEnrollmentPolicySettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultifactorEnrollmentPolicyAllOf) GetSettingsOk() (*MultifactorEnrollmentPolicySettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *MultifactorEnrollmentPolicyAllOf) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given MultifactorEnrollmentPolicySettings and assigns it to the Settings field.
func (o *MultifactorEnrollmentPolicyAllOf) SetSettings(v MultifactorEnrollmentPolicySettings) {
	o.Settings = &v
}

func (o MultifactorEnrollmentPolicyAllOf) MarshalJSON() ([]byte, error) {
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

func (o *MultifactorEnrollmentPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMultifactorEnrollmentPolicyAllOf := _MultifactorEnrollmentPolicyAllOf{}

	err = json.Unmarshal(bytes, &varMultifactorEnrollmentPolicyAllOf)
	if err == nil {
		*o = MultifactorEnrollmentPolicyAllOf(varMultifactorEnrollmentPolicyAllOf)
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

type NullableMultifactorEnrollmentPolicyAllOf struct {
	value *MultifactorEnrollmentPolicyAllOf
	isSet bool
}

func (v NullableMultifactorEnrollmentPolicyAllOf) Get() *MultifactorEnrollmentPolicyAllOf {
	return v.value
}

func (v *NullableMultifactorEnrollmentPolicyAllOf) Set(val *MultifactorEnrollmentPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMultifactorEnrollmentPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMultifactorEnrollmentPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultifactorEnrollmentPolicyAllOf(val *MultifactorEnrollmentPolicyAllOf) *NullableMultifactorEnrollmentPolicyAllOf {
	return &NullableMultifactorEnrollmentPolicyAllOf{value: val, isSet: true}
}

func (v NullableMultifactorEnrollmentPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultifactorEnrollmentPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

