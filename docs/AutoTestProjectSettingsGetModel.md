# AutoTestProjectSettingsGetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | Unique ID of the project. | 
**IsFlakyAuto** | Pointer to **bool** | Indicates if the status \&quot;Flaky/Stable\&quot; sets automatically | [optional] [default to false]
**FlakyStabilityPercentage** | Pointer to **int32** | Stability percentage for autotest flaky computing | [optional] [default to 100]
**FlakyTestRunCount** | Pointer to **int32** | Last test run count for autotest flaky computing | [optional] [default to 100]
**RerunEnabled** | **bool** | Auto rerun enabled | 
**RerunAttemptsCount** | **int32** | Auto rerun attempt count | 

## Methods

### NewAutoTestProjectSettingsGetModel

`func NewAutoTestProjectSettingsGetModel(projectId string, rerunEnabled bool, rerunAttemptsCount int32, ) *AutoTestProjectSettingsGetModel`

NewAutoTestProjectSettingsGetModel instantiates a new AutoTestProjectSettingsGetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestProjectSettingsGetModelWithDefaults

`func NewAutoTestProjectSettingsGetModelWithDefaults() *AutoTestProjectSettingsGetModel`

NewAutoTestProjectSettingsGetModelWithDefaults instantiates a new AutoTestProjectSettingsGetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *AutoTestProjectSettingsGetModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AutoTestProjectSettingsGetModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AutoTestProjectSettingsGetModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetIsFlakyAuto

`func (o *AutoTestProjectSettingsGetModel) GetIsFlakyAuto() bool`

GetIsFlakyAuto returns the IsFlakyAuto field if non-nil, zero value otherwise.

### GetIsFlakyAutoOk

`func (o *AutoTestProjectSettingsGetModel) GetIsFlakyAutoOk() (*bool, bool)`

GetIsFlakyAutoOk returns a tuple with the IsFlakyAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlakyAuto

`func (o *AutoTestProjectSettingsGetModel) SetIsFlakyAuto(v bool)`

SetIsFlakyAuto sets IsFlakyAuto field to given value.

### HasIsFlakyAuto

`func (o *AutoTestProjectSettingsGetModel) HasIsFlakyAuto() bool`

HasIsFlakyAuto returns a boolean if a field has been set.

### GetFlakyStabilityPercentage

`func (o *AutoTestProjectSettingsGetModel) GetFlakyStabilityPercentage() int32`

GetFlakyStabilityPercentage returns the FlakyStabilityPercentage field if non-nil, zero value otherwise.

### GetFlakyStabilityPercentageOk

`func (o *AutoTestProjectSettingsGetModel) GetFlakyStabilityPercentageOk() (*int32, bool)`

GetFlakyStabilityPercentageOk returns a tuple with the FlakyStabilityPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlakyStabilityPercentage

`func (o *AutoTestProjectSettingsGetModel) SetFlakyStabilityPercentage(v int32)`

SetFlakyStabilityPercentage sets FlakyStabilityPercentage field to given value.

### HasFlakyStabilityPercentage

`func (o *AutoTestProjectSettingsGetModel) HasFlakyStabilityPercentage() bool`

HasFlakyStabilityPercentage returns a boolean if a field has been set.

### GetFlakyTestRunCount

`func (o *AutoTestProjectSettingsGetModel) GetFlakyTestRunCount() int32`

GetFlakyTestRunCount returns the FlakyTestRunCount field if non-nil, zero value otherwise.

### GetFlakyTestRunCountOk

`func (o *AutoTestProjectSettingsGetModel) GetFlakyTestRunCountOk() (*int32, bool)`

GetFlakyTestRunCountOk returns a tuple with the FlakyTestRunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlakyTestRunCount

`func (o *AutoTestProjectSettingsGetModel) SetFlakyTestRunCount(v int32)`

SetFlakyTestRunCount sets FlakyTestRunCount field to given value.

### HasFlakyTestRunCount

`func (o *AutoTestProjectSettingsGetModel) HasFlakyTestRunCount() bool`

HasFlakyTestRunCount returns a boolean if a field has been set.

### GetRerunEnabled

`func (o *AutoTestProjectSettingsGetModel) GetRerunEnabled() bool`

GetRerunEnabled returns the RerunEnabled field if non-nil, zero value otherwise.

### GetRerunEnabledOk

`func (o *AutoTestProjectSettingsGetModel) GetRerunEnabledOk() (*bool, bool)`

GetRerunEnabledOk returns a tuple with the RerunEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerunEnabled

`func (o *AutoTestProjectSettingsGetModel) SetRerunEnabled(v bool)`

SetRerunEnabled sets RerunEnabled field to given value.


### GetRerunAttemptsCount

`func (o *AutoTestProjectSettingsGetModel) GetRerunAttemptsCount() int32`

GetRerunAttemptsCount returns the RerunAttemptsCount field if non-nil, zero value otherwise.

### GetRerunAttemptsCountOk

`func (o *AutoTestProjectSettingsGetModel) GetRerunAttemptsCountOk() (*int32, bool)`

GetRerunAttemptsCountOk returns a tuple with the RerunAttemptsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerunAttemptsCount

`func (o *AutoTestProjectSettingsGetModel) SetRerunAttemptsCount(v int32)`

SetRerunAttemptsCount sets RerunAttemptsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


