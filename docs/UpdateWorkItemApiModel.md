# UpdateWorkItemApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Workitem internal identifier | 
**SectionId** | **string** | Internal identifier of section where workitem is located | 
**Description** | Pointer to **NullableString** | Workitem description | [optional] 
**State** | [**WorkItemStates**](WorkItemStates.md) |  | 
**Priority** | [**WorkItemPriorityModel**](WorkItemPriorityModel.md) |  | 
**SourceType** | Pointer to [**NullableWorkItemSourceTypeModel**](WorkItemSourceTypeModel.md) |  | [optional] 
**Steps** | [**[]UpdateStepApiModel**](UpdateStepApiModel.md) | Collection of workitem steps | 
**PreconditionSteps** | [**[]UpdateStepApiModel**](UpdateStepApiModel.md) | Collection of workitem precondtion steps | 
**PostconditionSteps** | [**[]UpdateStepApiModel**](UpdateStepApiModel.md) | Collection of workitem postcondition steps | 
**Duration** | **int64** | Workitem duration in milliseconds | 
**Attributes** | **map[string]interface{}** | Key value pair of custom workitem attributes | 
**Tags** | [**[]TagModel**](TagModel.md) | Collection of workitem tags | 
**Links** | [**[]UpdateLinkApiModel**](UpdateLinkApiModel.md) | Collection of workitem links | 
**Name** | **string** | Workitem name | 
**Attachments** | [**[]AssignAttachmentApiModel**](AssignAttachmentApiModel.md) |  | 
**Iterations** | Pointer to [**[]AssignIterationApiModel**](AssignIterationApiModel.md) | Collection of parameter id sets | [optional] 
**AutoTests** | Pointer to [**[]AutoTestIdModel**](AutoTestIdModel.md) | Collection of autotest internal ids | [optional] 

## Methods

### NewUpdateWorkItemApiModel

`func NewUpdateWorkItemApiModel(id string, sectionId string, state WorkItemStates, priority WorkItemPriorityModel, steps []UpdateStepApiModel, preconditionSteps []UpdateStepApiModel, postconditionSteps []UpdateStepApiModel, duration int64, attributes map[string]interface{}, tags []TagModel, links []UpdateLinkApiModel, name string, attachments []AssignAttachmentApiModel, ) *UpdateWorkItemApiModel`

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
### GetState

`func (o *UpdateWorkItemApiModel) GetState() WorkItemStates`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateWorkItemApiModel) GetStateOk() (*WorkItemStates, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateWorkItemApiModel) SetState(v WorkItemStates)`

SetState sets State field to given value.


### GetPriority

`func (o *UpdateWorkItemApiModel) GetPriority() WorkItemPriorityModel`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpdateWorkItemApiModel) GetPriorityOk() (*WorkItemPriorityModel, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpdateWorkItemApiModel) SetPriority(v WorkItemPriorityModel)`

SetPriority sets Priority field to given value.


### GetSourceType

`func (o *UpdateWorkItemApiModel) GetSourceType() WorkItemSourceTypeModel`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *UpdateWorkItemApiModel) GetSourceTypeOk() (*WorkItemSourceTypeModel, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *UpdateWorkItemApiModel) SetSourceType(v WorkItemSourceTypeModel)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *UpdateWorkItemApiModel) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *UpdateWorkItemApiModel) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *UpdateWorkItemApiModel) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


