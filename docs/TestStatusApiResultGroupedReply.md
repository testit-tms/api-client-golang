# TestStatusApiResultGroupedReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | [**[]TestStatusApiResultGroup**](TestStatusApiResultGroup.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewTestStatusApiResultGroupedReply

`func NewTestStatusApiResultGroupedReply(groups []TestStatusApiResultGroup, totalCount int32, ) *TestStatusApiResultGroupedReply`

NewTestStatusApiResultGroupedReply instantiates a new TestStatusApiResultGroupedReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestStatusApiResultGroupedReplyWithDefaults

`func NewTestStatusApiResultGroupedReplyWithDefaults() *TestStatusApiResultGroupedReply`

NewTestStatusApiResultGroupedReplyWithDefaults instantiates a new TestStatusApiResultGroupedReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *TestStatusApiResultGroupedReply) GetGroups() []TestStatusApiResultGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *TestStatusApiResultGroupedReply) GetGroupsOk() (*[]TestStatusApiResultGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *TestStatusApiResultGroupedReply) SetGroups(v []TestStatusApiResultGroup)`

SetGroups sets Groups field to given value.


### GetTotalCount

`func (o *TestStatusApiResultGroupedReply) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TestStatusApiResultGroupedReply) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TestStatusApiResultGroupedReply) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


