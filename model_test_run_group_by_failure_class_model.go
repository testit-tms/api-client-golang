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

// checks if the TestRunGroupByFailureClassModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestRunGroupByFailureClassModel{}

// TestRunGroupByFailureClassModel struct for TestRunGroupByFailureClassModel
type TestRunGroupByFailureClassModel struct {
	FailureCategory NullableString `json:"failureCategory,omitempty"`
	Value *int32 `json:"value,omitempty"`
}

// NewTestRunGroupByFailureClassModel instantiates a new TestRunGroupByFailureClassModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestRunGroupByFailureClassModel() *TestRunGroupByFailureClassModel {
	this := TestRunGroupByFailureClassModel{}
	return &this
}

// NewTestRunGroupByFailureClassModelWithDefaults instantiates a new TestRunGroupByFailureClassModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestRunGroupByFailureClassModelWithDefaults() *TestRunGroupByFailureClassModel {
	this := TestRunGroupByFailureClassModel{}
	return &this
}

// GetFailureCategory returns the FailureCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunGroupByFailureClassModel) GetFailureCategory() string {
	if o == nil || IsNil(o.FailureCategory.Get()) {
		var ret string
		return ret
	}
	return *o.FailureCategory.Get()
}

// GetFailureCategoryOk returns a tuple with the FailureCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunGroupByFailureClassModel) GetFailureCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureCategory.Get(), o.FailureCategory.IsSet()
}

// HasFailureCategory returns a boolean if a field has been set.
func (o *TestRunGroupByFailureClassModel) HasFailureCategory() bool {
	if o != nil && o.FailureCategory.IsSet() {
		return true
	}

	return false
}

// SetFailureCategory gets a reference to the given NullableString and assigns it to the FailureCategory field.
func (o *TestRunGroupByFailureClassModel) SetFailureCategory(v string) {
	o.FailureCategory.Set(&v)
}
// SetFailureCategoryNil sets the value for FailureCategory to be an explicit nil
func (o *TestRunGroupByFailureClassModel) SetFailureCategoryNil() {
	o.FailureCategory.Set(nil)
}

// UnsetFailureCategory ensures that no value is present for FailureCategory, not even an explicit nil
func (o *TestRunGroupByFailureClassModel) UnsetFailureCategory() {
	o.FailureCategory.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TestRunGroupByFailureClassModel) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunGroupByFailureClassModel) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TestRunGroupByFailureClassModel) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *TestRunGroupByFailureClassModel) SetValue(v int32) {
	o.Value = &v
}

func (o TestRunGroupByFailureClassModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestRunGroupByFailureClassModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FailureCategory.IsSet() {
		toSerialize["failureCategory"] = o.FailureCategory.Get()
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTestRunGroupByFailureClassModel struct {
	value *TestRunGroupByFailureClassModel
	isSet bool
}

func (v NullableTestRunGroupByFailureClassModel) Get() *TestRunGroupByFailureClassModel {
	return v.value
}

func (v *NullableTestRunGroupByFailureClassModel) Set(val *TestRunGroupByFailureClassModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestRunGroupByFailureClassModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestRunGroupByFailureClassModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestRunGroupByFailureClassModel(val *TestRunGroupByFailureClassModel) *NullableTestRunGroupByFailureClassModel {
	return &NullableTestRunGroupByFailureClassModel{value: val, isSet: true}
}

func (v NullableTestRunGroupByFailureClassModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestRunGroupByFailureClassModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


