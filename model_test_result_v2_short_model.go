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
)

// checks if the TestResultV2ShortModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestResultV2ShortModel{}

// TestResultV2ShortModel struct for TestResultV2ShortModel
type TestResultV2ShortModel struct {
	Id string `json:"id"`
	ConfigurationId string `json:"configurationId"`
	WorkItemVersionId string `json:"workItemVersionId"`
	AutoTestId NullableString `json:"autoTestId,omitempty"`
	Message NullableString `json:"message,omitempty"`
	Traces NullableString `json:"traces,omitempty"`
	StartedOn NullableTime `json:"startedOn,omitempty"`
	CompletedOn NullableTime `json:"completedOn,omitempty"`
	RunByUserId NullableString `json:"runByUserId,omitempty"`
	StoppedByUserId NullableString `json:"stoppedByUserId,omitempty"`
	TestPointId NullableString `json:"testPointId,omitempty"`
	TestPoint NullableTestPointRelatedToTestResult `json:"testPoint,omitempty"`
	TestRunId string `json:"testRunId"`
	// Property can contain one of these values: Passed, Failed, InProgress, Blocked, Skipped
	Outcome string `json:"outcome"`
	Comment NullableString `json:"comment,omitempty"`
	Links []LinkModel `json:"links,omitempty"`
	Attachments []AttachmentModel `json:"attachments,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}

// NewTestResultV2ShortModel instantiates a new TestResultV2ShortModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestResultV2ShortModel(id string, configurationId string, workItemVersionId string, testRunId string, outcome string) *TestResultV2ShortModel {
	this := TestResultV2ShortModel{}
	this.Id = id
	this.ConfigurationId = configurationId
	this.WorkItemVersionId = workItemVersionId
	this.TestRunId = testRunId
	this.Outcome = outcome
	return &this
}

// NewTestResultV2ShortModelWithDefaults instantiates a new TestResultV2ShortModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestResultV2ShortModelWithDefaults() *TestResultV2ShortModel {
	this := TestResultV2ShortModel{}
	return &this
}

// GetId returns the Id field value
func (o *TestResultV2ShortModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestResultV2ShortModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestResultV2ShortModel) SetId(v string) {
	o.Id = v
}

// GetConfigurationId returns the ConfigurationId field value
func (o *TestResultV2ShortModel) GetConfigurationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigurationId
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value
// and a boolean to check if the value has been set.
func (o *TestResultV2ShortModel) GetConfigurationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationId, true
}

// SetConfigurationId sets field value
func (o *TestResultV2ShortModel) SetConfigurationId(v string) {
	o.ConfigurationId = v
}

// GetWorkItemVersionId returns the WorkItemVersionId field value
func (o *TestResultV2ShortModel) GetWorkItemVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkItemVersionId
}

// GetWorkItemVersionIdOk returns a tuple with the WorkItemVersionId field value
// and a boolean to check if the value has been set.
func (o *TestResultV2ShortModel) GetWorkItemVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkItemVersionId, true
}

// SetWorkItemVersionId sets field value
func (o *TestResultV2ShortModel) SetWorkItemVersionId(v string) {
	o.WorkItemVersionId = v
}

// GetAutoTestId returns the AutoTestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2ShortModel) GetAutoTestId() string {
	if o == nil || IsNil(o.AutoTestId.Get()) {
		var ret string
		return ret
	}
	return *o.AutoTestId.Get()
}

// GetAutoTestIdOk returns a tuple with the AutoTestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2ShortModel) GetAutoTestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoTestId.Get(), o.AutoTestId.IsSet()
}

// HasAutoTestId returns a boolean if a field has been set.
func (o *TestResultV2ShortModel) HasAutoTestId() bool {
	if o != nil && o.AutoTestId.IsSet() {
		return true
	}

	return false
}

// SetAutoTestId gets a reference to the given NullableString and assigns it to the AutoTestId field.
func (o *TestResultV2ShortModel) SetAutoTestId(v string) {
	o.AutoTestId.Set(&v)
}
// SetAutoTestIdNil sets the value for AutoTestId to be an explicit nil
func (o *TestResultV2ShortModel) SetAutoTestIdNil() {
	o.AutoTestId.Set(nil)
}

// UnsetAutoTestId ensures that no value is present for AutoTestId, not even an explicit nil
func (o *TestResultV2ShortModel) UnsetAutoTestId() {
	o.AutoTestId.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2ShortModel) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2ShortModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *TestResultV2ShortModel) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *TestResultV2ShortModel) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *TestResultV2ShortModel) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *TestResultV2ShortModel) UnsetMessage() {
	o.Message.Unset()
}

// GetTraces returns the Traces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2ShortModel) GetTraces() string {
	if o == nil || IsNil(o.Traces.Get()) {
		var ret string
		return ret
	}
	return *o.Traces.Get()
}

// GetTracesOk returns a tuple with the Traces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2ShortModel) GetTracesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Traces.Get(), o.Traces.IsSet()
}

// HasTraces returns a boolean if a field has been set.
func (o *TestResultV2ShortModel) HasTraces() bool {
	if o != nil && o.Traces.IsSet() {
		return true
	}

	return false
}

// SetTraces gets a reference to the given NullableString and assigns it to the Traces field.
func (o *TestResultV2ShortModel) SetTraces(v string) {
	o.Traces.Set(&v)
}
// SetTracesNil sets the value for Traces to be an explicit nil
func (o *TestResultV2ShortModel) SetTracesNil() {
	o.Traces.Set(nil)
}

// UnsetTraces ensures that no value is present for Traces, not even an explicit nil
func (o *TestResultV2ShortModel) UnsetTraces() {
	o.Traces.Unset()
}

// GetStartedOn returns the StartedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2ShortModel) GetStartedOn() time.Time {
	if o == nil || IsNil(o.StartedOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartedOn.Get()
}

// GetStartedOnOk returns a tuple with the StartedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2ShortModel) GetStartedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedOn.Get(), o.StartedOn.IsSet()
}

// HasStartedOn returns a boolean if a field has been set.
func (o *TestResultV2ShortModel) HasStartedOn() bool {
	if o != nil && o.StartedOn.IsSet() {
		return true
	}

	return false
}

// SetStartedOn gets a reference to the given NullableTime and assigns it to the StartedOn field.
func (o *TestResultV2ShortModel) SetStartedOn(v time.Time) {
	o.StartedOn.Set(&v)
}
// SetStartedOnNil sets the value for StartedOn to be an explicit nil
func (o *TestResultV2ShortModel) SetStartedOnNil() {
	o.StartedOn.Set(nil)
}

// UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
func (o *TestResultV2ShortModel) UnsetStartedOn() {
	o.StartedOn.Unset()
}

// GetCompletedOn returns the CompletedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2ShortModel) GetCompletedOn() time.Time {
	if o == nil || IsNil(o.CompletedOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CompletedOn.Get()
}

// GetCompletedOnOk returns a tuple with the CompletedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2ShortModel) GetCompletedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedOn.Get(), o.CompletedOn.IsSet()
}

// HasCompletedOn returns a boolean if a field has been set.
func (o *TestResultV2ShortModel) HasCompletedOn() bool {
	if o != nil && o.CompletedOn.IsSet() {
		return true
	}

	return false
}

// SetCompletedOn gets a reference to the given NullableTime and assigns it to the CompletedOn field.
func (o *TestResultV2ShortModel) SetCompletedOn(v time.Time) {
	o.CompletedOn.Set(&v)
}
// SetCompletedOnNil sets the value for CompletedOn to be an explicit nil
func (o *TestResultV2ShortModel) SetCompletedOnNil() {
	o.CompletedOn.Set(nil)
}

// UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
func (o *TestResultV2ShortModel) UnsetCompletedOn() {
	o.CompletedOn.Unset()
}

// GetRunByUserId returns the RunByUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2ShortModel) GetRunByUserId() string {
	if o == nil || IsNil(o.RunByUserId.Get()) {
		var ret string
		return ret
	}
	return *o.RunByUserId.Get()
}

// GetRunByUserIdOk returns a tuple with the RunByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2ShortModel) GetRunByUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunByUserId.Get(), o.RunByUserId.IsSet()
}

// HasRunByUserId returns a boolean if a field has been set.
func (o *TestResultV2ShortModel) HasRunByUserId() bool {
	if o != nil && o.RunByUserId.IsSet() {
		return true
	}

	return false
}

// SetRunByUserId gets a reference to the given NullableString and assigns it to the RunByUserId field.
func (o *TestResultV2ShortModel) SetRunByUserId(v string) {
	o.RunByUserId.Set(&v)
}
// SetRunByUserIdNil sets the value for RunByUserId to be an explicit nil
func (o *TestResultV2ShortModel) SetRunByUserIdNil() {
	o.RunByUserId.Set(nil)
}

// UnsetRunByUserId ensures that no value is present for RunByUserId, not even an explicit nil
func (o *TestResultV2ShortModel) UnsetRunByUserId() {
	o.RunByUserId.Unset()
}

// GetStoppedByUserId returns the StoppedByUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2ShortModel) GetStoppedByUserId() string {
	if o == nil || IsNil(o.StoppedByUserId.Get()) {
		var ret string
		return ret
	}
	return *o.StoppedByUserId.Get()
}

// GetStoppedByUserIdOk returns a tuple with the StoppedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2ShortModel) GetStoppedByUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoppedByUserId.Get(), o.StoppedByUserId.IsSet()
}

// HasStoppedByUserId returns a boolean if a field has been set.
func (o *TestResultV2ShortModel) HasStoppedByUserId() bool {
	if o != nil && o.StoppedByUserId.IsSet() {
		return true
	}

	return false
}

// SetStoppedByUserId gets a reference to the given NullableString and assigns it to the StoppedByUserId field.
func (o *TestResultV2ShortModel) SetStoppedByUserId(v string) {
	o.StoppedByUserId.Set(&v)
}
// SetStoppedByUserIdNil sets the value for StoppedByUserId to be an explicit nil
func (o *TestResultV2ShortModel) SetStoppedByUserIdNil() {
	o.StoppedByUserId.Set(nil)
}

// UnsetStoppedByUserId ensures that no value is present for StoppedByUserId, not even an explicit nil
func (o *TestResultV2ShortModel) UnsetStoppedByUserId() {
	o.StoppedByUserId.Unset()
}

// GetTestPointId returns the TestPointId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2ShortModel) GetTestPointId() string {
	if o == nil || IsNil(o.TestPointId.Get()) {
		var ret string
		return ret
	}
	return *o.TestPointId.Get()
}

// GetTestPointIdOk returns a tuple with the TestPointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2ShortModel) GetTestPointIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestPointId.Get(), o.TestPointId.IsSet()
}

// HasTestPointId returns a boolean if a field has been set.
func (o *TestResultV2ShortModel) HasTestPointId() bool {
	if o != nil && o.TestPointId.IsSet() {
		return true
	}

	return false
}

// SetTestPointId gets a reference to the given NullableString and assigns it to the TestPointId field.
func (o *TestResultV2ShortModel) SetTestPointId(v string) {
	o.TestPointId.Set(&v)
}
// SetTestPointIdNil sets the value for TestPointId to be an explicit nil
func (o *TestResultV2ShortModel) SetTestPointIdNil() {
	o.TestPointId.Set(nil)
}

// UnsetTestPointId ensures that no value is present for TestPointId, not even an explicit nil
func (o *TestResultV2ShortModel) UnsetTestPointId() {
	o.TestPointId.Unset()
}

// GetTestPoint returns the TestPoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2ShortModel) GetTestPoint() TestPointRelatedToTestResult {
	if o == nil || IsNil(o.TestPoint.Get()) {
		var ret TestPointRelatedToTestResult
		return ret
	}
	return *o.TestPoint.Get()
}

// GetTestPointOk returns a tuple with the TestPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2ShortModel) GetTestPointOk() (*TestPointRelatedToTestResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestPoint.Get(), o.TestPoint.IsSet()
}

// HasTestPoint returns a boolean if a field has been set.
func (o *TestResultV2ShortModel) HasTestPoint() bool {
	if o != nil && o.TestPoint.IsSet() {
		return true
	}

	return false
}

// SetTestPoint gets a reference to the given NullableTestPointRelatedToTestResult and assigns it to the TestPoint field.
func (o *TestResultV2ShortModel) SetTestPoint(v TestPointRelatedToTestResult) {
	o.TestPoint.Set(&v)
}
// SetTestPointNil sets the value for TestPoint to be an explicit nil
func (o *TestResultV2ShortModel) SetTestPointNil() {
	o.TestPoint.Set(nil)
}

// UnsetTestPoint ensures that no value is present for TestPoint, not even an explicit nil
func (o *TestResultV2ShortModel) UnsetTestPoint() {
	o.TestPoint.Unset()
}

// GetTestRunId returns the TestRunId field value
func (o *TestResultV2ShortModel) GetTestRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestRunId
}

// GetTestRunIdOk returns a tuple with the TestRunId field value
// and a boolean to check if the value has been set.
func (o *TestResultV2ShortModel) GetTestRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestRunId, true
}

// SetTestRunId sets field value
func (o *TestResultV2ShortModel) SetTestRunId(v string) {
	o.TestRunId = v
}

// GetOutcome returns the Outcome field value
func (o *TestResultV2ShortModel) GetOutcome() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value
// and a boolean to check if the value has been set.
func (o *TestResultV2ShortModel) GetOutcomeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outcome, true
}

// SetOutcome sets field value
func (o *TestResultV2ShortModel) SetOutcome(v string) {
	o.Outcome = v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2ShortModel) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2ShortModel) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *TestResultV2ShortModel) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *TestResultV2ShortModel) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *TestResultV2ShortModel) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *TestResultV2ShortModel) UnsetComment() {
	o.Comment.Unset()
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2ShortModel) GetLinks() []LinkModel {
	if o == nil {
		var ret []LinkModel
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2ShortModel) GetLinksOk() ([]LinkModel, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TestResultV2ShortModel) HasLinks() bool {
	if o != nil && IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkModel and assigns it to the Links field.
func (o *TestResultV2ShortModel) SetLinks(v []LinkModel) {
	o.Links = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2ShortModel) GetAttachments() []AttachmentModel {
	if o == nil {
		var ret []AttachmentModel
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2ShortModel) GetAttachmentsOk() ([]AttachmentModel, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *TestResultV2ShortModel) HasAttachments() bool {
	if o != nil && IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []AttachmentModel and assigns it to the Attachments field.
func (o *TestResultV2ShortModel) SetAttachments(v []AttachmentModel) {
	o.Attachments = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2ShortModel) GetParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2ShortModel) GetParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *TestResultV2ShortModel) HasParameters() bool {
	if o != nil && IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *TestResultV2ShortModel) SetParameters(v map[string]string) {
	o.Parameters = v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2ShortModel) GetProperties() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2ShortModel) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *TestResultV2ShortModel) HasProperties() bool {
	if o != nil && IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *TestResultV2ShortModel) SetProperties(v map[string]string) {
	o.Properties = v
}

func (o TestResultV2ShortModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestResultV2ShortModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["configurationId"] = o.ConfigurationId
	toSerialize["workItemVersionId"] = o.WorkItemVersionId
	if o.AutoTestId.IsSet() {
		toSerialize["autoTestId"] = o.AutoTestId.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Traces.IsSet() {
		toSerialize["traces"] = o.Traces.Get()
	}
	if o.StartedOn.IsSet() {
		toSerialize["startedOn"] = o.StartedOn.Get()
	}
	if o.CompletedOn.IsSet() {
		toSerialize["completedOn"] = o.CompletedOn.Get()
	}
	if o.RunByUserId.IsSet() {
		toSerialize["runByUserId"] = o.RunByUserId.Get()
	}
	if o.StoppedByUserId.IsSet() {
		toSerialize["stoppedByUserId"] = o.StoppedByUserId.Get()
	}
	if o.TestPointId.IsSet() {
		toSerialize["testPointId"] = o.TestPointId.Get()
	}
	if o.TestPoint.IsSet() {
		toSerialize["testPoint"] = o.TestPoint.Get()
	}
	toSerialize["testRunId"] = o.TestRunId
	toSerialize["outcome"] = o.Outcome
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableTestResultV2ShortModel struct {
	value *TestResultV2ShortModel
	isSet bool
}

func (v NullableTestResultV2ShortModel) Get() *TestResultV2ShortModel {
	return v.value
}

func (v *NullableTestResultV2ShortModel) Set(val *TestResultV2ShortModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestResultV2ShortModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestResultV2ShortModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestResultV2ShortModel(val *TestResultV2ShortModel) *NullableTestResultV2ShortModel {
	return &NullableTestResultV2ShortModel{value: val, isSet: true}
}

func (v NullableTestResultV2ShortModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestResultV2ShortModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


