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

// checks if the TestPlanWithTestSuiteTreeModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPlanWithTestSuiteTreeModel{}

// TestPlanWithTestSuiteTreeModel struct for TestPlanWithTestSuiteTreeModel
type TestPlanWithTestSuiteTreeModel struct {
	TestSuites []TestSuiteWithChildrenModel `json:"testSuites"`
	Status TestPlanStatusModel `json:"status"`
	// Set when test plan is starter (status changed to: In Progress)
	StartedOn NullableTime `json:"startedOn,omitempty"`
	// set when test plan status is completed (status changed to: Completed)
	CompletedOn NullableTime `json:"completedOn,omitempty"`
	CreatedDate NullableTime `json:"createdDate,omitempty"`
	ModifiedDate NullableTime `json:"modifiedDate,omitempty"`
	CreatedById string `json:"createdById"`
	ModifiedById NullableString `json:"modifiedById,omitempty"`
	// Used for search Test plan
	GlobalId int64 `json:"globalId"`
	IsDeleted bool `json:"isDeleted"`
	LockedDate NullableTime `json:"lockedDate,omitempty"`
	Id string `json:"id"`
	LockedById NullableString `json:"lockedById,omitempty"`
	Tags []TagShortModel `json:"tags,omitempty"`
	Name string `json:"name"`
	// Used for analytics
	StartDate NullableTime `json:"startDate,omitempty"`
	// Used for analytics
	EndDate NullableTime `json:"endDate,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Build NullableString `json:"build,omitempty"`
	ProjectId string `json:"projectId"`
	ProductName NullableString `json:"productName,omitempty"`
	HasAutomaticDurationTimer NullableBool `json:"hasAutomaticDurationTimer,omitempty"`
	Attributes map[string]interface{} `json:"attributes"`
}

// NewTestPlanWithTestSuiteTreeModel instantiates a new TestPlanWithTestSuiteTreeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPlanWithTestSuiteTreeModel(testSuites []TestSuiteWithChildrenModel, status TestPlanStatusModel, createdById string, globalId int64, isDeleted bool, id string, name string, projectId string, attributes map[string]interface{}) *TestPlanWithTestSuiteTreeModel {
	this := TestPlanWithTestSuiteTreeModel{}
	this.TestSuites = testSuites
	this.Status = status
	this.CreatedById = createdById
	this.GlobalId = globalId
	this.IsDeleted = isDeleted
	this.Id = id
	this.Name = name
	this.ProjectId = projectId
	this.Attributes = attributes
	return &this
}

// NewTestPlanWithTestSuiteTreeModelWithDefaults instantiates a new TestPlanWithTestSuiteTreeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPlanWithTestSuiteTreeModelWithDefaults() *TestPlanWithTestSuiteTreeModel {
	this := TestPlanWithTestSuiteTreeModel{}
	return &this
}

// GetTestSuites returns the TestSuites field value
func (o *TestPlanWithTestSuiteTreeModel) GetTestSuites() []TestSuiteWithChildrenModel {
	if o == nil {
		var ret []TestSuiteWithChildrenModel
		return ret
	}

	return o.TestSuites
}

// GetTestSuitesOk returns a tuple with the TestSuites field value
// and a boolean to check if the value has been set.
func (o *TestPlanWithTestSuiteTreeModel) GetTestSuitesOk() ([]TestSuiteWithChildrenModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestSuites, true
}

// SetTestSuites sets field value
func (o *TestPlanWithTestSuiteTreeModel) SetTestSuites(v []TestSuiteWithChildrenModel) {
	o.TestSuites = v
}

// GetStatus returns the Status field value
func (o *TestPlanWithTestSuiteTreeModel) GetStatus() TestPlanStatusModel {
	if o == nil {
		var ret TestPlanStatusModel
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TestPlanWithTestSuiteTreeModel) GetStatusOk() (*TestPlanStatusModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TestPlanWithTestSuiteTreeModel) SetStatus(v TestPlanStatusModel) {
	o.Status = v
}

// GetStartedOn returns the StartedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanWithTestSuiteTreeModel) GetStartedOn() time.Time {
	if o == nil || IsNil(o.StartedOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartedOn.Get()
}

// GetStartedOnOk returns a tuple with the StartedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanWithTestSuiteTreeModel) GetStartedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedOn.Get(), o.StartedOn.IsSet()
}

// HasStartedOn returns a boolean if a field has been set.
func (o *TestPlanWithTestSuiteTreeModel) HasStartedOn() bool {
	if o != nil && o.StartedOn.IsSet() {
		return true
	}

	return false
}

// SetStartedOn gets a reference to the given NullableTime and assigns it to the StartedOn field.
func (o *TestPlanWithTestSuiteTreeModel) SetStartedOn(v time.Time) {
	o.StartedOn.Set(&v)
}
// SetStartedOnNil sets the value for StartedOn to be an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) SetStartedOnNil() {
	o.StartedOn.Set(nil)
}

// UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) UnsetStartedOn() {
	o.StartedOn.Unset()
}

// GetCompletedOn returns the CompletedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanWithTestSuiteTreeModel) GetCompletedOn() time.Time {
	if o == nil || IsNil(o.CompletedOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CompletedOn.Get()
}

// GetCompletedOnOk returns a tuple with the CompletedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanWithTestSuiteTreeModel) GetCompletedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedOn.Get(), o.CompletedOn.IsSet()
}

// HasCompletedOn returns a boolean if a field has been set.
func (o *TestPlanWithTestSuiteTreeModel) HasCompletedOn() bool {
	if o != nil && o.CompletedOn.IsSet() {
		return true
	}

	return false
}

// SetCompletedOn gets a reference to the given NullableTime and assigns it to the CompletedOn field.
func (o *TestPlanWithTestSuiteTreeModel) SetCompletedOn(v time.Time) {
	o.CompletedOn.Set(&v)
}
// SetCompletedOnNil sets the value for CompletedOn to be an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) SetCompletedOnNil() {
	o.CompletedOn.Set(nil)
}

// UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) UnsetCompletedOn() {
	o.CompletedOn.Unset()
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanWithTestSuiteTreeModel) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate.Get()
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanWithTestSuiteTreeModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDate.Get(), o.CreatedDate.IsSet()
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *TestPlanWithTestSuiteTreeModel) HasCreatedDate() bool {
	if o != nil && o.CreatedDate.IsSet() {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given NullableTime and assigns it to the CreatedDate field.
func (o *TestPlanWithTestSuiteTreeModel) SetCreatedDate(v time.Time) {
	o.CreatedDate.Set(&v)
}
// SetCreatedDateNil sets the value for CreatedDate to be an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) SetCreatedDateNil() {
	o.CreatedDate.Set(nil)
}

// UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) UnsetCreatedDate() {
	o.CreatedDate.Unset()
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanWithTestSuiteTreeModel) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanWithTestSuiteTreeModel) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *TestPlanWithTestSuiteTreeModel) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTime and assigns it to the ModifiedDate field.
func (o *TestPlanWithTestSuiteTreeModel) SetModifiedDate(v time.Time) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetCreatedById returns the CreatedById field value
func (o *TestPlanWithTestSuiteTreeModel) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *TestPlanWithTestSuiteTreeModel) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *TestPlanWithTestSuiteTreeModel) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanWithTestSuiteTreeModel) GetModifiedById() string {
	if o == nil || IsNil(o.ModifiedById.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedById.Get()
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanWithTestSuiteTreeModel) GetModifiedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedById.Get(), o.ModifiedById.IsSet()
}

// HasModifiedById returns a boolean if a field has been set.
func (o *TestPlanWithTestSuiteTreeModel) HasModifiedById() bool {
	if o != nil && o.ModifiedById.IsSet() {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given NullableString and assigns it to the ModifiedById field.
func (o *TestPlanWithTestSuiteTreeModel) SetModifiedById(v string) {
	o.ModifiedById.Set(&v)
}
// SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) SetModifiedByIdNil() {
	o.ModifiedById.Set(nil)
}

// UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) UnsetModifiedById() {
	o.ModifiedById.Unset()
}

// GetGlobalId returns the GlobalId field value
func (o *TestPlanWithTestSuiteTreeModel) GetGlobalId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value
// and a boolean to check if the value has been set.
func (o *TestPlanWithTestSuiteTreeModel) GetGlobalIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalId, true
}

// SetGlobalId sets field value
func (o *TestPlanWithTestSuiteTreeModel) SetGlobalId(v int64) {
	o.GlobalId = v
}

// GetIsDeleted returns the IsDeleted field value
func (o *TestPlanWithTestSuiteTreeModel) GetIsDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value
// and a boolean to check if the value has been set.
func (o *TestPlanWithTestSuiteTreeModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeleted, true
}

// SetIsDeleted sets field value
func (o *TestPlanWithTestSuiteTreeModel) SetIsDeleted(v bool) {
	o.IsDeleted = v
}

// GetLockedDate returns the LockedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanWithTestSuiteTreeModel) GetLockedDate() time.Time {
	if o == nil || IsNil(o.LockedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LockedDate.Get()
}

// GetLockedDateOk returns a tuple with the LockedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanWithTestSuiteTreeModel) GetLockedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LockedDate.Get(), o.LockedDate.IsSet()
}

// HasLockedDate returns a boolean if a field has been set.
func (o *TestPlanWithTestSuiteTreeModel) HasLockedDate() bool {
	if o != nil && o.LockedDate.IsSet() {
		return true
	}

	return false
}

// SetLockedDate gets a reference to the given NullableTime and assigns it to the LockedDate field.
func (o *TestPlanWithTestSuiteTreeModel) SetLockedDate(v time.Time) {
	o.LockedDate.Set(&v)
}
// SetLockedDateNil sets the value for LockedDate to be an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) SetLockedDateNil() {
	o.LockedDate.Set(nil)
}

// UnsetLockedDate ensures that no value is present for LockedDate, not even an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) UnsetLockedDate() {
	o.LockedDate.Unset()
}

// GetId returns the Id field value
func (o *TestPlanWithTestSuiteTreeModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestPlanWithTestSuiteTreeModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestPlanWithTestSuiteTreeModel) SetId(v string) {
	o.Id = v
}

// GetLockedById returns the LockedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanWithTestSuiteTreeModel) GetLockedById() string {
	if o == nil || IsNil(o.LockedById.Get()) {
		var ret string
		return ret
	}
	return *o.LockedById.Get()
}

// GetLockedByIdOk returns a tuple with the LockedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanWithTestSuiteTreeModel) GetLockedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LockedById.Get(), o.LockedById.IsSet()
}

// HasLockedById returns a boolean if a field has been set.
func (o *TestPlanWithTestSuiteTreeModel) HasLockedById() bool {
	if o != nil && o.LockedById.IsSet() {
		return true
	}

	return false
}

// SetLockedById gets a reference to the given NullableString and assigns it to the LockedById field.
func (o *TestPlanWithTestSuiteTreeModel) SetLockedById(v string) {
	o.LockedById.Set(&v)
}
// SetLockedByIdNil sets the value for LockedById to be an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) SetLockedByIdNil() {
	o.LockedById.Set(nil)
}

// UnsetLockedById ensures that no value is present for LockedById, not even an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) UnsetLockedById() {
	o.LockedById.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanWithTestSuiteTreeModel) GetTags() []TagShortModel {
	if o == nil {
		var ret []TagShortModel
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanWithTestSuiteTreeModel) GetTagsOk() ([]TagShortModel, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TestPlanWithTestSuiteTreeModel) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagShortModel and assigns it to the Tags field.
func (o *TestPlanWithTestSuiteTreeModel) SetTags(v []TagShortModel) {
	o.Tags = v
}

// GetName returns the Name field value
func (o *TestPlanWithTestSuiteTreeModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TestPlanWithTestSuiteTreeModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TestPlanWithTestSuiteTreeModel) SetName(v string) {
	o.Name = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanWithTestSuiteTreeModel) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanWithTestSuiteTreeModel) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *TestPlanWithTestSuiteTreeModel) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableTime and assigns it to the StartDate field.
func (o *TestPlanWithTestSuiteTreeModel) SetStartDate(v time.Time) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanWithTestSuiteTreeModel) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanWithTestSuiteTreeModel) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *TestPlanWithTestSuiteTreeModel) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *TestPlanWithTestSuiteTreeModel) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanWithTestSuiteTreeModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanWithTestSuiteTreeModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TestPlanWithTestSuiteTreeModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TestPlanWithTestSuiteTreeModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) UnsetDescription() {
	o.Description.Unset()
}

// GetBuild returns the Build field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanWithTestSuiteTreeModel) GetBuild() string {
	if o == nil || IsNil(o.Build.Get()) {
		var ret string
		return ret
	}
	return *o.Build.Get()
}

// GetBuildOk returns a tuple with the Build field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanWithTestSuiteTreeModel) GetBuildOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Build.Get(), o.Build.IsSet()
}

// HasBuild returns a boolean if a field has been set.
func (o *TestPlanWithTestSuiteTreeModel) HasBuild() bool {
	if o != nil && o.Build.IsSet() {
		return true
	}

	return false
}

// SetBuild gets a reference to the given NullableString and assigns it to the Build field.
func (o *TestPlanWithTestSuiteTreeModel) SetBuild(v string) {
	o.Build.Set(&v)
}
// SetBuildNil sets the value for Build to be an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) SetBuildNil() {
	o.Build.Set(nil)
}

// UnsetBuild ensures that no value is present for Build, not even an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) UnsetBuild() {
	o.Build.Unset()
}

// GetProjectId returns the ProjectId field value
func (o *TestPlanWithTestSuiteTreeModel) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *TestPlanWithTestSuiteTreeModel) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *TestPlanWithTestSuiteTreeModel) SetProjectId(v string) {
	o.ProjectId = v
}

// GetProductName returns the ProductName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanWithTestSuiteTreeModel) GetProductName() string {
	if o == nil || IsNil(o.ProductName.Get()) {
		var ret string
		return ret
	}
	return *o.ProductName.Get()
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanWithTestSuiteTreeModel) GetProductNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductName.Get(), o.ProductName.IsSet()
}

// HasProductName returns a boolean if a field has been set.
func (o *TestPlanWithTestSuiteTreeModel) HasProductName() bool {
	if o != nil && o.ProductName.IsSet() {
		return true
	}

	return false
}

// SetProductName gets a reference to the given NullableString and assigns it to the ProductName field.
func (o *TestPlanWithTestSuiteTreeModel) SetProductName(v string) {
	o.ProductName.Set(&v)
}
// SetProductNameNil sets the value for ProductName to be an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) SetProductNameNil() {
	o.ProductName.Set(nil)
}

// UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) UnsetProductName() {
	o.ProductName.Unset()
}

// GetHasAutomaticDurationTimer returns the HasAutomaticDurationTimer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPlanWithTestSuiteTreeModel) GetHasAutomaticDurationTimer() bool {
	if o == nil || IsNil(o.HasAutomaticDurationTimer.Get()) {
		var ret bool
		return ret
	}
	return *o.HasAutomaticDurationTimer.Get()
}

// GetHasAutomaticDurationTimerOk returns a tuple with the HasAutomaticDurationTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPlanWithTestSuiteTreeModel) GetHasAutomaticDurationTimerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasAutomaticDurationTimer.Get(), o.HasAutomaticDurationTimer.IsSet()
}

// HasHasAutomaticDurationTimer returns a boolean if a field has been set.
func (o *TestPlanWithTestSuiteTreeModel) HasHasAutomaticDurationTimer() bool {
	if o != nil && o.HasAutomaticDurationTimer.IsSet() {
		return true
	}

	return false
}

// SetHasAutomaticDurationTimer gets a reference to the given NullableBool and assigns it to the HasAutomaticDurationTimer field.
func (o *TestPlanWithTestSuiteTreeModel) SetHasAutomaticDurationTimer(v bool) {
	o.HasAutomaticDurationTimer.Set(&v)
}
// SetHasAutomaticDurationTimerNil sets the value for HasAutomaticDurationTimer to be an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) SetHasAutomaticDurationTimerNil() {
	o.HasAutomaticDurationTimer.Set(nil)
}

// UnsetHasAutomaticDurationTimer ensures that no value is present for HasAutomaticDurationTimer, not even an explicit nil
func (o *TestPlanWithTestSuiteTreeModel) UnsetHasAutomaticDurationTimer() {
	o.HasAutomaticDurationTimer.Unset()
}

// GetAttributes returns the Attributes field value
func (o *TestPlanWithTestSuiteTreeModel) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TestPlanWithTestSuiteTreeModel) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *TestPlanWithTestSuiteTreeModel) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o TestPlanWithTestSuiteTreeModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPlanWithTestSuiteTreeModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["testSuites"] = o.TestSuites
	toSerialize["status"] = o.Status
	if o.StartedOn.IsSet() {
		toSerialize["startedOn"] = o.StartedOn.Get()
	}
	if o.CompletedOn.IsSet() {
		toSerialize["completedOn"] = o.CompletedOn.Get()
	}
	if o.CreatedDate.IsSet() {
		toSerialize["createdDate"] = o.CreatedDate.Get()
	}
	if o.ModifiedDate.IsSet() {
		toSerialize["modifiedDate"] = o.ModifiedDate.Get()
	}
	toSerialize["createdById"] = o.CreatedById
	if o.ModifiedById.IsSet() {
		toSerialize["modifiedById"] = o.ModifiedById.Get()
	}
	toSerialize["globalId"] = o.GlobalId
	toSerialize["isDeleted"] = o.IsDeleted
	if o.LockedDate.IsSet() {
		toSerialize["lockedDate"] = o.LockedDate.Get()
	}
	toSerialize["id"] = o.Id
	if o.LockedById.IsSet() {
		toSerialize["lockedById"] = o.LockedById.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["name"] = o.Name
	if o.StartDate.IsSet() {
		toSerialize["startDate"] = o.StartDate.Get()
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Build.IsSet() {
		toSerialize["build"] = o.Build.Get()
	}
	toSerialize["projectId"] = o.ProjectId
	if o.ProductName.IsSet() {
		toSerialize["productName"] = o.ProductName.Get()
	}
	if o.HasAutomaticDurationTimer.IsSet() {
		toSerialize["hasAutomaticDurationTimer"] = o.HasAutomaticDurationTimer.Get()
	}
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

type NullableTestPlanWithTestSuiteTreeModel struct {
	value *TestPlanWithTestSuiteTreeModel
	isSet bool
}

func (v NullableTestPlanWithTestSuiteTreeModel) Get() *TestPlanWithTestSuiteTreeModel {
	return v.value
}

func (v *NullableTestPlanWithTestSuiteTreeModel) Set(val *TestPlanWithTestSuiteTreeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPlanWithTestSuiteTreeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPlanWithTestSuiteTreeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPlanWithTestSuiteTreeModel(val *TestPlanWithTestSuiteTreeModel) *NullableTestPlanWithTestSuiteTreeModel {
	return &NullableTestPlanWithTestSuiteTreeModel{value: val, isSet: true}
}

func (v NullableTestPlanWithTestSuiteTreeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPlanWithTestSuiteTreeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


