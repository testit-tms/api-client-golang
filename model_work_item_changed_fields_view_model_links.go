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

// checks if the WorkItemChangedFieldsViewModelLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemChangedFieldsViewModelLinks{}

// WorkItemChangedFieldsViewModelLinks struct for WorkItemChangedFieldsViewModelLinks
type WorkItemChangedFieldsViewModelLinks struct {
	OldValue []WorkItemLinkChangeViewModel `json:"oldValue,omitempty"`
	NewValue []WorkItemLinkChangeViewModel `json:"newValue,omitempty"`
}

// NewWorkItemChangedFieldsViewModelLinks instantiates a new WorkItemChangedFieldsViewModelLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemChangedFieldsViewModelLinks() *WorkItemChangedFieldsViewModelLinks {
	this := WorkItemChangedFieldsViewModelLinks{}
	return &this
}

// NewWorkItemChangedFieldsViewModelLinksWithDefaults instantiates a new WorkItemChangedFieldsViewModelLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemChangedFieldsViewModelLinksWithDefaults() *WorkItemChangedFieldsViewModelLinks {
	this := WorkItemChangedFieldsViewModelLinks{}
	return &this
}

// GetOldValue returns the OldValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemChangedFieldsViewModelLinks) GetOldValue() []WorkItemLinkChangeViewModel {
	if o == nil {
		var ret []WorkItemLinkChangeViewModel
		return ret
	}
	return o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemChangedFieldsViewModelLinks) GetOldValueOk() ([]WorkItemLinkChangeViewModel, bool) {
	if o == nil || IsNil(o.OldValue) {
		return nil, false
	}
	return o.OldValue, true
}

// HasOldValue returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModelLinks) HasOldValue() bool {
	if o != nil && IsNil(o.OldValue) {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given []WorkItemLinkChangeViewModel and assigns it to the OldValue field.
func (o *WorkItemChangedFieldsViewModelLinks) SetOldValue(v []WorkItemLinkChangeViewModel) {
	o.OldValue = v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemChangedFieldsViewModelLinks) GetNewValue() []WorkItemLinkChangeViewModel {
	if o == nil {
		var ret []WorkItemLinkChangeViewModel
		return ret
	}
	return o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemChangedFieldsViewModelLinks) GetNewValueOk() ([]WorkItemLinkChangeViewModel, bool) {
	if o == nil || IsNil(o.NewValue) {
		return nil, false
	}
	return o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModelLinks) HasNewValue() bool {
	if o != nil && IsNil(o.NewValue) {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given []WorkItemLinkChangeViewModel and assigns it to the NewValue field.
func (o *WorkItemChangedFieldsViewModelLinks) SetNewValue(v []WorkItemLinkChangeViewModel) {
	o.NewValue = v
}

func (o WorkItemChangedFieldsViewModelLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemChangedFieldsViewModelLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OldValue != nil {
		toSerialize["oldValue"] = o.OldValue
	}
	if o.NewValue != nil {
		toSerialize["newValue"] = o.NewValue
	}
	return toSerialize, nil
}

type NullableWorkItemChangedFieldsViewModelLinks struct {
	value *WorkItemChangedFieldsViewModelLinks
	isSet bool
}

func (v NullableWorkItemChangedFieldsViewModelLinks) Get() *WorkItemChangedFieldsViewModelLinks {
	return v.value
}

func (v *NullableWorkItemChangedFieldsViewModelLinks) Set(val *WorkItemChangedFieldsViewModelLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemChangedFieldsViewModelLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemChangedFieldsViewModelLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemChangedFieldsViewModelLinks(val *WorkItemChangedFieldsViewModelLinks) *NullableWorkItemChangedFieldsViewModelLinks {
	return &NullableWorkItemChangedFieldsViewModelLinks{value: val, isSet: true}
}

func (v NullableWorkItemChangedFieldsViewModelLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemChangedFieldsViewModelLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


