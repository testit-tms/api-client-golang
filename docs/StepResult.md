# StepResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StepId** | **string** |  | 
**Outcome** | **string** |  | 
**SharedStepVersionId** | Pointer to **NullableString** |  | [optional] 
**SharedStepResults** | Pointer to [**[]SharedStepResult**](SharedStepResult.md) |  | [optional] 
**Comment** | Pointer to [**NullableStepComment**](StepComment.md) |  | [optional] 

## Methods

### NewStepResult

`func NewStepResult(stepId string, outcome string, ) *StepResult`

NewStepResult instantiates a new StepResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStepResultWithDefaults

`func NewStepResultWithDefaults() *StepResult`

NewStepResultWithDefaults instantiates a new StepResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStepId

`func (o *StepResult) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *StepResult) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *StepResult) SetStepId(v string)`

SetStepId sets StepId field to given value.


### GetOutcome

`func (o *StepResult) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *StepResult) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *StepResult) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### GetSharedStepVersionId

`func (o *StepResult) GetSharedStepVersionId() string`

GetSharedStepVersionId returns the SharedStepVersionId field if non-nil, zero value otherwise.

### GetSharedStepVersionIdOk

`func (o *StepResult) GetSharedStepVersionIdOk() (*string, bool)`

GetSharedStepVersionIdOk returns a tuple with the SharedStepVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedStepVersionId

`func (o *StepResult) SetSharedStepVersionId(v string)`

SetSharedStepVersionId sets SharedStepVersionId field to given value.

### HasSharedStepVersionId

`func (o *StepResult) HasSharedStepVersionId() bool`

HasSharedStepVersionId returns a boolean if a field has been set.

### SetSharedStepVersionIdNil

`func (o *StepResult) SetSharedStepVersionIdNil(b bool)`

 SetSharedStepVersionIdNil sets the value for SharedStepVersionId to be an explicit nil

### UnsetSharedStepVersionId
`func (o *StepResult) UnsetSharedStepVersionId()`

UnsetSharedStepVersionId ensures that no value is present for SharedStepVersionId, not even an explicit nil
### GetSharedStepResults

`func (o *StepResult) GetSharedStepResults() []SharedStepResult`

GetSharedStepResults returns the SharedStepResults field if non-nil, zero value otherwise.

### GetSharedStepResultsOk

`func (o *StepResult) GetSharedStepResultsOk() (*[]SharedStepResult, bool)`

GetSharedStepResultsOk returns a tuple with the SharedStepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedStepResults

`func (o *StepResult) SetSharedStepResults(v []SharedStepResult)`

SetSharedStepResults sets SharedStepResults field to given value.

### HasSharedStepResults

`func (o *StepResult) HasSharedStepResults() bool`

HasSharedStepResults returns a boolean if a field has been set.

### SetSharedStepResultsNil

`func (o *StepResult) SetSharedStepResultsNil(b bool)`

 SetSharedStepResultsNil sets the value for SharedStepResults to be an explicit nil

### UnsetSharedStepResults
`func (o *StepResult) UnsetSharedStepResults()`

UnsetSharedStepResults ensures that no value is present for SharedStepResults, not even an explicit nil
### GetComment

`func (o *StepResult) GetComment() StepComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *StepResult) GetCommentOk() (*StepComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *StepResult) SetComment(v StepComment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *StepResult) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *StepResult) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *StepResult) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


