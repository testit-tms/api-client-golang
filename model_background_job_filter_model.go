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

// checks if the BackgroundJobFilterModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackgroundJobFilterModel{}

// BackgroundJobFilterModel struct for BackgroundJobFilterModel
type BackgroundJobFilterModel struct {
	Types []BackgroundJobType `json:"types,omitempty"`
	States []BackgroundJobState `json:"states,omitempty"`
	IsDeleted NullableBool `json:"isDeleted,omitempty"`
	StartDate NullableDateTimeRangeSelectorModel `json:"startDate,omitempty"`
	EndDate NullableDateTimeRangeSelectorModel `json:"endDate,omitempty"`
}

// NewBackgroundJobFilterModel instantiates a new BackgroundJobFilterModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackgroundJobFilterModel() *BackgroundJobFilterModel {
	this := BackgroundJobFilterModel{}
	return &this
}

// NewBackgroundJobFilterModelWithDefaults instantiates a new BackgroundJobFilterModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackgroundJobFilterModelWithDefaults() *BackgroundJobFilterModel {
	this := BackgroundJobFilterModel{}
	return &this
}

// GetTypes returns the Types field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackgroundJobFilterModel) GetTypes() []BackgroundJobType {
	if o == nil {
		var ret []BackgroundJobType
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackgroundJobFilterModel) GetTypesOk() ([]BackgroundJobType, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *BackgroundJobFilterModel) HasTypes() bool {
	if o != nil && IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []BackgroundJobType and assigns it to the Types field.
func (o *BackgroundJobFilterModel) SetTypes(v []BackgroundJobType) {
	o.Types = v
}

// GetStates returns the States field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackgroundJobFilterModel) GetStates() []BackgroundJobState {
	if o == nil {
		var ret []BackgroundJobState
		return ret
	}
	return o.States
}

// GetStatesOk returns a tuple with the States field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackgroundJobFilterModel) GetStatesOk() ([]BackgroundJobState, bool) {
	if o == nil || IsNil(o.States) {
		return nil, false
	}
	return o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *BackgroundJobFilterModel) HasStates() bool {
	if o != nil && IsNil(o.States) {
		return true
	}

	return false
}

// SetStates gets a reference to the given []BackgroundJobState and assigns it to the States field.
func (o *BackgroundJobFilterModel) SetStates(v []BackgroundJobState) {
	o.States = v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackgroundJobFilterModel) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDeleted.Get()
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackgroundJobFilterModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDeleted.Get(), o.IsDeleted.IsSet()
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *BackgroundJobFilterModel) HasIsDeleted() bool {
	if o != nil && o.IsDeleted.IsSet() {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given NullableBool and assigns it to the IsDeleted field.
func (o *BackgroundJobFilterModel) SetIsDeleted(v bool) {
	o.IsDeleted.Set(&v)
}
// SetIsDeletedNil sets the value for IsDeleted to be an explicit nil
func (o *BackgroundJobFilterModel) SetIsDeletedNil() {
	o.IsDeleted.Set(nil)
}

// UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
func (o *BackgroundJobFilterModel) UnsetIsDeleted() {
	o.IsDeleted.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackgroundJobFilterModel) GetStartDate() DateTimeRangeSelectorModel {
	if o == nil || IsNil(o.StartDate.Get()) {
		var ret DateTimeRangeSelectorModel
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackgroundJobFilterModel) GetStartDateOk() (*DateTimeRangeSelectorModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *BackgroundJobFilterModel) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableDateTimeRangeSelectorModel and assigns it to the StartDate field.
func (o *BackgroundJobFilterModel) SetStartDate(v DateTimeRangeSelectorModel) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *BackgroundJobFilterModel) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *BackgroundJobFilterModel) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackgroundJobFilterModel) GetEndDate() DateTimeRangeSelectorModel {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret DateTimeRangeSelectorModel
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackgroundJobFilterModel) GetEndDateOk() (*DateTimeRangeSelectorModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *BackgroundJobFilterModel) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableDateTimeRangeSelectorModel and assigns it to the EndDate field.
func (o *BackgroundJobFilterModel) SetEndDate(v DateTimeRangeSelectorModel) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *BackgroundJobFilterModel) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *BackgroundJobFilterModel) UnsetEndDate() {
	o.EndDate.Unset()
}

func (o BackgroundJobFilterModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackgroundJobFilterModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	if o.States != nil {
		toSerialize["states"] = o.States
	}
	if o.IsDeleted.IsSet() {
		toSerialize["isDeleted"] = o.IsDeleted.Get()
	}
	if o.StartDate.IsSet() {
		toSerialize["startDate"] = o.StartDate.Get()
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	return toSerialize, nil
}

type NullableBackgroundJobFilterModel struct {
	value *BackgroundJobFilterModel
	isSet bool
}

func (v NullableBackgroundJobFilterModel) Get() *BackgroundJobFilterModel {
	return v.value
}

func (v *NullableBackgroundJobFilterModel) Set(val *BackgroundJobFilterModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackgroundJobFilterModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackgroundJobFilterModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackgroundJobFilterModel(val *BackgroundJobFilterModel) *NullableBackgroundJobFilterModel {
	return &NullableBackgroundJobFilterModel{value: val, isSet: true}
}

func (v NullableBackgroundJobFilterModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackgroundJobFilterModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


