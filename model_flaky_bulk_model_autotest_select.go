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

// checks if the FlakyBulkModelAutotestSelect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlakyBulkModelAutotestSelect{}

// FlakyBulkModelAutotestSelect struct for FlakyBulkModelAutotestSelect
type FlakyBulkModelAutotestSelect struct {
	Filter AutotestSelectModelFilter `json:"filter"`
	ExtractionModel AutotestSelectModelExtractionModel `json:"extractionModel"`
}

// NewFlakyBulkModelAutotestSelect instantiates a new FlakyBulkModelAutotestSelect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlakyBulkModelAutotestSelect(filter AutotestSelectModelFilter, extractionModel AutotestSelectModelExtractionModel) *FlakyBulkModelAutotestSelect {
	this := FlakyBulkModelAutotestSelect{}
	this.Filter = filter
	this.ExtractionModel = extractionModel
	return &this
}

// NewFlakyBulkModelAutotestSelectWithDefaults instantiates a new FlakyBulkModelAutotestSelect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlakyBulkModelAutotestSelectWithDefaults() *FlakyBulkModelAutotestSelect {
	this := FlakyBulkModelAutotestSelect{}
	return &this
}

// GetFilter returns the Filter field value
func (o *FlakyBulkModelAutotestSelect) GetFilter() AutotestSelectModelFilter {
	if o == nil {
		var ret AutotestSelectModelFilter
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *FlakyBulkModelAutotestSelect) GetFilterOk() (*AutotestSelectModelFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *FlakyBulkModelAutotestSelect) SetFilter(v AutotestSelectModelFilter) {
	o.Filter = v
}

// GetExtractionModel returns the ExtractionModel field value
func (o *FlakyBulkModelAutotestSelect) GetExtractionModel() AutotestSelectModelExtractionModel {
	if o == nil {
		var ret AutotestSelectModelExtractionModel
		return ret
	}

	return o.ExtractionModel
}

// GetExtractionModelOk returns a tuple with the ExtractionModel field value
// and a boolean to check if the value has been set.
func (o *FlakyBulkModelAutotestSelect) GetExtractionModelOk() (*AutotestSelectModelExtractionModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtractionModel, true
}

// SetExtractionModel sets field value
func (o *FlakyBulkModelAutotestSelect) SetExtractionModel(v AutotestSelectModelExtractionModel) {
	o.ExtractionModel = v
}

func (o FlakyBulkModelAutotestSelect) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlakyBulkModelAutotestSelect) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filter"] = o.Filter
	toSerialize["extractionModel"] = o.ExtractionModel
	return toSerialize, nil
}

type NullableFlakyBulkModelAutotestSelect struct {
	value *FlakyBulkModelAutotestSelect
	isSet bool
}

func (v NullableFlakyBulkModelAutotestSelect) Get() *FlakyBulkModelAutotestSelect {
	return v.value
}

func (v *NullableFlakyBulkModelAutotestSelect) Set(val *FlakyBulkModelAutotestSelect) {
	v.value = val
	v.isSet = true
}

func (v NullableFlakyBulkModelAutotestSelect) IsSet() bool {
	return v.isSet
}

func (v *NullableFlakyBulkModelAutotestSelect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlakyBulkModelAutotestSelect(val *FlakyBulkModelAutotestSelect) *NullableFlakyBulkModelAutotestSelect {
	return &NullableFlakyBulkModelAutotestSelect{value: val, isSet: true}
}

func (v NullableFlakyBulkModelAutotestSelect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlakyBulkModelAutotestSelect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


