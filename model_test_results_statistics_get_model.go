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

// checks if the TestResultsStatisticsGetModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestResultsStatisticsGetModel{}

// TestResultsStatisticsGetModel struct for TestResultsStatisticsGetModel
type TestResultsStatisticsGetModel struct {
	Statuses TestResultsStatisticsGetModelStatuses `json:"statuses"`
	FailureCategories TestResultsStatisticsGetModelFailureCategories `json:"failureCategories"`
}

type _TestResultsStatisticsGetModel TestResultsStatisticsGetModel

// NewTestResultsStatisticsGetModel instantiates a new TestResultsStatisticsGetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestResultsStatisticsGetModel(statuses TestResultsStatisticsGetModelStatuses, failureCategories TestResultsStatisticsGetModelFailureCategories) *TestResultsStatisticsGetModel {
	this := TestResultsStatisticsGetModel{}
	this.Statuses = statuses
	this.FailureCategories = failureCategories
	return &this
}

// NewTestResultsStatisticsGetModelWithDefaults instantiates a new TestResultsStatisticsGetModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestResultsStatisticsGetModelWithDefaults() *TestResultsStatisticsGetModel {
	this := TestResultsStatisticsGetModel{}
	return &this
}

// GetStatuses returns the Statuses field value
func (o *TestResultsStatisticsGetModel) GetStatuses() TestResultsStatisticsGetModelStatuses {
	if o == nil {
		var ret TestResultsStatisticsGetModelStatuses
		return ret
	}

	return o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value
// and a boolean to check if the value has been set.
func (o *TestResultsStatisticsGetModel) GetStatusesOk() (*TestResultsStatisticsGetModelStatuses, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Statuses, true
}

// SetStatuses sets field value
func (o *TestResultsStatisticsGetModel) SetStatuses(v TestResultsStatisticsGetModelStatuses) {
	o.Statuses = v
}

// GetFailureCategories returns the FailureCategories field value
func (o *TestResultsStatisticsGetModel) GetFailureCategories() TestResultsStatisticsGetModelFailureCategories {
	if o == nil {
		var ret TestResultsStatisticsGetModelFailureCategories
		return ret
	}

	return o.FailureCategories
}

// GetFailureCategoriesOk returns a tuple with the FailureCategories field value
// and a boolean to check if the value has been set.
func (o *TestResultsStatisticsGetModel) GetFailureCategoriesOk() (*TestResultsStatisticsGetModelFailureCategories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailureCategories, true
}

// SetFailureCategories sets field value
func (o *TestResultsStatisticsGetModel) SetFailureCategories(v TestResultsStatisticsGetModelFailureCategories) {
	o.FailureCategories = v
}

func (o TestResultsStatisticsGetModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestResultsStatisticsGetModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["statuses"] = o.Statuses
	toSerialize["failureCategories"] = o.FailureCategories
	return toSerialize, nil
}

func (o *TestResultsStatisticsGetModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"statuses",
		"failureCategories",
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

	varTestResultsStatisticsGetModel := _TestResultsStatisticsGetModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestResultsStatisticsGetModel)

	if err != nil {
		return err
	}

	*o = TestResultsStatisticsGetModel(varTestResultsStatisticsGetModel)

	return err
}

type NullableTestResultsStatisticsGetModel struct {
	value *TestResultsStatisticsGetModel
	isSet bool
}

func (v NullableTestResultsStatisticsGetModel) Get() *TestResultsStatisticsGetModel {
	return v.value
}

func (v *NullableTestResultsStatisticsGetModel) Set(val *TestResultsStatisticsGetModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestResultsStatisticsGetModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestResultsStatisticsGetModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestResultsStatisticsGetModel(val *TestResultsStatisticsGetModel) *NullableTestResultsStatisticsGetModel {
	return &NullableTestResultsStatisticsGetModel{value: val, isSet: true}
}

func (v NullableTestResultsStatisticsGetModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestResultsStatisticsGetModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


