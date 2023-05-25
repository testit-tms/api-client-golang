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

// FailureCategoryModel the model 'FailureCategoryModel'
type FailureCategoryModel string

// List of FailureCategoryModel
const (
	FAILURECATEGORYMODEL_INFRASTRUCTURE_DEFECT FailureCategoryModel = "InfrastructureDefect"
	FAILURECATEGORYMODEL_PRODUCT_DEFECT FailureCategoryModel = "ProductDefect"
	FAILURECATEGORYMODEL_TEST_DEFECT FailureCategoryModel = "TestDefect"
	FAILURECATEGORYMODEL_NO_DEFECT FailureCategoryModel = "NoDefect"
)

// All allowed values of FailureCategoryModel enum
var AllowedFailureCategoryModelEnumValues = []FailureCategoryModel{
	"InfrastructureDefect",
	"ProductDefect",
	"TestDefect",
	"NoDefect",
}

func (v *FailureCategoryModel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FailureCategoryModel(value)
	for _, existing := range AllowedFailureCategoryModelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FailureCategoryModel", value)
}

// NewFailureCategoryModelFromValue returns a pointer to a valid FailureCategoryModel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFailureCategoryModelFromValue(v string) (*FailureCategoryModel, error) {
	ev := FailureCategoryModel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FailureCategoryModel: valid values are %v", v, AllowedFailureCategoryModelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FailureCategoryModel) IsValid() bool {
	for _, existing := range AllowedFailureCategoryModelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FailureCategoryModel value
func (v FailureCategoryModel) Ptr() *FailureCategoryModel {
	return &v
}

type NullableFailureCategoryModel struct {
	value *FailureCategoryModel
	isSet bool
}

func (v NullableFailureCategoryModel) Get() *FailureCategoryModel {
	return v.value
}

func (v *NullableFailureCategoryModel) Set(val *FailureCategoryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFailureCategoryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFailureCategoryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailureCategoryModel(val *FailureCategoryModel) *NullableFailureCategoryModel {
	return &NullableFailureCategoryModel{value: val, isSet: true}
}

func (v NullableFailureCategoryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailureCategoryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

