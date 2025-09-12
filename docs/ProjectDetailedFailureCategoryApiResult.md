# ProjectDetailedFailureCategoryApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Failure category identifier | 
**Name** | Pointer to **NullableString** | Failure category name | [optional] 
**IsDeleted** | **bool** | Indicates if the entity is deleted | 
**FailureCategory** | [**FailureCategory**](FailureCategory.md) | Category type | 
**CreatedDate** | **time.Time** | Failure category creation date | 
**CreatedById** | **string** | Failure category creator identifier | 
**ModifiedDate** | Pointer to **NullableTime** | Failure category last modification date | [optional] 
**ModifiedById** | Pointer to **NullableString** | Failure category last modifier identifier | [optional] 
**FailureClassRegexes** | [**[]FailureClassRegexApiResult**](FailureClassRegexApiResult.md) | Failure category regexes | 
**ProjectsCount** | **int32** | Projects relations count | 

## Methods

### NewProjectDetailedFailureCategoryApiResult

`func NewProjectDetailedFailureCategoryApiResult(id string, isDeleted bool, failureCategory FailureCategory, createdDate time.Time, createdById string, failureClassRegexes []FailureClassRegexApiResult, projectsCount int32, ) *ProjectDetailedFailureCategoryApiResult`

NewProjectDetailedFailureCategoryApiResult instantiates a new ProjectDetailedFailureCategoryApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDetailedFailureCategoryApiResultWithDefaults

`func NewProjectDetailedFailureCategoryApiResultWithDefaults() *ProjectDetailedFailureCategoryApiResult`

NewProjectDetailedFailureCategoryApiResultWithDefaults instantiates a new ProjectDetailedFailureCategoryApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectDetailedFailureCategoryApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectDetailedFailureCategoryApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectDetailedFailureCategoryApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProjectDetailedFailureCategoryApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectDetailedFailureCategoryApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectDetailedFailureCategoryApiResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectDetailedFailureCategoryApiResult) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProjectDetailedFailureCategoryApiResult) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectDetailedFailureCategoryApiResult) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsDeleted

`func (o *ProjectDetailedFailureCategoryApiResult) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ProjectDetailedFailureCategoryApiResult) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ProjectDetailedFailureCategoryApiResult) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetFailureCategory

`func (o *ProjectDetailedFailureCategoryApiResult) GetFailureCategory() FailureCategory`

GetFailureCategory returns the FailureCategory field if non-nil, zero value otherwise.

### GetFailureCategoryOk

`func (o *ProjectDetailedFailureCategoryApiResult) GetFailureCategoryOk() (*FailureCategory, bool)`

GetFailureCategoryOk returns a tuple with the FailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategory

`func (o *ProjectDetailedFailureCategoryApiResult) SetFailureCategory(v FailureCategory)`

SetFailureCategory sets FailureCategory field to given value.


### GetCreatedDate

`func (o *ProjectDetailedFailureCategoryApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ProjectDetailedFailureCategoryApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ProjectDetailedFailureCategoryApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetCreatedById

`func (o *ProjectDetailedFailureCategoryApiResult) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *ProjectDetailedFailureCategoryApiResult) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *ProjectDetailedFailureCategoryApiResult) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedDate

`func (o *ProjectDetailedFailureCategoryApiResult) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ProjectDetailedFailureCategoryApiResult) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ProjectDetailedFailureCategoryApiResult) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *ProjectDetailedFailureCategoryApiResult) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *ProjectDetailedFailureCategoryApiResult) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *ProjectDetailedFailureCategoryApiResult) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedById

`func (o *ProjectDetailedFailureCategoryApiResult) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *ProjectDetailedFailureCategoryApiResult) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *ProjectDetailedFailureCategoryApiResult) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *ProjectDetailedFailureCategoryApiResult) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *ProjectDetailedFailureCategoryApiResult) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *ProjectDetailedFailureCategoryApiResult) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetFailureClassRegexes

`func (o *ProjectDetailedFailureCategoryApiResult) GetFailureClassRegexes() []FailureClassRegexApiResult`

GetFailureClassRegexes returns the FailureClassRegexes field if non-nil, zero value otherwise.

### GetFailureClassRegexesOk

`func (o *ProjectDetailedFailureCategoryApiResult) GetFailureClassRegexesOk() (*[]FailureClassRegexApiResult, bool)`

GetFailureClassRegexesOk returns a tuple with the FailureClassRegexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureClassRegexes

`func (o *ProjectDetailedFailureCategoryApiResult) SetFailureClassRegexes(v []FailureClassRegexApiResult)`

SetFailureClassRegexes sets FailureClassRegexes field to given value.


### GetProjectsCount

`func (o *ProjectDetailedFailureCategoryApiResult) GetProjectsCount() int32`

GetProjectsCount returns the ProjectsCount field if non-nil, zero value otherwise.

### GetProjectsCountOk

`func (o *ProjectDetailedFailureCategoryApiResult) GetProjectsCountOk() (*int32, bool)`

GetProjectsCountOk returns a tuple with the ProjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsCount

`func (o *ProjectDetailedFailureCategoryApiResult) SetProjectsCount(v int32)`

SetProjectsCount sets ProjectsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


