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

// checks if the ApiV2TestPointsSearchPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV2TestPointsSearchPostRequest{}

// ApiV2TestPointsSearchPostRequest struct for ApiV2TestPointsSearchPostRequest
type ApiV2TestPointsSearchPostRequest struct {
	// Specifies a test point test plan IDS to search for
	TestPlanIds []string `json:"testPlanIds,omitempty"`
	// Specifies a test point test suite IDs to search for
	TestSuiteIds []string `json:"testSuiteIds,omitempty"`
	// Specifies a test point work item global IDs to search for
	WorkItemGlobalIds []int64 `json:"workItemGlobalIds,omitempty"`
	WorkItemMedianDuration NullableTestPointFilterModelWorkItemMedianDuration `json:"workItemMedianDuration,omitempty"`
	// Specifies a test point statuses to search for
	Statuses []TestPointStatus `json:"statuses,omitempty"`
	// Specifies a test point priorities to search for
	Priorities []WorkItemPriorityModel `json:"priorities,omitempty"`
	// Specifies a test point automation status to search for
	IsAutomated NullableBool `json:"isAutomated,omitempty"`
	// Specifies a test point name to search for
	Name NullableString `json:"name,omitempty"`
	// Specifies a test point configuration IDs to search for
	ConfigurationIds []string `json:"configurationIds,omitempty"`
	// Specifies a test point assigned user IDs to search for
	TesterIds []string `json:"testerIds,omitempty"`
	Duration NullableTestPointFilterModelDuration `json:"duration,omitempty"`
	// Specifies a test point work item section IDs to search for
	SectionIds []string `json:"sectionIds,omitempty"`
	CreatedDate NullableTestPointFilterModelCreatedDate `json:"createdDate,omitempty"`
	// Specifies a test point creator IDs to search for
	CreatedByIds []string `json:"createdByIds,omitempty"`
	ModifiedDate NullableTestPointFilterModelModifiedDate `json:"modifiedDate,omitempty"`
	// Specifies a test point last editor IDs to search for
	ModifiedByIds []string `json:"modifiedByIds,omitempty"`
	// Specifies a test point tags to search for
	Tags []string `json:"tags,omitempty"`
	// Specifies a test point attributes to search for
	Attributes map[string][]string `json:"attributes,omitempty"`
	WorkItemCreatedDate NullableTestPointFilterModelWorkItemCreatedDate `json:"workItemCreatedDate,omitempty"`
	// Specifies a work item creator IDs to search for
	WorkItemCreatedByIds []string `json:"workItemCreatedByIds,omitempty"`
	WorkItemModifiedDate NullableTestPointFilterModelWorkItemModifiedDate `json:"workItemModifiedDate,omitempty"`
	// Specifies a work item last editor IDs to search for
	WorkItemModifiedByIds []string `json:"workItemModifiedByIds,omitempty"`
}

// NewApiV2TestPointsSearchPostRequest instantiates a new ApiV2TestPointsSearchPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2TestPointsSearchPostRequest() *ApiV2TestPointsSearchPostRequest {
	this := ApiV2TestPointsSearchPostRequest{}
	return &this
}

// NewApiV2TestPointsSearchPostRequestWithDefaults instantiates a new ApiV2TestPointsSearchPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2TestPointsSearchPostRequestWithDefaults() *ApiV2TestPointsSearchPostRequest {
	this := ApiV2TestPointsSearchPostRequest{}
	return &this
}

// GetTestPlanIds returns the TestPlanIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetTestPlanIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TestPlanIds
}

// GetTestPlanIdsOk returns a tuple with the TestPlanIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetTestPlanIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TestPlanIds) {
		return nil, false
	}
	return o.TestPlanIds, true
}

// HasTestPlanIds returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasTestPlanIds() bool {
	if o != nil && IsNil(o.TestPlanIds) {
		return true
	}

	return false
}

// SetTestPlanIds gets a reference to the given []string and assigns it to the TestPlanIds field.
func (o *ApiV2TestPointsSearchPostRequest) SetTestPlanIds(v []string) {
	o.TestPlanIds = v
}

// GetTestSuiteIds returns the TestSuiteIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetTestSuiteIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TestSuiteIds
}

// GetTestSuiteIdsOk returns a tuple with the TestSuiteIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetTestSuiteIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TestSuiteIds) {
		return nil, false
	}
	return o.TestSuiteIds, true
}

// HasTestSuiteIds returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasTestSuiteIds() bool {
	if o != nil && IsNil(o.TestSuiteIds) {
		return true
	}

	return false
}

// SetTestSuiteIds gets a reference to the given []string and assigns it to the TestSuiteIds field.
func (o *ApiV2TestPointsSearchPostRequest) SetTestSuiteIds(v []string) {
	o.TestSuiteIds = v
}

// GetWorkItemGlobalIds returns the WorkItemGlobalIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemGlobalIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.WorkItemGlobalIds
}

// GetWorkItemGlobalIdsOk returns a tuple with the WorkItemGlobalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemGlobalIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.WorkItemGlobalIds) {
		return nil, false
	}
	return o.WorkItemGlobalIds, true
}

// HasWorkItemGlobalIds returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasWorkItemGlobalIds() bool {
	if o != nil && IsNil(o.WorkItemGlobalIds) {
		return true
	}

	return false
}

// SetWorkItemGlobalIds gets a reference to the given []int64 and assigns it to the WorkItemGlobalIds field.
func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemGlobalIds(v []int64) {
	o.WorkItemGlobalIds = v
}

// GetWorkItemMedianDuration returns the WorkItemMedianDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemMedianDuration() TestPointFilterModelWorkItemMedianDuration {
	if o == nil || IsNil(o.WorkItemMedianDuration.Get()) {
		var ret TestPointFilterModelWorkItemMedianDuration
		return ret
	}
	return *o.WorkItemMedianDuration.Get()
}

// GetWorkItemMedianDurationOk returns a tuple with the WorkItemMedianDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemMedianDurationOk() (*TestPointFilterModelWorkItemMedianDuration, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkItemMedianDuration.Get(), o.WorkItemMedianDuration.IsSet()
}

// HasWorkItemMedianDuration returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasWorkItemMedianDuration() bool {
	if o != nil && o.WorkItemMedianDuration.IsSet() {
		return true
	}

	return false
}

// SetWorkItemMedianDuration gets a reference to the given NullableTestPointFilterModelWorkItemMedianDuration and assigns it to the WorkItemMedianDuration field.
func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemMedianDuration(v TestPointFilterModelWorkItemMedianDuration) {
	o.WorkItemMedianDuration.Set(&v)
}
// SetWorkItemMedianDurationNil sets the value for WorkItemMedianDuration to be an explicit nil
func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemMedianDurationNil() {
	o.WorkItemMedianDuration.Set(nil)
}

// UnsetWorkItemMedianDuration ensures that no value is present for WorkItemMedianDuration, not even an explicit nil
func (o *ApiV2TestPointsSearchPostRequest) UnsetWorkItemMedianDuration() {
	o.WorkItemMedianDuration.Unset()
}

// GetStatuses returns the Statuses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetStatuses() []TestPointStatus {
	if o == nil {
		var ret []TestPointStatus
		return ret
	}
	return o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetStatusesOk() ([]TestPointStatus, bool) {
	if o == nil || IsNil(o.Statuses) {
		return nil, false
	}
	return o.Statuses, true
}

// HasStatuses returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasStatuses() bool {
	if o != nil && IsNil(o.Statuses) {
		return true
	}

	return false
}

// SetStatuses gets a reference to the given []TestPointStatus and assigns it to the Statuses field.
func (o *ApiV2TestPointsSearchPostRequest) SetStatuses(v []TestPointStatus) {
	o.Statuses = v
}

// GetPriorities returns the Priorities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetPriorities() []WorkItemPriorityModel {
	if o == nil {
		var ret []WorkItemPriorityModel
		return ret
	}
	return o.Priorities
}

// GetPrioritiesOk returns a tuple with the Priorities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetPrioritiesOk() ([]WorkItemPriorityModel, bool) {
	if o == nil || IsNil(o.Priorities) {
		return nil, false
	}
	return o.Priorities, true
}

// HasPriorities returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasPriorities() bool {
	if o != nil && IsNil(o.Priorities) {
		return true
	}

	return false
}

// SetPriorities gets a reference to the given []WorkItemPriorityModel and assigns it to the Priorities field.
func (o *ApiV2TestPointsSearchPostRequest) SetPriorities(v []WorkItemPriorityModel) {
	o.Priorities = v
}

// GetIsAutomated returns the IsAutomated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetIsAutomated() bool {
	if o == nil || IsNil(o.IsAutomated.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAutomated.Get()
}

// GetIsAutomatedOk returns a tuple with the IsAutomated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetIsAutomatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAutomated.Get(), o.IsAutomated.IsSet()
}

// HasIsAutomated returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasIsAutomated() bool {
	if o != nil && o.IsAutomated.IsSet() {
		return true
	}

	return false
}

// SetIsAutomated gets a reference to the given NullableBool and assigns it to the IsAutomated field.
func (o *ApiV2TestPointsSearchPostRequest) SetIsAutomated(v bool) {
	o.IsAutomated.Set(&v)
}
// SetIsAutomatedNil sets the value for IsAutomated to be an explicit nil
func (o *ApiV2TestPointsSearchPostRequest) SetIsAutomatedNil() {
	o.IsAutomated.Set(nil)
}

// UnsetIsAutomated ensures that no value is present for IsAutomated, not even an explicit nil
func (o *ApiV2TestPointsSearchPostRequest) UnsetIsAutomated() {
	o.IsAutomated.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ApiV2TestPointsSearchPostRequest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ApiV2TestPointsSearchPostRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ApiV2TestPointsSearchPostRequest) UnsetName() {
	o.Name.Unset()
}

// GetConfigurationIds returns the ConfigurationIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetConfigurationIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ConfigurationIds
}

// GetConfigurationIdsOk returns a tuple with the ConfigurationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetConfigurationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ConfigurationIds) {
		return nil, false
	}
	return o.ConfigurationIds, true
}

// HasConfigurationIds returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasConfigurationIds() bool {
	if o != nil && IsNil(o.ConfigurationIds) {
		return true
	}

	return false
}

// SetConfigurationIds gets a reference to the given []string and assigns it to the ConfigurationIds field.
func (o *ApiV2TestPointsSearchPostRequest) SetConfigurationIds(v []string) {
	o.ConfigurationIds = v
}

// GetTesterIds returns the TesterIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetTesterIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TesterIds
}

// GetTesterIdsOk returns a tuple with the TesterIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetTesterIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TesterIds) {
		return nil, false
	}
	return o.TesterIds, true
}

// HasTesterIds returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasTesterIds() bool {
	if o != nil && IsNil(o.TesterIds) {
		return true
	}

	return false
}

// SetTesterIds gets a reference to the given []string and assigns it to the TesterIds field.
func (o *ApiV2TestPointsSearchPostRequest) SetTesterIds(v []string) {
	o.TesterIds = v
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetDuration() TestPointFilterModelDuration {
	if o == nil || IsNil(o.Duration.Get()) {
		var ret TestPointFilterModelDuration
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetDurationOk() (*TestPointFilterModelDuration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableTestPointFilterModelDuration and assigns it to the Duration field.
func (o *ApiV2TestPointsSearchPostRequest) SetDuration(v TestPointFilterModelDuration) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *ApiV2TestPointsSearchPostRequest) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *ApiV2TestPointsSearchPostRequest) UnsetDuration() {
	o.Duration.Unset()
}

// GetSectionIds returns the SectionIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetSectionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SectionIds
}

// GetSectionIdsOk returns a tuple with the SectionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetSectionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SectionIds) {
		return nil, false
	}
	return o.SectionIds, true
}

// HasSectionIds returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasSectionIds() bool {
	if o != nil && IsNil(o.SectionIds) {
		return true
	}

	return false
}

// SetSectionIds gets a reference to the given []string and assigns it to the SectionIds field.
func (o *ApiV2TestPointsSearchPostRequest) SetSectionIds(v []string) {
	o.SectionIds = v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetCreatedDate() TestPointFilterModelCreatedDate {
	if o == nil || IsNil(o.CreatedDate.Get()) {
		var ret TestPointFilterModelCreatedDate
		return ret
	}
	return *o.CreatedDate.Get()
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetCreatedDateOk() (*TestPointFilterModelCreatedDate, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDate.Get(), o.CreatedDate.IsSet()
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasCreatedDate() bool {
	if o != nil && o.CreatedDate.IsSet() {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given NullableTestPointFilterModelCreatedDate and assigns it to the CreatedDate field.
func (o *ApiV2TestPointsSearchPostRequest) SetCreatedDate(v TestPointFilterModelCreatedDate) {
	o.CreatedDate.Set(&v)
}
// SetCreatedDateNil sets the value for CreatedDate to be an explicit nil
func (o *ApiV2TestPointsSearchPostRequest) SetCreatedDateNil() {
	o.CreatedDate.Set(nil)
}

// UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
func (o *ApiV2TestPointsSearchPostRequest) UnsetCreatedDate() {
	o.CreatedDate.Unset()
}

// GetCreatedByIds returns the CreatedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetCreatedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CreatedByIds
}

// GetCreatedByIdsOk returns a tuple with the CreatedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetCreatedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CreatedByIds) {
		return nil, false
	}
	return o.CreatedByIds, true
}

// HasCreatedByIds returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasCreatedByIds() bool {
	if o != nil && IsNil(o.CreatedByIds) {
		return true
	}

	return false
}

// SetCreatedByIds gets a reference to the given []string and assigns it to the CreatedByIds field.
func (o *ApiV2TestPointsSearchPostRequest) SetCreatedByIds(v []string) {
	o.CreatedByIds = v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetModifiedDate() TestPointFilterModelModifiedDate {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret TestPointFilterModelModifiedDate
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetModifiedDateOk() (*TestPointFilterModelModifiedDate, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTestPointFilterModelModifiedDate and assigns it to the ModifiedDate field.
func (o *ApiV2TestPointsSearchPostRequest) SetModifiedDate(v TestPointFilterModelModifiedDate) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *ApiV2TestPointsSearchPostRequest) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *ApiV2TestPointsSearchPostRequest) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetModifiedByIds returns the ModifiedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetModifiedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ModifiedByIds
}

// GetModifiedByIdsOk returns a tuple with the ModifiedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetModifiedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ModifiedByIds) {
		return nil, false
	}
	return o.ModifiedByIds, true
}

// HasModifiedByIds returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasModifiedByIds() bool {
	if o != nil && IsNil(o.ModifiedByIds) {
		return true
	}

	return false
}

// SetModifiedByIds gets a reference to the given []string and assigns it to the ModifiedByIds field.
func (o *ApiV2TestPointsSearchPostRequest) SetModifiedByIds(v []string) {
	o.ModifiedByIds = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ApiV2TestPointsSearchPostRequest) SetTags(v []string) {
	o.Tags = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetAttributes() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetAttributesOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasAttributes() bool {
	if o != nil && IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string][]string and assigns it to the Attributes field.
func (o *ApiV2TestPointsSearchPostRequest) SetAttributes(v map[string][]string) {
	o.Attributes = v
}

// GetWorkItemCreatedDate returns the WorkItemCreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemCreatedDate() TestPointFilterModelWorkItemCreatedDate {
	if o == nil || IsNil(o.WorkItemCreatedDate.Get()) {
		var ret TestPointFilterModelWorkItemCreatedDate
		return ret
	}
	return *o.WorkItemCreatedDate.Get()
}

// GetWorkItemCreatedDateOk returns a tuple with the WorkItemCreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemCreatedDateOk() (*TestPointFilterModelWorkItemCreatedDate, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkItemCreatedDate.Get(), o.WorkItemCreatedDate.IsSet()
}

// HasWorkItemCreatedDate returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasWorkItemCreatedDate() bool {
	if o != nil && o.WorkItemCreatedDate.IsSet() {
		return true
	}

	return false
}

// SetWorkItemCreatedDate gets a reference to the given NullableTestPointFilterModelWorkItemCreatedDate and assigns it to the WorkItemCreatedDate field.
func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemCreatedDate(v TestPointFilterModelWorkItemCreatedDate) {
	o.WorkItemCreatedDate.Set(&v)
}
// SetWorkItemCreatedDateNil sets the value for WorkItemCreatedDate to be an explicit nil
func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemCreatedDateNil() {
	o.WorkItemCreatedDate.Set(nil)
}

// UnsetWorkItemCreatedDate ensures that no value is present for WorkItemCreatedDate, not even an explicit nil
func (o *ApiV2TestPointsSearchPostRequest) UnsetWorkItemCreatedDate() {
	o.WorkItemCreatedDate.Unset()
}

// GetWorkItemCreatedByIds returns the WorkItemCreatedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemCreatedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.WorkItemCreatedByIds
}

// GetWorkItemCreatedByIdsOk returns a tuple with the WorkItemCreatedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemCreatedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.WorkItemCreatedByIds) {
		return nil, false
	}
	return o.WorkItemCreatedByIds, true
}

// HasWorkItemCreatedByIds returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasWorkItemCreatedByIds() bool {
	if o != nil && IsNil(o.WorkItemCreatedByIds) {
		return true
	}

	return false
}

// SetWorkItemCreatedByIds gets a reference to the given []string and assigns it to the WorkItemCreatedByIds field.
func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemCreatedByIds(v []string) {
	o.WorkItemCreatedByIds = v
}

// GetWorkItemModifiedDate returns the WorkItemModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemModifiedDate() TestPointFilterModelWorkItemModifiedDate {
	if o == nil || IsNil(o.WorkItemModifiedDate.Get()) {
		var ret TestPointFilterModelWorkItemModifiedDate
		return ret
	}
	return *o.WorkItemModifiedDate.Get()
}

// GetWorkItemModifiedDateOk returns a tuple with the WorkItemModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemModifiedDateOk() (*TestPointFilterModelWorkItemModifiedDate, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkItemModifiedDate.Get(), o.WorkItemModifiedDate.IsSet()
}

// HasWorkItemModifiedDate returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasWorkItemModifiedDate() bool {
	if o != nil && o.WorkItemModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetWorkItemModifiedDate gets a reference to the given NullableTestPointFilterModelWorkItemModifiedDate and assigns it to the WorkItemModifiedDate field.
func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemModifiedDate(v TestPointFilterModelWorkItemModifiedDate) {
	o.WorkItemModifiedDate.Set(&v)
}
// SetWorkItemModifiedDateNil sets the value for WorkItemModifiedDate to be an explicit nil
func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemModifiedDateNil() {
	o.WorkItemModifiedDate.Set(nil)
}

// UnsetWorkItemModifiedDate ensures that no value is present for WorkItemModifiedDate, not even an explicit nil
func (o *ApiV2TestPointsSearchPostRequest) UnsetWorkItemModifiedDate() {
	o.WorkItemModifiedDate.Unset()
}

// GetWorkItemModifiedByIds returns the WorkItemModifiedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemModifiedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.WorkItemModifiedByIds
}

// GetWorkItemModifiedByIdsOk returns a tuple with the WorkItemModifiedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemModifiedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.WorkItemModifiedByIds) {
		return nil, false
	}
	return o.WorkItemModifiedByIds, true
}

// HasWorkItemModifiedByIds returns a boolean if a field has been set.
func (o *ApiV2TestPointsSearchPostRequest) HasWorkItemModifiedByIds() bool {
	if o != nil && IsNil(o.WorkItemModifiedByIds) {
		return true
	}

	return false
}

// SetWorkItemModifiedByIds gets a reference to the given []string and assigns it to the WorkItemModifiedByIds field.
func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemModifiedByIds(v []string) {
	o.WorkItemModifiedByIds = v
}

func (o ApiV2TestPointsSearchPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV2TestPointsSearchPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TestPlanIds != nil {
		toSerialize["testPlanIds"] = o.TestPlanIds
	}
	if o.TestSuiteIds != nil {
		toSerialize["testSuiteIds"] = o.TestSuiteIds
	}
	if o.WorkItemGlobalIds != nil {
		toSerialize["workItemGlobalIds"] = o.WorkItemGlobalIds
	}
	if o.WorkItemMedianDuration.IsSet() {
		toSerialize["workItemMedianDuration"] = o.WorkItemMedianDuration.Get()
	}
	if o.Statuses != nil {
		toSerialize["statuses"] = o.Statuses
	}
	if o.Priorities != nil {
		toSerialize["priorities"] = o.Priorities
	}
	if o.IsAutomated.IsSet() {
		toSerialize["isAutomated"] = o.IsAutomated.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ConfigurationIds != nil {
		toSerialize["configurationIds"] = o.ConfigurationIds
	}
	if o.TesterIds != nil {
		toSerialize["testerIds"] = o.TesterIds
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if o.SectionIds != nil {
		toSerialize["sectionIds"] = o.SectionIds
	}
	if o.CreatedDate.IsSet() {
		toSerialize["createdDate"] = o.CreatedDate.Get()
	}
	if o.CreatedByIds != nil {
		toSerialize["createdByIds"] = o.CreatedByIds
	}
	if o.ModifiedDate.IsSet() {
		toSerialize["modifiedDate"] = o.ModifiedDate.Get()
	}
	if o.ModifiedByIds != nil {
		toSerialize["modifiedByIds"] = o.ModifiedByIds
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.WorkItemCreatedDate.IsSet() {
		toSerialize["workItemCreatedDate"] = o.WorkItemCreatedDate.Get()
	}
	if o.WorkItemCreatedByIds != nil {
		toSerialize["workItemCreatedByIds"] = o.WorkItemCreatedByIds
	}
	if o.WorkItemModifiedDate.IsSet() {
		toSerialize["workItemModifiedDate"] = o.WorkItemModifiedDate.Get()
	}
	if o.WorkItemModifiedByIds != nil {
		toSerialize["workItemModifiedByIds"] = o.WorkItemModifiedByIds
	}
	return toSerialize, nil
}

type NullableApiV2TestPointsSearchPostRequest struct {
	value *ApiV2TestPointsSearchPostRequest
	isSet bool
}

func (v NullableApiV2TestPointsSearchPostRequest) Get() *ApiV2TestPointsSearchPostRequest {
	return v.value
}

func (v *NullableApiV2TestPointsSearchPostRequest) Set(val *ApiV2TestPointsSearchPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2TestPointsSearchPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2TestPointsSearchPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2TestPointsSearchPostRequest(val *ApiV2TestPointsSearchPostRequest) *NullableApiV2TestPointsSearchPostRequest {
	return &NullableApiV2TestPointsSearchPostRequest{value: val, isSet: true}
}

func (v NullableApiV2TestPointsSearchPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2TestPointsSearchPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

