# WorkItemChangedFieldsViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**NullableStringChangedFieldWithDiffsViewModel**](StringChangedFieldWithDiffsViewModel.md) |  | [optional] 
**IsDeleted** | [**BooleanChangedFieldViewModel**](BooleanChangedFieldViewModel.md) |  | 
**ProjectId** | [**GuidChangedFieldViewModel**](GuidChangedFieldViewModel.md) |  | 
**IsAutomated** | [**BooleanChangedFieldViewModel**](BooleanChangedFieldViewModel.md) |  | 
**SectionId** | [**GuidChangedFieldViewModel**](GuidChangedFieldViewModel.md) |  | 
**Description** | Pointer to [**NullableStringChangedFieldWithDiffsViewModel**](StringChangedFieldWithDiffsViewModel.md) |  | [optional] 
**State** | [**StringChangedFieldViewModel**](StringChangedFieldViewModel.md) |  | 
**Priority** | [**StringChangedFieldViewModel**](StringChangedFieldViewModel.md) |  | 
**Duration** | [**Int32ChangedFieldViewModel**](Int32ChangedFieldViewModel.md) |  | 
**Attributes** | [**map[string]WorkItemChangedAttributeViewModel**](WorkItemChangedAttributeViewModel.md) |  | 
**Steps** | [**WorkItemStepChangeViewModelArrayChangedFieldWithDiffsViewModel**](WorkItemStepChangeViewModelArrayChangedFieldWithDiffsViewModel.md) |  | 
**PreconditionSteps** | [**WorkItemStepChangeViewModelArrayChangedFieldWithDiffsViewModel**](WorkItemStepChangeViewModelArrayChangedFieldWithDiffsViewModel.md) |  | 
**PostconditionSteps** | [**WorkItemStepChangeViewModelArrayChangedFieldWithDiffsViewModel**](WorkItemStepChangeViewModelArrayChangedFieldWithDiffsViewModel.md) |  | 
**AutoTests** | [**AutoTestChangeViewModelArrayChangedFieldViewModel**](AutoTestChangeViewModelArrayChangedFieldViewModel.md) |  | 
**Attachments** | [**AttachmentChangeViewModelArrayChangedFieldViewModel**](AttachmentChangeViewModelArrayChangedFieldViewModel.md) |  | 
**Tags** | [**StringArrayChangedFieldViewModel**](StringArrayChangedFieldViewModel.md) |  | 
**Links** | [**WorkItemLinkChangeViewModelArrayChangedFieldViewModel**](WorkItemLinkChangeViewModelArrayChangedFieldViewModel.md) |  | 
**GlobalId** | [**Int64ChangedFieldViewModel**](Int64ChangedFieldViewModel.md) |  | 
**VersionNumber** | [**Int32ChangedFieldViewModel**](Int32ChangedFieldViewModel.md) |  | 
**EntityTypeName** | [**StringChangedFieldViewModel**](StringChangedFieldViewModel.md) |  | 

## Methods

### NewWorkItemChangedFieldsViewModel

`func NewWorkItemChangedFieldsViewModel(isDeleted BooleanChangedFieldViewModel, projectId GuidChangedFieldViewModel, isAutomated BooleanChangedFieldViewModel, sectionId GuidChangedFieldViewModel, state StringChangedFieldViewModel, priority StringChangedFieldViewModel, duration Int32ChangedFieldViewModel, attributes map[string]WorkItemChangedAttributeViewModel, steps WorkItemStepChangeViewModelArrayChangedFieldWithDiffsViewModel, preconditionSteps WorkItemStepChangeViewModelArrayChangedFieldWithDiffsViewModel, postconditionSteps WorkItemStepChangeViewModelArrayChangedFieldWithDiffsViewModel, autoTests AutoTestChangeViewModelArrayChangedFieldViewModel, attachments AttachmentChangeViewModelArrayChangedFieldViewModel, tags StringArrayChangedFieldViewModel, links WorkItemLinkChangeViewModelArrayChangedFieldViewModel, globalId Int64ChangedFieldViewModel, versionNumber Int32ChangedFieldViewModel, entityTypeName StringChangedFieldViewModel, ) *WorkItemChangedFieldsViewModel`

NewWorkItemChangedFieldsViewModel instantiates a new WorkItemChangedFieldsViewModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemChangedFieldsViewModelWithDefaults

`func NewWorkItemChangedFieldsViewModelWithDefaults() *WorkItemChangedFieldsViewModel`

NewWorkItemChangedFieldsViewModelWithDefaults instantiates a new WorkItemChangedFieldsViewModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkItemChangedFieldsViewModel) GetName() StringChangedFieldWithDiffsViewModel`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkItemChangedFieldsViewModel) GetNameOk() (*StringChangedFieldWithDiffsViewModel, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkItemChangedFieldsViewModel) SetName(v StringChangedFieldWithDiffsViewModel)`

SetName sets Name field to given value.

### HasName

`func (o *WorkItemChangedFieldsViewModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WorkItemChangedFieldsViewModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WorkItemChangedFieldsViewModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsDeleted

`func (o *WorkItemChangedFieldsViewModel) GetIsDeleted() BooleanChangedFieldViewModel`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *WorkItemChangedFieldsViewModel) GetIsDeletedOk() (*BooleanChangedFieldViewModel, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *WorkItemChangedFieldsViewModel) SetIsDeleted(v BooleanChangedFieldViewModel)`

SetIsDeleted sets IsDeleted field to given value.


### GetProjectId

`func (o *WorkItemChangedFieldsViewModel) GetProjectId() GuidChangedFieldViewModel`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WorkItemChangedFieldsViewModel) GetProjectIdOk() (*GuidChangedFieldViewModel, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WorkItemChangedFieldsViewModel) SetProjectId(v GuidChangedFieldViewModel)`

SetProjectId sets ProjectId field to given value.


### GetIsAutomated

`func (o *WorkItemChangedFieldsViewModel) GetIsAutomated() BooleanChangedFieldViewModel`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *WorkItemChangedFieldsViewModel) GetIsAutomatedOk() (*BooleanChangedFieldViewModel, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *WorkItemChangedFieldsViewModel) SetIsAutomated(v BooleanChangedFieldViewModel)`

SetIsAutomated sets IsAutomated field to given value.


### GetSectionId

`func (o *WorkItemChangedFieldsViewModel) GetSectionId() GuidChangedFieldViewModel`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *WorkItemChangedFieldsViewModel) GetSectionIdOk() (*GuidChangedFieldViewModel, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *WorkItemChangedFieldsViewModel) SetSectionId(v GuidChangedFieldViewModel)`

SetSectionId sets SectionId field to given value.


### GetDescription

`func (o *WorkItemChangedFieldsViewModel) GetDescription() StringChangedFieldWithDiffsViewModel`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkItemChangedFieldsViewModel) GetDescriptionOk() (*StringChangedFieldWithDiffsViewModel, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkItemChangedFieldsViewModel) SetDescription(v StringChangedFieldWithDiffsViewModel)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkItemChangedFieldsViewModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkItemChangedFieldsViewModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkItemChangedFieldsViewModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetState

`func (o *WorkItemChangedFieldsViewModel) GetState() StringChangedFieldViewModel`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkItemChangedFieldsViewModel) GetStateOk() (*StringChangedFieldViewModel, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkItemChangedFieldsViewModel) SetState(v StringChangedFieldViewModel)`

SetState sets State field to given value.


### GetPriority

`func (o *WorkItemChangedFieldsViewModel) GetPriority() StringChangedFieldViewModel`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WorkItemChangedFieldsViewModel) GetPriorityOk() (*StringChangedFieldViewModel, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WorkItemChangedFieldsViewModel) SetPriority(v StringChangedFieldViewModel)`

SetPriority sets Priority field to given value.


### GetDuration

`func (o *WorkItemChangedFieldsViewModel) GetDuration() Int32ChangedFieldViewModel`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WorkItemChangedFieldsViewModel) GetDurationOk() (*Int32ChangedFieldViewModel, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WorkItemChangedFieldsViewModel) SetDuration(v Int32ChangedFieldViewModel)`

SetDuration sets Duration field to given value.


### GetAttributes

`func (o *WorkItemChangedFieldsViewModel) GetAttributes() map[string]WorkItemChangedAttributeViewModel`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkItemChangedFieldsViewModel) GetAttributesOk() (*map[string]WorkItemChangedAttributeViewModel, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkItemChangedFieldsViewModel) SetAttributes(v map[string]WorkItemChangedAttributeViewModel)`

SetAttributes sets Attributes field to given value.


### GetSteps

`func (o *WorkItemChangedFieldsViewModel) GetSteps() WorkItemStepChangeViewModelArrayChangedFieldWithDiffsViewModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *WorkItemChangedFieldsViewModel) GetStepsOk() (*WorkItemStepChangeViewModelArrayChangedFieldWithDiffsViewModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *WorkItemChangedFieldsViewModel) SetSteps(v WorkItemStepChangeViewModelArrayChangedFieldWithDiffsViewModel)`

SetSteps sets Steps field to given value.


### GetPreconditionSteps

`func (o *WorkItemChangedFieldsViewModel) GetPreconditionSteps() WorkItemStepChangeViewModelArrayChangedFieldWithDiffsViewModel`

GetPreconditionSteps returns the PreconditionSteps field if non-nil, zero value otherwise.

### GetPreconditionStepsOk

`func (o *WorkItemChangedFieldsViewModel) GetPreconditionStepsOk() (*WorkItemStepChangeViewModelArrayChangedFieldWithDiffsViewModel, bool)`

GetPreconditionStepsOk returns a tuple with the PreconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditionSteps

`func (o *WorkItemChangedFieldsViewModel) SetPreconditionSteps(v WorkItemStepChangeViewModelArrayChangedFieldWithDiffsViewModel)`

SetPreconditionSteps sets PreconditionSteps field to given value.


### GetPostconditionSteps

`func (o *WorkItemChangedFieldsViewModel) GetPostconditionSteps() WorkItemStepChangeViewModelArrayChangedFieldWithDiffsViewModel`

GetPostconditionSteps returns the PostconditionSteps field if non-nil, zero value otherwise.

### GetPostconditionStepsOk

`func (o *WorkItemChangedFieldsViewModel) GetPostconditionStepsOk() (*WorkItemStepChangeViewModelArrayChangedFieldWithDiffsViewModel, bool)`

GetPostconditionStepsOk returns a tuple with the PostconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditionSteps

`func (o *WorkItemChangedFieldsViewModel) SetPostconditionSteps(v WorkItemStepChangeViewModelArrayChangedFieldWithDiffsViewModel)`

SetPostconditionSteps sets PostconditionSteps field to given value.


### GetAutoTests

`func (o *WorkItemChangedFieldsViewModel) GetAutoTests() AutoTestChangeViewModelArrayChangedFieldViewModel`

GetAutoTests returns the AutoTests field if non-nil, zero value otherwise.

### GetAutoTestsOk

`func (o *WorkItemChangedFieldsViewModel) GetAutoTestsOk() (*AutoTestChangeViewModelArrayChangedFieldViewModel, bool)`

GetAutoTestsOk returns a tuple with the AutoTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTests

`func (o *WorkItemChangedFieldsViewModel) SetAutoTests(v AutoTestChangeViewModelArrayChangedFieldViewModel)`

SetAutoTests sets AutoTests field to given value.


### GetAttachments

`func (o *WorkItemChangedFieldsViewModel) GetAttachments() AttachmentChangeViewModelArrayChangedFieldViewModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *WorkItemChangedFieldsViewModel) GetAttachmentsOk() (*AttachmentChangeViewModelArrayChangedFieldViewModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *WorkItemChangedFieldsViewModel) SetAttachments(v AttachmentChangeViewModelArrayChangedFieldViewModel)`

SetAttachments sets Attachments field to given value.


### GetTags

`func (o *WorkItemChangedFieldsViewModel) GetTags() StringArrayChangedFieldViewModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkItemChangedFieldsViewModel) GetTagsOk() (*StringArrayChangedFieldViewModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkItemChangedFieldsViewModel) SetTags(v StringArrayChangedFieldViewModel)`

SetTags sets Tags field to given value.


### GetLinks

`func (o *WorkItemChangedFieldsViewModel) GetLinks() WorkItemLinkChangeViewModelArrayChangedFieldViewModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WorkItemChangedFieldsViewModel) GetLinksOk() (*WorkItemLinkChangeViewModelArrayChangedFieldViewModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WorkItemChangedFieldsViewModel) SetLinks(v WorkItemLinkChangeViewModelArrayChangedFieldViewModel)`

SetLinks sets Links field to given value.


### GetGlobalId

`func (o *WorkItemChangedFieldsViewModel) GetGlobalId() Int64ChangedFieldViewModel`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *WorkItemChangedFieldsViewModel) GetGlobalIdOk() (*Int64ChangedFieldViewModel, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *WorkItemChangedFieldsViewModel) SetGlobalId(v Int64ChangedFieldViewModel)`

SetGlobalId sets GlobalId field to given value.


### GetVersionNumber

`func (o *WorkItemChangedFieldsViewModel) GetVersionNumber() Int32ChangedFieldViewModel`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *WorkItemChangedFieldsViewModel) GetVersionNumberOk() (*Int32ChangedFieldViewModel, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *WorkItemChangedFieldsViewModel) SetVersionNumber(v Int32ChangedFieldViewModel)`

SetVersionNumber sets VersionNumber field to given value.


### GetEntityTypeName

`func (o *WorkItemChangedFieldsViewModel) GetEntityTypeName() StringChangedFieldViewModel`

GetEntityTypeName returns the EntityTypeName field if non-nil, zero value otherwise.

### GetEntityTypeNameOk

`func (o *WorkItemChangedFieldsViewModel) GetEntityTypeNameOk() (*StringChangedFieldViewModel, bool)`

GetEntityTypeNameOk returns a tuple with the EntityTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypeName

`func (o *WorkItemChangedFieldsViewModel) SetEntityTypeName(v StringChangedFieldViewModel)`

SetEntityTypeName sets EntityTypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


