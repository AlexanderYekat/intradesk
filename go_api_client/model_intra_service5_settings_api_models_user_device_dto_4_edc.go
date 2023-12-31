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

// checks if the IntraService5SettingsApiModelsUserDeviceDto4EDC type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntraService5SettingsApiModelsUserDeviceDto4EDC{}

// IntraService5SettingsApiModelsUserDeviceDto4EDC struct for IntraService5SettingsApiModelsUserDeviceDto4EDC
type IntraService5SettingsApiModelsUserDeviceDto4EDC struct {
	DeviceType *IntraService5SharedEnumsDeviceType7442 `json:"deviceType,omitempty"`
	Token NullableString `json:"token,omitempty"`
}

// NewIntraService5SettingsApiModelsUserDeviceDto4EDC instantiates a new IntraService5SettingsApiModelsUserDeviceDto4EDC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntraService5SettingsApiModelsUserDeviceDto4EDC() *IntraService5SettingsApiModelsUserDeviceDto4EDC {
	this := IntraService5SettingsApiModelsUserDeviceDto4EDC{}
	return &this
}

// NewIntraService5SettingsApiModelsUserDeviceDto4EDCWithDefaults instantiates a new IntraService5SettingsApiModelsUserDeviceDto4EDC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntraService5SettingsApiModelsUserDeviceDto4EDCWithDefaults() *IntraService5SettingsApiModelsUserDeviceDto4EDC {
	this := IntraService5SettingsApiModelsUserDeviceDto4EDC{}
	return &this
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *IntraService5SettingsApiModelsUserDeviceDto4EDC) GetDeviceType() IntraService5SharedEnumsDeviceType7442 {
	if o == nil || IsNil(o.DeviceType) {
		var ret IntraService5SharedEnumsDeviceType7442
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntraService5SettingsApiModelsUserDeviceDto4EDC) GetDeviceTypeOk() (*IntraService5SharedEnumsDeviceType7442, bool) {
	if o == nil || IsNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *IntraService5SettingsApiModelsUserDeviceDto4EDC) HasDeviceType() bool {
	if o != nil && !IsNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given IntraService5SharedEnumsDeviceType7442 and assigns it to the DeviceType field.
func (o *IntraService5SettingsApiModelsUserDeviceDto4EDC) SetDeviceType(v IntraService5SharedEnumsDeviceType7442) {
	o.DeviceType = &v
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntraService5SettingsApiModelsUserDeviceDto4EDC) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntraService5SettingsApiModelsUserDeviceDto4EDC) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *IntraService5SettingsApiModelsUserDeviceDto4EDC) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *IntraService5SettingsApiModelsUserDeviceDto4EDC) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *IntraService5SettingsApiModelsUserDeviceDto4EDC) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *IntraService5SettingsApiModelsUserDeviceDto4EDC) UnsetToken() {
	o.Token.Unset()
}

func (o IntraService5SettingsApiModelsUserDeviceDto4EDC) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntraService5SettingsApiModelsUserDeviceDto4EDC) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceType) {
		toSerialize["deviceType"] = o.DeviceType
	}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	return toSerialize, nil
}

type NullableIntraService5SettingsApiModelsUserDeviceDto4EDC struct {
	value *IntraService5SettingsApiModelsUserDeviceDto4EDC
	isSet bool
}

func (v NullableIntraService5SettingsApiModelsUserDeviceDto4EDC) Get() *IntraService5SettingsApiModelsUserDeviceDto4EDC {
	return v.value
}

func (v *NullableIntraService5SettingsApiModelsUserDeviceDto4EDC) Set(val *IntraService5SettingsApiModelsUserDeviceDto4EDC) {
	v.value = val
	v.isSet = true
}

func (v NullableIntraService5SettingsApiModelsUserDeviceDto4EDC) IsSet() bool {
	return v.isSet
}

func (v *NullableIntraService5SettingsApiModelsUserDeviceDto4EDC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntraService5SettingsApiModelsUserDeviceDto4EDC(val *IntraService5SettingsApiModelsUserDeviceDto4EDC) *NullableIntraService5SettingsApiModelsUserDeviceDto4EDC {
	return &NullableIntraService5SettingsApiModelsUserDeviceDto4EDC{value: val, isSet: true}
}

func (v NullableIntraService5SettingsApiModelsUserDeviceDto4EDC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntraService5SettingsApiModelsUserDeviceDto4EDC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


