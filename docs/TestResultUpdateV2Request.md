# TestResultUpdateV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureClassIds** | Pointer to **[]string** |  | [optional] 
**Outcome** | Pointer to [**NullableTestResultOutcome**](TestResultOutcome.md) |  | [optional] 
**StatusCode** | Pointer to **NullableString** |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 
**StepResults** | Pointer to [**[]StepResultApiModel**](StepResultApiModel.md) |  | [optional] 
**Attachments** | Pointer to [**[]AttachmentUpdateRequest**](AttachmentUpdateRequest.md) |  | [optional] 
**DurationInMs** | Pointer to **NullableInt64** |  | [optional] 
**Duration** | Pointer to **NullableInt64** |  | [optional] 
**StepComments** | Pointer to [**[]TestResultStepCommentUpdateRequest**](TestResultStepCommentUpdateRequest.md) |  | [optional] 
**SetupResults** | Pointer to [**[]AttachmentPutModelAutoTestStepResultsModel**](AttachmentPutModelAutoTestStepResultsModel.md) |  | [optional] 
**TeardownResults** | Pointer to [**[]AttachmentPutModelAutoTestStepResultsModel**](AttachmentPutModelAutoTestStepResultsModel.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Trace** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTestResultUpdateV2Request

`func NewTestResultUpdateV2Request() *TestResultUpdateV2Request`

NewTestResultUpdateV2Request instantiates a new TestResultUpdateV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultUpdateV2RequestWithDefaults

`func NewTestResultUpdateV2RequestWithDefaults() *TestResultUpdateV2Request`

NewTestResultUpdateV2RequestWithDefaults instantiates a new TestResultUpdateV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureClassIds

`func (o *TestResultUpdateV2Request) GetFailureClassIds() []string`

GetFailureClassIds returns the FailureClassIds field if non-nil, zero value otherwise.

### GetFailureClassIdsOk

`func (o *TestResultUpdateV2Request) GetFailureClassIdsOk() (*[]string, bool)`

GetFailureClassIdsOk returns a tuple with the FailureClassIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureClassIds

`func (o *TestResultUpdateV2Request) SetFailureClassIds(v []string)`

SetFailureClassIds sets FailureClassIds field to given value.

### HasFailureClassIds

`func (o *TestResultUpdateV2Request) HasFailureClassIds() bool`

HasFailureClassIds returns a boolean if a field has been set.

### SetFailureClassIdsNil

`func (o *TestResultUpdateV2Request) SetFailureClassIdsNil(b bool)`

 SetFailureClassIdsNil sets the value for FailureClassIds to be an explicit nil

### UnsetFailureClassIds
`func (o *TestResultUpdateV2Request) UnsetFailureClassIds()`

UnsetFailureClassIds ensures that no value is present for FailureClassIds, not even an explicit nil
### GetOutcome

`func (o *TestResultUpdateV2Request) GetOutcome() TestResultOutcome`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *TestResultUpdateV2Request) GetOutcomeOk() (*TestResultOutcome, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *TestResultUpdateV2Request) SetOutcome(v TestResultOutcome)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *TestResultUpdateV2Request) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *TestResultUpdateV2Request) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *TestResultUpdateV2Request) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetStatusCode

`func (o *TestResultUpdateV2Request) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *TestResultUpdateV2Request) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *TestResultUpdateV2Request) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *TestResultUpdateV2Request) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### SetStatusCodeNil

`func (o *TestResultUpdateV2Request) SetStatusCodeNil(b bool)`

 SetStatusCodeNil sets the value for StatusCode to be an explicit nil

### UnsetStatusCode
`func (o *TestResultUpdateV2Request) UnsetStatusCode()`

UnsetStatusCode ensures that no value is present for StatusCode, not even an explicit nil
### GetComment

`func (o *TestResultUpdateV2Request) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TestResultUpdateV2Request) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TestResultUpdateV2Request) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TestResultUpdateV2Request) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *TestResultUpdateV2Request) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *TestResultUpdateV2Request) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetLinks

`func (o *TestResultUpdateV2Request) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TestResultUpdateV2Request) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TestResultUpdateV2Request) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TestResultUpdateV2Request) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TestResultUpdateV2Request) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TestResultUpdateV2Request) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetStepResults

`func (o *TestResultUpdateV2Request) GetStepResults() []StepResultApiModel`

GetStepResults returns the StepResults field if non-nil, zero value otherwise.

### GetStepResultsOk

`func (o *TestResultUpdateV2Request) GetStepResultsOk() (*[]StepResultApiModel, bool)`

GetStepResultsOk returns a tuple with the StepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepResults

`func (o *TestResultUpdateV2Request) SetStepResults(v []StepResultApiModel)`

SetStepResults sets StepResults field to given value.

### HasStepResults

`func (o *TestResultUpdateV2Request) HasStepResults() bool`

HasStepResults returns a boolean if a field has been set.

### SetStepResultsNil

`func (o *TestResultUpdateV2Request) SetStepResultsNil(b bool)`

 SetStepResultsNil sets the value for StepResults to be an explicit nil

### UnsetStepResults
`func (o *TestResultUpdateV2Request) UnsetStepResults()`

UnsetStepResults ensures that no value is present for StepResults, not even an explicit nil
### GetAttachments

`func (o *TestResultUpdateV2Request) GetAttachments() []AttachmentUpdateRequest`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestResultUpdateV2Request) GetAttachmentsOk() (*[]AttachmentUpdateRequest, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestResultUpdateV2Request) SetAttachments(v []AttachmentUpdateRequest)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestResultUpdateV2Request) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *TestResultUpdateV2Request) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *TestResultUpdateV2Request) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetDurationInMs

`func (o *TestResultUpdateV2Request) GetDurationInMs() int64`

GetDurationInMs returns the DurationInMs field if non-nil, zero value otherwise.

### GetDurationInMsOk

`func (o *TestResultUpdateV2Request) GetDurationInMsOk() (*int64, bool)`

GetDurationInMsOk returns a tuple with the DurationInMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMs

`func (o *TestResultUpdateV2Request) SetDurationInMs(v int64)`

SetDurationInMs sets DurationInMs field to given value.

### HasDurationInMs

`func (o *TestResultUpdateV2Request) HasDurationInMs() bool`

HasDurationInMs returns a boolean if a field has been set.

### SetDurationInMsNil

`func (o *TestResultUpdateV2Request) SetDurationInMsNil(b bool)`

 SetDurationInMsNil sets the value for DurationInMs to be an explicit nil

### UnsetDurationInMs
`func (o *TestResultUpdateV2Request) UnsetDurationInMs()`

UnsetDurationInMs ensures that no value is present for DurationInMs, not even an explicit nil
### GetDuration

`func (o *TestResultUpdateV2Request) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TestResultUpdateV2Request) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TestResultUpdateV2Request) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TestResultUpdateV2Request) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *TestResultUpdateV2Request) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *TestResultUpdateV2Request) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetStepComments

`func (o *TestResultUpdateV2Request) GetStepComments() []TestResultStepCommentUpdateRequest`

GetStepComments returns the StepComments field if non-nil, zero value otherwise.

### GetStepCommentsOk

`func (o *TestResultUpdateV2Request) GetStepCommentsOk() (*[]TestResultStepCommentUpdateRequest, bool)`

GetStepCommentsOk returns a tuple with the StepComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepComments

`func (o *TestResultUpdateV2Request) SetStepComments(v []TestResultStepCommentUpdateRequest)`

SetStepComments sets StepComments field to given value.

### HasStepComments

`func (o *TestResultUpdateV2Request) HasStepComments() bool`

HasStepComments returns a boolean if a field has been set.

### SetStepCommentsNil

`func (o *TestResultUpdateV2Request) SetStepCommentsNil(b bool)`

 SetStepCommentsNil sets the value for StepComments to be an explicit nil

### UnsetStepComments
`func (o *TestResultUpdateV2Request) UnsetStepComments()`

UnsetStepComments ensures that no value is present for StepComments, not even an explicit nil
### GetSetupResults

`func (o *TestResultUpdateV2Request) GetSetupResults() []AttachmentPutModelAutoTestStepResultsModel`

GetSetupResults returns the SetupResults field if non-nil, zero value otherwise.

### GetSetupResultsOk

`func (o *TestResultUpdateV2Request) GetSetupResultsOk() (*[]AttachmentPutModelAutoTestStepResultsModel, bool)`

GetSetupResultsOk returns a tuple with the SetupResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupResults

`func (o *TestResultUpdateV2Request) SetSetupResults(v []AttachmentPutModelAutoTestStepResultsModel)`

SetSetupResults sets SetupResults field to given value.

### HasSetupResults

`func (o *TestResultUpdateV2Request) HasSetupResults() bool`

HasSetupResults returns a boolean if a field has been set.

### SetSetupResultsNil

`func (o *TestResultUpdateV2Request) SetSetupResultsNil(b bool)`

 SetSetupResultsNil sets the value for SetupResults to be an explicit nil

### UnsetSetupResults
`func (o *TestResultUpdateV2Request) UnsetSetupResults()`

UnsetSetupResults ensures that no value is present for SetupResults, not even an explicit nil
### GetTeardownResults

`func (o *TestResultUpdateV2Request) GetTeardownResults() []AttachmentPutModelAutoTestStepResultsModel`

GetTeardownResults returns the TeardownResults field if non-nil, zero value otherwise.

### GetTeardownResultsOk

`func (o *TestResultUpdateV2Request) GetTeardownResultsOk() (*[]AttachmentPutModelAutoTestStepResultsModel, bool)`

GetTeardownResultsOk returns a tuple with the TeardownResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardownResults

`func (o *TestResultUpdateV2Request) SetTeardownResults(v []AttachmentPutModelAutoTestStepResultsModel)`

SetTeardownResults sets TeardownResults field to given value.

### HasTeardownResults

`func (o *TestResultUpdateV2Request) HasTeardownResults() bool`

HasTeardownResults returns a boolean if a field has been set.

### SetTeardownResultsNil

`func (o *TestResultUpdateV2Request) SetTeardownResultsNil(b bool)`

 SetTeardownResultsNil sets the value for TeardownResults to be an explicit nil

### UnsetTeardownResults
`func (o *TestResultUpdateV2Request) UnsetTeardownResults()`

UnsetTeardownResults ensures that no value is present for TeardownResults, not even an explicit nil
### GetMessage

`func (o *TestResultUpdateV2Request) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TestResultUpdateV2Request) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TestResultUpdateV2Request) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TestResultUpdateV2Request) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *TestResultUpdateV2Request) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *TestResultUpdateV2Request) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetTrace

`func (o *TestResultUpdateV2Request) GetTrace() string`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *TestResultUpdateV2Request) GetTraceOk() (*string, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *TestResultUpdateV2Request) SetTrace(v string)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *TestResultUpdateV2Request) HasTrace() bool`

HasTrace returns a boolean if a field has been set.

### SetTraceNil

`func (o *TestResultUpdateV2Request) SetTraceNil(b bool)`

 SetTraceNil sets the value for Trace to be an explicit nil

### UnsetTrace
`func (o *TestResultUpdateV2Request) UnsetTrace()`

UnsetTrace ensures that no value is present for Trace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


