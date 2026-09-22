# WorkflowShortApiResultGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**GroupKey**](GroupKey.md) |  | 
**Data** | [**[]WorkflowShortApiResult**](WorkflowShortApiResult.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewWorkflowShortApiResultGroup

`func NewWorkflowShortApiResultGroup(key GroupKey, data []WorkflowShortApiResult, totalCount int32, ) *WorkflowShortApiResultGroup`

NewWorkflowShortApiResultGroup instantiates a new WorkflowShortApiResultGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowShortApiResultGroupWithDefaults

`func NewWorkflowShortApiResultGroupWithDefaults() *WorkflowShortApiResultGroup`

NewWorkflowShortApiResultGroupWithDefaults instantiates a new WorkflowShortApiResultGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *WorkflowShortApiResultGroup) GetKey() GroupKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WorkflowShortApiResultGroup) GetKeyOk() (*GroupKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WorkflowShortApiResultGroup) SetKey(v GroupKey)`

SetKey sets Key field to given value.


### GetData

`func (o *WorkflowShortApiResultGroup) GetData() []WorkflowShortApiResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WorkflowShortApiResultGroup) GetDataOk() (*[]WorkflowShortApiResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WorkflowShortApiResultGroup) SetData(v []WorkflowShortApiResult)`

SetData sets Data field to given value.


### GetTotalCount

`func (o *WorkflowShortApiResultGroup) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *WorkflowShortApiResultGroup) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *WorkflowShortApiResultGroup) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


