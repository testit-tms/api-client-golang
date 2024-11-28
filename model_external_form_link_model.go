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

// checks if the ExternalFormLinkModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalFormLinkModel{}

// ExternalFormLinkModel struct for ExternalFormLinkModel
type ExternalFormLinkModel struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type _ExternalFormLinkModel ExternalFormLinkModel

// NewExternalFormLinkModel instantiates a new ExternalFormLinkModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalFormLinkModel(name string, url string) *ExternalFormLinkModel {
	this := ExternalFormLinkModel{}
	this.Name = name
	this.Url = url
	return &this
}

// NewExternalFormLinkModelWithDefaults instantiates a new ExternalFormLinkModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalFormLinkModelWithDefaults() *ExternalFormLinkModel {
	this := ExternalFormLinkModel{}
	return &this
}

// GetName returns the Name field value
func (o *ExternalFormLinkModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExternalFormLinkModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExternalFormLinkModel) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *ExternalFormLinkModel) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ExternalFormLinkModel) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ExternalFormLinkModel) SetUrl(v string) {
	o.Url = v
}

func (o ExternalFormLinkModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalFormLinkModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *ExternalFormLinkModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"url",
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

	varExternalFormLinkModel := _ExternalFormLinkModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalFormLinkModel)

	if err != nil {
		return err
	}

	*o = ExternalFormLinkModel(varExternalFormLinkModel)

	return err
}

type NullableExternalFormLinkModel struct {
	value *ExternalFormLinkModel
	isSet bool
}

func (v NullableExternalFormLinkModel) Get() *ExternalFormLinkModel {
	return v.value
}

func (v *NullableExternalFormLinkModel) Set(val *ExternalFormLinkModel) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalFormLinkModel) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalFormLinkModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalFormLinkModel(val *ExternalFormLinkModel) *NullableExternalFormLinkModel {
	return &NullableExternalFormLinkModel{value: val, isSet: true}
}

func (v NullableExternalFormLinkModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalFormLinkModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


