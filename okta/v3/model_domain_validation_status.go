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

// DomainValidationStatus the model 'DomainValidationStatus'
type DomainValidationStatus string

// List of DomainValidationStatus
const (
	DOMAINVALIDATIONSTATUS_COMPLETED DomainValidationStatus = "COMPLETED"
	DOMAINVALIDATIONSTATUS_IN_PROGRESS DomainValidationStatus = "IN_PROGRESS"
	DOMAINVALIDATIONSTATUS_NOT_STARTED DomainValidationStatus = "NOT_STARTED"
	DOMAINVALIDATIONSTATUS_VERIFIED DomainValidationStatus = "VERIFIED"
)

// All allowed values of DomainValidationStatus enum
var AllowedDomainValidationStatusEnumValues = []DomainValidationStatus{
	"COMPLETED",
	"IN_PROGRESS",
	"NOT_STARTED",
	"VERIFIED",
}

func (v *DomainValidationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DomainValidationStatus(value)
	for _, existing := range AllowedDomainValidationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DomainValidationStatus", value)
}

// NewDomainValidationStatusFromValue returns a pointer to a valid DomainValidationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDomainValidationStatusFromValue(v string) (*DomainValidationStatus, error) {
	ev := DomainValidationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DomainValidationStatus: valid values are %v", v, AllowedDomainValidationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DomainValidationStatus) IsValid() bool {
	for _, existing := range AllowedDomainValidationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DomainValidationStatus value
func (v DomainValidationStatus) Ptr() *DomainValidationStatus {
	return &v
}

type NullableDomainValidationStatus struct {
	value *DomainValidationStatus
	isSet bool
}

func (v NullableDomainValidationStatus) Get() *DomainValidationStatus {
	return v.value
}

func (v *NullableDomainValidationStatus) Set(val *DomainValidationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainValidationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainValidationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainValidationStatus(val *DomainValidationStatus) *NullableDomainValidationStatus {
	return &NullableDomainValidationStatus{value: val, isSet: true}
}

func (v NullableDomainValidationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainValidationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

