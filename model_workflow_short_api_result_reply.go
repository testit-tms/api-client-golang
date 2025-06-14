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

// checks if the WorkflowShortApiResultReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowShortApiResultReply{}

// WorkflowShortApiResultReply struct for WorkflowShortApiResultReply
type WorkflowShortApiResultReply struct {
	Data []WorkflowShortApiResult `json:"data"`
	TotalCount int32 `json:"totalCount"`
}

type _WorkflowShortApiResultReply WorkflowShortApiResultReply

// NewWorkflowShortApiResultReply instantiates a new WorkflowShortApiResultReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowShortApiResultReply(data []WorkflowShortApiResult, totalCount int32) *WorkflowShortApiResultReply {
	this := WorkflowShortApiResultReply{}
	this.Data = data
	this.TotalCount = totalCount
	return &this
}

// NewWorkflowShortApiResultReplyWithDefaults instantiates a new WorkflowShortApiResultReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowShortApiResultReplyWithDefaults() *WorkflowShortApiResultReply {
	this := WorkflowShortApiResultReply{}
	return &this
}

// GetData returns the Data field value
func (o *WorkflowShortApiResultReply) GetData() []WorkflowShortApiResult {
	if o == nil {
		var ret []WorkflowShortApiResult
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *WorkflowShortApiResultReply) GetDataOk() ([]WorkflowShortApiResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *WorkflowShortApiResultReply) SetData(v []WorkflowShortApiResult) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value
func (o *WorkflowShortApiResultReply) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *WorkflowShortApiResultReply) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *WorkflowShortApiResultReply) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o WorkflowShortApiResultReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowShortApiResultReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["totalCount"] = o.TotalCount
	return toSerialize, nil
}

func (o *WorkflowShortApiResultReply) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"totalCount",
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

	varWorkflowShortApiResultReply := _WorkflowShortApiResultReply{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkflowShortApiResultReply)

	if err != nil {
		return err
	}

	*o = WorkflowShortApiResultReply(varWorkflowShortApiResultReply)

	return err
}

type NullableWorkflowShortApiResultReply struct {
	value *WorkflowShortApiResultReply
	isSet bool
}

func (v NullableWorkflowShortApiResultReply) Get() *WorkflowShortApiResultReply {
	return v.value
}

func (v *NullableWorkflowShortApiResultReply) Set(val *WorkflowShortApiResultReply) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowShortApiResultReply) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowShortApiResultReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowShortApiResultReply(val *WorkflowShortApiResultReply) *NullableWorkflowShortApiResultReply {
	return &NullableWorkflowShortApiResultReply{value: val, isSet: true}
}

func (v NullableWorkflowShortApiResultReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowShortApiResultReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


