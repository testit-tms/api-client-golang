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

// checks if the ProjectCustomAttributeTemplateGetModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectCustomAttributeTemplateGetModel{}

// ProjectCustomAttributeTemplateGetModel struct for ProjectCustomAttributeTemplateGetModel
type ProjectCustomAttributeTemplateGetModel struct {
	// Unique ID of the custom attributes template
	Id *string `json:"id,omitempty"`
	// Indicates if the custom attribute template is deleted
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// Name of the custom attribute template
	Name *string `json:"name,omitempty"`
	// Attributes of the template
	CustomAttributeModels []CustomAttributeModel `json:"customAttributeModels,omitempty"`
}

// NewProjectCustomAttributeTemplateGetModel instantiates a new ProjectCustomAttributeTemplateGetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectCustomAttributeTemplateGetModel() *ProjectCustomAttributeTemplateGetModel {
	this := ProjectCustomAttributeTemplateGetModel{}
	return &this
}

// NewProjectCustomAttributeTemplateGetModelWithDefaults instantiates a new ProjectCustomAttributeTemplateGetModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectCustomAttributeTemplateGetModelWithDefaults() *ProjectCustomAttributeTemplateGetModel {
	this := ProjectCustomAttributeTemplateGetModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectCustomAttributeTemplateGetModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCustomAttributeTemplateGetModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectCustomAttributeTemplateGetModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProjectCustomAttributeTemplateGetModel) SetId(v string) {
	o.Id = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *ProjectCustomAttributeTemplateGetModel) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCustomAttributeTemplateGetModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *ProjectCustomAttributeTemplateGetModel) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *ProjectCustomAttributeTemplateGetModel) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectCustomAttributeTemplateGetModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCustomAttributeTemplateGetModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectCustomAttributeTemplateGetModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectCustomAttributeTemplateGetModel) SetName(v string) {
	o.Name = &v
}

// GetCustomAttributeModels returns the CustomAttributeModels field value if set, zero value otherwise.
func (o *ProjectCustomAttributeTemplateGetModel) GetCustomAttributeModels() []CustomAttributeModel {
	if o == nil || IsNil(o.CustomAttributeModels) {
		var ret []CustomAttributeModel
		return ret
	}
	return o.CustomAttributeModels
}

// GetCustomAttributeModelsOk returns a tuple with the CustomAttributeModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCustomAttributeTemplateGetModel) GetCustomAttributeModelsOk() ([]CustomAttributeModel, bool) {
	if o == nil || IsNil(o.CustomAttributeModels) {
		return nil, false
	}
	return o.CustomAttributeModels, true
}

// HasCustomAttributeModels returns a boolean if a field has been set.
func (o *ProjectCustomAttributeTemplateGetModel) HasCustomAttributeModels() bool {
	if o != nil && !IsNil(o.CustomAttributeModels) {
		return true
	}

	return false
}

// SetCustomAttributeModels gets a reference to the given []CustomAttributeModel and assigns it to the CustomAttributeModels field.
func (o *ProjectCustomAttributeTemplateGetModel) SetCustomAttributeModels(v []CustomAttributeModel) {
	o.CustomAttributeModels = v
}

func (o ProjectCustomAttributeTemplateGetModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectCustomAttributeTemplateGetModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CustomAttributeModels) {
		toSerialize["customAttributeModels"] = o.CustomAttributeModels
	}
	return toSerialize, nil
}

type NullableProjectCustomAttributeTemplateGetModel struct {
	value *ProjectCustomAttributeTemplateGetModel
	isSet bool
}

func (v NullableProjectCustomAttributeTemplateGetModel) Get() *ProjectCustomAttributeTemplateGetModel {
	return v.value
}

func (v *NullableProjectCustomAttributeTemplateGetModel) Set(val *ProjectCustomAttributeTemplateGetModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCustomAttributeTemplateGetModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCustomAttributeTemplateGetModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCustomAttributeTemplateGetModel(val *ProjectCustomAttributeTemplateGetModel) *NullableProjectCustomAttributeTemplateGetModel {
	return &NullableProjectCustomAttributeTemplateGetModel{value: val, isSet: true}
}

func (v NullableProjectCustomAttributeTemplateGetModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCustomAttributeTemplateGetModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


