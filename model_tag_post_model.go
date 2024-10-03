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

// checks if the TagPostModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagPostModel{}

// TagPostModel struct for TagPostModel
type TagPostModel struct {
	Name string `json:"name"`
}

// NewTagPostModel instantiates a new TagPostModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagPostModel(name string) *TagPostModel {
	this := TagPostModel{}
	this.Name = name
	return &this
}

// NewTagPostModelWithDefaults instantiates a new TagPostModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagPostModelWithDefaults() *TagPostModel {
	this := TagPostModel{}
	return &this
}

// GetName returns the Name field value
func (o *TagPostModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagPostModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagPostModel) SetName(v string) {
	o.Name = v
}

func (o TagPostModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagPostModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableTagPostModel struct {
	value *TagPostModel
	isSet bool
}

func (v NullableTagPostModel) Get() *TagPostModel {
	return v.value
}

func (v *NullableTagPostModel) Set(val *TagPostModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTagPostModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTagPostModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagPostModel(val *TagPostModel) *NullableTagPostModel {
	return &NullableTagPostModel{value: val, isSet: true}
}

func (v NullableTagPostModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagPostModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


