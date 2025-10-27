# FailureCategoryItemApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Failure category identifier | 
**Name** | Pointer to **NullableString** | Failure category name | [optional] 
**FailureCategory** | [**FailureCategory**](FailureCategory.md) | Category type | 
**CreatedDate** | **time.Time** | Failure category creation date | 
**CreatedById** | **string** | Failure category creator identifier | 
**ModifiedDate** | Pointer to **NullableTime** | Failure category last modification date | [optional] 
**ModifiedById** | Pointer to **NullableString** | Failure category last modifier identifier | [optional] 
**FailureCategoryId** | **int32** | Category type index | 
**RegexCount** | **int32** | Regexes count | 
**Projects** | [**[]ProjectNameApiResult**](ProjectNameApiResult.md) | Projects names | 
**ProjectIds** | **[]string** | Projects identifiers | 

## Methods

### NewFailureCategoryItemApiResult

`func NewFailureCategoryItemApiResult(id string, failureCategory FailureCategory, createdDate time.Time, createdById string, failureCategoryId int32, regexCount int32, projects []ProjectNameApiResult, projectIds []string, ) *FailureCategoryItemApiResult`

NewFailureCategoryItemApiResult instantiates a new FailureCategoryItemApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureCategoryItemApiResultWithDefaults

`func NewFailureCategoryItemApiResultWithDefaults() *FailureCategoryItemApiResult`

NewFailureCategoryItemApiResultWithDefaults instantiates a new FailureCategoryItemApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FailureCategoryItemApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FailureCategoryItemApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FailureCategoryItemApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FailureCategoryItemApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FailureCategoryItemApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FailureCategoryItemApiResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FailureCategoryItemApiResult) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FailureCategoryItemApiResult) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FailureCategoryItemApiResult) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFailureCategory

`func (o *FailureCategoryItemApiResult) GetFailureCategory() FailureCategory`

GetFailureCategory returns the FailureCategory field if non-nil, zero value otherwise.

### GetFailureCategoryOk

`func (o *FailureCategoryItemApiResult) GetFailureCategoryOk() (*FailureCategory, bool)`

GetFailureCategoryOk returns a tuple with the FailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategory

`func (o *FailureCategoryItemApiResult) SetFailureCategory(v FailureCategory)`

SetFailureCategory sets FailureCategory field to given value.


### GetCreatedDate

`func (o *FailureCategoryItemApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *FailureCategoryItemApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *FailureCategoryItemApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetCreatedById

`func (o *FailureCategoryItemApiResult) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *FailureCategoryItemApiResult) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *FailureCategoryItemApiResult) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedDate

`func (o *FailureCategoryItemApiResult) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *FailureCategoryItemApiResult) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *FailureCategoryItemApiResult) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *FailureCategoryItemApiResult) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *FailureCategoryItemApiResult) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *FailureCategoryItemApiResult) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedById

`func (o *FailureCategoryItemApiResult) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *FailureCategoryItemApiResult) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *FailureCategoryItemApiResult) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *FailureCategoryItemApiResult) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *FailureCategoryItemApiResult) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *FailureCategoryItemApiResult) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetFailureCategoryId

`func (o *FailureCategoryItemApiResult) GetFailureCategoryId() int32`

GetFailureCategoryId returns the FailureCategoryId field if non-nil, zero value otherwise.

### GetFailureCategoryIdOk

`func (o *FailureCategoryItemApiResult) GetFailureCategoryIdOk() (*int32, bool)`

GetFailureCategoryIdOk returns a tuple with the FailureCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategoryId

`func (o *FailureCategoryItemApiResult) SetFailureCategoryId(v int32)`

SetFailureCategoryId sets FailureCategoryId field to given value.


### GetRegexCount

`func (o *FailureCategoryItemApiResult) GetRegexCount() int32`

GetRegexCount returns the RegexCount field if non-nil, zero value otherwise.

### GetRegexCountOk

`func (o *FailureCategoryItemApiResult) GetRegexCountOk() (*int32, bool)`

GetRegexCountOk returns a tuple with the RegexCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexCount

`func (o *FailureCategoryItemApiResult) SetRegexCount(v int32)`

SetRegexCount sets RegexCount field to given value.


### GetProjects

`func (o *FailureCategoryItemApiResult) GetProjects() []ProjectNameApiResult`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *FailureCategoryItemApiResult) GetProjectsOk() (*[]ProjectNameApiResult, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *FailureCategoryItemApiResult) SetProjects(v []ProjectNameApiResult)`

SetProjects sets Projects field to given value.


### GetProjectIds

`func (o *FailureCategoryItemApiResult) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *FailureCategoryItemApiResult) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *FailureCategoryItemApiResult) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


