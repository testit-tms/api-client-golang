/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
)

// checks if the TestPlanWithAnalyticModelAnalytic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPlanWithAnalyticModelAnalytic{}

// TestPlanWithAnalyticModelAnalytic struct for TestPlanWithAnalyticModelAnalytic
type TestPlanWithAnalyticModelAnalytic struct {
	CountGroupByStatus []TestPlanGroupByStatus `json:"countGroupByStatus"`
	SumGroupByTester []TestPlanGroupByTester `json:"sumGroupByTester"`
	CountGroupByTester []TestPlanGroupByTester `json:"countGroupByTester"`
	CountGroupByTestSuite []TestPlanGroupByTestSuite `json:"countGroupByTestSuite"`
	CountGroupByTesterAndStatus []TestPlanGroupByTesterAndStatus `json:"countGroupByTesterAndStatus"`
}

// NewTestPlanWithAnalyticModelAnalytic instantiates a new TestPlanWithAnalyticModelAnalytic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPlanWithAnalyticModelAnalytic(countGroupByStatus []TestPlanGroupByStatus, sumGroupByTester []TestPlanGroupByTester, countGroupByTester []TestPlanGroupByTester, countGroupByTestSuite []TestPlanGroupByTestSuite, countGroupByTesterAndStatus []TestPlanGroupByTesterAndStatus) *TestPlanWithAnalyticModelAnalytic {
	this := TestPlanWithAnalyticModelAnalytic{}
	this.CountGroupByStatus = countGroupByStatus
	this.SumGroupByTester = sumGroupByTester
	this.CountGroupByTester = countGroupByTester
	this.CountGroupByTestSuite = countGroupByTestSuite
	this.CountGroupByTesterAndStatus = countGroupByTesterAndStatus
	return &this
}

// NewTestPlanWithAnalyticModelAnalyticWithDefaults instantiates a new TestPlanWithAnalyticModelAnalytic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPlanWithAnalyticModelAnalyticWithDefaults() *TestPlanWithAnalyticModelAnalytic {
	this := TestPlanWithAnalyticModelAnalytic{}
	return &this
}

// GetCountGroupByStatus returns the CountGroupByStatus field value
func (o *TestPlanWithAnalyticModelAnalytic) GetCountGroupByStatus() []TestPlanGroupByStatus {
	if o == nil {
		var ret []TestPlanGroupByStatus
		return ret
	}

	return o.CountGroupByStatus
}

// GetCountGroupByStatusOk returns a tuple with the CountGroupByStatus field value
// and a boolean to check if the value has been set.
func (o *TestPlanWithAnalyticModelAnalytic) GetCountGroupByStatusOk() ([]TestPlanGroupByStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountGroupByStatus, true
}

// SetCountGroupByStatus sets field value
func (o *TestPlanWithAnalyticModelAnalytic) SetCountGroupByStatus(v []TestPlanGroupByStatus) {
	o.CountGroupByStatus = v
}

// GetSumGroupByTester returns the SumGroupByTester field value
func (o *TestPlanWithAnalyticModelAnalytic) GetSumGroupByTester() []TestPlanGroupByTester {
	if o == nil {
		var ret []TestPlanGroupByTester
		return ret
	}

	return o.SumGroupByTester
}

// GetSumGroupByTesterOk returns a tuple with the SumGroupByTester field value
// and a boolean to check if the value has been set.
func (o *TestPlanWithAnalyticModelAnalytic) GetSumGroupByTesterOk() ([]TestPlanGroupByTester, bool) {
	if o == nil {
		return nil, false
	}
	return o.SumGroupByTester, true
}

// SetSumGroupByTester sets field value
func (o *TestPlanWithAnalyticModelAnalytic) SetSumGroupByTester(v []TestPlanGroupByTester) {
	o.SumGroupByTester = v
}

// GetCountGroupByTester returns the CountGroupByTester field value
func (o *TestPlanWithAnalyticModelAnalytic) GetCountGroupByTester() []TestPlanGroupByTester {
	if o == nil {
		var ret []TestPlanGroupByTester
		return ret
	}

	return o.CountGroupByTester
}

// GetCountGroupByTesterOk returns a tuple with the CountGroupByTester field value
// and a boolean to check if the value has been set.
func (o *TestPlanWithAnalyticModelAnalytic) GetCountGroupByTesterOk() ([]TestPlanGroupByTester, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountGroupByTester, true
}

// SetCountGroupByTester sets field value
func (o *TestPlanWithAnalyticModelAnalytic) SetCountGroupByTester(v []TestPlanGroupByTester) {
	o.CountGroupByTester = v
}

// GetCountGroupByTestSuite returns the CountGroupByTestSuite field value
func (o *TestPlanWithAnalyticModelAnalytic) GetCountGroupByTestSuite() []TestPlanGroupByTestSuite {
	if o == nil {
		var ret []TestPlanGroupByTestSuite
		return ret
	}

	return o.CountGroupByTestSuite
}

// GetCountGroupByTestSuiteOk returns a tuple with the CountGroupByTestSuite field value
// and a boolean to check if the value has been set.
func (o *TestPlanWithAnalyticModelAnalytic) GetCountGroupByTestSuiteOk() ([]TestPlanGroupByTestSuite, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountGroupByTestSuite, true
}

// SetCountGroupByTestSuite sets field value
func (o *TestPlanWithAnalyticModelAnalytic) SetCountGroupByTestSuite(v []TestPlanGroupByTestSuite) {
	o.CountGroupByTestSuite = v
}

// GetCountGroupByTesterAndStatus returns the CountGroupByTesterAndStatus field value
func (o *TestPlanWithAnalyticModelAnalytic) GetCountGroupByTesterAndStatus() []TestPlanGroupByTesterAndStatus {
	if o == nil {
		var ret []TestPlanGroupByTesterAndStatus
		return ret
	}

	return o.CountGroupByTesterAndStatus
}

// GetCountGroupByTesterAndStatusOk returns a tuple with the CountGroupByTesterAndStatus field value
// and a boolean to check if the value has been set.
func (o *TestPlanWithAnalyticModelAnalytic) GetCountGroupByTesterAndStatusOk() ([]TestPlanGroupByTesterAndStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountGroupByTesterAndStatus, true
}

// SetCountGroupByTesterAndStatus sets field value
func (o *TestPlanWithAnalyticModelAnalytic) SetCountGroupByTesterAndStatus(v []TestPlanGroupByTesterAndStatus) {
	o.CountGroupByTesterAndStatus = v
}

func (o TestPlanWithAnalyticModelAnalytic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPlanWithAnalyticModelAnalytic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["countGroupByStatus"] = o.CountGroupByStatus
	toSerialize["sumGroupByTester"] = o.SumGroupByTester
	toSerialize["countGroupByTester"] = o.CountGroupByTester
	toSerialize["countGroupByTestSuite"] = o.CountGroupByTestSuite
	toSerialize["countGroupByTesterAndStatus"] = o.CountGroupByTesterAndStatus
	return toSerialize, nil
}

type NullableTestPlanWithAnalyticModelAnalytic struct {
	value *TestPlanWithAnalyticModelAnalytic
	isSet bool
}

func (v NullableTestPlanWithAnalyticModelAnalytic) Get() *TestPlanWithAnalyticModelAnalytic {
	return v.value
}

func (v *NullableTestPlanWithAnalyticModelAnalytic) Set(val *TestPlanWithAnalyticModelAnalytic) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPlanWithAnalyticModelAnalytic) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPlanWithAnalyticModelAnalytic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPlanWithAnalyticModelAnalytic(val *TestPlanWithAnalyticModelAnalytic) *NullableTestPlanWithAnalyticModelAnalytic {
	return &NullableTestPlanWithAnalyticModelAnalytic{value: val, isSet: true}
}

func (v NullableTestPlanWithAnalyticModelAnalytic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPlanWithAnalyticModelAnalytic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

