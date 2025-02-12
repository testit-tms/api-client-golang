# TestPointShortApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Test point unique internal identifier | 
**TesterId** | Pointer to **NullableString** | Tester who is responded for the test unique internal identifier | [optional] 
**WorkItemId** | Pointer to **NullableString** | Workitem to which test point relates unique identifier | [optional] 
**ConfigurationId** | Pointer to **NullableString** | Configuration to which test point relates unique identifier | [optional] 
**Status** | Pointer to **NullableString** | Test point status   Applies one of these values: Blocked, NoResults, Failed, Passed | [optional] 
**StatusModel** | [**TestStatusApiResult**](TestStatusApiResult.md) | Test point status | 
**LastTestResultId** | Pointer to **NullableString** | Last test result unique identifier | [optional] 
**IterationId** | **string** | Iteration unique identifier | 
**WorkItemMedianDuration** | Pointer to **NullableInt64** | Median duration of work item the test point represents | [optional] 
**TestSuiteId** | **string** | Test suite to which test point relates unique identifier | 

## Methods

### NewTestPointShortApiResult

`func NewTestPointShortApiResult(id string, statusModel TestStatusApiResult, iterationId string, testSuiteId string, ) *TestPointShortApiResult`

NewTestPointShortApiResult instantiates a new TestPointShortApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPointShortApiResultWithDefaults

`func NewTestPointShortApiResultWithDefaults() *TestPointShortApiResult`

NewTestPointShortApiResultWithDefaults instantiates a new TestPointShortApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestPointShortApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestPointShortApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestPointShortApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetTesterId

`func (o *TestPointShortApiResult) GetTesterId() string`

GetTesterId returns the TesterId field if non-nil, zero value otherwise.

### GetTesterIdOk

`func (o *TestPointShortApiResult) GetTesterIdOk() (*string, bool)`

GetTesterIdOk returns a tuple with the TesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesterId

`func (o *TestPointShortApiResult) SetTesterId(v string)`

SetTesterId sets TesterId field to given value.

### HasTesterId

`func (o *TestPointShortApiResult) HasTesterId() bool`

HasTesterId returns a boolean if a field has been set.

### SetTesterIdNil

`func (o *TestPointShortApiResult) SetTesterIdNil(b bool)`

 SetTesterIdNil sets the value for TesterId to be an explicit nil

### UnsetTesterId
`func (o *TestPointShortApiResult) UnsetTesterId()`

UnsetTesterId ensures that no value is present for TesterId, not even an explicit nil
### GetWorkItemId

`func (o *TestPointShortApiResult) GetWorkItemId() string`

GetWorkItemId returns the WorkItemId field if non-nil, zero value otherwise.

### GetWorkItemIdOk

`func (o *TestPointShortApiResult) GetWorkItemIdOk() (*string, bool)`

GetWorkItemIdOk returns a tuple with the WorkItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemId

`func (o *TestPointShortApiResult) SetWorkItemId(v string)`

SetWorkItemId sets WorkItemId field to given value.

### HasWorkItemId

`func (o *TestPointShortApiResult) HasWorkItemId() bool`

HasWorkItemId returns a boolean if a field has been set.

### SetWorkItemIdNil

`func (o *TestPointShortApiResult) SetWorkItemIdNil(b bool)`

 SetWorkItemIdNil sets the value for WorkItemId to be an explicit nil

### UnsetWorkItemId
`func (o *TestPointShortApiResult) UnsetWorkItemId()`

UnsetWorkItemId ensures that no value is present for WorkItemId, not even an explicit nil
### GetConfigurationId

`func (o *TestPointShortApiResult) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *TestPointShortApiResult) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *TestPointShortApiResult) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *TestPointShortApiResult) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### SetConfigurationIdNil

`func (o *TestPointShortApiResult) SetConfigurationIdNil(b bool)`

 SetConfigurationIdNil sets the value for ConfigurationId to be an explicit nil

### UnsetConfigurationId
`func (o *TestPointShortApiResult) UnsetConfigurationId()`

UnsetConfigurationId ensures that no value is present for ConfigurationId, not even an explicit nil
### GetStatus

`func (o *TestPointShortApiResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestPointShortApiResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestPointShortApiResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestPointShortApiResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TestPointShortApiResult) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TestPointShortApiResult) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusModel

`func (o *TestPointShortApiResult) GetStatusModel() TestStatusApiResult`

GetStatusModel returns the StatusModel field if non-nil, zero value otherwise.

### GetStatusModelOk

`func (o *TestPointShortApiResult) GetStatusModelOk() (*TestStatusApiResult, bool)`

GetStatusModelOk returns a tuple with the StatusModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusModel

`func (o *TestPointShortApiResult) SetStatusModel(v TestStatusApiResult)`

SetStatusModel sets StatusModel field to given value.


### GetLastTestResultId

`func (o *TestPointShortApiResult) GetLastTestResultId() string`

GetLastTestResultId returns the LastTestResultId field if non-nil, zero value otherwise.

### GetLastTestResultIdOk

`func (o *TestPointShortApiResult) GetLastTestResultIdOk() (*string, bool)`

GetLastTestResultIdOk returns a tuple with the LastTestResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultId

`func (o *TestPointShortApiResult) SetLastTestResultId(v string)`

SetLastTestResultId sets LastTestResultId field to given value.

### HasLastTestResultId

`func (o *TestPointShortApiResult) HasLastTestResultId() bool`

HasLastTestResultId returns a boolean if a field has been set.

### SetLastTestResultIdNil

`func (o *TestPointShortApiResult) SetLastTestResultIdNil(b bool)`

 SetLastTestResultIdNil sets the value for LastTestResultId to be an explicit nil

### UnsetLastTestResultId
`func (o *TestPointShortApiResult) UnsetLastTestResultId()`

UnsetLastTestResultId ensures that no value is present for LastTestResultId, not even an explicit nil
### GetIterationId

`func (o *TestPointShortApiResult) GetIterationId() string`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *TestPointShortApiResult) GetIterationIdOk() (*string, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *TestPointShortApiResult) SetIterationId(v string)`

SetIterationId sets IterationId field to given value.


### GetWorkItemMedianDuration

`func (o *TestPointShortApiResult) GetWorkItemMedianDuration() int64`

GetWorkItemMedianDuration returns the WorkItemMedianDuration field if non-nil, zero value otherwise.

### GetWorkItemMedianDurationOk

`func (o *TestPointShortApiResult) GetWorkItemMedianDurationOk() (*int64, bool)`

GetWorkItemMedianDurationOk returns a tuple with the WorkItemMedianDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemMedianDuration

`func (o *TestPointShortApiResult) SetWorkItemMedianDuration(v int64)`

SetWorkItemMedianDuration sets WorkItemMedianDuration field to given value.

### HasWorkItemMedianDuration

`func (o *TestPointShortApiResult) HasWorkItemMedianDuration() bool`

HasWorkItemMedianDuration returns a boolean if a field has been set.

### SetWorkItemMedianDurationNil

`func (o *TestPointShortApiResult) SetWorkItemMedianDurationNil(b bool)`

 SetWorkItemMedianDurationNil sets the value for WorkItemMedianDuration to be an explicit nil

### UnsetWorkItemMedianDuration
`func (o *TestPointShortApiResult) UnsetWorkItemMedianDuration()`

UnsetWorkItemMedianDuration ensures that no value is present for WorkItemMedianDuration, not even an explicit nil
### GetTestSuiteId

`func (o *TestPointShortApiResult) GetTestSuiteId() string`

GetTestSuiteId returns the TestSuiteId field if non-nil, zero value otherwise.

### GetTestSuiteIdOk

`func (o *TestPointShortApiResult) GetTestSuiteIdOk() (*string, bool)`

GetTestSuiteIdOk returns a tuple with the TestSuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuiteId

`func (o *TestPointShortApiResult) SetTestSuiteId(v string)`

SetTestSuiteId sets TestSuiteId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


