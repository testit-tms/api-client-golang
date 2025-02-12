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

// checks if the TestPointShortApiResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPointShortApiResult{}

// TestPointShortApiResult struct for TestPointShortApiResult
type TestPointShortApiResult struct {
	// Test point unique internal identifier
	Id string `json:"id"`
	// Tester who is responded for the test unique internal identifier
	TesterId NullableString `json:"testerId,omitempty"`
	// Workitem to which test point relates unique identifier
	WorkItemId NullableString `json:"workItemId,omitempty"`
	// Configuration to which test point relates unique identifier
	ConfigurationId NullableString `json:"configurationId,omitempty"`
	// Test point status   Applies one of these values: Blocked, NoResults, Failed, Passed
	// Deprecated
	Status NullableString `json:"status,omitempty"`
	// Test point status
	StatusModel TestStatusApiResult `json:"statusModel"`
	// Last test result unique identifier
	LastTestResultId NullableString `json:"lastTestResultId,omitempty"`
	// Iteration unique identifier
	IterationId string `json:"iterationId"`
	// Median duration of work item the test point represents
	WorkItemMedianDuration NullableInt64 `json:"workItemMedianDuration,omitempty"`
	// Test suite to which test point relates unique identifier
	TestSuiteId string `json:"testSuiteId"`
}

type _TestPointShortApiResult TestPointShortApiResult

// NewTestPointShortApiResult instantiates a new TestPointShortApiResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPointShortApiResult(id string, statusModel TestStatusApiResult, iterationId string, testSuiteId string) *TestPointShortApiResult {
	this := TestPointShortApiResult{}
	this.Id = id
	this.StatusModel = statusModel
	this.IterationId = iterationId
	this.TestSuiteId = testSuiteId
	return &this
}

// NewTestPointShortApiResultWithDefaults instantiates a new TestPointShortApiResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPointShortApiResultWithDefaults() *TestPointShortApiResult {
	this := TestPointShortApiResult{}
	return &this
}

// GetId returns the Id field value
func (o *TestPointShortApiResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestPointShortApiResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestPointShortApiResult) SetId(v string) {
	o.Id = v
}

// GetTesterId returns the TesterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortApiResult) GetTesterId() string {
	if o == nil || IsNil(o.TesterId.Get()) {
		var ret string
		return ret
	}
	return *o.TesterId.Get()
}

// GetTesterIdOk returns a tuple with the TesterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortApiResult) GetTesterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TesterId.Get(), o.TesterId.IsSet()
}

// HasTesterId returns a boolean if a field has been set.
func (o *TestPointShortApiResult) HasTesterId() bool {
	if o != nil && o.TesterId.IsSet() {
		return true
	}

	return false
}

// SetTesterId gets a reference to the given NullableString and assigns it to the TesterId field.
func (o *TestPointShortApiResult) SetTesterId(v string) {
	o.TesterId.Set(&v)
}
// SetTesterIdNil sets the value for TesterId to be an explicit nil
func (o *TestPointShortApiResult) SetTesterIdNil() {
	o.TesterId.Set(nil)
}

// UnsetTesterId ensures that no value is present for TesterId, not even an explicit nil
func (o *TestPointShortApiResult) UnsetTesterId() {
	o.TesterId.Unset()
}

// GetWorkItemId returns the WorkItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortApiResult) GetWorkItemId() string {
	if o == nil || IsNil(o.WorkItemId.Get()) {
		var ret string
		return ret
	}
	return *o.WorkItemId.Get()
}

// GetWorkItemIdOk returns a tuple with the WorkItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortApiResult) GetWorkItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkItemId.Get(), o.WorkItemId.IsSet()
}

// HasWorkItemId returns a boolean if a field has been set.
func (o *TestPointShortApiResult) HasWorkItemId() bool {
	if o != nil && o.WorkItemId.IsSet() {
		return true
	}

	return false
}

// SetWorkItemId gets a reference to the given NullableString and assigns it to the WorkItemId field.
func (o *TestPointShortApiResult) SetWorkItemId(v string) {
	o.WorkItemId.Set(&v)
}
// SetWorkItemIdNil sets the value for WorkItemId to be an explicit nil
func (o *TestPointShortApiResult) SetWorkItemIdNil() {
	o.WorkItemId.Set(nil)
}

// UnsetWorkItemId ensures that no value is present for WorkItemId, not even an explicit nil
func (o *TestPointShortApiResult) UnsetWorkItemId() {
	o.WorkItemId.Unset()
}

// GetConfigurationId returns the ConfigurationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortApiResult) GetConfigurationId() string {
	if o == nil || IsNil(o.ConfigurationId.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigurationId.Get()
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortApiResult) GetConfigurationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigurationId.Get(), o.ConfigurationId.IsSet()
}

// HasConfigurationId returns a boolean if a field has been set.
func (o *TestPointShortApiResult) HasConfigurationId() bool {
	if o != nil && o.ConfigurationId.IsSet() {
		return true
	}

	return false
}

// SetConfigurationId gets a reference to the given NullableString and assigns it to the ConfigurationId field.
func (o *TestPointShortApiResult) SetConfigurationId(v string) {
	o.ConfigurationId.Set(&v)
}
// SetConfigurationIdNil sets the value for ConfigurationId to be an explicit nil
func (o *TestPointShortApiResult) SetConfigurationIdNil() {
	o.ConfigurationId.Set(nil)
}

// UnsetConfigurationId ensures that no value is present for ConfigurationId, not even an explicit nil
func (o *TestPointShortApiResult) UnsetConfigurationId() {
	o.ConfigurationId.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *TestPointShortApiResult) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *TestPointShortApiResult) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *TestPointShortApiResult) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
// Deprecated
func (o *TestPointShortApiResult) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *TestPointShortApiResult) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *TestPointShortApiResult) UnsetStatus() {
	o.Status.Unset()
}

// GetStatusModel returns the StatusModel field value
func (o *TestPointShortApiResult) GetStatusModel() TestStatusApiResult {
	if o == nil {
		var ret TestStatusApiResult
		return ret
	}

	return o.StatusModel
}

// GetStatusModelOk returns a tuple with the StatusModel field value
// and a boolean to check if the value has been set.
func (o *TestPointShortApiResult) GetStatusModelOk() (*TestStatusApiResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusModel, true
}

// SetStatusModel sets field value
func (o *TestPointShortApiResult) SetStatusModel(v TestStatusApiResult) {
	o.StatusModel = v
}

// GetLastTestResultId returns the LastTestResultId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortApiResult) GetLastTestResultId() string {
	if o == nil || IsNil(o.LastTestResultId.Get()) {
		var ret string
		return ret
	}
	return *o.LastTestResultId.Get()
}

// GetLastTestResultIdOk returns a tuple with the LastTestResultId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortApiResult) GetLastTestResultIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastTestResultId.Get(), o.LastTestResultId.IsSet()
}

// HasLastTestResultId returns a boolean if a field has been set.
func (o *TestPointShortApiResult) HasLastTestResultId() bool {
	if o != nil && o.LastTestResultId.IsSet() {
		return true
	}

	return false
}

// SetLastTestResultId gets a reference to the given NullableString and assigns it to the LastTestResultId field.
func (o *TestPointShortApiResult) SetLastTestResultId(v string) {
	o.LastTestResultId.Set(&v)
}
// SetLastTestResultIdNil sets the value for LastTestResultId to be an explicit nil
func (o *TestPointShortApiResult) SetLastTestResultIdNil() {
	o.LastTestResultId.Set(nil)
}

// UnsetLastTestResultId ensures that no value is present for LastTestResultId, not even an explicit nil
func (o *TestPointShortApiResult) UnsetLastTestResultId() {
	o.LastTestResultId.Unset()
}

// GetIterationId returns the IterationId field value
func (o *TestPointShortApiResult) GetIterationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IterationId
}

// GetIterationIdOk returns a tuple with the IterationId field value
// and a boolean to check if the value has been set.
func (o *TestPointShortApiResult) GetIterationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IterationId, true
}

// SetIterationId sets field value
func (o *TestPointShortApiResult) SetIterationId(v string) {
	o.IterationId = v
}

// GetWorkItemMedianDuration returns the WorkItemMedianDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortApiResult) GetWorkItemMedianDuration() int64 {
	if o == nil || IsNil(o.WorkItemMedianDuration.Get()) {
		var ret int64
		return ret
	}
	return *o.WorkItemMedianDuration.Get()
}

// GetWorkItemMedianDurationOk returns a tuple with the WorkItemMedianDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortApiResult) GetWorkItemMedianDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkItemMedianDuration.Get(), o.WorkItemMedianDuration.IsSet()
}

// HasWorkItemMedianDuration returns a boolean if a field has been set.
func (o *TestPointShortApiResult) HasWorkItemMedianDuration() bool {
	if o != nil && o.WorkItemMedianDuration.IsSet() {
		return true
	}

	return false
}

// SetWorkItemMedianDuration gets a reference to the given NullableInt64 and assigns it to the WorkItemMedianDuration field.
func (o *TestPointShortApiResult) SetWorkItemMedianDuration(v int64) {
	o.WorkItemMedianDuration.Set(&v)
}
// SetWorkItemMedianDurationNil sets the value for WorkItemMedianDuration to be an explicit nil
func (o *TestPointShortApiResult) SetWorkItemMedianDurationNil() {
	o.WorkItemMedianDuration.Set(nil)
}

// UnsetWorkItemMedianDuration ensures that no value is present for WorkItemMedianDuration, not even an explicit nil
func (o *TestPointShortApiResult) UnsetWorkItemMedianDuration() {
	o.WorkItemMedianDuration.Unset()
}

// GetTestSuiteId returns the TestSuiteId field value
func (o *TestPointShortApiResult) GetTestSuiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestSuiteId
}

// GetTestSuiteIdOk returns a tuple with the TestSuiteId field value
// and a boolean to check if the value has been set.
func (o *TestPointShortApiResult) GetTestSuiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestSuiteId, true
}

// SetTestSuiteId sets field value
func (o *TestPointShortApiResult) SetTestSuiteId(v string) {
	o.TestSuiteId = v
}

func (o TestPointShortApiResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPointShortApiResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.TesterId.IsSet() {
		toSerialize["testerId"] = o.TesterId.Get()
	}
	if o.WorkItemId.IsSet() {
		toSerialize["workItemId"] = o.WorkItemId.Get()
	}
	if o.ConfigurationId.IsSet() {
		toSerialize["configurationId"] = o.ConfigurationId.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	toSerialize["statusModel"] = o.StatusModel
	if o.LastTestResultId.IsSet() {
		toSerialize["lastTestResultId"] = o.LastTestResultId.Get()
	}
	toSerialize["iterationId"] = o.IterationId
	if o.WorkItemMedianDuration.IsSet() {
		toSerialize["workItemMedianDuration"] = o.WorkItemMedianDuration.Get()
	}
	toSerialize["testSuiteId"] = o.TestSuiteId
	return toSerialize, nil
}

func (o *TestPointShortApiResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"statusModel",
		"iterationId",
		"testSuiteId",
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

	varTestPointShortApiResult := _TestPointShortApiResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestPointShortApiResult)

	if err != nil {
		return err
	}

	*o = TestPointShortApiResult(varTestPointShortApiResult)

	return err
}

type NullableTestPointShortApiResult struct {
	value *TestPointShortApiResult
	isSet bool
}

func (v NullableTestPointShortApiResult) Get() *TestPointShortApiResult {
	return v.value
}

func (v *NullableTestPointShortApiResult) Set(val *TestPointShortApiResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPointShortApiResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPointShortApiResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPointShortApiResult(val *TestPointShortApiResult) *NullableTestPointShortApiResult {
	return &NullableTestPointShortApiResult{value: val, isSet: true}
}

func (v NullableTestPointShortApiResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPointShortApiResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


