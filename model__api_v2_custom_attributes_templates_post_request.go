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

// checks if the ApiV2CustomAttributesTemplatesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV2CustomAttributesTemplatesPostRequest{}

// ApiV2CustomAttributesTemplatesPostRequest struct for ApiV2CustomAttributesTemplatesPostRequest
type ApiV2CustomAttributesTemplatesPostRequest struct {
	// Collection of attribute IDs
	CustomAttributeIds []string `json:"customAttributeIds,omitempty"`
	// Custom attributes template name
	Name string `json:"name"`
}

// NewApiV2CustomAttributesTemplatesPostRequest instantiates a new ApiV2CustomAttributesTemplatesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2CustomAttributesTemplatesPostRequest(name string) *ApiV2CustomAttributesTemplatesPostRequest {
	this := ApiV2CustomAttributesTemplatesPostRequest{}
	this.Name = name
	return &this
}

// NewApiV2CustomAttributesTemplatesPostRequestWithDefaults instantiates a new ApiV2CustomAttributesTemplatesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2CustomAttributesTemplatesPostRequestWithDefaults() *ApiV2CustomAttributesTemplatesPostRequest {
	this := ApiV2CustomAttributesTemplatesPostRequest{}
	return &this
}

// GetCustomAttributeIds returns the CustomAttributeIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2CustomAttributesTemplatesPostRequest) GetCustomAttributeIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CustomAttributeIds
}

// GetCustomAttributeIdsOk returns a tuple with the CustomAttributeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2CustomAttributesTemplatesPostRequest) GetCustomAttributeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CustomAttributeIds) {
		return nil, false
	}
	return o.CustomAttributeIds, true
}

// HasCustomAttributeIds returns a boolean if a field has been set.
func (o *ApiV2CustomAttributesTemplatesPostRequest) HasCustomAttributeIds() bool {
	if o != nil && IsNil(o.CustomAttributeIds) {
		return true
	}

	return false
}

// SetCustomAttributeIds gets a reference to the given []string and assigns it to the CustomAttributeIds field.
func (o *ApiV2CustomAttributesTemplatesPostRequest) SetCustomAttributeIds(v []string) {
	o.CustomAttributeIds = v
}

// GetName returns the Name field value
func (o *ApiV2CustomAttributesTemplatesPostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiV2CustomAttributesTemplatesPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiV2CustomAttributesTemplatesPostRequest) SetName(v string) {
	o.Name = v
}

func (o ApiV2CustomAttributesTemplatesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV2CustomAttributesTemplatesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomAttributeIds != nil {
		toSerialize["customAttributeIds"] = o.CustomAttributeIds
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableApiV2CustomAttributesTemplatesPostRequest struct {
	value *ApiV2CustomAttributesTemplatesPostRequest
	isSet bool
}

func (v NullableApiV2CustomAttributesTemplatesPostRequest) Get() *ApiV2CustomAttributesTemplatesPostRequest {
	return v.value
}

func (v *NullableApiV2CustomAttributesTemplatesPostRequest) Set(val *ApiV2CustomAttributesTemplatesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2CustomAttributesTemplatesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2CustomAttributesTemplatesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2CustomAttributesTemplatesPostRequest(val *ApiV2CustomAttributesTemplatesPostRequest) *NullableApiV2CustomAttributesTemplatesPostRequest {
	return &NullableApiV2CustomAttributesTemplatesPostRequest{value: val, isSet: true}
}

func (v NullableApiV2CustomAttributesTemplatesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2CustomAttributesTemplatesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

