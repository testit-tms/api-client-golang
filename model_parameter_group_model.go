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

// checks if the ParameterGroupModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParameterGroupModel{}

// ParameterGroupModel struct for ParameterGroupModel
type ParameterGroupModel struct {
	Name *string `json:"name,omitempty"`
	Values *map[string]string `json:"values,omitempty"`
	ParameterKeyId *string `json:"parameterKeyId,omitempty"`
}

// NewParameterGroupModel instantiates a new ParameterGroupModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterGroupModel() *ParameterGroupModel {
	this := ParameterGroupModel{}
	return &this
}

// NewParameterGroupModelWithDefaults instantiates a new ParameterGroupModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterGroupModelWithDefaults() *ParameterGroupModel {
	this := ParameterGroupModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ParameterGroupModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterGroupModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ParameterGroupModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ParameterGroupModel) SetName(v string) {
	o.Name = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ParameterGroupModel) GetValues() map[string]string {
	if o == nil || IsNil(o.Values) {
		var ret map[string]string
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterGroupModel) GetValuesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ParameterGroupModel) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given map[string]string and assigns it to the Values field.
func (o *ParameterGroupModel) SetValues(v map[string]string) {
	o.Values = &v
}

// GetParameterKeyId returns the ParameterKeyId field value if set, zero value otherwise.
func (o *ParameterGroupModel) GetParameterKeyId() string {
	if o == nil || IsNil(o.ParameterKeyId) {
		var ret string
		return ret
	}
	return *o.ParameterKeyId
}

// GetParameterKeyIdOk returns a tuple with the ParameterKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterGroupModel) GetParameterKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParameterKeyId) {
		return nil, false
	}
	return o.ParameterKeyId, true
}

// HasParameterKeyId returns a boolean if a field has been set.
func (o *ParameterGroupModel) HasParameterKeyId() bool {
	if o != nil && !IsNil(o.ParameterKeyId) {
		return true
	}

	return false
}

// SetParameterKeyId gets a reference to the given string and assigns it to the ParameterKeyId field.
func (o *ParameterGroupModel) SetParameterKeyId(v string) {
	o.ParameterKeyId = &v
}

func (o ParameterGroupModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParameterGroupModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	if !IsNil(o.ParameterKeyId) {
		toSerialize["parameterKeyId"] = o.ParameterKeyId
	}
	return toSerialize, nil
}

type NullableParameterGroupModel struct {
	value *ParameterGroupModel
	isSet bool
}

func (v NullableParameterGroupModel) Get() *ParameterGroupModel {
	return v.value
}

func (v *NullableParameterGroupModel) Set(val *ParameterGroupModel) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterGroupModel) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterGroupModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterGroupModel(val *ParameterGroupModel) *NullableParameterGroupModel {
	return &NullableParameterGroupModel{value: val, isSet: true}
}

func (v NullableParameterGroupModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterGroupModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


