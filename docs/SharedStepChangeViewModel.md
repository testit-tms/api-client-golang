# SharedStepChangeViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**GlobalId** | **int64** |  | 
**Name** | **string** |  | 
**Steps** | [**[]WorkItemStepChangeViewModel**](WorkItemStepChangeViewModel.md) |  | 

## Methods

### NewSharedStepChangeViewModel

`func NewSharedStepChangeViewModel(id string, globalId int64, name string, steps []WorkItemStepChangeViewModel, ) *SharedStepChangeViewModel`

NewSharedStepChangeViewModel instantiates a new SharedStepChangeViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedStepChangeViewModelWithDefaults

`func NewSharedStepChangeViewModelWithDefaults() *SharedStepChangeViewModel`

NewSharedStepChangeViewModelWithDefaults instantiates a new SharedStepChangeViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SharedStepChangeViewModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharedStepChangeViewModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharedStepChangeViewModel) SetId(v string)`

SetId sets Id field to given value.


### GetGlobalId

`func (o *SharedStepChangeViewModel) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *SharedStepChangeViewModel) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *SharedStepChangeViewModel) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.


### GetName

`func (o *SharedStepChangeViewModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedStepChangeViewModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedStepChangeViewModel) SetName(v string)`

SetName sets Name field to given value.


### GetSteps

`func (o *SharedStepChangeViewModel) GetSteps() []WorkItemStepChangeViewModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *SharedStepChangeViewModel) GetStepsOk() (*[]WorkItemStepChangeViewModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *SharedStepChangeViewModel) SetSteps(v []WorkItemStepChangeViewModel)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


