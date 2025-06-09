# CreateWorkflowApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**Statuses** | [**[]WorkflowStatusApiModel**](WorkflowStatusApiModel.md) |  | 

## Methods

### NewCreateWorkflowApiModel

`func NewCreateWorkflowApiModel(name string, statuses []WorkflowStatusApiModel, ) *CreateWorkflowApiModel`

NewCreateWorkflowApiModel instantiates a new CreateWorkflowApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkflowApiModelWithDefaults

`func NewCreateWorkflowApiModelWithDefaults() *CreateWorkflowApiModel`

NewCreateWorkflowApiModelWithDefaults instantiates a new CreateWorkflowApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateWorkflowApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWorkflowApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWorkflowApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetIsDefault

`func (o *CreateWorkflowApiModel) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CreateWorkflowApiModel) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CreateWorkflowApiModel) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CreateWorkflowApiModel) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *CreateWorkflowApiModel) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *CreateWorkflowApiModel) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetStatuses

`func (o *CreateWorkflowApiModel) GetStatuses() []WorkflowStatusApiModel`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *CreateWorkflowApiModel) GetStatusesOk() (*[]WorkflowStatusApiModel, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *CreateWorkflowApiModel) SetStatuses(v []WorkflowStatusApiModel)`

SetStatuses sets Statuses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


