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

// checks if the AutoTestShortModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoTestShortModel{}

// AutoTestShortModel struct for AutoTestShortModel
type AutoTestShortModel struct {
	Id string `json:"id"`
	GlobalId int64 `json:"globalId"`
	ExternalId string `json:"externalId"`
	ProjectId string `json:"projectId"`
	Name string `json:"name"`
}

// NewAutoTestShortModel instantiates a new AutoTestShortModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoTestShortModel(id string, globalId int64, externalId string, projectId string, name string) *AutoTestShortModel {
	this := AutoTestShortModel{}
	this.Id = id
	this.GlobalId = globalId
	this.ExternalId = externalId
	this.ProjectId = projectId
	this.Name = name
	return &this
}

// NewAutoTestShortModelWithDefaults instantiates a new AutoTestShortModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoTestShortModelWithDefaults() *AutoTestShortModel {
	this := AutoTestShortModel{}
	return &this
}

// GetId returns the Id field value
func (o *AutoTestShortModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AutoTestShortModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AutoTestShortModel) SetId(v string) {
	o.Id = v
}

// GetGlobalId returns the GlobalId field value
func (o *AutoTestShortModel) GetGlobalId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value
// and a boolean to check if the value has been set.
func (o *AutoTestShortModel) GetGlobalIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalId, true
}

// SetGlobalId sets field value
func (o *AutoTestShortModel) SetGlobalId(v int64) {
	o.GlobalId = v
}

// GetExternalId returns the ExternalId field value
func (o *AutoTestShortModel) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *AutoTestShortModel) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *AutoTestShortModel) SetExternalId(v string) {
	o.ExternalId = v
}

// GetProjectId returns the ProjectId field value
func (o *AutoTestShortModel) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *AutoTestShortModel) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *AutoTestShortModel) SetProjectId(v string) {
	o.ProjectId = v
}

// GetName returns the Name field value
func (o *AutoTestShortModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AutoTestShortModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AutoTestShortModel) SetName(v string) {
	o.Name = v
}

func (o AutoTestShortModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoTestShortModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["globalId"] = o.GlobalId
	toSerialize["externalId"] = o.ExternalId
	toSerialize["projectId"] = o.ProjectId
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableAutoTestShortModel struct {
	value *AutoTestShortModel
	isSet bool
}

func (v NullableAutoTestShortModel) Get() *AutoTestShortModel {
	return v.value
}

func (v *NullableAutoTestShortModel) Set(val *AutoTestShortModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoTestShortModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoTestShortModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoTestShortModel(val *AutoTestShortModel) *NullableAutoTestShortModel {
	return &NullableAutoTestShortModel{value: val, isSet: true}
}

func (v NullableAutoTestShortModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoTestShortModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


