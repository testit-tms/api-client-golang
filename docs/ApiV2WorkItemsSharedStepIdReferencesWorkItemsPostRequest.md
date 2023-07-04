# ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of work item | [optional] 
**GlobalIds** | Pointer to **[]int64** | Collection of global (integer) identifiers | [optional] 
**SectionIds** | Pointer to **[]string** | Collection of section identifiers | [optional] 
**CreatedByIds** | Pointer to **[]string** | Collection of identifiers of users who created work item | [optional] 
**ModifiedByIds** | Pointer to **[]string** | Collection of identifiers of users who applied last modification to work item | [optional] 
**States** | Pointer to [**[]WorkItemStates**](WorkItemStates.md) | Collection of states of work item | [optional] 
**Priorities** | Pointer to [**[]WorkItemPriorityModel**](WorkItemPriorityModel.md) | Collection of priorities of work item | [optional] 
**EntityTypes** | Pointer to **[]string** | Collection of types of work item  &lt;br&gt;Allowed values: &#x60;TestCases&#x60;, &#x60;CheckLists&#x60;, &#x60;SharedSteps&#x60; | [optional] 
**CreatedDate** | Pointer to [**NullableSharedStepReferenceSectionsQueryFilterModelCreatedDate**](SharedStepReferenceSectionsQueryFilterModelCreatedDate.md) |  | [optional] 
**ModifiedDate** | Pointer to [**NullableSharedStepReferenceSectionsQueryFilterModelModifiedDate**](SharedStepReferenceSectionsQueryFilterModelModifiedDate.md) |  | [optional] 
**IsAutomated** | Pointer to **NullableBool** | Is result must consist of only manual/automated work items | [optional] 
**Tags** | Pointer to **[]string** | Collection of tags | [optional] 

## Methods

### NewApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest

`func NewApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest() *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest`

NewApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest instantiates a new ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequestWithDefaults

`func NewApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequestWithDefaults() *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest`

NewApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequestWithDefaults instantiates a new ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetGlobalIds

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetGlobalIds() []int64`

GetGlobalIds returns the GlobalIds field if non-nil, zero value otherwise.

### GetGlobalIdsOk

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetGlobalIdsOk() (*[]int64, bool)`

GetGlobalIdsOk returns a tuple with the GlobalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIds

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetGlobalIds(v []int64)`

SetGlobalIds sets GlobalIds field to given value.

### HasGlobalIds

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) HasGlobalIds() bool`

HasGlobalIds returns a boolean if a field has been set.

### SetGlobalIdsNil

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetGlobalIdsNil(b bool)`

 SetGlobalIdsNil sets the value for GlobalIds to be an explicit nil

### UnsetGlobalIds
`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) UnsetGlobalIds()`

UnsetGlobalIds ensures that no value is present for GlobalIds, not even an explicit nil
### GetSectionIds

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetSectionIds() []string`

GetSectionIds returns the SectionIds field if non-nil, zero value otherwise.

### GetSectionIdsOk

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetSectionIdsOk() (*[]string, bool)`

GetSectionIdsOk returns a tuple with the SectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionIds

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetSectionIds(v []string)`

SetSectionIds sets SectionIds field to given value.

### HasSectionIds

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) HasSectionIds() bool`

HasSectionIds returns a boolean if a field has been set.

### SetSectionIdsNil

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetSectionIdsNil(b bool)`

 SetSectionIdsNil sets the value for SectionIds to be an explicit nil

### UnsetSectionIds
`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) UnsetSectionIds()`

UnsetSectionIds ensures that no value is present for SectionIds, not even an explicit nil
### GetCreatedByIds

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetCreatedByIds() []string`

GetCreatedByIds returns the CreatedByIds field if non-nil, zero value otherwise.

### GetCreatedByIdsOk

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetCreatedByIdsOk() (*[]string, bool)`

GetCreatedByIdsOk returns a tuple with the CreatedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIds

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetCreatedByIds(v []string)`

SetCreatedByIds sets CreatedByIds field to given value.

### HasCreatedByIds

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) HasCreatedByIds() bool`

HasCreatedByIds returns a boolean if a field has been set.

### SetCreatedByIdsNil

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetCreatedByIdsNil(b bool)`

 SetCreatedByIdsNil sets the value for CreatedByIds to be an explicit nil

### UnsetCreatedByIds
`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) UnsetCreatedByIds()`

UnsetCreatedByIds ensures that no value is present for CreatedByIds, not even an explicit nil
### GetModifiedByIds

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetModifiedByIds() []string`

GetModifiedByIds returns the ModifiedByIds field if non-nil, zero value otherwise.

### GetModifiedByIdsOk

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetModifiedByIdsOk() (*[]string, bool)`

GetModifiedByIdsOk returns a tuple with the ModifiedByIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByIds

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetModifiedByIds(v []string)`

SetModifiedByIds sets ModifiedByIds field to given value.

### HasModifiedByIds

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) HasModifiedByIds() bool`

HasModifiedByIds returns a boolean if a field has been set.

### SetModifiedByIdsNil

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetModifiedByIdsNil(b bool)`

 SetModifiedByIdsNil sets the value for ModifiedByIds to be an explicit nil

### UnsetModifiedByIds
`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) UnsetModifiedByIds()`

UnsetModifiedByIds ensures that no value is present for ModifiedByIds, not even an explicit nil
### GetStates

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetStates() []WorkItemStates`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetStatesOk() (*[]WorkItemStates, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetStates(v []WorkItemStates)`

SetStates sets States field to given value.

### HasStates

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStatesNil

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetStatesNil(b bool)`

 SetStatesNil sets the value for States to be an explicit nil

### UnsetStates
`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) UnsetStates()`

UnsetStates ensures that no value is present for States, not even an explicit nil
### GetPriorities

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetPriorities() []WorkItemPriorityModel`

GetPriorities returns the Priorities field if non-nil, zero value otherwise.

### GetPrioritiesOk

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetPrioritiesOk() (*[]WorkItemPriorityModel, bool)`

GetPrioritiesOk returns a tuple with the Priorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorities

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetPriorities(v []WorkItemPriorityModel)`

SetPriorities sets Priorities field to given value.

### HasPriorities

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) HasPriorities() bool`

HasPriorities returns a boolean if a field has been set.

### SetPrioritiesNil

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetPrioritiesNil(b bool)`

 SetPrioritiesNil sets the value for Priorities to be an explicit nil

### UnsetPriorities
`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) UnsetPriorities()`

UnsetPriorities ensures that no value is present for Priorities, not even an explicit nil
### GetEntityTypes

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetEntityTypes() []string`

GetEntityTypes returns the EntityTypes field if non-nil, zero value otherwise.

### GetEntityTypesOk

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetEntityTypesOk() (*[]string, bool)`

GetEntityTypesOk returns a tuple with the EntityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypes

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetEntityTypes(v []string)`

SetEntityTypes sets EntityTypes field to given value.

### HasEntityTypes

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) HasEntityTypes() bool`

HasEntityTypes returns a boolean if a field has been set.

### SetEntityTypesNil

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetEntityTypesNil(b bool)`

 SetEntityTypesNil sets the value for EntityTypes to be an explicit nil

### UnsetEntityTypes
`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) UnsetEntityTypes()`

UnsetEntityTypes ensures that no value is present for EntityTypes, not even an explicit nil
### GetCreatedDate

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetCreatedDate() SharedStepReferenceSectionsQueryFilterModelCreatedDate`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetCreatedDateOk() (*SharedStepReferenceSectionsQueryFilterModelCreatedDate, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetCreatedDate(v SharedStepReferenceSectionsQueryFilterModelCreatedDate)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetModifiedDate

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetModifiedDate() SharedStepReferenceSectionsQueryFilterModelModifiedDate`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetModifiedDateOk() (*SharedStepReferenceSectionsQueryFilterModelModifiedDate, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetModifiedDate(v SharedStepReferenceSectionsQueryFilterModelModifiedDate)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetIsAutomated

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### SetIsAutomatedNil

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetIsAutomatedNil(b bool)`

 SetIsAutomatedNil sets the value for IsAutomated to be an explicit nil

### UnsetIsAutomated
`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) UnsetIsAutomated()`

UnsetIsAutomated ensures that no value is present for IsAutomated, not even an explicit nil
### GetTags

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ApiV2WorkItemsSharedStepIdReferencesWorkItemsPostRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


