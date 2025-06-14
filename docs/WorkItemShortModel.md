# WorkItemShortModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Work Item internal unique identifier | 
**VersionId** | **string** | Work Item version identifier | 
**VersionNumber** | **int32** | Work Item version number | 
**Name** | **string** | Work Item name | 
**EntityTypeName** | **string** | Work Item type. Possible values: CheckLists, SharedSteps, TestCases | 
**ProjectId** | **string** | Project unique identifier | 
**SectionId** | **string** | Identifier of Section where Work Item is located | 
**SectionName** | **string** | Section name of Work Item | 
**IsAutomated** | **bool** | Boolean flag determining whether Work Item is automated | 
**GlobalId** | **int64** | Work Item global identifier | 
**Duration** | **int32** | Work Item duration | 
**MedianDuration** | Pointer to **NullableInt64** | Work Item median duration | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | Work Item attributes | [optional] 
**CreatedById** | **string** | Unique identifier of user who created Work Item | 
**ModifiedById** | Pointer to **NullableString** | Unique identifier of user who applied the latest modification of Work Item | [optional] 
**CreatedDate** | Pointer to **NullableTime** | Date and time of Work Item creation | [optional] 
**ModifiedDate** | Pointer to **NullableTime** | Date and time of the latest modification of Work Item | [optional] 
**State** | [**WorkItemStates**](WorkItemStates.md) | The current state of Work Item | 
**Priority** | [**WorkItemPriorityModel**](WorkItemPriorityModel.md) | Work Item priority level | 
**SourceType** | [**WorkItemSourceTypeModel**](WorkItemSourceTypeModel.md) | Work Item source type | 
**IsDeleted** | **bool** | Flag determining whether Work Item is deleted | 
**TagNames** | Pointer to **[]string** | Array of tag names of Work Item | [optional] 
**Iterations** | [**[]IterationModel**](IterationModel.md) | Set of iterations related to Work Item | 
**Links** | [**[]LinkShortModel**](LinkShortModel.md) | Set of links related to Work Item | 

## Methods

### NewWorkItemShortModel

`func NewWorkItemShortModel(id string, versionId string, versionNumber int32, name string, entityTypeName string, projectId string, sectionId string, sectionName string, isAutomated bool, globalId int64, duration int32, createdById string, state WorkItemStates, priority WorkItemPriorityModel, sourceType WorkItemSourceTypeModel, isDeleted bool, iterations []IterationModel, links []LinkShortModel, ) *WorkItemShortModel`

NewWorkItemShortModel instantiates a new WorkItemShortModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemShortModelWithDefaults

`func NewWorkItemShortModelWithDefaults() *WorkItemShortModel`

NewWorkItemShortModelWithDefaults instantiates a new WorkItemShortModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkItemShortModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkItemShortModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkItemShortModel) SetId(v string)`

SetId sets Id field to given value.


### GetVersionId

`func (o *WorkItemShortModel) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *WorkItemShortModel) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *WorkItemShortModel) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetVersionNumber

`func (o *WorkItemShortModel) GetVersionNumber() int32`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *WorkItemShortModel) GetVersionNumberOk() (*int32, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *WorkItemShortModel) SetVersionNumber(v int32)`

SetVersionNumber sets VersionNumber field to given value.


### GetName

`func (o *WorkItemShortModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkItemShortModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkItemShortModel) SetName(v string)`

SetName sets Name field to given value.


### GetEntityTypeName

`func (o *WorkItemShortModel) GetEntityTypeName() string`

GetEntityTypeName returns the EntityTypeName field if non-nil, zero value otherwise.

### GetEntityTypeNameOk

`func (o *WorkItemShortModel) GetEntityTypeNameOk() (*string, bool)`

GetEntityTypeNameOk returns a tuple with the EntityTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypeName

`func (o *WorkItemShortModel) SetEntityTypeName(v string)`

SetEntityTypeName sets EntityTypeName field to given value.


### GetProjectId

`func (o *WorkItemShortModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WorkItemShortModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WorkItemShortModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSectionId

`func (o *WorkItemShortModel) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *WorkItemShortModel) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *WorkItemShortModel) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.


### GetSectionName

`func (o *WorkItemShortModel) GetSectionName() string`

GetSectionName returns the SectionName field if non-nil, zero value otherwise.

### GetSectionNameOk

`func (o *WorkItemShortModel) GetSectionNameOk() (*string, bool)`

GetSectionNameOk returns a tuple with the SectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionName

`func (o *WorkItemShortModel) SetSectionName(v string)`

SetSectionName sets SectionName field to given value.


### GetIsAutomated

`func (o *WorkItemShortModel) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *WorkItemShortModel) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *WorkItemShortModel) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.


### GetGlobalId

`func (o *WorkItemShortModel) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *WorkItemShortModel) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *WorkItemShortModel) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.


### GetDuration

`func (o *WorkItemShortModel) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WorkItemShortModel) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WorkItemShortModel) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetMedianDuration

`func (o *WorkItemShortModel) GetMedianDuration() int64`

GetMedianDuration returns the MedianDuration field if non-nil, zero value otherwise.

### GetMedianDurationOk

`func (o *WorkItemShortModel) GetMedianDurationOk() (*int64, bool)`

GetMedianDurationOk returns a tuple with the MedianDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianDuration

`func (o *WorkItemShortModel) SetMedianDuration(v int64)`

SetMedianDuration sets MedianDuration field to given value.

### HasMedianDuration

`func (o *WorkItemShortModel) HasMedianDuration() bool`

HasMedianDuration returns a boolean if a field has been set.

### SetMedianDurationNil

`func (o *WorkItemShortModel) SetMedianDurationNil(b bool)`

 SetMedianDurationNil sets the value for MedianDuration to be an explicit nil

### UnsetMedianDuration
`func (o *WorkItemShortModel) UnsetMedianDuration()`

UnsetMedianDuration ensures that no value is present for MedianDuration, not even an explicit nil
### GetAttributes

`func (o *WorkItemShortModel) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkItemShortModel) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkItemShortModel) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WorkItemShortModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *WorkItemShortModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *WorkItemShortModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetCreatedById

`func (o *WorkItemShortModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WorkItemShortModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WorkItemShortModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedById

`func (o *WorkItemShortModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *WorkItemShortModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *WorkItemShortModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *WorkItemShortModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *WorkItemShortModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *WorkItemShortModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetCreatedDate

`func (o *WorkItemShortModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *WorkItemShortModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *WorkItemShortModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *WorkItemShortModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *WorkItemShortModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *WorkItemShortModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetModifiedDate

`func (o *WorkItemShortModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *WorkItemShortModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *WorkItemShortModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *WorkItemShortModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *WorkItemShortModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *WorkItemShortModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetState

`func (o *WorkItemShortModel) GetState() WorkItemStates`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkItemShortModel) GetStateOk() (*WorkItemStates, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkItemShortModel) SetState(v WorkItemStates)`

SetState sets State field to given value.


### GetPriority

`func (o *WorkItemShortModel) GetPriority() WorkItemPriorityModel`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WorkItemShortModel) GetPriorityOk() (*WorkItemPriorityModel, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WorkItemShortModel) SetPriority(v WorkItemPriorityModel)`

SetPriority sets Priority field to given value.


### GetSourceType

`func (o *WorkItemShortModel) GetSourceType() WorkItemSourceTypeModel`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *WorkItemShortModel) GetSourceTypeOk() (*WorkItemSourceTypeModel, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *WorkItemShortModel) SetSourceType(v WorkItemSourceTypeModel)`

SetSourceType sets SourceType field to given value.


### GetIsDeleted

`func (o *WorkItemShortModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *WorkItemShortModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *WorkItemShortModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetTagNames

`func (o *WorkItemShortModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *WorkItemShortModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *WorkItemShortModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *WorkItemShortModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### SetTagNamesNil

`func (o *WorkItemShortModel) SetTagNamesNil(b bool)`

 SetTagNamesNil sets the value for TagNames to be an explicit nil

### UnsetTagNames
`func (o *WorkItemShortModel) UnsetTagNames()`

UnsetTagNames ensures that no value is present for TagNames, not even an explicit nil
### GetIterations

`func (o *WorkItemShortModel) GetIterations() []IterationModel`

GetIterations returns the Iterations field if non-nil, zero value otherwise.

### GetIterationsOk

`func (o *WorkItemShortModel) GetIterationsOk() (*[]IterationModel, bool)`

GetIterationsOk returns a tuple with the Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterations

`func (o *WorkItemShortModel) SetIterations(v []IterationModel)`

SetIterations sets Iterations field to given value.


### GetLinks

`func (o *WorkItemShortModel) GetLinks() []LinkShortModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WorkItemShortModel) GetLinksOk() (*[]LinkShortModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WorkItemShortModel) SetLinks(v []LinkShortModel)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


