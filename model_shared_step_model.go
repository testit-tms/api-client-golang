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

// checks if the SharedStepModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharedStepModel{}

// SharedStepModel struct for SharedStepModel
type SharedStepModel struct {
	VersionId string `json:"versionId"`
	GlobalId int64 `json:"globalId"`
	Name string `json:"name"`
	// Deprecated
	Steps []StepModel `json:"steps"`
	IsDeleted bool `json:"isDeleted"`
}

type _SharedStepModel SharedStepModel

// NewSharedStepModel instantiates a new SharedStepModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedStepModel(versionId string, globalId int64, name string, steps []StepModel, isDeleted bool) *SharedStepModel {
	this := SharedStepModel{}
	this.VersionId = versionId
	this.GlobalId = globalId
	this.Name = name
	this.Steps = steps
	this.IsDeleted = isDeleted
	return &this
}

// NewSharedStepModelWithDefaults instantiates a new SharedStepModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedStepModelWithDefaults() *SharedStepModel {
	this := SharedStepModel{}
	return &this
}

// GetVersionId returns the VersionId field value
func (o *SharedStepModel) GetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *SharedStepModel) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *SharedStepModel) SetVersionId(v string) {
	o.VersionId = v
}

// GetGlobalId returns the GlobalId field value
func (o *SharedStepModel) GetGlobalId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value
// and a boolean to check if the value has been set.
func (o *SharedStepModel) GetGlobalIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalId, true
}

// SetGlobalId sets field value
func (o *SharedStepModel) SetGlobalId(v int64) {
	o.GlobalId = v
}

// GetName returns the Name field value
func (o *SharedStepModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SharedStepModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SharedStepModel) SetName(v string) {
	o.Name = v
}

// GetSteps returns the Steps field value
// Deprecated
func (o *SharedStepModel) GetSteps() []StepModel {
	if o == nil {
		var ret []StepModel
		return ret
	}

	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *SharedStepModel) GetStepsOk() ([]StepModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Steps, true
}

// SetSteps sets field value
// Deprecated
func (o *SharedStepModel) SetSteps(v []StepModel) {
	o.Steps = v
}

// GetIsDeleted returns the IsDeleted field value
func (o *SharedStepModel) GetIsDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value
// and a boolean to check if the value has been set.
func (o *SharedStepModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeleted, true
}

// SetIsDeleted sets field value
func (o *SharedStepModel) SetIsDeleted(v bool) {
	o.IsDeleted = v
}

func (o SharedStepModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharedStepModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["versionId"] = o.VersionId
	toSerialize["globalId"] = o.GlobalId
	toSerialize["name"] = o.Name
	toSerialize["steps"] = o.Steps
	toSerialize["isDeleted"] = o.IsDeleted
	return toSerialize, nil
}

func (o *SharedStepModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"versionId",
		"globalId",
		"name",
		"steps",
		"isDeleted",
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

	varSharedStepModel := _SharedStepModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSharedStepModel)

	if err != nil {
		return err
	}

	*o = SharedStepModel(varSharedStepModel)

	return err
}

type NullableSharedStepModel struct {
	value *SharedStepModel
	isSet bool
}

func (v NullableSharedStepModel) Get() *SharedStepModel {
	return v.value
}

func (v *NullableSharedStepModel) Set(val *SharedStepModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedStepModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedStepModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedStepModel(val *SharedStepModel) *NullableSharedStepModel {
	return &NullableSharedStepModel{value: val, isSet: true}
}

func (v NullableSharedStepModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedStepModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


