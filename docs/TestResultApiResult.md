# TestResultApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**StartedOn** | Pointer to **NullableTime** |  | [optional] 
**CompletedOn** | Pointer to **NullableTime** |  | [optional] 
**DurationInMs** | Pointer to **NullableInt64** |  | [optional] 
**Traces** | Pointer to **NullableString** |  | [optional] 
**FailureType** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**RunByUserId** | Pointer to **NullableString** |  | [optional] 
**StoppedByUserId** | Pointer to **NullableString** |  | [optional] 
**Outcome** | **string** |  | 
**AutoTestId** | Pointer to **NullableString** |  | [optional] 
**TestPointId** | Pointer to **NullableString** |  | [optional] 
**TestRunId** | **string** |  | 
**ConfigurationId** | **string** |  | 
**Status** | [**TestStatusApiResult**](TestStatusApiResult.md) |  | 
**TestPoint** | Pointer to [**NullableTestPointShortApiResult**](TestPointShortApiResult.md) |  | [optional] 
**AutoTest** | Pointer to [**NullableAutoTestApiResult**](AutoTestApiResult.md) |  | [optional] 
**AutoTestStepResults** | Pointer to [**[]AutoTestStepResultsApiResult**](AutoTestStepResultsApiResult.md) |  | [optional] 
**SetupResults** | Pointer to [**[]AutoTestStepResultsApiResult**](AutoTestStepResultsApiResult.md) |  | [optional] 
**TeardownResults** | Pointer to [**[]AutoTestStepResultsApiResult**](AutoTestStepResultsApiResult.md) |  | [optional] 
**WorkItemVersionId** | Pointer to **NullableString** |  | [optional] 
**WorkItemVersionNumber** | Pointer to **NullableInt32** |  | [optional] 
**Attachments** | [**[]AttachmentApiResult**](AttachmentApiResult.md) |  | 
**Links** | [**[]LinkApiResult**](LinkApiResult.md) |  | 
**FailureClasses** | [**[]TestResultFailureClassApiResult**](TestResultFailureClassApiResult.md) |  | 
**StepComments** | Pointer to [**[]StepCommentApiResult**](StepCommentApiResult.md) |  | [optional] 
**Parameters** | Pointer to **map[string]string** |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**CreatedDate** | **time.Time** |  | 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**CreatedById** | **string** |  | 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**IsDeleted** | **bool** |  | 

## Methods

### NewTestResultApiResult

`func NewTestResultApiResult(id string, outcome string, testRunId string, configurationId string, status TestStatusApiResult, attachments []AttachmentApiResult, links []LinkApiResult, failureClasses []TestResultFailureClassApiResult, createdDate time.Time, createdById string, isDeleted bool, ) *TestResultApiResult`

NewTestResultApiResult instantiates a new TestResultApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultApiResultWithDefaults

`func NewTestResultApiResultWithDefaults() *TestResultApiResult`

NewTestResultApiResultWithDefaults instantiates a new TestResultApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestResultApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestResultApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestResultApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetStartedOn

`func (o *TestResultApiResult) GetStartedOn() time.Time`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *TestResultApiResult) GetStartedOnOk() (*time.Time, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *TestResultApiResult) SetStartedOn(v time.Time)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *TestResultApiResult) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *TestResultApiResult) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *TestResultApiResult) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetCompletedOn

`func (o *TestResultApiResult) GetCompletedOn() time.Time`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *TestResultApiResult) GetCompletedOnOk() (*time.Time, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *TestResultApiResult) SetCompletedOn(v time.Time)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *TestResultApiResult) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### SetCompletedOnNil

`func (o *TestResultApiResult) SetCompletedOnNil(b bool)`

 SetCompletedOnNil sets the value for CompletedOn to be an explicit nil

### UnsetCompletedOn
`func (o *TestResultApiResult) UnsetCompletedOn()`

UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
### GetDurationInMs

`func (o *TestResultApiResult) GetDurationInMs() int64`

GetDurationInMs returns the DurationInMs field if non-nil, zero value otherwise.

### GetDurationInMsOk

`func (o *TestResultApiResult) GetDurationInMsOk() (*int64, bool)`

GetDurationInMsOk returns a tuple with the DurationInMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMs

`func (o *TestResultApiResult) SetDurationInMs(v int64)`

SetDurationInMs sets DurationInMs field to given value.

### HasDurationInMs

`func (o *TestResultApiResult) HasDurationInMs() bool`

HasDurationInMs returns a boolean if a field has been set.

### SetDurationInMsNil

`func (o *TestResultApiResult) SetDurationInMsNil(b bool)`

 SetDurationInMsNil sets the value for DurationInMs to be an explicit nil

### UnsetDurationInMs
`func (o *TestResultApiResult) UnsetDurationInMs()`

UnsetDurationInMs ensures that no value is present for DurationInMs, not even an explicit nil
### GetTraces

`func (o *TestResultApiResult) GetTraces() string`

GetTraces returns the Traces field if non-nil, zero value otherwise.

### GetTracesOk

`func (o *TestResultApiResult) GetTracesOk() (*string, bool)`

GetTracesOk returns a tuple with the Traces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraces

`func (o *TestResultApiResult) SetTraces(v string)`

SetTraces sets Traces field to given value.

### HasTraces

`func (o *TestResultApiResult) HasTraces() bool`

HasTraces returns a boolean if a field has been set.

### SetTracesNil

`func (o *TestResultApiResult) SetTracesNil(b bool)`

 SetTracesNil sets the value for Traces to be an explicit nil

### UnsetTraces
`func (o *TestResultApiResult) UnsetTraces()`

UnsetTraces ensures that no value is present for Traces, not even an explicit nil
### GetFailureType

`func (o *TestResultApiResult) GetFailureType() string`

GetFailureType returns the FailureType field if non-nil, zero value otherwise.

### GetFailureTypeOk

`func (o *TestResultApiResult) GetFailureTypeOk() (*string, bool)`

GetFailureTypeOk returns a tuple with the FailureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureType

`func (o *TestResultApiResult) SetFailureType(v string)`

SetFailureType sets FailureType field to given value.

### HasFailureType

`func (o *TestResultApiResult) HasFailureType() bool`

HasFailureType returns a boolean if a field has been set.

### SetFailureTypeNil

`func (o *TestResultApiResult) SetFailureTypeNil(b bool)`

 SetFailureTypeNil sets the value for FailureType to be an explicit nil

### UnsetFailureType
`func (o *TestResultApiResult) UnsetFailureType()`

UnsetFailureType ensures that no value is present for FailureType, not even an explicit nil
### GetMessage

`func (o *TestResultApiResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TestResultApiResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TestResultApiResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TestResultApiResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *TestResultApiResult) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *TestResultApiResult) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetRunByUserId

`func (o *TestResultApiResult) GetRunByUserId() string`

GetRunByUserId returns the RunByUserId field if non-nil, zero value otherwise.

### GetRunByUserIdOk

`func (o *TestResultApiResult) GetRunByUserIdOk() (*string, bool)`

GetRunByUserIdOk returns a tuple with the RunByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunByUserId

`func (o *TestResultApiResult) SetRunByUserId(v string)`

SetRunByUserId sets RunByUserId field to given value.

### HasRunByUserId

`func (o *TestResultApiResult) HasRunByUserId() bool`

HasRunByUserId returns a boolean if a field has been set.

### SetRunByUserIdNil

`func (o *TestResultApiResult) SetRunByUserIdNil(b bool)`

 SetRunByUserIdNil sets the value for RunByUserId to be an explicit nil

### UnsetRunByUserId
`func (o *TestResultApiResult) UnsetRunByUserId()`

UnsetRunByUserId ensures that no value is present for RunByUserId, not even an explicit nil
### GetStoppedByUserId

`func (o *TestResultApiResult) GetStoppedByUserId() string`

GetStoppedByUserId returns the StoppedByUserId field if non-nil, zero value otherwise.

### GetStoppedByUserIdOk

`func (o *TestResultApiResult) GetStoppedByUserIdOk() (*string, bool)`

GetStoppedByUserIdOk returns a tuple with the StoppedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedByUserId

`func (o *TestResultApiResult) SetStoppedByUserId(v string)`

SetStoppedByUserId sets StoppedByUserId field to given value.

### HasStoppedByUserId

`func (o *TestResultApiResult) HasStoppedByUserId() bool`

HasStoppedByUserId returns a boolean if a field has been set.

### SetStoppedByUserIdNil

`func (o *TestResultApiResult) SetStoppedByUserIdNil(b bool)`

 SetStoppedByUserIdNil sets the value for StoppedByUserId to be an explicit nil

### UnsetStoppedByUserId
`func (o *TestResultApiResult) UnsetStoppedByUserId()`

UnsetStoppedByUserId ensures that no value is present for StoppedByUserId, not even an explicit nil
### GetOutcome

`func (o *TestResultApiResult) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *TestResultApiResult) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *TestResultApiResult) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### GetAutoTestId

`func (o *TestResultApiResult) GetAutoTestId() string`

GetAutoTestId returns the AutoTestId field if non-nil, zero value otherwise.

### GetAutoTestIdOk

`func (o *TestResultApiResult) GetAutoTestIdOk() (*string, bool)`

GetAutoTestIdOk returns a tuple with the AutoTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestId

`func (o *TestResultApiResult) SetAutoTestId(v string)`

SetAutoTestId sets AutoTestId field to given value.

### HasAutoTestId

`func (o *TestResultApiResult) HasAutoTestId() bool`

HasAutoTestId returns a boolean if a field has been set.

### SetAutoTestIdNil

`func (o *TestResultApiResult) SetAutoTestIdNil(b bool)`

 SetAutoTestIdNil sets the value for AutoTestId to be an explicit nil

### UnsetAutoTestId
`func (o *TestResultApiResult) UnsetAutoTestId()`

UnsetAutoTestId ensures that no value is present for AutoTestId, not even an explicit nil
### GetTestPointId

`func (o *TestResultApiResult) GetTestPointId() string`

GetTestPointId returns the TestPointId field if non-nil, zero value otherwise.

### GetTestPointIdOk

`func (o *TestResultApiResult) GetTestPointIdOk() (*string, bool)`

GetTestPointIdOk returns a tuple with the TestPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPointId

`func (o *TestResultApiResult) SetTestPointId(v string)`

SetTestPointId sets TestPointId field to given value.

### HasTestPointId

`func (o *TestResultApiResult) HasTestPointId() bool`

HasTestPointId returns a boolean if a field has been set.

### SetTestPointIdNil

`func (o *TestResultApiResult) SetTestPointIdNil(b bool)`

 SetTestPointIdNil sets the value for TestPointId to be an explicit nil

### UnsetTestPointId
`func (o *TestResultApiResult) UnsetTestPointId()`

UnsetTestPointId ensures that no value is present for TestPointId, not even an explicit nil
### GetTestRunId

`func (o *TestResultApiResult) GetTestRunId() string`

GetTestRunId returns the TestRunId field if non-nil, zero value otherwise.

### GetTestRunIdOk

`func (o *TestResultApiResult) GetTestRunIdOk() (*string, bool)`

GetTestRunIdOk returns a tuple with the TestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunId

`func (o *TestResultApiResult) SetTestRunId(v string)`

SetTestRunId sets TestRunId field to given value.


### GetConfigurationId

`func (o *TestResultApiResult) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *TestResultApiResult) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *TestResultApiResult) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.


### GetStatus

`func (o *TestResultApiResult) GetStatus() TestStatusApiResult`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestResultApiResult) GetStatusOk() (*TestStatusApiResult, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestResultApiResult) SetStatus(v TestStatusApiResult)`

SetStatus sets Status field to given value.


### GetTestPoint

`func (o *TestResultApiResult) GetTestPoint() TestPointShortApiResult`

GetTestPoint returns the TestPoint field if non-nil, zero value otherwise.

### GetTestPointOk

`func (o *TestResultApiResult) GetTestPointOk() (*TestPointShortApiResult, bool)`

GetTestPointOk returns a tuple with the TestPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPoint

`func (o *TestResultApiResult) SetTestPoint(v TestPointShortApiResult)`

SetTestPoint sets TestPoint field to given value.

### HasTestPoint

`func (o *TestResultApiResult) HasTestPoint() bool`

HasTestPoint returns a boolean if a field has been set.

### SetTestPointNil

`func (o *TestResultApiResult) SetTestPointNil(b bool)`

 SetTestPointNil sets the value for TestPoint to be an explicit nil

### UnsetTestPoint
`func (o *TestResultApiResult) UnsetTestPoint()`

UnsetTestPoint ensures that no value is present for TestPoint, not even an explicit nil
### GetAutoTest

`func (o *TestResultApiResult) GetAutoTest() AutoTestApiResult`

GetAutoTest returns the AutoTest field if non-nil, zero value otherwise.

### GetAutoTestOk

`func (o *TestResultApiResult) GetAutoTestOk() (*AutoTestApiResult, bool)`

GetAutoTestOk returns a tuple with the AutoTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTest

`func (o *TestResultApiResult) SetAutoTest(v AutoTestApiResult)`

SetAutoTest sets AutoTest field to given value.

### HasAutoTest

`func (o *TestResultApiResult) HasAutoTest() bool`

HasAutoTest returns a boolean if a field has been set.

### SetAutoTestNil

`func (o *TestResultApiResult) SetAutoTestNil(b bool)`

 SetAutoTestNil sets the value for AutoTest to be an explicit nil

### UnsetAutoTest
`func (o *TestResultApiResult) UnsetAutoTest()`

UnsetAutoTest ensures that no value is present for AutoTest, not even an explicit nil
### GetAutoTestStepResults

`func (o *TestResultApiResult) GetAutoTestStepResults() []AutoTestStepResultsApiResult`

GetAutoTestStepResults returns the AutoTestStepResults field if non-nil, zero value otherwise.

### GetAutoTestStepResultsOk

`func (o *TestResultApiResult) GetAutoTestStepResultsOk() (*[]AutoTestStepResultsApiResult, bool)`

GetAutoTestStepResultsOk returns a tuple with the AutoTestStepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestStepResults

`func (o *TestResultApiResult) SetAutoTestStepResults(v []AutoTestStepResultsApiResult)`

SetAutoTestStepResults sets AutoTestStepResults field to given value.

### HasAutoTestStepResults

`func (o *TestResultApiResult) HasAutoTestStepResults() bool`

HasAutoTestStepResults returns a boolean if a field has been set.

### SetAutoTestStepResultsNil

`func (o *TestResultApiResult) SetAutoTestStepResultsNil(b bool)`

 SetAutoTestStepResultsNil sets the value for AutoTestStepResults to be an explicit nil

### UnsetAutoTestStepResults
`func (o *TestResultApiResult) UnsetAutoTestStepResults()`

UnsetAutoTestStepResults ensures that no value is present for AutoTestStepResults, not even an explicit nil
### GetSetupResults

`func (o *TestResultApiResult) GetSetupResults() []AutoTestStepResultsApiResult`

GetSetupResults returns the SetupResults field if non-nil, zero value otherwise.

### GetSetupResultsOk

`func (o *TestResultApiResult) GetSetupResultsOk() (*[]AutoTestStepResultsApiResult, bool)`

GetSetupResultsOk returns a tuple with the SetupResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupResults

`func (o *TestResultApiResult) SetSetupResults(v []AutoTestStepResultsApiResult)`

SetSetupResults sets SetupResults field to given value.

### HasSetupResults

`func (o *TestResultApiResult) HasSetupResults() bool`

HasSetupResults returns a boolean if a field has been set.

### SetSetupResultsNil

`func (o *TestResultApiResult) SetSetupResultsNil(b bool)`

 SetSetupResultsNil sets the value for SetupResults to be an explicit nil

### UnsetSetupResults
`func (o *TestResultApiResult) UnsetSetupResults()`

UnsetSetupResults ensures that no value is present for SetupResults, not even an explicit nil
### GetTeardownResults

`func (o *TestResultApiResult) GetTeardownResults() []AutoTestStepResultsApiResult`

GetTeardownResults returns the TeardownResults field if non-nil, zero value otherwise.

### GetTeardownResultsOk

`func (o *TestResultApiResult) GetTeardownResultsOk() (*[]AutoTestStepResultsApiResult, bool)`

GetTeardownResultsOk returns a tuple with the TeardownResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardownResults

`func (o *TestResultApiResult) SetTeardownResults(v []AutoTestStepResultsApiResult)`

SetTeardownResults sets TeardownResults field to given value.

### HasTeardownResults

`func (o *TestResultApiResult) HasTeardownResults() bool`

HasTeardownResults returns a boolean if a field has been set.

### SetTeardownResultsNil

`func (o *TestResultApiResult) SetTeardownResultsNil(b bool)`

 SetTeardownResultsNil sets the value for TeardownResults to be an explicit nil

### UnsetTeardownResults
`func (o *TestResultApiResult) UnsetTeardownResults()`

UnsetTeardownResults ensures that no value is present for TeardownResults, not even an explicit nil
### GetWorkItemVersionId

`func (o *TestResultApiResult) GetWorkItemVersionId() string`

GetWorkItemVersionId returns the WorkItemVersionId field if non-nil, zero value otherwise.

### GetWorkItemVersionIdOk

`func (o *TestResultApiResult) GetWorkItemVersionIdOk() (*string, bool)`

GetWorkItemVersionIdOk returns a tuple with the WorkItemVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionId

`func (o *TestResultApiResult) SetWorkItemVersionId(v string)`

SetWorkItemVersionId sets WorkItemVersionId field to given value.

### HasWorkItemVersionId

`func (o *TestResultApiResult) HasWorkItemVersionId() bool`

HasWorkItemVersionId returns a boolean if a field has been set.

### SetWorkItemVersionIdNil

`func (o *TestResultApiResult) SetWorkItemVersionIdNil(b bool)`

 SetWorkItemVersionIdNil sets the value for WorkItemVersionId to be an explicit nil

### UnsetWorkItemVersionId
`func (o *TestResultApiResult) UnsetWorkItemVersionId()`

UnsetWorkItemVersionId ensures that no value is present for WorkItemVersionId, not even an explicit nil
### GetWorkItemVersionNumber

`func (o *TestResultApiResult) GetWorkItemVersionNumber() int32`

GetWorkItemVersionNumber returns the WorkItemVersionNumber field if non-nil, zero value otherwise.

### GetWorkItemVersionNumberOk

`func (o *TestResultApiResult) GetWorkItemVersionNumberOk() (*int32, bool)`

GetWorkItemVersionNumberOk returns a tuple with the WorkItemVersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionNumber

`func (o *TestResultApiResult) SetWorkItemVersionNumber(v int32)`

SetWorkItemVersionNumber sets WorkItemVersionNumber field to given value.

### HasWorkItemVersionNumber

`func (o *TestResultApiResult) HasWorkItemVersionNumber() bool`

HasWorkItemVersionNumber returns a boolean if a field has been set.

### SetWorkItemVersionNumberNil

`func (o *TestResultApiResult) SetWorkItemVersionNumberNil(b bool)`

 SetWorkItemVersionNumberNil sets the value for WorkItemVersionNumber to be an explicit nil

### UnsetWorkItemVersionNumber
`func (o *TestResultApiResult) UnsetWorkItemVersionNumber()`

UnsetWorkItemVersionNumber ensures that no value is present for WorkItemVersionNumber, not even an explicit nil
### GetAttachments

`func (o *TestResultApiResult) GetAttachments() []AttachmentApiResult`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestResultApiResult) GetAttachmentsOk() (*[]AttachmentApiResult, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestResultApiResult) SetAttachments(v []AttachmentApiResult)`

SetAttachments sets Attachments field to given value.


### GetLinks

`func (o *TestResultApiResult) GetLinks() []LinkApiResult`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TestResultApiResult) GetLinksOk() (*[]LinkApiResult, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TestResultApiResult) SetLinks(v []LinkApiResult)`

SetLinks sets Links field to given value.


### GetFailureClasses

`func (o *TestResultApiResult) GetFailureClasses() []TestResultFailureClassApiResult`

GetFailureClasses returns the FailureClasses field if non-nil, zero value otherwise.

### GetFailureClassesOk

`func (o *TestResultApiResult) GetFailureClassesOk() (*[]TestResultFailureClassApiResult, bool)`

GetFailureClassesOk returns a tuple with the FailureClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureClasses

`func (o *TestResultApiResult) SetFailureClasses(v []TestResultFailureClassApiResult)`

SetFailureClasses sets FailureClasses field to given value.


### GetStepComments

`func (o *TestResultApiResult) GetStepComments() []StepCommentApiResult`

GetStepComments returns the StepComments field if non-nil, zero value otherwise.

### GetStepCommentsOk

`func (o *TestResultApiResult) GetStepCommentsOk() (*[]StepCommentApiResult, bool)`

GetStepCommentsOk returns a tuple with the StepComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepComments

`func (o *TestResultApiResult) SetStepComments(v []StepCommentApiResult)`

SetStepComments sets StepComments field to given value.

### HasStepComments

`func (o *TestResultApiResult) HasStepComments() bool`

HasStepComments returns a boolean if a field has been set.

### SetStepCommentsNil

`func (o *TestResultApiResult) SetStepCommentsNil(b bool)`

 SetStepCommentsNil sets the value for StepComments to be an explicit nil

### UnsetStepComments
`func (o *TestResultApiResult) UnsetStepComments()`

UnsetStepComments ensures that no value is present for StepComments, not even an explicit nil
### GetParameters

`func (o *TestResultApiResult) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TestResultApiResult) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TestResultApiResult) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TestResultApiResult) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *TestResultApiResult) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *TestResultApiResult) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetProperties

`func (o *TestResultApiResult) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TestResultApiResult) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TestResultApiResult) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TestResultApiResult) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *TestResultApiResult) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *TestResultApiResult) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetCreatedDate

`func (o *TestResultApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestResultApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestResultApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *TestResultApiResult) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestResultApiResult) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestResultApiResult) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestResultApiResult) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *TestResultApiResult) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *TestResultApiResult) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *TestResultApiResult) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TestResultApiResult) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TestResultApiResult) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedById

`func (o *TestResultApiResult) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *TestResultApiResult) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *TestResultApiResult) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *TestResultApiResult) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *TestResultApiResult) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *TestResultApiResult) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetIsDeleted

`func (o *TestResultApiResult) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TestResultApiResult) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TestResultApiResult) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


