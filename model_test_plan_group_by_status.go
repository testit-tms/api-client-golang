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

// checks if the TestPlanGroupByStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPlanGroupByStatus{}

// TestPlanGroupByStatus struct for TestPlanGroupByStatus
type TestPlanGroupByStatus struct {
	Status string `json:"status"`
	Value int64 `json:"value"`
}

type _TestPlanGroupByStatus TestPlanGroupByStatus

// NewTestPlanGroupByStatus instantiates a new TestPlanGroupByStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPlanGroupByStatus(status string, value int64) *TestPlanGroupByStatus {
	this := TestPlanGroupByStatus{}
	this.Status = status
	this.Value = value
	return &this
}

// NewTestPlanGroupByStatusWithDefaults instantiates a new TestPlanGroupByStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPlanGroupByStatusWithDefaults() *TestPlanGroupByStatus {
	this := TestPlanGroupByStatus{}
	return &this
}

// GetStatus returns the Status field value
func (o *TestPlanGroupByStatus) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TestPlanGroupByStatus) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TestPlanGroupByStatus) SetStatus(v string) {
	o.Status = v
}

// GetValue returns the Value field value
func (o *TestPlanGroupByStatus) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TestPlanGroupByStatus) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TestPlanGroupByStatus) SetValue(v int64) {
	o.Value = v
}

func (o TestPlanGroupByStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPlanGroupByStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *TestPlanGroupByStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varTestPlanGroupByStatus := _TestPlanGroupByStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestPlanGroupByStatus)

	if err != nil {
		return err
	}

	*o = TestPlanGroupByStatus(varTestPlanGroupByStatus)

	return err
}

type NullableTestPlanGroupByStatus struct {
	value *TestPlanGroupByStatus
	isSet bool
}

func (v NullableTestPlanGroupByStatus) Get() *TestPlanGroupByStatus {
	return v.value
}

func (v *NullableTestPlanGroupByStatus) Set(val *TestPlanGroupByStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPlanGroupByStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPlanGroupByStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPlanGroupByStatus(val *TestPlanGroupByStatus) *NullableTestPlanGroupByStatus {
	return &NullableTestPlanGroupByStatus{value: val, isSet: true}
}

func (v NullableTestPlanGroupByStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPlanGroupByStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


