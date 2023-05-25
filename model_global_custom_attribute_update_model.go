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

// checks if the GlobalCustomAttributeUpdateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalCustomAttributeUpdateModel{}

// GlobalCustomAttributeUpdateModel struct for GlobalCustomAttributeUpdateModel
type GlobalCustomAttributeUpdateModel struct {
	// Name of attribute
	Name string `json:"name"`
	// Collection of attribute options  <br />  Available for attributes of type `options` and `multiple options` only
	Options []CustomAttributeOptionModel `json:"options,omitempty"`
	// Indicates whether the attribute is available
	IsEnabled NullableBool `json:"isEnabled,omitempty"`
	// Indicates whether the attribute value is mandatory to specify
	IsRequired NullableBool `json:"isRequired,omitempty"`
}

// NewGlobalCustomAttributeUpdateModel instantiates a new GlobalCustomAttributeUpdateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalCustomAttributeUpdateModel(name string) *GlobalCustomAttributeUpdateModel {
	this := GlobalCustomAttributeUpdateModel{}
	this.Name = name
	return &this
}

// NewGlobalCustomAttributeUpdateModelWithDefaults instantiates a new GlobalCustomAttributeUpdateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalCustomAttributeUpdateModelWithDefaults() *GlobalCustomAttributeUpdateModel {
	this := GlobalCustomAttributeUpdateModel{}
	return &this
}

// GetName returns the Name field value
func (o *GlobalCustomAttributeUpdateModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GlobalCustomAttributeUpdateModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GlobalCustomAttributeUpdateModel) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalCustomAttributeUpdateModel) GetOptions() []CustomAttributeOptionModel {
	if o == nil {
		var ret []CustomAttributeOptionModel
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GlobalCustomAttributeUpdateModel) GetOptionsOk() ([]CustomAttributeOptionModel, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *GlobalCustomAttributeUpdateModel) HasOptions() bool {
	if o != nil && IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []CustomAttributeOptionModel and assigns it to the Options field.
func (o *GlobalCustomAttributeUpdateModel) SetOptions(v []CustomAttributeOptionModel) {
	o.Options = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalCustomAttributeUpdateModel) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.IsEnabled.Get()
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GlobalCustomAttributeUpdateModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsEnabled.Get(), o.IsEnabled.IsSet()
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *GlobalCustomAttributeUpdateModel) HasIsEnabled() bool {
	if o != nil && o.IsEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given NullableBool and assigns it to the IsEnabled field.
func (o *GlobalCustomAttributeUpdateModel) SetIsEnabled(v bool) {
	o.IsEnabled.Set(&v)
}
// SetIsEnabledNil sets the value for IsEnabled to be an explicit nil
func (o *GlobalCustomAttributeUpdateModel) SetIsEnabledNil() {
	o.IsEnabled.Set(nil)
}

// UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
func (o *GlobalCustomAttributeUpdateModel) UnsetIsEnabled() {
	o.IsEnabled.Unset()
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalCustomAttributeUpdateModel) GetIsRequired() bool {
	if o == nil || IsNil(o.IsRequired.Get()) {
		var ret bool
		return ret
	}
	return *o.IsRequired.Get()
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GlobalCustomAttributeUpdateModel) GetIsRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRequired.Get(), o.IsRequired.IsSet()
}

// HasIsRequired returns a boolean if a field has been set.
func (o *GlobalCustomAttributeUpdateModel) HasIsRequired() bool {
	if o != nil && o.IsRequired.IsSet() {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given NullableBool and assigns it to the IsRequired field.
func (o *GlobalCustomAttributeUpdateModel) SetIsRequired(v bool) {
	o.IsRequired.Set(&v)
}
// SetIsRequiredNil sets the value for IsRequired to be an explicit nil
func (o *GlobalCustomAttributeUpdateModel) SetIsRequiredNil() {
	o.IsRequired.Set(nil)
}

// UnsetIsRequired ensures that no value is present for IsRequired, not even an explicit nil
func (o *GlobalCustomAttributeUpdateModel) UnsetIsRequired() {
	o.IsRequired.Unset()
}

func (o GlobalCustomAttributeUpdateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalCustomAttributeUpdateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.IsEnabled.IsSet() {
		toSerialize["isEnabled"] = o.IsEnabled.Get()
	}
	if o.IsRequired.IsSet() {
		toSerialize["isRequired"] = o.IsRequired.Get()
	}
	return toSerialize, nil
}

type NullableGlobalCustomAttributeUpdateModel struct {
	value *GlobalCustomAttributeUpdateModel
	isSet bool
}

func (v NullableGlobalCustomAttributeUpdateModel) Get() *GlobalCustomAttributeUpdateModel {
	return v.value
}

func (v *NullableGlobalCustomAttributeUpdateModel) Set(val *GlobalCustomAttributeUpdateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalCustomAttributeUpdateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalCustomAttributeUpdateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalCustomAttributeUpdateModel(val *GlobalCustomAttributeUpdateModel) *NullableGlobalCustomAttributeUpdateModel {
	return &NullableGlobalCustomAttributeUpdateModel{value: val, isSet: true}
}

func (v NullableGlobalCustomAttributeUpdateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalCustomAttributeUpdateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


