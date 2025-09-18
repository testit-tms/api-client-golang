# WorkflowShortApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**IsSystem** | **bool** |  | 
**IsDefault** | **bool** |  | 
**Projects** | [**WorkflowProjectApiResultApiCollectionPreview**](WorkflowProjectApiResultApiCollectionPreview.md) |  | 

## Methods

### NewWorkflowShortApiResult

`func NewWorkflowShortApiResult(id string, name string, isSystem bool, isDefault bool, projects WorkflowProjectApiResultApiCollectionPreview, ) *WorkflowShortApiResult`

NewWorkflowShortApiResult instantiates a new WorkflowShortApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowShortApiResultWithDefaults

`func NewWorkflowShortApiResultWithDefaults() *WorkflowShortApiResult`

NewWorkflowShortApiResultWithDefaults instantiates a new WorkflowShortApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowShortApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowShortApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowShortApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WorkflowShortApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowShortApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowShortApiResult) SetName(v string)`

SetName sets Name field to given value.


### GetIsSystem

`func (o *WorkflowShortApiResult) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *WorkflowShortApiResult) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *WorkflowShortApiResult) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.


### GetIsDefault

`func (o *WorkflowShortApiResult) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *WorkflowShortApiResult) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *WorkflowShortApiResult) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetProjects

`func (o *WorkflowShortApiResult) GetProjects() WorkflowProjectApiResultApiCollectionPreview`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *WorkflowShortApiResult) GetProjectsOk() (*WorkflowProjectApiResultApiCollectionPreview, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *WorkflowShortApiResult) SetProjects(v WorkflowProjectApiResultApiCollectionPreview)`

SetProjects sets Projects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


