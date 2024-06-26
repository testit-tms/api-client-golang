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

// checks if the ParameterShortModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParameterShortModel{}

// ParameterShortModel struct for ParameterShortModel
type ParameterShortModel struct {
	Id string `json:"id"`
	ParameterKeyId string `json:"parameterKeyId"`
	// Value of the parameter
	Value string `json:"value"`
	// Key of the parameter
	Name string `json:"name"`
}

type _ParameterShortModel ParameterShortModel

// NewParameterShortModel instantiates a new ParameterShortModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterShortModel(id string, parameterKeyId string, value string, name string) *ParameterShortModel {
	this := ParameterShortModel{}
	this.Id = id
	this.ParameterKeyId = parameterKeyId
	this.Value = value
	this.Name = name
	return &this
}

// NewParameterShortModelWithDefaults instantiates a new ParameterShortModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterShortModelWithDefaults() *ParameterShortModel {
	this := ParameterShortModel{}
	return &this
}

// GetId returns the Id field value
func (o *ParameterShortModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ParameterShortModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ParameterShortModel) SetId(v string) {
	o.Id = v
}

// GetParameterKeyId returns the ParameterKeyId field value
func (o *ParameterShortModel) GetParameterKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParameterKeyId
}

// GetParameterKeyIdOk returns a tuple with the ParameterKeyId field value
// and a boolean to check if the value has been set.
func (o *ParameterShortModel) GetParameterKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParameterKeyId, true
}

// SetParameterKeyId sets field value
func (o *ParameterShortModel) SetParameterKeyId(v string) {
	o.ParameterKeyId = v
}

// GetValue returns the Value field value
func (o *ParameterShortModel) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ParameterShortModel) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ParameterShortModel) SetValue(v string) {
	o.Value = v
}

// GetName returns the Name field value
func (o *ParameterShortModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParameterShortModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ParameterShortModel) SetName(v string) {
	o.Name = v
}

func (o ParameterShortModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParameterShortModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["parameterKeyId"] = o.ParameterKeyId
	toSerialize["value"] = o.Value
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *ParameterShortModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"parameterKeyId",
		"value",
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

	varParameterShortModel := _ParameterShortModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varParameterShortModel)

	if err != nil {
		return err
	}

	*o = ParameterShortModel(varParameterShortModel)

	return err
}

type NullableParameterShortModel struct {
	value *ParameterShortModel
	isSet bool
}

func (v NullableParameterShortModel) Get() *ParameterShortModel {
	return v.value
}

func (v *NullableParameterShortModel) Set(val *ParameterShortModel) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterShortModel) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterShortModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterShortModel(val *ParameterShortModel) *NullableParameterShortModel {
	return &NullableParameterShortModel{value: val, isSet: true}
}

func (v NullableParameterShortModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterShortModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


