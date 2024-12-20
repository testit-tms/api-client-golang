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

// checks if the TestStatusUpdateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestStatusUpdateModel{}

// TestStatusUpdateModel struct for TestStatusUpdateModel
type TestStatusUpdateModel struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
}

type _TestStatusUpdateModel TestStatusUpdateModel

// NewTestStatusUpdateModel instantiates a new TestStatusUpdateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestStatusUpdateModel(id string, name string) *TestStatusUpdateModel {
	this := TestStatusUpdateModel{}
	this.Id = id
	this.Name = name
	return &this
}

// NewTestStatusUpdateModelWithDefaults instantiates a new TestStatusUpdateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestStatusUpdateModelWithDefaults() *TestStatusUpdateModel {
	this := TestStatusUpdateModel{}
	return &this
}

// GetId returns the Id field value
func (o *TestStatusUpdateModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestStatusUpdateModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestStatusUpdateModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TestStatusUpdateModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TestStatusUpdateModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TestStatusUpdateModel) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestStatusUpdateModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestStatusUpdateModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TestStatusUpdateModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TestStatusUpdateModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TestStatusUpdateModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TestStatusUpdateModel) UnsetDescription() {
	o.Description.Unset()
}

func (o TestStatusUpdateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestStatusUpdateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return toSerialize, nil
}

func (o *TestStatusUpdateModel) UnmarshalJSON(data []byte) (err error) {
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

	varTestStatusUpdateModel := _TestStatusUpdateModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestStatusUpdateModel)

	if err != nil {
		return err
	}

	*o = TestStatusUpdateModel(varTestStatusUpdateModel)

	return err
}

type NullableTestStatusUpdateModel struct {
	value *TestStatusUpdateModel
	isSet bool
}

func (v NullableTestStatusUpdateModel) Get() *TestStatusUpdateModel {
	return v.value
}

func (v *NullableTestStatusUpdateModel) Set(val *TestStatusUpdateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestStatusUpdateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestStatusUpdateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestStatusUpdateModel(val *TestStatusUpdateModel) *NullableTestStatusUpdateModel {
	return &NullableTestStatusUpdateModel{value: val, isSet: true}
}

func (v NullableTestStatusUpdateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestStatusUpdateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


