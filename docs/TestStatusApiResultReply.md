# TestStatusApiResultReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TestStatusApiResult**](TestStatusApiResult.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewTestStatusApiResultReply

`func NewTestStatusApiResultReply(data []TestStatusApiResult, totalCount int32, ) *TestStatusApiResultReply`

NewTestStatusApiResultReply instantiates a new TestStatusApiResultReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestStatusApiResultReplyWithDefaults

`func NewTestStatusApiResultReplyWithDefaults() *TestStatusApiResultReply`

NewTestStatusApiResultReplyWithDefaults instantiates a new TestStatusApiResultReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TestStatusApiResultReply) GetData() []TestStatusApiResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TestStatusApiResultReply) GetDataOk() (*[]TestStatusApiResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TestStatusApiResultReply) SetData(v []TestStatusApiResult)`

SetData sets Data field to given value.


### GetTotalCount

`func (o *TestStatusApiResultReply) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TestStatusApiResultReply) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TestStatusApiResultReply) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


