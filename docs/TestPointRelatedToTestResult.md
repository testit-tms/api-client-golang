# TestPointRelatedToTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestSuiteId** | **string** |  | 
**Id** | **string** |  | 
**TesterId** | Pointer to **NullableString** |  | [optional] 
**WorkItemId** | Pointer to **NullableString** |  | [optional] 
**ConfigurationId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** | Applies one of these values: Blocked, NoResults, Failed, Passed | [optional] 
**LastTestResultId** | Pointer to **NullableString** |  | [optional] 
**IterationId** | **string** |  | 

## Methods

### NewTestPointRelatedToTestResult

`func NewTestPointRelatedToTestResult(testSuiteId string, id string, iterationId string, ) *TestPointRelatedToTestResult`

NewTestPointRelatedToTestResult instantiates a new TestPointRelatedToTestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPointRelatedToTestResultWithDefaults

`func NewTestPointRelatedToTestResultWithDefaults() *TestPointRelatedToTestResult`

NewTestPointRelatedToTestResultWithDefaults instantiates a new TestPointRelatedToTestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestSuiteId

`func (o *TestPointRelatedToTestResult) GetTestSuiteId() string`

GetTestSuiteId returns the TestSuiteId field if non-nil, zero value otherwise.

### GetTestSuiteIdOk

`func (o *TestPointRelatedToTestResult) GetTestSuiteIdOk() (*string, bool)`

GetTestSuiteIdOk returns a tuple with the TestSuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuiteId

`func (o *TestPointRelatedToTestResult) SetTestSuiteId(v string)`

SetTestSuiteId sets TestSuiteId field to given value.


### GetId

`func (o *TestPointRelatedToTestResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestPointRelatedToTestResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestPointRelatedToTestResult) SetId(v string)`

SetId sets Id field to given value.


### GetTesterId

`func (o *TestPointRelatedToTestResult) GetTesterId() string`

GetTesterId returns the TesterId field if non-nil, zero value otherwise.

### GetTesterIdOk

`func (o *TestPointRelatedToTestResult) GetTesterIdOk() (*string, bool)`

GetTesterIdOk returns a tuple with the TesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesterId

`func (o *TestPointRelatedToTestResult) SetTesterId(v string)`

SetTesterId sets TesterId field to given value.

### HasTesterId

`func (o *TestPointRelatedToTestResult) HasTesterId() bool`

HasTesterId returns a boolean if a field has been set.

### SetTesterIdNil

`func (o *TestPointRelatedToTestResult) SetTesterIdNil(b bool)`

 SetTesterIdNil sets the value for TesterId to be an explicit nil

### UnsetTesterId
`func (o *TestPointRelatedToTestResult) UnsetTesterId()`

UnsetTesterId ensures that no value is present for TesterId, not even an explicit nil
### GetWorkItemId

`func (o *TestPointRelatedToTestResult) GetWorkItemId() string`

GetWorkItemId returns the WorkItemId field if non-nil, zero value otherwise.

### GetWorkItemIdOk

`func (o *TestPointRelatedToTestResult) GetWorkItemIdOk() (*string, bool)`

GetWorkItemIdOk returns a tuple with the WorkItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemId

`func (o *TestPointRelatedToTestResult) SetWorkItemId(v string)`

SetWorkItemId sets WorkItemId field to given value.

### HasWorkItemId

`func (o *TestPointRelatedToTestResult) HasWorkItemId() bool`

HasWorkItemId returns a boolean if a field has been set.

### SetWorkItemIdNil

`func (o *TestPointRelatedToTestResult) SetWorkItemIdNil(b bool)`

 SetWorkItemIdNil sets the value for WorkItemId to be an explicit nil

### UnsetWorkItemId
`func (o *TestPointRelatedToTestResult) UnsetWorkItemId()`

UnsetWorkItemId ensures that no value is present for WorkItemId, not even an explicit nil
### GetConfigurationId

`func (o *TestPointRelatedToTestResult) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *TestPointRelatedToTestResult) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *TestPointRelatedToTestResult) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *TestPointRelatedToTestResult) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### SetConfigurationIdNil

`func (o *TestPointRelatedToTestResult) SetConfigurationIdNil(b bool)`

 SetConfigurationIdNil sets the value for ConfigurationId to be an explicit nil

### UnsetConfigurationId
`func (o *TestPointRelatedToTestResult) UnsetConfigurationId()`

UnsetConfigurationId ensures that no value is present for ConfigurationId, not even an explicit nil
### GetStatus

`func (o *TestPointRelatedToTestResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestPointRelatedToTestResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestPointRelatedToTestResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestPointRelatedToTestResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TestPointRelatedToTestResult) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TestPointRelatedToTestResult) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLastTestResultId

`func (o *TestPointRelatedToTestResult) GetLastTestResultId() string`

GetLastTestResultId returns the LastTestResultId field if non-nil, zero value otherwise.

### GetLastTestResultIdOk

`func (o *TestPointRelatedToTestResult) GetLastTestResultIdOk() (*string, bool)`

GetLastTestResultIdOk returns a tuple with the LastTestResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResultId

`func (o *TestPointRelatedToTestResult) SetLastTestResultId(v string)`

SetLastTestResultId sets LastTestResultId field to given value.

### HasLastTestResultId

`func (o *TestPointRelatedToTestResult) HasLastTestResultId() bool`

HasLastTestResultId returns a boolean if a field has been set.

### SetLastTestResultIdNil

`func (o *TestPointRelatedToTestResult) SetLastTestResultIdNil(b bool)`

 SetLastTestResultIdNil sets the value for LastTestResultId to be an explicit nil

### UnsetLastTestResultId
`func (o *TestPointRelatedToTestResult) UnsetLastTestResultId()`

UnsetLastTestResultId ensures that no value is present for LastTestResultId, not even an explicit nil
### GetIterationId

`func (o *TestPointRelatedToTestResult) GetIterationId() string`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *TestPointRelatedToTestResult) GetIterationIdOk() (*string, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *TestPointRelatedToTestResult) SetIterationId(v string)`

SetIterationId sets IterationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


