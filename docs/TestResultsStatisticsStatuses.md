# TestResultsStatisticsStatuses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InProgress** | **int32** | Number of test results which is running currently | 
**Passed** | **int32** | Number of test results which successfully passed | 
**Failed** | **int32** | Number of test results which failed with an error | 
**Skipped** | **int32** | Number of test results which did not run and were skipped | 
**Blocked** | **int32** | Number of test results which cannot be launched | 

## Methods

### NewTestResultsStatisticsStatuses

`func NewTestResultsStatisticsStatuses(inProgress int32, passed int32, failed int32, skipped int32, blocked int32, ) *TestResultsStatisticsStatuses`

NewTestResultsStatisticsStatuses instantiates a new TestResultsStatisticsStatuses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultsStatisticsStatusesWithDefaults

`func NewTestResultsStatisticsStatusesWithDefaults() *TestResultsStatisticsStatuses`

NewTestResultsStatisticsStatusesWithDefaults instantiates a new TestResultsStatisticsStatuses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInProgress

`func (o *TestResultsStatisticsStatuses) GetInProgress() int32`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *TestResultsStatisticsStatuses) GetInProgressOk() (*int32, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *TestResultsStatisticsStatuses) SetInProgress(v int32)`

SetInProgress sets InProgress field to given value.


### GetPassed

`func (o *TestResultsStatisticsStatuses) GetPassed() int32`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *TestResultsStatisticsStatuses) GetPassedOk() (*int32, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *TestResultsStatisticsStatuses) SetPassed(v int32)`

SetPassed sets Passed field to given value.


### GetFailed

`func (o *TestResultsStatisticsStatuses) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *TestResultsStatisticsStatuses) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *TestResultsStatisticsStatuses) SetFailed(v int32)`

SetFailed sets Failed field to given value.


### GetSkipped

`func (o *TestResultsStatisticsStatuses) GetSkipped() int32`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *TestResultsStatisticsStatuses) GetSkippedOk() (*int32, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *TestResultsStatisticsStatuses) SetSkipped(v int32)`

SetSkipped sets Skipped field to given value.


### GetBlocked

`func (o *TestResultsStatisticsStatuses) GetBlocked() int32`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *TestResultsStatisticsStatuses) GetBlockedOk() (*int32, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *TestResultsStatisticsStatuses) SetBlocked(v int32)`

SetBlocked sets Blocked field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


