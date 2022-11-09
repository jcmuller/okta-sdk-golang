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

// TotpUserFactorProfile struct for TotpUserFactorProfile
type TotpUserFactorProfile struct {
	CredentialId *string `json:"credentialId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TotpUserFactorProfile TotpUserFactorProfile

// NewTotpUserFactorProfile instantiates a new TotpUserFactorProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTotpUserFactorProfile() *TotpUserFactorProfile {
	this := TotpUserFactorProfile{}
	return &this
}

// NewTotpUserFactorProfileWithDefaults instantiates a new TotpUserFactorProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTotpUserFactorProfileWithDefaults() *TotpUserFactorProfile {
	this := TotpUserFactorProfile{}
	return &this
}

// GetCredentialId returns the CredentialId field value if set, zero value otherwise.
func (o *TotpUserFactorProfile) GetCredentialId() string {
	if o == nil || o.CredentialId == nil {
		var ret string
		return ret
	}
	return *o.CredentialId
}

// GetCredentialIdOk returns a tuple with the CredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TotpUserFactorProfile) GetCredentialIdOk() (*string, bool) {
	if o == nil || o.CredentialId == nil {
		return nil, false
	}
	return o.CredentialId, true
}

// HasCredentialId returns a boolean if a field has been set.
func (o *TotpUserFactorProfile) HasCredentialId() bool {
	if o != nil && o.CredentialId != nil {
		return true
	}

	return false
}

// SetCredentialId gets a reference to the given string and assigns it to the CredentialId field.
func (o *TotpUserFactorProfile) SetCredentialId(v string) {
	o.CredentialId = &v
}

func (o TotpUserFactorProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CredentialId != nil {
		toSerialize["credentialId"] = o.CredentialId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TotpUserFactorProfile) UnmarshalJSON(bytes []byte) (err error) {
	varTotpUserFactorProfile := _TotpUserFactorProfile{}

	err = json.Unmarshal(bytes, &varTotpUserFactorProfile)
	if err == nil {
		*o = TotpUserFactorProfile(varTotpUserFactorProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "credentialId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTotpUserFactorProfile struct {
	value *TotpUserFactorProfile
	isSet bool
}

func (v NullableTotpUserFactorProfile) Get() *TotpUserFactorProfile {
	return v.value
}

func (v *NullableTotpUserFactorProfile) Set(val *TotpUserFactorProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableTotpUserFactorProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableTotpUserFactorProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTotpUserFactorProfile(val *TotpUserFactorProfile) *NullableTotpUserFactorProfile {
	return &NullableTotpUserFactorProfile{value: val, isSet: true}
}

func (v NullableTotpUserFactorProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTotpUserFactorProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

