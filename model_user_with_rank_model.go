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

// checks if the UserWithRankModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserWithRankModel{}

// UserWithRankModel struct for UserWithRankModel
type UserWithRankModel struct {
	Id *string `json:"id,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	MiddleName *string `json:"middleName,omitempty"`
	UserName *string `json:"userName,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	UserType *string `json:"userType,omitempty"`
	AvatarUrl *string `json:"avatarUrl,omitempty"`
	AvatarMetadata *string `json:"avatarMetadata,omitempty"`
	IsDeleted *bool `json:"isDeleted,omitempty"`
	IsDisabled *bool `json:"isDisabled,omitempty"`
	ProviderId NullableString `json:"providerId,omitempty"`
	IsActiveStatusByEntity *bool `json:"isActiveStatusByEntity,omitempty"`
	UserRank *UserWithRankModelUserRank `json:"userRank,omitempty"`
}

// NewUserWithRankModel instantiates a new UserWithRankModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserWithRankModel() *UserWithRankModel {
	this := UserWithRankModel{}
	return &this
}

// NewUserWithRankModelWithDefaults instantiates a new UserWithRankModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithRankModelWithDefaults() *UserWithRankModel {
	this := UserWithRankModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserWithRankModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithRankModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserWithRankModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserWithRankModel) SetId(v string) {
	o.Id = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UserWithRankModel) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithRankModel) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserWithRankModel) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UserWithRankModel) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UserWithRankModel) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithRankModel) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UserWithRankModel) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UserWithRankModel) SetLastName(v string) {
	o.LastName = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *UserWithRankModel) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithRankModel) GetMiddleNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *UserWithRankModel) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *UserWithRankModel) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *UserWithRankModel) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithRankModel) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *UserWithRankModel) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *UserWithRankModel) SetUserName(v string) {
	o.UserName = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UserWithRankModel) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithRankModel) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UserWithRankModel) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UserWithRankModel) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *UserWithRankModel) GetUserType() string {
	if o == nil || IsNil(o.UserType) {
		var ret string
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithRankModel) GetUserTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UserType) {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *UserWithRankModel) HasUserType() bool {
	if o != nil && !IsNil(o.UserType) {
		return true
	}

	return false
}

// SetUserType gets a reference to the given string and assigns it to the UserType field.
func (o *UserWithRankModel) SetUserType(v string) {
	o.UserType = &v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *UserWithRankModel) GetAvatarUrl() string {
	if o == nil || IsNil(o.AvatarUrl) {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithRankModel) GetAvatarUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarUrl) {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *UserWithRankModel) HasAvatarUrl() bool {
	if o != nil && !IsNil(o.AvatarUrl) {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *UserWithRankModel) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetAvatarMetadata returns the AvatarMetadata field value if set, zero value otherwise.
func (o *UserWithRankModel) GetAvatarMetadata() string {
	if o == nil || IsNil(o.AvatarMetadata) {
		var ret string
		return ret
	}
	return *o.AvatarMetadata
}

// GetAvatarMetadataOk returns a tuple with the AvatarMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithRankModel) GetAvatarMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarMetadata) {
		return nil, false
	}
	return o.AvatarMetadata, true
}

// HasAvatarMetadata returns a boolean if a field has been set.
func (o *UserWithRankModel) HasAvatarMetadata() bool {
	if o != nil && !IsNil(o.AvatarMetadata) {
		return true
	}

	return false
}

// SetAvatarMetadata gets a reference to the given string and assigns it to the AvatarMetadata field.
func (o *UserWithRankModel) SetAvatarMetadata(v string) {
	o.AvatarMetadata = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *UserWithRankModel) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithRankModel) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *UserWithRankModel) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *UserWithRankModel) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetIsDisabled returns the IsDisabled field value if set, zero value otherwise.
func (o *UserWithRankModel) GetIsDisabled() bool {
	if o == nil || IsNil(o.IsDisabled) {
		var ret bool
		return ret
	}
	return *o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithRankModel) GetIsDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDisabled) {
		return nil, false
	}
	return o.IsDisabled, true
}

// HasIsDisabled returns a boolean if a field has been set.
func (o *UserWithRankModel) HasIsDisabled() bool {
	if o != nil && !IsNil(o.IsDisabled) {
		return true
	}

	return false
}

// SetIsDisabled gets a reference to the given bool and assigns it to the IsDisabled field.
func (o *UserWithRankModel) SetIsDisabled(v bool) {
	o.IsDisabled = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserWithRankModel) GetProviderId() string {
	if o == nil || IsNil(o.ProviderId.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderId.Get()
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserWithRankModel) GetProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderId.Get(), o.ProviderId.IsSet()
}

// HasProviderId returns a boolean if a field has been set.
func (o *UserWithRankModel) HasProviderId() bool {
	if o != nil && o.ProviderId.IsSet() {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given NullableString and assigns it to the ProviderId field.
func (o *UserWithRankModel) SetProviderId(v string) {
	o.ProviderId.Set(&v)
}
// SetProviderIdNil sets the value for ProviderId to be an explicit nil
func (o *UserWithRankModel) SetProviderIdNil() {
	o.ProviderId.Set(nil)
}

// UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
func (o *UserWithRankModel) UnsetProviderId() {
	o.ProviderId.Unset()
}

// GetIsActiveStatusByEntity returns the IsActiveStatusByEntity field value if set, zero value otherwise.
func (o *UserWithRankModel) GetIsActiveStatusByEntity() bool {
	if o == nil || IsNil(o.IsActiveStatusByEntity) {
		var ret bool
		return ret
	}
	return *o.IsActiveStatusByEntity
}

// GetIsActiveStatusByEntityOk returns a tuple with the IsActiveStatusByEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithRankModel) GetIsActiveStatusByEntityOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActiveStatusByEntity) {
		return nil, false
	}
	return o.IsActiveStatusByEntity, true
}

// HasIsActiveStatusByEntity returns a boolean if a field has been set.
func (o *UserWithRankModel) HasIsActiveStatusByEntity() bool {
	if o != nil && !IsNil(o.IsActiveStatusByEntity) {
		return true
	}

	return false
}

// SetIsActiveStatusByEntity gets a reference to the given bool and assigns it to the IsActiveStatusByEntity field.
func (o *UserWithRankModel) SetIsActiveStatusByEntity(v bool) {
	o.IsActiveStatusByEntity = &v
}

// GetUserRank returns the UserRank field value if set, zero value otherwise.
func (o *UserWithRankModel) GetUserRank() UserWithRankModelUserRank {
	if o == nil || IsNil(o.UserRank) {
		var ret UserWithRankModelUserRank
		return ret
	}
	return *o.UserRank
}

// GetUserRankOk returns a tuple with the UserRank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWithRankModel) GetUserRankOk() (*UserWithRankModelUserRank, bool) {
	if o == nil || IsNil(o.UserRank) {
		return nil, false
	}
	return o.UserRank, true
}

// HasUserRank returns a boolean if a field has been set.
func (o *UserWithRankModel) HasUserRank() bool {
	if o != nil && !IsNil(o.UserRank) {
		return true
	}

	return false
}

// SetUserRank gets a reference to the given UserWithRankModelUserRank and assigns it to the UserRank field.
func (o *UserWithRankModel) SetUserRank(v UserWithRankModelUserRank) {
	o.UserRank = &v
}

func (o UserWithRankModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserWithRankModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.MiddleName) {
		toSerialize["middleName"] = o.MiddleName
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.UserType) {
		toSerialize["userType"] = o.UserType
	}
	if !IsNil(o.AvatarUrl) {
		toSerialize["avatarUrl"] = o.AvatarUrl
	}
	if !IsNil(o.AvatarMetadata) {
		toSerialize["avatarMetadata"] = o.AvatarMetadata
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	if !IsNil(o.IsDisabled) {
		toSerialize["isDisabled"] = o.IsDisabled
	}
	if o.ProviderId.IsSet() {
		toSerialize["providerId"] = o.ProviderId.Get()
	}
	if !IsNil(o.IsActiveStatusByEntity) {
		toSerialize["isActiveStatusByEntity"] = o.IsActiveStatusByEntity
	}
	if !IsNil(o.UserRank) {
		toSerialize["userRank"] = o.UserRank
	}
	return toSerialize, nil
}

type NullableUserWithRankModel struct {
	value *UserWithRankModel
	isSet bool
}

func (v NullableUserWithRankModel) Get() *UserWithRankModel {
	return v.value
}

func (v *NullableUserWithRankModel) Set(val *UserWithRankModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUserWithRankModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUserWithRankModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserWithRankModel(val *UserWithRankModel) *NullableUserWithRankModel {
	return &NullableUserWithRankModel{value: val, isSet: true}
}

func (v NullableUserWithRankModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserWithRankModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


