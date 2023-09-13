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

// checks if the ApiV2WebhooksSearchPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV2WebhooksSearchPostRequest{}

// ApiV2WebhooksSearchPostRequest struct for ApiV2WebhooksSearchPostRequest
type ApiV2WebhooksSearchPostRequest struct {
	// Specifies a webhook name to search for
	Name NullableString `json:"name,omitempty"`
	// Specifies a webhook event types to search for
	EventTypes []WebHookEventTypeModel `json:"eventTypes,omitempty"`
	// Specifies a webhook methods to search for
	Methods []RequestTypeModel `json:"methods,omitempty"`
	// Specifies a webhook project IDs to search for
	ProjectIds []string `json:"projectIds,omitempty"`
	// Specifies a webhook deleted status to search for
	IsEnabled NullableBool `json:"isEnabled,omitempty"`
}

// NewApiV2WebhooksSearchPostRequest instantiates a new ApiV2WebhooksSearchPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2WebhooksSearchPostRequest() *ApiV2WebhooksSearchPostRequest {
	this := ApiV2WebhooksSearchPostRequest{}
	return &this
}

// NewApiV2WebhooksSearchPostRequestWithDefaults instantiates a new ApiV2WebhooksSearchPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2WebhooksSearchPostRequestWithDefaults() *ApiV2WebhooksSearchPostRequest {
	this := ApiV2WebhooksSearchPostRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2WebhooksSearchPostRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2WebhooksSearchPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ApiV2WebhooksSearchPostRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ApiV2WebhooksSearchPostRequest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ApiV2WebhooksSearchPostRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ApiV2WebhooksSearchPostRequest) UnsetName() {
	o.Name.Unset()
}

// GetEventTypes returns the EventTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2WebhooksSearchPostRequest) GetEventTypes() []WebHookEventTypeModel {
	if o == nil {
		var ret []WebHookEventTypeModel
		return ret
	}
	return o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2WebhooksSearchPostRequest) GetEventTypesOk() ([]WebHookEventTypeModel, bool) {
	if o == nil || IsNil(o.EventTypes) {
		return nil, false
	}
	return o.EventTypes, true
}

// HasEventTypes returns a boolean if a field has been set.
func (o *ApiV2WebhooksSearchPostRequest) HasEventTypes() bool {
	if o != nil && IsNil(o.EventTypes) {
		return true
	}

	return false
}

// SetEventTypes gets a reference to the given []WebHookEventTypeModel and assigns it to the EventTypes field.
func (o *ApiV2WebhooksSearchPostRequest) SetEventTypes(v []WebHookEventTypeModel) {
	o.EventTypes = v
}

// GetMethods returns the Methods field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2WebhooksSearchPostRequest) GetMethods() []RequestTypeModel {
	if o == nil {
		var ret []RequestTypeModel
		return ret
	}
	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2WebhooksSearchPostRequest) GetMethodsOk() ([]RequestTypeModel, bool) {
	if o == nil || IsNil(o.Methods) {
		return nil, false
	}
	return o.Methods, true
}

// HasMethods returns a boolean if a field has been set.
func (o *ApiV2WebhooksSearchPostRequest) HasMethods() bool {
	if o != nil && IsNil(o.Methods) {
		return true
	}

	return false
}

// SetMethods gets a reference to the given []RequestTypeModel and assigns it to the Methods field.
func (o *ApiV2WebhooksSearchPostRequest) SetMethods(v []RequestTypeModel) {
	o.Methods = v
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2WebhooksSearchPostRequest) GetProjectIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2WebhooksSearchPostRequest) GetProjectIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectIds) {
		return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *ApiV2WebhooksSearchPostRequest) HasProjectIds() bool {
	if o != nil && IsNil(o.ProjectIds) {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []string and assigns it to the ProjectIds field.
func (o *ApiV2WebhooksSearchPostRequest) SetProjectIds(v []string) {
	o.ProjectIds = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2WebhooksSearchPostRequest) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.IsEnabled.Get()
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2WebhooksSearchPostRequest) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsEnabled.Get(), o.IsEnabled.IsSet()
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *ApiV2WebhooksSearchPostRequest) HasIsEnabled() bool {
	if o != nil && o.IsEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given NullableBool and assigns it to the IsEnabled field.
func (o *ApiV2WebhooksSearchPostRequest) SetIsEnabled(v bool) {
	o.IsEnabled.Set(&v)
}
// SetIsEnabledNil sets the value for IsEnabled to be an explicit nil
func (o *ApiV2WebhooksSearchPostRequest) SetIsEnabledNil() {
	o.IsEnabled.Set(nil)
}

// UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
func (o *ApiV2WebhooksSearchPostRequest) UnsetIsEnabled() {
	o.IsEnabled.Unset()
}

func (o ApiV2WebhooksSearchPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV2WebhooksSearchPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.EventTypes != nil {
		toSerialize["eventTypes"] = o.EventTypes
	}
	if o.Methods != nil {
		toSerialize["methods"] = o.Methods
	}
	if o.ProjectIds != nil {
		toSerialize["projectIds"] = o.ProjectIds
	}
	if o.IsEnabled.IsSet() {
		toSerialize["isEnabled"] = o.IsEnabled.Get()
	}
	return toSerialize, nil
}

type NullableApiV2WebhooksSearchPostRequest struct {
	value *ApiV2WebhooksSearchPostRequest
	isSet bool
}

func (v NullableApiV2WebhooksSearchPostRequest) Get() *ApiV2WebhooksSearchPostRequest {
	return v.value
}

func (v *NullableApiV2WebhooksSearchPostRequest) Set(val *ApiV2WebhooksSearchPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2WebhooksSearchPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2WebhooksSearchPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2WebhooksSearchPostRequest(val *ApiV2WebhooksSearchPostRequest) *NullableApiV2WebhooksSearchPostRequest {
	return &NullableApiV2WebhooksSearchPostRequest{value: val, isSet: true}
}

func (v NullableApiV2WebhooksSearchPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2WebhooksSearchPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

