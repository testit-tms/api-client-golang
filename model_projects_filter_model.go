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

// checks if the ProjectsFilterModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectsFilterModel{}

// ProjectsFilterModel struct for ProjectsFilterModel
type ProjectsFilterModel struct {
	// Specifies a project name to search for
	Name NullableString `json:"name,omitempty"`
	// Specifies a project favorite status to search for
	IsFavorite NullableBool `json:"isFavorite,omitempty"`
	// Specifies a project deleted status to search for
	IsDeleted NullableBool `json:"isDeleted,omitempty"`
	// Specifies a project range of test cases count to search for
	TestCasesCount NullableInt32RangeSelectorModel `json:"testCasesCount,omitempty"`
	// Specifies a project range of checklists count to search for
	ChecklistsCount NullableInt32RangeSelectorModel `json:"checklistsCount,omitempty"`
	// Specifies a project range of shared steps count to search for
	SharedStepsCount NullableInt32RangeSelectorModel `json:"sharedStepsCount,omitempty"`
	// Specifies a project range of autotests count to search for
	AutotestsCount NullableInt32RangeSelectorModel `json:"autotestsCount,omitempty"`
	// Specifies a project global IDs to search for
	GlobalIds []int64 `json:"globalIds,omitempty"`
	// Specifies a project range of creation date to search for
	CreatedDate NullableDateTimeRangeSelectorModel `json:"createdDate,omitempty"`
	// Specifies an autotest creator IDs to search for
	CreatedByIds []string `json:"createdByIds,omitempty"`
	// Collection of project types to search for
	Types []ProjectTypeModel `json:"types,omitempty"`
}

// NewProjectsFilterModel instantiates a new ProjectsFilterModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsFilterModel() *ProjectsFilterModel {
	this := ProjectsFilterModel{}
	return &this
}

// NewProjectsFilterModelWithDefaults instantiates a new ProjectsFilterModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsFilterModelWithDefaults() *ProjectsFilterModel {
	this := ProjectsFilterModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsFilterModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsFilterModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ProjectsFilterModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ProjectsFilterModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ProjectsFilterModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ProjectsFilterModel) UnsetName() {
	o.Name.Unset()
}

// GetIsFavorite returns the IsFavorite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsFilterModel) GetIsFavorite() bool {
	if o == nil || IsNil(o.IsFavorite.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFavorite.Get()
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsFilterModel) GetIsFavoriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFavorite.Get(), o.IsFavorite.IsSet()
}

// HasIsFavorite returns a boolean if a field has been set.
func (o *ProjectsFilterModel) HasIsFavorite() bool {
	if o != nil && o.IsFavorite.IsSet() {
		return true
	}

	return false
}

// SetIsFavorite gets a reference to the given NullableBool and assigns it to the IsFavorite field.
func (o *ProjectsFilterModel) SetIsFavorite(v bool) {
	o.IsFavorite.Set(&v)
}
// SetIsFavoriteNil sets the value for IsFavorite to be an explicit nil
func (o *ProjectsFilterModel) SetIsFavoriteNil() {
	o.IsFavorite.Set(nil)
}

// UnsetIsFavorite ensures that no value is present for IsFavorite, not even an explicit nil
func (o *ProjectsFilterModel) UnsetIsFavorite() {
	o.IsFavorite.Unset()
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsFilterModel) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDeleted.Get()
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsFilterModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDeleted.Get(), o.IsDeleted.IsSet()
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *ProjectsFilterModel) HasIsDeleted() bool {
	if o != nil && o.IsDeleted.IsSet() {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given NullableBool and assigns it to the IsDeleted field.
func (o *ProjectsFilterModel) SetIsDeleted(v bool) {
	o.IsDeleted.Set(&v)
}
// SetIsDeletedNil sets the value for IsDeleted to be an explicit nil
func (o *ProjectsFilterModel) SetIsDeletedNil() {
	o.IsDeleted.Set(nil)
}

// UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
func (o *ProjectsFilterModel) UnsetIsDeleted() {
	o.IsDeleted.Unset()
}

// GetTestCasesCount returns the TestCasesCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsFilterModel) GetTestCasesCount() Int32RangeSelectorModel {
	if o == nil || IsNil(o.TestCasesCount.Get()) {
		var ret Int32RangeSelectorModel
		return ret
	}
	return *o.TestCasesCount.Get()
}

// GetTestCasesCountOk returns a tuple with the TestCasesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsFilterModel) GetTestCasesCountOk() (*Int32RangeSelectorModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestCasesCount.Get(), o.TestCasesCount.IsSet()
}

// HasTestCasesCount returns a boolean if a field has been set.
func (o *ProjectsFilterModel) HasTestCasesCount() bool {
	if o != nil && o.TestCasesCount.IsSet() {
		return true
	}

	return false
}

// SetTestCasesCount gets a reference to the given NullableInt32RangeSelectorModel and assigns it to the TestCasesCount field.
func (o *ProjectsFilterModel) SetTestCasesCount(v Int32RangeSelectorModel) {
	o.TestCasesCount.Set(&v)
}
// SetTestCasesCountNil sets the value for TestCasesCount to be an explicit nil
func (o *ProjectsFilterModel) SetTestCasesCountNil() {
	o.TestCasesCount.Set(nil)
}

// UnsetTestCasesCount ensures that no value is present for TestCasesCount, not even an explicit nil
func (o *ProjectsFilterModel) UnsetTestCasesCount() {
	o.TestCasesCount.Unset()
}

// GetChecklistsCount returns the ChecklistsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsFilterModel) GetChecklistsCount() Int32RangeSelectorModel {
	if o == nil || IsNil(o.ChecklistsCount.Get()) {
		var ret Int32RangeSelectorModel
		return ret
	}
	return *o.ChecklistsCount.Get()
}

// GetChecklistsCountOk returns a tuple with the ChecklistsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsFilterModel) GetChecklistsCountOk() (*Int32RangeSelectorModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChecklistsCount.Get(), o.ChecklistsCount.IsSet()
}

// HasChecklistsCount returns a boolean if a field has been set.
func (o *ProjectsFilterModel) HasChecklistsCount() bool {
	if o != nil && o.ChecklistsCount.IsSet() {
		return true
	}

	return false
}

// SetChecklistsCount gets a reference to the given NullableInt32RangeSelectorModel and assigns it to the ChecklistsCount field.
func (o *ProjectsFilterModel) SetChecklistsCount(v Int32RangeSelectorModel) {
	o.ChecklistsCount.Set(&v)
}
// SetChecklistsCountNil sets the value for ChecklistsCount to be an explicit nil
func (o *ProjectsFilterModel) SetChecklistsCountNil() {
	o.ChecklistsCount.Set(nil)
}

// UnsetChecklistsCount ensures that no value is present for ChecklistsCount, not even an explicit nil
func (o *ProjectsFilterModel) UnsetChecklistsCount() {
	o.ChecklistsCount.Unset()
}

// GetSharedStepsCount returns the SharedStepsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsFilterModel) GetSharedStepsCount() Int32RangeSelectorModel {
	if o == nil || IsNil(o.SharedStepsCount.Get()) {
		var ret Int32RangeSelectorModel
		return ret
	}
	return *o.SharedStepsCount.Get()
}

// GetSharedStepsCountOk returns a tuple with the SharedStepsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsFilterModel) GetSharedStepsCountOk() (*Int32RangeSelectorModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedStepsCount.Get(), o.SharedStepsCount.IsSet()
}

// HasSharedStepsCount returns a boolean if a field has been set.
func (o *ProjectsFilterModel) HasSharedStepsCount() bool {
	if o != nil && o.SharedStepsCount.IsSet() {
		return true
	}

	return false
}

// SetSharedStepsCount gets a reference to the given NullableInt32RangeSelectorModel and assigns it to the SharedStepsCount field.
func (o *ProjectsFilterModel) SetSharedStepsCount(v Int32RangeSelectorModel) {
	o.SharedStepsCount.Set(&v)
}
// SetSharedStepsCountNil sets the value for SharedStepsCount to be an explicit nil
func (o *ProjectsFilterModel) SetSharedStepsCountNil() {
	o.SharedStepsCount.Set(nil)
}

// UnsetSharedStepsCount ensures that no value is present for SharedStepsCount, not even an explicit nil
func (o *ProjectsFilterModel) UnsetSharedStepsCount() {
	o.SharedStepsCount.Unset()
}

// GetAutotestsCount returns the AutotestsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsFilterModel) GetAutotestsCount() Int32RangeSelectorModel {
	if o == nil || IsNil(o.AutotestsCount.Get()) {
		var ret Int32RangeSelectorModel
		return ret
	}
	return *o.AutotestsCount.Get()
}

// GetAutotestsCountOk returns a tuple with the AutotestsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsFilterModel) GetAutotestsCountOk() (*Int32RangeSelectorModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutotestsCount.Get(), o.AutotestsCount.IsSet()
}

// HasAutotestsCount returns a boolean if a field has been set.
func (o *ProjectsFilterModel) HasAutotestsCount() bool {
	if o != nil && o.AutotestsCount.IsSet() {
		return true
	}

	return false
}

// SetAutotestsCount gets a reference to the given NullableInt32RangeSelectorModel and assigns it to the AutotestsCount field.
func (o *ProjectsFilterModel) SetAutotestsCount(v Int32RangeSelectorModel) {
	o.AutotestsCount.Set(&v)
}
// SetAutotestsCountNil sets the value for AutotestsCount to be an explicit nil
func (o *ProjectsFilterModel) SetAutotestsCountNil() {
	o.AutotestsCount.Set(nil)
}

// UnsetAutotestsCount ensures that no value is present for AutotestsCount, not even an explicit nil
func (o *ProjectsFilterModel) UnsetAutotestsCount() {
	o.AutotestsCount.Unset()
}

// GetGlobalIds returns the GlobalIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsFilterModel) GetGlobalIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.GlobalIds
}

// GetGlobalIdsOk returns a tuple with the GlobalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsFilterModel) GetGlobalIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.GlobalIds) {
		return nil, false
	}
	return o.GlobalIds, true
}

// HasGlobalIds returns a boolean if a field has been set.
func (o *ProjectsFilterModel) HasGlobalIds() bool {
	if o != nil && !IsNil(o.GlobalIds) {
		return true
	}

	return false
}

// SetGlobalIds gets a reference to the given []int64 and assigns it to the GlobalIds field.
func (o *ProjectsFilterModel) SetGlobalIds(v []int64) {
	o.GlobalIds = v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsFilterModel) GetCreatedDate() DateTimeRangeSelectorModel {
	if o == nil || IsNil(o.CreatedDate.Get()) {
		var ret DateTimeRangeSelectorModel
		return ret
	}
	return *o.CreatedDate.Get()
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsFilterModel) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDate.Get(), o.CreatedDate.IsSet()
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ProjectsFilterModel) HasCreatedDate() bool {
	if o != nil && o.CreatedDate.IsSet() {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given NullableDateTimeRangeSelectorModel and assigns it to the CreatedDate field.
func (o *ProjectsFilterModel) SetCreatedDate(v DateTimeRangeSelectorModel) {
	o.CreatedDate.Set(&v)
}
// SetCreatedDateNil sets the value for CreatedDate to be an explicit nil
func (o *ProjectsFilterModel) SetCreatedDateNil() {
	o.CreatedDate.Set(nil)
}

// UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
func (o *ProjectsFilterModel) UnsetCreatedDate() {
	o.CreatedDate.Unset()
}

// GetCreatedByIds returns the CreatedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsFilterModel) GetCreatedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CreatedByIds
}

// GetCreatedByIdsOk returns a tuple with the CreatedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsFilterModel) GetCreatedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CreatedByIds) {
		return nil, false
	}
	return o.CreatedByIds, true
}

// HasCreatedByIds returns a boolean if a field has been set.
func (o *ProjectsFilterModel) HasCreatedByIds() bool {
	if o != nil && !IsNil(o.CreatedByIds) {
		return true
	}

	return false
}

// SetCreatedByIds gets a reference to the given []string and assigns it to the CreatedByIds field.
func (o *ProjectsFilterModel) SetCreatedByIds(v []string) {
	o.CreatedByIds = v
}

// GetTypes returns the Types field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsFilterModel) GetTypes() []ProjectTypeModel {
	if o == nil {
		var ret []ProjectTypeModel
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsFilterModel) GetTypesOk() ([]ProjectTypeModel, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *ProjectsFilterModel) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []ProjectTypeModel and assigns it to the Types field.
func (o *ProjectsFilterModel) SetTypes(v []ProjectTypeModel) {
	o.Types = v
}

func (o ProjectsFilterModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectsFilterModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.IsFavorite.IsSet() {
		toSerialize["isFavorite"] = o.IsFavorite.Get()
	}
	if o.IsDeleted.IsSet() {
		toSerialize["isDeleted"] = o.IsDeleted.Get()
	}
	if o.TestCasesCount.IsSet() {
		toSerialize["testCasesCount"] = o.TestCasesCount.Get()
	}
	if o.ChecklistsCount.IsSet() {
		toSerialize["checklistsCount"] = o.ChecklistsCount.Get()
	}
	if o.SharedStepsCount.IsSet() {
		toSerialize["sharedStepsCount"] = o.SharedStepsCount.Get()
	}
	if o.AutotestsCount.IsSet() {
		toSerialize["autotestsCount"] = o.AutotestsCount.Get()
	}
	if o.GlobalIds != nil {
		toSerialize["globalIds"] = o.GlobalIds
	}
	if o.CreatedDate.IsSet() {
		toSerialize["createdDate"] = o.CreatedDate.Get()
	}
	if o.CreatedByIds != nil {
		toSerialize["createdByIds"] = o.CreatedByIds
	}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	return toSerialize, nil
}

type NullableProjectsFilterModel struct {
	value *ProjectsFilterModel
	isSet bool
}

func (v NullableProjectsFilterModel) Get() *ProjectsFilterModel {
	return v.value
}

func (v *NullableProjectsFilterModel) Set(val *ProjectsFilterModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsFilterModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsFilterModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsFilterModel(val *ProjectsFilterModel) *NullableProjectsFilterModel {
	return &NullableProjectsFilterModel{value: val, isSet: true}
}

func (v NullableProjectsFilterModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsFilterModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


