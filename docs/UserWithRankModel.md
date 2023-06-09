# UserWithRankModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**MiddleName** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**UserType** | Pointer to **string** |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**AvatarMetadata** | Pointer to **string** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**IsDisabled** | Pointer to **bool** |  | [optional] 
**ProviderId** | Pointer to **NullableString** |  | [optional] 
**IsActiveStatusByEntity** | Pointer to **bool** |  | [optional] 
**UserRank** | Pointer to [**UserWithRankModelUserRank**](UserWithRankModelUserRank.md) |  | [optional] 

## Methods

### NewUserWithRankModel

`func NewUserWithRankModel() *UserWithRankModel`

NewUserWithRankModel instantiates a new UserWithRankModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithRankModelWithDefaults

`func NewUserWithRankModelWithDefaults() *UserWithRankModel`

NewUserWithRankModelWithDefaults instantiates a new UserWithRankModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserWithRankModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserWithRankModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserWithRankModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserWithRankModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *UserWithRankModel) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserWithRankModel) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserWithRankModel) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserWithRankModel) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserWithRankModel) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserWithRankModel) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserWithRankModel) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserWithRankModel) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMiddleName

`func (o *UserWithRankModel) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *UserWithRankModel) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *UserWithRankModel) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *UserWithRankModel) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetUserName

`func (o *UserWithRankModel) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UserWithRankModel) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UserWithRankModel) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UserWithRankModel) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserWithRankModel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserWithRankModel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserWithRankModel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserWithRankModel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetUserType

`func (o *UserWithRankModel) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UserWithRankModel) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UserWithRankModel) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UserWithRankModel) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *UserWithRankModel) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UserWithRankModel) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UserWithRankModel) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UserWithRankModel) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetAvatarMetadata

`func (o *UserWithRankModel) GetAvatarMetadata() string`

GetAvatarMetadata returns the AvatarMetadata field if non-nil, zero value otherwise.

### GetAvatarMetadataOk

`func (o *UserWithRankModel) GetAvatarMetadataOk() (*string, bool)`

GetAvatarMetadataOk returns a tuple with the AvatarMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarMetadata

`func (o *UserWithRankModel) SetAvatarMetadata(v string)`

SetAvatarMetadata sets AvatarMetadata field to given value.

### HasAvatarMetadata

`func (o *UserWithRankModel) HasAvatarMetadata() bool`

HasAvatarMetadata returns a boolean if a field has been set.

### GetIsDeleted

`func (o *UserWithRankModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UserWithRankModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UserWithRankModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *UserWithRankModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsDisabled

`func (o *UserWithRankModel) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *UserWithRankModel) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *UserWithRankModel) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *UserWithRankModel) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetProviderId

`func (o *UserWithRankModel) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *UserWithRankModel) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *UserWithRankModel) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *UserWithRankModel) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *UserWithRankModel) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *UserWithRankModel) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetIsActiveStatusByEntity

`func (o *UserWithRankModel) GetIsActiveStatusByEntity() bool`

GetIsActiveStatusByEntity returns the IsActiveStatusByEntity field if non-nil, zero value otherwise.

### GetIsActiveStatusByEntityOk

`func (o *UserWithRankModel) GetIsActiveStatusByEntityOk() (*bool, bool)`

GetIsActiveStatusByEntityOk returns a tuple with the IsActiveStatusByEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActiveStatusByEntity

`func (o *UserWithRankModel) SetIsActiveStatusByEntity(v bool)`

SetIsActiveStatusByEntity sets IsActiveStatusByEntity field to given value.

### HasIsActiveStatusByEntity

`func (o *UserWithRankModel) HasIsActiveStatusByEntity() bool`

HasIsActiveStatusByEntity returns a boolean if a field has been set.

### GetUserRank

`func (o *UserWithRankModel) GetUserRank() UserWithRankModelUserRank`

GetUserRank returns the UserRank field if non-nil, zero value otherwise.

### GetUserRankOk

`func (o *UserWithRankModel) GetUserRankOk() (*UserWithRankModelUserRank, bool)`

GetUserRankOk returns a tuple with the UserRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRank

`func (o *UserWithRankModel) SetUserRank(v UserWithRankModelUserRank)`

SetUserRank sets UserRank field to given value.

### HasUserRank

`func (o *UserWithRankModel) HasUserRank() bool`

HasUserRank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


