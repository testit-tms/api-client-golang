# TestPlanTestPointsSearchApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestSuiteIds** | Pointer to **[]string** | Specifies a test point test suite IDs to search for | [optional] 
**WorkItemGlobalIds** | Pointer to **[]int64** | Specifies a test point work item global IDs to search for | [optional] 
**WorkItemMedianDuration** | Pointer to [**NullableInt64RangeSelectorModel**](Int64RangeSelectorModel.md) | Specifies a test point work item median duration range to search for | [optional] 
**Statuses** | Pointer to [**[]TestPointStatus**](TestPointStatus.md) | Specifies a test point statuses to search for | [optional] 
**StatusCodes** | Pointer to **[]string** | Specifies a test point status codes to search for | [optional] 
**Priorities** | Pointer to [**[]WorkItemPriorityModel**](WorkItemPriorityModel.md) | Specifies a test point priorities to search for | [optional] 
**IsAutomated** | Pointer to **NullableBool** | Specifies a test point automation status to search for | [optional] 
**Name** | Pointer to **NullableString** | Specifies a test point name to search for | [optional] 
**ConfigurationIds** | Pointer to **[]string** | Specifies a test point configuration IDs to search for | [optional] 
**TesterIds** | Pointer to **[]string** | Specifies a test point assigned user IDs to search for | [optional] 
**Duration** | Pointer to [**NullableInt64RangeSelectorModel**](Int64RangeSelectorModel.md) | Specifies a test point range of duration to search for | [optional] 
**SectionIds** | Pointer to **[]string** | Specifies a test point work item section IDs to search for | [optional] 
**CreatedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) | Specifies a test point range of creation date to search for | [optional] 
**CreatedByIds** | Pointer to **[]string** | Specifies a test point creator IDs to search for | [optional] 
**ModifiedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) | Specifies a test point range of last modification date to search for | [optional] 
**ModifiedByIds** | Pointer to **[]string** | Specifies a test point last editor IDs to search for | [optional] 
**Tags** | Pointer to **[]string** | Specifies a test point tags to search for | [optional] 
**Attributes** | Pointer to **map[string][]string** | Specifies a test point attributes to search for | [optional] 
**WorkItemCreatedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) | Specifies a work item range of creation date to search for | [optional] 
**WorkItemCreatedByIds** | Pointer to **[]string** | Specifies a work item creator IDs to search for | [optional] 
**WorkItemModifiedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) | Specifies a work item range of last modification date to search for | [optional] 
**WorkItemModifiedByIds** | Pointer to **[]string** | Specifies a work item last editor IDs to search for | [optional] 

## Methods

### NewTestPlanTestPointsSearchApiModel

`func NewTestPlanTestPointsSearchApiModel() *TestPlanTestPointsSearchApiModel`

NewTestPlanTestPointsSearchApiModel instantiates a new TestPlanTestPointsSearchApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanTestPointsSearchApiModelWithDefaults

`func NewTestPlanTestPointsSearchApiModelWithDefaults() *TestPlanTestPointsSearchApiModel`

NewTestPlanTestPointsSearchApiModelWithDefaults instantiates a new TestPlanTestPointsSearchApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestSuiteIds

`func (o *TestPlanTestPointsSearchApiModel) GetTestSuiteIds() []string`

GetTestSuiteIds returns the TestSuiteIds field if non-nil, zero value otherwise.

### GetTestSuiteIdsOk

`func (o *TestPlanTestPointsSearchApiModel) GetTestSuiteIdsOk() (*[]string, bool)`

GetTestSuiteIdsOk returns a tuple with the TestSuiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuiteIds

`func (o *TestPlanTestPointsSearchApiModel) SetTestSuiteIds(v []string)`

SetTestSuiteIds sets TestSuiteIds field to given value.

### HasTestSuiteIds

`func (o *TestPlanTestPointsSearchApiModel) HasTestSuiteIds() bool`

HasTestSuiteIds returns a boolean if a field has been set.

### SetTestSuiteIdsNil

`func (o *TestPlanTestPointsSearchApiModel) SetTestSuiteIdsNil(b bool)`

 SetTestSuiteIdsNil sets the value for TestSuiteIds to be an explicit nil

### UnsetTestSuiteIds
`func (o *TestPlanTestPointsSearchApiModel) UnsetTestSuiteIds()`

UnsetTestSuiteIds ensures that no value is present for TestSuiteIds, not even an explicit nil
### GetWorkItemGlobalIds

`func (o *TestPlanTestPointsSearchApiModel) GetWorkItemGlobalIds() []int64`

GetWorkItemGlobalIds returns the WorkItemGlobalIds field if non-nil, zero value otherwise.

### GetWorkItemGlobalIdsOk

`func (o *TestPlanTestPointsSearchApiModel) GetWorkItemGlobalIdsOk() (*[]int64, bool)`

GetWorkItemGlobalIdsOk returns a tuple with the WorkItemGlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemGlobalIds

`func (o *TestPlanTestPointsSearchApiModel) SetWorkItemGlobalIds(v []int64)`

SetWorkItemGlobalIds sets WorkItemGlobalIds field to given value.

### HasWorkItemGlobalIds

`func (o *TestPlanTestPointsSearchApiModel) HasWorkItemGlobalIds() bool`

HasWorkItemGlobalIds returns a boolean if a field has been set.

### SetWorkItemGlobalIdsNil

`func (o *TestPlanTestPointsSearchApiModel) SetWorkItemGlobalIdsNil(b bool)`

 SetWorkItemGlobalIdsNil sets the value for WorkItemGlobalIds to be an explicit nil

### UnsetWorkItemGlobalIds
`func (o *TestPlanTestPointsSearchApiModel) UnsetWorkItemGlobalIds()`

UnsetWorkItemGlobalIds ensures that no value is present for WorkItemGlobalIds, not even an explicit nil
### GetWorkItemMedianDuration

`func (o *TestPlanTestPointsSearchApiModel) GetWorkItemMedianDuration() Int64RangeSelectorModel`

GetWorkItemMedianDuration returns the WorkItemMedianDuration field if non-nil, zero value otherwise.

### GetWorkItemMedianDurationOk

`func (o *TestPlanTestPointsSearchApiModel) GetWorkItemMedianDurationOk() (*Int64RangeSelectorModel, bool)`

GetWorkItemMedianDurationOk returns a tuple with the WorkItemMedianDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemMedianDuration

`func (o *TestPlanTestPointsSearchApiModel) SetWorkItemMedianDuration(v Int64RangeSelectorModel)`

SetWorkItemMedianDuration sets WorkItemMedianDuration field to given value.

### HasWorkItemMedianDuration

`func (o *TestPlanTestPointsSearchApiModel) HasWorkItemMedianDuration() bool`

HasWorkItemMedianDuration returns a boolean if a field has been set.

### SetWorkItemMedianDurationNil

`func (o *TestPlanTestPointsSearchApiModel) SetWorkItemMedianDurationNil(b bool)`

 SetWorkItemMedianDurationNil sets the value for WorkItemMedianDuration to be an explicit nil

### UnsetWorkItemMedianDuration
`func (o *TestPlanTestPointsSearchApiModel) UnsetWorkItemMedianDuration()`

UnsetWorkItemMedianDuration ensures that no value is present for WorkItemMedianDuration, not even an explicit nil
### GetStatuses

`func (o *TestPlanTestPointsSearchApiModel) GetStatuses() []TestPointStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *TestPlanTestPointsSearchApiModel) GetStatusesOk() (*[]TestPointStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *TestPlanTestPointsSearchApiModel) SetStatuses(v []TestPointStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *TestPlanTestPointsSearchApiModel) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### SetStatusesNil

`func (o *TestPlanTestPointsSearchApiModel) SetStatusesNil(b bool)`

 SetStatusesNil sets the value for Statuses to be an explicit nil

### UnsetStatuses
`func (o *TestPlanTestPointsSearchApiModel) UnsetStatuses()`

UnsetStatuses ensures that no value is present for Statuses, not even an explicit nil
### GetStatusCodes

`func (o *TestPlanTestPointsSearchApiModel) GetStatusCodes() []string`

GetStatusCodes returns the StatusCodes field if non-nil, zero value otherwise.

### GetStatusCodesOk

`func (o *TestPlanTestPointsSearchApiModel) GetStatusCodesOk() (*[]string, bool)`

GetStatusCodesOk returns a tuple with the StatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodes

`func (o *TestPlanTestPointsSearchApiModel) SetStatusCodes(v []string)`

SetStatusCodes sets StatusCodes field to given value.

### HasStatusCodes

`func (o *TestPlanTestPointsSearchApiModel) HasStatusCodes() bool`

HasStatusCodes returns a boolean if a field has been set.

### SetStatusCodesNil

`func (o *TestPlanTestPointsSearchApiModel) SetStatusCodesNil(b bool)`

 SetStatusCodesNil sets the value for StatusCodes to be an explicit nil

### UnsetStatusCodes
`func (o *TestPlanTestPointsSearchApiModel) UnsetStatusCodes()`

UnsetStatusCodes ensures that no value is present for StatusCodes, not even an explicit nil
### GetPriorities

`func (o *TestPlanTestPointsSearchApiModel) GetPriorities() []WorkItemPriorityModel`

GetPriorities returns the Priorities field if non-nil, zero value otherwise.

### GetPrioritiesOk

`func (o *TestPlanTestPointsSearchApiModel) GetPrioritiesOk() (*[]WorkItemPriorityModel, bool)`

GetPrioritiesOk returns a tuple with the Priorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorities

`func (o *TestPlanTestPointsSearchApiModel) SetPriorities(v []WorkItemPriorityModel)`

SetPriorities sets Priorities field to given value.

### HasPriorities

`func (o *TestPlanTestPointsSearchApiModel) HasPriorities() bool`

HasPriorities returns a boolean if a field has been set.

### SetPrioritiesNil

`func (o *TestPlanTestPointsSearchApiModel) SetPrioritiesNil(b bool)`

 SetPrioritiesNil sets the value for Priorities to be an explicit nil

### UnsetPriorities
`func (o *TestPlanTestPointsSearchApiModel) UnsetPriorities()`

UnsetPriorities ensures that no value is present for Priorities, not even an explicit nil
### GetIsAutomated

`func (o *TestPlanTestPointsSearchApiModel) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *TestPlanTestPointsSearchApiModel) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *TestPlanTestPointsSearchApiModel) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *TestPlanTestPointsSearchApiModel) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### SetIsAutomatedNil

`func (o *TestPlanTestPointsSearchApiModel) SetIsAutomatedNil(b bool)`

 SetIsAutomatedNil sets the value for IsAutomated to be an explicit nil

### UnsetIsAutomated
`func (o *TestPlanTestPointsSearchApiModel) UnsetIsAutomated()`

UnsetIsAutomated ensures that no value is present for IsAutomated, not even an explicit nil
### GetName

`func (o *TestPlanTestPointsSearchApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestPlanTestPointsSearchApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestPlanTestPointsSearchApiModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestPlanTestPointsSearchApiModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestPlanTestPointsSearchApiModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestPlanTestPointsSearchApiModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetConfigurationIds

`func (o *TestPlanTestPointsSearchApiModel) GetConfigurationIds() []string`

GetConfigurationIds returns the ConfigurationIds field if non-nil, zero value otherwise.

### GetConfigurationIdsOk

`func (o *TestPlanTestPointsSearchApiModel) GetConfigurationIdsOk() (*[]string, bool)`

GetConfigurationIdsOk returns a tuple with the ConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIds

`func (o *TestPlanTestPointsSearchApiModel) SetConfigurationIds(v []string)`

SetConfigurationIds sets ConfigurationIds field to given value.

### HasConfigurationIds

`func (o *TestPlanTestPointsSearchApiModel) HasConfigurationIds() bool`

HasConfigurationIds returns a boolean if a field has been set.

### SetConfigurationIdsNil

`func (o *TestPlanTestPointsSearchApiModel) SetConfigurationIdsNil(b bool)`

 SetConfigurationIdsNil sets the value for ConfigurationIds to be an explicit nil

### UnsetConfigurationIds
`func (o *TestPlanTestPointsSearchApiModel) UnsetConfigurationIds()`

UnsetConfigurationIds ensures that no value is present for ConfigurationIds, not even an explicit nil
### GetTesterIds

`func (o *TestPlanTestPointsSearchApiModel) GetTesterIds() []*string`

GetTesterIds returns the TesterIds field if non-nil, zero value otherwise.

### GetTesterIdsOk

`func (o *TestPlanTestPointsSearchApiModel) GetTesterIdsOk() (*[]*string, bool)`

GetTesterIdsOk returns a tuple with the TesterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesterIds

`func (o *TestPlanTestPointsSearchApiModel) SetTesterIds(v []*string)`

SetTesterIds sets TesterIds field to given value.

### HasTesterIds

`func (o *TestPlanTestPointsSearchApiModel) HasTesterIds() bool`

HasTesterIds returns a boolean if a field has been set.

### SetTesterIdsNil

`func (o *TestPlanTestPointsSearchApiModel) SetTesterIdsNil(b bool)`

 SetTesterIdsNil sets the value for TesterIds to be an explicit nil

### UnsetTesterIds
`func (o *TestPlanTestPointsSearchApiModel) UnsetTesterIds()`

UnsetTesterIds ensures that no value is present for TesterIds, not even an explicit nil
### GetDuration

`func (o *TestPlanTestPointsSearchApiModel) GetDuration() Int64RangeSelectorModel`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TestPlanTestPointsSearchApiModel) GetDurationOk() (*Int64RangeSelectorModel, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TestPlanTestPointsSearchApiModel) SetDuration(v Int64RangeSelectorModel)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TestPlanTestPointsSearchApiModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *TestPlanTestPointsSearchApiModel) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *TestPlanTestPointsSearchApiModel) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetSectionIds

`func (o *TestPlanTestPointsSearchApiModel) GetSectionIds() []string`

GetSectionIds returns the SectionIds field if non-nil, zero value otherwise.

### GetSectionIdsOk

`func (o *TestPlanTestPointsSearchApiModel) GetSectionIdsOk() (*[]string, bool)`

GetSectionIdsOk returns a tuple with the SectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionIds

`func (o *TestPlanTestPointsSearchApiModel) SetSectionIds(v []string)`

SetSectionIds sets SectionIds field to given value.

### HasSectionIds

`func (o *TestPlanTestPointsSearchApiModel) HasSectionIds() bool`

HasSectionIds returns a boolean if a field has been set.

### SetSectionIdsNil

`func (o *TestPlanTestPointsSearchApiModel) SetSectionIdsNil(b bool)`

 SetSectionIdsNil sets the value for SectionIds to be an explicit nil

### UnsetSectionIds
`func (o *TestPlanTestPointsSearchApiModel) UnsetSectionIds()`

UnsetSectionIds ensures that no value is present for SectionIds, not even an explicit nil
### GetCreatedDate

`func (o *TestPlanTestPointsSearchApiModel) GetCreatedDate() DateTimeRangeSelectorModel`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestPlanTestPointsSearchApiModel) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestPlanTestPointsSearchApiModel) SetCreatedDate(v DateTimeRangeSelectorModel)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TestPlanTestPointsSearchApiModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *TestPlanTestPointsSearchApiModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *TestPlanTestPointsSearchApiModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetCreatedByIds

`func (o *TestPlanTestPointsSearchApiModel) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *TestPlanTestPointsSearchApiModel) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *TestPlanTestPointsSearchApiModel) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *TestPlanTestPointsSearchApiModel) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *TestPlanTestPointsSearchApiModel) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *TestPlanTestPointsSearchApiModel) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedDate

`func (o *TestPlanTestPointsSearchApiModel) GetModifiedDate() DateTimeRangeSelectorModel`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestPlanTestPointsSearchApiModel) GetModifiedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestPlanTestPointsSearchApiModel) SetModifiedDate(v DateTimeRangeSelectorModel)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestPlanTestPointsSearchApiModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *TestPlanTestPointsSearchApiModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *TestPlanTestPointsSearchApiModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedByIds

`func (o *TestPlanTestPointsSearchApiModel) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *TestPlanTestPointsSearchApiModel) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *TestPlanTestPointsSearchApiModel) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *TestPlanTestPointsSearchApiModel) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *TestPlanTestPointsSearchApiModel) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *TestPlanTestPointsSearchApiModel) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil
### GetTags

`func (o *TestPlanTestPointsSearchApiModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestPlanTestPointsSearchApiModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestPlanTestPointsSearchApiModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TestPlanTestPointsSearchApiModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *TestPlanTestPointsSearchApiModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *TestPlanTestPointsSearchApiModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetAttributes

`func (o *TestPlanTestPointsSearchApiModel) GetAttributes() map[string][]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TestPlanTestPointsSearchApiModel) GetAttributesOk() (*map[string][]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TestPlanTestPointsSearchApiModel) SetAttributes(v map[string][]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TestPlanTestPointsSearchApiModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *TestPlanTestPointsSearchApiModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *TestPlanTestPointsSearchApiModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetWorkItemCreatedDate

`func (o *TestPlanTestPointsSearchApiModel) GetWorkItemCreatedDate() DateTimeRangeSelectorModel`

GetWorkItemCreatedDate returns the WorkItemCreatedDate field if non-nil, zero value otherwise.

### GetWorkItemCreatedDateOk

`func (o *TestPlanTestPointsSearchApiModel) GetWorkItemCreatedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetWorkItemCreatedDateOk returns a tuple with the WorkItemCreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemCreatedDate

`func (o *TestPlanTestPointsSearchApiModel) SetWorkItemCreatedDate(v DateTimeRangeSelectorModel)`

SetWorkItemCreatedDate sets WorkItemCreatedDate field to given value.

### HasWorkItemCreatedDate

`func (o *TestPlanTestPointsSearchApiModel) HasWorkItemCreatedDate() bool`

HasWorkItemCreatedDate returns a boolean if a field has been set.

### SetWorkItemCreatedDateNil

`func (o *TestPlanTestPointsSearchApiModel) SetWorkItemCreatedDateNil(b bool)`

 SetWorkItemCreatedDateNil sets the value for WorkItemCreatedDate to be an explicit nil

### UnsetWorkItemCreatedDate
`func (o *TestPlanTestPointsSearchApiModel) UnsetWorkItemCreatedDate()`

UnsetWorkItemCreatedDate ensures that no value is present for WorkItemCreatedDate, not even an explicit nil
### GetWorkItemCreatedByIds

`func (o *TestPlanTestPointsSearchApiModel) GetWorkItemCreatedByIds() []string`

GetWorkItemCreatedByIds returns the WorkItemCreatedByIds field if non-nil, zero value otherwise.

### GetWorkItemCreatedByIdsOk

`func (o *TestPlanTestPointsSearchApiModel) GetWorkItemCreatedByIdsOk() (*[]string, bool)`

GetWorkItemCreatedByIdsOk returns a tuple with the WorkItemCreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemCreatedByIds

`func (o *TestPlanTestPointsSearchApiModel) SetWorkItemCreatedByIds(v []string)`

SetWorkItemCreatedByIds sets WorkItemCreatedByIds field to given value.

### HasWorkItemCreatedByIds

`func (o *TestPlanTestPointsSearchApiModel) HasWorkItemCreatedByIds() bool`

HasWorkItemCreatedByIds returns a boolean if a field has been set.

### SetWorkItemCreatedByIdsNil

`func (o *TestPlanTestPointsSearchApiModel) SetWorkItemCreatedByIdsNil(b bool)`

 SetWorkItemCreatedByIdsNil sets the value for WorkItemCreatedByIds to be an explicit nil

### UnsetWorkItemCreatedByIds
`func (o *TestPlanTestPointsSearchApiModel) UnsetWorkItemCreatedByIds()`

UnsetWorkItemCreatedByIds ensures that no value is present for WorkItemCreatedByIds, not even an explicit nil
### GetWorkItemModifiedDate

`func (o *TestPlanTestPointsSearchApiModel) GetWorkItemModifiedDate() DateTimeRangeSelectorModel`

GetWorkItemModifiedDate returns the WorkItemModifiedDate field if non-nil, zero value otherwise.

### GetWorkItemModifiedDateOk

`func (o *TestPlanTestPointsSearchApiModel) GetWorkItemModifiedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetWorkItemModifiedDateOk returns a tuple with the WorkItemModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemModifiedDate

`func (o *TestPlanTestPointsSearchApiModel) SetWorkItemModifiedDate(v DateTimeRangeSelectorModel)`

SetWorkItemModifiedDate sets WorkItemModifiedDate field to given value.

### HasWorkItemModifiedDate

`func (o *TestPlanTestPointsSearchApiModel) HasWorkItemModifiedDate() bool`

HasWorkItemModifiedDate returns a boolean if a field has been set.

### SetWorkItemModifiedDateNil

`func (o *TestPlanTestPointsSearchApiModel) SetWorkItemModifiedDateNil(b bool)`

 SetWorkItemModifiedDateNil sets the value for WorkItemModifiedDate to be an explicit nil

### UnsetWorkItemModifiedDate
`func (o *TestPlanTestPointsSearchApiModel) UnsetWorkItemModifiedDate()`

UnsetWorkItemModifiedDate ensures that no value is present for WorkItemModifiedDate, not even an explicit nil
### GetWorkItemModifiedByIds

`func (o *TestPlanTestPointsSearchApiModel) GetWorkItemModifiedByIds() []string`

GetWorkItemModifiedByIds returns the WorkItemModifiedByIds field if non-nil, zero value otherwise.

### GetWorkItemModifiedByIdsOk

`func (o *TestPlanTestPointsSearchApiModel) GetWorkItemModifiedByIdsOk() (*[]string, bool)`

GetWorkItemModifiedByIdsOk returns a tuple with the WorkItemModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemModifiedByIds

`func (o *TestPlanTestPointsSearchApiModel) SetWorkItemModifiedByIds(v []string)`

SetWorkItemModifiedByIds sets WorkItemModifiedByIds field to given value.

### HasWorkItemModifiedByIds

`func (o *TestPlanTestPointsSearchApiModel) HasWorkItemModifiedByIds() bool`

HasWorkItemModifiedByIds returns a boolean if a field has been set.

### SetWorkItemModifiedByIdsNil

`func (o *TestPlanTestPointsSearchApiModel) SetWorkItemModifiedByIdsNil(b bool)`

 SetWorkItemModifiedByIdsNil sets the value for WorkItemModifiedByIds to be an explicit nil

### UnsetWorkItemModifiedByIds
`func (o *TestPlanTestPointsSearchApiModel) UnsetWorkItemModifiedByIds()`

UnsetWorkItemModifiedByIds ensures that no value is present for WorkItemModifiedByIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


