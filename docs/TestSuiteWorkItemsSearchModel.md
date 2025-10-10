# TestSuiteWorkItemsSearchModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagNames** | Pointer to **[]string** | Collection of tags | [optional] 
**EntityTypes** | Pointer to [**[]WorkItemEntityTypes**](WorkItemEntityTypes.md) | Collection of types of work item  Allowed values: &#x60;TestCases&#x60;, &#x60;CheckLists&#x60;, &#x60;SharedSteps&#x60; | [optional] 
**NameOrId** | Pointer to **NullableString** | Name or identifier (UUID) of work item | [optional] 
**IncludeIds** | Pointer to **[]string** | Collection of identifiers of work items which need to be included in result regardless of filtering | [optional] 
**ExcludeIds** | Pointer to **[]string** | Collection of identifiers of work items which need to be excluded from result regardless of filtering | [optional] 
**ExternalMetadata** | Pointer to [**NullableWorkItemExternalMetadataFilterModel**](WorkItemExternalMetadataFilterModel.md) | Specifies work item filter by its external metadata | [optional] 
**ProjectIds** | Pointer to **[]string** | Collection of project identifiers | [optional] 
**Links** | Pointer to [**NullableWorkItemLinkFilterModel**](WorkItemLinkFilterModel.md) | Specifies a work item filter by its links | [optional] 
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
**SourceTypes** | Pointer to [**[]WorkItemSourceTypeModel**](WorkItemSourceTypeModel.md) | Collection of priorities of work item | [optional] 
**Types** | Pointer to [**[]WorkItemEntityTypes**](WorkItemEntityTypes.md) | Collection of types of work item | [optional] 
**CreatedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) | Specifies a work item range of creation date to search for | [optional] 
**ModifiedDate** | Pointer to [**NullableDateTimeRangeSelectorModel**](DateTimeRangeSelectorModel.md) | Specifies a work item range of last modification date to search for | [optional] 
**Duration** | Pointer to [**NullableInt32RangeSelectorModel**](Int32RangeSelectorModel.md) | Specifies a work item duration range to search for | [optional] 
**MedianDuration** | Pointer to [**NullableInt64RangeSelectorModel**](Int64RangeSelectorModel.md) | Specifies a work item median duration range to search for | [optional] 
**IsAutomated** | Pointer to **NullableBool** | Is result must consist of only manual/automated work items | [optional] 
**Tags** | Pointer to **[]string** | Collection of tags | [optional] 
**AutoTestIds** | Pointer to **[]string** | Collection of identifiers of linked autotests | [optional] 
**WorkItemVersionIds** | Pointer to **[]string** | Collection of identifiers work items versions. | [optional] 

## Methods

### NewTestSuiteWorkItemsSearchModel

`func NewTestSuiteWorkItemsSearchModel() *TestSuiteWorkItemsSearchModel`

NewTestSuiteWorkItemsSearchModel instantiates a new TestSuiteWorkItemsSearchModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestSuiteWorkItemsSearchModelWithDefaults

`func NewTestSuiteWorkItemsSearchModelWithDefaults() *TestSuiteWorkItemsSearchModel`

NewTestSuiteWorkItemsSearchModelWithDefaults instantiates a new TestSuiteWorkItemsSearchModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagNames

`func (o *TestSuiteWorkItemsSearchModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *TestSuiteWorkItemsSearchModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *TestSuiteWorkItemsSearchModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *TestSuiteWorkItemsSearchModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### SetTagNamesNil

`func (o *TestSuiteWorkItemsSearchModel) SetTagNamesNil(b bool)`

 SetTagNamesNil sets the value for TagNames to be an explicit nil

### UnsetTagNames
`func (o *TestSuiteWorkItemsSearchModel) UnsetTagNames()`

UnsetTagNames ensures that no value is present for TagNames, not even an explicit nil
### GetEntityTypes

`func (o *TestSuiteWorkItemsSearchModel) GetEntityTypes() []WorkItemEntityTypes`

GetEntityTypes returns the EntityTypes field if non-nil, zero value otherwise.

### GetEntityTypesOk

`func (o *TestSuiteWorkItemsSearchModel) GetEntityTypesOk() (*[]WorkItemEntityTypes, bool)`

GetEntityTypesOk returns a tuple with the EntityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypes

`func (o *TestSuiteWorkItemsSearchModel) SetEntityTypes(v []WorkItemEntityTypes)`

SetEntityTypes sets EntityTypes field to given value.

### HasEntityTypes

`func (o *TestSuiteWorkItemsSearchModel) HasEntityTypes() bool`

HasEntityTypes returns a boolean if a field has been set.

### SetEntityTypesNil

`func (o *TestSuiteWorkItemsSearchModel) SetEntityTypesNil(b bool)`

 SetEntityTypesNil sets the value for EntityTypes to be an explicit nil

### UnsetEntityTypes
`func (o *TestSuiteWorkItemsSearchModel) UnsetEntityTypes()`

UnsetEntityTypes ensures that no value is present for EntityTypes, not even an explicit nil
### GetNameOrId

`func (o *TestSuiteWorkItemsSearchModel) GetNameOrId() string`

GetNameOrId returns the NameOrId field if non-nil, zero value otherwise.

### GetNameOrIdOk

`func (o *TestSuiteWorkItemsSearchModel) GetNameOrIdOk() (*string, bool)`

GetNameOrIdOk returns a tuple with the NameOrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOrId

`func (o *TestSuiteWorkItemsSearchModel) SetNameOrId(v string)`

SetNameOrId sets NameOrId field to given value.

### HasNameOrId

`func (o *TestSuiteWorkItemsSearchModel) HasNameOrId() bool`

HasNameOrId returns a boolean if a field has been set.

### SetNameOrIdNil

`func (o *TestSuiteWorkItemsSearchModel) SetNameOrIdNil(b bool)`

 SetNameOrIdNil sets the value for NameOrId to be an explicit nil

### UnsetNameOrId
`func (o *TestSuiteWorkItemsSearchModel) UnsetNameOrId()`

UnsetNameOrId ensures that no value is present for NameOrId, not even an explicit nil
### GetIncludeIds

`func (o *TestSuiteWorkItemsSearchModel) GetIncludeIds() []string`

GetIncludeIds returns the IncludeIds field if non-nil, zero value otherwise.

### GetIncludeIdsOk

`func (o *TestSuiteWorkItemsSearchModel) GetIncludeIdsOk() (*[]string, bool)`

GetIncludeIdsOk returns a tuple with the IncludeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeIds

`func (o *TestSuiteWorkItemsSearchModel) SetIncludeIds(v []string)`

SetIncludeIds sets IncludeIds field to given value.

### HasIncludeIds

`func (o *TestSuiteWorkItemsSearchModel) HasIncludeIds() bool`

HasIncludeIds returns a boolean if a field has been set.

### SetIncludeIdsNil

`func (o *TestSuiteWorkItemsSearchModel) SetIncludeIdsNil(b bool)`

 SetIncludeIdsNil sets the value for IncludeIds to be an explicit nil

### UnsetIncludeIds
`func (o *TestSuiteWorkItemsSearchModel) UnsetIncludeIds()`

UnsetIncludeIds ensures that no value is present for IncludeIds, not even an explicit nil
### GetExcludeIds

`func (o *TestSuiteWorkItemsSearchModel) GetExcludeIds() []string`

GetExcludeIds returns the ExcludeIds field if non-nil, zero value otherwise.

### GetExcludeIdsOk

`func (o *TestSuiteWorkItemsSearchModel) GetExcludeIdsOk() (*[]string, bool)`

GetExcludeIdsOk returns a tuple with the ExcludeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeIds

`func (o *TestSuiteWorkItemsSearchModel) SetExcludeIds(v []string)`

SetExcludeIds sets ExcludeIds field to given value.

### HasExcludeIds

`func (o *TestSuiteWorkItemsSearchModel) HasExcludeIds() bool`

HasExcludeIds returns a boolean if a field has been set.

### SetExcludeIdsNil

`func (o *TestSuiteWorkItemsSearchModel) SetExcludeIdsNil(b bool)`

 SetExcludeIdsNil sets the value for ExcludeIds to be an explicit nil

### UnsetExcludeIds
`func (o *TestSuiteWorkItemsSearchModel) UnsetExcludeIds()`

UnsetExcludeIds ensures that no value is present for ExcludeIds, not even an explicit nil
### GetExternalMetadata

`func (o *TestSuiteWorkItemsSearchModel) GetExternalMetadata() WorkItemExternalMetadataFilterModel`

GetExternalMetadata returns the ExternalMetadata field if non-nil, zero value otherwise.

### GetExternalMetadataOk

`func (o *TestSuiteWorkItemsSearchModel) GetExternalMetadataOk() (*WorkItemExternalMetadataFilterModel, bool)`

GetExternalMetadataOk returns a tuple with the ExternalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalMetadata

`func (o *TestSuiteWorkItemsSearchModel) SetExternalMetadata(v WorkItemExternalMetadataFilterModel)`

SetExternalMetadata sets ExternalMetadata field to given value.

### HasExternalMetadata

`func (o *TestSuiteWorkItemsSearchModel) HasExternalMetadata() bool`

HasExternalMetadata returns a boolean if a field has been set.

### SetExternalMetadataNil

`func (o *TestSuiteWorkItemsSearchModel) SetExternalMetadataNil(b bool)`

 SetExternalMetadataNil sets the value for ExternalMetadata to be an explicit nil

### UnsetExternalMetadata
`func (o *TestSuiteWorkItemsSearchModel) UnsetExternalMetadata()`

UnsetExternalMetadata ensures that no value is present for ExternalMetadata, not even an explicit nil
### GetProjectIds

`func (o *TestSuiteWorkItemsSearchModel) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *TestSuiteWorkItemsSearchModel) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *TestSuiteWorkItemsSearchModel) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *TestSuiteWorkItemsSearchModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *TestSuiteWorkItemsSearchModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *TestSuiteWorkItemsSearchModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetLinks

`func (o *TestSuiteWorkItemsSearchModel) GetLinks() WorkItemLinkFilterModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TestSuiteWorkItemsSearchModel) GetLinksOk() (*WorkItemLinkFilterModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TestSuiteWorkItemsSearchModel) SetLinks(v WorkItemLinkFilterModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TestSuiteWorkItemsSearchModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TestSuiteWorkItemsSearchModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TestSuiteWorkItemsSearchModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetName

`func (o *TestSuiteWorkItemsSearchModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestSuiteWorkItemsSearchModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestSuiteWorkItemsSearchModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestSuiteWorkItemsSearchModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestSuiteWorkItemsSearchModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestSuiteWorkItemsSearchModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIds

`func (o *TestSuiteWorkItemsSearchModel) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *TestSuiteWorkItemsSearchModel) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *TestSuiteWorkItemsSearchModel) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *TestSuiteWorkItemsSearchModel) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *TestSuiteWorkItemsSearchModel) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *TestSuiteWorkItemsSearchModel) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetGlobalIds

`func (o *TestSuiteWorkItemsSearchModel) GetGlobalIds() []int64`

GetGlobalIds returns the GlobalIds field if non-nil, zero value otherwise.

### GetGlobalIdsOk

`func (o *TestSuiteWorkItemsSearchModel) GetGlobalIdsOk() (*[]int64, bool)`

GetGlobalIdsOk returns a tuple with the GlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIds

`func (o *TestSuiteWorkItemsSearchModel) SetGlobalIds(v []int64)`

SetGlobalIds sets GlobalIds field to given value.

### HasGlobalIds

`func (o *TestSuiteWorkItemsSearchModel) HasGlobalIds() bool`

HasGlobalIds returns a boolean if a field has been set.

### SetGlobalIdsNil

`func (o *TestSuiteWorkItemsSearchModel) SetGlobalIdsNil(b bool)`

 SetGlobalIdsNil sets the value for GlobalIds to be an explicit nil

### UnsetGlobalIds
`func (o *TestSuiteWorkItemsSearchModel) UnsetGlobalIds()`

UnsetGlobalIds ensures that no value is present for GlobalIds, not even an explicit nil
### GetAttributes

`func (o *TestSuiteWorkItemsSearchModel) GetAttributes() map[string][]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TestSuiteWorkItemsSearchModel) GetAttributesOk() (*map[string][]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TestSuiteWorkItemsSearchModel) SetAttributes(v map[string][]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TestSuiteWorkItemsSearchModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *TestSuiteWorkItemsSearchModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *TestSuiteWorkItemsSearchModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetIsDeleted

`func (o *TestSuiteWorkItemsSearchModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TestSuiteWorkItemsSearchModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TestSuiteWorkItemsSearchModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *TestSuiteWorkItemsSearchModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *TestSuiteWorkItemsSearchModel) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *TestSuiteWorkItemsSearchModel) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetSectionIds

`func (o *TestSuiteWorkItemsSearchModel) GetSectionIds() []string`

GetSectionIds returns the SectionIds field if non-nil, zero value otherwise.

### GetSectionIdsOk

`func (o *TestSuiteWorkItemsSearchModel) GetSectionIdsOk() (*[]string, bool)`

GetSectionIdsOk returns a tuple with the SectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionIds

`func (o *TestSuiteWorkItemsSearchModel) SetSectionIds(v []string)`

SetSectionIds sets SectionIds field to given value.

### HasSectionIds

`func (o *TestSuiteWorkItemsSearchModel) HasSectionIds() bool`

HasSectionIds returns a boolean if a field has been set.

### SetSectionIdsNil

`func (o *TestSuiteWorkItemsSearchModel) SetSectionIdsNil(b bool)`

 SetSectionIdsNil sets the value for SectionIds to be an explicit nil

### UnsetSectionIds
`func (o *TestSuiteWorkItemsSearchModel) UnsetSectionIds()`

UnsetSectionIds ensures that no value is present for SectionIds, not even an explicit nil
### GetCreatedByIds

`func (o *TestSuiteWorkItemsSearchModel) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *TestSuiteWorkItemsSearchModel) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *TestSuiteWorkItemsSearchModel) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *TestSuiteWorkItemsSearchModel) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *TestSuiteWorkItemsSearchModel) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *TestSuiteWorkItemsSearchModel) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedByIds

`func (o *TestSuiteWorkItemsSearchModel) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *TestSuiteWorkItemsSearchModel) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *TestSuiteWorkItemsSearchModel) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *TestSuiteWorkItemsSearchModel) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *TestSuiteWorkItemsSearchModel) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *TestSuiteWorkItemsSearchModel) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil
### GetStates

`func (o *TestSuiteWorkItemsSearchModel) GetStates() []WorkItemStates`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *TestSuiteWorkItemsSearchModel) GetStatesOk() (*[]WorkItemStates, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *TestSuiteWorkItemsSearchModel) SetStates(v []WorkItemStates)`

SetStates sets States field to given value.

### HasStates

`func (o *TestSuiteWorkItemsSearchModel) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStatesNil

`func (o *TestSuiteWorkItemsSearchModel) SetStatesNil(b bool)`

 SetStatesNil sets the value for States to be an explicit nil

### UnsetStates
`func (o *TestSuiteWorkItemsSearchModel) UnsetStates()`

UnsetStates ensures that no value is present for States, not even an explicit nil
### GetPriorities

`func (o *TestSuiteWorkItemsSearchModel) GetPriorities() []WorkItemPriorityModel`

GetPriorities returns the Priorities field if non-nil, zero value otherwise.

### GetPrioritiesOk

`func (o *TestSuiteWorkItemsSearchModel) GetPrioritiesOk() (*[]WorkItemPriorityModel, bool)`

GetPrioritiesOk returns a tuple with the Priorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorities

`func (o *TestSuiteWorkItemsSearchModel) SetPriorities(v []WorkItemPriorityModel)`

SetPriorities sets Priorities field to given value.

### HasPriorities

`func (o *TestSuiteWorkItemsSearchModel) HasPriorities() bool`

HasPriorities returns a boolean if a field has been set.

### SetPrioritiesNil

`func (o *TestSuiteWorkItemsSearchModel) SetPrioritiesNil(b bool)`

 SetPrioritiesNil sets the value for Priorities to be an explicit nil

### UnsetPriorities
`func (o *TestSuiteWorkItemsSearchModel) UnsetPriorities()`

UnsetPriorities ensures that no value is present for Priorities, not even an explicit nil
### GetSourceTypes

`func (o *TestSuiteWorkItemsSearchModel) GetSourceTypes() []WorkItemSourceTypeModel`

GetSourceTypes returns the SourceTypes field if non-nil, zero value otherwise.

### GetSourceTypesOk

`func (o *TestSuiteWorkItemsSearchModel) GetSourceTypesOk() (*[]WorkItemSourceTypeModel, bool)`

GetSourceTypesOk returns a tuple with the SourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTypes

`func (o *TestSuiteWorkItemsSearchModel) SetSourceTypes(v []WorkItemSourceTypeModel)`

SetSourceTypes sets SourceTypes field to given value.

### HasSourceTypes

`func (o *TestSuiteWorkItemsSearchModel) HasSourceTypes() bool`

HasSourceTypes returns a boolean if a field has been set.

### SetSourceTypesNil

`func (o *TestSuiteWorkItemsSearchModel) SetSourceTypesNil(b bool)`

 SetSourceTypesNil sets the value for SourceTypes to be an explicit nil

### UnsetSourceTypes
`func (o *TestSuiteWorkItemsSearchModel) UnsetSourceTypes()`

UnsetSourceTypes ensures that no value is present for SourceTypes, not even an explicit nil
### GetTypes

`func (o *TestSuiteWorkItemsSearchModel) GetTypes() []WorkItemEntityTypes`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *TestSuiteWorkItemsSearchModel) GetTypesOk() (*[]WorkItemEntityTypes, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *TestSuiteWorkItemsSearchModel) SetTypes(v []WorkItemEntityTypes)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *TestSuiteWorkItemsSearchModel) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *TestSuiteWorkItemsSearchModel) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *TestSuiteWorkItemsSearchModel) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil
### GetCreatedDate

`func (o *TestSuiteWorkItemsSearchModel) GetCreatedDate() DateTimeRangeSelectorModel`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestSuiteWorkItemsSearchModel) GetCreatedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestSuiteWorkItemsSearchModel) SetCreatedDate(v DateTimeRangeSelectorModel)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TestSuiteWorkItemsSearchModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *TestSuiteWorkItemsSearchModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *TestSuiteWorkItemsSearchModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetModifiedDate

`func (o *TestSuiteWorkItemsSearchModel) GetModifiedDate() DateTimeRangeSelectorModel`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestSuiteWorkItemsSearchModel) GetModifiedDateOk() (*DateTimeRangeSelectorModel, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestSuiteWorkItemsSearchModel) SetModifiedDate(v DateTimeRangeSelectorModel)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestSuiteWorkItemsSearchModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *TestSuiteWorkItemsSearchModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *TestSuiteWorkItemsSearchModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetDuration

`func (o *TestSuiteWorkItemsSearchModel) GetDuration() Int32RangeSelectorModel`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TestSuiteWorkItemsSearchModel) GetDurationOk() (*Int32RangeSelectorModel, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TestSuiteWorkItemsSearchModel) SetDuration(v Int32RangeSelectorModel)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TestSuiteWorkItemsSearchModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *TestSuiteWorkItemsSearchModel) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *TestSuiteWorkItemsSearchModel) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetMedianDuration

`func (o *TestSuiteWorkItemsSearchModel) GetMedianDuration() Int64RangeSelectorModel`

GetMedianDuration returns the MedianDuration field if non-nil, zero value otherwise.

### GetMedianDurationOk

`func (o *TestSuiteWorkItemsSearchModel) GetMedianDurationOk() (*Int64RangeSelectorModel, bool)`

GetMedianDurationOk returns a tuple with the MedianDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianDuration

`func (o *TestSuiteWorkItemsSearchModel) SetMedianDuration(v Int64RangeSelectorModel)`

SetMedianDuration sets MedianDuration field to given value.

### HasMedianDuration

`func (o *TestSuiteWorkItemsSearchModel) HasMedianDuration() bool`

HasMedianDuration returns a boolean if a field has been set.

### SetMedianDurationNil

`func (o *TestSuiteWorkItemsSearchModel) SetMedianDurationNil(b bool)`

 SetMedianDurationNil sets the value for MedianDuration to be an explicit nil

### UnsetMedianDuration
`func (o *TestSuiteWorkItemsSearchModel) UnsetMedianDuration()`

UnsetMedianDuration ensures that no value is present for MedianDuration, not even an explicit nil
### GetIsAutomated

`func (o *TestSuiteWorkItemsSearchModel) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *TestSuiteWorkItemsSearchModel) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *TestSuiteWorkItemsSearchModel) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *TestSuiteWorkItemsSearchModel) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### SetIsAutomatedNil

`func (o *TestSuiteWorkItemsSearchModel) SetIsAutomatedNil(b bool)`

 SetIsAutomatedNil sets the value for IsAutomated to be an explicit nil

### UnsetIsAutomated
`func (o *TestSuiteWorkItemsSearchModel) UnsetIsAutomated()`

UnsetIsAutomated ensures that no value is present for IsAutomated, not even an explicit nil
### GetTags

`func (o *TestSuiteWorkItemsSearchModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestSuiteWorkItemsSearchModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestSuiteWorkItemsSearchModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TestSuiteWorkItemsSearchModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *TestSuiteWorkItemsSearchModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *TestSuiteWorkItemsSearchModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetAutoTestIds

`func (o *TestSuiteWorkItemsSearchModel) GetAutoTestIds() []string`

GetAutoTestIds returns the AutoTestIds field if non-nil, zero value otherwise.

### GetAutoTestIdsOk

`func (o *TestSuiteWorkItemsSearchModel) GetAutoTestIdsOk() (*[]string, bool)`

GetAutoTestIdsOk returns a tuple with the AutoTestIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestIds

`func (o *TestSuiteWorkItemsSearchModel) SetAutoTestIds(v []string)`

SetAutoTestIds sets AutoTestIds field to given value.

### HasAutoTestIds

`func (o *TestSuiteWorkItemsSearchModel) HasAutoTestIds() bool`

HasAutoTestIds returns a boolean if a field has been set.

### SetAutoTestIdsNil

`func (o *TestSuiteWorkItemsSearchModel) SetAutoTestIdsNil(b bool)`

 SetAutoTestIdsNil sets the value for AutoTestIds to be an explicit nil

### UnsetAutoTestIds
`func (o *TestSuiteWorkItemsSearchModel) UnsetAutoTestIds()`

UnsetAutoTestIds ensures that no value is present for AutoTestIds, not even an explicit nil
### GetWorkItemVersionIds

`func (o *TestSuiteWorkItemsSearchModel) GetWorkItemVersionIds() []string`

GetWorkItemVersionIds returns the WorkItemVersionIds field if non-nil, zero value otherwise.

### GetWorkItemVersionIdsOk

`func (o *TestSuiteWorkItemsSearchModel) GetWorkItemVersionIdsOk() (*[]string, bool)`

GetWorkItemVersionIdsOk returns a tuple with the WorkItemVersionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionIds

`func (o *TestSuiteWorkItemsSearchModel) SetWorkItemVersionIds(v []string)`

SetWorkItemVersionIds sets WorkItemVersionIds field to given value.

### HasWorkItemVersionIds

`func (o *TestSuiteWorkItemsSearchModel) HasWorkItemVersionIds() bool`

HasWorkItemVersionIds returns a boolean if a field has been set.

### SetWorkItemVersionIdsNil

`func (o *TestSuiteWorkItemsSearchModel) SetWorkItemVersionIdsNil(b bool)`

 SetWorkItemVersionIdsNil sets the value for WorkItemVersionIds to be an explicit nil

### UnsetWorkItemVersionIds
`func (o *TestSuiteWorkItemsSearchModel) UnsetWorkItemVersionIds()`

UnsetWorkItemVersionIds ensures that no value is present for WorkItemVersionIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


