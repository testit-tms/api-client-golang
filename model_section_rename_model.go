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

// checks if the SectionRenameModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SectionRenameModel{}

// SectionRenameModel struct for SectionRenameModel
type SectionRenameModel struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type _SectionRenameModel SectionRenameModel

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

func (o *SectionRenameModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varSectionRenameModel := _SectionRenameModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSectionRenameModel)

	if err != nil {
		return err
	}

	*o = SectionRenameModel(varSectionRenameModel)

	return err
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


