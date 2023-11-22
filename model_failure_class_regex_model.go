/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
)

// checks if the FailureClassRegexModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FailureClassRegexModel{}

// FailureClassRegexModel struct for FailureClassRegexModel
type FailureClassRegexModel struct {
	RegexText string `json:"regexText"`
	FailureClassId NullableString `json:"failureClassId,omitempty"`
	// Unique ID of the entity
	Id string `json:"id"`
	// Indicates if the entity is deleted
	IsDeleted bool `json:"isDeleted"`
}

// NewFailureClassRegexModel instantiates a new FailureClassRegexModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailureClassRegexModel(regexText string, id string, isDeleted bool) *FailureClassRegexModel {
	this := FailureClassRegexModel{}
	this.RegexText = regexText
	this.Id = id
	this.IsDeleted = isDeleted
	return &this
}

// NewFailureClassRegexModelWithDefaults instantiates a new FailureClassRegexModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailureClassRegexModelWithDefaults() *FailureClassRegexModel {
	this := FailureClassRegexModel{}
	return &this
}

// GetRegexText returns the RegexText field value
func (o *FailureClassRegexModel) GetRegexText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegexText
}

// GetRegexTextOk returns a tuple with the RegexText field value
// and a boolean to check if the value has been set.
func (o *FailureClassRegexModel) GetRegexTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegexText, true
}

// SetRegexText sets field value
func (o *FailureClassRegexModel) SetRegexText(v string) {
	o.RegexText = v
}

// GetFailureClassId returns the FailureClassId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FailureClassRegexModel) GetFailureClassId() string {
	if o == nil || IsNil(o.FailureClassId.Get()) {
		var ret string
		return ret
	}
	return *o.FailureClassId.Get()
}

// GetFailureClassIdOk returns a tuple with the FailureClassId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FailureClassRegexModel) GetFailureClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureClassId.Get(), o.FailureClassId.IsSet()
}

// HasFailureClassId returns a boolean if a field has been set.
func (o *FailureClassRegexModel) HasFailureClassId() bool {
	if o != nil && o.FailureClassId.IsSet() {
		return true
	}

	return false
}

// SetFailureClassId gets a reference to the given NullableString and assigns it to the FailureClassId field.
func (o *FailureClassRegexModel) SetFailureClassId(v string) {
	o.FailureClassId.Set(&v)
}
// SetFailureClassIdNil sets the value for FailureClassId to be an explicit nil
func (o *FailureClassRegexModel) SetFailureClassIdNil() {
	o.FailureClassId.Set(nil)
}

// UnsetFailureClassId ensures that no value is present for FailureClassId, not even an explicit nil
func (o *FailureClassRegexModel) UnsetFailureClassId() {
	o.FailureClassId.Unset()
}

// GetId returns the Id field value
func (o *FailureClassRegexModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FailureClassRegexModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FailureClassRegexModel) SetId(v string) {
	o.Id = v
}

// GetIsDeleted returns the IsDeleted field value
func (o *FailureClassRegexModel) GetIsDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value
// and a boolean to check if the value has been set.
func (o *FailureClassRegexModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeleted, true
}

// SetIsDeleted sets field value
func (o *FailureClassRegexModel) SetIsDeleted(v bool) {
	o.IsDeleted = v
}

func (o FailureClassRegexModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FailureClassRegexModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["regexText"] = o.RegexText
	if o.FailureClassId.IsSet() {
		toSerialize["failureClassId"] = o.FailureClassId.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["isDeleted"] = o.IsDeleted
	return toSerialize, nil
}

type NullableFailureClassRegexModel struct {
	value *FailureClassRegexModel
	isSet bool
}

func (v NullableFailureClassRegexModel) Get() *FailureClassRegexModel {
	return v.value
}

func (v *NullableFailureClassRegexModel) Set(val *FailureClassRegexModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFailureClassRegexModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFailureClassRegexModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailureClassRegexModel(val *FailureClassRegexModel) *NullableFailureClassRegexModel {
	return &NullableFailureClassRegexModel{value: val, isSet: true}
}

func (v NullableFailureClassRegexModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailureClassRegexModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


