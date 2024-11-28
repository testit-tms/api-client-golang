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

// checks if the ConfigurationShort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationShort{}

// ConfigurationShort struct for ConfigurationShort
type ConfigurationShort struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type _ConfigurationShort ConfigurationShort

// NewConfigurationShort instantiates a new ConfigurationShort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationShort(id string, name string) *ConfigurationShort {
	this := ConfigurationShort{}
	this.Id = id
	this.Name = name
	return &this
}

// NewConfigurationShortWithDefaults instantiates a new ConfigurationShort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationShortWithDefaults() *ConfigurationShort {
	this := ConfigurationShort{}
	return &this
}

// GetId returns the Id field value
func (o *ConfigurationShort) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConfigurationShort) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConfigurationShort) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ConfigurationShort) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfigurationShort) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConfigurationShort) SetName(v string) {
	o.Name = v
}

func (o ConfigurationShort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationShort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *ConfigurationShort) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varConfigurationShort := _ConfigurationShort{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConfigurationShort)

	if err != nil {
		return err
	}

	*o = ConfigurationShort(varConfigurationShort)

	return err
}

type NullableConfigurationShort struct {
	value *ConfigurationShort
	isSet bool
}

func (v NullableConfigurationShort) Get() *ConfigurationShort {
	return v.value
}

func (v *NullableConfigurationShort) Set(val *ConfigurationShort) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationShort) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationShort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationShort(val *ConfigurationShort) *NullableConfigurationShort {
	return &NullableConfigurationShort{value: val, isSet: true}
}

func (v NullableConfigurationShort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationShort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

