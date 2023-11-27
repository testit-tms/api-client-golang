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

// checks if the WorkItemLocalFilterModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemLocalFilterModel{}

// WorkItemLocalFilterModel Collection of filters to apply to search
type WorkItemLocalFilterModel struct {
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

// NewWorkItemLocalFilterModel instantiates a new WorkItemLocalFilterModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemLocalFilterModel() *WorkItemLocalFilterModel {
	this := WorkItemLocalFilterModel{}
	return &this
}

// NewWorkItemLocalFilterModelWithDefaults instantiates a new WorkItemLocalFilterModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemLocalFilterModelWithDefaults() *WorkItemLocalFilterModel {
	this := WorkItemLocalFilterModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemLocalFilterModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemLocalFilterModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *WorkItemLocalFilterModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *WorkItemLocalFilterModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *WorkItemLocalFilterModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *WorkItemLocalFilterModel) UnsetName() {
	o.Name.Unset()
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemLocalFilterModel) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemLocalFilterModel) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *WorkItemLocalFilterModel) HasIds() bool {
	if o != nil && IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *WorkItemLocalFilterModel) SetIds(v []string) {
	o.Ids = v
}

// GetGlobalIds returns the GlobalIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemLocalFilterModel) GetGlobalIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.GlobalIds
}

// GetGlobalIdsOk returns a tuple with the GlobalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemLocalFilterModel) GetGlobalIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.GlobalIds) {
		return nil, false
	}
	return o.GlobalIds, true
}

// HasGlobalIds returns a boolean if a field has been set.
func (o *WorkItemLocalFilterModel) HasGlobalIds() bool {
	if o != nil && IsNil(o.GlobalIds) {
		return true
	}

	return false
}

// SetGlobalIds gets a reference to the given []int64 and assigns it to the GlobalIds field.
func (o *WorkItemLocalFilterModel) SetGlobalIds(v []int64) {
	o.GlobalIds = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemLocalFilterModel) GetAttributes() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemLocalFilterModel) GetAttributesOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WorkItemLocalFilterModel) HasAttributes() bool {
	if o != nil && IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string][]string and assigns it to the Attributes field.
func (o *WorkItemLocalFilterModel) SetAttributes(v map[string][]string) {
	o.Attributes = v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemLocalFilterModel) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDeleted.Get()
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemLocalFilterModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDeleted.Get(), o.IsDeleted.IsSet()
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *WorkItemLocalFilterModel) HasIsDeleted() bool {
	if o != nil && o.IsDeleted.IsSet() {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given NullableBool and assigns it to the IsDeleted field.
func (o *WorkItemLocalFilterModel) SetIsDeleted(v bool) {
	o.IsDeleted.Set(&v)
}
// SetIsDeletedNil sets the value for IsDeleted to be an explicit nil
func (o *WorkItemLocalFilterModel) SetIsDeletedNil() {
	o.IsDeleted.Set(nil)
}

// UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
func (o *WorkItemLocalFilterModel) UnsetIsDeleted() {
	o.IsDeleted.Unset()
}

// GetSectionIds returns the SectionIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemLocalFilterModel) GetSectionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SectionIds
}

// GetSectionIdsOk returns a tuple with the SectionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemLocalFilterModel) GetSectionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SectionIds) {
		return nil, false
	}
	return o.SectionIds, true
}

// HasSectionIds returns a boolean if a field has been set.
func (o *WorkItemLocalFilterModel) HasSectionIds() bool {
	if o != nil && IsNil(o.SectionIds) {
		return true
	}

	return false
}

// SetSectionIds gets a reference to the given []string and assigns it to the SectionIds field.
func (o *WorkItemLocalFilterModel) SetSectionIds(v []string) {
	o.SectionIds = v
}

// GetCreatedByIds returns the CreatedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemLocalFilterModel) GetCreatedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CreatedByIds
}

// GetCreatedByIdsOk returns a tuple with the CreatedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemLocalFilterModel) GetCreatedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CreatedByIds) {
		return nil, false
	}
	return o.CreatedByIds, true
}

// HasCreatedByIds returns a boolean if a field has been set.
func (o *WorkItemLocalFilterModel) HasCreatedByIds() bool {
	if o != nil && IsNil(o.CreatedByIds) {
		return true
	}

	return false
}

// SetCreatedByIds gets a reference to the given []string and assigns it to the CreatedByIds field.
func (o *WorkItemLocalFilterModel) SetCreatedByIds(v []string) {
	o.CreatedByIds = v
}

// GetModifiedByIds returns the ModifiedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemLocalFilterModel) GetModifiedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ModifiedByIds
}

// GetModifiedByIdsOk returns a tuple with the ModifiedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemLocalFilterModel) GetModifiedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ModifiedByIds) {
		return nil, false
	}
	return o.ModifiedByIds, true
}

// HasModifiedByIds returns a boolean if a field has been set.
func (o *WorkItemLocalFilterModel) HasModifiedByIds() bool {
	if o != nil && IsNil(o.ModifiedByIds) {
		return true
	}

	return false
}

// SetModifiedByIds gets a reference to the given []string and assigns it to the ModifiedByIds field.
func (o *WorkItemLocalFilterModel) SetModifiedByIds(v []string) {
	o.ModifiedByIds = v
}

// GetStates returns the States field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemLocalFilterModel) GetStates() []WorkItemStates {
	if o == nil {
		var ret []WorkItemStates
		return ret
	}
	return o.States
}

// GetStatesOk returns a tuple with the States field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemLocalFilterModel) GetStatesOk() ([]WorkItemStates, bool) {
	if o == nil || IsNil(o.States) {
		return nil, false
	}
	return o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *WorkItemLocalFilterModel) HasStates() bool {
	if o != nil && IsNil(o.States) {
		return true
	}

	return false
}

// SetStates gets a reference to the given []WorkItemStates and assigns it to the States field.
func (o *WorkItemLocalFilterModel) SetStates(v []WorkItemStates) {
	o.States = v
}

// GetPriorities returns the Priorities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemLocalFilterModel) GetPriorities() []WorkItemPriorityModel {
	if o == nil {
		var ret []WorkItemPriorityModel
		return ret
	}
	return o.Priorities
}

// GetPrioritiesOk returns a tuple with the Priorities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemLocalFilterModel) GetPrioritiesOk() ([]WorkItemPriorityModel, bool) {
	if o == nil || IsNil(o.Priorities) {
		return nil, false
	}
	return o.Priorities, true
}

// HasPriorities returns a boolean if a field has been set.
func (o *WorkItemLocalFilterModel) HasPriorities() bool {
	if o != nil && IsNil(o.Priorities) {
		return true
	}

	return false
}

// SetPriorities gets a reference to the given []WorkItemPriorityModel and assigns it to the Priorities field.
func (o *WorkItemLocalFilterModel) SetPriorities(v []WorkItemPriorityModel) {
	o.Priorities = v
}

// GetTypes returns the Types field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemLocalFilterModel) GetTypes() []WorkItemEntityTypes {
	if o == nil {
		var ret []WorkItemEntityTypes
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemLocalFilterModel) GetTypesOk() ([]WorkItemEntityTypes, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *WorkItemLocalFilterModel) HasTypes() bool {
	if o != nil && IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []WorkItemEntityTypes and assigns it to the Types field.
func (o *WorkItemLocalFilterModel) SetTypes(v []WorkItemEntityTypes) {
	o.Types = v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemLocalFilterModel) GetCreatedDate() TestPointFilterModelWorkItemCreatedDate {
	if o == nil || IsNil(o.CreatedDate.Get()) {
		var ret TestPointFilterModelWorkItemCreatedDate
		return ret
	}
	return *o.CreatedDate.Get()
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemLocalFilterModel) GetCreatedDateOk() (*TestPointFilterModelWorkItemCreatedDate, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDate.Get(), o.CreatedDate.IsSet()
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *WorkItemLocalFilterModel) HasCreatedDate() bool {
	if o != nil && o.CreatedDate.IsSet() {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given NullableTestPointFilterModelWorkItemCreatedDate and assigns it to the CreatedDate field.
func (o *WorkItemLocalFilterModel) SetCreatedDate(v TestPointFilterModelWorkItemCreatedDate) {
	o.CreatedDate.Set(&v)
}
// SetCreatedDateNil sets the value for CreatedDate to be an explicit nil
func (o *WorkItemLocalFilterModel) SetCreatedDateNil() {
	o.CreatedDate.Set(nil)
}

// UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
func (o *WorkItemLocalFilterModel) UnsetCreatedDate() {
	o.CreatedDate.Unset()
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemLocalFilterModel) GetModifiedDate() TestPointFilterModelWorkItemModifiedDate {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret TestPointFilterModelWorkItemModifiedDate
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemLocalFilterModel) GetModifiedDateOk() (*TestPointFilterModelWorkItemModifiedDate, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *WorkItemLocalFilterModel) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTestPointFilterModelWorkItemModifiedDate and assigns it to the ModifiedDate field.
func (o *WorkItemLocalFilterModel) SetModifiedDate(v TestPointFilterModelWorkItemModifiedDate) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *WorkItemLocalFilterModel) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *WorkItemLocalFilterModel) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemLocalFilterModel) GetDuration() TestSuiteWorkItemsSearchModelDuration {
	if o == nil || IsNil(o.Duration.Get()) {
		var ret TestSuiteWorkItemsSearchModelDuration
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemLocalFilterModel) GetDurationOk() (*TestSuiteWorkItemsSearchModelDuration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *WorkItemLocalFilterModel) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableTestSuiteWorkItemsSearchModelDuration and assigns it to the Duration field.
func (o *WorkItemLocalFilterModel) SetDuration(v TestSuiteWorkItemsSearchModelDuration) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *WorkItemLocalFilterModel) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *WorkItemLocalFilterModel) UnsetDuration() {
	o.Duration.Unset()
}

// GetMedianDuration returns the MedianDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemLocalFilterModel) GetMedianDuration() TestSuiteWorkItemsSearchModelMedianDuration {
	if o == nil || IsNil(o.MedianDuration.Get()) {
		var ret TestSuiteWorkItemsSearchModelMedianDuration
		return ret
	}
	return *o.MedianDuration.Get()
}

// GetMedianDurationOk returns a tuple with the MedianDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemLocalFilterModel) GetMedianDurationOk() (*TestSuiteWorkItemsSearchModelMedianDuration, bool) {
	if o == nil {
		return nil, false
	}
	return o.MedianDuration.Get(), o.MedianDuration.IsSet()
}

// HasMedianDuration returns a boolean if a field has been set.
func (o *WorkItemLocalFilterModel) HasMedianDuration() bool {
	if o != nil && o.MedianDuration.IsSet() {
		return true
	}

	return false
}

// SetMedianDuration gets a reference to the given NullableTestSuiteWorkItemsSearchModelMedianDuration and assigns it to the MedianDuration field.
func (o *WorkItemLocalFilterModel) SetMedianDuration(v TestSuiteWorkItemsSearchModelMedianDuration) {
	o.MedianDuration.Set(&v)
}
// SetMedianDurationNil sets the value for MedianDuration to be an explicit nil
func (o *WorkItemLocalFilterModel) SetMedianDurationNil() {
	o.MedianDuration.Set(nil)
}

// UnsetMedianDuration ensures that no value is present for MedianDuration, not even an explicit nil
func (o *WorkItemLocalFilterModel) UnsetMedianDuration() {
	o.MedianDuration.Unset()
}

// GetIsAutomated returns the IsAutomated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemLocalFilterModel) GetIsAutomated() bool {
	if o == nil || IsNil(o.IsAutomated.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAutomated.Get()
}

// GetIsAutomatedOk returns a tuple with the IsAutomated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemLocalFilterModel) GetIsAutomatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAutomated.Get(), o.IsAutomated.IsSet()
}

// HasIsAutomated returns a boolean if a field has been set.
func (o *WorkItemLocalFilterModel) HasIsAutomated() bool {
	if o != nil && o.IsAutomated.IsSet() {
		return true
	}

	return false
}

// SetIsAutomated gets a reference to the given NullableBool and assigns it to the IsAutomated field.
func (o *WorkItemLocalFilterModel) SetIsAutomated(v bool) {
	o.IsAutomated.Set(&v)
}
// SetIsAutomatedNil sets the value for IsAutomated to be an explicit nil
func (o *WorkItemLocalFilterModel) SetIsAutomatedNil() {
	o.IsAutomated.Set(nil)
}

// UnsetIsAutomated ensures that no value is present for IsAutomated, not even an explicit nil
func (o *WorkItemLocalFilterModel) UnsetIsAutomated() {
	o.IsAutomated.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemLocalFilterModel) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemLocalFilterModel) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WorkItemLocalFilterModel) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *WorkItemLocalFilterModel) SetTags(v []string) {
	o.Tags = v
}

// GetAutoTestIds returns the AutoTestIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemLocalFilterModel) GetAutoTestIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AutoTestIds
}

// GetAutoTestIdsOk returns a tuple with the AutoTestIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemLocalFilterModel) GetAutoTestIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AutoTestIds) {
		return nil, false
	}
	return o.AutoTestIds, true
}

// HasAutoTestIds returns a boolean if a field has been set.
func (o *WorkItemLocalFilterModel) HasAutoTestIds() bool {
	if o != nil && IsNil(o.AutoTestIds) {
		return true
	}

	return false
}

// SetAutoTestIds gets a reference to the given []string and assigns it to the AutoTestIds field.
func (o *WorkItemLocalFilterModel) SetAutoTestIds(v []string) {
	o.AutoTestIds = v
}

func (o WorkItemLocalFilterModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemLocalFilterModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullableWorkItemLocalFilterModel struct {
	value *WorkItemLocalFilterModel
	isSet bool
}

func (v NullableWorkItemLocalFilterModel) Get() *WorkItemLocalFilterModel {
	return v.value
}

func (v *NullableWorkItemLocalFilterModel) Set(val *WorkItemLocalFilterModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemLocalFilterModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemLocalFilterModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemLocalFilterModel(val *WorkItemLocalFilterModel) *NullableWorkItemLocalFilterModel {
	return &NullableWorkItemLocalFilterModel{value: val, isSet: true}
}

func (v NullableWorkItemLocalFilterModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemLocalFilterModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


