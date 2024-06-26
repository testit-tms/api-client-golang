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

// checks if the TestResultChangeViewModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestResultChangeViewModel{}

// TestResultChangeViewModel struct for TestResultChangeViewModel
type TestResultChangeViewModel struct {
	TestPointCount int64 `json:"testPointCount"`
}

type _TestResultChangeViewModel TestResultChangeViewModel

// NewTestResultChangeViewModel instantiates a new TestResultChangeViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestResultChangeViewModel(testPointCount int64) *TestResultChangeViewModel {
	this := TestResultChangeViewModel{}
	this.TestPointCount = testPointCount
	return &this
}

// NewTestResultChangeViewModelWithDefaults instantiates a new TestResultChangeViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestResultChangeViewModelWithDefaults() *TestResultChangeViewModel {
	this := TestResultChangeViewModel{}
	return &this
}

// GetTestPointCount returns the TestPointCount field value
func (o *TestResultChangeViewModel) GetTestPointCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TestPointCount
}

// GetTestPointCountOk returns a tuple with the TestPointCount field value
// and a boolean to check if the value has been set.
func (o *TestResultChangeViewModel) GetTestPointCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestPointCount, true
}

// SetTestPointCount sets field value
func (o *TestResultChangeViewModel) SetTestPointCount(v int64) {
	o.TestPointCount = v
}

func (o TestResultChangeViewModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestResultChangeViewModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["testPointCount"] = o.TestPointCount
	return toSerialize, nil
}

func (o *TestResultChangeViewModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"testPointCount",
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

	varTestResultChangeViewModel := _TestResultChangeViewModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestResultChangeViewModel)

	if err != nil {
		return err
	}

	*o = TestResultChangeViewModel(varTestResultChangeViewModel)

	return err
}

type NullableTestResultChangeViewModel struct {
	value *TestResultChangeViewModel
	isSet bool
}

func (v NullableTestResultChangeViewModel) Get() *TestResultChangeViewModel {
	return v.value
}

func (v *NullableTestResultChangeViewModel) Set(val *TestResultChangeViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestResultChangeViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestResultChangeViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestResultChangeViewModel(val *TestResultChangeViewModel) *NullableTestResultChangeViewModel {
	return &NullableTestResultChangeViewModel{value: val, isSet: true}
}

func (v NullableTestResultChangeViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestResultChangeViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


