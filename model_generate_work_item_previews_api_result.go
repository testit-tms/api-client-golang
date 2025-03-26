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

// checks if the GenerateWorkItemPreviewsApiResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateWorkItemPreviewsApiResult{}

// GenerateWorkItemPreviewsApiResult struct for GenerateWorkItemPreviewsApiResult
type GenerateWorkItemPreviewsApiResult struct {
	Previews []WorkItemPreviewApiModel `json:"previews"`
}

type _GenerateWorkItemPreviewsApiResult GenerateWorkItemPreviewsApiResult

// NewGenerateWorkItemPreviewsApiResult instantiates a new GenerateWorkItemPreviewsApiResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateWorkItemPreviewsApiResult(previews []WorkItemPreviewApiModel) *GenerateWorkItemPreviewsApiResult {
	this := GenerateWorkItemPreviewsApiResult{}
	this.Previews = previews
	return &this
}

// NewGenerateWorkItemPreviewsApiResultWithDefaults instantiates a new GenerateWorkItemPreviewsApiResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateWorkItemPreviewsApiResultWithDefaults() *GenerateWorkItemPreviewsApiResult {
	this := GenerateWorkItemPreviewsApiResult{}
	return &this
}

// GetPreviews returns the Previews field value
func (o *GenerateWorkItemPreviewsApiResult) GetPreviews() []WorkItemPreviewApiModel {
	if o == nil {
		var ret []WorkItemPreviewApiModel
		return ret
	}

	return o.Previews
}

// GetPreviewsOk returns a tuple with the Previews field value
// and a boolean to check if the value has been set.
func (o *GenerateWorkItemPreviewsApiResult) GetPreviewsOk() ([]WorkItemPreviewApiModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previews, true
}

// SetPreviews sets field value
func (o *GenerateWorkItemPreviewsApiResult) SetPreviews(v []WorkItemPreviewApiModel) {
	o.Previews = v
}

func (o GenerateWorkItemPreviewsApiResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateWorkItemPreviewsApiResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["previews"] = o.Previews
	return toSerialize, nil
}

func (o *GenerateWorkItemPreviewsApiResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"previews",
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

	varGenerateWorkItemPreviewsApiResult := _GenerateWorkItemPreviewsApiResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGenerateWorkItemPreviewsApiResult)

	if err != nil {
		return err
	}

	*o = GenerateWorkItemPreviewsApiResult(varGenerateWorkItemPreviewsApiResult)

	return err
}

type NullableGenerateWorkItemPreviewsApiResult struct {
	value *GenerateWorkItemPreviewsApiResult
	isSet bool
}

func (v NullableGenerateWorkItemPreviewsApiResult) Get() *GenerateWorkItemPreviewsApiResult {
	return v.value
}

func (v *NullableGenerateWorkItemPreviewsApiResult) Set(val *GenerateWorkItemPreviewsApiResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateWorkItemPreviewsApiResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateWorkItemPreviewsApiResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateWorkItemPreviewsApiResult(val *GenerateWorkItemPreviewsApiResult) *NullableGenerateWorkItemPreviewsApiResult {
	return &NullableGenerateWorkItemPreviewsApiResult{value: val, isSet: true}
}

func (v NullableGenerateWorkItemPreviewsApiResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateWorkItemPreviewsApiResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


