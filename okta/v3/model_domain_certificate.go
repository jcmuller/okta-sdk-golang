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

// DomainCertificate struct for DomainCertificate
type DomainCertificate struct {
	Certificate *string `json:"certificate,omitempty"`
	CertificateChain *string `json:"certificateChain,omitempty"`
	PrivateKey *string `json:"privateKey,omitempty"`
	Type *DomainCertificateType `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DomainCertificate DomainCertificate

// NewDomainCertificate instantiates a new DomainCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainCertificate() *DomainCertificate {
	this := DomainCertificate{}
	return &this
}

// NewDomainCertificateWithDefaults instantiates a new DomainCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainCertificateWithDefaults() *DomainCertificate {
	this := DomainCertificate{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *DomainCertificate) GetCertificate() string {
	if o == nil || o.Certificate == nil {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainCertificate) GetCertificateOk() (*string, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *DomainCertificate) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *DomainCertificate) SetCertificate(v string) {
	o.Certificate = &v
}

// GetCertificateChain returns the CertificateChain field value if set, zero value otherwise.
func (o *DomainCertificate) GetCertificateChain() string {
	if o == nil || o.CertificateChain == nil {
		var ret string
		return ret
	}
	return *o.CertificateChain
}

// GetCertificateChainOk returns a tuple with the CertificateChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainCertificate) GetCertificateChainOk() (*string, bool) {
	if o == nil || o.CertificateChain == nil {
		return nil, false
	}
	return o.CertificateChain, true
}

// HasCertificateChain returns a boolean if a field has been set.
func (o *DomainCertificate) HasCertificateChain() bool {
	if o != nil && o.CertificateChain != nil {
		return true
	}

	return false
}

// SetCertificateChain gets a reference to the given string and assigns it to the CertificateChain field.
func (o *DomainCertificate) SetCertificateChain(v string) {
	o.CertificateChain = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *DomainCertificate) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainCertificate) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *DomainCertificate) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *DomainCertificate) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DomainCertificate) GetType() DomainCertificateType {
	if o == nil || o.Type == nil {
		var ret DomainCertificateType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainCertificate) GetTypeOk() (*DomainCertificateType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DomainCertificate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given DomainCertificateType and assigns it to the Type field.
func (o *DomainCertificate) SetType(v DomainCertificateType) {
	o.Type = &v
}

func (o DomainCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	if o.CertificateChain != nil {
		toSerialize["certificateChain"] = o.CertificateChain
	}
	if o.PrivateKey != nil {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DomainCertificate) UnmarshalJSON(bytes []byte) (err error) {
	varDomainCertificate := _DomainCertificate{}

	err = json.Unmarshal(bytes, &varDomainCertificate)
	if err == nil {
		*o = DomainCertificate(varDomainCertificate)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "certificateChain")
		delete(additionalProperties, "privateKey")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDomainCertificate struct {
	value *DomainCertificate
	isSet bool
}

func (v NullableDomainCertificate) Get() *DomainCertificate {
	return v.value
}

func (v *NullableDomainCertificate) Set(val *DomainCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainCertificate(val *DomainCertificate) *NullableDomainCertificate {
	return &NullableDomainCertificate{value: val, isSet: true}
}

func (v NullableDomainCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

