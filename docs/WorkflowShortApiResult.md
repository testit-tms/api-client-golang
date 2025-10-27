# WorkflowShortApiResult

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
**Projects** | [**WorkflowProjectApiResultApiCollectionPreview**](WorkflowProjectApiResultApiCollectionPreview.md) |  | 

## Methods

### NewWorkflowShortApiResult

`func NewWorkflowShortApiResult(id string, name string, isSystem bool, isDefault bool, createdDate time.Time, createdById string, modifiedDate time.Time, modifiedById string, projects WorkflowProjectApiResultApiCollectionPreview, ) *WorkflowShortApiResult`

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


### GetCreatedDate

`func (o *WorkflowShortApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *WorkflowShortApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *WorkflowShortApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetCreatedById

`func (o *WorkflowShortApiResult) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WorkflowShortApiResult) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WorkflowShortApiResult) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedDate

`func (o *WorkflowShortApiResult) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *WorkflowShortApiResult) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *WorkflowShortApiResult) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.


### GetModifiedById

`func (o *WorkflowShortApiResult) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *WorkflowShortApiResult) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *WorkflowShortApiResult) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.


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


