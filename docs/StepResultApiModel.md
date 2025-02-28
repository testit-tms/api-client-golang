# StepResultApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StepId** | **string** |  | 
**Outcome** | **string** |  | 
**SharedStepVersionId** | Pointer to **NullableString** |  | [optional] 
**SharedStepResults** | Pointer to [**[]SharedStepResultApiModel**](SharedStepResultApiModel.md) |  | [optional] 
**Comment** | Pointer to [**NullableStepCommentApiModel**](StepCommentApiModel.md) |  | [optional] 

## Methods

### NewStepResultApiModel

`func NewStepResultApiModel(stepId string, outcome string, ) *StepResultApiModel`

NewStepResultApiModel instantiates a new StepResultApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStepResultApiModelWithDefaults

`func NewStepResultApiModelWithDefaults() *StepResultApiModel`

NewStepResultApiModelWithDefaults instantiates a new StepResultApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStepId

`func (o *StepResultApiModel) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *StepResultApiModel) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *StepResultApiModel) SetStepId(v string)`

SetStepId sets StepId field to given value.


### GetOutcome

`func (o *StepResultApiModel) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *StepResultApiModel) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *StepResultApiModel) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### GetSharedStepVersionId

`func (o *StepResultApiModel) GetSharedStepVersionId() string`

GetSharedStepVersionId returns the SharedStepVersionId field if non-nil, zero value otherwise.

### GetSharedStepVersionIdOk

`func (o *StepResultApiModel) GetSharedStepVersionIdOk() (*string, bool)`

GetSharedStepVersionIdOk returns a tuple with the SharedStepVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedStepVersionId

`func (o *StepResultApiModel) SetSharedStepVersionId(v string)`

SetSharedStepVersionId sets SharedStepVersionId field to given value.

### HasSharedStepVersionId

`func (o *StepResultApiModel) HasSharedStepVersionId() bool`

HasSharedStepVersionId returns a boolean if a field has been set.

### SetSharedStepVersionIdNil

`func (o *StepResultApiModel) SetSharedStepVersionIdNil(b bool)`

 SetSharedStepVersionIdNil sets the value for SharedStepVersionId to be an explicit nil

### UnsetSharedStepVersionId
`func (o *StepResultApiModel) UnsetSharedStepVersionId()`

UnsetSharedStepVersionId ensures that no value is present for SharedStepVersionId, not even an explicit nil
### GetSharedStepResults

`func (o *StepResultApiModel) GetSharedStepResults() []SharedStepResultApiModel`

GetSharedStepResults returns the SharedStepResults field if non-nil, zero value otherwise.

### GetSharedStepResultsOk

`func (o *StepResultApiModel) GetSharedStepResultsOk() (*[]SharedStepResultApiModel, bool)`

GetSharedStepResultsOk returns a tuple with the SharedStepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedStepResults

`func (o *StepResultApiModel) SetSharedStepResults(v []SharedStepResultApiModel)`

SetSharedStepResults sets SharedStepResults field to given value.

### HasSharedStepResults

`func (o *StepResultApiModel) HasSharedStepResults() bool`

HasSharedStepResults returns a boolean if a field has been set.

### SetSharedStepResultsNil

`func (o *StepResultApiModel) SetSharedStepResultsNil(b bool)`

 SetSharedStepResultsNil sets the value for SharedStepResults to be an explicit nil

### UnsetSharedStepResults
`func (o *StepResultApiModel) UnsetSharedStepResults()`

UnsetSharedStepResults ensures that no value is present for SharedStepResults, not even an explicit nil
### GetComment

`func (o *StepResultApiModel) GetComment() StepCommentApiModel`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *StepResultApiModel) GetCommentOk() (*StepCommentApiModel, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *StepResultApiModel) SetComment(v StepCommentApiModel)`

SetComment sets Comment field to given value.

### HasComment

`func (o *StepResultApiModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *StepResultApiModel) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *StepResultApiModel) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


