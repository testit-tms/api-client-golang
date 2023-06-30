# ApiV2AutoTestsIdTestResultsSearchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outcomes** | Pointer to [**[]AutotestResultOutcome**](AutotestResultOutcome.md) |  | [optional] 
**TestPlanIds** | Pointer to **[]string** |  | [optional] 
**TestRunIds** | Pointer to **[]string** |  | [optional] 
**ConfigurationIds** | Pointer to **[]string** |  | [optional] 
**LaunchSource** | Pointer to **NullableString** |  | [optional] 
**UserIds** | Pointer to **[]string** |  | [optional] 
**Duration** | Pointer to [**NullableInt64RangeSelectorModel**](Int64RangeSelectorModel.md) |  | [optional] 

## Methods

### NewApiV2AutoTestsIdTestResultsSearchPostRequest

`func NewApiV2AutoTestsIdTestResultsSearchPostRequest() *ApiV2AutoTestsIdTestResultsSearchPostRequest`

NewApiV2AutoTestsIdTestResultsSearchPostRequest instantiates a new ApiV2AutoTestsIdTestResultsSearchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2AutoTestsIdTestResultsSearchPostRequestWithDefaults

`func NewApiV2AutoTestsIdTestResultsSearchPostRequestWithDefaults() *ApiV2AutoTestsIdTestResultsSearchPostRequest`

NewApiV2AutoTestsIdTestResultsSearchPostRequestWithDefaults instantiates a new ApiV2AutoTestsIdTestResultsSearchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutcomes

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) GetOutcomes() []AutotestResultOutcome`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) GetOutcomesOk() (*[]AutotestResultOutcome, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomes

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) SetOutcomes(v []AutotestResultOutcome)`

SetOutcomes sets Outcomes field to given value.

### HasOutcomes

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.

### SetOutcomesNil

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) SetOutcomesNil(b bool)`

 SetOutcomesNil sets the value for Outcomes to be an explicit nil

### UnsetOutcomes
`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) UnsetOutcomes()`

UnsetOutcomes ensures that no value is present for Outcomes, not even an explicit nil
### GetTestPlanIds

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) GetTestPlanIds() []string`

GetTestPlanIds returns the TestPlanIds field if non-nil, zero value otherwise.

### GetTestPlanIdsOk

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) GetTestPlanIdsOk() (*[]string, bool)`

GetTestPlanIdsOk returns a tuple with the TestPlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanIds

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) SetTestPlanIds(v []string)`

SetTestPlanIds sets TestPlanIds field to given value.

### HasTestPlanIds

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) HasTestPlanIds() bool`

HasTestPlanIds returns a boolean if a field has been set.

### SetTestPlanIdsNil

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) SetTestPlanIdsNil(b bool)`

 SetTestPlanIdsNil sets the value for TestPlanIds to be an explicit nil

### UnsetTestPlanIds
`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) UnsetTestPlanIds()`

UnsetTestPlanIds ensures that no value is present for TestPlanIds, not even an explicit nil
### GetTestRunIds

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) GetTestRunIds() []string`

GetTestRunIds returns the TestRunIds field if non-nil, zero value otherwise.

### GetTestRunIdsOk

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) GetTestRunIdsOk() (*[]string, bool)`

GetTestRunIdsOk returns a tuple with the TestRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunIds

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) SetTestRunIds(v []string)`

SetTestRunIds sets TestRunIds field to given value.

### HasTestRunIds

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) HasTestRunIds() bool`

HasTestRunIds returns a boolean if a field has been set.

### SetTestRunIdsNil

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) SetTestRunIdsNil(b bool)`

 SetTestRunIdsNil sets the value for TestRunIds to be an explicit nil

### UnsetTestRunIds
`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) UnsetTestRunIds()`

UnsetTestRunIds ensures that no value is present for TestRunIds, not even an explicit nil
### GetConfigurationIds

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) GetConfigurationIds() []string`

GetConfigurationIds returns the ConfigurationIds field if non-nil, zero value otherwise.

### GetConfigurationIdsOk

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) GetConfigurationIdsOk() (*[]string, bool)`

GetConfigurationIdsOk returns a tuple with the ConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIds

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) SetConfigurationIds(v []string)`

SetConfigurationIds sets ConfigurationIds field to given value.

### HasConfigurationIds

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) HasConfigurationIds() bool`

HasConfigurationIds returns a boolean if a field has been set.

### SetConfigurationIdsNil

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) SetConfigurationIdsNil(b bool)`

 SetConfigurationIdsNil sets the value for ConfigurationIds to be an explicit nil

### UnsetConfigurationIds
`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) UnsetConfigurationIds()`

UnsetConfigurationIds ensures that no value is present for ConfigurationIds, not even an explicit nil
### GetLaunchSource

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) GetLaunchSource() string`

GetLaunchSource returns the LaunchSource field if non-nil, zero value otherwise.

### GetLaunchSourceOk

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) GetLaunchSourceOk() (*string, bool)`

GetLaunchSourceOk returns a tuple with the LaunchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSource

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) SetLaunchSource(v string)`

SetLaunchSource sets LaunchSource field to given value.

### HasLaunchSource

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) HasLaunchSource() bool`

HasLaunchSource returns a boolean if a field has been set.

### SetLaunchSourceNil

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) SetLaunchSourceNil(b bool)`

 SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil

### UnsetLaunchSource
`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) UnsetLaunchSource()`

UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
### GetUserIds

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### SetUserIdsNil

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) SetUserIdsNil(b bool)`

 SetUserIdsNil sets the value for UserIds to be an explicit nil

### UnsetUserIds
`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) UnsetUserIds()`

UnsetUserIds ensures that no value is present for UserIds, not even an explicit nil
### GetDuration

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) GetDuration() Int64RangeSelectorModel`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) GetDurationOk() (*Int64RangeSelectorModel, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) SetDuration(v Int64RangeSelectorModel)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *ApiV2AutoTestsIdTestResultsSearchPostRequest) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


