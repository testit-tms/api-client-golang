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

// checks if the UserRankModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRankModel{}

// UserRankModel struct for UserRankModel
type UserRankModel struct {
	Score *int32 `json:"score,omitempty"`
	WorkItemsCreated *int32 `json:"workItemsCreated,omitempty"`
	PassedTestPoints *int32 `json:"passedTestPoints,omitempty"`
	FailedTestPoints *int32 `json:"failedTestPoints,omitempty"`
	SkippedTestPoints *int32 `json:"skippedTestPoints,omitempty"`
	BlockedTestPoints *int32 `json:"blockedTestPoints,omitempty"`
}

// NewUserRankModel instantiates a new UserRankModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRankModel() *UserRankModel {
	this := UserRankModel{}
	return &this
}

// NewUserRankModelWithDefaults instantiates a new UserRankModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRankModelWithDefaults() *UserRankModel {
	this := UserRankModel{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *UserRankModel) GetScore() int32 {
	if o == nil || IsNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRankModel) GetScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *UserRankModel) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *UserRankModel) SetScore(v int32) {
	o.Score = &v
}

// GetWorkItemsCreated returns the WorkItemsCreated field value if set, zero value otherwise.
func (o *UserRankModel) GetWorkItemsCreated() int32 {
	if o == nil || IsNil(o.WorkItemsCreated) {
		var ret int32
		return ret
	}
	return *o.WorkItemsCreated
}

// GetWorkItemsCreatedOk returns a tuple with the WorkItemsCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRankModel) GetWorkItemsCreatedOk() (*int32, bool) {
	if o == nil || IsNil(o.WorkItemsCreated) {
		return nil, false
	}
	return o.WorkItemsCreated, true
}

// HasWorkItemsCreated returns a boolean if a field has been set.
func (o *UserRankModel) HasWorkItemsCreated() bool {
	if o != nil && !IsNil(o.WorkItemsCreated) {
		return true
	}

	return false
}

// SetWorkItemsCreated gets a reference to the given int32 and assigns it to the WorkItemsCreated field.
func (o *UserRankModel) SetWorkItemsCreated(v int32) {
	o.WorkItemsCreated = &v
}

// GetPassedTestPoints returns the PassedTestPoints field value if set, zero value otherwise.
func (o *UserRankModel) GetPassedTestPoints() int32 {
	if o == nil || IsNil(o.PassedTestPoints) {
		var ret int32
		return ret
	}
	return *o.PassedTestPoints
}

// GetPassedTestPointsOk returns a tuple with the PassedTestPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRankModel) GetPassedTestPointsOk() (*int32, bool) {
	if o == nil || IsNil(o.PassedTestPoints) {
		return nil, false
	}
	return o.PassedTestPoints, true
}

// HasPassedTestPoints returns a boolean if a field has been set.
func (o *UserRankModel) HasPassedTestPoints() bool {
	if o != nil && !IsNil(o.PassedTestPoints) {
		return true
	}

	return false
}

// SetPassedTestPoints gets a reference to the given int32 and assigns it to the PassedTestPoints field.
func (o *UserRankModel) SetPassedTestPoints(v int32) {
	o.PassedTestPoints = &v
}

// GetFailedTestPoints returns the FailedTestPoints field value if set, zero value otherwise.
func (o *UserRankModel) GetFailedTestPoints() int32 {
	if o == nil || IsNil(o.FailedTestPoints) {
		var ret int32
		return ret
	}
	return *o.FailedTestPoints
}

// GetFailedTestPointsOk returns a tuple with the FailedTestPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRankModel) GetFailedTestPointsOk() (*int32, bool) {
	if o == nil || IsNil(o.FailedTestPoints) {
		return nil, false
	}
	return o.FailedTestPoints, true
}

// HasFailedTestPoints returns a boolean if a field has been set.
func (o *UserRankModel) HasFailedTestPoints() bool {
	if o != nil && !IsNil(o.FailedTestPoints) {
		return true
	}

	return false
}

// SetFailedTestPoints gets a reference to the given int32 and assigns it to the FailedTestPoints field.
func (o *UserRankModel) SetFailedTestPoints(v int32) {
	o.FailedTestPoints = &v
}

// GetSkippedTestPoints returns the SkippedTestPoints field value if set, zero value otherwise.
func (o *UserRankModel) GetSkippedTestPoints() int32 {
	if o == nil || IsNil(o.SkippedTestPoints) {
		var ret int32
		return ret
	}
	return *o.SkippedTestPoints
}

// GetSkippedTestPointsOk returns a tuple with the SkippedTestPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRankModel) GetSkippedTestPointsOk() (*int32, bool) {
	if o == nil || IsNil(o.SkippedTestPoints) {
		return nil, false
	}
	return o.SkippedTestPoints, true
}

// HasSkippedTestPoints returns a boolean if a field has been set.
func (o *UserRankModel) HasSkippedTestPoints() bool {
	if o != nil && !IsNil(o.SkippedTestPoints) {
		return true
	}

	return false
}

// SetSkippedTestPoints gets a reference to the given int32 and assigns it to the SkippedTestPoints field.
func (o *UserRankModel) SetSkippedTestPoints(v int32) {
	o.SkippedTestPoints = &v
}

// GetBlockedTestPoints returns the BlockedTestPoints field value if set, zero value otherwise.
func (o *UserRankModel) GetBlockedTestPoints() int32 {
	if o == nil || IsNil(o.BlockedTestPoints) {
		var ret int32
		return ret
	}
	return *o.BlockedTestPoints
}

// GetBlockedTestPointsOk returns a tuple with the BlockedTestPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRankModel) GetBlockedTestPointsOk() (*int32, bool) {
	if o == nil || IsNil(o.BlockedTestPoints) {
		return nil, false
	}
	return o.BlockedTestPoints, true
}

// HasBlockedTestPoints returns a boolean if a field has been set.
func (o *UserRankModel) HasBlockedTestPoints() bool {
	if o != nil && !IsNil(o.BlockedTestPoints) {
		return true
	}

	return false
}

// SetBlockedTestPoints gets a reference to the given int32 and assigns it to the BlockedTestPoints field.
func (o *UserRankModel) SetBlockedTestPoints(v int32) {
	o.BlockedTestPoints = &v
}

func (o UserRankModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRankModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.WorkItemsCreated) {
		toSerialize["workItemsCreated"] = o.WorkItemsCreated
	}
	if !IsNil(o.PassedTestPoints) {
		toSerialize["passedTestPoints"] = o.PassedTestPoints
	}
	if !IsNil(o.FailedTestPoints) {
		toSerialize["failedTestPoints"] = o.FailedTestPoints
	}
	if !IsNil(o.SkippedTestPoints) {
		toSerialize["skippedTestPoints"] = o.SkippedTestPoints
	}
	if !IsNil(o.BlockedTestPoints) {
		toSerialize["blockedTestPoints"] = o.BlockedTestPoints
	}
	return toSerialize, nil
}

type NullableUserRankModel struct {
	value *UserRankModel
	isSet bool
}

func (v NullableUserRankModel) Get() *UserRankModel {
	return v.value
}

func (v *NullableUserRankModel) Set(val *UserRankModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRankModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRankModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRankModel(val *UserRankModel) *NullableUserRankModel {
	return &NullableUserRankModel{value: val, isSet: true}
}

func (v NullableUserRankModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRankModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

