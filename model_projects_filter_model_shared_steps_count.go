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

// checks if the ProjectsFilterModelSharedStepsCount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectsFilterModelSharedStepsCount{}

// ProjectsFilterModelSharedStepsCount Specifies a project range of shared steps count to search for
type ProjectsFilterModelSharedStepsCount struct {
	From NullableInt32 `json:"from,omitempty"`
	To NullableInt32 `json:"to,omitempty"`
}

// NewProjectsFilterModelSharedStepsCount instantiates a new ProjectsFilterModelSharedStepsCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsFilterModelSharedStepsCount() *ProjectsFilterModelSharedStepsCount {
	this := ProjectsFilterModelSharedStepsCount{}
	return &this
}

// NewProjectsFilterModelSharedStepsCountWithDefaults instantiates a new ProjectsFilterModelSharedStepsCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsFilterModelSharedStepsCountWithDefaults() *ProjectsFilterModelSharedStepsCount {
	this := ProjectsFilterModelSharedStepsCount{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsFilterModelSharedStepsCount) GetFrom() int32 {
	if o == nil || IsNil(o.From.Get()) {
		var ret int32
		return ret
	}
	return *o.From.Get()
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsFilterModelSharedStepsCount) GetFromOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.From.Get(), o.From.IsSet()
}

// HasFrom returns a boolean if a field has been set.
func (o *ProjectsFilterModelSharedStepsCount) HasFrom() bool {
	if o != nil && o.From.IsSet() {
		return true
	}

	return false
}

// SetFrom gets a reference to the given NullableInt32 and assigns it to the From field.
func (o *ProjectsFilterModelSharedStepsCount) SetFrom(v int32) {
	o.From.Set(&v)
}
// SetFromNil sets the value for From to be an explicit nil
func (o *ProjectsFilterModelSharedStepsCount) SetFromNil() {
	o.From.Set(nil)
}

// UnsetFrom ensures that no value is present for From, not even an explicit nil
func (o *ProjectsFilterModelSharedStepsCount) UnsetFrom() {
	o.From.Unset()
}

// GetTo returns the To field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsFilterModelSharedStepsCount) GetTo() int32 {
	if o == nil || IsNil(o.To.Get()) {
		var ret int32
		return ret
	}
	return *o.To.Get()
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsFilterModelSharedStepsCount) GetToOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.To.Get(), o.To.IsSet()
}

// HasTo returns a boolean if a field has been set.
func (o *ProjectsFilterModelSharedStepsCount) HasTo() bool {
	if o != nil && o.To.IsSet() {
		return true
	}

	return false
}

// SetTo gets a reference to the given NullableInt32 and assigns it to the To field.
func (o *ProjectsFilterModelSharedStepsCount) SetTo(v int32) {
	o.To.Set(&v)
}
// SetToNil sets the value for To to be an explicit nil
func (o *ProjectsFilterModelSharedStepsCount) SetToNil() {
	o.To.Set(nil)
}

// UnsetTo ensures that no value is present for To, not even an explicit nil
func (o *ProjectsFilterModelSharedStepsCount) UnsetTo() {
	o.To.Unset()
}

func (o ProjectsFilterModelSharedStepsCount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectsFilterModelSharedStepsCount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.From.IsSet() {
		toSerialize["from"] = o.From.Get()
	}
	if o.To.IsSet() {
		toSerialize["to"] = o.To.Get()
	}
	return toSerialize, nil
}

type NullableProjectsFilterModelSharedStepsCount struct {
	value *ProjectsFilterModelSharedStepsCount
	isSet bool
}

func (v NullableProjectsFilterModelSharedStepsCount) Get() *ProjectsFilterModelSharedStepsCount {
	return v.value
}

func (v *NullableProjectsFilterModelSharedStepsCount) Set(val *ProjectsFilterModelSharedStepsCount) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsFilterModelSharedStepsCount) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsFilterModelSharedStepsCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsFilterModelSharedStepsCount(val *ProjectsFilterModelSharedStepsCount) *NullableProjectsFilterModelSharedStepsCount {
	return &NullableProjectsFilterModelSharedStepsCount{value: val, isSet: true}
}

func (v NullableProjectsFilterModelSharedStepsCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsFilterModelSharedStepsCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

