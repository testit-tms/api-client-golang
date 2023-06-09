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

// checks if the ConfigurationSelectModelExtractionModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationSelectModelExtractionModel{}

// ConfigurationSelectModelExtractionModel Rules for configurations extraction
type ConfigurationSelectModelExtractionModel struct {
	Ids NullableConfigurationExtractionModelIds `json:"ids,omitempty"`
	ProjectIds NullableConfigurationExtractionModelProjectIds `json:"projectIds,omitempty"`
}

// NewConfigurationSelectModelExtractionModel instantiates a new ConfigurationSelectModelExtractionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationSelectModelExtractionModel() *ConfigurationSelectModelExtractionModel {
	this := ConfigurationSelectModelExtractionModel{}
	return &this
}

// NewConfigurationSelectModelExtractionModelWithDefaults instantiates a new ConfigurationSelectModelExtractionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationSelectModelExtractionModelWithDefaults() *ConfigurationSelectModelExtractionModel {
	this := ConfigurationSelectModelExtractionModel{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationSelectModelExtractionModel) GetIds() ConfigurationExtractionModelIds {
	if o == nil || IsNil(o.Ids.Get()) {
		var ret ConfigurationExtractionModelIds
		return ret
	}
	return *o.Ids.Get()
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationSelectModelExtractionModel) GetIdsOk() (*ConfigurationExtractionModelIds, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids.Get(), o.Ids.IsSet()
}

// HasIds returns a boolean if a field has been set.
func (o *ConfigurationSelectModelExtractionModel) HasIds() bool {
	if o != nil && o.Ids.IsSet() {
		return true
	}

	return false
}

// SetIds gets a reference to the given NullableConfigurationExtractionModelIds and assigns it to the Ids field.
func (o *ConfigurationSelectModelExtractionModel) SetIds(v ConfigurationExtractionModelIds) {
	o.Ids.Set(&v)
}
// SetIdsNil sets the value for Ids to be an explicit nil
func (o *ConfigurationSelectModelExtractionModel) SetIdsNil() {
	o.Ids.Set(nil)
}

// UnsetIds ensures that no value is present for Ids, not even an explicit nil
func (o *ConfigurationSelectModelExtractionModel) UnsetIds() {
	o.Ids.Unset()
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationSelectModelExtractionModel) GetProjectIds() ConfigurationExtractionModelProjectIds {
	if o == nil || IsNil(o.ProjectIds.Get()) {
		var ret ConfigurationExtractionModelProjectIds
		return ret
	}
	return *o.ProjectIds.Get()
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationSelectModelExtractionModel) GetProjectIdsOk() (*ConfigurationExtractionModelProjectIds, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectIds.Get(), o.ProjectIds.IsSet()
}

// HasProjectIds returns a boolean if a field has been set.
func (o *ConfigurationSelectModelExtractionModel) HasProjectIds() bool {
	if o != nil && o.ProjectIds.IsSet() {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given NullableConfigurationExtractionModelProjectIds and assigns it to the ProjectIds field.
func (o *ConfigurationSelectModelExtractionModel) SetProjectIds(v ConfigurationExtractionModelProjectIds) {
	o.ProjectIds.Set(&v)
}
// SetProjectIdsNil sets the value for ProjectIds to be an explicit nil
func (o *ConfigurationSelectModelExtractionModel) SetProjectIdsNil() {
	o.ProjectIds.Set(nil)
}

// UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
func (o *ConfigurationSelectModelExtractionModel) UnsetProjectIds() {
	o.ProjectIds.Unset()
}

func (o ConfigurationSelectModelExtractionModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationSelectModelExtractionModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids.IsSet() {
		toSerialize["ids"] = o.Ids.Get()
	}
	if o.ProjectIds.IsSet() {
		toSerialize["projectIds"] = o.ProjectIds.Get()
	}
	return toSerialize, nil
}

type NullableConfigurationSelectModelExtractionModel struct {
	value *ConfigurationSelectModelExtractionModel
	isSet bool
}

func (v NullableConfigurationSelectModelExtractionModel) Get() *ConfigurationSelectModelExtractionModel {
	return v.value
}

func (v *NullableConfigurationSelectModelExtractionModel) Set(val *ConfigurationSelectModelExtractionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationSelectModelExtractionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationSelectModelExtractionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationSelectModelExtractionModel(val *ConfigurationSelectModelExtractionModel) *NullableConfigurationSelectModelExtractionModel {
	return &NullableConfigurationSelectModelExtractionModel{value: val, isSet: true}
}

func (v NullableConfigurationSelectModelExtractionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationSelectModelExtractionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


