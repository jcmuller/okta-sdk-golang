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

// SecurityQuestionUserFactorAllOf struct for SecurityQuestionUserFactorAllOf
type SecurityQuestionUserFactorAllOf struct {
	Profile *SecurityQuestionUserFactorProfile `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityQuestionUserFactorAllOf SecurityQuestionUserFactorAllOf

// NewSecurityQuestionUserFactorAllOf instantiates a new SecurityQuestionUserFactorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityQuestionUserFactorAllOf() *SecurityQuestionUserFactorAllOf {
	this := SecurityQuestionUserFactorAllOf{}
	return &this
}

// NewSecurityQuestionUserFactorAllOfWithDefaults instantiates a new SecurityQuestionUserFactorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityQuestionUserFactorAllOfWithDefaults() *SecurityQuestionUserFactorAllOf {
	this := SecurityQuestionUserFactorAllOf{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *SecurityQuestionUserFactorAllOf) GetProfile() SecurityQuestionUserFactorProfile {
	if o == nil || o.Profile == nil {
		var ret SecurityQuestionUserFactorProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityQuestionUserFactorAllOf) GetProfileOk() (*SecurityQuestionUserFactorProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *SecurityQuestionUserFactorAllOf) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given SecurityQuestionUserFactorProfile and assigns it to the Profile field.
func (o *SecurityQuestionUserFactorAllOf) SetProfile(v SecurityQuestionUserFactorProfile) {
	o.Profile = &v
}

func (o SecurityQuestionUserFactorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityQuestionUserFactorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityQuestionUserFactorAllOf := _SecurityQuestionUserFactorAllOf{}

	err = json.Unmarshal(bytes, &varSecurityQuestionUserFactorAllOf)
	if err == nil {
		*o = SecurityQuestionUserFactorAllOf(varSecurityQuestionUserFactorAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityQuestionUserFactorAllOf struct {
	value *SecurityQuestionUserFactorAllOf
	isSet bool
}

func (v NullableSecurityQuestionUserFactorAllOf) Get() *SecurityQuestionUserFactorAllOf {
	return v.value
}

func (v *NullableSecurityQuestionUserFactorAllOf) Set(val *SecurityQuestionUserFactorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityQuestionUserFactorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityQuestionUserFactorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityQuestionUserFactorAllOf(val *SecurityQuestionUserFactorAllOf) *NullableSecurityQuestionUserFactorAllOf {
	return &NullableSecurityQuestionUserFactorAllOf{value: val, isSet: true}
}

func (v NullableSecurityQuestionUserFactorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityQuestionUserFactorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

