# CreateWorkItemApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | Unique identifier of the project | 
**SectionId** | Pointer to **NullableString** | Unique identifier of the section within a project | [optional] 
**Name** | **string** | Name of the work item | 
**Description** | Pointer to **NullableString** | Description of the work item | [optional] 
**EntityTypeName** | [**WorkItemEntityTypeApiModel**](WorkItemEntityTypeApiModel.md) | Type of entity associated with this work item | 
**Duration** | **int64** | Duration of the work item in milliseconds | 
**State** | [**WorkItemStateApiModel**](WorkItemStateApiModel.md) | State of the work item | 
**Priority** | [**WorkItemPriorityApiModel**](WorkItemPriorityApiModel.md) | Priority level of the work item | 
**Attributes** | **map[string]interface{}** | Set of custom attributes associated with the work item | 
**Tags** | [**[]TagModel**](TagModel.md) | Set of tags applied to the work item | 
**PreconditionSteps** | [**[]CreateStepApiModel**](CreateStepApiModel.md) | Set of precondition steps that need to be executed before starting the main steps | 
**Steps** | [**[]CreateStepApiModel**](CreateStepApiModel.md) | Main steps or actions defined for the work item | 
**PostconditionSteps** | [**[]CreateStepApiModel**](CreateStepApiModel.md) | Set of postcondition steps that are executed after completing the main steps | 
**Iterations** | Pointer to [**[]AssignIterationApiModel**](AssignIterationApiModel.md) | Associated iterations linked to the work item | [optional] 
**AutoTests** | Pointer to [**[]AutoTestIdModel**](AutoTestIdModel.md) | Automated tests associated with the work item | [optional] 
**Attachments** | Pointer to [**[]AssignAttachmentApiModel**](AssignAttachmentApiModel.md) | Files attached to the work item | [optional] 
**Links** | [**[]CreateLinkApiModel**](CreateLinkApiModel.md) | Set of links related to the work item | 

## Methods

### NewCreateWorkItemApiModel

`func NewCreateWorkItemApiModel(projectId string, name string, entityTypeName WorkItemEntityTypeApiModel, duration int64, state WorkItemStateApiModel, priority WorkItemPriorityApiModel, attributes map[string]interface{}, tags []TagModel, preconditionSteps []CreateStepApiModel, steps []CreateStepApiModel, postconditionSteps []CreateStepApiModel, links []CreateLinkApiModel, ) *CreateWorkItemApiModel`

NewCreateWorkItemApiModel instantiates a new CreateWorkItemApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkItemApiModelWithDefaults

`func NewCreateWorkItemApiModelWithDefaults() *CreateWorkItemApiModel`

NewCreateWorkItemApiModelWithDefaults instantiates a new CreateWorkItemApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### HasSectionId

`func (o *CreateWorkItemApiModel) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### SetSectionIdNil

`func (o *CreateWorkItemApiModel) SetSectionIdNil(b bool)`

 SetSectionIdNil sets the value for SectionId to be an explicit nil

### UnsetSectionId
`func (o *CreateWorkItemApiModel) UnsetSectionId()`

UnsetSectionId ensures that no value is present for SectionId, not even an explicit nil
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
### GetEntityTypeName

`func (o *CreateWorkItemApiModel) GetEntityTypeName() WorkItemEntityTypeApiModel`

GetEntityTypeName returns the EntityTypeName field if non-nil, zero value otherwise.

### GetEntityTypeNameOk

`func (o *CreateWorkItemApiModel) GetEntityTypeNameOk() (*WorkItemEntityTypeApiModel, bool)`

GetEntityTypeNameOk returns a tuple with the EntityTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypeName

`func (o *CreateWorkItemApiModel) SetEntityTypeName(v WorkItemEntityTypeApiModel)`

SetEntityTypeName sets EntityTypeName field to given value.


### GetDuration

`func (o *CreateWorkItemApiModel) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CreateWorkItemApiModel) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CreateWorkItemApiModel) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetState

`func (o *CreateWorkItemApiModel) GetState() WorkItemStateApiModel`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateWorkItemApiModel) GetStateOk() (*WorkItemStateApiModel, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateWorkItemApiModel) SetState(v WorkItemStateApiModel)`

SetState sets State field to given value.


### GetPriority

`func (o *CreateWorkItemApiModel) GetPriority() WorkItemPriorityApiModel`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateWorkItemApiModel) GetPriorityOk() (*WorkItemPriorityApiModel, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateWorkItemApiModel) SetPriority(v WorkItemPriorityApiModel)`

SetPriority sets Priority field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


