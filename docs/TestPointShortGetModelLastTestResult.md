# TestPointShortGetModelLastTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TestRunId** | Pointer to **string** |  | [optional] 
**AutoTestId** | Pointer to **NullableString** |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**WorkItemVersionId** | Pointer to **string** |  | [optional] 
**Attachments** | Pointer to [**[]AttachmentModel**](AttachmentModel.md) |  | [optional] 

## Methods

### NewTestPointShortGetModelLastTestResult

`func NewTestPointShortGetModelLastTestResult() *TestPointShortGetModelLastTestResult`

NewTestPointShortGetModelLastTestResult instantiates a new TestPointShortGetModelLastTestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPointShortGetModelLastTestResultWithDefaults

`func NewTestPointShortGetModelLastTestResultWithDefaults() *TestPointShortGetModelLastTestResult`

NewTestPointShortGetModelLastTestResultWithDefaults instantiates a new TestPointShortGetModelLastTestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestPointShortGetModelLastTestResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestPointShortGetModelLastTestResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestPointShortGetModelLastTestResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestPointShortGetModelLastTestResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTestRunId

`func (o *TestPointShortGetModelLastTestResult) GetTestRunId() string`

GetTestRunId returns the TestRunId field if non-nil, zero value otherwise.

### GetTestRunIdOk

`func (o *TestPointShortGetModelLastTestResult) GetTestRunIdOk() (*string, bool)`

GetTestRunIdOk returns a tuple with the TestRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRunId

`func (o *TestPointShortGetModelLastTestResult) SetTestRunId(v string)`

SetTestRunId sets TestRunId field to given value.

### HasTestRunId

`func (o *TestPointShortGetModelLastTestResult) HasTestRunId() bool`

HasTestRunId returns a boolean if a field has been set.

### GetAutoTestId

`func (o *TestPointShortGetModelLastTestResult) GetAutoTestId() string`

GetAutoTestId returns the AutoTestId field if non-nil, zero value otherwise.

### GetAutoTestIdOk

`func (o *TestPointShortGetModelLastTestResult) GetAutoTestIdOk() (*string, bool)`

GetAutoTestIdOk returns a tuple with the AutoTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestId

`func (o *TestPointShortGetModelLastTestResult) SetAutoTestId(v string)`

SetAutoTestId sets AutoTestId field to given value.

### HasAutoTestId

`func (o *TestPointShortGetModelLastTestResult) HasAutoTestId() bool`

HasAutoTestId returns a boolean if a field has been set.

### SetAutoTestIdNil

`func (o *TestPointShortGetModelLastTestResult) SetAutoTestIdNil(b bool)`

 SetAutoTestIdNil sets the value for AutoTestId to be an explicit nil

### UnsetAutoTestId
`func (o *TestPointShortGetModelLastTestResult) UnsetAutoTestId()`

UnsetAutoTestId ensures that no value is present for AutoTestId, not even an explicit nil
### GetComment

`func (o *TestPointShortGetModelLastTestResult) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TestPointShortGetModelLastTestResult) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TestPointShortGetModelLastTestResult) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TestPointShortGetModelLastTestResult) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *TestPointShortGetModelLastTestResult) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *TestPointShortGetModelLastTestResult) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetLinks

`func (o *TestPointShortGetModelLastTestResult) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TestPointShortGetModelLastTestResult) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TestPointShortGetModelLastTestResult) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TestPointShortGetModelLastTestResult) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TestPointShortGetModelLastTestResult) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TestPointShortGetModelLastTestResult) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetWorkItemVersionId

`func (o *TestPointShortGetModelLastTestResult) GetWorkItemVersionId() string`

GetWorkItemVersionId returns the WorkItemVersionId field if non-nil, zero value otherwise.

### GetWorkItemVersionIdOk

`func (o *TestPointShortGetModelLastTestResult) GetWorkItemVersionIdOk() (*string, bool)`

GetWorkItemVersionIdOk returns a tuple with the WorkItemVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionId

`func (o *TestPointShortGetModelLastTestResult) SetWorkItemVersionId(v string)`

SetWorkItemVersionId sets WorkItemVersionId field to given value.

### HasWorkItemVersionId

`func (o *TestPointShortGetModelLastTestResult) HasWorkItemVersionId() bool`

HasWorkItemVersionId returns a boolean if a field has been set.

### GetAttachments

`func (o *TestPointShortGetModelLastTestResult) GetAttachments() []AttachmentModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestPointShortGetModelLastTestResult) GetAttachmentsOk() (*[]AttachmentModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestPointShortGetModelLastTestResult) SetAttachments(v []AttachmentModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestPointShortGetModelLastTestResult) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *TestPointShortGetModelLastTestResult) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *TestPointShortGetModelLastTestResult) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


