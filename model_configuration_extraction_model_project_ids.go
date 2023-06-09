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

// checks if the ConfigurationExtractionModelProjectIds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationExtractionModelProjectIds{}

// ConfigurationExtractionModelProjectIds Extraction parameters for projects
type ConfigurationExtractionModelProjectIds struct {
	Include []string `json:"include,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}

// NewConfigurationExtractionModelProjectIds instantiates a new ConfigurationExtractionModelProjectIds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationExtractionModelProjectIds() *ConfigurationExtractionModelProjectIds {
	this := ConfigurationExtractionModelProjectIds{}
	return &this
}

// NewConfigurationExtractionModelProjectIdsWithDefaults instantiates a new ConfigurationExtractionModelProjectIds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationExtractionModelProjectIdsWithDefaults() *ConfigurationExtractionModelProjectIds {
	this := ConfigurationExtractionModelProjectIds{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationExtractionModelProjectIds) GetInclude() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationExtractionModelProjectIds) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *ConfigurationExtractionModelProjectIds) HasInclude() bool {
	if o != nil && IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *ConfigurationExtractionModelProjectIds) SetInclude(v []string) {
	o.Include = v
}

// GetExclude returns the Exclude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationExtractionModelProjectIds) GetExclude() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationExtractionModelProjectIds) GetExcludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *ConfigurationExtractionModelProjectIds) HasExclude() bool {
	if o != nil && IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *ConfigurationExtractionModelProjectIds) SetExclude(v []string) {
	o.Exclude = v
}

func (o ConfigurationExtractionModelProjectIds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationExtractionModelProjectIds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}
	if o.Exclude != nil {
		toSerialize["exclude"] = o.Exclude
	}
	return toSerialize, nil
}

type NullableConfigurationExtractionModelProjectIds struct {
	value *ConfigurationExtractionModelProjectIds
	isSet bool
}

func (v NullableConfigurationExtractionModelProjectIds) Get() *ConfigurationExtractionModelProjectIds {
	return v.value
}

func (v *NullableConfigurationExtractionModelProjectIds) Set(val *ConfigurationExtractionModelProjectIds) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationExtractionModelProjectIds) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationExtractionModelProjectIds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationExtractionModelProjectIds(val *ConfigurationExtractionModelProjectIds) *NullableConfigurationExtractionModelProjectIds {
	return &NullableConfigurationExtractionModelProjectIds{value: val, isSet: true}
}

func (v NullableConfigurationExtractionModelProjectIds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationExtractionModelProjectIds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


