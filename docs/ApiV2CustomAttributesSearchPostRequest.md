# ApiV2CustomAttributesSearchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of attribute | [optional] 
**ProjectIds** | Pointer to **[]string** | Unique IDs of projects where attribute is in use | [optional] 
**CustomAttributeIds** | Pointer to **[]string** | Unique IDs of attributes for search restriction | [optional] 
**CustomAttributeTypes** | Pointer to [**[]CustomAttributeTypesEnum**](CustomAttributeTypesEnum.md) | Collection of attribute types | [optional] 
**IsGlobal** | Pointer to **NullableBool** | Indicates whether the attribute is available across all projects | [optional] 
**IsDeleted** | Pointer to **NullableBool** | Indicates whether the attribute is deleted | [optional] 

## Methods

### NewApiV2CustomAttributesSearchPostRequest

`func NewApiV2CustomAttributesSearchPostRequest() *ApiV2CustomAttributesSearchPostRequest`

NewApiV2CustomAttributesSearchPostRequest instantiates a new ApiV2CustomAttributesSearchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2CustomAttributesSearchPostRequestWithDefaults

`func NewApiV2CustomAttributesSearchPostRequestWithDefaults() *ApiV2CustomAttributesSearchPostRequest`

NewApiV2CustomAttributesSearchPostRequestWithDefaults instantiates a new ApiV2CustomAttributesSearchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiV2CustomAttributesSearchPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2CustomAttributesSearchPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2CustomAttributesSearchPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV2CustomAttributesSearchPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApiV2CustomAttributesSearchPostRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApiV2CustomAttributesSearchPostRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProjectIds

`func (o *ApiV2CustomAttributesSearchPostRequest) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *ApiV2CustomAttributesSearchPostRequest) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *ApiV2CustomAttributesSearchPostRequest) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *ApiV2CustomAttributesSearchPostRequest) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *ApiV2CustomAttributesSearchPostRequest) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *ApiV2CustomAttributesSearchPostRequest) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetCustomAttributeIds

`func (o *ApiV2CustomAttributesSearchPostRequest) GetCustomAttributeIds() []string`

GetCustomAttributeIds returns the CustomAttributeIds field if non-nil, zero value otherwise.

### GetCustomAttributeIdsOk

`func (o *ApiV2CustomAttributesSearchPostRequest) GetCustomAttributeIdsOk() (*[]string, bool)`

GetCustomAttributeIdsOk returns a tuple with the CustomAttributeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributeIds

`func (o *ApiV2CustomAttributesSearchPostRequest) SetCustomAttributeIds(v []string)`

SetCustomAttributeIds sets CustomAttributeIds field to given value.

### HasCustomAttributeIds

`func (o *ApiV2CustomAttributesSearchPostRequest) HasCustomAttributeIds() bool`

HasCustomAttributeIds returns a boolean if a field has been set.

### SetCustomAttributeIdsNil

`func (o *ApiV2CustomAttributesSearchPostRequest) SetCustomAttributeIdsNil(b bool)`

 SetCustomAttributeIdsNil sets the value for CustomAttributeIds to be an explicit nil

### UnsetCustomAttributeIds
`func (o *ApiV2CustomAttributesSearchPostRequest) UnsetCustomAttributeIds()`

UnsetCustomAttributeIds ensures that no value is present for CustomAttributeIds, not even an explicit nil
### GetCustomAttributeTypes

`func (o *ApiV2CustomAttributesSearchPostRequest) GetCustomAttributeTypes() []CustomAttributeTypesEnum`

GetCustomAttributeTypes returns the CustomAttributeTypes field if non-nil, zero value otherwise.

### GetCustomAttributeTypesOk

`func (o *ApiV2CustomAttributesSearchPostRequest) GetCustomAttributeTypesOk() (*[]CustomAttributeTypesEnum, bool)`

GetCustomAttributeTypesOk returns a tuple with the CustomAttributeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributeTypes

`func (o *ApiV2CustomAttributesSearchPostRequest) SetCustomAttributeTypes(v []CustomAttributeTypesEnum)`

SetCustomAttributeTypes sets CustomAttributeTypes field to given value.

### HasCustomAttributeTypes

`func (o *ApiV2CustomAttributesSearchPostRequest) HasCustomAttributeTypes() bool`

HasCustomAttributeTypes returns a boolean if a field has been set.

### SetCustomAttributeTypesNil

`func (o *ApiV2CustomAttributesSearchPostRequest) SetCustomAttributeTypesNil(b bool)`

 SetCustomAttributeTypesNil sets the value for CustomAttributeTypes to be an explicit nil

### UnsetCustomAttributeTypes
`func (o *ApiV2CustomAttributesSearchPostRequest) UnsetCustomAttributeTypes()`

UnsetCustomAttributeTypes ensures that no value is present for CustomAttributeTypes, not even an explicit nil
### GetIsGlobal

`func (o *ApiV2CustomAttributesSearchPostRequest) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *ApiV2CustomAttributesSearchPostRequest) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *ApiV2CustomAttributesSearchPostRequest) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.

### HasIsGlobal

`func (o *ApiV2CustomAttributesSearchPostRequest) HasIsGlobal() bool`

HasIsGlobal returns a boolean if a field has been set.

### SetIsGlobalNil

`func (o *ApiV2CustomAttributesSearchPostRequest) SetIsGlobalNil(b bool)`

 SetIsGlobalNil sets the value for IsGlobal to be an explicit nil

### UnsetIsGlobal
`func (o *ApiV2CustomAttributesSearchPostRequest) UnsetIsGlobal()`

UnsetIsGlobal ensures that no value is present for IsGlobal, not even an explicit nil
### GetIsDeleted

`func (o *ApiV2CustomAttributesSearchPostRequest) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ApiV2CustomAttributesSearchPostRequest) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ApiV2CustomAttributesSearchPostRequest) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ApiV2CustomAttributesSearchPostRequest) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *ApiV2CustomAttributesSearchPostRequest) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *ApiV2CustomAttributesSearchPostRequest) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


