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

// checks if the SharedStepChangeViewModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharedStepChangeViewModel{}

// SharedStepChangeViewModel struct for SharedStepChangeViewModel
type SharedStepChangeViewModel struct {
	Id *string `json:"id,omitempty"`
	GlobalId *int64 `json:"globalId,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Steps []WorkItemStepChangeViewModel `json:"steps,omitempty"`
}

// NewSharedStepChangeViewModel instantiates a new SharedStepChangeViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedStepChangeViewModel() *SharedStepChangeViewModel {
	this := SharedStepChangeViewModel{}
	return &this
}

// NewSharedStepChangeViewModelWithDefaults instantiates a new SharedStepChangeViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedStepChangeViewModelWithDefaults() *SharedStepChangeViewModel {
	this := SharedStepChangeViewModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SharedStepChangeViewModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedStepChangeViewModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SharedStepChangeViewModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SharedStepChangeViewModel) SetId(v string) {
	o.Id = &v
}

// GetGlobalId returns the GlobalId field value if set, zero value otherwise.
func (o *SharedStepChangeViewModel) GetGlobalId() int64 {
	if o == nil || IsNil(o.GlobalId) {
		var ret int64
		return ret
	}
	return *o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedStepChangeViewModel) GetGlobalIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GlobalId) {
		return nil, false
	}
	return o.GlobalId, true
}

// HasGlobalId returns a boolean if a field has been set.
func (o *SharedStepChangeViewModel) HasGlobalId() bool {
	if o != nil && !IsNil(o.GlobalId) {
		return true
	}

	return false
}

// SetGlobalId gets a reference to the given int64 and assigns it to the GlobalId field.
func (o *SharedStepChangeViewModel) SetGlobalId(v int64) {
	o.GlobalId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedStepChangeViewModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SharedStepChangeViewModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SharedStepChangeViewModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SharedStepChangeViewModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SharedStepChangeViewModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SharedStepChangeViewModel) UnsetName() {
	o.Name.Unset()
}

// GetSteps returns the Steps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedStepChangeViewModel) GetSteps() []WorkItemStepChangeViewModel {
	if o == nil {
		var ret []WorkItemStepChangeViewModel
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SharedStepChangeViewModel) GetStepsOk() ([]WorkItemStepChangeViewModel, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *SharedStepChangeViewModel) HasSteps() bool {
	if o != nil && IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []WorkItemStepChangeViewModel and assigns it to the Steps field.
func (o *SharedStepChangeViewModel) SetSteps(v []WorkItemStepChangeViewModel) {
	o.Steps = v
}

func (o SharedStepChangeViewModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharedStepChangeViewModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.GlobalId) {
		toSerialize["globalId"] = o.GlobalId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	return toSerialize, nil
}

type NullableSharedStepChangeViewModel struct {
	value *SharedStepChangeViewModel
	isSet bool
}

func (v NullableSharedStepChangeViewModel) Get() *SharedStepChangeViewModel {
	return v.value
}

func (v *NullableSharedStepChangeViewModel) Set(val *SharedStepChangeViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedStepChangeViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedStepChangeViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedStepChangeViewModel(val *SharedStepChangeViewModel) *NullableSharedStepChangeViewModel {
	return &NullableSharedStepChangeViewModel{value: val, isSet: true}
}

func (v NullableSharedStepChangeViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedStepChangeViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


