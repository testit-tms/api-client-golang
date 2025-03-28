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

// checks if the WorkItemLinkUrlFilterApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemLinkUrlFilterApiModel{}

// WorkItemLinkUrlFilterApiModel struct for WorkItemLinkUrlFilterApiModel
type WorkItemLinkUrlFilterApiModel struct {
	Types []WorkItemEntityTypes `json:"types,omitempty"`
	SearchUrl NullableString `json:"searchUrl,omitempty"`
}

// NewWorkItemLinkUrlFilterApiModel instantiates a new WorkItemLinkUrlFilterApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemLinkUrlFilterApiModel() *WorkItemLinkUrlFilterApiModel {
	this := WorkItemLinkUrlFilterApiModel{}
	return &this
}

// NewWorkItemLinkUrlFilterApiModelWithDefaults instantiates a new WorkItemLinkUrlFilterApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemLinkUrlFilterApiModelWithDefaults() *WorkItemLinkUrlFilterApiModel {
	this := WorkItemLinkUrlFilterApiModel{}
	return &this
}

// GetTypes returns the Types field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemLinkUrlFilterApiModel) GetTypes() []WorkItemEntityTypes {
	if o == nil {
		var ret []WorkItemEntityTypes
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemLinkUrlFilterApiModel) GetTypesOk() ([]WorkItemEntityTypes, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *WorkItemLinkUrlFilterApiModel) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []WorkItemEntityTypes and assigns it to the Types field.
func (o *WorkItemLinkUrlFilterApiModel) SetTypes(v []WorkItemEntityTypes) {
	o.Types = v
}

// GetSearchUrl returns the SearchUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemLinkUrlFilterApiModel) GetSearchUrl() string {
	if o == nil || IsNil(o.SearchUrl.Get()) {
		var ret string
		return ret
	}
	return *o.SearchUrl.Get()
}

// GetSearchUrlOk returns a tuple with the SearchUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemLinkUrlFilterApiModel) GetSearchUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchUrl.Get(), o.SearchUrl.IsSet()
}

// HasSearchUrl returns a boolean if a field has been set.
func (o *WorkItemLinkUrlFilterApiModel) HasSearchUrl() bool {
	if o != nil && o.SearchUrl.IsSet() {
		return true
	}

	return false
}

// SetSearchUrl gets a reference to the given NullableString and assigns it to the SearchUrl field.
func (o *WorkItemLinkUrlFilterApiModel) SetSearchUrl(v string) {
	o.SearchUrl.Set(&v)
}
// SetSearchUrlNil sets the value for SearchUrl to be an explicit nil
func (o *WorkItemLinkUrlFilterApiModel) SetSearchUrlNil() {
	o.SearchUrl.Set(nil)
}

// UnsetSearchUrl ensures that no value is present for SearchUrl, not even an explicit nil
func (o *WorkItemLinkUrlFilterApiModel) UnsetSearchUrl() {
	o.SearchUrl.Unset()
}

func (o WorkItemLinkUrlFilterApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemLinkUrlFilterApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	if o.SearchUrl.IsSet() {
		toSerialize["searchUrl"] = o.SearchUrl.Get()
	}
	return toSerialize, nil
}

type NullableWorkItemLinkUrlFilterApiModel struct {
	value *WorkItemLinkUrlFilterApiModel
	isSet bool
}

func (v NullableWorkItemLinkUrlFilterApiModel) Get() *WorkItemLinkUrlFilterApiModel {
	return v.value
}

func (v *NullableWorkItemLinkUrlFilterApiModel) Set(val *WorkItemLinkUrlFilterApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemLinkUrlFilterApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemLinkUrlFilterApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemLinkUrlFilterApiModel(val *WorkItemLinkUrlFilterApiModel) *NullableWorkItemLinkUrlFilterApiModel {
	return &NullableWorkItemLinkUrlFilterApiModel{value: val, isSet: true}
}

func (v NullableWorkItemLinkUrlFilterApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemLinkUrlFilterApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


