# TestRunShortGetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the test run | 
**Name** | Pointer to **NullableString** | Name of the test run | [optional] 
**State** | [**TestRunState**](TestRunState.md) |  | 
**CreatedDate** | **time.Time** | Date when the test run was created | 
**StartedDate** | Pointer to **NullableTime** | Date when the test run was started | [optional] 
**CompletedDate** | Pointer to **NullableTime** | Completion date of the test run | [optional] 
**CreatedById** | **string** | Unique ID of user who created the test run | 
**ModifiedById** | Pointer to **NullableString** | Unique ID of user who modified the test run last time | [optional] 
**IsDeleted** | **bool** | Is the test run is deleted | 
**AutoTestsCount** | **int32** | Number of AutoTests run in the test run | 
**Statistics** | Pointer to [**NullableTestRunShortGetModelStatistics**](TestRunShortGetModelStatistics.md) |  | [optional] 

## Methods

### NewTestRunShortGetModel

`func NewTestRunShortGetModel(id string, state TestRunState, createdDate time.Time, createdById string, isDeleted bool, autoTestsCount int32, ) *TestRunShortGetModel`

NewTestRunShortGetModel instantiates a new TestRunShortGetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunShortGetModelWithDefaults

`func NewTestRunShortGetModelWithDefaults() *TestRunShortGetModel`

NewTestRunShortGetModelWithDefaults instantiates a new TestRunShortGetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestRunShortGetModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestRunShortGetModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestRunShortGetModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TestRunShortGetModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestRunShortGetModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestRunShortGetModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestRunShortGetModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestRunShortGetModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestRunShortGetModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetState

`func (o *TestRunShortGetModel) GetState() TestRunState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TestRunShortGetModel) GetStateOk() (*TestRunState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TestRunShortGetModel) SetState(v TestRunState)`

SetState sets State field to given value.


### GetCreatedDate

`func (o *TestRunShortGetModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestRunShortGetModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestRunShortGetModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetStartedDate

`func (o *TestRunShortGetModel) GetStartedDate() time.Time`

GetStartedDate returns the StartedDate field if non-nil, zero value otherwise.

### GetStartedDateOk

`func (o *TestRunShortGetModel) GetStartedDateOk() (*time.Time, bool)`

GetStartedDateOk returns a tuple with the StartedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedDate

`func (o *TestRunShortGetModel) SetStartedDate(v time.Time)`

SetStartedDate sets StartedDate field to given value.

### HasStartedDate

`func (o *TestRunShortGetModel) HasStartedDate() bool`

HasStartedDate returns a boolean if a field has been set.

### SetStartedDateNil

`func (o *TestRunShortGetModel) SetStartedDateNil(b bool)`

 SetStartedDateNil sets the value for StartedDate to be an explicit nil

### UnsetStartedDate
`func (o *TestRunShortGetModel) UnsetStartedDate()`

UnsetStartedDate ensures that no value is present for StartedDate, not even an explicit nil
### GetCompletedDate

`func (o *TestRunShortGetModel) GetCompletedDate() time.Time`

GetCompletedDate returns the CompletedDate field if non-nil, zero value otherwise.

### GetCompletedDateOk

`func (o *TestRunShortGetModel) GetCompletedDateOk() (*time.Time, bool)`

GetCompletedDateOk returns a tuple with the CompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDate

`func (o *TestRunShortGetModel) SetCompletedDate(v time.Time)`

SetCompletedDate sets CompletedDate field to given value.

### HasCompletedDate

`func (o *TestRunShortGetModel) HasCompletedDate() bool`

HasCompletedDate returns a boolean if a field has been set.

### SetCompletedDateNil

`func (o *TestRunShortGetModel) SetCompletedDateNil(b bool)`

 SetCompletedDateNil sets the value for CompletedDate to be an explicit nil

### UnsetCompletedDate
`func (o *TestRunShortGetModel) UnsetCompletedDate()`

UnsetCompletedDate ensures that no value is present for CompletedDate, not even an explicit nil
### GetCreatedById

`func (o *TestRunShortGetModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TestRunShortGetModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TestRunShortGetModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedById

`func (o *TestRunShortGetModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *TestRunShortGetModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *TestRunShortGetModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *TestRunShortGetModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *TestRunShortGetModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *TestRunShortGetModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetIsDeleted

`func (o *TestRunShortGetModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TestRunShortGetModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TestRunShortGetModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetAutoTestsCount

`func (o *TestRunShortGetModel) GetAutoTestsCount() int32`

GetAutoTestsCount returns the AutoTestsCount field if non-nil, zero value otherwise.

### GetAutoTestsCountOk

`func (o *TestRunShortGetModel) GetAutoTestsCountOk() (*int32, bool)`

GetAutoTestsCountOk returns a tuple with the AutoTestsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestsCount

`func (o *TestRunShortGetModel) SetAutoTestsCount(v int32)`

SetAutoTestsCount sets AutoTestsCount field to given value.


### GetStatistics

`func (o *TestRunShortGetModel) GetStatistics() TestRunShortGetModelStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *TestRunShortGetModel) GetStatisticsOk() (*TestRunShortGetModelStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *TestRunShortGetModel) SetStatistics(v TestRunShortGetModelStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *TestRunShortGetModel) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### SetStatisticsNil

`func (o *TestRunShortGetModel) SetStatisticsNil(b bool)`

 SetStatisticsNil sets the value for Statistics to be an explicit nil

### UnsetStatistics
`func (o *TestRunShortGetModel) UnsetStatistics()`

UnsetStatistics ensures that no value is present for Statistics, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


