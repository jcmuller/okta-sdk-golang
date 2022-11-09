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

// LogStreamLinks struct for LogStreamLinks
type LogStreamLinks struct {
	Self *HrefObject `json:"self,omitempty"`
	Activate *HrefObject `json:"activate,omitempty"`
	Deactivate *HrefObject `json:"deactivate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogStreamLinks LogStreamLinks

// NewLogStreamLinks instantiates a new LogStreamLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogStreamLinks() *LogStreamLinks {
	this := LogStreamLinks{}
	return &this
}

// NewLogStreamLinksWithDefaults instantiates a new LogStreamLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogStreamLinksWithDefaults() *LogStreamLinks {
	this := LogStreamLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *LogStreamLinks) GetSelf() HrefObject {
	if o == nil || o.Self == nil {
		var ret HrefObject
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamLinks) GetSelfOk() (*HrefObject, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *LogStreamLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObject and assigns it to the Self field.
func (o *LogStreamLinks) SetSelf(v HrefObject) {
	o.Self = &v
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *LogStreamLinks) GetActivate() HrefObject {
	if o == nil || o.Activate == nil {
		var ret HrefObject
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamLinks) GetActivateOk() (*HrefObject, bool) {
	if o == nil || o.Activate == nil {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *LogStreamLinks) HasActivate() bool {
	if o != nil && o.Activate != nil {
		return true
	}

	return false
}

// SetActivate gets a reference to the given HrefObject and assigns it to the Activate field.
func (o *LogStreamLinks) SetActivate(v HrefObject) {
	o.Activate = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *LogStreamLinks) GetDeactivate() HrefObject {
	if o == nil || o.Deactivate == nil {
		var ret HrefObject
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogStreamLinks) GetDeactivateOk() (*HrefObject, bool) {
	if o == nil || o.Deactivate == nil {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *LogStreamLinks) HasDeactivate() bool {
	if o != nil && o.Deactivate != nil {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given HrefObject and assigns it to the Deactivate field.
func (o *LogStreamLinks) SetDeactivate(v HrefObject) {
	o.Deactivate = &v
}

func (o LogStreamLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Activate != nil {
		toSerialize["activate"] = o.Activate
	}
	if o.Deactivate != nil {
		toSerialize["deactivate"] = o.Deactivate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogStreamLinks) UnmarshalJSON(bytes []byte) (err error) {
	varLogStreamLinks := _LogStreamLinks{}

	err = json.Unmarshal(bytes, &varLogStreamLinks)
	if err == nil {
		*o = LogStreamLinks(varLogStreamLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "activate")
		delete(additionalProperties, "deactivate")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLogStreamLinks struct {
	value *LogStreamLinks
	isSet bool
}

func (v NullableLogStreamLinks) Get() *LogStreamLinks {
	return v.value
}

func (v *NullableLogStreamLinks) Set(val *LogStreamLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableLogStreamLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableLogStreamLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogStreamLinks(val *LogStreamLinks) *NullableLogStreamLinks {
	return &NullableLogStreamLinks{value: val, isSet: true}
}

func (v NullableLogStreamLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogStreamLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

