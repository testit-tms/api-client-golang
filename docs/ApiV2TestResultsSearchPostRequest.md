# ApiV2TestResultsSearchPostRequest

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

### NewApiV2TestResultsSearchPostRequest

`func NewApiV2TestResultsSearchPostRequest() *ApiV2TestResultsSearchPostRequest`

NewApiV2TestResultsSearchPostRequest instantiates a new ApiV2TestResultsSearchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2TestResultsSearchPostRequestWithDefaults

`func NewApiV2TestResultsSearchPostRequestWithDefaults() *ApiV2TestResultsSearchPostRequest`

NewApiV2TestResultsSearchPostRequestWithDefaults instantiates a new ApiV2TestResultsSearchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestRunIds

`func (o *ApiV2TestResultsSearchPostRequest) GetTestRunIds() []string`

GetTestRunIds returns the TestRunIds field if non-nil, zero value otherwise.

### GetTestRunIdsOk

`func (o *ApiV2TestResultsSearchPostRequest) GetTestRunIdsOk() (*[]string, bool)`

GetTestRunIdsOk returns a tuple with the TestRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunIds

`func (o *ApiV2TestResultsSearchPostRequest) SetTestRunIds(v []string)`

SetTestRunIds sets TestRunIds field to given value.

### HasTestRunIds

`func (o *ApiV2TestResultsSearchPostRequest) HasTestRunIds() bool`

HasTestRunIds returns a boolean if a field has been set.

### SetTestRunIdsNil

`func (o *ApiV2TestResultsSearchPostRequest) SetTestRunIdsNil(b bool)`

 SetTestRunIdsNil sets the value for TestRunIds to be an explicit nil

### UnsetTestRunIds
`func (o *ApiV2TestResultsSearchPostRequest) UnsetTestRunIds()`

UnsetTestRunIds ensures that no value is present for TestRunIds, not even an explicit nil
### GetAutoTestGlobalIds

`func (o *ApiV2TestResultsSearchPostRequest) GetAutoTestGlobalIds() []int64`

GetAutoTestGlobalIds returns the AutoTestGlobalIds field if non-nil, zero value otherwise.

### GetAutoTestGlobalIdsOk

`func (o *ApiV2TestResultsSearchPostRequest) GetAutoTestGlobalIdsOk() (*[]int64, bool)`

GetAutoTestGlobalIdsOk returns a tuple with the AutoTestGlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestGlobalIds

`func (o *ApiV2TestResultsSearchPostRequest) SetAutoTestGlobalIds(v []int64)`

SetAutoTestGlobalIds sets AutoTestGlobalIds field to given value.

### HasAutoTestGlobalIds

`func (o *ApiV2TestResultsSearchPostRequest) HasAutoTestGlobalIds() bool`

HasAutoTestGlobalIds returns a boolean if a field has been set.

### SetAutoTestGlobalIdsNil

`func (o *ApiV2TestResultsSearchPostRequest) SetAutoTestGlobalIdsNil(b bool)`

 SetAutoTestGlobalIdsNil sets the value for AutoTestGlobalIds to be an explicit nil

### UnsetAutoTestGlobalIds
`func (o *ApiV2TestResultsSearchPostRequest) UnsetAutoTestGlobalIds()`

UnsetAutoTestGlobalIds ensures that no value is present for AutoTestGlobalIds, not even an explicit nil
### GetName

`func (o *ApiV2TestResultsSearchPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2TestResultsSearchPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2TestResultsSearchPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV2TestResultsSearchPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApiV2TestResultsSearchPostRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApiV2TestResultsSearchPostRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreatedDate

`func (o *ApiV2TestResultsSearchPostRequest) GetCreatedDate() TestResultsFilterModelCreatedDate`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ApiV2TestResultsSearchPostRequest) GetCreatedDateOk() (*TestResultsFilterModelCreatedDate, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ApiV2TestResultsSearchPostRequest) SetCreatedDate(v TestResultsFilterModelCreatedDate)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ApiV2TestResultsSearchPostRequest) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *ApiV2TestResultsSearchPostRequest) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *ApiV2TestResultsSearchPostRequest) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetModifiedDate

`func (o *ApiV2TestResultsSearchPostRequest) GetModifiedDate() TestResultsFilterModelModifiedDate`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ApiV2TestResultsSearchPostRequest) GetModifiedDateOk() (*TestResultsFilterModelModifiedDate, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ApiV2TestResultsSearchPostRequest) SetModifiedDate(v TestResultsFilterModelModifiedDate)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *ApiV2TestResultsSearchPostRequest) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *ApiV2TestResultsSearchPostRequest) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *ApiV2TestResultsSearchPostRequest) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetStartedOn

`func (o *ApiV2TestResultsSearchPostRequest) GetStartedOn() TestResultsFilterModelStartedOn`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *ApiV2TestResultsSearchPostRequest) GetStartedOnOk() (*TestResultsFilterModelStartedOn, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *ApiV2TestResultsSearchPostRequest) SetStartedOn(v TestResultsFilterModelStartedOn)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *ApiV2TestResultsSearchPostRequest) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *ApiV2TestResultsSearchPostRequest) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *ApiV2TestResultsSearchPostRequest) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetCompletedOn

`func (o *ApiV2TestResultsSearchPostRequest) GetCompletedOn() TestResultsFilterModelCompletedOn`

GetCompletedOn returns the CompletedOn field if non-nil, zero value otherwise.

### GetCompletedOnOk

`func (o *ApiV2TestResultsSearchPostRequest) GetCompletedOnOk() (*TestResultsFilterModelCompletedOn, bool)`

GetCompletedOnOk returns a tuple with the CompletedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedOn

`func (o *ApiV2TestResultsSearchPostRequest) SetCompletedOn(v TestResultsFilterModelCompletedOn)`

SetCompletedOn sets CompletedOn field to given value.

### HasCompletedOn

`func (o *ApiV2TestResultsSearchPostRequest) HasCompletedOn() bool`

HasCompletedOn returns a boolean if a field has been set.

### SetCompletedOnNil

`func (o *ApiV2TestResultsSearchPostRequest) SetCompletedOnNil(b bool)`

 SetCompletedOnNil sets the value for CompletedOn to be an explicit nil

### UnsetCompletedOn
`func (o *ApiV2TestResultsSearchPostRequest) UnsetCompletedOn()`

UnsetCompletedOn ensures that no value is present for CompletedOn, not even an explicit nil
### GetDuration

`func (o *ApiV2TestResultsSearchPostRequest) GetDuration() TestResultsFilterModelDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ApiV2TestResultsSearchPostRequest) GetDurationOk() (*TestResultsFilterModelDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ApiV2TestResultsSearchPostRequest) SetDuration(v TestResultsFilterModelDuration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ApiV2TestResultsSearchPostRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *ApiV2TestResultsSearchPostRequest) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *ApiV2TestResultsSearchPostRequest) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetResultReasons

`func (o *ApiV2TestResultsSearchPostRequest) GetResultReasons() []string`

GetResultReasons returns the ResultReasons field if non-nil, zero value otherwise.

### GetResultReasonsOk

`func (o *ApiV2TestResultsSearchPostRequest) GetResultReasonsOk() (*[]string, bool)`

GetResultReasonsOk returns a tuple with the ResultReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultReasons

`func (o *ApiV2TestResultsSearchPostRequest) SetResultReasons(v []string)`

SetResultReasons sets ResultReasons field to given value.

### HasResultReasons

`func (o *ApiV2TestResultsSearchPostRequest) HasResultReasons() bool`

HasResultReasons returns a boolean if a field has been set.

### SetResultReasonsNil

`func (o *ApiV2TestResultsSearchPostRequest) SetResultReasonsNil(b bool)`

 SetResultReasonsNil sets the value for ResultReasons to be an explicit nil

### UnsetResultReasons
`func (o *ApiV2TestResultsSearchPostRequest) UnsetResultReasons()`

UnsetResultReasons ensures that no value is present for ResultReasons, not even an explicit nil
### GetConfigurationIds

`func (o *ApiV2TestResultsSearchPostRequest) GetConfigurationIds() []string`

GetConfigurationIds returns the ConfigurationIds field if non-nil, zero value otherwise.

### GetConfigurationIdsOk

`func (o *ApiV2TestResultsSearchPostRequest) GetConfigurationIdsOk() (*[]string, bool)`

GetConfigurationIdsOk returns a tuple with the ConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIds

`func (o *ApiV2TestResultsSearchPostRequest) SetConfigurationIds(v []string)`

SetConfigurationIds sets ConfigurationIds field to given value.

### HasConfigurationIds

`func (o *ApiV2TestResultsSearchPostRequest) HasConfigurationIds() bool`

HasConfigurationIds returns a boolean if a field has been set.

### SetConfigurationIdsNil

`func (o *ApiV2TestResultsSearchPostRequest) SetConfigurationIdsNil(b bool)`

 SetConfigurationIdsNil sets the value for ConfigurationIds to be an explicit nil

### UnsetConfigurationIds
`func (o *ApiV2TestResultsSearchPostRequest) UnsetConfigurationIds()`

UnsetConfigurationIds ensures that no value is present for ConfigurationIds, not even an explicit nil
### GetOutcomes

`func (o *ApiV2TestResultsSearchPostRequest) GetOutcomes() []TestResultOutcome`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *ApiV2TestResultsSearchPostRequest) GetOutcomesOk() (*[]TestResultOutcome, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomes

`func (o *ApiV2TestResultsSearchPostRequest) SetOutcomes(v []TestResultOutcome)`

SetOutcomes sets Outcomes field to given value.

### HasOutcomes

`func (o *ApiV2TestResultsSearchPostRequest) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.

### SetOutcomesNil

`func (o *ApiV2TestResultsSearchPostRequest) SetOutcomesNil(b bool)`

 SetOutcomesNil sets the value for Outcomes to be an explicit nil

### UnsetOutcomes
`func (o *ApiV2TestResultsSearchPostRequest) UnsetOutcomes()`

UnsetOutcomes ensures that no value is present for Outcomes, not even an explicit nil
### GetFailureCategories

`func (o *ApiV2TestResultsSearchPostRequest) GetFailureCategories() []FailureCategoryModel`

GetFailureCategories returns the FailureCategories field if non-nil, zero value otherwise.

### GetFailureCategoriesOk

`func (o *ApiV2TestResultsSearchPostRequest) GetFailureCategoriesOk() (*[]FailureCategoryModel, bool)`

GetFailureCategoriesOk returns a tuple with the FailureCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategories

`func (o *ApiV2TestResultsSearchPostRequest) SetFailureCategories(v []FailureCategoryModel)`

SetFailureCategories sets FailureCategories field to given value.

### HasFailureCategories

`func (o *ApiV2TestResultsSearchPostRequest) HasFailureCategories() bool`

HasFailureCategories returns a boolean if a field has been set.

### SetFailureCategoriesNil

`func (o *ApiV2TestResultsSearchPostRequest) SetFailureCategoriesNil(b bool)`

 SetFailureCategoriesNil sets the value for FailureCategories to be an explicit nil

### UnsetFailureCategories
`func (o *ApiV2TestResultsSearchPostRequest) UnsetFailureCategories()`

UnsetFailureCategories ensures that no value is present for FailureCategories, not even an explicit nil
### GetNamespace

`func (o *ApiV2TestResultsSearchPostRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ApiV2TestResultsSearchPostRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ApiV2TestResultsSearchPostRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ApiV2TestResultsSearchPostRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *ApiV2TestResultsSearchPostRequest) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *ApiV2TestResultsSearchPostRequest) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetClassName

`func (o *ApiV2TestResultsSearchPostRequest) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *ApiV2TestResultsSearchPostRequest) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *ApiV2TestResultsSearchPostRequest) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *ApiV2TestResultsSearchPostRequest) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### SetClassNameNil

`func (o *ApiV2TestResultsSearchPostRequest) SetClassNameNil(b bool)`

 SetClassNameNil sets the value for ClassName to be an explicit nil

### UnsetClassName
`func (o *ApiV2TestResultsSearchPostRequest) UnsetClassName()`

UnsetClassName ensures that no value is present for ClassName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


