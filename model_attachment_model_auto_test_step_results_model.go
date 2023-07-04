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

// checks if the AttachmentModelAutoTestStepResultsModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachmentModelAutoTestStepResultsModel{}

// AttachmentModelAutoTestStepResultsModel struct for AttachmentModelAutoTestStepResultsModel
type AttachmentModelAutoTestStepResultsModel struct {
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
	Outcome NullableAvailableTestResultOutcome `json:"outcome,omitempty"`
	// Nested step results. The maximum nesting level is 15.
	StepResults []AttachmentModelAutoTestStepResultsModel `json:"stepResults,omitempty"`
	// /// <summary>  Specifies an attachment GUID. Multiple values can be sent.  </summary>
	Attachments []AttachmentModel `json:"attachments,omitempty"`
	// \"<b>parameter</b>\": \"<b>value</b>\" pair with arbitrary custom parameters. Multiple parameters can be sent.
	Parameters map[string]string `json:"parameters,omitempty"`
}

// NewAttachmentModelAutoTestStepResultsModel instantiates a new AttachmentModelAutoTestStepResultsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentModelAutoTestStepResultsModel() *AttachmentModelAutoTestStepResultsModel {
	this := AttachmentModelAutoTestStepResultsModel{}
	return &this
}

// NewAttachmentModelAutoTestStepResultsModelWithDefaults instantiates a new AttachmentModelAutoTestStepResultsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentModelAutoTestStepResultsModelWithDefaults() *AttachmentModelAutoTestStepResultsModel {
	this := AttachmentModelAutoTestStepResultsModel{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentModelAutoTestStepResultsModel) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentModelAutoTestStepResultsModel) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *AttachmentModelAutoTestStepResultsModel) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *AttachmentModelAutoTestStepResultsModel) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *AttachmentModelAutoTestStepResultsModel) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *AttachmentModelAutoTestStepResultsModel) UnsetTitle() {
	o.Title.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentModelAutoTestStepResultsModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentModelAutoTestStepResultsModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AttachmentModelAutoTestStepResultsModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AttachmentModelAutoTestStepResultsModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AttachmentModelAutoTestStepResultsModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AttachmentModelAutoTestStepResultsModel) UnsetDescription() {
	o.Description.Unset()
}

// GetInfo returns the Info field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentModelAutoTestStepResultsModel) GetInfo() string {
	if o == nil || IsNil(o.Info.Get()) {
		var ret string
		return ret
	}
	return *o.Info.Get()
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentModelAutoTestStepResultsModel) GetInfoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Info.Get(), o.Info.IsSet()
}

// HasInfo returns a boolean if a field has been set.
func (o *AttachmentModelAutoTestStepResultsModel) HasInfo() bool {
	if o != nil && o.Info.IsSet() {
		return true
	}

	return false
}

// SetInfo gets a reference to the given NullableString and assigns it to the Info field.
func (o *AttachmentModelAutoTestStepResultsModel) SetInfo(v string) {
	o.Info.Set(&v)
}
// SetInfoNil sets the value for Info to be an explicit nil
func (o *AttachmentModelAutoTestStepResultsModel) SetInfoNil() {
	o.Info.Set(nil)
}

// UnsetInfo ensures that no value is present for Info, not even an explicit nil
func (o *AttachmentModelAutoTestStepResultsModel) UnsetInfo() {
	o.Info.Unset()
}

// GetStartedOn returns the StartedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentModelAutoTestStepResultsModel) GetStartedOn() time.Time {
	if o == nil || IsNil(o.StartedOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartedOn.Get()
}

// GetStartedOnOk returns a tuple with the StartedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentModelAutoTestStepResultsModel) GetStartedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedOn.Get(), o.StartedOn.IsSet()
}

// HasStartedOn returns a boolean if a field has been set.
func (o *AttachmentModelAutoTestStepResultsModel) HasStartedOn() bool {
	if o != nil && o.StartedOn.IsSet() {
		return true
	}

	return false
}

// SetStartedOn gets a reference to the given NullableTime and assigns it to the StartedOn field.
func (o *AttachmentModelAutoTestStepResultsModel) SetStartedOn(v time.Time) {
	o.StartedOn.Set(&v)
}
// SetStartedOnNil sets the value for StartedOn to be an explicit nil
func (o *AttachmentModelAutoTestStepResultsModel) SetStartedOnNil() {
	o.StartedOn.Set(nil)
}

// UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
func (o *AttachmentModelAutoTestStepResultsModel) UnsetStartedOn() {
	o.StartedOn.Unset()
}

// GetCompletedOn returns the CompletedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentModelAutoTestStepResultsModel) GetCompletedOn() time.Time {
	if o == nil || IsNil(o.CompletedOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CompletedOn.Get()
}

// GetCompletedOnOk returns a tuple with the CompletedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentModelAutoTestStepResultsModel) GetCompletedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedOn.Get(), o.CompletedOn.IsSet()
}

// HasCompletedOn returns a boolean if a field has been set.
func (o *AttachmentModelAutoTestStepResultsModel) HasCompletedOn() bool {
	if o != nil && o.CompletedOn.IsSet() {
		return true
	}

	return false
}

// SetCompletedOn gets a reference to the given NullableTime and assigns it to the CompletedOn field.
func (o *AttachmentModelAutoTestStepResultsModel) SetCompletedOn(v time.Time) {
	o.CompletedOn.Set(&v)
}
// SetCompletedOnNil sets the value for CompletedOn to be an explicit nil
func (o *AttachmentModelAutoTestStepResultsModel) SetCompletedOnNil() {
	o.CompletedOn.Set(nil)
}

// UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
func (o *AttachmentModelAutoTestStepResultsModel) UnsetCompletedOn() {
	o.CompletedOn.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentModelAutoTestStepResultsModel) GetDuration() int64 {
	if o == nil || IsNil(o.Duration.Get()) {
		var ret int64
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentModelAutoTestStepResultsModel) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *AttachmentModelAutoTestStepResultsModel) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableInt64 and assigns it to the Duration field.
func (o *AttachmentModelAutoTestStepResultsModel) SetDuration(v int64) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *AttachmentModelAutoTestStepResultsModel) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *AttachmentModelAutoTestStepResultsModel) UnsetDuration() {
	o.Duration.Unset()
}

// GetOutcome returns the Outcome field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentModelAutoTestStepResultsModel) GetOutcome() AvailableTestResultOutcome {
	if o == nil || IsNil(o.Outcome.Get()) {
		var ret AvailableTestResultOutcome
		return ret
	}
	return *o.Outcome.Get()
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentModelAutoTestStepResultsModel) GetOutcomeOk() (*AvailableTestResultOutcome, bool) {
	if o == nil {
		return nil, false
	}
	return o.Outcome.Get(), o.Outcome.IsSet()
}

// HasOutcome returns a boolean if a field has been set.
func (o *AttachmentModelAutoTestStepResultsModel) HasOutcome() bool {
	if o != nil && o.Outcome.IsSet() {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given NullableAvailableTestResultOutcome and assigns it to the Outcome field.
func (o *AttachmentModelAutoTestStepResultsModel) SetOutcome(v AvailableTestResultOutcome) {
	o.Outcome.Set(&v)
}
// SetOutcomeNil sets the value for Outcome to be an explicit nil
func (o *AttachmentModelAutoTestStepResultsModel) SetOutcomeNil() {
	o.Outcome.Set(nil)
}

// UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
func (o *AttachmentModelAutoTestStepResultsModel) UnsetOutcome() {
	o.Outcome.Unset()
}

// GetStepResults returns the StepResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentModelAutoTestStepResultsModel) GetStepResults() []AttachmentModelAutoTestStepResultsModel {
	if o == nil {
		var ret []AttachmentModelAutoTestStepResultsModel
		return ret
	}
	return o.StepResults
}

// GetStepResultsOk returns a tuple with the StepResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentModelAutoTestStepResultsModel) GetStepResultsOk() ([]AttachmentModelAutoTestStepResultsModel, bool) {
	if o == nil || IsNil(o.StepResults) {
		return nil, false
	}
	return o.StepResults, true
}

// HasStepResults returns a boolean if a field has been set.
func (o *AttachmentModelAutoTestStepResultsModel) HasStepResults() bool {
	if o != nil && IsNil(o.StepResults) {
		return true
	}

	return false
}

// SetStepResults gets a reference to the given []AttachmentModelAutoTestStepResultsModel and assigns it to the StepResults field.
func (o *AttachmentModelAutoTestStepResultsModel) SetStepResults(v []AttachmentModelAutoTestStepResultsModel) {
	o.StepResults = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentModelAutoTestStepResultsModel) GetAttachments() []AttachmentModel {
	if o == nil {
		var ret []AttachmentModel
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentModelAutoTestStepResultsModel) GetAttachmentsOk() ([]AttachmentModel, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *AttachmentModelAutoTestStepResultsModel) HasAttachments() bool {
	if o != nil && IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []AttachmentModel and assigns it to the Attachments field.
func (o *AttachmentModelAutoTestStepResultsModel) SetAttachments(v []AttachmentModel) {
	o.Attachments = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentModelAutoTestStepResultsModel) GetParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentModelAutoTestStepResultsModel) GetParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *AttachmentModelAutoTestStepResultsModel) HasParameters() bool {
	if o != nil && IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *AttachmentModelAutoTestStepResultsModel) SetParameters(v map[string]string) {
	o.Parameters = v
}

func (o AttachmentModelAutoTestStepResultsModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachmentModelAutoTestStepResultsModel) ToMap() (map[string]interface{}, error) {
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

type NullableAttachmentModelAutoTestStepResultsModel struct {
	value *AttachmentModelAutoTestStepResultsModel
	isSet bool
}

func (v NullableAttachmentModelAutoTestStepResultsModel) Get() *AttachmentModelAutoTestStepResultsModel {
	return v.value
}

func (v *NullableAttachmentModelAutoTestStepResultsModel) Set(val *AttachmentModelAutoTestStepResultsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentModelAutoTestStepResultsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentModelAutoTestStepResultsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentModelAutoTestStepResultsModel(val *AttachmentModelAutoTestStepResultsModel) *NullableAttachmentModelAutoTestStepResultsModel {
	return &NullableAttachmentModelAutoTestStepResultsModel{value: val, isSet: true}
}

func (v NullableAttachmentModelAutoTestStepResultsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentModelAutoTestStepResultsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


