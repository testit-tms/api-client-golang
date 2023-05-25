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

// checks if the CustomAttributeModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomAttributeModel{}

// CustomAttributeModel struct for CustomAttributeModel
type CustomAttributeModel struct {
	// Unique ID of the attribute
	Id *string `json:"id,omitempty"`
	// Collection of the attribute options  <br />  Available for attributes of type `options` and `multiple options` only
	Options []CustomAttributeOptionModel `json:"options,omitempty"`
	Type CustomAttributeTypesEnum `json:"type"`
	// Indicates if the attribute is deleted
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

// NewCustomAttributeModel instantiates a new CustomAttributeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomAttributeModel(type_ CustomAttributeTypesEnum, name string) *CustomAttributeModel {
	this := CustomAttributeModel{}
	this.Type = type_
	this.Name = name
	return &this
}

// NewCustomAttributeModelWithDefaults instantiates a new CustomAttributeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomAttributeModelWithDefaults() *CustomAttributeModel {
	this := CustomAttributeModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomAttributeModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAttributeModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomAttributeModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomAttributeModel) SetId(v string) {
	o.Id = &v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomAttributeModel) GetOptions() []CustomAttributeOptionModel {
	if o == nil {
		var ret []CustomAttributeOptionModel
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomAttributeModel) GetOptionsOk() ([]CustomAttributeOptionModel, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *CustomAttributeModel) HasOptions() bool {
	if o != nil && IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []CustomAttributeOptionModel and assigns it to the Options field.
func (o *CustomAttributeModel) SetOptions(v []CustomAttributeOptionModel) {
	o.Options = v
}

// GetType returns the Type field value
func (o *CustomAttributeModel) GetType() CustomAttributeTypesEnum {
	if o == nil {
		var ret CustomAttributeTypesEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomAttributeModel) GetTypeOk() (*CustomAttributeTypesEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomAttributeModel) SetType(v CustomAttributeTypesEnum) {
	o.Type = v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *CustomAttributeModel) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAttributeModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *CustomAttributeModel) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *CustomAttributeModel) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetName returns the Name field value
func (o *CustomAttributeModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomAttributeModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomAttributeModel) SetName(v string) {
	o.Name = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *CustomAttributeModel) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAttributeModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *CustomAttributeModel) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *CustomAttributeModel) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *CustomAttributeModel) GetIsRequired() bool {
	if o == nil || IsNil(o.IsRequired) {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAttributeModel) GetIsRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRequired) {
		return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *CustomAttributeModel) HasIsRequired() bool {
	if o != nil && !IsNil(o.IsRequired) {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *CustomAttributeModel) SetIsRequired(v bool) {
	o.IsRequired = &v
}

// GetIsGlobal returns the IsGlobal field value if set, zero value otherwise.
func (o *CustomAttributeModel) GetIsGlobal() bool {
	if o == nil || IsNil(o.IsGlobal) {
		var ret bool
		return ret
	}
	return *o.IsGlobal
}

// GetIsGlobalOk returns a tuple with the IsGlobal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAttributeModel) GetIsGlobalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGlobal) {
		return nil, false
	}
	return o.IsGlobal, true
}

// HasIsGlobal returns a boolean if a field has been set.
func (o *CustomAttributeModel) HasIsGlobal() bool {
	if o != nil && !IsNil(o.IsGlobal) {
		return true
	}

	return false
}

// SetIsGlobal gets a reference to the given bool and assigns it to the IsGlobal field.
func (o *CustomAttributeModel) SetIsGlobal(v bool) {
	o.IsGlobal = &v
}

func (o CustomAttributeModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomAttributeModel) ToMap() (map[string]interface{}, error) {
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

type NullableCustomAttributeModel struct {
	value *CustomAttributeModel
	isSet bool
}

func (v NullableCustomAttributeModel) Get() *CustomAttributeModel {
	return v.value
}

func (v *NullableCustomAttributeModel) Set(val *CustomAttributeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomAttributeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomAttributeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomAttributeModel(val *CustomAttributeModel) *NullableCustomAttributeModel {
	return &NullableCustomAttributeModel{value: val, isSet: true}
}

func (v NullableCustomAttributeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomAttributeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


