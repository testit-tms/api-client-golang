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

// checks if the WorkItemStepChangeViewModelWorkItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemStepChangeViewModelWorkItem{}

// WorkItemStepChangeViewModelWorkItem struct for WorkItemStepChangeViewModelWorkItem
type WorkItemStepChangeViewModelWorkItem struct {
	Id string `json:"id"`
	GlobalId int64 `json:"globalId"`
	Name string `json:"name"`
	Steps []WorkItemStepChangeViewModel `json:"steps"`
}

// NewWorkItemStepChangeViewModelWorkItem instantiates a new WorkItemStepChangeViewModelWorkItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemStepChangeViewModelWorkItem(id string, globalId int64, name string, steps []WorkItemStepChangeViewModel) *WorkItemStepChangeViewModelWorkItem {
	this := WorkItemStepChangeViewModelWorkItem{}
	this.Id = id
	this.GlobalId = globalId
	this.Name = name
	this.Steps = steps
	return &this
}

// NewWorkItemStepChangeViewModelWorkItemWithDefaults instantiates a new WorkItemStepChangeViewModelWorkItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemStepChangeViewModelWorkItemWithDefaults() *WorkItemStepChangeViewModelWorkItem {
	this := WorkItemStepChangeViewModelWorkItem{}
	return &this
}

// GetId returns the Id field value
func (o *WorkItemStepChangeViewModelWorkItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkItemStepChangeViewModelWorkItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkItemStepChangeViewModelWorkItem) SetId(v string) {
	o.Id = v
}

// GetGlobalId returns the GlobalId field value
func (o *WorkItemStepChangeViewModelWorkItem) GetGlobalId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value
// and a boolean to check if the value has been set.
func (o *WorkItemStepChangeViewModelWorkItem) GetGlobalIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalId, true
}

// SetGlobalId sets field value
func (o *WorkItemStepChangeViewModelWorkItem) SetGlobalId(v int64) {
	o.GlobalId = v
}

// GetName returns the Name field value
func (o *WorkItemStepChangeViewModelWorkItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkItemStepChangeViewModelWorkItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkItemStepChangeViewModelWorkItem) SetName(v string) {
	o.Name = v
}

// GetSteps returns the Steps field value
func (o *WorkItemStepChangeViewModelWorkItem) GetSteps() []WorkItemStepChangeViewModel {
	if o == nil {
		var ret []WorkItemStepChangeViewModel
		return ret
	}

	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *WorkItemStepChangeViewModelWorkItem) GetStepsOk() ([]WorkItemStepChangeViewModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Steps, true
}

// SetSteps sets field value
func (o *WorkItemStepChangeViewModelWorkItem) SetSteps(v []WorkItemStepChangeViewModel) {
	o.Steps = v
}

func (o WorkItemStepChangeViewModelWorkItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemStepChangeViewModelWorkItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["globalId"] = o.GlobalId
	toSerialize["name"] = o.Name
	toSerialize["steps"] = o.Steps
	return toSerialize, nil
}

type NullableWorkItemStepChangeViewModelWorkItem struct {
	value *WorkItemStepChangeViewModelWorkItem
	isSet bool
}

func (v NullableWorkItemStepChangeViewModelWorkItem) Get() *WorkItemStepChangeViewModelWorkItem {
	return v.value
}

func (v *NullableWorkItemStepChangeViewModelWorkItem) Set(val *WorkItemStepChangeViewModelWorkItem) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemStepChangeViewModelWorkItem) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemStepChangeViewModelWorkItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemStepChangeViewModelWorkItem(val *WorkItemStepChangeViewModelWorkItem) *NullableWorkItemStepChangeViewModelWorkItem {
	return &NullableWorkItemStepChangeViewModelWorkItem{value: val, isSet: true}
}

func (v NullableWorkItemStepChangeViewModelWorkItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemStepChangeViewModelWorkItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


