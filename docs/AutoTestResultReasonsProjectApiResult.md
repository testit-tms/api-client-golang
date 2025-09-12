# AutoTestResultReasonsProjectApiResult

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

### NewAutoTestResultReasonsProjectApiResult

`func NewAutoTestResultReasonsProjectApiResult(id string, failureCategory FailureCategory, createdDate time.Time, createdById string, failureCategoryId int32, regexCount int32, projects []ProjectNameApiResult, projectIds []string, ) *AutoTestResultReasonsProjectApiResult`

NewAutoTestResultReasonsProjectApiResult instantiates a new AutoTestResultReasonsProjectApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestResultReasonsProjectApiResultWithDefaults

`func NewAutoTestResultReasonsProjectApiResultWithDefaults() *AutoTestResultReasonsProjectApiResult`

NewAutoTestResultReasonsProjectApiResultWithDefaults instantiates a new AutoTestResultReasonsProjectApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoTestResultReasonsProjectApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoTestResultReasonsProjectApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoTestResultReasonsProjectApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AutoTestResultReasonsProjectApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoTestResultReasonsProjectApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoTestResultReasonsProjectApiResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoTestResultReasonsProjectApiResult) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AutoTestResultReasonsProjectApiResult) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AutoTestResultReasonsProjectApiResult) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFailureCategory

`func (o *AutoTestResultReasonsProjectApiResult) GetFailureCategory() FailureCategory`

GetFailureCategory returns the FailureCategory field if non-nil, zero value otherwise.

### GetFailureCategoryOk

`func (o *AutoTestResultReasonsProjectApiResult) GetFailureCategoryOk() (*FailureCategory, bool)`

GetFailureCategoryOk returns a tuple with the FailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategory

`func (o *AutoTestResultReasonsProjectApiResult) SetFailureCategory(v FailureCategory)`

SetFailureCategory sets FailureCategory field to given value.


### GetCreatedDate

`func (o *AutoTestResultReasonsProjectApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AutoTestResultReasonsProjectApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AutoTestResultReasonsProjectApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetCreatedById

`func (o *AutoTestResultReasonsProjectApiResult) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *AutoTestResultReasonsProjectApiResult) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *AutoTestResultReasonsProjectApiResult) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedDate

`func (o *AutoTestResultReasonsProjectApiResult) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *AutoTestResultReasonsProjectApiResult) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *AutoTestResultReasonsProjectApiResult) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *AutoTestResultReasonsProjectApiResult) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *AutoTestResultReasonsProjectApiResult) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *AutoTestResultReasonsProjectApiResult) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedById

`func (o *AutoTestResultReasonsProjectApiResult) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *AutoTestResultReasonsProjectApiResult) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *AutoTestResultReasonsProjectApiResult) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *AutoTestResultReasonsProjectApiResult) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *AutoTestResultReasonsProjectApiResult) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *AutoTestResultReasonsProjectApiResult) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetFailureCategoryId

`func (o *AutoTestResultReasonsProjectApiResult) GetFailureCategoryId() int32`

GetFailureCategoryId returns the FailureCategoryId field if non-nil, zero value otherwise.

### GetFailureCategoryIdOk

`func (o *AutoTestResultReasonsProjectApiResult) GetFailureCategoryIdOk() (*int32, bool)`

GetFailureCategoryIdOk returns a tuple with the FailureCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategoryId

`func (o *AutoTestResultReasonsProjectApiResult) SetFailureCategoryId(v int32)`

SetFailureCategoryId sets FailureCategoryId field to given value.


### GetRegexCount

`func (o *AutoTestResultReasonsProjectApiResult) GetRegexCount() int32`

GetRegexCount returns the RegexCount field if non-nil, zero value otherwise.

### GetRegexCountOk

`func (o *AutoTestResultReasonsProjectApiResult) GetRegexCountOk() (*int32, bool)`

GetRegexCountOk returns a tuple with the RegexCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexCount

`func (o *AutoTestResultReasonsProjectApiResult) SetRegexCount(v int32)`

SetRegexCount sets RegexCount field to given value.


### GetProjects

`func (o *AutoTestResultReasonsProjectApiResult) GetProjects() []ProjectNameApiResult`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *AutoTestResultReasonsProjectApiResult) GetProjectsOk() (*[]ProjectNameApiResult, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *AutoTestResultReasonsProjectApiResult) SetProjects(v []ProjectNameApiResult)`

SetProjects sets Projects field to given value.


### GetProjectIds

`func (o *AutoTestResultReasonsProjectApiResult) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *AutoTestResultReasonsProjectApiResult) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *AutoTestResultReasonsProjectApiResult) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


