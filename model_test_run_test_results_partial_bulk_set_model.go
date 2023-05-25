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

// checks if the TestRunTestResultsPartialBulkSetModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestRunTestResultsPartialBulkSetModel{}

// TestRunTestResultsPartialBulkSetModel struct for TestRunTestResultsPartialBulkSetModel
type TestRunTestResultsPartialBulkSetModel struct {
	Selector *TestRunTestResultsSelectModel `json:"selector,omitempty"`
	// Unique IDs of result reasons to be assigned to test results
	ResultReasonIds []string `json:"resultReasonIds,omitempty"`
	// Collection of links to be assigned to test results
	Links []LinkPostModel `json:"links,omitempty"`
	// Comment to be added to test results
	Comment NullableString `json:"comment,omitempty"`
	// Unique IDs of files to be attached to test results
	AttachmentIds []string `json:"attachmentIds,omitempty"`
}

// NewTestRunTestResultsPartialBulkSetModel instantiates a new TestRunTestResultsPartialBulkSetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestRunTestResultsPartialBulkSetModel() *TestRunTestResultsPartialBulkSetModel {
	this := TestRunTestResultsPartialBulkSetModel{}
	return &this
}

// NewTestRunTestResultsPartialBulkSetModelWithDefaults instantiates a new TestRunTestResultsPartialBulkSetModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestRunTestResultsPartialBulkSetModelWithDefaults() *TestRunTestResultsPartialBulkSetModel {
	this := TestRunTestResultsPartialBulkSetModel{}
	return &this
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *TestRunTestResultsPartialBulkSetModel) GetSelector() TestRunTestResultsSelectModel {
	if o == nil || IsNil(o.Selector) {
		var ret TestRunTestResultsSelectModel
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunTestResultsPartialBulkSetModel) GetSelectorOk() (*TestRunTestResultsSelectModel, bool) {
	if o == nil || IsNil(o.Selector) {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *TestRunTestResultsPartialBulkSetModel) HasSelector() bool {
	if o != nil && !IsNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given TestRunTestResultsSelectModel and assigns it to the Selector field.
func (o *TestRunTestResultsPartialBulkSetModel) SetSelector(v TestRunTestResultsSelectModel) {
	o.Selector = &v
}

// GetResultReasonIds returns the ResultReasonIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunTestResultsPartialBulkSetModel) GetResultReasonIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ResultReasonIds
}

// GetResultReasonIdsOk returns a tuple with the ResultReasonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunTestResultsPartialBulkSetModel) GetResultReasonIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ResultReasonIds) {
		return nil, false
	}
	return o.ResultReasonIds, true
}

// HasResultReasonIds returns a boolean if a field has been set.
func (o *TestRunTestResultsPartialBulkSetModel) HasResultReasonIds() bool {
	if o != nil && IsNil(o.ResultReasonIds) {
		return true
	}

	return false
}

// SetResultReasonIds gets a reference to the given []string and assigns it to the ResultReasonIds field.
func (o *TestRunTestResultsPartialBulkSetModel) SetResultReasonIds(v []string) {
	o.ResultReasonIds = v
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunTestResultsPartialBulkSetModel) GetLinks() []LinkPostModel {
	if o == nil {
		var ret []LinkPostModel
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunTestResultsPartialBulkSetModel) GetLinksOk() ([]LinkPostModel, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TestRunTestResultsPartialBulkSetModel) HasLinks() bool {
	if o != nil && IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkPostModel and assigns it to the Links field.
func (o *TestRunTestResultsPartialBulkSetModel) SetLinks(v []LinkPostModel) {
	o.Links = v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunTestResultsPartialBulkSetModel) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunTestResultsPartialBulkSetModel) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *TestRunTestResultsPartialBulkSetModel) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *TestRunTestResultsPartialBulkSetModel) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *TestRunTestResultsPartialBulkSetModel) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *TestRunTestResultsPartialBulkSetModel) UnsetComment() {
	o.Comment.Unset()
}

// GetAttachmentIds returns the AttachmentIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunTestResultsPartialBulkSetModel) GetAttachmentIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AttachmentIds
}

// GetAttachmentIdsOk returns a tuple with the AttachmentIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunTestResultsPartialBulkSetModel) GetAttachmentIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AttachmentIds) {
		return nil, false
	}
	return o.AttachmentIds, true
}

// HasAttachmentIds returns a boolean if a field has been set.
func (o *TestRunTestResultsPartialBulkSetModel) HasAttachmentIds() bool {
	if o != nil && IsNil(o.AttachmentIds) {
		return true
	}

	return false
}

// SetAttachmentIds gets a reference to the given []string and assigns it to the AttachmentIds field.
func (o *TestRunTestResultsPartialBulkSetModel) SetAttachmentIds(v []string) {
	o.AttachmentIds = v
}

func (o TestRunTestResultsPartialBulkSetModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestRunTestResultsPartialBulkSetModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	if o.ResultReasonIds != nil {
		toSerialize["resultReasonIds"] = o.ResultReasonIds
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.AttachmentIds != nil {
		toSerialize["attachmentIds"] = o.AttachmentIds
	}
	return toSerialize, nil
}

type NullableTestRunTestResultsPartialBulkSetModel struct {
	value *TestRunTestResultsPartialBulkSetModel
	isSet bool
}

func (v NullableTestRunTestResultsPartialBulkSetModel) Get() *TestRunTestResultsPartialBulkSetModel {
	return v.value
}

func (v *NullableTestRunTestResultsPartialBulkSetModel) Set(val *TestRunTestResultsPartialBulkSetModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestRunTestResultsPartialBulkSetModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestRunTestResultsPartialBulkSetModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestRunTestResultsPartialBulkSetModel(val *TestRunTestResultsPartialBulkSetModel) *NullableTestRunTestResultsPartialBulkSetModel {
	return &NullableTestRunTestResultsPartialBulkSetModel{value: val, isSet: true}
}

func (v NullableTestRunTestResultsPartialBulkSetModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestRunTestResultsPartialBulkSetModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


