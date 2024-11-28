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

// checks if the WebhooksFilterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhooksFilterRequest{}

// WebhooksFilterRequest struct for WebhooksFilterRequest
type WebhooksFilterRequest struct {
	// Specifies a webhook name to search for
	Name NullableString `json:"name,omitempty"`
	// Specifies a webhook event types to search for
	EventTypes []WebHookEventTypeRequest `json:"eventTypes,omitempty"`
	// Specifies a webhook methods to search for
	Methods []RequestTypeRequest `json:"methods,omitempty"`
	// Specifies a webhook project IDs to search for
	ProjectIds []string `json:"projectIds,omitempty"`
}

// NewWebhooksFilterRequest instantiates a new WebhooksFilterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhooksFilterRequest() *WebhooksFilterRequest {
	this := WebhooksFilterRequest{}
	return &this
}

// NewWebhooksFilterRequestWithDefaults instantiates a new WebhooksFilterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhooksFilterRequestWithDefaults() *WebhooksFilterRequest {
	this := WebhooksFilterRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksFilterRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksFilterRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *WebhooksFilterRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *WebhooksFilterRequest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *WebhooksFilterRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *WebhooksFilterRequest) UnsetName() {
	o.Name.Unset()
}

// GetEventTypes returns the EventTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksFilterRequest) GetEventTypes() []WebHookEventTypeRequest {
	if o == nil {
		var ret []WebHookEventTypeRequest
		return ret
	}
	return o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksFilterRequest) GetEventTypesOk() ([]WebHookEventTypeRequest, bool) {
	if o == nil || IsNil(o.EventTypes) {
		return nil, false
	}
	return o.EventTypes, true
}

// HasEventTypes returns a boolean if a field has been set.
func (o *WebhooksFilterRequest) HasEventTypes() bool {
	if o != nil && !IsNil(o.EventTypes) {
		return true
	}

	return false
}

// SetEventTypes gets a reference to the given []WebHookEventTypeRequest and assigns it to the EventTypes field.
func (o *WebhooksFilterRequest) SetEventTypes(v []WebHookEventTypeRequest) {
	o.EventTypes = v
}

// GetMethods returns the Methods field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksFilterRequest) GetMethods() []RequestTypeRequest {
	if o == nil {
		var ret []RequestTypeRequest
		return ret
	}
	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksFilterRequest) GetMethodsOk() ([]RequestTypeRequest, bool) {
	if o == nil || IsNil(o.Methods) {
		return nil, false
	}
	return o.Methods, true
}

// HasMethods returns a boolean if a field has been set.
func (o *WebhooksFilterRequest) HasMethods() bool {
	if o != nil && !IsNil(o.Methods) {
		return true
	}

	return false
}

// SetMethods gets a reference to the given []RequestTypeRequest and assigns it to the Methods field.
func (o *WebhooksFilterRequest) SetMethods(v []RequestTypeRequest) {
	o.Methods = v
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhooksFilterRequest) GetProjectIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhooksFilterRequest) GetProjectIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectIds) {
		return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *WebhooksFilterRequest) HasProjectIds() bool {
	if o != nil && !IsNil(o.ProjectIds) {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []string and assigns it to the ProjectIds field.
func (o *WebhooksFilterRequest) SetProjectIds(v []string) {
	o.ProjectIds = v
}

func (o WebhooksFilterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhooksFilterRequest) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableWebhooksFilterRequest struct {
	value *WebhooksFilterRequest
	isSet bool
}

func (v NullableWebhooksFilterRequest) Get() *WebhooksFilterRequest {
	return v.value
}

func (v *NullableWebhooksFilterRequest) Set(val *WebhooksFilterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhooksFilterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhooksFilterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhooksFilterRequest(val *WebhooksFilterRequest) *NullableWebhooksFilterRequest {
	return &NullableWebhooksFilterRequest{value: val, isSet: true}
}

func (v NullableWebhooksFilterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhooksFilterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

