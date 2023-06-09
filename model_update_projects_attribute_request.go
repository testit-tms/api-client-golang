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

// checks if the UpdateProjectsAttributeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProjectsAttributeRequest{}

// UpdateProjectsAttributeRequest struct for UpdateProjectsAttributeRequest
type UpdateProjectsAttributeRequest struct {
	// Unique ID of the attribute
	Id *string `json:"id,omitempty"`
	// Collection of the attribute options  <br />  Available for attributes of type `options` and `multiple options` only
	Options []CustomAttributeOptionModel `json:"options,omitempty"`
	Type CustomAttributeTypesEnum `json:"type"`
	// Indicates if the entity is deleted
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// Name of the attribute
	Name string `json:"name"`
	// Indicates if the attribute is enabled
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Indicates if the attribute value is mandatory to specify
	IsRequired *bool `json:"isRequired,omitempty"`
	// Indicates if the attribute is available across all projects
	IsGlobal *bool `json:"isGlobal,omitempty"`
}

// NewUpdateProjectsAttributeRequest instantiates a new UpdateProjectsAttributeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProjectsAttributeRequest(type_ CustomAttributeTypesEnum, name string) *UpdateProjectsAttributeRequest {
	this := UpdateProjectsAttributeRequest{}
	this.Type = type_
	this.Name = name
	return &this
}

// NewUpdateProjectsAttributeRequestWithDefaults instantiates a new UpdateProjectsAttributeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProjectsAttributeRequestWithDefaults() *UpdateProjectsAttributeRequest {
	this := UpdateProjectsAttributeRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateProjectsAttributeRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProjectsAttributeRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateProjectsAttributeRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateProjectsAttributeRequest) SetId(v string) {
	o.Id = &v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProjectsAttributeRequest) GetOptions() []CustomAttributeOptionModel {
	if o == nil {
		var ret []CustomAttributeOptionModel
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProjectsAttributeRequest) GetOptionsOk() ([]CustomAttributeOptionModel, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *UpdateProjectsAttributeRequest) HasOptions() bool {
	if o != nil && IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []CustomAttributeOptionModel and assigns it to the Options field.
func (o *UpdateProjectsAttributeRequest) SetOptions(v []CustomAttributeOptionModel) {
	o.Options = v
}

// GetType returns the Type field value
func (o *UpdateProjectsAttributeRequest) GetType() CustomAttributeTypesEnum {
	if o == nil {
		var ret CustomAttributeTypesEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateProjectsAttributeRequest) GetTypeOk() (*CustomAttributeTypesEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateProjectsAttributeRequest) SetType(v CustomAttributeTypesEnum) {
	o.Type = v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *UpdateProjectsAttributeRequest) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProjectsAttributeRequest) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *UpdateProjectsAttributeRequest) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *UpdateProjectsAttributeRequest) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetName returns the Name field value
func (o *UpdateProjectsAttributeRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateProjectsAttributeRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateProjectsAttributeRequest) SetName(v string) {
	o.Name = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *UpdateProjectsAttributeRequest) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProjectsAttributeRequest) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *UpdateProjectsAttributeRequest) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *UpdateProjectsAttributeRequest) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *UpdateProjectsAttributeRequest) GetIsRequired() bool {
	if o == nil || IsNil(o.IsRequired) {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProjectsAttributeRequest) GetIsRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRequired) {
		return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *UpdateProjectsAttributeRequest) HasIsRequired() bool {
	if o != nil && !IsNil(o.IsRequired) {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *UpdateProjectsAttributeRequest) SetIsRequired(v bool) {
	o.IsRequired = &v
}

// GetIsGlobal returns the IsGlobal field value if set, zero value otherwise.
func (o *UpdateProjectsAttributeRequest) GetIsGlobal() bool {
	if o == nil || IsNil(o.IsGlobal) {
		var ret bool
		return ret
	}
	return *o.IsGlobal
}

// GetIsGlobalOk returns a tuple with the IsGlobal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProjectsAttributeRequest) GetIsGlobalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGlobal) {
		return nil, false
	}
	return o.IsGlobal, true
}

// HasIsGlobal returns a boolean if a field has been set.
func (o *UpdateProjectsAttributeRequest) HasIsGlobal() bool {
	if o != nil && !IsNil(o.IsGlobal) {
		return true
	}

	return false
}

// SetIsGlobal gets a reference to the given bool and assigns it to the IsGlobal field.
func (o *UpdateProjectsAttributeRequest) SetIsGlobal(v bool) {
	o.IsGlobal = &v
}

func (o UpdateProjectsAttributeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProjectsAttributeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !IsNil(o.IsRequired) {
		toSerialize["isRequired"] = o.IsRequired
	}
	if !IsNil(o.IsGlobal) {
		toSerialize["isGlobal"] = o.IsGlobal
	}
	return toSerialize, nil
}

type NullableUpdateProjectsAttributeRequest struct {
	value *UpdateProjectsAttributeRequest
	isSet bool
}

func (v NullableUpdateProjectsAttributeRequest) Get() *UpdateProjectsAttributeRequest {
	return v.value
}

func (v *NullableUpdateProjectsAttributeRequest) Set(val *UpdateProjectsAttributeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProjectsAttributeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProjectsAttributeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProjectsAttributeRequest(val *UpdateProjectsAttributeRequest) *NullableUpdateProjectsAttributeRequest {
	return &NullableUpdateProjectsAttributeRequest{value: val, isSet: true}
}

func (v NullableUpdateProjectsAttributeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProjectsAttributeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


