# TestPointFilterRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestPlanIds** | Pointer to **[]string** | Specifies a test point test plan IDS to search for | [optional] 
**TestSuiteIds** | Pointer to **[]string** | Specifies a test point test suite IDs to search for | [optional] 
**WorkItemGlobalIds** | Pointer to **[]int64** | Specifies a test point work item global IDs to search for | [optional] 
**WorkItemMedianDuration** | Pointer to [**NullableInt64RangeSelectorModel**](Int64RangeSelectorModel.md) | Specifies a test point work item median duration range to search for | [optional] 
**WorkItemIsDeleted** | Pointer to **NullableBool** | Specifies a test point work item is deleted flag to search for | [optional] 
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

### NewTestPointFilterRequestModel

`func NewTestPointFilterRequestModel() *TestPointFilterRequestModel`

NewTestPointFilterRequestModel instantiates a new TestPointFilterRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPointFilterRequestModelWithDefaults

`func NewTestPointFilterRequestModelWithDefaults() *TestPointFilterRequestModel`

NewTestPointFilterRequestModelWithDefaults instantiates a new TestPointFilterRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestPlanIds

`func (o *TestPointFilterRequestModel) GetTestPlanIds() []string`

GetTestPlanIds returns the TestPlanIds field if non-nil, zero value otherwise.

### GetTestPlanIdsOk

`func (o *TestPointFilterRequestModel) GetTestPlanIdsOk() (*[]string, bool)`

GetTestPlanIdsOk returns a tuple with the TestPlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanIds

`func (o *TestPointFilterRequestModel) SetTestPlanIds(v []string)`

SetTestPlanIds sets TestPlanIds field to given value.

### HasTestPlanIds

`func (o *TestPointFilterRequestModel) HasTestPlanIds() bool`

HasTestPlanIds returns a boolean if a field has been set.

### SetTestPlanIdsNil

`func (o *TestPointFilterRequestModel) SetTestPlanIdsNil(b bool)`

 SetTestPlanIdsNil sets the value for TestPlanIds to be an explicit nil

### UnsetTestPlanIds
`func (o *TestPointFilterRequestModel) UnsetTestPlanIds()`

UnsetTestPlanIds ensures that no value is present for TestPlanIds, not even an explicit nil
### GetTestSuiteIds

`func (o *TestPointFilterRequestModel) GetTestSuiteIds() []string`

GetTestSuiteIds returns the TestSuiteIds field if non-nil, zero value otherwise.

### GetTestSuiteIdsOk

`func (o *TestPointFilterRequestModel) GetTestSuiteIdsOk() (*[]string, bool)`

GetTestSuiteIdsOk returns a tuple with the TestSuiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuiteIds

`func (o *TestPointFilterRequestModel) SetTestSuiteIds(v []string)`

SetTestSuiteIds sets TestSuiteIds field to given value.

### HasTestSuiteIds

`func (o *TestPointFilterRequestModel) HasTestSuiteIds() bool`

HasTestSuiteIds returns a boolean if a field has been set.

### SetTestSuiteIdsNil

`func (o *TestPointFilterRequestModel) SetTestSuiteIdsNil(b bool)`

 SetTestSuiteIdsNil sets the value for TestSuiteIds to be an explicit nil

### UnsetTestSuiteIds
`func (o *TestPointFilterRequestModel) UnsetTestSuiteIds()`

UnsetTestSuiteIds ensures that no value is present for TestSuiteIds, not even an explicit nil
### GetWorkItemGlobalIds

`func (o *TestPointFilterRequestModel) GetWorkItemGlobalIds() []int64`

GetWorkItemGlobalIds returns the WorkItemGlobalIds field if non-nil, zero value otherwise.

### GetWorkItemGlobalIdsOk

`func (o *TestPointFilterRequestModel) GetWorkItemGlobalIdsOk() (*[]int64, bool)`

GetWorkItemGlobalIdsOk returns a tuple with the WorkItemGlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemGlobalIds

`func (o *TestPointFilterRequestModel) SetWorkItemGlobalIds(v []int64)`

SetWorkItemGlobalIds sets WorkItemGlobalIds field to given value.

### HasWorkItemGlobalIds

`func (o *TestPointFilterRequestModel) HasWorkItemGlobalIds() bool`

HasWorkItemGlobalIds returns a boolean if a field has been set.

### SetWorkItemGlobalIdsNil

`func (o *TestPointFilterRequestModel) SetWorkItemGlobalIdsNil(b bool)`

 SetWorkItemGlobalIdsNil sets the value for WorkItemGlobalIds to be an explicit nil

### UnsetWorkItemGlobalIds
`func (o *TestPointFilterRequestModel) UnsetWorkItemGlobalIds()`

UnsetWorkItemGlobalIds ensures that no value is present for WorkItemGlobalIds, not even an explicit nil
### GetWorkItemMedianDuration

`func (o *TestPointFilterRequestModel) GetWorkItemMedianDuration() Int64RangeSelectorModel`

GetWorkItemMedianDuration returns the WorkItemMedianDuration field if non-nil, zero value otherwise.

### GetWorkItemMedianDurationOk

`func (o *TestPointFilterRequestModel) GetWorkItemMedianDurationOk() (*Int64RangeSelectorModel, bool)`

GetWorkItemMedianDurationOk returns a tuple with the WorkItemMedianDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemMedianDuration

`func (o *TestPointFilterRequestModel) SetWorkItemMedianDuration(v Int64RangeSelectorModel)`

SetWorkItemMedianDuration sets WorkItemMedianDuration field to given value.

### HasWorkItemMedianDuration

`func (o *TestPointFilterRequestModel) HasWorkItemMedianDuration() bool`

HasWorkItemMedianDuration returns a boolean if a field has been set.

### SetWorkItemMedianDurationNil

`func (o *TestPointFilterRequestModel) SetWorkItemMedianDurationNil(b bool)`

 SetWorkItemMedianDurationNil sets the value for WorkItemMedianDuration to be an explicit nil

### UnsetWorkItemMedianDuration
`func (o *TestPointFilterRequestModel) UnsetWorkItemMedianDuration()`

UnsetWorkItemMedianDuration ensures that no value is present for WorkItemMedianDuration, not even an explicit nil
### GetWorkItemIsDeleted

`func (o *TestPointFilterRequestModel) GetWorkItemIsDeleted() bool`

GetWorkItemIsDeleted returns the WorkItemIsDeleted field if non-nil, zero value otherwise.

### GetWorkItemIsDeletedOk

`func (o *TestPointFilterRequestModel) GetWorkItemIsDeletedOk() (*bool, bool)`

GetWorkItemIsDeletedOk returns a tuple with the WorkItemIsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemIsDeleted

`func (o *TestPointFilterRequestModel) SetWorkItemIsDeleted(v bool)`

SetWorkItemIsDeleted sets WorkItemIsDeleted field to given value.

### HasWorkItemIsDeleted

`func (o *TestPointFilterRequestModel) HasWorkItemIsDeleted() bool`

HasWorkItemIsDeleted returns a boolean if a field has been set.

### SetWorkItemIsDeletedNil

`func (o *TestPointFilterRequestModel) SetWorkItemIsDeletedNil(b bool)`

 SetWorkItemIsDeletedNil sets the value for WorkItemIsDeleted to be an explicit nil

### UnsetWorkItemIsDeleted
`func (o *TestPointFilterRequestModel) UnsetWorkItemIsDeleted()`

UnsetWorkItemIsDeleted ensures that no value is present for WorkItemIsDeleted, not even an explicit nil
### GetStatuses

`func (o *TestPointFilterRequestModel) GetStatuses() []TestPointStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *TestPointFilterRequestModel) GetStatusesOk() (*[]TestPointStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *TestPointFilterRequestModel) SetStatuses(v []TestPointStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *TestPointFilterRequestModel) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### SetStatusesNil

`func (o *TestPointFilterRequestModel) SetStatusesNil(b bool)`

 SetStatusesNil sets the value for Statuses to be an explicit nil

### UnsetStatuses
`func (o *TestPointFilterRequestModel) UnsetStatuses()`

UnsetStatuses ensures that no value is present for Statuses, not even an explicit nil
### GetStatusCodes

`func (o *TestPointFilterRequestModel) GetStatusCodes() []string`

GetStatusCodes returns the StatusCodes field if non-nil, zero value otherwise.

### GetStatusCodesOk

`func (o *TestPointFilterRequestModel) GetStatusCodesOk() (*[]string, bool)`

GetStatusCodesOk returns a tuple with the StatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodes

`func (o *TestPointFilterRequestModel) SetStatusCodes(v []string)`

SetStatusCodes sets StatusCodes field to given value.

### HasStatusCodes

`func (o *TestPointFilterRequestModel) HasStatusCodes() bool`

HasStatusCodes returns a boolean if a field has been set.

### SetStatusCodesNil

`func (o *TestPointFilterRequestModel) SetStatusCodesNil(b bool)`

 SetStatusCodesNil sets the value for StatusCodes to be an explicit nil

### UnsetStatusCodes
`func (o *TestPointFilterRequestModel) UnsetStatusCodes()`

UnsetStatusCodes ensures that no value is present for StatusCodes, not even an explicit nil
### GetPriorities

`func (o *TestPointFilterRequestModel) GetPriorities() []WorkItemPriorityModel`

GetPriorities returns the Priorities field if non-nil, zero value otherwise.

### GetPrioritiesOk

`func (o *TestPointFilterRequestModel) GetPrioritiesOk() (*[]WorkItemPriorityModel, bool)`

GetPrioritiesOk returns a tuple with the Priorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorities

`func (o *TestPointFilterRequestModel) SetPriorities(v []WorkItemPriorityModel)`

SetPriorities sets Priorities field to given value.

### HasPriorities

`func (o *TestPointFilterRequestModel) HasPriorities() bool`

HasPriorities returns a boolean if a field has been set.

### SetPrioritiesNil

`func (o *TestPointFilterRequestModel) SetPrioritiesNil(b bool)`

 SetPrioritiesNil sets the value for Priorities to be an explicit nil

### UnsetPriorities
`func (o *TestPointFilterRequestModel) UnsetPriorities()`

UnsetPriorities ensures that no value is present for Priorities, not even an explicit nil
### GetIsAutomated

`func (o *TestPointFilterRequestModel) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *TestPointFilterRequestModel) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *TestPointFilterRequestModel) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *TestPointFilterRequestModel) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### SetIsAutomatedNil

`func (o *TestPointFilterRequestModel) SetIsAutomatedNil(b bool)`

 SetIsAutomatedNil sets the value for IsAutomated to be an explicit nil

### UnsetIsAutomated
`func (o *TestPointFilterRequestModel) UnsetIsAutomated()`

UnsetIsAutomated ensures that no value is present for IsAutomated, not even an explicit nil
### GetName

`func (o *TestPointFilterRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestPointFilterRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestPointFilterRequestModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestPointFilterRequestModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestPointFilterRequestModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestPointFilterRequestModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetConfigurationIds

`func (o *TestPointFilterRequestModel) GetConfigurationIds() []string`

GetConfigurationIds returns the ConfigurationIds field if non-nil, zero value otherwise.

### GetConfigurationIdsOk

`func (o *TestPointFilterRequestModel) GetConfigurationIdsOk() (*[]string, bool)`

GetConfigurationIdsOk returns a tuple with the ConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIds

`func (o *TestPointFilterRequestModel) SetConfigurationIds(v []string)`

SetConfigurationIds sets ConfigurationIds field to given value.

### HasConfigurationIds

`func (o *TestPointFilterRequestModel) HasConfigurationIds() bool`

HasConfigurationIds returns a boolean if a field has been set.

### SetConfigurationIdsNil

`func (o *TestPointFilterRequestModel) SetConfigurationIdsNil(b bool)`

 SetConfigurationIdsNil sets the value for ConfigurationIds to be an explicit nil

### UnsetConfigurationIds
`func (o *TestPointFilterRequestModel) UnsetConfigurationIds()`

UnsetConfigurationIds ensures that no value is present for ConfigurationIds, not even an explicit nil
### GetTesterIds

`func (o *TestPointFilterRequestModel) GetTesterIds() []string`

GetTesterIds returns the TesterIds field if non-nil, zero value otherwise.

### GetTesterIdsOk

`func (o *TestPointFilterRequestModel) GetTesterIdsOk() (*[]string, bool)`

GetTesterIdsOk returns a tuple with the TesterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesterIds

`func (o *TestPointFilterRequestModel) SetTesterIds(v []string)`

SetTesterIds sets TesterIds field to given value.

### HasTesterIds

`func (o *TestPointFilterRequestModel) HasTesterIds() bool`

HasTesterIds returns a boolean if a field has been set.

### SetTesterIdsNil

`func (o *TestPointFilterRequestModel) SetTesterIdsNil(b bool)`

 SetTesterIdsNil sets the value for TesterIds to be an explicit nil

### UnsetTesterIds
`func (o *TestPointFilterRequestModel) UnsetTesterIds()`

UnsetTesterIds ensures that no value is present for TesterIds, not even an explicit nil
### GetDuration

`func (o *TestPointFilterRequestModel) GetDuration() Int64RangeSelectorModel`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TestPointFilterRequestModel) GetDurationOk() (*Int64RangeSelectorModel, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TestPointFilterRequestModel) SetDuration(v Int64RangeSelectorModel)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TestPointFilterRequestModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *TestPointFilterRequestModel) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *TestPointFilterRequestModel) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetSectionIds

`func (o *TestPointFilterRequestModel) GetSectionIds() []string`

GetSectionIds returns the SectionIds field if non-nil, zero value otherwise.

### GetSectionIdsOk

`func (o *TestPointFilterRequestModel) GetSectionIdsOk() (*[]string, bool)`

GetSectionIdsOk returns a tuple with the SectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionIds

`func (o *TestPointFilterRequestModel) SetSectionIds(v []string)`

SetSectionIds sets SectionIds field to given value.

### HasSectionIds

`func (o *TestPointFilterRequestModel) HasSectionIds() bool`

HasSectionIds returns a boolean if a field has been set.

### SetSectionIdsNil

`func (o *TestPointFilterRequestModel) SetSectionIdsNil(b bool)`

 SetSectionIdsNil sets the value for SectionIds to be an explicit nil

### UnsetSectionIds
`func (o *TestPointFilterRequestModel) UnsetSectionIds()`

UnsetSectionIds ensures that no value is present for SectionIds, not even an explicit nil
### GetCreatedDate

`func (o *TestPointFilterRequestModel) GetCreatedDate() DateTimeRangeSelectorModel`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestPointFilterRequestModel) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestPointFilterRequestModel) SetCreatedDate(v DateTimeRangeSelectorModel)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TestPointFilterRequestModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *TestPointFilterRequestModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *TestPointFilterRequestModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetCreatedByIds

`func (o *TestPointFilterRequestModel) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *TestPointFilterRequestModel) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *TestPointFilterRequestModel) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *TestPointFilterRequestModel) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *TestPointFilterRequestModel) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *TestPointFilterRequestModel) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedDate

`func (o *TestPointFilterRequestModel) GetModifiedDate() DateTimeRangeSelectorModel`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestPointFilterRequestModel) GetModifiedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestPointFilterRequestModel) SetModifiedDate(v DateTimeRangeSelectorModel)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestPointFilterRequestModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *TestPointFilterRequestModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *TestPointFilterRequestModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedByIds

`func (o *TestPointFilterRequestModel) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *TestPointFilterRequestModel) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *TestPointFilterRequestModel) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *TestPointFilterRequestModel) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *TestPointFilterRequestModel) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *TestPointFilterRequestModel) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil
### GetTags

`func (o *TestPointFilterRequestModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestPointFilterRequestModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestPointFilterRequestModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TestPointFilterRequestModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *TestPointFilterRequestModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *TestPointFilterRequestModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetAttributes

`func (o *TestPointFilterRequestModel) GetAttributes() map[string][]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TestPointFilterRequestModel) GetAttributesOk() (*map[string][]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TestPointFilterRequestModel) SetAttributes(v map[string][]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TestPointFilterRequestModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *TestPointFilterRequestModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *TestPointFilterRequestModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetWorkItemCreatedDate

`func (o *TestPointFilterRequestModel) GetWorkItemCreatedDate() DateTimeRangeSelectorModel`

GetWorkItemCreatedDate returns the WorkItemCreatedDate field if non-nil, zero value otherwise.

### GetWorkItemCreatedDateOk

`func (o *TestPointFilterRequestModel) GetWorkItemCreatedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetWorkItemCreatedDateOk returns a tuple with the WorkItemCreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemCreatedDate

`func (o *TestPointFilterRequestModel) SetWorkItemCreatedDate(v DateTimeRangeSelectorModel)`

SetWorkItemCreatedDate sets WorkItemCreatedDate field to given value.

### HasWorkItemCreatedDate

`func (o *TestPointFilterRequestModel) HasWorkItemCreatedDate() bool`

HasWorkItemCreatedDate returns a boolean if a field has been set.

### SetWorkItemCreatedDateNil

`func (o *TestPointFilterRequestModel) SetWorkItemCreatedDateNil(b bool)`

 SetWorkItemCreatedDateNil sets the value for WorkItemCreatedDate to be an explicit nil

### UnsetWorkItemCreatedDate
`func (o *TestPointFilterRequestModel) UnsetWorkItemCreatedDate()`

UnsetWorkItemCreatedDate ensures that no value is present for WorkItemCreatedDate, not even an explicit nil
### GetWorkItemCreatedByIds

`func (o *TestPointFilterRequestModel) GetWorkItemCreatedByIds() []string`

GetWorkItemCreatedByIds returns the WorkItemCreatedByIds field if non-nil, zero value otherwise.

### GetWorkItemCreatedByIdsOk

`func (o *TestPointFilterRequestModel) GetWorkItemCreatedByIdsOk() (*[]string, bool)`

GetWorkItemCreatedByIdsOk returns a tuple with the WorkItemCreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemCreatedByIds

`func (o *TestPointFilterRequestModel) SetWorkItemCreatedByIds(v []string)`

SetWorkItemCreatedByIds sets WorkItemCreatedByIds field to given value.

### HasWorkItemCreatedByIds

`func (o *TestPointFilterRequestModel) HasWorkItemCreatedByIds() bool`

HasWorkItemCreatedByIds returns a boolean if a field has been set.

### SetWorkItemCreatedByIdsNil

`func (o *TestPointFilterRequestModel) SetWorkItemCreatedByIdsNil(b bool)`

 SetWorkItemCreatedByIdsNil sets the value for WorkItemCreatedByIds to be an explicit nil

### UnsetWorkItemCreatedByIds
`func (o *TestPointFilterRequestModel) UnsetWorkItemCreatedByIds()`

UnsetWorkItemCreatedByIds ensures that no value is present for WorkItemCreatedByIds, not even an explicit nil
### GetWorkItemModifiedDate

`func (o *TestPointFilterRequestModel) GetWorkItemModifiedDate() DateTimeRangeSelectorModel`

GetWorkItemModifiedDate returns the WorkItemModifiedDate field if non-nil, zero value otherwise.

### GetWorkItemModifiedDateOk

`func (o *TestPointFilterRequestModel) GetWorkItemModifiedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetWorkItemModifiedDateOk returns a tuple with the WorkItemModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemModifiedDate

`func (o *TestPointFilterRequestModel) SetWorkItemModifiedDate(v DateTimeRangeSelectorModel)`

SetWorkItemModifiedDate sets WorkItemModifiedDate field to given value.

### HasWorkItemModifiedDate

`func (o *TestPointFilterRequestModel) HasWorkItemModifiedDate() bool`

HasWorkItemModifiedDate returns a boolean if a field has been set.

### SetWorkItemModifiedDateNil

`func (o *TestPointFilterRequestModel) SetWorkItemModifiedDateNil(b bool)`

 SetWorkItemModifiedDateNil sets the value for WorkItemModifiedDate to be an explicit nil

### UnsetWorkItemModifiedDate
`func (o *TestPointFilterRequestModel) UnsetWorkItemModifiedDate()`

UnsetWorkItemModifiedDate ensures that no value is present for WorkItemModifiedDate, not even an explicit nil
### GetWorkItemModifiedByIds

`func (o *TestPointFilterRequestModel) GetWorkItemModifiedByIds() []string`

GetWorkItemModifiedByIds returns the WorkItemModifiedByIds field if non-nil, zero value otherwise.

### GetWorkItemModifiedByIdsOk

`func (o *TestPointFilterRequestModel) GetWorkItemModifiedByIdsOk() (*[]string, bool)`

GetWorkItemModifiedByIdsOk returns a tuple with the WorkItemModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemModifiedByIds

`func (o *TestPointFilterRequestModel) SetWorkItemModifiedByIds(v []string)`

SetWorkItemModifiedByIds sets WorkItemModifiedByIds field to given value.

### HasWorkItemModifiedByIds

`func (o *TestPointFilterRequestModel) HasWorkItemModifiedByIds() bool`

HasWorkItemModifiedByIds returns a boolean if a field has been set.

### SetWorkItemModifiedByIdsNil

`func (o *TestPointFilterRequestModel) SetWorkItemModifiedByIdsNil(b bool)`

 SetWorkItemModifiedByIdsNil sets the value for WorkItemModifiedByIds to be an explicit nil

### UnsetWorkItemModifiedByIds
`func (o *TestPointFilterRequestModel) UnsetWorkItemModifiedByIds()`

UnsetWorkItemModifiedByIds ensures that no value is present for WorkItemModifiedByIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


