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

// checks if the WorkItemSelectApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemSelectApiModel{}

// WorkItemSelectApiModel struct for WorkItemSelectApiModel
type WorkItemSelectApiModel struct {
	Filter WorkItemFilterApiModel `json:"filter"`
	ExtractionModel NullableWorkItemExtractionApiModel `json:"extractionModel,omitempty"`
}

type _WorkItemSelectApiModel WorkItemSelectApiModel

// NewWorkItemSelectApiModel instantiates a new WorkItemSelectApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemSelectApiModel(filter WorkItemFilterApiModel) *WorkItemSelectApiModel {
	this := WorkItemSelectApiModel{}
	this.Filter = filter
	return &this
}

// NewWorkItemSelectApiModelWithDefaults instantiates a new WorkItemSelectApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemSelectApiModelWithDefaults() *WorkItemSelectApiModel {
	this := WorkItemSelectApiModel{}
	return &this
}

// GetFilter returns the Filter field value
func (o *WorkItemSelectApiModel) GetFilter() WorkItemFilterApiModel {
	if o == nil {
		var ret WorkItemFilterApiModel
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *WorkItemSelectApiModel) GetFilterOk() (*WorkItemFilterApiModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *WorkItemSelectApiModel) SetFilter(v WorkItemFilterApiModel) {
	o.Filter = v
}

// GetExtractionModel returns the ExtractionModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkItemSelectApiModel) GetExtractionModel() WorkItemExtractionApiModel {
	if o == nil || IsNil(o.ExtractionModel.Get()) {
		var ret WorkItemExtractionApiModel
		return ret
	}
	return *o.ExtractionModel.Get()
}

// GetExtractionModelOk returns a tuple with the ExtractionModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkItemSelectApiModel) GetExtractionModelOk() (*WorkItemExtractionApiModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtractionModel.Get(), o.ExtractionModel.IsSet()
}

// HasExtractionModel returns a boolean if a field has been set.
func (o *WorkItemSelectApiModel) HasExtractionModel() bool {
	if o != nil && o.ExtractionModel.IsSet() {
		return true
	}

	return false
}

// SetExtractionModel gets a reference to the given NullableWorkItemExtractionApiModel and assigns it to the ExtractionModel field.
func (o *WorkItemSelectApiModel) SetExtractionModel(v WorkItemExtractionApiModel) {
	o.ExtractionModel.Set(&v)
}
// SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil
func (o *WorkItemSelectApiModel) SetExtractionModelNil() {
	o.ExtractionModel.Set(nil)
}

// UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil
func (o *WorkItemSelectApiModel) UnsetExtractionModel() {
	o.ExtractionModel.Unset()
}

func (o WorkItemSelectApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemSelectApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filter"] = o.Filter
	if o.ExtractionModel.IsSet() {
		toSerialize["extractionModel"] = o.ExtractionModel.Get()
	}
	return toSerialize, nil
}

func (o *WorkItemSelectApiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filter",
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

	varWorkItemSelectApiModel := _WorkItemSelectApiModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkItemSelectApiModel)

	if err != nil {
		return err
	}

	*o = WorkItemSelectApiModel(varWorkItemSelectApiModel)

	return err
}

type NullableWorkItemSelectApiModel struct {
	value *WorkItemSelectApiModel
	isSet bool
}

func (v NullableWorkItemSelectApiModel) Get() *WorkItemSelectApiModel {
	return v.value
}

func (v *NullableWorkItemSelectApiModel) Set(val *WorkItemSelectApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemSelectApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemSelectApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemSelectApiModel(val *WorkItemSelectApiModel) *NullableWorkItemSelectApiModel {
	return &NullableWorkItemSelectApiModel{value: val, isSet: true}
}

func (v NullableWorkItemSelectApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemSelectApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


