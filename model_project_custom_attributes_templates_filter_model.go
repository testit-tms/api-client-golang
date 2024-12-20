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

// checks if the ProjectCustomAttributesTemplatesFilterModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectCustomAttributesTemplatesFilterModel{}

// ProjectCustomAttributesTemplatesFilterModel Collection of filters to apply to search
type ProjectCustomAttributesTemplatesFilterModel struct {
	// Name of custom attribute template
	Name NullableString `json:"name,omitempty"`
	// Collection of custom attributes types
	CustomAttributeTypes []CustomAttributeTypesEnum `json:"customAttributeTypes,omitempty"`
}

// NewProjectCustomAttributesTemplatesFilterModel instantiates a new ProjectCustomAttributesTemplatesFilterModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectCustomAttributesTemplatesFilterModel() *ProjectCustomAttributesTemplatesFilterModel {
	this := ProjectCustomAttributesTemplatesFilterModel{}
	return &this
}

// NewProjectCustomAttributesTemplatesFilterModelWithDefaults instantiates a new ProjectCustomAttributesTemplatesFilterModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectCustomAttributesTemplatesFilterModelWithDefaults() *ProjectCustomAttributesTemplatesFilterModel {
	this := ProjectCustomAttributesTemplatesFilterModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectCustomAttributesTemplatesFilterModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectCustomAttributesTemplatesFilterModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ProjectCustomAttributesTemplatesFilterModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ProjectCustomAttributesTemplatesFilterModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ProjectCustomAttributesTemplatesFilterModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ProjectCustomAttributesTemplatesFilterModel) UnsetName() {
	o.Name.Unset()
}

// GetCustomAttributeTypes returns the CustomAttributeTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectCustomAttributesTemplatesFilterModel) GetCustomAttributeTypes() []CustomAttributeTypesEnum {
	if o == nil {
		var ret []CustomAttributeTypesEnum
		return ret
	}
	return o.CustomAttributeTypes
}

// GetCustomAttributeTypesOk returns a tuple with the CustomAttributeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectCustomAttributesTemplatesFilterModel) GetCustomAttributeTypesOk() ([]CustomAttributeTypesEnum, bool) {
	if o == nil || IsNil(o.CustomAttributeTypes) {
		return nil, false
	}
	return o.CustomAttributeTypes, true
}

// HasCustomAttributeTypes returns a boolean if a field has been set.
func (o *ProjectCustomAttributesTemplatesFilterModel) HasCustomAttributeTypes() bool {
	if o != nil && !IsNil(o.CustomAttributeTypes) {
		return true
	}

	return false
}

// SetCustomAttributeTypes gets a reference to the given []CustomAttributeTypesEnum and assigns it to the CustomAttributeTypes field.
func (o *ProjectCustomAttributesTemplatesFilterModel) SetCustomAttributeTypes(v []CustomAttributeTypesEnum) {
	o.CustomAttributeTypes = v
}

func (o ProjectCustomAttributesTemplatesFilterModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectCustomAttributesTemplatesFilterModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.CustomAttributeTypes != nil {
		toSerialize["customAttributeTypes"] = o.CustomAttributeTypes
	}
	return toSerialize, nil
}

type NullableProjectCustomAttributesTemplatesFilterModel struct {
	value *ProjectCustomAttributesTemplatesFilterModel
	isSet bool
}

func (v NullableProjectCustomAttributesTemplatesFilterModel) Get() *ProjectCustomAttributesTemplatesFilterModel {
	return v.value
}

func (v *NullableProjectCustomAttributesTemplatesFilterModel) Set(val *ProjectCustomAttributesTemplatesFilterModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCustomAttributesTemplatesFilterModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCustomAttributesTemplatesFilterModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCustomAttributesTemplatesFilterModel(val *ProjectCustomAttributesTemplatesFilterModel) *NullableProjectCustomAttributesTemplatesFilterModel {
	return &NullableProjectCustomAttributesTemplatesFilterModel{value: val, isSet: true}
}

func (v NullableProjectCustomAttributesTemplatesFilterModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCustomAttributesTemplatesFilterModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


