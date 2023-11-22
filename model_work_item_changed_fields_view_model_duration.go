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

// checks if the WorkItemChangedFieldsViewModelDuration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemChangedFieldsViewModelDuration{}

// WorkItemChangedFieldsViewModelDuration struct for WorkItemChangedFieldsViewModelDuration
type WorkItemChangedFieldsViewModelDuration struct {
	OldValue int32 `json:"oldValue"`
	NewValue int32 `json:"newValue"`
}

// NewWorkItemChangedFieldsViewModelDuration instantiates a new WorkItemChangedFieldsViewModelDuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemChangedFieldsViewModelDuration(oldValue int32, newValue int32) *WorkItemChangedFieldsViewModelDuration {
	this := WorkItemChangedFieldsViewModelDuration{}
	this.OldValue = oldValue
	this.NewValue = newValue
	return &this
}

// NewWorkItemChangedFieldsViewModelDurationWithDefaults instantiates a new WorkItemChangedFieldsViewModelDuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemChangedFieldsViewModelDurationWithDefaults() *WorkItemChangedFieldsViewModelDuration {
	this := WorkItemChangedFieldsViewModelDuration{}
	return &this
}

// GetOldValue returns the OldValue field value
func (o *WorkItemChangedFieldsViewModelDuration) GetOldValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModelDuration) GetOldValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OldValue, true
}

// SetOldValue sets field value
func (o *WorkItemChangedFieldsViewModelDuration) SetOldValue(v int32) {
	o.OldValue = v
}

// GetNewValue returns the NewValue field value
func (o *WorkItemChangedFieldsViewModelDuration) GetNewValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModelDuration) GetNewValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewValue, true
}

// SetNewValue sets field value
func (o *WorkItemChangedFieldsViewModelDuration) SetNewValue(v int32) {
	o.NewValue = v
}

func (o WorkItemChangedFieldsViewModelDuration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemChangedFieldsViewModelDuration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["oldValue"] = o.OldValue
	toSerialize["newValue"] = o.NewValue
	return toSerialize, nil
}

type NullableWorkItemChangedFieldsViewModelDuration struct {
	value *WorkItemChangedFieldsViewModelDuration
	isSet bool
}

func (v NullableWorkItemChangedFieldsViewModelDuration) Get() *WorkItemChangedFieldsViewModelDuration {
	return v.value
}

func (v *NullableWorkItemChangedFieldsViewModelDuration) Set(val *WorkItemChangedFieldsViewModelDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemChangedFieldsViewModelDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemChangedFieldsViewModelDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemChangedFieldsViewModelDuration(val *WorkItemChangedFieldsViewModelDuration) *NullableWorkItemChangedFieldsViewModelDuration {
	return &NullableWorkItemChangedFieldsViewModelDuration{value: val, isSet: true}
}

func (v NullableWorkItemChangedFieldsViewModelDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemChangedFieldsViewModelDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

