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

// checks if the WorkItemIdentifierModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemIdentifierModel{}

// WorkItemIdentifierModel struct for WorkItemIdentifierModel
type WorkItemIdentifierModel struct {
	// Used for search WorkItem. Internal identifier has a Guid data format. Global identifier has an integer data format
	Id string `json:"id"`
	GlobalId int64 `json:"globalId"`
}

type _WorkItemIdentifierModel WorkItemIdentifierModel

// NewWorkItemIdentifierModel instantiates a new WorkItemIdentifierModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemIdentifierModel(id string, globalId int64) *WorkItemIdentifierModel {
	this := WorkItemIdentifierModel{}
	this.Id = id
	this.GlobalId = globalId
	return &this
}

// NewWorkItemIdentifierModelWithDefaults instantiates a new WorkItemIdentifierModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemIdentifierModelWithDefaults() *WorkItemIdentifierModel {
	this := WorkItemIdentifierModel{}
	return &this
}

// GetId returns the Id field value
func (o *WorkItemIdentifierModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkItemIdentifierModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkItemIdentifierModel) SetId(v string) {
	o.Id = v
}

// GetGlobalId returns the GlobalId field value
func (o *WorkItemIdentifierModel) GetGlobalId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value
// and a boolean to check if the value has been set.
func (o *WorkItemIdentifierModel) GetGlobalIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalId, true
}

// SetGlobalId sets field value
func (o *WorkItemIdentifierModel) SetGlobalId(v int64) {
	o.GlobalId = v
}

func (o WorkItemIdentifierModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemIdentifierModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["globalId"] = o.GlobalId
	return toSerialize, nil
}

func (o *WorkItemIdentifierModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varWorkItemIdentifierModel := _WorkItemIdentifierModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkItemIdentifierModel)

	if err != nil {
		return err
	}

	*o = WorkItemIdentifierModel(varWorkItemIdentifierModel)

	return err
}

type NullableWorkItemIdentifierModel struct {
	value *WorkItemIdentifierModel
	isSet bool
}

func (v NullableWorkItemIdentifierModel) Get() *WorkItemIdentifierModel {
	return v.value
}

func (v *NullableWorkItemIdentifierModel) Set(val *WorkItemIdentifierModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemIdentifierModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemIdentifierModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemIdentifierModel(val *WorkItemIdentifierModel) *NullableWorkItemIdentifierModel {
	return &NullableWorkItemIdentifierModel{value: val, isSet: true}
}

func (v NullableWorkItemIdentifierModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemIdentifierModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


