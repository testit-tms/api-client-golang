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

// checks if the StepResultModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StepResultModel{}

// StepResultModel struct for StepResultModel
type StepResultModel struct {
	StepId string `json:"stepId"`
	Outcome string `json:"outcome"`
	SharedStepVersionId NullableString `json:"sharedStepVersionId,omitempty"`
	SharedStepResults []SharedStepResultModel `json:"sharedStepResults,omitempty"`
	Comment NullableStepCommentModel `json:"comment,omitempty"`
}

// NewStepResultModel instantiates a new StepResultModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStepResultModel(stepId string, outcome string) *StepResultModel {
	this := StepResultModel{}
	this.StepId = stepId
	this.Outcome = outcome
	return &this
}

// NewStepResultModelWithDefaults instantiates a new StepResultModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStepResultModelWithDefaults() *StepResultModel {
	this := StepResultModel{}
	return &this
}

// GetStepId returns the StepId field value
func (o *StepResultModel) GetStepId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StepId
}

// GetStepIdOk returns a tuple with the StepId field value
// and a boolean to check if the value has been set.
func (o *StepResultModel) GetStepIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StepId, true
}

// SetStepId sets field value
func (o *StepResultModel) SetStepId(v string) {
	o.StepId = v
}

// GetOutcome returns the Outcome field value
func (o *StepResultModel) GetOutcome() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value
// and a boolean to check if the value has been set.
func (o *StepResultModel) GetOutcomeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outcome, true
}

// SetOutcome sets field value
func (o *StepResultModel) SetOutcome(v string) {
	o.Outcome = v
}

// GetSharedStepVersionId returns the SharedStepVersionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StepResultModel) GetSharedStepVersionId() string {
	if o == nil || IsNil(o.SharedStepVersionId.Get()) {
		var ret string
		return ret
	}
	return *o.SharedStepVersionId.Get()
}

// GetSharedStepVersionIdOk returns a tuple with the SharedStepVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StepResultModel) GetSharedStepVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedStepVersionId.Get(), o.SharedStepVersionId.IsSet()
}

// HasSharedStepVersionId returns a boolean if a field has been set.
func (o *StepResultModel) HasSharedStepVersionId() bool {
	if o != nil && o.SharedStepVersionId.IsSet() {
		return true
	}

	return false
}

// SetSharedStepVersionId gets a reference to the given NullableString and assigns it to the SharedStepVersionId field.
func (o *StepResultModel) SetSharedStepVersionId(v string) {
	o.SharedStepVersionId.Set(&v)
}
// SetSharedStepVersionIdNil sets the value for SharedStepVersionId to be an explicit nil
func (o *StepResultModel) SetSharedStepVersionIdNil() {
	o.SharedStepVersionId.Set(nil)
}

// UnsetSharedStepVersionId ensures that no value is present for SharedStepVersionId, not even an explicit nil
func (o *StepResultModel) UnsetSharedStepVersionId() {
	o.SharedStepVersionId.Unset()
}

// GetSharedStepResults returns the SharedStepResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StepResultModel) GetSharedStepResults() []SharedStepResultModel {
	if o == nil {
		var ret []SharedStepResultModel
		return ret
	}
	return o.SharedStepResults
}

// GetSharedStepResultsOk returns a tuple with the SharedStepResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StepResultModel) GetSharedStepResultsOk() ([]SharedStepResultModel, bool) {
	if o == nil || IsNil(o.SharedStepResults) {
		return nil, false
	}
	return o.SharedStepResults, true
}

// HasSharedStepResults returns a boolean if a field has been set.
func (o *StepResultModel) HasSharedStepResults() bool {
	if o != nil && IsNil(o.SharedStepResults) {
		return true
	}

	return false
}

// SetSharedStepResults gets a reference to the given []SharedStepResultModel and assigns it to the SharedStepResults field.
func (o *StepResultModel) SetSharedStepResults(v []SharedStepResultModel) {
	o.SharedStepResults = v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StepResultModel) GetComment() StepCommentModel {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret StepCommentModel
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StepResultModel) GetCommentOk() (*StepCommentModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *StepResultModel) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableStepCommentModel and assigns it to the Comment field.
func (o *StepResultModel) SetComment(v StepCommentModel) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *StepResultModel) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *StepResultModel) UnsetComment() {
	o.Comment.Unset()
}

func (o StepResultModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StepResultModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stepId"] = o.StepId
	toSerialize["outcome"] = o.Outcome
	if o.SharedStepVersionId.IsSet() {
		toSerialize["sharedStepVersionId"] = o.SharedStepVersionId.Get()
	}
	if o.SharedStepResults != nil {
		toSerialize["sharedStepResults"] = o.SharedStepResults
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	return toSerialize, nil
}

type NullableStepResultModel struct {
	value *StepResultModel
	isSet bool
}

func (v NullableStepResultModel) Get() *StepResultModel {
	return v.value
}

func (v *NullableStepResultModel) Set(val *StepResultModel) {
	v.value = val
	v.isSet = true
}

func (v NullableStepResultModel) IsSet() bool {
	return v.isSet
}

func (v *NullableStepResultModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStepResultModel(val *StepResultModel) *NullableStepResultModel {
	return &NullableStepResultModel{value: val, isSet: true}
}

func (v NullableStepResultModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStepResultModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


