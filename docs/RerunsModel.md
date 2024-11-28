# RerunsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RerunCount** | **int32** |  | 
**RerunTestResults** | [**[]RerunTestResultModel**](RerunTestResultModel.md) |  | 

## Methods

### NewRerunsModel

`func NewRerunsModel(rerunCount int32, rerunTestResults []RerunTestResultModel, ) *RerunsModel`

NewRerunsModel instantiates a new RerunsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRerunsModelWithDefaults

`func NewRerunsModelWithDefaults() *RerunsModel`

NewRerunsModelWithDefaults instantiates a new RerunsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRerunCount

`func (o *RerunsModel) GetRerunCount() int32`

GetRerunCount returns the RerunCount field if non-nil, zero value otherwise.

### GetRerunCountOk

`func (o *RerunsModel) GetRerunCountOk() (*int32, bool)`

GetRerunCountOk returns a tuple with the RerunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerunCount

`func (o *RerunsModel) SetRerunCount(v int32)`

SetRerunCount sets RerunCount field to given value.


### GetRerunTestResults

`func (o *RerunsModel) GetRerunTestResults() []RerunTestResultModel`

GetRerunTestResults returns the RerunTestResults field if non-nil, zero value otherwise.

### GetRerunTestResultsOk

`func (o *RerunsModel) GetRerunTestResultsOk() (*[]RerunTestResultModel, bool)`

GetRerunTestResultsOk returns a tuple with the RerunTestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerunTestResults

`func (o *RerunsModel) SetRerunTestResults(v []RerunTestResultModel)`

SetRerunTestResults sets RerunTestResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


