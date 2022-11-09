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

// BehaviorRuleSettingsAnomalousIPAllOf struct for BehaviorRuleSettingsAnomalousIPAllOf
type BehaviorRuleSettingsAnomalousIPAllOf struct {
	MaxEventsUsedForEvaluation *int32 `json:"maxEventsUsedForEvaluation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BehaviorRuleSettingsAnomalousIPAllOf BehaviorRuleSettingsAnomalousIPAllOf

// NewBehaviorRuleSettingsAnomalousIPAllOf instantiates a new BehaviorRuleSettingsAnomalousIPAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBehaviorRuleSettingsAnomalousIPAllOf() *BehaviorRuleSettingsAnomalousIPAllOf {
	this := BehaviorRuleSettingsAnomalousIPAllOf{}
	var maxEventsUsedForEvaluation int32 = 50
	this.MaxEventsUsedForEvaluation = &maxEventsUsedForEvaluation
	return &this
}

// NewBehaviorRuleSettingsAnomalousIPAllOfWithDefaults instantiates a new BehaviorRuleSettingsAnomalousIPAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBehaviorRuleSettingsAnomalousIPAllOfWithDefaults() *BehaviorRuleSettingsAnomalousIPAllOf {
	this := BehaviorRuleSettingsAnomalousIPAllOf{}
	var maxEventsUsedForEvaluation int32 = 50
	this.MaxEventsUsedForEvaluation = &maxEventsUsedForEvaluation
	return &this
}

// GetMaxEventsUsedForEvaluation returns the MaxEventsUsedForEvaluation field value if set, zero value otherwise.
func (o *BehaviorRuleSettingsAnomalousIPAllOf) GetMaxEventsUsedForEvaluation() int32 {
	if o == nil || o.MaxEventsUsedForEvaluation == nil {
		var ret int32
		return ret
	}
	return *o.MaxEventsUsedForEvaluation
}

// GetMaxEventsUsedForEvaluationOk returns a tuple with the MaxEventsUsedForEvaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRuleSettingsAnomalousIPAllOf) GetMaxEventsUsedForEvaluationOk() (*int32, bool) {
	if o == nil || o.MaxEventsUsedForEvaluation == nil {
		return nil, false
	}
	return o.MaxEventsUsedForEvaluation, true
}

// HasMaxEventsUsedForEvaluation returns a boolean if a field has been set.
func (o *BehaviorRuleSettingsAnomalousIPAllOf) HasMaxEventsUsedForEvaluation() bool {
	if o != nil && o.MaxEventsUsedForEvaluation != nil {
		return true
	}

	return false
}

// SetMaxEventsUsedForEvaluation gets a reference to the given int32 and assigns it to the MaxEventsUsedForEvaluation field.
func (o *BehaviorRuleSettingsAnomalousIPAllOf) SetMaxEventsUsedForEvaluation(v int32) {
	o.MaxEventsUsedForEvaluation = &v
}

func (o BehaviorRuleSettingsAnomalousIPAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxEventsUsedForEvaluation != nil {
		toSerialize["maxEventsUsedForEvaluation"] = o.MaxEventsUsedForEvaluation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BehaviorRuleSettingsAnomalousIPAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBehaviorRuleSettingsAnomalousIPAllOf := _BehaviorRuleSettingsAnomalousIPAllOf{}

	err = json.Unmarshal(bytes, &varBehaviorRuleSettingsAnomalousIPAllOf)
	if err == nil {
		*o = BehaviorRuleSettingsAnomalousIPAllOf(varBehaviorRuleSettingsAnomalousIPAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "maxEventsUsedForEvaluation")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBehaviorRuleSettingsAnomalousIPAllOf struct {
	value *BehaviorRuleSettingsAnomalousIPAllOf
	isSet bool
}

func (v NullableBehaviorRuleSettingsAnomalousIPAllOf) Get() *BehaviorRuleSettingsAnomalousIPAllOf {
	return v.value
}

func (v *NullableBehaviorRuleSettingsAnomalousIPAllOf) Set(val *BehaviorRuleSettingsAnomalousIPAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBehaviorRuleSettingsAnomalousIPAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBehaviorRuleSettingsAnomalousIPAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBehaviorRuleSettingsAnomalousIPAllOf(val *BehaviorRuleSettingsAnomalousIPAllOf) *NullableBehaviorRuleSettingsAnomalousIPAllOf {
	return &NullableBehaviorRuleSettingsAnomalousIPAllOf{value: val, isSet: true}
}

func (v NullableBehaviorRuleSettingsAnomalousIPAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBehaviorRuleSettingsAnomalousIPAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

