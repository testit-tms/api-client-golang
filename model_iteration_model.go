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

// checks if the IterationModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IterationModel{}

// IterationModel struct for IterationModel
type IterationModel struct {
	Id *string `json:"id,omitempty"`
	Parameters []ParameterShortModel `json:"parameters,omitempty"`
}

// NewIterationModel instantiates a new IterationModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIterationModel() *IterationModel {
	this := IterationModel{}
	return &this
}

// NewIterationModelWithDefaults instantiates a new IterationModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIterationModelWithDefaults() *IterationModel {
	this := IterationModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IterationModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IterationModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IterationModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IterationModel) SetId(v string) {
	o.Id = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *IterationModel) GetParameters() []ParameterShortModel {
	if o == nil || IsNil(o.Parameters) {
		var ret []ParameterShortModel
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IterationModel) GetParametersOk() ([]ParameterShortModel, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *IterationModel) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []ParameterShortModel and assigns it to the Parameters field.
func (o *IterationModel) SetParameters(v []ParameterShortModel) {
	o.Parameters = v
}

func (o IterationModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IterationModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	return toSerialize, nil
}

type NullableIterationModel struct {
	value *IterationModel
	isSet bool
}

func (v NullableIterationModel) Get() *IterationModel {
	return v.value
}

func (v *NullableIterationModel) Set(val *IterationModel) {
	v.value = val
	v.isSet = true
}

func (v NullableIterationModel) IsSet() bool {
	return v.isSet
}

func (v *NullableIterationModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIterationModel(val *IterationModel) *NullableIterationModel {
	return &NullableIterationModel{value: val, isSet: true}
}

func (v NullableIterationModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIterationModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


