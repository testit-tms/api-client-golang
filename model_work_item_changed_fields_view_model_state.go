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

// checks if the WorkItemChangedFieldsViewModelState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemChangedFieldsViewModelState{}

// WorkItemChangedFieldsViewModelState struct for WorkItemChangedFieldsViewModelState
type WorkItemChangedFieldsViewModelState struct {
	OldValue NullableString `json:"oldValue,omitempty"`
	NewValue NullableString `json:"newValue,omitempty"`
}

// NewWorkItemChangedFieldsViewModelState instantiates a new WorkItemChangedFieldsViewModelState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemChangedFieldsViewModelState() *WorkItemChangedFieldsViewModelState {
	this := WorkItemChangedFieldsViewModelState{}
	return &this
}

// NewWorkItemChangedFieldsViewModelStateWithDefaults instantiates a new WorkItemChangedFieldsViewModelState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemChangedFieldsViewModelStateWithDefaults() *WorkItemChangedFieldsViewModelState {
	this := WorkItemChangedFieldsViewModelState{}
	return &this
}

// GetOldValue returns the OldValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemChangedFieldsViewModelState) GetOldValue() string {
	if o == nil || IsNil(o.OldValue.Get()) {
		var ret string
		return ret
	}
	return *o.OldValue.Get()
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemChangedFieldsViewModelState) GetOldValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OldValue.Get(), o.OldValue.IsSet()
}

// HasOldValue returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModelState) HasOldValue() bool {
	if o != nil && o.OldValue.IsSet() {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given NullableString and assigns it to the OldValue field.
func (o *WorkItemChangedFieldsViewModelState) SetOldValue(v string) {
	o.OldValue.Set(&v)
}
// SetOldValueNil sets the value for OldValue to be an explicit nil
func (o *WorkItemChangedFieldsViewModelState) SetOldValueNil() {
	o.OldValue.Set(nil)
}

// UnsetOldValue ensures that no value is present for OldValue, not even an explicit nil
func (o *WorkItemChangedFieldsViewModelState) UnsetOldValue() {
	o.OldValue.Unset()
}

// GetNewValue returns the NewValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemChangedFieldsViewModelState) GetNewValue() string {
	if o == nil || IsNil(o.NewValue.Get()) {
		var ret string
		return ret
	}
	return *o.NewValue.Get()
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemChangedFieldsViewModelState) GetNewValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewValue.Get(), o.NewValue.IsSet()
}

// HasNewValue returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModelState) HasNewValue() bool {
	if o != nil && o.NewValue.IsSet() {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given NullableString and assigns it to the NewValue field.
func (o *WorkItemChangedFieldsViewModelState) SetNewValue(v string) {
	o.NewValue.Set(&v)
}
// SetNewValueNil sets the value for NewValue to be an explicit nil
func (o *WorkItemChangedFieldsViewModelState) SetNewValueNil() {
	o.NewValue.Set(nil)
}

// UnsetNewValue ensures that no value is present for NewValue, not even an explicit nil
func (o *WorkItemChangedFieldsViewModelState) UnsetNewValue() {
	o.NewValue.Unset()
}

func (o WorkItemChangedFieldsViewModelState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemChangedFieldsViewModelState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OldValue.IsSet() {
		toSerialize["oldValue"] = o.OldValue.Get()
	}
	if o.NewValue.IsSet() {
		toSerialize["newValue"] = o.NewValue.Get()
	}
	return toSerialize, nil
}

type NullableWorkItemChangedFieldsViewModelState struct {
	value *WorkItemChangedFieldsViewModelState
	isSet bool
}

func (v NullableWorkItemChangedFieldsViewModelState) Get() *WorkItemChangedFieldsViewModelState {
	return v.value
}

func (v *NullableWorkItemChangedFieldsViewModelState) Set(val *WorkItemChangedFieldsViewModelState) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemChangedFieldsViewModelState) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemChangedFieldsViewModelState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemChangedFieldsViewModelState(val *WorkItemChangedFieldsViewModelState) *NullableWorkItemChangedFieldsViewModelState {
	return &NullableWorkItemChangedFieldsViewModelState{value: val, isSet: true}
}

func (v NullableWorkItemChangedFieldsViewModelState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemChangedFieldsViewModelState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


