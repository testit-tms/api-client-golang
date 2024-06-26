/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the WorkItemChangedAttributeViewModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemChangedAttributeViewModel{}

// WorkItemChangedAttributeViewModel struct for WorkItemChangedAttributeViewModel
type WorkItemChangedAttributeViewModel struct {
	Type string `json:"type"`
	OldAttributeName string `json:"oldAttributeName"`
	NewAttributeName string `json:"newAttributeName"`
	OldValue interface{} `json:"oldValue"`
	NewValue interface{} `json:"newValue"`
}

type _WorkItemChangedAttributeViewModel WorkItemChangedAttributeViewModel

// NewWorkItemChangedAttributeViewModel instantiates a new WorkItemChangedAttributeViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemChangedAttributeViewModel(type_ string, oldAttributeName string, newAttributeName string, oldValue interface{}, newValue interface{}) *WorkItemChangedAttributeViewModel {
	this := WorkItemChangedAttributeViewModel{}
	this.Type = type_
	this.OldAttributeName = oldAttributeName
	this.NewAttributeName = newAttributeName
	this.OldValue = oldValue
	this.NewValue = newValue
	return &this
}

// NewWorkItemChangedAttributeViewModelWithDefaults instantiates a new WorkItemChangedAttributeViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemChangedAttributeViewModelWithDefaults() *WorkItemChangedAttributeViewModel {
	this := WorkItemChangedAttributeViewModel{}
	return &this
}

// GetType returns the Type field value
func (o *WorkItemChangedAttributeViewModel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WorkItemChangedAttributeViewModel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WorkItemChangedAttributeViewModel) SetType(v string) {
	o.Type = v
}

// GetOldAttributeName returns the OldAttributeName field value
func (o *WorkItemChangedAttributeViewModel) GetOldAttributeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OldAttributeName
}

// GetOldAttributeNameOk returns a tuple with the OldAttributeName field value
// and a boolean to check if the value has been set.
func (o *WorkItemChangedAttributeViewModel) GetOldAttributeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OldAttributeName, true
}

// SetOldAttributeName sets field value
func (o *WorkItemChangedAttributeViewModel) SetOldAttributeName(v string) {
	o.OldAttributeName = v
}

// GetNewAttributeName returns the NewAttributeName field value
func (o *WorkItemChangedAttributeViewModel) GetNewAttributeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewAttributeName
}

// GetNewAttributeNameOk returns a tuple with the NewAttributeName field value
// and a boolean to check if the value has been set.
func (o *WorkItemChangedAttributeViewModel) GetNewAttributeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewAttributeName, true
}

// SetNewAttributeName sets field value
func (o *WorkItemChangedAttributeViewModel) SetNewAttributeName(v string) {
	o.NewAttributeName = v
}

// GetOldValue returns the OldValue field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *WorkItemChangedAttributeViewModel) GetOldValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemChangedAttributeViewModel) GetOldValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.OldValue) {
		return nil, false
	}
	return &o.OldValue, true
}

// SetOldValue sets field value
func (o *WorkItemChangedAttributeViewModel) SetOldValue(v interface{}) {
	o.OldValue = v
}

// GetNewValue returns the NewValue field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *WorkItemChangedAttributeViewModel) GetNewValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemChangedAttributeViewModel) GetNewValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.NewValue) {
		return nil, false
	}
	return &o.NewValue, true
}

// SetNewValue sets field value
func (o *WorkItemChangedAttributeViewModel) SetNewValue(v interface{}) {
	o.NewValue = v
}

func (o WorkItemChangedAttributeViewModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemChangedAttributeViewModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["oldAttributeName"] = o.OldAttributeName
	toSerialize["newAttributeName"] = o.NewAttributeName
	if o.OldValue != nil {
		toSerialize["oldValue"] = o.OldValue
	}
	if o.NewValue != nil {
		toSerialize["newValue"] = o.NewValue
	}
	return toSerialize, nil
}

func (o *WorkItemChangedAttributeViewModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"oldAttributeName",
		"newAttributeName",
		"oldValue",
		"newValue",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWorkItemChangedAttributeViewModel := _WorkItemChangedAttributeViewModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkItemChangedAttributeViewModel)

	if err != nil {
		return err
	}

	*o = WorkItemChangedAttributeViewModel(varWorkItemChangedAttributeViewModel)

	return err
}

type NullableWorkItemChangedAttributeViewModel struct {
	value *WorkItemChangedAttributeViewModel
	isSet bool
}

func (v NullableWorkItemChangedAttributeViewModel) Get() *WorkItemChangedAttributeViewModel {
	return v.value
}

func (v *NullableWorkItemChangedAttributeViewModel) Set(val *WorkItemChangedAttributeViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemChangedAttributeViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemChangedAttributeViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemChangedAttributeViewModel(val *WorkItemChangedAttributeViewModel) *NullableWorkItemChangedAttributeViewModel {
	return &NullableWorkItemChangedAttributeViewModel{value: val, isSet: true}
}

func (v NullableWorkItemChangedAttributeViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemChangedAttributeViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


