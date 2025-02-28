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

// checks if the LinkShortApiResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkShortApiResult{}

// LinkShortApiResult struct for LinkShortApiResult
type LinkShortApiResult struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Url string `json:"url"`
	Type string `json:"type"`
}

type _LinkShortApiResult LinkShortApiResult

// NewLinkShortApiResult instantiates a new LinkShortApiResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkShortApiResult(id string, title string, url string, type_ string) *LinkShortApiResult {
	this := LinkShortApiResult{}
	this.Id = id
	this.Title = title
	this.Url = url
	this.Type = type_
	return &this
}

// NewLinkShortApiResultWithDefaults instantiates a new LinkShortApiResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkShortApiResultWithDefaults() *LinkShortApiResult {
	this := LinkShortApiResult{}
	return &this
}

// GetId returns the Id field value
func (o *LinkShortApiResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LinkShortApiResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LinkShortApiResult) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *LinkShortApiResult) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *LinkShortApiResult) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *LinkShortApiResult) SetTitle(v string) {
	o.Title = v
}

// GetUrl returns the Url field value
func (o *LinkShortApiResult) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *LinkShortApiResult) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *LinkShortApiResult) SetUrl(v string) {
	o.Url = v
}

// GetType returns the Type field value
func (o *LinkShortApiResult) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LinkShortApiResult) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LinkShortApiResult) SetType(v string) {
	o.Type = v
}

func (o LinkShortApiResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkShortApiResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title
	toSerialize["url"] = o.Url
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *LinkShortApiResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"title",
		"url",
		"type",
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

	varLinkShortApiResult := _LinkShortApiResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLinkShortApiResult)

	if err != nil {
		return err
	}

	*o = LinkShortApiResult(varLinkShortApiResult)

	return err
}

type NullableLinkShortApiResult struct {
	value *LinkShortApiResult
	isSet bool
}

func (v NullableLinkShortApiResult) Get() *LinkShortApiResult {
	return v.value
}

func (v *NullableLinkShortApiResult) Set(val *LinkShortApiResult) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkShortApiResult) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkShortApiResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkShortApiResult(val *LinkShortApiResult) *NullableLinkShortApiResult {
	return &NullableLinkShortApiResult{value: val, isSet: true}
}

func (v NullableLinkShortApiResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkShortApiResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


