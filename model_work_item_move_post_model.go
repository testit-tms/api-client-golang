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

// checks if the WorkItemMovePostModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemMovePostModel{}

// WorkItemMovePostModel struct for WorkItemMovePostModel
type WorkItemMovePostModel struct {
	Id string `json:"id"`
	NewSectionId string `json:"newSectionId"`
	OldSectionId NullableString `json:"oldSectionId,omitempty"`
	NextWorkItemId NullableString `json:"nextWorkItemId,omitempty"`
}

// NewWorkItemMovePostModel instantiates a new WorkItemMovePostModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemMovePostModel(id string, newSectionId string) *WorkItemMovePostModel {
	this := WorkItemMovePostModel{}
	this.Id = id
	this.NewSectionId = newSectionId
	return &this
}

// NewWorkItemMovePostModelWithDefaults instantiates a new WorkItemMovePostModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemMovePostModelWithDefaults() *WorkItemMovePostModel {
	this := WorkItemMovePostModel{}
	return &this
}

// GetId returns the Id field value
func (o *WorkItemMovePostModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkItemMovePostModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkItemMovePostModel) SetId(v string) {
	o.Id = v
}

// GetNewSectionId returns the NewSectionId field value
func (o *WorkItemMovePostModel) GetNewSectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewSectionId
}

// GetNewSectionIdOk returns a tuple with the NewSectionId field value
// and a boolean to check if the value has been set.
func (o *WorkItemMovePostModel) GetNewSectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewSectionId, true
}

// SetNewSectionId sets field value
func (o *WorkItemMovePostModel) SetNewSectionId(v string) {
	o.NewSectionId = v
}

// GetOldSectionId returns the OldSectionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemMovePostModel) GetOldSectionId() string {
	if o == nil || IsNil(o.OldSectionId.Get()) {
		var ret string
		return ret
	}
	return *o.OldSectionId.Get()
}

// GetOldSectionIdOk returns a tuple with the OldSectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemMovePostModel) GetOldSectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OldSectionId.Get(), o.OldSectionId.IsSet()
}

// HasOldSectionId returns a boolean if a field has been set.
func (o *WorkItemMovePostModel) HasOldSectionId() bool {
	if o != nil && o.OldSectionId.IsSet() {
		return true
	}

	return false
}

// SetOldSectionId gets a reference to the given NullableString and assigns it to the OldSectionId field.
func (o *WorkItemMovePostModel) SetOldSectionId(v string) {
	o.OldSectionId.Set(&v)
}
// SetOldSectionIdNil sets the value for OldSectionId to be an explicit nil
func (o *WorkItemMovePostModel) SetOldSectionIdNil() {
	o.OldSectionId.Set(nil)
}

// UnsetOldSectionId ensures that no value is present for OldSectionId, not even an explicit nil
func (o *WorkItemMovePostModel) UnsetOldSectionId() {
	o.OldSectionId.Unset()
}

// GetNextWorkItemId returns the NextWorkItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemMovePostModel) GetNextWorkItemId() string {
	if o == nil || IsNil(o.NextWorkItemId.Get()) {
		var ret string
		return ret
	}
	return *o.NextWorkItemId.Get()
}

// GetNextWorkItemIdOk returns a tuple with the NextWorkItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemMovePostModel) GetNextWorkItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextWorkItemId.Get(), o.NextWorkItemId.IsSet()
}

// HasNextWorkItemId returns a boolean if a field has been set.
func (o *WorkItemMovePostModel) HasNextWorkItemId() bool {
	if o != nil && o.NextWorkItemId.IsSet() {
		return true
	}

	return false
}

// SetNextWorkItemId gets a reference to the given NullableString and assigns it to the NextWorkItemId field.
func (o *WorkItemMovePostModel) SetNextWorkItemId(v string) {
	o.NextWorkItemId.Set(&v)
}
// SetNextWorkItemIdNil sets the value for NextWorkItemId to be an explicit nil
func (o *WorkItemMovePostModel) SetNextWorkItemIdNil() {
	o.NextWorkItemId.Set(nil)
}

// UnsetNextWorkItemId ensures that no value is present for NextWorkItemId, not even an explicit nil
func (o *WorkItemMovePostModel) UnsetNextWorkItemId() {
	o.NextWorkItemId.Unset()
}

func (o WorkItemMovePostModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemMovePostModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["newSectionId"] = o.NewSectionId
	if o.OldSectionId.IsSet() {
		toSerialize["oldSectionId"] = o.OldSectionId.Get()
	}
	if o.NextWorkItemId.IsSet() {
		toSerialize["nextWorkItemId"] = o.NextWorkItemId.Get()
	}
	return toSerialize, nil
}

type NullableWorkItemMovePostModel struct {
	value *WorkItemMovePostModel
	isSet bool
}

func (v NullableWorkItemMovePostModel) Get() *WorkItemMovePostModel {
	return v.value
}

func (v *NullableWorkItemMovePostModel) Set(val *WorkItemMovePostModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemMovePostModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemMovePostModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemMovePostModel(val *WorkItemMovePostModel) *NullableWorkItemMovePostModel {
	return &NullableWorkItemMovePostModel{value: val, isSet: true}
}

func (v NullableWorkItemMovePostModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemMovePostModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


