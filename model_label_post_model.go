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

// checks if the LabelPostModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelPostModel{}

// LabelPostModel struct for LabelPostModel
type LabelPostModel struct {
	// Name of the label
	Name string `json:"name"`
}

type _LabelPostModel LabelPostModel

// NewLabelPostModel instantiates a new LabelPostModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelPostModel(name string) *LabelPostModel {
	this := LabelPostModel{}
	this.Name = name
	return &this
}

// NewLabelPostModelWithDefaults instantiates a new LabelPostModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelPostModelWithDefaults() *LabelPostModel {
	this := LabelPostModel{}
	return &this
}

// GetName returns the Name field value
func (o *LabelPostModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LabelPostModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LabelPostModel) SetName(v string) {
	o.Name = v
}

func (o LabelPostModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelPostModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *LabelPostModel) UnmarshalJSON(data []byte) (err error) {
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

	varLabelPostModel := _LabelPostModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLabelPostModel)

	if err != nil {
		return err
	}

	*o = LabelPostModel(varLabelPostModel)

	return err
}

type NullableLabelPostModel struct {
	value *LabelPostModel
	isSet bool
}

func (v NullableLabelPostModel) Get() *LabelPostModel {
	return v.value
}

func (v *NullableLabelPostModel) Set(val *LabelPostModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelPostModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelPostModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelPostModel(val *LabelPostModel) *NullableLabelPostModel {
	return &NullableLabelPostModel{value: val, isSet: true}
}

func (v NullableLabelPostModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelPostModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


