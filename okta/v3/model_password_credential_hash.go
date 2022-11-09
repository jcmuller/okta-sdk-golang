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

// PasswordCredentialHash struct for PasswordCredentialHash
type PasswordCredentialHash struct {
	Algorithm *PasswordCredentialHashAlgorithm `json:"algorithm,omitempty"`
	Salt *string `json:"salt,omitempty"`
	SaltOrder *string `json:"saltOrder,omitempty"`
	Value *string `json:"value,omitempty"`
	WorkFactor *int32 `json:"workFactor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordCredentialHash PasswordCredentialHash

// NewPasswordCredentialHash instantiates a new PasswordCredentialHash object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordCredentialHash() *PasswordCredentialHash {
	this := PasswordCredentialHash{}
	return &this
}

// NewPasswordCredentialHashWithDefaults instantiates a new PasswordCredentialHash object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordCredentialHashWithDefaults() *PasswordCredentialHash {
	this := PasswordCredentialHash{}
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *PasswordCredentialHash) GetAlgorithm() PasswordCredentialHashAlgorithm {
	if o == nil || o.Algorithm == nil {
		var ret PasswordCredentialHashAlgorithm
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordCredentialHash) GetAlgorithmOk() (*PasswordCredentialHashAlgorithm, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *PasswordCredentialHash) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given PasswordCredentialHashAlgorithm and assigns it to the Algorithm field.
func (o *PasswordCredentialHash) SetAlgorithm(v PasswordCredentialHashAlgorithm) {
	o.Algorithm = &v
}

// GetSalt returns the Salt field value if set, zero value otherwise.
func (o *PasswordCredentialHash) GetSalt() string {
	if o == nil || o.Salt == nil {
		var ret string
		return ret
	}
	return *o.Salt
}

// GetSaltOk returns a tuple with the Salt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordCredentialHash) GetSaltOk() (*string, bool) {
	if o == nil || o.Salt == nil {
		return nil, false
	}
	return o.Salt, true
}

// HasSalt returns a boolean if a field has been set.
func (o *PasswordCredentialHash) HasSalt() bool {
	if o != nil && o.Salt != nil {
		return true
	}

	return false
}

// SetSalt gets a reference to the given string and assigns it to the Salt field.
func (o *PasswordCredentialHash) SetSalt(v string) {
	o.Salt = &v
}

// GetSaltOrder returns the SaltOrder field value if set, zero value otherwise.
func (o *PasswordCredentialHash) GetSaltOrder() string {
	if o == nil || o.SaltOrder == nil {
		var ret string
		return ret
	}
	return *o.SaltOrder
}

// GetSaltOrderOk returns a tuple with the SaltOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordCredentialHash) GetSaltOrderOk() (*string, bool) {
	if o == nil || o.SaltOrder == nil {
		return nil, false
	}
	return o.SaltOrder, true
}

// HasSaltOrder returns a boolean if a field has been set.
func (o *PasswordCredentialHash) HasSaltOrder() bool {
	if o != nil && o.SaltOrder != nil {
		return true
	}

	return false
}

// SetSaltOrder gets a reference to the given string and assigns it to the SaltOrder field.
func (o *PasswordCredentialHash) SetSaltOrder(v string) {
	o.SaltOrder = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PasswordCredentialHash) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordCredentialHash) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PasswordCredentialHash) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PasswordCredentialHash) SetValue(v string) {
	o.Value = &v
}

// GetWorkFactor returns the WorkFactor field value if set, zero value otherwise.
func (o *PasswordCredentialHash) GetWorkFactor() int32 {
	if o == nil || o.WorkFactor == nil {
		var ret int32
		return ret
	}
	return *o.WorkFactor
}

// GetWorkFactorOk returns a tuple with the WorkFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordCredentialHash) GetWorkFactorOk() (*int32, bool) {
	if o == nil || o.WorkFactor == nil {
		return nil, false
	}
	return o.WorkFactor, true
}

// HasWorkFactor returns a boolean if a field has been set.
func (o *PasswordCredentialHash) HasWorkFactor() bool {
	if o != nil && o.WorkFactor != nil {
		return true
	}

	return false
}

// SetWorkFactor gets a reference to the given int32 and assigns it to the WorkFactor field.
func (o *PasswordCredentialHash) SetWorkFactor(v int32) {
	o.WorkFactor = &v
}

func (o PasswordCredentialHash) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Algorithm != nil {
		toSerialize["algorithm"] = o.Algorithm
	}
	if o.Salt != nil {
		toSerialize["salt"] = o.Salt
	}
	if o.SaltOrder != nil {
		toSerialize["saltOrder"] = o.SaltOrder
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.WorkFactor != nil {
		toSerialize["workFactor"] = o.WorkFactor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordCredentialHash) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordCredentialHash := _PasswordCredentialHash{}

	err = json.Unmarshal(bytes, &varPasswordCredentialHash)
	if err == nil {
		*o = PasswordCredentialHash(varPasswordCredentialHash)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "algorithm")
		delete(additionalProperties, "salt")
		delete(additionalProperties, "saltOrder")
		delete(additionalProperties, "value")
		delete(additionalProperties, "workFactor")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordCredentialHash struct {
	value *PasswordCredentialHash
	isSet bool
}

func (v NullablePasswordCredentialHash) Get() *PasswordCredentialHash {
	return v.value
}

func (v *NullablePasswordCredentialHash) Set(val *PasswordCredentialHash) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordCredentialHash) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordCredentialHash) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordCredentialHash(val *PasswordCredentialHash) *NullablePasswordCredentialHash {
	return &NullablePasswordCredentialHash{value: val, isSet: true}
}

func (v NullablePasswordCredentialHash) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordCredentialHash) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

