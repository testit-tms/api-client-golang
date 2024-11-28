# TestPointByTestSuiteModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Test point unique internal identifier | 
**TesterId** | Pointer to **NullableString** | Tester who is responded for the test unique internal identifier | [optional] 
**WorkItemId** | Pointer to **NullableString** | Workitem to which test point relates unique identifier | [optional] 
**ConfigurationId** | Pointer to **NullableString** | Configuration to which test point relates unique identifier | [optional] 
**Status** | Pointer to **NullableString** | Test point status   Applies one of these values: Blocked, NoResults, Failed, Passed | [optional] 
**LastTestResultId** | Pointer to **NullableString** | Last test result unique identifier | [optional] 
**IterationId** | **string** | Iteration unique identifier | 
**WorkItemMedianDuration** | Pointer to **NullableInt64** | Median duration of work item the test point represents | [optional] 

## Methods

### NewTestPointByTestSuiteModel

`func NewTestPointByTestSuiteModel(id string, iterationId string, ) *TestPointByTestSuiteModel`

NewTestPointByTestSuiteModel instantiates a new TestPointByTestSuiteModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPointByTestSuiteModelWithDefaults

`func NewTestPointByTestSuiteModelWithDefaults() *TestPointByTestSuiteModel`

NewTestPointByTestSuiteModelWithDefaults instantiates a new TestPointByTestSuiteModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestPointByTestSuiteModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestPointByTestSuiteModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestPointByTestSuiteModel) SetId(v string)`

SetId sets Id field to given value.


### GetTesterId

`func (o *TestPointByTestSuiteModel) GetTesterId() string`

GetTesterId returns the TesterId field if non-nil, zero value otherwise.

### GetTesterIdOk

`func (o *TestPointByTestSuiteModel) GetTesterIdOk() (*string, bool)`

GetTesterIdOk returns a tuple with the TesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesterId

`func (o *TestPointByTestSuiteModel) SetTesterId(v string)`

SetTesterId sets TesterId field to given value.

### HasTesterId

`func (o *TestPointByTestSuiteModel) HasTesterId() bool`

HasTesterId returns a boolean if a field has been set.

### SetTesterIdNil

`func (o *TestPointByTestSuiteModel) SetTesterIdNil(b bool)`

 SetTesterIdNil sets the value for TesterId to be an explicit nil

### UnsetTesterId
`func (o *TestPointByTestSuiteModel) UnsetTesterId()`

UnsetTesterId ensures that no value is present for TesterId, not even an explicit nil
### GetWorkItemId

`func (o *TestPointByTestSuiteModel) GetWorkItemId() string`

GetWorkItemId returns the WorkItemId field if non-nil, zero value otherwise.

### GetWorkItemIdOk

`func (o *TestPointByTestSuiteModel) GetWorkItemIdOk() (*string, bool)`

GetWorkItemIdOk returns a tuple with the WorkItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemId

`func (o *TestPointByTestSuiteModel) SetWorkItemId(v string)`

SetWorkItemId sets WorkItemId field to given value.

### HasWorkItemId

`func (o *TestPointByTestSuiteModel) HasWorkItemId() bool`

HasWorkItemId returns a boolean if a field has been set.

### SetWorkItemIdNil

`func (o *TestPointByTestSuiteModel) SetWorkItemIdNil(b bool)`

 SetWorkItemIdNil sets the value for WorkItemId to be an explicit nil

### UnsetWorkItemId
`func (o *TestPointByTestSuiteModel) UnsetWorkItemId()`

UnsetWorkItemId ensures that no value is present for WorkItemId, not even an explicit nil
### GetConfigurationId

`func (o *TestPointByTestSuiteModel) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *TestPointByTestSuiteModel) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *TestPointByTestSuiteModel) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *TestPointByTestSuiteModel) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### SetConfigurationIdNil

`func (o *TestPointByTestSuiteModel) SetConfigurationIdNil(b bool)`

 SetConfigurationIdNil sets the value for ConfigurationId to be an explicit nil

### UnsetConfigurationId
`func (o *TestPointByTestSuiteModel) UnsetConfigurationId()`

UnsetConfigurationId ensures that no value is present for ConfigurationId, not even an explicit nil
### GetStatus

`func (o *TestPointByTestSuiteModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestPointByTestSuiteModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestPointByTestSuiteModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestPointByTestSuiteModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TestPointByTestSuiteModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TestPointByTestSuiteModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLastTestResultId

`func (o *TestPointByTestSuiteModel) GetLastTestResultId() string`

GetLastTestResultId returns the LastTestResultId field if non-nil, zero value otherwise.

### GetLastTestResultIdOk

`func (o *TestPointByTestSuiteModel) GetLastTestResultIdOk() (*string, bool)`

GetLastTestResultIdOk returns a tuple with the LastTestResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultId

`func (o *TestPointByTestSuiteModel) SetLastTestResultId(v string)`

SetLastTestResultId sets LastTestResultId field to given value.

### HasLastTestResultId

`func (o *TestPointByTestSuiteModel) HasLastTestResultId() bool`

HasLastTestResultId returns a boolean if a field has been set.

### SetLastTestResultIdNil

`func (o *TestPointByTestSuiteModel) SetLastTestResultIdNil(b bool)`

 SetLastTestResultIdNil sets the value for LastTestResultId to be an explicit nil

### UnsetLastTestResultId
`func (o *TestPointByTestSuiteModel) UnsetLastTestResultId()`

UnsetLastTestResultId ensures that no value is present for LastTestResultId, not even an explicit nil
### GetIterationId

`func (o *TestPointByTestSuiteModel) GetIterationId() string`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *TestPointByTestSuiteModel) GetIterationIdOk() (*string, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *TestPointByTestSuiteModel) SetIterationId(v string)`

SetIterationId sets IterationId field to given value.


### GetWorkItemMedianDuration

`func (o *TestPointByTestSuiteModel) GetWorkItemMedianDuration() int64`

GetWorkItemMedianDuration returns the WorkItemMedianDuration field if non-nil, zero value otherwise.

### GetWorkItemMedianDurationOk

`func (o *TestPointByTestSuiteModel) GetWorkItemMedianDurationOk() (*int64, bool)`

GetWorkItemMedianDurationOk returns a tuple with the WorkItemMedianDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemMedianDuration

`func (o *TestPointByTestSuiteModel) SetWorkItemMedianDuration(v int64)`

SetWorkItemMedianDuration sets WorkItemMedianDuration field to given value.

### HasWorkItemMedianDuration

`func (o *TestPointByTestSuiteModel) HasWorkItemMedianDuration() bool`

HasWorkItemMedianDuration returns a boolean if a field has been set.

### SetWorkItemMedianDurationNil

`func (o *TestPointByTestSuiteModel) SetWorkItemMedianDurationNil(b bool)`

 SetWorkItemMedianDurationNil sets the value for WorkItemMedianDuration to be an explicit nil

### UnsetWorkItemMedianDuration
`func (o *TestPointByTestSuiteModel) UnsetWorkItemMedianDuration()`

UnsetWorkItemMedianDuration ensures that no value is present for WorkItemMedianDuration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


