# TestResultsStatisticsGetModelStatuses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InProgress** | **int32** | Number of test results which is running currently | 
**Passed** | **int32** | Number of test results which successfully passed | 
**Failed** | **int32** | Number of test results which failed with an error | 
**Skipped** | **int32** | Number of test results which did not run and were skipped | 
**Blocked** | **int32** | Number of test results which cannot be launched | 

## Methods

### NewTestResultsStatisticsGetModelStatuses

`func NewTestResultsStatisticsGetModelStatuses(inProgress int32, passed int32, failed int32, skipped int32, blocked int32, ) *TestResultsStatisticsGetModelStatuses`

NewTestResultsStatisticsGetModelStatuses instantiates a new TestResultsStatisticsGetModelStatuses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultsStatisticsGetModelStatusesWithDefaults

`func NewTestResultsStatisticsGetModelStatusesWithDefaults() *TestResultsStatisticsGetModelStatuses`

NewTestResultsStatisticsGetModelStatusesWithDefaults instantiates a new TestResultsStatisticsGetModelStatuses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInProgress

`func (o *TestResultsStatisticsGetModelStatuses) GetInProgress() int32`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *TestResultsStatisticsGetModelStatuses) GetInProgressOk() (*int32, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *TestResultsStatisticsGetModelStatuses) SetInProgress(v int32)`

SetInProgress sets InProgress field to given value.


### GetPassed

`func (o *TestResultsStatisticsGetModelStatuses) GetPassed() int32`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *TestResultsStatisticsGetModelStatuses) GetPassedOk() (*int32, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *TestResultsStatisticsGetModelStatuses) SetPassed(v int32)`

SetPassed sets Passed field to given value.


### GetFailed

`func (o *TestResultsStatisticsGetModelStatuses) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *TestResultsStatisticsGetModelStatuses) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *TestResultsStatisticsGetModelStatuses) SetFailed(v int32)`

SetFailed sets Failed field to given value.


### GetSkipped

`func (o *TestResultsStatisticsGetModelStatuses) GetSkipped() int32`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *TestResultsStatisticsGetModelStatuses) GetSkippedOk() (*int32, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *TestResultsStatisticsGetModelStatuses) SetSkipped(v int32)`

SetSkipped sets Skipped field to given value.


### GetBlocked

`func (o *TestResultsStatisticsGetModelStatuses) GetBlocked() int32`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *TestResultsStatisticsGetModelStatuses) GetBlockedOk() (*int32, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *TestResultsStatisticsGetModelStatuses) SetBlocked(v int32)`

SetBlocked sets Blocked field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


