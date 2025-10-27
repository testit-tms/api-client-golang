# AutoTestStepResultUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** | The name of the step. | [optional] 
**Description** | Pointer to **NullableString** | Description of the step result. | [optional] 
**Info** | Pointer to **NullableString** | Extended description of the step result. | [optional] 
**StartedOn** | Pointer to **NullableTime** | Step start date. | [optional] 
**CompletedOn** | Pointer to **NullableTime** | Step end date. | [optional] 
**Duration** | Pointer to **NullableInt64** | Expected or actual duration of the test run execution in milliseconds. | [optional] 
**Outcome** | Pointer to [**NullableAvailableTestResultOutcome**](AvailableTestResultOutcome.md) | Specifies the result of the autotest execution. | [optional] 
**StepResults** | Pointer to [**[]AutoTestStepResultUpdateRequest**](AutoTestStepResultUpdateRequest.md) | Nested step results. The maximum nesting level is 15. | [optional] 
**Attachments** | Pointer to [**[]AttachmentUpdateRequest**](AttachmentUpdateRequest.md) | /// &lt;summary&gt;  Specifies an attachment GUID. Multiple values can be sent.  &lt;/summary&gt; | [optional] 
**Parameters** | Pointer to **map[string]string** | \&quot;&lt;b&gt;parameter&lt;/b&gt;\&quot;: \&quot;&lt;b&gt;value&lt;/b&gt;\&quot; pair with arbitrary custom parameters. Multiple parameters can be sent. | [optional] 

## Methods

### NewAutoTestStepResultUpdateRequest

`func NewAutoTestStepResultUpdateRequest() *AutoTestStepResultUpdateRequest`

NewAutoTestStepResultUpdateRequest instantiates a new AutoTestStepResultUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestStepResultUpdateRequestWithDefaults

`func NewAutoTestStepResultUpdateRequestWithDefaults() *AutoTestStepResultUpdateRequest`

NewAutoTestStepResultUpdateRequestWithDefaults instantiates a new AutoTestStepResultUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AutoTestStepResultUpdateRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AutoTestStepResultUpdateRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AutoTestStepResultUpdateRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AutoTestStepResultUpdateRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AutoTestStepResultUpdateRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AutoTestStepResultUpdateRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *AutoTestStepResultUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutoTestStepResultUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutoTestStepResultUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutoTestStepResultUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AutoTestStepResultUpdateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AutoTestStepResultUpdateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInfo

`func (o *AutoTestStepResultUpdateRequest) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AutoTestStepResultUpdateRequest) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AutoTestStepResultUpdateRequest) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AutoTestStepResultUpdateRequest) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### SetInfoNil

`func (o *AutoTestStepResultUpdateRequest) SetInfoNil(b bool)`

 SetInfoNil sets the value for Info to be an explicit nil

### UnsetInfo
`func (o *AutoTestStepResultUpdateRequest) UnsetInfo()`

UnsetInfo ensures that no value is present for Info, not even an explicit nil
### GetStartedOn

`func (o *AutoTestStepResultUpdateRequest) GetStartedOn() time.Time`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *AutoTestStepResultUpdateRequest) GetStartedOnOk() (*time.Time, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *AutoTestStepResultUpdateRequest) SetStartedOn(v time.Time)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *AutoTestStepResultUpdateRequest) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *AutoTestStepResultUpdateRequest) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *AutoTestStepResultUpdateRequest) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetCompletedOn

`func (o *AutoTestStepResultUpdateRequest) GetCompletedOn() time.Time`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *AutoTestStepResultUpdateRequest) GetCompletedOnOk() (*time.Time, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *AutoTestStepResultUpdateRequest) SetCompletedOn(v time.Time)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *AutoTestStepResultUpdateRequest) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### SetCompletedOnNil

`func (o *AutoTestStepResultUpdateRequest) SetCompletedOnNil(b bool)`

 SetCompletedOnNil sets the value for CompletedOn to be an explicit nil

### UnsetCompletedOn
`func (o *AutoTestStepResultUpdateRequest) UnsetCompletedOn()`

UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
### GetDuration

`func (o *AutoTestStepResultUpdateRequest) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AutoTestStepResultUpdateRequest) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AutoTestStepResultUpdateRequest) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AutoTestStepResultUpdateRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *AutoTestStepResultUpdateRequest) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *AutoTestStepResultUpdateRequest) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetOutcome

`func (o *AutoTestStepResultUpdateRequest) GetOutcome() AvailableTestResultOutcome`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *AutoTestStepResultUpdateRequest) GetOutcomeOk() (*AvailableTestResultOutcome, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *AutoTestStepResultUpdateRequest) SetOutcome(v AvailableTestResultOutcome)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *AutoTestStepResultUpdateRequest) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *AutoTestStepResultUpdateRequest) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *AutoTestStepResultUpdateRequest) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetStepResults

`func (o *AutoTestStepResultUpdateRequest) GetStepResults() []AutoTestStepResultUpdateRequest`

GetStepResults returns the StepResults field if non-nil, zero value otherwise.

### GetStepResultsOk

`func (o *AutoTestStepResultUpdateRequest) GetStepResultsOk() (*[]AutoTestStepResultUpdateRequest, bool)`

GetStepResultsOk returns a tuple with the StepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepResults

`func (o *AutoTestStepResultUpdateRequest) SetStepResults(v []AutoTestStepResultUpdateRequest)`

SetStepResults sets StepResults field to given value.

### HasStepResults

`func (o *AutoTestStepResultUpdateRequest) HasStepResults() bool`

HasStepResults returns a boolean if a field has been set.

### SetStepResultsNil

`func (o *AutoTestStepResultUpdateRequest) SetStepResultsNil(b bool)`

 SetStepResultsNil sets the value for StepResults to be an explicit nil

### UnsetStepResults
`func (o *AutoTestStepResultUpdateRequest) UnsetStepResults()`

UnsetStepResults ensures that no value is present for StepResults, not even an explicit nil
### GetAttachments

`func (o *AutoTestStepResultUpdateRequest) GetAttachments() []AttachmentUpdateRequest`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *AutoTestStepResultUpdateRequest) GetAttachmentsOk() (*[]AttachmentUpdateRequest, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *AutoTestStepResultUpdateRequest) SetAttachments(v []AttachmentUpdateRequest)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *AutoTestStepResultUpdateRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *AutoTestStepResultUpdateRequest) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *AutoTestStepResultUpdateRequest) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetParameters

`func (o *AutoTestStepResultUpdateRequest) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AutoTestStepResultUpdateRequest) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AutoTestStepResultUpdateRequest) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AutoTestStepResultUpdateRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *AutoTestStepResultUpdateRequest) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *AutoTestStepResultUpdateRequest) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


