# WorkItemApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the work item | 
**GlobalId** | **int64** | Global identifier of the work item | 
**VersionId** | **string** | Version identifier of the work item | 
**VersionNumber** | **int32** | Version number of the work item | 
**ProjectId** | **string** | Unique identifier of the project | 
**SectionId** | **string** | Unique identifier of the section within a project | 
**Name** | **string** | Name of the work item | 
**Description** | Pointer to **NullableString** | Description of the work item | [optional] 
**SourceType** | [**WorkItemSourceTypeApiModel**](WorkItemSourceTypeApiModel.md) | Source type of the work item | 
**EntityTypeName** | [**WorkItemEntityTypeApiModel**](WorkItemEntityTypeApiModel.md) | Type of entity associated with this work item | 
**Duration** | **int32** | Duration of the work item in milliseconds | 
**MedianDuration** | **int64** | Median duration of the work item in milliseconds | 
**State** | [**WorkItemStateApiModel**](WorkItemStateApiModel.md) | State of the work item | 
**Priority** | [**WorkItemPriorityApiModel**](WorkItemPriorityApiModel.md) | Priority level of the work item | 
**IsAutomated** | **bool** |  | 
**Attributes** | **map[string]interface{}** | Set of custom attributes associated with the work item | 
**Tags** | [**[]TagModel**](TagModel.md) | Set of tags applied to the work item | 
**SectionPreconditionSteps** | [**[]StepModel**](StepModel.md) | Set of section precondition steps that need to be executed before starting the work item steps | 
**SectionPostconditionSteps** | [**[]StepModel**](StepModel.md) | Set of section postcondition steps that need to be executed after completing the work item steps | 
**PreconditionSteps** | [**[]StepModel**](StepModel.md) | Set of precondition steps that need to be executed before starting the main steps | 
**Steps** | [**[]StepModel**](StepModel.md) | Main steps or actions defined for the work item | 
**PostconditionSteps** | [**[]StepModel**](StepModel.md) | Set of postcondition steps that are executed after completing the main steps | 
**Iterations** | [**[]IterationModel**](IterationModel.md) | Associated iterations linked to the work item | 
**AutoTests** | [**[]AutoTestModel**](AutoTestModel.md) | Automated tests associated with the work item | 
**Attachments** | [**[]AttachmentModel**](AttachmentModel.md) | Files attached to the work item | 
**Links** | [**[]LinkModel**](LinkModel.md) | Set of links related to the work item | 
**CreatedDate** | **time.Time** | Creation date of the work item | 
**CreatedById** | **string** | Unique identifier of the work item creator | 
**ModifiedDate** | Pointer to **NullableTime** | Modification date of the work item | [optional] 
**ModifiedById** | Pointer to **NullableString** | Unique identifier of the work item modifier | [optional] 
**IsDeleted** | **bool** | Indicates whether the work item is marked as deleted | 

## Methods

### NewWorkItemApiResult

`func NewWorkItemApiResult(id string, globalId int64, versionId string, versionNumber int32, projectId string, sectionId string, name string, sourceType WorkItemSourceTypeApiModel, entityTypeName WorkItemEntityTypeApiModel, duration int32, medianDuration int64, state WorkItemStateApiModel, priority WorkItemPriorityApiModel, isAutomated bool, attributes map[string]interface{}, tags []TagModel, sectionPreconditionSteps []StepModel, sectionPostconditionSteps []StepModel, preconditionSteps []StepModel, steps []StepModel, postconditionSteps []StepModel, iterations []IterationModel, autoTests []AutoTestModel, attachments []AttachmentModel, links []LinkModel, createdDate time.Time, createdById string, isDeleted bool, ) *WorkItemApiResult`

NewWorkItemApiResult instantiates a new WorkItemApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemApiResultWithDefaults

`func NewWorkItemApiResultWithDefaults() *WorkItemApiResult`

NewWorkItemApiResultWithDefaults instantiates a new WorkItemApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkItemApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkItemApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkItemApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetGlobalId

`func (o *WorkItemApiResult) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *WorkItemApiResult) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *WorkItemApiResult) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.


### GetVersionId

`func (o *WorkItemApiResult) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *WorkItemApiResult) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *WorkItemApiResult) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetVersionNumber

`func (o *WorkItemApiResult) GetVersionNumber() int32`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *WorkItemApiResult) GetVersionNumberOk() (*int32, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *WorkItemApiResult) SetVersionNumber(v int32)`

SetVersionNumber sets VersionNumber field to given value.


### GetProjectId

`func (o *WorkItemApiResult) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WorkItemApiResult) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WorkItemApiResult) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSectionId

`func (o *WorkItemApiResult) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *WorkItemApiResult) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *WorkItemApiResult) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.


### GetName

`func (o *WorkItemApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkItemApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkItemApiResult) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WorkItemApiResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkItemApiResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkItemApiResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkItemApiResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkItemApiResult) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkItemApiResult) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSourceType

`func (o *WorkItemApiResult) GetSourceType() WorkItemSourceTypeApiModel`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *WorkItemApiResult) GetSourceTypeOk() (*WorkItemSourceTypeApiModel, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *WorkItemApiResult) SetSourceType(v WorkItemSourceTypeApiModel)`

SetSourceType sets SourceType field to given value.


### GetEntityTypeName

`func (o *WorkItemApiResult) GetEntityTypeName() WorkItemEntityTypeApiModel`

GetEntityTypeName returns the EntityTypeName field if non-nil, zero value otherwise.

### GetEntityTypeNameOk

`func (o *WorkItemApiResult) GetEntityTypeNameOk() (*WorkItemEntityTypeApiModel, bool)`

GetEntityTypeNameOk returns a tuple with the EntityTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypeName

`func (o *WorkItemApiResult) SetEntityTypeName(v WorkItemEntityTypeApiModel)`

SetEntityTypeName sets EntityTypeName field to given value.


### GetDuration

`func (o *WorkItemApiResult) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WorkItemApiResult) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WorkItemApiResult) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetMedianDuration

`func (o *WorkItemApiResult) GetMedianDuration() int64`

GetMedianDuration returns the MedianDuration field if non-nil, zero value otherwise.

### GetMedianDurationOk

`func (o *WorkItemApiResult) GetMedianDurationOk() (*int64, bool)`

GetMedianDurationOk returns a tuple with the MedianDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianDuration

`func (o *WorkItemApiResult) SetMedianDuration(v int64)`

SetMedianDuration sets MedianDuration field to given value.


### GetState

`func (o *WorkItemApiResult) GetState() WorkItemStateApiModel`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkItemApiResult) GetStateOk() (*WorkItemStateApiModel, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkItemApiResult) SetState(v WorkItemStateApiModel)`

SetState sets State field to given value.


### GetPriority

`func (o *WorkItemApiResult) GetPriority() WorkItemPriorityApiModel`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WorkItemApiResult) GetPriorityOk() (*WorkItemPriorityApiModel, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WorkItemApiResult) SetPriority(v WorkItemPriorityApiModel)`

SetPriority sets Priority field to given value.


### GetIsAutomated

`func (o *WorkItemApiResult) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *WorkItemApiResult) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *WorkItemApiResult) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.


### GetAttributes

`func (o *WorkItemApiResult) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkItemApiResult) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkItemApiResult) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetTags

`func (o *WorkItemApiResult) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkItemApiResult) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkItemApiResult) SetTags(v []TagModel)`

SetTags sets Tags field to given value.


### GetSectionPreconditionSteps

`func (o *WorkItemApiResult) GetSectionPreconditionSteps() []StepModel`

GetSectionPreconditionSteps returns the SectionPreconditionSteps field if non-nil, zero value otherwise.

### GetSectionPreconditionStepsOk

`func (o *WorkItemApiResult) GetSectionPreconditionStepsOk() (*[]StepModel, bool)`

GetSectionPreconditionStepsOk returns a tuple with the SectionPreconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionPreconditionSteps

`func (o *WorkItemApiResult) SetSectionPreconditionSteps(v []StepModel)`

SetSectionPreconditionSteps sets SectionPreconditionSteps field to given value.


### GetSectionPostconditionSteps

`func (o *WorkItemApiResult) GetSectionPostconditionSteps() []StepModel`

GetSectionPostconditionSteps returns the SectionPostconditionSteps field if non-nil, zero value otherwise.

### GetSectionPostconditionStepsOk

`func (o *WorkItemApiResult) GetSectionPostconditionStepsOk() (*[]StepModel, bool)`

GetSectionPostconditionStepsOk returns a tuple with the SectionPostconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionPostconditionSteps

`func (o *WorkItemApiResult) SetSectionPostconditionSteps(v []StepModel)`

SetSectionPostconditionSteps sets SectionPostconditionSteps field to given value.


### GetPreconditionSteps

`func (o *WorkItemApiResult) GetPreconditionSteps() []StepModel`

GetPreconditionSteps returns the PreconditionSteps field if non-nil, zero value otherwise.

### GetPreconditionStepsOk

`func (o *WorkItemApiResult) GetPreconditionStepsOk() (*[]StepModel, bool)`

GetPreconditionStepsOk returns a tuple with the PreconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditionSteps

`func (o *WorkItemApiResult) SetPreconditionSteps(v []StepModel)`

SetPreconditionSteps sets PreconditionSteps field to given value.


### GetSteps

`func (o *WorkItemApiResult) GetSteps() []StepModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *WorkItemApiResult) GetStepsOk() (*[]StepModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *WorkItemApiResult) SetSteps(v []StepModel)`

SetSteps sets Steps field to given value.


### GetPostconditionSteps

`func (o *WorkItemApiResult) GetPostconditionSteps() []StepModel`

GetPostconditionSteps returns the PostconditionSteps field if non-nil, zero value otherwise.

### GetPostconditionStepsOk

`func (o *WorkItemApiResult) GetPostconditionStepsOk() (*[]StepModel, bool)`

GetPostconditionStepsOk returns a tuple with the PostconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditionSteps

`func (o *WorkItemApiResult) SetPostconditionSteps(v []StepModel)`

SetPostconditionSteps sets PostconditionSteps field to given value.


### GetIterations

`func (o *WorkItemApiResult) GetIterations() []IterationModel`

GetIterations returns the Iterations field if non-nil, zero value otherwise.

### GetIterationsOk

`func (o *WorkItemApiResult) GetIterationsOk() (*[]IterationModel, bool)`

GetIterationsOk returns a tuple with the Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterations

`func (o *WorkItemApiResult) SetIterations(v []IterationModel)`

SetIterations sets Iterations field to given value.


### GetAutoTests

`func (o *WorkItemApiResult) GetAutoTests() []AutoTestModel`

GetAutoTests returns the AutoTests field if non-nil, zero value otherwise.

### GetAutoTestsOk

`func (o *WorkItemApiResult) GetAutoTestsOk() (*[]AutoTestModel, bool)`

GetAutoTestsOk returns a tuple with the AutoTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTests

`func (o *WorkItemApiResult) SetAutoTests(v []AutoTestModel)`

SetAutoTests sets AutoTests field to given value.


### GetAttachments

`func (o *WorkItemApiResult) GetAttachments() []AttachmentModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *WorkItemApiResult) GetAttachmentsOk() (*[]AttachmentModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *WorkItemApiResult) SetAttachments(v []AttachmentModel)`

SetAttachments sets Attachments field to given value.


### GetLinks

`func (o *WorkItemApiResult) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WorkItemApiResult) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WorkItemApiResult) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.


### GetCreatedDate

`func (o *WorkItemApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *WorkItemApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *WorkItemApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetCreatedById

`func (o *WorkItemApiResult) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WorkItemApiResult) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WorkItemApiResult) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedDate

`func (o *WorkItemApiResult) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *WorkItemApiResult) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *WorkItemApiResult) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *WorkItemApiResult) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *WorkItemApiResult) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *WorkItemApiResult) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedById

`func (o *WorkItemApiResult) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *WorkItemApiResult) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *WorkItemApiResult) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *WorkItemApiResult) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *WorkItemApiResult) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *WorkItemApiResult) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetIsDeleted

`func (o *WorkItemApiResult) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *WorkItemApiResult) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *WorkItemApiResult) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


