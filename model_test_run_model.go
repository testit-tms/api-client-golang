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

// checks if the TestRunModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestRunModel{}

// TestRunModel struct for TestRunModel
type TestRunModel struct {
	AutoTests []AutoTestModel `json:"autoTests,omitempty"`
	AutoTestsCount int32 `json:"autoTestsCount"`
	TestSuiteIds []string `json:"testSuiteIds,omitempty"`
	IsAutomated bool `json:"isAutomated"`
	Analytic TestRunAnalyticResultModel `json:"analytic"`
	TestResults []TestResultModel `json:"testResults,omitempty"`
	TestPlan NullableTestPlanModel `json:"testPlan,omitempty"`
	CreatedDate time.Time `json:"createdDate"`
	ModifiedDate NullableTime `json:"modifiedDate,omitempty"`
	CreatedById string `json:"createdById"`
	ModifiedById NullableString `json:"modifiedById,omitempty"`
	CreatedByUserName NullableString `json:"createdByUserName,omitempty"`
	StartedDate NullableTime `json:"startedDate,omitempty"`
	CompletedDate NullableTime `json:"completedDate,omitempty"`
	Build string `json:"build"`
	Description string `json:"description"`
	StateName TestRunState `json:"stateName"`
	ProjectId string `json:"projectId"`
	TestPlanId NullableString `json:"testPlanId,omitempty"`
	RunByUserId NullableString `json:"runByUserId,omitempty"`
	StoppedByUserId NullableString `json:"stoppedByUserId,omitempty"`
	Name string `json:"name"`
	LaunchSource string `json:"launchSource"`
	// Unique ID of the entity
	Id string `json:"id"`
	// Indicates if the entity is deleted
	IsDeleted bool `json:"isDeleted"`
}

type _TestRunModel TestRunModel

// NewTestRunModel instantiates a new TestRunModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestRunModel(autoTestsCount int32, isAutomated bool, analytic TestRunAnalyticResultModel, createdDate time.Time, createdById string, build string, description string, stateName TestRunState, projectId string, name string, launchSource string, id string, isDeleted bool) *TestRunModel {
	this := TestRunModel{}
	this.AutoTestsCount = autoTestsCount
	this.IsAutomated = isAutomated
	this.Analytic = analytic
	this.CreatedDate = createdDate
	this.CreatedById = createdById
	this.Build = build
	this.Description = description
	this.StateName = stateName
	this.ProjectId = projectId
	this.Name = name
	this.LaunchSource = launchSource
	this.Id = id
	this.IsDeleted = isDeleted
	return &this
}

// NewTestRunModelWithDefaults instantiates a new TestRunModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestRunModelWithDefaults() *TestRunModel {
	this := TestRunModel{}
	return &this
}

// GetAutoTests returns the AutoTests field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunModel) GetAutoTests() []AutoTestModel {
	if o == nil {
		var ret []AutoTestModel
		return ret
	}
	return o.AutoTests
}

// GetAutoTestsOk returns a tuple with the AutoTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunModel) GetAutoTestsOk() ([]AutoTestModel, bool) {
	if o == nil || IsNil(o.AutoTests) {
		return nil, false
	}
	return o.AutoTests, true
}

// HasAutoTests returns a boolean if a field has been set.
func (o *TestRunModel) HasAutoTests() bool {
	if o != nil && !IsNil(o.AutoTests) {
		return true
	}

	return false
}

// SetAutoTests gets a reference to the given []AutoTestModel and assigns it to the AutoTests field.
func (o *TestRunModel) SetAutoTests(v []AutoTestModel) {
	o.AutoTests = v
}

// GetAutoTestsCount returns the AutoTestsCount field value
func (o *TestRunModel) GetAutoTestsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AutoTestsCount
}

// GetAutoTestsCountOk returns a tuple with the AutoTestsCount field value
// and a boolean to check if the value has been set.
func (o *TestRunModel) GetAutoTestsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoTestsCount, true
}

// SetAutoTestsCount sets field value
func (o *TestRunModel) SetAutoTestsCount(v int32) {
	o.AutoTestsCount = v
}

// GetTestSuiteIds returns the TestSuiteIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunModel) GetTestSuiteIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TestSuiteIds
}

// GetTestSuiteIdsOk returns a tuple with the TestSuiteIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunModel) GetTestSuiteIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TestSuiteIds) {
		return nil, false
	}
	return o.TestSuiteIds, true
}

// HasTestSuiteIds returns a boolean if a field has been set.
func (o *TestRunModel) HasTestSuiteIds() bool {
	if o != nil && !IsNil(o.TestSuiteIds) {
		return true
	}

	return false
}

// SetTestSuiteIds gets a reference to the given []string and assigns it to the TestSuiteIds field.
func (o *TestRunModel) SetTestSuiteIds(v []string) {
	o.TestSuiteIds = v
}

// GetIsAutomated returns the IsAutomated field value
func (o *TestRunModel) GetIsAutomated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAutomated
}

// GetIsAutomatedOk returns a tuple with the IsAutomated field value
// and a boolean to check if the value has been set.
func (o *TestRunModel) GetIsAutomatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAutomated, true
}

// SetIsAutomated sets field value
func (o *TestRunModel) SetIsAutomated(v bool) {
	o.IsAutomated = v
}

// GetAnalytic returns the Analytic field value
func (o *TestRunModel) GetAnalytic() TestRunAnalyticResultModel {
	if o == nil {
		var ret TestRunAnalyticResultModel
		return ret
	}

	return o.Analytic
}

// GetAnalyticOk returns a tuple with the Analytic field value
// and a boolean to check if the value has been set.
func (o *TestRunModel) GetAnalyticOk() (*TestRunAnalyticResultModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Analytic, true
}

// SetAnalytic sets field value
func (o *TestRunModel) SetAnalytic(v TestRunAnalyticResultModel) {
	o.Analytic = v
}

// GetTestResults returns the TestResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunModel) GetTestResults() []TestResultModel {
	if o == nil {
		var ret []TestResultModel
		return ret
	}
	return o.TestResults
}

// GetTestResultsOk returns a tuple with the TestResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunModel) GetTestResultsOk() ([]TestResultModel, bool) {
	if o == nil || IsNil(o.TestResults) {
		return nil, false
	}
	return o.TestResults, true
}

// HasTestResults returns a boolean if a field has been set.
func (o *TestRunModel) HasTestResults() bool {
	if o != nil && !IsNil(o.TestResults) {
		return true
	}

	return false
}

// SetTestResults gets a reference to the given []TestResultModel and assigns it to the TestResults field.
func (o *TestRunModel) SetTestResults(v []TestResultModel) {
	o.TestResults = v
}

// GetTestPlan returns the TestPlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunModel) GetTestPlan() TestPlanModel {
	if o == nil || IsNil(o.TestPlan.Get()) {
		var ret TestPlanModel
		return ret
	}
	return *o.TestPlan.Get()
}

// GetTestPlanOk returns a tuple with the TestPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunModel) GetTestPlanOk() (*TestPlanModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestPlan.Get(), o.TestPlan.IsSet()
}

// HasTestPlan returns a boolean if a field has been set.
func (o *TestRunModel) HasTestPlan() bool {
	if o != nil && o.TestPlan.IsSet() {
		return true
	}

	return false
}

// SetTestPlan gets a reference to the given NullableTestPlanModel and assigns it to the TestPlan field.
func (o *TestRunModel) SetTestPlan(v TestPlanModel) {
	o.TestPlan.Set(&v)
}
// SetTestPlanNil sets the value for TestPlan to be an explicit nil
func (o *TestRunModel) SetTestPlanNil() {
	o.TestPlan.Set(nil)
}

// UnsetTestPlan ensures that no value is present for TestPlan, not even an explicit nil
func (o *TestRunModel) UnsetTestPlan() {
	o.TestPlan.Unset()
}

// GetCreatedDate returns the CreatedDate field value
func (o *TestRunModel) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *TestRunModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *TestRunModel) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunModel) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunModel) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *TestRunModel) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTime and assigns it to the ModifiedDate field.
func (o *TestRunModel) SetModifiedDate(v time.Time) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *TestRunModel) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *TestRunModel) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetCreatedById returns the CreatedById field value
func (o *TestRunModel) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *TestRunModel) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *TestRunModel) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunModel) GetModifiedById() string {
	if o == nil || IsNil(o.ModifiedById.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedById.Get()
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunModel) GetModifiedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedById.Get(), o.ModifiedById.IsSet()
}

// HasModifiedById returns a boolean if a field has been set.
func (o *TestRunModel) HasModifiedById() bool {
	if o != nil && o.ModifiedById.IsSet() {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given NullableString and assigns it to the ModifiedById field.
func (o *TestRunModel) SetModifiedById(v string) {
	o.ModifiedById.Set(&v)
}
// SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil
func (o *TestRunModel) SetModifiedByIdNil() {
	o.ModifiedById.Set(nil)
}

// UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
func (o *TestRunModel) UnsetModifiedById() {
	o.ModifiedById.Unset()
}

// GetCreatedByUserName returns the CreatedByUserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunModel) GetCreatedByUserName() string {
	if o == nil || IsNil(o.CreatedByUserName.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedByUserName.Get()
}

// GetCreatedByUserNameOk returns a tuple with the CreatedByUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunModel) GetCreatedByUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedByUserName.Get(), o.CreatedByUserName.IsSet()
}

// HasCreatedByUserName returns a boolean if a field has been set.
func (o *TestRunModel) HasCreatedByUserName() bool {
	if o != nil && o.CreatedByUserName.IsSet() {
		return true
	}

	return false
}

// SetCreatedByUserName gets a reference to the given NullableString and assigns it to the CreatedByUserName field.
func (o *TestRunModel) SetCreatedByUserName(v string) {
	o.CreatedByUserName.Set(&v)
}
// SetCreatedByUserNameNil sets the value for CreatedByUserName to be an explicit nil
func (o *TestRunModel) SetCreatedByUserNameNil() {
	o.CreatedByUserName.Set(nil)
}

// UnsetCreatedByUserName ensures that no value is present for CreatedByUserName, not even an explicit nil
func (o *TestRunModel) UnsetCreatedByUserName() {
	o.CreatedByUserName.Unset()
}

// GetStartedDate returns the StartedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunModel) GetStartedDate() time.Time {
	if o == nil || IsNil(o.StartedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartedDate.Get()
}

// GetStartedDateOk returns a tuple with the StartedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunModel) GetStartedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedDate.Get(), o.StartedDate.IsSet()
}

// HasStartedDate returns a boolean if a field has been set.
func (o *TestRunModel) HasStartedDate() bool {
	if o != nil && o.StartedDate.IsSet() {
		return true
	}

	return false
}

// SetStartedDate gets a reference to the given NullableTime and assigns it to the StartedDate field.
func (o *TestRunModel) SetStartedDate(v time.Time) {
	o.StartedDate.Set(&v)
}
// SetStartedDateNil sets the value for StartedDate to be an explicit nil
func (o *TestRunModel) SetStartedDateNil() {
	o.StartedDate.Set(nil)
}

// UnsetStartedDate ensures that no value is present for StartedDate, not even an explicit nil
func (o *TestRunModel) UnsetStartedDate() {
	o.StartedDate.Unset()
}

// GetCompletedDate returns the CompletedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunModel) GetCompletedDate() time.Time {
	if o == nil || IsNil(o.CompletedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CompletedDate.Get()
}

// GetCompletedDateOk returns a tuple with the CompletedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunModel) GetCompletedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedDate.Get(), o.CompletedDate.IsSet()
}

// HasCompletedDate returns a boolean if a field has been set.
func (o *TestRunModel) HasCompletedDate() bool {
	if o != nil && o.CompletedDate.IsSet() {
		return true
	}

	return false
}

// SetCompletedDate gets a reference to the given NullableTime and assigns it to the CompletedDate field.
func (o *TestRunModel) SetCompletedDate(v time.Time) {
	o.CompletedDate.Set(&v)
}
// SetCompletedDateNil sets the value for CompletedDate to be an explicit nil
func (o *TestRunModel) SetCompletedDateNil() {
	o.CompletedDate.Set(nil)
}

// UnsetCompletedDate ensures that no value is present for CompletedDate, not even an explicit nil
func (o *TestRunModel) UnsetCompletedDate() {
	o.CompletedDate.Unset()
}

// GetBuild returns the Build field value
func (o *TestRunModel) GetBuild() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Build
}

// GetBuildOk returns a tuple with the Build field value
// and a boolean to check if the value has been set.
func (o *TestRunModel) GetBuildOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Build, true
}

// SetBuild sets field value
func (o *TestRunModel) SetBuild(v string) {
	o.Build = v
}

// GetDescription returns the Description field value
func (o *TestRunModel) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TestRunModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TestRunModel) SetDescription(v string) {
	o.Description = v
}

// GetStateName returns the StateName field value
func (o *TestRunModel) GetStateName() TestRunState {
	if o == nil {
		var ret TestRunState
		return ret
	}

	return o.StateName
}

// GetStateNameOk returns a tuple with the StateName field value
// and a boolean to check if the value has been set.
func (o *TestRunModel) GetStateNameOk() (*TestRunState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateName, true
}

// SetStateName sets field value
func (o *TestRunModel) SetStateName(v TestRunState) {
	o.StateName = v
}

// GetProjectId returns the ProjectId field value
func (o *TestRunModel) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *TestRunModel) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *TestRunModel) SetProjectId(v string) {
	o.ProjectId = v
}

// GetTestPlanId returns the TestPlanId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunModel) GetTestPlanId() string {
	if o == nil || IsNil(o.TestPlanId.Get()) {
		var ret string
		return ret
	}
	return *o.TestPlanId.Get()
}

// GetTestPlanIdOk returns a tuple with the TestPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunModel) GetTestPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestPlanId.Get(), o.TestPlanId.IsSet()
}

// HasTestPlanId returns a boolean if a field has been set.
func (o *TestRunModel) HasTestPlanId() bool {
	if o != nil && o.TestPlanId.IsSet() {
		return true
	}

	return false
}

// SetTestPlanId gets a reference to the given NullableString and assigns it to the TestPlanId field.
func (o *TestRunModel) SetTestPlanId(v string) {
	o.TestPlanId.Set(&v)
}
// SetTestPlanIdNil sets the value for TestPlanId to be an explicit nil
func (o *TestRunModel) SetTestPlanIdNil() {
	o.TestPlanId.Set(nil)
}

// UnsetTestPlanId ensures that no value is present for TestPlanId, not even an explicit nil
func (o *TestRunModel) UnsetTestPlanId() {
	o.TestPlanId.Unset()
}

// GetRunByUserId returns the RunByUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunModel) GetRunByUserId() string {
	if o == nil || IsNil(o.RunByUserId.Get()) {
		var ret string
		return ret
	}
	return *o.RunByUserId.Get()
}

// GetRunByUserIdOk returns a tuple with the RunByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunModel) GetRunByUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunByUserId.Get(), o.RunByUserId.IsSet()
}

// HasRunByUserId returns a boolean if a field has been set.
func (o *TestRunModel) HasRunByUserId() bool {
	if o != nil && o.RunByUserId.IsSet() {
		return true
	}

	return false
}

// SetRunByUserId gets a reference to the given NullableString and assigns it to the RunByUserId field.
func (o *TestRunModel) SetRunByUserId(v string) {
	o.RunByUserId.Set(&v)
}
// SetRunByUserIdNil sets the value for RunByUserId to be an explicit nil
func (o *TestRunModel) SetRunByUserIdNil() {
	o.RunByUserId.Set(nil)
}

// UnsetRunByUserId ensures that no value is present for RunByUserId, not even an explicit nil
func (o *TestRunModel) UnsetRunByUserId() {
	o.RunByUserId.Unset()
}

// GetStoppedByUserId returns the StoppedByUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunModel) GetStoppedByUserId() string {
	if o == nil || IsNil(o.StoppedByUserId.Get()) {
		var ret string
		return ret
	}
	return *o.StoppedByUserId.Get()
}

// GetStoppedByUserIdOk returns a tuple with the StoppedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunModel) GetStoppedByUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoppedByUserId.Get(), o.StoppedByUserId.IsSet()
}

// HasStoppedByUserId returns a boolean if a field has been set.
func (o *TestRunModel) HasStoppedByUserId() bool {
	if o != nil && o.StoppedByUserId.IsSet() {
		return true
	}

	return false
}

// SetStoppedByUserId gets a reference to the given NullableString and assigns it to the StoppedByUserId field.
func (o *TestRunModel) SetStoppedByUserId(v string) {
	o.StoppedByUserId.Set(&v)
}
// SetStoppedByUserIdNil sets the value for StoppedByUserId to be an explicit nil
func (o *TestRunModel) SetStoppedByUserIdNil() {
	o.StoppedByUserId.Set(nil)
}

// UnsetStoppedByUserId ensures that no value is present for StoppedByUserId, not even an explicit nil
func (o *TestRunModel) UnsetStoppedByUserId() {
	o.StoppedByUserId.Unset()
}

// GetName returns the Name field value
func (o *TestRunModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TestRunModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TestRunModel) SetName(v string) {
	o.Name = v
}

// GetLaunchSource returns the LaunchSource field value
func (o *TestRunModel) GetLaunchSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LaunchSource
}

// GetLaunchSourceOk returns a tuple with the LaunchSource field value
// and a boolean to check if the value has been set.
func (o *TestRunModel) GetLaunchSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LaunchSource, true
}

// SetLaunchSource sets field value
func (o *TestRunModel) SetLaunchSource(v string) {
	o.LaunchSource = v
}

// GetId returns the Id field value
func (o *TestRunModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestRunModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestRunModel) SetId(v string) {
	o.Id = v
}

// GetIsDeleted returns the IsDeleted field value
func (o *TestRunModel) GetIsDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value
// and a boolean to check if the value has been set.
func (o *TestRunModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeleted, true
}

// SetIsDeleted sets field value
func (o *TestRunModel) SetIsDeleted(v bool) {
	o.IsDeleted = v
}

func (o TestRunModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestRunModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoTests != nil {
		toSerialize["autoTests"] = o.AutoTests
	}
	toSerialize["autoTestsCount"] = o.AutoTestsCount
	if o.TestSuiteIds != nil {
		toSerialize["testSuiteIds"] = o.TestSuiteIds
	}
	toSerialize["isAutomated"] = o.IsAutomated
	toSerialize["analytic"] = o.Analytic
	if o.TestResults != nil {
		toSerialize["testResults"] = o.TestResults
	}
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
	if o.StartedDate.IsSet() {
		toSerialize["startedDate"] = o.StartedDate.Get()
	}
	if o.CompletedDate.IsSet() {
		toSerialize["completedDate"] = o.CompletedDate.Get()
	}
	toSerialize["build"] = o.Build
	toSerialize["description"] = o.Description
	toSerialize["stateName"] = o.StateName
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
	toSerialize["name"] = o.Name
	toSerialize["launchSource"] = o.LaunchSource
	toSerialize["id"] = o.Id
	toSerialize["isDeleted"] = o.IsDeleted
	return toSerialize, nil
}

func (o *TestRunModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"autoTestsCount",
		"isAutomated",
		"analytic",
		"createdDate",
		"createdById",
		"build",
		"description",
		"stateName",
		"projectId",
		"name",
		"launchSource",
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

	varTestRunModel := _TestRunModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestRunModel)

	if err != nil {
		return err
	}

	*o = TestRunModel(varTestRunModel)

	return err
}

type NullableTestRunModel struct {
	value *TestRunModel
	isSet bool
}

func (v NullableTestRunModel) Get() *TestRunModel {
	return v.value
}

func (v *NullableTestRunModel) Set(val *TestRunModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestRunModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestRunModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestRunModel(val *TestRunModel) *NullableTestRunModel {
	return &NullableTestRunModel{value: val, isSet: true}
}

func (v NullableTestRunModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestRunModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


