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

// checks if the WorkItemChangedFieldsViewModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemChangedFieldsViewModel{}

// WorkItemChangedFieldsViewModel struct for WorkItemChangedFieldsViewModel
type WorkItemChangedFieldsViewModel struct {
	Name *StringWorkItemChangedFieldViewModel `json:"name,omitempty"`
	IsDeleted *BooleanWorkItemChangedFieldViewModel `json:"isDeleted,omitempty"`
	ProjectId *GuidWorkItemChangedFieldViewModel `json:"projectId,omitempty"`
	IsAutomated *BooleanWorkItemChangedFieldViewModel `json:"isAutomated,omitempty"`
	SectionId *GuidWorkItemChangedFieldViewModel `json:"sectionId,omitempty"`
	Description *StringWorkItemChangedFieldViewModel `json:"description,omitempty"`
	State *StringWorkItemChangedFieldViewModel `json:"state,omitempty"`
	Priority *StringWorkItemChangedFieldViewModel `json:"priority,omitempty"`
	Duration *Int32WorkItemChangedFieldViewModel `json:"duration,omitempty"`
	Attributes map[string]WorkItemChangedAttributeViewModel `json:"attributes,omitempty"`
	Steps *WorkItemStepChangeViewModelArrayWorkItemChangedFieldViewModel `json:"steps,omitempty"`
	PreconditionSteps *WorkItemStepChangeViewModelArrayWorkItemChangedFieldViewModel `json:"preconditionSteps,omitempty"`
	PostconditionSteps *WorkItemStepChangeViewModelArrayWorkItemChangedFieldViewModel `json:"postconditionSteps,omitempty"`
	AutoTests *AutoTestChangeViewModelArrayWorkItemChangedFieldViewModel `json:"autoTests,omitempty"`
	Attachments *AttachmentChangeViewModelArrayWorkItemChangedFieldViewModel `json:"attachments,omitempty"`
	Tags *StringArrayWorkItemChangedFieldViewModel `json:"tags,omitempty"`
	Links *WorkItemLinkChangeViewModelArrayWorkItemChangedFieldViewModel `json:"links,omitempty"`
	GlobalId *Int64WorkItemChangedFieldViewModel `json:"globalId,omitempty"`
	VersionNumber *Int32WorkItemChangedFieldViewModel `json:"versionNumber,omitempty"`
	EntityTypeName *StringWorkItemChangedFieldViewModel `json:"entityTypeName,omitempty"`
}

// NewWorkItemChangedFieldsViewModel instantiates a new WorkItemChangedFieldsViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemChangedFieldsViewModel() *WorkItemChangedFieldsViewModel {
	this := WorkItemChangedFieldsViewModel{}
	return &this
}

// NewWorkItemChangedFieldsViewModelWithDefaults instantiates a new WorkItemChangedFieldsViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemChangedFieldsViewModelWithDefaults() *WorkItemChangedFieldsViewModel {
	this := WorkItemChangedFieldsViewModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkItemChangedFieldsViewModel) GetName() StringWorkItemChangedFieldViewModel {
	if o == nil || IsNil(o.Name) {
		var ret StringWorkItemChangedFieldViewModel
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModel) GetNameOk() (*StringWorkItemChangedFieldViewModel, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given StringWorkItemChangedFieldViewModel and assigns it to the Name field.
func (o *WorkItemChangedFieldsViewModel) SetName(v StringWorkItemChangedFieldViewModel) {
	o.Name = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *WorkItemChangedFieldsViewModel) GetIsDeleted() BooleanWorkItemChangedFieldViewModel {
	if o == nil || IsNil(o.IsDeleted) {
		var ret BooleanWorkItemChangedFieldViewModel
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModel) GetIsDeletedOk() (*BooleanWorkItemChangedFieldViewModel, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModel) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given BooleanWorkItemChangedFieldViewModel and assigns it to the IsDeleted field.
func (o *WorkItemChangedFieldsViewModel) SetIsDeleted(v BooleanWorkItemChangedFieldViewModel) {
	o.IsDeleted = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *WorkItemChangedFieldsViewModel) GetProjectId() GuidWorkItemChangedFieldViewModel {
	if o == nil || IsNil(o.ProjectId) {
		var ret GuidWorkItemChangedFieldViewModel
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModel) GetProjectIdOk() (*GuidWorkItemChangedFieldViewModel, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModel) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given GuidWorkItemChangedFieldViewModel and assigns it to the ProjectId field.
func (o *WorkItemChangedFieldsViewModel) SetProjectId(v GuidWorkItemChangedFieldViewModel) {
	o.ProjectId = &v
}

// GetIsAutomated returns the IsAutomated field value if set, zero value otherwise.
func (o *WorkItemChangedFieldsViewModel) GetIsAutomated() BooleanWorkItemChangedFieldViewModel {
	if o == nil || IsNil(o.IsAutomated) {
		var ret BooleanWorkItemChangedFieldViewModel
		return ret
	}
	return *o.IsAutomated
}

// GetIsAutomatedOk returns a tuple with the IsAutomated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModel) GetIsAutomatedOk() (*BooleanWorkItemChangedFieldViewModel, bool) {
	if o == nil || IsNil(o.IsAutomated) {
		return nil, false
	}
	return o.IsAutomated, true
}

// HasIsAutomated returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModel) HasIsAutomated() bool {
	if o != nil && !IsNil(o.IsAutomated) {
		return true
	}

	return false
}

// SetIsAutomated gets a reference to the given BooleanWorkItemChangedFieldViewModel and assigns it to the IsAutomated field.
func (o *WorkItemChangedFieldsViewModel) SetIsAutomated(v BooleanWorkItemChangedFieldViewModel) {
	o.IsAutomated = &v
}

// GetSectionId returns the SectionId field value if set, zero value otherwise.
func (o *WorkItemChangedFieldsViewModel) GetSectionId() GuidWorkItemChangedFieldViewModel {
	if o == nil || IsNil(o.SectionId) {
		var ret GuidWorkItemChangedFieldViewModel
		return ret
	}
	return *o.SectionId
}

// GetSectionIdOk returns a tuple with the SectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModel) GetSectionIdOk() (*GuidWorkItemChangedFieldViewModel, bool) {
	if o == nil || IsNil(o.SectionId) {
		return nil, false
	}
	return o.SectionId, true
}

// HasSectionId returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModel) HasSectionId() bool {
	if o != nil && !IsNil(o.SectionId) {
		return true
	}

	return false
}

// SetSectionId gets a reference to the given GuidWorkItemChangedFieldViewModel and assigns it to the SectionId field.
func (o *WorkItemChangedFieldsViewModel) SetSectionId(v GuidWorkItemChangedFieldViewModel) {
	o.SectionId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkItemChangedFieldsViewModel) GetDescription() StringWorkItemChangedFieldViewModel {
	if o == nil || IsNil(o.Description) {
		var ret StringWorkItemChangedFieldViewModel
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModel) GetDescriptionOk() (*StringWorkItemChangedFieldViewModel, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given StringWorkItemChangedFieldViewModel and assigns it to the Description field.
func (o *WorkItemChangedFieldsViewModel) SetDescription(v StringWorkItemChangedFieldViewModel) {
	o.Description = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *WorkItemChangedFieldsViewModel) GetState() StringWorkItemChangedFieldViewModel {
	if o == nil || IsNil(o.State) {
		var ret StringWorkItemChangedFieldViewModel
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModel) GetStateOk() (*StringWorkItemChangedFieldViewModel, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModel) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given StringWorkItemChangedFieldViewModel and assigns it to the State field.
func (o *WorkItemChangedFieldsViewModel) SetState(v StringWorkItemChangedFieldViewModel) {
	o.State = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *WorkItemChangedFieldsViewModel) GetPriority() StringWorkItemChangedFieldViewModel {
	if o == nil || IsNil(o.Priority) {
		var ret StringWorkItemChangedFieldViewModel
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModel) GetPriorityOk() (*StringWorkItemChangedFieldViewModel, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModel) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given StringWorkItemChangedFieldViewModel and assigns it to the Priority field.
func (o *WorkItemChangedFieldsViewModel) SetPriority(v StringWorkItemChangedFieldViewModel) {
	o.Priority = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *WorkItemChangedFieldsViewModel) GetDuration() Int32WorkItemChangedFieldViewModel {
	if o == nil || IsNil(o.Duration) {
		var ret Int32WorkItemChangedFieldViewModel
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModel) GetDurationOk() (*Int32WorkItemChangedFieldViewModel, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModel) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given Int32WorkItemChangedFieldViewModel and assigns it to the Duration field.
func (o *WorkItemChangedFieldsViewModel) SetDuration(v Int32WorkItemChangedFieldViewModel) {
	o.Duration = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemChangedFieldsViewModel) GetAttributes() map[string]WorkItemChangedAttributeViewModel {
	if o == nil {
		var ret map[string]WorkItemChangedAttributeViewModel
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemChangedFieldsViewModel) GetAttributesOk() (*map[string]WorkItemChangedAttributeViewModel, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModel) HasAttributes() bool {
	if o != nil && IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]WorkItemChangedAttributeViewModel and assigns it to the Attributes field.
func (o *WorkItemChangedFieldsViewModel) SetAttributes(v map[string]WorkItemChangedAttributeViewModel) {
	o.Attributes = v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *WorkItemChangedFieldsViewModel) GetSteps() WorkItemStepChangeViewModelArrayWorkItemChangedFieldViewModel {
	if o == nil || IsNil(o.Steps) {
		var ret WorkItemStepChangeViewModelArrayWorkItemChangedFieldViewModel
		return ret
	}
	return *o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModel) GetStepsOk() (*WorkItemStepChangeViewModelArrayWorkItemChangedFieldViewModel, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModel) HasSteps() bool {
	if o != nil && !IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given WorkItemStepChangeViewModelArrayWorkItemChangedFieldViewModel and assigns it to the Steps field.
func (o *WorkItemChangedFieldsViewModel) SetSteps(v WorkItemStepChangeViewModelArrayWorkItemChangedFieldViewModel) {
	o.Steps = &v
}

// GetPreconditionSteps returns the PreconditionSteps field value if set, zero value otherwise.
func (o *WorkItemChangedFieldsViewModel) GetPreconditionSteps() WorkItemStepChangeViewModelArrayWorkItemChangedFieldViewModel {
	if o == nil || IsNil(o.PreconditionSteps) {
		var ret WorkItemStepChangeViewModelArrayWorkItemChangedFieldViewModel
		return ret
	}
	return *o.PreconditionSteps
}

// GetPreconditionStepsOk returns a tuple with the PreconditionSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModel) GetPreconditionStepsOk() (*WorkItemStepChangeViewModelArrayWorkItemChangedFieldViewModel, bool) {
	if o == nil || IsNil(o.PreconditionSteps) {
		return nil, false
	}
	return o.PreconditionSteps, true
}

// HasPreconditionSteps returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModel) HasPreconditionSteps() bool {
	if o != nil && !IsNil(o.PreconditionSteps) {
		return true
	}

	return false
}

// SetPreconditionSteps gets a reference to the given WorkItemStepChangeViewModelArrayWorkItemChangedFieldViewModel and assigns it to the PreconditionSteps field.
func (o *WorkItemChangedFieldsViewModel) SetPreconditionSteps(v WorkItemStepChangeViewModelArrayWorkItemChangedFieldViewModel) {
	o.PreconditionSteps = &v
}

// GetPostconditionSteps returns the PostconditionSteps field value if set, zero value otherwise.
func (o *WorkItemChangedFieldsViewModel) GetPostconditionSteps() WorkItemStepChangeViewModelArrayWorkItemChangedFieldViewModel {
	if o == nil || IsNil(o.PostconditionSteps) {
		var ret WorkItemStepChangeViewModelArrayWorkItemChangedFieldViewModel
		return ret
	}
	return *o.PostconditionSteps
}

// GetPostconditionStepsOk returns a tuple with the PostconditionSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModel) GetPostconditionStepsOk() (*WorkItemStepChangeViewModelArrayWorkItemChangedFieldViewModel, bool) {
	if o == nil || IsNil(o.PostconditionSteps) {
		return nil, false
	}
	return o.PostconditionSteps, true
}

// HasPostconditionSteps returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModel) HasPostconditionSteps() bool {
	if o != nil && !IsNil(o.PostconditionSteps) {
		return true
	}

	return false
}

// SetPostconditionSteps gets a reference to the given WorkItemStepChangeViewModelArrayWorkItemChangedFieldViewModel and assigns it to the PostconditionSteps field.
func (o *WorkItemChangedFieldsViewModel) SetPostconditionSteps(v WorkItemStepChangeViewModelArrayWorkItemChangedFieldViewModel) {
	o.PostconditionSteps = &v
}

// GetAutoTests returns the AutoTests field value if set, zero value otherwise.
func (o *WorkItemChangedFieldsViewModel) GetAutoTests() AutoTestChangeViewModelArrayWorkItemChangedFieldViewModel {
	if o == nil || IsNil(o.AutoTests) {
		var ret AutoTestChangeViewModelArrayWorkItemChangedFieldViewModel
		return ret
	}
	return *o.AutoTests
}

// GetAutoTestsOk returns a tuple with the AutoTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModel) GetAutoTestsOk() (*AutoTestChangeViewModelArrayWorkItemChangedFieldViewModel, bool) {
	if o == nil || IsNil(o.AutoTests) {
		return nil, false
	}
	return o.AutoTests, true
}

// HasAutoTests returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModel) HasAutoTests() bool {
	if o != nil && !IsNil(o.AutoTests) {
		return true
	}

	return false
}

// SetAutoTests gets a reference to the given AutoTestChangeViewModelArrayWorkItemChangedFieldViewModel and assigns it to the AutoTests field.
func (o *WorkItemChangedFieldsViewModel) SetAutoTests(v AutoTestChangeViewModelArrayWorkItemChangedFieldViewModel) {
	o.AutoTests = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *WorkItemChangedFieldsViewModel) GetAttachments() AttachmentChangeViewModelArrayWorkItemChangedFieldViewModel {
	if o == nil || IsNil(o.Attachments) {
		var ret AttachmentChangeViewModelArrayWorkItemChangedFieldViewModel
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModel) GetAttachmentsOk() (*AttachmentChangeViewModelArrayWorkItemChangedFieldViewModel, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModel) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AttachmentChangeViewModelArrayWorkItemChangedFieldViewModel and assigns it to the Attachments field.
func (o *WorkItemChangedFieldsViewModel) SetAttachments(v AttachmentChangeViewModelArrayWorkItemChangedFieldViewModel) {
	o.Attachments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WorkItemChangedFieldsViewModel) GetTags() StringArrayWorkItemChangedFieldViewModel {
	if o == nil || IsNil(o.Tags) {
		var ret StringArrayWorkItemChangedFieldViewModel
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModel) GetTagsOk() (*StringArrayWorkItemChangedFieldViewModel, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModel) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given StringArrayWorkItemChangedFieldViewModel and assigns it to the Tags field.
func (o *WorkItemChangedFieldsViewModel) SetTags(v StringArrayWorkItemChangedFieldViewModel) {
	o.Tags = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *WorkItemChangedFieldsViewModel) GetLinks() WorkItemLinkChangeViewModelArrayWorkItemChangedFieldViewModel {
	if o == nil || IsNil(o.Links) {
		var ret WorkItemLinkChangeViewModelArrayWorkItemChangedFieldViewModel
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModel) GetLinksOk() (*WorkItemLinkChangeViewModelArrayWorkItemChangedFieldViewModel, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given WorkItemLinkChangeViewModelArrayWorkItemChangedFieldViewModel and assigns it to the Links field.
func (o *WorkItemChangedFieldsViewModel) SetLinks(v WorkItemLinkChangeViewModelArrayWorkItemChangedFieldViewModel) {
	o.Links = &v
}

// GetGlobalId returns the GlobalId field value if set, zero value otherwise.
func (o *WorkItemChangedFieldsViewModel) GetGlobalId() Int64WorkItemChangedFieldViewModel {
	if o == nil || IsNil(o.GlobalId) {
		var ret Int64WorkItemChangedFieldViewModel
		return ret
	}
	return *o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModel) GetGlobalIdOk() (*Int64WorkItemChangedFieldViewModel, bool) {
	if o == nil || IsNil(o.GlobalId) {
		return nil, false
	}
	return o.GlobalId, true
}

// HasGlobalId returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModel) HasGlobalId() bool {
	if o != nil && !IsNil(o.GlobalId) {
		return true
	}

	return false
}

// SetGlobalId gets a reference to the given Int64WorkItemChangedFieldViewModel and assigns it to the GlobalId field.
func (o *WorkItemChangedFieldsViewModel) SetGlobalId(v Int64WorkItemChangedFieldViewModel) {
	o.GlobalId = &v
}

// GetVersionNumber returns the VersionNumber field value if set, zero value otherwise.
func (o *WorkItemChangedFieldsViewModel) GetVersionNumber() Int32WorkItemChangedFieldViewModel {
	if o == nil || IsNil(o.VersionNumber) {
		var ret Int32WorkItemChangedFieldViewModel
		return ret
	}
	return *o.VersionNumber
}

// GetVersionNumberOk returns a tuple with the VersionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModel) GetVersionNumberOk() (*Int32WorkItemChangedFieldViewModel, bool) {
	if o == nil || IsNil(o.VersionNumber) {
		return nil, false
	}
	return o.VersionNumber, true
}

// HasVersionNumber returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModel) HasVersionNumber() bool {
	if o != nil && !IsNil(o.VersionNumber) {
		return true
	}

	return false
}

// SetVersionNumber gets a reference to the given Int32WorkItemChangedFieldViewModel and assigns it to the VersionNumber field.
func (o *WorkItemChangedFieldsViewModel) SetVersionNumber(v Int32WorkItemChangedFieldViewModel) {
	o.VersionNumber = &v
}

// GetEntityTypeName returns the EntityTypeName field value if set, zero value otherwise.
func (o *WorkItemChangedFieldsViewModel) GetEntityTypeName() StringWorkItemChangedFieldViewModel {
	if o == nil || IsNil(o.EntityTypeName) {
		var ret StringWorkItemChangedFieldViewModel
		return ret
	}
	return *o.EntityTypeName
}

// GetEntityTypeNameOk returns a tuple with the EntityTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemChangedFieldsViewModel) GetEntityTypeNameOk() (*StringWorkItemChangedFieldViewModel, bool) {
	if o == nil || IsNil(o.EntityTypeName) {
		return nil, false
	}
	return o.EntityTypeName, true
}

// HasEntityTypeName returns a boolean if a field has been set.
func (o *WorkItemChangedFieldsViewModel) HasEntityTypeName() bool {
	if o != nil && !IsNil(o.EntityTypeName) {
		return true
	}

	return false
}

// SetEntityTypeName gets a reference to the given StringWorkItemChangedFieldViewModel and assigns it to the EntityTypeName field.
func (o *WorkItemChangedFieldsViewModel) SetEntityTypeName(v StringWorkItemChangedFieldViewModel) {
	o.EntityTypeName = &v
}

func (o WorkItemChangedFieldsViewModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemChangedFieldsViewModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.IsAutomated) {
		toSerialize["isAutomated"] = o.IsAutomated
	}
	if !IsNil(o.SectionId) {
		toSerialize["sectionId"] = o.SectionId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}
	if !IsNil(o.PreconditionSteps) {
		toSerialize["preconditionSteps"] = o.PreconditionSteps
	}
	if !IsNil(o.PostconditionSteps) {
		toSerialize["postconditionSteps"] = o.PostconditionSteps
	}
	if !IsNil(o.AutoTests) {
		toSerialize["autoTests"] = o.AutoTests
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.GlobalId) {
		toSerialize["globalId"] = o.GlobalId
	}
	if !IsNil(o.VersionNumber) {
		toSerialize["versionNumber"] = o.VersionNumber
	}
	if !IsNil(o.EntityTypeName) {
		toSerialize["entityTypeName"] = o.EntityTypeName
	}
	return toSerialize, nil
}

type NullableWorkItemChangedFieldsViewModel struct {
	value *WorkItemChangedFieldsViewModel
	isSet bool
}

func (v NullableWorkItemChangedFieldsViewModel) Get() *WorkItemChangedFieldsViewModel {
	return v.value
}

func (v *NullableWorkItemChangedFieldsViewModel) Set(val *WorkItemChangedFieldsViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemChangedFieldsViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemChangedFieldsViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemChangedFieldsViewModel(val *WorkItemChangedFieldsViewModel) *NullableWorkItemChangedFieldsViewModel {
	return &NullableWorkItemChangedFieldsViewModel{value: val, isSet: true}
}

func (v NullableWorkItemChangedFieldsViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemChangedFieldsViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


