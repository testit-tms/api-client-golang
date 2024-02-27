# TestPointFilterModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestPlanIds** | Pointer to **[]string** | Specifies a test point test plan IDS to search for | [optional] 
**TestSuiteIds** | Pointer to **[]string** | Specifies a test point test suite IDs to search for | [optional] 
**WorkItemGlobalIds** | Pointer to **[]int64** | Specifies a test point work item global IDs to search for | [optional] 
**WorkItemMedianDuration** | Pointer to [**NullableTestPointFilterModelWorkItemMedianDuration**](TestPointFilterModelWorkItemMedianDuration.md) |  | [optional] 
**WorkItemIsDeleted** | Pointer to **NullableBool** | Specifies a test point work item is deleted flag to search for | [optional] 
**Statuses** | Pointer to [**[]TestPointStatus**](TestPointStatus.md) | Specifies a test point statuses to search for | [optional] 
**Priorities** | Pointer to [**[]WorkItemPriorityModel**](WorkItemPriorityModel.md) | Specifies a test point priorities to search for | [optional] 
**IsAutomated** | Pointer to **NullableBool** | Specifies a test point automation status to search for | [optional] 
**Name** | Pointer to **NullableString** | Specifies a test point name to search for | [optional] 
**ConfigurationIds** | Pointer to **[]string** | Specifies a test point configuration IDs to search for | [optional] 
**TesterIds** | Pointer to **[]string** | Specifies a test point assigned user IDs to search for | [optional] 
**Duration** | Pointer to [**NullableTestPointFilterModelDuration**](TestPointFilterModelDuration.md) |  | [optional] 
**SectionIds** | Pointer to **[]string** | Specifies a test point work item section IDs to search for | [optional] 
**CreatedDate** | Pointer to [**NullableTestPointFilterModelCreatedDate**](TestPointFilterModelCreatedDate.md) |  | [optional] 
**CreatedByIds** | Pointer to **[]string** | Specifies a test point creator IDs to search for | [optional] 
**ModifiedDate** | Pointer to [**NullableTestPointFilterModelModifiedDate**](TestPointFilterModelModifiedDate.md) |  | [optional] 
**ModifiedByIds** | Pointer to **[]string** | Specifies a test point last editor IDs to search for | [optional] 
**Tags** | Pointer to **[]string** | Specifies a test point tags to search for | [optional] 
**Attributes** | Pointer to **map[string][]string** | Specifies a test point attributes to search for | [optional] 
**WorkItemCreatedDate** | Pointer to [**NullableTestPointFilterModelWorkItemCreatedDate**](TestPointFilterModelWorkItemCreatedDate.md) |  | [optional] 
**WorkItemCreatedByIds** | Pointer to **[]string** | Specifies a work item creator IDs to search for | [optional] 
**WorkItemModifiedDate** | Pointer to [**NullableTestPointFilterModelWorkItemModifiedDate**](TestPointFilterModelWorkItemModifiedDate.md) |  | [optional] 
**WorkItemModifiedByIds** | Pointer to **[]string** | Specifies a work item last editor IDs to search for | [optional] 

## Methods

### NewTestPointFilterModel

`func NewTestPointFilterModel() *TestPointFilterModel`

NewTestPointFilterModel instantiates a new TestPointFilterModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPointFilterModelWithDefaults

`func NewTestPointFilterModelWithDefaults() *TestPointFilterModel`

NewTestPointFilterModelWithDefaults instantiates a new TestPointFilterModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestPlanIds

`func (o *TestPointFilterModel) GetTestPlanIds() []string`

GetTestPlanIds returns the TestPlanIds field if non-nil, zero value otherwise.

### GetTestPlanIdsOk

`func (o *TestPointFilterModel) GetTestPlanIdsOk() (*[]string, bool)`

GetTestPlanIdsOk returns a tuple with the TestPlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanIds

`func (o *TestPointFilterModel) SetTestPlanIds(v []string)`

SetTestPlanIds sets TestPlanIds field to given value.

### HasTestPlanIds

`func (o *TestPointFilterModel) HasTestPlanIds() bool`

HasTestPlanIds returns a boolean if a field has been set.

### SetTestPlanIdsNil

`func (o *TestPointFilterModel) SetTestPlanIdsNil(b bool)`

 SetTestPlanIdsNil sets the value for TestPlanIds to be an explicit nil

### UnsetTestPlanIds
`func (o *TestPointFilterModel) UnsetTestPlanIds()`

UnsetTestPlanIds ensures that no value is present for TestPlanIds, not even an explicit nil
### GetTestSuiteIds

`func (o *TestPointFilterModel) GetTestSuiteIds() []string`

GetTestSuiteIds returns the TestSuiteIds field if non-nil, zero value otherwise.

### GetTestSuiteIdsOk

`func (o *TestPointFilterModel) GetTestSuiteIdsOk() (*[]string, bool)`

GetTestSuiteIdsOk returns a tuple with the TestSuiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuiteIds

`func (o *TestPointFilterModel) SetTestSuiteIds(v []string)`

SetTestSuiteIds sets TestSuiteIds field to given value.

### HasTestSuiteIds

`func (o *TestPointFilterModel) HasTestSuiteIds() bool`

HasTestSuiteIds returns a boolean if a field has been set.

### SetTestSuiteIdsNil

`func (o *TestPointFilterModel) SetTestSuiteIdsNil(b bool)`

 SetTestSuiteIdsNil sets the value for TestSuiteIds to be an explicit nil

### UnsetTestSuiteIds
`func (o *TestPointFilterModel) UnsetTestSuiteIds()`

UnsetTestSuiteIds ensures that no value is present for TestSuiteIds, not even an explicit nil
### GetWorkItemGlobalIds

`func (o *TestPointFilterModel) GetWorkItemGlobalIds() []int64`

GetWorkItemGlobalIds returns the WorkItemGlobalIds field if non-nil, zero value otherwise.

### GetWorkItemGlobalIdsOk

`func (o *TestPointFilterModel) GetWorkItemGlobalIdsOk() (*[]int64, bool)`

GetWorkItemGlobalIdsOk returns a tuple with the WorkItemGlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemGlobalIds

`func (o *TestPointFilterModel) SetWorkItemGlobalIds(v []int64)`

SetWorkItemGlobalIds sets WorkItemGlobalIds field to given value.

### HasWorkItemGlobalIds

`func (o *TestPointFilterModel) HasWorkItemGlobalIds() bool`

HasWorkItemGlobalIds returns a boolean if a field has been set.

### SetWorkItemGlobalIdsNil

`func (o *TestPointFilterModel) SetWorkItemGlobalIdsNil(b bool)`

 SetWorkItemGlobalIdsNil sets the value for WorkItemGlobalIds to be an explicit nil

### UnsetWorkItemGlobalIds
`func (o *TestPointFilterModel) UnsetWorkItemGlobalIds()`

UnsetWorkItemGlobalIds ensures that no value is present for WorkItemGlobalIds, not even an explicit nil
### GetWorkItemMedianDuration

`func (o *TestPointFilterModel) GetWorkItemMedianDuration() TestPointFilterModelWorkItemMedianDuration`

GetWorkItemMedianDuration returns the WorkItemMedianDuration field if non-nil, zero value otherwise.

### GetWorkItemMedianDurationOk

`func (o *TestPointFilterModel) GetWorkItemMedianDurationOk() (*TestPointFilterModelWorkItemMedianDuration, bool)`

GetWorkItemMedianDurationOk returns a tuple with the WorkItemMedianDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemMedianDuration

`func (o *TestPointFilterModel) SetWorkItemMedianDuration(v TestPointFilterModelWorkItemMedianDuration)`

SetWorkItemMedianDuration sets WorkItemMedianDuration field to given value.

### HasWorkItemMedianDuration

`func (o *TestPointFilterModel) HasWorkItemMedianDuration() bool`

HasWorkItemMedianDuration returns a boolean if a field has been set.

### SetWorkItemMedianDurationNil

`func (o *TestPointFilterModel) SetWorkItemMedianDurationNil(b bool)`

 SetWorkItemMedianDurationNil sets the value for WorkItemMedianDuration to be an explicit nil

### UnsetWorkItemMedianDuration
`func (o *TestPointFilterModel) UnsetWorkItemMedianDuration()`

UnsetWorkItemMedianDuration ensures that no value is present for WorkItemMedianDuration, not even an explicit nil
### GetWorkItemIsDeleted

`func (o *TestPointFilterModel) GetWorkItemIsDeleted() bool`

GetWorkItemIsDeleted returns the WorkItemIsDeleted field if non-nil, zero value otherwise.

### GetWorkItemIsDeletedOk

`func (o *TestPointFilterModel) GetWorkItemIsDeletedOk() (*bool, bool)`

GetWorkItemIsDeletedOk returns a tuple with the WorkItemIsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemIsDeleted

`func (o *TestPointFilterModel) SetWorkItemIsDeleted(v bool)`

SetWorkItemIsDeleted sets WorkItemIsDeleted field to given value.

### HasWorkItemIsDeleted

`func (o *TestPointFilterModel) HasWorkItemIsDeleted() bool`

HasWorkItemIsDeleted returns a boolean if a field has been set.

### SetWorkItemIsDeletedNil

`func (o *TestPointFilterModel) SetWorkItemIsDeletedNil(b bool)`

 SetWorkItemIsDeletedNil sets the value for WorkItemIsDeleted to be an explicit nil

### UnsetWorkItemIsDeleted
`func (o *TestPointFilterModel) UnsetWorkItemIsDeleted()`

UnsetWorkItemIsDeleted ensures that no value is present for WorkItemIsDeleted, not even an explicit nil
### GetStatuses

`func (o *TestPointFilterModel) GetStatuses() []TestPointStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *TestPointFilterModel) GetStatusesOk() (*[]TestPointStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *TestPointFilterModel) SetStatuses(v []TestPointStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *TestPointFilterModel) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### SetStatusesNil

`func (o *TestPointFilterModel) SetStatusesNil(b bool)`

 SetStatusesNil sets the value for Statuses to be an explicit nil

### UnsetStatuses
`func (o *TestPointFilterModel) UnsetStatuses()`

UnsetStatuses ensures that no value is present for Statuses, not even an explicit nil
### GetPriorities

`func (o *TestPointFilterModel) GetPriorities() []WorkItemPriorityModel`

GetPriorities returns the Priorities field if non-nil, zero value otherwise.

### GetPrioritiesOk

`func (o *TestPointFilterModel) GetPrioritiesOk() (*[]WorkItemPriorityModel, bool)`

GetPrioritiesOk returns a tuple with the Priorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorities

`func (o *TestPointFilterModel) SetPriorities(v []WorkItemPriorityModel)`

SetPriorities sets Priorities field to given value.

### HasPriorities

`func (o *TestPointFilterModel) HasPriorities() bool`

HasPriorities returns a boolean if a field has been set.

### SetPrioritiesNil

`func (o *TestPointFilterModel) SetPrioritiesNil(b bool)`

 SetPrioritiesNil sets the value for Priorities to be an explicit nil

### UnsetPriorities
`func (o *TestPointFilterModel) UnsetPriorities()`

UnsetPriorities ensures that no value is present for Priorities, not even an explicit nil
### GetIsAutomated

`func (o *TestPointFilterModel) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *TestPointFilterModel) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *TestPointFilterModel) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *TestPointFilterModel) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### SetIsAutomatedNil

`func (o *TestPointFilterModel) SetIsAutomatedNil(b bool)`

 SetIsAutomatedNil sets the value for IsAutomated to be an explicit nil

### UnsetIsAutomated
`func (o *TestPointFilterModel) UnsetIsAutomated()`

UnsetIsAutomated ensures that no value is present for IsAutomated, not even an explicit nil
### GetName

`func (o *TestPointFilterModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestPointFilterModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestPointFilterModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestPointFilterModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestPointFilterModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestPointFilterModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetConfigurationIds

`func (o *TestPointFilterModel) GetConfigurationIds() []string`

GetConfigurationIds returns the ConfigurationIds field if non-nil, zero value otherwise.

### GetConfigurationIdsOk

`func (o *TestPointFilterModel) GetConfigurationIdsOk() (*[]string, bool)`

GetConfigurationIdsOk returns a tuple with the ConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIds

`func (o *TestPointFilterModel) SetConfigurationIds(v []string)`

SetConfigurationIds sets ConfigurationIds field to given value.

### HasConfigurationIds

`func (o *TestPointFilterModel) HasConfigurationIds() bool`

HasConfigurationIds returns a boolean if a field has been set.

### SetConfigurationIdsNil

`func (o *TestPointFilterModel) SetConfigurationIdsNil(b bool)`

 SetConfigurationIdsNil sets the value for ConfigurationIds to be an explicit nil

### UnsetConfigurationIds
`func (o *TestPointFilterModel) UnsetConfigurationIds()`

UnsetConfigurationIds ensures that no value is present for ConfigurationIds, not even an explicit nil
### GetTesterIds

`func (o *TestPointFilterModel) GetTesterIds() []string`

GetTesterIds returns the TesterIds field if non-nil, zero value otherwise.

### GetTesterIdsOk

`func (o *TestPointFilterModel) GetTesterIdsOk() (*[]string, bool)`

GetTesterIdsOk returns a tuple with the TesterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesterIds

`func (o *TestPointFilterModel) SetTesterIds(v []string)`

SetTesterIds sets TesterIds field to given value.

### HasTesterIds

`func (o *TestPointFilterModel) HasTesterIds() bool`

HasTesterIds returns a boolean if a field has been set.

### SetTesterIdsNil

`func (o *TestPointFilterModel) SetTesterIdsNil(b bool)`

 SetTesterIdsNil sets the value for TesterIds to be an explicit nil

### UnsetTesterIds
`func (o *TestPointFilterModel) UnsetTesterIds()`

UnsetTesterIds ensures that no value is present for TesterIds, not even an explicit nil
### GetDuration

`func (o *TestPointFilterModel) GetDuration() TestPointFilterModelDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TestPointFilterModel) GetDurationOk() (*TestPointFilterModelDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TestPointFilterModel) SetDuration(v TestPointFilterModelDuration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TestPointFilterModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *TestPointFilterModel) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *TestPointFilterModel) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetSectionIds

`func (o *TestPointFilterModel) GetSectionIds() []string`

GetSectionIds returns the SectionIds field if non-nil, zero value otherwise.

### GetSectionIdsOk

`func (o *TestPointFilterModel) GetSectionIdsOk() (*[]string, bool)`

GetSectionIdsOk returns a tuple with the SectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionIds

`func (o *TestPointFilterModel) SetSectionIds(v []string)`

SetSectionIds sets SectionIds field to given value.

### HasSectionIds

`func (o *TestPointFilterModel) HasSectionIds() bool`

HasSectionIds returns a boolean if a field has been set.

### SetSectionIdsNil

`func (o *TestPointFilterModel) SetSectionIdsNil(b bool)`

 SetSectionIdsNil sets the value for SectionIds to be an explicit nil

### UnsetSectionIds
`func (o *TestPointFilterModel) UnsetSectionIds()`

UnsetSectionIds ensures that no value is present for SectionIds, not even an explicit nil
### GetCreatedDate

`func (o *TestPointFilterModel) GetCreatedDate() TestPointFilterModelCreatedDate`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestPointFilterModel) GetCreatedDateOk() (*TestPointFilterModelCreatedDate, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestPointFilterModel) SetCreatedDate(v TestPointFilterModelCreatedDate)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TestPointFilterModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *TestPointFilterModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *TestPointFilterModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetCreatedByIds

`func (o *TestPointFilterModel) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *TestPointFilterModel) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *TestPointFilterModel) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *TestPointFilterModel) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *TestPointFilterModel) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *TestPointFilterModel) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedDate

`func (o *TestPointFilterModel) GetModifiedDate() TestPointFilterModelModifiedDate`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestPointFilterModel) GetModifiedDateOk() (*TestPointFilterModelModifiedDate, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestPointFilterModel) SetModifiedDate(v TestPointFilterModelModifiedDate)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestPointFilterModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *TestPointFilterModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *TestPointFilterModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedByIds

`func (o *TestPointFilterModel) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *TestPointFilterModel) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *TestPointFilterModel) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *TestPointFilterModel) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *TestPointFilterModel) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *TestPointFilterModel) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil
### GetTags

`func (o *TestPointFilterModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestPointFilterModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestPointFilterModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TestPointFilterModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *TestPointFilterModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *TestPointFilterModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetAttributes

`func (o *TestPointFilterModel) GetAttributes() map[string][]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TestPointFilterModel) GetAttributesOk() (*map[string][]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TestPointFilterModel) SetAttributes(v map[string][]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TestPointFilterModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *TestPointFilterModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *TestPointFilterModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetWorkItemCreatedDate

`func (o *TestPointFilterModel) GetWorkItemCreatedDate() TestPointFilterModelWorkItemCreatedDate`

GetWorkItemCreatedDate returns the WorkItemCreatedDate field if non-nil, zero value otherwise.

### GetWorkItemCreatedDateOk

`func (o *TestPointFilterModel) GetWorkItemCreatedDateOk() (*TestPointFilterModelWorkItemCreatedDate, bool)`

GetWorkItemCreatedDateOk returns a tuple with the WorkItemCreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemCreatedDate

`func (o *TestPointFilterModel) SetWorkItemCreatedDate(v TestPointFilterModelWorkItemCreatedDate)`

SetWorkItemCreatedDate sets WorkItemCreatedDate field to given value.

### HasWorkItemCreatedDate

`func (o *TestPointFilterModel) HasWorkItemCreatedDate() bool`

HasWorkItemCreatedDate returns a boolean if a field has been set.

### SetWorkItemCreatedDateNil

`func (o *TestPointFilterModel) SetWorkItemCreatedDateNil(b bool)`

 SetWorkItemCreatedDateNil sets the value for WorkItemCreatedDate to be an explicit nil

### UnsetWorkItemCreatedDate
`func (o *TestPointFilterModel) UnsetWorkItemCreatedDate()`

UnsetWorkItemCreatedDate ensures that no value is present for WorkItemCreatedDate, not even an explicit nil
### GetWorkItemCreatedByIds

`func (o *TestPointFilterModel) GetWorkItemCreatedByIds() []string`

GetWorkItemCreatedByIds returns the WorkItemCreatedByIds field if non-nil, zero value otherwise.

### GetWorkItemCreatedByIdsOk

`func (o *TestPointFilterModel) GetWorkItemCreatedByIdsOk() (*[]string, bool)`

GetWorkItemCreatedByIdsOk returns a tuple with the WorkItemCreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemCreatedByIds

`func (o *TestPointFilterModel) SetWorkItemCreatedByIds(v []string)`

SetWorkItemCreatedByIds sets WorkItemCreatedByIds field to given value.

### HasWorkItemCreatedByIds

`func (o *TestPointFilterModel) HasWorkItemCreatedByIds() bool`

HasWorkItemCreatedByIds returns a boolean if a field has been set.

### SetWorkItemCreatedByIdsNil

`func (o *TestPointFilterModel) SetWorkItemCreatedByIdsNil(b bool)`

 SetWorkItemCreatedByIdsNil sets the value for WorkItemCreatedByIds to be an explicit nil

### UnsetWorkItemCreatedByIds
`func (o *TestPointFilterModel) UnsetWorkItemCreatedByIds()`

UnsetWorkItemCreatedByIds ensures that no value is present for WorkItemCreatedByIds, not even an explicit nil
### GetWorkItemModifiedDate

`func (o *TestPointFilterModel) GetWorkItemModifiedDate() TestPointFilterModelWorkItemModifiedDate`

GetWorkItemModifiedDate returns the WorkItemModifiedDate field if non-nil, zero value otherwise.

### GetWorkItemModifiedDateOk

`func (o *TestPointFilterModel) GetWorkItemModifiedDateOk() (*TestPointFilterModelWorkItemModifiedDate, bool)`

GetWorkItemModifiedDateOk returns a tuple with the WorkItemModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemModifiedDate

`func (o *TestPointFilterModel) SetWorkItemModifiedDate(v TestPointFilterModelWorkItemModifiedDate)`

SetWorkItemModifiedDate sets WorkItemModifiedDate field to given value.

### HasWorkItemModifiedDate

`func (o *TestPointFilterModel) HasWorkItemModifiedDate() bool`

HasWorkItemModifiedDate returns a boolean if a field has been set.

### SetWorkItemModifiedDateNil

`func (o *TestPointFilterModel) SetWorkItemModifiedDateNil(b bool)`

 SetWorkItemModifiedDateNil sets the value for WorkItemModifiedDate to be an explicit nil

### UnsetWorkItemModifiedDate
`func (o *TestPointFilterModel) UnsetWorkItemModifiedDate()`

UnsetWorkItemModifiedDate ensures that no value is present for WorkItemModifiedDate, not even an explicit nil
### GetWorkItemModifiedByIds

`func (o *TestPointFilterModel) GetWorkItemModifiedByIds() []string`

GetWorkItemModifiedByIds returns the WorkItemModifiedByIds field if non-nil, zero value otherwise.

### GetWorkItemModifiedByIdsOk

`func (o *TestPointFilterModel) GetWorkItemModifiedByIdsOk() (*[]string, bool)`

GetWorkItemModifiedByIdsOk returns a tuple with the WorkItemModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemModifiedByIds

`func (o *TestPointFilterModel) SetWorkItemModifiedByIds(v []string)`

SetWorkItemModifiedByIds sets WorkItemModifiedByIds field to given value.

### HasWorkItemModifiedByIds

`func (o *TestPointFilterModel) HasWorkItemModifiedByIds() bool`

HasWorkItemModifiedByIds returns a boolean if a field has been set.

### SetWorkItemModifiedByIdsNil

`func (o *TestPointFilterModel) SetWorkItemModifiedByIdsNil(b bool)`

 SetWorkItemModifiedByIdsNil sets the value for WorkItemModifiedByIds to be an explicit nil

### UnsetWorkItemModifiedByIds
`func (o *TestPointFilterModel) UnsetWorkItemModifiedByIds()`

UnsetWorkItemModifiedByIds ensures that no value is present for WorkItemModifiedByIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


