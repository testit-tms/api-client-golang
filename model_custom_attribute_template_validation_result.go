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

// checks if the CustomAttributeTemplateValidationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomAttributeTemplateValidationResult{}

// CustomAttributeTemplateValidationResult struct for CustomAttributeTemplateValidationResult
type CustomAttributeTemplateValidationResult struct {
	Exists bool `json:"exists"`
}

type _CustomAttributeTemplateValidationResult CustomAttributeTemplateValidationResult

// NewCustomAttributeTemplateValidationResult instantiates a new CustomAttributeTemplateValidationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomAttributeTemplateValidationResult(exists bool) *CustomAttributeTemplateValidationResult {
	this := CustomAttributeTemplateValidationResult{}
	this.Exists = exists
	return &this
}

// NewCustomAttributeTemplateValidationResultWithDefaults instantiates a new CustomAttributeTemplateValidationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomAttributeTemplateValidationResultWithDefaults() *CustomAttributeTemplateValidationResult {
	this := CustomAttributeTemplateValidationResult{}
	return &this
}

// GetExists returns the Exists field value
func (o *CustomAttributeTemplateValidationResult) GetExists() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Exists
}

// GetExistsOk returns a tuple with the Exists field value
// and a boolean to check if the value has been set.
func (o *CustomAttributeTemplateValidationResult) GetExistsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exists, true
}

// SetExists sets field value
func (o *CustomAttributeTemplateValidationResult) SetExists(v bool) {
	o.Exists = v
}

func (o CustomAttributeTemplateValidationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomAttributeTemplateValidationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["exists"] = o.Exists
	return toSerialize, nil
}

func (o *CustomAttributeTemplateValidationResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"exists",
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

	varCustomAttributeTemplateValidationResult := _CustomAttributeTemplateValidationResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomAttributeTemplateValidationResult)

	if err != nil {
		return err
	}

	*o = CustomAttributeTemplateValidationResult(varCustomAttributeTemplateValidationResult)

	return err
}

type NullableCustomAttributeTemplateValidationResult struct {
	value *CustomAttributeTemplateValidationResult
	isSet bool
}

func (v NullableCustomAttributeTemplateValidationResult) Get() *CustomAttributeTemplateValidationResult {
	return v.value
}

func (v *NullableCustomAttributeTemplateValidationResult) Set(val *CustomAttributeTemplateValidationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomAttributeTemplateValidationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomAttributeTemplateValidationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomAttributeTemplateValidationResult(val *CustomAttributeTemplateValidationResult) *NullableCustomAttributeTemplateValidationResult {
	return &NullableCustomAttributeTemplateValidationResult{value: val, isSet: true}
}

func (v NullableCustomAttributeTemplateValidationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomAttributeTemplateValidationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


