# CreateWorkItemApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityTypeName** | [**WorkItemEntityTypes**](WorkItemEntityTypes.md) |  | 
**Description** | Pointer to **NullableString** | Workitem description | [optional] 
**State** | [**WorkItemStates**](WorkItemStates.md) |  | 
**Priority** | [**WorkItemPriorityModel**](WorkItemPriorityModel.md) |  | 
**Steps** | [**[]CreateStepApiModel**](CreateStepApiModel.md) | Collection of workitem steps | 
**PreconditionSteps** | [**[]CreateStepApiModel**](CreateStepApiModel.md) | Collection of workitem precondition steps | 
**PostconditionSteps** | [**[]CreateStepApiModel**](CreateStepApiModel.md) | Collection of workitem postcondition steps | 
**Duration** | **int32** | WorkItem duration in milliseconds, must be 0 for shared steps and greater than 0 for the other types of work items | 
**Attributes** | **map[string]interface{}** | Key value pair of custom workitem attributes | 
**Tags** | [**[]TagModel**](TagModel.md) | Collection of workitem tags | 
**Attachments** | Pointer to [**[]AssignAttachmentApiModel**](AssignAttachmentApiModel.md) | Collection of workitem attachments | [optional] 
**Iterations** | Pointer to [**[]AssignIterationApiModel**](AssignIterationApiModel.md) | Collection of parameter sets | [optional] 
**Links** | [**[]CreateLinkApiModel**](CreateLinkApiModel.md) | Collection of workitem links | 
**Name** | **string** | Workitem name | 
**ProjectId** | **string** | Project unique identifier - used to link workitem with project | 
**SectionId** | **string** | Internal identifier of section where workitem is located | 
**AutoTests** | Pointer to [**[]AutoTestIdModel**](AutoTestIdModel.md) | Collection of autotest internal ids | [optional] 

## Methods

### NewCreateWorkItemApiModel

`func NewCreateWorkItemApiModel(entityTypeName WorkItemEntityTypes, state WorkItemStates, priority WorkItemPriorityModel, steps []CreateStepApiModel, preconditionSteps []CreateStepApiModel, postconditionSteps []CreateStepApiModel, duration int32, attributes map[string]interface{}, tags []TagModel, links []CreateLinkApiModel, name string, projectId string, sectionId string, ) *CreateWorkItemApiModel`

NewCreateWorkItemApiModel instantiates a new CreateWorkItemApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkItemApiModelWithDefaults

`func NewCreateWorkItemApiModelWithDefaults() *CreateWorkItemApiModel`

NewCreateWorkItemApiModelWithDefaults instantiates a new CreateWorkItemApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityTypeName

`func (o *CreateWorkItemApiModel) GetEntityTypeName() WorkItemEntityTypes`

GetEntityTypeName returns the EntityTypeName field if non-nil, zero value otherwise.

### GetEntityTypeNameOk

`func (o *CreateWorkItemApiModel) GetEntityTypeNameOk() (*WorkItemEntityTypes, bool)`

GetEntityTypeNameOk returns a tuple with the EntityTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypeName

`func (o *CreateWorkItemApiModel) SetEntityTypeName(v WorkItemEntityTypes)`

SetEntityTypeName sets EntityTypeName field to given value.


### GetDescription

`func (o *CreateWorkItemApiModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateWorkItemApiModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateWorkItemApiModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateWorkItemApiModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateWorkItemApiModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateWorkItemApiModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetState

`func (o *CreateWorkItemApiModel) GetState() WorkItemStates`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateWorkItemApiModel) GetStateOk() (*WorkItemStates, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateWorkItemApiModel) SetState(v WorkItemStates)`

SetState sets State field to given value.


### GetPriority

`func (o *CreateWorkItemApiModel) GetPriority() WorkItemPriorityModel`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateWorkItemApiModel) GetPriorityOk() (*WorkItemPriorityModel, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateWorkItemApiModel) SetPriority(v WorkItemPriorityModel)`

SetPriority sets Priority field to given value.


### GetSteps

`func (o *CreateWorkItemApiModel) GetSteps() []CreateStepApiModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *CreateWorkItemApiModel) GetStepsOk() (*[]CreateStepApiModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *CreateWorkItemApiModel) SetSteps(v []CreateStepApiModel)`

SetSteps sets Steps field to given value.


### GetPreconditionSteps

`func (o *CreateWorkItemApiModel) GetPreconditionSteps() []CreateStepApiModel`

GetPreconditionSteps returns the PreconditionSteps field if non-nil, zero value otherwise.

### GetPreconditionStepsOk

`func (o *CreateWorkItemApiModel) GetPreconditionStepsOk() (*[]CreateStepApiModel, bool)`

GetPreconditionStepsOk returns a tuple with the PreconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditionSteps

`func (o *CreateWorkItemApiModel) SetPreconditionSteps(v []CreateStepApiModel)`

SetPreconditionSteps sets PreconditionSteps field to given value.


### GetPostconditionSteps

`func (o *CreateWorkItemApiModel) GetPostconditionSteps() []CreateStepApiModel`

GetPostconditionSteps returns the PostconditionSteps field if non-nil, zero value otherwise.

### GetPostconditionStepsOk

`func (o *CreateWorkItemApiModel) GetPostconditionStepsOk() (*[]CreateStepApiModel, bool)`

GetPostconditionStepsOk returns a tuple with the PostconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditionSteps

`func (o *CreateWorkItemApiModel) SetPostconditionSteps(v []CreateStepApiModel)`

SetPostconditionSteps sets PostconditionSteps field to given value.


### GetDuration

`func (o *CreateWorkItemApiModel) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CreateWorkItemApiModel) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CreateWorkItemApiModel) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetAttributes

`func (o *CreateWorkItemApiModel) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CreateWorkItemApiModel) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CreateWorkItemApiModel) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetTags

`func (o *CreateWorkItemApiModel) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateWorkItemApiModel) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateWorkItemApiModel) SetTags(v []TagModel)`

SetTags sets Tags field to given value.


### GetAttachments

`func (o *CreateWorkItemApiModel) GetAttachments() []AssignAttachmentApiModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *CreateWorkItemApiModel) GetAttachmentsOk() (*[]AssignAttachmentApiModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *CreateWorkItemApiModel) SetAttachments(v []AssignAttachmentApiModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *CreateWorkItemApiModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *CreateWorkItemApiModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *CreateWorkItemApiModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetIterations

`func (o *CreateWorkItemApiModel) GetIterations() []AssignIterationApiModel`

GetIterations returns the Iterations field if non-nil, zero value otherwise.

### GetIterationsOk

`func (o *CreateWorkItemApiModel) GetIterationsOk() (*[]AssignIterationApiModel, bool)`

GetIterationsOk returns a tuple with the Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterations

`func (o *CreateWorkItemApiModel) SetIterations(v []AssignIterationApiModel)`

SetIterations sets Iterations field to given value.

### HasIterations

`func (o *CreateWorkItemApiModel) HasIterations() bool`

HasIterations returns a boolean if a field has been set.

### SetIterationsNil

`func (o *CreateWorkItemApiModel) SetIterationsNil(b bool)`

 SetIterationsNil sets the value for Iterations to be an explicit nil

### UnsetIterations
`func (o *CreateWorkItemApiModel) UnsetIterations()`

UnsetIterations ensures that no value is present for Iterations, not even an explicit nil
### GetLinks

`func (o *CreateWorkItemApiModel) GetLinks() []CreateLinkApiModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateWorkItemApiModel) GetLinksOk() (*[]CreateLinkApiModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateWorkItemApiModel) SetLinks(v []CreateLinkApiModel)`

SetLinks sets Links field to given value.


### GetName

`func (o *CreateWorkItemApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWorkItemApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWorkItemApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *CreateWorkItemApiModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateWorkItemApiModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateWorkItemApiModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSectionId

`func (o *CreateWorkItemApiModel) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *CreateWorkItemApiModel) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *CreateWorkItemApiModel) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.


### GetAutoTests

`func (o *CreateWorkItemApiModel) GetAutoTests() []AutoTestIdModel`

GetAutoTests returns the AutoTests field if non-nil, zero value otherwise.

### GetAutoTestsOk

`func (o *CreateWorkItemApiModel) GetAutoTestsOk() (*[]AutoTestIdModel, bool)`

GetAutoTestsOk returns a tuple with the AutoTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTests

`func (o *CreateWorkItemApiModel) SetAutoTests(v []AutoTestIdModel)`

SetAutoTests sets AutoTests field to given value.

### HasAutoTests

`func (o *CreateWorkItemApiModel) HasAutoTests() bool`

HasAutoTests returns a boolean if a field has been set.

### SetAutoTestsNil

`func (o *CreateWorkItemApiModel) SetAutoTestsNil(b bool)`

 SetAutoTestsNil sets the value for AutoTests to be an explicit nil

### UnsetAutoTests
`func (o *CreateWorkItemApiModel) UnsetAutoTests()`

UnsetAutoTests ensures that no value is present for AutoTests, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


