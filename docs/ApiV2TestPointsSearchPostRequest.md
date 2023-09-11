# ApiV2TestPointsSearchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestPlanIds** | Pointer to **[]string** | Specifies a test point test plan IDS to search for | [optional] 
**TestSuiteIds** | Pointer to **[]string** | Specifies a test point test suite IDs to search for | [optional] 
**WorkItemGlobalIds** | Pointer to **[]int64** | Specifies a test point work item global IDs to search for | [optional] 
**WorkItemMedianDuration** | Pointer to [**NullableTestPointFilterModelWorkItemMedianDuration**](TestPointFilterModelWorkItemMedianDuration.md) |  | [optional] 
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

### NewApiV2TestPointsSearchPostRequest

`func NewApiV2TestPointsSearchPostRequest() *ApiV2TestPointsSearchPostRequest`

NewApiV2TestPointsSearchPostRequest instantiates a new ApiV2TestPointsSearchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2TestPointsSearchPostRequestWithDefaults

`func NewApiV2TestPointsSearchPostRequestWithDefaults() *ApiV2TestPointsSearchPostRequest`

NewApiV2TestPointsSearchPostRequestWithDefaults instantiates a new ApiV2TestPointsSearchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestPlanIds

`func (o *ApiV2TestPointsSearchPostRequest) GetTestPlanIds() []string`

GetTestPlanIds returns the TestPlanIds field if non-nil, zero value otherwise.

### GetTestPlanIdsOk

`func (o *ApiV2TestPointsSearchPostRequest) GetTestPlanIdsOk() (*[]string, bool)`

GetTestPlanIdsOk returns a tuple with the TestPlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanIds

`func (o *ApiV2TestPointsSearchPostRequest) SetTestPlanIds(v []string)`

SetTestPlanIds sets TestPlanIds field to given value.

### HasTestPlanIds

`func (o *ApiV2TestPointsSearchPostRequest) HasTestPlanIds() bool`

HasTestPlanIds returns a boolean if a field has been set.

### SetTestPlanIdsNil

`func (o *ApiV2TestPointsSearchPostRequest) SetTestPlanIdsNil(b bool)`

 SetTestPlanIdsNil sets the value for TestPlanIds to be an explicit nil

### UnsetTestPlanIds
`func (o *ApiV2TestPointsSearchPostRequest) UnsetTestPlanIds()`

UnsetTestPlanIds ensures that no value is present for TestPlanIds, not even an explicit nil
### GetTestSuiteIds

`func (o *ApiV2TestPointsSearchPostRequest) GetTestSuiteIds() []string`

GetTestSuiteIds returns the TestSuiteIds field if non-nil, zero value otherwise.

### GetTestSuiteIdsOk

`func (o *ApiV2TestPointsSearchPostRequest) GetTestSuiteIdsOk() (*[]string, bool)`

GetTestSuiteIdsOk returns a tuple with the TestSuiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuiteIds

`func (o *ApiV2TestPointsSearchPostRequest) SetTestSuiteIds(v []string)`

SetTestSuiteIds sets TestSuiteIds field to given value.

### HasTestSuiteIds

`func (o *ApiV2TestPointsSearchPostRequest) HasTestSuiteIds() bool`

HasTestSuiteIds returns a boolean if a field has been set.

### SetTestSuiteIdsNil

`func (o *ApiV2TestPointsSearchPostRequest) SetTestSuiteIdsNil(b bool)`

 SetTestSuiteIdsNil sets the value for TestSuiteIds to be an explicit nil

### UnsetTestSuiteIds
`func (o *ApiV2TestPointsSearchPostRequest) UnsetTestSuiteIds()`

UnsetTestSuiteIds ensures that no value is present for TestSuiteIds, not even an explicit nil
### GetWorkItemGlobalIds

`func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemGlobalIds() []int64`

GetWorkItemGlobalIds returns the WorkItemGlobalIds field if non-nil, zero value otherwise.

### GetWorkItemGlobalIdsOk

`func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemGlobalIdsOk() (*[]int64, bool)`

GetWorkItemGlobalIdsOk returns a tuple with the WorkItemGlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemGlobalIds

`func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemGlobalIds(v []int64)`

SetWorkItemGlobalIds sets WorkItemGlobalIds field to given value.

### HasWorkItemGlobalIds

`func (o *ApiV2TestPointsSearchPostRequest) HasWorkItemGlobalIds() bool`

HasWorkItemGlobalIds returns a boolean if a field has been set.

### SetWorkItemGlobalIdsNil

`func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemGlobalIdsNil(b bool)`

 SetWorkItemGlobalIdsNil sets the value for WorkItemGlobalIds to be an explicit nil

### UnsetWorkItemGlobalIds
`func (o *ApiV2TestPointsSearchPostRequest) UnsetWorkItemGlobalIds()`

UnsetWorkItemGlobalIds ensures that no value is present for WorkItemGlobalIds, not even an explicit nil
### GetWorkItemMedianDuration

`func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemMedianDuration() TestPointFilterModelWorkItemMedianDuration`

GetWorkItemMedianDuration returns the WorkItemMedianDuration field if non-nil, zero value otherwise.

### GetWorkItemMedianDurationOk

`func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemMedianDurationOk() (*TestPointFilterModelWorkItemMedianDuration, bool)`

GetWorkItemMedianDurationOk returns a tuple with the WorkItemMedianDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemMedianDuration

`func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemMedianDuration(v TestPointFilterModelWorkItemMedianDuration)`

SetWorkItemMedianDuration sets WorkItemMedianDuration field to given value.

### HasWorkItemMedianDuration

`func (o *ApiV2TestPointsSearchPostRequest) HasWorkItemMedianDuration() bool`

HasWorkItemMedianDuration returns a boolean if a field has been set.

### SetWorkItemMedianDurationNil

`func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemMedianDurationNil(b bool)`

 SetWorkItemMedianDurationNil sets the value for WorkItemMedianDuration to be an explicit nil

### UnsetWorkItemMedianDuration
`func (o *ApiV2TestPointsSearchPostRequest) UnsetWorkItemMedianDuration()`

UnsetWorkItemMedianDuration ensures that no value is present for WorkItemMedianDuration, not even an explicit nil
### GetStatuses

`func (o *ApiV2TestPointsSearchPostRequest) GetStatuses() []TestPointStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *ApiV2TestPointsSearchPostRequest) GetStatusesOk() (*[]TestPointStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *ApiV2TestPointsSearchPostRequest) SetStatuses(v []TestPointStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *ApiV2TestPointsSearchPostRequest) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### SetStatusesNil

`func (o *ApiV2TestPointsSearchPostRequest) SetStatusesNil(b bool)`

 SetStatusesNil sets the value for Statuses to be an explicit nil

### UnsetStatuses
`func (o *ApiV2TestPointsSearchPostRequest) UnsetStatuses()`

UnsetStatuses ensures that no value is present for Statuses, not even an explicit nil
### GetPriorities

`func (o *ApiV2TestPointsSearchPostRequest) GetPriorities() []WorkItemPriorityModel`

GetPriorities returns the Priorities field if non-nil, zero value otherwise.

### GetPrioritiesOk

`func (o *ApiV2TestPointsSearchPostRequest) GetPrioritiesOk() (*[]WorkItemPriorityModel, bool)`

GetPrioritiesOk returns a tuple with the Priorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorities

`func (o *ApiV2TestPointsSearchPostRequest) SetPriorities(v []WorkItemPriorityModel)`

SetPriorities sets Priorities field to given value.

### HasPriorities

`func (o *ApiV2TestPointsSearchPostRequest) HasPriorities() bool`

HasPriorities returns a boolean if a field has been set.

### SetPrioritiesNil

`func (o *ApiV2TestPointsSearchPostRequest) SetPrioritiesNil(b bool)`

 SetPrioritiesNil sets the value for Priorities to be an explicit nil

### UnsetPriorities
`func (o *ApiV2TestPointsSearchPostRequest) UnsetPriorities()`

UnsetPriorities ensures that no value is present for Priorities, not even an explicit nil
### GetIsAutomated

`func (o *ApiV2TestPointsSearchPostRequest) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *ApiV2TestPointsSearchPostRequest) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *ApiV2TestPointsSearchPostRequest) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *ApiV2TestPointsSearchPostRequest) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### SetIsAutomatedNil

`func (o *ApiV2TestPointsSearchPostRequest) SetIsAutomatedNil(b bool)`

 SetIsAutomatedNil sets the value for IsAutomated to be an explicit nil

### UnsetIsAutomated
`func (o *ApiV2TestPointsSearchPostRequest) UnsetIsAutomated()`

UnsetIsAutomated ensures that no value is present for IsAutomated, not even an explicit nil
### GetName

`func (o *ApiV2TestPointsSearchPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2TestPointsSearchPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2TestPointsSearchPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV2TestPointsSearchPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApiV2TestPointsSearchPostRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApiV2TestPointsSearchPostRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetConfigurationIds

`func (o *ApiV2TestPointsSearchPostRequest) GetConfigurationIds() []string`

GetConfigurationIds returns the ConfigurationIds field if non-nil, zero value otherwise.

### GetConfigurationIdsOk

`func (o *ApiV2TestPointsSearchPostRequest) GetConfigurationIdsOk() (*[]string, bool)`

GetConfigurationIdsOk returns a tuple with the ConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIds

`func (o *ApiV2TestPointsSearchPostRequest) SetConfigurationIds(v []string)`

SetConfigurationIds sets ConfigurationIds field to given value.

### HasConfigurationIds

`func (o *ApiV2TestPointsSearchPostRequest) HasConfigurationIds() bool`

HasConfigurationIds returns a boolean if a field has been set.

### SetConfigurationIdsNil

`func (o *ApiV2TestPointsSearchPostRequest) SetConfigurationIdsNil(b bool)`

 SetConfigurationIdsNil sets the value for ConfigurationIds to be an explicit nil

### UnsetConfigurationIds
`func (o *ApiV2TestPointsSearchPostRequest) UnsetConfigurationIds()`

UnsetConfigurationIds ensures that no value is present for ConfigurationIds, not even an explicit nil
### GetTesterIds

`func (o *ApiV2TestPointsSearchPostRequest) GetTesterIds() []string`

GetTesterIds returns the TesterIds field if non-nil, zero value otherwise.

### GetTesterIdsOk

`func (o *ApiV2TestPointsSearchPostRequest) GetTesterIdsOk() (*[]string, bool)`

GetTesterIdsOk returns a tuple with the TesterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesterIds

`func (o *ApiV2TestPointsSearchPostRequest) SetTesterIds(v []string)`

SetTesterIds sets TesterIds field to given value.

### HasTesterIds

`func (o *ApiV2TestPointsSearchPostRequest) HasTesterIds() bool`

HasTesterIds returns a boolean if a field has been set.

### SetTesterIdsNil

`func (o *ApiV2TestPointsSearchPostRequest) SetTesterIdsNil(b bool)`

 SetTesterIdsNil sets the value for TesterIds to be an explicit nil

### UnsetTesterIds
`func (o *ApiV2TestPointsSearchPostRequest) UnsetTesterIds()`

UnsetTesterIds ensures that no value is present for TesterIds, not even an explicit nil
### GetDuration

`func (o *ApiV2TestPointsSearchPostRequest) GetDuration() TestPointFilterModelDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ApiV2TestPointsSearchPostRequest) GetDurationOk() (*TestPointFilterModelDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ApiV2TestPointsSearchPostRequest) SetDuration(v TestPointFilterModelDuration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ApiV2TestPointsSearchPostRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *ApiV2TestPointsSearchPostRequest) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *ApiV2TestPointsSearchPostRequest) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetSectionIds

`func (o *ApiV2TestPointsSearchPostRequest) GetSectionIds() []string`

GetSectionIds returns the SectionIds field if non-nil, zero value otherwise.

### GetSectionIdsOk

`func (o *ApiV2TestPointsSearchPostRequest) GetSectionIdsOk() (*[]string, bool)`

GetSectionIdsOk returns a tuple with the SectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionIds

`func (o *ApiV2TestPointsSearchPostRequest) SetSectionIds(v []string)`

SetSectionIds sets SectionIds field to given value.

### HasSectionIds

`func (o *ApiV2TestPointsSearchPostRequest) HasSectionIds() bool`

HasSectionIds returns a boolean if a field has been set.

### SetSectionIdsNil

`func (o *ApiV2TestPointsSearchPostRequest) SetSectionIdsNil(b bool)`

 SetSectionIdsNil sets the value for SectionIds to be an explicit nil

### UnsetSectionIds
`func (o *ApiV2TestPointsSearchPostRequest) UnsetSectionIds()`

UnsetSectionIds ensures that no value is present for SectionIds, not even an explicit nil
### GetCreatedDate

`func (o *ApiV2TestPointsSearchPostRequest) GetCreatedDate() TestPointFilterModelCreatedDate`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ApiV2TestPointsSearchPostRequest) GetCreatedDateOk() (*TestPointFilterModelCreatedDate, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ApiV2TestPointsSearchPostRequest) SetCreatedDate(v TestPointFilterModelCreatedDate)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ApiV2TestPointsSearchPostRequest) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *ApiV2TestPointsSearchPostRequest) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *ApiV2TestPointsSearchPostRequest) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetCreatedByIds

`func (o *ApiV2TestPointsSearchPostRequest) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *ApiV2TestPointsSearchPostRequest) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *ApiV2TestPointsSearchPostRequest) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *ApiV2TestPointsSearchPostRequest) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *ApiV2TestPointsSearchPostRequest) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *ApiV2TestPointsSearchPostRequest) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedDate

`func (o *ApiV2TestPointsSearchPostRequest) GetModifiedDate() TestPointFilterModelModifiedDate`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ApiV2TestPointsSearchPostRequest) GetModifiedDateOk() (*TestPointFilterModelModifiedDate, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ApiV2TestPointsSearchPostRequest) SetModifiedDate(v TestPointFilterModelModifiedDate)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *ApiV2TestPointsSearchPostRequest) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *ApiV2TestPointsSearchPostRequest) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *ApiV2TestPointsSearchPostRequest) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedByIds

`func (o *ApiV2TestPointsSearchPostRequest) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *ApiV2TestPointsSearchPostRequest) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *ApiV2TestPointsSearchPostRequest) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *ApiV2TestPointsSearchPostRequest) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *ApiV2TestPointsSearchPostRequest) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *ApiV2TestPointsSearchPostRequest) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil
### GetTags

`func (o *ApiV2TestPointsSearchPostRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApiV2TestPointsSearchPostRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApiV2TestPointsSearchPostRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApiV2TestPointsSearchPostRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ApiV2TestPointsSearchPostRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ApiV2TestPointsSearchPostRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetAttributes

`func (o *ApiV2TestPointsSearchPostRequest) GetAttributes() map[string][]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApiV2TestPointsSearchPostRequest) GetAttributesOk() (*map[string][]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApiV2TestPointsSearchPostRequest) SetAttributes(v map[string][]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ApiV2TestPointsSearchPostRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *ApiV2TestPointsSearchPostRequest) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *ApiV2TestPointsSearchPostRequest) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetWorkItemCreatedDate

`func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemCreatedDate() TestPointFilterModelWorkItemCreatedDate`

GetWorkItemCreatedDate returns the WorkItemCreatedDate field if non-nil, zero value otherwise.

### GetWorkItemCreatedDateOk

`func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemCreatedDateOk() (*TestPointFilterModelWorkItemCreatedDate, bool)`

GetWorkItemCreatedDateOk returns a tuple with the WorkItemCreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemCreatedDate

`func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemCreatedDate(v TestPointFilterModelWorkItemCreatedDate)`

SetWorkItemCreatedDate sets WorkItemCreatedDate field to given value.

### HasWorkItemCreatedDate

`func (o *ApiV2TestPointsSearchPostRequest) HasWorkItemCreatedDate() bool`

HasWorkItemCreatedDate returns a boolean if a field has been set.

### SetWorkItemCreatedDateNil

`func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemCreatedDateNil(b bool)`

 SetWorkItemCreatedDateNil sets the value for WorkItemCreatedDate to be an explicit nil

### UnsetWorkItemCreatedDate
`func (o *ApiV2TestPointsSearchPostRequest) UnsetWorkItemCreatedDate()`

UnsetWorkItemCreatedDate ensures that no value is present for WorkItemCreatedDate, not even an explicit nil
### GetWorkItemCreatedByIds

`func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemCreatedByIds() []string`

GetWorkItemCreatedByIds returns the WorkItemCreatedByIds field if non-nil, zero value otherwise.

### GetWorkItemCreatedByIdsOk

`func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemCreatedByIdsOk() (*[]string, bool)`

GetWorkItemCreatedByIdsOk returns a tuple with the WorkItemCreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemCreatedByIds

`func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemCreatedByIds(v []string)`

SetWorkItemCreatedByIds sets WorkItemCreatedByIds field to given value.

### HasWorkItemCreatedByIds

`func (o *ApiV2TestPointsSearchPostRequest) HasWorkItemCreatedByIds() bool`

HasWorkItemCreatedByIds returns a boolean if a field has been set.

### SetWorkItemCreatedByIdsNil

`func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemCreatedByIdsNil(b bool)`

 SetWorkItemCreatedByIdsNil sets the value for WorkItemCreatedByIds to be an explicit nil

### UnsetWorkItemCreatedByIds
`func (o *ApiV2TestPointsSearchPostRequest) UnsetWorkItemCreatedByIds()`

UnsetWorkItemCreatedByIds ensures that no value is present for WorkItemCreatedByIds, not even an explicit nil
### GetWorkItemModifiedDate

`func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemModifiedDate() TestPointFilterModelWorkItemModifiedDate`

GetWorkItemModifiedDate returns the WorkItemModifiedDate field if non-nil, zero value otherwise.

### GetWorkItemModifiedDateOk

`func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemModifiedDateOk() (*TestPointFilterModelWorkItemModifiedDate, bool)`

GetWorkItemModifiedDateOk returns a tuple with the WorkItemModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemModifiedDate

`func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemModifiedDate(v TestPointFilterModelWorkItemModifiedDate)`

SetWorkItemModifiedDate sets WorkItemModifiedDate field to given value.

### HasWorkItemModifiedDate

`func (o *ApiV2TestPointsSearchPostRequest) HasWorkItemModifiedDate() bool`

HasWorkItemModifiedDate returns a boolean if a field has been set.

### SetWorkItemModifiedDateNil

`func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemModifiedDateNil(b bool)`

 SetWorkItemModifiedDateNil sets the value for WorkItemModifiedDate to be an explicit nil

### UnsetWorkItemModifiedDate
`func (o *ApiV2TestPointsSearchPostRequest) UnsetWorkItemModifiedDate()`

UnsetWorkItemModifiedDate ensures that no value is present for WorkItemModifiedDate, not even an explicit nil
### GetWorkItemModifiedByIds

`func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemModifiedByIds() []string`

GetWorkItemModifiedByIds returns the WorkItemModifiedByIds field if non-nil, zero value otherwise.

### GetWorkItemModifiedByIdsOk

`func (o *ApiV2TestPointsSearchPostRequest) GetWorkItemModifiedByIdsOk() (*[]string, bool)`

GetWorkItemModifiedByIdsOk returns a tuple with the WorkItemModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemModifiedByIds

`func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemModifiedByIds(v []string)`

SetWorkItemModifiedByIds sets WorkItemModifiedByIds field to given value.

### HasWorkItemModifiedByIds

`func (o *ApiV2TestPointsSearchPostRequest) HasWorkItemModifiedByIds() bool`

HasWorkItemModifiedByIds returns a boolean if a field has been set.

### SetWorkItemModifiedByIdsNil

`func (o *ApiV2TestPointsSearchPostRequest) SetWorkItemModifiedByIdsNil(b bool)`

 SetWorkItemModifiedByIdsNil sets the value for WorkItemModifiedByIds to be an explicit nil

### UnsetWorkItemModifiedByIds
`func (o *ApiV2TestPointsSearchPostRequest) UnsetWorkItemModifiedByIds()`

UnsetWorkItemModifiedByIds ensures that no value is present for WorkItemModifiedByIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


