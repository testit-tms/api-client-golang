/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the TestResultShortModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestResultShortModel{}

// TestResultShortModel struct for TestResultShortModel
type TestResultShortModel struct {
	Id string `json:"id"`
	Outcome string `json:"outcome"`
	Traces NullableString `json:"traces,omitempty"`
	FailureType string `json:"failureType"`
	Message NullableString `json:"message,omitempty"`
	TestPoint NullableTestPointPutModel `json:"testPoint,omitempty"`
	CreatedDate NullableTime `json:"createdDate,omitempty"`
	AutoTest NullableAutoTestShortModel `json:"autoTest,omitempty"`
	Attachments []AttachmentModel `json:"attachments,omitempty"`
}

type _TestResultShortModel TestResultShortModel

// NewTestResultShortModel instantiates a new TestResultShortModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestResultShortModel(id string, outcome string, failureType string) *TestResultShortModel {
	this := TestResultShortModel{}
	this.Id = id
	this.Outcome = outcome
	this.FailureType = failureType
	return &this
}

// NewTestResultShortModelWithDefaults instantiates a new TestResultShortModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestResultShortModelWithDefaults() *TestResultShortModel {
	this := TestResultShortModel{}
	return &this
}

// GetId returns the Id field value
func (o *TestResultShortModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestResultShortModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestResultShortModel) SetId(v string) {
	o.Id = v
}

// GetOutcome returns the Outcome field value
func (o *TestResultShortModel) GetOutcome() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value
// and a boolean to check if the value has been set.
func (o *TestResultShortModel) GetOutcomeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outcome, true
}

// SetOutcome sets field value
func (o *TestResultShortModel) SetOutcome(v string) {
	o.Outcome = v
}

// GetTraces returns the Traces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultShortModel) GetTraces() string {
	if o == nil || IsNil(o.Traces.Get()) {
		var ret string
		return ret
	}
	return *o.Traces.Get()
}

// GetTracesOk returns a tuple with the Traces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultShortModel) GetTracesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Traces.Get(), o.Traces.IsSet()
}

// HasTraces returns a boolean if a field has been set.
func (o *TestResultShortModel) HasTraces() bool {
	if o != nil && o.Traces.IsSet() {
		return true
	}

	return false
}

// SetTraces gets a reference to the given NullableString and assigns it to the Traces field.
func (o *TestResultShortModel) SetTraces(v string) {
	o.Traces.Set(&v)
}
// SetTracesNil sets the value for Traces to be an explicit nil
func (o *TestResultShortModel) SetTracesNil() {
	o.Traces.Set(nil)
}

// UnsetTraces ensures that no value is present for Traces, not even an explicit nil
func (o *TestResultShortModel) UnsetTraces() {
	o.Traces.Unset()
}

// GetFailureType returns the FailureType field value
func (o *TestResultShortModel) GetFailureType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FailureType
}

// GetFailureTypeOk returns a tuple with the FailureType field value
// and a boolean to check if the value has been set.
func (o *TestResultShortModel) GetFailureTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailureType, true
}

// SetFailureType sets field value
func (o *TestResultShortModel) SetFailureType(v string) {
	o.FailureType = v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultShortModel) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultShortModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *TestResultShortModel) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *TestResultShortModel) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *TestResultShortModel) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *TestResultShortModel) UnsetMessage() {
	o.Message.Unset()
}

// GetTestPoint returns the TestPoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultShortModel) GetTestPoint() TestPointPutModel {
	if o == nil || IsNil(o.TestPoint.Get()) {
		var ret TestPointPutModel
		return ret
	}
	return *o.TestPoint.Get()
}

// GetTestPointOk returns a tuple with the TestPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultShortModel) GetTestPointOk() (*TestPointPutModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestPoint.Get(), o.TestPoint.IsSet()
}

// HasTestPoint returns a boolean if a field has been set.
func (o *TestResultShortModel) HasTestPoint() bool {
	if o != nil && o.TestPoint.IsSet() {
		return true
	}

	return false
}

// SetTestPoint gets a reference to the given NullableTestPointPutModel and assigns it to the TestPoint field.
func (o *TestResultShortModel) SetTestPoint(v TestPointPutModel) {
	o.TestPoint.Set(&v)
}
// SetTestPointNil sets the value for TestPoint to be an explicit nil
func (o *TestResultShortModel) SetTestPointNil() {
	o.TestPoint.Set(nil)
}

// UnsetTestPoint ensures that no value is present for TestPoint, not even an explicit nil
func (o *TestResultShortModel) UnsetTestPoint() {
	o.TestPoint.Unset()
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultShortModel) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate.Get()
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultShortModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDate.Get(), o.CreatedDate.IsSet()
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *TestResultShortModel) HasCreatedDate() bool {
	if o != nil && o.CreatedDate.IsSet() {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given NullableTime and assigns it to the CreatedDate field.
func (o *TestResultShortModel) SetCreatedDate(v time.Time) {
	o.CreatedDate.Set(&v)
}
// SetCreatedDateNil sets the value for CreatedDate to be an explicit nil
func (o *TestResultShortModel) SetCreatedDateNil() {
	o.CreatedDate.Set(nil)
}

// UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
func (o *TestResultShortModel) UnsetCreatedDate() {
	o.CreatedDate.Unset()
}

// GetAutoTest returns the AutoTest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultShortModel) GetAutoTest() AutoTestShortModel {
	if o == nil || IsNil(o.AutoTest.Get()) {
		var ret AutoTestShortModel
		return ret
	}
	return *o.AutoTest.Get()
}

// GetAutoTestOk returns a tuple with the AutoTest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultShortModel) GetAutoTestOk() (*AutoTestShortModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoTest.Get(), o.AutoTest.IsSet()
}

// HasAutoTest returns a boolean if a field has been set.
func (o *TestResultShortModel) HasAutoTest() bool {
	if o != nil && o.AutoTest.IsSet() {
		return true
	}

	return false
}

// SetAutoTest gets a reference to the given NullableAutoTestShortModel and assigns it to the AutoTest field.
func (o *TestResultShortModel) SetAutoTest(v AutoTestShortModel) {
	o.AutoTest.Set(&v)
}
// SetAutoTestNil sets the value for AutoTest to be an explicit nil
func (o *TestResultShortModel) SetAutoTestNil() {
	o.AutoTest.Set(nil)
}

// UnsetAutoTest ensures that no value is present for AutoTest, not even an explicit nil
func (o *TestResultShortModel) UnsetAutoTest() {
	o.AutoTest.Unset()
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultShortModel) GetAttachments() []AttachmentModel {
	if o == nil {
		var ret []AttachmentModel
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultShortModel) GetAttachmentsOk() ([]AttachmentModel, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *TestResultShortModel) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []AttachmentModel and assigns it to the Attachments field.
func (o *TestResultShortModel) SetAttachments(v []AttachmentModel) {
	o.Attachments = v
}

func (o TestResultShortModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestResultShortModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["outcome"] = o.Outcome
	if o.Traces.IsSet() {
		toSerialize["traces"] = o.Traces.Get()
	}
	toSerialize["failureType"] = o.FailureType
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.TestPoint.IsSet() {
		toSerialize["testPoint"] = o.TestPoint.Get()
	}
	if o.CreatedDate.IsSet() {
		toSerialize["createdDate"] = o.CreatedDate.Get()
	}
	if o.AutoTest.IsSet() {
		toSerialize["autoTest"] = o.AutoTest.Get()
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

func (o *TestResultShortModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"outcome",
		"failureType",
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

	varTestResultShortModel := _TestResultShortModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestResultShortModel)

	if err != nil {
		return err
	}

	*o = TestResultShortModel(varTestResultShortModel)

	return err
}

type NullableTestResultShortModel struct {
	value *TestResultShortModel
	isSet bool
}

func (v NullableTestResultShortModel) Get() *TestResultShortModel {
	return v.value
}

func (v *NullableTestResultShortModel) Set(val *TestResultShortModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestResultShortModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestResultShortModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestResultShortModel(val *TestResultShortModel) *NullableTestResultShortModel {
	return &NullableTestResultShortModel{value: val, isSet: true}
}

func (v NullableTestResultShortModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestResultShortModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


