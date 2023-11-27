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

// checks if the TestResultStepCommentPutModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestResultStepCommentPutModel{}

// TestResultStepCommentPutModel struct for TestResultStepCommentPutModel
type TestResultStepCommentPutModel struct {
	Id string `json:"id"`
	Text string `json:"text"`
	StepId string `json:"stepId"`
	ParentStepId NullableString `json:"parentStepId,omitempty"`
	Attachments []AttachmentPutModel `json:"attachments"`
}

// NewTestResultStepCommentPutModel instantiates a new TestResultStepCommentPutModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestResultStepCommentPutModel(id string, text string, stepId string, attachments []AttachmentPutModel) *TestResultStepCommentPutModel {
	this := TestResultStepCommentPutModel{}
	this.Id = id
	this.Text = text
	this.StepId = stepId
	this.Attachments = attachments
	return &this
}

// NewTestResultStepCommentPutModelWithDefaults instantiates a new TestResultStepCommentPutModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestResultStepCommentPutModelWithDefaults() *TestResultStepCommentPutModel {
	this := TestResultStepCommentPutModel{}
	return &this
}

// GetId returns the Id field value
func (o *TestResultStepCommentPutModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestResultStepCommentPutModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestResultStepCommentPutModel) SetId(v string) {
	o.Id = v
}

// GetText returns the Text field value
func (o *TestResultStepCommentPutModel) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *TestResultStepCommentPutModel) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *TestResultStepCommentPutModel) SetText(v string) {
	o.Text = v
}

// GetStepId returns the StepId field value
func (o *TestResultStepCommentPutModel) GetStepId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StepId
}

// GetStepIdOk returns a tuple with the StepId field value
// and a boolean to check if the value has been set.
func (o *TestResultStepCommentPutModel) GetStepIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StepId, true
}

// SetStepId sets field value
func (o *TestResultStepCommentPutModel) SetStepId(v string) {
	o.StepId = v
}

// GetParentStepId returns the ParentStepId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultStepCommentPutModel) GetParentStepId() string {
	if o == nil || IsNil(o.ParentStepId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentStepId.Get()
}

// GetParentStepIdOk returns a tuple with the ParentStepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultStepCommentPutModel) GetParentStepIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentStepId.Get(), o.ParentStepId.IsSet()
}

// HasParentStepId returns a boolean if a field has been set.
func (o *TestResultStepCommentPutModel) HasParentStepId() bool {
	if o != nil && o.ParentStepId.IsSet() {
		return true
	}

	return false
}

// SetParentStepId gets a reference to the given NullableString and assigns it to the ParentStepId field.
func (o *TestResultStepCommentPutModel) SetParentStepId(v string) {
	o.ParentStepId.Set(&v)
}
// SetParentStepIdNil sets the value for ParentStepId to be an explicit nil
func (o *TestResultStepCommentPutModel) SetParentStepIdNil() {
	o.ParentStepId.Set(nil)
}

// UnsetParentStepId ensures that no value is present for ParentStepId, not even an explicit nil
func (o *TestResultStepCommentPutModel) UnsetParentStepId() {
	o.ParentStepId.Unset()
}

// GetAttachments returns the Attachments field value
func (o *TestResultStepCommentPutModel) GetAttachments() []AttachmentPutModel {
	if o == nil {
		var ret []AttachmentPutModel
		return ret
	}

	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value
// and a boolean to check if the value has been set.
func (o *TestResultStepCommentPutModel) GetAttachmentsOk() ([]AttachmentPutModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attachments, true
}

// SetAttachments sets field value
func (o *TestResultStepCommentPutModel) SetAttachments(v []AttachmentPutModel) {
	o.Attachments = v
}

func (o TestResultStepCommentPutModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestResultStepCommentPutModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["text"] = o.Text
	toSerialize["stepId"] = o.StepId
	if o.ParentStepId.IsSet() {
		toSerialize["parentStepId"] = o.ParentStepId.Get()
	}
	toSerialize["attachments"] = o.Attachments
	return toSerialize, nil
}

type NullableTestResultStepCommentPutModel struct {
	value *TestResultStepCommentPutModel
	isSet bool
}

func (v NullableTestResultStepCommentPutModel) Get() *TestResultStepCommentPutModel {
	return v.value
}

func (v *NullableTestResultStepCommentPutModel) Set(val *TestResultStepCommentPutModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestResultStepCommentPutModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestResultStepCommentPutModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestResultStepCommentPutModel(val *TestResultStepCommentPutModel) *NullableTestResultStepCommentPutModel {
	return &NullableTestResultStepCommentPutModel{value: val, isSet: true}
}

func (v NullableTestResultStepCommentPutModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestResultStepCommentPutModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


