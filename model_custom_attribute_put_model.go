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

// checks if the CustomAttributePutModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomAttributePutModel{}

// CustomAttributePutModel struct for CustomAttributePutModel
type CustomAttributePutModel struct {
	// Unique ID of the attribute
	Id string `json:"id"`
	// Collection of the attribute options   Available for attributes of type `options` and `multiple options` only
	Options []CustomAttributeOptionModel `json:"options,omitempty"`
	// Type of the attribute
	Type CustomAttributeTypesEnum `json:"type"`
	// Indicates if the entity is deleted
	IsDeleted bool `json:"isDeleted"`
	// Name of the attribute
	Name string `json:"name"`
	// Indicates if the attribute is enabled
	IsEnabled bool `json:"isEnabled"`
	// Indicates if the attribute value is mandatory to specify
	IsRequired bool `json:"isRequired"`
	// Indicates if the attribute is available across all projects
	IsGlobal bool `json:"isGlobal"`
}

type _CustomAttributePutModel CustomAttributePutModel

// NewCustomAttributePutModel instantiates a new CustomAttributePutModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomAttributePutModel(id string, type_ CustomAttributeTypesEnum, isDeleted bool, name string, isEnabled bool, isRequired bool, isGlobal bool) *CustomAttributePutModel {
	this := CustomAttributePutModel{}
	this.Id = id
	this.Type = type_
	this.IsDeleted = isDeleted
	this.Name = name
	this.IsEnabled = isEnabled
	this.IsRequired = isRequired
	this.IsGlobal = isGlobal
	return &this
}

// NewCustomAttributePutModelWithDefaults instantiates a new CustomAttributePutModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomAttributePutModelWithDefaults() *CustomAttributePutModel {
	this := CustomAttributePutModel{}
	return &this
}

// GetId returns the Id field value
func (o *CustomAttributePutModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomAttributePutModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomAttributePutModel) SetId(v string) {
	o.Id = v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomAttributePutModel) GetOptions() []CustomAttributeOptionModel {
	if o == nil {
		var ret []CustomAttributeOptionModel
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomAttributePutModel) GetOptionsOk() ([]CustomAttributeOptionModel, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *CustomAttributePutModel) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []CustomAttributeOptionModel and assigns it to the Options field.
func (o *CustomAttributePutModel) SetOptions(v []CustomAttributeOptionModel) {
	o.Options = v
}

// GetType returns the Type field value
func (o *CustomAttributePutModel) GetType() CustomAttributeTypesEnum {
	if o == nil {
		var ret CustomAttributeTypesEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomAttributePutModel) GetTypeOk() (*CustomAttributeTypesEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomAttributePutModel) SetType(v CustomAttributeTypesEnum) {
	o.Type = v
}

// GetIsDeleted returns the IsDeleted field value
func (o *CustomAttributePutModel) GetIsDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value
// and a boolean to check if the value has been set.
func (o *CustomAttributePutModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeleted, true
}

// SetIsDeleted sets field value
func (o *CustomAttributePutModel) SetIsDeleted(v bool) {
	o.IsDeleted = v
}

// GetName returns the Name field value
func (o *CustomAttributePutModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomAttributePutModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomAttributePutModel) SetName(v string) {
	o.Name = v
}

// GetIsEnabled returns the IsEnabled field value
func (o *CustomAttributePutModel) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *CustomAttributePutModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *CustomAttributePutModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetIsRequired returns the IsRequired field value
func (o *CustomAttributePutModel) GetIsRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value
// and a boolean to check if the value has been set.
func (o *CustomAttributePutModel) GetIsRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRequired, true
}

// SetIsRequired sets field value
func (o *CustomAttributePutModel) SetIsRequired(v bool) {
	o.IsRequired = v
}

// GetIsGlobal returns the IsGlobal field value
func (o *CustomAttributePutModel) GetIsGlobal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsGlobal
}

// GetIsGlobalOk returns a tuple with the IsGlobal field value
// and a boolean to check if the value has been set.
func (o *CustomAttributePutModel) GetIsGlobalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsGlobal, true
}

// SetIsGlobal sets field value
func (o *CustomAttributePutModel) SetIsGlobal(v bool) {
	o.IsGlobal = v
}

func (o CustomAttributePutModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomAttributePutModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	toSerialize["type"] = o.Type
	toSerialize["isDeleted"] = o.IsDeleted
	toSerialize["name"] = o.Name
	toSerialize["isEnabled"] = o.IsEnabled
	toSerialize["isRequired"] = o.IsRequired
	toSerialize["isGlobal"] = o.IsGlobal
	return toSerialize, nil
}

func (o *CustomAttributePutModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"isDeleted",
		"name",
		"isEnabled",
		"isRequired",
		"isGlobal",
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

	varCustomAttributePutModel := _CustomAttributePutModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomAttributePutModel)

	if err != nil {
		return err
	}

	*o = CustomAttributePutModel(varCustomAttributePutModel)

	return err
}

type NullableCustomAttributePutModel struct {
	value *CustomAttributePutModel
	isSet bool
}

func (v NullableCustomAttributePutModel) Get() *CustomAttributePutModel {
	return v.value
}

func (v *NullableCustomAttributePutModel) Set(val *CustomAttributePutModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomAttributePutModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomAttributePutModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomAttributePutModel(val *CustomAttributePutModel) *NullableCustomAttributePutModel {
	return &NullableCustomAttributePutModel{value: val, isSet: true}
}

func (v NullableCustomAttributePutModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomAttributePutModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


