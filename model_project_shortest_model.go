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

// checks if the ProjectShortestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectShortestModel{}

// ProjectShortestModel struct for ProjectShortestModel
type ProjectShortestModel struct {
	// Unique ID of project
	Id string `json:"id"`
	// Indicates whether the project is deleted
	IsDeleted bool `json:"isDeleted"`
	// Global ID of project
	GlobalId int64 `json:"globalId"`
	// Name of project
	Name string `json:"name"`
}

type _ProjectShortestModel ProjectShortestModel

// NewProjectShortestModel instantiates a new ProjectShortestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectShortestModel(id string, isDeleted bool, globalId int64, name string) *ProjectShortestModel {
	this := ProjectShortestModel{}
	this.Id = id
	this.IsDeleted = isDeleted
	this.GlobalId = globalId
	this.Name = name
	return &this
}

// NewProjectShortestModelWithDefaults instantiates a new ProjectShortestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectShortestModelWithDefaults() *ProjectShortestModel {
	this := ProjectShortestModel{}
	return &this
}

// GetId returns the Id field value
func (o *ProjectShortestModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectShortestModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectShortestModel) SetId(v string) {
	o.Id = v
}

// GetIsDeleted returns the IsDeleted field value
func (o *ProjectShortestModel) GetIsDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value
// and a boolean to check if the value has been set.
func (o *ProjectShortestModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeleted, true
}

// SetIsDeleted sets field value
func (o *ProjectShortestModel) SetIsDeleted(v bool) {
	o.IsDeleted = v
}

// GetGlobalId returns the GlobalId field value
func (o *ProjectShortestModel) GetGlobalId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value
// and a boolean to check if the value has been set.
func (o *ProjectShortestModel) GetGlobalIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalId, true
}

// SetGlobalId sets field value
func (o *ProjectShortestModel) SetGlobalId(v int64) {
	o.GlobalId = v
}

// GetName returns the Name field value
func (o *ProjectShortestModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectShortestModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectShortestModel) SetName(v string) {
	o.Name = v
}

func (o ProjectShortestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectShortestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["isDeleted"] = o.IsDeleted
	toSerialize["globalId"] = o.GlobalId
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *ProjectShortestModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"isDeleted",
		"globalId",
		"name",
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

	varProjectShortestModel := _ProjectShortestModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectShortestModel)

	if err != nil {
		return err
	}

	*o = ProjectShortestModel(varProjectShortestModel)

	return err
}

type NullableProjectShortestModel struct {
	value *ProjectShortestModel
	isSet bool
}

func (v NullableProjectShortestModel) Get() *ProjectShortestModel {
	return v.value
}

func (v *NullableProjectShortestModel) Set(val *ProjectShortestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectShortestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectShortestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectShortestModel(val *ProjectShortestModel) *NullableProjectShortestModel {
	return &NullableProjectShortestModel{value: val, isSet: true}
}

func (v NullableProjectShortestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectShortestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


