# WorkflowStatusApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Code** | **string** |  | 
**Type** | [**TestStatusApiType**](TestStatusApiType.md) | Collection of possible status types | 
**Description** | **NullableString** |  | 
**IsSystem** | **bool** |  | 
**Priority** | **int32** |  | 

## Methods

### NewWorkflowStatusApiResult

`func NewWorkflowStatusApiResult(id string, name string, code string, type_ TestStatusApiType, description NullableString, isSystem bool, priority int32, ) *WorkflowStatusApiResult`

NewWorkflowStatusApiResult instantiates a new WorkflowStatusApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStatusApiResultWithDefaults

`func NewWorkflowStatusApiResultWithDefaults() *WorkflowStatusApiResult`

NewWorkflowStatusApiResultWithDefaults instantiates a new WorkflowStatusApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowStatusApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowStatusApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowStatusApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WorkflowStatusApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowStatusApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowStatusApiResult) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *WorkflowStatusApiResult) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *WorkflowStatusApiResult) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *WorkflowStatusApiResult) SetCode(v string)`

SetCode sets Code field to given value.


### GetType

`func (o *WorkflowStatusApiResult) GetType() TestStatusApiType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowStatusApiResult) GetTypeOk() (*TestStatusApiType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowStatusApiResult) SetType(v TestStatusApiType)`

SetType sets Type field to given value.


### GetDescription

`func (o *WorkflowStatusApiResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowStatusApiResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowStatusApiResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *WorkflowStatusApiResult) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkflowStatusApiResult) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsSystem

`func (o *WorkflowStatusApiResult) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *WorkflowStatusApiResult) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *WorkflowStatusApiResult) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.


### GetPriority

`func (o *WorkflowStatusApiResult) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WorkflowStatusApiResult) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WorkflowStatusApiResult) SetPriority(v int32)`

SetPriority sets Priority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


