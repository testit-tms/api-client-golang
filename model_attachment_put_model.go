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

// checks if the AttachmentPutModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachmentPutModel{}

// AttachmentPutModel struct for AttachmentPutModel
type AttachmentPutModel struct {
	// Unique ID of the attachment
	Id string `json:"id"`
}

type _AttachmentPutModel AttachmentPutModel

// NewAttachmentPutModel instantiates a new AttachmentPutModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentPutModel(id string) *AttachmentPutModel {
	this := AttachmentPutModel{}
	this.Id = id
	return &this
}

// NewAttachmentPutModelWithDefaults instantiates a new AttachmentPutModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentPutModelWithDefaults() *AttachmentPutModel {
	this := AttachmentPutModel{}
	return &this
}

// GetId returns the Id field value
func (o *AttachmentPutModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AttachmentPutModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AttachmentPutModel) SetId(v string) {
	o.Id = v
}

func (o AttachmentPutModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachmentPutModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *AttachmentPutModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varAttachmentPutModel := _AttachmentPutModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAttachmentPutModel)

	if err != nil {
		return err
	}

	*o = AttachmentPutModel(varAttachmentPutModel)

	return err
}

type NullableAttachmentPutModel struct {
	value *AttachmentPutModel
	isSet bool
}

func (v NullableAttachmentPutModel) Get() *AttachmentPutModel {
	return v.value
}

func (v *NullableAttachmentPutModel) Set(val *AttachmentPutModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentPutModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentPutModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentPutModel(val *AttachmentPutModel) *NullableAttachmentPutModel {
	return &NullableAttachmentPutModel{value: val, isSet: true}
}

func (v NullableAttachmentPutModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentPutModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


