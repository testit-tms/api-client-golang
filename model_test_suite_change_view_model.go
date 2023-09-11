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

// checks if the TestSuiteChangeViewModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestSuiteChangeViewModel{}

// TestSuiteChangeViewModel struct for TestSuiteChangeViewModel
type TestSuiteChangeViewModel struct {
	Id string `json:"id"`
	Name NullableString `json:"name,omitempty"`
	Configurations []ShortConfiguration `json:"configurations,omitempty"`
	WorkItemCount int64 `json:"workItemCount"`
}

// NewTestSuiteChangeViewModel instantiates a new TestSuiteChangeViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestSuiteChangeViewModel(id string, workItemCount int64) *TestSuiteChangeViewModel {
	this := TestSuiteChangeViewModel{}
	this.Id = id
	this.WorkItemCount = workItemCount
	return &this
}

// NewTestSuiteChangeViewModelWithDefaults instantiates a new TestSuiteChangeViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestSuiteChangeViewModelWithDefaults() *TestSuiteChangeViewModel {
	this := TestSuiteChangeViewModel{}
	return &this
}

// GetId returns the Id field value
func (o *TestSuiteChangeViewModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestSuiteChangeViewModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestSuiteChangeViewModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestSuiteChangeViewModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestSuiteChangeViewModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TestSuiteChangeViewModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TestSuiteChangeViewModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TestSuiteChangeViewModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TestSuiteChangeViewModel) UnsetName() {
	o.Name.Unset()
}

// GetConfigurations returns the Configurations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestSuiteChangeViewModel) GetConfigurations() []ShortConfiguration {
	if o == nil {
		var ret []ShortConfiguration
		return ret
	}
	return o.Configurations
}

// GetConfigurationsOk returns a tuple with the Configurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestSuiteChangeViewModel) GetConfigurationsOk() ([]ShortConfiguration, bool) {
	if o == nil || IsNil(o.Configurations) {
		return nil, false
	}
	return o.Configurations, true
}

// HasConfigurations returns a boolean if a field has been set.
func (o *TestSuiteChangeViewModel) HasConfigurations() bool {
	if o != nil && IsNil(o.Configurations) {
		return true
	}

	return false
}

// SetConfigurations gets a reference to the given []ShortConfiguration and assigns it to the Configurations field.
func (o *TestSuiteChangeViewModel) SetConfigurations(v []ShortConfiguration) {
	o.Configurations = v
}

// GetWorkItemCount returns the WorkItemCount field value
func (o *TestSuiteChangeViewModel) GetWorkItemCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.WorkItemCount
}

// GetWorkItemCountOk returns a tuple with the WorkItemCount field value
// and a boolean to check if the value has been set.
func (o *TestSuiteChangeViewModel) GetWorkItemCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkItemCount, true
}

// SetWorkItemCount sets field value
func (o *TestSuiteChangeViewModel) SetWorkItemCount(v int64) {
	o.WorkItemCount = v
}

func (o TestSuiteChangeViewModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestSuiteChangeViewModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Configurations != nil {
		toSerialize["configurations"] = o.Configurations
	}
	toSerialize["workItemCount"] = o.WorkItemCount
	return toSerialize, nil
}

type NullableTestSuiteChangeViewModel struct {
	value *TestSuiteChangeViewModel
	isSet bool
}

func (v NullableTestSuiteChangeViewModel) Get() *TestSuiteChangeViewModel {
	return v.value
}

func (v *NullableTestSuiteChangeViewModel) Set(val *TestSuiteChangeViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestSuiteChangeViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestSuiteChangeViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestSuiteChangeViewModel(val *TestSuiteChangeViewModel) *NullableTestSuiteChangeViewModel {
	return &NullableTestSuiteChangeViewModel{value: val, isSet: true}
}

func (v NullableTestSuiteChangeViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestSuiteChangeViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


