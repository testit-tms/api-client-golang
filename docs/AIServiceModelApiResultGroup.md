# AIServiceModelApiResultGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**GroupKey**](GroupKey.md) |  | 
**Data** | [**[]AIServiceModelApiResult**](AIServiceModelApiResult.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewAIServiceModelApiResultGroup

`func NewAIServiceModelApiResultGroup(key GroupKey, data []AIServiceModelApiResult, totalCount int32, ) *AIServiceModelApiResultGroup`

NewAIServiceModelApiResultGroup instantiates a new AIServiceModelApiResultGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIServiceModelApiResultGroupWithDefaults

`func NewAIServiceModelApiResultGroupWithDefaults() *AIServiceModelApiResultGroup`

NewAIServiceModelApiResultGroupWithDefaults instantiates a new AIServiceModelApiResultGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AIServiceModelApiResultGroup) GetKey() GroupKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AIServiceModelApiResultGroup) GetKeyOk() (*GroupKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AIServiceModelApiResultGroup) SetKey(v GroupKey)`

SetKey sets Key field to given value.


### GetData

`func (o *AIServiceModelApiResultGroup) GetData() []AIServiceModelApiResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AIServiceModelApiResultGroup) GetDataOk() (*[]AIServiceModelApiResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AIServiceModelApiResultGroup) SetData(v []AIServiceModelApiResult)`

SetData sets Data field to given value.


### GetTotalCount

`func (o *AIServiceModelApiResultGroup) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AIServiceModelApiResultGroup) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AIServiceModelApiResultGroup) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


