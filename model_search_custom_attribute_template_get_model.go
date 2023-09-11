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
	Id string `json:"id"`
	IsDeleted bool `json:"isDeleted"`
	Name NullableString `json:"name,omitempty"`
	ProjectShortestModels []ProjectShortestModel `json:"projectShortestModels,omitempty"`
	CustomAttributeModels []CustomAttributeModel `json:"customAttributeModels,omitempty"`
}

// NewSearchCustomAttributeTemplateGetModel instantiates a new SearchCustomAttributeTemplateGetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCustomAttributeTemplateGetModel(id string, isDeleted bool) *SearchCustomAttributeTemplateGetModel {
	this := SearchCustomAttributeTemplateGetModel{}
	this.Id = id
	this.IsDeleted = isDeleted
	return &this
}

// NewSearchCustomAttributeTemplateGetModelWithDefaults instantiates a new SearchCustomAttributeTemplateGetModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCustomAttributeTemplateGetModelWithDefaults() *SearchCustomAttributeTemplateGetModel {
	this := SearchCustomAttributeTemplateGetModel{}
	return &this
}

// GetId returns the Id field value
func (o *SearchCustomAttributeTemplateGetModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SearchCustomAttributeTemplateGetModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SearchCustomAttributeTemplateGetModel) SetId(v string) {
	o.Id = v
}

// GetIsDeleted returns the IsDeleted field value
func (o *SearchCustomAttributeTemplateGetModel) GetIsDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value
// and a boolean to check if the value has been set.
func (o *SearchCustomAttributeTemplateGetModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeleted, true
}

// SetIsDeleted sets field value
func (o *SearchCustomAttributeTemplateGetModel) SetIsDeleted(v bool) {
	o.IsDeleted = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchCustomAttributeTemplateGetModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchCustomAttributeTemplateGetModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SearchCustomAttributeTemplateGetModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SearchCustomAttributeTemplateGetModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SearchCustomAttributeTemplateGetModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SearchCustomAttributeTemplateGetModel) UnsetName() {
	o.Name.Unset()
}

// GetProjectShortestModels returns the ProjectShortestModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchCustomAttributeTemplateGetModel) GetProjectShortestModels() []ProjectShortestModel {
	if o == nil {
		var ret []ProjectShortestModel
		return ret
	}
	return o.ProjectShortestModels
}

// GetProjectShortestModelsOk returns a tuple with the ProjectShortestModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchCustomAttributeTemplateGetModel) GetProjectShortestModelsOk() ([]ProjectShortestModel, bool) {
	if o == nil || IsNil(o.ProjectShortestModels) {
		return nil, false
	}
	return o.ProjectShortestModels, true
}

// HasProjectShortestModels returns a boolean if a field has been set.
func (o *SearchCustomAttributeTemplateGetModel) HasProjectShortestModels() bool {
	if o != nil && IsNil(o.ProjectShortestModels) {
		return true
	}

	return false
}

// SetProjectShortestModels gets a reference to the given []ProjectShortestModel and assigns it to the ProjectShortestModels field.
func (o *SearchCustomAttributeTemplateGetModel) SetProjectShortestModels(v []ProjectShortestModel) {
	o.ProjectShortestModels = v
}

// GetCustomAttributeModels returns the CustomAttributeModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchCustomAttributeTemplateGetModel) GetCustomAttributeModels() []CustomAttributeModel {
	if o == nil {
		var ret []CustomAttributeModel
		return ret
	}
	return o.CustomAttributeModels
}

// GetCustomAttributeModelsOk returns a tuple with the CustomAttributeModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchCustomAttributeTemplateGetModel) GetCustomAttributeModelsOk() ([]CustomAttributeModel, bool) {
	if o == nil || IsNil(o.CustomAttributeModels) {
		return nil, false
	}
	return o.CustomAttributeModels, true
}

// HasCustomAttributeModels returns a boolean if a field has been set.
func (o *SearchCustomAttributeTemplateGetModel) HasCustomAttributeModels() bool {
	if o != nil && IsNil(o.CustomAttributeModels) {
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
	toSerialize["id"] = o.Id
	toSerialize["isDeleted"] = o.IsDeleted
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ProjectShortestModels != nil {
		toSerialize["projectShortestModels"] = o.ProjectShortestModels
	}
	if o.CustomAttributeModels != nil {
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


