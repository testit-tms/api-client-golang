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

// checks if the FilterModelData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilterModelData{}

// FilterModelData struct for FilterModelData
type FilterModelData struct {
	// Collection of project identifiers
	ProjectIds []string `json:"projectIds,omitempty"`
	// Name of work item
	Name NullableString `json:"name,omitempty"`
	// Specifies a work item unique IDs to search for
	Ids []string `json:"ids,omitempty"`
	// Collection of global (integer) identifiers
	GlobalIds []int64 `json:"globalIds,omitempty"`
	// Custom attributes of work item
	Attributes map[string][]string `json:"attributes,omitempty"`
	// Is result must consist of only actual/deleted work items
	IsDeleted NullableBool `json:"isDeleted,omitempty"`
	// Collection of section identifiers
	SectionIds []string `json:"sectionIds,omitempty"`
	// Collection of identifiers of users who created work item
	CreatedByIds []string `json:"createdByIds,omitempty"`
	// Collection of identifiers of users who applied last modification to work item
	ModifiedByIds []string `json:"modifiedByIds,omitempty"`
	// Collection of states of work item
	States []WorkItemStates `json:"states,omitempty"`
	// Collection of priorities of work item
	Priorities []WorkItemPriorityModel `json:"priorities,omitempty"`
	// Collection of types of work item
	Types []WorkItemEntityTypes `json:"types,omitempty"`
	CreatedDate NullableTestPointFilterModelWorkItemCreatedDate `json:"createdDate,omitempty"`
	ModifiedDate NullableTestPointFilterModelWorkItemModifiedDate `json:"modifiedDate,omitempty"`
	Duration NullableTestSuiteWorkItemsSearchModelDuration `json:"duration,omitempty"`
	MedianDuration NullableTestSuiteWorkItemsSearchModelMedianDuration `json:"medianDuration,omitempty"`
	// Is result must consist of only manual/automated work items
	IsAutomated NullableBool `json:"isAutomated,omitempty"`
	// Collection of tags
	Tags []string `json:"tags,omitempty"`
	// Collection of identifiers of linked autotests
	AutoTestIds []string `json:"autoTestIds,omitempty"`
}

// NewFilterModelData instantiates a new FilterModelData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilterModelData() *FilterModelData {
	this := FilterModelData{}
	return &this
}

// NewFilterModelDataWithDefaults instantiates a new FilterModelData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterModelDataWithDefaults() *FilterModelData {
	this := FilterModelData{}
	return &this
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModelData) GetProjectIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModelData) GetProjectIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectIds) {
		return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *FilterModelData) HasProjectIds() bool {
	if o != nil && IsNil(o.ProjectIds) {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []string and assigns it to the ProjectIds field.
func (o *FilterModelData) SetProjectIds(v []string) {
	o.ProjectIds = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModelData) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModelData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *FilterModelData) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *FilterModelData) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *FilterModelData) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *FilterModelData) UnsetName() {
	o.Name.Unset()
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModelData) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModelData) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *FilterModelData) HasIds() bool {
	if o != nil && IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *FilterModelData) SetIds(v []string) {
	o.Ids = v
}

// GetGlobalIds returns the GlobalIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModelData) GetGlobalIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.GlobalIds
}

// GetGlobalIdsOk returns a tuple with the GlobalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModelData) GetGlobalIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.GlobalIds) {
		return nil, false
	}
	return o.GlobalIds, true
}

// HasGlobalIds returns a boolean if a field has been set.
func (o *FilterModelData) HasGlobalIds() bool {
	if o != nil && IsNil(o.GlobalIds) {
		return true
	}

	return false
}

// SetGlobalIds gets a reference to the given []int64 and assigns it to the GlobalIds field.
func (o *FilterModelData) SetGlobalIds(v []int64) {
	o.GlobalIds = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModelData) GetAttributes() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModelData) GetAttributesOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *FilterModelData) HasAttributes() bool {
	if o != nil && IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string][]string and assigns it to the Attributes field.
func (o *FilterModelData) SetAttributes(v map[string][]string) {
	o.Attributes = v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModelData) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDeleted.Get()
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModelData) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDeleted.Get(), o.IsDeleted.IsSet()
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *FilterModelData) HasIsDeleted() bool {
	if o != nil && o.IsDeleted.IsSet() {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given NullableBool and assigns it to the IsDeleted field.
func (o *FilterModelData) SetIsDeleted(v bool) {
	o.IsDeleted.Set(&v)
}
// SetIsDeletedNil sets the value for IsDeleted to be an explicit nil
func (o *FilterModelData) SetIsDeletedNil() {
	o.IsDeleted.Set(nil)
}

// UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
func (o *FilterModelData) UnsetIsDeleted() {
	o.IsDeleted.Unset()
}

// GetSectionIds returns the SectionIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModelData) GetSectionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SectionIds
}

// GetSectionIdsOk returns a tuple with the SectionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModelData) GetSectionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SectionIds) {
		return nil, false
	}
	return o.SectionIds, true
}

// HasSectionIds returns a boolean if a field has been set.
func (o *FilterModelData) HasSectionIds() bool {
	if o != nil && IsNil(o.SectionIds) {
		return true
	}

	return false
}

// SetSectionIds gets a reference to the given []string and assigns it to the SectionIds field.
func (o *FilterModelData) SetSectionIds(v []string) {
	o.SectionIds = v
}

// GetCreatedByIds returns the CreatedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModelData) GetCreatedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CreatedByIds
}

// GetCreatedByIdsOk returns a tuple with the CreatedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModelData) GetCreatedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CreatedByIds) {
		return nil, false
	}
	return o.CreatedByIds, true
}

// HasCreatedByIds returns a boolean if a field has been set.
func (o *FilterModelData) HasCreatedByIds() bool {
	if o != nil && IsNil(o.CreatedByIds) {
		return true
	}

	return false
}

// SetCreatedByIds gets a reference to the given []string and assigns it to the CreatedByIds field.
func (o *FilterModelData) SetCreatedByIds(v []string) {
	o.CreatedByIds = v
}

// GetModifiedByIds returns the ModifiedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModelData) GetModifiedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ModifiedByIds
}

// GetModifiedByIdsOk returns a tuple with the ModifiedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModelData) GetModifiedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ModifiedByIds) {
		return nil, false
	}
	return o.ModifiedByIds, true
}

// HasModifiedByIds returns a boolean if a field has been set.
func (o *FilterModelData) HasModifiedByIds() bool {
	if o != nil && IsNil(o.ModifiedByIds) {
		return true
	}

	return false
}

// SetModifiedByIds gets a reference to the given []string and assigns it to the ModifiedByIds field.
func (o *FilterModelData) SetModifiedByIds(v []string) {
	o.ModifiedByIds = v
}

// GetStates returns the States field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModelData) GetStates() []WorkItemStates {
	if o == nil {
		var ret []WorkItemStates
		return ret
	}
	return o.States
}

// GetStatesOk returns a tuple with the States field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModelData) GetStatesOk() ([]WorkItemStates, bool) {
	if o == nil || IsNil(o.States) {
		return nil, false
	}
	return o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *FilterModelData) HasStates() bool {
	if o != nil && IsNil(o.States) {
		return true
	}

	return false
}

// SetStates gets a reference to the given []WorkItemStates and assigns it to the States field.
func (o *FilterModelData) SetStates(v []WorkItemStates) {
	o.States = v
}

// GetPriorities returns the Priorities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModelData) GetPriorities() []WorkItemPriorityModel {
	if o == nil {
		var ret []WorkItemPriorityModel
		return ret
	}
	return o.Priorities
}

// GetPrioritiesOk returns a tuple with the Priorities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModelData) GetPrioritiesOk() ([]WorkItemPriorityModel, bool) {
	if o == nil || IsNil(o.Priorities) {
		return nil, false
	}
	return o.Priorities, true
}

// HasPriorities returns a boolean if a field has been set.
func (o *FilterModelData) HasPriorities() bool {
	if o != nil && IsNil(o.Priorities) {
		return true
	}

	return false
}

// SetPriorities gets a reference to the given []WorkItemPriorityModel and assigns it to the Priorities field.
func (o *FilterModelData) SetPriorities(v []WorkItemPriorityModel) {
	o.Priorities = v
}

// GetTypes returns the Types field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModelData) GetTypes() []WorkItemEntityTypes {
	if o == nil {
		var ret []WorkItemEntityTypes
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModelData) GetTypesOk() ([]WorkItemEntityTypes, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *FilterModelData) HasTypes() bool {
	if o != nil && IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []WorkItemEntityTypes and assigns it to the Types field.
func (o *FilterModelData) SetTypes(v []WorkItemEntityTypes) {
	o.Types = v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModelData) GetCreatedDate() TestPointFilterModelWorkItemCreatedDate {
	if o == nil || IsNil(o.CreatedDate.Get()) {
		var ret TestPointFilterModelWorkItemCreatedDate
		return ret
	}
	return *o.CreatedDate.Get()
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModelData) GetCreatedDateOk() (*TestPointFilterModelWorkItemCreatedDate, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDate.Get(), o.CreatedDate.IsSet()
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *FilterModelData) HasCreatedDate() bool {
	if o != nil && o.CreatedDate.IsSet() {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given NullableTestPointFilterModelWorkItemCreatedDate and assigns it to the CreatedDate field.
func (o *FilterModelData) SetCreatedDate(v TestPointFilterModelWorkItemCreatedDate) {
	o.CreatedDate.Set(&v)
}
// SetCreatedDateNil sets the value for CreatedDate to be an explicit nil
func (o *FilterModelData) SetCreatedDateNil() {
	o.CreatedDate.Set(nil)
}

// UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
func (o *FilterModelData) UnsetCreatedDate() {
	o.CreatedDate.Unset()
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModelData) GetModifiedDate() TestPointFilterModelWorkItemModifiedDate {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret TestPointFilterModelWorkItemModifiedDate
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModelData) GetModifiedDateOk() (*TestPointFilterModelWorkItemModifiedDate, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *FilterModelData) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTestPointFilterModelWorkItemModifiedDate and assigns it to the ModifiedDate field.
func (o *FilterModelData) SetModifiedDate(v TestPointFilterModelWorkItemModifiedDate) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *FilterModelData) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *FilterModelData) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModelData) GetDuration() TestSuiteWorkItemsSearchModelDuration {
	if o == nil || IsNil(o.Duration.Get()) {
		var ret TestSuiteWorkItemsSearchModelDuration
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModelData) GetDurationOk() (*TestSuiteWorkItemsSearchModelDuration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *FilterModelData) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableTestSuiteWorkItemsSearchModelDuration and assigns it to the Duration field.
func (o *FilterModelData) SetDuration(v TestSuiteWorkItemsSearchModelDuration) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *FilterModelData) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *FilterModelData) UnsetDuration() {
	o.Duration.Unset()
}

// GetMedianDuration returns the MedianDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModelData) GetMedianDuration() TestSuiteWorkItemsSearchModelMedianDuration {
	if o == nil || IsNil(o.MedianDuration.Get()) {
		var ret TestSuiteWorkItemsSearchModelMedianDuration
		return ret
	}
	return *o.MedianDuration.Get()
}

// GetMedianDurationOk returns a tuple with the MedianDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModelData) GetMedianDurationOk() (*TestSuiteWorkItemsSearchModelMedianDuration, bool) {
	if o == nil {
		return nil, false
	}
	return o.MedianDuration.Get(), o.MedianDuration.IsSet()
}

// HasMedianDuration returns a boolean if a field has been set.
func (o *FilterModelData) HasMedianDuration() bool {
	if o != nil && o.MedianDuration.IsSet() {
		return true
	}

	return false
}

// SetMedianDuration gets a reference to the given NullableTestSuiteWorkItemsSearchModelMedianDuration and assigns it to the MedianDuration field.
func (o *FilterModelData) SetMedianDuration(v TestSuiteWorkItemsSearchModelMedianDuration) {
	o.MedianDuration.Set(&v)
}
// SetMedianDurationNil sets the value for MedianDuration to be an explicit nil
func (o *FilterModelData) SetMedianDurationNil() {
	o.MedianDuration.Set(nil)
}

// UnsetMedianDuration ensures that no value is present for MedianDuration, not even an explicit nil
func (o *FilterModelData) UnsetMedianDuration() {
	o.MedianDuration.Unset()
}

// GetIsAutomated returns the IsAutomated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModelData) GetIsAutomated() bool {
	if o == nil || IsNil(o.IsAutomated.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAutomated.Get()
}

// GetIsAutomatedOk returns a tuple with the IsAutomated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModelData) GetIsAutomatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAutomated.Get(), o.IsAutomated.IsSet()
}

// HasIsAutomated returns a boolean if a field has been set.
func (o *FilterModelData) HasIsAutomated() bool {
	if o != nil && o.IsAutomated.IsSet() {
		return true
	}

	return false
}

// SetIsAutomated gets a reference to the given NullableBool and assigns it to the IsAutomated field.
func (o *FilterModelData) SetIsAutomated(v bool) {
	o.IsAutomated.Set(&v)
}
// SetIsAutomatedNil sets the value for IsAutomated to be an explicit nil
func (o *FilterModelData) SetIsAutomatedNil() {
	o.IsAutomated.Set(nil)
}

// UnsetIsAutomated ensures that no value is present for IsAutomated, not even an explicit nil
func (o *FilterModelData) UnsetIsAutomated() {
	o.IsAutomated.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModelData) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModelData) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FilterModelData) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FilterModelData) SetTags(v []string) {
	o.Tags = v
}

// GetAutoTestIds returns the AutoTestIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilterModelData) GetAutoTestIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AutoTestIds
}

// GetAutoTestIdsOk returns a tuple with the AutoTestIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilterModelData) GetAutoTestIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AutoTestIds) {
		return nil, false
	}
	return o.AutoTestIds, true
}

// HasAutoTestIds returns a boolean if a field has been set.
func (o *FilterModelData) HasAutoTestIds() bool {
	if o != nil && IsNil(o.AutoTestIds) {
		return true
	}

	return false
}

// SetAutoTestIds gets a reference to the given []string and assigns it to the AutoTestIds field.
func (o *FilterModelData) SetAutoTestIds(v []string) {
	o.AutoTestIds = v
}

func (o FilterModelData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilterModelData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectIds != nil {
		toSerialize["projectIds"] = o.ProjectIds
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.GlobalIds != nil {
		toSerialize["globalIds"] = o.GlobalIds
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.IsDeleted.IsSet() {
		toSerialize["isDeleted"] = o.IsDeleted.Get()
	}
	if o.SectionIds != nil {
		toSerialize["sectionIds"] = o.SectionIds
	}
	if o.CreatedByIds != nil {
		toSerialize["createdByIds"] = o.CreatedByIds
	}
	if o.ModifiedByIds != nil {
		toSerialize["modifiedByIds"] = o.ModifiedByIds
	}
	if o.States != nil {
		toSerialize["states"] = o.States
	}
	if o.Priorities != nil {
		toSerialize["priorities"] = o.Priorities
	}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	if o.CreatedDate.IsSet() {
		toSerialize["createdDate"] = o.CreatedDate.Get()
	}
	if o.ModifiedDate.IsSet() {
		toSerialize["modifiedDate"] = o.ModifiedDate.Get()
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if o.MedianDuration.IsSet() {
		toSerialize["medianDuration"] = o.MedianDuration.Get()
	}
	if o.IsAutomated.IsSet() {
		toSerialize["isAutomated"] = o.IsAutomated.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.AutoTestIds != nil {
		toSerialize["autoTestIds"] = o.AutoTestIds
	}
	return toSerialize, nil
}

type NullableFilterModelData struct {
	value *FilterModelData
	isSet bool
}

func (v NullableFilterModelData) Get() *FilterModelData {
	return v.value
}

func (v *NullableFilterModelData) Set(val *FilterModelData) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterModelData) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterModelData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterModelData(val *FilterModelData) *NullableFilterModelData {
	return &NullableFilterModelData{value: val, isSet: true}
}

func (v NullableFilterModelData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterModelData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


