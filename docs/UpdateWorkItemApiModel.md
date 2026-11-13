# UpdateWorkItemApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the work item | 
**SectionId** | **string** | Unique identifier of the section within a project | 
**Name** | **string** | Name of the work item | 
**Description** | Pointer to **NullableString** | Description of the work item | [optional] 
**Duration** | **int64** | Duration of the work item in milliseconds | 
**State** | [**WorkItemStateApiModel**](WorkItemStateApiModel.md) | Current state of the work item | 
**Priority** | [**WorkItemPriorityApiModel**](WorkItemPriorityApiModel.md) | Priority level assigned to the work item | 
**Attributes** | Pointer to **map[string]interface{}** | Set of custom attributes associated with the work item | [optional] 
**Tags** | Pointer to [**[]TagModel**](TagModel.md) | Set of tags applied to the work item | [optional] 
**PreconditionSteps** | Pointer to [**[]UpdateStepApiModel**](UpdateStepApiModel.md) | Set of precondition steps that must be executed before the main steps | [optional] 
**Steps** | Pointer to [**[]UpdateStepApiModel**](UpdateStepApiModel.md) | Set of main steps or actions defined for the work item | [optional] 
**PostconditionSteps** | Pointer to [**[]UpdateStepApiModel**](UpdateStepApiModel.md) | Set of postcondition steps that are executed after completing the main steps | [optional] 
**Iterations** | Pointer to [**[]AssignIterationApiModel**](AssignIterationApiModel.md) | Set of iterations associated with the work item | [optional] 
**AutoTests** | Pointer to [**[]AutoTestIdModel**](AutoTestIdModel.md) | Set of automated tests linked to the work item | [optional] 
**ManualTestId** | Pointer to **NullableString** | Identifier of the linked manual test | [optional] 
**AutoTestCases** | Pointer to [**[]AssignAutoTestCaseIdApiModel**](AssignAutoTestCaseIdApiModel.md) | Set of automated test cases linked to the work item | [optional] 
**Attachments** | Pointer to [**[]AssignAttachmentApiModel**](AssignAttachmentApiModel.md) | Set of files attached to the work item | [optional] 
**Links** | Pointer to [**[]UpdateLinkApiModel**](UpdateLinkApiModel.md) | Set of links related to the work item | [optional] 
**Parameters** | Pointer to [**[]WorkItemParameterKeyApiModel**](WorkItemParameterKeyApiModel.md) | Set of parameter keys associated with the work item | [optional] 
**LayerName** | Pointer to **NullableString** | Name of test pyramid layer associated with the work item | [optional] 
**LayerSource** | Pointer to [**NullableLayerSource**](LayerSource.md) | Source of test pyramid layer associated with the work item | [optional] 
**ResetLayer** | Pointer to **NullableBool** | Reset current test pyramid layer associated with the work item and re-evaluates it | [optional] 

## Methods

### NewUpdateWorkItemApiModel

`func NewUpdateWorkItemApiModel(id string, sectionId string, name string, duration int64, state WorkItemStateApiModel, priority WorkItemPriorityApiModel, ) *UpdateWorkItemApiModel`

NewUpdateWorkItemApiModel instantiates a new UpdateWorkItemApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWorkItemApiModelWithDefaults

`func NewUpdateWorkItemApiModelWithDefaults() *UpdateWorkItemApiModel`

NewUpdateWorkItemApiModelWithDefaults instantiates a new UpdateWorkItemApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateWorkItemApiModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateWorkItemApiModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateWorkItemApiModel) SetId(v string)`

SetId sets Id field to given value.


### GetSectionId

`func (o *UpdateWorkItemApiModel) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *UpdateWorkItemApiModel) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *UpdateWorkItemApiModel) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.


### GetName

`func (o *UpdateWorkItemApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWorkItemApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWorkItemApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateWorkItemApiModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateWorkItemApiModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateWorkItemApiModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateWorkItemApiModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateWorkItemApiModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateWorkItemApiModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDuration

`func (o *UpdateWorkItemApiModel) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UpdateWorkItemApiModel) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UpdateWorkItemApiModel) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetState

`func (o *UpdateWorkItemApiModel) GetState() WorkItemStateApiModel`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateWorkItemApiModel) GetStateOk() (*WorkItemStateApiModel, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateWorkItemApiModel) SetState(v WorkItemStateApiModel)`

SetState sets State field to given value.


### GetPriority

`func (o *UpdateWorkItemApiModel) GetPriority() WorkItemPriorityApiModel`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpdateWorkItemApiModel) GetPriorityOk() (*WorkItemPriorityApiModel, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpdateWorkItemApiModel) SetPriority(v WorkItemPriorityApiModel)`

SetPriority sets Priority field to given value.


### GetAttributes

`func (o *UpdateWorkItemApiModel) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateWorkItemApiModel) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateWorkItemApiModel) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UpdateWorkItemApiModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *UpdateWorkItemApiModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *UpdateWorkItemApiModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetTags

`func (o *UpdateWorkItemApiModel) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateWorkItemApiModel) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateWorkItemApiModel) SetTags(v []TagModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateWorkItemApiModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *UpdateWorkItemApiModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *UpdateWorkItemApiModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetPreconditionSteps

`func (o *UpdateWorkItemApiModel) GetPreconditionSteps() []UpdateStepApiModel`

GetPreconditionSteps returns the PreconditionSteps field if non-nil, zero value otherwise.

### GetPreconditionStepsOk

`func (o *UpdateWorkItemApiModel) GetPreconditionStepsOk() (*[]UpdateStepApiModel, bool)`

GetPreconditionStepsOk returns a tuple with the PreconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditionSteps

`func (o *UpdateWorkItemApiModel) SetPreconditionSteps(v []UpdateStepApiModel)`

SetPreconditionSteps sets PreconditionSteps field to given value.

### HasPreconditionSteps

`func (o *UpdateWorkItemApiModel) HasPreconditionSteps() bool`

HasPreconditionSteps returns a boolean if a field has been set.

### SetPreconditionStepsNil

`func (o *UpdateWorkItemApiModel) SetPreconditionStepsNil(b bool)`

 SetPreconditionStepsNil sets the value for PreconditionSteps to be an explicit nil

### UnsetPreconditionSteps
`func (o *UpdateWorkItemApiModel) UnsetPreconditionSteps()`

UnsetPreconditionSteps ensures that no value is present for PreconditionSteps, not even an explicit nil
### GetSteps

`func (o *UpdateWorkItemApiModel) GetSteps() []UpdateStepApiModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *UpdateWorkItemApiModel) GetStepsOk() (*[]UpdateStepApiModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *UpdateWorkItemApiModel) SetSteps(v []UpdateStepApiModel)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *UpdateWorkItemApiModel) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *UpdateWorkItemApiModel) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *UpdateWorkItemApiModel) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil
### GetPostconditionSteps

`func (o *UpdateWorkItemApiModel) GetPostconditionSteps() []UpdateStepApiModel`

GetPostconditionSteps returns the PostconditionSteps field if non-nil, zero value otherwise.

### GetPostconditionStepsOk

`func (o *UpdateWorkItemApiModel) GetPostconditionStepsOk() (*[]UpdateStepApiModel, bool)`

GetPostconditionStepsOk returns a tuple with the PostconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditionSteps

`func (o *UpdateWorkItemApiModel) SetPostconditionSteps(v []UpdateStepApiModel)`

SetPostconditionSteps sets PostconditionSteps field to given value.

### HasPostconditionSteps

`func (o *UpdateWorkItemApiModel) HasPostconditionSteps() bool`

HasPostconditionSteps returns a boolean if a field has been set.

### SetPostconditionStepsNil

`func (o *UpdateWorkItemApiModel) SetPostconditionStepsNil(b bool)`

 SetPostconditionStepsNil sets the value for PostconditionSteps to be an explicit nil

### UnsetPostconditionSteps
`func (o *UpdateWorkItemApiModel) UnsetPostconditionSteps()`

UnsetPostconditionSteps ensures that no value is present for PostconditionSteps, not even an explicit nil
### GetIterations

`func (o *UpdateWorkItemApiModel) GetIterations() []AssignIterationApiModel`

GetIterations returns the Iterations field if non-nil, zero value otherwise.

### GetIterationsOk

`func (o *UpdateWorkItemApiModel) GetIterationsOk() (*[]AssignIterationApiModel, bool)`

GetIterationsOk returns a tuple with the Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterations

`func (o *UpdateWorkItemApiModel) SetIterations(v []AssignIterationApiModel)`

SetIterations sets Iterations field to given value.

### HasIterations

`func (o *UpdateWorkItemApiModel) HasIterations() bool`

HasIterations returns a boolean if a field has been set.

### SetIterationsNil

`func (o *UpdateWorkItemApiModel) SetIterationsNil(b bool)`

 SetIterationsNil sets the value for Iterations to be an explicit nil

### UnsetIterations
`func (o *UpdateWorkItemApiModel) UnsetIterations()`

UnsetIterations ensures that no value is present for Iterations, not even an explicit nil
### GetAutoTests

`func (o *UpdateWorkItemApiModel) GetAutoTests() []AutoTestIdModel`

GetAutoTests returns the AutoTests field if non-nil, zero value otherwise.

### GetAutoTestsOk

`func (o *UpdateWorkItemApiModel) GetAutoTestsOk() (*[]AutoTestIdModel, bool)`

GetAutoTestsOk returns a tuple with the AutoTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTests

`func (o *UpdateWorkItemApiModel) SetAutoTests(v []AutoTestIdModel)`

SetAutoTests sets AutoTests field to given value.

### HasAutoTests

`func (o *UpdateWorkItemApiModel) HasAutoTests() bool`

HasAutoTests returns a boolean if a field has been set.

### SetAutoTestsNil

`func (o *UpdateWorkItemApiModel) SetAutoTestsNil(b bool)`

 SetAutoTestsNil sets the value for AutoTests to be an explicit nil

### UnsetAutoTests
`func (o *UpdateWorkItemApiModel) UnsetAutoTests()`

UnsetAutoTests ensures that no value is present for AutoTests, not even an explicit nil
### GetManualTestId

`func (o *UpdateWorkItemApiModel) GetManualTestId() string`

GetManualTestId returns the ManualTestId field if non-nil, zero value otherwise.

### GetManualTestIdOk

`func (o *UpdateWorkItemApiModel) GetManualTestIdOk() (*string, bool)`

GetManualTestIdOk returns a tuple with the ManualTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualTestId

`func (o *UpdateWorkItemApiModel) SetManualTestId(v string)`

SetManualTestId sets ManualTestId field to given value.

### HasManualTestId

`func (o *UpdateWorkItemApiModel) HasManualTestId() bool`

HasManualTestId returns a boolean if a field has been set.

### SetManualTestIdNil

`func (o *UpdateWorkItemApiModel) SetManualTestIdNil(b bool)`

 SetManualTestIdNil sets the value for ManualTestId to be an explicit nil

### UnsetManualTestId
`func (o *UpdateWorkItemApiModel) UnsetManualTestId()`

UnsetManualTestId ensures that no value is present for ManualTestId, not even an explicit nil
### GetAutoTestCases

`func (o *UpdateWorkItemApiModel) GetAutoTestCases() []AssignAutoTestCaseIdApiModel`

GetAutoTestCases returns the AutoTestCases field if non-nil, zero value otherwise.

### GetAutoTestCasesOk

`func (o *UpdateWorkItemApiModel) GetAutoTestCasesOk() (*[]AssignAutoTestCaseIdApiModel, bool)`

GetAutoTestCasesOk returns a tuple with the AutoTestCases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestCases

`func (o *UpdateWorkItemApiModel) SetAutoTestCases(v []AssignAutoTestCaseIdApiModel)`

SetAutoTestCases sets AutoTestCases field to given value.

### HasAutoTestCases

`func (o *UpdateWorkItemApiModel) HasAutoTestCases() bool`

HasAutoTestCases returns a boolean if a field has been set.

### SetAutoTestCasesNil

`func (o *UpdateWorkItemApiModel) SetAutoTestCasesNil(b bool)`

 SetAutoTestCasesNil sets the value for AutoTestCases to be an explicit nil

### UnsetAutoTestCases
`func (o *UpdateWorkItemApiModel) UnsetAutoTestCases()`

UnsetAutoTestCases ensures that no value is present for AutoTestCases, not even an explicit nil
### GetAttachments

`func (o *UpdateWorkItemApiModel) GetAttachments() []AssignAttachmentApiModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *UpdateWorkItemApiModel) GetAttachmentsOk() (*[]AssignAttachmentApiModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *UpdateWorkItemApiModel) SetAttachments(v []AssignAttachmentApiModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *UpdateWorkItemApiModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *UpdateWorkItemApiModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *UpdateWorkItemApiModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetLinks

`func (o *UpdateWorkItemApiModel) GetLinks() []UpdateLinkApiModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UpdateWorkItemApiModel) GetLinksOk() (*[]UpdateLinkApiModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UpdateWorkItemApiModel) SetLinks(v []UpdateLinkApiModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UpdateWorkItemApiModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *UpdateWorkItemApiModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *UpdateWorkItemApiModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetParameters

`func (o *UpdateWorkItemApiModel) GetParameters() []WorkItemParameterKeyApiModel`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *UpdateWorkItemApiModel) GetParametersOk() (*[]WorkItemParameterKeyApiModel, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *UpdateWorkItemApiModel) SetParameters(v []WorkItemParameterKeyApiModel)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *UpdateWorkItemApiModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *UpdateWorkItemApiModel) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *UpdateWorkItemApiModel) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetLayerName

`func (o *UpdateWorkItemApiModel) GetLayerName() string`

GetLayerName returns the LayerName field if non-nil, zero value otherwise.

### GetLayerNameOk

`func (o *UpdateWorkItemApiModel) GetLayerNameOk() (*string, bool)`

GetLayerNameOk returns a tuple with the LayerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerName

`func (o *UpdateWorkItemApiModel) SetLayerName(v string)`

SetLayerName sets LayerName field to given value.

### HasLayerName

`func (o *UpdateWorkItemApiModel) HasLayerName() bool`

HasLayerName returns a boolean if a field has been set.

### SetLayerNameNil

`func (o *UpdateWorkItemApiModel) SetLayerNameNil(b bool)`

 SetLayerNameNil sets the value for LayerName to be an explicit nil

### UnsetLayerName
`func (o *UpdateWorkItemApiModel) UnsetLayerName()`

UnsetLayerName ensures that no value is present for LayerName, not even an explicit nil
### GetLayerSource

`func (o *UpdateWorkItemApiModel) GetLayerSource() LayerSource`

GetLayerSource returns the LayerSource field if non-nil, zero value otherwise.

### GetLayerSourceOk

`func (o *UpdateWorkItemApiModel) GetLayerSourceOk() (*LayerSource, bool)`

GetLayerSourceOk returns a tuple with the LayerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerSource

`func (o *UpdateWorkItemApiModel) SetLayerSource(v LayerSource)`

SetLayerSource sets LayerSource field to given value.

### HasLayerSource

`func (o *UpdateWorkItemApiModel) HasLayerSource() bool`

HasLayerSource returns a boolean if a field has been set.

### SetLayerSourceNil

`func (o *UpdateWorkItemApiModel) SetLayerSourceNil(b bool)`

 SetLayerSourceNil sets the value for LayerSource to be an explicit nil

### UnsetLayerSource
`func (o *UpdateWorkItemApiModel) UnsetLayerSource()`

UnsetLayerSource ensures that no value is present for LayerSource, not even an explicit nil
### GetResetLayer

`func (o *UpdateWorkItemApiModel) GetResetLayer() bool`

GetResetLayer returns the ResetLayer field if non-nil, zero value otherwise.

### GetResetLayerOk

`func (o *UpdateWorkItemApiModel) GetResetLayerOk() (*bool, bool)`

GetResetLayerOk returns a tuple with the ResetLayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetLayer

`func (o *UpdateWorkItemApiModel) SetResetLayer(v bool)`

SetResetLayer sets ResetLayer field to given value.

### HasResetLayer

`func (o *UpdateWorkItemApiModel) HasResetLayer() bool`

HasResetLayer returns a boolean if a field has been set.

### SetResetLayerNil

`func (o *UpdateWorkItemApiModel) SetResetLayerNil(b bool)`

 SetResetLayerNil sets the value for ResetLayer to be an explicit nil

### UnsetResetLayer
`func (o *UpdateWorkItemApiModel) UnsetResetLayer()`

UnsetResetLayer ensures that no value is present for ResetLayer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


