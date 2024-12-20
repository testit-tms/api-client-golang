/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AutotestsSelectModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutotestsSelectModel{}

// AutotestsSelectModel struct for AutotestsSelectModel
type AutotestsSelectModel struct {
	// Object containing different filters to adjust search
	Filter AutotestFilterModel `json:"filter"`
	// Object specifying data to be included
	Includes SearchAutoTestsQueryIncludesModel `json:"includes"`
}

type _AutotestsSelectModel AutotestsSelectModel

// NewAutotestsSelectModel instantiates a new AutotestsSelectModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutotestsSelectModel(filter AutotestFilterModel, includes SearchAutoTestsQueryIncludesModel) *AutotestsSelectModel {
	this := AutotestsSelectModel{}
	this.Filter = filter
	this.Includes = includes
	return &this
}

// NewAutotestsSelectModelWithDefaults instantiates a new AutotestsSelectModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutotestsSelectModelWithDefaults() *AutotestsSelectModel {
	this := AutotestsSelectModel{}
	return &this
}

// GetFilter returns the Filter field value
func (o *AutotestsSelectModel) GetFilter() AutotestFilterModel {
	if o == nil {
		var ret AutotestFilterModel
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *AutotestsSelectModel) GetFilterOk() (*AutotestFilterModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *AutotestsSelectModel) SetFilter(v AutotestFilterModel) {
	o.Filter = v
}

// GetIncludes returns the Includes field value
func (o *AutotestsSelectModel) GetIncludes() SearchAutoTestsQueryIncludesModel {
	if o == nil {
		var ret SearchAutoTestsQueryIncludesModel
		return ret
	}

	return o.Includes
}

// GetIncludesOk returns a tuple with the Includes field value
// and a boolean to check if the value has been set.
func (o *AutotestsSelectModel) GetIncludesOk() (*SearchAutoTestsQueryIncludesModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Includes, true
}

// SetIncludes sets field value
func (o *AutotestsSelectModel) SetIncludes(v SearchAutoTestsQueryIncludesModel) {
	o.Includes = v
}

func (o AutotestsSelectModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutotestsSelectModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filter"] = o.Filter
	toSerialize["includes"] = o.Includes
	return toSerialize, nil
}

func (o *AutotestsSelectModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filter",
		"includes",
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

	varAutotestsSelectModel := _AutotestsSelectModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAutotestsSelectModel)

	if err != nil {
		return err
	}

	*o = AutotestsSelectModel(varAutotestsSelectModel)

	return err
}

type NullableAutotestsSelectModel struct {
	value *AutotestsSelectModel
	isSet bool
}

func (v NullableAutotestsSelectModel) Get() *AutotestsSelectModel {
	return v.value
}

func (v *NullableAutotestsSelectModel) Set(val *AutotestsSelectModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAutotestsSelectModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAutotestsSelectModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutotestsSelectModel(val *AutotestsSelectModel) *NullableAutotestsSelectModel {
	return &NullableAutotestsSelectModel{value: val, isSet: true}
}

func (v NullableAutotestsSelectModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutotestsSelectModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


