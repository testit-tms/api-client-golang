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

// checks if the SharedStepResultModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharedStepResultModel{}

// SharedStepResultModel struct for SharedStepResultModel
type SharedStepResultModel struct {
	StepId *string `json:"stepId,omitempty"`
	Outcome *string `json:"outcome,omitempty"`
}

// NewSharedStepResultModel instantiates a new SharedStepResultModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedStepResultModel() *SharedStepResultModel {
	this := SharedStepResultModel{}
	return &this
}

// NewSharedStepResultModelWithDefaults instantiates a new SharedStepResultModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedStepResultModelWithDefaults() *SharedStepResultModel {
	this := SharedStepResultModel{}
	return &this
}

// GetStepId returns the StepId field value if set, zero value otherwise.
func (o *SharedStepResultModel) GetStepId() string {
	if o == nil || IsNil(o.StepId) {
		var ret string
		return ret
	}
	return *o.StepId
}

// GetStepIdOk returns a tuple with the StepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedStepResultModel) GetStepIdOk() (*string, bool) {
	if o == nil || IsNil(o.StepId) {
		return nil, false
	}
	return o.StepId, true
}

// HasStepId returns a boolean if a field has been set.
func (o *SharedStepResultModel) HasStepId() bool {
	if o != nil && !IsNil(o.StepId) {
		return true
	}

	return false
}

// SetStepId gets a reference to the given string and assigns it to the StepId field.
func (o *SharedStepResultModel) SetStepId(v string) {
	o.StepId = &v
}

// GetOutcome returns the Outcome field value if set, zero value otherwise.
func (o *SharedStepResultModel) GetOutcome() string {
	if o == nil || IsNil(o.Outcome) {
		var ret string
		return ret
	}
	return *o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedStepResultModel) GetOutcomeOk() (*string, bool) {
	if o == nil || IsNil(o.Outcome) {
		return nil, false
	}
	return o.Outcome, true
}

// HasOutcome returns a boolean if a field has been set.
func (o *SharedStepResultModel) HasOutcome() bool {
	if o != nil && !IsNil(o.Outcome) {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given string and assigns it to the Outcome field.
func (o *SharedStepResultModel) SetOutcome(v string) {
	o.Outcome = &v
}

func (o SharedStepResultModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharedStepResultModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StepId) {
		toSerialize["stepId"] = o.StepId
	}
	if !IsNil(o.Outcome) {
		toSerialize["outcome"] = o.Outcome
	}
	return toSerialize, nil
}

type NullableSharedStepResultModel struct {
	value *SharedStepResultModel
	isSet bool
}

func (v NullableSharedStepResultModel) Get() *SharedStepResultModel {
	return v.value
}

func (v *NullableSharedStepResultModel) Set(val *SharedStepResultModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedStepResultModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedStepResultModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedStepResultModel(val *SharedStepResultModel) *NullableSharedStepResultModel {
	return &NullableSharedStepResultModel{value: val, isSet: true}
}

func (v NullableSharedStepResultModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedStepResultModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


