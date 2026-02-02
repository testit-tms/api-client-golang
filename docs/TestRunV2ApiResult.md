# TestRunV2ApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Test run unique identifier | 
**Name** | **string** | Test run name | 
**Description** | Pointer to **NullableString** | Test run description | [optional] 
**LaunchSource** | Pointer to **NullableString** | Test run launch source              Once launch source is specified it cannot be updated. | [optional] 
**StartedOn** | Pointer to **NullableTime** | Date and time of test run start | [optional] 
**CompletedOn** | Pointer to **NullableTime** | Date and time of test run end | [optional] 
**StateName** | [**TestRunState**](TestRunState.md) | Test run state | 
**Status** | [**TestStatusApiResult**](TestStatusApiResult.md) | Test run status | 
**ProjectId** | **string** | Project unique identifier              This property is used to link test run with project. | 
**TestPlanId** | Pointer to **NullableString** | Test plan unique identifier              This property is used to link test run with test plan. | [optional] 
**TestResults** | Pointer to [**[]TestResultV2GetModel**](TestResultV2GetModel.md) | Enumeration of test results related to test run | [optional] 
**CreatedDate** | **time.Time** | Date and time of test run creation | 
**ModifiedDate** | Pointer to **NullableTime** | Date and time of last test run  modification | [optional] 
**CreatedById** | **string** | Unique identifier of user who created test run | 
**ModifiedById** | Pointer to **NullableString** | Unique identifier of user who applied last test run  modification | [optional] 
**CreatedByUserName** | Pointer to **NullableString** | Username of user who created test run | [optional] 
**Attachments** | [**[]AttachmentApiResult**](AttachmentApiResult.md) | Collection of attachments related to the test run | 
**Links** | [**[]LinkApiResult**](LinkApiResult.md) | Collection of links related to the test run | 
**CustomParameters** | Pointer to **map[string]string** | Customers test run parameters | [optional] 
**Webhooks** | [**[]NamedEntityApiModel**](NamedEntityApiModel.md) | Enabled webhooks | 
**RunCount** | **int32** | Run count | 

## Methods

### NewTestRunV2ApiResult

`func NewTestRunV2ApiResult(id string, name string, stateName TestRunState, status TestStatusApiResult, projectId string, createdDate time.Time, createdById string, attachments []AttachmentApiResult, links []LinkApiResult, webhooks []NamedEntityApiModel, runCount int32, ) *TestRunV2ApiResult`

NewTestRunV2ApiResult instantiates a new TestRunV2ApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunV2ApiResultWithDefaults

`func NewTestRunV2ApiResultWithDefaults() *TestRunV2ApiResult`

NewTestRunV2ApiResultWithDefaults instantiates a new TestRunV2ApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestRunV2ApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestRunV2ApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestRunV2ApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TestRunV2ApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestRunV2ApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestRunV2ApiResult) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TestRunV2ApiResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestRunV2ApiResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestRunV2ApiResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestRunV2ApiResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestRunV2ApiResult) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestRunV2ApiResult) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLaunchSource

`func (o *TestRunV2ApiResult) GetLaunchSource() string`

GetLaunchSource returns the LaunchSource field if non-nil, zero value otherwise.

### GetLaunchSourceOk

`func (o *TestRunV2ApiResult) GetLaunchSourceOk() (*string, bool)`

GetLaunchSourceOk returns a tuple with the LaunchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSource

`func (o *TestRunV2ApiResult) SetLaunchSource(v string)`

SetLaunchSource sets LaunchSource field to given value.

### HasLaunchSource

`func (o *TestRunV2ApiResult) HasLaunchSource() bool`

HasLaunchSource returns a boolean if a field has been set.

### SetLaunchSourceNil

`func (o *TestRunV2ApiResult) SetLaunchSourceNil(b bool)`

 SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil

### UnsetLaunchSource
`func (o *TestRunV2ApiResult) UnsetLaunchSource()`

UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
### GetStartedOn

`func (o *TestRunV2ApiResult) GetStartedOn() time.Time`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *TestRunV2ApiResult) GetStartedOnOk() (*time.Time, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *TestRunV2ApiResult) SetStartedOn(v time.Time)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *TestRunV2ApiResult) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *TestRunV2ApiResult) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *TestRunV2ApiResult) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetCompletedOn

`func (o *TestRunV2ApiResult) GetCompletedOn() time.Time`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *TestRunV2ApiResult) GetCompletedOnOk() (*time.Time, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *TestRunV2ApiResult) SetCompletedOn(v time.Time)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *TestRunV2ApiResult) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### SetCompletedOnNil

`func (o *TestRunV2ApiResult) SetCompletedOnNil(b bool)`

 SetCompletedOnNil sets the value for CompletedOn to be an explicit nil

### UnsetCompletedOn
`func (o *TestRunV2ApiResult) UnsetCompletedOn()`

UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
### GetStateName

`func (o *TestRunV2ApiResult) GetStateName() TestRunState`

GetStateName returns the StateName field if non-nil, zero value otherwise.

### GetStateNameOk

`func (o *TestRunV2ApiResult) GetStateNameOk() (*TestRunState, bool)`

GetStateNameOk returns a tuple with the StateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateName

`func (o *TestRunV2ApiResult) SetStateName(v TestRunState)`

SetStateName sets StateName field to given value.


### GetStatus

`func (o *TestRunV2ApiResult) GetStatus() TestStatusApiResult`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestRunV2ApiResult) GetStatusOk() (*TestStatusApiResult, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestRunV2ApiResult) SetStatus(v TestStatusApiResult)`

SetStatus sets Status field to given value.


### GetProjectId

`func (o *TestRunV2ApiResult) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TestRunV2ApiResult) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TestRunV2ApiResult) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetTestPlanId

`func (o *TestRunV2ApiResult) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *TestRunV2ApiResult) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *TestRunV2ApiResult) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.

### HasTestPlanId

`func (o *TestRunV2ApiResult) HasTestPlanId() bool`

HasTestPlanId returns a boolean if a field has been set.

### SetTestPlanIdNil

`func (o *TestRunV2ApiResult) SetTestPlanIdNil(b bool)`

 SetTestPlanIdNil sets the value for TestPlanId to be an explicit nil

### UnsetTestPlanId
`func (o *TestRunV2ApiResult) UnsetTestPlanId()`

UnsetTestPlanId ensures that no value is present for TestPlanId, not even an explicit nil
### GetTestResults

`func (o *TestRunV2ApiResult) GetTestResults() []TestResultV2GetModel`

GetTestResults returns the TestResults field if non-nil, zero value otherwise.

### GetTestResultsOk

`func (o *TestRunV2ApiResult) GetTestResultsOk() (*[]TestResultV2GetModel, bool)`

GetTestResultsOk returns a tuple with the TestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResults

`func (o *TestRunV2ApiResult) SetTestResults(v []TestResultV2GetModel)`

SetTestResults sets TestResults field to given value.

### HasTestResults

`func (o *TestRunV2ApiResult) HasTestResults() bool`

HasTestResults returns a boolean if a field has been set.

### SetTestResultsNil

`func (o *TestRunV2ApiResult) SetTestResultsNil(b bool)`

 SetTestResultsNil sets the value for TestResults to be an explicit nil

### UnsetTestResults
`func (o *TestRunV2ApiResult) UnsetTestResults()`

UnsetTestResults ensures that no value is present for TestResults, not even an explicit nil
### GetCreatedDate

`func (o *TestRunV2ApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestRunV2ApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestRunV2ApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *TestRunV2ApiResult) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestRunV2ApiResult) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestRunV2ApiResult) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestRunV2ApiResult) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *TestRunV2ApiResult) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *TestRunV2ApiResult) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *TestRunV2ApiResult) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TestRunV2ApiResult) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TestRunV2ApiResult) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedById

`func (o *TestRunV2ApiResult) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *TestRunV2ApiResult) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *TestRunV2ApiResult) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *TestRunV2ApiResult) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *TestRunV2ApiResult) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *TestRunV2ApiResult) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetCreatedByUserName

`func (o *TestRunV2ApiResult) GetCreatedByUserName() string`

GetCreatedByUserName returns the CreatedByUserName field if non-nil, zero value otherwise.

### GetCreatedByUserNameOk

`func (o *TestRunV2ApiResult) GetCreatedByUserNameOk() (*string, bool)`

GetCreatedByUserNameOk returns a tuple with the CreatedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserName

`func (o *TestRunV2ApiResult) SetCreatedByUserName(v string)`

SetCreatedByUserName sets CreatedByUserName field to given value.

### HasCreatedByUserName

`func (o *TestRunV2ApiResult) HasCreatedByUserName() bool`

HasCreatedByUserName returns a boolean if a field has been set.

### SetCreatedByUserNameNil

`func (o *TestRunV2ApiResult) SetCreatedByUserNameNil(b bool)`

 SetCreatedByUserNameNil sets the value for CreatedByUserName to be an explicit nil

### UnsetCreatedByUserName
`func (o *TestRunV2ApiResult) UnsetCreatedByUserName()`

UnsetCreatedByUserName ensures that no value is present for CreatedByUserName, not even an explicit nil
### GetAttachments

`func (o *TestRunV2ApiResult) GetAttachments() []AttachmentApiResult`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestRunV2ApiResult) GetAttachmentsOk() (*[]AttachmentApiResult, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestRunV2ApiResult) SetAttachments(v []AttachmentApiResult)`

SetAttachments sets Attachments field to given value.


### GetLinks

`func (o *TestRunV2ApiResult) GetLinks() []LinkApiResult`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TestRunV2ApiResult) GetLinksOk() (*[]LinkApiResult, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TestRunV2ApiResult) SetLinks(v []LinkApiResult)`

SetLinks sets Links field to given value.


### GetCustomParameters

`func (o *TestRunV2ApiResult) GetCustomParameters() map[string]string`

GetCustomParameters returns the CustomParameters field if non-nil, zero value otherwise.

### GetCustomParametersOk

`func (o *TestRunV2ApiResult) GetCustomParametersOk() (*map[string]string, bool)`

GetCustomParametersOk returns a tuple with the CustomParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParameters

`func (o *TestRunV2ApiResult) SetCustomParameters(v map[string]string)`

SetCustomParameters sets CustomParameters field to given value.

### HasCustomParameters

`func (o *TestRunV2ApiResult) HasCustomParameters() bool`

HasCustomParameters returns a boolean if a field has been set.

### SetCustomParametersNil

`func (o *TestRunV2ApiResult) SetCustomParametersNil(b bool)`

 SetCustomParametersNil sets the value for CustomParameters to be an explicit nil

### UnsetCustomParameters
`func (o *TestRunV2ApiResult) UnsetCustomParameters()`

UnsetCustomParameters ensures that no value is present for CustomParameters, not even an explicit nil
### GetWebhooks

`func (o *TestRunV2ApiResult) GetWebhooks() []NamedEntityApiModel`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *TestRunV2ApiResult) GetWebhooksOk() (*[]NamedEntityApiModel, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *TestRunV2ApiResult) SetWebhooks(v []NamedEntityApiModel)`

SetWebhooks sets Webhooks field to given value.


### GetRunCount

`func (o *TestRunV2ApiResult) GetRunCount() int32`

GetRunCount returns the RunCount field if non-nil, zero value otherwise.

### GetRunCountOk

`func (o *TestRunV2ApiResult) GetRunCountOk() (*int32, bool)`

GetRunCountOk returns a tuple with the RunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCount

`func (o *TestRunV2ApiResult) SetRunCount(v int32)`

SetRunCount sets RunCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


