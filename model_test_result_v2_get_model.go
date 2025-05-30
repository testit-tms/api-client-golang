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

// checks if the TestResultV2GetModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestResultV2GetModel{}

// TestResultV2GetModel struct for TestResultV2GetModel
type TestResultV2GetModel struct {
	Configuration NullableConfigurationModel `json:"configuration,omitempty"`
	AutoTest NullableAutoTestModelV2GetModel `json:"autoTest,omitempty"`
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
	TestPoint NullableTestPointShortModel `json:"testPoint,omitempty"`
	TestRunId string `json:"testRunId"`
	// Property can contain one of these values: Passed, Failed, InProgress, Blocked, Skipped
	Outcome string `json:"outcome"`
	Comment NullableString `json:"comment,omitempty"`
	Links []LinkModel `json:"links,omitempty"`
	Attachments []AttachmentModel `json:"attachments,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}

type _TestResultV2GetModel TestResultV2GetModel

// NewTestResultV2GetModel instantiates a new TestResultV2GetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestResultV2GetModel(id string, configurationId string, workItemVersionId string, testRunId string, outcome string) *TestResultV2GetModel {
	this := TestResultV2GetModel{}
	this.Id = id
	this.ConfigurationId = configurationId
	this.WorkItemVersionId = workItemVersionId
	this.TestRunId = testRunId
	this.Outcome = outcome
	return &this
}

// NewTestResultV2GetModelWithDefaults instantiates a new TestResultV2GetModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestResultV2GetModelWithDefaults() *TestResultV2GetModel {
	this := TestResultV2GetModel{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2GetModel) GetConfiguration() ConfigurationModel {
	if o == nil || IsNil(o.Configuration.Get()) {
		var ret ConfigurationModel
		return ret
	}
	return *o.Configuration.Get()
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2GetModel) GetConfigurationOk() (*ConfigurationModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Configuration.Get(), o.Configuration.IsSet()
}

// HasConfiguration returns a boolean if a field has been set.
func (o *TestResultV2GetModel) HasConfiguration() bool {
	if o != nil && o.Configuration.IsSet() {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given NullableConfigurationModel and assigns it to the Configuration field.
func (o *TestResultV2GetModel) SetConfiguration(v ConfigurationModel) {
	o.Configuration.Set(&v)
}
// SetConfigurationNil sets the value for Configuration to be an explicit nil
func (o *TestResultV2GetModel) SetConfigurationNil() {
	o.Configuration.Set(nil)
}

// UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
func (o *TestResultV2GetModel) UnsetConfiguration() {
	o.Configuration.Unset()
}

// GetAutoTest returns the AutoTest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2GetModel) GetAutoTest() AutoTestModelV2GetModel {
	if o == nil || IsNil(o.AutoTest.Get()) {
		var ret AutoTestModelV2GetModel
		return ret
	}
	return *o.AutoTest.Get()
}

// GetAutoTestOk returns a tuple with the AutoTest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2GetModel) GetAutoTestOk() (*AutoTestModelV2GetModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoTest.Get(), o.AutoTest.IsSet()
}

// HasAutoTest returns a boolean if a field has been set.
func (o *TestResultV2GetModel) HasAutoTest() bool {
	if o != nil && o.AutoTest.IsSet() {
		return true
	}

	return false
}

// SetAutoTest gets a reference to the given NullableAutoTestModelV2GetModel and assigns it to the AutoTest field.
func (o *TestResultV2GetModel) SetAutoTest(v AutoTestModelV2GetModel) {
	o.AutoTest.Set(&v)
}
// SetAutoTestNil sets the value for AutoTest to be an explicit nil
func (o *TestResultV2GetModel) SetAutoTestNil() {
	o.AutoTest.Set(nil)
}

// UnsetAutoTest ensures that no value is present for AutoTest, not even an explicit nil
func (o *TestResultV2GetModel) UnsetAutoTest() {
	o.AutoTest.Unset()
}

// GetId returns the Id field value
func (o *TestResultV2GetModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestResultV2GetModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestResultV2GetModel) SetId(v string) {
	o.Id = v
}

// GetConfigurationId returns the ConfigurationId field value
func (o *TestResultV2GetModel) GetConfigurationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigurationId
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value
// and a boolean to check if the value has been set.
func (o *TestResultV2GetModel) GetConfigurationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationId, true
}

// SetConfigurationId sets field value
func (o *TestResultV2GetModel) SetConfigurationId(v string) {
	o.ConfigurationId = v
}

// GetWorkItemVersionId returns the WorkItemVersionId field value
func (o *TestResultV2GetModel) GetWorkItemVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkItemVersionId
}

// GetWorkItemVersionIdOk returns a tuple with the WorkItemVersionId field value
// and a boolean to check if the value has been set.
func (o *TestResultV2GetModel) GetWorkItemVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkItemVersionId, true
}

// SetWorkItemVersionId sets field value
func (o *TestResultV2GetModel) SetWorkItemVersionId(v string) {
	o.WorkItemVersionId = v
}

// GetAutoTestId returns the AutoTestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2GetModel) GetAutoTestId() string {
	if o == nil || IsNil(o.AutoTestId.Get()) {
		var ret string
		return ret
	}
	return *o.AutoTestId.Get()
}

// GetAutoTestIdOk returns a tuple with the AutoTestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2GetModel) GetAutoTestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoTestId.Get(), o.AutoTestId.IsSet()
}

// HasAutoTestId returns a boolean if a field has been set.
func (o *TestResultV2GetModel) HasAutoTestId() bool {
	if o != nil && o.AutoTestId.IsSet() {
		return true
	}

	return false
}

// SetAutoTestId gets a reference to the given NullableString and assigns it to the AutoTestId field.
func (o *TestResultV2GetModel) SetAutoTestId(v string) {
	o.AutoTestId.Set(&v)
}
// SetAutoTestIdNil sets the value for AutoTestId to be an explicit nil
func (o *TestResultV2GetModel) SetAutoTestIdNil() {
	o.AutoTestId.Set(nil)
}

// UnsetAutoTestId ensures that no value is present for AutoTestId, not even an explicit nil
func (o *TestResultV2GetModel) UnsetAutoTestId() {
	o.AutoTestId.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2GetModel) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2GetModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *TestResultV2GetModel) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *TestResultV2GetModel) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *TestResultV2GetModel) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *TestResultV2GetModel) UnsetMessage() {
	o.Message.Unset()
}

// GetTraces returns the Traces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2GetModel) GetTraces() string {
	if o == nil || IsNil(o.Traces.Get()) {
		var ret string
		return ret
	}
	return *o.Traces.Get()
}

// GetTracesOk returns a tuple with the Traces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2GetModel) GetTracesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Traces.Get(), o.Traces.IsSet()
}

// HasTraces returns a boolean if a field has been set.
func (o *TestResultV2GetModel) HasTraces() bool {
	if o != nil && o.Traces.IsSet() {
		return true
	}

	return false
}

// SetTraces gets a reference to the given NullableString and assigns it to the Traces field.
func (o *TestResultV2GetModel) SetTraces(v string) {
	o.Traces.Set(&v)
}
// SetTracesNil sets the value for Traces to be an explicit nil
func (o *TestResultV2GetModel) SetTracesNil() {
	o.Traces.Set(nil)
}

// UnsetTraces ensures that no value is present for Traces, not even an explicit nil
func (o *TestResultV2GetModel) UnsetTraces() {
	o.Traces.Unset()
}

// GetStartedOn returns the StartedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2GetModel) GetStartedOn() time.Time {
	if o == nil || IsNil(o.StartedOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartedOn.Get()
}

// GetStartedOnOk returns a tuple with the StartedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2GetModel) GetStartedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedOn.Get(), o.StartedOn.IsSet()
}

// HasStartedOn returns a boolean if a field has been set.
func (o *TestResultV2GetModel) HasStartedOn() bool {
	if o != nil && o.StartedOn.IsSet() {
		return true
	}

	return false
}

// SetStartedOn gets a reference to the given NullableTime and assigns it to the StartedOn field.
func (o *TestResultV2GetModel) SetStartedOn(v time.Time) {
	o.StartedOn.Set(&v)
}
// SetStartedOnNil sets the value for StartedOn to be an explicit nil
func (o *TestResultV2GetModel) SetStartedOnNil() {
	o.StartedOn.Set(nil)
}

// UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
func (o *TestResultV2GetModel) UnsetStartedOn() {
	o.StartedOn.Unset()
}

// GetCompletedOn returns the CompletedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2GetModel) GetCompletedOn() time.Time {
	if o == nil || IsNil(o.CompletedOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CompletedOn.Get()
}

// GetCompletedOnOk returns a tuple with the CompletedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2GetModel) GetCompletedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedOn.Get(), o.CompletedOn.IsSet()
}

// HasCompletedOn returns a boolean if a field has been set.
func (o *TestResultV2GetModel) HasCompletedOn() bool {
	if o != nil && o.CompletedOn.IsSet() {
		return true
	}

	return false
}

// SetCompletedOn gets a reference to the given NullableTime and assigns it to the CompletedOn field.
func (o *TestResultV2GetModel) SetCompletedOn(v time.Time) {
	o.CompletedOn.Set(&v)
}
// SetCompletedOnNil sets the value for CompletedOn to be an explicit nil
func (o *TestResultV2GetModel) SetCompletedOnNil() {
	o.CompletedOn.Set(nil)
}

// UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
func (o *TestResultV2GetModel) UnsetCompletedOn() {
	o.CompletedOn.Unset()
}

// GetRunByUserId returns the RunByUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2GetModel) GetRunByUserId() string {
	if o == nil || IsNil(o.RunByUserId.Get()) {
		var ret string
		return ret
	}
	return *o.RunByUserId.Get()
}

// GetRunByUserIdOk returns a tuple with the RunByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2GetModel) GetRunByUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunByUserId.Get(), o.RunByUserId.IsSet()
}

// HasRunByUserId returns a boolean if a field has been set.
func (o *TestResultV2GetModel) HasRunByUserId() bool {
	if o != nil && o.RunByUserId.IsSet() {
		return true
	}

	return false
}

// SetRunByUserId gets a reference to the given NullableString and assigns it to the RunByUserId field.
func (o *TestResultV2GetModel) SetRunByUserId(v string) {
	o.RunByUserId.Set(&v)
}
// SetRunByUserIdNil sets the value for RunByUserId to be an explicit nil
func (o *TestResultV2GetModel) SetRunByUserIdNil() {
	o.RunByUserId.Set(nil)
}

// UnsetRunByUserId ensures that no value is present for RunByUserId, not even an explicit nil
func (o *TestResultV2GetModel) UnsetRunByUserId() {
	o.RunByUserId.Unset()
}

// GetStoppedByUserId returns the StoppedByUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2GetModel) GetStoppedByUserId() string {
	if o == nil || IsNil(o.StoppedByUserId.Get()) {
		var ret string
		return ret
	}
	return *o.StoppedByUserId.Get()
}

// GetStoppedByUserIdOk returns a tuple with the StoppedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2GetModel) GetStoppedByUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoppedByUserId.Get(), o.StoppedByUserId.IsSet()
}

// HasStoppedByUserId returns a boolean if a field has been set.
func (o *TestResultV2GetModel) HasStoppedByUserId() bool {
	if o != nil && o.StoppedByUserId.IsSet() {
		return true
	}

	return false
}

// SetStoppedByUserId gets a reference to the given NullableString and assigns it to the StoppedByUserId field.
func (o *TestResultV2GetModel) SetStoppedByUserId(v string) {
	o.StoppedByUserId.Set(&v)
}
// SetStoppedByUserIdNil sets the value for StoppedByUserId to be an explicit nil
func (o *TestResultV2GetModel) SetStoppedByUserIdNil() {
	o.StoppedByUserId.Set(nil)
}

// UnsetStoppedByUserId ensures that no value is present for StoppedByUserId, not even an explicit nil
func (o *TestResultV2GetModel) UnsetStoppedByUserId() {
	o.StoppedByUserId.Unset()
}

// GetTestPointId returns the TestPointId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2GetModel) GetTestPointId() string {
	if o == nil || IsNil(o.TestPointId.Get()) {
		var ret string
		return ret
	}
	return *o.TestPointId.Get()
}

// GetTestPointIdOk returns a tuple with the TestPointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2GetModel) GetTestPointIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestPointId.Get(), o.TestPointId.IsSet()
}

// HasTestPointId returns a boolean if a field has been set.
func (o *TestResultV2GetModel) HasTestPointId() bool {
	if o != nil && o.TestPointId.IsSet() {
		return true
	}

	return false
}

// SetTestPointId gets a reference to the given NullableString and assigns it to the TestPointId field.
func (o *TestResultV2GetModel) SetTestPointId(v string) {
	o.TestPointId.Set(&v)
}
// SetTestPointIdNil sets the value for TestPointId to be an explicit nil
func (o *TestResultV2GetModel) SetTestPointIdNil() {
	o.TestPointId.Set(nil)
}

// UnsetTestPointId ensures that no value is present for TestPointId, not even an explicit nil
func (o *TestResultV2GetModel) UnsetTestPointId() {
	o.TestPointId.Unset()
}

// GetTestPoint returns the TestPoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2GetModel) GetTestPoint() TestPointShortModel {
	if o == nil || IsNil(o.TestPoint.Get()) {
		var ret TestPointShortModel
		return ret
	}
	return *o.TestPoint.Get()
}

// GetTestPointOk returns a tuple with the TestPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2GetModel) GetTestPointOk() (*TestPointShortModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestPoint.Get(), o.TestPoint.IsSet()
}

// HasTestPoint returns a boolean if a field has been set.
func (o *TestResultV2GetModel) HasTestPoint() bool {
	if o != nil && o.TestPoint.IsSet() {
		return true
	}

	return false
}

// SetTestPoint gets a reference to the given NullableTestPointShortModel and assigns it to the TestPoint field.
func (o *TestResultV2GetModel) SetTestPoint(v TestPointShortModel) {
	o.TestPoint.Set(&v)
}
// SetTestPointNil sets the value for TestPoint to be an explicit nil
func (o *TestResultV2GetModel) SetTestPointNil() {
	o.TestPoint.Set(nil)
}

// UnsetTestPoint ensures that no value is present for TestPoint, not even an explicit nil
func (o *TestResultV2GetModel) UnsetTestPoint() {
	o.TestPoint.Unset()
}

// GetTestRunId returns the TestRunId field value
func (o *TestResultV2GetModel) GetTestRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestRunId
}

// GetTestRunIdOk returns a tuple with the TestRunId field value
// and a boolean to check if the value has been set.
func (o *TestResultV2GetModel) GetTestRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestRunId, true
}

// SetTestRunId sets field value
func (o *TestResultV2GetModel) SetTestRunId(v string) {
	o.TestRunId = v
}

// GetOutcome returns the Outcome field value
func (o *TestResultV2GetModel) GetOutcome() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value
// and a boolean to check if the value has been set.
func (o *TestResultV2GetModel) GetOutcomeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outcome, true
}

// SetOutcome sets field value
func (o *TestResultV2GetModel) SetOutcome(v string) {
	o.Outcome = v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2GetModel) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2GetModel) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *TestResultV2GetModel) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *TestResultV2GetModel) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *TestResultV2GetModel) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *TestResultV2GetModel) UnsetComment() {
	o.Comment.Unset()
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2GetModel) GetLinks() []LinkModel {
	if o == nil {
		var ret []LinkModel
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2GetModel) GetLinksOk() ([]LinkModel, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TestResultV2GetModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkModel and assigns it to the Links field.
func (o *TestResultV2GetModel) SetLinks(v []LinkModel) {
	o.Links = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2GetModel) GetAttachments() []AttachmentModel {
	if o == nil {
		var ret []AttachmentModel
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2GetModel) GetAttachmentsOk() ([]AttachmentModel, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *TestResultV2GetModel) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []AttachmentModel and assigns it to the Attachments field.
func (o *TestResultV2GetModel) SetAttachments(v []AttachmentModel) {
	o.Attachments = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2GetModel) GetParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2GetModel) GetParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *TestResultV2GetModel) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *TestResultV2GetModel) SetParameters(v map[string]string) {
	o.Parameters = v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultV2GetModel) GetProperties() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultV2GetModel) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *TestResultV2GetModel) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *TestResultV2GetModel) SetProperties(v map[string]string) {
	o.Properties = v
}

func (o TestResultV2GetModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestResultV2GetModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Configuration.IsSet() {
		toSerialize["configuration"] = o.Configuration.Get()
	}
	if o.AutoTest.IsSet() {
		toSerialize["autoTest"] = o.AutoTest.Get()
	}
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

func (o *TestResultV2GetModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"configurationId",
		"workItemVersionId",
		"testRunId",
		"outcome",
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

	varTestResultV2GetModel := _TestResultV2GetModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestResultV2GetModel)

	if err != nil {
		return err
	}

	*o = TestResultV2GetModel(varTestResultV2GetModel)

	return err
}

type NullableTestResultV2GetModel struct {
	value *TestResultV2GetModel
	isSet bool
}

func (v NullableTestResultV2GetModel) Get() *TestResultV2GetModel {
	return v.value
}

func (v *NullableTestResultV2GetModel) Set(val *TestResultV2GetModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestResultV2GetModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestResultV2GetModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestResultV2GetModel(val *TestResultV2GetModel) *NullableTestResultV2GetModel {
	return &NullableTestResultV2GetModel{value: val, isSet: true}
}

func (v NullableTestResultV2GetModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestResultV2GetModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


