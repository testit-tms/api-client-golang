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

// checks if the WorkItemSearchQueryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemSearchQueryModel{}

// WorkItemSearchQueryModel struct for WorkItemSearchQueryModel
type WorkItemSearchQueryModel struct {
	// Collection of project identifiers
	ProjectIds []string `json:"projectIds,omitempty"`
	// Specifies a work item filter by its links
	Links NullableWorkItemLinkFilterModel `json:"links,omitempty"`
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
	// Specifies a work item range of creation date to search for
	CreatedDate NullableDateTimeRangeSelectorModel `json:"createdDate,omitempty"`
	// Specifies a work item range of last modification date to search for
	ModifiedDate NullableDateTimeRangeSelectorModel `json:"modifiedDate,omitempty"`
	// Specifies a work item duration range to search for
	Duration NullableInt32RangeSelectorModel `json:"duration,omitempty"`
	// Specifies a work item median duration range to search for
	MedianDuration NullableInt64RangeSelectorModel `json:"medianDuration,omitempty"`
	// Is result must consist of only manual/automated work items
	IsAutomated NullableBool `json:"isAutomated,omitempty"`
	// Collection of tags
	Tags []string `json:"tags,omitempty"`
	// Collection of identifiers of linked autotests
	AutoTestIds []string `json:"autoTestIds,omitempty"`
	// Collection of identifiers work items versions.
	WorkItemVersionIds []string `json:"workItemVersionIds,omitempty"`
}

// NewWorkItemSearchQueryModel instantiates a new WorkItemSearchQueryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemSearchQueryModel() *WorkItemSearchQueryModel {
	this := WorkItemSearchQueryModel{}
	return &this
}

// NewWorkItemSearchQueryModelWithDefaults instantiates a new WorkItemSearchQueryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemSearchQueryModelWithDefaults() *WorkItemSearchQueryModel {
	this := WorkItemSearchQueryModel{}
	return &this
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetProjectIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetProjectIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectIds) {
		return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasProjectIds() bool {
	if o != nil && !IsNil(o.ProjectIds) {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []string and assigns it to the ProjectIds field.
func (o *WorkItemSearchQueryModel) SetProjectIds(v []string) {
	o.ProjectIds = v
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetLinks() WorkItemLinkFilterModel {
	if o == nil || IsNil(o.Links.Get()) {
		var ret WorkItemLinkFilterModel
		return ret
	}
	return *o.Links.Get()
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetLinksOk() (*WorkItemLinkFilterModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links.Get(), o.Links.IsSet()
}

// HasLinks returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasLinks() bool {
	if o != nil && o.Links.IsSet() {
		return true
	}

	return false
}

// SetLinks gets a reference to the given NullableWorkItemLinkFilterModel and assigns it to the Links field.
func (o *WorkItemSearchQueryModel) SetLinks(v WorkItemLinkFilterModel) {
	o.Links.Set(&v)
}
// SetLinksNil sets the value for Links to be an explicit nil
func (o *WorkItemSearchQueryModel) SetLinksNil() {
	o.Links.Set(nil)
}

// UnsetLinks ensures that no value is present for Links, not even an explicit nil
func (o *WorkItemSearchQueryModel) UnsetLinks() {
	o.Links.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *WorkItemSearchQueryModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *WorkItemSearchQueryModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *WorkItemSearchQueryModel) UnsetName() {
	o.Name.Unset()
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *WorkItemSearchQueryModel) SetIds(v []string) {
	o.Ids = v
}

// GetGlobalIds returns the GlobalIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetGlobalIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.GlobalIds
}

// GetGlobalIdsOk returns a tuple with the GlobalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetGlobalIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.GlobalIds) {
		return nil, false
	}
	return o.GlobalIds, true
}

// HasGlobalIds returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasGlobalIds() bool {
	if o != nil && !IsNil(o.GlobalIds) {
		return true
	}

	return false
}

// SetGlobalIds gets a reference to the given []int64 and assigns it to the GlobalIds field.
func (o *WorkItemSearchQueryModel) SetGlobalIds(v []int64) {
	o.GlobalIds = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetAttributes() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetAttributesOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string][]string and assigns it to the Attributes field.
func (o *WorkItemSearchQueryModel) SetAttributes(v map[string][]string) {
	o.Attributes = v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDeleted.Get()
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDeleted.Get(), o.IsDeleted.IsSet()
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasIsDeleted() bool {
	if o != nil && o.IsDeleted.IsSet() {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given NullableBool and assigns it to the IsDeleted field.
func (o *WorkItemSearchQueryModel) SetIsDeleted(v bool) {
	o.IsDeleted.Set(&v)
}
// SetIsDeletedNil sets the value for IsDeleted to be an explicit nil
func (o *WorkItemSearchQueryModel) SetIsDeletedNil() {
	o.IsDeleted.Set(nil)
}

// UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
func (o *WorkItemSearchQueryModel) UnsetIsDeleted() {
	o.IsDeleted.Unset()
}

// GetSectionIds returns the SectionIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetSectionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SectionIds
}

// GetSectionIdsOk returns a tuple with the SectionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetSectionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SectionIds) {
		return nil, false
	}
	return o.SectionIds, true
}

// HasSectionIds returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasSectionIds() bool {
	if o != nil && !IsNil(o.SectionIds) {
		return true
	}

	return false
}

// SetSectionIds gets a reference to the given []string and assigns it to the SectionIds field.
func (o *WorkItemSearchQueryModel) SetSectionIds(v []string) {
	o.SectionIds = v
}

// GetCreatedByIds returns the CreatedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetCreatedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CreatedByIds
}

// GetCreatedByIdsOk returns a tuple with the CreatedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetCreatedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CreatedByIds) {
		return nil, false
	}
	return o.CreatedByIds, true
}

// HasCreatedByIds returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasCreatedByIds() bool {
	if o != nil && !IsNil(o.CreatedByIds) {
		return true
	}

	return false
}

// SetCreatedByIds gets a reference to the given []string and assigns it to the CreatedByIds field.
func (o *WorkItemSearchQueryModel) SetCreatedByIds(v []string) {
	o.CreatedByIds = v
}

// GetModifiedByIds returns the ModifiedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetModifiedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ModifiedByIds
}

// GetModifiedByIdsOk returns a tuple with the ModifiedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetModifiedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ModifiedByIds) {
		return nil, false
	}
	return o.ModifiedByIds, true
}

// HasModifiedByIds returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasModifiedByIds() bool {
	if o != nil && !IsNil(o.ModifiedByIds) {
		return true
	}

	return false
}

// SetModifiedByIds gets a reference to the given []string and assigns it to the ModifiedByIds field.
func (o *WorkItemSearchQueryModel) SetModifiedByIds(v []string) {
	o.ModifiedByIds = v
}

// GetStates returns the States field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetStates() []WorkItemStates {
	if o == nil {
		var ret []WorkItemStates
		return ret
	}
	return o.States
}

// GetStatesOk returns a tuple with the States field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetStatesOk() ([]WorkItemStates, bool) {
	if o == nil || IsNil(o.States) {
		return nil, false
	}
	return o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasStates() bool {
	if o != nil && !IsNil(o.States) {
		return true
	}

	return false
}

// SetStates gets a reference to the given []WorkItemStates and assigns it to the States field.
func (o *WorkItemSearchQueryModel) SetStates(v []WorkItemStates) {
	o.States = v
}

// GetPriorities returns the Priorities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetPriorities() []WorkItemPriorityModel {
	if o == nil {
		var ret []WorkItemPriorityModel
		return ret
	}
	return o.Priorities
}

// GetPrioritiesOk returns a tuple with the Priorities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetPrioritiesOk() ([]WorkItemPriorityModel, bool) {
	if o == nil || IsNil(o.Priorities) {
		return nil, false
	}
	return o.Priorities, true
}

// HasPriorities returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasPriorities() bool {
	if o != nil && !IsNil(o.Priorities) {
		return true
	}

	return false
}

// SetPriorities gets a reference to the given []WorkItemPriorityModel and assigns it to the Priorities field.
func (o *WorkItemSearchQueryModel) SetPriorities(v []WorkItemPriorityModel) {
	o.Priorities = v
}

// GetTypes returns the Types field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetTypes() []WorkItemEntityTypes {
	if o == nil {
		var ret []WorkItemEntityTypes
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetTypesOk() ([]WorkItemEntityTypes, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []WorkItemEntityTypes and assigns it to the Types field.
func (o *WorkItemSearchQueryModel) SetTypes(v []WorkItemEntityTypes) {
	o.Types = v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetCreatedDate() DateTimeRangeSelectorModel {
	if o == nil || IsNil(o.CreatedDate.Get()) {
		var ret DateTimeRangeSelectorModel
		return ret
	}
	return *o.CreatedDate.Get()
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDate.Get(), o.CreatedDate.IsSet()
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasCreatedDate() bool {
	if o != nil && o.CreatedDate.IsSet() {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given NullableDateTimeRangeSelectorModel and assigns it to the CreatedDate field.
func (o *WorkItemSearchQueryModel) SetCreatedDate(v DateTimeRangeSelectorModel) {
	o.CreatedDate.Set(&v)
}
// SetCreatedDateNil sets the value for CreatedDate to be an explicit nil
func (o *WorkItemSearchQueryModel) SetCreatedDateNil() {
	o.CreatedDate.Set(nil)
}

// UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
func (o *WorkItemSearchQueryModel) UnsetCreatedDate() {
	o.CreatedDate.Unset()
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetModifiedDate() DateTimeRangeSelectorModel {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret DateTimeRangeSelectorModel
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetModifiedDateOk() (*DateTimeRangeSelectorModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableDateTimeRangeSelectorModel and assigns it to the ModifiedDate field.
func (o *WorkItemSearchQueryModel) SetModifiedDate(v DateTimeRangeSelectorModel) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *WorkItemSearchQueryModel) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *WorkItemSearchQueryModel) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetDuration() Int32RangeSelectorModel {
	if o == nil || IsNil(o.Duration.Get()) {
		var ret Int32RangeSelectorModel
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetDurationOk() (*Int32RangeSelectorModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableInt32RangeSelectorModel and assigns it to the Duration field.
func (o *WorkItemSearchQueryModel) SetDuration(v Int32RangeSelectorModel) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *WorkItemSearchQueryModel) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *WorkItemSearchQueryModel) UnsetDuration() {
	o.Duration.Unset()
}

// GetMedianDuration returns the MedianDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetMedianDuration() Int64RangeSelectorModel {
	if o == nil || IsNil(o.MedianDuration.Get()) {
		var ret Int64RangeSelectorModel
		return ret
	}
	return *o.MedianDuration.Get()
}

// GetMedianDurationOk returns a tuple with the MedianDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetMedianDurationOk() (*Int64RangeSelectorModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.MedianDuration.Get(), o.MedianDuration.IsSet()
}

// HasMedianDuration returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasMedianDuration() bool {
	if o != nil && o.MedianDuration.IsSet() {
		return true
	}

	return false
}

// SetMedianDuration gets a reference to the given NullableInt64RangeSelectorModel and assigns it to the MedianDuration field.
func (o *WorkItemSearchQueryModel) SetMedianDuration(v Int64RangeSelectorModel) {
	o.MedianDuration.Set(&v)
}
// SetMedianDurationNil sets the value for MedianDuration to be an explicit nil
func (o *WorkItemSearchQueryModel) SetMedianDurationNil() {
	o.MedianDuration.Set(nil)
}

// UnsetMedianDuration ensures that no value is present for MedianDuration, not even an explicit nil
func (o *WorkItemSearchQueryModel) UnsetMedianDuration() {
	o.MedianDuration.Unset()
}

// GetIsAutomated returns the IsAutomated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetIsAutomated() bool {
	if o == nil || IsNil(o.IsAutomated.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAutomated.Get()
}

// GetIsAutomatedOk returns a tuple with the IsAutomated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetIsAutomatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAutomated.Get(), o.IsAutomated.IsSet()
}

// HasIsAutomated returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasIsAutomated() bool {
	if o != nil && o.IsAutomated.IsSet() {
		return true
	}

	return false
}

// SetIsAutomated gets a reference to the given NullableBool and assigns it to the IsAutomated field.
func (o *WorkItemSearchQueryModel) SetIsAutomated(v bool) {
	o.IsAutomated.Set(&v)
}
// SetIsAutomatedNil sets the value for IsAutomated to be an explicit nil
func (o *WorkItemSearchQueryModel) SetIsAutomatedNil() {
	o.IsAutomated.Set(nil)
}

// UnsetIsAutomated ensures that no value is present for IsAutomated, not even an explicit nil
func (o *WorkItemSearchQueryModel) UnsetIsAutomated() {
	o.IsAutomated.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *WorkItemSearchQueryModel) SetTags(v []string) {
	o.Tags = v
}

// GetAutoTestIds returns the AutoTestIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetAutoTestIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AutoTestIds
}

// GetAutoTestIdsOk returns a tuple with the AutoTestIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetAutoTestIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AutoTestIds) {
		return nil, false
	}
	return o.AutoTestIds, true
}

// HasAutoTestIds returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasAutoTestIds() bool {
	if o != nil && !IsNil(o.AutoTestIds) {
		return true
	}

	return false
}

// SetAutoTestIds gets a reference to the given []string and assigns it to the AutoTestIds field.
func (o *WorkItemSearchQueryModel) SetAutoTestIds(v []string) {
	o.AutoTestIds = v
}

// GetWorkItemVersionIds returns the WorkItemVersionIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSearchQueryModel) GetWorkItemVersionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.WorkItemVersionIds
}

// GetWorkItemVersionIdsOk returns a tuple with the WorkItemVersionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSearchQueryModel) GetWorkItemVersionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.WorkItemVersionIds) {
		return nil, false
	}
	return o.WorkItemVersionIds, true
}

// HasWorkItemVersionIds returns a boolean if a field has been set.
func (o *WorkItemSearchQueryModel) HasWorkItemVersionIds() bool {
	if o != nil && !IsNil(o.WorkItemVersionIds) {
		return true
	}

	return false
}

// SetWorkItemVersionIds gets a reference to the given []string and assigns it to the WorkItemVersionIds field.
func (o *WorkItemSearchQueryModel) SetWorkItemVersionIds(v []string) {
	o.WorkItemVersionIds = v
}

func (o WorkItemSearchQueryModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemSearchQueryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectIds != nil {
		toSerialize["projectIds"] = o.ProjectIds
	}
	if o.Links.IsSet() {
		toSerialize["links"] = o.Links.Get()
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
	if o.WorkItemVersionIds != nil {
		toSerialize["workItemVersionIds"] = o.WorkItemVersionIds
	}
	return toSerialize, nil
}

type NullableWorkItemSearchQueryModel struct {
	value *WorkItemSearchQueryModel
	isSet bool
}

func (v NullableWorkItemSearchQueryModel) Get() *WorkItemSearchQueryModel {
	return v.value
}

func (v *NullableWorkItemSearchQueryModel) Set(val *WorkItemSearchQueryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemSearchQueryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemSearchQueryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemSearchQueryModel(val *WorkItemSearchQueryModel) *NullableWorkItemSearchQueryModel {
	return &NullableWorkItemSearchQueryModel{value: val, isSet: true}
}

func (v NullableWorkItemSearchQueryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemSearchQueryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


