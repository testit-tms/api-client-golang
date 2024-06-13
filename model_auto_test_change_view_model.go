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

// checks if the AutoTestChangeViewModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoTestChangeViewModel{}

// AutoTestChangeViewModel struct for AutoTestChangeViewModel
type AutoTestChangeViewModel struct {
	Id string `json:"id"`
	ProjectId string `json:"projectId"`
	ExternalId string `json:"externalId"`
	GlobalId int64 `json:"globalId"`
}

type _AutoTestChangeViewModel AutoTestChangeViewModel

// NewAutoTestChangeViewModel instantiates a new AutoTestChangeViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoTestChangeViewModel(id string, projectId string, externalId string, globalId int64) *AutoTestChangeViewModel {
	this := AutoTestChangeViewModel{}
	this.Id = id
	this.ProjectId = projectId
	this.ExternalId = externalId
	this.GlobalId = globalId
	return &this
}

// NewAutoTestChangeViewModelWithDefaults instantiates a new AutoTestChangeViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoTestChangeViewModelWithDefaults() *AutoTestChangeViewModel {
	this := AutoTestChangeViewModel{}
	return &this
}

// GetId returns the Id field value
func (o *AutoTestChangeViewModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AutoTestChangeViewModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AutoTestChangeViewModel) SetId(v string) {
	o.Id = v
}

// GetProjectId returns the ProjectId field value
func (o *AutoTestChangeViewModel) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *AutoTestChangeViewModel) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *AutoTestChangeViewModel) SetProjectId(v string) {
	o.ProjectId = v
}

// GetExternalId returns the ExternalId field value
func (o *AutoTestChangeViewModel) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *AutoTestChangeViewModel) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *AutoTestChangeViewModel) SetExternalId(v string) {
	o.ExternalId = v
}

// GetGlobalId returns the GlobalId field value
func (o *AutoTestChangeViewModel) GetGlobalId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value
// and a boolean to check if the value has been set.
func (o *AutoTestChangeViewModel) GetGlobalIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalId, true
}

// SetGlobalId sets field value
func (o *AutoTestChangeViewModel) SetGlobalId(v int64) {
	o.GlobalId = v
}

func (o AutoTestChangeViewModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoTestChangeViewModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["projectId"] = o.ProjectId
	toSerialize["externalId"] = o.ExternalId
	toSerialize["globalId"] = o.GlobalId
	return toSerialize, nil
}

func (o *AutoTestChangeViewModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"projectId",
		"externalId",
		"globalId",
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

	varAutoTestChangeViewModel := _AutoTestChangeViewModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAutoTestChangeViewModel)

	if err != nil {
		return err
	}

	*o = AutoTestChangeViewModel(varAutoTestChangeViewModel)

	return err
}

type NullableAutoTestChangeViewModel struct {
	value *AutoTestChangeViewModel
	isSet bool
}

func (v NullableAutoTestChangeViewModel) Get() *AutoTestChangeViewModel {
	return v.value
}

func (v *NullableAutoTestChangeViewModel) Set(val *AutoTestChangeViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoTestChangeViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoTestChangeViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoTestChangeViewModel(val *AutoTestChangeViewModel) *NullableAutoTestChangeViewModel {
	return &NullableAutoTestChangeViewModel{value: val, isSet: true}
}

func (v NullableAutoTestChangeViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoTestChangeViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


