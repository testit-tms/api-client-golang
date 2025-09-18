# UpdateWorkflowApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**IsDefault** | **bool** |  | 
**Statuses** | [**[]WorkflowStatusApiModel**](WorkflowStatusApiModel.md) |  | 
**ProjectIds** | **[]string** |  | 

## Methods

### NewUpdateWorkflowApiModel

`func NewUpdateWorkflowApiModel(name string, isDefault bool, statuses []WorkflowStatusApiModel, projectIds []string, ) *UpdateWorkflowApiModel`

NewUpdateWorkflowApiModel instantiates a new UpdateWorkflowApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWorkflowApiModelWithDefaults

`func NewUpdateWorkflowApiModelWithDefaults() *UpdateWorkflowApiModel`

NewUpdateWorkflowApiModelWithDefaults instantiates a new UpdateWorkflowApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateWorkflowApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWorkflowApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWorkflowApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetIsDefault

`func (o *UpdateWorkflowApiModel) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *UpdateWorkflowApiModel) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *UpdateWorkflowApiModel) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetStatuses

`func (o *UpdateWorkflowApiModel) GetStatuses() []WorkflowStatusApiModel`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *UpdateWorkflowApiModel) GetStatusesOk() (*[]WorkflowStatusApiModel, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *UpdateWorkflowApiModel) SetStatuses(v []WorkflowStatusApiModel)`

SetStatuses sets Statuses field to given value.


### GetProjectIds

`func (o *UpdateWorkflowApiModel) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *UpdateWorkflowApiModel) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *UpdateWorkflowApiModel) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


