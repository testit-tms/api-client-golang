# WorkItemPreviewApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**Steps** | [**[]WorkItemPreviewStepApiModel**](WorkItemPreviewStepApiModel.md) |  | 
**Action** | **string** |  | 
**Expected** | **string** |  | 

## Methods

### NewWorkItemPreviewApiModel

`func NewWorkItemPreviewApiModel(name string, description string, steps []WorkItemPreviewStepApiModel, action string, expected string, ) *WorkItemPreviewApiModel`

NewWorkItemPreviewApiModel instantiates a new WorkItemPreviewApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemPreviewApiModelWithDefaults

`func NewWorkItemPreviewApiModelWithDefaults() *WorkItemPreviewApiModel`

NewWorkItemPreviewApiModelWithDefaults instantiates a new WorkItemPreviewApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkItemPreviewApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkItemPreviewApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkItemPreviewApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WorkItemPreviewApiModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkItemPreviewApiModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkItemPreviewApiModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSteps

`func (o *WorkItemPreviewApiModel) GetSteps() []WorkItemPreviewStepApiModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *WorkItemPreviewApiModel) GetStepsOk() (*[]WorkItemPreviewStepApiModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *WorkItemPreviewApiModel) SetSteps(v []WorkItemPreviewStepApiModel)`

SetSteps sets Steps field to given value.


### GetAction

`func (o *WorkItemPreviewApiModel) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WorkItemPreviewApiModel) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WorkItemPreviewApiModel) SetAction(v string)`

SetAction sets Action field to given value.


### GetExpected

`func (o *WorkItemPreviewApiModel) GetExpected() string`

GetExpected returns the Expected field if non-nil, zero value otherwise.

### GetExpectedOk

`func (o *WorkItemPreviewApiModel) GetExpectedOk() (*string, bool)`

GetExpectedOk returns a tuple with the Expected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpected

`func (o *WorkItemPreviewApiModel) SetExpected(v string)`

SetExpected sets Expected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


