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
	Score int32 `json:"score"`
	WorkItemsCreated int32 `json:"workItemsCreated"`
	PassedTestPoints int32 `json:"passedTestPoints"`
	FailedTestPoints int32 `json:"failedTestPoints"`
	SkippedTestPoints int32 `json:"skippedTestPoints"`
	BlockedTestPoints int32 `json:"blockedTestPoints"`
}

// NewUserRankModel instantiates a new UserRankModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRankModel(score int32, workItemsCreated int32, passedTestPoints int32, failedTestPoints int32, skippedTestPoints int32, blockedTestPoints int32) *UserRankModel {
	this := UserRankModel{}
	this.Score = score
	this.WorkItemsCreated = workItemsCreated
	this.PassedTestPoints = passedTestPoints
	this.FailedTestPoints = failedTestPoints
	this.SkippedTestPoints = skippedTestPoints
	this.BlockedTestPoints = blockedTestPoints
	return &this
}

// NewUserRankModelWithDefaults instantiates a new UserRankModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRankModelWithDefaults() *UserRankModel {
	this := UserRankModel{}
	return &this
}

// GetScore returns the Score field value
func (o *UserRankModel) GetScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *UserRankModel) GetScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *UserRankModel) SetScore(v int32) {
	o.Score = v
}

// GetWorkItemsCreated returns the WorkItemsCreated field value
func (o *UserRankModel) GetWorkItemsCreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WorkItemsCreated
}

// GetWorkItemsCreatedOk returns a tuple with the WorkItemsCreated field value
// and a boolean to check if the value has been set.
func (o *UserRankModel) GetWorkItemsCreatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkItemsCreated, true
}

// SetWorkItemsCreated sets field value
func (o *UserRankModel) SetWorkItemsCreated(v int32) {
	o.WorkItemsCreated = v
}

// GetPassedTestPoints returns the PassedTestPoints field value
func (o *UserRankModel) GetPassedTestPoints() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PassedTestPoints
}

// GetPassedTestPointsOk returns a tuple with the PassedTestPoints field value
// and a boolean to check if the value has been set.
func (o *UserRankModel) GetPassedTestPointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PassedTestPoints, true
}

// SetPassedTestPoints sets field value
func (o *UserRankModel) SetPassedTestPoints(v int32) {
	o.PassedTestPoints = v
}

// GetFailedTestPoints returns the FailedTestPoints field value
func (o *UserRankModel) GetFailedTestPoints() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FailedTestPoints
}

// GetFailedTestPointsOk returns a tuple with the FailedTestPoints field value
// and a boolean to check if the value has been set.
func (o *UserRankModel) GetFailedTestPointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailedTestPoints, true
}

// SetFailedTestPoints sets field value
func (o *UserRankModel) SetFailedTestPoints(v int32) {
	o.FailedTestPoints = v
}

// GetSkippedTestPoints returns the SkippedTestPoints field value
func (o *UserRankModel) GetSkippedTestPoints() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SkippedTestPoints
}

// GetSkippedTestPointsOk returns a tuple with the SkippedTestPoints field value
// and a boolean to check if the value has been set.
func (o *UserRankModel) GetSkippedTestPointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkippedTestPoints, true
}

// SetSkippedTestPoints sets field value
func (o *UserRankModel) SetSkippedTestPoints(v int32) {
	o.SkippedTestPoints = v
}

// GetBlockedTestPoints returns the BlockedTestPoints field value
func (o *UserRankModel) GetBlockedTestPoints() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockedTestPoints
}

// GetBlockedTestPointsOk returns a tuple with the BlockedTestPoints field value
// and a boolean to check if the value has been set.
func (o *UserRankModel) GetBlockedTestPointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockedTestPoints, true
}

// SetBlockedTestPoints sets field value
func (o *UserRankModel) SetBlockedTestPoints(v int32) {
	o.BlockedTestPoints = v
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
	toSerialize["score"] = o.Score
	toSerialize["workItemsCreated"] = o.WorkItemsCreated
	toSerialize["passedTestPoints"] = o.PassedTestPoints
	toSerialize["failedTestPoints"] = o.FailedTestPoints
	toSerialize["skippedTestPoints"] = o.SkippedTestPoints
	toSerialize["blockedTestPoints"] = o.BlockedTestPoints
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


