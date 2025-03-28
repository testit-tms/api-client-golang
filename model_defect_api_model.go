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

// checks if the DefectApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefectApiModel{}

// DefectApiModel struct for DefectApiModel
type DefectApiModel struct {
	// Link to created issue
	ExternalUrl string `json:"externalUrl"`
}

type _DefectApiModel DefectApiModel

// NewDefectApiModel instantiates a new DefectApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefectApiModel(externalUrl string) *DefectApiModel {
	this := DefectApiModel{}
	this.ExternalUrl = externalUrl
	return &this
}

// NewDefectApiModelWithDefaults instantiates a new DefectApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefectApiModelWithDefaults() *DefectApiModel {
	this := DefectApiModel{}
	return &this
}

// GetExternalUrl returns the ExternalUrl field value
func (o *DefectApiModel) GetExternalUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalUrl
}

// GetExternalUrlOk returns a tuple with the ExternalUrl field value
// and a boolean to check if the value has been set.
func (o *DefectApiModel) GetExternalUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUrl, true
}

// SetExternalUrl sets field value
func (o *DefectApiModel) SetExternalUrl(v string) {
	o.ExternalUrl = v
}

func (o DefectApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefectApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["externalUrl"] = o.ExternalUrl
	return toSerialize, nil
}

func (o *DefectApiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"externalUrl",
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

	varDefectApiModel := _DefectApiModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDefectApiModel)

	if err != nil {
		return err
	}

	*o = DefectApiModel(varDefectApiModel)

	return err
}

type NullableDefectApiModel struct {
	value *DefectApiModel
	isSet bool
}

func (v NullableDefectApiModel) Get() *DefectApiModel {
	return v.value
}

func (v *NullableDefectApiModel) Set(val *DefectApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDefectApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDefectApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefectApiModel(val *DefectApiModel) *NullableDefectApiModel {
	return &NullableDefectApiModel{value: val, isSet: true}
}

func (v NullableDefectApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefectApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


