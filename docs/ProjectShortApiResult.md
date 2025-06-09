# ProjectShortApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of project | 
**IsDeleted** | **bool** | Indicates whether the project is deleted | 
**GlobalId** | **int64** | Global ID of project | 
**Name** | **string** | Name of project | 
**Type** | [**ProjectType**](ProjectType.md) | Type of the project | 
**IsFavorite** | **bool** | Indicates if the project is marked as favorite | 

## Methods

### NewProjectShortApiResult

`func NewProjectShortApiResult(id string, isDeleted bool, globalId int64, name string, type_ ProjectType, isFavorite bool, ) *ProjectShortApiResult`

NewProjectShortApiResult instantiates a new ProjectShortApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectShortApiResultWithDefaults

`func NewProjectShortApiResultWithDefaults() *ProjectShortApiResult`

NewProjectShortApiResultWithDefaults instantiates a new ProjectShortApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectShortApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectShortApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectShortApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleted

`func (o *ProjectShortApiResult) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ProjectShortApiResult) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ProjectShortApiResult) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetGlobalId

`func (o *ProjectShortApiResult) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *ProjectShortApiResult) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *ProjectShortApiResult) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.


### GetName

`func (o *ProjectShortApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectShortApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectShortApiResult) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ProjectShortApiResult) GetType() ProjectType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProjectShortApiResult) GetTypeOk() (*ProjectType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProjectShortApiResult) SetType(v ProjectType)`

SetType sets Type field to given value.


### GetIsFavorite

`func (o *ProjectShortApiResult) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *ProjectShortApiResult) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *ProjectShortApiResult) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


