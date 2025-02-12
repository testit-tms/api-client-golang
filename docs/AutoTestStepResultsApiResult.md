# AutoTestStepResultsApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Info** | Pointer to **NullableString** |  | [optional] 
**StartedOn** | Pointer to **NullableTime** |  | [optional] 
**CompletedOn** | Pointer to **NullableTime** |  | [optional] 
**Duration** | Pointer to **NullableInt64** |  | [optional] 
**Outcome** | Pointer to [**NullableAutoTestOutcome**](AutoTestOutcome.md) |  | [optional] 
**StepResults** | Pointer to [**[]AutoTestStepResultsApiResult**](AutoTestStepResultsApiResult.md) |  | [optional] 
**Attachments** | Pointer to [**[]AttachmentApiResult**](AttachmentApiResult.md) |  | [optional] 
**Parameters** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAutoTestStepResultsApiResult

`func NewAutoTestStepResultsApiResult() *AutoTestStepResultsApiResult`

NewAutoTestStepResultsApiResult instantiates a new AutoTestStepResultsApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestStepResultsApiResultWithDefaults

`func NewAutoTestStepResultsApiResultWithDefaults() *AutoTestStepResultsApiResult`

NewAutoTestStepResultsApiResultWithDefaults instantiates a new AutoTestStepResultsApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AutoTestStepResultsApiResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AutoTestStepResultsApiResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AutoTestStepResultsApiResult) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AutoTestStepResultsApiResult) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AutoTestStepResultsApiResult) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AutoTestStepResultsApiResult) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *AutoTestStepResultsApiResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutoTestStepResultsApiResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutoTestStepResultsApiResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutoTestStepResultsApiResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AutoTestStepResultsApiResult) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AutoTestStepResultsApiResult) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInfo

`func (o *AutoTestStepResultsApiResult) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AutoTestStepResultsApiResult) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AutoTestStepResultsApiResult) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AutoTestStepResultsApiResult) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### SetInfoNil

`func (o *AutoTestStepResultsApiResult) SetInfoNil(b bool)`

 SetInfoNil sets the value for Info to be an explicit nil

### UnsetInfo
`func (o *AutoTestStepResultsApiResult) UnsetInfo()`

UnsetInfo ensures that no value is present for Info, not even an explicit nil
### GetStartedOn

`func (o *AutoTestStepResultsApiResult) GetStartedOn() time.Time`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *AutoTestStepResultsApiResult) GetStartedOnOk() (*time.Time, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *AutoTestStepResultsApiResult) SetStartedOn(v time.Time)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *AutoTestStepResultsApiResult) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *AutoTestStepResultsApiResult) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *AutoTestStepResultsApiResult) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetCompletedOn

`func (o *AutoTestStepResultsApiResult) GetCompletedOn() time.Time`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *AutoTestStepResultsApiResult) GetCompletedOnOk() (*time.Time, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *AutoTestStepResultsApiResult) SetCompletedOn(v time.Time)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *AutoTestStepResultsApiResult) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### SetCompletedOnNil

`func (o *AutoTestStepResultsApiResult) SetCompletedOnNil(b bool)`

 SetCompletedOnNil sets the value for CompletedOn to be an explicit nil

### UnsetCompletedOn
`func (o *AutoTestStepResultsApiResult) UnsetCompletedOn()`

UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
### GetDuration

`func (o *AutoTestStepResultsApiResult) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AutoTestStepResultsApiResult) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AutoTestStepResultsApiResult) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AutoTestStepResultsApiResult) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *AutoTestStepResultsApiResult) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *AutoTestStepResultsApiResult) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetOutcome

`func (o *AutoTestStepResultsApiResult) GetOutcome() AutoTestOutcome`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *AutoTestStepResultsApiResult) GetOutcomeOk() (*AutoTestOutcome, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *AutoTestStepResultsApiResult) SetOutcome(v AutoTestOutcome)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *AutoTestStepResultsApiResult) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *AutoTestStepResultsApiResult) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *AutoTestStepResultsApiResult) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetStepResults

`func (o *AutoTestStepResultsApiResult) GetStepResults() []AutoTestStepResultsApiResult`

GetStepResults returns the StepResults field if non-nil, zero value otherwise.

### GetStepResultsOk

`func (o *AutoTestStepResultsApiResult) GetStepResultsOk() (*[]AutoTestStepResultsApiResult, bool)`

GetStepResultsOk returns a tuple with the StepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepResults

`func (o *AutoTestStepResultsApiResult) SetStepResults(v []AutoTestStepResultsApiResult)`

SetStepResults sets StepResults field to given value.

### HasStepResults

`func (o *AutoTestStepResultsApiResult) HasStepResults() bool`

HasStepResults returns a boolean if a field has been set.

### SetStepResultsNil

`func (o *AutoTestStepResultsApiResult) SetStepResultsNil(b bool)`

 SetStepResultsNil sets the value for StepResults to be an explicit nil

### UnsetStepResults
`func (o *AutoTestStepResultsApiResult) UnsetStepResults()`

UnsetStepResults ensures that no value is present for StepResults, not even an explicit nil
### GetAttachments

`func (o *AutoTestStepResultsApiResult) GetAttachments() []AttachmentApiResult`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *AutoTestStepResultsApiResult) GetAttachmentsOk() (*[]AttachmentApiResult, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *AutoTestStepResultsApiResult) SetAttachments(v []AttachmentApiResult)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *AutoTestStepResultsApiResult) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *AutoTestStepResultsApiResult) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *AutoTestStepResultsApiResult) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetParameters

`func (o *AutoTestStepResultsApiResult) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AutoTestStepResultsApiResult) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AutoTestStepResultsApiResult) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AutoTestStepResultsApiResult) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *AutoTestStepResultsApiResult) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *AutoTestStepResultsApiResult) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


