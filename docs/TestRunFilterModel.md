# TestRunFilterModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectIds** | Pointer to **[]string** | Specifies a test run project IDs to search for | [optional] 
**Name** | Pointer to **NullableString** | Specifies test run name | [optional] 
**States** | Pointer to [**[]TestRunState**](TestRunState.md) | Specifies a test run states to search for | [optional] 
**CreatedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) | Specifies a test run range of created date to search for | [optional] 
**StartedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) | Specifies a test run range of started date to search for | [optional] 
**CreatedByIds** | Pointer to **[]string** | Specifies a test run creator IDs to search for | [optional] 
**ModifiedByIds** | Pointer to **[]string** | Specifies a test run last editor IDs to search for | [optional] 
**IsDeleted** | Pointer to **NullableBool** | Specifies a test run deleted status to search for | [optional] 
**AutoTestsCount** | Pointer to [**NullableInt32RangeSelectorModel**](Int32RangeSelectorModel.md) | Number of autoTests run in the test run | [optional] 
**TestResultsOutcome** | Pointer to [**[]TestResultOutcome**](TestResultOutcome.md) | Specifies test results outcomes | [optional] 
**FailureCategory** | Pointer to [**[]FailureCategoryModel**](FailureCategoryModel.md) | Specifies failure categories | [optional] 
**CompletedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) | Specifies a test run range of completed date to search for | [optional] 
**TestResultsConfigurationIds** | Pointer to **[]string** | Specifies a test result configuration IDs to search for | [optional] 

## Methods

### NewTestRunFilterModel

`func NewTestRunFilterModel() *TestRunFilterModel`

NewTestRunFilterModel instantiates a new TestRunFilterModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunFilterModelWithDefaults

`func NewTestRunFilterModelWithDefaults() *TestRunFilterModel`

NewTestRunFilterModelWithDefaults instantiates a new TestRunFilterModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectIds

`func (o *TestRunFilterModel) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *TestRunFilterModel) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *TestRunFilterModel) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *TestRunFilterModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *TestRunFilterModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *TestRunFilterModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetName

`func (o *TestRunFilterModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestRunFilterModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestRunFilterModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestRunFilterModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestRunFilterModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestRunFilterModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStates

`func (o *TestRunFilterModel) GetStates() []TestRunState`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *TestRunFilterModel) GetStatesOk() (*[]TestRunState, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *TestRunFilterModel) SetStates(v []TestRunState)`

SetStates sets States field to given value.

### HasStates

`func (o *TestRunFilterModel) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStatesNil

`func (o *TestRunFilterModel) SetStatesNil(b bool)`

 SetStatesNil sets the value for States to be an explicit nil

### UnsetStates
`func (o *TestRunFilterModel) UnsetStates()`

UnsetStates ensures that no value is present for States, not even an explicit nil
### GetCreatedDate

`func (o *TestRunFilterModel) GetCreatedDate() DateTimeRangeSelectorModel`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestRunFilterModel) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestRunFilterModel) SetCreatedDate(v DateTimeRangeSelectorModel)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TestRunFilterModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *TestRunFilterModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *TestRunFilterModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetStartedDate

`func (o *TestRunFilterModel) GetStartedDate() DateTimeRangeSelectorModel`

GetStartedDate returns the StartedDate field if non-nil, zero value otherwise.

### GetStartedDateOk

`func (o *TestRunFilterModel) GetStartedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetStartedDateOk returns a tuple with the StartedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedDate

`func (o *TestRunFilterModel) SetStartedDate(v DateTimeRangeSelectorModel)`

SetStartedDate sets StartedDate field to given value.

### HasStartedDate

`func (o *TestRunFilterModel) HasStartedDate() bool`

HasStartedDate returns a boolean if a field has been set.

### SetStartedDateNil

`func (o *TestRunFilterModel) SetStartedDateNil(b bool)`

 SetStartedDateNil sets the value for StartedDate to be an explicit nil

### UnsetStartedDate
`func (o *TestRunFilterModel) UnsetStartedDate()`

UnsetStartedDate ensures that no value is present for StartedDate, not even an explicit nil
### GetCreatedByIds

`func (o *TestRunFilterModel) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *TestRunFilterModel) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *TestRunFilterModel) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *TestRunFilterModel) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *TestRunFilterModel) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *TestRunFilterModel) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedByIds

`func (o *TestRunFilterModel) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *TestRunFilterModel) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *TestRunFilterModel) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *TestRunFilterModel) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *TestRunFilterModel) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *TestRunFilterModel) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil
### GetIsDeleted

`func (o *TestRunFilterModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TestRunFilterModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TestRunFilterModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *TestRunFilterModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *TestRunFilterModel) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *TestRunFilterModel) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetAutoTestsCount

`func (o *TestRunFilterModel) GetAutoTestsCount() Int32RangeSelectorModel`

GetAutoTestsCount returns the AutoTestsCount field if non-nil, zero value otherwise.

### GetAutoTestsCountOk

`func (o *TestRunFilterModel) GetAutoTestsCountOk() (*Int32RangeSelectorModel, bool)`

GetAutoTestsCountOk returns a tuple with the AutoTestsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestsCount

`func (o *TestRunFilterModel) SetAutoTestsCount(v Int32RangeSelectorModel)`

SetAutoTestsCount sets AutoTestsCount field to given value.

### HasAutoTestsCount

`func (o *TestRunFilterModel) HasAutoTestsCount() bool`

HasAutoTestsCount returns a boolean if a field has been set.

### SetAutoTestsCountNil

`func (o *TestRunFilterModel) SetAutoTestsCountNil(b bool)`

 SetAutoTestsCountNil sets the value for AutoTestsCount to be an explicit nil

### UnsetAutoTestsCount
`func (o *TestRunFilterModel) UnsetAutoTestsCount()`

UnsetAutoTestsCount ensures that no value is present for AutoTestsCount, not even an explicit nil
### GetTestResultsOutcome

`func (o *TestRunFilterModel) GetTestResultsOutcome() []TestResultOutcome`

GetTestResultsOutcome returns the TestResultsOutcome field if non-nil, zero value otherwise.

### GetTestResultsOutcomeOk

`func (o *TestRunFilterModel) GetTestResultsOutcomeOk() (*[]TestResultOutcome, bool)`

GetTestResultsOutcomeOk returns a tuple with the TestResultsOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResultsOutcome

`func (o *TestRunFilterModel) SetTestResultsOutcome(v []TestResultOutcome)`

SetTestResultsOutcome sets TestResultsOutcome field to given value.

### HasTestResultsOutcome

`func (o *TestRunFilterModel) HasTestResultsOutcome() bool`

HasTestResultsOutcome returns a boolean if a field has been set.

### SetTestResultsOutcomeNil

`func (o *TestRunFilterModel) SetTestResultsOutcomeNil(b bool)`

 SetTestResultsOutcomeNil sets the value for TestResultsOutcome to be an explicit nil

### UnsetTestResultsOutcome
`func (o *TestRunFilterModel) UnsetTestResultsOutcome()`

UnsetTestResultsOutcome ensures that no value is present for TestResultsOutcome, not even an explicit nil
### GetFailureCategory

`func (o *TestRunFilterModel) GetFailureCategory() []FailureCategoryModel`

GetFailureCategory returns the FailureCategory field if non-nil, zero value otherwise.

### GetFailureCategoryOk

`func (o *TestRunFilterModel) GetFailureCategoryOk() (*[]FailureCategoryModel, bool)`

GetFailureCategoryOk returns a tuple with the FailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategory

`func (o *TestRunFilterModel) SetFailureCategory(v []FailureCategoryModel)`

SetFailureCategory sets FailureCategory field to given value.

### HasFailureCategory

`func (o *TestRunFilterModel) HasFailureCategory() bool`

HasFailureCategory returns a boolean if a field has been set.

### SetFailureCategoryNil

`func (o *TestRunFilterModel) SetFailureCategoryNil(b bool)`

 SetFailureCategoryNil sets the value for FailureCategory to be an explicit nil

### UnsetFailureCategory
`func (o *TestRunFilterModel) UnsetFailureCategory()`

UnsetFailureCategory ensures that no value is present for FailureCategory, not even an explicit nil
### GetCompletedDate

`func (o *TestRunFilterModel) GetCompletedDate() DateTimeRangeSelectorModel`

GetCompletedDate returns the CompletedDate field if non-nil, zero value otherwise.

### GetCompletedDateOk

`func (o *TestRunFilterModel) GetCompletedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCompletedDateOk returns a tuple with the CompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDate

`func (o *TestRunFilterModel) SetCompletedDate(v DateTimeRangeSelectorModel)`

SetCompletedDate sets CompletedDate field to given value.

### HasCompletedDate

`func (o *TestRunFilterModel) HasCompletedDate() bool`

HasCompletedDate returns a boolean if a field has been set.

### SetCompletedDateNil

`func (o *TestRunFilterModel) SetCompletedDateNil(b bool)`

 SetCompletedDateNil sets the value for CompletedDate to be an explicit nil

### UnsetCompletedDate
`func (o *TestRunFilterModel) UnsetCompletedDate()`

UnsetCompletedDate ensures that no value is present for CompletedDate, not even an explicit nil
### GetTestResultsConfigurationIds

`func (o *TestRunFilterModel) GetTestResultsConfigurationIds() []string`

GetTestResultsConfigurationIds returns the TestResultsConfigurationIds field if non-nil, zero value otherwise.

### GetTestResultsConfigurationIdsOk

`func (o *TestRunFilterModel) GetTestResultsConfigurationIdsOk() (*[]string, bool)`

GetTestResultsConfigurationIdsOk returns a tuple with the TestResultsConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResultsConfigurationIds

`func (o *TestRunFilterModel) SetTestResultsConfigurationIds(v []string)`

SetTestResultsConfigurationIds sets TestResultsConfigurationIds field to given value.

### HasTestResultsConfigurationIds

`func (o *TestRunFilterModel) HasTestResultsConfigurationIds() bool`

HasTestResultsConfigurationIds returns a boolean if a field has been set.

### SetTestResultsConfigurationIdsNil

`func (o *TestRunFilterModel) SetTestResultsConfigurationIdsNil(b bool)`

 SetTestResultsConfigurationIdsNil sets the value for TestResultsConfigurationIds to be an explicit nil

### UnsetTestResultsConfigurationIds
`func (o *TestRunFilterModel) UnsetTestResultsConfigurationIds()`

UnsetTestResultsConfigurationIds ensures that no value is present for TestResultsConfigurationIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


