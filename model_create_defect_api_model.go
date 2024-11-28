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

// checks if the CreateDefectApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDefectApiModel{}

// CreateDefectApiModel struct for CreateDefectApiModel
type CreateDefectApiModel struct {
	// Linked test result IDs
	TestResultIds []string `json:"testResultIds"`
	// External form definition
	Form ExternalFormCreateModel `json:"form"`
}

type _CreateDefectApiModel CreateDefectApiModel

// NewCreateDefectApiModel instantiates a new CreateDefectApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDefectApiModel(testResultIds []string, form ExternalFormCreateModel) *CreateDefectApiModel {
	this := CreateDefectApiModel{}
	this.TestResultIds = testResultIds
	this.Form = form
	return &this
}

// NewCreateDefectApiModelWithDefaults instantiates a new CreateDefectApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDefectApiModelWithDefaults() *CreateDefectApiModel {
	this := CreateDefectApiModel{}
	return &this
}

// GetTestResultIds returns the TestResultIds field value
func (o *CreateDefectApiModel) GetTestResultIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TestResultIds
}

// GetTestResultIdsOk returns a tuple with the TestResultIds field value
// and a boolean to check if the value has been set.
func (o *CreateDefectApiModel) GetTestResultIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestResultIds, true
}

// SetTestResultIds sets field value
func (o *CreateDefectApiModel) SetTestResultIds(v []string) {
	o.TestResultIds = v
}

// GetForm returns the Form field value
func (o *CreateDefectApiModel) GetForm() ExternalFormCreateModel {
	if o == nil {
		var ret ExternalFormCreateModel
		return ret
	}

	return o.Form
}

// GetFormOk returns a tuple with the Form field value
// and a boolean to check if the value has been set.
func (o *CreateDefectApiModel) GetFormOk() (*ExternalFormCreateModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Form, true
}

// SetForm sets field value
func (o *CreateDefectApiModel) SetForm(v ExternalFormCreateModel) {
	o.Form = v
}

func (o CreateDefectApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDefectApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["testResultIds"] = o.TestResultIds
	toSerialize["form"] = o.Form
	return toSerialize, nil
}

func (o *CreateDefectApiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"testResultIds",
		"form",
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

	varCreateDefectApiModel := _CreateDefectApiModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateDefectApiModel)

	if err != nil {
		return err
	}

	*o = CreateDefectApiModel(varCreateDefectApiModel)

	return err
}

type NullableCreateDefectApiModel struct {
	value *CreateDefectApiModel
	isSet bool
}

func (v NullableCreateDefectApiModel) Get() *CreateDefectApiModel {
	return v.value
}

func (v *NullableCreateDefectApiModel) Set(val *CreateDefectApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDefectApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDefectApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDefectApiModel(val *CreateDefectApiModel) *NullableCreateDefectApiModel {
	return &NullableCreateDefectApiModel{value: val, isSet: true}
}

func (v NullableCreateDefectApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDefectApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


