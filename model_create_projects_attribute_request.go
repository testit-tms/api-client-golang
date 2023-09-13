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

// checks if the CreateProjectsAttributeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProjectsAttributeRequest{}

// CreateProjectsAttributeRequest struct for CreateProjectsAttributeRequest
type CreateProjectsAttributeRequest struct {
	// Collection of attribute options  <br />  Available for attributes of type `options` and `multiple options` only
	Options []CustomAttributeOptionPostModel `json:"options,omitempty"`
	Type CustomAttributeTypesEnum `json:"type"`
	// Name of the attribute
	Name string `json:"name"`
	// Indicates if the attribute is enabled
	IsEnabled bool `json:"isEnabled"`
	// Indicates if the attribute value is mandatory to specify
	IsRequired bool `json:"isRequired"`
	// Indicates if the attribute is available across all projects
	IsGlobal bool `json:"isGlobal"`
}

// NewCreateProjectsAttributeRequest instantiates a new CreateProjectsAttributeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectsAttributeRequest(type_ CustomAttributeTypesEnum, name string, isEnabled bool, isRequired bool, isGlobal bool) *CreateProjectsAttributeRequest {
	this := CreateProjectsAttributeRequest{}
	this.Type = type_
	this.Name = name
	this.IsEnabled = isEnabled
	this.IsRequired = isRequired
	this.IsGlobal = isGlobal
	return &this
}

// NewCreateProjectsAttributeRequestWithDefaults instantiates a new CreateProjectsAttributeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectsAttributeRequestWithDefaults() *CreateProjectsAttributeRequest {
	this := CreateProjectsAttributeRequest{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProjectsAttributeRequest) GetOptions() []CustomAttributeOptionPostModel {
	if o == nil {
		var ret []CustomAttributeOptionPostModel
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProjectsAttributeRequest) GetOptionsOk() ([]CustomAttributeOptionPostModel, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *CreateProjectsAttributeRequest) HasOptions() bool {
	if o != nil && IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []CustomAttributeOptionPostModel and assigns it to the Options field.
func (o *CreateProjectsAttributeRequest) SetOptions(v []CustomAttributeOptionPostModel) {
	o.Options = v
}

// GetType returns the Type field value
func (o *CreateProjectsAttributeRequest) GetType() CustomAttributeTypesEnum {
	if o == nil {
		var ret CustomAttributeTypesEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateProjectsAttributeRequest) GetTypeOk() (*CustomAttributeTypesEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateProjectsAttributeRequest) SetType(v CustomAttributeTypesEnum) {
	o.Type = v
}

// GetName returns the Name field value
func (o *CreateProjectsAttributeRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateProjectsAttributeRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateProjectsAttributeRequest) SetName(v string) {
	o.Name = v
}

// GetIsEnabled returns the IsEnabled field value
func (o *CreateProjectsAttributeRequest) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *CreateProjectsAttributeRequest) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *CreateProjectsAttributeRequest) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetIsRequired returns the IsRequired field value
func (o *CreateProjectsAttributeRequest) GetIsRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value
// and a boolean to check if the value has been set.
func (o *CreateProjectsAttributeRequest) GetIsRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRequired, true
}

// SetIsRequired sets field value
func (o *CreateProjectsAttributeRequest) SetIsRequired(v bool) {
	o.IsRequired = v
}

// GetIsGlobal returns the IsGlobal field value
func (o *CreateProjectsAttributeRequest) GetIsGlobal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsGlobal
}

// GetIsGlobalOk returns a tuple with the IsGlobal field value
// and a boolean to check if the value has been set.
func (o *CreateProjectsAttributeRequest) GetIsGlobalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsGlobal, true
}

// SetIsGlobal sets field value
func (o *CreateProjectsAttributeRequest) SetIsGlobal(v bool) {
	o.IsGlobal = v
}

func (o CreateProjectsAttributeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProjectsAttributeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	toSerialize["isEnabled"] = o.IsEnabled
	toSerialize["isRequired"] = o.IsRequired
	toSerialize["isGlobal"] = o.IsGlobal
	return toSerialize, nil
}

type NullableCreateProjectsAttributeRequest struct {
	value *CreateProjectsAttributeRequest
	isSet bool
}

func (v NullableCreateProjectsAttributeRequest) Get() *CreateProjectsAttributeRequest {
	return v.value
}

func (v *NullableCreateProjectsAttributeRequest) Set(val *CreateProjectsAttributeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectsAttributeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectsAttributeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectsAttributeRequest(val *CreateProjectsAttributeRequest) *NullableCreateProjectsAttributeRequest {
	return &NullableCreateProjectsAttributeRequest{value: val, isSet: true}
}

func (v NullableCreateProjectsAttributeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectsAttributeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

