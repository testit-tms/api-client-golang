/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
	"fmt"
)

// RequestTypeModel the model 'RequestTypeModel'
type RequestTypeModel string

// List of RequestTypeModel
const (
	REQUESTTYPEMODEL_POST RequestTypeModel = "Post"
	REQUESTTYPEMODEL_PUT RequestTypeModel = "Put"
	REQUESTTYPEMODEL_DELETE RequestTypeModel = "Delete"
)

// All allowed values of RequestTypeModel enum
var AllowedRequestTypeModelEnumValues = []RequestTypeModel{
	"Post",
	"Put",
	"Delete",
}

func (v *RequestTypeModel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestTypeModel(value)
	for _, existing := range AllowedRequestTypeModelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestTypeModel", value)
}

// NewRequestTypeModelFromValue returns a pointer to a valid RequestTypeModel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestTypeModelFromValue(v string) (*RequestTypeModel, error) {
	ev := RequestTypeModel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestTypeModel: valid values are %v", v, AllowedRequestTypeModelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestTypeModel) IsValid() bool {
	for _, existing := range AllowedRequestTypeModelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RequestTypeModel value
func (v RequestTypeModel) Ptr() *RequestTypeModel {
	return &v
}

type NullableRequestTypeModel struct {
	value *RequestTypeModel
	isSet bool
}

func (v NullableRequestTypeModel) Get() *RequestTypeModel {
	return v.value
}

func (v *NullableRequestTypeModel) Set(val *RequestTypeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeModel(val *RequestTypeModel) *NullableRequestTypeModel {
	return &NullableRequestTypeModel{value: val, isSet: true}
}

func (v NullableRequestTypeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
