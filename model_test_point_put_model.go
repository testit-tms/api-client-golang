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

// checks if the TestPointPutModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPointPutModel{}

// TestPointPutModel struct for TestPointPutModel
type TestPointPutModel struct {
	TesterId NullableString `json:"testerId,omitempty"`
	IterationId string `json:"iterationId"`
	WorkItemId NullableString `json:"workItemId,omitempty"`
	ConfigurationId NullableString `json:"configurationId,omitempty"`
	TestSuiteId string `json:"testSuiteId"`
	// Deprecated
	Status NullableString `json:"status,omitempty"`
	StatusModel TestStatusModel `json:"statusModel"`
	LastTestResultId NullableString `json:"lastTestResultId,omitempty"`
	// Unique ID of the entity
	Id string `json:"id"`
	// Indicates if the entity is deleted
	IsDeleted bool `json:"isDeleted"`
}

type _TestPointPutModel TestPointPutModel

// NewTestPointPutModel instantiates a new TestPointPutModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPointPutModel(iterationId string, testSuiteId string, statusModel TestStatusModel, id string, isDeleted bool) *TestPointPutModel {
	this := TestPointPutModel{}
	this.IterationId = iterationId
	this.TestSuiteId = testSuiteId
	this.StatusModel = statusModel
	this.Id = id
	this.IsDeleted = isDeleted
	return &this
}

// NewTestPointPutModelWithDefaults instantiates a new TestPointPutModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPointPutModelWithDefaults() *TestPointPutModel {
	this := TestPointPutModel{}
	return &this
}

// GetTesterId returns the TesterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointPutModel) GetTesterId() string {
	if o == nil || IsNil(o.TesterId.Get()) {
		var ret string
		return ret
	}
	return *o.TesterId.Get()
}

// GetTesterIdOk returns a tuple with the TesterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointPutModel) GetTesterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TesterId.Get(), o.TesterId.IsSet()
}

// HasTesterId returns a boolean if a field has been set.
func (o *TestPointPutModel) HasTesterId() bool {
	if o != nil && o.TesterId.IsSet() {
		return true
	}

	return false
}

// SetTesterId gets a reference to the given NullableString and assigns it to the TesterId field.
func (o *TestPointPutModel) SetTesterId(v string) {
	o.TesterId.Set(&v)
}
// SetTesterIdNil sets the value for TesterId to be an explicit nil
func (o *TestPointPutModel) SetTesterIdNil() {
	o.TesterId.Set(nil)
}

// UnsetTesterId ensures that no value is present for TesterId, not even an explicit nil
func (o *TestPointPutModel) UnsetTesterId() {
	o.TesterId.Unset()
}

// GetIterationId returns the IterationId field value
func (o *TestPointPutModel) GetIterationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IterationId
}

// GetIterationIdOk returns a tuple with the IterationId field value
// and a boolean to check if the value has been set.
func (o *TestPointPutModel) GetIterationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IterationId, true
}

// SetIterationId sets field value
func (o *TestPointPutModel) SetIterationId(v string) {
	o.IterationId = v
}

// GetWorkItemId returns the WorkItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointPutModel) GetWorkItemId() string {
	if o == nil || IsNil(o.WorkItemId.Get()) {
		var ret string
		return ret
	}
	return *o.WorkItemId.Get()
}

// GetWorkItemIdOk returns a tuple with the WorkItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointPutModel) GetWorkItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkItemId.Get(), o.WorkItemId.IsSet()
}

// HasWorkItemId returns a boolean if a field has been set.
func (o *TestPointPutModel) HasWorkItemId() bool {
	if o != nil && o.WorkItemId.IsSet() {
		return true
	}

	return false
}

// SetWorkItemId gets a reference to the given NullableString and assigns it to the WorkItemId field.
func (o *TestPointPutModel) SetWorkItemId(v string) {
	o.WorkItemId.Set(&v)
}
// SetWorkItemIdNil sets the value for WorkItemId to be an explicit nil
func (o *TestPointPutModel) SetWorkItemIdNil() {
	o.WorkItemId.Set(nil)
}

// UnsetWorkItemId ensures that no value is present for WorkItemId, not even an explicit nil
func (o *TestPointPutModel) UnsetWorkItemId() {
	o.WorkItemId.Unset()
}

// GetConfigurationId returns the ConfigurationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointPutModel) GetConfigurationId() string {
	if o == nil || IsNil(o.ConfigurationId.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigurationId.Get()
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointPutModel) GetConfigurationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigurationId.Get(), o.ConfigurationId.IsSet()
}

// HasConfigurationId returns a boolean if a field has been set.
func (o *TestPointPutModel) HasConfigurationId() bool {
	if o != nil && o.ConfigurationId.IsSet() {
		return true
	}

	return false
}

// SetConfigurationId gets a reference to the given NullableString and assigns it to the ConfigurationId field.
func (o *TestPointPutModel) SetConfigurationId(v string) {
	o.ConfigurationId.Set(&v)
}
// SetConfigurationIdNil sets the value for ConfigurationId to be an explicit nil
func (o *TestPointPutModel) SetConfigurationIdNil() {
	o.ConfigurationId.Set(nil)
}

// UnsetConfigurationId ensures that no value is present for ConfigurationId, not even an explicit nil
func (o *TestPointPutModel) UnsetConfigurationId() {
	o.ConfigurationId.Unset()
}

// GetTestSuiteId returns the TestSuiteId field value
func (o *TestPointPutModel) GetTestSuiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestSuiteId
}

// GetTestSuiteIdOk returns a tuple with the TestSuiteId field value
// and a boolean to check if the value has been set.
func (o *TestPointPutModel) GetTestSuiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestSuiteId, true
}

// SetTestSuiteId sets field value
func (o *TestPointPutModel) SetTestSuiteId(v string) {
	o.TestSuiteId = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *TestPointPutModel) GetStatus() string {
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
func (o *TestPointPutModel) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *TestPointPutModel) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
// Deprecated
func (o *TestPointPutModel) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *TestPointPutModel) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *TestPointPutModel) UnsetStatus() {
	o.Status.Unset()
}

// GetStatusModel returns the StatusModel field value
func (o *TestPointPutModel) GetStatusModel() TestStatusModel {
	if o == nil {
		var ret TestStatusModel
		return ret
	}

	return o.StatusModel
}

// GetStatusModelOk returns a tuple with the StatusModel field value
// and a boolean to check if the value has been set.
func (o *TestPointPutModel) GetStatusModelOk() (*TestStatusModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusModel, true
}

// SetStatusModel sets field value
func (o *TestPointPutModel) SetStatusModel(v TestStatusModel) {
	o.StatusModel = v
}

// GetLastTestResultId returns the LastTestResultId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointPutModel) GetLastTestResultId() string {
	if o == nil || IsNil(o.LastTestResultId.Get()) {
		var ret string
		return ret
	}
	return *o.LastTestResultId.Get()
}

// GetLastTestResultIdOk returns a tuple with the LastTestResultId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointPutModel) GetLastTestResultIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastTestResultId.Get(), o.LastTestResultId.IsSet()
}

// HasLastTestResultId returns a boolean if a field has been set.
func (o *TestPointPutModel) HasLastTestResultId() bool {
	if o != nil && o.LastTestResultId.IsSet() {
		return true
	}

	return false
}

// SetLastTestResultId gets a reference to the given NullableString and assigns it to the LastTestResultId field.
func (o *TestPointPutModel) SetLastTestResultId(v string) {
	o.LastTestResultId.Set(&v)
}
// SetLastTestResultIdNil sets the value for LastTestResultId to be an explicit nil
func (o *TestPointPutModel) SetLastTestResultIdNil() {
	o.LastTestResultId.Set(nil)
}

// UnsetLastTestResultId ensures that no value is present for LastTestResultId, not even an explicit nil
func (o *TestPointPutModel) UnsetLastTestResultId() {
	o.LastTestResultId.Unset()
}

// GetId returns the Id field value
func (o *TestPointPutModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestPointPutModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestPointPutModel) SetId(v string) {
	o.Id = v
}

// GetIsDeleted returns the IsDeleted field value
func (o *TestPointPutModel) GetIsDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value
// and a boolean to check if the value has been set.
func (o *TestPointPutModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeleted, true
}

// SetIsDeleted sets field value
func (o *TestPointPutModel) SetIsDeleted(v bool) {
	o.IsDeleted = v
}

func (o TestPointPutModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPointPutModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TesterId.IsSet() {
		toSerialize["testerId"] = o.TesterId.Get()
	}
	toSerialize["iterationId"] = o.IterationId
	if o.WorkItemId.IsSet() {
		toSerialize["workItemId"] = o.WorkItemId.Get()
	}
	if o.ConfigurationId.IsSet() {
		toSerialize["configurationId"] = o.ConfigurationId.Get()
	}
	toSerialize["testSuiteId"] = o.TestSuiteId
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	toSerialize["statusModel"] = o.StatusModel
	if o.LastTestResultId.IsSet() {
		toSerialize["lastTestResultId"] = o.LastTestResultId.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["isDeleted"] = o.IsDeleted
	return toSerialize, nil
}

func (o *TestPointPutModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"iterationId",
		"testSuiteId",
		"statusModel",
		"id",
		"isDeleted",
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

	varTestPointPutModel := _TestPointPutModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestPointPutModel)

	if err != nil {
		return err
	}

	*o = TestPointPutModel(varTestPointPutModel)

	return err
}

type NullableTestPointPutModel struct {
	value *TestPointPutModel
	isSet bool
}

func (v NullableTestPointPutModel) Get() *TestPointPutModel {
	return v.value
}

func (v *NullableTestPointPutModel) Set(val *TestPointPutModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPointPutModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPointPutModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPointPutModel(val *TestPointPutModel) *NullableTestPointPutModel {
	return &NullableTestPointPutModel{value: val, isSet: true}
}

func (v NullableTestPointPutModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPointPutModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


