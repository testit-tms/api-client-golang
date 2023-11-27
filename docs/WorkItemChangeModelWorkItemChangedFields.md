# WorkItemChangeModelWorkItemChangedFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**NullableStringChangedFieldWithDiffsViewModel**](StringChangedFieldWithDiffsViewModel.md) |  | [optional] 
**IsDeleted** | [**WorkItemChangedFieldsViewModelIsDeleted**](WorkItemChangedFieldsViewModelIsDeleted.md) |  | 
**ProjectId** | [**WorkItemChangedFieldsViewModelProjectId**](WorkItemChangedFieldsViewModelProjectId.md) |  | 
**IsAutomated** | [**WorkItemChangedFieldsViewModelIsDeleted**](WorkItemChangedFieldsViewModelIsDeleted.md) |  | 
**SectionId** | [**WorkItemChangedFieldsViewModelProjectId**](WorkItemChangedFieldsViewModelProjectId.md) |  | 
**Description** | Pointer to [**NullableStringChangedFieldWithDiffsViewModel**](StringChangedFieldWithDiffsViewModel.md) |  | [optional] 
**State** | [**WorkItemChangedFieldsViewModelState**](WorkItemChangedFieldsViewModelState.md) |  | 
**Priority** | [**WorkItemChangedFieldsViewModelState**](WorkItemChangedFieldsViewModelState.md) |  | 
**Duration** | [**WorkItemChangedFieldsViewModelDuration**](WorkItemChangedFieldsViewModelDuration.md) |  | 
**Attributes** | [**map[string]WorkItemChangedAttributeViewModel**](WorkItemChangedAttributeViewModel.md) |  | 
**Steps** | [**WorkItemChangedFieldsViewModelSteps**](WorkItemChangedFieldsViewModelSteps.md) |  | 
**PreconditionSteps** | [**WorkItemChangedFieldsViewModelSteps**](WorkItemChangedFieldsViewModelSteps.md) |  | 
**PostconditionSteps** | [**WorkItemChangedFieldsViewModelSteps**](WorkItemChangedFieldsViewModelSteps.md) |  | 
**AutoTests** | [**WorkItemChangedFieldsViewModelAutoTests**](WorkItemChangedFieldsViewModelAutoTests.md) |  | 
**Attachments** | [**WorkItemChangedFieldsViewModelAttachments**](WorkItemChangedFieldsViewModelAttachments.md) |  | 
**Tags** | [**WorkItemChangedFieldsViewModelTags**](WorkItemChangedFieldsViewModelTags.md) |  | 
**Links** | [**WorkItemChangedFieldsViewModelLinks**](WorkItemChangedFieldsViewModelLinks.md) |  | 
**GlobalId** | [**WorkItemChangedFieldsViewModelGlobalId**](WorkItemChangedFieldsViewModelGlobalId.md) |  | 
**VersionNumber** | [**WorkItemChangedFieldsViewModelDuration**](WorkItemChangedFieldsViewModelDuration.md) |  | 
**EntityTypeName** | [**WorkItemChangedFieldsViewModelState**](WorkItemChangedFieldsViewModelState.md) |  | 

## Methods

### NewWorkItemChangeModelWorkItemChangedFields

`func NewWorkItemChangeModelWorkItemChangedFields(isDeleted WorkItemChangedFieldsViewModelIsDeleted, projectId WorkItemChangedFieldsViewModelProjectId, isAutomated WorkItemChangedFieldsViewModelIsDeleted, sectionId WorkItemChangedFieldsViewModelProjectId, state WorkItemChangedFieldsViewModelState, priority WorkItemChangedFieldsViewModelState, duration WorkItemChangedFieldsViewModelDuration, attributes map[string]WorkItemChangedAttributeViewModel, steps WorkItemChangedFieldsViewModelSteps, preconditionSteps WorkItemChangedFieldsViewModelSteps, postconditionSteps WorkItemChangedFieldsViewModelSteps, autoTests WorkItemChangedFieldsViewModelAutoTests, attachments WorkItemChangedFieldsViewModelAttachments, tags WorkItemChangedFieldsViewModelTags, links WorkItemChangedFieldsViewModelLinks, globalId WorkItemChangedFieldsViewModelGlobalId, versionNumber WorkItemChangedFieldsViewModelDuration, entityTypeName WorkItemChangedFieldsViewModelState, ) *WorkItemChangeModelWorkItemChangedFields`

NewWorkItemChangeModelWorkItemChangedFields instantiates a new WorkItemChangeModelWorkItemChangedFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemChangeModelWorkItemChangedFieldsWithDefaults

`func NewWorkItemChangeModelWorkItemChangedFieldsWithDefaults() *WorkItemChangeModelWorkItemChangedFields`

NewWorkItemChangeModelWorkItemChangedFieldsWithDefaults instantiates a new WorkItemChangeModelWorkItemChangedFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkItemChangeModelWorkItemChangedFields) GetName() StringChangedFieldWithDiffsViewModel`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkItemChangeModelWorkItemChangedFields) GetNameOk() (*StringChangedFieldWithDiffsViewModel, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkItemChangeModelWorkItemChangedFields) SetName(v StringChangedFieldWithDiffsViewModel)`

SetName sets Name field to given value.

### HasName

`func (o *WorkItemChangeModelWorkItemChangedFields) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WorkItemChangeModelWorkItemChangedFields) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WorkItemChangeModelWorkItemChangedFields) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsDeleted

`func (o *WorkItemChangeModelWorkItemChangedFields) GetIsDeleted() WorkItemChangedFieldsViewModelIsDeleted`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *WorkItemChangeModelWorkItemChangedFields) GetIsDeletedOk() (*WorkItemChangedFieldsViewModelIsDeleted, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *WorkItemChangeModelWorkItemChangedFields) SetIsDeleted(v WorkItemChangedFieldsViewModelIsDeleted)`

SetIsDeleted sets IsDeleted field to given value.


### GetProjectId

`func (o *WorkItemChangeModelWorkItemChangedFields) GetProjectId() WorkItemChangedFieldsViewModelProjectId`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WorkItemChangeModelWorkItemChangedFields) GetProjectIdOk() (*WorkItemChangedFieldsViewModelProjectId, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WorkItemChangeModelWorkItemChangedFields) SetProjectId(v WorkItemChangedFieldsViewModelProjectId)`

SetProjectId sets ProjectId field to given value.


### GetIsAutomated

`func (o *WorkItemChangeModelWorkItemChangedFields) GetIsAutomated() WorkItemChangedFieldsViewModelIsDeleted`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *WorkItemChangeModelWorkItemChangedFields) GetIsAutomatedOk() (*WorkItemChangedFieldsViewModelIsDeleted, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *WorkItemChangeModelWorkItemChangedFields) SetIsAutomated(v WorkItemChangedFieldsViewModelIsDeleted)`

SetIsAutomated sets IsAutomated field to given value.


### GetSectionId

`func (o *WorkItemChangeModelWorkItemChangedFields) GetSectionId() WorkItemChangedFieldsViewModelProjectId`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *WorkItemChangeModelWorkItemChangedFields) GetSectionIdOk() (*WorkItemChangedFieldsViewModelProjectId, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *WorkItemChangeModelWorkItemChangedFields) SetSectionId(v WorkItemChangedFieldsViewModelProjectId)`

SetSectionId sets SectionId field to given value.


### GetDescription

`func (o *WorkItemChangeModelWorkItemChangedFields) GetDescription() StringChangedFieldWithDiffsViewModel`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkItemChangeModelWorkItemChangedFields) GetDescriptionOk() (*StringChangedFieldWithDiffsViewModel, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkItemChangeModelWorkItemChangedFields) SetDescription(v StringChangedFieldWithDiffsViewModel)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkItemChangeModelWorkItemChangedFields) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkItemChangeModelWorkItemChangedFields) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkItemChangeModelWorkItemChangedFields) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetState

`func (o *WorkItemChangeModelWorkItemChangedFields) GetState() WorkItemChangedFieldsViewModelState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkItemChangeModelWorkItemChangedFields) GetStateOk() (*WorkItemChangedFieldsViewModelState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkItemChangeModelWorkItemChangedFields) SetState(v WorkItemChangedFieldsViewModelState)`

SetState sets State field to given value.


### GetPriority

`func (o *WorkItemChangeModelWorkItemChangedFields) GetPriority() WorkItemChangedFieldsViewModelState`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WorkItemChangeModelWorkItemChangedFields) GetPriorityOk() (*WorkItemChangedFieldsViewModelState, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WorkItemChangeModelWorkItemChangedFields) SetPriority(v WorkItemChangedFieldsViewModelState)`

SetPriority sets Priority field to given value.


### GetDuration

`func (o *WorkItemChangeModelWorkItemChangedFields) GetDuration() WorkItemChangedFieldsViewModelDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WorkItemChangeModelWorkItemChangedFields) GetDurationOk() (*WorkItemChangedFieldsViewModelDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WorkItemChangeModelWorkItemChangedFields) SetDuration(v WorkItemChangedFieldsViewModelDuration)`

SetDuration sets Duration field to given value.


### GetAttributes

`func (o *WorkItemChangeModelWorkItemChangedFields) GetAttributes() map[string]WorkItemChangedAttributeViewModel`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkItemChangeModelWorkItemChangedFields) GetAttributesOk() (*map[string]WorkItemChangedAttributeViewModel, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkItemChangeModelWorkItemChangedFields) SetAttributes(v map[string]WorkItemChangedAttributeViewModel)`

SetAttributes sets Attributes field to given value.


### GetSteps

`func (o *WorkItemChangeModelWorkItemChangedFields) GetSteps() WorkItemChangedFieldsViewModelSteps`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *WorkItemChangeModelWorkItemChangedFields) GetStepsOk() (*WorkItemChangedFieldsViewModelSteps, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *WorkItemChangeModelWorkItemChangedFields) SetSteps(v WorkItemChangedFieldsViewModelSteps)`

SetSteps sets Steps field to given value.


### GetPreconditionSteps

`func (o *WorkItemChangeModelWorkItemChangedFields) GetPreconditionSteps() WorkItemChangedFieldsViewModelSteps`

GetPreconditionSteps returns the PreconditionSteps field if non-nil, zero value otherwise.

### GetPreconditionStepsOk

`func (o *WorkItemChangeModelWorkItemChangedFields) GetPreconditionStepsOk() (*WorkItemChangedFieldsViewModelSteps, bool)`

GetPreconditionStepsOk returns a tuple with the PreconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditionSteps

`func (o *WorkItemChangeModelWorkItemChangedFields) SetPreconditionSteps(v WorkItemChangedFieldsViewModelSteps)`

SetPreconditionSteps sets PreconditionSteps field to given value.


### GetPostconditionSteps

`func (o *WorkItemChangeModelWorkItemChangedFields) GetPostconditionSteps() WorkItemChangedFieldsViewModelSteps`

GetPostconditionSteps returns the PostconditionSteps field if non-nil, zero value otherwise.

### GetPostconditionStepsOk

`func (o *WorkItemChangeModelWorkItemChangedFields) GetPostconditionStepsOk() (*WorkItemChangedFieldsViewModelSteps, bool)`

GetPostconditionStepsOk returns a tuple with the PostconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditionSteps

`func (o *WorkItemChangeModelWorkItemChangedFields) SetPostconditionSteps(v WorkItemChangedFieldsViewModelSteps)`

SetPostconditionSteps sets PostconditionSteps field to given value.


### GetAutoTests

`func (o *WorkItemChangeModelWorkItemChangedFields) GetAutoTests() WorkItemChangedFieldsViewModelAutoTests`

GetAutoTests returns the AutoTests field if non-nil, zero value otherwise.

### GetAutoTestsOk

`func (o *WorkItemChangeModelWorkItemChangedFields) GetAutoTestsOk() (*WorkItemChangedFieldsViewModelAutoTests, bool)`

GetAutoTestsOk returns a tuple with the AutoTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTests

`func (o *WorkItemChangeModelWorkItemChangedFields) SetAutoTests(v WorkItemChangedFieldsViewModelAutoTests)`

SetAutoTests sets AutoTests field to given value.


### GetAttachments

`func (o *WorkItemChangeModelWorkItemChangedFields) GetAttachments() WorkItemChangedFieldsViewModelAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *WorkItemChangeModelWorkItemChangedFields) GetAttachmentsOk() (*WorkItemChangedFieldsViewModelAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *WorkItemChangeModelWorkItemChangedFields) SetAttachments(v WorkItemChangedFieldsViewModelAttachments)`

SetAttachments sets Attachments field to given value.


### GetTags

`func (o *WorkItemChangeModelWorkItemChangedFields) GetTags() WorkItemChangedFieldsViewModelTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkItemChangeModelWorkItemChangedFields) GetTagsOk() (*WorkItemChangedFieldsViewModelTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkItemChangeModelWorkItemChangedFields) SetTags(v WorkItemChangedFieldsViewModelTags)`

SetTags sets Tags field to given value.


### GetLinks

`func (o *WorkItemChangeModelWorkItemChangedFields) GetLinks() WorkItemChangedFieldsViewModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WorkItemChangeModelWorkItemChangedFields) GetLinksOk() (*WorkItemChangedFieldsViewModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WorkItemChangeModelWorkItemChangedFields) SetLinks(v WorkItemChangedFieldsViewModelLinks)`

SetLinks sets Links field to given value.


### GetGlobalId

`func (o *WorkItemChangeModelWorkItemChangedFields) GetGlobalId() WorkItemChangedFieldsViewModelGlobalId`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *WorkItemChangeModelWorkItemChangedFields) GetGlobalIdOk() (*WorkItemChangedFieldsViewModelGlobalId, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *WorkItemChangeModelWorkItemChangedFields) SetGlobalId(v WorkItemChangedFieldsViewModelGlobalId)`

SetGlobalId sets GlobalId field to given value.


### GetVersionNumber

`func (o *WorkItemChangeModelWorkItemChangedFields) GetVersionNumber() WorkItemChangedFieldsViewModelDuration`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *WorkItemChangeModelWorkItemChangedFields) GetVersionNumberOk() (*WorkItemChangedFieldsViewModelDuration, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *WorkItemChangeModelWorkItemChangedFields) SetVersionNumber(v WorkItemChangedFieldsViewModelDuration)`

SetVersionNumber sets VersionNumber field to given value.


### GetEntityTypeName

`func (o *WorkItemChangeModelWorkItemChangedFields) GetEntityTypeName() WorkItemChangedFieldsViewModelState`

GetEntityTypeName returns the EntityTypeName field if non-nil, zero value otherwise.

### GetEntityTypeNameOk

`func (o *WorkItemChangeModelWorkItemChangedFields) GetEntityTypeNameOk() (*WorkItemChangedFieldsViewModelState, bool)`

GetEntityTypeNameOk returns a tuple with the EntityTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypeName

`func (o *WorkItemChangeModelWorkItemChangedFields) SetEntityTypeName(v WorkItemChangedFieldsViewModelState)`

SetEntityTypeName sets EntityTypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


