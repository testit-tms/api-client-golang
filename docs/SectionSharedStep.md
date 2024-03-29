# SectionSharedStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **string** |  | 
**GlobalId** | **int64** |  | 
**Name** | **string** |  | 
**Steps** | [**[]StepModel**](StepModel.md) |  | 
**IsDeleted** | **bool** |  | 

## Methods

### NewSectionSharedStep

`func NewSectionSharedStep(versionId string, globalId int64, name string, steps []StepModel, isDeleted bool, ) *SectionSharedStep`

NewSectionSharedStep instantiates a new SectionSharedStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionSharedStepWithDefaults

`func NewSectionSharedStepWithDefaults() *SectionSharedStep`

NewSectionSharedStepWithDefaults instantiates a new SectionSharedStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *SectionSharedStep) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *SectionSharedStep) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *SectionSharedStep) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetGlobalId

`func (o *SectionSharedStep) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *SectionSharedStep) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *SectionSharedStep) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.


### GetName

`func (o *SectionSharedStep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SectionSharedStep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SectionSharedStep) SetName(v string)`

SetName sets Name field to given value.


### GetSteps

`func (o *SectionSharedStep) GetSteps() []StepModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *SectionSharedStep) GetStepsOk() (*[]StepModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *SectionSharedStep) SetSteps(v []StepModel)`

SetSteps sets Steps field to given value.


### GetIsDeleted

`func (o *SectionSharedStep) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *SectionSharedStep) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *SectionSharedStep) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


