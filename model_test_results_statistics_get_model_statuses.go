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

// checks if the TestResultsStatisticsGetModelStatuses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestResultsStatisticsGetModelStatuses{}

// TestResultsStatisticsGetModelStatuses Test results counts aggregated by outcome
type TestResultsStatisticsGetModelStatuses struct {
	// Number of test results which is running currently
	InProgress int32 `json:"inProgress"`
	// Number of test results which successfully passed
	Passed int32 `json:"passed"`
	// Number of test results which failed with an error
	Failed int32 `json:"failed"`
	// Number of test results which did not run and were skipped
	Skipped int32 `json:"skipped"`
	// Number of test results which cannot be launched
	Blocked int32 `json:"blocked"`
}

// NewTestResultsStatisticsGetModelStatuses instantiates a new TestResultsStatisticsGetModelStatuses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestResultsStatisticsGetModelStatuses(inProgress int32, passed int32, failed int32, skipped int32, blocked int32) *TestResultsStatisticsGetModelStatuses {
	this := TestResultsStatisticsGetModelStatuses{}
	this.InProgress = inProgress
	this.Passed = passed
	this.Failed = failed
	this.Skipped = skipped
	this.Blocked = blocked
	return &this
}

// NewTestResultsStatisticsGetModelStatusesWithDefaults instantiates a new TestResultsStatisticsGetModelStatuses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestResultsStatisticsGetModelStatusesWithDefaults() *TestResultsStatisticsGetModelStatuses {
	this := TestResultsStatisticsGetModelStatuses{}
	return &this
}

// GetInProgress returns the InProgress field value
func (o *TestResultsStatisticsGetModelStatuses) GetInProgress() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InProgress
}

// GetInProgressOk returns a tuple with the InProgress field value
// and a boolean to check if the value has been set.
func (o *TestResultsStatisticsGetModelStatuses) GetInProgressOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InProgress, true
}

// SetInProgress sets field value
func (o *TestResultsStatisticsGetModelStatuses) SetInProgress(v int32) {
	o.InProgress = v
}

// GetPassed returns the Passed field value
func (o *TestResultsStatisticsGetModelStatuses) GetPassed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Passed
}

// GetPassedOk returns a tuple with the Passed field value
// and a boolean to check if the value has been set.
func (o *TestResultsStatisticsGetModelStatuses) GetPassedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Passed, true
}

// SetPassed sets field value
func (o *TestResultsStatisticsGetModelStatuses) SetPassed(v int32) {
	o.Passed = v
}

// GetFailed returns the Failed field value
func (o *TestResultsStatisticsGetModelStatuses) GetFailed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Failed
}

// GetFailedOk returns a tuple with the Failed field value
// and a boolean to check if the value has been set.
func (o *TestResultsStatisticsGetModelStatuses) GetFailedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Failed, true
}

// SetFailed sets field value
func (o *TestResultsStatisticsGetModelStatuses) SetFailed(v int32) {
	o.Failed = v
}

// GetSkipped returns the Skipped field value
func (o *TestResultsStatisticsGetModelStatuses) GetSkipped() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Skipped
}

// GetSkippedOk returns a tuple with the Skipped field value
// and a boolean to check if the value has been set.
func (o *TestResultsStatisticsGetModelStatuses) GetSkippedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Skipped, true
}

// SetSkipped sets field value
func (o *TestResultsStatisticsGetModelStatuses) SetSkipped(v int32) {
	o.Skipped = v
}

// GetBlocked returns the Blocked field value
func (o *TestResultsStatisticsGetModelStatuses) GetBlocked() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Blocked
}

// GetBlockedOk returns a tuple with the Blocked field value
// and a boolean to check if the value has been set.
func (o *TestResultsStatisticsGetModelStatuses) GetBlockedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blocked, true
}

// SetBlocked sets field value
func (o *TestResultsStatisticsGetModelStatuses) SetBlocked(v int32) {
	o.Blocked = v
}

func (o TestResultsStatisticsGetModelStatuses) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestResultsStatisticsGetModelStatuses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inProgress"] = o.InProgress
	toSerialize["passed"] = o.Passed
	toSerialize["failed"] = o.Failed
	toSerialize["skipped"] = o.Skipped
	toSerialize["blocked"] = o.Blocked
	return toSerialize, nil
}

type NullableTestResultsStatisticsGetModelStatuses struct {
	value *TestResultsStatisticsGetModelStatuses
	isSet bool
}

func (v NullableTestResultsStatisticsGetModelStatuses) Get() *TestResultsStatisticsGetModelStatuses {
	return v.value
}

func (v *NullableTestResultsStatisticsGetModelStatuses) Set(val *TestResultsStatisticsGetModelStatuses) {
	v.value = val
	v.isSet = true
}

func (v NullableTestResultsStatisticsGetModelStatuses) IsSet() bool {
	return v.isSet
}

func (v *NullableTestResultsStatisticsGetModelStatuses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestResultsStatisticsGetModelStatuses(val *TestResultsStatisticsGetModelStatuses) *NullableTestResultsStatisticsGetModelStatuses {
	return &NullableTestResultsStatisticsGetModelStatuses{value: val, isSet: true}
}

func (v NullableTestResultsStatisticsGetModelStatuses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestResultsStatisticsGetModelStatuses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

