# WorkItemCommentApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the comment | 
**Text** | **string** | Text of the comment | 
**CreatedById** | **string** | ID of user created comment | 
**ModifiedById** | **string** | ID of user modified comment | 
**CreatedDate** | **time.Time** | Comment created date | 
**ModifiedDate** | **time.Time** | Comment modified date | 

## Methods

### NewWorkItemCommentApiResult

`func NewWorkItemCommentApiResult(id string, text string, createdById string, modifiedById string, createdDate time.Time, modifiedDate time.Time, ) *WorkItemCommentApiResult`

NewWorkItemCommentApiResult instantiates a new WorkItemCommentApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemCommentApiResultWithDefaults

`func NewWorkItemCommentApiResultWithDefaults() *WorkItemCommentApiResult`

NewWorkItemCommentApiResultWithDefaults instantiates a new WorkItemCommentApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkItemCommentApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkItemCommentApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkItemCommentApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetText

`func (o *WorkItemCommentApiResult) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WorkItemCommentApiResult) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WorkItemCommentApiResult) SetText(v string)`

SetText sets Text field to given value.


### GetCreatedById

`func (o *WorkItemCommentApiResult) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WorkItemCommentApiResult) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WorkItemCommentApiResult) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedById

`func (o *WorkItemCommentApiResult) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *WorkItemCommentApiResult) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *WorkItemCommentApiResult) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.


### GetCreatedDate

`func (o *WorkItemCommentApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *WorkItemCommentApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *WorkItemCommentApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *WorkItemCommentApiResult) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *WorkItemCommentApiResult) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *WorkItemCommentApiResult) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


