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

// checks if the TestPlanTestPointsAnalyticsApiResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPlanTestPointsAnalyticsApiResult{}

// TestPlanTestPointsAnalyticsApiResult struct for TestPlanTestPointsAnalyticsApiResult
type TestPlanTestPointsAnalyticsApiResult struct {
	CountGroupByStatus []TestPlanTestPointsStatusGroupApiResult `json:"countGroupByStatus"`
	SumGroupByTester []TestPlanTestPointsTesterGroupApiResult `json:"sumGroupByTester"`
	CountGroupByTester []TestPlanTestPointsTesterGroupApiResult `json:"countGroupByTester"`
	CountGroupByTesterAndStatus []TestPlanTestPointsTesterAndStatusGroupApiResult `json:"countGroupByTesterAndStatus"`
}

type _TestPlanTestPointsAnalyticsApiResult TestPlanTestPointsAnalyticsApiResult

// NewTestPlanTestPointsAnalyticsApiResult instantiates a new TestPlanTestPointsAnalyticsApiResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPlanTestPointsAnalyticsApiResult(countGroupByStatus []TestPlanTestPointsStatusGroupApiResult, sumGroupByTester []TestPlanTestPointsTesterGroupApiResult, countGroupByTester []TestPlanTestPointsTesterGroupApiResult, countGroupByTesterAndStatus []TestPlanTestPointsTesterAndStatusGroupApiResult) *TestPlanTestPointsAnalyticsApiResult {
	this := TestPlanTestPointsAnalyticsApiResult{}
	this.CountGroupByStatus = countGroupByStatus
	this.SumGroupByTester = sumGroupByTester
	this.CountGroupByTester = countGroupByTester
	this.CountGroupByTesterAndStatus = countGroupByTesterAndStatus
	return &this
}

// NewTestPlanTestPointsAnalyticsApiResultWithDefaults instantiates a new TestPlanTestPointsAnalyticsApiResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPlanTestPointsAnalyticsApiResultWithDefaults() *TestPlanTestPointsAnalyticsApiResult {
	this := TestPlanTestPointsAnalyticsApiResult{}
	return &this
}

// GetCountGroupByStatus returns the CountGroupByStatus field value
func (o *TestPlanTestPointsAnalyticsApiResult) GetCountGroupByStatus() []TestPlanTestPointsStatusGroupApiResult {
	if o == nil {
		var ret []TestPlanTestPointsStatusGroupApiResult
		return ret
	}

	return o.CountGroupByStatus
}

// GetCountGroupByStatusOk returns a tuple with the CountGroupByStatus field value
// and a boolean to check if the value has been set.
func (o *TestPlanTestPointsAnalyticsApiResult) GetCountGroupByStatusOk() ([]TestPlanTestPointsStatusGroupApiResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountGroupByStatus, true
}

// SetCountGroupByStatus sets field value
func (o *TestPlanTestPointsAnalyticsApiResult) SetCountGroupByStatus(v []TestPlanTestPointsStatusGroupApiResult) {
	o.CountGroupByStatus = v
}

// GetSumGroupByTester returns the SumGroupByTester field value
func (o *TestPlanTestPointsAnalyticsApiResult) GetSumGroupByTester() []TestPlanTestPointsTesterGroupApiResult {
	if o == nil {
		var ret []TestPlanTestPointsTesterGroupApiResult
		return ret
	}

	return o.SumGroupByTester
}

// GetSumGroupByTesterOk returns a tuple with the SumGroupByTester field value
// and a boolean to check if the value has been set.
func (o *TestPlanTestPointsAnalyticsApiResult) GetSumGroupByTesterOk() ([]TestPlanTestPointsTesterGroupApiResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.SumGroupByTester, true
}

// SetSumGroupByTester sets field value
func (o *TestPlanTestPointsAnalyticsApiResult) SetSumGroupByTester(v []TestPlanTestPointsTesterGroupApiResult) {
	o.SumGroupByTester = v
}

// GetCountGroupByTester returns the CountGroupByTester field value
func (o *TestPlanTestPointsAnalyticsApiResult) GetCountGroupByTester() []TestPlanTestPointsTesterGroupApiResult {
	if o == nil {
		var ret []TestPlanTestPointsTesterGroupApiResult
		return ret
	}

	return o.CountGroupByTester
}

// GetCountGroupByTesterOk returns a tuple with the CountGroupByTester field value
// and a boolean to check if the value has been set.
func (o *TestPlanTestPointsAnalyticsApiResult) GetCountGroupByTesterOk() ([]TestPlanTestPointsTesterGroupApiResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountGroupByTester, true
}

// SetCountGroupByTester sets field value
func (o *TestPlanTestPointsAnalyticsApiResult) SetCountGroupByTester(v []TestPlanTestPointsTesterGroupApiResult) {
	o.CountGroupByTester = v
}

// GetCountGroupByTesterAndStatus returns the CountGroupByTesterAndStatus field value
func (o *TestPlanTestPointsAnalyticsApiResult) GetCountGroupByTesterAndStatus() []TestPlanTestPointsTesterAndStatusGroupApiResult {
	if o == nil {
		var ret []TestPlanTestPointsTesterAndStatusGroupApiResult
		return ret
	}

	return o.CountGroupByTesterAndStatus
}

// GetCountGroupByTesterAndStatusOk returns a tuple with the CountGroupByTesterAndStatus field value
// and a boolean to check if the value has been set.
func (o *TestPlanTestPointsAnalyticsApiResult) GetCountGroupByTesterAndStatusOk() ([]TestPlanTestPointsTesterAndStatusGroupApiResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountGroupByTesterAndStatus, true
}

// SetCountGroupByTesterAndStatus sets field value
func (o *TestPlanTestPointsAnalyticsApiResult) SetCountGroupByTesterAndStatus(v []TestPlanTestPointsTesterAndStatusGroupApiResult) {
	o.CountGroupByTesterAndStatus = v
}

func (o TestPlanTestPointsAnalyticsApiResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPlanTestPointsAnalyticsApiResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["countGroupByStatus"] = o.CountGroupByStatus
	toSerialize["sumGroupByTester"] = o.SumGroupByTester
	toSerialize["countGroupByTester"] = o.CountGroupByTester
	toSerialize["countGroupByTesterAndStatus"] = o.CountGroupByTesterAndStatus
	return toSerialize, nil
}

func (o *TestPlanTestPointsAnalyticsApiResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"countGroupByStatus",
		"sumGroupByTester",
		"countGroupByTester",
		"countGroupByTesterAndStatus",
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

	varTestPlanTestPointsAnalyticsApiResult := _TestPlanTestPointsAnalyticsApiResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestPlanTestPointsAnalyticsApiResult)

	if err != nil {
		return err
	}

	*o = TestPlanTestPointsAnalyticsApiResult(varTestPlanTestPointsAnalyticsApiResult)

	return err
}

type NullableTestPlanTestPointsAnalyticsApiResult struct {
	value *TestPlanTestPointsAnalyticsApiResult
	isSet bool
}

func (v NullableTestPlanTestPointsAnalyticsApiResult) Get() *TestPlanTestPointsAnalyticsApiResult {
	return v.value
}

func (v *NullableTestPlanTestPointsAnalyticsApiResult) Set(val *TestPlanTestPointsAnalyticsApiResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPlanTestPointsAnalyticsApiResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPlanTestPointsAnalyticsApiResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPlanTestPointsAnalyticsApiResult(val *TestPlanTestPointsAnalyticsApiResult) *NullableTestPlanTestPointsAnalyticsApiResult {
	return &NullableTestPlanTestPointsAnalyticsApiResult{value: val, isSet: true}
}

func (v NullableTestPlanTestPointsAnalyticsApiResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPlanTestPointsAnalyticsApiResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


