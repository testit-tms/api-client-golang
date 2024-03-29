# ConfigurationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Parameters** | Pointer to **map[string]string** |  | [optional] 
**ProjectId** | **string** | This property is used to link configuration with project | 
**IsDefault** | **bool** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**CreatedDate** | **time.Time** |  | 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**CreatedById** | **string** |  | 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**GlobalId** | **int64** |  | 
**Id** | **string** | Unique ID of the entity | 
**IsDeleted** | **bool** | Indicates if the entity is deleted | 

## Methods

### NewConfigurationModel

`func NewConfigurationModel(projectId string, isDefault bool, createdDate time.Time, createdById string, globalId int64, id string, isDeleted bool, ) *ConfigurationModel`

NewConfigurationModel instantiates a new ConfigurationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationModelWithDefaults

`func NewConfigurationModelWithDefaults() *ConfigurationModel`

NewConfigurationModelWithDefaults instantiates a new ConfigurationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ConfigurationModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigurationModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigurationModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigurationModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ConfigurationModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConfigurationModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetParameters

`func (o *ConfigurationModel) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ConfigurationModel) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ConfigurationModel) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ConfigurationModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *ConfigurationModel) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *ConfigurationModel) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetProjectId

`func (o *ConfigurationModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ConfigurationModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ConfigurationModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetIsDefault

`func (o *ConfigurationModel) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ConfigurationModel) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ConfigurationModel) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetName

`func (o *ConfigurationModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigurationModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigurationModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigurationModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConfigurationModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConfigurationModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreatedDate

`func (o *ConfigurationModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ConfigurationModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ConfigurationModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *ConfigurationModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ConfigurationModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ConfigurationModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *ConfigurationModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *ConfigurationModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *ConfigurationModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *ConfigurationModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *ConfigurationModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *ConfigurationModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedById

`func (o *ConfigurationModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *ConfigurationModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *ConfigurationModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *ConfigurationModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *ConfigurationModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *ConfigurationModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetGlobalId

`func (o *ConfigurationModel) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *ConfigurationModel) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *ConfigurationModel) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.


### GetId

`func (o *ConfigurationModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigurationModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigurationModel) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleted

`func (o *ConfigurationModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ConfigurationModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ConfigurationModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


