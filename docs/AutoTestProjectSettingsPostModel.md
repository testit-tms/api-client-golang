# AutoTestProjectSettingsPostModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsFlakyAuto** | Pointer to **bool** | Indicates if the status \&quot;Flaky/Stable\&quot; sets automatically | [optional] [default to false]
**FlakyStabilityPercentage** | Pointer to **int32** | Stability percentage for autotest flaky computing | [optional] [default to 100]
**FlakyTestRunCount** | Pointer to **int32** | Last test run count for autotest flaky computing | [optional] [default to 100]
**RerunEnabled** | **bool** | Auto rerun enabled | 
**RerunAttemptsCount** | **int32** | Auto rerun attempt count | 

## Methods

### NewAutoTestProjectSettingsPostModel

`func NewAutoTestProjectSettingsPostModel(rerunEnabled bool, rerunAttemptsCount int32, ) *AutoTestProjectSettingsPostModel`

NewAutoTestProjectSettingsPostModel instantiates a new AutoTestProjectSettingsPostModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestProjectSettingsPostModelWithDefaults

`func NewAutoTestProjectSettingsPostModelWithDefaults() *AutoTestProjectSettingsPostModel`

NewAutoTestProjectSettingsPostModelWithDefaults instantiates a new AutoTestProjectSettingsPostModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsFlakyAuto

`func (o *AutoTestProjectSettingsPostModel) GetIsFlakyAuto() bool`

GetIsFlakyAuto returns the IsFlakyAuto field if non-nil, zero value otherwise.

### GetIsFlakyAutoOk

`func (o *AutoTestProjectSettingsPostModel) GetIsFlakyAutoOk() (*bool, bool)`

GetIsFlakyAutoOk returns a tuple with the IsFlakyAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlakyAuto

`func (o *AutoTestProjectSettingsPostModel) SetIsFlakyAuto(v bool)`

SetIsFlakyAuto sets IsFlakyAuto field to given value.

### HasIsFlakyAuto

`func (o *AutoTestProjectSettingsPostModel) HasIsFlakyAuto() bool`

HasIsFlakyAuto returns a boolean if a field has been set.

### GetFlakyStabilityPercentage

`func (o *AutoTestProjectSettingsPostModel) GetFlakyStabilityPercentage() int32`

GetFlakyStabilityPercentage returns the FlakyStabilityPercentage field if non-nil, zero value otherwise.

### GetFlakyStabilityPercentageOk

`func (o *AutoTestProjectSettingsPostModel) GetFlakyStabilityPercentageOk() (*int32, bool)`

GetFlakyStabilityPercentageOk returns a tuple with the FlakyStabilityPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlakyStabilityPercentage

`func (o *AutoTestProjectSettingsPostModel) SetFlakyStabilityPercentage(v int32)`

SetFlakyStabilityPercentage sets FlakyStabilityPercentage field to given value.

### HasFlakyStabilityPercentage

`func (o *AutoTestProjectSettingsPostModel) HasFlakyStabilityPercentage() bool`

HasFlakyStabilityPercentage returns a boolean if a field has been set.

### GetFlakyTestRunCount

`func (o *AutoTestProjectSettingsPostModel) GetFlakyTestRunCount() int32`

GetFlakyTestRunCount returns the FlakyTestRunCount field if non-nil, zero value otherwise.

### GetFlakyTestRunCountOk

`func (o *AutoTestProjectSettingsPostModel) GetFlakyTestRunCountOk() (*int32, bool)`

GetFlakyTestRunCountOk returns a tuple with the FlakyTestRunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlakyTestRunCount

`func (o *AutoTestProjectSettingsPostModel) SetFlakyTestRunCount(v int32)`

SetFlakyTestRunCount sets FlakyTestRunCount field to given value.

### HasFlakyTestRunCount

`func (o *AutoTestProjectSettingsPostModel) HasFlakyTestRunCount() bool`

HasFlakyTestRunCount returns a boolean if a field has been set.

### GetRerunEnabled

`func (o *AutoTestProjectSettingsPostModel) GetRerunEnabled() bool`

GetRerunEnabled returns the RerunEnabled field if non-nil, zero value otherwise.

### GetRerunEnabledOk

`func (o *AutoTestProjectSettingsPostModel) GetRerunEnabledOk() (*bool, bool)`

GetRerunEnabledOk returns a tuple with the RerunEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerunEnabled

`func (o *AutoTestProjectSettingsPostModel) SetRerunEnabled(v bool)`

SetRerunEnabled sets RerunEnabled field to given value.


### GetRerunAttemptsCount

`func (o *AutoTestProjectSettingsPostModel) GetRerunAttemptsCount() int32`

GetRerunAttemptsCount returns the RerunAttemptsCount field if non-nil, zero value otherwise.

### GetRerunAttemptsCountOk

`func (o *AutoTestProjectSettingsPostModel) GetRerunAttemptsCountOk() (*int32, bool)`

GetRerunAttemptsCountOk returns a tuple with the RerunAttemptsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerunAttemptsCount

`func (o *AutoTestProjectSettingsPostModel) SetRerunAttemptsCount(v int32)`

SetRerunAttemptsCount sets RerunAttemptsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


