# RerunsApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RerunCount** | **int32** |  | 
**RerunTestResults** | [**[]RerunTestResultApiResult**](RerunTestResultApiResult.md) |  | 

## Methods

### NewRerunsApiResult

`func NewRerunsApiResult(rerunCount int32, rerunTestResults []RerunTestResultApiResult, ) *RerunsApiResult`

NewRerunsApiResult instantiates a new RerunsApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRerunsApiResultWithDefaults

`func NewRerunsApiResultWithDefaults() *RerunsApiResult`

NewRerunsApiResultWithDefaults instantiates a new RerunsApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRerunCount

`func (o *RerunsApiResult) GetRerunCount() int32`

GetRerunCount returns the RerunCount field if non-nil, zero value otherwise.

### GetRerunCountOk

`func (o *RerunsApiResult) GetRerunCountOk() (*int32, bool)`

GetRerunCountOk returns a tuple with the RerunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerunCount

`func (o *RerunsApiResult) SetRerunCount(v int32)`

SetRerunCount sets RerunCount field to given value.


### GetRerunTestResults

`func (o *RerunsApiResult) GetRerunTestResults() []RerunTestResultApiResult`

GetRerunTestResults returns the RerunTestResults field if non-nil, zero value otherwise.

### GetRerunTestResultsOk

`func (o *RerunsApiResult) GetRerunTestResultsOk() (*[]RerunTestResultApiResult, bool)`

GetRerunTestResultsOk returns a tuple with the RerunTestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerunTestResults

`func (o *RerunsApiResult) SetRerunTestResults(v []RerunTestResultApiResult)`

SetRerunTestResults sets RerunTestResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


