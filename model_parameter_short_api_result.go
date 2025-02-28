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

// checks if the ParameterShortApiResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParameterShortApiResult{}

// ParameterShortApiResult struct for ParameterShortApiResult
type ParameterShortApiResult struct {
	Id string `json:"id"`
	ParameterKeyId string `json:"parameterKeyId"`
	// Value of the parameter
	Value string `json:"value"`
	// Key of the parameter
	Name string `json:"name"`
}

type _ParameterShortApiResult ParameterShortApiResult

// NewParameterShortApiResult instantiates a new ParameterShortApiResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterShortApiResult(id string, parameterKeyId string, value string, name string) *ParameterShortApiResult {
	this := ParameterShortApiResult{}
	this.Id = id
	this.ParameterKeyId = parameterKeyId
	this.Value = value
	this.Name = name
	return &this
}

// NewParameterShortApiResultWithDefaults instantiates a new ParameterShortApiResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterShortApiResultWithDefaults() *ParameterShortApiResult {
	this := ParameterShortApiResult{}
	return &this
}

// GetId returns the Id field value
func (o *ParameterShortApiResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ParameterShortApiResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ParameterShortApiResult) SetId(v string) {
	o.Id = v
}

// GetParameterKeyId returns the ParameterKeyId field value
func (o *ParameterShortApiResult) GetParameterKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParameterKeyId
}

// GetParameterKeyIdOk returns a tuple with the ParameterKeyId field value
// and a boolean to check if the value has been set.
func (o *ParameterShortApiResult) GetParameterKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParameterKeyId, true
}

// SetParameterKeyId sets field value
func (o *ParameterShortApiResult) SetParameterKeyId(v string) {
	o.ParameterKeyId = v
}

// GetValue returns the Value field value
func (o *ParameterShortApiResult) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ParameterShortApiResult) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ParameterShortApiResult) SetValue(v string) {
	o.Value = v
}

// GetName returns the Name field value
func (o *ParameterShortApiResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParameterShortApiResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ParameterShortApiResult) SetName(v string) {
	o.Name = v
}

func (o ParameterShortApiResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParameterShortApiResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["parameterKeyId"] = o.ParameterKeyId
	toSerialize["value"] = o.Value
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *ParameterShortApiResult) UnmarshalJSON(data []byte) (err error) {
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

	varParameterShortApiResult := _ParameterShortApiResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varParameterShortApiResult)

	if err != nil {
		return err
	}

	*o = ParameterShortApiResult(varParameterShortApiResult)

	return err
}

type NullableParameterShortApiResult struct {
	value *ParameterShortApiResult
	isSet bool
}

func (v NullableParameterShortApiResult) Get() *ParameterShortApiResult {
	return v.value
}

func (v *NullableParameterShortApiResult) Set(val *ParameterShortApiResult) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterShortApiResult) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterShortApiResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterShortApiResult(val *ParameterShortApiResult) *NullableParameterShortApiResult {
	return &NullableParameterShortApiResult{value: val, isSet: true}
}

func (v NullableParameterShortApiResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterShortApiResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


