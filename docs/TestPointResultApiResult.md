# TestPointResultApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestPointId** | Pointer to **NullableString** |  | [optional] 
**AggregatedOutcome** | Pointer to **NullableString** |  | [optional] 
**AggregatedStatus** | Pointer to [**NullableTestStatusApiResult**](TestStatusApiResult.md) |  | [optional] 
**WorkItemGlobalId** | Pointer to **NullableInt64** |  | [optional] 
**WorkItemName** | Pointer to **NullableString** |  | [optional] 
**ConfigurationName** | Pointer to **NullableString** |  | [optional] 
**TestResults** | [**[]TestResultShortApiResult**](TestResultShortApiResult.md) |  | 

## Methods

### NewTestPointResultApiResult

`func NewTestPointResultApiResult(testResults []TestResultShortApiResult, ) *TestPointResultApiResult`

NewTestPointResultApiResult instantiates a new TestPointResultApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPointResultApiResultWithDefaults

`func NewTestPointResultApiResultWithDefaults() *TestPointResultApiResult`

NewTestPointResultApiResultWithDefaults instantiates a new TestPointResultApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestPointId

`func (o *TestPointResultApiResult) GetTestPointId() string`

GetTestPointId returns the TestPointId field if non-nil, zero value otherwise.

### GetTestPointIdOk

`func (o *TestPointResultApiResult) GetTestPointIdOk() (*string, bool)`

GetTestPointIdOk returns a tuple with the TestPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPointId

`func (o *TestPointResultApiResult) SetTestPointId(v string)`

SetTestPointId sets TestPointId field to given value.

### HasTestPointId

`func (o *TestPointResultApiResult) HasTestPointId() bool`

HasTestPointId returns a boolean if a field has been set.

### SetTestPointIdNil

`func (o *TestPointResultApiResult) SetTestPointIdNil(b bool)`

 SetTestPointIdNil sets the value for TestPointId to be an explicit nil

### UnsetTestPointId
`func (o *TestPointResultApiResult) UnsetTestPointId()`

UnsetTestPointId ensures that no value is present for TestPointId, not even an explicit nil
### GetAggregatedOutcome

`func (o *TestPointResultApiResult) GetAggregatedOutcome() string`

GetAggregatedOutcome returns the AggregatedOutcome field if non-nil, zero value otherwise.

### GetAggregatedOutcomeOk

`func (o *TestPointResultApiResult) GetAggregatedOutcomeOk() (*string, bool)`

GetAggregatedOutcomeOk returns a tuple with the AggregatedOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedOutcome

`func (o *TestPointResultApiResult) SetAggregatedOutcome(v string)`

SetAggregatedOutcome sets AggregatedOutcome field to given value.

### HasAggregatedOutcome

`func (o *TestPointResultApiResult) HasAggregatedOutcome() bool`

HasAggregatedOutcome returns a boolean if a field has been set.

### SetAggregatedOutcomeNil

`func (o *TestPointResultApiResult) SetAggregatedOutcomeNil(b bool)`

 SetAggregatedOutcomeNil sets the value for AggregatedOutcome to be an explicit nil

### UnsetAggregatedOutcome
`func (o *TestPointResultApiResult) UnsetAggregatedOutcome()`

UnsetAggregatedOutcome ensures that no value is present for AggregatedOutcome, not even an explicit nil
### GetAggregatedStatus

`func (o *TestPointResultApiResult) GetAggregatedStatus() TestStatusApiResult`

GetAggregatedStatus returns the AggregatedStatus field if non-nil, zero value otherwise.

### GetAggregatedStatusOk

`func (o *TestPointResultApiResult) GetAggregatedStatusOk() (*TestStatusApiResult, bool)`

GetAggregatedStatusOk returns a tuple with the AggregatedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedStatus

`func (o *TestPointResultApiResult) SetAggregatedStatus(v TestStatusApiResult)`

SetAggregatedStatus sets AggregatedStatus field to given value.

### HasAggregatedStatus

`func (o *TestPointResultApiResult) HasAggregatedStatus() bool`

HasAggregatedStatus returns a boolean if a field has been set.

### SetAggregatedStatusNil

`func (o *TestPointResultApiResult) SetAggregatedStatusNil(b bool)`

 SetAggregatedStatusNil sets the value for AggregatedStatus to be an explicit nil

### UnsetAggregatedStatus
`func (o *TestPointResultApiResult) UnsetAggregatedStatus()`

UnsetAggregatedStatus ensures that no value is present for AggregatedStatus, not even an explicit nil
### GetWorkItemGlobalId

`func (o *TestPointResultApiResult) GetWorkItemGlobalId() int64`

GetWorkItemGlobalId returns the WorkItemGlobalId field if non-nil, zero value otherwise.

### GetWorkItemGlobalIdOk

`func (o *TestPointResultApiResult) GetWorkItemGlobalIdOk() (*int64, bool)`

GetWorkItemGlobalIdOk returns a tuple with the WorkItemGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemGlobalId

`func (o *TestPointResultApiResult) SetWorkItemGlobalId(v int64)`

SetWorkItemGlobalId sets WorkItemGlobalId field to given value.

### HasWorkItemGlobalId

`func (o *TestPointResultApiResult) HasWorkItemGlobalId() bool`

HasWorkItemGlobalId returns a boolean if a field has been set.

### SetWorkItemGlobalIdNil

`func (o *TestPointResultApiResult) SetWorkItemGlobalIdNil(b bool)`

 SetWorkItemGlobalIdNil sets the value for WorkItemGlobalId to be an explicit nil

### UnsetWorkItemGlobalId
`func (o *TestPointResultApiResult) UnsetWorkItemGlobalId()`

UnsetWorkItemGlobalId ensures that no value is present for WorkItemGlobalId, not even an explicit nil
### GetWorkItemName

`func (o *TestPointResultApiResult) GetWorkItemName() string`

GetWorkItemName returns the WorkItemName field if non-nil, zero value otherwise.

### GetWorkItemNameOk

`func (o *TestPointResultApiResult) GetWorkItemNameOk() (*string, bool)`

GetWorkItemNameOk returns a tuple with the WorkItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemName

`func (o *TestPointResultApiResult) SetWorkItemName(v string)`

SetWorkItemName sets WorkItemName field to given value.

### HasWorkItemName

`func (o *TestPointResultApiResult) HasWorkItemName() bool`

HasWorkItemName returns a boolean if a field has been set.

### SetWorkItemNameNil

`func (o *TestPointResultApiResult) SetWorkItemNameNil(b bool)`

 SetWorkItemNameNil sets the value for WorkItemName to be an explicit nil

### UnsetWorkItemName
`func (o *TestPointResultApiResult) UnsetWorkItemName()`

UnsetWorkItemName ensures that no value is present for WorkItemName, not even an explicit nil
### GetConfigurationName

`func (o *TestPointResultApiResult) GetConfigurationName() string`

GetConfigurationName returns the ConfigurationName field if non-nil, zero value otherwise.

### GetConfigurationNameOk

`func (o *TestPointResultApiResult) GetConfigurationNameOk() (*string, bool)`

GetConfigurationNameOk returns a tuple with the ConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationName

`func (o *TestPointResultApiResult) SetConfigurationName(v string)`

SetConfigurationName sets ConfigurationName field to given value.

### HasConfigurationName

`func (o *TestPointResultApiResult) HasConfigurationName() bool`

HasConfigurationName returns a boolean if a field has been set.

### SetConfigurationNameNil

`func (o *TestPointResultApiResult) SetConfigurationNameNil(b bool)`

 SetConfigurationNameNil sets the value for ConfigurationName to be an explicit nil

### UnsetConfigurationName
`func (o *TestPointResultApiResult) UnsetConfigurationName()`

UnsetConfigurationName ensures that no value is present for ConfigurationName, not even an explicit nil
### GetTestResults

`func (o *TestPointResultApiResult) GetTestResults() []TestResultShortApiResult`

GetTestResults returns the TestResults field if non-nil, zero value otherwise.

### GetTestResultsOk

`func (o *TestPointResultApiResult) GetTestResultsOk() (*[]TestResultShortApiResult, bool)`

GetTestResultsOk returns a tuple with the TestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResults

`func (o *TestPointResultApiResult) SetTestResults(v []TestResultShortApiResult)`

SetTestResults sets TestResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


