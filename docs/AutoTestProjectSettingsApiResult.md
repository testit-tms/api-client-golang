# AutoTestProjectSettingsApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | Unique ID of the project. | 
**IsFlakyAuto** | **bool** | Indicates if the status \&quot;Flaky/Stable\&quot; sets automatically | 
**FlakyStabilityPercentage** | **int32** | Stability percentage for autotest flaky computing | 
**FlakyTestRunCount** | **int32** | Last test run count for autotest flaky computing | 
**RerunEnabled** | **bool** | Auto rerun enabled | 
**RerunAttemptsCount** | **int32** | Auto rerun attempt count | 

## Methods

### NewAutoTestProjectSettingsApiResult

`func NewAutoTestProjectSettingsApiResult(projectId string, isFlakyAuto bool, flakyStabilityPercentage int32, flakyTestRunCount int32, rerunEnabled bool, rerunAttemptsCount int32, ) *AutoTestProjectSettingsApiResult`

NewAutoTestProjectSettingsApiResult instantiates a new AutoTestProjectSettingsApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestProjectSettingsApiResultWithDefaults

`func NewAutoTestProjectSettingsApiResultWithDefaults() *AutoTestProjectSettingsApiResult`

NewAutoTestProjectSettingsApiResultWithDefaults instantiates a new AutoTestProjectSettingsApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *AutoTestProjectSettingsApiResult) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AutoTestProjectSettingsApiResult) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AutoTestProjectSettingsApiResult) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetIsFlakyAuto

`func (o *AutoTestProjectSettingsApiResult) GetIsFlakyAuto() bool`

GetIsFlakyAuto returns the IsFlakyAuto field if non-nil, zero value otherwise.

### GetIsFlakyAutoOk

`func (o *AutoTestProjectSettingsApiResult) GetIsFlakyAutoOk() (*bool, bool)`

GetIsFlakyAutoOk returns a tuple with the IsFlakyAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlakyAuto

`func (o *AutoTestProjectSettingsApiResult) SetIsFlakyAuto(v bool)`

SetIsFlakyAuto sets IsFlakyAuto field to given value.


### GetFlakyStabilityPercentage

`func (o *AutoTestProjectSettingsApiResult) GetFlakyStabilityPercentage() int32`

GetFlakyStabilityPercentage returns the FlakyStabilityPercentage field if non-nil, zero value otherwise.

### GetFlakyStabilityPercentageOk

`func (o *AutoTestProjectSettingsApiResult) GetFlakyStabilityPercentageOk() (*int32, bool)`

GetFlakyStabilityPercentageOk returns a tuple with the FlakyStabilityPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlakyStabilityPercentage

`func (o *AutoTestProjectSettingsApiResult) SetFlakyStabilityPercentage(v int32)`

SetFlakyStabilityPercentage sets FlakyStabilityPercentage field to given value.


### GetFlakyTestRunCount

`func (o *AutoTestProjectSettingsApiResult) GetFlakyTestRunCount() int32`

GetFlakyTestRunCount returns the FlakyTestRunCount field if non-nil, zero value otherwise.

### GetFlakyTestRunCountOk

`func (o *AutoTestProjectSettingsApiResult) GetFlakyTestRunCountOk() (*int32, bool)`

GetFlakyTestRunCountOk returns a tuple with the FlakyTestRunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlakyTestRunCount

`func (o *AutoTestProjectSettingsApiResult) SetFlakyTestRunCount(v int32)`

SetFlakyTestRunCount sets FlakyTestRunCount field to given value.


### GetRerunEnabled

`func (o *AutoTestProjectSettingsApiResult) GetRerunEnabled() bool`

GetRerunEnabled returns the RerunEnabled field if non-nil, zero value otherwise.

### GetRerunEnabledOk

`func (o *AutoTestProjectSettingsApiResult) GetRerunEnabledOk() (*bool, bool)`

GetRerunEnabledOk returns a tuple with the RerunEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerunEnabled

`func (o *AutoTestProjectSettingsApiResult) SetRerunEnabled(v bool)`

SetRerunEnabled sets RerunEnabled field to given value.


### GetRerunAttemptsCount

`func (o *AutoTestProjectSettingsApiResult) GetRerunAttemptsCount() int32`

GetRerunAttemptsCount returns the RerunAttemptsCount field if non-nil, zero value otherwise.

### GetRerunAttemptsCountOk

`func (o *AutoTestProjectSettingsApiResult) GetRerunAttemptsCountOk() (*int32, bool)`

GetRerunAttemptsCountOk returns a tuple with the RerunAttemptsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerunAttemptsCount

`func (o *AutoTestProjectSettingsApiResult) SetRerunAttemptsCount(v int32)`

SetRerunAttemptsCount sets RerunAttemptsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


