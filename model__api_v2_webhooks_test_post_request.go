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

// checks if the ApiV2WebhooksTestPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV2WebhooksTestPostRequest{}

// ApiV2WebhooksTestPostRequest struct for ApiV2WebhooksTestPostRequest
type ApiV2WebhooksTestPostRequest struct {
	RequestType RequestTypeModel `json:"requestType"`
	// Request URL of the webhook
	Url string `json:"url"`
}

// NewApiV2WebhooksTestPostRequest instantiates a new ApiV2WebhooksTestPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV2WebhooksTestPostRequest(requestType RequestTypeModel, url string) *ApiV2WebhooksTestPostRequest {
	this := ApiV2WebhooksTestPostRequest{}
	this.RequestType = requestType
	this.Url = url
	return &this
}

// NewApiV2WebhooksTestPostRequestWithDefaults instantiates a new ApiV2WebhooksTestPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV2WebhooksTestPostRequestWithDefaults() *ApiV2WebhooksTestPostRequest {
	this := ApiV2WebhooksTestPostRequest{}
	return &this
}

// GetRequestType returns the RequestType field value
func (o *ApiV2WebhooksTestPostRequest) GetRequestType() RequestTypeModel {
	if o == nil {
		var ret RequestTypeModel
		return ret
	}

	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *ApiV2WebhooksTestPostRequest) GetRequestTypeOk() (*RequestTypeModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value
func (o *ApiV2WebhooksTestPostRequest) SetRequestType(v RequestTypeModel) {
	o.RequestType = v
}

// GetUrl returns the Url field value
func (o *ApiV2WebhooksTestPostRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ApiV2WebhooksTestPostRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ApiV2WebhooksTestPostRequest) SetUrl(v string) {
	o.Url = v
}

func (o ApiV2WebhooksTestPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV2WebhooksTestPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requestType"] = o.RequestType
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableApiV2WebhooksTestPostRequest struct {
	value *ApiV2WebhooksTestPostRequest
	isSet bool
}

func (v NullableApiV2WebhooksTestPostRequest) Get() *ApiV2WebhooksTestPostRequest {
	return v.value
}

func (v *NullableApiV2WebhooksTestPostRequest) Set(val *ApiV2WebhooksTestPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV2WebhooksTestPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV2WebhooksTestPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV2WebhooksTestPostRequest(val *ApiV2WebhooksTestPostRequest) *NullableApiV2WebhooksTestPostRequest {
	return &NullableApiV2WebhooksTestPostRequest{value: val, isSet: true}
}

func (v NullableApiV2WebhooksTestPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV2WebhooksTestPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


