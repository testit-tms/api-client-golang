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
)

// checks if the AutoTestStepResultUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoTestStepResultUpdateRequest{}

// AutoTestStepResultUpdateRequest struct for AutoTestStepResultUpdateRequest
type AutoTestStepResultUpdateRequest struct {
	// The name of the step.
	Title NullableString `json:"title,omitempty"`
	// Description of the step result.
	Description NullableString `json:"description,omitempty"`
	// Extended description of the step result.
	Info NullableString `json:"info,omitempty"`
	// Step start date.
	StartedOn NullableTime `json:"startedOn,omitempty"`
	// Step end date.
	CompletedOn NullableTime `json:"completedOn,omitempty"`
	// Expected or actual duration of the test run execution in milliseconds.
	Duration NullableInt64 `json:"duration,omitempty"`
	// Specifies the result of the autotest execution.
	Outcome NullableAvailableTestResultOutcome `json:"outcome,omitempty"`
	// Nested step results. The maximum nesting level is 15.
	StepResults []AttachmentPutModelAutoTestStepResultsModel `json:"stepResults,omitempty"`
	// /// <summary>  Specifies an attachment GUID. Multiple values can be sent.  </summary>
	Attachments []AttachmentUpdateRequest `json:"attachments,omitempty"`
	// \"<b>parameter</b>\": \"<b>value</b>\" pair with arbitrary custom parameters. Multiple parameters can be sent.
	Parameters map[string]string `json:"parameters,omitempty"`
}

// NewAutoTestStepResultUpdateRequest instantiates a new AutoTestStepResultUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoTestStepResultUpdateRequest() *AutoTestStepResultUpdateRequest {
	this := AutoTestStepResultUpdateRequest{}
	return &this
}

// NewAutoTestStepResultUpdateRequestWithDefaults instantiates a new AutoTestStepResultUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoTestStepResultUpdateRequestWithDefaults() *AutoTestStepResultUpdateRequest {
	this := AutoTestStepResultUpdateRequest{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestStepResultUpdateRequest) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestStepResultUpdateRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *AutoTestStepResultUpdateRequest) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *AutoTestStepResultUpdateRequest) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *AutoTestStepResultUpdateRequest) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *AutoTestStepResultUpdateRequest) UnsetTitle() {
	o.Title.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestStepResultUpdateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestStepResultUpdateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AutoTestStepResultUpdateRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AutoTestStepResultUpdateRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AutoTestStepResultUpdateRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AutoTestStepResultUpdateRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetInfo returns the Info field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestStepResultUpdateRequest) GetInfo() string {
	if o == nil || IsNil(o.Info.Get()) {
		var ret string
		return ret
	}
	return *o.Info.Get()
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestStepResultUpdateRequest) GetInfoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Info.Get(), o.Info.IsSet()
}

// HasInfo returns a boolean if a field has been set.
func (o *AutoTestStepResultUpdateRequest) HasInfo() bool {
	if o != nil && o.Info.IsSet() {
		return true
	}

	return false
}

// SetInfo gets a reference to the given NullableString and assigns it to the Info field.
func (o *AutoTestStepResultUpdateRequest) SetInfo(v string) {
	o.Info.Set(&v)
}
// SetInfoNil sets the value for Info to be an explicit nil
func (o *AutoTestStepResultUpdateRequest) SetInfoNil() {
	o.Info.Set(nil)
}

// UnsetInfo ensures that no value is present for Info, not even an explicit nil
func (o *AutoTestStepResultUpdateRequest) UnsetInfo() {
	o.Info.Unset()
}

// GetStartedOn returns the StartedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestStepResultUpdateRequest) GetStartedOn() time.Time {
	if o == nil || IsNil(o.StartedOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartedOn.Get()
}

// GetStartedOnOk returns a tuple with the StartedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestStepResultUpdateRequest) GetStartedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedOn.Get(), o.StartedOn.IsSet()
}

// HasStartedOn returns a boolean if a field has been set.
func (o *AutoTestStepResultUpdateRequest) HasStartedOn() bool {
	if o != nil && o.StartedOn.IsSet() {
		return true
	}

	return false
}

// SetStartedOn gets a reference to the given NullableTime and assigns it to the StartedOn field.
func (o *AutoTestStepResultUpdateRequest) SetStartedOn(v time.Time) {
	o.StartedOn.Set(&v)
}
// SetStartedOnNil sets the value for StartedOn to be an explicit nil
func (o *AutoTestStepResultUpdateRequest) SetStartedOnNil() {
	o.StartedOn.Set(nil)
}

// UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
func (o *AutoTestStepResultUpdateRequest) UnsetStartedOn() {
	o.StartedOn.Unset()
}

// GetCompletedOn returns the CompletedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestStepResultUpdateRequest) GetCompletedOn() time.Time {
	if o == nil || IsNil(o.CompletedOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CompletedOn.Get()
}

// GetCompletedOnOk returns a tuple with the CompletedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestStepResultUpdateRequest) GetCompletedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedOn.Get(), o.CompletedOn.IsSet()
}

// HasCompletedOn returns a boolean if a field has been set.
func (o *AutoTestStepResultUpdateRequest) HasCompletedOn() bool {
	if o != nil && o.CompletedOn.IsSet() {
		return true
	}

	return false
}

// SetCompletedOn gets a reference to the given NullableTime and assigns it to the CompletedOn field.
func (o *AutoTestStepResultUpdateRequest) SetCompletedOn(v time.Time) {
	o.CompletedOn.Set(&v)
}
// SetCompletedOnNil sets the value for CompletedOn to be an explicit nil
func (o *AutoTestStepResultUpdateRequest) SetCompletedOnNil() {
	o.CompletedOn.Set(nil)
}

// UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
func (o *AutoTestStepResultUpdateRequest) UnsetCompletedOn() {
	o.CompletedOn.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestStepResultUpdateRequest) GetDuration() int64 {
	if o == nil || IsNil(o.Duration.Get()) {
		var ret int64
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestStepResultUpdateRequest) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *AutoTestStepResultUpdateRequest) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableInt64 and assigns it to the Duration field.
func (o *AutoTestStepResultUpdateRequest) SetDuration(v int64) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *AutoTestStepResultUpdateRequest) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *AutoTestStepResultUpdateRequest) UnsetDuration() {
	o.Duration.Unset()
}

// GetOutcome returns the Outcome field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestStepResultUpdateRequest) GetOutcome() AvailableTestResultOutcome {
	if o == nil || IsNil(o.Outcome.Get()) {
		var ret AvailableTestResultOutcome
		return ret
	}
	return *o.Outcome.Get()
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestStepResultUpdateRequest) GetOutcomeOk() (*AvailableTestResultOutcome, bool) {
	if o == nil {
		return nil, false
	}
	return o.Outcome.Get(), o.Outcome.IsSet()
}

// HasOutcome returns a boolean if a field has been set.
func (o *AutoTestStepResultUpdateRequest) HasOutcome() bool {
	if o != nil && o.Outcome.IsSet() {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given NullableAvailableTestResultOutcome and assigns it to the Outcome field.
func (o *AutoTestStepResultUpdateRequest) SetOutcome(v AvailableTestResultOutcome) {
	o.Outcome.Set(&v)
}
// SetOutcomeNil sets the value for Outcome to be an explicit nil
func (o *AutoTestStepResultUpdateRequest) SetOutcomeNil() {
	o.Outcome.Set(nil)
}

// UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
func (o *AutoTestStepResultUpdateRequest) UnsetOutcome() {
	o.Outcome.Unset()
}

// GetStepResults returns the StepResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestStepResultUpdateRequest) GetStepResults() []AttachmentPutModelAutoTestStepResultsModel {
	if o == nil {
		var ret []AttachmentPutModelAutoTestStepResultsModel
		return ret
	}
	return o.StepResults
}

// GetStepResultsOk returns a tuple with the StepResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestStepResultUpdateRequest) GetStepResultsOk() ([]AttachmentPutModelAutoTestStepResultsModel, bool) {
	if o == nil || IsNil(o.StepResults) {
		return nil, false
	}
	return o.StepResults, true
}

// HasStepResults returns a boolean if a field has been set.
func (o *AutoTestStepResultUpdateRequest) HasStepResults() bool {
	if o != nil && !IsNil(o.StepResults) {
		return true
	}

	return false
}

// SetStepResults gets a reference to the given []AttachmentPutModelAutoTestStepResultsModel and assigns it to the StepResults field.
func (o *AutoTestStepResultUpdateRequest) SetStepResults(v []AttachmentPutModelAutoTestStepResultsModel) {
	o.StepResults = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestStepResultUpdateRequest) GetAttachments() []AttachmentUpdateRequest {
	if o == nil {
		var ret []AttachmentUpdateRequest
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestStepResultUpdateRequest) GetAttachmentsOk() ([]AttachmentUpdateRequest, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *AutoTestStepResultUpdateRequest) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []AttachmentUpdateRequest and assigns it to the Attachments field.
func (o *AutoTestStepResultUpdateRequest) SetAttachments(v []AttachmentUpdateRequest) {
	o.Attachments = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestStepResultUpdateRequest) GetParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestStepResultUpdateRequest) GetParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *AutoTestStepResultUpdateRequest) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *AutoTestStepResultUpdateRequest) SetParameters(v map[string]string) {
	o.Parameters = v
}

func (o AutoTestStepResultUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoTestStepResultUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Info.IsSet() {
		toSerialize["info"] = o.Info.Get()
	}
	if o.StartedOn.IsSet() {
		toSerialize["startedOn"] = o.StartedOn.Get()
	}
	if o.CompletedOn.IsSet() {
		toSerialize["completedOn"] = o.CompletedOn.Get()
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if o.Outcome.IsSet() {
		toSerialize["outcome"] = o.Outcome.Get()
	}
	if o.StepResults != nil {
		toSerialize["stepResults"] = o.StepResults
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	return toSerialize, nil
}

type NullableAutoTestStepResultUpdateRequest struct {
	value *AutoTestStepResultUpdateRequest
	isSet bool
}

func (v NullableAutoTestStepResultUpdateRequest) Get() *AutoTestStepResultUpdateRequest {
	return v.value
}

func (v *NullableAutoTestStepResultUpdateRequest) Set(val *AutoTestStepResultUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoTestStepResultUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoTestStepResultUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoTestStepResultUpdateRequest(val *AutoTestStepResultUpdateRequest) *NullableAutoTestStepResultUpdateRequest {
	return &NullableAutoTestStepResultUpdateRequest{value: val, isSet: true}
}

func (v NullableAutoTestStepResultUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoTestStepResultUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


