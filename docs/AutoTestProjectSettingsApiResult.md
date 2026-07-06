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
**WorkItemUpdatingEnabled** | **bool** | Autotest to work item updating enabled | 
**WorkItemUpdatingFields** | [**WorkItemUpdatingFieldsApiResult**](WorkItemUpdatingFieldsApiResult.md) | Autotest to work item updating fields | 
**ArchiveOutdatedTestRunsEnabled** | **bool** | Indicates whether archiving of outdated test runs is enabled for the project. | 
**TestRunsArchiveLimitEnabled** | **bool** | Indicates whether a limit is enforced on the number of archived test runs. | 
**TestRunsRetentionPeriodDays** | **int32** |  The retention period in days for test runs. After this period,  outdated test runs may be archived based on project settings | 
**MaxActiveTestRunsCount** | **int32** | Maximum number of active test runs to keep. When this limit is exceeded,  older test runs are automatically archived | 

## Methods

### NewAutoTestProjectSettingsApiResult

`func NewAutoTestProjectSettingsApiResult(projectId string, isFlakyAuto bool, flakyStabilityPercentage int32, flakyTestRunCount int32, rerunEnabled bool, rerunAttemptsCount int32, workItemUpdatingEnabled bool, workItemUpdatingFields WorkItemUpdatingFieldsApiResult, archiveOutdatedTestRunsEnabled bool, testRunsArchiveLimitEnabled bool, testRunsRetentionPeriodDays int32, maxActiveTestRunsCount int32, ) *AutoTestProjectSettingsApiResult`

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


### GetWorkItemUpdatingEnabled

`func (o *AutoTestProjectSettingsApiResult) GetWorkItemUpdatingEnabled() bool`

GetWorkItemUpdatingEnabled returns the WorkItemUpdatingEnabled field if non-nil, zero value otherwise.

### GetWorkItemUpdatingEnabledOk

`func (o *AutoTestProjectSettingsApiResult) GetWorkItemUpdatingEnabledOk() (*bool, bool)`

GetWorkItemUpdatingEnabledOk returns a tuple with the WorkItemUpdatingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemUpdatingEnabled

`func (o *AutoTestProjectSettingsApiResult) SetWorkItemUpdatingEnabled(v bool)`

SetWorkItemUpdatingEnabled sets WorkItemUpdatingEnabled field to given value.


### GetWorkItemUpdatingFields

`func (o *AutoTestProjectSettingsApiResult) GetWorkItemUpdatingFields() WorkItemUpdatingFieldsApiResult`

GetWorkItemUpdatingFields returns the WorkItemUpdatingFields field if non-nil, zero value otherwise.

### GetWorkItemUpdatingFieldsOk

`func (o *AutoTestProjectSettingsApiResult) GetWorkItemUpdatingFieldsOk() (*WorkItemUpdatingFieldsApiResult, bool)`

GetWorkItemUpdatingFieldsOk returns a tuple with the WorkItemUpdatingFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemUpdatingFields

`func (o *AutoTestProjectSettingsApiResult) SetWorkItemUpdatingFields(v WorkItemUpdatingFieldsApiResult)`

SetWorkItemUpdatingFields sets WorkItemUpdatingFields field to given value.


### GetArchiveOutdatedTestRunsEnabled

`func (o *AutoTestProjectSettingsApiResult) GetArchiveOutdatedTestRunsEnabled() bool`

GetArchiveOutdatedTestRunsEnabled returns the ArchiveOutdatedTestRunsEnabled field if non-nil, zero value otherwise.

### GetArchiveOutdatedTestRunsEnabledOk

`func (o *AutoTestProjectSettingsApiResult) GetArchiveOutdatedTestRunsEnabledOk() (*bool, bool)`

GetArchiveOutdatedTestRunsEnabledOk returns a tuple with the ArchiveOutdatedTestRunsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveOutdatedTestRunsEnabled

`func (o *AutoTestProjectSettingsApiResult) SetArchiveOutdatedTestRunsEnabled(v bool)`

SetArchiveOutdatedTestRunsEnabled sets ArchiveOutdatedTestRunsEnabled field to given value.


### GetTestRunsArchiveLimitEnabled

`func (o *AutoTestProjectSettingsApiResult) GetTestRunsArchiveLimitEnabled() bool`

GetTestRunsArchiveLimitEnabled returns the TestRunsArchiveLimitEnabled field if non-nil, zero value otherwise.

### GetTestRunsArchiveLimitEnabledOk

`func (o *AutoTestProjectSettingsApiResult) GetTestRunsArchiveLimitEnabledOk() (*bool, bool)`

GetTestRunsArchiveLimitEnabledOk returns a tuple with the TestRunsArchiveLimitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunsArchiveLimitEnabled

`func (o *AutoTestProjectSettingsApiResult) SetTestRunsArchiveLimitEnabled(v bool)`

SetTestRunsArchiveLimitEnabled sets TestRunsArchiveLimitEnabled field to given value.


### GetTestRunsRetentionPeriodDays

`func (o *AutoTestProjectSettingsApiResult) GetTestRunsRetentionPeriodDays() int32`

GetTestRunsRetentionPeriodDays returns the TestRunsRetentionPeriodDays field if non-nil, zero value otherwise.

### GetTestRunsRetentionPeriodDaysOk

`func (o *AutoTestProjectSettingsApiResult) GetTestRunsRetentionPeriodDaysOk() (*int32, bool)`

GetTestRunsRetentionPeriodDaysOk returns a tuple with the TestRunsRetentionPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunsRetentionPeriodDays

`func (o *AutoTestProjectSettingsApiResult) SetTestRunsRetentionPeriodDays(v int32)`

SetTestRunsRetentionPeriodDays sets TestRunsRetentionPeriodDays field to given value.


### GetMaxActiveTestRunsCount

`func (o *AutoTestProjectSettingsApiResult) GetMaxActiveTestRunsCount() int32`

GetMaxActiveTestRunsCount returns the MaxActiveTestRunsCount field if non-nil, zero value otherwise.

### GetMaxActiveTestRunsCountOk

`func (o *AutoTestProjectSettingsApiResult) GetMaxActiveTestRunsCountOk() (*int32, bool)`

GetMaxActiveTestRunsCountOk returns a tuple with the MaxActiveTestRunsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxActiveTestRunsCount

`func (o *AutoTestProjectSettingsApiResult) SetMaxActiveTestRunsCount(v int32)`

SetMaxActiveTestRunsCount sets MaxActiveTestRunsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


