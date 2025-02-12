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

// checks if the TestResultFailureClassApiResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestResultFailureClassApiResult{}

// TestResultFailureClassApiResult struct for TestResultFailureClassApiResult
type TestResultFailureClassApiResult struct {
	FailureCategory FailureCategory `json:"failureCategory"`
}

type _TestResultFailureClassApiResult TestResultFailureClassApiResult

// NewTestResultFailureClassApiResult instantiates a new TestResultFailureClassApiResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestResultFailureClassApiResult(failureCategory FailureCategory) *TestResultFailureClassApiResult {
	this := TestResultFailureClassApiResult{}
	this.FailureCategory = failureCategory
	return &this
}

// NewTestResultFailureClassApiResultWithDefaults instantiates a new TestResultFailureClassApiResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestResultFailureClassApiResultWithDefaults() *TestResultFailureClassApiResult {
	this := TestResultFailureClassApiResult{}
	return &this
}

// GetFailureCategory returns the FailureCategory field value
func (o *TestResultFailureClassApiResult) GetFailureCategory() FailureCategory {
	if o == nil {
		var ret FailureCategory
		return ret
	}

	return o.FailureCategory
}

// GetFailureCategoryOk returns a tuple with the FailureCategory field value
// and a boolean to check if the value has been set.
func (o *TestResultFailureClassApiResult) GetFailureCategoryOk() (*FailureCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailureCategory, true
}

// SetFailureCategory sets field value
func (o *TestResultFailureClassApiResult) SetFailureCategory(v FailureCategory) {
	o.FailureCategory = v
}

func (o TestResultFailureClassApiResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestResultFailureClassApiResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["failureCategory"] = o.FailureCategory
	return toSerialize, nil
}

func (o *TestResultFailureClassApiResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"failureCategory",
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

	varTestResultFailureClassApiResult := _TestResultFailureClassApiResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestResultFailureClassApiResult)

	if err != nil {
		return err
	}

	*o = TestResultFailureClassApiResult(varTestResultFailureClassApiResult)

	return err
}

type NullableTestResultFailureClassApiResult struct {
	value *TestResultFailureClassApiResult
	isSet bool
}

func (v NullableTestResultFailureClassApiResult) Get() *TestResultFailureClassApiResult {
	return v.value
}

func (v *NullableTestResultFailureClassApiResult) Set(val *TestResultFailureClassApiResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTestResultFailureClassApiResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTestResultFailureClassApiResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestResultFailureClassApiResult(val *TestResultFailureClassApiResult) *NullableTestResultFailureClassApiResult {
	return &NullableTestResultFailureClassApiResult{value: val, isSet: true}
}

func (v NullableTestResultFailureClassApiResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestResultFailureClassApiResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


