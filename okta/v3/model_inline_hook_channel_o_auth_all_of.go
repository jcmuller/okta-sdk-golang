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

// InlineHookChannelOAuthAllOf struct for InlineHookChannelOAuthAllOf
type InlineHookChannelOAuthAllOf struct {
	Config *InlineHookOAuthChannelConfig `json:"config,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookChannelOAuthAllOf InlineHookChannelOAuthAllOf

// NewInlineHookChannelOAuthAllOf instantiates a new InlineHookChannelOAuthAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookChannelOAuthAllOf() *InlineHookChannelOAuthAllOf {
	this := InlineHookChannelOAuthAllOf{}
	return &this
}

// NewInlineHookChannelOAuthAllOfWithDefaults instantiates a new InlineHookChannelOAuthAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookChannelOAuthAllOfWithDefaults() *InlineHookChannelOAuthAllOf {
	this := InlineHookChannelOAuthAllOf{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *InlineHookChannelOAuthAllOf) GetConfig() InlineHookOAuthChannelConfig {
	if o == nil || o.Config == nil {
		var ret InlineHookOAuthChannelConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannelOAuthAllOf) GetConfigOk() (*InlineHookOAuthChannelConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *InlineHookChannelOAuthAllOf) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given InlineHookOAuthChannelConfig and assigns it to the Config field.
func (o *InlineHookChannelOAuthAllOf) SetConfig(v InlineHookOAuthChannelConfig) {
	o.Config = &v
}

func (o InlineHookChannelOAuthAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InlineHookChannelOAuthAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varInlineHookChannelOAuthAllOf := _InlineHookChannelOAuthAllOf{}

	err = json.Unmarshal(bytes, &varInlineHookChannelOAuthAllOf)
	if err == nil {
		*o = InlineHookChannelOAuthAllOf(varInlineHookChannelOAuthAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "config")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableInlineHookChannelOAuthAllOf struct {
	value *InlineHookChannelOAuthAllOf
	isSet bool
}

func (v NullableInlineHookChannelOAuthAllOf) Get() *InlineHookChannelOAuthAllOf {
	return v.value
}

func (v *NullableInlineHookChannelOAuthAllOf) Set(val *InlineHookChannelOAuthAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookChannelOAuthAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookChannelOAuthAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookChannelOAuthAllOf(val *InlineHookChannelOAuthAllOf) *NullableInlineHookChannelOAuthAllOf {
	return &NullableInlineHookChannelOAuthAllOf{value: val, isSet: true}
}

func (v NullableInlineHookChannelOAuthAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookChannelOAuthAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

