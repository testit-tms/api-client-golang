# TestResultHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal test result identifier | 
**CreatedDate** | **time.Time** | Test result creation date | 
**ModifiedDate** | **time.Time** | Test result last modification date | 
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
**Comment** | Pointer to **NullableString** | Test result comment                If any test result related to the test run is linked with autotest, comment will have default value.  Otherwise, the comment equals to the comment of first created test result in the test run | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) | Test result links                If any test result related to the test run is linked with autotest, link will be equal to the links of last modified test result.  Otherwise, the links equals to the links of first created test result in the test run. | [optional] 
**StartedOn** | Pointer to **NullableTime** | Start date time from test result or from test run (if test run new state is Running or Completed state) | [optional] 
**CompletedOn** | Pointer to **NullableTime** | End date time from test result or from test run (if test run new state is In progress, Stopped or Completed) | [optional] 
**Duration** | Pointer to **NullableInt64** | Duration of first created test result in the test run | [optional] 
**CreatedById** | **string** | Unique identifier of user who created first test result in the test run | 
**ModifiedById** | Pointer to **NullableString** | Unique identifier of user who applied last modification of first test result in the test run | [optional] 
**Attachments** | Pointer to [**[]AttachmentModel**](AttachmentModel.md) | Attachments related to the test result                If any test result related to the test run is linked with autotest, attachments will be equal to the attachments of last modified  test result. Otherwise, the attachments equals to the attachments of first created test result in the test run. | [optional] 
**WorkItemVersionId** | Pointer to **NullableString** | Unique identifier of workitem version related to the first test result in the test run | [optional] 
**WorkItemVersionNumber** | Pointer to **NullableInt32** | Number of workitem version related to the first test result in the test run | [optional] 
**LaunchSource** | Pointer to **NullableString** |  | [optional] 
**FailureClassIds** | **[]string** | Unique identifier of failure classes related to the first test result in the test run | 
**Parameters** | Pointer to **map[string]string** | Parameters of test result | [optional] 

## Methods

### NewTestResultHistoryResponse

`func NewTestResultHistoryResponse(id string, createdDate time.Time, modifiedDate time.Time, userId string, isAutomated bool, createdById string, failureClassIds []string, ) *TestResultHistoryResponse`

NewTestResultHistoryResponse instantiates a new TestResultHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultHistoryResponseWithDefaults

`func NewTestResultHistoryResponseWithDefaults() *TestResultHistoryResponse`

NewTestResultHistoryResponseWithDefaults instantiates a new TestResultHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestResultHistoryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestResultHistoryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestResultHistoryResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedDate

`func (o *TestResultHistoryResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestResultHistoryResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestResultHistoryResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *TestResultHistoryResponse) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestResultHistoryResponse) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestResultHistoryResponse) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.


### GetUserId

`func (o *TestResultHistoryResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TestResultHistoryResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TestResultHistoryResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetTestRunId

`func (o *TestResultHistoryResponse) GetTestRunId() string`

GetTestRunId returns the TestRunId field if non-nil, zero value otherwise.

### GetTestRunIdOk

`func (o *TestResultHistoryResponse) GetTestRunIdOk() (*string, bool)`

GetTestRunIdOk returns a tuple with the TestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunId

`func (o *TestResultHistoryResponse) SetTestRunId(v string)`

SetTestRunId sets TestRunId field to given value.

### HasTestRunId

`func (o *TestResultHistoryResponse) HasTestRunId() bool`

HasTestRunId returns a boolean if a field has been set.

### SetTestRunIdNil

`func (o *TestResultHistoryResponse) SetTestRunIdNil(b bool)`

 SetTestRunIdNil sets the value for TestRunId to be an explicit nil

### UnsetTestRunId
`func (o *TestResultHistoryResponse) UnsetTestRunId()`

UnsetTestRunId ensures that no value is present for TestRunId, not even an explicit nil
### GetTestRunName

`func (o *TestResultHistoryResponse) GetTestRunName() string`

GetTestRunName returns the TestRunName field if non-nil, zero value otherwise.

### GetTestRunNameOk

`func (o *TestResultHistoryResponse) GetTestRunNameOk() (*string, bool)`

GetTestRunNameOk returns a tuple with the TestRunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunName

`func (o *TestResultHistoryResponse) SetTestRunName(v string)`

SetTestRunName sets TestRunName field to given value.

### HasTestRunName

`func (o *TestResultHistoryResponse) HasTestRunName() bool`

HasTestRunName returns a boolean if a field has been set.

### SetTestRunNameNil

`func (o *TestResultHistoryResponse) SetTestRunNameNil(b bool)`

 SetTestRunNameNil sets the value for TestRunName to be an explicit nil

### UnsetTestRunName
`func (o *TestResultHistoryResponse) UnsetTestRunName()`

UnsetTestRunName ensures that no value is present for TestRunName, not even an explicit nil
### GetCreatedByUserName

`func (o *TestResultHistoryResponse) GetCreatedByUserName() string`

GetCreatedByUserName returns the CreatedByUserName field if non-nil, zero value otherwise.

### GetCreatedByUserNameOk

`func (o *TestResultHistoryResponse) GetCreatedByUserNameOk() (*string, bool)`

GetCreatedByUserNameOk returns a tuple with the CreatedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserName

`func (o *TestResultHistoryResponse) SetCreatedByUserName(v string)`

SetCreatedByUserName sets CreatedByUserName field to given value.

### HasCreatedByUserName

`func (o *TestResultHistoryResponse) HasCreatedByUserName() bool`

HasCreatedByUserName returns a boolean if a field has been set.

### SetCreatedByUserNameNil

`func (o *TestResultHistoryResponse) SetCreatedByUserNameNil(b bool)`

 SetCreatedByUserNameNil sets the value for CreatedByUserName to be an explicit nil

### UnsetCreatedByUserName
`func (o *TestResultHistoryResponse) UnsetCreatedByUserName()`

UnsetCreatedByUserName ensures that no value is present for CreatedByUserName, not even an explicit nil
### GetTestPlanId

`func (o *TestResultHistoryResponse) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *TestResultHistoryResponse) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *TestResultHistoryResponse) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.

### HasTestPlanId

`func (o *TestResultHistoryResponse) HasTestPlanId() bool`

HasTestPlanId returns a boolean if a field has been set.

### SetTestPlanIdNil

`func (o *TestResultHistoryResponse) SetTestPlanIdNil(b bool)`

 SetTestPlanIdNil sets the value for TestPlanId to be an explicit nil

### UnsetTestPlanId
`func (o *TestResultHistoryResponse) UnsetTestPlanId()`

UnsetTestPlanId ensures that no value is present for TestPlanId, not even an explicit nil
### GetTestPlanGlobalId

`func (o *TestResultHistoryResponse) GetTestPlanGlobalId() int64`

GetTestPlanGlobalId returns the TestPlanGlobalId field if non-nil, zero value otherwise.

### GetTestPlanGlobalIdOk

`func (o *TestResultHistoryResponse) GetTestPlanGlobalIdOk() (*int64, bool)`

GetTestPlanGlobalIdOk returns a tuple with the TestPlanGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanGlobalId

`func (o *TestResultHistoryResponse) SetTestPlanGlobalId(v int64)`

SetTestPlanGlobalId sets TestPlanGlobalId field to given value.

### HasTestPlanGlobalId

`func (o *TestResultHistoryResponse) HasTestPlanGlobalId() bool`

HasTestPlanGlobalId returns a boolean if a field has been set.

### SetTestPlanGlobalIdNil

`func (o *TestResultHistoryResponse) SetTestPlanGlobalIdNil(b bool)`

 SetTestPlanGlobalIdNil sets the value for TestPlanGlobalId to be an explicit nil

### UnsetTestPlanGlobalId
`func (o *TestResultHistoryResponse) UnsetTestPlanGlobalId()`

UnsetTestPlanGlobalId ensures that no value is present for TestPlanGlobalId, not even an explicit nil
### GetTestPlanName

`func (o *TestResultHistoryResponse) GetTestPlanName() string`

GetTestPlanName returns the TestPlanName field if non-nil, zero value otherwise.

### GetTestPlanNameOk

`func (o *TestResultHistoryResponse) GetTestPlanNameOk() (*string, bool)`

GetTestPlanNameOk returns a tuple with the TestPlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanName

`func (o *TestResultHistoryResponse) SetTestPlanName(v string)`

SetTestPlanName sets TestPlanName field to given value.

### HasTestPlanName

`func (o *TestResultHistoryResponse) HasTestPlanName() bool`

HasTestPlanName returns a boolean if a field has been set.

### SetTestPlanNameNil

`func (o *TestResultHistoryResponse) SetTestPlanNameNil(b bool)`

 SetTestPlanNameNil sets the value for TestPlanName to be an explicit nil

### UnsetTestPlanName
`func (o *TestResultHistoryResponse) UnsetTestPlanName()`

UnsetTestPlanName ensures that no value is present for TestPlanName, not even an explicit nil
### GetConfigurationName

`func (o *TestResultHistoryResponse) GetConfigurationName() string`

GetConfigurationName returns the ConfigurationName field if non-nil, zero value otherwise.

### GetConfigurationNameOk

`func (o *TestResultHistoryResponse) GetConfigurationNameOk() (*string, bool)`

GetConfigurationNameOk returns a tuple with the ConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationName

`func (o *TestResultHistoryResponse) SetConfigurationName(v string)`

SetConfigurationName sets ConfigurationName field to given value.

### HasConfigurationName

`func (o *TestResultHistoryResponse) HasConfigurationName() bool`

HasConfigurationName returns a boolean if a field has been set.

### SetConfigurationNameNil

`func (o *TestResultHistoryResponse) SetConfigurationNameNil(b bool)`

 SetConfigurationNameNil sets the value for ConfigurationName to be an explicit nil

### UnsetConfigurationName
`func (o *TestResultHistoryResponse) UnsetConfigurationName()`

UnsetConfigurationName ensures that no value is present for ConfigurationName, not even an explicit nil
### GetIsAutomated

`func (o *TestResultHistoryResponse) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *TestResultHistoryResponse) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *TestResultHistoryResponse) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.


### GetOutcome

`func (o *TestResultHistoryResponse) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *TestResultHistoryResponse) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *TestResultHistoryResponse) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *TestResultHistoryResponse) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *TestResultHistoryResponse) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *TestResultHistoryResponse) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetComment

`func (o *TestResultHistoryResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TestResultHistoryResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TestResultHistoryResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TestResultHistoryResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *TestResultHistoryResponse) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *TestResultHistoryResponse) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetLinks

`func (o *TestResultHistoryResponse) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TestResultHistoryResponse) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TestResultHistoryResponse) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TestResultHistoryResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TestResultHistoryResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TestResultHistoryResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetStartedOn

`func (o *TestResultHistoryResponse) GetStartedOn() time.Time`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *TestResultHistoryResponse) GetStartedOnOk() (*time.Time, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *TestResultHistoryResponse) SetStartedOn(v time.Time)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *TestResultHistoryResponse) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *TestResultHistoryResponse) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *TestResultHistoryResponse) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetCompletedOn

`func (o *TestResultHistoryResponse) GetCompletedOn() time.Time`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *TestResultHistoryResponse) GetCompletedOnOk() (*time.Time, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *TestResultHistoryResponse) SetCompletedOn(v time.Time)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *TestResultHistoryResponse) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### SetCompletedOnNil

`func (o *TestResultHistoryResponse) SetCompletedOnNil(b bool)`

 SetCompletedOnNil sets the value for CompletedOn to be an explicit nil

### UnsetCompletedOn
`func (o *TestResultHistoryResponse) UnsetCompletedOn()`

UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
### GetDuration

`func (o *TestResultHistoryResponse) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TestResultHistoryResponse) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TestResultHistoryResponse) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TestResultHistoryResponse) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *TestResultHistoryResponse) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *TestResultHistoryResponse) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetCreatedById

`func (o *TestResultHistoryResponse) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TestResultHistoryResponse) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TestResultHistoryResponse) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedById

`func (o *TestResultHistoryResponse) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *TestResultHistoryResponse) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *TestResultHistoryResponse) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *TestResultHistoryResponse) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *TestResultHistoryResponse) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *TestResultHistoryResponse) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetAttachments

`func (o *TestResultHistoryResponse) GetAttachments() []AttachmentModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestResultHistoryResponse) GetAttachmentsOk() (*[]AttachmentModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestResultHistoryResponse) SetAttachments(v []AttachmentModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestResultHistoryResponse) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *TestResultHistoryResponse) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *TestResultHistoryResponse) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetWorkItemVersionId

`func (o *TestResultHistoryResponse) GetWorkItemVersionId() string`

GetWorkItemVersionId returns the WorkItemVersionId field if non-nil, zero value otherwise.

### GetWorkItemVersionIdOk

`func (o *TestResultHistoryResponse) GetWorkItemVersionIdOk() (*string, bool)`

GetWorkItemVersionIdOk returns a tuple with the WorkItemVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionId

`func (o *TestResultHistoryResponse) SetWorkItemVersionId(v string)`

SetWorkItemVersionId sets WorkItemVersionId field to given value.

### HasWorkItemVersionId

`func (o *TestResultHistoryResponse) HasWorkItemVersionId() bool`

HasWorkItemVersionId returns a boolean if a field has been set.

### SetWorkItemVersionIdNil

`func (o *TestResultHistoryResponse) SetWorkItemVersionIdNil(b bool)`

 SetWorkItemVersionIdNil sets the value for WorkItemVersionId to be an explicit nil

### UnsetWorkItemVersionId
`func (o *TestResultHistoryResponse) UnsetWorkItemVersionId()`

UnsetWorkItemVersionId ensures that no value is present for WorkItemVersionId, not even an explicit nil
### GetWorkItemVersionNumber

`func (o *TestResultHistoryResponse) GetWorkItemVersionNumber() int32`

GetWorkItemVersionNumber returns the WorkItemVersionNumber field if non-nil, zero value otherwise.

### GetWorkItemVersionNumberOk

`func (o *TestResultHistoryResponse) GetWorkItemVersionNumberOk() (*int32, bool)`

GetWorkItemVersionNumberOk returns a tuple with the WorkItemVersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionNumber

`func (o *TestResultHistoryResponse) SetWorkItemVersionNumber(v int32)`

SetWorkItemVersionNumber sets WorkItemVersionNumber field to given value.

### HasWorkItemVersionNumber

`func (o *TestResultHistoryResponse) HasWorkItemVersionNumber() bool`

HasWorkItemVersionNumber returns a boolean if a field has been set.

### SetWorkItemVersionNumberNil

`func (o *TestResultHistoryResponse) SetWorkItemVersionNumberNil(b bool)`

 SetWorkItemVersionNumberNil sets the value for WorkItemVersionNumber to be an explicit nil

### UnsetWorkItemVersionNumber
`func (o *TestResultHistoryResponse) UnsetWorkItemVersionNumber()`

UnsetWorkItemVersionNumber ensures that no value is present for WorkItemVersionNumber, not even an explicit nil
### GetLaunchSource

`func (o *TestResultHistoryResponse) GetLaunchSource() string`

GetLaunchSource returns the LaunchSource field if non-nil, zero value otherwise.

### GetLaunchSourceOk

`func (o *TestResultHistoryResponse) GetLaunchSourceOk() (*string, bool)`

GetLaunchSourceOk returns a tuple with the LaunchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSource

`func (o *TestResultHistoryResponse) SetLaunchSource(v string)`

SetLaunchSource sets LaunchSource field to given value.

### HasLaunchSource

`func (o *TestResultHistoryResponse) HasLaunchSource() bool`

HasLaunchSource returns a boolean if a field has been set.

### SetLaunchSourceNil

`func (o *TestResultHistoryResponse) SetLaunchSourceNil(b bool)`

 SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil

### UnsetLaunchSource
`func (o *TestResultHistoryResponse) UnsetLaunchSource()`

UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
### GetFailureClassIds

`func (o *TestResultHistoryResponse) GetFailureClassIds() []string`

GetFailureClassIds returns the FailureClassIds field if non-nil, zero value otherwise.

### GetFailureClassIdsOk

`func (o *TestResultHistoryResponse) GetFailureClassIdsOk() (*[]string, bool)`

GetFailureClassIdsOk returns a tuple with the FailureClassIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureClassIds

`func (o *TestResultHistoryResponse) SetFailureClassIds(v []string)`

SetFailureClassIds sets FailureClassIds field to given value.


### GetParameters

`func (o *TestResultHistoryResponse) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TestResultHistoryResponse) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TestResultHistoryResponse) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TestResultHistoryResponse) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *TestResultHistoryResponse) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *TestResultHistoryResponse) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


