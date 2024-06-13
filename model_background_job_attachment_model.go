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

// checks if the BackgroundJobAttachmentModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackgroundJobAttachmentModel{}

// BackgroundJobAttachmentModel struct for BackgroundJobAttachmentModel
type BackgroundJobAttachmentModel struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type _BackgroundJobAttachmentModel BackgroundJobAttachmentModel

// NewBackgroundJobAttachmentModel instantiates a new BackgroundJobAttachmentModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackgroundJobAttachmentModel(id string, name string, type_ string) *BackgroundJobAttachmentModel {
	this := BackgroundJobAttachmentModel{}
	this.Id = id
	this.Name = name
	this.Type = type_
	return &this
}

// NewBackgroundJobAttachmentModelWithDefaults instantiates a new BackgroundJobAttachmentModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackgroundJobAttachmentModelWithDefaults() *BackgroundJobAttachmentModel {
	this := BackgroundJobAttachmentModel{}
	return &this
}

// GetId returns the Id field value
func (o *BackgroundJobAttachmentModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BackgroundJobAttachmentModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BackgroundJobAttachmentModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *BackgroundJobAttachmentModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BackgroundJobAttachmentModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BackgroundJobAttachmentModel) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *BackgroundJobAttachmentModel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BackgroundJobAttachmentModel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BackgroundJobAttachmentModel) SetType(v string) {
	o.Type = v
}

func (o BackgroundJobAttachmentModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackgroundJobAttachmentModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *BackgroundJobAttachmentModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"type",
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

	varBackgroundJobAttachmentModel := _BackgroundJobAttachmentModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBackgroundJobAttachmentModel)

	if err != nil {
		return err
	}

	*o = BackgroundJobAttachmentModel(varBackgroundJobAttachmentModel)

	return err
}

type NullableBackgroundJobAttachmentModel struct {
	value *BackgroundJobAttachmentModel
	isSet bool
}

func (v NullableBackgroundJobAttachmentModel) Get() *BackgroundJobAttachmentModel {
	return v.value
}

func (v *NullableBackgroundJobAttachmentModel) Set(val *BackgroundJobAttachmentModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackgroundJobAttachmentModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackgroundJobAttachmentModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackgroundJobAttachmentModel(val *BackgroundJobAttachmentModel) *NullableBackgroundJobAttachmentModel {
	return &NullableBackgroundJobAttachmentModel{value: val, isSet: true}
}

func (v NullableBackgroundJobAttachmentModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackgroundJobAttachmentModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


