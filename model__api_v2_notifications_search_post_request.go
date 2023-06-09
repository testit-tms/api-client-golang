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

// checks if the ApiV2NotificationsSearchPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV2NotificationsSearchPostRequest{}

// ApiV2NotificationsSearchPostRequest struct for ApiV2NotificationsSearchPostRequest
type ApiV2NotificationsSearchPostRequest struct {
	Types []NotificationTypeModel `json:"types,omitempty"`
	IsRead NullableBool `json:"isRead,omitempty"`
	CreatedDate NullableDateTimeRangeSelectorModel `json:"createdDate,omitempty"`
}

// NewApiV2NotificationsSearchPostRequest instantiates a new ApiV2NotificationsSearchPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2NotificationsSearchPostRequest() *ApiV2NotificationsSearchPostRequest {
	this := ApiV2NotificationsSearchPostRequest{}
	return &this
}

// NewApiV2NotificationsSearchPostRequestWithDefaults instantiates a new ApiV2NotificationsSearchPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2NotificationsSearchPostRequestWithDefaults() *ApiV2NotificationsSearchPostRequest {
	this := ApiV2NotificationsSearchPostRequest{}
	return &this
}

// GetTypes returns the Types field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2NotificationsSearchPostRequest) GetTypes() []NotificationTypeModel {
	if o == nil {
		var ret []NotificationTypeModel
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2NotificationsSearchPostRequest) GetTypesOk() ([]NotificationTypeModel, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *ApiV2NotificationsSearchPostRequest) HasTypes() bool {
	if o != nil && IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []NotificationTypeModel and assigns it to the Types field.
func (o *ApiV2NotificationsSearchPostRequest) SetTypes(v []NotificationTypeModel) {
	o.Types = v
}

// GetIsRead returns the IsRead field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2NotificationsSearchPostRequest) GetIsRead() bool {
	if o == nil || IsNil(o.IsRead.Get()) {
		var ret bool
		return ret
	}
	return *o.IsRead.Get()
}

// GetIsReadOk returns a tuple with the IsRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2NotificationsSearchPostRequest) GetIsReadOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRead.Get(), o.IsRead.IsSet()
}

// HasIsRead returns a boolean if a field has been set.
func (o *ApiV2NotificationsSearchPostRequest) HasIsRead() bool {
	if o != nil && o.IsRead.IsSet() {
		return true
	}

	return false
}

// SetIsRead gets a reference to the given NullableBool and assigns it to the IsRead field.
func (o *ApiV2NotificationsSearchPostRequest) SetIsRead(v bool) {
	o.IsRead.Set(&v)
}
// SetIsReadNil sets the value for IsRead to be an explicit nil
func (o *ApiV2NotificationsSearchPostRequest) SetIsReadNil() {
	o.IsRead.Set(nil)
}

// UnsetIsRead ensures that no value is present for IsRead, not even an explicit nil
func (o *ApiV2NotificationsSearchPostRequest) UnsetIsRead() {
	o.IsRead.Unset()
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV2NotificationsSearchPostRequest) GetCreatedDate() DateTimeRangeSelectorModel {
	if o == nil || IsNil(o.CreatedDate.Get()) {
		var ret DateTimeRangeSelectorModel
		return ret
	}
	return *o.CreatedDate.Get()
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV2NotificationsSearchPostRequest) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedDate.Get(), o.CreatedDate.IsSet()
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ApiV2NotificationsSearchPostRequest) HasCreatedDate() bool {
	if o != nil && o.CreatedDate.IsSet() {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given NullableDateTimeRangeSelectorModel and assigns it to the CreatedDate field.
func (o *ApiV2NotificationsSearchPostRequest) SetCreatedDate(v DateTimeRangeSelectorModel) {
	o.CreatedDate.Set(&v)
}
// SetCreatedDateNil sets the value for CreatedDate to be an explicit nil
func (o *ApiV2NotificationsSearchPostRequest) SetCreatedDateNil() {
	o.CreatedDate.Set(nil)
}

// UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
func (o *ApiV2NotificationsSearchPostRequest) UnsetCreatedDate() {
	o.CreatedDate.Unset()
}

func (o ApiV2NotificationsSearchPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV2NotificationsSearchPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	if o.IsRead.IsSet() {
		toSerialize["isRead"] = o.IsRead.Get()
	}
	if o.CreatedDate.IsSet() {
		toSerialize["createdDate"] = o.CreatedDate.Get()
	}
	return toSerialize, nil
}

type NullableApiV2NotificationsSearchPostRequest struct {
	value *ApiV2NotificationsSearchPostRequest
	isSet bool
}

func (v NullableApiV2NotificationsSearchPostRequest) Get() *ApiV2NotificationsSearchPostRequest {
	return v.value
}

func (v *NullableApiV2NotificationsSearchPostRequest) Set(val *ApiV2NotificationsSearchPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2NotificationsSearchPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2NotificationsSearchPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2NotificationsSearchPostRequest(val *ApiV2NotificationsSearchPostRequest) *NullableApiV2NotificationsSearchPostRequest {
	return &NullableApiV2NotificationsSearchPostRequest{value: val, isSet: true}
}

func (v NullableApiV2NotificationsSearchPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2NotificationsSearchPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


