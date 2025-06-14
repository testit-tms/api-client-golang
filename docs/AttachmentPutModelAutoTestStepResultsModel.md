# AttachmentPutModelAutoTestStepResultsModel

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
**StepResults** | Pointer to [**[]AttachmentPutModelAutoTestStepResultsModel**](AttachmentPutModelAutoTestStepResultsModel.md) | Nested step results. The maximum nesting level is 15. | [optional] 
**Attachments** | Pointer to [**[]AttachmentPutModel**](AttachmentPutModel.md) | /// &lt;summary&gt; Specifies an attachment GUID. Multiple values can be sent. &lt;/summary&gt; | [optional] 
**Parameters** | Pointer to **map[string]string** | \&quot;&lt;b&gt;parameter&lt;/b&gt;\&quot;: \&quot;&lt;b&gt;value&lt;/b&gt;\&quot; pair with arbitrary custom parameters. Multiple parameters can be sent. | [optional] 

## Methods

### NewAttachmentPutModelAutoTestStepResultsModel

`func NewAttachmentPutModelAutoTestStepResultsModel() *AttachmentPutModelAutoTestStepResultsModel`

NewAttachmentPutModelAutoTestStepResultsModel instantiates a new AttachmentPutModelAutoTestStepResultsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentPutModelAutoTestStepResultsModelWithDefaults

`func NewAttachmentPutModelAutoTestStepResultsModelWithDefaults() *AttachmentPutModelAutoTestStepResultsModel`

NewAttachmentPutModelAutoTestStepResultsModelWithDefaults instantiates a new AttachmentPutModelAutoTestStepResultsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AttachmentPutModelAutoTestStepResultsModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AttachmentPutModelAutoTestStepResultsModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AttachmentPutModelAutoTestStepResultsModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AttachmentPutModelAutoTestStepResultsModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AttachmentPutModelAutoTestStepResultsModel) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AttachmentPutModelAutoTestStepResultsModel) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *AttachmentPutModelAutoTestStepResultsModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AttachmentPutModelAutoTestStepResultsModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AttachmentPutModelAutoTestStepResultsModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AttachmentPutModelAutoTestStepResultsModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AttachmentPutModelAutoTestStepResultsModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AttachmentPutModelAutoTestStepResultsModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInfo

`func (o *AttachmentPutModelAutoTestStepResultsModel) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AttachmentPutModelAutoTestStepResultsModel) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AttachmentPutModelAutoTestStepResultsModel) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AttachmentPutModelAutoTestStepResultsModel) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### SetInfoNil

`func (o *AttachmentPutModelAutoTestStepResultsModel) SetInfoNil(b bool)`

 SetInfoNil sets the value for Info to be an explicit nil

### UnsetInfo
`func (o *AttachmentPutModelAutoTestStepResultsModel) UnsetInfo()`

UnsetInfo ensures that no value is present for Info, not even an explicit nil
### GetStartedOn

`func (o *AttachmentPutModelAutoTestStepResultsModel) GetStartedOn() time.Time`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *AttachmentPutModelAutoTestStepResultsModel) GetStartedOnOk() (*time.Time, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *AttachmentPutModelAutoTestStepResultsModel) SetStartedOn(v time.Time)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *AttachmentPutModelAutoTestStepResultsModel) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *AttachmentPutModelAutoTestStepResultsModel) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *AttachmentPutModelAutoTestStepResultsModel) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetCompletedOn

`func (o *AttachmentPutModelAutoTestStepResultsModel) GetCompletedOn() time.Time`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *AttachmentPutModelAutoTestStepResultsModel) GetCompletedOnOk() (*time.Time, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *AttachmentPutModelAutoTestStepResultsModel) SetCompletedOn(v time.Time)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *AttachmentPutModelAutoTestStepResultsModel) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### SetCompletedOnNil

`func (o *AttachmentPutModelAutoTestStepResultsModel) SetCompletedOnNil(b bool)`

 SetCompletedOnNil sets the value for CompletedOn to be an explicit nil

### UnsetCompletedOn
`func (o *AttachmentPutModelAutoTestStepResultsModel) UnsetCompletedOn()`

UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
### GetDuration

`func (o *AttachmentPutModelAutoTestStepResultsModel) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AttachmentPutModelAutoTestStepResultsModel) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AttachmentPutModelAutoTestStepResultsModel) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AttachmentPutModelAutoTestStepResultsModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *AttachmentPutModelAutoTestStepResultsModel) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *AttachmentPutModelAutoTestStepResultsModel) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetOutcome

`func (o *AttachmentPutModelAutoTestStepResultsModel) GetOutcome() AvailableTestResultOutcome`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *AttachmentPutModelAutoTestStepResultsModel) GetOutcomeOk() (*AvailableTestResultOutcome, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *AttachmentPutModelAutoTestStepResultsModel) SetOutcome(v AvailableTestResultOutcome)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *AttachmentPutModelAutoTestStepResultsModel) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *AttachmentPutModelAutoTestStepResultsModel) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *AttachmentPutModelAutoTestStepResultsModel) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetStepResults

`func (o *AttachmentPutModelAutoTestStepResultsModel) GetStepResults() []AttachmentPutModelAutoTestStepResultsModel`

GetStepResults returns the StepResults field if non-nil, zero value otherwise.

### GetStepResultsOk

`func (o *AttachmentPutModelAutoTestStepResultsModel) GetStepResultsOk() (*[]AttachmentPutModelAutoTestStepResultsModel, bool)`

GetStepResultsOk returns a tuple with the StepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepResults

`func (o *AttachmentPutModelAutoTestStepResultsModel) SetStepResults(v []AttachmentPutModelAutoTestStepResultsModel)`

SetStepResults sets StepResults field to given value.

### HasStepResults

`func (o *AttachmentPutModelAutoTestStepResultsModel) HasStepResults() bool`

HasStepResults returns a boolean if a field has been set.

### SetStepResultsNil

`func (o *AttachmentPutModelAutoTestStepResultsModel) SetStepResultsNil(b bool)`

 SetStepResultsNil sets the value for StepResults to be an explicit nil

### UnsetStepResults
`func (o *AttachmentPutModelAutoTestStepResultsModel) UnsetStepResults()`

UnsetStepResults ensures that no value is present for StepResults, not even an explicit nil
### GetAttachments

`func (o *AttachmentPutModelAutoTestStepResultsModel) GetAttachments() []AttachmentPutModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *AttachmentPutModelAutoTestStepResultsModel) GetAttachmentsOk() (*[]AttachmentPutModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *AttachmentPutModelAutoTestStepResultsModel) SetAttachments(v []AttachmentPutModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *AttachmentPutModelAutoTestStepResultsModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *AttachmentPutModelAutoTestStepResultsModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *AttachmentPutModelAutoTestStepResultsModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetParameters

`func (o *AttachmentPutModelAutoTestStepResultsModel) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AttachmentPutModelAutoTestStepResultsModel) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AttachmentPutModelAutoTestStepResultsModel) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AttachmentPutModelAutoTestStepResultsModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *AttachmentPutModelAutoTestStepResultsModel) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *AttachmentPutModelAutoTestStepResultsModel) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


