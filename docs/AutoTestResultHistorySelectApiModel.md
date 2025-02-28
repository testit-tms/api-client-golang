# AutoTestResultHistorySelectApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outcomes** | Pointer to [**[]AutotestResultOutcome**](AutotestResultOutcome.md) |  | [optional] 
**StatusCodes** | Pointer to **[]string** |  | [optional] 
**TestPlanIds** | Pointer to **[]string** |  | [optional] 
**TestRunIds** | Pointer to **[]string** |  | [optional] 
**ConfigurationIds** | Pointer to **[]string** |  | [optional] 
**LaunchSource** | Pointer to **NullableString** |  | [optional] 
**UserIds** | Pointer to **[]string** |  | [optional] 
**Duration** | Pointer to [**NullableInt64RangeSelectorModel**](Int64RangeSelectorModel.md) |  | [optional] 

## Methods

### NewAutoTestResultHistorySelectApiModel

`func NewAutoTestResultHistorySelectApiModel() *AutoTestResultHistorySelectApiModel`

NewAutoTestResultHistorySelectApiModel instantiates a new AutoTestResultHistorySelectApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestResultHistorySelectApiModelWithDefaults

`func NewAutoTestResultHistorySelectApiModelWithDefaults() *AutoTestResultHistorySelectApiModel`

NewAutoTestResultHistorySelectApiModelWithDefaults instantiates a new AutoTestResultHistorySelectApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutcomes

`func (o *AutoTestResultHistorySelectApiModel) GetOutcomes() []AutotestResultOutcome`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *AutoTestResultHistorySelectApiModel) GetOutcomesOk() (*[]AutotestResultOutcome, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomes

`func (o *AutoTestResultHistorySelectApiModel) SetOutcomes(v []AutotestResultOutcome)`

SetOutcomes sets Outcomes field to given value.

### HasOutcomes

`func (o *AutoTestResultHistorySelectApiModel) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.

### SetOutcomesNil

`func (o *AutoTestResultHistorySelectApiModel) SetOutcomesNil(b bool)`

 SetOutcomesNil sets the value for Outcomes to be an explicit nil

### UnsetOutcomes
`func (o *AutoTestResultHistorySelectApiModel) UnsetOutcomes()`

UnsetOutcomes ensures that no value is present for Outcomes, not even an explicit nil
### GetStatusCodes

`func (o *AutoTestResultHistorySelectApiModel) GetStatusCodes() []string`

GetStatusCodes returns the StatusCodes field if non-nil, zero value otherwise.

### GetStatusCodesOk

`func (o *AutoTestResultHistorySelectApiModel) GetStatusCodesOk() (*[]string, bool)`

GetStatusCodesOk returns a tuple with the StatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodes

`func (o *AutoTestResultHistorySelectApiModel) SetStatusCodes(v []string)`

SetStatusCodes sets StatusCodes field to given value.

### HasStatusCodes

`func (o *AutoTestResultHistorySelectApiModel) HasStatusCodes() bool`

HasStatusCodes returns a boolean if a field has been set.

### SetStatusCodesNil

`func (o *AutoTestResultHistorySelectApiModel) SetStatusCodesNil(b bool)`

 SetStatusCodesNil sets the value for StatusCodes to be an explicit nil

### UnsetStatusCodes
`func (o *AutoTestResultHistorySelectApiModel) UnsetStatusCodes()`

UnsetStatusCodes ensures that no value is present for StatusCodes, not even an explicit nil
### GetTestPlanIds

`func (o *AutoTestResultHistorySelectApiModel) GetTestPlanIds() []string`

GetTestPlanIds returns the TestPlanIds field if non-nil, zero value otherwise.

### GetTestPlanIdsOk

`func (o *AutoTestResultHistorySelectApiModel) GetTestPlanIdsOk() (*[]string, bool)`

GetTestPlanIdsOk returns a tuple with the TestPlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanIds

`func (o *AutoTestResultHistorySelectApiModel) SetTestPlanIds(v []string)`

SetTestPlanIds sets TestPlanIds field to given value.

### HasTestPlanIds

`func (o *AutoTestResultHistorySelectApiModel) HasTestPlanIds() bool`

HasTestPlanIds returns a boolean if a field has been set.

### SetTestPlanIdsNil

`func (o *AutoTestResultHistorySelectApiModel) SetTestPlanIdsNil(b bool)`

 SetTestPlanIdsNil sets the value for TestPlanIds to be an explicit nil

### UnsetTestPlanIds
`func (o *AutoTestResultHistorySelectApiModel) UnsetTestPlanIds()`

UnsetTestPlanIds ensures that no value is present for TestPlanIds, not even an explicit nil
### GetTestRunIds

`func (o *AutoTestResultHistorySelectApiModel) GetTestRunIds() []string`

GetTestRunIds returns the TestRunIds field if non-nil, zero value otherwise.

### GetTestRunIdsOk

`func (o *AutoTestResultHistorySelectApiModel) GetTestRunIdsOk() (*[]string, bool)`

GetTestRunIdsOk returns a tuple with the TestRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunIds

`func (o *AutoTestResultHistorySelectApiModel) SetTestRunIds(v []string)`

SetTestRunIds sets TestRunIds field to given value.

### HasTestRunIds

`func (o *AutoTestResultHistorySelectApiModel) HasTestRunIds() bool`

HasTestRunIds returns a boolean if a field has been set.

### SetTestRunIdsNil

`func (o *AutoTestResultHistorySelectApiModel) SetTestRunIdsNil(b bool)`

 SetTestRunIdsNil sets the value for TestRunIds to be an explicit nil

### UnsetTestRunIds
`func (o *AutoTestResultHistorySelectApiModel) UnsetTestRunIds()`

UnsetTestRunIds ensures that no value is present for TestRunIds, not even an explicit nil
### GetConfigurationIds

`func (o *AutoTestResultHistorySelectApiModel) GetConfigurationIds() []string`

GetConfigurationIds returns the ConfigurationIds field if non-nil, zero value otherwise.

### GetConfigurationIdsOk

`func (o *AutoTestResultHistorySelectApiModel) GetConfigurationIdsOk() (*[]string, bool)`

GetConfigurationIdsOk returns a tuple with the ConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIds

`func (o *AutoTestResultHistorySelectApiModel) SetConfigurationIds(v []string)`

SetConfigurationIds sets ConfigurationIds field to given value.

### HasConfigurationIds

`func (o *AutoTestResultHistorySelectApiModel) HasConfigurationIds() bool`

HasConfigurationIds returns a boolean if a field has been set.

### SetConfigurationIdsNil

`func (o *AutoTestResultHistorySelectApiModel) SetConfigurationIdsNil(b bool)`

 SetConfigurationIdsNil sets the value for ConfigurationIds to be an explicit nil

### UnsetConfigurationIds
`func (o *AutoTestResultHistorySelectApiModel) UnsetConfigurationIds()`

UnsetConfigurationIds ensures that no value is present for ConfigurationIds, not even an explicit nil
### GetLaunchSource

`func (o *AutoTestResultHistorySelectApiModel) GetLaunchSource() string`

GetLaunchSource returns the LaunchSource field if non-nil, zero value otherwise.

### GetLaunchSourceOk

`func (o *AutoTestResultHistorySelectApiModel) GetLaunchSourceOk() (*string, bool)`

GetLaunchSourceOk returns a tuple with the LaunchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSource

`func (o *AutoTestResultHistorySelectApiModel) SetLaunchSource(v string)`

SetLaunchSource sets LaunchSource field to given value.

### HasLaunchSource

`func (o *AutoTestResultHistorySelectApiModel) HasLaunchSource() bool`

HasLaunchSource returns a boolean if a field has been set.

### SetLaunchSourceNil

`func (o *AutoTestResultHistorySelectApiModel) SetLaunchSourceNil(b bool)`

 SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil

### UnsetLaunchSource
`func (o *AutoTestResultHistorySelectApiModel) UnsetLaunchSource()`

UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
### GetUserIds

`func (o *AutoTestResultHistorySelectApiModel) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *AutoTestResultHistorySelectApiModel) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *AutoTestResultHistorySelectApiModel) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *AutoTestResultHistorySelectApiModel) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### SetUserIdsNil

`func (o *AutoTestResultHistorySelectApiModel) SetUserIdsNil(b bool)`

 SetUserIdsNil sets the value for UserIds to be an explicit nil

### UnsetUserIds
`func (o *AutoTestResultHistorySelectApiModel) UnsetUserIds()`

UnsetUserIds ensures that no value is present for UserIds, not even an explicit nil
### GetDuration

`func (o *AutoTestResultHistorySelectApiModel) GetDuration() Int64RangeSelectorModel`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AutoTestResultHistorySelectApiModel) GetDurationOk() (*Int64RangeSelectorModel, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AutoTestResultHistorySelectApiModel) SetDuration(v Int64RangeSelectorModel)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AutoTestResultHistorySelectApiModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *AutoTestResultHistorySelectApiModel) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *AutoTestResultHistorySelectApiModel) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


