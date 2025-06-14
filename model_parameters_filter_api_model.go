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

// checks if the ParametersFilterApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParametersFilterApiModel{}

// ParametersFilterApiModel struct for ParametersFilterApiModel
type ParametersFilterApiModel struct {
	Name NullableString `json:"name,omitempty"`
	IsDeleted NullableBool `json:"isDeleted,omitempty"`
}

// NewParametersFilterApiModel instantiates a new ParametersFilterApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParametersFilterApiModel() *ParametersFilterApiModel {
	this := ParametersFilterApiModel{}
	return &this
}

// NewParametersFilterApiModelWithDefaults instantiates a new ParametersFilterApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParametersFilterApiModelWithDefaults() *ParametersFilterApiModel {
	this := ParametersFilterApiModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParametersFilterApiModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParametersFilterApiModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ParametersFilterApiModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ParametersFilterApiModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ParametersFilterApiModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ParametersFilterApiModel) UnsetName() {
	o.Name.Unset()
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParametersFilterApiModel) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDeleted.Get()
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParametersFilterApiModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDeleted.Get(), o.IsDeleted.IsSet()
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *ParametersFilterApiModel) HasIsDeleted() bool {
	if o != nil && o.IsDeleted.IsSet() {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given NullableBool and assigns it to the IsDeleted field.
func (o *ParametersFilterApiModel) SetIsDeleted(v bool) {
	o.IsDeleted.Set(&v)
}
// SetIsDeletedNil sets the value for IsDeleted to be an explicit nil
func (o *ParametersFilterApiModel) SetIsDeletedNil() {
	o.IsDeleted.Set(nil)
}

// UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
func (o *ParametersFilterApiModel) UnsetIsDeleted() {
	o.IsDeleted.Unset()
}

func (o ParametersFilterApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParametersFilterApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.IsDeleted.IsSet() {
		toSerialize["isDeleted"] = o.IsDeleted.Get()
	}
	return toSerialize, nil
}

type NullableParametersFilterApiModel struct {
	value *ParametersFilterApiModel
	isSet bool
}

func (v NullableParametersFilterApiModel) Get() *ParametersFilterApiModel {
	return v.value
}

func (v *NullableParametersFilterApiModel) Set(val *ParametersFilterApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableParametersFilterApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableParametersFilterApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParametersFilterApiModel(val *ParametersFilterApiModel) *NullableParametersFilterApiModel {
	return &NullableParametersFilterApiModel{value: val, isSet: true}
}

func (v NullableParametersFilterApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParametersFilterApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


