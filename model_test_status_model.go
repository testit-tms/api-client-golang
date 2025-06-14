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

// checks if the TestStatusModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestStatusModel{}

// TestStatusModel struct for TestStatusModel
type TestStatusModel struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Type TestStatusType `json:"type"`
	IsSystem bool `json:"isSystem"`
	Code string `json:"code"`
	Description NullableString `json:"description,omitempty"`
}

type _TestStatusModel TestStatusModel

// NewTestStatusModel instantiates a new TestStatusModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestStatusModel(id string, name string, type_ TestStatusType, isSystem bool, code string) *TestStatusModel {
	this := TestStatusModel{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.IsSystem = isSystem
	this.Code = code
	return &this
}

// NewTestStatusModelWithDefaults instantiates a new TestStatusModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestStatusModelWithDefaults() *TestStatusModel {
	this := TestStatusModel{}
	return &this
}

// GetId returns the Id field value
func (o *TestStatusModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestStatusModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestStatusModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TestStatusModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TestStatusModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TestStatusModel) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *TestStatusModel) GetType() TestStatusType {
	if o == nil {
		var ret TestStatusType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TestStatusModel) GetTypeOk() (*TestStatusType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TestStatusModel) SetType(v TestStatusType) {
	o.Type = v
}

// GetIsSystem returns the IsSystem field value
func (o *TestStatusModel) GetIsSystem() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSystem
}

// GetIsSystemOk returns a tuple with the IsSystem field value
// and a boolean to check if the value has been set.
func (o *TestStatusModel) GetIsSystemOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSystem, true
}

// SetIsSystem sets field value
func (o *TestStatusModel) SetIsSystem(v bool) {
	o.IsSystem = v
}

// GetCode returns the Code field value
func (o *TestStatusModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *TestStatusModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *TestStatusModel) SetCode(v string) {
	o.Code = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestStatusModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestStatusModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TestStatusModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TestStatusModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TestStatusModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TestStatusModel) UnsetDescription() {
	o.Description.Unset()
}

func (o TestStatusModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestStatusModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["isSystem"] = o.IsSystem
	toSerialize["code"] = o.Code
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return toSerialize, nil
}

func (o *TestStatusModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"type",
		"isSystem",
		"code",
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

	varTestStatusModel := _TestStatusModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestStatusModel)

	if err != nil {
		return err
	}

	*o = TestStatusModel(varTestStatusModel)

	return err
}

type NullableTestStatusModel struct {
	value *TestStatusModel
	isSet bool
}

func (v NullableTestStatusModel) Get() *TestStatusModel {
	return v.value
}

func (v *NullableTestStatusModel) Set(val *TestStatusModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestStatusModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestStatusModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestStatusModel(val *TestStatusModel) *NullableTestStatusModel {
	return &NullableTestStatusModel{value: val, isSet: true}
}

func (v NullableTestStatusModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestStatusModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


