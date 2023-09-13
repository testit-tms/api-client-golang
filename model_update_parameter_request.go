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

// checks if the UpdateParameterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateParameterRequest{}

// UpdateParameterRequest struct for UpdateParameterRequest
type UpdateParameterRequest struct {
	Id string `json:"id"`
	Value string `json:"value"`
	Name string `json:"name"`
}

// NewUpdateParameterRequest instantiates a new UpdateParameterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateParameterRequest(id string, value string, name string) *UpdateParameterRequest {
	this := UpdateParameterRequest{}
	this.Id = id
	this.Value = value
	this.Name = name
	return &this
}

// NewUpdateParameterRequestWithDefaults instantiates a new UpdateParameterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateParameterRequestWithDefaults() *UpdateParameterRequest {
	this := UpdateParameterRequest{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateParameterRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateParameterRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateParameterRequest) SetId(v string) {
	o.Id = v
}

// GetValue returns the Value field value
func (o *UpdateParameterRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UpdateParameterRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UpdateParameterRequest) SetValue(v string) {
	o.Value = v
}

// GetName returns the Name field value
func (o *UpdateParameterRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateParameterRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateParameterRequest) SetName(v string) {
	o.Name = v
}

func (o UpdateParameterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateParameterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["value"] = o.Value
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableUpdateParameterRequest struct {
	value *UpdateParameterRequest
	isSet bool
}

func (v NullableUpdateParameterRequest) Get() *UpdateParameterRequest {
	return v.value
}

func (v *NullableUpdateParameterRequest) Set(val *UpdateParameterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateParameterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateParameterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateParameterRequest(val *UpdateParameterRequest) *NullableUpdateParameterRequest {
	return &NullableUpdateParameterRequest{value: val, isSet: true}
}

func (v NullableUpdateParameterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateParameterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

