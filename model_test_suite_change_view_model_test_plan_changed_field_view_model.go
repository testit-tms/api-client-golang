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

// checks if the TestSuiteChangeViewModelTestPlanChangedFieldViewModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestSuiteChangeViewModelTestPlanChangedFieldViewModel{}

// TestSuiteChangeViewModelTestPlanChangedFieldViewModel struct for TestSuiteChangeViewModelTestPlanChangedFieldViewModel
type TestSuiteChangeViewModelTestPlanChangedFieldViewModel struct {
	OldValue *TestSuiteChangeViewModel `json:"oldValue,omitempty"`
	NewValue *TestSuiteChangeViewModel `json:"newValue,omitempty"`
}

// NewTestSuiteChangeViewModelTestPlanChangedFieldViewModel instantiates a new TestSuiteChangeViewModelTestPlanChangedFieldViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestSuiteChangeViewModelTestPlanChangedFieldViewModel() *TestSuiteChangeViewModelTestPlanChangedFieldViewModel {
	this := TestSuiteChangeViewModelTestPlanChangedFieldViewModel{}
	return &this
}

// NewTestSuiteChangeViewModelTestPlanChangedFieldViewModelWithDefaults instantiates a new TestSuiteChangeViewModelTestPlanChangedFieldViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestSuiteChangeViewModelTestPlanChangedFieldViewModelWithDefaults() *TestSuiteChangeViewModelTestPlanChangedFieldViewModel {
	this := TestSuiteChangeViewModelTestPlanChangedFieldViewModel{}
	return &this
}

// GetOldValue returns the OldValue field value if set, zero value otherwise.
func (o *TestSuiteChangeViewModelTestPlanChangedFieldViewModel) GetOldValue() TestSuiteChangeViewModel {
	if o == nil || IsNil(o.OldValue) {
		var ret TestSuiteChangeViewModel
		return ret
	}
	return *o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSuiteChangeViewModelTestPlanChangedFieldViewModel) GetOldValueOk() (*TestSuiteChangeViewModel, bool) {
	if o == nil || IsNil(o.OldValue) {
		return nil, false
	}
	return o.OldValue, true
}

// HasOldValue returns a boolean if a field has been set.
func (o *TestSuiteChangeViewModelTestPlanChangedFieldViewModel) HasOldValue() bool {
	if o != nil && !IsNil(o.OldValue) {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given TestSuiteChangeViewModel and assigns it to the OldValue field.
func (o *TestSuiteChangeViewModelTestPlanChangedFieldViewModel) SetOldValue(v TestSuiteChangeViewModel) {
	o.OldValue = &v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise.
func (o *TestSuiteChangeViewModelTestPlanChangedFieldViewModel) GetNewValue() TestSuiteChangeViewModel {
	if o == nil || IsNil(o.NewValue) {
		var ret TestSuiteChangeViewModel
		return ret
	}
	return *o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSuiteChangeViewModelTestPlanChangedFieldViewModel) GetNewValueOk() (*TestSuiteChangeViewModel, bool) {
	if o == nil || IsNil(o.NewValue) {
		return nil, false
	}
	return o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *TestSuiteChangeViewModelTestPlanChangedFieldViewModel) HasNewValue() bool {
	if o != nil && !IsNil(o.NewValue) {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given TestSuiteChangeViewModel and assigns it to the NewValue field.
func (o *TestSuiteChangeViewModelTestPlanChangedFieldViewModel) SetNewValue(v TestSuiteChangeViewModel) {
	o.NewValue = &v
}

func (o TestSuiteChangeViewModelTestPlanChangedFieldViewModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestSuiteChangeViewModelTestPlanChangedFieldViewModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OldValue) {
		toSerialize["oldValue"] = o.OldValue
	}
	if !IsNil(o.NewValue) {
		toSerialize["newValue"] = o.NewValue
	}
	return toSerialize, nil
}

type NullableTestSuiteChangeViewModelTestPlanChangedFieldViewModel struct {
	value *TestSuiteChangeViewModelTestPlanChangedFieldViewModel
	isSet bool
}

func (v NullableTestSuiteChangeViewModelTestPlanChangedFieldViewModel) Get() *TestSuiteChangeViewModelTestPlanChangedFieldViewModel {
	return v.value
}

func (v *NullableTestSuiteChangeViewModelTestPlanChangedFieldViewModel) Set(val *TestSuiteChangeViewModelTestPlanChangedFieldViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestSuiteChangeViewModelTestPlanChangedFieldViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestSuiteChangeViewModelTestPlanChangedFieldViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestSuiteChangeViewModelTestPlanChangedFieldViewModel(val *TestSuiteChangeViewModelTestPlanChangedFieldViewModel) *NullableTestSuiteChangeViewModelTestPlanChangedFieldViewModel {
	return &NullableTestSuiteChangeViewModelTestPlanChangedFieldViewModel{value: val, isSet: true}
}

func (v NullableTestSuiteChangeViewModelTestPlanChangedFieldViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestSuiteChangeViewModelTestPlanChangedFieldViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

