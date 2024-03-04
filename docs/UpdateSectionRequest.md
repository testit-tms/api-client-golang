# UpdateSectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**ProjectId** | **string** |  | 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**PreconditionSteps** | Pointer to [**[]StepPutModel**](StepPutModel.md) |  | [optional] 
**PostconditionSteps** | Pointer to [**[]StepPutModel**](StepPutModel.md) |  | [optional] 
**Attachments** | [**[]AttachmentPutModel**](AttachmentPutModel.md) |  | 

## Methods

### NewUpdateSectionRequest

`func NewUpdateSectionRequest(id string, name string, projectId string, attachments []AttachmentPutModel, ) *UpdateSectionRequest`

NewUpdateSectionRequest instantiates a new UpdateSectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSectionRequestWithDefaults

`func NewUpdateSectionRequestWithDefaults() *UpdateSectionRequest`

NewUpdateSectionRequestWithDefaults instantiates a new UpdateSectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateSectionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateSectionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateSectionRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateSectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSectionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *UpdateSectionRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateSectionRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateSectionRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetParentId

`func (o *UpdateSectionRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *UpdateSectionRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *UpdateSectionRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *UpdateSectionRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *UpdateSectionRequest) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *UpdateSectionRequest) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetPreconditionSteps

`func (o *UpdateSectionRequest) GetPreconditionSteps() []StepPutModel`

GetPreconditionSteps returns the PreconditionSteps field if non-nil, zero value otherwise.

### GetPreconditionStepsOk

`func (o *UpdateSectionRequest) GetPreconditionStepsOk() (*[]StepPutModel, bool)`

GetPreconditionStepsOk returns a tuple with the PreconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditionSteps

`func (o *UpdateSectionRequest) SetPreconditionSteps(v []StepPutModel)`

SetPreconditionSteps sets PreconditionSteps field to given value.

### HasPreconditionSteps

`func (o *UpdateSectionRequest) HasPreconditionSteps() bool`

HasPreconditionSteps returns a boolean if a field has been set.

### SetPreconditionStepsNil

`func (o *UpdateSectionRequest) SetPreconditionStepsNil(b bool)`

 SetPreconditionStepsNil sets the value for PreconditionSteps to be an explicit nil

### UnsetPreconditionSteps
`func (o *UpdateSectionRequest) UnsetPreconditionSteps()`

UnsetPreconditionSteps ensures that no value is present for PreconditionSteps, not even an explicit nil
### GetPostconditionSteps

`func (o *UpdateSectionRequest) GetPostconditionSteps() []StepPutModel`

GetPostconditionSteps returns the PostconditionSteps field if non-nil, zero value otherwise.

### GetPostconditionStepsOk

`func (o *UpdateSectionRequest) GetPostconditionStepsOk() (*[]StepPutModel, bool)`

GetPostconditionStepsOk returns a tuple with the PostconditionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostconditionSteps

`func (o *UpdateSectionRequest) SetPostconditionSteps(v []StepPutModel)`

SetPostconditionSteps sets PostconditionSteps field to given value.

### HasPostconditionSteps

`func (o *UpdateSectionRequest) HasPostconditionSteps() bool`

HasPostconditionSteps returns a boolean if a field has been set.

### SetPostconditionStepsNil

`func (o *UpdateSectionRequest) SetPostconditionStepsNil(b bool)`

 SetPostconditionStepsNil sets the value for PostconditionSteps to be an explicit nil

### UnsetPostconditionSteps
`func (o *UpdateSectionRequest) UnsetPostconditionSteps()`

UnsetPostconditionSteps ensures that no value is present for PostconditionSteps, not even an explicit nil
### GetAttachments

`func (o *UpdateSectionRequest) GetAttachments() []AttachmentPutModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *UpdateSectionRequest) GetAttachmentsOk() (*[]AttachmentPutModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *UpdateSectionRequest) SetAttachments(v []AttachmentPutModel)`

SetAttachments sets Attachments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


