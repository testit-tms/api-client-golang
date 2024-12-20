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

// checks if the TestRunAnalyticResultModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestRunAnalyticResultModel{}

// TestRunAnalyticResultModel struct for TestRunAnalyticResultModel
type TestRunAnalyticResultModel struct {
	CountGroupByStatus []TestRunGroupByStatusModel `json:"countGroupByStatus,omitempty"`
	CountGroupByFailureClass []TestRunGroupByFailureClassModel `json:"countGroupByFailureClass,omitempty"`
}

// NewTestRunAnalyticResultModel instantiates a new TestRunAnalyticResultModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestRunAnalyticResultModel() *TestRunAnalyticResultModel {
	this := TestRunAnalyticResultModel{}
	return &this
}

// NewTestRunAnalyticResultModelWithDefaults instantiates a new TestRunAnalyticResultModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestRunAnalyticResultModelWithDefaults() *TestRunAnalyticResultModel {
	this := TestRunAnalyticResultModel{}
	return &this
}

// GetCountGroupByStatus returns the CountGroupByStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunAnalyticResultModel) GetCountGroupByStatus() []TestRunGroupByStatusModel {
	if o == nil {
		var ret []TestRunGroupByStatusModel
		return ret
	}
	return o.CountGroupByStatus
}

// GetCountGroupByStatusOk returns a tuple with the CountGroupByStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunAnalyticResultModel) GetCountGroupByStatusOk() ([]TestRunGroupByStatusModel, bool) {
	if o == nil || IsNil(o.CountGroupByStatus) {
		return nil, false
	}
	return o.CountGroupByStatus, true
}

// HasCountGroupByStatus returns a boolean if a field has been set.
func (o *TestRunAnalyticResultModel) HasCountGroupByStatus() bool {
	if o != nil && !IsNil(o.CountGroupByStatus) {
		return true
	}

	return false
}

// SetCountGroupByStatus gets a reference to the given []TestRunGroupByStatusModel and assigns it to the CountGroupByStatus field.
func (o *TestRunAnalyticResultModel) SetCountGroupByStatus(v []TestRunGroupByStatusModel) {
	o.CountGroupByStatus = v
}

// GetCountGroupByFailureClass returns the CountGroupByFailureClass field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunAnalyticResultModel) GetCountGroupByFailureClass() []TestRunGroupByFailureClassModel {
	if o == nil {
		var ret []TestRunGroupByFailureClassModel
		return ret
	}
	return o.CountGroupByFailureClass
}

// GetCountGroupByFailureClassOk returns a tuple with the CountGroupByFailureClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunAnalyticResultModel) GetCountGroupByFailureClassOk() ([]TestRunGroupByFailureClassModel, bool) {
	if o == nil || IsNil(o.CountGroupByFailureClass) {
		return nil, false
	}
	return o.CountGroupByFailureClass, true
}

// HasCountGroupByFailureClass returns a boolean if a field has been set.
func (o *TestRunAnalyticResultModel) HasCountGroupByFailureClass() bool {
	if o != nil && !IsNil(o.CountGroupByFailureClass) {
		return true
	}

	return false
}

// SetCountGroupByFailureClass gets a reference to the given []TestRunGroupByFailureClassModel and assigns it to the CountGroupByFailureClass field.
func (o *TestRunAnalyticResultModel) SetCountGroupByFailureClass(v []TestRunGroupByFailureClassModel) {
	o.CountGroupByFailureClass = v
}

func (o TestRunAnalyticResultModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestRunAnalyticResultModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CountGroupByStatus != nil {
		toSerialize["countGroupByStatus"] = o.CountGroupByStatus
	}
	if o.CountGroupByFailureClass != nil {
		toSerialize["countGroupByFailureClass"] = o.CountGroupByFailureClass
	}
	return toSerialize, nil
}

type NullableTestRunAnalyticResultModel struct {
	value *TestRunAnalyticResultModel
	isSet bool
}

func (v NullableTestRunAnalyticResultModel) Get() *TestRunAnalyticResultModel {
	return v.value
}

func (v *NullableTestRunAnalyticResultModel) Set(val *TestRunAnalyticResultModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestRunAnalyticResultModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestRunAnalyticResultModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestRunAnalyticResultModel(val *TestRunAnalyticResultModel) *NullableTestRunAnalyticResultModel {
	return &NullableTestRunAnalyticResultModel{value: val, isSet: true}
}

func (v NullableTestRunAnalyticResultModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestRunAnalyticResultModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


