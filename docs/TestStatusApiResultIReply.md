# TestStatusApiResultIReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | [readonly] 
**Data** | [**[]TestStatusApiResult**](TestStatusApiResult.md) |  | 
**Groups** | [**[]TestStatusApiResultGroup**](TestStatusApiResultGroup.md) |  | 

## Methods

### NewTestStatusApiResultIReply

`func NewTestStatusApiResultIReply(totalCount int32, data []TestStatusApiResult, groups []TestStatusApiResultGroup, ) *TestStatusApiResultIReply`

NewTestStatusApiResultIReply instantiates a new TestStatusApiResultIReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestStatusApiResultIReplyWithDefaults

`func NewTestStatusApiResultIReplyWithDefaults() *TestStatusApiResultIReply`

NewTestStatusApiResultIReplyWithDefaults instantiates a new TestStatusApiResultIReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *TestStatusApiResultIReply) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TestStatusApiResultIReply) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TestStatusApiResultIReply) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetData

`func (o *TestStatusApiResultIReply) GetData() []TestStatusApiResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TestStatusApiResultIReply) GetDataOk() (*[]TestStatusApiResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TestStatusApiResultIReply) SetData(v []TestStatusApiResult)`

SetData sets Data field to given value.


### GetGroups

`func (o *TestStatusApiResultIReply) GetGroups() []TestStatusApiResultGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *TestStatusApiResultIReply) GetGroupsOk() (*[]TestStatusApiResultGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *TestStatusApiResultIReply) SetGroups(v []TestStatusApiResultGroup)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


