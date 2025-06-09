# LastTestResultApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TestRunId** | **string** |  | 
**AutoTestId** | Pointer to **NullableString** |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**Links** | [**[]LinkApiResult**](LinkApiResult.md) |  | 
**WorkItemVersionId** | Pointer to **NullableString** |  | [optional] 
**Attachments** | [**[]AttachmentApiResult**](AttachmentApiResult.md) |  | 

## Methods

### NewLastTestResultApiResult

`func NewLastTestResultApiResult(id string, testRunId string, links []LinkApiResult, attachments []AttachmentApiResult, ) *LastTestResultApiResult`

NewLastTestResultApiResult instantiates a new LastTestResultApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastTestResultApiResultWithDefaults

`func NewLastTestResultApiResultWithDefaults() *LastTestResultApiResult`

NewLastTestResultApiResultWithDefaults instantiates a new LastTestResultApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LastTestResultApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LastTestResultApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LastTestResultApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetTestRunId

`func (o *LastTestResultApiResult) GetTestRunId() string`

GetTestRunId returns the TestRunId field if non-nil, zero value otherwise.

### GetTestRunIdOk

`func (o *LastTestResultApiResult) GetTestRunIdOk() (*string, bool)`

GetTestRunIdOk returns a tuple with the TestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunId

`func (o *LastTestResultApiResult) SetTestRunId(v string)`

SetTestRunId sets TestRunId field to given value.


### GetAutoTestId

`func (o *LastTestResultApiResult) GetAutoTestId() string`

GetAutoTestId returns the AutoTestId field if non-nil, zero value otherwise.

### GetAutoTestIdOk

`func (o *LastTestResultApiResult) GetAutoTestIdOk() (*string, bool)`

GetAutoTestIdOk returns a tuple with the AutoTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestId

`func (o *LastTestResultApiResult) SetAutoTestId(v string)`

SetAutoTestId sets AutoTestId field to given value.

### HasAutoTestId

`func (o *LastTestResultApiResult) HasAutoTestId() bool`

HasAutoTestId returns a boolean if a field has been set.

### SetAutoTestIdNil

`func (o *LastTestResultApiResult) SetAutoTestIdNil(b bool)`

 SetAutoTestIdNil sets the value for AutoTestId to be an explicit nil

### UnsetAutoTestId
`func (o *LastTestResultApiResult) UnsetAutoTestId()`

UnsetAutoTestId ensures that no value is present for AutoTestId, not even an explicit nil
### GetComment

`func (o *LastTestResultApiResult) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *LastTestResultApiResult) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *LastTestResultApiResult) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *LastTestResultApiResult) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *LastTestResultApiResult) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *LastTestResultApiResult) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetLinks

`func (o *LastTestResultApiResult) GetLinks() []LinkApiResult`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LastTestResultApiResult) GetLinksOk() (*[]LinkApiResult, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LastTestResultApiResult) SetLinks(v []LinkApiResult)`

SetLinks sets Links field to given value.


### GetWorkItemVersionId

`func (o *LastTestResultApiResult) GetWorkItemVersionId() string`

GetWorkItemVersionId returns the WorkItemVersionId field if non-nil, zero value otherwise.

### GetWorkItemVersionIdOk

`func (o *LastTestResultApiResult) GetWorkItemVersionIdOk() (*string, bool)`

GetWorkItemVersionIdOk returns a tuple with the WorkItemVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionId

`func (o *LastTestResultApiResult) SetWorkItemVersionId(v string)`

SetWorkItemVersionId sets WorkItemVersionId field to given value.

### HasWorkItemVersionId

`func (o *LastTestResultApiResult) HasWorkItemVersionId() bool`

HasWorkItemVersionId returns a boolean if a field has been set.

### SetWorkItemVersionIdNil

`func (o *LastTestResultApiResult) SetWorkItemVersionIdNil(b bool)`

 SetWorkItemVersionIdNil sets the value for WorkItemVersionId to be an explicit nil

### UnsetWorkItemVersionId
`func (o *LastTestResultApiResult) UnsetWorkItemVersionId()`

UnsetWorkItemVersionId ensures that no value is present for WorkItemVersionId, not even an explicit nil
### GetAttachments

`func (o *LastTestResultApiResult) GetAttachments() []AttachmentApiResult`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *LastTestResultApiResult) GetAttachmentsOk() (*[]AttachmentApiResult, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *LastTestResultApiResult) SetAttachments(v []AttachmentApiResult)`

SetAttachments sets Attachments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


