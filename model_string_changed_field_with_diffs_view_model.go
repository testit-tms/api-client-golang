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

// checks if the StringChangedFieldWithDiffsViewModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringChangedFieldWithDiffsViewModel{}

// StringChangedFieldWithDiffsViewModel struct for StringChangedFieldWithDiffsViewModel
type StringChangedFieldWithDiffsViewModel struct {
	DiffValue NullableString `json:"diffValue,omitempty"`
	OldValue NullableString `json:"oldValue,omitempty"`
	NewValue NullableString `json:"newValue,omitempty"`
}

// NewStringChangedFieldWithDiffsViewModel instantiates a new StringChangedFieldWithDiffsViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringChangedFieldWithDiffsViewModel() *StringChangedFieldWithDiffsViewModel {
	this := StringChangedFieldWithDiffsViewModel{}
	return &this
}

// NewStringChangedFieldWithDiffsViewModelWithDefaults instantiates a new StringChangedFieldWithDiffsViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringChangedFieldWithDiffsViewModelWithDefaults() *StringChangedFieldWithDiffsViewModel {
	this := StringChangedFieldWithDiffsViewModel{}
	return &this
}

// GetDiffValue returns the DiffValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StringChangedFieldWithDiffsViewModel) GetDiffValue() string {
	if o == nil || IsNil(o.DiffValue.Get()) {
		var ret string
		return ret
	}
	return *o.DiffValue.Get()
}

// GetDiffValueOk returns a tuple with the DiffValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StringChangedFieldWithDiffsViewModel) GetDiffValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiffValue.Get(), o.DiffValue.IsSet()
}

// HasDiffValue returns a boolean if a field has been set.
func (o *StringChangedFieldWithDiffsViewModel) HasDiffValue() bool {
	if o != nil && o.DiffValue.IsSet() {
		return true
	}

	return false
}

// SetDiffValue gets a reference to the given NullableString and assigns it to the DiffValue field.
func (o *StringChangedFieldWithDiffsViewModel) SetDiffValue(v string) {
	o.DiffValue.Set(&v)
}
// SetDiffValueNil sets the value for DiffValue to be an explicit nil
func (o *StringChangedFieldWithDiffsViewModel) SetDiffValueNil() {
	o.DiffValue.Set(nil)
}

// UnsetDiffValue ensures that no value is present for DiffValue, not even an explicit nil
func (o *StringChangedFieldWithDiffsViewModel) UnsetDiffValue() {
	o.DiffValue.Unset()
}

// GetOldValue returns the OldValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StringChangedFieldWithDiffsViewModel) GetOldValue() string {
	if o == nil || IsNil(o.OldValue.Get()) {
		var ret string
		return ret
	}
	return *o.OldValue.Get()
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StringChangedFieldWithDiffsViewModel) GetOldValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OldValue.Get(), o.OldValue.IsSet()
}

// HasOldValue returns a boolean if a field has been set.
func (o *StringChangedFieldWithDiffsViewModel) HasOldValue() bool {
	if o != nil && o.OldValue.IsSet() {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given NullableString and assigns it to the OldValue field.
func (o *StringChangedFieldWithDiffsViewModel) SetOldValue(v string) {
	o.OldValue.Set(&v)
}
// SetOldValueNil sets the value for OldValue to be an explicit nil
func (o *StringChangedFieldWithDiffsViewModel) SetOldValueNil() {
	o.OldValue.Set(nil)
}

// UnsetOldValue ensures that no value is present for OldValue, not even an explicit nil
func (o *StringChangedFieldWithDiffsViewModel) UnsetOldValue() {
	o.OldValue.Unset()
}

// GetNewValue returns the NewValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StringChangedFieldWithDiffsViewModel) GetNewValue() string {
	if o == nil || IsNil(o.NewValue.Get()) {
		var ret string
		return ret
	}
	return *o.NewValue.Get()
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StringChangedFieldWithDiffsViewModel) GetNewValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewValue.Get(), o.NewValue.IsSet()
}

// HasNewValue returns a boolean if a field has been set.
func (o *StringChangedFieldWithDiffsViewModel) HasNewValue() bool {
	if o != nil && o.NewValue.IsSet() {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given NullableString and assigns it to the NewValue field.
func (o *StringChangedFieldWithDiffsViewModel) SetNewValue(v string) {
	o.NewValue.Set(&v)
}
// SetNewValueNil sets the value for NewValue to be an explicit nil
func (o *StringChangedFieldWithDiffsViewModel) SetNewValueNil() {
	o.NewValue.Set(nil)
}

// UnsetNewValue ensures that no value is present for NewValue, not even an explicit nil
func (o *StringChangedFieldWithDiffsViewModel) UnsetNewValue() {
	o.NewValue.Unset()
}

func (o StringChangedFieldWithDiffsViewModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringChangedFieldWithDiffsViewModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DiffValue.IsSet() {
		toSerialize["diffValue"] = o.DiffValue.Get()
	}
	if o.OldValue.IsSet() {
		toSerialize["oldValue"] = o.OldValue.Get()
	}
	if o.NewValue.IsSet() {
		toSerialize["newValue"] = o.NewValue.Get()
	}
	return toSerialize, nil
}

type NullableStringChangedFieldWithDiffsViewModel struct {
	value *StringChangedFieldWithDiffsViewModel
	isSet bool
}

func (v NullableStringChangedFieldWithDiffsViewModel) Get() *StringChangedFieldWithDiffsViewModel {
	return v.value
}

func (v *NullableStringChangedFieldWithDiffsViewModel) Set(val *StringChangedFieldWithDiffsViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableStringChangedFieldWithDiffsViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableStringChangedFieldWithDiffsViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringChangedFieldWithDiffsViewModel(val *StringChangedFieldWithDiffsViewModel) *NullableStringChangedFieldWithDiffsViewModel {
	return &NullableStringChangedFieldWithDiffsViewModel{value: val, isSet: true}
}

func (v NullableStringChangedFieldWithDiffsViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringChangedFieldWithDiffsViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


