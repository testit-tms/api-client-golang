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

// checks if the ApiV2TestRunsSearchPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV2TestRunsSearchPostRequest{}

// ApiV2TestRunsSearchPostRequest struct for ApiV2TestRunsSearchPostRequest
type ApiV2TestRunsSearchPostRequest struct {
	// Specifies a test run project IDs to search for
	ProjectIds []string `json:"projectIds,omitempty"`
	// Specifies a test run states to search for
	States []TestRunState `json:"states,omitempty"`
	CreatedDate NullableTestRunFilterModelCreatedDate `json:"createdDate,omitempty"`
	// Specifies a test run last editor IDs to search for
	ModifiedByIds []string `json:"modifiedByIds,omitempty"`
	// Specifies a test run deleted status to search for
	IsDeleted NullableBool `json:"isDeleted,omitempty"`
}

// NewApiV2TestRunsSearchPostRequest instantiates a new ApiV2TestRunsSearchPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2TestRunsSearchPostRequest() *ApiV2TestRunsSearchPostRequest {
	this := ApiV2TestRunsSearchPostRequest{}
	return &this
}

// NewApiV2TestRunsSearchPostRequestWithDefaults instantiates a new ApiV2TestRunsSearchPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2TestRunsSearchPostRequestWithDefaults() *ApiV2TestRunsSearchPostRequest {
	this := ApiV2TestRunsSearchPostRequest{}
	return &this
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestRunsSearchPostRequest) GetProjectIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestRunsSearchPostRequest) GetProjectIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectIds) {
		return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *ApiV2TestRunsSearchPostRequest) HasProjectIds() bool {
	if o != nil && IsNil(o.ProjectIds) {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []string and assigns it to the ProjectIds field.
func (o *ApiV2TestRunsSearchPostRequest) SetProjectIds(v []string) {
	o.ProjectIds = v
}

// GetStates returns the States field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestRunsSearchPostRequest) GetStates() []TestRunState {
	if o == nil {
		var ret []TestRunState
		return ret
	}
	return o.States
}

// GetStatesOk returns a tuple with the States field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestRunsSearchPostRequest) GetStatesOk() ([]TestRunState, bool) {
	if o == nil || IsNil(o.States) {
		return nil, false
	}
	return o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *ApiV2TestRunsSearchPostRequest) HasStates() bool {
	if o != nil && IsNil(o.States) {
		return true
	}

	return false
}

// SetStates gets a reference to the given []TestRunState and assigns it to the States field.
func (o *ApiV2TestRunsSearchPostRequest) SetStates(v []TestRunState) {
	o.States = v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestRunsSearchPostRequest) GetCreatedDate() TestRunFilterModelCreatedDate {
	if o == nil || IsNil(o.CreatedDate.Get()) {
		var ret TestRunFilterModelCreatedDate
		return ret
	}
	return *o.CreatedDate.Get()
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestRunsSearchPostRequest) GetCreatedDateOk() (*TestRunFilterModelCreatedDate, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDate.Get(), o.CreatedDate.IsSet()
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ApiV2TestRunsSearchPostRequest) HasCreatedDate() bool {
	if o != nil && o.CreatedDate.IsSet() {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given NullableTestRunFilterModelCreatedDate and assigns it to the CreatedDate field.
func (o *ApiV2TestRunsSearchPostRequest) SetCreatedDate(v TestRunFilterModelCreatedDate) {
	o.CreatedDate.Set(&v)
}
// SetCreatedDateNil sets the value for CreatedDate to be an explicit nil
func (o *ApiV2TestRunsSearchPostRequest) SetCreatedDateNil() {
	o.CreatedDate.Set(nil)
}

// UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
func (o *ApiV2TestRunsSearchPostRequest) UnsetCreatedDate() {
	o.CreatedDate.Unset()
}

// GetModifiedByIds returns the ModifiedByIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestRunsSearchPostRequest) GetModifiedByIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ModifiedByIds
}

// GetModifiedByIdsOk returns a tuple with the ModifiedByIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestRunsSearchPostRequest) GetModifiedByIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ModifiedByIds) {
		return nil, false
	}
	return o.ModifiedByIds, true
}

// HasModifiedByIds returns a boolean if a field has been set.
func (o *ApiV2TestRunsSearchPostRequest) HasModifiedByIds() bool {
	if o != nil && IsNil(o.ModifiedByIds) {
		return true
	}

	return false
}

// SetModifiedByIds gets a reference to the given []string and assigns it to the ModifiedByIds field.
func (o *ApiV2TestRunsSearchPostRequest) SetModifiedByIds(v []string) {
	o.ModifiedByIds = v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2TestRunsSearchPostRequest) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDeleted.Get()
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2TestRunsSearchPostRequest) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDeleted.Get(), o.IsDeleted.IsSet()
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *ApiV2TestRunsSearchPostRequest) HasIsDeleted() bool {
	if o != nil && o.IsDeleted.IsSet() {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given NullableBool and assigns it to the IsDeleted field.
func (o *ApiV2TestRunsSearchPostRequest) SetIsDeleted(v bool) {
	o.IsDeleted.Set(&v)
}
// SetIsDeletedNil sets the value for IsDeleted to be an explicit nil
func (o *ApiV2TestRunsSearchPostRequest) SetIsDeletedNil() {
	o.IsDeleted.Set(nil)
}

// UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
func (o *ApiV2TestRunsSearchPostRequest) UnsetIsDeleted() {
	o.IsDeleted.Unset()
}

func (o ApiV2TestRunsSearchPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV2TestRunsSearchPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectIds != nil {
		toSerialize["projectIds"] = o.ProjectIds
	}
	if o.States != nil {
		toSerialize["states"] = o.States
	}
	if o.CreatedDate.IsSet() {
		toSerialize["createdDate"] = o.CreatedDate.Get()
	}
	if o.ModifiedByIds != nil {
		toSerialize["modifiedByIds"] = o.ModifiedByIds
	}
	if o.IsDeleted.IsSet() {
		toSerialize["isDeleted"] = o.IsDeleted.Get()
	}
	return toSerialize, nil
}

type NullableApiV2TestRunsSearchPostRequest struct {
	value *ApiV2TestRunsSearchPostRequest
	isSet bool
}

func (v NullableApiV2TestRunsSearchPostRequest) Get() *ApiV2TestRunsSearchPostRequest {
	return v.value
}

func (v *NullableApiV2TestRunsSearchPostRequest) Set(val *ApiV2TestRunsSearchPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2TestRunsSearchPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2TestRunsSearchPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2TestRunsSearchPostRequest(val *ApiV2TestRunsSearchPostRequest) *NullableApiV2TestRunsSearchPostRequest {
	return &NullableApiV2TestRunsSearchPostRequest{value: val, isSet: true}
}

func (v NullableApiV2TestRunsSearchPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2TestRunsSearchPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


