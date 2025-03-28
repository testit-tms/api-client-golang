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

// FailureCategory the model 'FailureCategory'
type FailureCategory string

// List of FailureCategory
const (
	FAILURECATEGORY_INFRASTRUCTURE_DEFECT FailureCategory = "InfrastructureDefect"
	FAILURECATEGORY_PRODUCT_DEFECT FailureCategory = "ProductDefect"
	FAILURECATEGORY_TEST_DEFECT FailureCategory = "TestDefect"
	FAILURECATEGORY_NO_DEFECT FailureCategory = "NoDefect"
	FAILURECATEGORY_NO_ANALYTICS FailureCategory = "NoAnalytics"
)

// All allowed values of FailureCategory enum
var AllowedFailureCategoryEnumValues = []FailureCategory{
	"InfrastructureDefect",
	"ProductDefect",
	"TestDefect",
	"NoDefect",
	"NoAnalytics",
}

func (v *FailureCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FailureCategory(value)
	for _, existing := range AllowedFailureCategoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FailureCategory", value)
}

// NewFailureCategoryFromValue returns a pointer to a valid FailureCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFailureCategoryFromValue(v string) (*FailureCategory, error) {
	ev := FailureCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FailureCategory: valid values are %v", v, AllowedFailureCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FailureCategory) IsValid() bool {
	for _, existing := range AllowedFailureCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FailureCategory value
func (v FailureCategory) Ptr() *FailureCategory {
	return &v
}

type NullableFailureCategory struct {
	value *FailureCategory
	isSet bool
}

func (v NullableFailureCategory) Get() *FailureCategory {
	return v.value
}

func (v *NullableFailureCategory) Set(val *FailureCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableFailureCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableFailureCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailureCategory(val *FailureCategory) *NullableFailureCategory {
	return &NullableFailureCategory{value: val, isSet: true}
}

func (v NullableFailureCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailureCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

