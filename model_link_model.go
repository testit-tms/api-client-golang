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

// checks if the LinkModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkModel{}

// LinkModel struct for LinkModel
type LinkModel struct {
	Id *string `json:"id,omitempty"`
	// Link name.
	Title NullableString `json:"title,omitempty"`
	// Address can be specified without protocol, but necessarily with the domain.
	Url string `json:"url"`
	// Link description.
	Description NullableString `json:"description,omitempty"`
	Type NullableLinkType `json:"type,omitempty"`
	HasInfo *bool `json:"hasInfo,omitempty"`
}

// NewLinkModel instantiates a new LinkModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkModel(url string) *LinkModel {
	this := LinkModel{}
	this.Url = url
	return &this
}

// NewLinkModelWithDefaults instantiates a new LinkModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkModelWithDefaults() *LinkModel {
	this := LinkModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LinkModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LinkModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LinkModel) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LinkModel) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkModel) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *LinkModel) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *LinkModel) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *LinkModel) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *LinkModel) UnsetTitle() {
	o.Title.Unset()
}

// GetUrl returns the Url field value
func (o *LinkModel) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *LinkModel) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *LinkModel) SetUrl(v string) {
	o.Url = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LinkModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *LinkModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *LinkModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *LinkModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *LinkModel) UnsetDescription() {
	o.Description.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LinkModel) GetType() LinkType {
	if o == nil || IsNil(o.Type.Get()) {
		var ret LinkType
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkModel) GetTypeOk() (*LinkType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *LinkModel) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableLinkType and assigns it to the Type field.
func (o *LinkModel) SetType(v LinkType) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *LinkModel) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *LinkModel) UnsetType() {
	o.Type.Unset()
}

// GetHasInfo returns the HasInfo field value if set, zero value otherwise.
func (o *LinkModel) GetHasInfo() bool {
	if o == nil || IsNil(o.HasInfo) {
		var ret bool
		return ret
	}
	return *o.HasInfo
}

// GetHasInfoOk returns a tuple with the HasInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkModel) GetHasInfoOk() (*bool, bool) {
	if o == nil || IsNil(o.HasInfo) {
		return nil, false
	}
	return o.HasInfo, true
}

// HasHasInfo returns a boolean if a field has been set.
func (o *LinkModel) HasHasInfo() bool {
	if o != nil && !IsNil(o.HasInfo) {
		return true
	}

	return false
}

// SetHasInfo gets a reference to the given bool and assigns it to the HasInfo field.
func (o *LinkModel) SetHasInfo(v bool) {
	o.HasInfo = &v
}

func (o LinkModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	toSerialize["url"] = o.Url
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if !IsNil(o.HasInfo) {
		toSerialize["hasInfo"] = o.HasInfo
	}
	return toSerialize, nil
}

type NullableLinkModel struct {
	value *LinkModel
	isSet bool
}

func (v NullableLinkModel) Get() *LinkModel {
	return v.value
}

func (v *NullableLinkModel) Set(val *LinkModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkModel(val *LinkModel) *NullableLinkModel {
	return &NullableLinkModel{value: val, isSet: true}
}

func (v NullableLinkModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


