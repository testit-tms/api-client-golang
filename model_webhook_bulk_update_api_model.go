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

// checks if the WebhookBulkUpdateApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookBulkUpdateApiModel{}

// WebhookBulkUpdateApiModel struct for WebhookBulkUpdateApiModel
type WebhookBulkUpdateApiModel struct {
	IsEnabled bool `json:"isEnabled"`
}

type _WebhookBulkUpdateApiModel WebhookBulkUpdateApiModel

// NewWebhookBulkUpdateApiModel instantiates a new WebhookBulkUpdateApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookBulkUpdateApiModel(isEnabled bool) *WebhookBulkUpdateApiModel {
	this := WebhookBulkUpdateApiModel{}
	this.IsEnabled = isEnabled
	return &this
}

// NewWebhookBulkUpdateApiModelWithDefaults instantiates a new WebhookBulkUpdateApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookBulkUpdateApiModelWithDefaults() *WebhookBulkUpdateApiModel {
	this := WebhookBulkUpdateApiModel{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *WebhookBulkUpdateApiModel) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *WebhookBulkUpdateApiModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *WebhookBulkUpdateApiModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

func (o WebhookBulkUpdateApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookBulkUpdateApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isEnabled"] = o.IsEnabled
	return toSerialize, nil
}

func (o *WebhookBulkUpdateApiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isEnabled",
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

	varWebhookBulkUpdateApiModel := _WebhookBulkUpdateApiModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookBulkUpdateApiModel)

	if err != nil {
		return err
	}

	*o = WebhookBulkUpdateApiModel(varWebhookBulkUpdateApiModel)

	return err
}

type NullableWebhookBulkUpdateApiModel struct {
	value *WebhookBulkUpdateApiModel
	isSet bool
}

func (v NullableWebhookBulkUpdateApiModel) Get() *WebhookBulkUpdateApiModel {
	return v.value
}

func (v *NullableWebhookBulkUpdateApiModel) Set(val *WebhookBulkUpdateApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookBulkUpdateApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookBulkUpdateApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookBulkUpdateApiModel(val *WebhookBulkUpdateApiModel) *NullableWebhookBulkUpdateApiModel {
	return &NullableWebhookBulkUpdateApiModel{value: val, isSet: true}
}

func (v NullableWebhookBulkUpdateApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookBulkUpdateApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


