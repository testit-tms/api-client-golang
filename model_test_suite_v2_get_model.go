/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
	"time"
)

// checks if the TestSuiteV2GetModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestSuiteV2GetModel{}

// TestSuiteV2GetModel struct for TestSuiteV2GetModel
type TestSuiteV2GetModel struct {
	// Unique ID of the test suite
	Id *string `json:"id,omitempty"`
	// Date of the last refresh of the test suite
	RefreshDate NullableTime `json:"refreshDate,omitempty"`
	// Unique ID of the parent test suite in hierarchy
	ParentId NullableString `json:"parentId,omitempty"`
	// Unique ID of test plan to which the test suite belongs
	TestPlanId string `json:"testPlanId"`
	// Name of the test suite
	Name string `json:"name"`
	Type NullableTestSuiteType `json:"type,omitempty"`
	// Indicates if the test suite retains section tree structure
	SaveStructure NullableBool `json:"saveStructure,omitempty"`
	// Indicates if scheduled auto refresh is enabled for the test suite
	AutoRefresh NullableBool `json:"autoRefresh,omitempty"`
}

// NewTestSuiteV2GetModel instantiates a new TestSuiteV2GetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestSuiteV2GetModel(testPlanId string, name string) *TestSuiteV2GetModel {
	this := TestSuiteV2GetModel{}
	this.TestPlanId = testPlanId
	this.Name = name
	return &this
}

// NewTestSuiteV2GetModelWithDefaults instantiates a new TestSuiteV2GetModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestSuiteV2GetModelWithDefaults() *TestSuiteV2GetModel {
	this := TestSuiteV2GetModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TestSuiteV2GetModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSuiteV2GetModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TestSuiteV2GetModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TestSuiteV2GetModel) SetId(v string) {
	o.Id = &v
}

// GetRefreshDate returns the RefreshDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestSuiteV2GetModel) GetRefreshDate() time.Time {
	if o == nil || IsNil(o.RefreshDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RefreshDate.Get()
}

// GetRefreshDateOk returns a tuple with the RefreshDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestSuiteV2GetModel) GetRefreshDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshDate.Get(), o.RefreshDate.IsSet()
}

// HasRefreshDate returns a boolean if a field has been set.
func (o *TestSuiteV2GetModel) HasRefreshDate() bool {
	if o != nil && o.RefreshDate.IsSet() {
		return true
	}

	return false
}

// SetRefreshDate gets a reference to the given NullableTime and assigns it to the RefreshDate field.
func (o *TestSuiteV2GetModel) SetRefreshDate(v time.Time) {
	o.RefreshDate.Set(&v)
}
// SetRefreshDateNil sets the value for RefreshDate to be an explicit nil
func (o *TestSuiteV2GetModel) SetRefreshDateNil() {
	o.RefreshDate.Set(nil)
}

// UnsetRefreshDate ensures that no value is present for RefreshDate, not even an explicit nil
func (o *TestSuiteV2GetModel) UnsetRefreshDate() {
	o.RefreshDate.Unset()
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestSuiteV2GetModel) GetParentId() string {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestSuiteV2GetModel) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *TestSuiteV2GetModel) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *TestSuiteV2GetModel) SetParentId(v string) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *TestSuiteV2GetModel) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *TestSuiteV2GetModel) UnsetParentId() {
	o.ParentId.Unset()
}

// GetTestPlanId returns the TestPlanId field value
func (o *TestSuiteV2GetModel) GetTestPlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestPlanId
}

// GetTestPlanIdOk returns a tuple with the TestPlanId field value
// and a boolean to check if the value has been set.
func (o *TestSuiteV2GetModel) GetTestPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestPlanId, true
}

// SetTestPlanId sets field value
func (o *TestSuiteV2GetModel) SetTestPlanId(v string) {
	o.TestPlanId = v
}

// GetName returns the Name field value
func (o *TestSuiteV2GetModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TestSuiteV2GetModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TestSuiteV2GetModel) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestSuiteV2GetModel) GetType() TestSuiteType {
	if o == nil || IsNil(o.Type.Get()) {
		var ret TestSuiteType
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestSuiteV2GetModel) GetTypeOk() (*TestSuiteType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *TestSuiteV2GetModel) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableTestSuiteType and assigns it to the Type field.
func (o *TestSuiteV2GetModel) SetType(v TestSuiteType) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *TestSuiteV2GetModel) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *TestSuiteV2GetModel) UnsetType() {
	o.Type.Unset()
}

// GetSaveStructure returns the SaveStructure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestSuiteV2GetModel) GetSaveStructure() bool {
	if o == nil || IsNil(o.SaveStructure.Get()) {
		var ret bool
		return ret
	}
	return *o.SaveStructure.Get()
}

// GetSaveStructureOk returns a tuple with the SaveStructure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestSuiteV2GetModel) GetSaveStructureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaveStructure.Get(), o.SaveStructure.IsSet()
}

// HasSaveStructure returns a boolean if a field has been set.
func (o *TestSuiteV2GetModel) HasSaveStructure() bool {
	if o != nil && o.SaveStructure.IsSet() {
		return true
	}

	return false
}

// SetSaveStructure gets a reference to the given NullableBool and assigns it to the SaveStructure field.
func (o *TestSuiteV2GetModel) SetSaveStructure(v bool) {
	o.SaveStructure.Set(&v)
}
// SetSaveStructureNil sets the value for SaveStructure to be an explicit nil
func (o *TestSuiteV2GetModel) SetSaveStructureNil() {
	o.SaveStructure.Set(nil)
}

// UnsetSaveStructure ensures that no value is present for SaveStructure, not even an explicit nil
func (o *TestSuiteV2GetModel) UnsetSaveStructure() {
	o.SaveStructure.Unset()
}

// GetAutoRefresh returns the AutoRefresh field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestSuiteV2GetModel) GetAutoRefresh() bool {
	if o == nil || IsNil(o.AutoRefresh.Get()) {
		var ret bool
		return ret
	}
	return *o.AutoRefresh.Get()
}

// GetAutoRefreshOk returns a tuple with the AutoRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestSuiteV2GetModel) GetAutoRefreshOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoRefresh.Get(), o.AutoRefresh.IsSet()
}

// HasAutoRefresh returns a boolean if a field has been set.
func (o *TestSuiteV2GetModel) HasAutoRefresh() bool {
	if o != nil && o.AutoRefresh.IsSet() {
		return true
	}

	return false
}

// SetAutoRefresh gets a reference to the given NullableBool and assigns it to the AutoRefresh field.
func (o *TestSuiteV2GetModel) SetAutoRefresh(v bool) {
	o.AutoRefresh.Set(&v)
}
// SetAutoRefreshNil sets the value for AutoRefresh to be an explicit nil
func (o *TestSuiteV2GetModel) SetAutoRefreshNil() {
	o.AutoRefresh.Set(nil)
}

// UnsetAutoRefresh ensures that no value is present for AutoRefresh, not even an explicit nil
func (o *TestSuiteV2GetModel) UnsetAutoRefresh() {
	o.AutoRefresh.Unset()
}

func (o TestSuiteV2GetModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestSuiteV2GetModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.RefreshDate.IsSet() {
		toSerialize["refreshDate"] = o.RefreshDate.Get()
	}
	if o.ParentId.IsSet() {
		toSerialize["parentId"] = o.ParentId.Get()
	}
	toSerialize["testPlanId"] = o.TestPlanId
	toSerialize["name"] = o.Name
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.SaveStructure.IsSet() {
		toSerialize["saveStructure"] = o.SaveStructure.Get()
	}
	if o.AutoRefresh.IsSet() {
		toSerialize["autoRefresh"] = o.AutoRefresh.Get()
	}
	return toSerialize, nil
}

type NullableTestSuiteV2GetModel struct {
	value *TestSuiteV2GetModel
	isSet bool
}

func (v NullableTestSuiteV2GetModel) Get() *TestSuiteV2GetModel {
	return v.value
}

func (v *NullableTestSuiteV2GetModel) Set(val *TestSuiteV2GetModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestSuiteV2GetModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestSuiteV2GetModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestSuiteV2GetModel(val *TestSuiteV2GetModel) *NullableTestSuiteV2GetModel {
	return &NullableTestSuiteV2GetModel{value: val, isSet: true}
}

func (v NullableTestSuiteV2GetModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestSuiteV2GetModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


