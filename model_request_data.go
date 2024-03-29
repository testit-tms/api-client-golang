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

// checks if the RequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestData{}

// RequestData struct for RequestData
type RequestData struct {
	Uri NullableString `json:"uri,omitempty"`
	StatusCode int32 `json:"statusCode"`
	RequestBody NullableString `json:"requestBody,omitempty"`
	RequestMeta string `json:"requestMeta"`
	ResponseBody string `json:"responseBody"`
	ResponseMeta string `json:"responseMeta"`
}

// NewRequestData instantiates a new RequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestData(statusCode int32, requestMeta string, responseBody string, responseMeta string) *RequestData {
	this := RequestData{}
	this.StatusCode = statusCode
	this.RequestMeta = requestMeta
	this.ResponseBody = responseBody
	this.ResponseMeta = responseMeta
	return &this
}

// NewRequestDataWithDefaults instantiates a new RequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestDataWithDefaults() *RequestData {
	this := RequestData{}
	return &this
}

// GetUri returns the Uri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestData) GetUri() string {
	if o == nil || IsNil(o.Uri.Get()) {
		var ret string
		return ret
	}
	return *o.Uri.Get()
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestData) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uri.Get(), o.Uri.IsSet()
}

// HasUri returns a boolean if a field has been set.
func (o *RequestData) HasUri() bool {
	if o != nil && o.Uri.IsSet() {
		return true
	}

	return false
}

// SetUri gets a reference to the given NullableString and assigns it to the Uri field.
func (o *RequestData) SetUri(v string) {
	o.Uri.Set(&v)
}
// SetUriNil sets the value for Uri to be an explicit nil
func (o *RequestData) SetUriNil() {
	o.Uri.Set(nil)
}

// UnsetUri ensures that no value is present for Uri, not even an explicit nil
func (o *RequestData) UnsetUri() {
	o.Uri.Unset()
}

// GetStatusCode returns the StatusCode field value
func (o *RequestData) GetStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *RequestData) GetStatusCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *RequestData) SetStatusCode(v int32) {
	o.StatusCode = v
}

// GetRequestBody returns the RequestBody field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestData) GetRequestBody() string {
	if o == nil || IsNil(o.RequestBody.Get()) {
		var ret string
		return ret
	}
	return *o.RequestBody.Get()
}

// GetRequestBodyOk returns a tuple with the RequestBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestData) GetRequestBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestBody.Get(), o.RequestBody.IsSet()
}

// HasRequestBody returns a boolean if a field has been set.
func (o *RequestData) HasRequestBody() bool {
	if o != nil && o.RequestBody.IsSet() {
		return true
	}

	return false
}

// SetRequestBody gets a reference to the given NullableString and assigns it to the RequestBody field.
func (o *RequestData) SetRequestBody(v string) {
	o.RequestBody.Set(&v)
}
// SetRequestBodyNil sets the value for RequestBody to be an explicit nil
func (o *RequestData) SetRequestBodyNil() {
	o.RequestBody.Set(nil)
}

// UnsetRequestBody ensures that no value is present for RequestBody, not even an explicit nil
func (o *RequestData) UnsetRequestBody() {
	o.RequestBody.Unset()
}

// GetRequestMeta returns the RequestMeta field value
func (o *RequestData) GetRequestMeta() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestMeta
}

// GetRequestMetaOk returns a tuple with the RequestMeta field value
// and a boolean to check if the value has been set.
func (o *RequestData) GetRequestMetaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestMeta, true
}

// SetRequestMeta sets field value
func (o *RequestData) SetRequestMeta(v string) {
	o.RequestMeta = v
}

// GetResponseBody returns the ResponseBody field value
func (o *RequestData) GetResponseBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseBody
}

// GetResponseBodyOk returns a tuple with the ResponseBody field value
// and a boolean to check if the value has been set.
func (o *RequestData) GetResponseBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseBody, true
}

// SetResponseBody sets field value
func (o *RequestData) SetResponseBody(v string) {
	o.ResponseBody = v
}

// GetResponseMeta returns the ResponseMeta field value
func (o *RequestData) GetResponseMeta() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseMeta
}

// GetResponseMetaOk returns a tuple with the ResponseMeta field value
// and a boolean to check if the value has been set.
func (o *RequestData) GetResponseMetaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseMeta, true
}

// SetResponseMeta sets field value
func (o *RequestData) SetResponseMeta(v string) {
	o.ResponseMeta = v
}

func (o RequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Uri.IsSet() {
		toSerialize["uri"] = o.Uri.Get()
	}
	toSerialize["statusCode"] = o.StatusCode
	if o.RequestBody.IsSet() {
		toSerialize["requestBody"] = o.RequestBody.Get()
	}
	toSerialize["requestMeta"] = o.RequestMeta
	toSerialize["responseBody"] = o.ResponseBody
	toSerialize["responseMeta"] = o.ResponseMeta
	return toSerialize, nil
}

type NullableRequestData struct {
	value *RequestData
	isSet bool
}

func (v NullableRequestData) Get() *RequestData {
	return v.value
}

func (v *NullableRequestData) Set(val *RequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestData(val *RequestData) *NullableRequestData {
	return &NullableRequestData{value: val, isSet: true}
}

func (v NullableRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


