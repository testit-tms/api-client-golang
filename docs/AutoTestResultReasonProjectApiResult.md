# AutoTestResultReasonProjectApiResult

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
**Projects** | [**[]ProjectNameApiResult**](ProjectNameApiResult.md) | Projects names | 
**FailureClassRegexes** | [**[]FailureClassRegexApiResult**](FailureClassRegexApiResult.md) | Failure category regexes | 

## Methods

### NewAutoTestResultReasonProjectApiResult

`func NewAutoTestResultReasonProjectApiResult(id string, isDeleted bool, failureCategory FailureCategory, createdDate time.Time, createdById string, projects []ProjectNameApiResult, failureClassRegexes []FailureClassRegexApiResult, ) *AutoTestResultReasonProjectApiResult`

NewAutoTestResultReasonProjectApiResult instantiates a new AutoTestResultReasonProjectApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestResultReasonProjectApiResultWithDefaults

`func NewAutoTestResultReasonProjectApiResultWithDefaults() *AutoTestResultReasonProjectApiResult`

NewAutoTestResultReasonProjectApiResultWithDefaults instantiates a new AutoTestResultReasonProjectApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoTestResultReasonProjectApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoTestResultReasonProjectApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoTestResultReasonProjectApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AutoTestResultReasonProjectApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoTestResultReasonProjectApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoTestResultReasonProjectApiResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoTestResultReasonProjectApiResult) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AutoTestResultReasonProjectApiResult) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AutoTestResultReasonProjectApiResult) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsDeleted

`func (o *AutoTestResultReasonProjectApiResult) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *AutoTestResultReasonProjectApiResult) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *AutoTestResultReasonProjectApiResult) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetFailureCategory

`func (o *AutoTestResultReasonProjectApiResult) GetFailureCategory() FailureCategory`

GetFailureCategory returns the FailureCategory field if non-nil, zero value otherwise.

### GetFailureCategoryOk

`func (o *AutoTestResultReasonProjectApiResult) GetFailureCategoryOk() (*FailureCategory, bool)`

GetFailureCategoryOk returns a tuple with the FailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategory

`func (o *AutoTestResultReasonProjectApiResult) SetFailureCategory(v FailureCategory)`

SetFailureCategory sets FailureCategory field to given value.


### GetCreatedDate

`func (o *AutoTestResultReasonProjectApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AutoTestResultReasonProjectApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AutoTestResultReasonProjectApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetCreatedById

`func (o *AutoTestResultReasonProjectApiResult) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *AutoTestResultReasonProjectApiResult) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *AutoTestResultReasonProjectApiResult) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedDate

`func (o *AutoTestResultReasonProjectApiResult) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *AutoTestResultReasonProjectApiResult) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *AutoTestResultReasonProjectApiResult) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *AutoTestResultReasonProjectApiResult) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *AutoTestResultReasonProjectApiResult) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *AutoTestResultReasonProjectApiResult) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedById

`func (o *AutoTestResultReasonProjectApiResult) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *AutoTestResultReasonProjectApiResult) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *AutoTestResultReasonProjectApiResult) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *AutoTestResultReasonProjectApiResult) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *AutoTestResultReasonProjectApiResult) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *AutoTestResultReasonProjectApiResult) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetProjects

`func (o *AutoTestResultReasonProjectApiResult) GetProjects() []ProjectNameApiResult`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *AutoTestResultReasonProjectApiResult) GetProjectsOk() (*[]ProjectNameApiResult, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *AutoTestResultReasonProjectApiResult) SetProjects(v []ProjectNameApiResult)`

SetProjects sets Projects field to given value.


### GetFailureClassRegexes

`func (o *AutoTestResultReasonProjectApiResult) GetFailureClassRegexes() []FailureClassRegexApiResult`

GetFailureClassRegexes returns the FailureClassRegexes field if non-nil, zero value otherwise.

### GetFailureClassRegexesOk

`func (o *AutoTestResultReasonProjectApiResult) GetFailureClassRegexesOk() (*[]FailureClassRegexApiResult, bool)`

GetFailureClassRegexesOk returns a tuple with the FailureClassRegexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureClassRegexes

`func (o *AutoTestResultReasonProjectApiResult) SetFailureClassRegexes(v []FailureClassRegexApiResult)`

SetFailureClassRegexes sets FailureClassRegexes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


