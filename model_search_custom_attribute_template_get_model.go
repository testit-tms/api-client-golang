/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SearchCustomAttributeTemplateGetModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCustomAttributeTemplateGetModel{}

// SearchCustomAttributeTemplateGetModel struct for SearchCustomAttributeTemplateGetModel
type SearchCustomAttributeTemplateGetModel struct {
	Id string `json:"id"`
	IsDeleted bool `json:"isDeleted"`
	Name string `json:"name"`
	ProjectShortestModels []ProjectShortestModel `json:"projectShortestModels"`
	CustomAttributeModels []CustomAttributeModel `json:"customAttributeModels"`
}

type _SearchCustomAttributeTemplateGetModel SearchCustomAttributeTemplateGetModel

// NewSearchCustomAttributeTemplateGetModel instantiates a new SearchCustomAttributeTemplateGetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCustomAttributeTemplateGetModel(id string, isDeleted bool, name string, projectShortestModels []ProjectShortestModel, customAttributeModels []CustomAttributeModel) *SearchCustomAttributeTemplateGetModel {
	this := SearchCustomAttributeTemplateGetModel{}
	this.Id = id
	this.IsDeleted = isDeleted
	this.Name = name
	this.ProjectShortestModels = projectShortestModels
	this.CustomAttributeModels = customAttributeModels
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

// GetName returns the Name field value
func (o *SearchCustomAttributeTemplateGetModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SearchCustomAttributeTemplateGetModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SearchCustomAttributeTemplateGetModel) SetName(v string) {
	o.Name = v
}

// GetProjectShortestModels returns the ProjectShortestModels field value
func (o *SearchCustomAttributeTemplateGetModel) GetProjectShortestModels() []ProjectShortestModel {
	if o == nil {
		var ret []ProjectShortestModel
		return ret
	}

	return o.ProjectShortestModels
}

// GetProjectShortestModelsOk returns a tuple with the ProjectShortestModels field value
// and a boolean to check if the value has been set.
func (o *SearchCustomAttributeTemplateGetModel) GetProjectShortestModelsOk() ([]ProjectShortestModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectShortestModels, true
}

// SetProjectShortestModels sets field value
func (o *SearchCustomAttributeTemplateGetModel) SetProjectShortestModels(v []ProjectShortestModel) {
	o.ProjectShortestModels = v
}

// GetCustomAttributeModels returns the CustomAttributeModels field value
func (o *SearchCustomAttributeTemplateGetModel) GetCustomAttributeModels() []CustomAttributeModel {
	if o == nil {
		var ret []CustomAttributeModel
		return ret
	}

	return o.CustomAttributeModels
}

// GetCustomAttributeModelsOk returns a tuple with the CustomAttributeModels field value
// and a boolean to check if the value has been set.
func (o *SearchCustomAttributeTemplateGetModel) GetCustomAttributeModelsOk() ([]CustomAttributeModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomAttributeModels, true
}

// SetCustomAttributeModels sets field value
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
	toSerialize["name"] = o.Name
	toSerialize["projectShortestModels"] = o.ProjectShortestModels
	toSerialize["customAttributeModels"] = o.CustomAttributeModels
	return toSerialize, nil
}

func (o *SearchCustomAttributeTemplateGetModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"isDeleted",
		"name",
		"projectShortestModels",
		"customAttributeModels",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSearchCustomAttributeTemplateGetModel := _SearchCustomAttributeTemplateGetModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchCustomAttributeTemplateGetModel)

	if err != nil {
		return err
	}

	*o = SearchCustomAttributeTemplateGetModel(varSearchCustomAttributeTemplateGetModel)

	return err
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


