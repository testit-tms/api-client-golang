# TestResultsFilterModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestRunIds** | Pointer to **[]string** | Specifies a test result test run IDs to search for | [optional] 
**AutoTestGlobalIds** | Pointer to **[]int64** | Specifies an autotest global IDs to search results for | [optional] 
**Name** | Pointer to **NullableString** | Specifies an autotest name to search results for | [optional] 
**CreatedDate** | Pointer to [**NullableTestResultsFilterModelCreatedDate**](TestResultsFilterModelCreatedDate.md) |  | [optional] 
**ModifiedDate** | Pointer to [**NullableTestResultsFilterModelModifiedDate**](TestResultsFilterModelModifiedDate.md) |  | [optional] 
**StartedOn** | Pointer to [**NullableTestResultsFilterModelStartedOn**](TestResultsFilterModelStartedOn.md) |  | [optional] 
**CompletedOn** | Pointer to [**NullableTestResultsFilterModelCompletedOn**](TestResultsFilterModelCompletedOn.md) |  | [optional] 
**Duration** | Pointer to [**NullableTestResultsFilterModelDuration**](TestResultsFilterModelDuration.md) |  | [optional] 
**ResultReasons** | Pointer to **[]string** | Specifies result reasons for searching test results | [optional] 
**ConfigurationIds** | Pointer to **[]string** | Specifies a test result configuration IDs to search for | [optional] 
**Outcomes** | Pointer to [**[]TestResultOutcome**](TestResultOutcome.md) | Specifies a test result outcomes to search for | [optional] 
**FailureCategories** | Pointer to [**[]FailureCategoryModel**](FailureCategoryModel.md) | Specifies a test result failure categories to search for | [optional] 
**Namespace** | Pointer to **NullableString** | Specifies a test result namespace to search for | [optional] 
**ClassName** | Pointer to **NullableString** | Specifies a test result class name to search for | [optional] 

## Methods

### NewTestResultsFilterModel

`func NewTestResultsFilterModel() *TestResultsFilterModel`

NewTestResultsFilterModel instantiates a new TestResultsFilterModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultsFilterModelWithDefaults

`func NewTestResultsFilterModelWithDefaults() *TestResultsFilterModel`

NewTestResultsFilterModelWithDefaults instantiates a new TestResultsFilterModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestRunIds

`func (o *TestResultsFilterModel) GetTestRunIds() []string`

GetTestRunIds returns the TestRunIds field if non-nil, zero value otherwise.

### GetTestRunIdsOk

`func (o *TestResultsFilterModel) GetTestRunIdsOk() (*[]string, bool)`

GetTestRunIdsOk returns a tuple with the TestRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunIds

`func (o *TestResultsFilterModel) SetTestRunIds(v []string)`

SetTestRunIds sets TestRunIds field to given value.

### HasTestRunIds

`func (o *TestResultsFilterModel) HasTestRunIds() bool`

HasTestRunIds returns a boolean if a field has been set.

### SetTestRunIdsNil

`func (o *TestResultsFilterModel) SetTestRunIdsNil(b bool)`

 SetTestRunIdsNil sets the value for TestRunIds to be an explicit nil

### UnsetTestRunIds
`func (o *TestResultsFilterModel) UnsetTestRunIds()`

UnsetTestRunIds ensures that no value is present for TestRunIds, not even an explicit nil
### GetAutoTestGlobalIds

`func (o *TestResultsFilterModel) GetAutoTestGlobalIds() []int64`

GetAutoTestGlobalIds returns the AutoTestGlobalIds field if non-nil, zero value otherwise.

### GetAutoTestGlobalIdsOk

`func (o *TestResultsFilterModel) GetAutoTestGlobalIdsOk() (*[]int64, bool)`

GetAutoTestGlobalIdsOk returns a tuple with the AutoTestGlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestGlobalIds

`func (o *TestResultsFilterModel) SetAutoTestGlobalIds(v []int64)`

SetAutoTestGlobalIds sets AutoTestGlobalIds field to given value.

### HasAutoTestGlobalIds

`func (o *TestResultsFilterModel) HasAutoTestGlobalIds() bool`

HasAutoTestGlobalIds returns a boolean if a field has been set.

### SetAutoTestGlobalIdsNil

`func (o *TestResultsFilterModel) SetAutoTestGlobalIdsNil(b bool)`

 SetAutoTestGlobalIdsNil sets the value for AutoTestGlobalIds to be an explicit nil

### UnsetAutoTestGlobalIds
`func (o *TestResultsFilterModel) UnsetAutoTestGlobalIds()`

UnsetAutoTestGlobalIds ensures that no value is present for AutoTestGlobalIds, not even an explicit nil
### GetName

`func (o *TestResultsFilterModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestResultsFilterModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestResultsFilterModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestResultsFilterModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestResultsFilterModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestResultsFilterModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreatedDate

`func (o *TestResultsFilterModel) GetCreatedDate() TestResultsFilterModelCreatedDate`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestResultsFilterModel) GetCreatedDateOk() (*TestResultsFilterModelCreatedDate, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestResultsFilterModel) SetCreatedDate(v TestResultsFilterModelCreatedDate)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TestResultsFilterModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *TestResultsFilterModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *TestResultsFilterModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetModifiedDate

`func (o *TestResultsFilterModel) GetModifiedDate() TestResultsFilterModelModifiedDate`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestResultsFilterModel) GetModifiedDateOk() (*TestResultsFilterModelModifiedDate, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestResultsFilterModel) SetModifiedDate(v TestResultsFilterModelModifiedDate)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestResultsFilterModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *TestResultsFilterModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *TestResultsFilterModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetStartedOn

`func (o *TestResultsFilterModel) GetStartedOn() TestResultsFilterModelStartedOn`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *TestResultsFilterModel) GetStartedOnOk() (*TestResultsFilterModelStartedOn, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *TestResultsFilterModel) SetStartedOn(v TestResultsFilterModelStartedOn)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *TestResultsFilterModel) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *TestResultsFilterModel) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *TestResultsFilterModel) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetCompletedOn

`func (o *TestResultsFilterModel) GetCompletedOn() TestResultsFilterModelCompletedOn`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *TestResultsFilterModel) GetCompletedOnOk() (*TestResultsFilterModelCompletedOn, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *TestResultsFilterModel) SetCompletedOn(v TestResultsFilterModelCompletedOn)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *TestResultsFilterModel) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### SetCompletedOnNil

`func (o *TestResultsFilterModel) SetCompletedOnNil(b bool)`

 SetCompletedOnNil sets the value for CompletedOn to be an explicit nil

### UnsetCompletedOn
`func (o *TestResultsFilterModel) UnsetCompletedOn()`

UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
### GetDuration

`func (o *TestResultsFilterModel) GetDuration() TestResultsFilterModelDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TestResultsFilterModel) GetDurationOk() (*TestResultsFilterModelDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TestResultsFilterModel) SetDuration(v TestResultsFilterModelDuration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TestResultsFilterModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *TestResultsFilterModel) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *TestResultsFilterModel) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetResultReasons

`func (o *TestResultsFilterModel) GetResultReasons() []string`

GetResultReasons returns the ResultReasons field if non-nil, zero value otherwise.

### GetResultReasonsOk

`func (o *TestResultsFilterModel) GetResultReasonsOk() (*[]string, bool)`

GetResultReasonsOk returns a tuple with the ResultReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultReasons

`func (o *TestResultsFilterModel) SetResultReasons(v []string)`

SetResultReasons sets ResultReasons field to given value.

### HasResultReasons

`func (o *TestResultsFilterModel) HasResultReasons() bool`

HasResultReasons returns a boolean if a field has been set.

### SetResultReasonsNil

`func (o *TestResultsFilterModel) SetResultReasonsNil(b bool)`

 SetResultReasonsNil sets the value for ResultReasons to be an explicit nil

### UnsetResultReasons
`func (o *TestResultsFilterModel) UnsetResultReasons()`

UnsetResultReasons ensures that no value is present for ResultReasons, not even an explicit nil
### GetConfigurationIds

`func (o *TestResultsFilterModel) GetConfigurationIds() []string`

GetConfigurationIds returns the ConfigurationIds field if non-nil, zero value otherwise.

### GetConfigurationIdsOk

`func (o *TestResultsFilterModel) GetConfigurationIdsOk() (*[]string, bool)`

GetConfigurationIdsOk returns a tuple with the ConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIds

`func (o *TestResultsFilterModel) SetConfigurationIds(v []string)`

SetConfigurationIds sets ConfigurationIds field to given value.

### HasConfigurationIds

`func (o *TestResultsFilterModel) HasConfigurationIds() bool`

HasConfigurationIds returns a boolean if a field has been set.

### SetConfigurationIdsNil

`func (o *TestResultsFilterModel) SetConfigurationIdsNil(b bool)`

 SetConfigurationIdsNil sets the value for ConfigurationIds to be an explicit nil

### UnsetConfigurationIds
`func (o *TestResultsFilterModel) UnsetConfigurationIds()`

UnsetConfigurationIds ensures that no value is present for ConfigurationIds, not even an explicit nil
### GetOutcomes

`func (o *TestResultsFilterModel) GetOutcomes() []TestResultOutcome`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *TestResultsFilterModel) GetOutcomesOk() (*[]TestResultOutcome, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomes

`func (o *TestResultsFilterModel) SetOutcomes(v []TestResultOutcome)`

SetOutcomes sets Outcomes field to given value.

### HasOutcomes

`func (o *TestResultsFilterModel) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.

### SetOutcomesNil

`func (o *TestResultsFilterModel) SetOutcomesNil(b bool)`

 SetOutcomesNil sets the value for Outcomes to be an explicit nil

### UnsetOutcomes
`func (o *TestResultsFilterModel) UnsetOutcomes()`

UnsetOutcomes ensures that no value is present for Outcomes, not even an explicit nil
### GetFailureCategories

`func (o *TestResultsFilterModel) GetFailureCategories() []FailureCategoryModel`

GetFailureCategories returns the FailureCategories field if non-nil, zero value otherwise.

### GetFailureCategoriesOk

`func (o *TestResultsFilterModel) GetFailureCategoriesOk() (*[]FailureCategoryModel, bool)`

GetFailureCategoriesOk returns a tuple with the FailureCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategories

`func (o *TestResultsFilterModel) SetFailureCategories(v []FailureCategoryModel)`

SetFailureCategories sets FailureCategories field to given value.

### HasFailureCategories

`func (o *TestResultsFilterModel) HasFailureCategories() bool`

HasFailureCategories returns a boolean if a field has been set.

### SetFailureCategoriesNil

`func (o *TestResultsFilterModel) SetFailureCategoriesNil(b bool)`

 SetFailureCategoriesNil sets the value for FailureCategories to be an explicit nil

### UnsetFailureCategories
`func (o *TestResultsFilterModel) UnsetFailureCategories()`

UnsetFailureCategories ensures that no value is present for FailureCategories, not even an explicit nil
### GetNamespace

`func (o *TestResultsFilterModel) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *TestResultsFilterModel) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *TestResultsFilterModel) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *TestResultsFilterModel) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *TestResultsFilterModel) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *TestResultsFilterModel) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetClassName

`func (o *TestResultsFilterModel) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *TestResultsFilterModel) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *TestResultsFilterModel) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *TestResultsFilterModel) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### SetClassNameNil

`func (o *TestResultsFilterModel) SetClassNameNil(b bool)`

 SetClassNameNil sets the value for ClassName to be an explicit nil

### UnsetClassName
`func (o *TestResultsFilterModel) UnsetClassName()`

UnsetClassName ensures that no value is present for ClassName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


