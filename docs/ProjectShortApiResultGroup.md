# ProjectShortApiResultGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**GroupKey**](GroupKey.md) |  | 
**Data** | [**[]ProjectShortApiResult**](ProjectShortApiResult.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewProjectShortApiResultGroup

`func NewProjectShortApiResultGroup(key GroupKey, data []ProjectShortApiResult, totalCount int32, ) *ProjectShortApiResultGroup`

NewProjectShortApiResultGroup instantiates a new ProjectShortApiResultGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectShortApiResultGroupWithDefaults

`func NewProjectShortApiResultGroupWithDefaults() *ProjectShortApiResultGroup`

NewProjectShortApiResultGroupWithDefaults instantiates a new ProjectShortApiResultGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ProjectShortApiResultGroup) GetKey() GroupKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProjectShortApiResultGroup) GetKeyOk() (*GroupKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProjectShortApiResultGroup) SetKey(v GroupKey)`

SetKey sets Key field to given value.


### GetData

`func (o *ProjectShortApiResultGroup) GetData() []ProjectShortApiResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProjectShortApiResultGroup) GetDataOk() (*[]ProjectShortApiResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProjectShortApiResultGroup) SetData(v []ProjectShortApiResult)`

SetData sets Data field to given value.


### GetTotalCount

`func (o *ProjectShortApiResultGroup) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ProjectShortApiResultGroup) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ProjectShortApiResultGroup) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


