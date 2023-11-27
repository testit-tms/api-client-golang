/*
API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmsclient

import (
	"encoding/json"
	"time"
)

// checks if the AttachmentModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachmentModel{}

// AttachmentModel struct for AttachmentModel
type AttachmentModel struct {
	// Unique ID of the attachment file
	FileId string `json:"fileId"`
	// MIME type of the attachment
	Type string `json:"type"`
	// Size in bytes of the attachment file
	Size float32 `json:"size"`
	// Creation date of the attachment
	CreatedDate time.Time `json:"createdDate"`
	// Last modification date of the attachment
	ModifiedDate NullableTime `json:"modifiedDate,omitempty"`
	// Unique ID of the attachment creator
	CreatedById string `json:"createdById"`
	// Unique ID of the attachment last editor
	ModifiedById NullableString `json:"modifiedById,omitempty"`
	// Name of the attachment file
	Name string `json:"name"`
	// Indicates whether the attachment is temporary (may be automatically deleted)
	IsTemp bool `json:"isTemp"`
	// Unique ID of the attachment
	Id string `json:"id"`
}

// NewAttachmentModel instantiates a new AttachmentModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentModel(fileId string, type_ string, size float32, createdDate time.Time, createdById string, name string, isTemp bool, id string) *AttachmentModel {
	this := AttachmentModel{}
	this.FileId = fileId
	this.Type = type_
	this.Size = size
	this.CreatedDate = createdDate
	this.CreatedById = createdById
	this.Name = name
	this.IsTemp = isTemp
	this.Id = id
	return &this
}

// NewAttachmentModelWithDefaults instantiates a new AttachmentModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentModelWithDefaults() *AttachmentModel {
	this := AttachmentModel{}
	return &this
}

// GetFileId returns the FileId field value
func (o *AttachmentModel) GetFileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value
// and a boolean to check if the value has been set.
func (o *AttachmentModel) GetFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileId, true
}

// SetFileId sets field value
func (o *AttachmentModel) SetFileId(v string) {
	o.FileId = v
}

// GetType returns the Type field value
func (o *AttachmentModel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AttachmentModel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AttachmentModel) SetType(v string) {
	o.Type = v
}

// GetSize returns the Size field value
func (o *AttachmentModel) GetSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *AttachmentModel) GetSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *AttachmentModel) SetSize(v float32) {
	o.Size = v
}

// GetCreatedDate returns the CreatedDate field value
func (o *AttachmentModel) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *AttachmentModel) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *AttachmentModel) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentModel) GetModifiedDate() time.Time {
	if o == nil || IsNil(o.ModifiedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate.Get()
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentModel) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedDate.Get(), o.ModifiedDate.IsSet()
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *AttachmentModel) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate.IsSet() {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given NullableTime and assigns it to the ModifiedDate field.
func (o *AttachmentModel) SetModifiedDate(v time.Time) {
	o.ModifiedDate.Set(&v)
}
// SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil
func (o *AttachmentModel) SetModifiedDateNil() {
	o.ModifiedDate.Set(nil)
}

// UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
func (o *AttachmentModel) UnsetModifiedDate() {
	o.ModifiedDate.Unset()
}

// GetCreatedById returns the CreatedById field value
func (o *AttachmentModel) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *AttachmentModel) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *AttachmentModel) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetModifiedById returns the ModifiedById field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentModel) GetModifiedById() string {
	if o == nil || IsNil(o.ModifiedById.Get()) {
		var ret string
		return ret
	}
	return *o.ModifiedById.Get()
}

// GetModifiedByIdOk returns a tuple with the ModifiedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentModel) GetModifiedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedById.Get(), o.ModifiedById.IsSet()
}

// HasModifiedById returns a boolean if a field has been set.
func (o *AttachmentModel) HasModifiedById() bool {
	if o != nil && o.ModifiedById.IsSet() {
		return true
	}

	return false
}

// SetModifiedById gets a reference to the given NullableString and assigns it to the ModifiedById field.
func (o *AttachmentModel) SetModifiedById(v string) {
	o.ModifiedById.Set(&v)
}
// SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil
func (o *AttachmentModel) SetModifiedByIdNil() {
	o.ModifiedById.Set(nil)
}

// UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
func (o *AttachmentModel) UnsetModifiedById() {
	o.ModifiedById.Unset()
}

// GetName returns the Name field value
func (o *AttachmentModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AttachmentModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AttachmentModel) SetName(v string) {
	o.Name = v
}

// GetIsTemp returns the IsTemp field value
func (o *AttachmentModel) GetIsTemp() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsTemp
}

// GetIsTempOk returns a tuple with the IsTemp field value
// and a boolean to check if the value has been set.
func (o *AttachmentModel) GetIsTempOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTemp, true
}

// SetIsTemp sets field value
func (o *AttachmentModel) SetIsTemp(v bool) {
	o.IsTemp = v
}

// GetId returns the Id field value
func (o *AttachmentModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AttachmentModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AttachmentModel) SetId(v string) {
	o.Id = v
}

func (o AttachmentModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachmentModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fileId"] = o.FileId
	toSerialize["type"] = o.Type
	toSerialize["size"] = o.Size
	toSerialize["createdDate"] = o.CreatedDate
	if o.ModifiedDate.IsSet() {
		toSerialize["modifiedDate"] = o.ModifiedDate.Get()
	}
	toSerialize["createdById"] = o.CreatedById
	if o.ModifiedById.IsSet() {
		toSerialize["modifiedById"] = o.ModifiedById.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["isTemp"] = o.IsTemp
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAttachmentModel struct {
	value *AttachmentModel
	isSet bool
}

func (v NullableAttachmentModel) Get() *AttachmentModel {
	return v.value
}

func (v *NullableAttachmentModel) Set(val *AttachmentModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentModel(val *AttachmentModel) *NullableAttachmentModel {
	return &NullableAttachmentModel{value: val, isSet: true}
}

func (v NullableAttachmentModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


