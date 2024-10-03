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

// checks if the TestPointShortGetModelLastTestResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPointShortGetModelLastTestResult{}

// TestPointShortGetModelLastTestResult Model of the test point last test result
type TestPointShortGetModelLastTestResult struct {
	Id string `json:"id"`
	TestRunId string `json:"testRunId"`
	AutoTestId NullableString `json:"autoTestId,omitempty"`
	Comment NullableString `json:"comment,omitempty"`
	Links []LinkModel `json:"links,omitempty"`
	WorkItemVersionId string `json:"workItemVersionId"`
	Attachments []AttachmentModel `json:"attachments,omitempty"`
}

// NewTestPointShortGetModelLastTestResult instantiates a new TestPointShortGetModelLastTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPointShortGetModelLastTestResult(id string, testRunId string, workItemVersionId string) *TestPointShortGetModelLastTestResult {
	this := TestPointShortGetModelLastTestResult{}
	this.Id = id
	this.TestRunId = testRunId
	this.WorkItemVersionId = workItemVersionId
	return &this
}

// NewTestPointShortGetModelLastTestResultWithDefaults instantiates a new TestPointShortGetModelLastTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPointShortGetModelLastTestResultWithDefaults() *TestPointShortGetModelLastTestResult {
	this := TestPointShortGetModelLastTestResult{}
	return &this
}

// GetId returns the Id field value
func (o *TestPointShortGetModelLastTestResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestPointShortGetModelLastTestResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestPointShortGetModelLastTestResult) SetId(v string) {
	o.Id = v
}

// GetTestRunId returns the TestRunId field value
func (o *TestPointShortGetModelLastTestResult) GetTestRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestRunId
}

// GetTestRunIdOk returns a tuple with the TestRunId field value
// and a boolean to check if the value has been set.
func (o *TestPointShortGetModelLastTestResult) GetTestRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestRunId, true
}

// SetTestRunId sets field value
func (o *TestPointShortGetModelLastTestResult) SetTestRunId(v string) {
	o.TestRunId = v
}

// GetAutoTestId returns the AutoTestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortGetModelLastTestResult) GetAutoTestId() string {
	if o == nil || IsNil(o.AutoTestId.Get()) {
		var ret string
		return ret
	}
	return *o.AutoTestId.Get()
}

// GetAutoTestIdOk returns a tuple with the AutoTestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortGetModelLastTestResult) GetAutoTestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoTestId.Get(), o.AutoTestId.IsSet()
}

// HasAutoTestId returns a boolean if a field has been set.
func (o *TestPointShortGetModelLastTestResult) HasAutoTestId() bool {
	if o != nil && o.AutoTestId.IsSet() {
		return true
	}

	return false
}

// SetAutoTestId gets a reference to the given NullableString and assigns it to the AutoTestId field.
func (o *TestPointShortGetModelLastTestResult) SetAutoTestId(v string) {
	o.AutoTestId.Set(&v)
}
// SetAutoTestIdNil sets the value for AutoTestId to be an explicit nil
func (o *TestPointShortGetModelLastTestResult) SetAutoTestIdNil() {
	o.AutoTestId.Set(nil)
}

// UnsetAutoTestId ensures that no value is present for AutoTestId, not even an explicit nil
func (o *TestPointShortGetModelLastTestResult) UnsetAutoTestId() {
	o.AutoTestId.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortGetModelLastTestResult) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortGetModelLastTestResult) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *TestPointShortGetModelLastTestResult) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *TestPointShortGetModelLastTestResult) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *TestPointShortGetModelLastTestResult) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *TestPointShortGetModelLastTestResult) UnsetComment() {
	o.Comment.Unset()
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortGetModelLastTestResult) GetLinks() []LinkModel {
	if o == nil {
		var ret []LinkModel
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortGetModelLastTestResult) GetLinksOk() ([]LinkModel, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TestPointShortGetModelLastTestResult) HasLinks() bool {
	if o != nil && IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkModel and assigns it to the Links field.
func (o *TestPointShortGetModelLastTestResult) SetLinks(v []LinkModel) {
	o.Links = v
}

// GetWorkItemVersionId returns the WorkItemVersionId field value
func (o *TestPointShortGetModelLastTestResult) GetWorkItemVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkItemVersionId
}

// GetWorkItemVersionIdOk returns a tuple with the WorkItemVersionId field value
// and a boolean to check if the value has been set.
func (o *TestPointShortGetModelLastTestResult) GetWorkItemVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkItemVersionId, true
}

// SetWorkItemVersionId sets field value
func (o *TestPointShortGetModelLastTestResult) SetWorkItemVersionId(v string) {
	o.WorkItemVersionId = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointShortGetModelLastTestResult) GetAttachments() []AttachmentModel {
	if o == nil {
		var ret []AttachmentModel
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointShortGetModelLastTestResult) GetAttachmentsOk() ([]AttachmentModel, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *TestPointShortGetModelLastTestResult) HasAttachments() bool {
	if o != nil && IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []AttachmentModel and assigns it to the Attachments field.
func (o *TestPointShortGetModelLastTestResult) SetAttachments(v []AttachmentModel) {
	o.Attachments = v
}

func (o TestPointShortGetModelLastTestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPointShortGetModelLastTestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["testRunId"] = o.TestRunId
	if o.AutoTestId.IsSet() {
		toSerialize["autoTestId"] = o.AutoTestId.Get()
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	toSerialize["workItemVersionId"] = o.WorkItemVersionId
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

type NullableTestPointShortGetModelLastTestResult struct {
	value *TestPointShortGetModelLastTestResult
	isSet bool
}

func (v NullableTestPointShortGetModelLastTestResult) Get() *TestPointShortGetModelLastTestResult {
	return v.value
}

func (v *NullableTestPointShortGetModelLastTestResult) Set(val *TestPointShortGetModelLastTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPointShortGetModelLastTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPointShortGetModelLastTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPointShortGetModelLastTestResult(val *TestPointShortGetModelLastTestResult) *NullableTestPointShortGetModelLastTestResult {
	return &NullableTestPointShortGetModelLastTestResult{value: val, isSet: true}
}

func (v NullableTestPointShortGetModelLastTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPointShortGetModelLastTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


