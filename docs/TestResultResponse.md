# TestResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedDate** | **time.Time** |  | 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**CreatedById** | **string** |  | 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**StepComments** | Pointer to [**[]StepComment**](StepComment.md) |  | [optional] 
**FailureClassIds** | **[]string** |  | 
**Outcome** | Pointer to [**NullableTestResultOutcome**](TestResultOutcome.md) |  | [optional] 
**Status** | Pointer to [**NullableTestStatusApiResult**](TestStatusApiResult.md) |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 
**StepResults** | Pointer to [**[]StepResult**](StepResult.md) |  | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) |  | [optional] 
**AutoTestId** | Pointer to **NullableString** |  | [optional] 
**ConfigurationId** | **string** |  | 
**StartedOn** | Pointer to **NullableTime** |  | [optional] 
**CompletedOn** | Pointer to **NullableTime** |  | [optional] 
**DurationInMs** | Pointer to **NullableInt64** |  | [optional] 
**Traces** | Pointer to **NullableString** |  | [optional] 
**FailureType** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**RunByUserId** | Pointer to **NullableString** |  | [optional] 
**StoppedByUserId** | Pointer to **NullableString** |  | [optional] 
**TestPointId** | **string** |  | 
**TestRunId** | **string** |  | 
**TestPoint** | Pointer to [**NullableTestPoint**](TestPoint.md) |  | [optional] 
**AutoTest** | Pointer to [**NullableAutoTest**](AutoTest.md) |  | [optional] 
**AutoTestStepResults** | Pointer to [**[]AutoTestStepResult**](AutoTestStepResult.md) |  | [optional] 
**SetupResults** | Pointer to [**[]AutoTestStepResult**](AutoTestStepResult.md) |  | [optional] 
**TeardownResults** | Pointer to [**[]AutoTestStepResult**](AutoTestStepResult.md) |  | [optional] 
**WorkItemVersionId** | **string** |  | 
**WorkItemVersionNumber** | Pointer to **NullableInt32** |  | [optional] 
**Parameters** | Pointer to **map[string]string** |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTestResultResponse

`func NewTestResultResponse(id string, createdDate time.Time, createdById string, failureClassIds []string, configurationId string, testPointId string, testRunId string, workItemVersionId string, ) *TestResultResponse`

NewTestResultResponse instantiates a new TestResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultResponseWithDefaults

`func NewTestResultResponseWithDefaults() *TestResultResponse`

NewTestResultResponseWithDefaults instantiates a new TestResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestResultResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestResultResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestResultResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedDate

`func (o *TestResultResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestResultResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestResultResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *TestResultResponse) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestResultResponse) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestResultResponse) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestResultResponse) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *TestResultResponse) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *TestResultResponse) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *TestResultResponse) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TestResultResponse) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TestResultResponse) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedById

`func (o *TestResultResponse) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *TestResultResponse) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *TestResultResponse) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *TestResultResponse) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *TestResultResponse) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *TestResultResponse) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetStepComments

`func (o *TestResultResponse) GetStepComments() []StepComment`

GetStepComments returns the StepComments field if non-nil, zero value otherwise.

### GetStepCommentsOk

`func (o *TestResultResponse) GetStepCommentsOk() (*[]StepComment, bool)`

GetStepCommentsOk returns a tuple with the StepComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepComments

`func (o *TestResultResponse) SetStepComments(v []StepComment)`

SetStepComments sets StepComments field to given value.

### HasStepComments

`func (o *TestResultResponse) HasStepComments() bool`

HasStepComments returns a boolean if a field has been set.

### SetStepCommentsNil

`func (o *TestResultResponse) SetStepCommentsNil(b bool)`

 SetStepCommentsNil sets the value for StepComments to be an explicit nil

### UnsetStepComments
`func (o *TestResultResponse) UnsetStepComments()`

UnsetStepComments ensures that no value is present for StepComments, not even an explicit nil
### GetFailureClassIds

`func (o *TestResultResponse) GetFailureClassIds() []string`

GetFailureClassIds returns the FailureClassIds field if non-nil, zero value otherwise.

### GetFailureClassIdsOk

`func (o *TestResultResponse) GetFailureClassIdsOk() (*[]string, bool)`

GetFailureClassIdsOk returns a tuple with the FailureClassIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureClassIds

`func (o *TestResultResponse) SetFailureClassIds(v []string)`

SetFailureClassIds sets FailureClassIds field to given value.


### GetOutcome

`func (o *TestResultResponse) GetOutcome() TestResultOutcome`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *TestResultResponse) GetOutcomeOk() (*TestResultOutcome, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *TestResultResponse) SetOutcome(v TestResultOutcome)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *TestResultResponse) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *TestResultResponse) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *TestResultResponse) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetStatus

`func (o *TestResultResponse) GetStatus() TestStatusApiResult`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestResultResponse) GetStatusOk() (*TestStatusApiResult, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestResultResponse) SetStatus(v TestStatusApiResult)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestResultResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TestResultResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TestResultResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetComment

`func (o *TestResultResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TestResultResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TestResultResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TestResultResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *TestResultResponse) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *TestResultResponse) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetLinks

`func (o *TestResultResponse) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TestResultResponse) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TestResultResponse) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TestResultResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TestResultResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TestResultResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetStepResults

`func (o *TestResultResponse) GetStepResults() []StepResult`

GetStepResults returns the StepResults field if non-nil, zero value otherwise.

### GetStepResultsOk

`func (o *TestResultResponse) GetStepResultsOk() (*[]StepResult, bool)`

GetStepResultsOk returns a tuple with the StepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepResults

`func (o *TestResultResponse) SetStepResults(v []StepResult)`

SetStepResults sets StepResults field to given value.

### HasStepResults

`func (o *TestResultResponse) HasStepResults() bool`

HasStepResults returns a boolean if a field has been set.

### SetStepResultsNil

`func (o *TestResultResponse) SetStepResultsNil(b bool)`

 SetStepResultsNil sets the value for StepResults to be an explicit nil

### UnsetStepResults
`func (o *TestResultResponse) UnsetStepResults()`

UnsetStepResults ensures that no value is present for StepResults, not even an explicit nil
### GetAttachments

`func (o *TestResultResponse) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestResultResponse) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestResultResponse) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestResultResponse) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *TestResultResponse) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *TestResultResponse) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetAutoTestId

`func (o *TestResultResponse) GetAutoTestId() string`

GetAutoTestId returns the AutoTestId field if non-nil, zero value otherwise.

### GetAutoTestIdOk

`func (o *TestResultResponse) GetAutoTestIdOk() (*string, bool)`

GetAutoTestIdOk returns a tuple with the AutoTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestId

`func (o *TestResultResponse) SetAutoTestId(v string)`

SetAutoTestId sets AutoTestId field to given value.

### HasAutoTestId

`func (o *TestResultResponse) HasAutoTestId() bool`

HasAutoTestId returns a boolean if a field has been set.

### SetAutoTestIdNil

`func (o *TestResultResponse) SetAutoTestIdNil(b bool)`

 SetAutoTestIdNil sets the value for AutoTestId to be an explicit nil

### UnsetAutoTestId
`func (o *TestResultResponse) UnsetAutoTestId()`

UnsetAutoTestId ensures that no value is present for AutoTestId, not even an explicit nil
### GetConfigurationId

`func (o *TestResultResponse) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *TestResultResponse) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *TestResultResponse) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.


### GetStartedOn

`func (o *TestResultResponse) GetStartedOn() time.Time`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *TestResultResponse) GetStartedOnOk() (*time.Time, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *TestResultResponse) SetStartedOn(v time.Time)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *TestResultResponse) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *TestResultResponse) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *TestResultResponse) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetCompletedOn

`func (o *TestResultResponse) GetCompletedOn() time.Time`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *TestResultResponse) GetCompletedOnOk() (*time.Time, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *TestResultResponse) SetCompletedOn(v time.Time)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *TestResultResponse) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### SetCompletedOnNil

`func (o *TestResultResponse) SetCompletedOnNil(b bool)`

 SetCompletedOnNil sets the value for CompletedOn to be an explicit nil

### UnsetCompletedOn
`func (o *TestResultResponse) UnsetCompletedOn()`

UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
### GetDurationInMs

`func (o *TestResultResponse) GetDurationInMs() int64`

GetDurationInMs returns the DurationInMs field if non-nil, zero value otherwise.

### GetDurationInMsOk

`func (o *TestResultResponse) GetDurationInMsOk() (*int64, bool)`

GetDurationInMsOk returns a tuple with the DurationInMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMs

`func (o *TestResultResponse) SetDurationInMs(v int64)`

SetDurationInMs sets DurationInMs field to given value.

### HasDurationInMs

`func (o *TestResultResponse) HasDurationInMs() bool`

HasDurationInMs returns a boolean if a field has been set.

### SetDurationInMsNil

`func (o *TestResultResponse) SetDurationInMsNil(b bool)`

 SetDurationInMsNil sets the value for DurationInMs to be an explicit nil

### UnsetDurationInMs
`func (o *TestResultResponse) UnsetDurationInMs()`

UnsetDurationInMs ensures that no value is present for DurationInMs, not even an explicit nil
### GetTraces

`func (o *TestResultResponse) GetTraces() string`

GetTraces returns the Traces field if non-nil, zero value otherwise.

### GetTracesOk

`func (o *TestResultResponse) GetTracesOk() (*string, bool)`

GetTracesOk returns a tuple with the Traces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraces

`func (o *TestResultResponse) SetTraces(v string)`

SetTraces sets Traces field to given value.

### HasTraces

`func (o *TestResultResponse) HasTraces() bool`

HasTraces returns a boolean if a field has been set.

### SetTracesNil

`func (o *TestResultResponse) SetTracesNil(b bool)`

 SetTracesNil sets the value for Traces to be an explicit nil

### UnsetTraces
`func (o *TestResultResponse) UnsetTraces()`

UnsetTraces ensures that no value is present for Traces, not even an explicit nil
### GetFailureType

`func (o *TestResultResponse) GetFailureType() string`

GetFailureType returns the FailureType field if non-nil, zero value otherwise.

### GetFailureTypeOk

`func (o *TestResultResponse) GetFailureTypeOk() (*string, bool)`

GetFailureTypeOk returns a tuple with the FailureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureType

`func (o *TestResultResponse) SetFailureType(v string)`

SetFailureType sets FailureType field to given value.

### HasFailureType

`func (o *TestResultResponse) HasFailureType() bool`

HasFailureType returns a boolean if a field has been set.

### SetFailureTypeNil

`func (o *TestResultResponse) SetFailureTypeNil(b bool)`

 SetFailureTypeNil sets the value for FailureType to be an explicit nil

### UnsetFailureType
`func (o *TestResultResponse) UnsetFailureType()`

UnsetFailureType ensures that no value is present for FailureType, not even an explicit nil
### GetMessage

`func (o *TestResultResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TestResultResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TestResultResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TestResultResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *TestResultResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *TestResultResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetRunByUserId

`func (o *TestResultResponse) GetRunByUserId() string`

GetRunByUserId returns the RunByUserId field if non-nil, zero value otherwise.

### GetRunByUserIdOk

`func (o *TestResultResponse) GetRunByUserIdOk() (*string, bool)`

GetRunByUserIdOk returns a tuple with the RunByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunByUserId

`func (o *TestResultResponse) SetRunByUserId(v string)`

SetRunByUserId sets RunByUserId field to given value.

### HasRunByUserId

`func (o *TestResultResponse) HasRunByUserId() bool`

HasRunByUserId returns a boolean if a field has been set.

### SetRunByUserIdNil

`func (o *TestResultResponse) SetRunByUserIdNil(b bool)`

 SetRunByUserIdNil sets the value for RunByUserId to be an explicit nil

### UnsetRunByUserId
`func (o *TestResultResponse) UnsetRunByUserId()`

UnsetRunByUserId ensures that no value is present for RunByUserId, not even an explicit nil
### GetStoppedByUserId

`func (o *TestResultResponse) GetStoppedByUserId() string`

GetStoppedByUserId returns the StoppedByUserId field if non-nil, zero value otherwise.

### GetStoppedByUserIdOk

`func (o *TestResultResponse) GetStoppedByUserIdOk() (*string, bool)`

GetStoppedByUserIdOk returns a tuple with the StoppedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedByUserId

`func (o *TestResultResponse) SetStoppedByUserId(v string)`

SetStoppedByUserId sets StoppedByUserId field to given value.

### HasStoppedByUserId

`func (o *TestResultResponse) HasStoppedByUserId() bool`

HasStoppedByUserId returns a boolean if a field has been set.

### SetStoppedByUserIdNil

`func (o *TestResultResponse) SetStoppedByUserIdNil(b bool)`

 SetStoppedByUserIdNil sets the value for StoppedByUserId to be an explicit nil

### UnsetStoppedByUserId
`func (o *TestResultResponse) UnsetStoppedByUserId()`

UnsetStoppedByUserId ensures that no value is present for StoppedByUserId, not even an explicit nil
### GetTestPointId

`func (o *TestResultResponse) GetTestPointId() string`

GetTestPointId returns the TestPointId field if non-nil, zero value otherwise.

### GetTestPointIdOk

`func (o *TestResultResponse) GetTestPointIdOk() (*string, bool)`

GetTestPointIdOk returns a tuple with the TestPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPointId

`func (o *TestResultResponse) SetTestPointId(v string)`

SetTestPointId sets TestPointId field to given value.


### GetTestRunId

`func (o *TestResultResponse) GetTestRunId() string`

GetTestRunId returns the TestRunId field if non-nil, zero value otherwise.

### GetTestRunIdOk

`func (o *TestResultResponse) GetTestRunIdOk() (*string, bool)`

GetTestRunIdOk returns a tuple with the TestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunId

`func (o *TestResultResponse) SetTestRunId(v string)`

SetTestRunId sets TestRunId field to given value.


### GetTestPoint

`func (o *TestResultResponse) GetTestPoint() TestPoint`

GetTestPoint returns the TestPoint field if non-nil, zero value otherwise.

### GetTestPointOk

`func (o *TestResultResponse) GetTestPointOk() (*TestPoint, bool)`

GetTestPointOk returns a tuple with the TestPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPoint

`func (o *TestResultResponse) SetTestPoint(v TestPoint)`

SetTestPoint sets TestPoint field to given value.

### HasTestPoint

`func (o *TestResultResponse) HasTestPoint() bool`

HasTestPoint returns a boolean if a field has been set.

### SetTestPointNil

`func (o *TestResultResponse) SetTestPointNil(b bool)`

 SetTestPointNil sets the value for TestPoint to be an explicit nil

### UnsetTestPoint
`func (o *TestResultResponse) UnsetTestPoint()`

UnsetTestPoint ensures that no value is present for TestPoint, not even an explicit nil
### GetAutoTest

`func (o *TestResultResponse) GetAutoTest() AutoTest`

GetAutoTest returns the AutoTest field if non-nil, zero value otherwise.

### GetAutoTestOk

`func (o *TestResultResponse) GetAutoTestOk() (*AutoTest, bool)`

GetAutoTestOk returns a tuple with the AutoTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTest

`func (o *TestResultResponse) SetAutoTest(v AutoTest)`

SetAutoTest sets AutoTest field to given value.

### HasAutoTest

`func (o *TestResultResponse) HasAutoTest() bool`

HasAutoTest returns a boolean if a field has been set.

### SetAutoTestNil

`func (o *TestResultResponse) SetAutoTestNil(b bool)`

 SetAutoTestNil sets the value for AutoTest to be an explicit nil

### UnsetAutoTest
`func (o *TestResultResponse) UnsetAutoTest()`

UnsetAutoTest ensures that no value is present for AutoTest, not even an explicit nil
### GetAutoTestStepResults

`func (o *TestResultResponse) GetAutoTestStepResults() []AutoTestStepResult`

GetAutoTestStepResults returns the AutoTestStepResults field if non-nil, zero value otherwise.

### GetAutoTestStepResultsOk

`func (o *TestResultResponse) GetAutoTestStepResultsOk() (*[]AutoTestStepResult, bool)`

GetAutoTestStepResultsOk returns a tuple with the AutoTestStepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestStepResults

`func (o *TestResultResponse) SetAutoTestStepResults(v []AutoTestStepResult)`

SetAutoTestStepResults sets AutoTestStepResults field to given value.

### HasAutoTestStepResults

`func (o *TestResultResponse) HasAutoTestStepResults() bool`

HasAutoTestStepResults returns a boolean if a field has been set.

### SetAutoTestStepResultsNil

`func (o *TestResultResponse) SetAutoTestStepResultsNil(b bool)`

 SetAutoTestStepResultsNil sets the value for AutoTestStepResults to be an explicit nil

### UnsetAutoTestStepResults
`func (o *TestResultResponse) UnsetAutoTestStepResults()`

UnsetAutoTestStepResults ensures that no value is present for AutoTestStepResults, not even an explicit nil
### GetSetupResults

`func (o *TestResultResponse) GetSetupResults() []AutoTestStepResult`

GetSetupResults returns the SetupResults field if non-nil, zero value otherwise.

### GetSetupResultsOk

`func (o *TestResultResponse) GetSetupResultsOk() (*[]AutoTestStepResult, bool)`

GetSetupResultsOk returns a tuple with the SetupResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupResults

`func (o *TestResultResponse) SetSetupResults(v []AutoTestStepResult)`

SetSetupResults sets SetupResults field to given value.

### HasSetupResults

`func (o *TestResultResponse) HasSetupResults() bool`

HasSetupResults returns a boolean if a field has been set.

### SetSetupResultsNil

`func (o *TestResultResponse) SetSetupResultsNil(b bool)`

 SetSetupResultsNil sets the value for SetupResults to be an explicit nil

### UnsetSetupResults
`func (o *TestResultResponse) UnsetSetupResults()`

UnsetSetupResults ensures that no value is present for SetupResults, not even an explicit nil
### GetTeardownResults

`func (o *TestResultResponse) GetTeardownResults() []AutoTestStepResult`

GetTeardownResults returns the TeardownResults field if non-nil, zero value otherwise.

### GetTeardownResultsOk

`func (o *TestResultResponse) GetTeardownResultsOk() (*[]AutoTestStepResult, bool)`

GetTeardownResultsOk returns a tuple with the TeardownResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardownResults

`func (o *TestResultResponse) SetTeardownResults(v []AutoTestStepResult)`

SetTeardownResults sets TeardownResults field to given value.

### HasTeardownResults

`func (o *TestResultResponse) HasTeardownResults() bool`

HasTeardownResults returns a boolean if a field has been set.

### SetTeardownResultsNil

`func (o *TestResultResponse) SetTeardownResultsNil(b bool)`

 SetTeardownResultsNil sets the value for TeardownResults to be an explicit nil

### UnsetTeardownResults
`func (o *TestResultResponse) UnsetTeardownResults()`

UnsetTeardownResults ensures that no value is present for TeardownResults, not even an explicit nil
### GetWorkItemVersionId

`func (o *TestResultResponse) GetWorkItemVersionId() string`

GetWorkItemVersionId returns the WorkItemVersionId field if non-nil, zero value otherwise.

### GetWorkItemVersionIdOk

`func (o *TestResultResponse) GetWorkItemVersionIdOk() (*string, bool)`

GetWorkItemVersionIdOk returns a tuple with the WorkItemVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionId

`func (o *TestResultResponse) SetWorkItemVersionId(v string)`

SetWorkItemVersionId sets WorkItemVersionId field to given value.


### GetWorkItemVersionNumber

`func (o *TestResultResponse) GetWorkItemVersionNumber() int32`

GetWorkItemVersionNumber returns the WorkItemVersionNumber field if non-nil, zero value otherwise.

### GetWorkItemVersionNumberOk

`func (o *TestResultResponse) GetWorkItemVersionNumberOk() (*int32, bool)`

GetWorkItemVersionNumberOk returns a tuple with the WorkItemVersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionNumber

`func (o *TestResultResponse) SetWorkItemVersionNumber(v int32)`

SetWorkItemVersionNumber sets WorkItemVersionNumber field to given value.

### HasWorkItemVersionNumber

`func (o *TestResultResponse) HasWorkItemVersionNumber() bool`

HasWorkItemVersionNumber returns a boolean if a field has been set.

### SetWorkItemVersionNumberNil

`func (o *TestResultResponse) SetWorkItemVersionNumberNil(b bool)`

 SetWorkItemVersionNumberNil sets the value for WorkItemVersionNumber to be an explicit nil

### UnsetWorkItemVersionNumber
`func (o *TestResultResponse) UnsetWorkItemVersionNumber()`

UnsetWorkItemVersionNumber ensures that no value is present for WorkItemVersionNumber, not even an explicit nil
### GetParameters

`func (o *TestResultResponse) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TestResultResponse) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TestResultResponse) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TestResultResponse) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *TestResultResponse) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *TestResultResponse) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetProperties

`func (o *TestResultResponse) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TestResultResponse) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TestResultResponse) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TestResultResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *TestResultResponse) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *TestResultResponse) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


