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
**State** | [**WorkItemStateApiModel**](WorkItemStateApiModel.md) | Current state of the work item | 
**Priority** | [**WorkItemPriorityApiModel**](WorkItemPriorityApiModel.md) | Priority level assigned to the work item | 
**Attributes** | Pointer to **map[string]interface{}** | Set of custom attributes associated with the work item | [optional] 
**Tags** | Pointer to [**[]TagModel**](TagModel.md) | Set of tags applied to the work item | [optional] 
**PreconditionSteps** | Pointer to [**[]CreateStepApiModel**](CreateStepApiModel.md) | Set of precondition steps that must be executed before the main steps | [optional] 
**Steps** | Pointer to [**[]CreateStepApiModel**](CreateStepApiModel.md) | Set of main steps or actions defined for the work item | [optional] 
**PostconditionSteps** | Pointer to [**[]CreateStepApiModel**](CreateStepApiModel.md) | Set of postcondition steps that are executed after completing the main steps | [optional] 
**Iterations** | Pointer to [**[]AssignIterationApiModel**](AssignIterationApiModel.md) | Set of iterations associated with the work item | [optional] 
**AutoTests** | Pointer to [**[]AutoTestIdModel**](AutoTestIdModel.md) | Set of automated tests linked to the work item | [optional] 
**Attachments** | Pointer to [**[]AssignAttachmentApiModel**](AssignAttachmentApiModel.md) | Set of files attached to the work item | [optional] 
**Links** | Pointer to [**[]CreateLinkApiModel**](CreateLinkApiModel.md) | Set of links related to the work item | [optional] 
**Parameters** | Pointer to [**[]WorkItemParameterKeyApiModel**](WorkItemParameterKeyApiModel.md) | Set of parameter keys associated with the work item | [optional] 

## Methods

### NewCreateWorkItemApiModel

`func NewCreateWorkItemApiModel(projectId string, name string, entityTypeName WorkItemEntityTypeApiModel, duration int64, state WorkItemStateApiModel, priority WorkItemPriorityApiModel, ) *CreateWorkItemApiModel`

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

### HasAttributes

`func (o *CreateWorkItemApiModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *CreateWorkItemApiModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *CreateWorkItemApiModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
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

### HasTags

`func (o *CreateWorkItemApiModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateWorkItemApiModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateWorkItemApiModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
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

### HasPreconditionSteps

`func (o *CreateWorkItemApiModel) HasPreconditionSteps() bool`

HasPreconditionSteps returns a boolean if a field has been set.

### SetPreconditionStepsNil

`func (o *CreateWorkItemApiModel) SetPreconditionStepsNil(b bool)`

 SetPreconditionStepsNil sets the value for PreconditionSteps to be an explicit nil

### UnsetPreconditionSteps
`func (o *CreateWorkItemApiModel) UnsetPreconditionSteps()`

UnsetPreconditionSteps ensures that no value is present for PreconditionSteps, not even an explicit nil
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

### HasSteps

`func (o *CreateWorkItemApiModel) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *CreateWorkItemApiModel) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *CreateWorkItemApiModel) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil
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

### HasPostconditionSteps

`func (o *CreateWorkItemApiModel) HasPostconditionSteps() bool`

HasPostconditionSteps returns a boolean if a field has been set.

### SetPostconditionStepsNil

`func (o *CreateWorkItemApiModel) SetPostconditionStepsNil(b bool)`

 SetPostconditionStepsNil sets the value for PostconditionSteps to be an explicit nil

### UnsetPostconditionSteps
`func (o *CreateWorkItemApiModel) UnsetPostconditionSteps()`

UnsetPostconditionSteps ensures that no value is present for PostconditionSteps, not even an explicit nil
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

### HasLinks

`func (o *CreateWorkItemApiModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *CreateWorkItemApiModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *CreateWorkItemApiModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetParameters

`func (o *CreateWorkItemApiModel) GetParameters() []WorkItemParameterKeyApiModel`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CreateWorkItemApiModel) GetParametersOk() (*[]WorkItemParameterKeyApiModel, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CreateWorkItemApiModel) SetParameters(v []WorkItemParameterKeyApiModel)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *CreateWorkItemApiModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *CreateWorkItemApiModel) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *CreateWorkItemApiModel) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


