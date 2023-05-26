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

// checks if the TestSuiteV2PostModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestSuiteV2PostModel{}

// TestSuiteV2PostModel struct for TestSuiteV2PostModel
type TestSuiteV2PostModel struct {
	ParentId NullableString `json:"parentId,omitempty"`
	TestPlanId string `json:"testPlanId"`
	Name string `json:"name"`
	Type *TestSuiteType `json:"type,omitempty"`
	SaveStructure NullableBool `json:"saveStructure,omitempty"`
}

// NewTestSuiteV2PostModel instantiates a new TestSuiteV2PostModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestSuiteV2PostModel(testPlanId string, name string) *TestSuiteV2PostModel {
	this := TestSuiteV2PostModel{}
	this.TestPlanId = testPlanId
	this.Name = name
	return &this
}

// NewTestSuiteV2PostModelWithDefaults instantiates a new TestSuiteV2PostModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestSuiteV2PostModelWithDefaults() *TestSuiteV2PostModel {
	this := TestSuiteV2PostModel{}
	return &this
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestSuiteV2PostModel) GetParentId() string {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestSuiteV2PostModel) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *TestSuiteV2PostModel) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *TestSuiteV2PostModel) SetParentId(v string) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *TestSuiteV2PostModel) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *TestSuiteV2PostModel) UnsetParentId() {
	o.ParentId.Unset()
}

// GetTestPlanId returns the TestPlanId field value
func (o *TestSuiteV2PostModel) GetTestPlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestPlanId
}

// GetTestPlanIdOk returns a tuple with the TestPlanId field value
// and a boolean to check if the value has been set.
func (o *TestSuiteV2PostModel) GetTestPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestPlanId, true
}

// SetTestPlanId sets field value
func (o *TestSuiteV2PostModel) SetTestPlanId(v string) {
	o.TestPlanId = v
}

// GetName returns the Name field value
func (o *TestSuiteV2PostModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TestSuiteV2PostModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TestSuiteV2PostModel) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TestSuiteV2PostModel) GetType() TestSuiteType {
	if o == nil || IsNil(o.Type) {
		var ret TestSuiteType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSuiteV2PostModel) GetTypeOk() (*TestSuiteType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TestSuiteV2PostModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TestSuiteType and assigns it to the Type field.
func (o *TestSuiteV2PostModel) SetType(v TestSuiteType) {
	o.Type = &v
}

// GetSaveStructure returns the SaveStructure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestSuiteV2PostModel) GetSaveStructure() bool {
	if o == nil || IsNil(o.SaveStructure.Get()) {
		var ret bool
		return ret
	}
	return *o.SaveStructure.Get()
}

// GetSaveStructureOk returns a tuple with the SaveStructure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestSuiteV2PostModel) GetSaveStructureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaveStructure.Get(), o.SaveStructure.IsSet()
}

// HasSaveStructure returns a boolean if a field has been set.
func (o *TestSuiteV2PostModel) HasSaveStructure() bool {
	if o != nil && o.SaveStructure.IsSet() {
		return true
	}

	return false
}

// SetSaveStructure gets a reference to the given NullableBool and assigns it to the SaveStructure field.
func (o *TestSuiteV2PostModel) SetSaveStructure(v bool) {
	o.SaveStructure.Set(&v)
}
// SetSaveStructureNil sets the value for SaveStructure to be an explicit nil
func (o *TestSuiteV2PostModel) SetSaveStructureNil() {
	o.SaveStructure.Set(nil)
}

// UnsetSaveStructure ensures that no value is present for SaveStructure, not even an explicit nil
func (o *TestSuiteV2PostModel) UnsetSaveStructure() {
	o.SaveStructure.Unset()
}

func (o TestSuiteV2PostModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestSuiteV2PostModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ParentId.IsSet() {
		toSerialize["parentId"] = o.ParentId.Get()
	}
	toSerialize["testPlanId"] = o.TestPlanId
	toSerialize["name"] = o.Name
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.SaveStructure.IsSet() {
		toSerialize["saveStructure"] = o.SaveStructure.Get()
	}
	return toSerialize, nil
}

type NullableTestSuiteV2PostModel struct {
	value *TestSuiteV2PostModel
	isSet bool
}

func (v NullableTestSuiteV2PostModel) Get() *TestSuiteV2PostModel {
	return v.value
}

func (v *NullableTestSuiteV2PostModel) Set(val *TestSuiteV2PostModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestSuiteV2PostModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestSuiteV2PostModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestSuiteV2PostModel(val *TestSuiteV2PostModel) *NullableTestSuiteV2PostModel {
	return &NullableTestSuiteV2PostModel{value: val, isSet: true}
}

func (v NullableTestSuiteV2PostModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestSuiteV2PostModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

