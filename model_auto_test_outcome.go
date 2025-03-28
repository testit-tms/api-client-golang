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

// AutoTestOutcome the model 'AutoTestOutcome'
type AutoTestOutcome string

// List of AutoTestOutcome
const (
	AUTOTESTOUTCOME_PASSED AutoTestOutcome = "Passed"
	AUTOTESTOUTCOME_FAILED AutoTestOutcome = "Failed"
	AUTOTESTOUTCOME_BLOCKED AutoTestOutcome = "Blocked"
	AUTOTESTOUTCOME_SKIPPED AutoTestOutcome = "Skipped"
)

// All allowed values of AutoTestOutcome enum
var AllowedAutoTestOutcomeEnumValues = []AutoTestOutcome{
	"Passed",
	"Failed",
	"Blocked",
	"Skipped",
}

func (v *AutoTestOutcome) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AutoTestOutcome(value)
	for _, existing := range AllowedAutoTestOutcomeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AutoTestOutcome", value)
}

// NewAutoTestOutcomeFromValue returns a pointer to a valid AutoTestOutcome
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAutoTestOutcomeFromValue(v string) (*AutoTestOutcome, error) {
	ev := AutoTestOutcome(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AutoTestOutcome: valid values are %v", v, AllowedAutoTestOutcomeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AutoTestOutcome) IsValid() bool {
	for _, existing := range AllowedAutoTestOutcomeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AutoTestOutcome value
func (v AutoTestOutcome) Ptr() *AutoTestOutcome {
	return &v
}

type NullableAutoTestOutcome struct {
	value *AutoTestOutcome
	isSet bool
}

func (v NullableAutoTestOutcome) Get() *AutoTestOutcome {
	return v.value
}

func (v *NullableAutoTestOutcome) Set(val *AutoTestOutcome) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoTestOutcome) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoTestOutcome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoTestOutcome(val *AutoTestOutcome) *NullableAutoTestOutcome {
	return &NullableAutoTestOutcome{value: val, isSet: true}
}

func (v NullableAutoTestOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoTestOutcome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

