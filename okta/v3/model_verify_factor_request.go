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

// VerifyFactorRequest struct for VerifyFactorRequest
type VerifyFactorRequest struct {
	ActivationToken *string `json:"activationToken,omitempty"`
	Answer *string `json:"answer,omitempty"`
	Attestation *string `json:"attestation,omitempty"`
	ClientData *string `json:"clientData,omitempty"`
	NextPassCode *string `json:"nextPassCode,omitempty"`
	PassCode *string `json:"passCode,omitempty"`
	RegistrationData *string `json:"registrationData,omitempty"`
	StateToken *string `json:"stateToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VerifyFactorRequest VerifyFactorRequest

// NewVerifyFactorRequest instantiates a new VerifyFactorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyFactorRequest() *VerifyFactorRequest {
	this := VerifyFactorRequest{}
	return &this
}

// NewVerifyFactorRequestWithDefaults instantiates a new VerifyFactorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyFactorRequestWithDefaults() *VerifyFactorRequest {
	this := VerifyFactorRequest{}
	return &this
}

// GetActivationToken returns the ActivationToken field value if set, zero value otherwise.
func (o *VerifyFactorRequest) GetActivationToken() string {
	if o == nil || o.ActivationToken == nil {
		var ret string
		return ret
	}
	return *o.ActivationToken
}

// GetActivationTokenOk returns a tuple with the ActivationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyFactorRequest) GetActivationTokenOk() (*string, bool) {
	if o == nil || o.ActivationToken == nil {
		return nil, false
	}
	return o.ActivationToken, true
}

// HasActivationToken returns a boolean if a field has been set.
func (o *VerifyFactorRequest) HasActivationToken() bool {
	if o != nil && o.ActivationToken != nil {
		return true
	}

	return false
}

// SetActivationToken gets a reference to the given string and assigns it to the ActivationToken field.
func (o *VerifyFactorRequest) SetActivationToken(v string) {
	o.ActivationToken = &v
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *VerifyFactorRequest) GetAnswer() string {
	if o == nil || o.Answer == nil {
		var ret string
		return ret
	}
	return *o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyFactorRequest) GetAnswerOk() (*string, bool) {
	if o == nil || o.Answer == nil {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *VerifyFactorRequest) HasAnswer() bool {
	if o != nil && o.Answer != nil {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given string and assigns it to the Answer field.
func (o *VerifyFactorRequest) SetAnswer(v string) {
	o.Answer = &v
}

// GetAttestation returns the Attestation field value if set, zero value otherwise.
func (o *VerifyFactorRequest) GetAttestation() string {
	if o == nil || o.Attestation == nil {
		var ret string
		return ret
	}
	return *o.Attestation
}

// GetAttestationOk returns a tuple with the Attestation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyFactorRequest) GetAttestationOk() (*string, bool) {
	if o == nil || o.Attestation == nil {
		return nil, false
	}
	return o.Attestation, true
}

// HasAttestation returns a boolean if a field has been set.
func (o *VerifyFactorRequest) HasAttestation() bool {
	if o != nil && o.Attestation != nil {
		return true
	}

	return false
}

// SetAttestation gets a reference to the given string and assigns it to the Attestation field.
func (o *VerifyFactorRequest) SetAttestation(v string) {
	o.Attestation = &v
}

// GetClientData returns the ClientData field value if set, zero value otherwise.
func (o *VerifyFactorRequest) GetClientData() string {
	if o == nil || o.ClientData == nil {
		var ret string
		return ret
	}
	return *o.ClientData
}

// GetClientDataOk returns a tuple with the ClientData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyFactorRequest) GetClientDataOk() (*string, bool) {
	if o == nil || o.ClientData == nil {
		return nil, false
	}
	return o.ClientData, true
}

// HasClientData returns a boolean if a field has been set.
func (o *VerifyFactorRequest) HasClientData() bool {
	if o != nil && o.ClientData != nil {
		return true
	}

	return false
}

// SetClientData gets a reference to the given string and assigns it to the ClientData field.
func (o *VerifyFactorRequest) SetClientData(v string) {
	o.ClientData = &v
}

// GetNextPassCode returns the NextPassCode field value if set, zero value otherwise.
func (o *VerifyFactorRequest) GetNextPassCode() string {
	if o == nil || o.NextPassCode == nil {
		var ret string
		return ret
	}
	return *o.NextPassCode
}

// GetNextPassCodeOk returns a tuple with the NextPassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyFactorRequest) GetNextPassCodeOk() (*string, bool) {
	if o == nil || o.NextPassCode == nil {
		return nil, false
	}
	return o.NextPassCode, true
}

// HasNextPassCode returns a boolean if a field has been set.
func (o *VerifyFactorRequest) HasNextPassCode() bool {
	if o != nil && o.NextPassCode != nil {
		return true
	}

	return false
}

// SetNextPassCode gets a reference to the given string and assigns it to the NextPassCode field.
func (o *VerifyFactorRequest) SetNextPassCode(v string) {
	o.NextPassCode = &v
}

// GetPassCode returns the PassCode field value if set, zero value otherwise.
func (o *VerifyFactorRequest) GetPassCode() string {
	if o == nil || o.PassCode == nil {
		var ret string
		return ret
	}
	return *o.PassCode
}

// GetPassCodeOk returns a tuple with the PassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyFactorRequest) GetPassCodeOk() (*string, bool) {
	if o == nil || o.PassCode == nil {
		return nil, false
	}
	return o.PassCode, true
}

// HasPassCode returns a boolean if a field has been set.
func (o *VerifyFactorRequest) HasPassCode() bool {
	if o != nil && o.PassCode != nil {
		return true
	}

	return false
}

// SetPassCode gets a reference to the given string and assigns it to the PassCode field.
func (o *VerifyFactorRequest) SetPassCode(v string) {
	o.PassCode = &v
}

// GetRegistrationData returns the RegistrationData field value if set, zero value otherwise.
func (o *VerifyFactorRequest) GetRegistrationData() string {
	if o == nil || o.RegistrationData == nil {
		var ret string
		return ret
	}
	return *o.RegistrationData
}

// GetRegistrationDataOk returns a tuple with the RegistrationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyFactorRequest) GetRegistrationDataOk() (*string, bool) {
	if o == nil || o.RegistrationData == nil {
		return nil, false
	}
	return o.RegistrationData, true
}

// HasRegistrationData returns a boolean if a field has been set.
func (o *VerifyFactorRequest) HasRegistrationData() bool {
	if o != nil && o.RegistrationData != nil {
		return true
	}

	return false
}

// SetRegistrationData gets a reference to the given string and assigns it to the RegistrationData field.
func (o *VerifyFactorRequest) SetRegistrationData(v string) {
	o.RegistrationData = &v
}

// GetStateToken returns the StateToken field value if set, zero value otherwise.
func (o *VerifyFactorRequest) GetStateToken() string {
	if o == nil || o.StateToken == nil {
		var ret string
		return ret
	}
	return *o.StateToken
}

// GetStateTokenOk returns a tuple with the StateToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyFactorRequest) GetStateTokenOk() (*string, bool) {
	if o == nil || o.StateToken == nil {
		return nil, false
	}
	return o.StateToken, true
}

// HasStateToken returns a boolean if a field has been set.
func (o *VerifyFactorRequest) HasStateToken() bool {
	if o != nil && o.StateToken != nil {
		return true
	}

	return false
}

// SetStateToken gets a reference to the given string and assigns it to the StateToken field.
func (o *VerifyFactorRequest) SetStateToken(v string) {
	o.StateToken = &v
}

func (o VerifyFactorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActivationToken != nil {
		toSerialize["activationToken"] = o.ActivationToken
	}
	if o.Answer != nil {
		toSerialize["answer"] = o.Answer
	}
	if o.Attestation != nil {
		toSerialize["attestation"] = o.Attestation
	}
	if o.ClientData != nil {
		toSerialize["clientData"] = o.ClientData
	}
	if o.NextPassCode != nil {
		toSerialize["nextPassCode"] = o.NextPassCode
	}
	if o.PassCode != nil {
		toSerialize["passCode"] = o.PassCode
	}
	if o.RegistrationData != nil {
		toSerialize["registrationData"] = o.RegistrationData
	}
	if o.StateToken != nil {
		toSerialize["stateToken"] = o.StateToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VerifyFactorRequest) UnmarshalJSON(bytes []byte) (err error) {
	varVerifyFactorRequest := _VerifyFactorRequest{}

	err = json.Unmarshal(bytes, &varVerifyFactorRequest)
	if err == nil {
		*o = VerifyFactorRequest(varVerifyFactorRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "activationToken")
		delete(additionalProperties, "answer")
		delete(additionalProperties, "attestation")
		delete(additionalProperties, "clientData")
		delete(additionalProperties, "nextPassCode")
		delete(additionalProperties, "passCode")
		delete(additionalProperties, "registrationData")
		delete(additionalProperties, "stateToken")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableVerifyFactorRequest struct {
	value *VerifyFactorRequest
	isSet bool
}

func (v NullableVerifyFactorRequest) Get() *VerifyFactorRequest {
	return v.value
}

func (v *NullableVerifyFactorRequest) Set(val *VerifyFactorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyFactorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyFactorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyFactorRequest(val *VerifyFactorRequest) *NullableVerifyFactorRequest {
	return &NullableVerifyFactorRequest{value: val, isSet: true}
}

func (v NullableVerifyFactorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyFactorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

