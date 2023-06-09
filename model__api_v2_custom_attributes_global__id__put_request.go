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

// checks if the ApiV2CustomAttributesGlobalIdPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV2CustomAttributesGlobalIdPutRequest{}

// ApiV2CustomAttributesGlobalIdPutRequest struct for ApiV2CustomAttributesGlobalIdPutRequest
type ApiV2CustomAttributesGlobalIdPutRequest struct {
	// Name of attribute
	Name string `json:"name"`
	// Collection of attribute options  <br />  Available for attributes of type `options` and `multiple options` only
	Options []CustomAttributeOptionModel `json:"options,omitempty"`
	// Indicates whether the attribute is available
	IsEnabled NullableBool `json:"isEnabled,omitempty"`
	// Indicates whether the attribute value is mandatory to specify
	IsRequired NullableBool `json:"isRequired,omitempty"`
}

// NewApiV2CustomAttributesGlobalIdPutRequest instantiates a new ApiV2CustomAttributesGlobalIdPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2CustomAttributesGlobalIdPutRequest(name string) *ApiV2CustomAttributesGlobalIdPutRequest {
	this := ApiV2CustomAttributesGlobalIdPutRequest{}
	this.Name = name
	return &this
}

// NewApiV2CustomAttributesGlobalIdPutRequestWithDefaults instantiates a new ApiV2CustomAttributesGlobalIdPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2CustomAttributesGlobalIdPutRequestWithDefaults() *ApiV2CustomAttributesGlobalIdPutRequest {
	this := ApiV2CustomAttributesGlobalIdPutRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ApiV2CustomAttributesGlobalIdPutRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiV2CustomAttributesGlobalIdPutRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiV2CustomAttributesGlobalIdPutRequest) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2CustomAttributesGlobalIdPutRequest) GetOptions() []CustomAttributeOptionModel {
	if o == nil {
		var ret []CustomAttributeOptionModel
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2CustomAttributesGlobalIdPutRequest) GetOptionsOk() ([]CustomAttributeOptionModel, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ApiV2CustomAttributesGlobalIdPutRequest) HasOptions() bool {
	if o != nil && IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []CustomAttributeOptionModel and assigns it to the Options field.
func (o *ApiV2CustomAttributesGlobalIdPutRequest) SetOptions(v []CustomAttributeOptionModel) {
	o.Options = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2CustomAttributesGlobalIdPutRequest) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.IsEnabled.Get()
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2CustomAttributesGlobalIdPutRequest) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsEnabled.Get(), o.IsEnabled.IsSet()
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *ApiV2CustomAttributesGlobalIdPutRequest) HasIsEnabled() bool {
	if o != nil && o.IsEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given NullableBool and assigns it to the IsEnabled field.
func (o *ApiV2CustomAttributesGlobalIdPutRequest) SetIsEnabled(v bool) {
	o.IsEnabled.Set(&v)
}
// SetIsEnabledNil sets the value for IsEnabled to be an explicit nil
func (o *ApiV2CustomAttributesGlobalIdPutRequest) SetIsEnabledNil() {
	o.IsEnabled.Set(nil)
}

// UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
func (o *ApiV2CustomAttributesGlobalIdPutRequest) UnsetIsEnabled() {
	o.IsEnabled.Unset()
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2CustomAttributesGlobalIdPutRequest) GetIsRequired() bool {
	if o == nil || IsNil(o.IsRequired.Get()) {
		var ret bool
		return ret
	}
	return *o.IsRequired.Get()
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2CustomAttributesGlobalIdPutRequest) GetIsRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRequired.Get(), o.IsRequired.IsSet()
}

// HasIsRequired returns a boolean if a field has been set.
func (o *ApiV2CustomAttributesGlobalIdPutRequest) HasIsRequired() bool {
	if o != nil && o.IsRequired.IsSet() {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given NullableBool and assigns it to the IsRequired field.
func (o *ApiV2CustomAttributesGlobalIdPutRequest) SetIsRequired(v bool) {
	o.IsRequired.Set(&v)
}
// SetIsRequiredNil sets the value for IsRequired to be an explicit nil
func (o *ApiV2CustomAttributesGlobalIdPutRequest) SetIsRequiredNil() {
	o.IsRequired.Set(nil)
}

// UnsetIsRequired ensures that no value is present for IsRequired, not even an explicit nil
func (o *ApiV2CustomAttributesGlobalIdPutRequest) UnsetIsRequired() {
	o.IsRequired.Unset()
}

func (o ApiV2CustomAttributesGlobalIdPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV2CustomAttributesGlobalIdPutRequest) ToMap() (map[string]interface{}, error) {
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

type NullableApiV2CustomAttributesGlobalIdPutRequest struct {
	value *ApiV2CustomAttributesGlobalIdPutRequest
	isSet bool
}

func (v NullableApiV2CustomAttributesGlobalIdPutRequest) Get() *ApiV2CustomAttributesGlobalIdPutRequest {
	return v.value
}

func (v *NullableApiV2CustomAttributesGlobalIdPutRequest) Set(val *ApiV2CustomAttributesGlobalIdPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2CustomAttributesGlobalIdPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2CustomAttributesGlobalIdPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2CustomAttributesGlobalIdPutRequest(val *ApiV2CustomAttributesGlobalIdPutRequest) *NullableApiV2CustomAttributesGlobalIdPutRequest {
	return &NullableApiV2CustomAttributesGlobalIdPutRequest{value: val, isSet: true}
}

func (v NullableApiV2CustomAttributesGlobalIdPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2CustomAttributesGlobalIdPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


