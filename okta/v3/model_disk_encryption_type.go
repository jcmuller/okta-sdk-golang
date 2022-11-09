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

// DiskEncryptionType the model 'DiskEncryptionType'
type DiskEncryptionType string

// List of DiskEncryptionType
const (
	DISKENCRYPTIONTYPE_ALL_INTERNAL_VOLUMES DiskEncryptionType = "ALL_INTERNAL_VOLUMES"
	DISKENCRYPTIONTYPE_FULL DiskEncryptionType = "FULL"
	DISKENCRYPTIONTYPE_USER DiskEncryptionType = "USER"
)

// All allowed values of DiskEncryptionType enum
var AllowedDiskEncryptionTypeEnumValues = []DiskEncryptionType{
	"ALL_INTERNAL_VOLUMES",
	"FULL",
	"USER",
}

func (v *DiskEncryptionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DiskEncryptionType(value)
	for _, existing := range AllowedDiskEncryptionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DiskEncryptionType", value)
}

// NewDiskEncryptionTypeFromValue returns a pointer to a valid DiskEncryptionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDiskEncryptionTypeFromValue(v string) (*DiskEncryptionType, error) {
	ev := DiskEncryptionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DiskEncryptionType: valid values are %v", v, AllowedDiskEncryptionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DiskEncryptionType) IsValid() bool {
	for _, existing := range AllowedDiskEncryptionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiskEncryptionType value
func (v DiskEncryptionType) Ptr() *DiskEncryptionType {
	return &v
}

type NullableDiskEncryptionType struct {
	value *DiskEncryptionType
	isSet bool
}

func (v NullableDiskEncryptionType) Get() *DiskEncryptionType {
	return v.value
}

func (v *NullableDiskEncryptionType) Set(val *DiskEncryptionType) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskEncryptionType) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskEncryptionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskEncryptionType(val *DiskEncryptionType) *NullableDiskEncryptionType {
	return &NullableDiskEncryptionType{value: val, isSet: true}
}

func (v NullableDiskEncryptionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskEncryptionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

