# UpdateWorkItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | [**[]AttachmentPutModel**](AttachmentPutModel.md) |  | 
**Iterations** | Pointer to [**[]IterationPutModel**](IterationPutModel.md) |  | [optional] 
**AutoTests** | Pointer to [**[]AutoTestIdModel**](AutoTestIdModel.md) |  | [optional] 
**Id** | **string** |  | 
**SectionId** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**State** | [**WorkItemStates**](WorkItemStates.md) |  | 
**Priority** | [**WorkItemPriorityModel**](WorkItemPriorityModel.md) |  | 
**Steps** | [**[]StepPutModel**](StepPutModel.md) |  | 
**PreconditionSteps** | [**[]StepPutModel**](StepPutModel.md) |  | 
**PostconditionSteps** | [**[]StepPutModel**](StepPutModel.md) |  | 
**Duration** | **int32** |  | 
**Attributes** | **map[string]interface{}** |  | 
**Tags** | [**[]TagShortModel**](TagShortModel.md) |  | 
**Links** | [**[]LinkPutModel**](LinkPutModel.md) |  | 
**Name** | **string** |  | 

## Methods

### NewUpdateWorkItemRequest

`func NewUpdateWorkItemRequest(attachments []AttachmentPutModel, id string, sectionId string, state WorkItemStates, priority WorkItemPriorityModel, steps []StepPutModel, preconditionSteps []StepPutModel, postconditionSteps []StepPutModel, duration int32, attributes map[string]interface{}, tags []TagShortModel, links []LinkPutModel, name string, ) *UpdateWorkItemRequest`

NewUpdateWorkItemRequest instantiates a new UpdateWorkItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWorkItemRequestWithDefaults

`func NewUpdateWorkItemRequestWithDefaults() *UpdateWorkItemRequest`

NewUpdateWorkItemRequestWithDefaults instantiates a new UpdateWorkItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *UpdateWorkItemRequest) GetAttachments() []AttachmentPutModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *UpdateWorkItemRequest) GetAttachmentsOk() (*[]AttachmentPutModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *UpdateWorkItemRequest) SetAttachments(v []AttachmentPutModel)`

SetAttachments sets Attachments field to given value.


### GetIterations

`func (o *UpdateWorkItemRequest) GetIterations() []IterationPutModel`

GetIterations returns the Iterations field if non-nil, zero value otherwise.

### GetIterationsOk

`func (o *UpdateWorkItemRequest) GetIterationsOk() (*[]IterationPutModel, bool)`

GetIterationsOk returns a tuple with the Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterations

`func (o *UpdateWorkItemRequest) SetIterations(v []IterationPutModel)`

SetIterations sets Iterations field to given value.

### HasIterations

`func (o *UpdateWorkItemRequest) HasIterations() bool`

HasIterations returns a boolean if a field has been set.

### SetIterationsNil

`func (o *UpdateWorkItemRequest) SetIterationsNil(b bool)`

 SetIterationsNil sets the value for Iterations to be an explicit nil

### UnsetIterations
`func (o *UpdateWorkItemRequest) UnsetIterations()`

UnsetIterations ensures that no value is present for Iterations, not even an explicit nil
### GetAutoTests

`func (o *UpdateWorkItemRequest) GetAutoTests() []AutoTestIdModel`

GetAutoTests returns the AutoTests field if non-nil, zero value otherwise.

### GetAutoTestsOk

`func (o *UpdateWorkItemRequest) GetAutoTestsOk() (*[]AutoTestIdModel, bool)`

GetAutoTestsOk returns a tuple with the AutoTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTests

`func (o *UpdateWorkItemRequest) SetAutoTests(v []AutoTestIdModel)`

SetAutoTests sets AutoTests field to given value.

### HasAutoTests

`func (o *UpdateWorkItemRequest) HasAutoTests() bool`

HasAutoTests returns a boolean if a field has been set.

### SetAutoTestsNil

`func (o *UpdateWorkItemRequest) SetAutoTestsNil(b bool)`

 SetAutoTestsNil sets the value for AutoTests to be an explicit nil

### UnsetAutoTests
`func (o *UpdateWorkItemRequest) UnsetAutoTests()`

UnsetAutoTests ensures that no value is present for AutoTests, not even an explicit nil
### GetId

`func (o *UpdateWorkItemRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateWorkItemRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateWorkItemRequest) SetId(v string)`

SetId sets Id field to given value.


### GetSectionId

`func (o *UpdateWorkItemRequest) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *UpdateWorkItemRequest) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *UpdateWorkItemRequest) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.


### GetDescription

`func (o *UpdateWorkItemRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateWorkItemRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateWorkItemRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateWorkItemRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateWorkItemRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateWorkItemRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetState

`func (o *UpdateWorkItemRequest) GetState() WorkItemStates`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateWorkItemRequest) GetStateOk() (*WorkItemStates, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateWorkItemRequest) SetState(v WorkItemStates)`

SetState sets State field to given value.


### GetPriority

`func (o *UpdateWorkItemRequest) GetPriority() WorkItemPriorityModel`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpdateWorkItemRequest) GetPriorityOk() (*WorkItemPriorityModel, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpdateWorkItemRequest) SetPriority(v WorkItemPriorityModel)`

SetPriority sets Priority field to given value.


### GetSteps

`func (o *UpdateWorkItemRequest) GetSteps() []StepPutModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *UpdateWorkItemRequest) GetStepsOk() (*[]StepPutModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *UpdateWorkItemRequest) SetSteps(v []StepPutModel)`

SetSteps sets Steps field to given value.


### GetPreconditionSteps

`func (o *UpdateWorkItemRequest) GetPreconditionSteps() []StepPutModel`

GetPreconditionSteps returns the PreconditionSteps field if non-nil, zero value otherwise.

### GetPreconditionStepsOk

`func (o *UpdateWorkItemRequest) GetPreconditionStepsOk() (*[]StepPutModel, bool)`

GetPreconditionStepsOk returns a tuple with the PreconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditionSteps

`func (o *UpdateWorkItemRequest) SetPreconditionSteps(v []StepPutModel)`

SetPreconditionSteps sets PreconditionSteps field to given value.


### GetPostconditionSteps

`func (o *UpdateWorkItemRequest) GetPostconditionSteps() []StepPutModel`

GetPostconditionSteps returns the PostconditionSteps field if non-nil, zero value otherwise.

### GetPostconditionStepsOk

`func (o *UpdateWorkItemRequest) GetPostconditionStepsOk() (*[]StepPutModel, bool)`

GetPostconditionStepsOk returns a tuple with the PostconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditionSteps

`func (o *UpdateWorkItemRequest) SetPostconditionSteps(v []StepPutModel)`

SetPostconditionSteps sets PostconditionSteps field to given value.


### GetDuration

`func (o *UpdateWorkItemRequest) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UpdateWorkItemRequest) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UpdateWorkItemRequest) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetAttributes

`func (o *UpdateWorkItemRequest) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateWorkItemRequest) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateWorkItemRequest) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetTags

`func (o *UpdateWorkItemRequest) GetTags() []TagShortModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateWorkItemRequest) GetTagsOk() (*[]TagShortModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateWorkItemRequest) SetTags(v []TagShortModel)`

SetTags sets Tags field to given value.


### GetLinks

`func (o *UpdateWorkItemRequest) GetLinks() []LinkPutModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UpdateWorkItemRequest) GetLinksOk() (*[]LinkPutModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UpdateWorkItemRequest) SetLinks(v []LinkPutModel)`

SetLinks sets Links field to given value.


### GetName

`func (o *UpdateWorkItemRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWorkItemRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWorkItemRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


