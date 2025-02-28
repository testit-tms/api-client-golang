# StepCommentApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Text** | Pointer to **NullableString** |  | [optional] 
**StepId** | **string** |  | 
**ParentStepId** | Pointer to **NullableString** |  | [optional] 
**Attachments** | [**[]AttachmentApiResult**](AttachmentApiResult.md) |  | 
**TestResultId** | **string** |  | 
**CreatedById** | **string** |  | 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**CreatedDate** | **time.Time** |  | 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewStepCommentApiModel

`func NewStepCommentApiModel(id string, stepId string, attachments []AttachmentApiResult, testResultId string, createdById string, createdDate time.Time, ) *StepCommentApiModel`

NewStepCommentApiModel instantiates a new StepCommentApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStepCommentApiModelWithDefaults

`func NewStepCommentApiModelWithDefaults() *StepCommentApiModel`

NewStepCommentApiModelWithDefaults instantiates a new StepCommentApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StepCommentApiModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StepCommentApiModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StepCommentApiModel) SetId(v string)`

SetId sets Id field to given value.


### GetText

`func (o *StepCommentApiModel) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *StepCommentApiModel) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *StepCommentApiModel) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *StepCommentApiModel) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *StepCommentApiModel) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *StepCommentApiModel) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetStepId

`func (o *StepCommentApiModel) GetStepId() string`

GetStepId returns the StepId field if non-nil, zero value otherwise.

### GetStepIdOk

`func (o *StepCommentApiModel) GetStepIdOk() (*string, bool)`

GetStepIdOk returns a tuple with the StepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepId

`func (o *StepCommentApiModel) SetStepId(v string)`

SetStepId sets StepId field to given value.


### GetParentStepId

`func (o *StepCommentApiModel) GetParentStepId() string`

GetParentStepId returns the ParentStepId field if non-nil, zero value otherwise.

### GetParentStepIdOk

`func (o *StepCommentApiModel) GetParentStepIdOk() (*string, bool)`

GetParentStepIdOk returns a tuple with the ParentStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStepId

`func (o *StepCommentApiModel) SetParentStepId(v string)`

SetParentStepId sets ParentStepId field to given value.

### HasParentStepId

`func (o *StepCommentApiModel) HasParentStepId() bool`

HasParentStepId returns a boolean if a field has been set.

### SetParentStepIdNil

`func (o *StepCommentApiModel) SetParentStepIdNil(b bool)`

 SetParentStepIdNil sets the value for ParentStepId to be an explicit nil

### UnsetParentStepId
`func (o *StepCommentApiModel) UnsetParentStepId()`

UnsetParentStepId ensures that no value is present for ParentStepId, not even an explicit nil
### GetAttachments

`func (o *StepCommentApiModel) GetAttachments() []AttachmentApiResult`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *StepCommentApiModel) GetAttachmentsOk() (*[]AttachmentApiResult, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *StepCommentApiModel) SetAttachments(v []AttachmentApiResult)`

SetAttachments sets Attachments field to given value.


### GetTestResultId

`func (o *StepCommentApiModel) GetTestResultId() string`

GetTestResultId returns the TestResultId field if non-nil, zero value otherwise.

### GetTestResultIdOk

`func (o *StepCommentApiModel) GetTestResultIdOk() (*string, bool)`

GetTestResultIdOk returns a tuple with the TestResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResultId

`func (o *StepCommentApiModel) SetTestResultId(v string)`

SetTestResultId sets TestResultId field to given value.


### GetCreatedById

`func (o *StepCommentApiModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *StepCommentApiModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *StepCommentApiModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedById

`func (o *StepCommentApiModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *StepCommentApiModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *StepCommentApiModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *StepCommentApiModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *StepCommentApiModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *StepCommentApiModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetCreatedDate

`func (o *StepCommentApiModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *StepCommentApiModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *StepCommentApiModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *StepCommentApiModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *StepCommentApiModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *StepCommentApiModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *StepCommentApiModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *StepCommentApiModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *StepCommentApiModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


