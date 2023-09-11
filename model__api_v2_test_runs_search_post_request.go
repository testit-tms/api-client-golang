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

// checks if the ApiV2TestRunsSearchPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV2TestRunsSearchPostRequest{}

// ApiV2TestRunsSearchPostRequest struct for ApiV2TestRunsSearchPostRequest
type ApiV2TestRunsSearchPostRequest struct {
	// Specifies a test run project IDs to search for
	ProjectIds []string `json:"projectIds,omitempty"`
	// Specifies test run name
	Name NullableString `json:"name,omitempty"`
	// Specifies a test run states to search for
	States []TestRunState `json:"states,omitempty"`
	StartedDate NullableTestRunFilterModelStartedDate `json:"startedDate,omitempty"`
	// Specifies a test run creator IDs to search for
	CreatedByIds []string `json:"createdByIds,omitempty"`
	// Specifies a test run last editor IDs to search for
	ModifiedByIds []string `json:"modifiedByIds,omitempty"`
	// Specifies a test run deleted status to search for
	IsDeleted NullableBool `json:"isDeleted,omitempty"`
	AutoTestsCount NullableTestRunFilterModelAutoTestsCount `json:"autoTestsCount,omitempty"`
	// Specifies test results outcomes
	TestResultsOutcome []TestResultOutcome `json:"testResultsOutcome,omitempty"`
	// Specifies failure categories
	FailureCategory []FailureCategoryModel `json:"failureCategory,omitempty"`
	CompletedDate NullableTestRunFilterModelCompletedDate `json:"completedDate,omitempty"`
}

// NewApiV2TestRunsSearchPostRequest instantiates a new ApiV2TestRunsSearchPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2TestRunsSearchPostRequest() *ApiV2TestRunsSearchPostRequest {
	this := ApiV2TestRunsSearchPostRequest{}
	return &this
}

// NewApiV2TestRunsSearchPostRequestWithDefaults instantiates a new ApiV2TestRunsSearchPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2TestRunsSearchPostRequestWithDefaults() *ApiV2TestRunsSearchPostRequest {
	this := ApiV2TestRunsSearchPostRequest{}
	return &this
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestRunsSearchPostRequest) GetProjectIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestRunsSearchPostRequest) GetProjectIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectIds) {
		return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *ApiV2TestRunsSearchPostRequest) HasProjectIds() bool {
	if o != nil && IsNil(o.ProjectIds) {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []string and assigns it to the ProjectIds field.
func (o *ApiV2TestRunsSearchPostRequest) SetProjectIds(v []string) {
	o.ProjectIds = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestRunsSearchPostRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestRunsSearchPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ApiV2TestRunsSearchPostRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ApiV2TestRunsSearchPostRequest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ApiV2TestRunsSearchPostRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ApiV2TestRunsSearchPostRequest) UnsetName() {
	o.Name.Unset()
}

// GetStates returns the States field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestRunsSearchPostRequest) GetStates() []TestRunState {
	if o == nil {
		var ret []TestRunState
		return ret
	}
	return o.States
}

// GetStatesOk returns a tuple with the States field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestRunsSearchPostRequest) GetStatesOk() ([]TestRunState, bool) {
	if o == nil || IsNil(o.States) {
		return nil, false
	}
	return o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *ApiV2TestRunsSearchPostRequest) HasStates() bool {
	if o != nil && IsNil(o.States) {
		return true
	}

	return false
}

// SetStates gets a reference to the given []TestRunState and assigns it to the States field.
func (o *ApiV2TestRunsSearchPostRequest) SetStates(v []TestRunState) {
	o.States = v
}

// GetStartedDate returns the StartedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestRunsSearchPostRequest) GetStartedDate() TestRunFilterModelStartedDate {
	if o == nil || IsNil(o.StartedDate.Get()) {
		var ret TestRunFilterModelStartedDate
		return ret
	}
	return *o.StartedDate.Get()
}

// GetStartedDateOk returns a tuple with the StartedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestRunsSearchPostRequest) GetStartedDateOk() (*TestRunFilterModelStartedDate, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedDate.Get(), o.StartedDate.IsSet()
}

// HasStartedDate returns a boolean if a field has been set.
func (o *ApiV2TestRunsSearchPostRequest) HasStartedDate() bool {
	if o != nil && o.StartedDate.IsSet() {
		return true
	}

	return false
}

// SetStartedDate gets a reference to the given NullableTestRunFilterModelStartedDate and assigns it to the StartedDate field.
func (o *ApiV2TestRunsSearchPostRequest) SetStartedDate(v TestRunFilterModelStartedDate) {
	o.StartedDate.Set(&v)
}
// SetStartedDateNil sets the value for StartedDate to be an explicit nil
func (o *ApiV2TestRunsSearchPostRequest) SetStartedDateNil() {
	o.StartedDate.Set(nil)
}

// UnsetStartedDate ensures that no value is present for StartedDate, not even an explicit nil
func (o *ApiV2TestRunsSearchPostRequest) UnsetStartedDate() {
	o.StartedDate.Unset()
}

// GetCreatedByIds returns the CreatedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestRunsSearchPostRequest) GetCreatedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CreatedByIds
}

// GetCreatedByIdsOk returns a tuple with the CreatedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestRunsSearchPostRequest) GetCreatedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CreatedByIds) {
		return nil, false
	}
	return o.CreatedByIds, true
}

// HasCreatedByIds returns a boolean if a field has been set.
func (o *ApiV2TestRunsSearchPostRequest) HasCreatedByIds() bool {
	if o != nil && IsNil(o.CreatedByIds) {
		return true
	}

	return false
}

// SetCreatedByIds gets a reference to the given []string and assigns it to the CreatedByIds field.
func (o *ApiV2TestRunsSearchPostRequest) SetCreatedByIds(v []string) {
	o.CreatedByIds = v
}

// GetModifiedByIds returns the ModifiedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestRunsSearchPostRequest) GetModifiedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ModifiedByIds
}

// GetModifiedByIdsOk returns a tuple with the ModifiedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestRunsSearchPostRequest) GetModifiedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ModifiedByIds) {
		return nil, false
	}
	return o.ModifiedByIds, true
}

// HasModifiedByIds returns a boolean if a field has been set.
func (o *ApiV2TestRunsSearchPostRequest) HasModifiedByIds() bool {
	if o != nil && IsNil(o.ModifiedByIds) {
		return true
	}

	return false
}

// SetModifiedByIds gets a reference to the given []string and assigns it to the ModifiedByIds field.
func (o *ApiV2TestRunsSearchPostRequest) SetModifiedByIds(v []string) {
	o.ModifiedByIds = v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestRunsSearchPostRequest) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDeleted.Get()
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestRunsSearchPostRequest) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDeleted.Get(), o.IsDeleted.IsSet()
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *ApiV2TestRunsSearchPostRequest) HasIsDeleted() bool {
	if o != nil && o.IsDeleted.IsSet() {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given NullableBool and assigns it to the IsDeleted field.
func (o *ApiV2TestRunsSearchPostRequest) SetIsDeleted(v bool) {
	o.IsDeleted.Set(&v)
}
// SetIsDeletedNil sets the value for IsDeleted to be an explicit nil
func (o *ApiV2TestRunsSearchPostRequest) SetIsDeletedNil() {
	o.IsDeleted.Set(nil)
}

// UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
func (o *ApiV2TestRunsSearchPostRequest) UnsetIsDeleted() {
	o.IsDeleted.Unset()
}

// GetAutoTestsCount returns the AutoTestsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestRunsSearchPostRequest) GetAutoTestsCount() TestRunFilterModelAutoTestsCount {
	if o == nil || IsNil(o.AutoTestsCount.Get()) {
		var ret TestRunFilterModelAutoTestsCount
		return ret
	}
	return *o.AutoTestsCount.Get()
}

// GetAutoTestsCountOk returns a tuple with the AutoTestsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestRunsSearchPostRequest) GetAutoTestsCountOk() (*TestRunFilterModelAutoTestsCount, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoTestsCount.Get(), o.AutoTestsCount.IsSet()
}

// HasAutoTestsCount returns a boolean if a field has been set.
func (o *ApiV2TestRunsSearchPostRequest) HasAutoTestsCount() bool {
	if o != nil && o.AutoTestsCount.IsSet() {
		return true
	}

	return false
}

// SetAutoTestsCount gets a reference to the given NullableTestRunFilterModelAutoTestsCount and assigns it to the AutoTestsCount field.
func (o *ApiV2TestRunsSearchPostRequest) SetAutoTestsCount(v TestRunFilterModelAutoTestsCount) {
	o.AutoTestsCount.Set(&v)
}
// SetAutoTestsCountNil sets the value for AutoTestsCount to be an explicit nil
func (o *ApiV2TestRunsSearchPostRequest) SetAutoTestsCountNil() {
	o.AutoTestsCount.Set(nil)
}

// UnsetAutoTestsCount ensures that no value is present for AutoTestsCount, not even an explicit nil
func (o *ApiV2TestRunsSearchPostRequest) UnsetAutoTestsCount() {
	o.AutoTestsCount.Unset()
}

// GetTestResultsOutcome returns the TestResultsOutcome field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestRunsSearchPostRequest) GetTestResultsOutcome() []TestResultOutcome {
	if o == nil {
		var ret []TestResultOutcome
		return ret
	}
	return o.TestResultsOutcome
}

// GetTestResultsOutcomeOk returns a tuple with the TestResultsOutcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestRunsSearchPostRequest) GetTestResultsOutcomeOk() ([]TestResultOutcome, bool) {
	if o == nil || IsNil(o.TestResultsOutcome) {
		return nil, false
	}
	return o.TestResultsOutcome, true
}

// HasTestResultsOutcome returns a boolean if a field has been set.
func (o *ApiV2TestRunsSearchPostRequest) HasTestResultsOutcome() bool {
	if o != nil && IsNil(o.TestResultsOutcome) {
		return true
	}

	return false
}

// SetTestResultsOutcome gets a reference to the given []TestResultOutcome and assigns it to the TestResultsOutcome field.
func (o *ApiV2TestRunsSearchPostRequest) SetTestResultsOutcome(v []TestResultOutcome) {
	o.TestResultsOutcome = v
}

// GetFailureCategory returns the FailureCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestRunsSearchPostRequest) GetFailureCategory() []FailureCategoryModel {
	if o == nil {
		var ret []FailureCategoryModel
		return ret
	}
	return o.FailureCategory
}

// GetFailureCategoryOk returns a tuple with the FailureCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestRunsSearchPostRequest) GetFailureCategoryOk() ([]FailureCategoryModel, bool) {
	if o == nil || IsNil(o.FailureCategory) {
		return nil, false
	}
	return o.FailureCategory, true
}

// HasFailureCategory returns a boolean if a field has been set.
func (o *ApiV2TestRunsSearchPostRequest) HasFailureCategory() bool {
	if o != nil && IsNil(o.FailureCategory) {
		return true
	}

	return false
}

// SetFailureCategory gets a reference to the given []FailureCategoryModel and assigns it to the FailureCategory field.
func (o *ApiV2TestRunsSearchPostRequest) SetFailureCategory(v []FailureCategoryModel) {
	o.FailureCategory = v
}

// GetCompletedDate returns the CompletedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestRunsSearchPostRequest) GetCompletedDate() TestRunFilterModelCompletedDate {
	if o == nil || IsNil(o.CompletedDate.Get()) {
		var ret TestRunFilterModelCompletedDate
		return ret
	}
	return *o.CompletedDate.Get()
}

// GetCompletedDateOk returns a tuple with the CompletedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestRunsSearchPostRequest) GetCompletedDateOk() (*TestRunFilterModelCompletedDate, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedDate.Get(), o.CompletedDate.IsSet()
}

// HasCompletedDate returns a boolean if a field has been set.
func (o *ApiV2TestRunsSearchPostRequest) HasCompletedDate() bool {
	if o != nil && o.CompletedDate.IsSet() {
		return true
	}

	return false
}

// SetCompletedDate gets a reference to the given NullableTestRunFilterModelCompletedDate and assigns it to the CompletedDate field.
func (o *ApiV2TestRunsSearchPostRequest) SetCompletedDate(v TestRunFilterModelCompletedDate) {
	o.CompletedDate.Set(&v)
}
// SetCompletedDateNil sets the value for CompletedDate to be an explicit nil
func (o *ApiV2TestRunsSearchPostRequest) SetCompletedDateNil() {
	o.CompletedDate.Set(nil)
}

// UnsetCompletedDate ensures that no value is present for CompletedDate, not even an explicit nil
func (o *ApiV2TestRunsSearchPostRequest) UnsetCompletedDate() {
	o.CompletedDate.Unset()
}

func (o ApiV2TestRunsSearchPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV2TestRunsSearchPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectIds != nil {
		toSerialize["projectIds"] = o.ProjectIds
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.States != nil {
		toSerialize["states"] = o.States
	}
	if o.StartedDate.IsSet() {
		toSerialize["startedDate"] = o.StartedDate.Get()
	}
	if o.CreatedByIds != nil {
		toSerialize["createdByIds"] = o.CreatedByIds
	}
	if o.ModifiedByIds != nil {
		toSerialize["modifiedByIds"] = o.ModifiedByIds
	}
	if o.IsDeleted.IsSet() {
		toSerialize["isDeleted"] = o.IsDeleted.Get()
	}
	if o.AutoTestsCount.IsSet() {
		toSerialize["autoTestsCount"] = o.AutoTestsCount.Get()
	}
	if o.TestResultsOutcome != nil {
		toSerialize["testResultsOutcome"] = o.TestResultsOutcome
	}
	if o.FailureCategory != nil {
		toSerialize["failureCategory"] = o.FailureCategory
	}
	if o.CompletedDate.IsSet() {
		toSerialize["completedDate"] = o.CompletedDate.Get()
	}
	return toSerialize, nil
}

type NullableApiV2TestRunsSearchPostRequest struct {
	value *ApiV2TestRunsSearchPostRequest
	isSet bool
}

func (v NullableApiV2TestRunsSearchPostRequest) Get() *ApiV2TestRunsSearchPostRequest {
	return v.value
}

func (v *NullableApiV2TestRunsSearchPostRequest) Set(val *ApiV2TestRunsSearchPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2TestRunsSearchPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2TestRunsSearchPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2TestRunsSearchPostRequest(val *ApiV2TestRunsSearchPostRequest) *NullableApiV2TestRunsSearchPostRequest {
	return &NullableApiV2TestRunsSearchPostRequest{value: val, isSet: true}
}

func (v NullableApiV2TestRunsSearchPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2TestRunsSearchPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


