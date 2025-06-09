# WorkflowApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**IsSystem** | **bool** |  | 
**IsDefault** | **bool** |  | 
**Statuses** | [**[]WorkflowStatusApiResult**](WorkflowStatusApiResult.md) |  | 

## Methods

### NewWorkflowApiResult

`func NewWorkflowApiResult(id string, name string, isSystem bool, isDefault bool, statuses []WorkflowStatusApiResult, ) *WorkflowApiResult`

NewWorkflowApiResult instantiates a new WorkflowApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowApiResultWithDefaults

`func NewWorkflowApiResultWithDefaults() *WorkflowApiResult`

NewWorkflowApiResultWithDefaults instantiates a new WorkflowApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WorkflowApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowApiResult) SetName(v string)`

SetName sets Name field to given value.


### GetIsSystem

`func (o *WorkflowApiResult) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *WorkflowApiResult) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *WorkflowApiResult) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.


### GetIsDefault

`func (o *WorkflowApiResult) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *WorkflowApiResult) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *WorkflowApiResult) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetStatuses

`func (o *WorkflowApiResult) GetStatuses() []WorkflowStatusApiResult`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *WorkflowApiResult) GetStatusesOk() (*[]WorkflowStatusApiResult, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *WorkflowApiResult) SetStatuses(v []WorkflowStatusApiResult)`

SetStatuses sets Statuses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


