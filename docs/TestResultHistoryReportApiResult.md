# TestResultHistoryReportApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal test result identifier | 
**CreatedDate** | **time.Time** | Test result creation date | 
**ModifiedDate** | Pointer to **NullableTime** | Test result last modification date | [optional] 
**UserId** | **string** | Internal identifier of user who stopped test run related to the test result or user who created the test result                If test run was stopped, this property equals identifier of user who stopped it.  Otherwise, the property equals identifier of user who created the test result | 
**TestRunId** | Pointer to **NullableString** | Identifier of test run related to the test result | [optional] 
**TestRunName** | Pointer to **NullableString** | Name of test run related to the test result | [optional] 
**CreatedByUserName** | Pointer to **NullableString** | Username of user who created test run | [optional] 
**TestPlanId** | Pointer to **NullableString** | Internal identifier of test plan related to the test result&#39;s test run | [optional] 
**TestPlanGlobalId** | Pointer to **NullableInt64** | Global identifier of test plan related to the test result&#39;s test run | [optional] 
**TestPlanName** | Pointer to **NullableString** | Name of test plan related to the test result&#39;s test run | [optional] 
**ConfigurationName** | Pointer to **NullableString** | Configuration name of test point related to the test result or from test result itself                If test point related to the test result has configuration, this property will be equal to the test point configuration name.  Otherwise, this property will be equal to the test result configuration name | [optional] 
**IsAutomated** | **bool** | Boolean flag defines if test point related to the test result is automated or not | 
**Outcome** | Pointer to **NullableString** | Outcome from test result with max modified date or from first created test result                Property can contain one of these values: Passed, Failed, InProgress, Blocked, Skipped.                If any test result related to the test run is linked with autotest and the run has an outcome, the outcome value equals to the  worst outcome of the last modified test result. Otherwise, the outcome equals to the outcome of first created test result in the  test run. | [optional] 
**Status** | [**TestStatusApiResult**](TestStatusApiResult.md) | Status from test result with max modified date or from first created test result | 
**Comment** | Pointer to **NullableString** | Test result comment                If any test result related to the test run is linked with autotest, comment will have default value.  Otherwise, the comment equals to the comment of first created test result in the test run | [optional] 
**Links** | Pointer to [**[]LinkApiResult**](LinkApiResult.md) | Test result links                If any test result related to the test run is linked with autotest, link will be equal to the links of last modified test result.  Otherwise, the links equals to the links of first created test result in the test run. | [optional] 
**StartedOn** | Pointer to **NullableTime** | Start date time from test result or from test run (if test run new state is Running or Completed state) | [optional] 
**CompletedOn** | Pointer to **NullableTime** | End date time from test result or from test run (if test run new state is In progress, Stopped or Completed) | [optional] 
**Duration** | Pointer to **NullableInt64** | Duration of first created test result in the test run | [optional] 
**CreatedById** | **string** | Unique identifier of user who created first test result in the test run | 
**ModifiedById** | Pointer to **NullableString** | Unique identifier of user who applied last modification of first test result in the test run | [optional] 
**Attachments** | Pointer to [**[]AttachmentApiResult**](AttachmentApiResult.md) | Attachments related to the test result                If any test result related to the test run is linked with autotest, attachments will be equal to the attachments of last modified  test result. Otherwise, the attachments equals to the attachments of first created test result in the test run. | [optional] 
**WorkItemVersionId** | Pointer to **NullableString** | Unique identifier of workitem version related to the first test result in the test run | [optional] 
**WorkItemVersionNumber** | Pointer to **NullableInt32** | Number of workitem version related to the first test result in the test run | [optional] 
**LaunchSource** | Pointer to **NullableString** |  | [optional] 
**FailureClassIds** | **[]string** | Unique identifier of failure classes related to the first test result in the test run | 
**Parameters** | Pointer to **map[string]string** | Parameters of test result | [optional] 

## Methods

### NewTestResultHistoryReportApiResult

`func NewTestResultHistoryReportApiResult(id string, createdDate time.Time, userId string, isAutomated bool, status TestStatusApiResult, createdById string, failureClassIds []string, ) *TestResultHistoryReportApiResult`

NewTestResultHistoryReportApiResult instantiates a new TestResultHistoryReportApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultHistoryReportApiResultWithDefaults

`func NewTestResultHistoryReportApiResultWithDefaults() *TestResultHistoryReportApiResult`

NewTestResultHistoryReportApiResultWithDefaults instantiates a new TestResultHistoryReportApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestResultHistoryReportApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestResultHistoryReportApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestResultHistoryReportApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedDate

`func (o *TestResultHistoryReportApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestResultHistoryReportApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestResultHistoryReportApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *TestResultHistoryReportApiResult) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestResultHistoryReportApiResult) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestResultHistoryReportApiResult) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestResultHistoryReportApiResult) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *TestResultHistoryReportApiResult) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *TestResultHistoryReportApiResult) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetUserId

`func (o *TestResultHistoryReportApiResult) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TestResultHistoryReportApiResult) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TestResultHistoryReportApiResult) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetTestRunId

`func (o *TestResultHistoryReportApiResult) GetTestRunId() string`

GetTestRunId returns the TestRunId field if non-nil, zero value otherwise.

### GetTestRunIdOk

`func (o *TestResultHistoryReportApiResult) GetTestRunIdOk() (*string, bool)`

GetTestRunIdOk returns a tuple with the TestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunId

`func (o *TestResultHistoryReportApiResult) SetTestRunId(v string)`

SetTestRunId sets TestRunId field to given value.

### HasTestRunId

`func (o *TestResultHistoryReportApiResult) HasTestRunId() bool`

HasTestRunId returns a boolean if a field has been set.

### SetTestRunIdNil

`func (o *TestResultHistoryReportApiResult) SetTestRunIdNil(b bool)`

 SetTestRunIdNil sets the value for TestRunId to be an explicit nil

### UnsetTestRunId
`func (o *TestResultHistoryReportApiResult) UnsetTestRunId()`

UnsetTestRunId ensures that no value is present for TestRunId, not even an explicit nil
### GetTestRunName

`func (o *TestResultHistoryReportApiResult) GetTestRunName() string`

GetTestRunName returns the TestRunName field if non-nil, zero value otherwise.

### GetTestRunNameOk

`func (o *TestResultHistoryReportApiResult) GetTestRunNameOk() (*string, bool)`

GetTestRunNameOk returns a tuple with the TestRunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunName

`func (o *TestResultHistoryReportApiResult) SetTestRunName(v string)`

SetTestRunName sets TestRunName field to given value.

### HasTestRunName

`func (o *TestResultHistoryReportApiResult) HasTestRunName() bool`

HasTestRunName returns a boolean if a field has been set.

### SetTestRunNameNil

`func (o *TestResultHistoryReportApiResult) SetTestRunNameNil(b bool)`

 SetTestRunNameNil sets the value for TestRunName to be an explicit nil

### UnsetTestRunName
`func (o *TestResultHistoryReportApiResult) UnsetTestRunName()`

UnsetTestRunName ensures that no value is present for TestRunName, not even an explicit nil
### GetCreatedByUserName

`func (o *TestResultHistoryReportApiResult) GetCreatedByUserName() string`

GetCreatedByUserName returns the CreatedByUserName field if non-nil, zero value otherwise.

### GetCreatedByUserNameOk

`func (o *TestResultHistoryReportApiResult) GetCreatedByUserNameOk() (*string, bool)`

GetCreatedByUserNameOk returns a tuple with the CreatedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserName

`func (o *TestResultHistoryReportApiResult) SetCreatedByUserName(v string)`

SetCreatedByUserName sets CreatedByUserName field to given value.

### HasCreatedByUserName

`func (o *TestResultHistoryReportApiResult) HasCreatedByUserName() bool`

HasCreatedByUserName returns a boolean if a field has been set.

### SetCreatedByUserNameNil

`func (o *TestResultHistoryReportApiResult) SetCreatedByUserNameNil(b bool)`

 SetCreatedByUserNameNil sets the value for CreatedByUserName to be an explicit nil

### UnsetCreatedByUserName
`func (o *TestResultHistoryReportApiResult) UnsetCreatedByUserName()`

UnsetCreatedByUserName ensures that no value is present for CreatedByUserName, not even an explicit nil
### GetTestPlanId

`func (o *TestResultHistoryReportApiResult) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *TestResultHistoryReportApiResult) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *TestResultHistoryReportApiResult) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.

### HasTestPlanId

`func (o *TestResultHistoryReportApiResult) HasTestPlanId() bool`

HasTestPlanId returns a boolean if a field has been set.

### SetTestPlanIdNil

`func (o *TestResultHistoryReportApiResult) SetTestPlanIdNil(b bool)`

 SetTestPlanIdNil sets the value for TestPlanId to be an explicit nil

### UnsetTestPlanId
`func (o *TestResultHistoryReportApiResult) UnsetTestPlanId()`

UnsetTestPlanId ensures that no value is present for TestPlanId, not even an explicit nil
### GetTestPlanGlobalId

`func (o *TestResultHistoryReportApiResult) GetTestPlanGlobalId() int64`

GetTestPlanGlobalId returns the TestPlanGlobalId field if non-nil, zero value otherwise.

### GetTestPlanGlobalIdOk

`func (o *TestResultHistoryReportApiResult) GetTestPlanGlobalIdOk() (*int64, bool)`

GetTestPlanGlobalIdOk returns a tuple with the TestPlanGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanGlobalId

`func (o *TestResultHistoryReportApiResult) SetTestPlanGlobalId(v int64)`

SetTestPlanGlobalId sets TestPlanGlobalId field to given value.

### HasTestPlanGlobalId

`func (o *TestResultHistoryReportApiResult) HasTestPlanGlobalId() bool`

HasTestPlanGlobalId returns a boolean if a field has been set.

### SetTestPlanGlobalIdNil

`func (o *TestResultHistoryReportApiResult) SetTestPlanGlobalIdNil(b bool)`

 SetTestPlanGlobalIdNil sets the value for TestPlanGlobalId to be an explicit nil

### UnsetTestPlanGlobalId
`func (o *TestResultHistoryReportApiResult) UnsetTestPlanGlobalId()`

UnsetTestPlanGlobalId ensures that no value is present for TestPlanGlobalId, not even an explicit nil
### GetTestPlanName

`func (o *TestResultHistoryReportApiResult) GetTestPlanName() string`

GetTestPlanName returns the TestPlanName field if non-nil, zero value otherwise.

### GetTestPlanNameOk

`func (o *TestResultHistoryReportApiResult) GetTestPlanNameOk() (*string, bool)`

GetTestPlanNameOk returns a tuple with the TestPlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanName

`func (o *TestResultHistoryReportApiResult) SetTestPlanName(v string)`

SetTestPlanName sets TestPlanName field to given value.

### HasTestPlanName

`func (o *TestResultHistoryReportApiResult) HasTestPlanName() bool`

HasTestPlanName returns a boolean if a field has been set.

### SetTestPlanNameNil

`func (o *TestResultHistoryReportApiResult) SetTestPlanNameNil(b bool)`

 SetTestPlanNameNil sets the value for TestPlanName to be an explicit nil

### UnsetTestPlanName
`func (o *TestResultHistoryReportApiResult) UnsetTestPlanName()`

UnsetTestPlanName ensures that no value is present for TestPlanName, not even an explicit nil
### GetConfigurationName

`func (o *TestResultHistoryReportApiResult) GetConfigurationName() string`

GetConfigurationName returns the ConfigurationName field if non-nil, zero value otherwise.

### GetConfigurationNameOk

`func (o *TestResultHistoryReportApiResult) GetConfigurationNameOk() (*string, bool)`

GetConfigurationNameOk returns a tuple with the ConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationName

`func (o *TestResultHistoryReportApiResult) SetConfigurationName(v string)`

SetConfigurationName sets ConfigurationName field to given value.

### HasConfigurationName

`func (o *TestResultHistoryReportApiResult) HasConfigurationName() bool`

HasConfigurationName returns a boolean if a field has been set.

### SetConfigurationNameNil

`func (o *TestResultHistoryReportApiResult) SetConfigurationNameNil(b bool)`

 SetConfigurationNameNil sets the value for ConfigurationName to be an explicit nil

### UnsetConfigurationName
`func (o *TestResultHistoryReportApiResult) UnsetConfigurationName()`

UnsetConfigurationName ensures that no value is present for ConfigurationName, not even an explicit nil
### GetIsAutomated

`func (o *TestResultHistoryReportApiResult) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *TestResultHistoryReportApiResult) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *TestResultHistoryReportApiResult) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.


### GetOutcome

`func (o *TestResultHistoryReportApiResult) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *TestResultHistoryReportApiResult) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *TestResultHistoryReportApiResult) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *TestResultHistoryReportApiResult) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *TestResultHistoryReportApiResult) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *TestResultHistoryReportApiResult) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetStatus

`func (o *TestResultHistoryReportApiResult) GetStatus() TestStatusApiResult`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestResultHistoryReportApiResult) GetStatusOk() (*TestStatusApiResult, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestResultHistoryReportApiResult) SetStatus(v TestStatusApiResult)`

SetStatus sets Status field to given value.


### GetComment

`func (o *TestResultHistoryReportApiResult) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TestResultHistoryReportApiResult) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TestResultHistoryReportApiResult) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TestResultHistoryReportApiResult) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *TestResultHistoryReportApiResult) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *TestResultHistoryReportApiResult) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetLinks

`func (o *TestResultHistoryReportApiResult) GetLinks() []LinkApiResult`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TestResultHistoryReportApiResult) GetLinksOk() (*[]LinkApiResult, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TestResultHistoryReportApiResult) SetLinks(v []LinkApiResult)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TestResultHistoryReportApiResult) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TestResultHistoryReportApiResult) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TestResultHistoryReportApiResult) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetStartedOn

`func (o *TestResultHistoryReportApiResult) GetStartedOn() time.Time`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *TestResultHistoryReportApiResult) GetStartedOnOk() (*time.Time, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *TestResultHistoryReportApiResult) SetStartedOn(v time.Time)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *TestResultHistoryReportApiResult) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *TestResultHistoryReportApiResult) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *TestResultHistoryReportApiResult) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetCompletedOn

`func (o *TestResultHistoryReportApiResult) GetCompletedOn() time.Time`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *TestResultHistoryReportApiResult) GetCompletedOnOk() (*time.Time, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *TestResultHistoryReportApiResult) SetCompletedOn(v time.Time)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *TestResultHistoryReportApiResult) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### SetCompletedOnNil

`func (o *TestResultHistoryReportApiResult) SetCompletedOnNil(b bool)`

 SetCompletedOnNil sets the value for CompletedOn to be an explicit nil

### UnsetCompletedOn
`func (o *TestResultHistoryReportApiResult) UnsetCompletedOn()`

UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
### GetDuration

`func (o *TestResultHistoryReportApiResult) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TestResultHistoryReportApiResult) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TestResultHistoryReportApiResult) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TestResultHistoryReportApiResult) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *TestResultHistoryReportApiResult) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *TestResultHistoryReportApiResult) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetCreatedById

`func (o *TestResultHistoryReportApiResult) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TestResultHistoryReportApiResult) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TestResultHistoryReportApiResult) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedById

`func (o *TestResultHistoryReportApiResult) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *TestResultHistoryReportApiResult) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *TestResultHistoryReportApiResult) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *TestResultHistoryReportApiResult) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *TestResultHistoryReportApiResult) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *TestResultHistoryReportApiResult) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetAttachments

`func (o *TestResultHistoryReportApiResult) GetAttachments() []AttachmentApiResult`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestResultHistoryReportApiResult) GetAttachmentsOk() (*[]AttachmentApiResult, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestResultHistoryReportApiResult) SetAttachments(v []AttachmentApiResult)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestResultHistoryReportApiResult) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *TestResultHistoryReportApiResult) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *TestResultHistoryReportApiResult) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetWorkItemVersionId

`func (o *TestResultHistoryReportApiResult) GetWorkItemVersionId() string`

GetWorkItemVersionId returns the WorkItemVersionId field if non-nil, zero value otherwise.

### GetWorkItemVersionIdOk

`func (o *TestResultHistoryReportApiResult) GetWorkItemVersionIdOk() (*string, bool)`

GetWorkItemVersionIdOk returns a tuple with the WorkItemVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionId

`func (o *TestResultHistoryReportApiResult) SetWorkItemVersionId(v string)`

SetWorkItemVersionId sets WorkItemVersionId field to given value.

### HasWorkItemVersionId

`func (o *TestResultHistoryReportApiResult) HasWorkItemVersionId() bool`

HasWorkItemVersionId returns a boolean if a field has been set.

### SetWorkItemVersionIdNil

`func (o *TestResultHistoryReportApiResult) SetWorkItemVersionIdNil(b bool)`

 SetWorkItemVersionIdNil sets the value for WorkItemVersionId to be an explicit nil

### UnsetWorkItemVersionId
`func (o *TestResultHistoryReportApiResult) UnsetWorkItemVersionId()`

UnsetWorkItemVersionId ensures that no value is present for WorkItemVersionId, not even an explicit nil
### GetWorkItemVersionNumber

`func (o *TestResultHistoryReportApiResult) GetWorkItemVersionNumber() int32`

GetWorkItemVersionNumber returns the WorkItemVersionNumber field if non-nil, zero value otherwise.

### GetWorkItemVersionNumberOk

`func (o *TestResultHistoryReportApiResult) GetWorkItemVersionNumberOk() (*int32, bool)`

GetWorkItemVersionNumberOk returns a tuple with the WorkItemVersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionNumber

`func (o *TestResultHistoryReportApiResult) SetWorkItemVersionNumber(v int32)`

SetWorkItemVersionNumber sets WorkItemVersionNumber field to given value.

### HasWorkItemVersionNumber

`func (o *TestResultHistoryReportApiResult) HasWorkItemVersionNumber() bool`

HasWorkItemVersionNumber returns a boolean if a field has been set.

### SetWorkItemVersionNumberNil

`func (o *TestResultHistoryReportApiResult) SetWorkItemVersionNumberNil(b bool)`

 SetWorkItemVersionNumberNil sets the value for WorkItemVersionNumber to be an explicit nil

### UnsetWorkItemVersionNumber
`func (o *TestResultHistoryReportApiResult) UnsetWorkItemVersionNumber()`

UnsetWorkItemVersionNumber ensures that no value is present for WorkItemVersionNumber, not even an explicit nil
### GetLaunchSource

`func (o *TestResultHistoryReportApiResult) GetLaunchSource() string`

GetLaunchSource returns the LaunchSource field if non-nil, zero value otherwise.

### GetLaunchSourceOk

`func (o *TestResultHistoryReportApiResult) GetLaunchSourceOk() (*string, bool)`

GetLaunchSourceOk returns a tuple with the LaunchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSource

`func (o *TestResultHistoryReportApiResult) SetLaunchSource(v string)`

SetLaunchSource sets LaunchSource field to given value.

### HasLaunchSource

`func (o *TestResultHistoryReportApiResult) HasLaunchSource() bool`

HasLaunchSource returns a boolean if a field has been set.

### SetLaunchSourceNil

`func (o *TestResultHistoryReportApiResult) SetLaunchSourceNil(b bool)`

 SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil

### UnsetLaunchSource
`func (o *TestResultHistoryReportApiResult) UnsetLaunchSource()`

UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
### GetFailureClassIds

`func (o *TestResultHistoryReportApiResult) GetFailureClassIds() []string`

GetFailureClassIds returns the FailureClassIds field if non-nil, zero value otherwise.

### GetFailureClassIdsOk

`func (o *TestResultHistoryReportApiResult) GetFailureClassIdsOk() (*[]string, bool)`

GetFailureClassIdsOk returns a tuple with the FailureClassIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureClassIds

`func (o *TestResultHistoryReportApiResult) SetFailureClassIds(v []string)`

SetFailureClassIds sets FailureClassIds field to given value.


### GetParameters

`func (o *TestResultHistoryReportApiResult) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TestResultHistoryReportApiResult) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TestResultHistoryReportApiResult) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TestResultHistoryReportApiResult) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *TestResultHistoryReportApiResult) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *TestResultHistoryReportApiResult) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


