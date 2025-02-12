# TestRunShortApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the test run | 
**Name** | **string** | Name of the test run | 
**State** | [**TestRunState**](TestRunState.md) | Current state of the test run | 
**Status** | [**TestStatusApiResult**](TestStatusApiResult.md) | Current status of the test run | 
**CreatedDate** | **time.Time** | Date when the test run was created | 
**StartedDate** | Pointer to **NullableTime** | Date when the test run was started | [optional] 
**CompletedDate** | Pointer to **NullableTime** | Completion date of the test run | [optional] 
**CreatedById** | **string** | Unique ID of user who created the test run | 
**ModifiedById** | Pointer to **NullableString** | Unique ID of user who modified the test run last time | [optional] 
**IsDeleted** | **bool** | Is the test run is deleted | 
**AutoTestsCount** | **int32** | Number of AutoTests run in the test run | 
**Statistics** | [**TestResultsStatisticsApiResult**](TestResultsStatisticsApiResult.md) | Statistics of the test run | 
**TestResultsConfigurations** | [**[]ConfigurationShortApiResult**](ConfigurationShortApiResult.md) | Test results configurations | 

## Methods

### NewTestRunShortApiResult

`func NewTestRunShortApiResult(id string, name string, state TestRunState, status TestStatusApiResult, createdDate time.Time, createdById string, isDeleted bool, autoTestsCount int32, statistics TestResultsStatisticsApiResult, testResultsConfigurations []ConfigurationShortApiResult, ) *TestRunShortApiResult`

NewTestRunShortApiResult instantiates a new TestRunShortApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunShortApiResultWithDefaults

`func NewTestRunShortApiResultWithDefaults() *TestRunShortApiResult`

NewTestRunShortApiResultWithDefaults instantiates a new TestRunShortApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestRunShortApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestRunShortApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestRunShortApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TestRunShortApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestRunShortApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestRunShortApiResult) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *TestRunShortApiResult) GetState() TestRunState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TestRunShortApiResult) GetStateOk() (*TestRunState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TestRunShortApiResult) SetState(v TestRunState)`

SetState sets State field to given value.


### GetStatus

`func (o *TestRunShortApiResult) GetStatus() TestStatusApiResult`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestRunShortApiResult) GetStatusOk() (*TestStatusApiResult, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestRunShortApiResult) SetStatus(v TestStatusApiResult)`

SetStatus sets Status field to given value.


### GetCreatedDate

`func (o *TestRunShortApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestRunShortApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestRunShortApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetStartedDate

`func (o *TestRunShortApiResult) GetStartedDate() time.Time`

GetStartedDate returns the StartedDate field if non-nil, zero value otherwise.

### GetStartedDateOk

`func (o *TestRunShortApiResult) GetStartedDateOk() (*time.Time, bool)`

GetStartedDateOk returns a tuple with the StartedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedDate

`func (o *TestRunShortApiResult) SetStartedDate(v time.Time)`

SetStartedDate sets StartedDate field to given value.

### HasStartedDate

`func (o *TestRunShortApiResult) HasStartedDate() bool`

HasStartedDate returns a boolean if a field has been set.

### SetStartedDateNil

`func (o *TestRunShortApiResult) SetStartedDateNil(b bool)`

 SetStartedDateNil sets the value for StartedDate to be an explicit nil

### UnsetStartedDate
`func (o *TestRunShortApiResult) UnsetStartedDate()`

UnsetStartedDate ensures that no value is present for StartedDate, not even an explicit nil
### GetCompletedDate

`func (o *TestRunShortApiResult) GetCompletedDate() time.Time`

GetCompletedDate returns the CompletedDate field if non-nil, zero value otherwise.

### GetCompletedDateOk

`func (o *TestRunShortApiResult) GetCompletedDateOk() (*time.Time, bool)`

GetCompletedDateOk returns a tuple with the CompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDate

`func (o *TestRunShortApiResult) SetCompletedDate(v time.Time)`

SetCompletedDate sets CompletedDate field to given value.

### HasCompletedDate

`func (o *TestRunShortApiResult) HasCompletedDate() bool`

HasCompletedDate returns a boolean if a field has been set.

### SetCompletedDateNil

`func (o *TestRunShortApiResult) SetCompletedDateNil(b bool)`

 SetCompletedDateNil sets the value for CompletedDate to be an explicit nil

### UnsetCompletedDate
`func (o *TestRunShortApiResult) UnsetCompletedDate()`

UnsetCompletedDate ensures that no value is present for CompletedDate, not even an explicit nil
### GetCreatedById

`func (o *TestRunShortApiResult) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TestRunShortApiResult) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TestRunShortApiResult) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedById

`func (o *TestRunShortApiResult) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *TestRunShortApiResult) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *TestRunShortApiResult) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *TestRunShortApiResult) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *TestRunShortApiResult) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *TestRunShortApiResult) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetIsDeleted

`func (o *TestRunShortApiResult) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TestRunShortApiResult) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TestRunShortApiResult) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetAutoTestsCount

`func (o *TestRunShortApiResult) GetAutoTestsCount() int32`

GetAutoTestsCount returns the AutoTestsCount field if non-nil, zero value otherwise.

### GetAutoTestsCountOk

`func (o *TestRunShortApiResult) GetAutoTestsCountOk() (*int32, bool)`

GetAutoTestsCountOk returns a tuple with the AutoTestsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestsCount

`func (o *TestRunShortApiResult) SetAutoTestsCount(v int32)`

SetAutoTestsCount sets AutoTestsCount field to given value.


### GetStatistics

`func (o *TestRunShortApiResult) GetStatistics() TestResultsStatisticsApiResult`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *TestRunShortApiResult) GetStatisticsOk() (*TestResultsStatisticsApiResult, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *TestRunShortApiResult) SetStatistics(v TestResultsStatisticsApiResult)`

SetStatistics sets Statistics field to given value.


### GetTestResultsConfigurations

`func (o *TestRunShortApiResult) GetTestResultsConfigurations() []ConfigurationShortApiResult`

GetTestResultsConfigurations returns the TestResultsConfigurations field if non-nil, zero value otherwise.

### GetTestResultsConfigurationsOk

`func (o *TestRunShortApiResult) GetTestResultsConfigurationsOk() (*[]ConfigurationShortApiResult, bool)`

GetTestResultsConfigurationsOk returns a tuple with the TestResultsConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResultsConfigurations

`func (o *TestRunShortApiResult) SetTestResultsConfigurations(v []ConfigurationShortApiResult)`

SetTestResultsConfigurations sets TestResultsConfigurations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


