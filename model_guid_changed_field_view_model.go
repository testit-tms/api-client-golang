/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
)

// checks if the GuidChangedFieldViewModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuidChangedFieldViewModel{}

// GuidChangedFieldViewModel struct for GuidChangedFieldViewModel
type GuidChangedFieldViewModel struct {
	OldValue string `json:"oldValue"`
	NewValue string `json:"newValue"`
}

// NewGuidChangedFieldViewModel instantiates a new GuidChangedFieldViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuidChangedFieldViewModel(oldValue string, newValue string) *GuidChangedFieldViewModel {
	this := GuidChangedFieldViewModel{}
	this.OldValue = oldValue
	this.NewValue = newValue
	return &this
}

// NewGuidChangedFieldViewModelWithDefaults instantiates a new GuidChangedFieldViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuidChangedFieldViewModelWithDefaults() *GuidChangedFieldViewModel {
	this := GuidChangedFieldViewModel{}
	return &this
}

// GetOldValue returns the OldValue field value
func (o *GuidChangedFieldViewModel) GetOldValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value
// and a boolean to check if the value has been set.
func (o *GuidChangedFieldViewModel) GetOldValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OldValue, true
}

// SetOldValue sets field value
func (o *GuidChangedFieldViewModel) SetOldValue(v string) {
	o.OldValue = v
}

// GetNewValue returns the NewValue field value
func (o *GuidChangedFieldViewModel) GetNewValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value
// and a boolean to check if the value has been set.
func (o *GuidChangedFieldViewModel) GetNewValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewValue, true
}

// SetNewValue sets field value
func (o *GuidChangedFieldViewModel) SetNewValue(v string) {
	o.NewValue = v
}

func (o GuidChangedFieldViewModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuidChangedFieldViewModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["oldValue"] = o.OldValue
	toSerialize["newValue"] = o.NewValue
	return toSerialize, nil
}

type NullableGuidChangedFieldViewModel struct {
	value *GuidChangedFieldViewModel
	isSet bool
}

func (v NullableGuidChangedFieldViewModel) Get() *GuidChangedFieldViewModel {
	return v.value
}

func (v *NullableGuidChangedFieldViewModel) Set(val *GuidChangedFieldViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGuidChangedFieldViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGuidChangedFieldViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuidChangedFieldViewModel(val *GuidChangedFieldViewModel) *NullableGuidChangedFieldViewModel {
	return &NullableGuidChangedFieldViewModel{value: val, isSet: true}
}

func (v NullableGuidChangedFieldViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuidChangedFieldViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


