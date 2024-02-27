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

// checks if the ApiV2TagsDeleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV2TagsDeleteRequest{}

// ApiV2TagsDeleteRequest struct for ApiV2TagsDeleteRequest
type ApiV2TagsDeleteRequest struct {
	Filter NullableTagsFilterModel `json:"filter,omitempty"`
	ExtractionModel NullableTagExtractionModel `json:"extractionModel,omitempty"`
}

// NewApiV2TagsDeleteRequest instantiates a new ApiV2TagsDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2TagsDeleteRequest() *ApiV2TagsDeleteRequest {
	this := ApiV2TagsDeleteRequest{}
	return &this
}

// NewApiV2TagsDeleteRequestWithDefaults instantiates a new ApiV2TagsDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2TagsDeleteRequestWithDefaults() *ApiV2TagsDeleteRequest {
	this := ApiV2TagsDeleteRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TagsDeleteRequest) GetFilter() TagsFilterModel {
	if o == nil || IsNil(o.Filter.Get()) {
		var ret TagsFilterModel
		return ret
	}
	return *o.Filter.Get()
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TagsDeleteRequest) GetFilterOk() (*TagsFilterModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter.Get(), o.Filter.IsSet()
}

// HasFilter returns a boolean if a field has been set.
func (o *ApiV2TagsDeleteRequest) HasFilter() bool {
	if o != nil && o.Filter.IsSet() {
		return true
	}

	return false
}

// SetFilter gets a reference to the given NullableTagsFilterModel and assigns it to the Filter field.
func (o *ApiV2TagsDeleteRequest) SetFilter(v TagsFilterModel) {
	o.Filter.Set(&v)
}
// SetFilterNil sets the value for Filter to be an explicit nil
func (o *ApiV2TagsDeleteRequest) SetFilterNil() {
	o.Filter.Set(nil)
}

// UnsetFilter ensures that no value is present for Filter, not even an explicit nil
func (o *ApiV2TagsDeleteRequest) UnsetFilter() {
	o.Filter.Unset()
}

// GetExtractionModel returns the ExtractionModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TagsDeleteRequest) GetExtractionModel() TagExtractionModel {
	if o == nil || IsNil(o.ExtractionModel.Get()) {
		var ret TagExtractionModel
		return ret
	}
	return *o.ExtractionModel.Get()
}

// GetExtractionModelOk returns a tuple with the ExtractionModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TagsDeleteRequest) GetExtractionModelOk() (*TagExtractionModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtractionModel.Get(), o.ExtractionModel.IsSet()
}

// HasExtractionModel returns a boolean if a field has been set.
func (o *ApiV2TagsDeleteRequest) HasExtractionModel() bool {
	if o != nil && o.ExtractionModel.IsSet() {
		return true
	}

	return false
}

// SetExtractionModel gets a reference to the given NullableTagExtractionModel and assigns it to the ExtractionModel field.
func (o *ApiV2TagsDeleteRequest) SetExtractionModel(v TagExtractionModel) {
	o.ExtractionModel.Set(&v)
}
// SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil
func (o *ApiV2TagsDeleteRequest) SetExtractionModelNil() {
	o.ExtractionModel.Set(nil)
}

// UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil
func (o *ApiV2TagsDeleteRequest) UnsetExtractionModel() {
	o.ExtractionModel.Unset()
}

func (o ApiV2TagsDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV2TagsDeleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter.IsSet() {
		toSerialize["filter"] = o.Filter.Get()
	}
	if o.ExtractionModel.IsSet() {
		toSerialize["extractionModel"] = o.ExtractionModel.Get()
	}
	return toSerialize, nil
}

type NullableApiV2TagsDeleteRequest struct {
	value *ApiV2TagsDeleteRequest
	isSet bool
}

func (v NullableApiV2TagsDeleteRequest) Get() *ApiV2TagsDeleteRequest {
	return v.value
}

func (v *NullableApiV2TagsDeleteRequest) Set(val *ApiV2TagsDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2TagsDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2TagsDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2TagsDeleteRequest(val *ApiV2TagsDeleteRequest) *NullableApiV2TagsDeleteRequest {
	return &NullableApiV2TagsDeleteRequest{value: val, isSet: true}
}

func (v NullableApiV2TagsDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2TagsDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

