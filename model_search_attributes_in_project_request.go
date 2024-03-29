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

// checks if the SearchAttributesInProjectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchAttributesInProjectRequest{}

// SearchAttributesInProjectRequest struct for SearchAttributesInProjectRequest
type SearchAttributesInProjectRequest struct {
	// Specifies an attribute name to search for
	Name string `json:"name"`
	// Specifies an attribute mandatory status to search for
	IsRequired NullableBool `json:"isRequired,omitempty"`
	// Specifies an attribute global status to search for
	IsGlobal NullableBool `json:"isGlobal,omitempty"`
	// Specifies an attribute types to search for
	Types []CustomAttributeTypesEnum `json:"types"`
	// Specifies an attribute enabled status to search for
	IsEnabled NullableBool `json:"isEnabled,omitempty"`
}

// NewSearchAttributesInProjectRequest instantiates a new SearchAttributesInProjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchAttributesInProjectRequest(name string, types []CustomAttributeTypesEnum) *SearchAttributesInProjectRequest {
	this := SearchAttributesInProjectRequest{}
	this.Name = name
	this.Types = types
	return &this
}

// NewSearchAttributesInProjectRequestWithDefaults instantiates a new SearchAttributesInProjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchAttributesInProjectRequestWithDefaults() *SearchAttributesInProjectRequest {
	this := SearchAttributesInProjectRequest{}
	return &this
}

// GetName returns the Name field value
func (o *SearchAttributesInProjectRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SearchAttributesInProjectRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SearchAttributesInProjectRequest) SetName(v string) {
	o.Name = v
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchAttributesInProjectRequest) GetIsRequired() bool {
	if o == nil || IsNil(o.IsRequired.Get()) {
		var ret bool
		return ret
	}
	return *o.IsRequired.Get()
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchAttributesInProjectRequest) GetIsRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRequired.Get(), o.IsRequired.IsSet()
}

// HasIsRequired returns a boolean if a field has been set.
func (o *SearchAttributesInProjectRequest) HasIsRequired() bool {
	if o != nil && o.IsRequired.IsSet() {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given NullableBool and assigns it to the IsRequired field.
func (o *SearchAttributesInProjectRequest) SetIsRequired(v bool) {
	o.IsRequired.Set(&v)
}
// SetIsRequiredNil sets the value for IsRequired to be an explicit nil
func (o *SearchAttributesInProjectRequest) SetIsRequiredNil() {
	o.IsRequired.Set(nil)
}

// UnsetIsRequired ensures that no value is present for IsRequired, not even an explicit nil
func (o *SearchAttributesInProjectRequest) UnsetIsRequired() {
	o.IsRequired.Unset()
}

// GetIsGlobal returns the IsGlobal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchAttributesInProjectRequest) GetIsGlobal() bool {
	if o == nil || IsNil(o.IsGlobal.Get()) {
		var ret bool
		return ret
	}
	return *o.IsGlobal.Get()
}

// GetIsGlobalOk returns a tuple with the IsGlobal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchAttributesInProjectRequest) GetIsGlobalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsGlobal.Get(), o.IsGlobal.IsSet()
}

// HasIsGlobal returns a boolean if a field has been set.
func (o *SearchAttributesInProjectRequest) HasIsGlobal() bool {
	if o != nil && o.IsGlobal.IsSet() {
		return true
	}

	return false
}

// SetIsGlobal gets a reference to the given NullableBool and assigns it to the IsGlobal field.
func (o *SearchAttributesInProjectRequest) SetIsGlobal(v bool) {
	o.IsGlobal.Set(&v)
}
// SetIsGlobalNil sets the value for IsGlobal to be an explicit nil
func (o *SearchAttributesInProjectRequest) SetIsGlobalNil() {
	o.IsGlobal.Set(nil)
}

// UnsetIsGlobal ensures that no value is present for IsGlobal, not even an explicit nil
func (o *SearchAttributesInProjectRequest) UnsetIsGlobal() {
	o.IsGlobal.Unset()
}

// GetTypes returns the Types field value
func (o *SearchAttributesInProjectRequest) GetTypes() []CustomAttributeTypesEnum {
	if o == nil {
		var ret []CustomAttributeTypesEnum
		return ret
	}

	return o.Types
}

// GetTypesOk returns a tuple with the Types field value
// and a boolean to check if the value has been set.
func (o *SearchAttributesInProjectRequest) GetTypesOk() ([]CustomAttributeTypesEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Types, true
}

// SetTypes sets field value
func (o *SearchAttributesInProjectRequest) SetTypes(v []CustomAttributeTypesEnum) {
	o.Types = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchAttributesInProjectRequest) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.IsEnabled.Get()
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchAttributesInProjectRequest) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsEnabled.Get(), o.IsEnabled.IsSet()
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *SearchAttributesInProjectRequest) HasIsEnabled() bool {
	if o != nil && o.IsEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given NullableBool and assigns it to the IsEnabled field.
func (o *SearchAttributesInProjectRequest) SetIsEnabled(v bool) {
	o.IsEnabled.Set(&v)
}
// SetIsEnabledNil sets the value for IsEnabled to be an explicit nil
func (o *SearchAttributesInProjectRequest) SetIsEnabledNil() {
	o.IsEnabled.Set(nil)
}

// UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
func (o *SearchAttributesInProjectRequest) UnsetIsEnabled() {
	o.IsEnabled.Unset()
}

func (o SearchAttributesInProjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchAttributesInProjectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.IsRequired.IsSet() {
		toSerialize["isRequired"] = o.IsRequired.Get()
	}
	if o.IsGlobal.IsSet() {
		toSerialize["isGlobal"] = o.IsGlobal.Get()
	}
	toSerialize["types"] = o.Types
	if o.IsEnabled.IsSet() {
		toSerialize["isEnabled"] = o.IsEnabled.Get()
	}
	return toSerialize, nil
}

type NullableSearchAttributesInProjectRequest struct {
	value *SearchAttributesInProjectRequest
	isSet bool
}

func (v NullableSearchAttributesInProjectRequest) Get() *SearchAttributesInProjectRequest {
	return v.value
}

func (v *NullableSearchAttributesInProjectRequest) Set(val *SearchAttributesInProjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchAttributesInProjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchAttributesInProjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchAttributesInProjectRequest(val *SearchAttributesInProjectRequest) *NullableSearchAttributesInProjectRequest {
	return &NullableSearchAttributesInProjectRequest{value: val, isSet: true}
}

func (v NullableSearchAttributesInProjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchAttributesInProjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


