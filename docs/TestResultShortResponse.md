# TestResultShortResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the test result | 
**Name** | **string** | Name of autotest represented by the test result | 
**AutotestGlobalId** | **int64** | Global ID of autotest represented by the test result | 
**TestRunId** | **string** | Unique ID of test run where the test result is located | 
**ConfigurationId** | **string** | Unique ID of configuration which the test result uses | 
**ConfigurationName** | **string** | Name of configuration which the test result uses | 
**Outcome** | Pointer to **NullableString** | Outcome of the test result | [optional] 
**Status** | Pointer to [**NullableTestStatus**](TestStatus.md) |  | [optional] 
**ResultReasons** | [**[]AutoTestResultReasonShort**](AutoTestResultReasonShort.md) | Collection of result reasons which the test result have | 
**Comment** | Pointer to **NullableString** | Comment to the test result | [optional] 
**Date** | **time.Time** | Date when the test result was completed or started or created | 
**CreatedDate** | **time.Time** | Date when the test result has been created | 
**ModifiedDate** | Pointer to **NullableTime** | Date when the test result has been modified | [optional] 
**StartedOn** | Pointer to **NullableTime** | Date when the test result has been started | [optional] 
**CompletedOn** | Pointer to **NullableTime** | Date when the test result has been completed | [optional] 
**Duration** | Pointer to **NullableInt64** | Time which it took to run the test | [optional] 
**Links** | [**[]LinkShort**](LinkShort.md) | Collection of links attached to the test result | 
**Attachments** | [**[]Attachment**](Attachment.md) | Collection of files attached to the test result | 
**RerunCompletedCount** | **int32** | Run count | 

## Methods

### NewTestResultShortResponse

`func NewTestResultShortResponse(id string, name string, autotestGlobalId int64, testRunId string, configurationId string, configurationName string, resultReasons []AutoTestResultReasonShort, date time.Time, createdDate time.Time, links []LinkShort, attachments []Attachment, rerunCompletedCount int32, ) *TestResultShortResponse`

NewTestResultShortResponse instantiates a new TestResultShortResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultShortResponseWithDefaults

`func NewTestResultShortResponseWithDefaults() *TestResultShortResponse`

NewTestResultShortResponseWithDefaults instantiates a new TestResultShortResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestResultShortResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestResultShortResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestResultShortResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TestResultShortResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestResultShortResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestResultShortResponse) SetName(v string)`

SetName sets Name field to given value.


### GetAutotestGlobalId

`func (o *TestResultShortResponse) GetAutotestGlobalId() int64`

GetAutotestGlobalId returns the AutotestGlobalId field if non-nil, zero value otherwise.

### GetAutotestGlobalIdOk

`func (o *TestResultShortResponse) GetAutotestGlobalIdOk() (*int64, bool)`

GetAutotestGlobalIdOk returns a tuple with the AutotestGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutotestGlobalId

`func (o *TestResultShortResponse) SetAutotestGlobalId(v int64)`

SetAutotestGlobalId sets AutotestGlobalId field to given value.


### GetTestRunId

`func (o *TestResultShortResponse) GetTestRunId() string`

GetTestRunId returns the TestRunId field if non-nil, zero value otherwise.

### GetTestRunIdOk

`func (o *TestResultShortResponse) GetTestRunIdOk() (*string, bool)`

GetTestRunIdOk returns a tuple with the TestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunId

`func (o *TestResultShortResponse) SetTestRunId(v string)`

SetTestRunId sets TestRunId field to given value.


### GetConfigurationId

`func (o *TestResultShortResponse) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *TestResultShortResponse) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *TestResultShortResponse) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.


### GetConfigurationName

`func (o *TestResultShortResponse) GetConfigurationName() string`

GetConfigurationName returns the ConfigurationName field if non-nil, zero value otherwise.

### GetConfigurationNameOk

`func (o *TestResultShortResponse) GetConfigurationNameOk() (*string, bool)`

GetConfigurationNameOk returns a tuple with the ConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationName

`func (o *TestResultShortResponse) SetConfigurationName(v string)`

SetConfigurationName sets ConfigurationName field to given value.


### GetOutcome

`func (o *TestResultShortResponse) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *TestResultShortResponse) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *TestResultShortResponse) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *TestResultShortResponse) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *TestResultShortResponse) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *TestResultShortResponse) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetStatus

`func (o *TestResultShortResponse) GetStatus() TestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestResultShortResponse) GetStatusOk() (*TestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestResultShortResponse) SetStatus(v TestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestResultShortResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TestResultShortResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TestResultShortResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetResultReasons

`func (o *TestResultShortResponse) GetResultReasons() []AutoTestResultReasonShort`

GetResultReasons returns the ResultReasons field if non-nil, zero value otherwise.

### GetResultReasonsOk

`func (o *TestResultShortResponse) GetResultReasonsOk() (*[]AutoTestResultReasonShort, bool)`

GetResultReasonsOk returns a tuple with the ResultReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultReasons

`func (o *TestResultShortResponse) SetResultReasons(v []AutoTestResultReasonShort)`

SetResultReasons sets ResultReasons field to given value.


### GetComment

`func (o *TestResultShortResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TestResultShortResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TestResultShortResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TestResultShortResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *TestResultShortResponse) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *TestResultShortResponse) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetDate

`func (o *TestResultShortResponse) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TestResultShortResponse) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TestResultShortResponse) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetCreatedDate

`func (o *TestResultShortResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestResultShortResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestResultShortResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *TestResultShortResponse) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestResultShortResponse) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestResultShortResponse) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestResultShortResponse) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *TestResultShortResponse) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *TestResultShortResponse) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetStartedOn

`func (o *TestResultShortResponse) GetStartedOn() time.Time`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *TestResultShortResponse) GetStartedOnOk() (*time.Time, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *TestResultShortResponse) SetStartedOn(v time.Time)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *TestResultShortResponse) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *TestResultShortResponse) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *TestResultShortResponse) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetCompletedOn

`func (o *TestResultShortResponse) GetCompletedOn() time.Time`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *TestResultShortResponse) GetCompletedOnOk() (*time.Time, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *TestResultShortResponse) SetCompletedOn(v time.Time)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *TestResultShortResponse) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### SetCompletedOnNil

`func (o *TestResultShortResponse) SetCompletedOnNil(b bool)`

 SetCompletedOnNil sets the value for CompletedOn to be an explicit nil

### UnsetCompletedOn
`func (o *TestResultShortResponse) UnsetCompletedOn()`

UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
### GetDuration

`func (o *TestResultShortResponse) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TestResultShortResponse) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TestResultShortResponse) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TestResultShortResponse) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *TestResultShortResponse) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *TestResultShortResponse) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetLinks

`func (o *TestResultShortResponse) GetLinks() []LinkShort`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TestResultShortResponse) GetLinksOk() (*[]LinkShort, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TestResultShortResponse) SetLinks(v []LinkShort)`

SetLinks sets Links field to given value.


### GetAttachments

`func (o *TestResultShortResponse) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestResultShortResponse) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestResultShortResponse) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.


### GetRerunCompletedCount

`func (o *TestResultShortResponse) GetRerunCompletedCount() int32`

GetRerunCompletedCount returns the RerunCompletedCount field if non-nil, zero value otherwise.

### GetRerunCompletedCountOk

`func (o *TestResultShortResponse) GetRerunCompletedCountOk() (*int32, bool)`

GetRerunCompletedCountOk returns a tuple with the RerunCompletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerunCompletedCount

`func (o *TestResultShortResponse) SetRerunCompletedCount(v int32)`

SetRerunCompletedCount sets RerunCompletedCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


