# TestStatusApiResultGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**GroupKey**](GroupKey.md) |  | 
**Data** | [**[]TestStatusApiResult**](TestStatusApiResult.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewTestStatusApiResultGroup

`func NewTestStatusApiResultGroup(key GroupKey, data []TestStatusApiResult, totalCount int32, ) *TestStatusApiResultGroup`

NewTestStatusApiResultGroup instantiates a new TestStatusApiResultGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestStatusApiResultGroupWithDefaults

`func NewTestStatusApiResultGroupWithDefaults() *TestStatusApiResultGroup`

NewTestStatusApiResultGroupWithDefaults instantiates a new TestStatusApiResultGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *TestStatusApiResultGroup) GetKey() GroupKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TestStatusApiResultGroup) GetKeyOk() (*GroupKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TestStatusApiResultGroup) SetKey(v GroupKey)`

SetKey sets Key field to given value.


### GetData

`func (o *TestStatusApiResultGroup) GetData() []TestStatusApiResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TestStatusApiResultGroup) GetDataOk() (*[]TestStatusApiResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TestStatusApiResultGroup) SetData(v []TestStatusApiResult)`

SetData sets Data field to given value.


### GetTotalCount

`func (o *TestStatusApiResultGroup) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TestStatusApiResultGroup) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TestStatusApiResultGroup) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


