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

// checks if the CreateTagApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTagApiModel{}

// CreateTagApiModel struct for CreateTagApiModel
type CreateTagApiModel struct {
	// Name of the tag
	Name string `json:"name"`
}

type _CreateTagApiModel CreateTagApiModel

// NewCreateTagApiModel instantiates a new CreateTagApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTagApiModel(name string) *CreateTagApiModel {
	this := CreateTagApiModel{}
	this.Name = name
	return &this
}

// NewCreateTagApiModelWithDefaults instantiates a new CreateTagApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTagApiModelWithDefaults() *CreateTagApiModel {
	this := CreateTagApiModel{}
	return &this
}

// GetName returns the Name field value
func (o *CreateTagApiModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateTagApiModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateTagApiModel) SetName(v string) {
	o.Name = v
}

func (o CreateTagApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTagApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *CreateTagApiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varCreateTagApiModel := _CreateTagApiModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTagApiModel)

	if err != nil {
		return err
	}

	*o = CreateTagApiModel(varCreateTagApiModel)

	return err
}

type NullableCreateTagApiModel struct {
	value *CreateTagApiModel
	isSet bool
}

func (v NullableCreateTagApiModel) Get() *CreateTagApiModel {
	return v.value
}

func (v *NullableCreateTagApiModel) Set(val *CreateTagApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTagApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTagApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTagApiModel(val *CreateTagApiModel) *NullableCreateTagApiModel {
	return &NullableCreateTagApiModel{value: val, isSet: true}
}

func (v NullableCreateTagApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTagApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


