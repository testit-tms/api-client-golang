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

// checks if the UpdateCustomAttributeTestPlanProjectRelationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCustomAttributeTestPlanProjectRelationsRequest{}

// UpdateCustomAttributeTestPlanProjectRelationsRequest struct for UpdateCustomAttributeTestPlanProjectRelationsRequest
type UpdateCustomAttributeTestPlanProjectRelationsRequest struct {
	// Custom attribute internal unique identifier
	Id string `json:"id"`
	// Flag that defines if custom attribute is enabled
	IsEnabled bool `json:"isEnabled"`
	// Flag that defines if custom attribute is required
	IsRequired bool `json:"isRequired"`
}

// NewUpdateCustomAttributeTestPlanProjectRelationsRequest instantiates a new UpdateCustomAttributeTestPlanProjectRelationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCustomAttributeTestPlanProjectRelationsRequest(id string, isEnabled bool, isRequired bool) *UpdateCustomAttributeTestPlanProjectRelationsRequest {
	this := UpdateCustomAttributeTestPlanProjectRelationsRequest{}
	this.Id = id
	this.IsEnabled = isEnabled
	this.IsRequired = isRequired
	return &this
}

// NewUpdateCustomAttributeTestPlanProjectRelationsRequestWithDefaults instantiates a new UpdateCustomAttributeTestPlanProjectRelationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCustomAttributeTestPlanProjectRelationsRequestWithDefaults() *UpdateCustomAttributeTestPlanProjectRelationsRequest {
	this := UpdateCustomAttributeTestPlanProjectRelationsRequest{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateCustomAttributeTestPlanProjectRelationsRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateCustomAttributeTestPlanProjectRelationsRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateCustomAttributeTestPlanProjectRelationsRequest) SetId(v string) {
	o.Id = v
}

// GetIsEnabled returns the IsEnabled field value
func (o *UpdateCustomAttributeTestPlanProjectRelationsRequest) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *UpdateCustomAttributeTestPlanProjectRelationsRequest) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *UpdateCustomAttributeTestPlanProjectRelationsRequest) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetIsRequired returns the IsRequired field value
func (o *UpdateCustomAttributeTestPlanProjectRelationsRequest) GetIsRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value
// and a boolean to check if the value has been set.
func (o *UpdateCustomAttributeTestPlanProjectRelationsRequest) GetIsRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRequired, true
}

// SetIsRequired sets field value
func (o *UpdateCustomAttributeTestPlanProjectRelationsRequest) SetIsRequired(v bool) {
	o.IsRequired = v
}

func (o UpdateCustomAttributeTestPlanProjectRelationsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCustomAttributeTestPlanProjectRelationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["isEnabled"] = o.IsEnabled
	toSerialize["isRequired"] = o.IsRequired
	return toSerialize, nil
}

type NullableUpdateCustomAttributeTestPlanProjectRelationsRequest struct {
	value *UpdateCustomAttributeTestPlanProjectRelationsRequest
	isSet bool
}

func (v NullableUpdateCustomAttributeTestPlanProjectRelationsRequest) Get() *UpdateCustomAttributeTestPlanProjectRelationsRequest {
	return v.value
}

func (v *NullableUpdateCustomAttributeTestPlanProjectRelationsRequest) Set(val *UpdateCustomAttributeTestPlanProjectRelationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCustomAttributeTestPlanProjectRelationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCustomAttributeTestPlanProjectRelationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCustomAttributeTestPlanProjectRelationsRequest(val *UpdateCustomAttributeTestPlanProjectRelationsRequest) *NullableUpdateCustomAttributeTestPlanProjectRelationsRequest {
	return &NullableUpdateCustomAttributeTestPlanProjectRelationsRequest{value: val, isSet: true}
}

func (v NullableUpdateCustomAttributeTestPlanProjectRelationsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCustomAttributeTestPlanProjectRelationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


