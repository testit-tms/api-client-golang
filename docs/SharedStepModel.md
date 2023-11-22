# SharedStepModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **string** |  | 
**GlobalId** | **int64** |  | 
**Name** | **string** |  | 
**Steps** | [**[]StepModel**](StepModel.md) |  | 
**IsDeleted** | **bool** |  | 

## Methods

### NewSharedStepModel

`func NewSharedStepModel(versionId string, globalId int64, name string, steps []StepModel, isDeleted bool, ) *SharedStepModel`

NewSharedStepModel instantiates a new SharedStepModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedStepModelWithDefaults

`func NewSharedStepModelWithDefaults() *SharedStepModel`

NewSharedStepModelWithDefaults instantiates a new SharedStepModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *SharedStepModel) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *SharedStepModel) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *SharedStepModel) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetGlobalId

`func (o *SharedStepModel) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *SharedStepModel) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *SharedStepModel) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.


### GetName

`func (o *SharedStepModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedStepModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedStepModel) SetName(v string)`

SetName sets Name field to given value.


### GetSteps

`func (o *SharedStepModel) GetSteps() []StepModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *SharedStepModel) GetStepsOk() (*[]StepModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *SharedStepModel) SetSteps(v []StepModel)`

SetSteps sets Steps field to given value.


### GetIsDeleted

`func (o *SharedStepModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *SharedStepModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *SharedStepModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


