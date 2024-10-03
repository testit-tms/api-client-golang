# ApiV2TestResultsIdPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureClassIds** | Pointer to **[]string** |  | [optional] 
**Outcome** | Pointer to [**NullableTestResultOutcome**](TestResultOutcome.md) |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**StepResults** | Pointer to [**[]StepResultModel**](StepResultModel.md) |  | [optional] 
**Attachments** | Pointer to [**[]AttachmentPutModel**](AttachmentPutModel.md) |  | [optional] 
**DurationInMs** | Pointer to **NullableInt64** |  | [optional] 
**Duration** | Pointer to **NullableInt64** |  | [optional] 
**StepComments** | Pointer to [**[]TestResultStepCommentPutModel**](TestResultStepCommentPutModel.md) |  | [optional] 
**SetupResults** | Pointer to [**[]AttachmentPutModelAutoTestStepResultsModel**](AttachmentPutModelAutoTestStepResultsModel.md) |  | [optional] 
**TeardownResults** | Pointer to [**[]AttachmentPutModelAutoTestStepResultsModel**](AttachmentPutModelAutoTestStepResultsModel.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Trace** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewApiV2TestResultsIdPutRequest

`func NewApiV2TestResultsIdPutRequest() *ApiV2TestResultsIdPutRequest`

NewApiV2TestResultsIdPutRequest instantiates a new ApiV2TestResultsIdPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2TestResultsIdPutRequestWithDefaults

`func NewApiV2TestResultsIdPutRequestWithDefaults() *ApiV2TestResultsIdPutRequest`

NewApiV2TestResultsIdPutRequestWithDefaults instantiates a new ApiV2TestResultsIdPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureClassIds

`func (o *ApiV2TestResultsIdPutRequest) GetFailureClassIds() []string`

GetFailureClassIds returns the FailureClassIds field if non-nil, zero value otherwise.

### GetFailureClassIdsOk

`func (o *ApiV2TestResultsIdPutRequest) GetFailureClassIdsOk() (*[]string, bool)`

GetFailureClassIdsOk returns a tuple with the FailureClassIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureClassIds

`func (o *ApiV2TestResultsIdPutRequest) SetFailureClassIds(v []string)`

SetFailureClassIds sets FailureClassIds field to given value.

### HasFailureClassIds

`func (o *ApiV2TestResultsIdPutRequest) HasFailureClassIds() bool`

HasFailureClassIds returns a boolean if a field has been set.

### SetFailureClassIdsNil

`func (o *ApiV2TestResultsIdPutRequest) SetFailureClassIdsNil(b bool)`

 SetFailureClassIdsNil sets the value for FailureClassIds to be an explicit nil

### UnsetFailureClassIds
`func (o *ApiV2TestResultsIdPutRequest) UnsetFailureClassIds()`

UnsetFailureClassIds ensures that no value is present for FailureClassIds, not even an explicit nil
### GetOutcome

`func (o *ApiV2TestResultsIdPutRequest) GetOutcome() TestResultOutcome`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *ApiV2TestResultsIdPutRequest) GetOutcomeOk() (*TestResultOutcome, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *ApiV2TestResultsIdPutRequest) SetOutcome(v TestResultOutcome)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *ApiV2TestResultsIdPutRequest) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *ApiV2TestResultsIdPutRequest) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *ApiV2TestResultsIdPutRequest) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetComment

`func (o *ApiV2TestResultsIdPutRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApiV2TestResultsIdPutRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApiV2TestResultsIdPutRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApiV2TestResultsIdPutRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ApiV2TestResultsIdPutRequest) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ApiV2TestResultsIdPutRequest) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetLinks

`func (o *ApiV2TestResultsIdPutRequest) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiV2TestResultsIdPutRequest) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiV2TestResultsIdPutRequest) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApiV2TestResultsIdPutRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ApiV2TestResultsIdPutRequest) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ApiV2TestResultsIdPutRequest) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetStepResults

`func (o *ApiV2TestResultsIdPutRequest) GetStepResults() []StepResultModel`

GetStepResults returns the StepResults field if non-nil, zero value otherwise.

### GetStepResultsOk

`func (o *ApiV2TestResultsIdPutRequest) GetStepResultsOk() (*[]StepResultModel, bool)`

GetStepResultsOk returns a tuple with the StepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepResults

`func (o *ApiV2TestResultsIdPutRequest) SetStepResults(v []StepResultModel)`

SetStepResults sets StepResults field to given value.

### HasStepResults

`func (o *ApiV2TestResultsIdPutRequest) HasStepResults() bool`

HasStepResults returns a boolean if a field has been set.

### SetStepResultsNil

`func (o *ApiV2TestResultsIdPutRequest) SetStepResultsNil(b bool)`

 SetStepResultsNil sets the value for StepResults to be an explicit nil

### UnsetStepResults
`func (o *ApiV2TestResultsIdPutRequest) UnsetStepResults()`

UnsetStepResults ensures that no value is present for StepResults, not even an explicit nil
### GetAttachments

`func (o *ApiV2TestResultsIdPutRequest) GetAttachments() []AttachmentPutModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ApiV2TestResultsIdPutRequest) GetAttachmentsOk() (*[]AttachmentPutModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ApiV2TestResultsIdPutRequest) SetAttachments(v []AttachmentPutModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ApiV2TestResultsIdPutRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *ApiV2TestResultsIdPutRequest) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *ApiV2TestResultsIdPutRequest) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetDurationInMs

`func (o *ApiV2TestResultsIdPutRequest) GetDurationInMs() int64`

GetDurationInMs returns the DurationInMs field if non-nil, zero value otherwise.

### GetDurationInMsOk

`func (o *ApiV2TestResultsIdPutRequest) GetDurationInMsOk() (*int64, bool)`

GetDurationInMsOk returns a tuple with the DurationInMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMs

`func (o *ApiV2TestResultsIdPutRequest) SetDurationInMs(v int64)`

SetDurationInMs sets DurationInMs field to given value.

### HasDurationInMs

`func (o *ApiV2TestResultsIdPutRequest) HasDurationInMs() bool`

HasDurationInMs returns a boolean if a field has been set.

### SetDurationInMsNil

`func (o *ApiV2TestResultsIdPutRequest) SetDurationInMsNil(b bool)`

 SetDurationInMsNil sets the value for DurationInMs to be an explicit nil

### UnsetDurationInMs
`func (o *ApiV2TestResultsIdPutRequest) UnsetDurationInMs()`

UnsetDurationInMs ensures that no value is present for DurationInMs, not even an explicit nil
### GetDuration

`func (o *ApiV2TestResultsIdPutRequest) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ApiV2TestResultsIdPutRequest) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ApiV2TestResultsIdPutRequest) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ApiV2TestResultsIdPutRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *ApiV2TestResultsIdPutRequest) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *ApiV2TestResultsIdPutRequest) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetStepComments

`func (o *ApiV2TestResultsIdPutRequest) GetStepComments() []TestResultStepCommentPutModel`

GetStepComments returns the StepComments field if non-nil, zero value otherwise.

### GetStepCommentsOk

`func (o *ApiV2TestResultsIdPutRequest) GetStepCommentsOk() (*[]TestResultStepCommentPutModel, bool)`

GetStepCommentsOk returns a tuple with the StepComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepComments

`func (o *ApiV2TestResultsIdPutRequest) SetStepComments(v []TestResultStepCommentPutModel)`

SetStepComments sets StepComments field to given value.

### HasStepComments

`func (o *ApiV2TestResultsIdPutRequest) HasStepComments() bool`

HasStepComments returns a boolean if a field has been set.

### SetStepCommentsNil

`func (o *ApiV2TestResultsIdPutRequest) SetStepCommentsNil(b bool)`

 SetStepCommentsNil sets the value for StepComments to be an explicit nil

### UnsetStepComments
`func (o *ApiV2TestResultsIdPutRequest) UnsetStepComments()`

UnsetStepComments ensures that no value is present for StepComments, not even an explicit nil
### GetSetupResults

`func (o *ApiV2TestResultsIdPutRequest) GetSetupResults() []AttachmentPutModelAutoTestStepResultsModel`

GetSetupResults returns the SetupResults field if non-nil, zero value otherwise.

### GetSetupResultsOk

`func (o *ApiV2TestResultsIdPutRequest) GetSetupResultsOk() (*[]AttachmentPutModelAutoTestStepResultsModel, bool)`

GetSetupResultsOk returns a tuple with the SetupResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupResults

`func (o *ApiV2TestResultsIdPutRequest) SetSetupResults(v []AttachmentPutModelAutoTestStepResultsModel)`

SetSetupResults sets SetupResults field to given value.

### HasSetupResults

`func (o *ApiV2TestResultsIdPutRequest) HasSetupResults() bool`

HasSetupResults returns a boolean if a field has been set.

### SetSetupResultsNil

`func (o *ApiV2TestResultsIdPutRequest) SetSetupResultsNil(b bool)`

 SetSetupResultsNil sets the value for SetupResults to be an explicit nil

### UnsetSetupResults
`func (o *ApiV2TestResultsIdPutRequest) UnsetSetupResults()`

UnsetSetupResults ensures that no value is present for SetupResults, not even an explicit nil
### GetTeardownResults

`func (o *ApiV2TestResultsIdPutRequest) GetTeardownResults() []AttachmentPutModelAutoTestStepResultsModel`

GetTeardownResults returns the TeardownResults field if non-nil, zero value otherwise.

### GetTeardownResultsOk

`func (o *ApiV2TestResultsIdPutRequest) GetTeardownResultsOk() (*[]AttachmentPutModelAutoTestStepResultsModel, bool)`

GetTeardownResultsOk returns a tuple with the TeardownResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardownResults

`func (o *ApiV2TestResultsIdPutRequest) SetTeardownResults(v []AttachmentPutModelAutoTestStepResultsModel)`

SetTeardownResults sets TeardownResults field to given value.

### HasTeardownResults

`func (o *ApiV2TestResultsIdPutRequest) HasTeardownResults() bool`

HasTeardownResults returns a boolean if a field has been set.

### SetTeardownResultsNil

`func (o *ApiV2TestResultsIdPutRequest) SetTeardownResultsNil(b bool)`

 SetTeardownResultsNil sets the value for TeardownResults to be an explicit nil

### UnsetTeardownResults
`func (o *ApiV2TestResultsIdPutRequest) UnsetTeardownResults()`

UnsetTeardownResults ensures that no value is present for TeardownResults, not even an explicit nil
### GetMessage

`func (o *ApiV2TestResultsIdPutRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiV2TestResultsIdPutRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiV2TestResultsIdPutRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiV2TestResultsIdPutRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ApiV2TestResultsIdPutRequest) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ApiV2TestResultsIdPutRequest) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetTrace

`func (o *ApiV2TestResultsIdPutRequest) GetTrace() string`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *ApiV2TestResultsIdPutRequest) GetTraceOk() (*string, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *ApiV2TestResultsIdPutRequest) SetTrace(v string)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *ApiV2TestResultsIdPutRequest) HasTrace() bool`

HasTrace returns a boolean if a field has been set.

### SetTraceNil

`func (o *ApiV2TestResultsIdPutRequest) SetTraceNil(b bool)`

 SetTraceNil sets the value for Trace to be an explicit nil

### UnsetTrace
`func (o *ApiV2TestResultsIdPutRequest) UnsetTrace()`

UnsetTrace ensures that no value is present for Trace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


