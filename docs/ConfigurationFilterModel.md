# ConfigurationFilterModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectIds** | Pointer to **[]string** | Collection of identifiers of projects from which configurations will be taken | [optional] 
**Name** | Pointer to **NullableString** | Filter to search by name (case-insensitive, partial match) | [optional] 
**IsDeleted** | Pointer to **NullableBool** | Is configurations deleted or existing | [optional] 
**GlobalIds** | Pointer to **[]int64** | Collection of global (integer) identifiers to filter configurations | [optional] 

## Methods

### NewConfigurationFilterModel

`func NewConfigurationFilterModel() *ConfigurationFilterModel`

NewConfigurationFilterModel instantiates a new ConfigurationFilterModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationFilterModelWithDefaults

`func NewConfigurationFilterModelWithDefaults() *ConfigurationFilterModel`

NewConfigurationFilterModelWithDefaults instantiates a new ConfigurationFilterModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectIds

`func (o *ConfigurationFilterModel) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *ConfigurationFilterModel) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *ConfigurationFilterModel) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *ConfigurationFilterModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *ConfigurationFilterModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *ConfigurationFilterModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetName

`func (o *ConfigurationFilterModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigurationFilterModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigurationFilterModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigurationFilterModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConfigurationFilterModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConfigurationFilterModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsDeleted

`func (o *ConfigurationFilterModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ConfigurationFilterModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ConfigurationFilterModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ConfigurationFilterModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *ConfigurationFilterModel) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *ConfigurationFilterModel) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetGlobalIds

`func (o *ConfigurationFilterModel) GetGlobalIds() []int64`

GetGlobalIds returns the GlobalIds field if non-nil, zero value otherwise.

### GetGlobalIdsOk

`func (o *ConfigurationFilterModel) GetGlobalIdsOk() (*[]int64, bool)`

GetGlobalIdsOk returns a tuple with the GlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIds

`func (o *ConfigurationFilterModel) SetGlobalIds(v []int64)`

SetGlobalIds sets GlobalIds field to given value.

### HasGlobalIds

`func (o *ConfigurationFilterModel) HasGlobalIds() bool`

HasGlobalIds returns a boolean if a field has been set.

### SetGlobalIdsNil

`func (o *ConfigurationFilterModel) SetGlobalIdsNil(b bool)`

 SetGlobalIdsNil sets the value for GlobalIds to be an explicit nil

### UnsetGlobalIds
`func (o *ConfigurationFilterModel) UnsetGlobalIds()`

UnsetGlobalIds ensures that no value is present for GlobalIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


