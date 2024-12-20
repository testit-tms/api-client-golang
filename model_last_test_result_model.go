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

// checks if the LastTestResultModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LastTestResultModel{}

// LastTestResultModel struct for LastTestResultModel
type LastTestResultModel struct {
	Id string `json:"id"`
	TestRunId string `json:"testRunId"`
	AutoTestId NullableString `json:"autoTestId,omitempty"`
	Comment NullableString `json:"comment,omitempty"`
	Links []LinkModel `json:"links,omitempty"`
	WorkItemVersionId string `json:"workItemVersionId"`
	Attachments []AttachmentModel `json:"attachments,omitempty"`
}

type _LastTestResultModel LastTestResultModel

// NewLastTestResultModel instantiates a new LastTestResultModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLastTestResultModel(id string, testRunId string, workItemVersionId string) *LastTestResultModel {
	this := LastTestResultModel{}
	this.Id = id
	this.TestRunId = testRunId
	this.WorkItemVersionId = workItemVersionId
	return &this
}

// NewLastTestResultModelWithDefaults instantiates a new LastTestResultModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLastTestResultModelWithDefaults() *LastTestResultModel {
	this := LastTestResultModel{}
	return &this
}

// GetId returns the Id field value
func (o *LastTestResultModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LastTestResultModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LastTestResultModel) SetId(v string) {
	o.Id = v
}

// GetTestRunId returns the TestRunId field value
func (o *LastTestResultModel) GetTestRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestRunId
}

// GetTestRunIdOk returns a tuple with the TestRunId field value
// and a boolean to check if the value has been set.
func (o *LastTestResultModel) GetTestRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestRunId, true
}

// SetTestRunId sets field value
func (o *LastTestResultModel) SetTestRunId(v string) {
	o.TestRunId = v
}

// GetAutoTestId returns the AutoTestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LastTestResultModel) GetAutoTestId() string {
	if o == nil || IsNil(o.AutoTestId.Get()) {
		var ret string
		return ret
	}
	return *o.AutoTestId.Get()
}

// GetAutoTestIdOk returns a tuple with the AutoTestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LastTestResultModel) GetAutoTestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoTestId.Get(), o.AutoTestId.IsSet()
}

// HasAutoTestId returns a boolean if a field has been set.
func (o *LastTestResultModel) HasAutoTestId() bool {
	if o != nil && o.AutoTestId.IsSet() {
		return true
	}

	return false
}

// SetAutoTestId gets a reference to the given NullableString and assigns it to the AutoTestId field.
func (o *LastTestResultModel) SetAutoTestId(v string) {
	o.AutoTestId.Set(&v)
}
// SetAutoTestIdNil sets the value for AutoTestId to be an explicit nil
func (o *LastTestResultModel) SetAutoTestIdNil() {
	o.AutoTestId.Set(nil)
}

// UnsetAutoTestId ensures that no value is present for AutoTestId, not even an explicit nil
func (o *LastTestResultModel) UnsetAutoTestId() {
	o.AutoTestId.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LastTestResultModel) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LastTestResultModel) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *LastTestResultModel) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *LastTestResultModel) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *LastTestResultModel) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *LastTestResultModel) UnsetComment() {
	o.Comment.Unset()
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LastTestResultModel) GetLinks() []LinkModel {
	if o == nil {
		var ret []LinkModel
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LastTestResultModel) GetLinksOk() ([]LinkModel, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *LastTestResultModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkModel and assigns it to the Links field.
func (o *LastTestResultModel) SetLinks(v []LinkModel) {
	o.Links = v
}

// GetWorkItemVersionId returns the WorkItemVersionId field value
func (o *LastTestResultModel) GetWorkItemVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkItemVersionId
}

// GetWorkItemVersionIdOk returns a tuple with the WorkItemVersionId field value
// and a boolean to check if the value has been set.
func (o *LastTestResultModel) GetWorkItemVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkItemVersionId, true
}

// SetWorkItemVersionId sets field value
func (o *LastTestResultModel) SetWorkItemVersionId(v string) {
	o.WorkItemVersionId = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LastTestResultModel) GetAttachments() []AttachmentModel {
	if o == nil {
		var ret []AttachmentModel
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LastTestResultModel) GetAttachmentsOk() ([]AttachmentModel, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *LastTestResultModel) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []AttachmentModel and assigns it to the Attachments field.
func (o *LastTestResultModel) SetAttachments(v []AttachmentModel) {
	o.Attachments = v
}

func (o LastTestResultModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LastTestResultModel) ToMap() (map[string]interface{}, error) {
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

func (o *LastTestResultModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"testRunId",
		"workItemVersionId",
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

	varLastTestResultModel := _LastTestResultModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLastTestResultModel)

	if err != nil {
		return err
	}

	*o = LastTestResultModel(varLastTestResultModel)

	return err
}

type NullableLastTestResultModel struct {
	value *LastTestResultModel
	isSet bool
}

func (v NullableLastTestResultModel) Get() *LastTestResultModel {
	return v.value
}

func (v *NullableLastTestResultModel) Set(val *LastTestResultModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLastTestResultModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLastTestResultModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLastTestResultModel(val *LastTestResultModel) *NullableLastTestResultModel {
	return &NullableLastTestResultModel{value: val, isSet: true}
}

func (v NullableLastTestResultModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLastTestResultModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


