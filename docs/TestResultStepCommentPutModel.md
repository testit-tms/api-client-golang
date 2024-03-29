# TestResultStepCommentPutModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Text** | **string** |  | 
**StepId** | **string** |  | 
**ParentStepId** | Pointer to **NullableString** |  | [optional] 
**Attachments** | [**[]AttachmentPutModel**](AttachmentPutModel.md) |  | 

## Methods

### NewTestResultStepCommentPutModel

`func NewTestResultStepCommentPutModel(id string, text string, stepId string, attachments []AttachmentPutModel, ) *TestResultStepCommentPutModel`

NewTestResultStepCommentPutModel instantiates a new TestResultStepCommentPutModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultStepCommentPutModelWithDefaults

`func NewTestResultStepCommentPutModelWithDefaults() *TestResultStepCommentPutModel`

NewTestResultStepCommentPutModelWithDefaults instantiates a new TestResultStepCommentPutModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestResultStepCommentPutModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestResultStepCommentPutModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestResultStepCommentPutModel) SetId(v string)`

SetId sets Id field to given value.


### GetText

`func (o *TestResultStepCommentPutModel) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TestResultStepCommentPutModel) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TestResultStepCommentPutModel) SetText(v string)`

SetText sets Text field to given value.


### GetStepId

`func (o *TestResultStepCommentPutModel) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *TestResultStepCommentPutModel) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *TestResultStepCommentPutModel) SetStepId(v string)`

SetStepId sets StepId field to given value.


### GetParentStepId

`func (o *TestResultStepCommentPutModel) GetParentStepId() string`

GetParentStepId returns the ParentStepId field if non-nil, zero value otherwise.

### GetParentStepIdOk

`func (o *TestResultStepCommentPutModel) GetParentStepIdOk() (*string, bool)`

GetParentStepIdOk returns a tuple with the ParentStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStepId

`func (o *TestResultStepCommentPutModel) SetParentStepId(v string)`

SetParentStepId sets ParentStepId field to given value.

### HasParentStepId

`func (o *TestResultStepCommentPutModel) HasParentStepId() bool`

HasParentStepId returns a boolean if a field has been set.

### SetParentStepIdNil

`func (o *TestResultStepCommentPutModel) SetParentStepIdNil(b bool)`

 SetParentStepIdNil sets the value for ParentStepId to be an explicit nil

### UnsetParentStepId
`func (o *TestResultStepCommentPutModel) UnsetParentStepId()`

UnsetParentStepId ensures that no value is present for ParentStepId, not even an explicit nil
### GetAttachments

`func (o *TestResultStepCommentPutModel) GetAttachments() []AttachmentPutModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestResultStepCommentPutModel) GetAttachmentsOk() (*[]AttachmentPutModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestResultStepCommentPutModel) SetAttachments(v []AttachmentPutModel)`

SetAttachments sets Attachments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


