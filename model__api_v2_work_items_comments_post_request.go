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

// checks if the ApiV2WorkItemsCommentsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV2WorkItemsCommentsPostRequest{}

// ApiV2WorkItemsCommentsPostRequest struct for ApiV2WorkItemsCommentsPostRequest
type ApiV2WorkItemsCommentsPostRequest struct {
	Text string `json:"text"`
	WorkItemId string `json:"workItemId"`
}

// NewApiV2WorkItemsCommentsPostRequest instantiates a new ApiV2WorkItemsCommentsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2WorkItemsCommentsPostRequest(text string, workItemId string) *ApiV2WorkItemsCommentsPostRequest {
	this := ApiV2WorkItemsCommentsPostRequest{}
	this.Text = text
	this.WorkItemId = workItemId
	return &this
}

// NewApiV2WorkItemsCommentsPostRequestWithDefaults instantiates a new ApiV2WorkItemsCommentsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2WorkItemsCommentsPostRequestWithDefaults() *ApiV2WorkItemsCommentsPostRequest {
	this := ApiV2WorkItemsCommentsPostRequest{}
	return &this
}

// GetText returns the Text field value
func (o *ApiV2WorkItemsCommentsPostRequest) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *ApiV2WorkItemsCommentsPostRequest) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *ApiV2WorkItemsCommentsPostRequest) SetText(v string) {
	o.Text = v
}

// GetWorkItemId returns the WorkItemId field value
func (o *ApiV2WorkItemsCommentsPostRequest) GetWorkItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkItemId
}

// GetWorkItemIdOk returns a tuple with the WorkItemId field value
// and a boolean to check if the value has been set.
func (o *ApiV2WorkItemsCommentsPostRequest) GetWorkItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkItemId, true
}

// SetWorkItemId sets field value
func (o *ApiV2WorkItemsCommentsPostRequest) SetWorkItemId(v string) {
	o.WorkItemId = v
}

func (o ApiV2WorkItemsCommentsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV2WorkItemsCommentsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["text"] = o.Text
	toSerialize["workItemId"] = o.WorkItemId
	return toSerialize, nil
}

type NullableApiV2WorkItemsCommentsPostRequest struct {
	value *ApiV2WorkItemsCommentsPostRequest
	isSet bool
}

func (v NullableApiV2WorkItemsCommentsPostRequest) Get() *ApiV2WorkItemsCommentsPostRequest {
	return v.value
}

func (v *NullableApiV2WorkItemsCommentsPostRequest) Set(val *ApiV2WorkItemsCommentsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2WorkItemsCommentsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2WorkItemsCommentsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2WorkItemsCommentsPostRequest(val *ApiV2WorkItemsCommentsPostRequest) *NullableApiV2WorkItemsCommentsPostRequest {
	return &NullableApiV2WorkItemsCommentsPostRequest{value: val, isSet: true}
}

func (v NullableApiV2WorkItemsCommentsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2WorkItemsCommentsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


