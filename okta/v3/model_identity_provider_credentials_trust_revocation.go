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

// IdentityProviderCredentialsTrustRevocation the model 'IdentityProviderCredentialsTrustRevocation'
type IdentityProviderCredentialsTrustRevocation string

// List of IdentityProviderCredentialsTrustRevocation
const (
	IDENTITYPROVIDERCREDENTIALSTRUSTREVOCATION_CRL IdentityProviderCredentialsTrustRevocation = "CRL"
	IDENTITYPROVIDERCREDENTIALSTRUSTREVOCATION_DELTA_CRL IdentityProviderCredentialsTrustRevocation = "DELTA_CRL"
	IDENTITYPROVIDERCREDENTIALSTRUSTREVOCATION_OCSP IdentityProviderCredentialsTrustRevocation = "OCSP"
)

// All allowed values of IdentityProviderCredentialsTrustRevocation enum
var AllowedIdentityProviderCredentialsTrustRevocationEnumValues = []IdentityProviderCredentialsTrustRevocation{
	"CRL",
	"DELTA_CRL",
	"OCSP",
}

func (v *IdentityProviderCredentialsTrustRevocation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IdentityProviderCredentialsTrustRevocation(value)
	for _, existing := range AllowedIdentityProviderCredentialsTrustRevocationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IdentityProviderCredentialsTrustRevocation", value)
}

// NewIdentityProviderCredentialsTrustRevocationFromValue returns a pointer to a valid IdentityProviderCredentialsTrustRevocation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIdentityProviderCredentialsTrustRevocationFromValue(v string) (*IdentityProviderCredentialsTrustRevocation, error) {
	ev := IdentityProviderCredentialsTrustRevocation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IdentityProviderCredentialsTrustRevocation: valid values are %v", v, AllowedIdentityProviderCredentialsTrustRevocationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IdentityProviderCredentialsTrustRevocation) IsValid() bool {
	for _, existing := range AllowedIdentityProviderCredentialsTrustRevocationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IdentityProviderCredentialsTrustRevocation value
func (v IdentityProviderCredentialsTrustRevocation) Ptr() *IdentityProviderCredentialsTrustRevocation {
	return &v
}

type NullableIdentityProviderCredentialsTrustRevocation struct {
	value *IdentityProviderCredentialsTrustRevocation
	isSet bool
}

func (v NullableIdentityProviderCredentialsTrustRevocation) Get() *IdentityProviderCredentialsTrustRevocation {
	return v.value
}

func (v *NullableIdentityProviderCredentialsTrustRevocation) Set(val *IdentityProviderCredentialsTrustRevocation) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderCredentialsTrustRevocation) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderCredentialsTrustRevocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderCredentialsTrustRevocation(val *IdentityProviderCredentialsTrustRevocation) *NullableIdentityProviderCredentialsTrustRevocation {
	return &NullableIdentityProviderCredentialsTrustRevocation{value: val, isSet: true}
}

func (v NullableIdentityProviderCredentialsTrustRevocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderCredentialsTrustRevocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

