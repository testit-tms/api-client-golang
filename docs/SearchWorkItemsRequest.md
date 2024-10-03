# SearchWorkItemsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagNames** | Pointer to **[]string** | Collection of tags | [optional] 
**EntityTypes** | Pointer to [**[]WorkItemEntityTypes**](WorkItemEntityTypes.md) | Collection of types of work item   Allowed values: &#x60;TestCases&#x60;, &#x60;CheckLists&#x60;, &#x60;SharedSteps&#x60; | [optional] 
**NameOrId** | Pointer to **NullableString** | Name or identifier (UUID) of work item | [optional] 
**IncludeIds** | Pointer to **[]string** | Collection of identifiers of work items which need to be included in result regardless of filtering | [optional] 
**ExcludeIds** | Pointer to **[]string** | Collection of identifiers of work items which need to be excluded from result regardless of filtering | [optional] 
**ProjectIds** | Pointer to **[]string** | Collection of project identifiers | [optional] 
**Name** | Pointer to **NullableString** | Name of work item | [optional] 
**Ids** | Pointer to **[]string** | Specifies a work item unique IDs to search for | [optional] 
**GlobalIds** | Pointer to **[]int64** | Collection of global (integer) identifiers | [optional] 
**Attributes** | Pointer to **map[string][]string** | Custom attributes of work item | [optional] 
**IsDeleted** | Pointer to **NullableBool** | Is result must consist of only actual/deleted work items | [optional] 
**SectionIds** | Pointer to **[]string** | Collection of section identifiers | [optional] 
**CreatedByIds** | Pointer to **[]string** | Collection of identifiers of users who created work item | [optional] 
**ModifiedByIds** | Pointer to **[]string** | Collection of identifiers of users who applied last modification to work item | [optional] 
**States** | Pointer to [**[]WorkItemStates**](WorkItemStates.md) | Collection of states of work item | [optional] 
**Priorities** | Pointer to [**[]WorkItemPriorityModel**](WorkItemPriorityModel.md) | Collection of priorities of work item | [optional] 
**Types** | Pointer to [**[]WorkItemEntityTypes**](WorkItemEntityTypes.md) | Collection of types of work item | [optional] 
**CreatedDate** | Pointer to [**NullableTestPointFilterModelWorkItemCreatedDate**](TestPointFilterModelWorkItemCreatedDate.md) |  | [optional] 
**ModifiedDate** | Pointer to [**NullableTestPointFilterModelWorkItemModifiedDate**](TestPointFilterModelWorkItemModifiedDate.md) |  | [optional] 
**Duration** | Pointer to [**NullableTestSuiteWorkItemsSearchModelDuration**](TestSuiteWorkItemsSearchModelDuration.md) |  | [optional] 
**MedianDuration** | Pointer to [**NullableTestSuiteWorkItemsSearchModelMedianDuration**](TestSuiteWorkItemsSearchModelMedianDuration.md) |  | [optional] 
**IsAutomated** | Pointer to **NullableBool** | Is result must consist of only manual/automated work items | [optional] 
**Tags** | Pointer to **[]string** | Collection of tags | [optional] 
**AutoTestIds** | Pointer to **[]string** | Collection of identifiers of linked autotests | [optional] 
**WorkItemVersionIds** | Pointer to **[]string** | Collection of identifiers work items versions. | [optional] 

## Methods

### NewSearchWorkItemsRequest

`func NewSearchWorkItemsRequest() *SearchWorkItemsRequest`

NewSearchWorkItemsRequest instantiates a new SearchWorkItemsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchWorkItemsRequestWithDefaults

`func NewSearchWorkItemsRequestWithDefaults() *SearchWorkItemsRequest`

NewSearchWorkItemsRequestWithDefaults instantiates a new SearchWorkItemsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagNames

`func (o *SearchWorkItemsRequest) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *SearchWorkItemsRequest) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *SearchWorkItemsRequest) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *SearchWorkItemsRequest) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### SetTagNamesNil

`func (o *SearchWorkItemsRequest) SetTagNamesNil(b bool)`

 SetTagNamesNil sets the value for TagNames to be an explicit nil

### UnsetTagNames
`func (o *SearchWorkItemsRequest) UnsetTagNames()`

UnsetTagNames ensures that no value is present for TagNames, not even an explicit nil
### GetEntityTypes

`func (o *SearchWorkItemsRequest) GetEntityTypes() []WorkItemEntityTypes`

GetEntityTypes returns the EntityTypes field if non-nil, zero value otherwise.

### GetEntityTypesOk

`func (o *SearchWorkItemsRequest) GetEntityTypesOk() (*[]WorkItemEntityTypes, bool)`

GetEntityTypesOk returns a tuple with the EntityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypes

`func (o *SearchWorkItemsRequest) SetEntityTypes(v []WorkItemEntityTypes)`

SetEntityTypes sets EntityTypes field to given value.

### HasEntityTypes

`func (o *SearchWorkItemsRequest) HasEntityTypes() bool`

HasEntityTypes returns a boolean if a field has been set.

### SetEntityTypesNil

`func (o *SearchWorkItemsRequest) SetEntityTypesNil(b bool)`

 SetEntityTypesNil sets the value for EntityTypes to be an explicit nil

### UnsetEntityTypes
`func (o *SearchWorkItemsRequest) UnsetEntityTypes()`

UnsetEntityTypes ensures that no value is present for EntityTypes, not even an explicit nil
### GetNameOrId

`func (o *SearchWorkItemsRequest) GetNameOrId() string`

GetNameOrId returns the NameOrId field if non-nil, zero value otherwise.

### GetNameOrIdOk

`func (o *SearchWorkItemsRequest) GetNameOrIdOk() (*string, bool)`

GetNameOrIdOk returns a tuple with the NameOrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOrId

`func (o *SearchWorkItemsRequest) SetNameOrId(v string)`

SetNameOrId sets NameOrId field to given value.

### HasNameOrId

`func (o *SearchWorkItemsRequest) HasNameOrId() bool`

HasNameOrId returns a boolean if a field has been set.

### SetNameOrIdNil

`func (o *SearchWorkItemsRequest) SetNameOrIdNil(b bool)`

 SetNameOrIdNil sets the value for NameOrId to be an explicit nil

### UnsetNameOrId
`func (o *SearchWorkItemsRequest) UnsetNameOrId()`

UnsetNameOrId ensures that no value is present for NameOrId, not even an explicit nil
### GetIncludeIds

`func (o *SearchWorkItemsRequest) GetIncludeIds() []string`

GetIncludeIds returns the IncludeIds field if non-nil, zero value otherwise.

### GetIncludeIdsOk

`func (o *SearchWorkItemsRequest) GetIncludeIdsOk() (*[]string, bool)`

GetIncludeIdsOk returns a tuple with the IncludeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeIds

`func (o *SearchWorkItemsRequest) SetIncludeIds(v []string)`

SetIncludeIds sets IncludeIds field to given value.

### HasIncludeIds

`func (o *SearchWorkItemsRequest) HasIncludeIds() bool`

HasIncludeIds returns a boolean if a field has been set.

### SetIncludeIdsNil

`func (o *SearchWorkItemsRequest) SetIncludeIdsNil(b bool)`

 SetIncludeIdsNil sets the value for IncludeIds to be an explicit nil

### UnsetIncludeIds
`func (o *SearchWorkItemsRequest) UnsetIncludeIds()`

UnsetIncludeIds ensures that no value is present for IncludeIds, not even an explicit nil
### GetExcludeIds

`func (o *SearchWorkItemsRequest) GetExcludeIds() []string`

GetExcludeIds returns the ExcludeIds field if non-nil, zero value otherwise.

### GetExcludeIdsOk

`func (o *SearchWorkItemsRequest) GetExcludeIdsOk() (*[]string, bool)`

GetExcludeIdsOk returns a tuple with the ExcludeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeIds

`func (o *SearchWorkItemsRequest) SetExcludeIds(v []string)`

SetExcludeIds sets ExcludeIds field to given value.

### HasExcludeIds

`func (o *SearchWorkItemsRequest) HasExcludeIds() bool`

HasExcludeIds returns a boolean if a field has been set.

### SetExcludeIdsNil

`func (o *SearchWorkItemsRequest) SetExcludeIdsNil(b bool)`

 SetExcludeIdsNil sets the value for ExcludeIds to be an explicit nil

### UnsetExcludeIds
`func (o *SearchWorkItemsRequest) UnsetExcludeIds()`

UnsetExcludeIds ensures that no value is present for ExcludeIds, not even an explicit nil
### GetProjectIds

`func (o *SearchWorkItemsRequest) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *SearchWorkItemsRequest) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *SearchWorkItemsRequest) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *SearchWorkItemsRequest) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *SearchWorkItemsRequest) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *SearchWorkItemsRequest) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetName

`func (o *SearchWorkItemsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchWorkItemsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchWorkItemsRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchWorkItemsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SearchWorkItemsRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SearchWorkItemsRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIds

`func (o *SearchWorkItemsRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *SearchWorkItemsRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *SearchWorkItemsRequest) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *SearchWorkItemsRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *SearchWorkItemsRequest) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *SearchWorkItemsRequest) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetGlobalIds

`func (o *SearchWorkItemsRequest) GetGlobalIds() []int64`

GetGlobalIds returns the GlobalIds field if non-nil, zero value otherwise.

### GetGlobalIdsOk

`func (o *SearchWorkItemsRequest) GetGlobalIdsOk() (*[]int64, bool)`

GetGlobalIdsOk returns a tuple with the GlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIds

`func (o *SearchWorkItemsRequest) SetGlobalIds(v []int64)`

SetGlobalIds sets GlobalIds field to given value.

### HasGlobalIds

`func (o *SearchWorkItemsRequest) HasGlobalIds() bool`

HasGlobalIds returns a boolean if a field has been set.

### SetGlobalIdsNil

`func (o *SearchWorkItemsRequest) SetGlobalIdsNil(b bool)`

 SetGlobalIdsNil sets the value for GlobalIds to be an explicit nil

### UnsetGlobalIds
`func (o *SearchWorkItemsRequest) UnsetGlobalIds()`

UnsetGlobalIds ensures that no value is present for GlobalIds, not even an explicit nil
### GetAttributes

`func (o *SearchWorkItemsRequest) GetAttributes() map[string][]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SearchWorkItemsRequest) GetAttributesOk() (*map[string][]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SearchWorkItemsRequest) SetAttributes(v map[string][]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SearchWorkItemsRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *SearchWorkItemsRequest) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *SearchWorkItemsRequest) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetIsDeleted

`func (o *SearchWorkItemsRequest) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *SearchWorkItemsRequest) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *SearchWorkItemsRequest) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *SearchWorkItemsRequest) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *SearchWorkItemsRequest) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *SearchWorkItemsRequest) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetSectionIds

`func (o *SearchWorkItemsRequest) GetSectionIds() []string`

GetSectionIds returns the SectionIds field if non-nil, zero value otherwise.

### GetSectionIdsOk

`func (o *SearchWorkItemsRequest) GetSectionIdsOk() (*[]string, bool)`

GetSectionIdsOk returns a tuple with the SectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionIds

`func (o *SearchWorkItemsRequest) SetSectionIds(v []string)`

SetSectionIds sets SectionIds field to given value.

### HasSectionIds

`func (o *SearchWorkItemsRequest) HasSectionIds() bool`

HasSectionIds returns a boolean if a field has been set.

### SetSectionIdsNil

`func (o *SearchWorkItemsRequest) SetSectionIdsNil(b bool)`

 SetSectionIdsNil sets the value for SectionIds to be an explicit nil

### UnsetSectionIds
`func (o *SearchWorkItemsRequest) UnsetSectionIds()`

UnsetSectionIds ensures that no value is present for SectionIds, not even an explicit nil
### GetCreatedByIds

`func (o *SearchWorkItemsRequest) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *SearchWorkItemsRequest) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *SearchWorkItemsRequest) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *SearchWorkItemsRequest) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *SearchWorkItemsRequest) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *SearchWorkItemsRequest) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedByIds

`func (o *SearchWorkItemsRequest) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *SearchWorkItemsRequest) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *SearchWorkItemsRequest) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *SearchWorkItemsRequest) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *SearchWorkItemsRequest) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *SearchWorkItemsRequest) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil
### GetStates

`func (o *SearchWorkItemsRequest) GetStates() []WorkItemStates`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *SearchWorkItemsRequest) GetStatesOk() (*[]WorkItemStates, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *SearchWorkItemsRequest) SetStates(v []WorkItemStates)`

SetStates sets States field to given value.

### HasStates

`func (o *SearchWorkItemsRequest) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStatesNil

`func (o *SearchWorkItemsRequest) SetStatesNil(b bool)`

 SetStatesNil sets the value for States to be an explicit nil

### UnsetStates
`func (o *SearchWorkItemsRequest) UnsetStates()`

UnsetStates ensures that no value is present for States, not even an explicit nil
### GetPriorities

`func (o *SearchWorkItemsRequest) GetPriorities() []WorkItemPriorityModel`

GetPriorities returns the Priorities field if non-nil, zero value otherwise.

### GetPrioritiesOk

`func (o *SearchWorkItemsRequest) GetPrioritiesOk() (*[]WorkItemPriorityModel, bool)`

GetPrioritiesOk returns a tuple with the Priorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorities

`func (o *SearchWorkItemsRequest) SetPriorities(v []WorkItemPriorityModel)`

SetPriorities sets Priorities field to given value.

### HasPriorities

`func (o *SearchWorkItemsRequest) HasPriorities() bool`

HasPriorities returns a boolean if a field has been set.

### SetPrioritiesNil

`func (o *SearchWorkItemsRequest) SetPrioritiesNil(b bool)`

 SetPrioritiesNil sets the value for Priorities to be an explicit nil

### UnsetPriorities
`func (o *SearchWorkItemsRequest) UnsetPriorities()`

UnsetPriorities ensures that no value is present for Priorities, not even an explicit nil
### GetTypes

`func (o *SearchWorkItemsRequest) GetTypes() []WorkItemEntityTypes`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *SearchWorkItemsRequest) GetTypesOk() (*[]WorkItemEntityTypes, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *SearchWorkItemsRequest) SetTypes(v []WorkItemEntityTypes)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *SearchWorkItemsRequest) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *SearchWorkItemsRequest) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *SearchWorkItemsRequest) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil
### GetCreatedDate

`func (o *SearchWorkItemsRequest) GetCreatedDate() TestPointFilterModelWorkItemCreatedDate`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *SearchWorkItemsRequest) GetCreatedDateOk() (*TestPointFilterModelWorkItemCreatedDate, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *SearchWorkItemsRequest) SetCreatedDate(v TestPointFilterModelWorkItemCreatedDate)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *SearchWorkItemsRequest) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *SearchWorkItemsRequest) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *SearchWorkItemsRequest) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetModifiedDate

`func (o *SearchWorkItemsRequest) GetModifiedDate() TestPointFilterModelWorkItemModifiedDate`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *SearchWorkItemsRequest) GetModifiedDateOk() (*TestPointFilterModelWorkItemModifiedDate, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *SearchWorkItemsRequest) SetModifiedDate(v TestPointFilterModelWorkItemModifiedDate)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *SearchWorkItemsRequest) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *SearchWorkItemsRequest) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *SearchWorkItemsRequest) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetDuration

`func (o *SearchWorkItemsRequest) GetDuration() TestSuiteWorkItemsSearchModelDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SearchWorkItemsRequest) GetDurationOk() (*TestSuiteWorkItemsSearchModelDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SearchWorkItemsRequest) SetDuration(v TestSuiteWorkItemsSearchModelDuration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SearchWorkItemsRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *SearchWorkItemsRequest) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *SearchWorkItemsRequest) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetMedianDuration

`func (o *SearchWorkItemsRequest) GetMedianDuration() TestSuiteWorkItemsSearchModelMedianDuration`

GetMedianDuration returns the MedianDuration field if non-nil, zero value otherwise.

### GetMedianDurationOk

`func (o *SearchWorkItemsRequest) GetMedianDurationOk() (*TestSuiteWorkItemsSearchModelMedianDuration, bool)`

GetMedianDurationOk returns a tuple with the MedianDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianDuration

`func (o *SearchWorkItemsRequest) SetMedianDuration(v TestSuiteWorkItemsSearchModelMedianDuration)`

SetMedianDuration sets MedianDuration field to given value.

### HasMedianDuration

`func (o *SearchWorkItemsRequest) HasMedianDuration() bool`

HasMedianDuration returns a boolean if a field has been set.

### SetMedianDurationNil

`func (o *SearchWorkItemsRequest) SetMedianDurationNil(b bool)`

 SetMedianDurationNil sets the value for MedianDuration to be an explicit nil

### UnsetMedianDuration
`func (o *SearchWorkItemsRequest) UnsetMedianDuration()`

UnsetMedianDuration ensures that no value is present for MedianDuration, not even an explicit nil
### GetIsAutomated

`func (o *SearchWorkItemsRequest) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *SearchWorkItemsRequest) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *SearchWorkItemsRequest) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *SearchWorkItemsRequest) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### SetIsAutomatedNil

`func (o *SearchWorkItemsRequest) SetIsAutomatedNil(b bool)`

 SetIsAutomatedNil sets the value for IsAutomated to be an explicit nil

### UnsetIsAutomated
`func (o *SearchWorkItemsRequest) UnsetIsAutomated()`

UnsetIsAutomated ensures that no value is present for IsAutomated, not even an explicit nil
### GetTags

`func (o *SearchWorkItemsRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SearchWorkItemsRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SearchWorkItemsRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SearchWorkItemsRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *SearchWorkItemsRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *SearchWorkItemsRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetAutoTestIds

`func (o *SearchWorkItemsRequest) GetAutoTestIds() []string`

GetAutoTestIds returns the AutoTestIds field if non-nil, zero value otherwise.

### GetAutoTestIdsOk

`func (o *SearchWorkItemsRequest) GetAutoTestIdsOk() (*[]string, bool)`

GetAutoTestIdsOk returns a tuple with the AutoTestIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestIds

`func (o *SearchWorkItemsRequest) SetAutoTestIds(v []string)`

SetAutoTestIds sets AutoTestIds field to given value.

### HasAutoTestIds

`func (o *SearchWorkItemsRequest) HasAutoTestIds() bool`

HasAutoTestIds returns a boolean if a field has been set.

### SetAutoTestIdsNil

`func (o *SearchWorkItemsRequest) SetAutoTestIdsNil(b bool)`

 SetAutoTestIdsNil sets the value for AutoTestIds to be an explicit nil

### UnsetAutoTestIds
`func (o *SearchWorkItemsRequest) UnsetAutoTestIds()`

UnsetAutoTestIds ensures that no value is present for AutoTestIds, not even an explicit nil
### GetWorkItemVersionIds

`func (o *SearchWorkItemsRequest) GetWorkItemVersionIds() []string`

GetWorkItemVersionIds returns the WorkItemVersionIds field if non-nil, zero value otherwise.

### GetWorkItemVersionIdsOk

`func (o *SearchWorkItemsRequest) GetWorkItemVersionIdsOk() (*[]string, bool)`

GetWorkItemVersionIdsOk returns a tuple with the WorkItemVersionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionIds

`func (o *SearchWorkItemsRequest) SetWorkItemVersionIds(v []string)`

SetWorkItemVersionIds sets WorkItemVersionIds field to given value.

### HasWorkItemVersionIds

`func (o *SearchWorkItemsRequest) HasWorkItemVersionIds() bool`

HasWorkItemVersionIds returns a boolean if a field has been set.

### SetWorkItemVersionIdsNil

`func (o *SearchWorkItemsRequest) SetWorkItemVersionIdsNil(b bool)`

 SetWorkItemVersionIdsNil sets the value for WorkItemVersionIds to be an explicit nil

### UnsetWorkItemVersionIds
`func (o *SearchWorkItemsRequest) UnsetWorkItemVersionIds()`

UnsetWorkItemVersionIds ensures that no value is present for WorkItemVersionIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


