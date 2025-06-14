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

// checks if the TestPlanTestPointsTestSuiteSearchApiResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPlanTestPointsTestSuiteSearchApiResult{}

// TestPlanTestPointsTestSuiteSearchApiResult struct for TestPlanTestPointsTestSuiteSearchApiResult
type TestPlanTestPointsTestSuiteSearchApiResult struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type _TestPlanTestPointsTestSuiteSearchApiResult TestPlanTestPointsTestSuiteSearchApiResult

// NewTestPlanTestPointsTestSuiteSearchApiResult instantiates a new TestPlanTestPointsTestSuiteSearchApiResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPlanTestPointsTestSuiteSearchApiResult(id string, name string) *TestPlanTestPointsTestSuiteSearchApiResult {
	this := TestPlanTestPointsTestSuiteSearchApiResult{}
	this.Id = id
	this.Name = name
	return &this
}

// NewTestPlanTestPointsTestSuiteSearchApiResultWithDefaults instantiates a new TestPlanTestPointsTestSuiteSearchApiResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPlanTestPointsTestSuiteSearchApiResultWithDefaults() *TestPlanTestPointsTestSuiteSearchApiResult {
	this := TestPlanTestPointsTestSuiteSearchApiResult{}
	return &this
}

// GetId returns the Id field value
func (o *TestPlanTestPointsTestSuiteSearchApiResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestPlanTestPointsTestSuiteSearchApiResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestPlanTestPointsTestSuiteSearchApiResult) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TestPlanTestPointsTestSuiteSearchApiResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TestPlanTestPointsTestSuiteSearchApiResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TestPlanTestPointsTestSuiteSearchApiResult) SetName(v string) {
	o.Name = v
}

func (o TestPlanTestPointsTestSuiteSearchApiResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPlanTestPointsTestSuiteSearchApiResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *TestPlanTestPointsTestSuiteSearchApiResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varTestPlanTestPointsTestSuiteSearchApiResult := _TestPlanTestPointsTestSuiteSearchApiResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestPlanTestPointsTestSuiteSearchApiResult)

	if err != nil {
		return err
	}

	*o = TestPlanTestPointsTestSuiteSearchApiResult(varTestPlanTestPointsTestSuiteSearchApiResult)

	return err
}

type NullableTestPlanTestPointsTestSuiteSearchApiResult struct {
	value *TestPlanTestPointsTestSuiteSearchApiResult
	isSet bool
}

func (v NullableTestPlanTestPointsTestSuiteSearchApiResult) Get() *TestPlanTestPointsTestSuiteSearchApiResult {
	return v.value
}

func (v *NullableTestPlanTestPointsTestSuiteSearchApiResult) Set(val *TestPlanTestPointsTestSuiteSearchApiResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPlanTestPointsTestSuiteSearchApiResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPlanTestPointsTestSuiteSearchApiResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPlanTestPointsTestSuiteSearchApiResult(val *TestPlanTestPointsTestSuiteSearchApiResult) *NullableTestPlanTestPointsTestSuiteSearchApiResult {
	return &NullableTestPlanTestPointsTestSuiteSearchApiResult{value: val, isSet: true}
}

func (v NullableTestPlanTestPointsTestSuiteSearchApiResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPlanTestPointsTestSuiteSearchApiResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


