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

// checks if the TagSelectModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagSelectModel{}

// TagSelectModel struct for TagSelectModel
type TagSelectModel struct {
	Filter NullableTagsFilterModel `json:"filter,omitempty"`
	ExtractionModel NullableTagExtractionModel `json:"extractionModel,omitempty"`
}

// NewTagSelectModel instantiates a new TagSelectModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagSelectModel() *TagSelectModel {
	this := TagSelectModel{}
	return &this
}

// NewTagSelectModelWithDefaults instantiates a new TagSelectModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagSelectModelWithDefaults() *TagSelectModel {
	this := TagSelectModel{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagSelectModel) GetFilter() TagsFilterModel {
	if o == nil || IsNil(o.Filter.Get()) {
		var ret TagsFilterModel
		return ret
	}
	return *o.Filter.Get()
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagSelectModel) GetFilterOk() (*TagsFilterModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter.Get(), o.Filter.IsSet()
}

// HasFilter returns a boolean if a field has been set.
func (o *TagSelectModel) HasFilter() bool {
	if o != nil && o.Filter.IsSet() {
		return true
	}

	return false
}

// SetFilter gets a reference to the given NullableTagsFilterModel and assigns it to the Filter field.
func (o *TagSelectModel) SetFilter(v TagsFilterModel) {
	o.Filter.Set(&v)
}
// SetFilterNil sets the value for Filter to be an explicit nil
func (o *TagSelectModel) SetFilterNil() {
	o.Filter.Set(nil)
}

// UnsetFilter ensures that no value is present for Filter, not even an explicit nil
func (o *TagSelectModel) UnsetFilter() {
	o.Filter.Unset()
}

// GetExtractionModel returns the ExtractionModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TagSelectModel) GetExtractionModel() TagExtractionModel {
	if o == nil || IsNil(o.ExtractionModel.Get()) {
		var ret TagExtractionModel
		return ret
	}
	return *o.ExtractionModel.Get()
}

// GetExtractionModelOk returns a tuple with the ExtractionModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagSelectModel) GetExtractionModelOk() (*TagExtractionModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtractionModel.Get(), o.ExtractionModel.IsSet()
}

// HasExtractionModel returns a boolean if a field has been set.
func (o *TagSelectModel) HasExtractionModel() bool {
	if o != nil && o.ExtractionModel.IsSet() {
		return true
	}

	return false
}

// SetExtractionModel gets a reference to the given NullableTagExtractionModel and assigns it to the ExtractionModel field.
func (o *TagSelectModel) SetExtractionModel(v TagExtractionModel) {
	o.ExtractionModel.Set(&v)
}
// SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil
func (o *TagSelectModel) SetExtractionModelNil() {
	o.ExtractionModel.Set(nil)
}

// UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil
func (o *TagSelectModel) UnsetExtractionModel() {
	o.ExtractionModel.Unset()
}

func (o TagSelectModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagSelectModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter.IsSet() {
		toSerialize["filter"] = o.Filter.Get()
	}
	if o.ExtractionModel.IsSet() {
		toSerialize["extractionModel"] = o.ExtractionModel.Get()
	}
	return toSerialize, nil
}

type NullableTagSelectModel struct {
	value *TagSelectModel
	isSet bool
}

func (v NullableTagSelectModel) Get() *TagSelectModel {
	return v.value
}

func (v *NullableTagSelectModel) Set(val *TagSelectModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTagSelectModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTagSelectModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagSelectModel(val *TagSelectModel) *NullableTagSelectModel {
	return &NullableTagSelectModel{value: val, isSet: true}
}

func (v NullableTagSelectModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagSelectModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

