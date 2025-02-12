# TestRunFilterApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectIds** | Pointer to **[]string** | Specifies a test run project IDs to search for | [optional] 
**Name** | Pointer to **NullableString** | Specifies test run name | [optional] 
**States** | Pointer to [**[]TestRunState**](TestRunState.md) | Specifies a test run states to search for | [optional] 
**StatusCodes** | Pointer to **[]string** | Specifies a test run status codes to search for | [optional] 
**CreatedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) | Specifies a test run range of created date to search for | [optional] 
**StartedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) | Specifies a test run range of started date to search for | [optional] 
**CreatedByIds** | Pointer to **[]string** | Specifies a test run creator IDs to search for | [optional] 
**ModifiedByIds** | Pointer to **[]string** | Specifies a test run last editor IDs to search for | [optional] 
**IsDeleted** | Pointer to **NullableBool** | Specifies a test run deleted status to search for | [optional] 
**AutoTestsCount** | Pointer to [**NullableInt32RangeSelectorModel**](Int32RangeSelectorModel.md) | Number of autoTests run in the test run | [optional] 
**TestResultsOutcome** | Pointer to [**[]TestResultOutcome**](TestResultOutcome.md) | Specifies test results outcomes | [optional] 
**TestResultsStatusCodes** | Pointer to **[]string** | Specifies test results status codes | [optional] 
**FailureCategory** | Pointer to [**[]FailureCategory**](FailureCategory.md) | Specifies failure categories | [optional] 
**CompletedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) | Specifies a test run range of completed date to search for | [optional] 
**TestResultsConfigurationIds** | Pointer to **[]string** | Specifies a test result configuration IDs to search for | [optional] 

## Methods

### NewTestRunFilterApiModel

`func NewTestRunFilterApiModel() *TestRunFilterApiModel`

NewTestRunFilterApiModel instantiates a new TestRunFilterApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunFilterApiModelWithDefaults

`func NewTestRunFilterApiModelWithDefaults() *TestRunFilterApiModel`

NewTestRunFilterApiModelWithDefaults instantiates a new TestRunFilterApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectIds

`func (o *TestRunFilterApiModel) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *TestRunFilterApiModel) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *TestRunFilterApiModel) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *TestRunFilterApiModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *TestRunFilterApiModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *TestRunFilterApiModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetName

`func (o *TestRunFilterApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestRunFilterApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestRunFilterApiModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestRunFilterApiModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestRunFilterApiModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestRunFilterApiModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStates

`func (o *TestRunFilterApiModel) GetStates() []TestRunState`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *TestRunFilterApiModel) GetStatesOk() (*[]TestRunState, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *TestRunFilterApiModel) SetStates(v []TestRunState)`

SetStates sets States field to given value.

### HasStates

`func (o *TestRunFilterApiModel) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStatesNil

`func (o *TestRunFilterApiModel) SetStatesNil(b bool)`

 SetStatesNil sets the value for States to be an explicit nil

### UnsetStates
`func (o *TestRunFilterApiModel) UnsetStates()`

UnsetStates ensures that no value is present for States, not even an explicit nil
### GetStatusCodes

`func (o *TestRunFilterApiModel) GetStatusCodes() []string`

GetStatusCodes returns the StatusCodes field if non-nil, zero value otherwise.

### GetStatusCodesOk

`func (o *TestRunFilterApiModel) GetStatusCodesOk() (*[]string, bool)`

GetStatusCodesOk returns a tuple with the StatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodes

`func (o *TestRunFilterApiModel) SetStatusCodes(v []string)`

SetStatusCodes sets StatusCodes field to given value.

### HasStatusCodes

`func (o *TestRunFilterApiModel) HasStatusCodes() bool`

HasStatusCodes returns a boolean if a field has been set.

### SetStatusCodesNil

`func (o *TestRunFilterApiModel) SetStatusCodesNil(b bool)`

 SetStatusCodesNil sets the value for StatusCodes to be an explicit nil

### UnsetStatusCodes
`func (o *TestRunFilterApiModel) UnsetStatusCodes()`

UnsetStatusCodes ensures that no value is present for StatusCodes, not even an explicit nil
### GetCreatedDate

`func (o *TestRunFilterApiModel) GetCreatedDate() DateTimeRangeSelectorModel`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestRunFilterApiModel) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestRunFilterApiModel) SetCreatedDate(v DateTimeRangeSelectorModel)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TestRunFilterApiModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *TestRunFilterApiModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *TestRunFilterApiModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetStartedDate

`func (o *TestRunFilterApiModel) GetStartedDate() DateTimeRangeSelectorModel`

GetStartedDate returns the StartedDate field if non-nil, zero value otherwise.

### GetStartedDateOk

`func (o *TestRunFilterApiModel) GetStartedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetStartedDateOk returns a tuple with the StartedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedDate

`func (o *TestRunFilterApiModel) SetStartedDate(v DateTimeRangeSelectorModel)`

SetStartedDate sets StartedDate field to given value.

### HasStartedDate

`func (o *TestRunFilterApiModel) HasStartedDate() bool`

HasStartedDate returns a boolean if a field has been set.

### SetStartedDateNil

`func (o *TestRunFilterApiModel) SetStartedDateNil(b bool)`

 SetStartedDateNil sets the value for StartedDate to be an explicit nil

### UnsetStartedDate
`func (o *TestRunFilterApiModel) UnsetStartedDate()`

UnsetStartedDate ensures that no value is present for StartedDate, not even an explicit nil
### GetCreatedByIds

`func (o *TestRunFilterApiModel) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *TestRunFilterApiModel) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *TestRunFilterApiModel) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *TestRunFilterApiModel) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *TestRunFilterApiModel) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *TestRunFilterApiModel) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedByIds

`func (o *TestRunFilterApiModel) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *TestRunFilterApiModel) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *TestRunFilterApiModel) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *TestRunFilterApiModel) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *TestRunFilterApiModel) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *TestRunFilterApiModel) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil
### GetIsDeleted

`func (o *TestRunFilterApiModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TestRunFilterApiModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TestRunFilterApiModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *TestRunFilterApiModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *TestRunFilterApiModel) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *TestRunFilterApiModel) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetAutoTestsCount

`func (o *TestRunFilterApiModel) GetAutoTestsCount() Int32RangeSelectorModel`

GetAutoTestsCount returns the AutoTestsCount field if non-nil, zero value otherwise.

### GetAutoTestsCountOk

`func (o *TestRunFilterApiModel) GetAutoTestsCountOk() (*Int32RangeSelectorModel, bool)`

GetAutoTestsCountOk returns a tuple with the AutoTestsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestsCount

`func (o *TestRunFilterApiModel) SetAutoTestsCount(v Int32RangeSelectorModel)`

SetAutoTestsCount sets AutoTestsCount field to given value.

### HasAutoTestsCount

`func (o *TestRunFilterApiModel) HasAutoTestsCount() bool`

HasAutoTestsCount returns a boolean if a field has been set.

### SetAutoTestsCountNil

`func (o *TestRunFilterApiModel) SetAutoTestsCountNil(b bool)`

 SetAutoTestsCountNil sets the value for AutoTestsCount to be an explicit nil

### UnsetAutoTestsCount
`func (o *TestRunFilterApiModel) UnsetAutoTestsCount()`

UnsetAutoTestsCount ensures that no value is present for AutoTestsCount, not even an explicit nil
### GetTestResultsOutcome

`func (o *TestRunFilterApiModel) GetTestResultsOutcome() []TestResultOutcome`

GetTestResultsOutcome returns the TestResultsOutcome field if non-nil, zero value otherwise.

### GetTestResultsOutcomeOk

`func (o *TestRunFilterApiModel) GetTestResultsOutcomeOk() (*[]TestResultOutcome, bool)`

GetTestResultsOutcomeOk returns a tuple with the TestResultsOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResultsOutcome

`func (o *TestRunFilterApiModel) SetTestResultsOutcome(v []TestResultOutcome)`

SetTestResultsOutcome sets TestResultsOutcome field to given value.

### HasTestResultsOutcome

`func (o *TestRunFilterApiModel) HasTestResultsOutcome() bool`

HasTestResultsOutcome returns a boolean if a field has been set.

### SetTestResultsOutcomeNil

`func (o *TestRunFilterApiModel) SetTestResultsOutcomeNil(b bool)`

 SetTestResultsOutcomeNil sets the value for TestResultsOutcome to be an explicit nil

### UnsetTestResultsOutcome
`func (o *TestRunFilterApiModel) UnsetTestResultsOutcome()`

UnsetTestResultsOutcome ensures that no value is present for TestResultsOutcome, not even an explicit nil
### GetTestResultsStatusCodes

`func (o *TestRunFilterApiModel) GetTestResultsStatusCodes() []string`

GetTestResultsStatusCodes returns the TestResultsStatusCodes field if non-nil, zero value otherwise.

### GetTestResultsStatusCodesOk

`func (o *TestRunFilterApiModel) GetTestResultsStatusCodesOk() (*[]string, bool)`

GetTestResultsStatusCodesOk returns a tuple with the TestResultsStatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResultsStatusCodes

`func (o *TestRunFilterApiModel) SetTestResultsStatusCodes(v []string)`

SetTestResultsStatusCodes sets TestResultsStatusCodes field to given value.

### HasTestResultsStatusCodes

`func (o *TestRunFilterApiModel) HasTestResultsStatusCodes() bool`

HasTestResultsStatusCodes returns a boolean if a field has been set.

### SetTestResultsStatusCodesNil

`func (o *TestRunFilterApiModel) SetTestResultsStatusCodesNil(b bool)`

 SetTestResultsStatusCodesNil sets the value for TestResultsStatusCodes to be an explicit nil

### UnsetTestResultsStatusCodes
`func (o *TestRunFilterApiModel) UnsetTestResultsStatusCodes()`

UnsetTestResultsStatusCodes ensures that no value is present for TestResultsStatusCodes, not even an explicit nil
### GetFailureCategory

`func (o *TestRunFilterApiModel) GetFailureCategory() []FailureCategory`

GetFailureCategory returns the FailureCategory field if non-nil, zero value otherwise.

### GetFailureCategoryOk

`func (o *TestRunFilterApiModel) GetFailureCategoryOk() (*[]FailureCategory, bool)`

GetFailureCategoryOk returns a tuple with the FailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategory

`func (o *TestRunFilterApiModel) SetFailureCategory(v []FailureCategory)`

SetFailureCategory sets FailureCategory field to given value.

### HasFailureCategory

`func (o *TestRunFilterApiModel) HasFailureCategory() bool`

HasFailureCategory returns a boolean if a field has been set.

### SetFailureCategoryNil

`func (o *TestRunFilterApiModel) SetFailureCategoryNil(b bool)`

 SetFailureCategoryNil sets the value for FailureCategory to be an explicit nil

### UnsetFailureCategory
`func (o *TestRunFilterApiModel) UnsetFailureCategory()`

UnsetFailureCategory ensures that no value is present for FailureCategory, not even an explicit nil
### GetCompletedDate

`func (o *TestRunFilterApiModel) GetCompletedDate() DateTimeRangeSelectorModel`

GetCompletedDate returns the CompletedDate field if non-nil, zero value otherwise.

### GetCompletedDateOk

`func (o *TestRunFilterApiModel) GetCompletedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCompletedDateOk returns a tuple with the CompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDate

`func (o *TestRunFilterApiModel) SetCompletedDate(v DateTimeRangeSelectorModel)`

SetCompletedDate sets CompletedDate field to given value.

### HasCompletedDate

`func (o *TestRunFilterApiModel) HasCompletedDate() bool`

HasCompletedDate returns a boolean if a field has been set.

### SetCompletedDateNil

`func (o *TestRunFilterApiModel) SetCompletedDateNil(b bool)`

 SetCompletedDateNil sets the value for CompletedDate to be an explicit nil

### UnsetCompletedDate
`func (o *TestRunFilterApiModel) UnsetCompletedDate()`

UnsetCompletedDate ensures that no value is present for CompletedDate, not even an explicit nil
### GetTestResultsConfigurationIds

`func (o *TestRunFilterApiModel) GetTestResultsConfigurationIds() []string`

GetTestResultsConfigurationIds returns the TestResultsConfigurationIds field if non-nil, zero value otherwise.

### GetTestResultsConfigurationIdsOk

`func (o *TestRunFilterApiModel) GetTestResultsConfigurationIdsOk() (*[]string, bool)`

GetTestResultsConfigurationIdsOk returns a tuple with the TestResultsConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResultsConfigurationIds

`func (o *TestRunFilterApiModel) SetTestResultsConfigurationIds(v []string)`

SetTestResultsConfigurationIds sets TestResultsConfigurationIds field to given value.

### HasTestResultsConfigurationIds

`func (o *TestRunFilterApiModel) HasTestResultsConfigurationIds() bool`

HasTestResultsConfigurationIds returns a boolean if a field has been set.

### SetTestResultsConfigurationIdsNil

`func (o *TestRunFilterApiModel) SetTestResultsConfigurationIdsNil(b bool)`

 SetTestResultsConfigurationIdsNil sets the value for TestResultsConfigurationIds to be an explicit nil

### UnsetTestResultsConfigurationIds
`func (o *TestRunFilterApiModel) UnsetTestResultsConfigurationIds()`

UnsetTestResultsConfigurationIds ensures that no value is present for TestResultsConfigurationIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


