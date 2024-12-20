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

// checks if the WorkItemPutModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemPutModel{}

// WorkItemPutModel struct for WorkItemPutModel
type WorkItemPutModel struct {
	Attachments []AttachmentPutModel `json:"attachments"`
	Iterations []IterationPutModel `json:"iterations,omitempty"`
	AutoTests []AutoTestIdModel `json:"autoTests,omitempty"`
	Id string `json:"id"`
	SectionId string `json:"sectionId"`
	Description NullableString `json:"description,omitempty"`
	State WorkItemStates `json:"state"`
	Priority WorkItemPriorityModel `json:"priority"`
	Steps []StepPutModel `json:"steps"`
	PreconditionSteps []StepPutModel `json:"preconditionSteps"`
	PostconditionSteps []StepPutModel `json:"postconditionSteps"`
	Duration int32 `json:"duration"`
	Attributes map[string]interface{} `json:"attributes"`
	Tags []TagPutModel `json:"tags"`
	Links []LinkPutModel `json:"links"`
	Name string `json:"name"`
}

type _WorkItemPutModel WorkItemPutModel

// NewWorkItemPutModel instantiates a new WorkItemPutModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemPutModel(attachments []AttachmentPutModel, id string, sectionId string, state WorkItemStates, priority WorkItemPriorityModel, steps []StepPutModel, preconditionSteps []StepPutModel, postconditionSteps []StepPutModel, duration int32, attributes map[string]interface{}, tags []TagPutModel, links []LinkPutModel, name string) *WorkItemPutModel {
	this := WorkItemPutModel{}
	this.Attachments = attachments
	this.Id = id
	this.SectionId = sectionId
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
	return &this
}

// NewWorkItemPutModelWithDefaults instantiates a new WorkItemPutModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemPutModelWithDefaults() *WorkItemPutModel {
	this := WorkItemPutModel{}
	return &this
}

// GetAttachments returns the Attachments field value
func (o *WorkItemPutModel) GetAttachments() []AttachmentPutModel {
	if o == nil {
		var ret []AttachmentPutModel
		return ret
	}

	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value
// and a boolean to check if the value has been set.
func (o *WorkItemPutModel) GetAttachmentsOk() ([]AttachmentPutModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachments, true
}

// SetAttachments sets field value
func (o *WorkItemPutModel) SetAttachments(v []AttachmentPutModel) {
	o.Attachments = v
}

// GetIterations returns the Iterations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemPutModel) GetIterations() []IterationPutModel {
	if o == nil {
		var ret []IterationPutModel
		return ret
	}
	return o.Iterations
}

// GetIterationsOk returns a tuple with the Iterations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemPutModel) GetIterationsOk() ([]IterationPutModel, bool) {
	if o == nil || IsNil(o.Iterations) {
		return nil, false
	}
	return o.Iterations, true
}

// HasIterations returns a boolean if a field has been set.
func (o *WorkItemPutModel) HasIterations() bool {
	if o != nil && !IsNil(o.Iterations) {
		return true
	}

	return false
}

// SetIterations gets a reference to the given []IterationPutModel and assigns it to the Iterations field.
func (o *WorkItemPutModel) SetIterations(v []IterationPutModel) {
	o.Iterations = v
}

// GetAutoTests returns the AutoTests field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemPutModel) GetAutoTests() []AutoTestIdModel {
	if o == nil {
		var ret []AutoTestIdModel
		return ret
	}
	return o.AutoTests
}

// GetAutoTestsOk returns a tuple with the AutoTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemPutModel) GetAutoTestsOk() ([]AutoTestIdModel, bool) {
	if o == nil || IsNil(o.AutoTests) {
		return nil, false
	}
	return o.AutoTests, true
}

// HasAutoTests returns a boolean if a field has been set.
func (o *WorkItemPutModel) HasAutoTests() bool {
	if o != nil && !IsNil(o.AutoTests) {
		return true
	}

	return false
}

// SetAutoTests gets a reference to the given []AutoTestIdModel and assigns it to the AutoTests field.
func (o *WorkItemPutModel) SetAutoTests(v []AutoTestIdModel) {
	o.AutoTests = v
}

// GetId returns the Id field value
func (o *WorkItemPutModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkItemPutModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkItemPutModel) SetId(v string) {
	o.Id = v
}

// GetSectionId returns the SectionId field value
func (o *WorkItemPutModel) GetSectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SectionId
}

// GetSectionIdOk returns a tuple with the SectionId field value
// and a boolean to check if the value has been set.
func (o *WorkItemPutModel) GetSectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SectionId, true
}

// SetSectionId sets field value
func (o *WorkItemPutModel) SetSectionId(v string) {
	o.SectionId = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemPutModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemPutModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkItemPutModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *WorkItemPutModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *WorkItemPutModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *WorkItemPutModel) UnsetDescription() {
	o.Description.Unset()
}

// GetState returns the State field value
func (o *WorkItemPutModel) GetState() WorkItemStates {
	if o == nil {
		var ret WorkItemStates
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *WorkItemPutModel) GetStateOk() (*WorkItemStates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *WorkItemPutModel) SetState(v WorkItemStates) {
	o.State = v
}

// GetPriority returns the Priority field value
func (o *WorkItemPutModel) GetPriority() WorkItemPriorityModel {
	if o == nil {
		var ret WorkItemPriorityModel
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *WorkItemPutModel) GetPriorityOk() (*WorkItemPriorityModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *WorkItemPutModel) SetPriority(v WorkItemPriorityModel) {
	o.Priority = v
}

// GetSteps returns the Steps field value
func (o *WorkItemPutModel) GetSteps() []StepPutModel {
	if o == nil {
		var ret []StepPutModel
		return ret
	}

	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *WorkItemPutModel) GetStepsOk() ([]StepPutModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Steps, true
}

// SetSteps sets field value
func (o *WorkItemPutModel) SetSteps(v []StepPutModel) {
	o.Steps = v
}

// GetPreconditionSteps returns the PreconditionSteps field value
func (o *WorkItemPutModel) GetPreconditionSteps() []StepPutModel {
	if o == nil {
		var ret []StepPutModel
		return ret
	}

	return o.PreconditionSteps
}

// GetPreconditionStepsOk returns a tuple with the PreconditionSteps field value
// and a boolean to check if the value has been set.
func (o *WorkItemPutModel) GetPreconditionStepsOk() ([]StepPutModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreconditionSteps, true
}

// SetPreconditionSteps sets field value
func (o *WorkItemPutModel) SetPreconditionSteps(v []StepPutModel) {
	o.PreconditionSteps = v
}

// GetPostconditionSteps returns the PostconditionSteps field value
func (o *WorkItemPutModel) GetPostconditionSteps() []StepPutModel {
	if o == nil {
		var ret []StepPutModel
		return ret
	}

	return o.PostconditionSteps
}

// GetPostconditionStepsOk returns a tuple with the PostconditionSteps field value
// and a boolean to check if the value has been set.
func (o *WorkItemPutModel) GetPostconditionStepsOk() ([]StepPutModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostconditionSteps, true
}

// SetPostconditionSteps sets field value
func (o *WorkItemPutModel) SetPostconditionSteps(v []StepPutModel) {
	o.PostconditionSteps = v
}

// GetDuration returns the Duration field value
func (o *WorkItemPutModel) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *WorkItemPutModel) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *WorkItemPutModel) SetDuration(v int32) {
	o.Duration = v
}

// GetAttributes returns the Attributes field value
func (o *WorkItemPutModel) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *WorkItemPutModel) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *WorkItemPutModel) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetTags returns the Tags field value
func (o *WorkItemPutModel) GetTags() []TagPutModel {
	if o == nil {
		var ret []TagPutModel
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *WorkItemPutModel) GetTagsOk() ([]TagPutModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *WorkItemPutModel) SetTags(v []TagPutModel) {
	o.Tags = v
}

// GetLinks returns the Links field value
func (o *WorkItemPutModel) GetLinks() []LinkPutModel {
	if o == nil {
		var ret []LinkPutModel
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *WorkItemPutModel) GetLinksOk() ([]LinkPutModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *WorkItemPutModel) SetLinks(v []LinkPutModel) {
	o.Links = v
}

// GetName returns the Name field value
func (o *WorkItemPutModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkItemPutModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkItemPutModel) SetName(v string) {
	o.Name = v
}

func (o WorkItemPutModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemPutModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attachments"] = o.Attachments
	if o.Iterations != nil {
		toSerialize["iterations"] = o.Iterations
	}
	if o.AutoTests != nil {
		toSerialize["autoTests"] = o.AutoTests
	}
	toSerialize["id"] = o.Id
	toSerialize["sectionId"] = o.SectionId
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
	toSerialize["links"] = o.Links
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *WorkItemPutModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attachments",
		"id",
		"sectionId",
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

	varWorkItemPutModel := _WorkItemPutModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkItemPutModel)

	if err != nil {
		return err
	}

	*o = WorkItemPutModel(varWorkItemPutModel)

	return err
}

type NullableWorkItemPutModel struct {
	value *WorkItemPutModel
	isSet bool
}

func (v NullableWorkItemPutModel) Get() *WorkItemPutModel {
	return v.value
}

func (v *NullableWorkItemPutModel) Set(val *WorkItemPutModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemPutModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemPutModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemPutModel(val *WorkItemPutModel) *NullableWorkItemPutModel {
	return &NullableWorkItemPutModel{value: val, isSet: true}
}

func (v NullableWorkItemPutModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemPutModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


