# ApiV2TestRunsSearchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectIds** | Pointer to **[]string** | Specifies a test run project IDs to search for | [optional] 
**Name** | Pointer to **NullableString** | Specifies test run name | [optional] 
**States** | Pointer to [**[]TestRunState**](TestRunState.md) | Specifies a test run states to search for | [optional] 
**StartedDate** | Pointer to [**NullableTestRunFilterModelStartedDate**](TestRunFilterModelStartedDate.md) |  | [optional] 
**CreatedByIds** | Pointer to **[]string** | Specifies a test run creator IDs to search for | [optional] 
**ModifiedByIds** | Pointer to **[]string** | Specifies a test run last editor IDs to search for | [optional] 
**IsDeleted** | Pointer to **NullableBool** | Specifies a test run deleted status to search for | [optional] 
**AutoTestsCount** | Pointer to [**NullableTestRunFilterModelAutoTestsCount**](TestRunFilterModelAutoTestsCount.md) |  | [optional] 
**TestResultsOutcome** | Pointer to [**[]TestResultOutcome**](TestResultOutcome.md) | Specifies test results outcomes | [optional] 
**FailureCategory** | Pointer to [**[]FailureCategoryModel**](FailureCategoryModel.md) | Specifies failure categories | [optional] 
**CompletedDate** | Pointer to [**NullableTestRunFilterModelCompletedDate**](TestRunFilterModelCompletedDate.md) |  | [optional] 

## Methods

### NewApiV2TestRunsSearchPostRequest

`func NewApiV2TestRunsSearchPostRequest() *ApiV2TestRunsSearchPostRequest`

NewApiV2TestRunsSearchPostRequest instantiates a new ApiV2TestRunsSearchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2TestRunsSearchPostRequestWithDefaults

`func NewApiV2TestRunsSearchPostRequestWithDefaults() *ApiV2TestRunsSearchPostRequest`

NewApiV2TestRunsSearchPostRequestWithDefaults instantiates a new ApiV2TestRunsSearchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectIds

`func (o *ApiV2TestRunsSearchPostRequest) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *ApiV2TestRunsSearchPostRequest) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *ApiV2TestRunsSearchPostRequest) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *ApiV2TestRunsSearchPostRequest) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *ApiV2TestRunsSearchPostRequest) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *ApiV2TestRunsSearchPostRequest) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetName

`func (o *ApiV2TestRunsSearchPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2TestRunsSearchPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2TestRunsSearchPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV2TestRunsSearchPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApiV2TestRunsSearchPostRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApiV2TestRunsSearchPostRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStates

`func (o *ApiV2TestRunsSearchPostRequest) GetStates() []TestRunState`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *ApiV2TestRunsSearchPostRequest) GetStatesOk() (*[]TestRunState, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *ApiV2TestRunsSearchPostRequest) SetStates(v []TestRunState)`

SetStates sets States field to given value.

### HasStates

`func (o *ApiV2TestRunsSearchPostRequest) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStatesNil

`func (o *ApiV2TestRunsSearchPostRequest) SetStatesNil(b bool)`

 SetStatesNil sets the value for States to be an explicit nil

### UnsetStates
`func (o *ApiV2TestRunsSearchPostRequest) UnsetStates()`

UnsetStates ensures that no value is present for States, not even an explicit nil
### GetStartedDate

`func (o *ApiV2TestRunsSearchPostRequest) GetStartedDate() TestRunFilterModelStartedDate`

GetStartedDate returns the StartedDate field if non-nil, zero value otherwise.

### GetStartedDateOk

`func (o *ApiV2TestRunsSearchPostRequest) GetStartedDateOk() (*TestRunFilterModelStartedDate, bool)`

GetStartedDateOk returns a tuple with the StartedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedDate

`func (o *ApiV2TestRunsSearchPostRequest) SetStartedDate(v TestRunFilterModelStartedDate)`

SetStartedDate sets StartedDate field to given value.

### HasStartedDate

`func (o *ApiV2TestRunsSearchPostRequest) HasStartedDate() bool`

HasStartedDate returns a boolean if a field has been set.

### SetStartedDateNil

`func (o *ApiV2TestRunsSearchPostRequest) SetStartedDateNil(b bool)`

 SetStartedDateNil sets the value for StartedDate to be an explicit nil

### UnsetStartedDate
`func (o *ApiV2TestRunsSearchPostRequest) UnsetStartedDate()`

UnsetStartedDate ensures that no value is present for StartedDate, not even an explicit nil
### GetCreatedByIds

`func (o *ApiV2TestRunsSearchPostRequest) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *ApiV2TestRunsSearchPostRequest) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *ApiV2TestRunsSearchPostRequest) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *ApiV2TestRunsSearchPostRequest) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *ApiV2TestRunsSearchPostRequest) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *ApiV2TestRunsSearchPostRequest) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedByIds

`func (o *ApiV2TestRunsSearchPostRequest) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *ApiV2TestRunsSearchPostRequest) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *ApiV2TestRunsSearchPostRequest) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *ApiV2TestRunsSearchPostRequest) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *ApiV2TestRunsSearchPostRequest) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *ApiV2TestRunsSearchPostRequest) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil
### GetIsDeleted

`func (o *ApiV2TestRunsSearchPostRequest) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ApiV2TestRunsSearchPostRequest) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ApiV2TestRunsSearchPostRequest) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ApiV2TestRunsSearchPostRequest) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *ApiV2TestRunsSearchPostRequest) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *ApiV2TestRunsSearchPostRequest) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetAutoTestsCount

`func (o *ApiV2TestRunsSearchPostRequest) GetAutoTestsCount() TestRunFilterModelAutoTestsCount`

GetAutoTestsCount returns the AutoTestsCount field if non-nil, zero value otherwise.

### GetAutoTestsCountOk

`func (o *ApiV2TestRunsSearchPostRequest) GetAutoTestsCountOk() (*TestRunFilterModelAutoTestsCount, bool)`

GetAutoTestsCountOk returns a tuple with the AutoTestsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestsCount

`func (o *ApiV2TestRunsSearchPostRequest) SetAutoTestsCount(v TestRunFilterModelAutoTestsCount)`

SetAutoTestsCount sets AutoTestsCount field to given value.

### HasAutoTestsCount

`func (o *ApiV2TestRunsSearchPostRequest) HasAutoTestsCount() bool`

HasAutoTestsCount returns a boolean if a field has been set.

### SetAutoTestsCountNil

`func (o *ApiV2TestRunsSearchPostRequest) SetAutoTestsCountNil(b bool)`

 SetAutoTestsCountNil sets the value for AutoTestsCount to be an explicit nil

### UnsetAutoTestsCount
`func (o *ApiV2TestRunsSearchPostRequest) UnsetAutoTestsCount()`

UnsetAutoTestsCount ensures that no value is present for AutoTestsCount, not even an explicit nil
### GetTestResultsOutcome

`func (o *ApiV2TestRunsSearchPostRequest) GetTestResultsOutcome() []TestResultOutcome`

GetTestResultsOutcome returns the TestResultsOutcome field if non-nil, zero value otherwise.

### GetTestResultsOutcomeOk

`func (o *ApiV2TestRunsSearchPostRequest) GetTestResultsOutcomeOk() (*[]TestResultOutcome, bool)`

GetTestResultsOutcomeOk returns a tuple with the TestResultsOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResultsOutcome

`func (o *ApiV2TestRunsSearchPostRequest) SetTestResultsOutcome(v []TestResultOutcome)`

SetTestResultsOutcome sets TestResultsOutcome field to given value.

### HasTestResultsOutcome

`func (o *ApiV2TestRunsSearchPostRequest) HasTestResultsOutcome() bool`

HasTestResultsOutcome returns a boolean if a field has been set.

### SetTestResultsOutcomeNil

`func (o *ApiV2TestRunsSearchPostRequest) SetTestResultsOutcomeNil(b bool)`

 SetTestResultsOutcomeNil sets the value for TestResultsOutcome to be an explicit nil

### UnsetTestResultsOutcome
`func (o *ApiV2TestRunsSearchPostRequest) UnsetTestResultsOutcome()`

UnsetTestResultsOutcome ensures that no value is present for TestResultsOutcome, not even an explicit nil
### GetFailureCategory

`func (o *ApiV2TestRunsSearchPostRequest) GetFailureCategory() []FailureCategoryModel`

GetFailureCategory returns the FailureCategory field if non-nil, zero value otherwise.

### GetFailureCategoryOk

`func (o *ApiV2TestRunsSearchPostRequest) GetFailureCategoryOk() (*[]FailureCategoryModel, bool)`

GetFailureCategoryOk returns a tuple with the FailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategory

`func (o *ApiV2TestRunsSearchPostRequest) SetFailureCategory(v []FailureCategoryModel)`

SetFailureCategory sets FailureCategory field to given value.

### HasFailureCategory

`func (o *ApiV2TestRunsSearchPostRequest) HasFailureCategory() bool`

HasFailureCategory returns a boolean if a field has been set.

### SetFailureCategoryNil

`func (o *ApiV2TestRunsSearchPostRequest) SetFailureCategoryNil(b bool)`

 SetFailureCategoryNil sets the value for FailureCategory to be an explicit nil

### UnsetFailureCategory
`func (o *ApiV2TestRunsSearchPostRequest) UnsetFailureCategory()`

UnsetFailureCategory ensures that no value is present for FailureCategory, not even an explicit nil
### GetCompletedDate

`func (o *ApiV2TestRunsSearchPostRequest) GetCompletedDate() TestRunFilterModelCompletedDate`

GetCompletedDate returns the CompletedDate field if non-nil, zero value otherwise.

### GetCompletedDateOk

`func (o *ApiV2TestRunsSearchPostRequest) GetCompletedDateOk() (*TestRunFilterModelCompletedDate, bool)`

GetCompletedDateOk returns a tuple with the CompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDate

`func (o *ApiV2TestRunsSearchPostRequest) SetCompletedDate(v TestRunFilterModelCompletedDate)`

SetCompletedDate sets CompletedDate field to given value.

### HasCompletedDate

`func (o *ApiV2TestRunsSearchPostRequest) HasCompletedDate() bool`

HasCompletedDate returns a boolean if a field has been set.

### SetCompletedDateNil

`func (o *ApiV2TestRunsSearchPostRequest) SetCompletedDateNil(b bool)`

 SetCompletedDateNil sets the value for CompletedDate to be an explicit nil

### UnsetCompletedDate
`func (o *ApiV2TestRunsSearchPostRequest) UnsetCompletedDate()`

UnsetCompletedDate ensures that no value is present for CompletedDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


