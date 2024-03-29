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

// checks if the CreateParameterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateParameterRequest{}

// CreateParameterRequest struct for CreateParameterRequest
type CreateParameterRequest struct {
	// Value of the parameter
	Value string `json:"value"`
	// Key of the parameter
	Name string `json:"name"`
}

// NewCreateParameterRequest instantiates a new CreateParameterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateParameterRequest(value string, name string) *CreateParameterRequest {
	this := CreateParameterRequest{}
	this.Value = value
	this.Name = name
	return &this
}

// NewCreateParameterRequestWithDefaults instantiates a new CreateParameterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateParameterRequestWithDefaults() *CreateParameterRequest {
	this := CreateParameterRequest{}
	return &this
}

// GetValue returns the Value field value
func (o *CreateParameterRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CreateParameterRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CreateParameterRequest) SetValue(v string) {
	o.Value = v
}

// GetName returns the Name field value
func (o *CreateParameterRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateParameterRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateParameterRequest) SetName(v string) {
	o.Name = v
}

func (o CreateParameterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateParameterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableCreateParameterRequest struct {
	value *CreateParameterRequest
	isSet bool
}

func (v NullableCreateParameterRequest) Get() *CreateParameterRequest {
	return v.value
}

func (v *NullableCreateParameterRequest) Set(val *CreateParameterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateParameterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateParameterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateParameterRequest(val *CreateParameterRequest) *NullableCreateParameterRequest {
	return &NullableCreateParameterRequest{value: val, isSet: true}
}

func (v NullableCreateParameterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateParameterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


