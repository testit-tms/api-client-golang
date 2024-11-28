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

// checks if the StepComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StepComment{}

// StepComment struct for StepComment
type StepComment struct {
	// Entity unique identifier
	Id string `json:"id"`
	Text NullableString `json:"text,omitempty"`
	StepId string `json:"stepId"`
	ParentStepId NullableString `json:"parentStepId,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	TestResultId string `json:"testResultId"`
	CreatedById string `json:"createdById"`
	ModifiedById NullableString `json:"modifiedById,omitempty"`
	CreatedDate time.Time `json:"createdDate"`
	ModifiedDate NullableTime `json:"modifiedDate,omitempty"`
}

type _StepComment StepComment

// NewStepComment instantiates a new StepComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStepComment(id string, stepId string, testResultId string, createdById string, createdDate time.Time) *StepComment {
	this := StepComment{}
	this.Id = id
	this.StepId = stepId
	this.TestResultId = testResultId
	this.CreatedById = createdById
	this.CreatedDate = createdDate
	return &this
}

// NewStepCommentWithDefaults instantiates a new StepComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStepCommentWithDefaults() *StepComment {
	this := StepComment{}
	return &this
}

// GetId returns the Id field value
func (o *StepComment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StepComment) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StepComment) SetId(v string) {
	o.Id = v
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StepComment) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StepComment) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *StepComment) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *StepComment) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *StepComment) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *StepComment) UnsetText() {
	o.Text.Unset()
}

// GetStepId returns the StepId field value
func (o *StepComment) GetStepId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StepId
}

// GetStepIdOk returns a tuple with the StepId field value
// and a boolean to check if the value has been set.
func (o *StepComment) GetStepIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StepId, true
}

// SetStepId sets field value
func (o *StepComment) SetStepId(v string) {
	o.StepId = v
}

// GetParentStepId returns the ParentStepId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StepComment) GetParentStepId() string {
	if o == nil || IsNil(o.ParentStepId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentStepId.Get()
}

// GetParentStepIdOk returns a tuple with the ParentStepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StepComment) GetParentStepIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentStepId.Get(), o.ParentStepId.IsSet()
}

// HasParentStepId returns a boolean if a field has been set.
func (o *StepComment) HasParentStepId() bool {
	if o != nil && o.ParentStepId.IsSet() {
		return true
	}

	return false
}

// SetParentStepId gets a reference to the given NullableString and assigns it to the ParentStepId field.
func (o *StepComment) SetParentStepId(v string) {
	o.ParentStepId.Set(&v)
}
// SetParentStepIdNil sets the value for ParentStepId to be an explicit nil
func (o *StepComment) SetParentStepIdNil() {
	o.ParentStepId.Set(nil)
}

// UnsetParentStepId ensures that no value is present for ParentStepId, not even an explicit nil
func (o *StepComment) UnsetParentStepId() {
	o.ParentStepId.Unset()
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StepComment) GetAttachments() []Attachment {
	if o == nil {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StepComment) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *StepComment) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *StepComment) SetAttachments(v []Attachment) {
	o.Attachments = v
}

// GetTestResultId returns the TestResultId field value
func (o *StepComment) GetTestResultId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestResultId
}

// GetTestResultIdOk returns a tuple with the TestResultId field value
// and a boolean to check if the value has been set.
func (o *StepComment) GetTestResultIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestResultId, true
}

// SetTestResultId sets field value
func (o *StepComment) SetTestResultId(v string) {
	o.TestResultId = v
}

// GetCreatedById returns the CreatedById field value
func (o *StepComment) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *StepComment) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *StepComment) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StepComment) GetModifiedById() string {
	if o == nil || IsNil(o.ModifiedById.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedById.Get()
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StepComment) GetModifiedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedById.Get(), o.ModifiedById.IsSet()
}

// HasModifiedById returns a boolean if a field has been set.
func (o *StepComment) HasModifiedById() bool {
	if o != nil && o.ModifiedById.IsSet() {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given NullableString and assigns it to the ModifiedById field.
func (o *StepComment) SetModifiedById(v string) {
	o.ModifiedById.Set(&v)
}
// SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil
func (o *StepComment) SetModifiedByIdNil() {
	o.ModifiedById.Set(nil)
}

// UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
func (o *StepComment) UnsetModifiedById() {
	o.ModifiedById.Unset()
}

// GetCreatedDate returns the CreatedDate field value
func (o *StepComment) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *StepComment) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *StepComment) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StepComment) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StepComment) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *StepComment) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTime and assigns it to the ModifiedDate field.
func (o *StepComment) SetModifiedDate(v time.Time) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *StepComment) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *StepComment) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

func (o StepComment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StepComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	toSerialize["stepId"] = o.StepId
	if o.ParentStepId.IsSet() {
		toSerialize["parentStepId"] = o.ParentStepId.Get()
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	toSerialize["testResultId"] = o.TestResultId
	toSerialize["createdById"] = o.CreatedById
	if o.ModifiedById.IsSet() {
		toSerialize["modifiedById"] = o.ModifiedById.Get()
	}
	toSerialize["createdDate"] = o.CreatedDate
	if o.ModifiedDate.IsSet() {
		toSerialize["modifiedDate"] = o.ModifiedDate.Get()
	}
	return toSerialize, nil
}

func (o *StepComment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"stepId",
		"testResultId",
		"createdById",
		"createdDate",
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

	varStepComment := _StepComment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStepComment)

	if err != nil {
		return err
	}

	*o = StepComment(varStepComment)

	return err
}

type NullableStepComment struct {
	value *StepComment
	isSet bool
}

func (v NullableStepComment) Get() *StepComment {
	return v.value
}

func (v *NullableStepComment) Set(val *StepComment) {
	v.value = val
	v.isSet = true
}

func (v NullableStepComment) IsSet() bool {
	return v.isSet
}

func (v *NullableStepComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStepComment(val *StepComment) *NullableStepComment {
	return &NullableStepComment{value: val, isSet: true}
}

func (v NullableStepComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStepComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

