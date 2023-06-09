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

// checks if the SearchCustomAttributeTemplateGetModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCustomAttributeTemplateGetModel{}

// SearchCustomAttributeTemplateGetModel struct for SearchCustomAttributeTemplateGetModel
type SearchCustomAttributeTemplateGetModel struct {
	Id *string `json:"id,omitempty"`
	IsDeleted *bool `json:"isDeleted,omitempty"`
	Name *string `json:"name,omitempty"`
	ProjectShortestModels []ProjectShortestModel `json:"projectShortestModels,omitempty"`
	CustomAttributeModels []CustomAttributeModel `json:"customAttributeModels,omitempty"`
}

// NewSearchCustomAttributeTemplateGetModel instantiates a new SearchCustomAttributeTemplateGetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCustomAttributeTemplateGetModel() *SearchCustomAttributeTemplateGetModel {
	this := SearchCustomAttributeTemplateGetModel{}
	return &this
}

// NewSearchCustomAttributeTemplateGetModelWithDefaults instantiates a new SearchCustomAttributeTemplateGetModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCustomAttributeTemplateGetModelWithDefaults() *SearchCustomAttributeTemplateGetModel {
	this := SearchCustomAttributeTemplateGetModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SearchCustomAttributeTemplateGetModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCustomAttributeTemplateGetModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SearchCustomAttributeTemplateGetModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SearchCustomAttributeTemplateGetModel) SetId(v string) {
	o.Id = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *SearchCustomAttributeTemplateGetModel) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCustomAttributeTemplateGetModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *SearchCustomAttributeTemplateGetModel) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *SearchCustomAttributeTemplateGetModel) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SearchCustomAttributeTemplateGetModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCustomAttributeTemplateGetModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SearchCustomAttributeTemplateGetModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SearchCustomAttributeTemplateGetModel) SetName(v string) {
	o.Name = &v
}

// GetProjectShortestModels returns the ProjectShortestModels field value if set, zero value otherwise.
func (o *SearchCustomAttributeTemplateGetModel) GetProjectShortestModels() []ProjectShortestModel {
	if o == nil || IsNil(o.ProjectShortestModels) {
		var ret []ProjectShortestModel
		return ret
	}
	return o.ProjectShortestModels
}

// GetProjectShortestModelsOk returns a tuple with the ProjectShortestModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCustomAttributeTemplateGetModel) GetProjectShortestModelsOk() ([]ProjectShortestModel, bool) {
	if o == nil || IsNil(o.ProjectShortestModels) {
		return nil, false
	}
	return o.ProjectShortestModels, true
}

// HasProjectShortestModels returns a boolean if a field has been set.
func (o *SearchCustomAttributeTemplateGetModel) HasProjectShortestModels() bool {
	if o != nil && !IsNil(o.ProjectShortestModels) {
		return true
	}

	return false
}

// SetProjectShortestModels gets a reference to the given []ProjectShortestModel and assigns it to the ProjectShortestModels field.
func (o *SearchCustomAttributeTemplateGetModel) SetProjectShortestModels(v []ProjectShortestModel) {
	o.ProjectShortestModels = v
}

// GetCustomAttributeModels returns the CustomAttributeModels field value if set, zero value otherwise.
func (o *SearchCustomAttributeTemplateGetModel) GetCustomAttributeModels() []CustomAttributeModel {
	if o == nil || IsNil(o.CustomAttributeModels) {
		var ret []CustomAttributeModel
		return ret
	}
	return o.CustomAttributeModels
}

// GetCustomAttributeModelsOk returns a tuple with the CustomAttributeModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCustomAttributeTemplateGetModel) GetCustomAttributeModelsOk() ([]CustomAttributeModel, bool) {
	if o == nil || IsNil(o.CustomAttributeModels) {
		return nil, false
	}
	return o.CustomAttributeModels, true
}

// HasCustomAttributeModels returns a boolean if a field has been set.
func (o *SearchCustomAttributeTemplateGetModel) HasCustomAttributeModels() bool {
	if o != nil && !IsNil(o.CustomAttributeModels) {
		return true
	}

	return false
}

// SetCustomAttributeModels gets a reference to the given []CustomAttributeModel and assigns it to the CustomAttributeModels field.
func (o *SearchCustomAttributeTemplateGetModel) SetCustomAttributeModels(v []CustomAttributeModel) {
	o.CustomAttributeModels = v
}

func (o SearchCustomAttributeTemplateGetModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCustomAttributeTemplateGetModel) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ProjectShortestModels) {
		toSerialize["projectShortestModels"] = o.ProjectShortestModels
	}
	if !IsNil(o.CustomAttributeModels) {
		toSerialize["customAttributeModels"] = o.CustomAttributeModels
	}
	return toSerialize, nil
}

type NullableSearchCustomAttributeTemplateGetModel struct {
	value *SearchCustomAttributeTemplateGetModel
	isSet bool
}

func (v NullableSearchCustomAttributeTemplateGetModel) Get() *SearchCustomAttributeTemplateGetModel {
	return v.value
}

func (v *NullableSearchCustomAttributeTemplateGetModel) Set(val *SearchCustomAttributeTemplateGetModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCustomAttributeTemplateGetModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCustomAttributeTemplateGetModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCustomAttributeTemplateGetModel(val *SearchCustomAttributeTemplateGetModel) *NullableSearchCustomAttributeTemplateGetModel {
	return &NullableSearchCustomAttributeTemplateGetModel{value: val, isSet: true}
}

func (v NullableSearchCustomAttributeTemplateGetModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCustomAttributeTemplateGetModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


