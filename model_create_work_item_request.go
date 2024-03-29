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

// checks if the CreateWorkItemRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWorkItemRequest{}

// CreateWorkItemRequest struct for CreateWorkItemRequest
type CreateWorkItemRequest struct {
	EntityTypeName WorkItemEntityTypes `json:"entityTypeName"`
	Description NullableString `json:"description,omitempty"`
	State WorkItemStates `json:"state"`
	Priority WorkItemPriorityModel `json:"priority"`
	Steps []StepPostModel `json:"steps"`
	PreconditionSteps []StepPostModel `json:"preconditionSteps"`
	PostconditionSteps []StepPostModel `json:"postconditionSteps"`
	// Must be 0 for shared steps and greater than 0 for the other types of work items
	Duration int32 `json:"duration"`
	Attributes map[string]interface{} `json:"attributes"`
	Tags []TagPostModel `json:"tags"`
	Attachments []AttachmentPutModel `json:"attachments,omitempty"`
	Iterations []IterationPutModel `json:"iterations,omitempty"`
	Links []LinkPostModel `json:"links"`
	Name string `json:"name"`
	// This property is used to link workitem with project
	ProjectId string `json:"projectId"`
	SectionId string `json:"sectionId"`
	AutoTests []AutoTestIdModel `json:"autoTests,omitempty"`
}

// NewCreateWorkItemRequest instantiates a new CreateWorkItemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWorkItemRequest(entityTypeName WorkItemEntityTypes, state WorkItemStates, priority WorkItemPriorityModel, steps []StepPostModel, preconditionSteps []StepPostModel, postconditionSteps []StepPostModel, duration int32, attributes map[string]interface{}, tags []TagPostModel, links []LinkPostModel, name string, projectId string, sectionId string) *CreateWorkItemRequest {
	this := CreateWorkItemRequest{}
	this.EntityTypeName = entityTypeName
	this.State = state
	this.Priority = priority
	this.Steps = steps
	this.PreconditionSteps = preconditionSteps
	this.PostconditionSteps = postconditionSteps
	this.Duration = duration
	this.Attributes = attributes
	this.Tags = tags
	this.Links = links
	this.Name = name
	this.ProjectId = projectId
	this.SectionId = sectionId
	return &this
}

// NewCreateWorkItemRequestWithDefaults instantiates a new CreateWorkItemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWorkItemRequestWithDefaults() *CreateWorkItemRequest {
	this := CreateWorkItemRequest{}
	return &this
}

// GetEntityTypeName returns the EntityTypeName field value
func (o *CreateWorkItemRequest) GetEntityTypeName() WorkItemEntityTypes {
	if o == nil {
		var ret WorkItemEntityTypes
		return ret
	}

	return o.EntityTypeName
}

// GetEntityTypeNameOk returns a tuple with the EntityTypeName field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemRequest) GetEntityTypeNameOk() (*WorkItemEntityTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityTypeName, true
}

// SetEntityTypeName sets field value
func (o *CreateWorkItemRequest) SetEntityTypeName(v WorkItemEntityTypes) {
	o.EntityTypeName = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateWorkItemRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateWorkItemRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateWorkItemRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateWorkItemRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateWorkItemRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateWorkItemRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetState returns the State field value
func (o *CreateWorkItemRequest) GetState() WorkItemStates {
	if o == nil {
		var ret WorkItemStates
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemRequest) GetStateOk() (*WorkItemStates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CreateWorkItemRequest) SetState(v WorkItemStates) {
	o.State = v
}

// GetPriority returns the Priority field value
func (o *CreateWorkItemRequest) GetPriority() WorkItemPriorityModel {
	if o == nil {
		var ret WorkItemPriorityModel
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemRequest) GetPriorityOk() (*WorkItemPriorityModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *CreateWorkItemRequest) SetPriority(v WorkItemPriorityModel) {
	o.Priority = v
}

// GetSteps returns the Steps field value
func (o *CreateWorkItemRequest) GetSteps() []StepPostModel {
	if o == nil {
		var ret []StepPostModel
		return ret
	}

	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemRequest) GetStepsOk() ([]StepPostModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Steps, true
}

// SetSteps sets field value
func (o *CreateWorkItemRequest) SetSteps(v []StepPostModel) {
	o.Steps = v
}

// GetPreconditionSteps returns the PreconditionSteps field value
func (o *CreateWorkItemRequest) GetPreconditionSteps() []StepPostModel {
	if o == nil {
		var ret []StepPostModel
		return ret
	}

	return o.PreconditionSteps
}

// GetPreconditionStepsOk returns a tuple with the PreconditionSteps field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemRequest) GetPreconditionStepsOk() ([]StepPostModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreconditionSteps, true
}

// SetPreconditionSteps sets field value
func (o *CreateWorkItemRequest) SetPreconditionSteps(v []StepPostModel) {
	o.PreconditionSteps = v
}

// GetPostconditionSteps returns the PostconditionSteps field value
func (o *CreateWorkItemRequest) GetPostconditionSteps() []StepPostModel {
	if o == nil {
		var ret []StepPostModel
		return ret
	}

	return o.PostconditionSteps
}

// GetPostconditionStepsOk returns a tuple with the PostconditionSteps field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemRequest) GetPostconditionStepsOk() ([]StepPostModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostconditionSteps, true
}

// SetPostconditionSteps sets field value
func (o *CreateWorkItemRequest) SetPostconditionSteps(v []StepPostModel) {
	o.PostconditionSteps = v
}

// GetDuration returns the Duration field value
func (o *CreateWorkItemRequest) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemRequest) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *CreateWorkItemRequest) SetDuration(v int32) {
	o.Duration = v
}

// GetAttributes returns the Attributes field value
func (o *CreateWorkItemRequest) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemRequest) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *CreateWorkItemRequest) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetTags returns the Tags field value
func (o *CreateWorkItemRequest) GetTags() []TagPostModel {
	if o == nil {
		var ret []TagPostModel
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemRequest) GetTagsOk() ([]TagPostModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *CreateWorkItemRequest) SetTags(v []TagPostModel) {
	o.Tags = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateWorkItemRequest) GetAttachments() []AttachmentPutModel {
	if o == nil {
		var ret []AttachmentPutModel
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateWorkItemRequest) GetAttachmentsOk() ([]AttachmentPutModel, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *CreateWorkItemRequest) HasAttachments() bool {
	if o != nil && IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []AttachmentPutModel and assigns it to the Attachments field.
func (o *CreateWorkItemRequest) SetAttachments(v []AttachmentPutModel) {
	o.Attachments = v
}

// GetIterations returns the Iterations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateWorkItemRequest) GetIterations() []IterationPutModel {
	if o == nil {
		var ret []IterationPutModel
		return ret
	}
	return o.Iterations
}

// GetIterationsOk returns a tuple with the Iterations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateWorkItemRequest) GetIterationsOk() ([]IterationPutModel, bool) {
	if o == nil || IsNil(o.Iterations) {
		return nil, false
	}
	return o.Iterations, true
}

// HasIterations returns a boolean if a field has been set.
func (o *CreateWorkItemRequest) HasIterations() bool {
	if o != nil && IsNil(o.Iterations) {
		return true
	}

	return false
}

// SetIterations gets a reference to the given []IterationPutModel and assigns it to the Iterations field.
func (o *CreateWorkItemRequest) SetIterations(v []IterationPutModel) {
	o.Iterations = v
}

// GetLinks returns the Links field value
func (o *CreateWorkItemRequest) GetLinks() []LinkPostModel {
	if o == nil {
		var ret []LinkPostModel
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemRequest) GetLinksOk() ([]LinkPostModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *CreateWorkItemRequest) SetLinks(v []LinkPostModel) {
	o.Links = v
}

// GetName returns the Name field value
func (o *CreateWorkItemRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateWorkItemRequest) SetName(v string) {
	o.Name = v
}

// GetProjectId returns the ProjectId field value
func (o *CreateWorkItemRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemRequest) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *CreateWorkItemRequest) SetProjectId(v string) {
	o.ProjectId = v
}

// GetSectionId returns the SectionId field value
func (o *CreateWorkItemRequest) GetSectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SectionId
}

// GetSectionIdOk returns a tuple with the SectionId field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemRequest) GetSectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SectionId, true
}

// SetSectionId sets field value
func (o *CreateWorkItemRequest) SetSectionId(v string) {
	o.SectionId = v
}

// GetAutoTests returns the AutoTests field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateWorkItemRequest) GetAutoTests() []AutoTestIdModel {
	if o == nil {
		var ret []AutoTestIdModel
		return ret
	}
	return o.AutoTests
}

// GetAutoTestsOk returns a tuple with the AutoTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateWorkItemRequest) GetAutoTestsOk() ([]AutoTestIdModel, bool) {
	if o == nil || IsNil(o.AutoTests) {
		return nil, false
	}
	return o.AutoTests, true
}

// HasAutoTests returns a boolean if a field has been set.
func (o *CreateWorkItemRequest) HasAutoTests() bool {
	if o != nil && IsNil(o.AutoTests) {
		return true
	}

	return false
}

// SetAutoTests gets a reference to the given []AutoTestIdModel and assigns it to the AutoTests field.
func (o *CreateWorkItemRequest) SetAutoTests(v []AutoTestIdModel) {
	o.AutoTests = v
}

func (o CreateWorkItemRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWorkItemRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entityTypeName"] = o.EntityTypeName
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["state"] = o.State
	toSerialize["priority"] = o.Priority
	toSerialize["steps"] = o.Steps
	toSerialize["preconditionSteps"] = o.PreconditionSteps
	toSerialize["postconditionSteps"] = o.PostconditionSteps
	toSerialize["duration"] = o.Duration
	toSerialize["attributes"] = o.Attributes
	toSerialize["tags"] = o.Tags
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Iterations != nil {
		toSerialize["iterations"] = o.Iterations
	}
	toSerialize["links"] = o.Links
	toSerialize["name"] = o.Name
	toSerialize["projectId"] = o.ProjectId
	toSerialize["sectionId"] = o.SectionId
	if o.AutoTests != nil {
		toSerialize["autoTests"] = o.AutoTests
	}
	return toSerialize, nil
}

type NullableCreateWorkItemRequest struct {
	value *CreateWorkItemRequest
	isSet bool
}

func (v NullableCreateWorkItemRequest) Get() *CreateWorkItemRequest {
	return v.value
}

func (v *NullableCreateWorkItemRequest) Set(val *CreateWorkItemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkItemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkItemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkItemRequest(val *CreateWorkItemRequest) *NullableCreateWorkItemRequest {
	return &NullableCreateWorkItemRequest{value: val, isSet: true}
}

func (v NullableCreateWorkItemRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkItemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


