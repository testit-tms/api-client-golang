# WorkItemChangedFieldsViewModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**NullableStringChangedFieldWithDiffsViewModel**](StringChangedFieldWithDiffsViewModel.md) |  | [optional] 
**IsDeleted** | Pointer to [**WorkItemChangedFieldsViewModelIsDeleted**](WorkItemChangedFieldsViewModelIsDeleted.md) |  | [optional] 
**ProjectId** | Pointer to [**WorkItemChangedFieldsViewModelProjectId**](WorkItemChangedFieldsViewModelProjectId.md) |  | [optional] 
**IsAutomated** | Pointer to [**WorkItemChangedFieldsViewModelIsDeleted**](WorkItemChangedFieldsViewModelIsDeleted.md) |  | [optional] 
**SectionId** | Pointer to [**WorkItemChangedFieldsViewModelProjectId**](WorkItemChangedFieldsViewModelProjectId.md) |  | [optional] 
**Description** | Pointer to [**NullableStringChangedFieldWithDiffsViewModel**](StringChangedFieldWithDiffsViewModel.md) |  | [optional] 
**State** | Pointer to [**WorkItemChangedFieldsViewModelState**](WorkItemChangedFieldsViewModelState.md) |  | [optional] 
**Priority** | Pointer to [**WorkItemChangedFieldsViewModelState**](WorkItemChangedFieldsViewModelState.md) |  | [optional] 
**Duration** | Pointer to [**WorkItemChangedFieldsViewModelDuration**](WorkItemChangedFieldsViewModelDuration.md) |  | [optional] 
**Attributes** | Pointer to [**map[string]WorkItemChangedAttributeViewModel**](WorkItemChangedAttributeViewModel.md) |  | [optional] 
**Steps** | Pointer to [**WorkItemChangedFieldsViewModelSteps**](WorkItemChangedFieldsViewModelSteps.md) |  | [optional] 
**PreconditionSteps** | Pointer to [**WorkItemChangedFieldsViewModelSteps**](WorkItemChangedFieldsViewModelSteps.md) |  | [optional] 
**PostconditionSteps** | Pointer to [**WorkItemChangedFieldsViewModelSteps**](WorkItemChangedFieldsViewModelSteps.md) |  | [optional] 
**AutoTests** | Pointer to [**WorkItemChangedFieldsViewModelAutoTests**](WorkItemChangedFieldsViewModelAutoTests.md) |  | [optional] 
**Attachments** | Pointer to [**WorkItemChangedFieldsViewModelAttachments**](WorkItemChangedFieldsViewModelAttachments.md) |  | [optional] 
**Tags** | Pointer to [**WorkItemChangedFieldsViewModelTags**](WorkItemChangedFieldsViewModelTags.md) |  | [optional] 
**Links** | Pointer to [**WorkItemChangedFieldsViewModelLinks**](WorkItemChangedFieldsViewModelLinks.md) |  | [optional] 
**GlobalId** | Pointer to [**WorkItemChangedFieldsViewModelGlobalId**](WorkItemChangedFieldsViewModelGlobalId.md) |  | [optional] 
**VersionNumber** | Pointer to [**WorkItemChangedFieldsViewModelDuration**](WorkItemChangedFieldsViewModelDuration.md) |  | [optional] 
**EntityTypeName** | Pointer to [**WorkItemChangedFieldsViewModelState**](WorkItemChangedFieldsViewModelState.md) |  | [optional] 

## Methods

### NewWorkItemChangedFieldsViewModel

`func NewWorkItemChangedFieldsViewModel() *WorkItemChangedFieldsViewModel`

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

`func (o *WorkItemChangedFieldsViewModel) GetIsDeleted() WorkItemChangedFieldsViewModelIsDeleted`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *WorkItemChangedFieldsViewModel) GetIsDeletedOk() (*WorkItemChangedFieldsViewModelIsDeleted, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *WorkItemChangedFieldsViewModel) SetIsDeleted(v WorkItemChangedFieldsViewModelIsDeleted)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *WorkItemChangedFieldsViewModel) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetProjectId

`func (o *WorkItemChangedFieldsViewModel) GetProjectId() WorkItemChangedFieldsViewModelProjectId`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WorkItemChangedFieldsViewModel) GetProjectIdOk() (*WorkItemChangedFieldsViewModelProjectId, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WorkItemChangedFieldsViewModel) SetProjectId(v WorkItemChangedFieldsViewModelProjectId)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *WorkItemChangedFieldsViewModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetIsAutomated

`func (o *WorkItemChangedFieldsViewModel) GetIsAutomated() WorkItemChangedFieldsViewModelIsDeleted`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *WorkItemChangedFieldsViewModel) GetIsAutomatedOk() (*WorkItemChangedFieldsViewModelIsDeleted, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *WorkItemChangedFieldsViewModel) SetIsAutomated(v WorkItemChangedFieldsViewModelIsDeleted)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *WorkItemChangedFieldsViewModel) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### GetSectionId

`func (o *WorkItemChangedFieldsViewModel) GetSectionId() WorkItemChangedFieldsViewModelProjectId`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *WorkItemChangedFieldsViewModel) GetSectionIdOk() (*WorkItemChangedFieldsViewModelProjectId, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *WorkItemChangedFieldsViewModel) SetSectionId(v WorkItemChangedFieldsViewModelProjectId)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *WorkItemChangedFieldsViewModel) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

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

`func (o *WorkItemChangedFieldsViewModel) GetState() WorkItemChangedFieldsViewModelState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkItemChangedFieldsViewModel) GetStateOk() (*WorkItemChangedFieldsViewModelState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkItemChangedFieldsViewModel) SetState(v WorkItemChangedFieldsViewModelState)`

SetState sets State field to given value.

### HasState

`func (o *WorkItemChangedFieldsViewModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPriority

`func (o *WorkItemChangedFieldsViewModel) GetPriority() WorkItemChangedFieldsViewModelState`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WorkItemChangedFieldsViewModel) GetPriorityOk() (*WorkItemChangedFieldsViewModelState, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WorkItemChangedFieldsViewModel) SetPriority(v WorkItemChangedFieldsViewModelState)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *WorkItemChangedFieldsViewModel) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetDuration

`func (o *WorkItemChangedFieldsViewModel) GetDuration() WorkItemChangedFieldsViewModelDuration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WorkItemChangedFieldsViewModel) GetDurationOk() (*WorkItemChangedFieldsViewModelDuration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WorkItemChangedFieldsViewModel) SetDuration(v WorkItemChangedFieldsViewModelDuration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *WorkItemChangedFieldsViewModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

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

### HasAttributes

`func (o *WorkItemChangedFieldsViewModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetSteps

`func (o *WorkItemChangedFieldsViewModel) GetSteps() WorkItemChangedFieldsViewModelSteps`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *WorkItemChangedFieldsViewModel) GetStepsOk() (*WorkItemChangedFieldsViewModelSteps, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *WorkItemChangedFieldsViewModel) SetSteps(v WorkItemChangedFieldsViewModelSteps)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *WorkItemChangedFieldsViewModel) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetPreconditionSteps

`func (o *WorkItemChangedFieldsViewModel) GetPreconditionSteps() WorkItemChangedFieldsViewModelSteps`

GetPreconditionSteps returns the PreconditionSteps field if non-nil, zero value otherwise.

### GetPreconditionStepsOk

`func (o *WorkItemChangedFieldsViewModel) GetPreconditionStepsOk() (*WorkItemChangedFieldsViewModelSteps, bool)`

GetPreconditionStepsOk returns a tuple with the PreconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditionSteps

`func (o *WorkItemChangedFieldsViewModel) SetPreconditionSteps(v WorkItemChangedFieldsViewModelSteps)`

SetPreconditionSteps sets PreconditionSteps field to given value.

### HasPreconditionSteps

`func (o *WorkItemChangedFieldsViewModel) HasPreconditionSteps() bool`

HasPreconditionSteps returns a boolean if a field has been set.

### GetPostconditionSteps

`func (o *WorkItemChangedFieldsViewModel) GetPostconditionSteps() WorkItemChangedFieldsViewModelSteps`

GetPostconditionSteps returns the PostconditionSteps field if non-nil, zero value otherwise.

### GetPostconditionStepsOk

`func (o *WorkItemChangedFieldsViewModel) GetPostconditionStepsOk() (*WorkItemChangedFieldsViewModelSteps, bool)`

GetPostconditionStepsOk returns a tuple with the PostconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditionSteps

`func (o *WorkItemChangedFieldsViewModel) SetPostconditionSteps(v WorkItemChangedFieldsViewModelSteps)`

SetPostconditionSteps sets PostconditionSteps field to given value.

### HasPostconditionSteps

`func (o *WorkItemChangedFieldsViewModel) HasPostconditionSteps() bool`

HasPostconditionSteps returns a boolean if a field has been set.

### GetAutoTests

`func (o *WorkItemChangedFieldsViewModel) GetAutoTests() WorkItemChangedFieldsViewModelAutoTests`

GetAutoTests returns the AutoTests field if non-nil, zero value otherwise.

### GetAutoTestsOk

`func (o *WorkItemChangedFieldsViewModel) GetAutoTestsOk() (*WorkItemChangedFieldsViewModelAutoTests, bool)`

GetAutoTestsOk returns a tuple with the AutoTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTests

`func (o *WorkItemChangedFieldsViewModel) SetAutoTests(v WorkItemChangedFieldsViewModelAutoTests)`

SetAutoTests sets AutoTests field to given value.

### HasAutoTests

`func (o *WorkItemChangedFieldsViewModel) HasAutoTests() bool`

HasAutoTests returns a boolean if a field has been set.

### GetAttachments

`func (o *WorkItemChangedFieldsViewModel) GetAttachments() WorkItemChangedFieldsViewModelAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *WorkItemChangedFieldsViewModel) GetAttachmentsOk() (*WorkItemChangedFieldsViewModelAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *WorkItemChangedFieldsViewModel) SetAttachments(v WorkItemChangedFieldsViewModelAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *WorkItemChangedFieldsViewModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetTags

`func (o *WorkItemChangedFieldsViewModel) GetTags() WorkItemChangedFieldsViewModelTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkItemChangedFieldsViewModel) GetTagsOk() (*WorkItemChangedFieldsViewModelTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkItemChangedFieldsViewModel) SetTags(v WorkItemChangedFieldsViewModelTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WorkItemChangedFieldsViewModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLinks

`func (o *WorkItemChangedFieldsViewModel) GetLinks() WorkItemChangedFieldsViewModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WorkItemChangedFieldsViewModel) GetLinksOk() (*WorkItemChangedFieldsViewModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WorkItemChangedFieldsViewModel) SetLinks(v WorkItemChangedFieldsViewModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WorkItemChangedFieldsViewModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetGlobalId

`func (o *WorkItemChangedFieldsViewModel) GetGlobalId() WorkItemChangedFieldsViewModelGlobalId`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *WorkItemChangedFieldsViewModel) GetGlobalIdOk() (*WorkItemChangedFieldsViewModelGlobalId, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *WorkItemChangedFieldsViewModel) SetGlobalId(v WorkItemChangedFieldsViewModelGlobalId)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *WorkItemChangedFieldsViewModel) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### GetVersionNumber

`func (o *WorkItemChangedFieldsViewModel) GetVersionNumber() WorkItemChangedFieldsViewModelDuration`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *WorkItemChangedFieldsViewModel) GetVersionNumberOk() (*WorkItemChangedFieldsViewModelDuration, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *WorkItemChangedFieldsViewModel) SetVersionNumber(v WorkItemChangedFieldsViewModelDuration)`

SetVersionNumber sets VersionNumber field to given value.

### HasVersionNumber

`func (o *WorkItemChangedFieldsViewModel) HasVersionNumber() bool`

HasVersionNumber returns a boolean if a field has been set.

### GetEntityTypeName

`func (o *WorkItemChangedFieldsViewModel) GetEntityTypeName() WorkItemChangedFieldsViewModelState`

GetEntityTypeName returns the EntityTypeName field if non-nil, zero value otherwise.

### GetEntityTypeNameOk

`func (o *WorkItemChangedFieldsViewModel) GetEntityTypeNameOk() (*WorkItemChangedFieldsViewModelState, bool)`

GetEntityTypeNameOk returns a tuple with the EntityTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypeName

`func (o *WorkItemChangedFieldsViewModel) SetEntityTypeName(v WorkItemChangedFieldsViewModelState)`

SetEntityTypeName sets EntityTypeName field to given value.

### HasEntityTypeName

`func (o *WorkItemChangedFieldsViewModel) HasEntityTypeName() bool`

HasEntityTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


