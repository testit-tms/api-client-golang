/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UpdateMultipleLinksApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMultipleLinksApiModel{}

// UpdateMultipleLinksApiModel struct for UpdateMultipleLinksApiModel
type UpdateMultipleLinksApiModel struct {
	Action ActionUpdate `json:"action"`
	Links []CreateLinkApiModel `json:"links,omitempty"`
}

type _UpdateMultipleLinksApiModel UpdateMultipleLinksApiModel

// NewUpdateMultipleLinksApiModel instantiates a new UpdateMultipleLinksApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMultipleLinksApiModel(action ActionUpdate) *UpdateMultipleLinksApiModel {
	this := UpdateMultipleLinksApiModel{}
	this.Action = action
	return &this
}

// NewUpdateMultipleLinksApiModelWithDefaults instantiates a new UpdateMultipleLinksApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMultipleLinksApiModelWithDefaults() *UpdateMultipleLinksApiModel {
	this := UpdateMultipleLinksApiModel{}
	return &this
}

// GetAction returns the Action field value
func (o *UpdateMultipleLinksApiModel) GetAction() ActionUpdate {
	if o == nil {
		var ret ActionUpdate
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *UpdateMultipleLinksApiModel) GetActionOk() (*ActionUpdate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *UpdateMultipleLinksApiModel) SetAction(v ActionUpdate) {
	o.Action = v
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMultipleLinksApiModel) GetLinks() []CreateLinkApiModel {
	if o == nil {
		var ret []CreateLinkApiModel
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMultipleLinksApiModel) GetLinksOk() ([]CreateLinkApiModel, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UpdateMultipleLinksApiModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []CreateLinkApiModel and assigns it to the Links field.
func (o *UpdateMultipleLinksApiModel) SetLinks(v []CreateLinkApiModel) {
	o.Links = v
}

func (o UpdateMultipleLinksApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMultipleLinksApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *UpdateMultipleLinksApiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateMultipleLinksApiModel := _UpdateMultipleLinksApiModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateMultipleLinksApiModel)

	if err != nil {
		return err
	}

	*o = UpdateMultipleLinksApiModel(varUpdateMultipleLinksApiModel)

	return err
}

type NullableUpdateMultipleLinksApiModel struct {
	value *UpdateMultipleLinksApiModel
	isSet bool
}

func (v NullableUpdateMultipleLinksApiModel) Get() *UpdateMultipleLinksApiModel {
	return v.value
}

func (v *NullableUpdateMultipleLinksApiModel) Set(val *UpdateMultipleLinksApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMultipleLinksApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMultipleLinksApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMultipleLinksApiModel(val *UpdateMultipleLinksApiModel) *NullableUpdateMultipleLinksApiModel {
	return &NullableUpdateMultipleLinksApiModel{value: val, isSet: true}
}

func (v NullableUpdateMultipleLinksApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMultipleLinksApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


