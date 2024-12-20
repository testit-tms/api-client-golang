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

// checks if the UserCustomNameValidationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserCustomNameValidationResponse{}

// UserCustomNameValidationResponse struct for UserCustomNameValidationResponse
type UserCustomNameValidationResponse struct {
	Exists bool `json:"exists"`
}

type _UserCustomNameValidationResponse UserCustomNameValidationResponse

// NewUserCustomNameValidationResponse instantiates a new UserCustomNameValidationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCustomNameValidationResponse(exists bool) *UserCustomNameValidationResponse {
	this := UserCustomNameValidationResponse{}
	this.Exists = exists
	return &this
}

// NewUserCustomNameValidationResponseWithDefaults instantiates a new UserCustomNameValidationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCustomNameValidationResponseWithDefaults() *UserCustomNameValidationResponse {
	this := UserCustomNameValidationResponse{}
	return &this
}

// GetExists returns the Exists field value
func (o *UserCustomNameValidationResponse) GetExists() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Exists
}

// GetExistsOk returns a tuple with the Exists field value
// and a boolean to check if the value has been set.
func (o *UserCustomNameValidationResponse) GetExistsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exists, true
}

// SetExists sets field value
func (o *UserCustomNameValidationResponse) SetExists(v bool) {
	o.Exists = v
}

func (o UserCustomNameValidationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCustomNameValidationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["exists"] = o.Exists
	return toSerialize, nil
}

func (o *UserCustomNameValidationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"exists",
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

	varUserCustomNameValidationResponse := _UserCustomNameValidationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserCustomNameValidationResponse)

	if err != nil {
		return err
	}

	*o = UserCustomNameValidationResponse(varUserCustomNameValidationResponse)

	return err
}

type NullableUserCustomNameValidationResponse struct {
	value *UserCustomNameValidationResponse
	isSet bool
}

func (v NullableUserCustomNameValidationResponse) Get() *UserCustomNameValidationResponse {
	return v.value
}

func (v *NullableUserCustomNameValidationResponse) Set(val *UserCustomNameValidationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCustomNameValidationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCustomNameValidationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCustomNameValidationResponse(val *UserCustomNameValidationResponse) *NullableUserCustomNameValidationResponse {
	return &NullableUserCustomNameValidationResponse{value: val, isSet: true}
}

func (v NullableUserCustomNameValidationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCustomNameValidationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


