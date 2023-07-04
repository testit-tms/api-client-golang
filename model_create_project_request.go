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

// checks if the CreateProjectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProjectRequest{}

// CreateProjectRequest struct for CreateProjectRequest
type CreateProjectRequest struct {
	// Description of the project
	Description NullableString `json:"description,omitempty"`
	// Name of the project
	Name string `json:"name"`
	// Indicates if the project is marked as favorite
	IsFavorite NullableBool `json:"isFavorite,omitempty"`
}

// NewCreateProjectRequest instantiates a new CreateProjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectRequest(name string) *CreateProjectRequest {
	this := CreateProjectRequest{}
	this.Name = name
	return &this
}

// NewCreateProjectRequestWithDefaults instantiates a new CreateProjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectRequestWithDefaults() *CreateProjectRequest {
	this := CreateProjectRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProjectRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProjectRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateProjectRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateProjectRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateProjectRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateProjectRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value
func (o *CreateProjectRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateProjectRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateProjectRequest) SetName(v string) {
	o.Name = v
}

// GetIsFavorite returns the IsFavorite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProjectRequest) GetIsFavorite() bool {
	if o == nil || IsNil(o.IsFavorite.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFavorite.Get()
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProjectRequest) GetIsFavoriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFavorite.Get(), o.IsFavorite.IsSet()
}

// HasIsFavorite returns a boolean if a field has been set.
func (o *CreateProjectRequest) HasIsFavorite() bool {
	if o != nil && o.IsFavorite.IsSet() {
		return true
	}

	return false
}

// SetIsFavorite gets a reference to the given NullableBool and assigns it to the IsFavorite field.
func (o *CreateProjectRequest) SetIsFavorite(v bool) {
	o.IsFavorite.Set(&v)
}
// SetIsFavoriteNil sets the value for IsFavorite to be an explicit nil
func (o *CreateProjectRequest) SetIsFavoriteNil() {
	o.IsFavorite.Set(nil)
}

// UnsetIsFavorite ensures that no value is present for IsFavorite, not even an explicit nil
func (o *CreateProjectRequest) UnsetIsFavorite() {
	o.IsFavorite.Unset()
}

func (o CreateProjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProjectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["name"] = o.Name
	if o.IsFavorite.IsSet() {
		toSerialize["isFavorite"] = o.IsFavorite.Get()
	}
	return toSerialize, nil
}

type NullableCreateProjectRequest struct {
	value *CreateProjectRequest
	isSet bool
}

func (v NullableCreateProjectRequest) Get() *CreateProjectRequest {
	return v.value
}

func (v *NullableCreateProjectRequest) Set(val *CreateProjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectRequest(val *CreateProjectRequest) *NullableCreateProjectRequest {
	return &NullableCreateProjectRequest{value: val, isSet: true}
}

func (v NullableCreateProjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


