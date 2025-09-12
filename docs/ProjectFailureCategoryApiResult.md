# ProjectFailureCategoryApiResult

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
**ProjectsCount** | **int32** | Projects relations count | 
**FailureCategoryId** | **int32** | Category type index | 
**RegexCount** | **int32** | Regexes count | 

## Methods

### NewProjectFailureCategoryApiResult

`func NewProjectFailureCategoryApiResult(id string, failureCategory FailureCategory, createdDate time.Time, createdById string, projectsCount int32, failureCategoryId int32, regexCount int32, ) *ProjectFailureCategoryApiResult`

NewProjectFailureCategoryApiResult instantiates a new ProjectFailureCategoryApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectFailureCategoryApiResultWithDefaults

`func NewProjectFailureCategoryApiResultWithDefaults() *ProjectFailureCategoryApiResult`

NewProjectFailureCategoryApiResultWithDefaults instantiates a new ProjectFailureCategoryApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectFailureCategoryApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectFailureCategoryApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectFailureCategoryApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProjectFailureCategoryApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectFailureCategoryApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectFailureCategoryApiResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectFailureCategoryApiResult) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProjectFailureCategoryApiResult) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectFailureCategoryApiResult) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFailureCategory

`func (o *ProjectFailureCategoryApiResult) GetFailureCategory() FailureCategory`

GetFailureCategory returns the FailureCategory field if non-nil, zero value otherwise.

### GetFailureCategoryOk

`func (o *ProjectFailureCategoryApiResult) GetFailureCategoryOk() (*FailureCategory, bool)`

GetFailureCategoryOk returns a tuple with the FailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategory

`func (o *ProjectFailureCategoryApiResult) SetFailureCategory(v FailureCategory)`

SetFailureCategory sets FailureCategory field to given value.


### GetCreatedDate

`func (o *ProjectFailureCategoryApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ProjectFailureCategoryApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ProjectFailureCategoryApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetCreatedById

`func (o *ProjectFailureCategoryApiResult) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *ProjectFailureCategoryApiResult) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *ProjectFailureCategoryApiResult) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedDate

`func (o *ProjectFailureCategoryApiResult) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ProjectFailureCategoryApiResult) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ProjectFailureCategoryApiResult) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *ProjectFailureCategoryApiResult) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *ProjectFailureCategoryApiResult) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *ProjectFailureCategoryApiResult) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedById

`func (o *ProjectFailureCategoryApiResult) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *ProjectFailureCategoryApiResult) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *ProjectFailureCategoryApiResult) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *ProjectFailureCategoryApiResult) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *ProjectFailureCategoryApiResult) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *ProjectFailureCategoryApiResult) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetProjectsCount

`func (o *ProjectFailureCategoryApiResult) GetProjectsCount() int32`

GetProjectsCount returns the ProjectsCount field if non-nil, zero value otherwise.

### GetProjectsCountOk

`func (o *ProjectFailureCategoryApiResult) GetProjectsCountOk() (*int32, bool)`

GetProjectsCountOk returns a tuple with the ProjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsCount

`func (o *ProjectFailureCategoryApiResult) SetProjectsCount(v int32)`

SetProjectsCount sets ProjectsCount field to given value.


### GetFailureCategoryId

`func (o *ProjectFailureCategoryApiResult) GetFailureCategoryId() int32`

GetFailureCategoryId returns the FailureCategoryId field if non-nil, zero value otherwise.

### GetFailureCategoryIdOk

`func (o *ProjectFailureCategoryApiResult) GetFailureCategoryIdOk() (*int32, bool)`

GetFailureCategoryIdOk returns a tuple with the FailureCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategoryId

`func (o *ProjectFailureCategoryApiResult) SetFailureCategoryId(v int32)`

SetFailureCategoryId sets FailureCategoryId field to given value.


### GetRegexCount

`func (o *ProjectFailureCategoryApiResult) GetRegexCount() int32`

GetRegexCount returns the RegexCount field if non-nil, zero value otherwise.

### GetRegexCountOk

`func (o *ProjectFailureCategoryApiResult) GetRegexCountOk() (*int32, bool)`

GetRegexCountOk returns a tuple with the RegexCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexCount

`func (o *ProjectFailureCategoryApiResult) SetRegexCount(v int32)`

SetRegexCount sets RegexCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


