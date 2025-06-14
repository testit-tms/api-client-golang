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

// checks if the TestStatusShortApiResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestStatusShortApiResult{}

// TestStatusShortApiResult struct for TestStatusShortApiResult
type TestStatusShortApiResult struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type _TestStatusShortApiResult TestStatusShortApiResult

// NewTestStatusShortApiResult instantiates a new TestStatusShortApiResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestStatusShortApiResult(id string, name string, code string) *TestStatusShortApiResult {
	this := TestStatusShortApiResult{}
	this.Id = id
	this.Name = name
	this.Code = code
	return &this
}

// NewTestStatusShortApiResultWithDefaults instantiates a new TestStatusShortApiResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestStatusShortApiResultWithDefaults() *TestStatusShortApiResult {
	this := TestStatusShortApiResult{}
	return &this
}

// GetId returns the Id field value
func (o *TestStatusShortApiResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestStatusShortApiResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestStatusShortApiResult) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TestStatusShortApiResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TestStatusShortApiResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TestStatusShortApiResult) SetName(v string) {
	o.Name = v
}

// GetCode returns the Code field value
func (o *TestStatusShortApiResult) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *TestStatusShortApiResult) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *TestStatusShortApiResult) SetCode(v string) {
	o.Code = v
}

func (o TestStatusShortApiResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestStatusShortApiResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["code"] = o.Code
	return toSerialize, nil
}

func (o *TestStatusShortApiResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varTestStatusShortApiResult := _TestStatusShortApiResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestStatusShortApiResult)

	if err != nil {
		return err
	}

	*o = TestStatusShortApiResult(varTestStatusShortApiResult)

	return err
}

type NullableTestStatusShortApiResult struct {
	value *TestStatusShortApiResult
	isSet bool
}

func (v NullableTestStatusShortApiResult) Get() *TestStatusShortApiResult {
	return v.value
}

func (v *NullableTestStatusShortApiResult) Set(val *TestStatusShortApiResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTestStatusShortApiResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTestStatusShortApiResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestStatusShortApiResult(val *TestStatusShortApiResult) *NullableTestStatusShortApiResult {
	return &NullableTestStatusShortApiResult{value: val, isSet: true}
}

func (v NullableTestStatusShortApiResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestStatusShortApiResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


