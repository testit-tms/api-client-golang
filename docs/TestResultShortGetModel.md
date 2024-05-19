# TestResultShortGetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the test result | 
**Name** | **string** | Name of autotest represented by the test result | 
**AutotestGlobalId** | **int64** | Global ID of autotest represented by the test result | 
**TestRunId** | **string** | Unique ID of test run where the test result is located | 
**ConfigurationId** | **string** | Unique ID of configuration which the test result uses | 
**ConfigurationName** | **string** | Name of configuration which the test result uses | 
**Outcome** | **string** | Outcome of the test result | 
**ResultReasons** | [**[]AutotestResultReasonSubGetModel**](AutotestResultReasonSubGetModel.md) | Collection of result reasons which the test result have | 
**Comment** | Pointer to **NullableString** | Comment to the test result | [optional] 
**Date** | **time.Time** | Date when the test result was completed or started or created | 
**CreatedDate** | **time.Time** | Date when the test result has been created | 
**ModifiedDate** | Pointer to **NullableTime** | Date when the test result has been modified | [optional] 
**StartedOn** | Pointer to **NullableTime** | Date when the test result has been started | [optional] 
**CompletedOn** | Pointer to **NullableTime** | Date when the test result has been completed | [optional] 
**Duration** | Pointer to **NullableInt64** | Time which it took to run the test | [optional] 
**Links** | [**[]LinkSubGetModel**](LinkSubGetModel.md) | Collection of links attached to the test result | 
**Attachments** | [**[]AttachmentModel**](AttachmentModel.md) | Collection of files attached to the test result | 

## Methods

### NewTestResultShortGetModel

`func NewTestResultShortGetModel(id string, name string, autotestGlobalId int64, testRunId string, configurationId string, configurationName string, outcome string, resultReasons []AutotestResultReasonSubGetModel, date time.Time, createdDate time.Time, links []LinkSubGetModel, attachments []AttachmentModel, ) *TestResultShortGetModel`

NewTestResultShortGetModel instantiates a new TestResultShortGetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultShortGetModelWithDefaults

`func NewTestResultShortGetModelWithDefaults() *TestResultShortGetModel`

NewTestResultShortGetModelWithDefaults instantiates a new TestResultShortGetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestResultShortGetModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestResultShortGetModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestResultShortGetModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TestResultShortGetModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestResultShortGetModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestResultShortGetModel) SetName(v string)`

SetName sets Name field to given value.


### GetAutotestGlobalId

`func (o *TestResultShortGetModel) GetAutotestGlobalId() int64`

GetAutotestGlobalId returns the AutotestGlobalId field if non-nil, zero value otherwise.

### GetAutotestGlobalIdOk

`func (o *TestResultShortGetModel) GetAutotestGlobalIdOk() (*int64, bool)`

GetAutotestGlobalIdOk returns a tuple with the AutotestGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutotestGlobalId

`func (o *TestResultShortGetModel) SetAutotestGlobalId(v int64)`

SetAutotestGlobalId sets AutotestGlobalId field to given value.


### GetTestRunId

`func (o *TestResultShortGetModel) GetTestRunId() string`

GetTestRunId returns the TestRunId field if non-nil, zero value otherwise.

### GetTestRunIdOk

`func (o *TestResultShortGetModel) GetTestRunIdOk() (*string, bool)`

GetTestRunIdOk returns a tuple with the TestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunId

`func (o *TestResultShortGetModel) SetTestRunId(v string)`

SetTestRunId sets TestRunId field to given value.


### GetConfigurationId

`func (o *TestResultShortGetModel) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *TestResultShortGetModel) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *TestResultShortGetModel) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.


### GetConfigurationName

`func (o *TestResultShortGetModel) GetConfigurationName() string`

GetConfigurationName returns the ConfigurationName field if non-nil, zero value otherwise.

### GetConfigurationNameOk

`func (o *TestResultShortGetModel) GetConfigurationNameOk() (*string, bool)`

GetConfigurationNameOk returns a tuple with the ConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationName

`func (o *TestResultShortGetModel) SetConfigurationName(v string)`

SetConfigurationName sets ConfigurationName field to given value.


### GetOutcome

`func (o *TestResultShortGetModel) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *TestResultShortGetModel) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *TestResultShortGetModel) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### GetResultReasons

`func (o *TestResultShortGetModel) GetResultReasons() []AutotestResultReasonSubGetModel`

GetResultReasons returns the ResultReasons field if non-nil, zero value otherwise.

### GetResultReasonsOk

`func (o *TestResultShortGetModel) GetResultReasonsOk() (*[]AutotestResultReasonSubGetModel, bool)`

GetResultReasonsOk returns a tuple with the ResultReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultReasons

`func (o *TestResultShortGetModel) SetResultReasons(v []AutotestResultReasonSubGetModel)`

SetResultReasons sets ResultReasons field to given value.


### GetComment

`func (o *TestResultShortGetModel) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TestResultShortGetModel) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TestResultShortGetModel) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TestResultShortGetModel) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *TestResultShortGetModel) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *TestResultShortGetModel) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetDate

`func (o *TestResultShortGetModel) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TestResultShortGetModel) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TestResultShortGetModel) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetCreatedDate

`func (o *TestResultShortGetModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestResultShortGetModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestResultShortGetModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *TestResultShortGetModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestResultShortGetModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestResultShortGetModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestResultShortGetModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *TestResultShortGetModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *TestResultShortGetModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetStartedOn

`func (o *TestResultShortGetModel) GetStartedOn() time.Time`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *TestResultShortGetModel) GetStartedOnOk() (*time.Time, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *TestResultShortGetModel) SetStartedOn(v time.Time)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *TestResultShortGetModel) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *TestResultShortGetModel) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *TestResultShortGetModel) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetCompletedOn

`func (o *TestResultShortGetModel) GetCompletedOn() time.Time`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *TestResultShortGetModel) GetCompletedOnOk() (*time.Time, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *TestResultShortGetModel) SetCompletedOn(v time.Time)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *TestResultShortGetModel) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### SetCompletedOnNil

`func (o *TestResultShortGetModel) SetCompletedOnNil(b bool)`

 SetCompletedOnNil sets the value for CompletedOn to be an explicit nil

### UnsetCompletedOn
`func (o *TestResultShortGetModel) UnsetCompletedOn()`

UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
### GetDuration

`func (o *TestResultShortGetModel) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TestResultShortGetModel) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TestResultShortGetModel) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TestResultShortGetModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *TestResultShortGetModel) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *TestResultShortGetModel) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetLinks

`func (o *TestResultShortGetModel) GetLinks() []LinkSubGetModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TestResultShortGetModel) GetLinksOk() (*[]LinkSubGetModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TestResultShortGetModel) SetLinks(v []LinkSubGetModel)`

SetLinks sets Links field to given value.


### GetAttachments

`func (o *TestResultShortGetModel) GetAttachments() []AttachmentModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestResultShortGetModel) GetAttachmentsOk() (*[]AttachmentModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestResultShortGetModel) SetAttachments(v []AttachmentModel)`

SetAttachments sets Attachments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


