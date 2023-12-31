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

// checks if the IntraService5SettingsApiModelsAdditionalFieldValueModel86F9 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntraService5SettingsApiModelsAdditionalFieldValueModel86F9{}

// IntraService5SettingsApiModelsAdditionalFieldValueModel86F9 struct for IntraService5SettingsApiModelsAdditionalFieldValueModel86F9
type IntraService5SettingsApiModelsAdditionalFieldValueModel86F9 struct {
	Alias NullableString `json:"alias,omitempty"`
	Value NullableString `json:"value,omitempty"`
}

// NewIntraService5SettingsApiModelsAdditionalFieldValueModel86F9 instantiates a new IntraService5SettingsApiModelsAdditionalFieldValueModel86F9 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntraService5SettingsApiModelsAdditionalFieldValueModel86F9() *IntraService5SettingsApiModelsAdditionalFieldValueModel86F9 {
	this := IntraService5SettingsApiModelsAdditionalFieldValueModel86F9{}
	return &this
}

// NewIntraService5SettingsApiModelsAdditionalFieldValueModel86F9WithDefaults instantiates a new IntraService5SettingsApiModelsAdditionalFieldValueModel86F9 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntraService5SettingsApiModelsAdditionalFieldValueModel86F9WithDefaults() *IntraService5SettingsApiModelsAdditionalFieldValueModel86F9 {
	this := IntraService5SettingsApiModelsAdditionalFieldValueModel86F9{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntraService5SettingsApiModelsAdditionalFieldValueModel86F9) GetAlias() string {
	if o == nil || IsNil(o.Alias.Get()) {
		var ret string
		return ret
	}
	return *o.Alias.Get()
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntraService5SettingsApiModelsAdditionalFieldValueModel86F9) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alias.Get(), o.Alias.IsSet()
}

// HasAlias returns a boolean if a field has been set.
func (o *IntraService5SettingsApiModelsAdditionalFieldValueModel86F9) HasAlias() bool {
	if o != nil && o.Alias.IsSet() {
		return true
	}

	return false
}

// SetAlias gets a reference to the given NullableString and assigns it to the Alias field.
func (o *IntraService5SettingsApiModelsAdditionalFieldValueModel86F9) SetAlias(v string) {
	o.Alias.Set(&v)
}
// SetAliasNil sets the value for Alias to be an explicit nil
func (o *IntraService5SettingsApiModelsAdditionalFieldValueModel86F9) SetAliasNil() {
	o.Alias.Set(nil)
}

// UnsetAlias ensures that no value is present for Alias, not even an explicit nil
func (o *IntraService5SettingsApiModelsAdditionalFieldValueModel86F9) UnsetAlias() {
	o.Alias.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntraService5SettingsApiModelsAdditionalFieldValueModel86F9) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntraService5SettingsApiModelsAdditionalFieldValueModel86F9) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *IntraService5SettingsApiModelsAdditionalFieldValueModel86F9) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *IntraService5SettingsApiModelsAdditionalFieldValueModel86F9) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *IntraService5SettingsApiModelsAdditionalFieldValueModel86F9) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *IntraService5SettingsApiModelsAdditionalFieldValueModel86F9) UnsetValue() {
	o.Value.Unset()
}

func (o IntraService5SettingsApiModelsAdditionalFieldValueModel86F9) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntraService5SettingsApiModelsAdditionalFieldValueModel86F9) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Alias.IsSet() {
		toSerialize["alias"] = o.Alias.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableIntraService5SettingsApiModelsAdditionalFieldValueModel86F9 struct {
	value *IntraService5SettingsApiModelsAdditionalFieldValueModel86F9
	isSet bool
}

func (v NullableIntraService5SettingsApiModelsAdditionalFieldValueModel86F9) Get() *IntraService5SettingsApiModelsAdditionalFieldValueModel86F9 {
	return v.value
}

func (v *NullableIntraService5SettingsApiModelsAdditionalFieldValueModel86F9) Set(val *IntraService5SettingsApiModelsAdditionalFieldValueModel86F9) {
	v.value = val
	v.isSet = true
}

func (v NullableIntraService5SettingsApiModelsAdditionalFieldValueModel86F9) IsSet() bool {
	return v.isSet
}

func (v *NullableIntraService5SettingsApiModelsAdditionalFieldValueModel86F9) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntraService5SettingsApiModelsAdditionalFieldValueModel86F9(val *IntraService5SettingsApiModelsAdditionalFieldValueModel86F9) *NullableIntraService5SettingsApiModelsAdditionalFieldValueModel86F9 {
	return &NullableIntraService5SettingsApiModelsAdditionalFieldValueModel86F9{value: val, isSet: true}
}

func (v NullableIntraService5SettingsApiModelsAdditionalFieldValueModel86F9) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntraService5SettingsApiModelsAdditionalFieldValueModel86F9) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


