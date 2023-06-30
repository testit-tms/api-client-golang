# ApiV2TestRunsIdStatisticsFilterPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationIds** | Pointer to **[]string** | Specifies a test result configuration IDs to search for | [optional] 
**Outcomes** | Pointer to [**[]TestResultOutcome**](TestResultOutcome.md) | Specifies a test result outcomes to search for | [optional] 
**FailureCategories** | Pointer to [**[]FailureCategoryModel**](FailureCategoryModel.md) | Specifies a test result failure categories to search for | [optional] 
**Namespace** | Pointer to **NullableString** | Specifies a test result namespace to search for | [optional] 
**ClassName** | Pointer to **NullableString** | Specifies a test result class name to search for | [optional] 

## Methods

### NewApiV2TestRunsIdStatisticsFilterPostRequest

`func NewApiV2TestRunsIdStatisticsFilterPostRequest() *ApiV2TestRunsIdStatisticsFilterPostRequest`

NewApiV2TestRunsIdStatisticsFilterPostRequest instantiates a new ApiV2TestRunsIdStatisticsFilterPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2TestRunsIdStatisticsFilterPostRequestWithDefaults

`func NewApiV2TestRunsIdStatisticsFilterPostRequestWithDefaults() *ApiV2TestRunsIdStatisticsFilterPostRequest`

NewApiV2TestRunsIdStatisticsFilterPostRequestWithDefaults instantiates a new ApiV2TestRunsIdStatisticsFilterPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationIds

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) GetConfigurationIds() []string`

GetConfigurationIds returns the ConfigurationIds field if non-nil, zero value otherwise.

### GetConfigurationIdsOk

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) GetConfigurationIdsOk() (*[]string, bool)`

GetConfigurationIdsOk returns a tuple with the ConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIds

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) SetConfigurationIds(v []string)`

SetConfigurationIds sets ConfigurationIds field to given value.

### HasConfigurationIds

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) HasConfigurationIds() bool`

HasConfigurationIds returns a boolean if a field has been set.

### SetConfigurationIdsNil

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) SetConfigurationIdsNil(b bool)`

 SetConfigurationIdsNil sets the value for ConfigurationIds to be an explicit nil

### UnsetConfigurationIds
`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) UnsetConfigurationIds()`

UnsetConfigurationIds ensures that no value is present for ConfigurationIds, not even an explicit nil
### GetOutcomes

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) GetOutcomes() []TestResultOutcome`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) GetOutcomesOk() (*[]TestResultOutcome, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomes

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) SetOutcomes(v []TestResultOutcome)`

SetOutcomes sets Outcomes field to given value.

### HasOutcomes

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.

### SetOutcomesNil

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) SetOutcomesNil(b bool)`

 SetOutcomesNil sets the value for Outcomes to be an explicit nil

### UnsetOutcomes
`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) UnsetOutcomes()`

UnsetOutcomes ensures that no value is present for Outcomes, not even an explicit nil
### GetFailureCategories

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) GetFailureCategories() []FailureCategoryModel`

GetFailureCategories returns the FailureCategories field if non-nil, zero value otherwise.

### GetFailureCategoriesOk

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) GetFailureCategoriesOk() (*[]FailureCategoryModel, bool)`

GetFailureCategoriesOk returns a tuple with the FailureCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategories

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) SetFailureCategories(v []FailureCategoryModel)`

SetFailureCategories sets FailureCategories field to given value.

### HasFailureCategories

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) HasFailureCategories() bool`

HasFailureCategories returns a boolean if a field has been set.

### SetFailureCategoriesNil

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) SetFailureCategoriesNil(b bool)`

 SetFailureCategoriesNil sets the value for FailureCategories to be an explicit nil

### UnsetFailureCategories
`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) UnsetFailureCategories()`

UnsetFailureCategories ensures that no value is present for FailureCategories, not even an explicit nil
### GetNamespace

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetClassName

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### SetClassNameNil

`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) SetClassNameNil(b bool)`

 SetClassNameNil sets the value for ClassName to be an explicit nil

### UnsetClassName
`func (o *ApiV2TestRunsIdStatisticsFilterPostRequest) UnsetClassName()`

UnsetClassName ensures that no value is present for ClassName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


