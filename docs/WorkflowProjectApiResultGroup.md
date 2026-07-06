# WorkflowProjectApiResultGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**GroupKey**](GroupKey.md) |  | 
**Data** | [**[]WorkflowProjectApiResult**](WorkflowProjectApiResult.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewWorkflowProjectApiResultGroup

`func NewWorkflowProjectApiResultGroup(key GroupKey, data []WorkflowProjectApiResult, totalCount int32, ) *WorkflowProjectApiResultGroup`

NewWorkflowProjectApiResultGroup instantiates a new WorkflowProjectApiResultGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowProjectApiResultGroupWithDefaults

`func NewWorkflowProjectApiResultGroupWithDefaults() *WorkflowProjectApiResultGroup`

NewWorkflowProjectApiResultGroupWithDefaults instantiates a new WorkflowProjectApiResultGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *WorkflowProjectApiResultGroup) GetKey() GroupKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WorkflowProjectApiResultGroup) GetKeyOk() (*GroupKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WorkflowProjectApiResultGroup) SetKey(v GroupKey)`

SetKey sets Key field to given value.


### GetData

`func (o *WorkflowProjectApiResultGroup) GetData() []WorkflowProjectApiResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WorkflowProjectApiResultGroup) GetDataOk() (*[]WorkflowProjectApiResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WorkflowProjectApiResultGroup) SetData(v []WorkflowProjectApiResult)`

SetData sets Data field to given value.


### GetTotalCount

`func (o *WorkflowProjectApiResultGroup) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *WorkflowProjectApiResultGroup) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *WorkflowProjectApiResultGroup) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


