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

// checks if the TestResultHistoryReportModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestResultHistoryReportModel{}

// TestResultHistoryReportModel struct for TestResultHistoryReportModel
type TestResultHistoryReportModel struct {
	Id string `json:"id"`
	CreatedDate time.Time `json:"createdDate"`
	ModifiedDate time.Time `json:"modifiedDate"`
	// If test run was stopped, this property equals identifier of user who stopped it.Otherwise, the property equals identifier of user who created the test result
	UserId string `json:"userId"`
	TestRunId NullableString `json:"testRunId,omitempty"`
	TestRunName NullableString `json:"testRunName,omitempty"`
	CreatedByUserName NullableString `json:"createdByUserName,omitempty"`
	TestPlanId NullableString `json:"testPlanId,omitempty"`
	TestPlanGlobalId NullableInt64 `json:"testPlanGlobalId,omitempty"`
	TestPlanName NullableString `json:"testPlanName,omitempty"`
	// If test point related to the test result has configuration, this property will be equal to the test point configuration name. Otherwise, this property will be equal to the test result configuration name
	ConfigurationName NullableString `json:"configurationName,omitempty"`
	IsAutomated bool `json:"isAutomated"`
	// If any test result related to the test run is linked with autotest and the run has an outcome, the outcome value equalsto the worst outcome of the last modified test result.Otherwise, the outcome equals to the outcome of first created test result in the test run
	Outcome NullableString `json:"outcome,omitempty"`
	// If any test result related to the test run is linked with autotest, comment will have default valueOtherwise, the comment equals to the comment of first created test result in the test run
	Comment NullableString `json:"comment,omitempty"`
	// If any test result related to the test run is linked with autotest, link will be equal to the links of last modified test result.Otherwise, the links equals to the links of first created test result in the test run
	Links []LinkModel `json:"links,omitempty"`
	StartedOn NullableTime `json:"startedOn,omitempty"`
	CompletedOn NullableTime `json:"completedOn,omitempty"`
	Duration NullableInt64 `json:"duration,omitempty"`
	CreatedById string `json:"createdById"`
	ModifiedById NullableString `json:"modifiedById,omitempty"`
	// If any test result related to the test run is linked with autotest, attachments will be equal to the attachments of last modified test result.Otherwise, the attachments equals to the attachments of first created test result in the test run
	Attachments []AttachmentModel `json:"attachments,omitempty"`
	WorkItemVersionId NullableString `json:"workItemVersionId,omitempty"`
	WorkItemVersionNumber NullableInt32 `json:"workItemVersionNumber,omitempty"`
	LaunchSource NullableString `json:"launchSource,omitempty"`
	FailureClassIds []string `json:"failureClassIds"`
	Parameters map[string]string `json:"parameters,omitempty"`
}

type _TestResultHistoryReportModel TestResultHistoryReportModel

// NewTestResultHistoryReportModel instantiates a new TestResultHistoryReportModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestResultHistoryReportModel(id string, createdDate time.Time, modifiedDate time.Time, userId string, isAutomated bool, createdById string, failureClassIds []string) *TestResultHistoryReportModel {
	this := TestResultHistoryReportModel{}
	this.Id = id
	this.CreatedDate = createdDate
	this.ModifiedDate = modifiedDate
	this.UserId = userId
	this.IsAutomated = isAutomated
	this.CreatedById = createdById
	this.FailureClassIds = failureClassIds
	return &this
}

// NewTestResultHistoryReportModelWithDefaults instantiates a new TestResultHistoryReportModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestResultHistoryReportModelWithDefaults() *TestResultHistoryReportModel {
	this := TestResultHistoryReportModel{}
	return &this
}

// GetId returns the Id field value
func (o *TestResultHistoryReportModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestResultHistoryReportModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestResultHistoryReportModel) SetId(v string) {
	o.Id = v
}

// GetCreatedDate returns the CreatedDate field value
func (o *TestResultHistoryReportModel) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *TestResultHistoryReportModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *TestResultHistoryReportModel) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetModifiedDate returns the ModifiedDate field value
func (o *TestResultHistoryReportModel) GetModifiedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value
// and a boolean to check if the value has been set.
func (o *TestResultHistoryReportModel) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedDate, true
}

// SetModifiedDate sets field value
func (o *TestResultHistoryReportModel) SetModifiedDate(v time.Time) {
	o.ModifiedDate = v
}

// GetUserId returns the UserId field value
func (o *TestResultHistoryReportModel) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *TestResultHistoryReportModel) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *TestResultHistoryReportModel) SetUserId(v string) {
	o.UserId = v
}

// GetTestRunId returns the TestRunId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultHistoryReportModel) GetTestRunId() string {
	if o == nil || IsNil(o.TestRunId.Get()) {
		var ret string
		return ret
	}
	return *o.TestRunId.Get()
}

// GetTestRunIdOk returns a tuple with the TestRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultHistoryReportModel) GetTestRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestRunId.Get(), o.TestRunId.IsSet()
}

// HasTestRunId returns a boolean if a field has been set.
func (o *TestResultHistoryReportModel) HasTestRunId() bool {
	if o != nil && o.TestRunId.IsSet() {
		return true
	}

	return false
}

// SetTestRunId gets a reference to the given NullableString and assigns it to the TestRunId field.
func (o *TestResultHistoryReportModel) SetTestRunId(v string) {
	o.TestRunId.Set(&v)
}
// SetTestRunIdNil sets the value for TestRunId to be an explicit nil
func (o *TestResultHistoryReportModel) SetTestRunIdNil() {
	o.TestRunId.Set(nil)
}

// UnsetTestRunId ensures that no value is present for TestRunId, not even an explicit nil
func (o *TestResultHistoryReportModel) UnsetTestRunId() {
	o.TestRunId.Unset()
}

// GetTestRunName returns the TestRunName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultHistoryReportModel) GetTestRunName() string {
	if o == nil || IsNil(o.TestRunName.Get()) {
		var ret string
		return ret
	}
	return *o.TestRunName.Get()
}

// GetTestRunNameOk returns a tuple with the TestRunName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultHistoryReportModel) GetTestRunNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestRunName.Get(), o.TestRunName.IsSet()
}

// HasTestRunName returns a boolean if a field has been set.
func (o *TestResultHistoryReportModel) HasTestRunName() bool {
	if o != nil && o.TestRunName.IsSet() {
		return true
	}

	return false
}

// SetTestRunName gets a reference to the given NullableString and assigns it to the TestRunName field.
func (o *TestResultHistoryReportModel) SetTestRunName(v string) {
	o.TestRunName.Set(&v)
}
// SetTestRunNameNil sets the value for TestRunName to be an explicit nil
func (o *TestResultHistoryReportModel) SetTestRunNameNil() {
	o.TestRunName.Set(nil)
}

// UnsetTestRunName ensures that no value is present for TestRunName, not even an explicit nil
func (o *TestResultHistoryReportModel) UnsetTestRunName() {
	o.TestRunName.Unset()
}

// GetCreatedByUserName returns the CreatedByUserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultHistoryReportModel) GetCreatedByUserName() string {
	if o == nil || IsNil(o.CreatedByUserName.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedByUserName.Get()
}

// GetCreatedByUserNameOk returns a tuple with the CreatedByUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultHistoryReportModel) GetCreatedByUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedByUserName.Get(), o.CreatedByUserName.IsSet()
}

// HasCreatedByUserName returns a boolean if a field has been set.
func (o *TestResultHistoryReportModel) HasCreatedByUserName() bool {
	if o != nil && o.CreatedByUserName.IsSet() {
		return true
	}

	return false
}

// SetCreatedByUserName gets a reference to the given NullableString and assigns it to the CreatedByUserName field.
func (o *TestResultHistoryReportModel) SetCreatedByUserName(v string) {
	o.CreatedByUserName.Set(&v)
}
// SetCreatedByUserNameNil sets the value for CreatedByUserName to be an explicit nil
func (o *TestResultHistoryReportModel) SetCreatedByUserNameNil() {
	o.CreatedByUserName.Set(nil)
}

// UnsetCreatedByUserName ensures that no value is present for CreatedByUserName, not even an explicit nil
func (o *TestResultHistoryReportModel) UnsetCreatedByUserName() {
	o.CreatedByUserName.Unset()
}

// GetTestPlanId returns the TestPlanId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultHistoryReportModel) GetTestPlanId() string {
	if o == nil || IsNil(o.TestPlanId.Get()) {
		var ret string
		return ret
	}
	return *o.TestPlanId.Get()
}

// GetTestPlanIdOk returns a tuple with the TestPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultHistoryReportModel) GetTestPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestPlanId.Get(), o.TestPlanId.IsSet()
}

// HasTestPlanId returns a boolean if a field has been set.
func (o *TestResultHistoryReportModel) HasTestPlanId() bool {
	if o != nil && o.TestPlanId.IsSet() {
		return true
	}

	return false
}

// SetTestPlanId gets a reference to the given NullableString and assigns it to the TestPlanId field.
func (o *TestResultHistoryReportModel) SetTestPlanId(v string) {
	o.TestPlanId.Set(&v)
}
// SetTestPlanIdNil sets the value for TestPlanId to be an explicit nil
func (o *TestResultHistoryReportModel) SetTestPlanIdNil() {
	o.TestPlanId.Set(nil)
}

// UnsetTestPlanId ensures that no value is present for TestPlanId, not even an explicit nil
func (o *TestResultHistoryReportModel) UnsetTestPlanId() {
	o.TestPlanId.Unset()
}

// GetTestPlanGlobalId returns the TestPlanGlobalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultHistoryReportModel) GetTestPlanGlobalId() int64 {
	if o == nil || IsNil(o.TestPlanGlobalId.Get()) {
		var ret int64
		return ret
	}
	return *o.TestPlanGlobalId.Get()
}

// GetTestPlanGlobalIdOk returns a tuple with the TestPlanGlobalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultHistoryReportModel) GetTestPlanGlobalIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestPlanGlobalId.Get(), o.TestPlanGlobalId.IsSet()
}

// HasTestPlanGlobalId returns a boolean if a field has been set.
func (o *TestResultHistoryReportModel) HasTestPlanGlobalId() bool {
	if o != nil && o.TestPlanGlobalId.IsSet() {
		return true
	}

	return false
}

// SetTestPlanGlobalId gets a reference to the given NullableInt64 and assigns it to the TestPlanGlobalId field.
func (o *TestResultHistoryReportModel) SetTestPlanGlobalId(v int64) {
	o.TestPlanGlobalId.Set(&v)
}
// SetTestPlanGlobalIdNil sets the value for TestPlanGlobalId to be an explicit nil
func (o *TestResultHistoryReportModel) SetTestPlanGlobalIdNil() {
	o.TestPlanGlobalId.Set(nil)
}

// UnsetTestPlanGlobalId ensures that no value is present for TestPlanGlobalId, not even an explicit nil
func (o *TestResultHistoryReportModel) UnsetTestPlanGlobalId() {
	o.TestPlanGlobalId.Unset()
}

// GetTestPlanName returns the TestPlanName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultHistoryReportModel) GetTestPlanName() string {
	if o == nil || IsNil(o.TestPlanName.Get()) {
		var ret string
		return ret
	}
	return *o.TestPlanName.Get()
}

// GetTestPlanNameOk returns a tuple with the TestPlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultHistoryReportModel) GetTestPlanNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestPlanName.Get(), o.TestPlanName.IsSet()
}

// HasTestPlanName returns a boolean if a field has been set.
func (o *TestResultHistoryReportModel) HasTestPlanName() bool {
	if o != nil && o.TestPlanName.IsSet() {
		return true
	}

	return false
}

// SetTestPlanName gets a reference to the given NullableString and assigns it to the TestPlanName field.
func (o *TestResultHistoryReportModel) SetTestPlanName(v string) {
	o.TestPlanName.Set(&v)
}
// SetTestPlanNameNil sets the value for TestPlanName to be an explicit nil
func (o *TestResultHistoryReportModel) SetTestPlanNameNil() {
	o.TestPlanName.Set(nil)
}

// UnsetTestPlanName ensures that no value is present for TestPlanName, not even an explicit nil
func (o *TestResultHistoryReportModel) UnsetTestPlanName() {
	o.TestPlanName.Unset()
}

// GetConfigurationName returns the ConfigurationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultHistoryReportModel) GetConfigurationName() string {
	if o == nil || IsNil(o.ConfigurationName.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigurationName.Get()
}

// GetConfigurationNameOk returns a tuple with the ConfigurationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultHistoryReportModel) GetConfigurationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigurationName.Get(), o.ConfigurationName.IsSet()
}

// HasConfigurationName returns a boolean if a field has been set.
func (o *TestResultHistoryReportModel) HasConfigurationName() bool {
	if o != nil && o.ConfigurationName.IsSet() {
		return true
	}

	return false
}

// SetConfigurationName gets a reference to the given NullableString and assigns it to the ConfigurationName field.
func (o *TestResultHistoryReportModel) SetConfigurationName(v string) {
	o.ConfigurationName.Set(&v)
}
// SetConfigurationNameNil sets the value for ConfigurationName to be an explicit nil
func (o *TestResultHistoryReportModel) SetConfigurationNameNil() {
	o.ConfigurationName.Set(nil)
}

// UnsetConfigurationName ensures that no value is present for ConfigurationName, not even an explicit nil
func (o *TestResultHistoryReportModel) UnsetConfigurationName() {
	o.ConfigurationName.Unset()
}

// GetIsAutomated returns the IsAutomated field value
func (o *TestResultHistoryReportModel) GetIsAutomated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAutomated
}

// GetIsAutomatedOk returns a tuple with the IsAutomated field value
// and a boolean to check if the value has been set.
func (o *TestResultHistoryReportModel) GetIsAutomatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAutomated, true
}

// SetIsAutomated sets field value
func (o *TestResultHistoryReportModel) SetIsAutomated(v bool) {
	o.IsAutomated = v
}

// GetOutcome returns the Outcome field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultHistoryReportModel) GetOutcome() string {
	if o == nil || IsNil(o.Outcome.Get()) {
		var ret string
		return ret
	}
	return *o.Outcome.Get()
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultHistoryReportModel) GetOutcomeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Outcome.Get(), o.Outcome.IsSet()
}

// HasOutcome returns a boolean if a field has been set.
func (o *TestResultHistoryReportModel) HasOutcome() bool {
	if o != nil && o.Outcome.IsSet() {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given NullableString and assigns it to the Outcome field.
func (o *TestResultHistoryReportModel) SetOutcome(v string) {
	o.Outcome.Set(&v)
}
// SetOutcomeNil sets the value for Outcome to be an explicit nil
func (o *TestResultHistoryReportModel) SetOutcomeNil() {
	o.Outcome.Set(nil)
}

// UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
func (o *TestResultHistoryReportModel) UnsetOutcome() {
	o.Outcome.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultHistoryReportModel) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultHistoryReportModel) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *TestResultHistoryReportModel) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *TestResultHistoryReportModel) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *TestResultHistoryReportModel) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *TestResultHistoryReportModel) UnsetComment() {
	o.Comment.Unset()
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultHistoryReportModel) GetLinks() []LinkModel {
	if o == nil {
		var ret []LinkModel
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultHistoryReportModel) GetLinksOk() ([]LinkModel, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TestResultHistoryReportModel) HasLinks() bool {
	if o != nil && IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkModel and assigns it to the Links field.
func (o *TestResultHistoryReportModel) SetLinks(v []LinkModel) {
	o.Links = v
}

// GetStartedOn returns the StartedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultHistoryReportModel) GetStartedOn() time.Time {
	if o == nil || IsNil(o.StartedOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartedOn.Get()
}

// GetStartedOnOk returns a tuple with the StartedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultHistoryReportModel) GetStartedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedOn.Get(), o.StartedOn.IsSet()
}

// HasStartedOn returns a boolean if a field has been set.
func (o *TestResultHistoryReportModel) HasStartedOn() bool {
	if o != nil && o.StartedOn.IsSet() {
		return true
	}

	return false
}

// SetStartedOn gets a reference to the given NullableTime and assigns it to the StartedOn field.
func (o *TestResultHistoryReportModel) SetStartedOn(v time.Time) {
	o.StartedOn.Set(&v)
}
// SetStartedOnNil sets the value for StartedOn to be an explicit nil
func (o *TestResultHistoryReportModel) SetStartedOnNil() {
	o.StartedOn.Set(nil)
}

// UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
func (o *TestResultHistoryReportModel) UnsetStartedOn() {
	o.StartedOn.Unset()
}

// GetCompletedOn returns the CompletedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultHistoryReportModel) GetCompletedOn() time.Time {
	if o == nil || IsNil(o.CompletedOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CompletedOn.Get()
}

// GetCompletedOnOk returns a tuple with the CompletedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultHistoryReportModel) GetCompletedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedOn.Get(), o.CompletedOn.IsSet()
}

// HasCompletedOn returns a boolean if a field has been set.
func (o *TestResultHistoryReportModel) HasCompletedOn() bool {
	if o != nil && o.CompletedOn.IsSet() {
		return true
	}

	return false
}

// SetCompletedOn gets a reference to the given NullableTime and assigns it to the CompletedOn field.
func (o *TestResultHistoryReportModel) SetCompletedOn(v time.Time) {
	o.CompletedOn.Set(&v)
}
// SetCompletedOnNil sets the value for CompletedOn to be an explicit nil
func (o *TestResultHistoryReportModel) SetCompletedOnNil() {
	o.CompletedOn.Set(nil)
}

// UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
func (o *TestResultHistoryReportModel) UnsetCompletedOn() {
	o.CompletedOn.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultHistoryReportModel) GetDuration() int64 {
	if o == nil || IsNil(o.Duration.Get()) {
		var ret int64
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultHistoryReportModel) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *TestResultHistoryReportModel) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableInt64 and assigns it to the Duration field.
func (o *TestResultHistoryReportModel) SetDuration(v int64) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *TestResultHistoryReportModel) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *TestResultHistoryReportModel) UnsetDuration() {
	o.Duration.Unset()
}

// GetCreatedById returns the CreatedById field value
func (o *TestResultHistoryReportModel) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *TestResultHistoryReportModel) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *TestResultHistoryReportModel) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultHistoryReportModel) GetModifiedById() string {
	if o == nil || IsNil(o.ModifiedById.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedById.Get()
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultHistoryReportModel) GetModifiedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedById.Get(), o.ModifiedById.IsSet()
}

// HasModifiedById returns a boolean if a field has been set.
func (o *TestResultHistoryReportModel) HasModifiedById() bool {
	if o != nil && o.ModifiedById.IsSet() {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given NullableString and assigns it to the ModifiedById field.
func (o *TestResultHistoryReportModel) SetModifiedById(v string) {
	o.ModifiedById.Set(&v)
}
// SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil
func (o *TestResultHistoryReportModel) SetModifiedByIdNil() {
	o.ModifiedById.Set(nil)
}

// UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
func (o *TestResultHistoryReportModel) UnsetModifiedById() {
	o.ModifiedById.Unset()
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultHistoryReportModel) GetAttachments() []AttachmentModel {
	if o == nil {
		var ret []AttachmentModel
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultHistoryReportModel) GetAttachmentsOk() ([]AttachmentModel, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *TestResultHistoryReportModel) HasAttachments() bool {
	if o != nil && IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []AttachmentModel and assigns it to the Attachments field.
func (o *TestResultHistoryReportModel) SetAttachments(v []AttachmentModel) {
	o.Attachments = v
}

// GetWorkItemVersionId returns the WorkItemVersionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultHistoryReportModel) GetWorkItemVersionId() string {
	if o == nil || IsNil(o.WorkItemVersionId.Get()) {
		var ret string
		return ret
	}
	return *o.WorkItemVersionId.Get()
}

// GetWorkItemVersionIdOk returns a tuple with the WorkItemVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultHistoryReportModel) GetWorkItemVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkItemVersionId.Get(), o.WorkItemVersionId.IsSet()
}

// HasWorkItemVersionId returns a boolean if a field has been set.
func (o *TestResultHistoryReportModel) HasWorkItemVersionId() bool {
	if o != nil && o.WorkItemVersionId.IsSet() {
		return true
	}

	return false
}

// SetWorkItemVersionId gets a reference to the given NullableString and assigns it to the WorkItemVersionId field.
func (o *TestResultHistoryReportModel) SetWorkItemVersionId(v string) {
	o.WorkItemVersionId.Set(&v)
}
// SetWorkItemVersionIdNil sets the value for WorkItemVersionId to be an explicit nil
func (o *TestResultHistoryReportModel) SetWorkItemVersionIdNil() {
	o.WorkItemVersionId.Set(nil)
}

// UnsetWorkItemVersionId ensures that no value is present for WorkItemVersionId, not even an explicit nil
func (o *TestResultHistoryReportModel) UnsetWorkItemVersionId() {
	o.WorkItemVersionId.Unset()
}

// GetWorkItemVersionNumber returns the WorkItemVersionNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultHistoryReportModel) GetWorkItemVersionNumber() int32 {
	if o == nil || IsNil(o.WorkItemVersionNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.WorkItemVersionNumber.Get()
}

// GetWorkItemVersionNumberOk returns a tuple with the WorkItemVersionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultHistoryReportModel) GetWorkItemVersionNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkItemVersionNumber.Get(), o.WorkItemVersionNumber.IsSet()
}

// HasWorkItemVersionNumber returns a boolean if a field has been set.
func (o *TestResultHistoryReportModel) HasWorkItemVersionNumber() bool {
	if o != nil && o.WorkItemVersionNumber.IsSet() {
		return true
	}

	return false
}

// SetWorkItemVersionNumber gets a reference to the given NullableInt32 and assigns it to the WorkItemVersionNumber field.
func (o *TestResultHistoryReportModel) SetWorkItemVersionNumber(v int32) {
	o.WorkItemVersionNumber.Set(&v)
}
// SetWorkItemVersionNumberNil sets the value for WorkItemVersionNumber to be an explicit nil
func (o *TestResultHistoryReportModel) SetWorkItemVersionNumberNil() {
	o.WorkItemVersionNumber.Set(nil)
}

// UnsetWorkItemVersionNumber ensures that no value is present for WorkItemVersionNumber, not even an explicit nil
func (o *TestResultHistoryReportModel) UnsetWorkItemVersionNumber() {
	o.WorkItemVersionNumber.Unset()
}

// GetLaunchSource returns the LaunchSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultHistoryReportModel) GetLaunchSource() string {
	if o == nil || IsNil(o.LaunchSource.Get()) {
		var ret string
		return ret
	}
	return *o.LaunchSource.Get()
}

// GetLaunchSourceOk returns a tuple with the LaunchSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultHistoryReportModel) GetLaunchSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LaunchSource.Get(), o.LaunchSource.IsSet()
}

// HasLaunchSource returns a boolean if a field has been set.
func (o *TestResultHistoryReportModel) HasLaunchSource() bool {
	if o != nil && o.LaunchSource.IsSet() {
		return true
	}

	return false
}

// SetLaunchSource gets a reference to the given NullableString and assigns it to the LaunchSource field.
func (o *TestResultHistoryReportModel) SetLaunchSource(v string) {
	o.LaunchSource.Set(&v)
}
// SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil
func (o *TestResultHistoryReportModel) SetLaunchSourceNil() {
	o.LaunchSource.Set(nil)
}

// UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
func (o *TestResultHistoryReportModel) UnsetLaunchSource() {
	o.LaunchSource.Unset()
}

// GetFailureClassIds returns the FailureClassIds field value
func (o *TestResultHistoryReportModel) GetFailureClassIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.FailureClassIds
}

// GetFailureClassIdsOk returns a tuple with the FailureClassIds field value
// and a boolean to check if the value has been set.
func (o *TestResultHistoryReportModel) GetFailureClassIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureClassIds, true
}

// SetFailureClassIds sets field value
func (o *TestResultHistoryReportModel) SetFailureClassIds(v []string) {
	o.FailureClassIds = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultHistoryReportModel) GetParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultHistoryReportModel) GetParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *TestResultHistoryReportModel) HasParameters() bool {
	if o != nil && IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *TestResultHistoryReportModel) SetParameters(v map[string]string) {
	o.Parameters = v
}

func (o TestResultHistoryReportModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestResultHistoryReportModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["createdDate"] = o.CreatedDate
	toSerialize["modifiedDate"] = o.ModifiedDate
	toSerialize["userId"] = o.UserId
	if o.TestRunId.IsSet() {
		toSerialize["testRunId"] = o.TestRunId.Get()
	}
	if o.TestRunName.IsSet() {
		toSerialize["testRunName"] = o.TestRunName.Get()
	}
	if o.CreatedByUserName.IsSet() {
		toSerialize["createdByUserName"] = o.CreatedByUserName.Get()
	}
	if o.TestPlanId.IsSet() {
		toSerialize["testPlanId"] = o.TestPlanId.Get()
	}
	if o.TestPlanGlobalId.IsSet() {
		toSerialize["testPlanGlobalId"] = o.TestPlanGlobalId.Get()
	}
	if o.TestPlanName.IsSet() {
		toSerialize["testPlanName"] = o.TestPlanName.Get()
	}
	if o.ConfigurationName.IsSet() {
		toSerialize["configurationName"] = o.ConfigurationName.Get()
	}
	toSerialize["isAutomated"] = o.IsAutomated
	if o.Outcome.IsSet() {
		toSerialize["outcome"] = o.Outcome.Get()
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.StartedOn.IsSet() {
		toSerialize["startedOn"] = o.StartedOn.Get()
	}
	if o.CompletedOn.IsSet() {
		toSerialize["completedOn"] = o.CompletedOn.Get()
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	toSerialize["createdById"] = o.CreatedById
	if o.ModifiedById.IsSet() {
		toSerialize["modifiedById"] = o.ModifiedById.Get()
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.WorkItemVersionId.IsSet() {
		toSerialize["workItemVersionId"] = o.WorkItemVersionId.Get()
	}
	if o.WorkItemVersionNumber.IsSet() {
		toSerialize["workItemVersionNumber"] = o.WorkItemVersionNumber.Get()
	}
	if o.LaunchSource.IsSet() {
		toSerialize["launchSource"] = o.LaunchSource.Get()
	}
	toSerialize["failureClassIds"] = o.FailureClassIds
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	return toSerialize, nil
}

func (o *TestResultHistoryReportModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"createdDate",
		"modifiedDate",
		"userId",
		"isAutomated",
		"createdById",
		"failureClassIds",
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

	varTestResultHistoryReportModel := _TestResultHistoryReportModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestResultHistoryReportModel)

	if err != nil {
		return err
	}

	*o = TestResultHistoryReportModel(varTestResultHistoryReportModel)

	return err
}

type NullableTestResultHistoryReportModel struct {
	value *TestResultHistoryReportModel
	isSet bool
}

func (v NullableTestResultHistoryReportModel) Get() *TestResultHistoryReportModel {
	return v.value
}

func (v *NullableTestResultHistoryReportModel) Set(val *TestResultHistoryReportModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestResultHistoryReportModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestResultHistoryReportModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestResultHistoryReportModel(val *TestResultHistoryReportModel) *NullableTestResultHistoryReportModel {
	return &NullableTestResultHistoryReportModel{value: val, isSet: true}
}

func (v NullableTestResultHistoryReportModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestResultHistoryReportModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


