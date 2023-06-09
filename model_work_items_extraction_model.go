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

// checks if the WorkItemsExtractionModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemsExtractionModel{}

// WorkItemsExtractionModel Rules for different level entities inclusion/exclusion
type WorkItemsExtractionModel struct {
	Ids NullableWorkItemsExtractionModelIds `json:"ids,omitempty"`
	SectionIds NullableWorkItemsExtractionModelSectionIds `json:"sectionIds,omitempty"`
	ProjectIds NullableConfigurationExtractionModelProjectIds `json:"projectIds,omitempty"`
}

// NewWorkItemsExtractionModel instantiates a new WorkItemsExtractionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemsExtractionModel() *WorkItemsExtractionModel {
	this := WorkItemsExtractionModel{}
	return &this
}

// NewWorkItemsExtractionModelWithDefaults instantiates a new WorkItemsExtractionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemsExtractionModelWithDefaults() *WorkItemsExtractionModel {
	this := WorkItemsExtractionModel{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemsExtractionModel) GetIds() WorkItemsExtractionModelIds {
	if o == nil || IsNil(o.Ids.Get()) {
		var ret WorkItemsExtractionModelIds
		return ret
	}
	return *o.Ids.Get()
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemsExtractionModel) GetIdsOk() (*WorkItemsExtractionModelIds, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids.Get(), o.Ids.IsSet()
}

// HasIds returns a boolean if a field has been set.
func (o *WorkItemsExtractionModel) HasIds() bool {
	if o != nil && o.Ids.IsSet() {
		return true
	}

	return false
}

// SetIds gets a reference to the given NullableWorkItemsExtractionModelIds and assigns it to the Ids field.
func (o *WorkItemsExtractionModel) SetIds(v WorkItemsExtractionModelIds) {
	o.Ids.Set(&v)
}
// SetIdsNil sets the value for Ids to be an explicit nil
func (o *WorkItemsExtractionModel) SetIdsNil() {
	o.Ids.Set(nil)
}

// UnsetIds ensures that no value is present for Ids, not even an explicit nil
func (o *WorkItemsExtractionModel) UnsetIds() {
	o.Ids.Unset()
}

// GetSectionIds returns the SectionIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemsExtractionModel) GetSectionIds() WorkItemsExtractionModelSectionIds {
	if o == nil || IsNil(o.SectionIds.Get()) {
		var ret WorkItemsExtractionModelSectionIds
		return ret
	}
	return *o.SectionIds.Get()
}

// GetSectionIdsOk returns a tuple with the SectionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemsExtractionModel) GetSectionIdsOk() (*WorkItemsExtractionModelSectionIds, bool) {
	if o == nil {
		return nil, false
	}
	return o.SectionIds.Get(), o.SectionIds.IsSet()
}

// HasSectionIds returns a boolean if a field has been set.
func (o *WorkItemsExtractionModel) HasSectionIds() bool {
	if o != nil && o.SectionIds.IsSet() {
		return true
	}

	return false
}

// SetSectionIds gets a reference to the given NullableWorkItemsExtractionModelSectionIds and assigns it to the SectionIds field.
func (o *WorkItemsExtractionModel) SetSectionIds(v WorkItemsExtractionModelSectionIds) {
	o.SectionIds.Set(&v)
}
// SetSectionIdsNil sets the value for SectionIds to be an explicit nil
func (o *WorkItemsExtractionModel) SetSectionIdsNil() {
	o.SectionIds.Set(nil)
}

// UnsetSectionIds ensures that no value is present for SectionIds, not even an explicit nil
func (o *WorkItemsExtractionModel) UnsetSectionIds() {
	o.SectionIds.Unset()
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemsExtractionModel) GetProjectIds() ConfigurationExtractionModelProjectIds {
	if o == nil || IsNil(o.ProjectIds.Get()) {
		var ret ConfigurationExtractionModelProjectIds
		return ret
	}
	return *o.ProjectIds.Get()
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemsExtractionModel) GetProjectIdsOk() (*ConfigurationExtractionModelProjectIds, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectIds.Get(), o.ProjectIds.IsSet()
}

// HasProjectIds returns a boolean if a field has been set.
func (o *WorkItemsExtractionModel) HasProjectIds() bool {
	if o != nil && o.ProjectIds.IsSet() {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given NullableConfigurationExtractionModelProjectIds and assigns it to the ProjectIds field.
func (o *WorkItemsExtractionModel) SetProjectIds(v ConfigurationExtractionModelProjectIds) {
	o.ProjectIds.Set(&v)
}
// SetProjectIdsNil sets the value for ProjectIds to be an explicit nil
func (o *WorkItemsExtractionModel) SetProjectIdsNil() {
	o.ProjectIds.Set(nil)
}

// UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
func (o *WorkItemsExtractionModel) UnsetProjectIds() {
	o.ProjectIds.Unset()
}

func (o WorkItemsExtractionModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemsExtractionModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids.IsSet() {
		toSerialize["ids"] = o.Ids.Get()
	}
	if o.SectionIds.IsSet() {
		toSerialize["sectionIds"] = o.SectionIds.Get()
	}
	if o.ProjectIds.IsSet() {
		toSerialize["projectIds"] = o.ProjectIds.Get()
	}
	return toSerialize, nil
}

type NullableWorkItemsExtractionModel struct {
	value *WorkItemsExtractionModel
	isSet bool
}

func (v NullableWorkItemsExtractionModel) Get() *WorkItemsExtractionModel {
	return v.value
}

func (v *NullableWorkItemsExtractionModel) Set(val *WorkItemsExtractionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemsExtractionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemsExtractionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemsExtractionModel(val *WorkItemsExtractionModel) *NullableWorkItemsExtractionModel {
	return &NullableWorkItemsExtractionModel{value: val, isSet: true}
}

func (v NullableWorkItemsExtractionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemsExtractionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


