/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the WebhooksUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhooksUpdateRequest{}

// WebhooksUpdateRequest struct for WebhooksUpdateRequest
type WebhooksUpdateRequest struct {
	Filter WebhooksFilterRequest `json:"filter"`
	Model WebhookBulkUpdateApiModel `json:"model"`
	Extractor WebhooksExtractionRequest `json:"extractor"`
}

type _WebhooksUpdateRequest WebhooksUpdateRequest

// NewWebhooksUpdateRequest instantiates a new WebhooksUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhooksUpdateRequest(filter WebhooksFilterRequest, model WebhookBulkUpdateApiModel, extractor WebhooksExtractionRequest) *WebhooksUpdateRequest {
	this := WebhooksUpdateRequest{}
	this.Filter = filter
	this.Model = model
	this.Extractor = extractor
	return &this
}

// NewWebhooksUpdateRequestWithDefaults instantiates a new WebhooksUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhooksUpdateRequestWithDefaults() *WebhooksUpdateRequest {
	this := WebhooksUpdateRequest{}
	return &this
}

// GetFilter returns the Filter field value
func (o *WebhooksUpdateRequest) GetFilter() WebhooksFilterRequest {
	if o == nil {
		var ret WebhooksFilterRequest
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *WebhooksUpdateRequest) GetFilterOk() (*WebhooksFilterRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *WebhooksUpdateRequest) SetFilter(v WebhooksFilterRequest) {
	o.Filter = v
}

// GetModel returns the Model field value
func (o *WebhooksUpdateRequest) GetModel() WebhookBulkUpdateApiModel {
	if o == nil {
		var ret WebhookBulkUpdateApiModel
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *WebhooksUpdateRequest) GetModelOk() (*WebhookBulkUpdateApiModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *WebhooksUpdateRequest) SetModel(v WebhookBulkUpdateApiModel) {
	o.Model = v
}

// GetExtractor returns the Extractor field value
func (o *WebhooksUpdateRequest) GetExtractor() WebhooksExtractionRequest {
	if o == nil {
		var ret WebhooksExtractionRequest
		return ret
	}

	return o.Extractor
}

// GetExtractorOk returns a tuple with the Extractor field value
// and a boolean to check if the value has been set.
func (o *WebhooksUpdateRequest) GetExtractorOk() (*WebhooksExtractionRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Extractor, true
}

// SetExtractor sets field value
func (o *WebhooksUpdateRequest) SetExtractor(v WebhooksExtractionRequest) {
	o.Extractor = v
}

func (o WebhooksUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhooksUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filter"] = o.Filter
	toSerialize["model"] = o.Model
	toSerialize["extractor"] = o.Extractor
	return toSerialize, nil
}

func (o *WebhooksUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filter",
		"model",
		"extractor",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWebhooksUpdateRequest := _WebhooksUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhooksUpdateRequest)

	if err != nil {
		return err
	}

	*o = WebhooksUpdateRequest(varWebhooksUpdateRequest)

	return err
}

type NullableWebhooksUpdateRequest struct {
	value *WebhooksUpdateRequest
	isSet bool
}

func (v NullableWebhooksUpdateRequest) Get() *WebhooksUpdateRequest {
	return v.value
}

func (v *NullableWebhooksUpdateRequest) Set(val *WebhooksUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhooksUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhooksUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhooksUpdateRequest(val *WebhooksUpdateRequest) *NullableWebhooksUpdateRequest {
	return &NullableWebhooksUpdateRequest{value: val, isSet: true}
}

func (v NullableWebhooksUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhooksUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

