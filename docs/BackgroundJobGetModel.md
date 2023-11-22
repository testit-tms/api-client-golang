# BackgroundJobGetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**JobId** | **string** |  | 
**JobType** | [**BackgroundJobType**](BackgroundJobType.md) |  | 
**State** | [**BackgroundJobState**](BackgroundJobState.md) |  | 
**IsDeleted** | **bool** |  | 
**Progress** | **int64** |  | 
**CreatedDate** | **time.Time** |  | 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 
**Attachments** | [**[]BackgroundJobAttachmentModel**](BackgroundJobAttachmentModel.md) |  | 

## Methods

### NewBackgroundJobGetModel

`func NewBackgroundJobGetModel(id string, jobId string, jobType BackgroundJobType, state BackgroundJobState, isDeleted bool, progress int64, createdDate time.Time, attachments []BackgroundJobAttachmentModel, ) *BackgroundJobGetModel`

NewBackgroundJobGetModel instantiates a new BackgroundJobGetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackgroundJobGetModelWithDefaults

`func NewBackgroundJobGetModelWithDefaults() *BackgroundJobGetModel`

NewBackgroundJobGetModelWithDefaults instantiates a new BackgroundJobGetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackgroundJobGetModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackgroundJobGetModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackgroundJobGetModel) SetId(v string)`

SetId sets Id field to given value.


### GetJobId

`func (o *BackgroundJobGetModel) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *BackgroundJobGetModel) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *BackgroundJobGetModel) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetJobType

`func (o *BackgroundJobGetModel) GetJobType() BackgroundJobType`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *BackgroundJobGetModel) GetJobTypeOk() (*BackgroundJobType, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *BackgroundJobGetModel) SetJobType(v BackgroundJobType)`

SetJobType sets JobType field to given value.


### GetState

`func (o *BackgroundJobGetModel) GetState() BackgroundJobState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BackgroundJobGetModel) GetStateOk() (*BackgroundJobState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BackgroundJobGetModel) SetState(v BackgroundJobState)`

SetState sets State field to given value.


### GetIsDeleted

`func (o *BackgroundJobGetModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *BackgroundJobGetModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *BackgroundJobGetModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetProgress

`func (o *BackgroundJobGetModel) GetProgress() int64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *BackgroundJobGetModel) GetProgressOk() (*int64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *BackgroundJobGetModel) SetProgress(v int64)`

SetProgress sets Progress field to given value.


### GetCreatedDate

`func (o *BackgroundJobGetModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *BackgroundJobGetModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *BackgroundJobGetModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetStartDate

`func (o *BackgroundJobGetModel) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BackgroundJobGetModel) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BackgroundJobGetModel) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BackgroundJobGetModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *BackgroundJobGetModel) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *BackgroundJobGetModel) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *BackgroundJobGetModel) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BackgroundJobGetModel) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BackgroundJobGetModel) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BackgroundJobGetModel) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *BackgroundJobGetModel) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *BackgroundJobGetModel) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetError

`func (o *BackgroundJobGetModel) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BackgroundJobGetModel) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BackgroundJobGetModel) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BackgroundJobGetModel) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *BackgroundJobGetModel) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *BackgroundJobGetModel) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetAttachments

`func (o *BackgroundJobGetModel) GetAttachments() []BackgroundJobAttachmentModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *BackgroundJobGetModel) GetAttachmentsOk() (*[]BackgroundJobAttachmentModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *BackgroundJobGetModel) SetAttachments(v []BackgroundJobAttachmentModel)`

SetAttachments sets Attachments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


