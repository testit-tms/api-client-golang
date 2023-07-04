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

// BackgroundJobState the model 'BackgroundJobState'
type BackgroundJobState string

// List of BackgroundJobState
const (
	BACKGROUNDJOBSTATE_ENQUEUED BackgroundJobState = "Enqueued"
	BACKGROUNDJOBSTATE_IN_PROGRESS BackgroundJobState = "InProgress"
	BACKGROUNDJOBSTATE_COMPLETED BackgroundJobState = "Completed"
	BACKGROUNDJOBSTATE_FAILED BackgroundJobState = "Failed"
	BACKGROUNDJOBSTATE_CANCELED BackgroundJobState = "Canceled"
)

// All allowed values of BackgroundJobState enum
var AllowedBackgroundJobStateEnumValues = []BackgroundJobState{
	"Enqueued",
	"InProgress",
	"Completed",
	"Failed",
	"Canceled",
}

func (v *BackgroundJobState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BackgroundJobState(value)
	for _, existing := range AllowedBackgroundJobStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BackgroundJobState", value)
}

// NewBackgroundJobStateFromValue returns a pointer to a valid BackgroundJobState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBackgroundJobStateFromValue(v string) (*BackgroundJobState, error) {
	ev := BackgroundJobState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BackgroundJobState: valid values are %v", v, AllowedBackgroundJobStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BackgroundJobState) IsValid() bool {
	for _, existing := range AllowedBackgroundJobStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackgroundJobState value
func (v BackgroundJobState) Ptr() *BackgroundJobState {
	return &v
}

type NullableBackgroundJobState struct {
	value *BackgroundJobState
	isSet bool
}

func (v NullableBackgroundJobState) Get() *BackgroundJobState {
	return v.value
}

func (v *NullableBackgroundJobState) Set(val *BackgroundJobState) {
	v.value = val
	v.isSet = true
}

func (v NullableBackgroundJobState) IsSet() bool {
	return v.isSet
}

func (v *NullableBackgroundJobState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackgroundJobState(val *BackgroundJobState) *NullableBackgroundJobState {
	return &NullableBackgroundJobState{value: val, isSet: true}
}

func (v NullableBackgroundJobState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackgroundJobState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

