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

// checks if the ExternalFormCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalFormCreateModel{}

// ExternalFormCreateModel struct for ExternalFormCreateModel
type ExternalFormCreateModel struct {
	PossibleValues map[string][]ExternalFormAllowedValueModel `json:"possibleValues"`
	Fields []ExternalFormFieldModel `json:"fields"`
	Links []ExternalFormLinkModel `json:"links"`
	Values map[string]interface{} `json:"values"`
}

type _ExternalFormCreateModel ExternalFormCreateModel

// NewExternalFormCreateModel instantiates a new ExternalFormCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalFormCreateModel(possibleValues map[string][]ExternalFormAllowedValueModel, fields []ExternalFormFieldModel, links []ExternalFormLinkModel, values map[string]interface{}) *ExternalFormCreateModel {
	this := ExternalFormCreateModel{}
	this.PossibleValues = possibleValues
	this.Fields = fields
	this.Links = links
	this.Values = values
	return &this
}

// NewExternalFormCreateModelWithDefaults instantiates a new ExternalFormCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalFormCreateModelWithDefaults() *ExternalFormCreateModel {
	this := ExternalFormCreateModel{}
	return &this
}

// GetPossibleValues returns the PossibleValues field value
func (o *ExternalFormCreateModel) GetPossibleValues() map[string][]ExternalFormAllowedValueModel {
	if o == nil {
		var ret map[string][]ExternalFormAllowedValueModel
		return ret
	}

	return o.PossibleValues
}

// GetPossibleValuesOk returns a tuple with the PossibleValues field value
// and a boolean to check if the value has been set.
func (o *ExternalFormCreateModel) GetPossibleValuesOk() (*map[string][]ExternalFormAllowedValueModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PossibleValues, true
}

// SetPossibleValues sets field value
func (o *ExternalFormCreateModel) SetPossibleValues(v map[string][]ExternalFormAllowedValueModel) {
	o.PossibleValues = v
}

// GetFields returns the Fields field value
func (o *ExternalFormCreateModel) GetFields() []ExternalFormFieldModel {
	if o == nil {
		var ret []ExternalFormFieldModel
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *ExternalFormCreateModel) GetFieldsOk() ([]ExternalFormFieldModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *ExternalFormCreateModel) SetFields(v []ExternalFormFieldModel) {
	o.Fields = v
}

// GetLinks returns the Links field value
func (o *ExternalFormCreateModel) GetLinks() []ExternalFormLinkModel {
	if o == nil {
		var ret []ExternalFormLinkModel
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ExternalFormCreateModel) GetLinksOk() ([]ExternalFormLinkModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *ExternalFormCreateModel) SetLinks(v []ExternalFormLinkModel) {
	o.Links = v
}

// GetValues returns the Values field value
func (o *ExternalFormCreateModel) GetValues() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *ExternalFormCreateModel) GetValuesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *ExternalFormCreateModel) SetValues(v map[string]interface{}) {
	o.Values = v
}

func (o ExternalFormCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalFormCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["possibleValues"] = o.PossibleValues
	toSerialize["fields"] = o.Fields
	toSerialize["links"] = o.Links
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

func (o *ExternalFormCreateModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"possibleValues",
		"fields",
		"links",
		"values",
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

	varExternalFormCreateModel := _ExternalFormCreateModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalFormCreateModel)

	if err != nil {
		return err
	}

	*o = ExternalFormCreateModel(varExternalFormCreateModel)

	return err
}

type NullableExternalFormCreateModel struct {
	value *ExternalFormCreateModel
	isSet bool
}

func (v NullableExternalFormCreateModel) Get() *ExternalFormCreateModel {
	return v.value
}

func (v *NullableExternalFormCreateModel) Set(val *ExternalFormCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalFormCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalFormCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalFormCreateModel(val *ExternalFormCreateModel) *NullableExternalFormCreateModel {
	return &NullableExternalFormCreateModel{value: val, isSet: true}
}

func (v NullableExternalFormCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalFormCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


