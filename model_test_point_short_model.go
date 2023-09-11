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

// checks if the TestPointShortModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPointShortModel{}

// TestPointShortModel struct for TestPointShortModel
type TestPointShortModel struct {
	TestSuiteId string `json:"testSuiteId"`
	Id string `json:"id"`
	TesterId NullableString `json:"testerId,omitempty"`
	WorkItemId NullableString `json:"workItemId,omitempty"`
	ConfigurationId NullableString `json:"configurationId,omitempty"`
	// Applies one of these values: Blocked, NoResults, Failed, Passed
	Status NullableString `json:"status,omitempty"`
	LastTestResultId NullableString `json:"lastTestResultId,omitempty"`
	IterationId string `json:"iterationId"`
}

// NewTestPointShortModel instantiates a new TestPointShortModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPointShortModel(testSuiteId string, id string, iterationId string) *TestPointShortModel {
	this := TestPointShortModel{}
	this.TestSuiteId = testSuiteId
	this.Id = id
	this.IterationId = iterationId
	return &this
}

// NewTestPointShortModelWithDefaults instantiates a new TestPointShortModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPointShortModelWithDefaults() *TestPointShortModel {
	this := TestPointShortModel{}
	return &this
}

// GetTestSuiteId returns the TestSuiteId field value
func (o *TestPointShortModel) GetTestSuiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestSuiteId
}

// GetTestSuiteIdOk returns a tuple with the TestSuiteId field value
// and a boolean to check if the value has been set.
func (o *TestPointShortModel) GetTestSuiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestSuiteId, true
}

// SetTestSuiteId sets field value
func (o *TestPointShortModel) SetTestSuiteId(v string) {
	o.TestSuiteId = v
}

// GetId returns the Id field value
func (o *TestPointShortModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestPointShortModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestPointShortModel) SetId(v string) {
	o.Id = v
}

// GetTesterId returns the TesterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortModel) GetTesterId() string {
	if o == nil || IsNil(o.TesterId.Get()) {
		var ret string
		return ret
	}
	return *o.TesterId.Get()
}

// GetTesterIdOk returns a tuple with the TesterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortModel) GetTesterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TesterId.Get(), o.TesterId.IsSet()
}

// HasTesterId returns a boolean if a field has been set.
func (o *TestPointShortModel) HasTesterId() bool {
	if o != nil && o.TesterId.IsSet() {
		return true
	}

	return false
}

// SetTesterId gets a reference to the given NullableString and assigns it to the TesterId field.
func (o *TestPointShortModel) SetTesterId(v string) {
	o.TesterId.Set(&v)
}
// SetTesterIdNil sets the value for TesterId to be an explicit nil
func (o *TestPointShortModel) SetTesterIdNil() {
	o.TesterId.Set(nil)
}

// UnsetTesterId ensures that no value is present for TesterId, not even an explicit nil
func (o *TestPointShortModel) UnsetTesterId() {
	o.TesterId.Unset()
}

// GetWorkItemId returns the WorkItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortModel) GetWorkItemId() string {
	if o == nil || IsNil(o.WorkItemId.Get()) {
		var ret string
		return ret
	}
	return *o.WorkItemId.Get()
}

// GetWorkItemIdOk returns a tuple with the WorkItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortModel) GetWorkItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkItemId.Get(), o.WorkItemId.IsSet()
}

// HasWorkItemId returns a boolean if a field has been set.
func (o *TestPointShortModel) HasWorkItemId() bool {
	if o != nil && o.WorkItemId.IsSet() {
		return true
	}

	return false
}

// SetWorkItemId gets a reference to the given NullableString and assigns it to the WorkItemId field.
func (o *TestPointShortModel) SetWorkItemId(v string) {
	o.WorkItemId.Set(&v)
}
// SetWorkItemIdNil sets the value for WorkItemId to be an explicit nil
func (o *TestPointShortModel) SetWorkItemIdNil() {
	o.WorkItemId.Set(nil)
}

// UnsetWorkItemId ensures that no value is present for WorkItemId, not even an explicit nil
func (o *TestPointShortModel) UnsetWorkItemId() {
	o.WorkItemId.Unset()
}

// GetConfigurationId returns the ConfigurationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortModel) GetConfigurationId() string {
	if o == nil || IsNil(o.ConfigurationId.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigurationId.Get()
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortModel) GetConfigurationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigurationId.Get(), o.ConfigurationId.IsSet()
}

// HasConfigurationId returns a boolean if a field has been set.
func (o *TestPointShortModel) HasConfigurationId() bool {
	if o != nil && o.ConfigurationId.IsSet() {
		return true
	}

	return false
}

// SetConfigurationId gets a reference to the given NullableString and assigns it to the ConfigurationId field.
func (o *TestPointShortModel) SetConfigurationId(v string) {
	o.ConfigurationId.Set(&v)
}
// SetConfigurationIdNil sets the value for ConfigurationId to be an explicit nil
func (o *TestPointShortModel) SetConfigurationIdNil() {
	o.ConfigurationId.Set(nil)
}

// UnsetConfigurationId ensures that no value is present for ConfigurationId, not even an explicit nil
func (o *TestPointShortModel) UnsetConfigurationId() {
	o.ConfigurationId.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortModel) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortModel) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *TestPointShortModel) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *TestPointShortModel) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *TestPointShortModel) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *TestPointShortModel) UnsetStatus() {
	o.Status.Unset()
}

// GetLastTestResultId returns the LastTestResultId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortModel) GetLastTestResultId() string {
	if o == nil || IsNil(o.LastTestResultId.Get()) {
		var ret string
		return ret
	}
	return *o.LastTestResultId.Get()
}

// GetLastTestResultIdOk returns a tuple with the LastTestResultId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortModel) GetLastTestResultIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastTestResultId.Get(), o.LastTestResultId.IsSet()
}

// HasLastTestResultId returns a boolean if a field has been set.
func (o *TestPointShortModel) HasLastTestResultId() bool {
	if o != nil && o.LastTestResultId.IsSet() {
		return true
	}

	return false
}

// SetLastTestResultId gets a reference to the given NullableString and assigns it to the LastTestResultId field.
func (o *TestPointShortModel) SetLastTestResultId(v string) {
	o.LastTestResultId.Set(&v)
}
// SetLastTestResultIdNil sets the value for LastTestResultId to be an explicit nil
func (o *TestPointShortModel) SetLastTestResultIdNil() {
	o.LastTestResultId.Set(nil)
}

// UnsetLastTestResultId ensures that no value is present for LastTestResultId, not even an explicit nil
func (o *TestPointShortModel) UnsetLastTestResultId() {
	o.LastTestResultId.Unset()
}

// GetIterationId returns the IterationId field value
func (o *TestPointShortModel) GetIterationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IterationId
}

// GetIterationIdOk returns a tuple with the IterationId field value
// and a boolean to check if the value has been set.
func (o *TestPointShortModel) GetIterationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IterationId, true
}

// SetIterationId sets field value
func (o *TestPointShortModel) SetIterationId(v string) {
	o.IterationId = v
}

func (o TestPointShortModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPointShortModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["testSuiteId"] = o.TestSuiteId
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
	if o.LastTestResultId.IsSet() {
		toSerialize["lastTestResultId"] = o.LastTestResultId.Get()
	}
	toSerialize["iterationId"] = o.IterationId
	return toSerialize, nil
}

type NullableTestPointShortModel struct {
	value *TestPointShortModel
	isSet bool
}

func (v NullableTestPointShortModel) Get() *TestPointShortModel {
	return v.value
}

func (v *NullableTestPointShortModel) Set(val *TestPointShortModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPointShortModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPointShortModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPointShortModel(val *TestPointShortModel) *NullableTestPointShortModel {
	return &NullableTestPointShortModel{value: val, isSet: true}
}

func (v NullableTestPointShortModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPointShortModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


