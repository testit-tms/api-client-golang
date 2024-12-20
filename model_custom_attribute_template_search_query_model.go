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

// checks if the CustomAttributeTemplateSearchQueryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomAttributeTemplateSearchQueryModel{}

// CustomAttributeTemplateSearchQueryModel struct for CustomAttributeTemplateSearchQueryModel
type CustomAttributeTemplateSearchQueryModel struct {
	Name NullableString `json:"name,omitempty"`
	ProjectIds []string `json:"projectIds,omitempty"`
	CustomAttributeTypes []CustomAttributeTypesEnum `json:"customAttributeTypes,omitempty"`
	IsDeleted NullableBool `json:"isDeleted,omitempty"`
}

// NewCustomAttributeTemplateSearchQueryModel instantiates a new CustomAttributeTemplateSearchQueryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomAttributeTemplateSearchQueryModel() *CustomAttributeTemplateSearchQueryModel {
	this := CustomAttributeTemplateSearchQueryModel{}
	return &this
}

// NewCustomAttributeTemplateSearchQueryModelWithDefaults instantiates a new CustomAttributeTemplateSearchQueryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomAttributeTemplateSearchQueryModelWithDefaults() *CustomAttributeTemplateSearchQueryModel {
	this := CustomAttributeTemplateSearchQueryModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomAttributeTemplateSearchQueryModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomAttributeTemplateSearchQueryModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CustomAttributeTemplateSearchQueryModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CustomAttributeTemplateSearchQueryModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CustomAttributeTemplateSearchQueryModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CustomAttributeTemplateSearchQueryModel) UnsetName() {
	o.Name.Unset()
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomAttributeTemplateSearchQueryModel) GetProjectIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomAttributeTemplateSearchQueryModel) GetProjectIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectIds) {
		return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *CustomAttributeTemplateSearchQueryModel) HasProjectIds() bool {
	if o != nil && !IsNil(o.ProjectIds) {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []string and assigns it to the ProjectIds field.
func (o *CustomAttributeTemplateSearchQueryModel) SetProjectIds(v []string) {
	o.ProjectIds = v
}

// GetCustomAttributeTypes returns the CustomAttributeTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomAttributeTemplateSearchQueryModel) GetCustomAttributeTypes() []CustomAttributeTypesEnum {
	if o == nil {
		var ret []CustomAttributeTypesEnum
		return ret
	}
	return o.CustomAttributeTypes
}

// GetCustomAttributeTypesOk returns a tuple with the CustomAttributeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomAttributeTemplateSearchQueryModel) GetCustomAttributeTypesOk() ([]CustomAttributeTypesEnum, bool) {
	if o == nil || IsNil(o.CustomAttributeTypes) {
		return nil, false
	}
	return o.CustomAttributeTypes, true
}

// HasCustomAttributeTypes returns a boolean if a field has been set.
func (o *CustomAttributeTemplateSearchQueryModel) HasCustomAttributeTypes() bool {
	if o != nil && !IsNil(o.CustomAttributeTypes) {
		return true
	}

	return false
}

// SetCustomAttributeTypes gets a reference to the given []CustomAttributeTypesEnum and assigns it to the CustomAttributeTypes field.
func (o *CustomAttributeTemplateSearchQueryModel) SetCustomAttributeTypes(v []CustomAttributeTypesEnum) {
	o.CustomAttributeTypes = v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomAttributeTemplateSearchQueryModel) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDeleted.Get()
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomAttributeTemplateSearchQueryModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDeleted.Get(), o.IsDeleted.IsSet()
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *CustomAttributeTemplateSearchQueryModel) HasIsDeleted() bool {
	if o != nil && o.IsDeleted.IsSet() {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given NullableBool and assigns it to the IsDeleted field.
func (o *CustomAttributeTemplateSearchQueryModel) SetIsDeleted(v bool) {
	o.IsDeleted.Set(&v)
}
// SetIsDeletedNil sets the value for IsDeleted to be an explicit nil
func (o *CustomAttributeTemplateSearchQueryModel) SetIsDeletedNil() {
	o.IsDeleted.Set(nil)
}

// UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
func (o *CustomAttributeTemplateSearchQueryModel) UnsetIsDeleted() {
	o.IsDeleted.Unset()
}

func (o CustomAttributeTemplateSearchQueryModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomAttributeTemplateSearchQueryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ProjectIds != nil {
		toSerialize["projectIds"] = o.ProjectIds
	}
	if o.CustomAttributeTypes != nil {
		toSerialize["customAttributeTypes"] = o.CustomAttributeTypes
	}
	if o.IsDeleted.IsSet() {
		toSerialize["isDeleted"] = o.IsDeleted.Get()
	}
	return toSerialize, nil
}

type NullableCustomAttributeTemplateSearchQueryModel struct {
	value *CustomAttributeTemplateSearchQueryModel
	isSet bool
}

func (v NullableCustomAttributeTemplateSearchQueryModel) Get() *CustomAttributeTemplateSearchQueryModel {
	return v.value
}

func (v *NullableCustomAttributeTemplateSearchQueryModel) Set(val *CustomAttributeTemplateSearchQueryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomAttributeTemplateSearchQueryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomAttributeTemplateSearchQueryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomAttributeTemplateSearchQueryModel(val *CustomAttributeTemplateSearchQueryModel) *NullableCustomAttributeTemplateSearchQueryModel {
	return &NullableCustomAttributeTemplateSearchQueryModel{value: val, isSet: true}
}

func (v NullableCustomAttributeTemplateSearchQueryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomAttributeTemplateSearchQueryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


