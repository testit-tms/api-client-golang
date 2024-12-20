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

// checks if the AutoTestNamespaceModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoTestNamespaceModel{}

// AutoTestNamespaceModel struct for AutoTestNamespaceModel
type AutoTestNamespaceModel struct {
	Name NullableString `json:"name,omitempty"`
	Classes []string `json:"classes,omitempty"`
}

// NewAutoTestNamespaceModel instantiates a new AutoTestNamespaceModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoTestNamespaceModel() *AutoTestNamespaceModel {
	this := AutoTestNamespaceModel{}
	return &this
}

// NewAutoTestNamespaceModelWithDefaults instantiates a new AutoTestNamespaceModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoTestNamespaceModelWithDefaults() *AutoTestNamespaceModel {
	this := AutoTestNamespaceModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestNamespaceModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestNamespaceModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AutoTestNamespaceModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AutoTestNamespaceModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AutoTestNamespaceModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AutoTestNamespaceModel) UnsetName() {
	o.Name.Unset()
}

// GetClasses returns the Classes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestNamespaceModel) GetClasses() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Classes
}

// GetClassesOk returns a tuple with the Classes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestNamespaceModel) GetClassesOk() ([]string, bool) {
	if o == nil || IsNil(o.Classes) {
		return nil, false
	}
	return o.Classes, true
}

// HasClasses returns a boolean if a field has been set.
func (o *AutoTestNamespaceModel) HasClasses() bool {
	if o != nil && !IsNil(o.Classes) {
		return true
	}

	return false
}

// SetClasses gets a reference to the given []string and assigns it to the Classes field.
func (o *AutoTestNamespaceModel) SetClasses(v []string) {
	o.Classes = v
}

func (o AutoTestNamespaceModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoTestNamespaceModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Classes != nil {
		toSerialize["classes"] = o.Classes
	}
	return toSerialize, nil
}

type NullableAutoTestNamespaceModel struct {
	value *AutoTestNamespaceModel
	isSet bool
}

func (v NullableAutoTestNamespaceModel) Get() *AutoTestNamespaceModel {
	return v.value
}

func (v *NullableAutoTestNamespaceModel) Set(val *AutoTestNamespaceModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoTestNamespaceModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoTestNamespaceModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoTestNamespaceModel(val *AutoTestNamespaceModel) *NullableAutoTestNamespaceModel {
	return &NullableAutoTestNamespaceModel{value: val, isSet: true}
}

func (v NullableAutoTestNamespaceModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoTestNamespaceModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


