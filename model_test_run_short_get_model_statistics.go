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

// checks if the TestRunShortGetModelStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestRunShortGetModelStatistics{}

// TestRunShortGetModelStatistics Statistics of the test run
type TestRunShortGetModelStatistics struct {
	Statuses TestResultsStatisticsGetModelStatuses `json:"statuses"`
	FailureCategories TestResultsStatisticsGetModelFailureCategories `json:"failureCategories"`
}

// NewTestRunShortGetModelStatistics instantiates a new TestRunShortGetModelStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestRunShortGetModelStatistics(statuses TestResultsStatisticsGetModelStatuses, failureCategories TestResultsStatisticsGetModelFailureCategories) *TestRunShortGetModelStatistics {
	this := TestRunShortGetModelStatistics{}
	this.Statuses = statuses
	this.FailureCategories = failureCategories
	return &this
}

// NewTestRunShortGetModelStatisticsWithDefaults instantiates a new TestRunShortGetModelStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestRunShortGetModelStatisticsWithDefaults() *TestRunShortGetModelStatistics {
	this := TestRunShortGetModelStatistics{}
	return &this
}

// GetStatuses returns the Statuses field value
func (o *TestRunShortGetModelStatistics) GetStatuses() TestResultsStatisticsGetModelStatuses {
	if o == nil {
		var ret TestResultsStatisticsGetModelStatuses
		return ret
	}

	return o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value
// and a boolean to check if the value has been set.
func (o *TestRunShortGetModelStatistics) GetStatusesOk() (*TestResultsStatisticsGetModelStatuses, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Statuses, true
}

// SetStatuses sets field value
func (o *TestRunShortGetModelStatistics) SetStatuses(v TestResultsStatisticsGetModelStatuses) {
	o.Statuses = v
}

// GetFailureCategories returns the FailureCategories field value
func (o *TestRunShortGetModelStatistics) GetFailureCategories() TestResultsStatisticsGetModelFailureCategories {
	if o == nil {
		var ret TestResultsStatisticsGetModelFailureCategories
		return ret
	}

	return o.FailureCategories
}

// GetFailureCategoriesOk returns a tuple with the FailureCategories field value
// and a boolean to check if the value has been set.
func (o *TestRunShortGetModelStatistics) GetFailureCategoriesOk() (*TestResultsStatisticsGetModelFailureCategories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailureCategories, true
}

// SetFailureCategories sets field value
func (o *TestRunShortGetModelStatistics) SetFailureCategories(v TestResultsStatisticsGetModelFailureCategories) {
	o.FailureCategories = v
}

func (o TestRunShortGetModelStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestRunShortGetModelStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["statuses"] = o.Statuses
	toSerialize["failureCategories"] = o.FailureCategories
	return toSerialize, nil
}

type NullableTestRunShortGetModelStatistics struct {
	value *TestRunShortGetModelStatistics
	isSet bool
}

func (v NullableTestRunShortGetModelStatistics) Get() *TestRunShortGetModelStatistics {
	return v.value
}

func (v *NullableTestRunShortGetModelStatistics) Set(val *TestRunShortGetModelStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableTestRunShortGetModelStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableTestRunShortGetModelStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestRunShortGetModelStatistics(val *TestRunShortGetModelStatistics) *NullableTestRunShortGetModelStatistics {
	return &NullableTestRunShortGetModelStatistics{value: val, isSet: true}
}

func (v NullableTestRunShortGetModelStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestRunShortGetModelStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


