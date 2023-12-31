/*
Settings API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the IntraService5SettingsApiModelsClientUserToEmployeeModel1D42 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntraService5SettingsApiModelsClientUserToEmployeeModel1D42{}

// IntraService5SettingsApiModelsClientUserToEmployeeModel1D42 struct for IntraService5SettingsApiModelsClientUserToEmployeeModel1D42
type IntraService5SettingsApiModelsClientUserToEmployeeModel1D42 struct {
	UserId *int64 `json:"userId,omitempty"`
	RoleId NullableInt64 `json:"roleId,omitempty"`
}

// NewIntraService5SettingsApiModelsClientUserToEmployeeModel1D42 instantiates a new IntraService5SettingsApiModelsClientUserToEmployeeModel1D42 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntraService5SettingsApiModelsClientUserToEmployeeModel1D42() *IntraService5SettingsApiModelsClientUserToEmployeeModel1D42 {
	this := IntraService5SettingsApiModelsClientUserToEmployeeModel1D42{}
	return &this
}

// NewIntraService5SettingsApiModelsClientUserToEmployeeModel1D42WithDefaults instantiates a new IntraService5SettingsApiModelsClientUserToEmployeeModel1D42 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntraService5SettingsApiModelsClientUserToEmployeeModel1D42WithDefaults() *IntraService5SettingsApiModelsClientUserToEmployeeModel1D42 {
	this := IntraService5SettingsApiModelsClientUserToEmployeeModel1D42{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *IntraService5SettingsApiModelsClientUserToEmployeeModel1D42) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntraService5SettingsApiModelsClientUserToEmployeeModel1D42) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *IntraService5SettingsApiModelsClientUserToEmployeeModel1D42) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *IntraService5SettingsApiModelsClientUserToEmployeeModel1D42) SetUserId(v int64) {
	o.UserId = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntraService5SettingsApiModelsClientUserToEmployeeModel1D42) GetRoleId() int64 {
	if o == nil || IsNil(o.RoleId.Get()) {
		var ret int64
		return ret
	}
	return *o.RoleId.Get()
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntraService5SettingsApiModelsClientUserToEmployeeModel1D42) GetRoleIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleId.Get(), o.RoleId.IsSet()
}

// HasRoleId returns a boolean if a field has been set.
func (o *IntraService5SettingsApiModelsClientUserToEmployeeModel1D42) HasRoleId() bool {
	if o != nil && o.RoleId.IsSet() {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given NullableInt64 and assigns it to the RoleId field.
func (o *IntraService5SettingsApiModelsClientUserToEmployeeModel1D42) SetRoleId(v int64) {
	o.RoleId.Set(&v)
}
// SetRoleIdNil sets the value for RoleId to be an explicit nil
func (o *IntraService5SettingsApiModelsClientUserToEmployeeModel1D42) SetRoleIdNil() {
	o.RoleId.Set(nil)
}

// UnsetRoleId ensures that no value is present for RoleId, not even an explicit nil
func (o *IntraService5SettingsApiModelsClientUserToEmployeeModel1D42) UnsetRoleId() {
	o.RoleId.Unset()
}

func (o IntraService5SettingsApiModelsClientUserToEmployeeModel1D42) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntraService5SettingsApiModelsClientUserToEmployeeModel1D42) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if o.RoleId.IsSet() {
		toSerialize["roleId"] = o.RoleId.Get()
	}
	return toSerialize, nil
}

type NullableIntraService5SettingsApiModelsClientUserToEmployeeModel1D42 struct {
	value *IntraService5SettingsApiModelsClientUserToEmployeeModel1D42
	isSet bool
}

func (v NullableIntraService5SettingsApiModelsClientUserToEmployeeModel1D42) Get() *IntraService5SettingsApiModelsClientUserToEmployeeModel1D42 {
	return v.value
}

func (v *NullableIntraService5SettingsApiModelsClientUserToEmployeeModel1D42) Set(val *IntraService5SettingsApiModelsClientUserToEmployeeModel1D42) {
	v.value = val
	v.isSet = true
}

func (v NullableIntraService5SettingsApiModelsClientUserToEmployeeModel1D42) IsSet() bool {
	return v.isSet
}

func (v *NullableIntraService5SettingsApiModelsClientUserToEmployeeModel1D42) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntraService5SettingsApiModelsClientUserToEmployeeModel1D42(val *IntraService5SettingsApiModelsClientUserToEmployeeModel1D42) *NullableIntraService5SettingsApiModelsClientUserToEmployeeModel1D42 {
	return &NullableIntraService5SettingsApiModelsClientUserToEmployeeModel1D42{value: val, isSet: true}
}

func (v NullableIntraService5SettingsApiModelsClientUserToEmployeeModel1D42) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntraService5SettingsApiModelsClientUserToEmployeeModel1D42) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


