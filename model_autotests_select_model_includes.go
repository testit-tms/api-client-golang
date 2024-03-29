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

// checks if the AutotestsSelectModelIncludes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutotestsSelectModelIncludes{}

// AutotestsSelectModelIncludes Object specifying data to be included
type AutotestsSelectModelIncludes struct {
	// If autotest steps will be included
	IncludeSteps bool `json:"includeSteps"`
	// If autotest links will be included
	IncludeLinks bool `json:"includeLinks"`
	// If autotest labels will be included
	IncludeLabels bool `json:"includeLabels"`
}

// NewAutotestsSelectModelIncludes instantiates a new AutotestsSelectModelIncludes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutotestsSelectModelIncludes(includeSteps bool, includeLinks bool, includeLabels bool) *AutotestsSelectModelIncludes {
	this := AutotestsSelectModelIncludes{}
	this.IncludeSteps = includeSteps
	this.IncludeLinks = includeLinks
	this.IncludeLabels = includeLabels
	return &this
}

// NewAutotestsSelectModelIncludesWithDefaults instantiates a new AutotestsSelectModelIncludes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutotestsSelectModelIncludesWithDefaults() *AutotestsSelectModelIncludes {
	this := AutotestsSelectModelIncludes{}
	return &this
}

// GetIncludeSteps returns the IncludeSteps field value
func (o *AutotestsSelectModelIncludes) GetIncludeSteps() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludeSteps
}

// GetIncludeStepsOk returns a tuple with the IncludeSteps field value
// and a boolean to check if the value has been set.
func (o *AutotestsSelectModelIncludes) GetIncludeStepsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeSteps, true
}

// SetIncludeSteps sets field value
func (o *AutotestsSelectModelIncludes) SetIncludeSteps(v bool) {
	o.IncludeSteps = v
}

// GetIncludeLinks returns the IncludeLinks field value
func (o *AutotestsSelectModelIncludes) GetIncludeLinks() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludeLinks
}

// GetIncludeLinksOk returns a tuple with the IncludeLinks field value
// and a boolean to check if the value has been set.
func (o *AutotestsSelectModelIncludes) GetIncludeLinksOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeLinks, true
}

// SetIncludeLinks sets field value
func (o *AutotestsSelectModelIncludes) SetIncludeLinks(v bool) {
	o.IncludeLinks = v
}

// GetIncludeLabels returns the IncludeLabels field value
func (o *AutotestsSelectModelIncludes) GetIncludeLabels() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludeLabels
}

// GetIncludeLabelsOk returns a tuple with the IncludeLabels field value
// and a boolean to check if the value has been set.
func (o *AutotestsSelectModelIncludes) GetIncludeLabelsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeLabels, true
}

// SetIncludeLabels sets field value
func (o *AutotestsSelectModelIncludes) SetIncludeLabels(v bool) {
	o.IncludeLabels = v
}

func (o AutotestsSelectModelIncludes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutotestsSelectModelIncludes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["includeSteps"] = o.IncludeSteps
	toSerialize["includeLinks"] = o.IncludeLinks
	toSerialize["includeLabels"] = o.IncludeLabels
	return toSerialize, nil
}

type NullableAutotestsSelectModelIncludes struct {
	value *AutotestsSelectModelIncludes
	isSet bool
}

func (v NullableAutotestsSelectModelIncludes) Get() *AutotestsSelectModelIncludes {
	return v.value
}

func (v *NullableAutotestsSelectModelIncludes) Set(val *AutotestsSelectModelIncludes) {
	v.value = val
	v.isSet = true
}

func (v NullableAutotestsSelectModelIncludes) IsSet() bool {
	return v.isSet
}

func (v *NullableAutotestsSelectModelIncludes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutotestsSelectModelIncludes(val *AutotestsSelectModelIncludes) *NullableAutotestsSelectModelIncludes {
	return &NullableAutotestsSelectModelIncludes{value: val, isSet: true}
}

func (v NullableAutotestsSelectModelIncludes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutotestsSelectModelIncludes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


