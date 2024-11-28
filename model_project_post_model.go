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

// checks if the ProjectPostModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectPostModel{}

// ProjectPostModel struct for ProjectPostModel
type ProjectPostModel struct {
	// Description of the project
	Description NullableString `json:"description,omitempty"`
	// Name of the project
	Name string `json:"name"`
	// Indicates if the project is marked as favorite
	IsFavorite NullableBool `json:"isFavorite,omitempty"`
	// Indicates if the status \"Flaky/Stable\" sets automatically
	// Deprecated
	IsFlakyAuto NullableBool `json:"isFlakyAuto,omitempty"`
}

type _ProjectPostModel ProjectPostModel

// NewProjectPostModel instantiates a new ProjectPostModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectPostModel(name string) *ProjectPostModel {
	this := ProjectPostModel{}
	this.Name = name
	return &this
}

// NewProjectPostModelWithDefaults instantiates a new ProjectPostModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectPostModelWithDefaults() *ProjectPostModel {
	this := ProjectPostModel{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectPostModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectPostModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectPostModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ProjectPostModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ProjectPostModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ProjectPostModel) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value
func (o *ProjectPostModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectPostModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectPostModel) SetName(v string) {
	o.Name = v
}

// GetIsFavorite returns the IsFavorite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectPostModel) GetIsFavorite() bool {
	if o == nil || IsNil(o.IsFavorite.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFavorite.Get()
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectPostModel) GetIsFavoriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFavorite.Get(), o.IsFavorite.IsSet()
}

// HasIsFavorite returns a boolean if a field has been set.
func (o *ProjectPostModel) HasIsFavorite() bool {
	if o != nil && o.IsFavorite.IsSet() {
		return true
	}

	return false
}

// SetIsFavorite gets a reference to the given NullableBool and assigns it to the IsFavorite field.
func (o *ProjectPostModel) SetIsFavorite(v bool) {
	o.IsFavorite.Set(&v)
}
// SetIsFavoriteNil sets the value for IsFavorite to be an explicit nil
func (o *ProjectPostModel) SetIsFavoriteNil() {
	o.IsFavorite.Set(nil)
}

// UnsetIsFavorite ensures that no value is present for IsFavorite, not even an explicit nil
func (o *ProjectPostModel) UnsetIsFavorite() {
	o.IsFavorite.Unset()
}

// GetIsFlakyAuto returns the IsFlakyAuto field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *ProjectPostModel) GetIsFlakyAuto() bool {
	if o == nil || IsNil(o.IsFlakyAuto.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFlakyAuto.Get()
}

// GetIsFlakyAutoOk returns a tuple with the IsFlakyAuto field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *ProjectPostModel) GetIsFlakyAutoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFlakyAuto.Get(), o.IsFlakyAuto.IsSet()
}

// HasIsFlakyAuto returns a boolean if a field has been set.
func (o *ProjectPostModel) HasIsFlakyAuto() bool {
	if o != nil && o.IsFlakyAuto.IsSet() {
		return true
	}

	return false
}

// SetIsFlakyAuto gets a reference to the given NullableBool and assigns it to the IsFlakyAuto field.
// Deprecated
func (o *ProjectPostModel) SetIsFlakyAuto(v bool) {
	o.IsFlakyAuto.Set(&v)
}
// SetIsFlakyAutoNil sets the value for IsFlakyAuto to be an explicit nil
func (o *ProjectPostModel) SetIsFlakyAutoNil() {
	o.IsFlakyAuto.Set(nil)
}

// UnsetIsFlakyAuto ensures that no value is present for IsFlakyAuto, not even an explicit nil
func (o *ProjectPostModel) UnsetIsFlakyAuto() {
	o.IsFlakyAuto.Unset()
}

func (o ProjectPostModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectPostModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["name"] = o.Name
	if o.IsFavorite.IsSet() {
		toSerialize["isFavorite"] = o.IsFavorite.Get()
	}
	if o.IsFlakyAuto.IsSet() {
		toSerialize["isFlakyAuto"] = o.IsFlakyAuto.Get()
	}
	return toSerialize, nil
}

func (o *ProjectPostModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varProjectPostModel := _ProjectPostModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectPostModel)

	if err != nil {
		return err
	}

	*o = ProjectPostModel(varProjectPostModel)

	return err
}

type NullableProjectPostModel struct {
	value *ProjectPostModel
	isSet bool
}

func (v NullableProjectPostModel) Get() *ProjectPostModel {
	return v.value
}

func (v *NullableProjectPostModel) Set(val *ProjectPostModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectPostModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectPostModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectPostModel(val *ProjectPostModel) *NullableProjectPostModel {
	return &NullableProjectPostModel{value: val, isSet: true}
}

func (v NullableProjectPostModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectPostModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


