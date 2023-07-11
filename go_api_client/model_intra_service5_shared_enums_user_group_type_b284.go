/*
Settings API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// IntraService5SharedEnumsUserGroupTypeB284 the model 'IntraService5SharedEnumsUserGroupTypeB284'
type IntraService5SharedEnumsUserGroupTypeB284 int32

// List of IntraService5.Shared.Enums.UserGroupType_B284
const (
	_10 IntraService5SharedEnumsUserGroupTypeB284 = 10
	_20 IntraService5SharedEnumsUserGroupTypeB284 = 20
	_30 IntraService5SharedEnumsUserGroupTypeB284 = 30
	_40 IntraService5SharedEnumsUserGroupTypeB284 = 40
)

// All allowed values of IntraService5SharedEnumsUserGroupTypeB284 enum
var AllowedIntraService5SharedEnumsUserGroupTypeB284EnumValues = []IntraService5SharedEnumsUserGroupTypeB284{
	10,
	20,
	30,
	40,
}

func (v *IntraService5SharedEnumsUserGroupTypeB284) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IntraService5SharedEnumsUserGroupTypeB284(value)
	for _, existing := range AllowedIntraService5SharedEnumsUserGroupTypeB284EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IntraService5SharedEnumsUserGroupTypeB284", value)
}

// NewIntraService5SharedEnumsUserGroupTypeB284FromValue returns a pointer to a valid IntraService5SharedEnumsUserGroupTypeB284
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIntraService5SharedEnumsUserGroupTypeB284FromValue(v int32) (*IntraService5SharedEnumsUserGroupTypeB284, error) {
	ev := IntraService5SharedEnumsUserGroupTypeB284(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IntraService5SharedEnumsUserGroupTypeB284: valid values are %v", v, AllowedIntraService5SharedEnumsUserGroupTypeB284EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IntraService5SharedEnumsUserGroupTypeB284) IsValid() bool {
	for _, existing := range AllowedIntraService5SharedEnumsUserGroupTypeB284EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntraService5.Shared.Enums.UserGroupType_B284 value
func (v IntraService5SharedEnumsUserGroupTypeB284) Ptr() *IntraService5SharedEnumsUserGroupTypeB284 {
	return &v
}

type NullableIntraService5SharedEnumsUserGroupTypeB284 struct {
	value *IntraService5SharedEnumsUserGroupTypeB284
	isSet bool
}

func (v NullableIntraService5SharedEnumsUserGroupTypeB284) Get() *IntraService5SharedEnumsUserGroupTypeB284 {
	return v.value
}

func (v *NullableIntraService5SharedEnumsUserGroupTypeB284) Set(val *IntraService5SharedEnumsUserGroupTypeB284) {
	v.value = val
	v.isSet = true
}

func (v NullableIntraService5SharedEnumsUserGroupTypeB284) IsSet() bool {
	return v.isSet
}

func (v *NullableIntraService5SharedEnumsUserGroupTypeB284) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntraService5SharedEnumsUserGroupTypeB284(val *IntraService5SharedEnumsUserGroupTypeB284) *NullableIntraService5SharedEnumsUserGroupTypeB284 {
	return &NullableIntraService5SharedEnumsUserGroupTypeB284{value: val, isSet: true}
}

func (v NullableIntraService5SharedEnumsUserGroupTypeB284) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntraService5SharedEnumsUserGroupTypeB284) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
