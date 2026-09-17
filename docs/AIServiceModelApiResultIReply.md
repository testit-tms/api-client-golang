# AIServiceModelApiResultIReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | [readonly] 
**Data** | [**[]AIServiceModelApiResult**](AIServiceModelApiResult.md) |  | 
**Groups** | [**[]AIServiceModelApiResultGroup**](AIServiceModelApiResultGroup.md) |  | 

## Methods

### NewAIServiceModelApiResultIReply

`func NewAIServiceModelApiResultIReply(totalCount int32, data []AIServiceModelApiResult, groups []AIServiceModelApiResultGroup, ) *AIServiceModelApiResultIReply`

NewAIServiceModelApiResultIReply instantiates a new AIServiceModelApiResultIReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIServiceModelApiResultIReplyWithDefaults

`func NewAIServiceModelApiResultIReplyWithDefaults() *AIServiceModelApiResultIReply`

NewAIServiceModelApiResultIReplyWithDefaults instantiates a new AIServiceModelApiResultIReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *AIServiceModelApiResultIReply) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AIServiceModelApiResultIReply) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AIServiceModelApiResultIReply) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetData

`func (o *AIServiceModelApiResultIReply) GetData() []AIServiceModelApiResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AIServiceModelApiResultIReply) GetDataOk() (*[]AIServiceModelApiResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AIServiceModelApiResultIReply) SetData(v []AIServiceModelApiResult)`

SetData sets Data field to given value.


### GetGroups

`func (o *AIServiceModelApiResultIReply) GetGroups() []AIServiceModelApiResultGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AIServiceModelApiResultIReply) GetGroupsOk() (*[]AIServiceModelApiResultGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AIServiceModelApiResultIReply) SetGroups(v []AIServiceModelApiResultGroup)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


