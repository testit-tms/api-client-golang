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

// checks if the TestRunUpdateMultipleModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestRunUpdateMultipleModel{}

// TestRunUpdateMultipleModel struct for TestRunUpdateMultipleModel
type TestRunUpdateMultipleModel struct {
	SelectModel TestRunSelectionModel `json:"selectModel"`
	Description NullableString `json:"description,omitempty"`
	AttachmentUpdateScheme NullableSetOfAttachmentIds `json:"attachmentUpdateScheme,omitempty"`
	LinkUpdateScheme NullableSetOfLinks `json:"linkUpdateScheme,omitempty"`
}

// NewTestRunUpdateMultipleModel instantiates a new TestRunUpdateMultipleModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestRunUpdateMultipleModel(selectModel TestRunSelectionModel) *TestRunUpdateMultipleModel {
	this := TestRunUpdateMultipleModel{}
	this.SelectModel = selectModel
	return &this
}

// NewTestRunUpdateMultipleModelWithDefaults instantiates a new TestRunUpdateMultipleModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestRunUpdateMultipleModelWithDefaults() *TestRunUpdateMultipleModel {
	this := TestRunUpdateMultipleModel{}
	return &this
}

// GetSelectModel returns the SelectModel field value
func (o *TestRunUpdateMultipleModel) GetSelectModel() TestRunSelectionModel {
	if o == nil {
		var ret TestRunSelectionModel
		return ret
	}

	return o.SelectModel
}

// GetSelectModelOk returns a tuple with the SelectModel field value
// and a boolean to check if the value has been set.
func (o *TestRunUpdateMultipleModel) GetSelectModelOk() (*TestRunSelectionModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectModel, true
}

// SetSelectModel sets field value
func (o *TestRunUpdateMultipleModel) SetSelectModel(v TestRunSelectionModel) {
	o.SelectModel = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunUpdateMultipleModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunUpdateMultipleModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TestRunUpdateMultipleModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TestRunUpdateMultipleModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TestRunUpdateMultipleModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TestRunUpdateMultipleModel) UnsetDescription() {
	o.Description.Unset()
}

// GetAttachmentUpdateScheme returns the AttachmentUpdateScheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunUpdateMultipleModel) GetAttachmentUpdateScheme() SetOfAttachmentIds {
	if o == nil || IsNil(o.AttachmentUpdateScheme.Get()) {
		var ret SetOfAttachmentIds
		return ret
	}
	return *o.AttachmentUpdateScheme.Get()
}

// GetAttachmentUpdateSchemeOk returns a tuple with the AttachmentUpdateScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunUpdateMultipleModel) GetAttachmentUpdateSchemeOk() (*SetOfAttachmentIds, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachmentUpdateScheme.Get(), o.AttachmentUpdateScheme.IsSet()
}

// HasAttachmentUpdateScheme returns a boolean if a field has been set.
func (o *TestRunUpdateMultipleModel) HasAttachmentUpdateScheme() bool {
	if o != nil && o.AttachmentUpdateScheme.IsSet() {
		return true
	}

	return false
}

// SetAttachmentUpdateScheme gets a reference to the given NullableSetOfAttachmentIds and assigns it to the AttachmentUpdateScheme field.
func (o *TestRunUpdateMultipleModel) SetAttachmentUpdateScheme(v SetOfAttachmentIds) {
	o.AttachmentUpdateScheme.Set(&v)
}
// SetAttachmentUpdateSchemeNil sets the value for AttachmentUpdateScheme to be an explicit nil
func (o *TestRunUpdateMultipleModel) SetAttachmentUpdateSchemeNil() {
	o.AttachmentUpdateScheme.Set(nil)
}

// UnsetAttachmentUpdateScheme ensures that no value is present for AttachmentUpdateScheme, not even an explicit nil
func (o *TestRunUpdateMultipleModel) UnsetAttachmentUpdateScheme() {
	o.AttachmentUpdateScheme.Unset()
}

// GetLinkUpdateScheme returns the LinkUpdateScheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRunUpdateMultipleModel) GetLinkUpdateScheme() SetOfLinks {
	if o == nil || IsNil(o.LinkUpdateScheme.Get()) {
		var ret SetOfLinks
		return ret
	}
	return *o.LinkUpdateScheme.Get()
}

// GetLinkUpdateSchemeOk returns a tuple with the LinkUpdateScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRunUpdateMultipleModel) GetLinkUpdateSchemeOk() (*SetOfLinks, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkUpdateScheme.Get(), o.LinkUpdateScheme.IsSet()
}

// HasLinkUpdateScheme returns a boolean if a field has been set.
func (o *TestRunUpdateMultipleModel) HasLinkUpdateScheme() bool {
	if o != nil && o.LinkUpdateScheme.IsSet() {
		return true
	}

	return false
}

// SetLinkUpdateScheme gets a reference to the given NullableSetOfLinks and assigns it to the LinkUpdateScheme field.
func (o *TestRunUpdateMultipleModel) SetLinkUpdateScheme(v SetOfLinks) {
	o.LinkUpdateScheme.Set(&v)
}
// SetLinkUpdateSchemeNil sets the value for LinkUpdateScheme to be an explicit nil
func (o *TestRunUpdateMultipleModel) SetLinkUpdateSchemeNil() {
	o.LinkUpdateScheme.Set(nil)
}

// UnsetLinkUpdateScheme ensures that no value is present for LinkUpdateScheme, not even an explicit nil
func (o *TestRunUpdateMultipleModel) UnsetLinkUpdateScheme() {
	o.LinkUpdateScheme.Unset()
}

func (o TestRunUpdateMultipleModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestRunUpdateMultipleModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["selectModel"] = o.SelectModel
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.AttachmentUpdateScheme.IsSet() {
		toSerialize["attachmentUpdateScheme"] = o.AttachmentUpdateScheme.Get()
	}
	if o.LinkUpdateScheme.IsSet() {
		toSerialize["linkUpdateScheme"] = o.LinkUpdateScheme.Get()
	}
	return toSerialize, nil
}

type NullableTestRunUpdateMultipleModel struct {
	value *TestRunUpdateMultipleModel
	isSet bool
}

func (v NullableTestRunUpdateMultipleModel) Get() *TestRunUpdateMultipleModel {
	return v.value
}

func (v *NullableTestRunUpdateMultipleModel) Set(val *TestRunUpdateMultipleModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestRunUpdateMultipleModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestRunUpdateMultipleModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestRunUpdateMultipleModel(val *TestRunUpdateMultipleModel) *NullableTestRunUpdateMultipleModel {
	return &NullableTestRunUpdateMultipleModel{value: val, isSet: true}
}

func (v NullableTestRunUpdateMultipleModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestRunUpdateMultipleModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


