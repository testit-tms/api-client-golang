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

// checks if the WorkItemModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemModel{}

// WorkItemModel struct for WorkItemModel
type WorkItemModel struct {
	// used for versioning changes in workitem
	VersionId string `json:"versionId"`
	// used for getting a median duration of all autotests related to this workitem
	MedianDuration int64 `json:"medianDuration"`
	IsDeleted bool `json:"isDeleted"`
	ProjectId string `json:"projectId"`
	EntityTypeName WorkItemEntityTypes `json:"entityTypeName"`
	IsAutomated bool `json:"isAutomated"`
	AutoTests []AutoTestModel `json:"autoTests,omitempty"`
	Attachments []AttachmentModel `json:"attachments,omitempty"`
	SectionPreconditionSteps []StepModel `json:"sectionPreconditionSteps,omitempty"`
	SectionPostconditionSteps []StepModel `json:"sectionPostconditionSteps,omitempty"`
	// used for define chronology of workitem state in each version
	VersionNumber int32 `json:"versionNumber"`
	Iterations []IterationModel `json:"iterations,omitempty"`
	CreatedDate time.Time `json:"createdDate"`
	ModifiedDate NullableTime `json:"modifiedDate,omitempty"`
	CreatedById string `json:"createdById"`
	ModifiedById NullableString `json:"modifiedById,omitempty"`
	GlobalId int64 `json:"globalId"`
	Id string `json:"id"`
	SectionId string `json:"sectionId"`
	Description NullableString `json:"description,omitempty"`
	State WorkItemStates `json:"state"`
	Priority WorkItemPriorityModel `json:"priority"`
	SourceType WorkItemSourceTypeModel `json:"sourceType"`
	Steps []StepModel `json:"steps"`
	PreconditionSteps []StepModel `json:"preconditionSteps"`
	PostconditionSteps []StepModel `json:"postconditionSteps"`
	Duration int32 `json:"duration"`
	Attributes map[string]interface{} `json:"attributes"`
	Tags []TagModel `json:"tags"`
	Links []LinkModel `json:"links"`
	Name string `json:"name"`
}

type _WorkItemModel WorkItemModel

// NewWorkItemModel instantiates a new WorkItemModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemModel(versionId string, medianDuration int64, isDeleted bool, projectId string, entityTypeName WorkItemEntityTypes, isAutomated bool, versionNumber int32, createdDate time.Time, createdById string, globalId int64, id string, sectionId string, state WorkItemStates, priority WorkItemPriorityModel, sourceType WorkItemSourceTypeModel, steps []StepModel, preconditionSteps []StepModel, postconditionSteps []StepModel, duration int32, attributes map[string]interface{}, tags []TagModel, links []LinkModel, name string) *WorkItemModel {
	this := WorkItemModel{}
	this.VersionId = versionId
	this.MedianDuration = medianDuration
	this.IsDeleted = isDeleted
	this.ProjectId = projectId
	this.EntityTypeName = entityTypeName
	this.IsAutomated = isAutomated
	this.VersionNumber = versionNumber
	this.CreatedDate = createdDate
	this.CreatedById = createdById
	this.GlobalId = globalId
	this.Id = id
	this.SectionId = sectionId
	this.State = state
	this.Priority = priority
	this.SourceType = sourceType
	this.Steps = steps
	this.PreconditionSteps = preconditionSteps
	this.PostconditionSteps = postconditionSteps
	this.Duration = duration
	this.Attributes = attributes
	this.Tags = tags
	this.Links = links
	this.Name = name
	return &this
}

// NewWorkItemModelWithDefaults instantiates a new WorkItemModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemModelWithDefaults() *WorkItemModel {
	this := WorkItemModel{}
	return &this
}

// GetVersionId returns the VersionId field value
func (o *WorkItemModel) GetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *WorkItemModel) SetVersionId(v string) {
	o.VersionId = v
}

// GetMedianDuration returns the MedianDuration field value
func (o *WorkItemModel) GetMedianDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MedianDuration
}

// GetMedianDurationOk returns a tuple with the MedianDuration field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetMedianDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MedianDuration, true
}

// SetMedianDuration sets field value
func (o *WorkItemModel) SetMedianDuration(v int64) {
	o.MedianDuration = v
}

// GetIsDeleted returns the IsDeleted field value
func (o *WorkItemModel) GetIsDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeleted, true
}

// SetIsDeleted sets field value
func (o *WorkItemModel) SetIsDeleted(v bool) {
	o.IsDeleted = v
}

// GetProjectId returns the ProjectId field value
func (o *WorkItemModel) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *WorkItemModel) SetProjectId(v string) {
	o.ProjectId = v
}

// GetEntityTypeName returns the EntityTypeName field value
func (o *WorkItemModel) GetEntityTypeName() WorkItemEntityTypes {
	if o == nil {
		var ret WorkItemEntityTypes
		return ret
	}

	return o.EntityTypeName
}

// GetEntityTypeNameOk returns a tuple with the EntityTypeName field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetEntityTypeNameOk() (*WorkItemEntityTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityTypeName, true
}

// SetEntityTypeName sets field value
func (o *WorkItemModel) SetEntityTypeName(v WorkItemEntityTypes) {
	o.EntityTypeName = v
}

// GetIsAutomated returns the IsAutomated field value
func (o *WorkItemModel) GetIsAutomated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAutomated
}

// GetIsAutomatedOk returns a tuple with the IsAutomated field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetIsAutomatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAutomated, true
}

// SetIsAutomated sets field value
func (o *WorkItemModel) SetIsAutomated(v bool) {
	o.IsAutomated = v
}

// GetAutoTests returns the AutoTests field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemModel) GetAutoTests() []AutoTestModel {
	if o == nil {
		var ret []AutoTestModel
		return ret
	}
	return o.AutoTests
}

// GetAutoTestsOk returns a tuple with the AutoTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemModel) GetAutoTestsOk() ([]AutoTestModel, bool) {
	if o == nil || IsNil(o.AutoTests) {
		return nil, false
	}
	return o.AutoTests, true
}

// HasAutoTests returns a boolean if a field has been set.
func (o *WorkItemModel) HasAutoTests() bool {
	if o != nil && !IsNil(o.AutoTests) {
		return true
	}

	return false
}

// SetAutoTests gets a reference to the given []AutoTestModel and assigns it to the AutoTests field.
func (o *WorkItemModel) SetAutoTests(v []AutoTestModel) {
	o.AutoTests = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemModel) GetAttachments() []AttachmentModel {
	if o == nil {
		var ret []AttachmentModel
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemModel) GetAttachmentsOk() ([]AttachmentModel, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *WorkItemModel) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []AttachmentModel and assigns it to the Attachments field.
func (o *WorkItemModel) SetAttachments(v []AttachmentModel) {
	o.Attachments = v
}

// GetSectionPreconditionSteps returns the SectionPreconditionSteps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemModel) GetSectionPreconditionSteps() []StepModel {
	if o == nil {
		var ret []StepModel
		return ret
	}
	return o.SectionPreconditionSteps
}

// GetSectionPreconditionStepsOk returns a tuple with the SectionPreconditionSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemModel) GetSectionPreconditionStepsOk() ([]StepModel, bool) {
	if o == nil || IsNil(o.SectionPreconditionSteps) {
		return nil, false
	}
	return o.SectionPreconditionSteps, true
}

// HasSectionPreconditionSteps returns a boolean if a field has been set.
func (o *WorkItemModel) HasSectionPreconditionSteps() bool {
	if o != nil && !IsNil(o.SectionPreconditionSteps) {
		return true
	}

	return false
}

// SetSectionPreconditionSteps gets a reference to the given []StepModel and assigns it to the SectionPreconditionSteps field.
func (o *WorkItemModel) SetSectionPreconditionSteps(v []StepModel) {
	o.SectionPreconditionSteps = v
}

// GetSectionPostconditionSteps returns the SectionPostconditionSteps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemModel) GetSectionPostconditionSteps() []StepModel {
	if o == nil {
		var ret []StepModel
		return ret
	}
	return o.SectionPostconditionSteps
}

// GetSectionPostconditionStepsOk returns a tuple with the SectionPostconditionSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemModel) GetSectionPostconditionStepsOk() ([]StepModel, bool) {
	if o == nil || IsNil(o.SectionPostconditionSteps) {
		return nil, false
	}
	return o.SectionPostconditionSteps, true
}

// HasSectionPostconditionSteps returns a boolean if a field has been set.
func (o *WorkItemModel) HasSectionPostconditionSteps() bool {
	if o != nil && !IsNil(o.SectionPostconditionSteps) {
		return true
	}

	return false
}

// SetSectionPostconditionSteps gets a reference to the given []StepModel and assigns it to the SectionPostconditionSteps field.
func (o *WorkItemModel) SetSectionPostconditionSteps(v []StepModel) {
	o.SectionPostconditionSteps = v
}

// GetVersionNumber returns the VersionNumber field value
func (o *WorkItemModel) GetVersionNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VersionNumber
}

// GetVersionNumberOk returns a tuple with the VersionNumber field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetVersionNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionNumber, true
}

// SetVersionNumber sets field value
func (o *WorkItemModel) SetVersionNumber(v int32) {
	o.VersionNumber = v
}

// GetIterations returns the Iterations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemModel) GetIterations() []IterationModel {
	if o == nil {
		var ret []IterationModel
		return ret
	}
	return o.Iterations
}

// GetIterationsOk returns a tuple with the Iterations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemModel) GetIterationsOk() ([]IterationModel, bool) {
	if o == nil || IsNil(o.Iterations) {
		return nil, false
	}
	return o.Iterations, true
}

// HasIterations returns a boolean if a field has been set.
func (o *WorkItemModel) HasIterations() bool {
	if o != nil && !IsNil(o.Iterations) {
		return true
	}

	return false
}

// SetIterations gets a reference to the given []IterationModel and assigns it to the Iterations field.
func (o *WorkItemModel) SetIterations(v []IterationModel) {
	o.Iterations = v
}

// GetCreatedDate returns the CreatedDate field value
func (o *WorkItemModel) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *WorkItemModel) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemModel) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemModel) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *WorkItemModel) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTime and assigns it to the ModifiedDate field.
func (o *WorkItemModel) SetModifiedDate(v time.Time) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *WorkItemModel) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *WorkItemModel) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetCreatedById returns the CreatedById field value
func (o *WorkItemModel) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *WorkItemModel) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemModel) GetModifiedById() string {
	if o == nil || IsNil(o.ModifiedById.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedById.Get()
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemModel) GetModifiedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedById.Get(), o.ModifiedById.IsSet()
}

// HasModifiedById returns a boolean if a field has been set.
func (o *WorkItemModel) HasModifiedById() bool {
	if o != nil && o.ModifiedById.IsSet() {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given NullableString and assigns it to the ModifiedById field.
func (o *WorkItemModel) SetModifiedById(v string) {
	o.ModifiedById.Set(&v)
}
// SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil
func (o *WorkItemModel) SetModifiedByIdNil() {
	o.ModifiedById.Set(nil)
}

// UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
func (o *WorkItemModel) UnsetModifiedById() {
	o.ModifiedById.Unset()
}

// GetGlobalId returns the GlobalId field value
func (o *WorkItemModel) GetGlobalId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetGlobalIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalId, true
}

// SetGlobalId sets field value
func (o *WorkItemModel) SetGlobalId(v int64) {
	o.GlobalId = v
}

// GetId returns the Id field value
func (o *WorkItemModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkItemModel) SetId(v string) {
	o.Id = v
}

// GetSectionId returns the SectionId field value
func (o *WorkItemModel) GetSectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SectionId
}

// GetSectionIdOk returns a tuple with the SectionId field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetSectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SectionId, true
}

// SetSectionId sets field value
func (o *WorkItemModel) SetSectionId(v string) {
	o.SectionId = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkItemModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *WorkItemModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *WorkItemModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *WorkItemModel) UnsetDescription() {
	o.Description.Unset()
}

// GetState returns the State field value
func (o *WorkItemModel) GetState() WorkItemStates {
	if o == nil {
		var ret WorkItemStates
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetStateOk() (*WorkItemStates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *WorkItemModel) SetState(v WorkItemStates) {
	o.State = v
}

// GetPriority returns the Priority field value
func (o *WorkItemModel) GetPriority() WorkItemPriorityModel {
	if o == nil {
		var ret WorkItemPriorityModel
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetPriorityOk() (*WorkItemPriorityModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *WorkItemModel) SetPriority(v WorkItemPriorityModel) {
	o.Priority = v
}

// GetSourceType returns the SourceType field value
func (o *WorkItemModel) GetSourceType() WorkItemSourceTypeModel {
	if o == nil {
		var ret WorkItemSourceTypeModel
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetSourceTypeOk() (*WorkItemSourceTypeModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *WorkItemModel) SetSourceType(v WorkItemSourceTypeModel) {
	o.SourceType = v
}

// GetSteps returns the Steps field value
func (o *WorkItemModel) GetSteps() []StepModel {
	if o == nil {
		var ret []StepModel
		return ret
	}

	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetStepsOk() ([]StepModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Steps, true
}

// SetSteps sets field value
func (o *WorkItemModel) SetSteps(v []StepModel) {
	o.Steps = v
}

// GetPreconditionSteps returns the PreconditionSteps field value
func (o *WorkItemModel) GetPreconditionSteps() []StepModel {
	if o == nil {
		var ret []StepModel
		return ret
	}

	return o.PreconditionSteps
}

// GetPreconditionStepsOk returns a tuple with the PreconditionSteps field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetPreconditionStepsOk() ([]StepModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreconditionSteps, true
}

// SetPreconditionSteps sets field value
func (o *WorkItemModel) SetPreconditionSteps(v []StepModel) {
	o.PreconditionSteps = v
}

// GetPostconditionSteps returns the PostconditionSteps field value
func (o *WorkItemModel) GetPostconditionSteps() []StepModel {
	if o == nil {
		var ret []StepModel
		return ret
	}

	return o.PostconditionSteps
}

// GetPostconditionStepsOk returns a tuple with the PostconditionSteps field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetPostconditionStepsOk() ([]StepModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostconditionSteps, true
}

// SetPostconditionSteps sets field value
func (o *WorkItemModel) SetPostconditionSteps(v []StepModel) {
	o.PostconditionSteps = v
}

// GetDuration returns the Duration field value
func (o *WorkItemModel) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *WorkItemModel) SetDuration(v int32) {
	o.Duration = v
}

// GetAttributes returns the Attributes field value
func (o *WorkItemModel) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *WorkItemModel) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetTags returns the Tags field value
func (o *WorkItemModel) GetTags() []TagModel {
	if o == nil {
		var ret []TagModel
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetTagsOk() ([]TagModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *WorkItemModel) SetTags(v []TagModel) {
	o.Tags = v
}

// GetLinks returns the Links field value
func (o *WorkItemModel) GetLinks() []LinkModel {
	if o == nil {
		var ret []LinkModel
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetLinksOk() ([]LinkModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *WorkItemModel) SetLinks(v []LinkModel) {
	o.Links = v
}

// GetName returns the Name field value
func (o *WorkItemModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkItemModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkItemModel) SetName(v string) {
	o.Name = v
}

func (o WorkItemModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["versionId"] = o.VersionId
	toSerialize["medianDuration"] = o.MedianDuration
	toSerialize["isDeleted"] = o.IsDeleted
	toSerialize["projectId"] = o.ProjectId
	toSerialize["entityTypeName"] = o.EntityTypeName
	toSerialize["isAutomated"] = o.IsAutomated
	if o.AutoTests != nil {
		toSerialize["autoTests"] = o.AutoTests
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.SectionPreconditionSteps != nil {
		toSerialize["sectionPreconditionSteps"] = o.SectionPreconditionSteps
	}
	if o.SectionPostconditionSteps != nil {
		toSerialize["sectionPostconditionSteps"] = o.SectionPostconditionSteps
	}
	toSerialize["versionNumber"] = o.VersionNumber
	if o.Iterations != nil {
		toSerialize["iterations"] = o.Iterations
	}
	toSerialize["createdDate"] = o.CreatedDate
	if o.ModifiedDate.IsSet() {
		toSerialize["modifiedDate"] = o.ModifiedDate.Get()
	}
	toSerialize["createdById"] = o.CreatedById
	if o.ModifiedById.IsSet() {
		toSerialize["modifiedById"] = o.ModifiedById.Get()
	}
	toSerialize["globalId"] = o.GlobalId
	toSerialize["id"] = o.Id
	toSerialize["sectionId"] = o.SectionId
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["state"] = o.State
	toSerialize["priority"] = o.Priority
	toSerialize["sourceType"] = o.SourceType
	toSerialize["steps"] = o.Steps
	toSerialize["preconditionSteps"] = o.PreconditionSteps
	toSerialize["postconditionSteps"] = o.PostconditionSteps
	toSerialize["duration"] = o.Duration
	toSerialize["attributes"] = o.Attributes
	toSerialize["tags"] = o.Tags
	toSerialize["links"] = o.Links
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *WorkItemModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"versionId",
		"medianDuration",
		"isDeleted",
		"projectId",
		"entityTypeName",
		"isAutomated",
		"versionNumber",
		"createdDate",
		"createdById",
		"globalId",
		"id",
		"sectionId",
		"state",
		"priority",
		"sourceType",
		"steps",
		"preconditionSteps",
		"postconditionSteps",
		"duration",
		"attributes",
		"tags",
		"links",
		"name",
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

	varWorkItemModel := _WorkItemModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkItemModel)

	if err != nil {
		return err
	}

	*o = WorkItemModel(varWorkItemModel)

	return err
}

type NullableWorkItemModel struct {
	value *WorkItemModel
	isSet bool
}

func (v NullableWorkItemModel) Get() *WorkItemModel {
	return v.value
}

func (v *NullableWorkItemModel) Set(val *WorkItemModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemModel(val *WorkItemModel) *NullableWorkItemModel {
	return &NullableWorkItemModel{value: val, isSet: true}
}

func (v NullableWorkItemModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


