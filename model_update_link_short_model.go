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

// checks if the UpdateLinkShortModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateLinkShortModel{}

// UpdateLinkShortModel struct for UpdateLinkShortModel
type UpdateLinkShortModel struct {
	Action ActionUpdate `json:"action"`
	Links []LinkPostModel `json:"links,omitempty"`
}

// NewUpdateLinkShortModel instantiates a new UpdateLinkShortModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLinkShortModel(action ActionUpdate) *UpdateLinkShortModel {
	this := UpdateLinkShortModel{}
	this.Action = action
	return &this
}

// NewUpdateLinkShortModelWithDefaults instantiates a new UpdateLinkShortModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLinkShortModelWithDefaults() *UpdateLinkShortModel {
	this := UpdateLinkShortModel{}
	return &this
}

// GetAction returns the Action field value
func (o *UpdateLinkShortModel) GetAction() ActionUpdate {
	if o == nil {
		var ret ActionUpdate
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *UpdateLinkShortModel) GetActionOk() (*ActionUpdate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *UpdateLinkShortModel) SetAction(v ActionUpdate) {
	o.Action = v
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateLinkShortModel) GetLinks() []LinkPostModel {
	if o == nil {
		var ret []LinkPostModel
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateLinkShortModel) GetLinksOk() ([]LinkPostModel, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UpdateLinkShortModel) HasLinks() bool {
	if o != nil && IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkPostModel and assigns it to the Links field.
func (o *UpdateLinkShortModel) SetLinks(v []LinkPostModel) {
	o.Links = v
}

func (o UpdateLinkShortModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateLinkShortModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableUpdateLinkShortModel struct {
	value *UpdateLinkShortModel
	isSet bool
}

func (v NullableUpdateLinkShortModel) Get() *UpdateLinkShortModel {
	return v.value
}

func (v *NullableUpdateLinkShortModel) Set(val *UpdateLinkShortModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLinkShortModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLinkShortModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLinkShortModel(val *UpdateLinkShortModel) *NullableUpdateLinkShortModel {
	return &NullableUpdateLinkShortModel{value: val, isSet: true}
}

func (v NullableUpdateLinkShortModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLinkShortModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

