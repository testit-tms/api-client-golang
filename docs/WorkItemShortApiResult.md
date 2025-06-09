# WorkItemShortApiResult

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
**SourceType** | [**WorkItemSourceTypeModel**](WorkItemSourceTypeModel.md) | Work Item priority level | 
**IsDeleted** | **bool** | Flag determining whether Work Item is deleted | 
**TagNames** | Pointer to **[]string** | Array of tag names of Work Item | [optional] 
**Iterations** | [**[]IterationApiResult**](IterationApiResult.md) | Set of iterations related to Work Item | 
**Links** | [**[]LinkShortApiResult**](LinkShortApiResult.md) | Set of links related to Work Item | 

## Methods

### NewWorkItemShortApiResult

`func NewWorkItemShortApiResult(id string, versionId string, versionNumber int32, name string, entityTypeName string, projectId string, sectionId string, sectionName string, isAutomated bool, globalId int64, duration int32, createdById string, state WorkItemStates, priority WorkItemPriorityModel, sourceType WorkItemSourceTypeModel, isDeleted bool, iterations []IterationApiResult, links []LinkShortApiResult, ) *WorkItemShortApiResult`

NewWorkItemShortApiResult instantiates a new WorkItemShortApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemShortApiResultWithDefaults

`func NewWorkItemShortApiResultWithDefaults() *WorkItemShortApiResult`

NewWorkItemShortApiResultWithDefaults instantiates a new WorkItemShortApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkItemShortApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkItemShortApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkItemShortApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetVersionId

`func (o *WorkItemShortApiResult) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *WorkItemShortApiResult) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *WorkItemShortApiResult) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetVersionNumber

`func (o *WorkItemShortApiResult) GetVersionNumber() int32`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *WorkItemShortApiResult) GetVersionNumberOk() (*int32, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *WorkItemShortApiResult) SetVersionNumber(v int32)`

SetVersionNumber sets VersionNumber field to given value.


### GetName

`func (o *WorkItemShortApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkItemShortApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkItemShortApiResult) SetName(v string)`

SetName sets Name field to given value.


### GetEntityTypeName

`func (o *WorkItemShortApiResult) GetEntityTypeName() string`

GetEntityTypeName returns the EntityTypeName field if non-nil, zero value otherwise.

### GetEntityTypeNameOk

`func (o *WorkItemShortApiResult) GetEntityTypeNameOk() (*string, bool)`

GetEntityTypeNameOk returns a tuple with the EntityTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypeName

`func (o *WorkItemShortApiResult) SetEntityTypeName(v string)`

SetEntityTypeName sets EntityTypeName field to given value.


### GetProjectId

`func (o *WorkItemShortApiResult) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WorkItemShortApiResult) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WorkItemShortApiResult) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSectionId

`func (o *WorkItemShortApiResult) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *WorkItemShortApiResult) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *WorkItemShortApiResult) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.


### GetSectionName

`func (o *WorkItemShortApiResult) GetSectionName() string`

GetSectionName returns the SectionName field if non-nil, zero value otherwise.

### GetSectionNameOk

`func (o *WorkItemShortApiResult) GetSectionNameOk() (*string, bool)`

GetSectionNameOk returns a tuple with the SectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionName

`func (o *WorkItemShortApiResult) SetSectionName(v string)`

SetSectionName sets SectionName field to given value.


### GetIsAutomated

`func (o *WorkItemShortApiResult) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *WorkItemShortApiResult) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *WorkItemShortApiResult) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.


### GetGlobalId

`func (o *WorkItemShortApiResult) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *WorkItemShortApiResult) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *WorkItemShortApiResult) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.


### GetDuration

`func (o *WorkItemShortApiResult) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WorkItemShortApiResult) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WorkItemShortApiResult) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetMedianDuration

`func (o *WorkItemShortApiResult) GetMedianDuration() int64`

GetMedianDuration returns the MedianDuration field if non-nil, zero value otherwise.

### GetMedianDurationOk

`func (o *WorkItemShortApiResult) GetMedianDurationOk() (*int64, bool)`

GetMedianDurationOk returns a tuple with the MedianDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianDuration

`func (o *WorkItemShortApiResult) SetMedianDuration(v int64)`

SetMedianDuration sets MedianDuration field to given value.

### HasMedianDuration

`func (o *WorkItemShortApiResult) HasMedianDuration() bool`

HasMedianDuration returns a boolean if a field has been set.

### SetMedianDurationNil

`func (o *WorkItemShortApiResult) SetMedianDurationNil(b bool)`

 SetMedianDurationNil sets the value for MedianDuration to be an explicit nil

### UnsetMedianDuration
`func (o *WorkItemShortApiResult) UnsetMedianDuration()`

UnsetMedianDuration ensures that no value is present for MedianDuration, not even an explicit nil
### GetAttributes

`func (o *WorkItemShortApiResult) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkItemShortApiResult) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkItemShortApiResult) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WorkItemShortApiResult) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *WorkItemShortApiResult) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *WorkItemShortApiResult) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetCreatedById

`func (o *WorkItemShortApiResult) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WorkItemShortApiResult) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WorkItemShortApiResult) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedById

`func (o *WorkItemShortApiResult) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *WorkItemShortApiResult) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *WorkItemShortApiResult) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *WorkItemShortApiResult) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *WorkItemShortApiResult) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *WorkItemShortApiResult) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetCreatedDate

`func (o *WorkItemShortApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *WorkItemShortApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *WorkItemShortApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *WorkItemShortApiResult) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *WorkItemShortApiResult) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *WorkItemShortApiResult) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetModifiedDate

`func (o *WorkItemShortApiResult) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *WorkItemShortApiResult) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *WorkItemShortApiResult) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *WorkItemShortApiResult) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *WorkItemShortApiResult) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *WorkItemShortApiResult) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetState

`func (o *WorkItemShortApiResult) GetState() WorkItemStates`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkItemShortApiResult) GetStateOk() (*WorkItemStates, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkItemShortApiResult) SetState(v WorkItemStates)`

SetState sets State field to given value.


### GetPriority

`func (o *WorkItemShortApiResult) GetPriority() WorkItemPriorityModel`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WorkItemShortApiResult) GetPriorityOk() (*WorkItemPriorityModel, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WorkItemShortApiResult) SetPriority(v WorkItemPriorityModel)`

SetPriority sets Priority field to given value.


### GetSourceType

`func (o *WorkItemShortApiResult) GetSourceType() WorkItemSourceTypeModel`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *WorkItemShortApiResult) GetSourceTypeOk() (*WorkItemSourceTypeModel, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *WorkItemShortApiResult) SetSourceType(v WorkItemSourceTypeModel)`

SetSourceType sets SourceType field to given value.


### GetIsDeleted

`func (o *WorkItemShortApiResult) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *WorkItemShortApiResult) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *WorkItemShortApiResult) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetTagNames

`func (o *WorkItemShortApiResult) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *WorkItemShortApiResult) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *WorkItemShortApiResult) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *WorkItemShortApiResult) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### SetTagNamesNil

`func (o *WorkItemShortApiResult) SetTagNamesNil(b bool)`

 SetTagNamesNil sets the value for TagNames to be an explicit nil

### UnsetTagNames
`func (o *WorkItemShortApiResult) UnsetTagNames()`

UnsetTagNames ensures that no value is present for TagNames, not even an explicit nil
### GetIterations

`func (o *WorkItemShortApiResult) GetIterations() []IterationApiResult`

GetIterations returns the Iterations field if non-nil, zero value otherwise.

### GetIterationsOk

`func (o *WorkItemShortApiResult) GetIterationsOk() (*[]IterationApiResult, bool)`

GetIterationsOk returns a tuple with the Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterations

`func (o *WorkItemShortApiResult) SetIterations(v []IterationApiResult)`

SetIterations sets Iterations field to given value.


### GetLinks

`func (o *WorkItemShortApiResult) GetLinks() []LinkShortApiResult`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WorkItemShortApiResult) GetLinksOk() (*[]LinkShortApiResult, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WorkItemShortApiResult) SetLinks(v []LinkShortApiResult)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


