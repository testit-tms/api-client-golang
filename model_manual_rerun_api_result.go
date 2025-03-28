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

// checks if the ManualRerunApiResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManualRerunApiResult{}

// ManualRerunApiResult struct for ManualRerunApiResult
type ManualRerunApiResult struct {
	TestResultsCount int32 `json:"testResultsCount"`
}

type _ManualRerunApiResult ManualRerunApiResult

// NewManualRerunApiResult instantiates a new ManualRerunApiResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualRerunApiResult(testResultsCount int32) *ManualRerunApiResult {
	this := ManualRerunApiResult{}
	this.TestResultsCount = testResultsCount
	return &this
}

// NewManualRerunApiResultWithDefaults instantiates a new ManualRerunApiResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualRerunApiResultWithDefaults() *ManualRerunApiResult {
	this := ManualRerunApiResult{}
	return &this
}

// GetTestResultsCount returns the TestResultsCount field value
func (o *ManualRerunApiResult) GetTestResultsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TestResultsCount
}

// GetTestResultsCountOk returns a tuple with the TestResultsCount field value
// and a boolean to check if the value has been set.
func (o *ManualRerunApiResult) GetTestResultsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestResultsCount, true
}

// SetTestResultsCount sets field value
func (o *ManualRerunApiResult) SetTestResultsCount(v int32) {
	o.TestResultsCount = v
}

func (o ManualRerunApiResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManualRerunApiResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["testResultsCount"] = o.TestResultsCount
	return toSerialize, nil
}

func (o *ManualRerunApiResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"testResultsCount",
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

	varManualRerunApiResult := _ManualRerunApiResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManualRerunApiResult)

	if err != nil {
		return err
	}

	*o = ManualRerunApiResult(varManualRerunApiResult)

	return err
}

type NullableManualRerunApiResult struct {
	value *ManualRerunApiResult
	isSet bool
}

func (v NullableManualRerunApiResult) Get() *ManualRerunApiResult {
	return v.value
}

func (v *NullableManualRerunApiResult) Set(val *ManualRerunApiResult) {
	v.value = val
	v.isSet = true
}

func (v NullableManualRerunApiResult) IsSet() bool {
	return v.isSet
}

func (v *NullableManualRerunApiResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualRerunApiResult(val *ManualRerunApiResult) *NullableManualRerunApiResult {
	return &NullableManualRerunApiResult{value: val, isSet: true}
}

func (v NullableManualRerunApiResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualRerunApiResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


