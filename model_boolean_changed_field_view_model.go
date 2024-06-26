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

// checks if the BooleanChangedFieldViewModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BooleanChangedFieldViewModel{}

// BooleanChangedFieldViewModel struct for BooleanChangedFieldViewModel
type BooleanChangedFieldViewModel struct {
	OldValue bool `json:"oldValue"`
	NewValue bool `json:"newValue"`
}

type _BooleanChangedFieldViewModel BooleanChangedFieldViewModel

// NewBooleanChangedFieldViewModel instantiates a new BooleanChangedFieldViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBooleanChangedFieldViewModel(oldValue bool, newValue bool) *BooleanChangedFieldViewModel {
	this := BooleanChangedFieldViewModel{}
	this.OldValue = oldValue
	this.NewValue = newValue
	return &this
}

// NewBooleanChangedFieldViewModelWithDefaults instantiates a new BooleanChangedFieldViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBooleanChangedFieldViewModelWithDefaults() *BooleanChangedFieldViewModel {
	this := BooleanChangedFieldViewModel{}
	return &this
}

// GetOldValue returns the OldValue field value
func (o *BooleanChangedFieldViewModel) GetOldValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value
// and a boolean to check if the value has been set.
func (o *BooleanChangedFieldViewModel) GetOldValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OldValue, true
}

// SetOldValue sets field value
func (o *BooleanChangedFieldViewModel) SetOldValue(v bool) {
	o.OldValue = v
}

// GetNewValue returns the NewValue field value
func (o *BooleanChangedFieldViewModel) GetNewValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value
// and a boolean to check if the value has been set.
func (o *BooleanChangedFieldViewModel) GetNewValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewValue, true
}

// SetNewValue sets field value
func (o *BooleanChangedFieldViewModel) SetNewValue(v bool) {
	o.NewValue = v
}

func (o BooleanChangedFieldViewModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BooleanChangedFieldViewModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["oldValue"] = o.OldValue
	toSerialize["newValue"] = o.NewValue
	return toSerialize, nil
}

func (o *BooleanChangedFieldViewModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"oldValue",
		"newValue",
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

	varBooleanChangedFieldViewModel := _BooleanChangedFieldViewModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBooleanChangedFieldViewModel)

	if err != nil {
		return err
	}

	*o = BooleanChangedFieldViewModel(varBooleanChangedFieldViewModel)

	return err
}

type NullableBooleanChangedFieldViewModel struct {
	value *BooleanChangedFieldViewModel
	isSet bool
}

func (v NullableBooleanChangedFieldViewModel) Get() *BooleanChangedFieldViewModel {
	return v.value
}

func (v *NullableBooleanChangedFieldViewModel) Set(val *BooleanChangedFieldViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBooleanChangedFieldViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBooleanChangedFieldViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBooleanChangedFieldViewModel(val *BooleanChangedFieldViewModel) *NullableBooleanChangedFieldViewModel {
	return &NullableBooleanChangedFieldViewModel{value: val, isSet: true}
}

func (v NullableBooleanChangedFieldViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBooleanChangedFieldViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


