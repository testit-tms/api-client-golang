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

// checks if the TestResultUpdateV2Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestResultUpdateV2Request{}

// TestResultUpdateV2Request struct for TestResultUpdateV2Request
type TestResultUpdateV2Request struct {
	FailureClassIds []string `json:"failureClassIds,omitempty"`
	// Deprecated
	Outcome NullableTestResultOutcome `json:"outcome,omitempty"`
	StatusCode NullableString `json:"statusCode,omitempty"`
	Comment NullableString `json:"comment,omitempty"`
	Links []Link `json:"links,omitempty"`
	StepResults []StepResultApiModel `json:"stepResults,omitempty"`
	Attachments []AttachmentUpdateRequest `json:"attachments,omitempty"`
	// Deprecated
	DurationInMs NullableInt64 `json:"durationInMs,omitempty"`
	Duration NullableInt64 `json:"duration,omitempty"`
	StepComments []TestResultStepCommentUpdateRequest `json:"stepComments,omitempty"`
	SetupResults []AutoTestStepResultUpdateRequest `json:"setupResults,omitempty"`
	TeardownResults []AutoTestStepResultUpdateRequest `json:"teardownResults,omitempty"`
	Message NullableString `json:"message,omitempty"`
	Trace NullableString `json:"trace,omitempty"`
}

// NewTestResultUpdateV2Request instantiates a new TestResultUpdateV2Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestResultUpdateV2Request() *TestResultUpdateV2Request {
	this := TestResultUpdateV2Request{}
	return &this
}

// NewTestResultUpdateV2RequestWithDefaults instantiates a new TestResultUpdateV2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestResultUpdateV2RequestWithDefaults() *TestResultUpdateV2Request {
	this := TestResultUpdateV2Request{}
	return &this
}

// GetFailureClassIds returns the FailureClassIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultUpdateV2Request) GetFailureClassIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.FailureClassIds
}

// GetFailureClassIdsOk returns a tuple with the FailureClassIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultUpdateV2Request) GetFailureClassIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.FailureClassIds) {
		return nil, false
	}
	return o.FailureClassIds, true
}

// HasFailureClassIds returns a boolean if a field has been set.
func (o *TestResultUpdateV2Request) HasFailureClassIds() bool {
	if o != nil && !IsNil(o.FailureClassIds) {
		return true
	}

	return false
}

// SetFailureClassIds gets a reference to the given []string and assigns it to the FailureClassIds field.
func (o *TestResultUpdateV2Request) SetFailureClassIds(v []string) {
	o.FailureClassIds = v
}

// GetOutcome returns the Outcome field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *TestResultUpdateV2Request) GetOutcome() TestResultOutcome {
	if o == nil || IsNil(o.Outcome.Get()) {
		var ret TestResultOutcome
		return ret
	}
	return *o.Outcome.Get()
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *TestResultUpdateV2Request) GetOutcomeOk() (*TestResultOutcome, bool) {
	if o == nil {
		return nil, false
	}
	return o.Outcome.Get(), o.Outcome.IsSet()
}

// HasOutcome returns a boolean if a field has been set.
func (o *TestResultUpdateV2Request) HasOutcome() bool {
	if o != nil && o.Outcome.IsSet() {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given NullableTestResultOutcome and assigns it to the Outcome field.
// Deprecated
func (o *TestResultUpdateV2Request) SetOutcome(v TestResultOutcome) {
	o.Outcome.Set(&v)
}
// SetOutcomeNil sets the value for Outcome to be an explicit nil
func (o *TestResultUpdateV2Request) SetOutcomeNil() {
	o.Outcome.Set(nil)
}

// UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
func (o *TestResultUpdateV2Request) UnsetOutcome() {
	o.Outcome.Unset()
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultUpdateV2Request) GetStatusCode() string {
	if o == nil || IsNil(o.StatusCode.Get()) {
		var ret string
		return ret
	}
	return *o.StatusCode.Get()
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultUpdateV2Request) GetStatusCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatusCode.Get(), o.StatusCode.IsSet()
}

// HasStatusCode returns a boolean if a field has been set.
func (o *TestResultUpdateV2Request) HasStatusCode() bool {
	if o != nil && o.StatusCode.IsSet() {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given NullableString and assigns it to the StatusCode field.
func (o *TestResultUpdateV2Request) SetStatusCode(v string) {
	o.StatusCode.Set(&v)
}
// SetStatusCodeNil sets the value for StatusCode to be an explicit nil
func (o *TestResultUpdateV2Request) SetStatusCodeNil() {
	o.StatusCode.Set(nil)
}

// UnsetStatusCode ensures that no value is present for StatusCode, not even an explicit nil
func (o *TestResultUpdateV2Request) UnsetStatusCode() {
	o.StatusCode.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultUpdateV2Request) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultUpdateV2Request) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *TestResultUpdateV2Request) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *TestResultUpdateV2Request) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *TestResultUpdateV2Request) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *TestResultUpdateV2Request) UnsetComment() {
	o.Comment.Unset()
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultUpdateV2Request) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultUpdateV2Request) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TestResultUpdateV2Request) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *TestResultUpdateV2Request) SetLinks(v []Link) {
	o.Links = v
}

// GetStepResults returns the StepResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultUpdateV2Request) GetStepResults() []StepResultApiModel {
	if o == nil {
		var ret []StepResultApiModel
		return ret
	}
	return o.StepResults
}

// GetStepResultsOk returns a tuple with the StepResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultUpdateV2Request) GetStepResultsOk() ([]StepResultApiModel, bool) {
	if o == nil || IsNil(o.StepResults) {
		return nil, false
	}
	return o.StepResults, true
}

// HasStepResults returns a boolean if a field has been set.
func (o *TestResultUpdateV2Request) HasStepResults() bool {
	if o != nil && !IsNil(o.StepResults) {
		return true
	}

	return false
}

// SetStepResults gets a reference to the given []StepResultApiModel and assigns it to the StepResults field.
func (o *TestResultUpdateV2Request) SetStepResults(v []StepResultApiModel) {
	o.StepResults = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultUpdateV2Request) GetAttachments() []AttachmentUpdateRequest {
	if o == nil {
		var ret []AttachmentUpdateRequest
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultUpdateV2Request) GetAttachmentsOk() ([]AttachmentUpdateRequest, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *TestResultUpdateV2Request) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []AttachmentUpdateRequest and assigns it to the Attachments field.
func (o *TestResultUpdateV2Request) SetAttachments(v []AttachmentUpdateRequest) {
	o.Attachments = v
}

// GetDurationInMs returns the DurationInMs field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *TestResultUpdateV2Request) GetDurationInMs() int64 {
	if o == nil || IsNil(o.DurationInMs.Get()) {
		var ret int64
		return ret
	}
	return *o.DurationInMs.Get()
}

// GetDurationInMsOk returns a tuple with the DurationInMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *TestResultUpdateV2Request) GetDurationInMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DurationInMs.Get(), o.DurationInMs.IsSet()
}

// HasDurationInMs returns a boolean if a field has been set.
func (o *TestResultUpdateV2Request) HasDurationInMs() bool {
	if o != nil && o.DurationInMs.IsSet() {
		return true
	}

	return false
}

// SetDurationInMs gets a reference to the given NullableInt64 and assigns it to the DurationInMs field.
// Deprecated
func (o *TestResultUpdateV2Request) SetDurationInMs(v int64) {
	o.DurationInMs.Set(&v)
}
// SetDurationInMsNil sets the value for DurationInMs to be an explicit nil
func (o *TestResultUpdateV2Request) SetDurationInMsNil() {
	o.DurationInMs.Set(nil)
}

// UnsetDurationInMs ensures that no value is present for DurationInMs, not even an explicit nil
func (o *TestResultUpdateV2Request) UnsetDurationInMs() {
	o.DurationInMs.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultUpdateV2Request) GetDuration() int64 {
	if o == nil || IsNil(o.Duration.Get()) {
		var ret int64
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultUpdateV2Request) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *TestResultUpdateV2Request) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableInt64 and assigns it to the Duration field.
func (o *TestResultUpdateV2Request) SetDuration(v int64) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *TestResultUpdateV2Request) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *TestResultUpdateV2Request) UnsetDuration() {
	o.Duration.Unset()
}

// GetStepComments returns the StepComments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultUpdateV2Request) GetStepComments() []TestResultStepCommentUpdateRequest {
	if o == nil {
		var ret []TestResultStepCommentUpdateRequest
		return ret
	}
	return o.StepComments
}

// GetStepCommentsOk returns a tuple with the StepComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultUpdateV2Request) GetStepCommentsOk() ([]TestResultStepCommentUpdateRequest, bool) {
	if o == nil || IsNil(o.StepComments) {
		return nil, false
	}
	return o.StepComments, true
}

// HasStepComments returns a boolean if a field has been set.
func (o *TestResultUpdateV2Request) HasStepComments() bool {
	if o != nil && !IsNil(o.StepComments) {
		return true
	}

	return false
}

// SetStepComments gets a reference to the given []TestResultStepCommentUpdateRequest and assigns it to the StepComments field.
func (o *TestResultUpdateV2Request) SetStepComments(v []TestResultStepCommentUpdateRequest) {
	o.StepComments = v
}

// GetSetupResults returns the SetupResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultUpdateV2Request) GetSetupResults() []AutoTestStepResultUpdateRequest {
	if o == nil {
		var ret []AutoTestStepResultUpdateRequest
		return ret
	}
	return o.SetupResults
}

// GetSetupResultsOk returns a tuple with the SetupResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultUpdateV2Request) GetSetupResultsOk() ([]AutoTestStepResultUpdateRequest, bool) {
	if o == nil || IsNil(o.SetupResults) {
		return nil, false
	}
	return o.SetupResults, true
}

// HasSetupResults returns a boolean if a field has been set.
func (o *TestResultUpdateV2Request) HasSetupResults() bool {
	if o != nil && !IsNil(o.SetupResults) {
		return true
	}

	return false
}

// SetSetupResults gets a reference to the given []AutoTestStepResultUpdateRequest and assigns it to the SetupResults field.
func (o *TestResultUpdateV2Request) SetSetupResults(v []AutoTestStepResultUpdateRequest) {
	o.SetupResults = v
}

// GetTeardownResults returns the TeardownResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultUpdateV2Request) GetTeardownResults() []AutoTestStepResultUpdateRequest {
	if o == nil {
		var ret []AutoTestStepResultUpdateRequest
		return ret
	}
	return o.TeardownResults
}

// GetTeardownResultsOk returns a tuple with the TeardownResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultUpdateV2Request) GetTeardownResultsOk() ([]AutoTestStepResultUpdateRequest, bool) {
	if o == nil || IsNil(o.TeardownResults) {
		return nil, false
	}
	return o.TeardownResults, true
}

// HasTeardownResults returns a boolean if a field has been set.
func (o *TestResultUpdateV2Request) HasTeardownResults() bool {
	if o != nil && !IsNil(o.TeardownResults) {
		return true
	}

	return false
}

// SetTeardownResults gets a reference to the given []AutoTestStepResultUpdateRequest and assigns it to the TeardownResults field.
func (o *TestResultUpdateV2Request) SetTeardownResults(v []AutoTestStepResultUpdateRequest) {
	o.TeardownResults = v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultUpdateV2Request) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultUpdateV2Request) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *TestResultUpdateV2Request) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *TestResultUpdateV2Request) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *TestResultUpdateV2Request) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *TestResultUpdateV2Request) UnsetMessage() {
	o.Message.Unset()
}

// GetTrace returns the Trace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestResultUpdateV2Request) GetTrace() string {
	if o == nil || IsNil(o.Trace.Get()) {
		var ret string
		return ret
	}
	return *o.Trace.Get()
}

// GetTraceOk returns a tuple with the Trace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestResultUpdateV2Request) GetTraceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Trace.Get(), o.Trace.IsSet()
}

// HasTrace returns a boolean if a field has been set.
func (o *TestResultUpdateV2Request) HasTrace() bool {
	if o != nil && o.Trace.IsSet() {
		return true
	}

	return false
}

// SetTrace gets a reference to the given NullableString and assigns it to the Trace field.
func (o *TestResultUpdateV2Request) SetTrace(v string) {
	o.Trace.Set(&v)
}
// SetTraceNil sets the value for Trace to be an explicit nil
func (o *TestResultUpdateV2Request) SetTraceNil() {
	o.Trace.Set(nil)
}

// UnsetTrace ensures that no value is present for Trace, not even an explicit nil
func (o *TestResultUpdateV2Request) UnsetTrace() {
	o.Trace.Unset()
}

func (o TestResultUpdateV2Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestResultUpdateV2Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FailureClassIds != nil {
		toSerialize["failureClassIds"] = o.FailureClassIds
	}
	if o.Outcome.IsSet() {
		toSerialize["outcome"] = o.Outcome.Get()
	}
	if o.StatusCode.IsSet() {
		toSerialize["statusCode"] = o.StatusCode.Get()
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.StepResults != nil {
		toSerialize["stepResults"] = o.StepResults
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.DurationInMs.IsSet() {
		toSerialize["durationInMs"] = o.DurationInMs.Get()
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if o.StepComments != nil {
		toSerialize["stepComments"] = o.StepComments
	}
	if o.SetupResults != nil {
		toSerialize["setupResults"] = o.SetupResults
	}
	if o.TeardownResults != nil {
		toSerialize["teardownResults"] = o.TeardownResults
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Trace.IsSet() {
		toSerialize["trace"] = o.Trace.Get()
	}
	return toSerialize, nil
}

type NullableTestResultUpdateV2Request struct {
	value *TestResultUpdateV2Request
	isSet bool
}

func (v NullableTestResultUpdateV2Request) Get() *TestResultUpdateV2Request {
	return v.value
}

func (v *NullableTestResultUpdateV2Request) Set(val *TestResultUpdateV2Request) {
	v.value = val
	v.isSet = true
}

func (v NullableTestResultUpdateV2Request) IsSet() bool {
	return v.isSet
}

func (v *NullableTestResultUpdateV2Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestResultUpdateV2Request(val *TestResultUpdateV2Request) *NullableTestResultUpdateV2Request {
	return &NullableTestResultUpdateV2Request{value: val, isSet: true}
}

func (v NullableTestResultUpdateV2Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestResultUpdateV2Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


