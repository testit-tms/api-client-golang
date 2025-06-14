/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateWorkItemApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWorkItemApiModel{}

// CreateWorkItemApiModel struct for CreateWorkItemApiModel
type CreateWorkItemApiModel struct {
	EntityTypeName WorkItemEntityTypes `json:"entityTypeName"`
	// Workitem description
	Description NullableString `json:"description,omitempty"`
	State WorkItemStates `json:"state"`
	Priority WorkItemPriorityModel `json:"priority"`
	// Collection of workitem steps
	Steps []CreateStepApiModel `json:"steps"`
	// Collection of workitem precondition steps
	PreconditionSteps []CreateStepApiModel `json:"preconditionSteps"`
	// Collection of workitem postcondition steps
	PostconditionSteps []CreateStepApiModel `json:"postconditionSteps"`
	// WorkItem duration in milliseconds, must be 0 for shared steps and greater than 0 for the other types of work items
	Duration int32 `json:"duration"`
	// Key value pair of custom workitem attributes
	Attributes map[string]interface{} `json:"attributes"`
	// Collection of workitem tags
	Tags []TagModel `json:"tags"`
	// Collection of workitem attachments
	Attachments []AssignAttachmentApiModel `json:"attachments,omitempty"`
	// Collection of parameter sets
	Iterations []AssignIterationApiModel `json:"iterations,omitempty"`
	// Collection of workitem links
	Links []CreateLinkApiModel `json:"links"`
	// Workitem name
	Name string `json:"name"`
	// Project unique identifier - used to link workitem with project
	ProjectId string `json:"projectId"`
	// Internal identifier of section where workitem is located
	SectionId string `json:"sectionId"`
	// Collection of autotest internal ids
	AutoTests []AutoTestIdModel `json:"autoTests,omitempty"`
}

type _CreateWorkItemApiModel CreateWorkItemApiModel

// NewCreateWorkItemApiModel instantiates a new CreateWorkItemApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWorkItemApiModel(entityTypeName WorkItemEntityTypes, state WorkItemStates, priority WorkItemPriorityModel, steps []CreateStepApiModel, preconditionSteps []CreateStepApiModel, postconditionSteps []CreateStepApiModel, duration int32, attributes map[string]interface{}, tags []TagModel, links []CreateLinkApiModel, name string, projectId string, sectionId string) *CreateWorkItemApiModel {
	this := CreateWorkItemApiModel{}
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

// NewCreateWorkItemApiModelWithDefaults instantiates a new CreateWorkItemApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWorkItemApiModelWithDefaults() *CreateWorkItemApiModel {
	this := CreateWorkItemApiModel{}
	return &this
}

// GetEntityTypeName returns the EntityTypeName field value
func (o *CreateWorkItemApiModel) GetEntityTypeName() WorkItemEntityTypes {
	if o == nil {
		var ret WorkItemEntityTypes
		return ret
	}

	return o.EntityTypeName
}

// GetEntityTypeNameOk returns a tuple with the EntityTypeName field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemApiModel) GetEntityTypeNameOk() (*WorkItemEntityTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityTypeName, true
}

// SetEntityTypeName sets field value
func (o *CreateWorkItemApiModel) SetEntityTypeName(v WorkItemEntityTypes) {
	o.EntityTypeName = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateWorkItemApiModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateWorkItemApiModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateWorkItemApiModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateWorkItemApiModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateWorkItemApiModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateWorkItemApiModel) UnsetDescription() {
	o.Description.Unset()
}

// GetState returns the State field value
func (o *CreateWorkItemApiModel) GetState() WorkItemStates {
	if o == nil {
		var ret WorkItemStates
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemApiModel) GetStateOk() (*WorkItemStates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CreateWorkItemApiModel) SetState(v WorkItemStates) {
	o.State = v
}

// GetPriority returns the Priority field value
func (o *CreateWorkItemApiModel) GetPriority() WorkItemPriorityModel {
	if o == nil {
		var ret WorkItemPriorityModel
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemApiModel) GetPriorityOk() (*WorkItemPriorityModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *CreateWorkItemApiModel) SetPriority(v WorkItemPriorityModel) {
	o.Priority = v
}

// GetSteps returns the Steps field value
func (o *CreateWorkItemApiModel) GetSteps() []CreateStepApiModel {
	if o == nil {
		var ret []CreateStepApiModel
		return ret
	}

	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemApiModel) GetStepsOk() ([]CreateStepApiModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Steps, true
}

// SetSteps sets field value
func (o *CreateWorkItemApiModel) SetSteps(v []CreateStepApiModel) {
	o.Steps = v
}

// GetPreconditionSteps returns the PreconditionSteps field value
func (o *CreateWorkItemApiModel) GetPreconditionSteps() []CreateStepApiModel {
	if o == nil {
		var ret []CreateStepApiModel
		return ret
	}

	return o.PreconditionSteps
}

// GetPreconditionStepsOk returns a tuple with the PreconditionSteps field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemApiModel) GetPreconditionStepsOk() ([]CreateStepApiModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreconditionSteps, true
}

// SetPreconditionSteps sets field value
func (o *CreateWorkItemApiModel) SetPreconditionSteps(v []CreateStepApiModel) {
	o.PreconditionSteps = v
}

// GetPostconditionSteps returns the PostconditionSteps field value
func (o *CreateWorkItemApiModel) GetPostconditionSteps() []CreateStepApiModel {
	if o == nil {
		var ret []CreateStepApiModel
		return ret
	}

	return o.PostconditionSteps
}

// GetPostconditionStepsOk returns a tuple with the PostconditionSteps field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemApiModel) GetPostconditionStepsOk() ([]CreateStepApiModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostconditionSteps, true
}

// SetPostconditionSteps sets field value
func (o *CreateWorkItemApiModel) SetPostconditionSteps(v []CreateStepApiModel) {
	o.PostconditionSteps = v
}

// GetDuration returns the Duration field value
func (o *CreateWorkItemApiModel) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemApiModel) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *CreateWorkItemApiModel) SetDuration(v int32) {
	o.Duration = v
}

// GetAttributes returns the Attributes field value
func (o *CreateWorkItemApiModel) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemApiModel) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *CreateWorkItemApiModel) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetTags returns the Tags field value
func (o *CreateWorkItemApiModel) GetTags() []TagModel {
	if o == nil {
		var ret []TagModel
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemApiModel) GetTagsOk() ([]TagModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *CreateWorkItemApiModel) SetTags(v []TagModel) {
	o.Tags = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateWorkItemApiModel) GetAttachments() []AssignAttachmentApiModel {
	if o == nil {
		var ret []AssignAttachmentApiModel
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateWorkItemApiModel) GetAttachmentsOk() ([]AssignAttachmentApiModel, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *CreateWorkItemApiModel) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []AssignAttachmentApiModel and assigns it to the Attachments field.
func (o *CreateWorkItemApiModel) SetAttachments(v []AssignAttachmentApiModel) {
	o.Attachments = v
}

// GetIterations returns the Iterations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateWorkItemApiModel) GetIterations() []AssignIterationApiModel {
	if o == nil {
		var ret []AssignIterationApiModel
		return ret
	}
	return o.Iterations
}

// GetIterationsOk returns a tuple with the Iterations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateWorkItemApiModel) GetIterationsOk() ([]AssignIterationApiModel, bool) {
	if o == nil || IsNil(o.Iterations) {
		return nil, false
	}
	return o.Iterations, true
}

// HasIterations returns a boolean if a field has been set.
func (o *CreateWorkItemApiModel) HasIterations() bool {
	if o != nil && !IsNil(o.Iterations) {
		return true
	}

	return false
}

// SetIterations gets a reference to the given []AssignIterationApiModel and assigns it to the Iterations field.
func (o *CreateWorkItemApiModel) SetIterations(v []AssignIterationApiModel) {
	o.Iterations = v
}

// GetLinks returns the Links field value
func (o *CreateWorkItemApiModel) GetLinks() []CreateLinkApiModel {
	if o == nil {
		var ret []CreateLinkApiModel
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemApiModel) GetLinksOk() ([]CreateLinkApiModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *CreateWorkItemApiModel) SetLinks(v []CreateLinkApiModel) {
	o.Links = v
}

// GetName returns the Name field value
func (o *CreateWorkItemApiModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemApiModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateWorkItemApiModel) SetName(v string) {
	o.Name = v
}

// GetProjectId returns the ProjectId field value
func (o *CreateWorkItemApiModel) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemApiModel) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *CreateWorkItemApiModel) SetProjectId(v string) {
	o.ProjectId = v
}

// GetSectionId returns the SectionId field value
func (o *CreateWorkItemApiModel) GetSectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SectionId
}

// GetSectionIdOk returns a tuple with the SectionId field value
// and a boolean to check if the value has been set.
func (o *CreateWorkItemApiModel) GetSectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SectionId, true
}

// SetSectionId sets field value
func (o *CreateWorkItemApiModel) SetSectionId(v string) {
	o.SectionId = v
}

// GetAutoTests returns the AutoTests field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateWorkItemApiModel) GetAutoTests() []AutoTestIdModel {
	if o == nil {
		var ret []AutoTestIdModel
		return ret
	}
	return o.AutoTests
}

// GetAutoTestsOk returns a tuple with the AutoTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateWorkItemApiModel) GetAutoTestsOk() ([]AutoTestIdModel, bool) {
	if o == nil || IsNil(o.AutoTests) {
		return nil, false
	}
	return o.AutoTests, true
}

// HasAutoTests returns a boolean if a field has been set.
func (o *CreateWorkItemApiModel) HasAutoTests() bool {
	if o != nil && !IsNil(o.AutoTests) {
		return true
	}

	return false
}

// SetAutoTests gets a reference to the given []AutoTestIdModel and assigns it to the AutoTests field.
func (o *CreateWorkItemApiModel) SetAutoTests(v []AutoTestIdModel) {
	o.AutoTests = v
}

func (o CreateWorkItemApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWorkItemApiModel) ToMap() (map[string]interface{}, error) {
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

func (o *CreateWorkItemApiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entityTypeName",
		"state",
		"priority",
		"steps",
		"preconditionSteps",
		"postconditionSteps",
		"duration",
		"attributes",
		"tags",
		"links",
		"name",
		"projectId",
		"sectionId",
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

	varCreateWorkItemApiModel := _CreateWorkItemApiModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateWorkItemApiModel)

	if err != nil {
		return err
	}

	*o = CreateWorkItemApiModel(varCreateWorkItemApiModel)

	return err
}

type NullableCreateWorkItemApiModel struct {
	value *CreateWorkItemApiModel
	isSet bool
}

func (v NullableCreateWorkItemApiModel) Get() *CreateWorkItemApiModel {
	return v.value
}

func (v *NullableCreateWorkItemApiModel) Set(val *CreateWorkItemApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkItemApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkItemApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkItemApiModel(val *CreateWorkItemApiModel) *NullableCreateWorkItemApiModel {
	return &NullableCreateWorkItemApiModel{value: val, isSet: true}
}

func (v NullableCreateWorkItemApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkItemApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


