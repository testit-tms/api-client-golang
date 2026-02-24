# CreateWorkItemCommentApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkItemId** | **string** | ID of work item to comment | 
**Text** | **string** | Text of the comment | 

## Methods

### NewCreateWorkItemCommentApiModel

`func NewCreateWorkItemCommentApiModel(workItemId string, text string, ) *CreateWorkItemCommentApiModel`

NewCreateWorkItemCommentApiModel instantiates a new CreateWorkItemCommentApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkItemCommentApiModelWithDefaults

`func NewCreateWorkItemCommentApiModelWithDefaults() *CreateWorkItemCommentApiModel`

NewCreateWorkItemCommentApiModelWithDefaults instantiates a new CreateWorkItemCommentApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkItemId

`func (o *CreateWorkItemCommentApiModel) GetWorkItemId() string`

GetWorkItemId returns the WorkItemId field if non-nil, zero value otherwise.

### GetWorkItemIdOk

`func (o *CreateWorkItemCommentApiModel) GetWorkItemIdOk() (*string, bool)`

GetWorkItemIdOk returns a tuple with the WorkItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemId

`func (o *CreateWorkItemCommentApiModel) SetWorkItemId(v string)`

SetWorkItemId sets WorkItemId field to given value.


### GetText

`func (o *CreateWorkItemCommentApiModel) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CreateWorkItemCommentApiModel) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CreateWorkItemCommentApiModel) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


