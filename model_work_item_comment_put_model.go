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

// checks if the WorkItemCommentPutModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemCommentPutModel{}

// WorkItemCommentPutModel struct for WorkItemCommentPutModel
type WorkItemCommentPutModel struct {
	// Text of the comment
	Text string `json:"text"`
	// Unique ID of the comment
	Id string `json:"id"`
}

type _WorkItemCommentPutModel WorkItemCommentPutModel

// NewWorkItemCommentPutModel instantiates a new WorkItemCommentPutModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemCommentPutModel(text string, id string) *WorkItemCommentPutModel {
	this := WorkItemCommentPutModel{}
	this.Text = text
	this.Id = id
	return &this
}

// NewWorkItemCommentPutModelWithDefaults instantiates a new WorkItemCommentPutModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemCommentPutModelWithDefaults() *WorkItemCommentPutModel {
	this := WorkItemCommentPutModel{}
	return &this
}

// GetText returns the Text field value
func (o *WorkItemCommentPutModel) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *WorkItemCommentPutModel) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *WorkItemCommentPutModel) SetText(v string) {
	o.Text = v
}

// GetId returns the Id field value
func (o *WorkItemCommentPutModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkItemCommentPutModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkItemCommentPutModel) SetId(v string) {
	o.Id = v
}

func (o WorkItemCommentPutModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemCommentPutModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["text"] = o.Text
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *WorkItemCommentPutModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"text",
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

	varWorkItemCommentPutModel := _WorkItemCommentPutModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkItemCommentPutModel)

	if err != nil {
		return err
	}

	*o = WorkItemCommentPutModel(varWorkItemCommentPutModel)

	return err
}

type NullableWorkItemCommentPutModel struct {
	value *WorkItemCommentPutModel
	isSet bool
}

func (v NullableWorkItemCommentPutModel) Get() *WorkItemCommentPutModel {
	return v.value
}

func (v *NullableWorkItemCommentPutModel) Set(val *WorkItemCommentPutModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemCommentPutModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemCommentPutModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemCommentPutModel(val *WorkItemCommentPutModel) *NullableWorkItemCommentPutModel {
	return &NullableWorkItemCommentPutModel{value: val, isSet: true}
}

func (v NullableWorkItemCommentPutModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemCommentPutModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


