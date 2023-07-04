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

// checks if the WorkItemLinkChangeViewModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkItemLinkChangeViewModel{}

// WorkItemLinkChangeViewModel struct for WorkItemLinkChangeViewModel
type WorkItemLinkChangeViewModel struct {
	Description *string `json:"description,omitempty"`
	Url *string `json:"url,omitempty"`
	Title *string `json:"title,omitempty"`
	HasInfo *bool `json:"hasInfo,omitempty"`
	Id *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewWorkItemLinkChangeViewModel instantiates a new WorkItemLinkChangeViewModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItemLinkChangeViewModel() *WorkItemLinkChangeViewModel {
	this := WorkItemLinkChangeViewModel{}
	return &this
}

// NewWorkItemLinkChangeViewModelWithDefaults instantiates a new WorkItemLinkChangeViewModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemLinkChangeViewModelWithDefaults() *WorkItemLinkChangeViewModel {
	this := WorkItemLinkChangeViewModel{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkItemLinkChangeViewModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemLinkChangeViewModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkItemLinkChangeViewModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkItemLinkChangeViewModel) SetDescription(v string) {
	o.Description = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WorkItemLinkChangeViewModel) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemLinkChangeViewModel) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WorkItemLinkChangeViewModel) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WorkItemLinkChangeViewModel) SetUrl(v string) {
	o.Url = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *WorkItemLinkChangeViewModel) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemLinkChangeViewModel) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *WorkItemLinkChangeViewModel) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *WorkItemLinkChangeViewModel) SetTitle(v string) {
	o.Title = &v
}

// GetHasInfo returns the HasInfo field value if set, zero value otherwise.
func (o *WorkItemLinkChangeViewModel) GetHasInfo() bool {
	if o == nil || IsNil(o.HasInfo) {
		var ret bool
		return ret
	}
	return *o.HasInfo
}

// GetHasInfoOk returns a tuple with the HasInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemLinkChangeViewModel) GetHasInfoOk() (*bool, bool) {
	if o == nil || IsNil(o.HasInfo) {
		return nil, false
	}
	return o.HasInfo, true
}

// HasHasInfo returns a boolean if a field has been set.
func (o *WorkItemLinkChangeViewModel) HasHasInfo() bool {
	if o != nil && !IsNil(o.HasInfo) {
		return true
	}

	return false
}

// SetHasInfo gets a reference to the given bool and assigns it to the HasInfo field.
func (o *WorkItemLinkChangeViewModel) SetHasInfo(v bool) {
	o.HasInfo = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkItemLinkChangeViewModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemLinkChangeViewModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkItemLinkChangeViewModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkItemLinkChangeViewModel) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkItemLinkChangeViewModel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItemLinkChangeViewModel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkItemLinkChangeViewModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkItemLinkChangeViewModel) SetType(v string) {
	o.Type = &v
}

func (o WorkItemLinkChangeViewModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkItemLinkChangeViewModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.HasInfo) {
		toSerialize["hasInfo"] = o.HasInfo
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableWorkItemLinkChangeViewModel struct {
	value *WorkItemLinkChangeViewModel
	isSet bool
}

func (v NullableWorkItemLinkChangeViewModel) Get() *WorkItemLinkChangeViewModel {
	return v.value
}

func (v *NullableWorkItemLinkChangeViewModel) Set(val *WorkItemLinkChangeViewModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItemLinkChangeViewModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItemLinkChangeViewModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItemLinkChangeViewModel(val *WorkItemLinkChangeViewModel) *NullableWorkItemLinkChangeViewModel {
	return &NullableWorkItemLinkChangeViewModel{value: val, isSet: true}
}

func (v NullableWorkItemLinkChangeViewModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItemLinkChangeViewModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


