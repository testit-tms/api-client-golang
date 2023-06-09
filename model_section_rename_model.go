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

// checks if the SectionRenameModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SectionRenameModel{}

// SectionRenameModel struct for SectionRenameModel
type SectionRenameModel struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

// NewSectionRenameModel instantiates a new SectionRenameModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSectionRenameModel(id string, name string) *SectionRenameModel {
	this := SectionRenameModel{}
	this.Id = id
	this.Name = name
	return &this
}

// NewSectionRenameModelWithDefaults instantiates a new SectionRenameModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSectionRenameModelWithDefaults() *SectionRenameModel {
	this := SectionRenameModel{}
	return &this
}

// GetId returns the Id field value
func (o *SectionRenameModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SectionRenameModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SectionRenameModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SectionRenameModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SectionRenameModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SectionRenameModel) SetName(v string) {
	o.Name = v
}

func (o SectionRenameModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SectionRenameModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableSectionRenameModel struct {
	value *SectionRenameModel
	isSet bool
}

func (v NullableSectionRenameModel) Get() *SectionRenameModel {
	return v.value
}

func (v *NullableSectionRenameModel) Set(val *SectionRenameModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSectionRenameModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSectionRenameModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSectionRenameModel(val *SectionRenameModel) *NullableSectionRenameModel {
	return &NullableSectionRenameModel{value: val, isSet: true}
}

func (v NullableSectionRenameModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSectionRenameModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


