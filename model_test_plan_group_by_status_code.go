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

// checks if the TestPlanGroupByStatusCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPlanGroupByStatusCode{}

// TestPlanGroupByStatusCode struct for TestPlanGroupByStatusCode
type TestPlanGroupByStatusCode struct {
	StatusCode string `json:"statusCode"`
	Value int64 `json:"value"`
}

type _TestPlanGroupByStatusCode TestPlanGroupByStatusCode

// NewTestPlanGroupByStatusCode instantiates a new TestPlanGroupByStatusCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPlanGroupByStatusCode(statusCode string, value int64) *TestPlanGroupByStatusCode {
	this := TestPlanGroupByStatusCode{}
	this.StatusCode = statusCode
	this.Value = value
	return &this
}

// NewTestPlanGroupByStatusCodeWithDefaults instantiates a new TestPlanGroupByStatusCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPlanGroupByStatusCodeWithDefaults() *TestPlanGroupByStatusCode {
	this := TestPlanGroupByStatusCode{}
	return &this
}

// GetStatusCode returns the StatusCode field value
func (o *TestPlanGroupByStatusCode) GetStatusCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *TestPlanGroupByStatusCode) GetStatusCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *TestPlanGroupByStatusCode) SetStatusCode(v string) {
	o.StatusCode = v
}

// GetValue returns the Value field value
func (o *TestPlanGroupByStatusCode) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TestPlanGroupByStatusCode) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TestPlanGroupByStatusCode) SetValue(v int64) {
	o.Value = v
}

func (o TestPlanGroupByStatusCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPlanGroupByStatusCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["statusCode"] = o.StatusCode
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *TestPlanGroupByStatusCode) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"statusCode",
		"value",
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

	varTestPlanGroupByStatusCode := _TestPlanGroupByStatusCode{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestPlanGroupByStatusCode)

	if err != nil {
		return err
	}

	*o = TestPlanGroupByStatusCode(varTestPlanGroupByStatusCode)

	return err
}

type NullableTestPlanGroupByStatusCode struct {
	value *TestPlanGroupByStatusCode
	isSet bool
}

func (v NullableTestPlanGroupByStatusCode) Get() *TestPlanGroupByStatusCode {
	return v.value
}

func (v *NullableTestPlanGroupByStatusCode) Set(val *TestPlanGroupByStatusCode) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPlanGroupByStatusCode) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPlanGroupByStatusCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPlanGroupByStatusCode(val *TestPlanGroupByStatusCode) *NullableTestPlanGroupByStatusCode {
	return &NullableTestPlanGroupByStatusCode{value: val, isSet: true}
}

func (v NullableTestPlanGroupByStatusCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPlanGroupByStatusCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


