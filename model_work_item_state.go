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

// WorkItemState the model 'WorkItemState'
type WorkItemState string

// List of WorkItemState
const (
	WORKITEMSTATE_NEEDS_WORK WorkItemState = "NeedsWork"
	WORKITEMSTATE_NOT_READY WorkItemState = "NotReady"
	WORKITEMSTATE_READY WorkItemState = "Ready"
)

// All allowed values of WorkItemState enum
var AllowedWorkItemStateEnumValues = []WorkItemState{
	"NeedsWork",
	"NotReady",
	"Ready",
}

func (v *WorkItemState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WorkItemState(value)
	for _, existing := range AllowedWorkItemStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WorkItemState", value)
}

// NewWorkItemStateFromValue returns a pointer to a valid WorkItemState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWorkItemStateFromValue(v string) (*WorkItemState, error) {
	ev := WorkItemState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WorkItemState: valid values are %v", v, AllowedWorkItemStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WorkItemState) IsValid() bool {
	for _, existing := range AllowedWorkItemStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WorkItemState value
func (v WorkItemState) Ptr() *WorkItemState {
	return &v
}

type NullableWorkItemState struct {
	value *WorkItemState
	isSet bool
}

func (v NullableWorkItemState) Get() *WorkItemState {
	return v.value
}

func (v *NullableWorkItemState) Set(val *WorkItemState) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemState) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemState(val *WorkItemState) *NullableWorkItemState {
	return &NullableWorkItemState{value: val, isSet: true}
}

func (v NullableWorkItemState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

