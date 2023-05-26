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

// checks if the ProjectExportQueryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectExportQueryModel{}

// ProjectExportQueryModel struct for ProjectExportQueryModel
type ProjectExportQueryModel struct {
	// Specifies the IDs of the sections you want to export.<br />  Use this parameter if you want to export certain parts of the project.<br />  In this parameter, \"<b>string</b>\" values are IDs of the test library sections.
	SectionIds []string `json:"sectionIds,omitempty"`
	// Specifies the work items you want to export from a project.<br />  Use this parameter if you want to export certain work items.<br />  In this parameter, \"<b>string</b>\" values are IDs of the work items.
	WorkItemIds []string `json:"workItemIds,omitempty"`
}

// NewProjectExportQueryModel instantiates a new ProjectExportQueryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectExportQueryModel() *ProjectExportQueryModel {
	this := ProjectExportQueryModel{}
	return &this
}

// NewProjectExportQueryModelWithDefaults instantiates a new ProjectExportQueryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectExportQueryModelWithDefaults() *ProjectExportQueryModel {
	this := ProjectExportQueryModel{}
	return &this
}

// GetSectionIds returns the SectionIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectExportQueryModel) GetSectionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SectionIds
}

// GetSectionIdsOk returns a tuple with the SectionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectExportQueryModel) GetSectionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SectionIds) {
		return nil, false
	}
	return o.SectionIds, true
}

// HasSectionIds returns a boolean if a field has been set.
func (o *ProjectExportQueryModel) HasSectionIds() bool {
	if o != nil && IsNil(o.SectionIds) {
		return true
	}

	return false
}

// SetSectionIds gets a reference to the given []string and assigns it to the SectionIds field.
func (o *ProjectExportQueryModel) SetSectionIds(v []string) {
	o.SectionIds = v
}

// GetWorkItemIds returns the WorkItemIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectExportQueryModel) GetWorkItemIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.WorkItemIds
}

// GetWorkItemIdsOk returns a tuple with the WorkItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectExportQueryModel) GetWorkItemIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.WorkItemIds) {
		return nil, false
	}
	return o.WorkItemIds, true
}

// HasWorkItemIds returns a boolean if a field has been set.
func (o *ProjectExportQueryModel) HasWorkItemIds() bool {
	if o != nil && IsNil(o.WorkItemIds) {
		return true
	}

	return false
}

// SetWorkItemIds gets a reference to the given []string and assigns it to the WorkItemIds field.
func (o *ProjectExportQueryModel) SetWorkItemIds(v []string) {
	o.WorkItemIds = v
}

func (o ProjectExportQueryModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectExportQueryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SectionIds != nil {
		toSerialize["sectionIds"] = o.SectionIds
	}
	if o.WorkItemIds != nil {
		toSerialize["workItemIds"] = o.WorkItemIds
	}
	return toSerialize, nil
}

type NullableProjectExportQueryModel struct {
	value *ProjectExportQueryModel
	isSet bool
}

func (v NullableProjectExportQueryModel) Get() *ProjectExportQueryModel {
	return v.value
}

func (v *NullableProjectExportQueryModel) Set(val *ProjectExportQueryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectExportQueryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectExportQueryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectExportQueryModel(val *ProjectExportQueryModel) *NullableProjectExportQueryModel {
	return &NullableProjectExportQueryModel{value: val, isSet: true}
}

func (v NullableProjectExportQueryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectExportQueryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

