# WorkItemStepChangeViewModelWorkItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**GlobalId** | **int64** |  | 
**Name** | **string** |  | 
**Steps** | [**[]WorkItemStepChangeViewModel**](WorkItemStepChangeViewModel.md) |  | 

## Methods

### NewWorkItemStepChangeViewModelWorkItem

`func NewWorkItemStepChangeViewModelWorkItem(id string, globalId int64, name string, steps []WorkItemStepChangeViewModel, ) *WorkItemStepChangeViewModelWorkItem`

NewWorkItemStepChangeViewModelWorkItem instantiates a new WorkItemStepChangeViewModelWorkItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemStepChangeViewModelWorkItemWithDefaults

`func NewWorkItemStepChangeViewModelWorkItemWithDefaults() *WorkItemStepChangeViewModelWorkItem`

NewWorkItemStepChangeViewModelWorkItemWithDefaults instantiates a new WorkItemStepChangeViewModelWorkItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkItemStepChangeViewModelWorkItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkItemStepChangeViewModelWorkItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkItemStepChangeViewModelWorkItem) SetId(v string)`

SetId sets Id field to given value.


### GetGlobalId

`func (o *WorkItemStepChangeViewModelWorkItem) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *WorkItemStepChangeViewModelWorkItem) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *WorkItemStepChangeViewModelWorkItem) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.


### GetName

`func (o *WorkItemStepChangeViewModelWorkItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkItemStepChangeViewModelWorkItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkItemStepChangeViewModelWorkItem) SetName(v string)`

SetName sets Name field to given value.


### GetSteps

`func (o *WorkItemStepChangeViewModelWorkItem) GetSteps() []WorkItemStepChangeViewModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *WorkItemStepChangeViewModelWorkItem) GetStepsOk() (*[]WorkItemStepChangeViewModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *WorkItemStepChangeViewModelWorkItem) SetSteps(v []WorkItemStepChangeViewModel)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


