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

// checks if the AutoTestStepModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoTestStepModel{}

// AutoTestStepModel struct for AutoTestStepModel
type AutoTestStepModel struct {
	// Step name.
	Title string `json:"title"`
	// Detailed step description. It appears when the step is unfolded.
	Description NullableString `json:"description,omitempty"`
	// Includes a nested step inside another step. The maximum nesting level is 15.
	Steps []AutoTestStepModel `json:"steps,omitempty"`
}

type _AutoTestStepModel AutoTestStepModel

// NewAutoTestStepModel instantiates a new AutoTestStepModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoTestStepModel(title string) *AutoTestStepModel {
	this := AutoTestStepModel{}
	this.Title = title
	return &this
}

// NewAutoTestStepModelWithDefaults instantiates a new AutoTestStepModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoTestStepModelWithDefaults() *AutoTestStepModel {
	this := AutoTestStepModel{}
	return &this
}

// GetTitle returns the Title field value
func (o *AutoTestStepModel) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AutoTestStepModel) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *AutoTestStepModel) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestStepModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestStepModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AutoTestStepModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AutoTestStepModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AutoTestStepModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AutoTestStepModel) UnsetDescription() {
	o.Description.Unset()
}

// GetSteps returns the Steps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTestStepModel) GetSteps() []AutoTestStepModel {
	if o == nil {
		var ret []AutoTestStepModel
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTestStepModel) GetStepsOk() ([]AutoTestStepModel, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *AutoTestStepModel) HasSteps() bool {
	if o != nil && !IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []AutoTestStepModel and assigns it to the Steps field.
func (o *AutoTestStepModel) SetSteps(v []AutoTestStepModel) {
	o.Steps = v
}

func (o AutoTestStepModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoTestStepModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	return toSerialize, nil
}

func (o *AutoTestStepModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
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

	varAutoTestStepModel := _AutoTestStepModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAutoTestStepModel)

	if err != nil {
		return err
	}

	*o = AutoTestStepModel(varAutoTestStepModel)

	return err
}

type NullableAutoTestStepModel struct {
	value *AutoTestStepModel
	isSet bool
}

func (v NullableAutoTestStepModel) Get() *AutoTestStepModel {
	return v.value
}

func (v *NullableAutoTestStepModel) Set(val *AutoTestStepModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoTestStepModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoTestStepModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoTestStepModel(val *AutoTestStepModel) *NullableAutoTestStepModel {
	return &NullableAutoTestStepModel{value: val, isSet: true}
}

func (v NullableAutoTestStepModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoTestStepModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


