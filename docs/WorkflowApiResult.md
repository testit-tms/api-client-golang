# WorkflowApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**IsSystem** | **bool** |  | 
**IsDefault** | **bool** |  | 
**CreatedDate** | **time.Time** |  | 
**CreatedById** | **string** |  | 
**ModifiedDate** | **time.Time** |  | 
**ModifiedById** | **string** |  | 
**Statuses** | [**[]WorkflowStatusApiResult**](WorkflowStatusApiResult.md) |  | 
**Projects** | [**[]WorkflowProjectApiResult**](WorkflowProjectApiResult.md) |  | 

## Methods

### NewWorkflowApiResult

`func NewWorkflowApiResult(id string, name string, isSystem bool, isDefault bool, createdDate time.Time, createdById string, modifiedDate time.Time, modifiedById string, statuses []WorkflowStatusApiResult, projects []WorkflowProjectApiResult, ) *WorkflowApiResult`

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


### GetCreatedDate

`func (o *WorkflowApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *WorkflowApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *WorkflowApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetCreatedById

`func (o *WorkflowApiResult) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WorkflowApiResult) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WorkflowApiResult) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedDate

`func (o *WorkflowApiResult) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *WorkflowApiResult) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *WorkflowApiResult) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.


### GetModifiedById

`func (o *WorkflowApiResult) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *WorkflowApiResult) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *WorkflowApiResult) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.


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


### GetProjects

`func (o *WorkflowApiResult) GetProjects() []WorkflowProjectApiResult`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *WorkflowApiResult) GetProjectsOk() (*[]WorkflowProjectApiResult, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *WorkflowApiResult) SetProjects(v []WorkflowProjectApiResult)`

SetProjects sets Projects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


