# AutoTestProjectSettingsApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsFlakyAuto** | Pointer to **bool** | Indicates if the status \&quot;Flaky/Stable\&quot; sets automatically | [optional] [default to false]
**FlakyStabilityPercentage** | Pointer to **int32** | Stability percentage for autotest flaky computing | [optional] [default to 100]
**FlakyTestRunCount** | Pointer to **int32** | Last test run count for autotest flaky computing | [optional] [default to 100]
**RerunEnabled** | **bool** | Auto rerun enabled | 
**RerunAttemptsCount** | **int32** | Auto rerun attempt count | 
**WorkItemUpdatingEnabled** | Pointer to **bool** | Autotest to work item updating enabled | [optional] [default to false]
**WorkItemUpdatingFields** | [**WorkItemUpdatingFieldsApiModel**](WorkItemUpdatingFieldsApiModel.md) | Autotest to work item updating fields | 

## Methods

### NewAutoTestProjectSettingsApiModel

`func NewAutoTestProjectSettingsApiModel(rerunEnabled bool, rerunAttemptsCount int32, workItemUpdatingFields WorkItemUpdatingFieldsApiModel, ) *AutoTestProjectSettingsApiModel`

NewAutoTestProjectSettingsApiModel instantiates a new AutoTestProjectSettingsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestProjectSettingsApiModelWithDefaults

`func NewAutoTestProjectSettingsApiModelWithDefaults() *AutoTestProjectSettingsApiModel`

NewAutoTestProjectSettingsApiModelWithDefaults instantiates a new AutoTestProjectSettingsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsFlakyAuto

`func (o *AutoTestProjectSettingsApiModel) GetIsFlakyAuto() bool`

GetIsFlakyAuto returns the IsFlakyAuto field if non-nil, zero value otherwise.

### GetIsFlakyAutoOk

`func (o *AutoTestProjectSettingsApiModel) GetIsFlakyAutoOk() (*bool, bool)`

GetIsFlakyAutoOk returns a tuple with the IsFlakyAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlakyAuto

`func (o *AutoTestProjectSettingsApiModel) SetIsFlakyAuto(v bool)`

SetIsFlakyAuto sets IsFlakyAuto field to given value.

### HasIsFlakyAuto

`func (o *AutoTestProjectSettingsApiModel) HasIsFlakyAuto() bool`

HasIsFlakyAuto returns a boolean if a field has been set.

### GetFlakyStabilityPercentage

`func (o *AutoTestProjectSettingsApiModel) GetFlakyStabilityPercentage() int32`

GetFlakyStabilityPercentage returns the FlakyStabilityPercentage field if non-nil, zero value otherwise.

### GetFlakyStabilityPercentageOk

`func (o *AutoTestProjectSettingsApiModel) GetFlakyStabilityPercentageOk() (*int32, bool)`

GetFlakyStabilityPercentageOk returns a tuple with the FlakyStabilityPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlakyStabilityPercentage

`func (o *AutoTestProjectSettingsApiModel) SetFlakyStabilityPercentage(v int32)`

SetFlakyStabilityPercentage sets FlakyStabilityPercentage field to given value.

### HasFlakyStabilityPercentage

`func (o *AutoTestProjectSettingsApiModel) HasFlakyStabilityPercentage() bool`

HasFlakyStabilityPercentage returns a boolean if a field has been set.

### GetFlakyTestRunCount

`func (o *AutoTestProjectSettingsApiModel) GetFlakyTestRunCount() int32`

GetFlakyTestRunCount returns the FlakyTestRunCount field if non-nil, zero value otherwise.

### GetFlakyTestRunCountOk

`func (o *AutoTestProjectSettingsApiModel) GetFlakyTestRunCountOk() (*int32, bool)`

GetFlakyTestRunCountOk returns a tuple with the FlakyTestRunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlakyTestRunCount

`func (o *AutoTestProjectSettingsApiModel) SetFlakyTestRunCount(v int32)`

SetFlakyTestRunCount sets FlakyTestRunCount field to given value.

### HasFlakyTestRunCount

`func (o *AutoTestProjectSettingsApiModel) HasFlakyTestRunCount() bool`

HasFlakyTestRunCount returns a boolean if a field has been set.

### GetRerunEnabled

`func (o *AutoTestProjectSettingsApiModel) GetRerunEnabled() bool`

GetRerunEnabled returns the RerunEnabled field if non-nil, zero value otherwise.

### GetRerunEnabledOk

`func (o *AutoTestProjectSettingsApiModel) GetRerunEnabledOk() (*bool, bool)`

GetRerunEnabledOk returns a tuple with the RerunEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerunEnabled

`func (o *AutoTestProjectSettingsApiModel) SetRerunEnabled(v bool)`

SetRerunEnabled sets RerunEnabled field to given value.


### GetRerunAttemptsCount

`func (o *AutoTestProjectSettingsApiModel) GetRerunAttemptsCount() int32`

GetRerunAttemptsCount returns the RerunAttemptsCount field if non-nil, zero value otherwise.

### GetRerunAttemptsCountOk

`func (o *AutoTestProjectSettingsApiModel) GetRerunAttemptsCountOk() (*int32, bool)`

GetRerunAttemptsCountOk returns a tuple with the RerunAttemptsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerunAttemptsCount

`func (o *AutoTestProjectSettingsApiModel) SetRerunAttemptsCount(v int32)`

SetRerunAttemptsCount sets RerunAttemptsCount field to given value.


### GetWorkItemUpdatingEnabled

`func (o *AutoTestProjectSettingsApiModel) GetWorkItemUpdatingEnabled() bool`

GetWorkItemUpdatingEnabled returns the WorkItemUpdatingEnabled field if non-nil, zero value otherwise.

### GetWorkItemUpdatingEnabledOk

`func (o *AutoTestProjectSettingsApiModel) GetWorkItemUpdatingEnabledOk() (*bool, bool)`

GetWorkItemUpdatingEnabledOk returns a tuple with the WorkItemUpdatingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemUpdatingEnabled

`func (o *AutoTestProjectSettingsApiModel) SetWorkItemUpdatingEnabled(v bool)`

SetWorkItemUpdatingEnabled sets WorkItemUpdatingEnabled field to given value.

### HasWorkItemUpdatingEnabled

`func (o *AutoTestProjectSettingsApiModel) HasWorkItemUpdatingEnabled() bool`

HasWorkItemUpdatingEnabled returns a boolean if a field has been set.

### GetWorkItemUpdatingFields

`func (o *AutoTestProjectSettingsApiModel) GetWorkItemUpdatingFields() WorkItemUpdatingFieldsApiModel`

GetWorkItemUpdatingFields returns the WorkItemUpdatingFields field if non-nil, zero value otherwise.

### GetWorkItemUpdatingFieldsOk

`func (o *AutoTestProjectSettingsApiModel) GetWorkItemUpdatingFieldsOk() (*WorkItemUpdatingFieldsApiModel, bool)`

GetWorkItemUpdatingFieldsOk returns a tuple with the WorkItemUpdatingFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemUpdatingFields

`func (o *AutoTestProjectSettingsApiModel) SetWorkItemUpdatingFields(v WorkItemUpdatingFieldsApiModel)`

SetWorkItemUpdatingFields sets WorkItemUpdatingFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


