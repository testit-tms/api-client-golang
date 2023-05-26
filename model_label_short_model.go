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

// checks if the LabelShortModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelShortModel{}

// LabelShortModel struct for LabelShortModel
type LabelShortModel struct {
	GlobalId *int64 `json:"globalId,omitempty"`
	// Label name.
	Name string `json:"name"`
}

// NewLabelShortModel instantiates a new LabelShortModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelShortModel(name string) *LabelShortModel {
	this := LabelShortModel{}
	this.Name = name
	return &this
}

// NewLabelShortModelWithDefaults instantiates a new LabelShortModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelShortModelWithDefaults() *LabelShortModel {
	this := LabelShortModel{}
	return &this
}

// GetGlobalId returns the GlobalId field value if set, zero value otherwise.
func (o *LabelShortModel) GetGlobalId() int64 {
	if o == nil || IsNil(o.GlobalId) {
		var ret int64
		return ret
	}
	return *o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelShortModel) GetGlobalIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GlobalId) {
		return nil, false
	}
	return o.GlobalId, true
}

// HasGlobalId returns a boolean if a field has been set.
func (o *LabelShortModel) HasGlobalId() bool {
	if o != nil && !IsNil(o.GlobalId) {
		return true
	}

	return false
}

// SetGlobalId gets a reference to the given int64 and assigns it to the GlobalId field.
func (o *LabelShortModel) SetGlobalId(v int64) {
	o.GlobalId = &v
}

// GetName returns the Name field value
func (o *LabelShortModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LabelShortModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LabelShortModel) SetName(v string) {
	o.Name = v
}

func (o LabelShortModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelShortModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GlobalId) {
		toSerialize["globalId"] = o.GlobalId
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableLabelShortModel struct {
	value *LabelShortModel
	isSet bool
}

func (v NullableLabelShortModel) Get() *LabelShortModel {
	return v.value
}

func (v *NullableLabelShortModel) Set(val *LabelShortModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelShortModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelShortModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelShortModel(val *LabelShortModel) *NullableLabelShortModel {
	return &NullableLabelShortModel{value: val, isSet: true}
}

func (v NullableLabelShortModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelShortModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

