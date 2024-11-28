# TestResultStepCommentUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Entity unique identifier | 
**Text** | **string** |  | 
**StepId** | **string** |  | 
**ParentStepId** | Pointer to **NullableString** |  | [optional] 
**Attachments** | [**[]AttachmentUpdateRequest**](AttachmentUpdateRequest.md) |  | 

## Methods

### NewTestResultStepCommentUpdateRequest

`func NewTestResultStepCommentUpdateRequest(id string, text string, stepId string, attachments []AttachmentUpdateRequest, ) *TestResultStepCommentUpdateRequest`

NewTestResultStepCommentUpdateRequest instantiates a new TestResultStepCommentUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultStepCommentUpdateRequestWithDefaults

`func NewTestResultStepCommentUpdateRequestWithDefaults() *TestResultStepCommentUpdateRequest`

NewTestResultStepCommentUpdateRequestWithDefaults instantiates a new TestResultStepCommentUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestResultStepCommentUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestResultStepCommentUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestResultStepCommentUpdateRequest) SetId(v string)`

SetId sets Id field to given value.


### GetText

`func (o *TestResultStepCommentUpdateRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TestResultStepCommentUpdateRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TestResultStepCommentUpdateRequest) SetText(v string)`

SetText sets Text field to given value.


### GetStepId

`func (o *TestResultStepCommentUpdateRequest) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *TestResultStepCommentUpdateRequest) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *TestResultStepCommentUpdateRequest) SetStepId(v string)`

SetStepId sets StepId field to given value.


### GetParentStepId

`func (o *TestResultStepCommentUpdateRequest) GetParentStepId() string`

GetParentStepId returns the ParentStepId field if non-nil, zero value otherwise.

### GetParentStepIdOk

`func (o *TestResultStepCommentUpdateRequest) GetParentStepIdOk() (*string, bool)`

GetParentStepIdOk returns a tuple with the ParentStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStepId

`func (o *TestResultStepCommentUpdateRequest) SetParentStepId(v string)`

SetParentStepId sets ParentStepId field to given value.

### HasParentStepId

`func (o *TestResultStepCommentUpdateRequest) HasParentStepId() bool`

HasParentStepId returns a boolean if a field has been set.

### SetParentStepIdNil

`func (o *TestResultStepCommentUpdateRequest) SetParentStepIdNil(b bool)`

 SetParentStepIdNil sets the value for ParentStepId to be an explicit nil

### UnsetParentStepId
`func (o *TestResultStepCommentUpdateRequest) UnsetParentStepId()`

UnsetParentStepId ensures that no value is present for ParentStepId, not even an explicit nil
### GetAttachments

`func (o *TestResultStepCommentUpdateRequest) GetAttachments() []AttachmentUpdateRequest`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestResultStepCommentUpdateRequest) GetAttachmentsOk() (*[]AttachmentUpdateRequest, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestResultStepCommentUpdateRequest) SetAttachments(v []AttachmentUpdateRequest)`

SetAttachments sets Attachments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


