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

// checks if the TestRunApiResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestRunApiResult{}

// TestRunApiResult struct for TestRunApiResult
type TestRunApiResult struct {
	// Unique ID of the entity
	Id string `json:"id"`
	// Indicates if the entity is deleted
	IsDeleted bool `json:"isDeleted"`
	StartedDate NullableTime `json:"startedDate,omitempty"`
	CompletedDate NullableTime `json:"completedDate,omitempty"`
	Build string `json:"build"`
	Description NullableString `json:"description,omitempty"`
	// Deprecated
	StateName TestRunState `json:"stateName"`
	Status TestStatusApiResult `json:"status"`
	ProjectId string `json:"projectId"`
	TestPlanId NullableString `json:"testPlanId,omitempty"`
	RunByUserId NullableString `json:"runByUserId,omitempty"`
	StoppedByUserId NullableString `json:"stoppedByUserId,omitempty"`
	Name NullableString `json:"name,omitempty"`
	LaunchSource NullableString `json:"launchSource,omitempty"`
	AutoTests []AutoTestApiResult `json:"autoTests"`
	AutoTestsCount int32 `json:"autoTestsCount"`
	TestSuiteIds []string `json:"testSuiteIds"`
	IsAutomated bool `json:"isAutomated"`
	Analytic TestRunAnalyticApiResult `json:"analytic"`
	TestResults []TestResultApiResult `json:"testResults"`
	TestPlan NullableTestPlanApiResult `json:"testPlan,omitempty"`
	CreatedDate time.Time `json:"createdDate"`
	ModifiedDate NullableTime `json:"modifiedDate,omitempty"`
	CreatedById string `json:"createdById"`
	ModifiedById NullableString `json:"modifiedById,omitempty"`
	CreatedByUserName NullableString `json:"createdByUserName,omitempty"`
}

type _TestRunApiResult TestRunApiResult

// NewTestRunApiResult instantiates a new TestRunApiResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestRunApiResult(id string, isDeleted bool, build string, stateName TestRunState, status TestStatusApiResult, projectId string, autoTests []AutoTestApiResult, autoTestsCount int32, testSuiteIds []string, isAutomated bool, analytic TestRunAnalyticApiResult, testResults []TestResultApiResult, createdDate time.Time, createdById string) *TestRunApiResult {
	this := TestRunApiResult{}
	this.Id = id
	this.IsDeleted = isDeleted
	this.Build = build
	this.StateName = stateName
	this.Status = status
	this.ProjectId = projectId
	this.AutoTests = autoTests
	this.AutoTestsCount = autoTestsCount
	this.TestSuiteIds = testSuiteIds
	this.IsAutomated = isAutomated
	this.Analytic = analytic
	this.TestResults = testResults
	this.CreatedDate = createdDate
	this.CreatedById = createdById
	return &this
}

// NewTestRunApiResultWithDefaults instantiates a new TestRunApiResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestRunApiResultWithDefaults() *TestRunApiResult {
	this := TestRunApiResult{}
	return &this
}

// GetId returns the Id field value
func (o *TestRunApiResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestRunApiResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestRunApiResult) SetId(v string) {
	o.Id = v
}

// GetIsDeleted returns the IsDeleted field value
func (o *TestRunApiResult) GetIsDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value
// and a boolean to check if the value has been set.
func (o *TestRunApiResult) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeleted, true
}

// SetIsDeleted sets field value
func (o *TestRunApiResult) SetIsDeleted(v bool) {
	o.IsDeleted = v
}

// GetStartedDate returns the StartedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunApiResult) GetStartedDate() time.Time {
	if o == nil || IsNil(o.StartedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartedDate.Get()
}

// GetStartedDateOk returns a tuple with the StartedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunApiResult) GetStartedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedDate.Get(), o.StartedDate.IsSet()
}

// HasStartedDate returns a boolean if a field has been set.
func (o *TestRunApiResult) HasStartedDate() bool {
	if o != nil && o.StartedDate.IsSet() {
		return true
	}

	return false
}

// SetStartedDate gets a reference to the given NullableTime and assigns it to the StartedDate field.
func (o *TestRunApiResult) SetStartedDate(v time.Time) {
	o.StartedDate.Set(&v)
}
// SetStartedDateNil sets the value for StartedDate to be an explicit nil
func (o *TestRunApiResult) SetStartedDateNil() {
	o.StartedDate.Set(nil)
}

// UnsetStartedDate ensures that no value is present for StartedDate, not even an explicit nil
func (o *TestRunApiResult) UnsetStartedDate() {
	o.StartedDate.Unset()
}

// GetCompletedDate returns the CompletedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunApiResult) GetCompletedDate() time.Time {
	if o == nil || IsNil(o.CompletedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CompletedDate.Get()
}

// GetCompletedDateOk returns a tuple with the CompletedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunApiResult) GetCompletedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedDate.Get(), o.CompletedDate.IsSet()
}

// HasCompletedDate returns a boolean if a field has been set.
func (o *TestRunApiResult) HasCompletedDate() bool {
	if o != nil && o.CompletedDate.IsSet() {
		return true
	}

	return false
}

// SetCompletedDate gets a reference to the given NullableTime and assigns it to the CompletedDate field.
func (o *TestRunApiResult) SetCompletedDate(v time.Time) {
	o.CompletedDate.Set(&v)
}
// SetCompletedDateNil sets the value for CompletedDate to be an explicit nil
func (o *TestRunApiResult) SetCompletedDateNil() {
	o.CompletedDate.Set(nil)
}

// UnsetCompletedDate ensures that no value is present for CompletedDate, not even an explicit nil
func (o *TestRunApiResult) UnsetCompletedDate() {
	o.CompletedDate.Unset()
}

// GetBuild returns the Build field value
func (o *TestRunApiResult) GetBuild() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Build
}

// GetBuildOk returns a tuple with the Build field value
// and a boolean to check if the value has been set.
func (o *TestRunApiResult) GetBuildOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Build, true
}

// SetBuild sets field value
func (o *TestRunApiResult) SetBuild(v string) {
	o.Build = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunApiResult) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunApiResult) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TestRunApiResult) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TestRunApiResult) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TestRunApiResult) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TestRunApiResult) UnsetDescription() {
	o.Description.Unset()
}

// GetStateName returns the StateName field value
// Deprecated
func (o *TestRunApiResult) GetStateName() TestRunState {
	if o == nil {
		var ret TestRunState
		return ret
	}

	return o.StateName
}

// GetStateNameOk returns a tuple with the StateName field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *TestRunApiResult) GetStateNameOk() (*TestRunState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateName, true
}

// SetStateName sets field value
// Deprecated
func (o *TestRunApiResult) SetStateName(v TestRunState) {
	o.StateName = v
}

// GetStatus returns the Status field value
func (o *TestRunApiResult) GetStatus() TestStatusApiResult {
	if o == nil {
		var ret TestStatusApiResult
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TestRunApiResult) GetStatusOk() (*TestStatusApiResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TestRunApiResult) SetStatus(v TestStatusApiResult) {
	o.Status = v
}

// GetProjectId returns the ProjectId field value
func (o *TestRunApiResult) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *TestRunApiResult) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *TestRunApiResult) SetProjectId(v string) {
	o.ProjectId = v
}

// GetTestPlanId returns the TestPlanId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunApiResult) GetTestPlanId() string {
	if o == nil || IsNil(o.TestPlanId.Get()) {
		var ret string
		return ret
	}
	return *o.TestPlanId.Get()
}

// GetTestPlanIdOk returns a tuple with the TestPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunApiResult) GetTestPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestPlanId.Get(), o.TestPlanId.IsSet()
}

// HasTestPlanId returns a boolean if a field has been set.
func (o *TestRunApiResult) HasTestPlanId() bool {
	if o != nil && o.TestPlanId.IsSet() {
		return true
	}

	return false
}

// SetTestPlanId gets a reference to the given NullableString and assigns it to the TestPlanId field.
func (o *TestRunApiResult) SetTestPlanId(v string) {
	o.TestPlanId.Set(&v)
}
// SetTestPlanIdNil sets the value for TestPlanId to be an explicit nil
func (o *TestRunApiResult) SetTestPlanIdNil() {
	o.TestPlanId.Set(nil)
}

// UnsetTestPlanId ensures that no value is present for TestPlanId, not even an explicit nil
func (o *TestRunApiResult) UnsetTestPlanId() {
	o.TestPlanId.Unset()
}

// GetRunByUserId returns the RunByUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunApiResult) GetRunByUserId() string {
	if o == nil || IsNil(o.RunByUserId.Get()) {
		var ret string
		return ret
	}
	return *o.RunByUserId.Get()
}

// GetRunByUserIdOk returns a tuple with the RunByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunApiResult) GetRunByUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunByUserId.Get(), o.RunByUserId.IsSet()
}

// HasRunByUserId returns a boolean if a field has been set.
func (o *TestRunApiResult) HasRunByUserId() bool {
	if o != nil && o.RunByUserId.IsSet() {
		return true
	}

	return false
}

// SetRunByUserId gets a reference to the given NullableString and assigns it to the RunByUserId field.
func (o *TestRunApiResult) SetRunByUserId(v string) {
	o.RunByUserId.Set(&v)
}
// SetRunByUserIdNil sets the value for RunByUserId to be an explicit nil
func (o *TestRunApiResult) SetRunByUserIdNil() {
	o.RunByUserId.Set(nil)
}

// UnsetRunByUserId ensures that no value is present for RunByUserId, not even an explicit nil
func (o *TestRunApiResult) UnsetRunByUserId() {
	o.RunByUserId.Unset()
}

// GetStoppedByUserId returns the StoppedByUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunApiResult) GetStoppedByUserId() string {
	if o == nil || IsNil(o.StoppedByUserId.Get()) {
		var ret string
		return ret
	}
	return *o.StoppedByUserId.Get()
}

// GetStoppedByUserIdOk returns a tuple with the StoppedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunApiResult) GetStoppedByUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoppedByUserId.Get(), o.StoppedByUserId.IsSet()
}

// HasStoppedByUserId returns a boolean if a field has been set.
func (o *TestRunApiResult) HasStoppedByUserId() bool {
	if o != nil && o.StoppedByUserId.IsSet() {
		return true
	}

	return false
}

// SetStoppedByUserId gets a reference to the given NullableString and assigns it to the StoppedByUserId field.
func (o *TestRunApiResult) SetStoppedByUserId(v string) {
	o.StoppedByUserId.Set(&v)
}
// SetStoppedByUserIdNil sets the value for StoppedByUserId to be an explicit nil
func (o *TestRunApiResult) SetStoppedByUserIdNil() {
	o.StoppedByUserId.Set(nil)
}

// UnsetStoppedByUserId ensures that no value is present for StoppedByUserId, not even an explicit nil
func (o *TestRunApiResult) UnsetStoppedByUserId() {
	o.StoppedByUserId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunApiResult) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunApiResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TestRunApiResult) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TestRunApiResult) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TestRunApiResult) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TestRunApiResult) UnsetName() {
	o.Name.Unset()
}

// GetLaunchSource returns the LaunchSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunApiResult) GetLaunchSource() string {
	if o == nil || IsNil(o.LaunchSource.Get()) {
		var ret string
		return ret
	}
	return *o.LaunchSource.Get()
}

// GetLaunchSourceOk returns a tuple with the LaunchSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunApiResult) GetLaunchSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LaunchSource.Get(), o.LaunchSource.IsSet()
}

// HasLaunchSource returns a boolean if a field has been set.
func (o *TestRunApiResult) HasLaunchSource() bool {
	if o != nil && o.LaunchSource.IsSet() {
		return true
	}

	return false
}

// SetLaunchSource gets a reference to the given NullableString and assigns it to the LaunchSource field.
func (o *TestRunApiResult) SetLaunchSource(v string) {
	o.LaunchSource.Set(&v)
}
// SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil
func (o *TestRunApiResult) SetLaunchSourceNil() {
	o.LaunchSource.Set(nil)
}

// UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
func (o *TestRunApiResult) UnsetLaunchSource() {
	o.LaunchSource.Unset()
}

// GetAutoTests returns the AutoTests field value
func (o *TestRunApiResult) GetAutoTests() []AutoTestApiResult {
	if o == nil {
		var ret []AutoTestApiResult
		return ret
	}

	return o.AutoTests
}

// GetAutoTestsOk returns a tuple with the AutoTests field value
// and a boolean to check if the value has been set.
func (o *TestRunApiResult) GetAutoTestsOk() ([]AutoTestApiResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoTests, true
}

// SetAutoTests sets field value
func (o *TestRunApiResult) SetAutoTests(v []AutoTestApiResult) {
	o.AutoTests = v
}

// GetAutoTestsCount returns the AutoTestsCount field value
func (o *TestRunApiResult) GetAutoTestsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AutoTestsCount
}

// GetAutoTestsCountOk returns a tuple with the AutoTestsCount field value
// and a boolean to check if the value has been set.
func (o *TestRunApiResult) GetAutoTestsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoTestsCount, true
}

// SetAutoTestsCount sets field value
func (o *TestRunApiResult) SetAutoTestsCount(v int32) {
	o.AutoTestsCount = v
}

// GetTestSuiteIds returns the TestSuiteIds field value
func (o *TestRunApiResult) GetTestSuiteIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TestSuiteIds
}

// GetTestSuiteIdsOk returns a tuple with the TestSuiteIds field value
// and a boolean to check if the value has been set.
func (o *TestRunApiResult) GetTestSuiteIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestSuiteIds, true
}

// SetTestSuiteIds sets field value
func (o *TestRunApiResult) SetTestSuiteIds(v []string) {
	o.TestSuiteIds = v
}

// GetIsAutomated returns the IsAutomated field value
func (o *TestRunApiResult) GetIsAutomated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAutomated
}

// GetIsAutomatedOk returns a tuple with the IsAutomated field value
// and a boolean to check if the value has been set.
func (o *TestRunApiResult) GetIsAutomatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAutomated, true
}

// SetIsAutomated sets field value
func (o *TestRunApiResult) SetIsAutomated(v bool) {
	o.IsAutomated = v
}

// GetAnalytic returns the Analytic field value
func (o *TestRunApiResult) GetAnalytic() TestRunAnalyticApiResult {
	if o == nil {
		var ret TestRunAnalyticApiResult
		return ret
	}

	return o.Analytic
}

// GetAnalyticOk returns a tuple with the Analytic field value
// and a boolean to check if the value has been set.
func (o *TestRunApiResult) GetAnalyticOk() (*TestRunAnalyticApiResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Analytic, true
}

// SetAnalytic sets field value
func (o *TestRunApiResult) SetAnalytic(v TestRunAnalyticApiResult) {
	o.Analytic = v
}

// GetTestResults returns the TestResults field value
func (o *TestRunApiResult) GetTestResults() []TestResultApiResult {
	if o == nil {
		var ret []TestResultApiResult
		return ret
	}

	return o.TestResults
}

// GetTestResultsOk returns a tuple with the TestResults field value
// and a boolean to check if the value has been set.
func (o *TestRunApiResult) GetTestResultsOk() ([]TestResultApiResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestResults, true
}

// SetTestResults sets field value
func (o *TestRunApiResult) SetTestResults(v []TestResultApiResult) {
	o.TestResults = v
}

// GetTestPlan returns the TestPlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunApiResult) GetTestPlan() TestPlanApiResult {
	if o == nil || IsNil(o.TestPlan.Get()) {
		var ret TestPlanApiResult
		return ret
	}
	return *o.TestPlan.Get()
}

// GetTestPlanOk returns a tuple with the TestPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunApiResult) GetTestPlanOk() (*TestPlanApiResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestPlan.Get(), o.TestPlan.IsSet()
}

// HasTestPlan returns a boolean if a field has been set.
func (o *TestRunApiResult) HasTestPlan() bool {
	if o != nil && o.TestPlan.IsSet() {
		return true
	}

	return false
}

// SetTestPlan gets a reference to the given NullableTestPlanApiResult and assigns it to the TestPlan field.
func (o *TestRunApiResult) SetTestPlan(v TestPlanApiResult) {
	o.TestPlan.Set(&v)
}
// SetTestPlanNil sets the value for TestPlan to be an explicit nil
func (o *TestRunApiResult) SetTestPlanNil() {
	o.TestPlan.Set(nil)
}

// UnsetTestPlan ensures that no value is present for TestPlan, not even an explicit nil
func (o *TestRunApiResult) UnsetTestPlan() {
	o.TestPlan.Unset()
}

// GetCreatedDate returns the CreatedDate field value
func (o *TestRunApiResult) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *TestRunApiResult) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *TestRunApiResult) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunApiResult) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunApiResult) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *TestRunApiResult) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTime and assigns it to the ModifiedDate field.
func (o *TestRunApiResult) SetModifiedDate(v time.Time) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *TestRunApiResult) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *TestRunApiResult) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetCreatedById returns the CreatedById field value
func (o *TestRunApiResult) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *TestRunApiResult) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *TestRunApiResult) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunApiResult) GetModifiedById() string {
	if o == nil || IsNil(o.ModifiedById.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedById.Get()
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunApiResult) GetModifiedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedById.Get(), o.ModifiedById.IsSet()
}

// HasModifiedById returns a boolean if a field has been set.
func (o *TestRunApiResult) HasModifiedById() bool {
	if o != nil && o.ModifiedById.IsSet() {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given NullableString and assigns it to the ModifiedById field.
func (o *TestRunApiResult) SetModifiedById(v string) {
	o.ModifiedById.Set(&v)
}
// SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil
func (o *TestRunApiResult) SetModifiedByIdNil() {
	o.ModifiedById.Set(nil)
}

// UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
func (o *TestRunApiResult) UnsetModifiedById() {
	o.ModifiedById.Unset()
}

// GetCreatedByUserName returns the CreatedByUserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunApiResult) GetCreatedByUserName() string {
	if o == nil || IsNil(o.CreatedByUserName.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedByUserName.Get()
}

// GetCreatedByUserNameOk returns a tuple with the CreatedByUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunApiResult) GetCreatedByUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedByUserName.Get(), o.CreatedByUserName.IsSet()
}

// HasCreatedByUserName returns a boolean if a field has been set.
func (o *TestRunApiResult) HasCreatedByUserName() bool {
	if o != nil && o.CreatedByUserName.IsSet() {
		return true
	}

	return false
}

// SetCreatedByUserName gets a reference to the given NullableString and assigns it to the CreatedByUserName field.
func (o *TestRunApiResult) SetCreatedByUserName(v string) {
	o.CreatedByUserName.Set(&v)
}
// SetCreatedByUserNameNil sets the value for CreatedByUserName to be an explicit nil
func (o *TestRunApiResult) SetCreatedByUserNameNil() {
	o.CreatedByUserName.Set(nil)
}

// UnsetCreatedByUserName ensures that no value is present for CreatedByUserName, not even an explicit nil
func (o *TestRunApiResult) UnsetCreatedByUserName() {
	o.CreatedByUserName.Unset()
}

func (o TestRunApiResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestRunApiResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["isDeleted"] = o.IsDeleted
	if o.StartedDate.IsSet() {
		toSerialize["startedDate"] = o.StartedDate.Get()
	}
	if o.CompletedDate.IsSet() {
		toSerialize["completedDate"] = o.CompletedDate.Get()
	}
	toSerialize["build"] = o.Build
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["stateName"] = o.StateName
	toSerialize["status"] = o.Status
	toSerialize["projectId"] = o.ProjectId
	if o.TestPlanId.IsSet() {
		toSerialize["testPlanId"] = o.TestPlanId.Get()
	}
	if o.RunByUserId.IsSet() {
		toSerialize["runByUserId"] = o.RunByUserId.Get()
	}
	if o.StoppedByUserId.IsSet() {
		toSerialize["stoppedByUserId"] = o.StoppedByUserId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.LaunchSource.IsSet() {
		toSerialize["launchSource"] = o.LaunchSource.Get()
	}
	toSerialize["autoTests"] = o.AutoTests
	toSerialize["autoTestsCount"] = o.AutoTestsCount
	toSerialize["testSuiteIds"] = o.TestSuiteIds
	toSerialize["isAutomated"] = o.IsAutomated
	toSerialize["analytic"] = o.Analytic
	toSerialize["testResults"] = o.TestResults
	if o.TestPlan.IsSet() {
		toSerialize["testPlan"] = o.TestPlan.Get()
	}
	toSerialize["createdDate"] = o.CreatedDate
	if o.ModifiedDate.IsSet() {
		toSerialize["modifiedDate"] = o.ModifiedDate.Get()
	}
	toSerialize["createdById"] = o.CreatedById
	if o.ModifiedById.IsSet() {
		toSerialize["modifiedById"] = o.ModifiedById.Get()
	}
	if o.CreatedByUserName.IsSet() {
		toSerialize["createdByUserName"] = o.CreatedByUserName.Get()
	}
	return toSerialize, nil
}

func (o *TestRunApiResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"isDeleted",
		"build",
		"stateName",
		"status",
		"projectId",
		"autoTests",
		"autoTestsCount",
		"testSuiteIds",
		"isAutomated",
		"analytic",
		"testResults",
		"createdDate",
		"createdById",
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

	varTestRunApiResult := _TestRunApiResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestRunApiResult)

	if err != nil {
		return err
	}

	*o = TestRunApiResult(varTestRunApiResult)

	return err
}

type NullableTestRunApiResult struct {
	value *TestRunApiResult
	isSet bool
}

func (v NullableTestRunApiResult) Get() *TestRunApiResult {
	return v.value
}

func (v *NullableTestRunApiResult) Set(val *TestRunApiResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTestRunApiResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTestRunApiResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestRunApiResult(val *TestRunApiResult) *NullableTestRunApiResult {
	return &NullableTestRunApiResult{value: val, isSet: true}
}

func (v NullableTestRunApiResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestRunApiResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


