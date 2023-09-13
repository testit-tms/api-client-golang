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

// checks if the ApiV2ProjectsIdTestPlansDeleteBulkPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV2ProjectsIdTestPlansDeleteBulkPostRequest{}

// ApiV2ProjectsIdTestPlansDeleteBulkPostRequest struct for ApiV2ProjectsIdTestPlansDeleteBulkPostRequest
type ApiV2ProjectsIdTestPlansDeleteBulkPostRequest struct {
	Filter NullableProjectTestPlansFilterModel `json:"filter,omitempty"`
	ExtractionModel NullableTestPlanExtractionModel `json:"extractionModel,omitempty"`
}

// NewApiV2ProjectsIdTestPlansDeleteBulkPostRequest instantiates a new ApiV2ProjectsIdTestPlansDeleteBulkPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2ProjectsIdTestPlansDeleteBulkPostRequest() *ApiV2ProjectsIdTestPlansDeleteBulkPostRequest {
	this := ApiV2ProjectsIdTestPlansDeleteBulkPostRequest{}
	return &this
}

// NewApiV2ProjectsIdTestPlansDeleteBulkPostRequestWithDefaults instantiates a new ApiV2ProjectsIdTestPlansDeleteBulkPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2ProjectsIdTestPlansDeleteBulkPostRequestWithDefaults() *ApiV2ProjectsIdTestPlansDeleteBulkPostRequest {
	this := ApiV2ProjectsIdTestPlansDeleteBulkPostRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsIdTestPlansDeleteBulkPostRequest) GetFilter() ProjectTestPlansFilterModel {
	if o == nil || IsNil(o.Filter.Get()) {
		var ret ProjectTestPlansFilterModel
		return ret
	}
	return *o.Filter.Get()
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsIdTestPlansDeleteBulkPostRequest) GetFilterOk() (*ProjectTestPlansFilterModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter.Get(), o.Filter.IsSet()
}

// HasFilter returns a boolean if a field has been set.
func (o *ApiV2ProjectsIdTestPlansDeleteBulkPostRequest) HasFilter() bool {
	if o != nil && o.Filter.IsSet() {
		return true
	}

	return false
}

// SetFilter gets a reference to the given NullableProjectTestPlansFilterModel and assigns it to the Filter field.
func (o *ApiV2ProjectsIdTestPlansDeleteBulkPostRequest) SetFilter(v ProjectTestPlansFilterModel) {
	o.Filter.Set(&v)
}
// SetFilterNil sets the value for Filter to be an explicit nil
func (o *ApiV2ProjectsIdTestPlansDeleteBulkPostRequest) SetFilterNil() {
	o.Filter.Set(nil)
}

// UnsetFilter ensures that no value is present for Filter, not even an explicit nil
func (o *ApiV2ProjectsIdTestPlansDeleteBulkPostRequest) UnsetFilter() {
	o.Filter.Unset()
}

// GetExtractionModel returns the ExtractionModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2ProjectsIdTestPlansDeleteBulkPostRequest) GetExtractionModel() TestPlanExtractionModel {
	if o == nil || IsNil(o.ExtractionModel.Get()) {
		var ret TestPlanExtractionModel
		return ret
	}
	return *o.ExtractionModel.Get()
}

// GetExtractionModelOk returns a tuple with the ExtractionModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2ProjectsIdTestPlansDeleteBulkPostRequest) GetExtractionModelOk() (*TestPlanExtractionModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtractionModel.Get(), o.ExtractionModel.IsSet()
}

// HasExtractionModel returns a boolean if a field has been set.
func (o *ApiV2ProjectsIdTestPlansDeleteBulkPostRequest) HasExtractionModel() bool {
	if o != nil && o.ExtractionModel.IsSet() {
		return true
	}

	return false
}

// SetExtractionModel gets a reference to the given NullableTestPlanExtractionModel and assigns it to the ExtractionModel field.
func (o *ApiV2ProjectsIdTestPlansDeleteBulkPostRequest) SetExtractionModel(v TestPlanExtractionModel) {
	o.ExtractionModel.Set(&v)
}
// SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil
func (o *ApiV2ProjectsIdTestPlansDeleteBulkPostRequest) SetExtractionModelNil() {
	o.ExtractionModel.Set(nil)
}

// UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil
func (o *ApiV2ProjectsIdTestPlansDeleteBulkPostRequest) UnsetExtractionModel() {
	o.ExtractionModel.Unset()
}

func (o ApiV2ProjectsIdTestPlansDeleteBulkPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV2ProjectsIdTestPlansDeleteBulkPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter.IsSet() {
		toSerialize["filter"] = o.Filter.Get()
	}
	if o.ExtractionModel.IsSet() {
		toSerialize["extractionModel"] = o.ExtractionModel.Get()
	}
	return toSerialize, nil
}

type NullableApiV2ProjectsIdTestPlansDeleteBulkPostRequest struct {
	value *ApiV2ProjectsIdTestPlansDeleteBulkPostRequest
	isSet bool
}

func (v NullableApiV2ProjectsIdTestPlansDeleteBulkPostRequest) Get() *ApiV2ProjectsIdTestPlansDeleteBulkPostRequest {
	return v.value
}

func (v *NullableApiV2ProjectsIdTestPlansDeleteBulkPostRequest) Set(val *ApiV2ProjectsIdTestPlansDeleteBulkPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2ProjectsIdTestPlansDeleteBulkPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2ProjectsIdTestPlansDeleteBulkPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2ProjectsIdTestPlansDeleteBulkPostRequest(val *ApiV2ProjectsIdTestPlansDeleteBulkPostRequest) *NullableApiV2ProjectsIdTestPlansDeleteBulkPostRequest {
	return &NullableApiV2ProjectsIdTestPlansDeleteBulkPostRequest{value: val, isSet: true}
}

func (v NullableApiV2ProjectsIdTestPlansDeleteBulkPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2ProjectsIdTestPlansDeleteBulkPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

