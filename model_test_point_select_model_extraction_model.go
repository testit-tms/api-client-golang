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

// checks if the TestPointSelectModelExtractionModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestPointSelectModelExtractionModel{}

// TestPointSelectModelExtractionModel struct for TestPointSelectModelExtractionModel
type TestPointSelectModelExtractionModel struct {
	Ids NullableTestPointsExtractionModelIds `json:"ids,omitempty"`
}

// NewTestPointSelectModelExtractionModel instantiates a new TestPointSelectModelExtractionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestPointSelectModelExtractionModel() *TestPointSelectModelExtractionModel {
	this := TestPointSelectModelExtractionModel{}
	return &this
}

// NewTestPointSelectModelExtractionModelWithDefaults instantiates a new TestPointSelectModelExtractionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestPointSelectModelExtractionModelWithDefaults() *TestPointSelectModelExtractionModel {
	this := TestPointSelectModelExtractionModel{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestPointSelectModelExtractionModel) GetIds() TestPointsExtractionModelIds {
	if o == nil || IsNil(o.Ids.Get()) {
		var ret TestPointsExtractionModelIds
		return ret
	}
	return *o.Ids.Get()
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestPointSelectModelExtractionModel) GetIdsOk() (*TestPointsExtractionModelIds, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids.Get(), o.Ids.IsSet()
}

// HasIds returns a boolean if a field has been set.
func (o *TestPointSelectModelExtractionModel) HasIds() bool {
	if o != nil && o.Ids.IsSet() {
		return true
	}

	return false
}

// SetIds gets a reference to the given NullableTestPointsExtractionModelIds and assigns it to the Ids field.
func (o *TestPointSelectModelExtractionModel) SetIds(v TestPointsExtractionModelIds) {
	o.Ids.Set(&v)
}
// SetIdsNil sets the value for Ids to be an explicit nil
func (o *TestPointSelectModelExtractionModel) SetIdsNil() {
	o.Ids.Set(nil)
}

// UnsetIds ensures that no value is present for Ids, not even an explicit nil
func (o *TestPointSelectModelExtractionModel) UnsetIds() {
	o.Ids.Unset()
}

func (o TestPointSelectModelExtractionModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestPointSelectModelExtractionModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids.IsSet() {
		toSerialize["ids"] = o.Ids.Get()
	}
	return toSerialize, nil
}

type NullableTestPointSelectModelExtractionModel struct {
	value *TestPointSelectModelExtractionModel
	isSet bool
}

func (v NullableTestPointSelectModelExtractionModel) Get() *TestPointSelectModelExtractionModel {
	return v.value
}

func (v *NullableTestPointSelectModelExtractionModel) Set(val *TestPointSelectModelExtractionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestPointSelectModelExtractionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestPointSelectModelExtractionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestPointSelectModelExtractionModel(val *TestPointSelectModelExtractionModel) *NullableTestPointSelectModelExtractionModel {
	return &NullableTestPointSelectModelExtractionModel{value: val, isSet: true}
}

func (v NullableTestPointSelectModelExtractionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestPointSelectModelExtractionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


