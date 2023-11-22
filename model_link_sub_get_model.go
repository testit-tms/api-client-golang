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

// checks if the LinkSubGetModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkSubGetModel{}

// LinkSubGetModel struct for LinkSubGetModel
type LinkSubGetModel struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

// NewLinkSubGetModel instantiates a new LinkSubGetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkSubGetModel(name string, url string) *LinkSubGetModel {
	this := LinkSubGetModel{}
	this.Name = name
	this.Url = url
	return &this
}

// NewLinkSubGetModelWithDefaults instantiates a new LinkSubGetModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkSubGetModelWithDefaults() *LinkSubGetModel {
	this := LinkSubGetModel{}
	return &this
}

// GetName returns the Name field value
func (o *LinkSubGetModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LinkSubGetModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LinkSubGetModel) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *LinkSubGetModel) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *LinkSubGetModel) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *LinkSubGetModel) SetUrl(v string) {
	o.Url = v
}

func (o LinkSubGetModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkSubGetModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableLinkSubGetModel struct {
	value *LinkSubGetModel
	isSet bool
}

func (v NullableLinkSubGetModel) Get() *LinkSubGetModel {
	return v.value
}

func (v *NullableLinkSubGetModel) Set(val *LinkSubGetModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkSubGetModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkSubGetModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkSubGetModel(val *LinkSubGetModel) *NullableLinkSubGetModel {
	return &NullableLinkSubGetModel{value: val, isSet: true}
}

func (v NullableLinkSubGetModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkSubGetModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


