# RerunTestResultApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Outcome** | **string** |  | 
**Status** | [**TestStatusApiResult**](TestStatusApiResult.md) |  | 
**RunNumber** | **int32** |  | 

## Methods

### NewRerunTestResultApiResult

`func NewRerunTestResultApiResult(id string, outcome string, status TestStatusApiResult, runNumber int32, ) *RerunTestResultApiResult`

NewRerunTestResultApiResult instantiates a new RerunTestResultApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRerunTestResultApiResultWithDefaults

`func NewRerunTestResultApiResultWithDefaults() *RerunTestResultApiResult`

NewRerunTestResultApiResultWithDefaults instantiates a new RerunTestResultApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RerunTestResultApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RerunTestResultApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RerunTestResultApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetOutcome

`func (o *RerunTestResultApiResult) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *RerunTestResultApiResult) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *RerunTestResultApiResult) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### GetStatus

`func (o *RerunTestResultApiResult) GetStatus() TestStatusApiResult`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RerunTestResultApiResult) GetStatusOk() (*TestStatusApiResult, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RerunTestResultApiResult) SetStatus(v TestStatusApiResult)`

SetStatus sets Status field to given value.


### GetRunNumber

`func (o *RerunTestResultApiResult) GetRunNumber() int32`

GetRunNumber returns the RunNumber field if non-nil, zero value otherwise.

### GetRunNumberOk

`func (o *RerunTestResultApiResult) GetRunNumberOk() (*int32, bool)`

GetRunNumberOk returns a tuple with the RunNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunNumber

`func (o *RerunTestResultApiResult) SetRunNumber(v int32)`

SetRunNumber sets RunNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


