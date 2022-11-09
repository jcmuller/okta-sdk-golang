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

// SchemeApplicationCredentialsAllOf struct for SchemeApplicationCredentialsAllOf
type SchemeApplicationCredentialsAllOf struct {
	Password *PasswordCredential `json:"password,omitempty"`
	RevealPassword *bool `json:"revealPassword,omitempty"`
	Scheme *ApplicationCredentialsScheme `json:"scheme,omitempty"`
	Signing *ApplicationCredentialsSigning `json:"signing,omitempty"`
	UserName *string `json:"userName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SchemeApplicationCredentialsAllOf SchemeApplicationCredentialsAllOf

// NewSchemeApplicationCredentialsAllOf instantiates a new SchemeApplicationCredentialsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemeApplicationCredentialsAllOf() *SchemeApplicationCredentialsAllOf {
	this := SchemeApplicationCredentialsAllOf{}
	return &this
}

// NewSchemeApplicationCredentialsAllOfWithDefaults instantiates a new SchemeApplicationCredentialsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemeApplicationCredentialsAllOfWithDefaults() *SchemeApplicationCredentialsAllOf {
	this := SchemeApplicationCredentialsAllOf{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SchemeApplicationCredentialsAllOf) GetPassword() PasswordCredential {
	if o == nil || o.Password == nil {
		var ret PasswordCredential
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemeApplicationCredentialsAllOf) GetPasswordOk() (*PasswordCredential, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SchemeApplicationCredentialsAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given PasswordCredential and assigns it to the Password field.
func (o *SchemeApplicationCredentialsAllOf) SetPassword(v PasswordCredential) {
	o.Password = &v
}

// GetRevealPassword returns the RevealPassword field value if set, zero value otherwise.
func (o *SchemeApplicationCredentialsAllOf) GetRevealPassword() bool {
	if o == nil || o.RevealPassword == nil {
		var ret bool
		return ret
	}
	return *o.RevealPassword
}

// GetRevealPasswordOk returns a tuple with the RevealPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemeApplicationCredentialsAllOf) GetRevealPasswordOk() (*bool, bool) {
	if o == nil || o.RevealPassword == nil {
		return nil, false
	}
	return o.RevealPassword, true
}

// HasRevealPassword returns a boolean if a field has been set.
func (o *SchemeApplicationCredentialsAllOf) HasRevealPassword() bool {
	if o != nil && o.RevealPassword != nil {
		return true
	}

	return false
}

// SetRevealPassword gets a reference to the given bool and assigns it to the RevealPassword field.
func (o *SchemeApplicationCredentialsAllOf) SetRevealPassword(v bool) {
	o.RevealPassword = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *SchemeApplicationCredentialsAllOf) GetScheme() ApplicationCredentialsScheme {
	if o == nil || o.Scheme == nil {
		var ret ApplicationCredentialsScheme
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemeApplicationCredentialsAllOf) GetSchemeOk() (*ApplicationCredentialsScheme, bool) {
	if o == nil || o.Scheme == nil {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *SchemeApplicationCredentialsAllOf) HasScheme() bool {
	if o != nil && o.Scheme != nil {
		return true
	}

	return false
}

// SetScheme gets a reference to the given ApplicationCredentialsScheme and assigns it to the Scheme field.
func (o *SchemeApplicationCredentialsAllOf) SetScheme(v ApplicationCredentialsScheme) {
	o.Scheme = &v
}

// GetSigning returns the Signing field value if set, zero value otherwise.
func (o *SchemeApplicationCredentialsAllOf) GetSigning() ApplicationCredentialsSigning {
	if o == nil || o.Signing == nil {
		var ret ApplicationCredentialsSigning
		return ret
	}
	return *o.Signing
}

// GetSigningOk returns a tuple with the Signing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemeApplicationCredentialsAllOf) GetSigningOk() (*ApplicationCredentialsSigning, bool) {
	if o == nil || o.Signing == nil {
		return nil, false
	}
	return o.Signing, true
}

// HasSigning returns a boolean if a field has been set.
func (o *SchemeApplicationCredentialsAllOf) HasSigning() bool {
	if o != nil && o.Signing != nil {
		return true
	}

	return false
}

// SetSigning gets a reference to the given ApplicationCredentialsSigning and assigns it to the Signing field.
func (o *SchemeApplicationCredentialsAllOf) SetSigning(v ApplicationCredentialsSigning) {
	o.Signing = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *SchemeApplicationCredentialsAllOf) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemeApplicationCredentialsAllOf) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *SchemeApplicationCredentialsAllOf) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *SchemeApplicationCredentialsAllOf) SetUserName(v string) {
	o.UserName = &v
}

func (o SchemeApplicationCredentialsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.RevealPassword != nil {
		toSerialize["revealPassword"] = o.RevealPassword
	}
	if o.Scheme != nil {
		toSerialize["scheme"] = o.Scheme
	}
	if o.Signing != nil {
		toSerialize["signing"] = o.Signing
	}
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SchemeApplicationCredentialsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSchemeApplicationCredentialsAllOf := _SchemeApplicationCredentialsAllOf{}

	err = json.Unmarshal(bytes, &varSchemeApplicationCredentialsAllOf)
	if err == nil {
		*o = SchemeApplicationCredentialsAllOf(varSchemeApplicationCredentialsAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "password")
		delete(additionalProperties, "revealPassword")
		delete(additionalProperties, "scheme")
		delete(additionalProperties, "signing")
		delete(additionalProperties, "userName")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSchemeApplicationCredentialsAllOf struct {
	value *SchemeApplicationCredentialsAllOf
	isSet bool
}

func (v NullableSchemeApplicationCredentialsAllOf) Get() *SchemeApplicationCredentialsAllOf {
	return v.value
}

func (v *NullableSchemeApplicationCredentialsAllOf) Set(val *SchemeApplicationCredentialsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemeApplicationCredentialsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemeApplicationCredentialsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemeApplicationCredentialsAllOf(val *SchemeApplicationCredentialsAllOf) *NullableSchemeApplicationCredentialsAllOf {
	return &NullableSchemeApplicationCredentialsAllOf{value: val, isSet: true}
}

func (v NullableSchemeApplicationCredentialsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemeApplicationCredentialsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

