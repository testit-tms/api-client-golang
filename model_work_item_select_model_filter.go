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

// checks if the WorkItemSelectModelFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemSelectModelFilter{}

// WorkItemSelectModelFilter Collection of filters to apply to search
type WorkItemSelectModelFilter struct {
	// Name or identifier (UUID) of work item
	NameOrId NullableString `json:"nameOrId,omitempty"`
	// Collection of identifiers of work items which need to be included in result regardless of filtering
	IncludeIds []string `json:"includeIds,omitempty"`
	// Collection of identifiers of work items which need to be excluded from result regardless of filtering
	ExcludeIds []string `json:"excludeIds,omitempty"`
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
	// Collection of project identifiers
	ProjectIds []string `json:"projectIds,omitempty"`
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
	Duration NullableWorkItemFilterModelDuration `json:"duration,omitempty"`
	// Is result must consist of only manual/automated work items
	IsAutomated NullableBool `json:"isAutomated,omitempty"`
	// Collection of tags
	Tags []string `json:"tags,omitempty"`
	// Collection of identifiers of linked autotests
	AutoTestIds []string `json:"autoTestIds,omitempty"`
}

// NewWorkItemSelectModelFilter instantiates a new WorkItemSelectModelFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemSelectModelFilter() *WorkItemSelectModelFilter {
	this := WorkItemSelectModelFilter{}
	return &this
}

// NewWorkItemSelectModelFilterWithDefaults instantiates a new WorkItemSelectModelFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemSelectModelFilterWithDefaults() *WorkItemSelectModelFilter {
	this := WorkItemSelectModelFilter{}
	return &this
}

// GetNameOrId returns the NameOrId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetNameOrId() string {
	if o == nil || IsNil(o.NameOrId.Get()) {
		var ret string
		return ret
	}
	return *o.NameOrId.Get()
}

// GetNameOrIdOk returns a tuple with the NameOrId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetNameOrIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NameOrId.Get(), o.NameOrId.IsSet()
}

// HasNameOrId returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasNameOrId() bool {
	if o != nil && o.NameOrId.IsSet() {
		return true
	}

	return false
}

// SetNameOrId gets a reference to the given NullableString and assigns it to the NameOrId field.
func (o *WorkItemSelectModelFilter) SetNameOrId(v string) {
	o.NameOrId.Set(&v)
}
// SetNameOrIdNil sets the value for NameOrId to be an explicit nil
func (o *WorkItemSelectModelFilter) SetNameOrIdNil() {
	o.NameOrId.Set(nil)
}

// UnsetNameOrId ensures that no value is present for NameOrId, not even an explicit nil
func (o *WorkItemSelectModelFilter) UnsetNameOrId() {
	o.NameOrId.Unset()
}

// GetIncludeIds returns the IncludeIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetIncludeIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IncludeIds
}

// GetIncludeIdsOk returns a tuple with the IncludeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetIncludeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeIds) {
		return nil, false
	}
	return o.IncludeIds, true
}

// HasIncludeIds returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasIncludeIds() bool {
	if o != nil && IsNil(o.IncludeIds) {
		return true
	}

	return false
}

// SetIncludeIds gets a reference to the given []string and assigns it to the IncludeIds field.
func (o *WorkItemSelectModelFilter) SetIncludeIds(v []string) {
	o.IncludeIds = v
}

// GetExcludeIds returns the ExcludeIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetExcludeIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludeIds
}

// GetExcludeIdsOk returns a tuple with the ExcludeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetExcludeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeIds) {
		return nil, false
	}
	return o.ExcludeIds, true
}

// HasExcludeIds returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasExcludeIds() bool {
	if o != nil && IsNil(o.ExcludeIds) {
		return true
	}

	return false
}

// SetExcludeIds gets a reference to the given []string and assigns it to the ExcludeIds field.
func (o *WorkItemSelectModelFilter) SetExcludeIds(v []string) {
	o.ExcludeIds = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *WorkItemSelectModelFilter) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *WorkItemSelectModelFilter) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *WorkItemSelectModelFilter) UnsetName() {
	o.Name.Unset()
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasIds() bool {
	if o != nil && IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *WorkItemSelectModelFilter) SetIds(v []string) {
	o.Ids = v
}

// GetGlobalIds returns the GlobalIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetGlobalIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.GlobalIds
}

// GetGlobalIdsOk returns a tuple with the GlobalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetGlobalIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.GlobalIds) {
		return nil, false
	}
	return o.GlobalIds, true
}

// HasGlobalIds returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasGlobalIds() bool {
	if o != nil && IsNil(o.GlobalIds) {
		return true
	}

	return false
}

// SetGlobalIds gets a reference to the given []int64 and assigns it to the GlobalIds field.
func (o *WorkItemSelectModelFilter) SetGlobalIds(v []int64) {
	o.GlobalIds = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetAttributes() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetAttributesOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasAttributes() bool {
	if o != nil && IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string][]string and assigns it to the Attributes field.
func (o *WorkItemSelectModelFilter) SetAttributes(v map[string][]string) {
	o.Attributes = v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDeleted.Get()
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDeleted.Get(), o.IsDeleted.IsSet()
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasIsDeleted() bool {
	if o != nil && o.IsDeleted.IsSet() {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given NullableBool and assigns it to the IsDeleted field.
func (o *WorkItemSelectModelFilter) SetIsDeleted(v bool) {
	o.IsDeleted.Set(&v)
}
// SetIsDeletedNil sets the value for IsDeleted to be an explicit nil
func (o *WorkItemSelectModelFilter) SetIsDeletedNil() {
	o.IsDeleted.Set(nil)
}

// UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
func (o *WorkItemSelectModelFilter) UnsetIsDeleted() {
	o.IsDeleted.Unset()
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetProjectIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetProjectIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectIds) {
		return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasProjectIds() bool {
	if o != nil && IsNil(o.ProjectIds) {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []string and assigns it to the ProjectIds field.
func (o *WorkItemSelectModelFilter) SetProjectIds(v []string) {
	o.ProjectIds = v
}

// GetSectionIds returns the SectionIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetSectionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SectionIds
}

// GetSectionIdsOk returns a tuple with the SectionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetSectionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SectionIds) {
		return nil, false
	}
	return o.SectionIds, true
}

// HasSectionIds returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasSectionIds() bool {
	if o != nil && IsNil(o.SectionIds) {
		return true
	}

	return false
}

// SetSectionIds gets a reference to the given []string and assigns it to the SectionIds field.
func (o *WorkItemSelectModelFilter) SetSectionIds(v []string) {
	o.SectionIds = v
}

// GetCreatedByIds returns the CreatedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetCreatedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CreatedByIds
}

// GetCreatedByIdsOk returns a tuple with the CreatedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetCreatedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CreatedByIds) {
		return nil, false
	}
	return o.CreatedByIds, true
}

// HasCreatedByIds returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasCreatedByIds() bool {
	if o != nil && IsNil(o.CreatedByIds) {
		return true
	}

	return false
}

// SetCreatedByIds gets a reference to the given []string and assigns it to the CreatedByIds field.
func (o *WorkItemSelectModelFilter) SetCreatedByIds(v []string) {
	o.CreatedByIds = v
}

// GetModifiedByIds returns the ModifiedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetModifiedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ModifiedByIds
}

// GetModifiedByIdsOk returns a tuple with the ModifiedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetModifiedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ModifiedByIds) {
		return nil, false
	}
	return o.ModifiedByIds, true
}

// HasModifiedByIds returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasModifiedByIds() bool {
	if o != nil && IsNil(o.ModifiedByIds) {
		return true
	}

	return false
}

// SetModifiedByIds gets a reference to the given []string and assigns it to the ModifiedByIds field.
func (o *WorkItemSelectModelFilter) SetModifiedByIds(v []string) {
	o.ModifiedByIds = v
}

// GetStates returns the States field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetStates() []WorkItemStates {
	if o == nil {
		var ret []WorkItemStates
		return ret
	}
	return o.States
}

// GetStatesOk returns a tuple with the States field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetStatesOk() ([]WorkItemStates, bool) {
	if o == nil || IsNil(o.States) {
		return nil, false
	}
	return o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasStates() bool {
	if o != nil && IsNil(o.States) {
		return true
	}

	return false
}

// SetStates gets a reference to the given []WorkItemStates and assigns it to the States field.
func (o *WorkItemSelectModelFilter) SetStates(v []WorkItemStates) {
	o.States = v
}

// GetPriorities returns the Priorities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetPriorities() []WorkItemPriorityModel {
	if o == nil {
		var ret []WorkItemPriorityModel
		return ret
	}
	return o.Priorities
}

// GetPrioritiesOk returns a tuple with the Priorities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetPrioritiesOk() ([]WorkItemPriorityModel, bool) {
	if o == nil || IsNil(o.Priorities) {
		return nil, false
	}
	return o.Priorities, true
}

// HasPriorities returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasPriorities() bool {
	if o != nil && IsNil(o.Priorities) {
		return true
	}

	return false
}

// SetPriorities gets a reference to the given []WorkItemPriorityModel and assigns it to the Priorities field.
func (o *WorkItemSelectModelFilter) SetPriorities(v []WorkItemPriorityModel) {
	o.Priorities = v
}

// GetTypes returns the Types field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetTypes() []WorkItemEntityTypes {
	if o == nil {
		var ret []WorkItemEntityTypes
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetTypesOk() ([]WorkItemEntityTypes, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasTypes() bool {
	if o != nil && IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []WorkItemEntityTypes and assigns it to the Types field.
func (o *WorkItemSelectModelFilter) SetTypes(v []WorkItemEntityTypes) {
	o.Types = v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetCreatedDate() TestPointFilterModelWorkItemCreatedDate {
	if o == nil || IsNil(o.CreatedDate.Get()) {
		var ret TestPointFilterModelWorkItemCreatedDate
		return ret
	}
	return *o.CreatedDate.Get()
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetCreatedDateOk() (*TestPointFilterModelWorkItemCreatedDate, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDate.Get(), o.CreatedDate.IsSet()
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasCreatedDate() bool {
	if o != nil && o.CreatedDate.IsSet() {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given NullableTestPointFilterModelWorkItemCreatedDate and assigns it to the CreatedDate field.
func (o *WorkItemSelectModelFilter) SetCreatedDate(v TestPointFilterModelWorkItemCreatedDate) {
	o.CreatedDate.Set(&v)
}
// SetCreatedDateNil sets the value for CreatedDate to be an explicit nil
func (o *WorkItemSelectModelFilter) SetCreatedDateNil() {
	o.CreatedDate.Set(nil)
}

// UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
func (o *WorkItemSelectModelFilter) UnsetCreatedDate() {
	o.CreatedDate.Unset()
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetModifiedDate() TestPointFilterModelWorkItemModifiedDate {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret TestPointFilterModelWorkItemModifiedDate
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetModifiedDateOk() (*TestPointFilterModelWorkItemModifiedDate, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTestPointFilterModelWorkItemModifiedDate and assigns it to the ModifiedDate field.
func (o *WorkItemSelectModelFilter) SetModifiedDate(v TestPointFilterModelWorkItemModifiedDate) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *WorkItemSelectModelFilter) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *WorkItemSelectModelFilter) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetDuration() WorkItemFilterModelDuration {
	if o == nil || IsNil(o.Duration.Get()) {
		var ret WorkItemFilterModelDuration
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetDurationOk() (*WorkItemFilterModelDuration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableWorkItemFilterModelDuration and assigns it to the Duration field.
func (o *WorkItemSelectModelFilter) SetDuration(v WorkItemFilterModelDuration) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *WorkItemSelectModelFilter) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *WorkItemSelectModelFilter) UnsetDuration() {
	o.Duration.Unset()
}

// GetIsAutomated returns the IsAutomated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetIsAutomated() bool {
	if o == nil || IsNil(o.IsAutomated.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAutomated.Get()
}

// GetIsAutomatedOk returns a tuple with the IsAutomated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetIsAutomatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAutomated.Get(), o.IsAutomated.IsSet()
}

// HasIsAutomated returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasIsAutomated() bool {
	if o != nil && o.IsAutomated.IsSet() {
		return true
	}

	return false
}

// SetIsAutomated gets a reference to the given NullableBool and assigns it to the IsAutomated field.
func (o *WorkItemSelectModelFilter) SetIsAutomated(v bool) {
	o.IsAutomated.Set(&v)
}
// SetIsAutomatedNil sets the value for IsAutomated to be an explicit nil
func (o *WorkItemSelectModelFilter) SetIsAutomatedNil() {
	o.IsAutomated.Set(nil)
}

// UnsetIsAutomated ensures that no value is present for IsAutomated, not even an explicit nil
func (o *WorkItemSelectModelFilter) UnsetIsAutomated() {
	o.IsAutomated.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *WorkItemSelectModelFilter) SetTags(v []string) {
	o.Tags = v
}

// GetAutoTestIds returns the AutoTestIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectModelFilter) GetAutoTestIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AutoTestIds
}

// GetAutoTestIdsOk returns a tuple with the AutoTestIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectModelFilter) GetAutoTestIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AutoTestIds) {
		return nil, false
	}
	return o.AutoTestIds, true
}

// HasAutoTestIds returns a boolean if a field has been set.
func (o *WorkItemSelectModelFilter) HasAutoTestIds() bool {
	if o != nil && IsNil(o.AutoTestIds) {
		return true
	}

	return false
}

// SetAutoTestIds gets a reference to the given []string and assigns it to the AutoTestIds field.
func (o *WorkItemSelectModelFilter) SetAutoTestIds(v []string) {
	o.AutoTestIds = v
}

func (o WorkItemSelectModelFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemSelectModelFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NameOrId.IsSet() {
		toSerialize["nameOrId"] = o.NameOrId.Get()
	}
	if o.IncludeIds != nil {
		toSerialize["includeIds"] = o.IncludeIds
	}
	if o.ExcludeIds != nil {
		toSerialize["excludeIds"] = o.ExcludeIds
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
	if o.ProjectIds != nil {
		toSerialize["projectIds"] = o.ProjectIds
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

type NullableWorkItemSelectModelFilter struct {
	value *WorkItemSelectModelFilter
	isSet bool
}

func (v NullableWorkItemSelectModelFilter) Get() *WorkItemSelectModelFilter {
	return v.value
}

func (v *NullableWorkItemSelectModelFilter) Set(val *WorkItemSelectModelFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemSelectModelFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemSelectModelFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemSelectModelFilter(val *WorkItemSelectModelFilter) *NullableWorkItemSelectModelFilter {
	return &NullableWorkItemSelectModelFilter{value: val, isSet: true}
}

func (v NullableWorkItemSelectModelFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemSelectModelFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


