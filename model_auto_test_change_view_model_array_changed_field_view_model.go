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

// checks if the AutoTestChangeViewModelArrayChangedFieldViewModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoTestChangeViewModelArrayChangedFieldViewModel{}

// AutoTestChangeViewModelArrayChangedFieldViewModel struct for AutoTestChangeViewModelArrayChangedFieldViewModel
type AutoTestChangeViewModelArrayChangedFieldViewModel struct {
	OldValue []AutoTestChangeViewModel `json:"oldValue,omitempty"`
	NewValue []AutoTestChangeViewModel `json:"newValue,omitempty"`
}

// NewAutoTestChangeViewModelArrayChangedFieldViewModel instantiates a new AutoTestChangeViewModelArrayChangedFieldViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoTestChangeViewModelArrayChangedFieldViewModel() *AutoTestChangeViewModelArrayChangedFieldViewModel {
	this := AutoTestChangeViewModelArrayChangedFieldViewModel{}
	return &this
}

// NewAutoTestChangeViewModelArrayChangedFieldViewModelWithDefaults instantiates a new AutoTestChangeViewModelArrayChangedFieldViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoTestChangeViewModelArrayChangedFieldViewModelWithDefaults() *AutoTestChangeViewModelArrayChangedFieldViewModel {
	this := AutoTestChangeViewModelArrayChangedFieldViewModel{}
	return &this
}

// GetOldValue returns the OldValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestChangeViewModelArrayChangedFieldViewModel) GetOldValue() []AutoTestChangeViewModel {
	if o == nil {
		var ret []AutoTestChangeViewModel
		return ret
	}
	return o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestChangeViewModelArrayChangedFieldViewModel) GetOldValueOk() ([]AutoTestChangeViewModel, bool) {
	if o == nil || IsNil(o.OldValue) {
		return nil, false
	}
	return o.OldValue, true
}

// HasOldValue returns a boolean if a field has been set.
func (o *AutoTestChangeViewModelArrayChangedFieldViewModel) HasOldValue() bool {
	if o != nil && IsNil(o.OldValue) {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given []AutoTestChangeViewModel and assigns it to the OldValue field.
func (o *AutoTestChangeViewModelArrayChangedFieldViewModel) SetOldValue(v []AutoTestChangeViewModel) {
	o.OldValue = v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestChangeViewModelArrayChangedFieldViewModel) GetNewValue() []AutoTestChangeViewModel {
	if o == nil {
		var ret []AutoTestChangeViewModel
		return ret
	}
	return o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestChangeViewModelArrayChangedFieldViewModel) GetNewValueOk() ([]AutoTestChangeViewModel, bool) {
	if o == nil || IsNil(o.NewValue) {
		return nil, false
	}
	return o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *AutoTestChangeViewModelArrayChangedFieldViewModel) HasNewValue() bool {
	if o != nil && IsNil(o.NewValue) {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given []AutoTestChangeViewModel and assigns it to the NewValue field.
func (o *AutoTestChangeViewModelArrayChangedFieldViewModel) SetNewValue(v []AutoTestChangeViewModel) {
	o.NewValue = v
}

func (o AutoTestChangeViewModelArrayChangedFieldViewModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoTestChangeViewModelArrayChangedFieldViewModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OldValue != nil {
		toSerialize["oldValue"] = o.OldValue
	}
	if o.NewValue != nil {
		toSerialize["newValue"] = o.NewValue
	}
	return toSerialize, nil
}

type NullableAutoTestChangeViewModelArrayChangedFieldViewModel struct {
	value *AutoTestChangeViewModelArrayChangedFieldViewModel
	isSet bool
}

func (v NullableAutoTestChangeViewModelArrayChangedFieldViewModel) Get() *AutoTestChangeViewModelArrayChangedFieldViewModel {
	return v.value
}

func (v *NullableAutoTestChangeViewModelArrayChangedFieldViewModel) Set(val *AutoTestChangeViewModelArrayChangedFieldViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoTestChangeViewModelArrayChangedFieldViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoTestChangeViewModelArrayChangedFieldViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoTestChangeViewModelArrayChangedFieldViewModel(val *AutoTestChangeViewModelArrayChangedFieldViewModel) *NullableAutoTestChangeViewModelArrayChangedFieldViewModel {
	return &NullableAutoTestChangeViewModelArrayChangedFieldViewModel{value: val, isSet: true}
}

func (v NullableAutoTestChangeViewModelArrayChangedFieldViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoTestChangeViewModelArrayChangedFieldViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

