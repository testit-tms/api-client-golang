# UpdateMultipleAttachmentsApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**ActionUpdate**](ActionUpdate.md) |  | 
**AttachmentIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateMultipleAttachmentsApiModel

`func NewUpdateMultipleAttachmentsApiModel(action ActionUpdate, ) *UpdateMultipleAttachmentsApiModel`

NewUpdateMultipleAttachmentsApiModel instantiates a new UpdateMultipleAttachmentsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMultipleAttachmentsApiModelWithDefaults

`func NewUpdateMultipleAttachmentsApiModelWithDefaults() *UpdateMultipleAttachmentsApiModel`

NewUpdateMultipleAttachmentsApiModelWithDefaults instantiates a new UpdateMultipleAttachmentsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UpdateMultipleAttachmentsApiModel) GetAction() ActionUpdate`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateMultipleAttachmentsApiModel) GetActionOk() (*ActionUpdate, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateMultipleAttachmentsApiModel) SetAction(v ActionUpdate)`

SetAction sets Action field to given value.


### GetAttachmentIds

`func (o *UpdateMultipleAttachmentsApiModel) GetAttachmentIds() []string`

GetAttachmentIds returns the AttachmentIds field if non-nil, zero value otherwise.

### GetAttachmentIdsOk

`func (o *UpdateMultipleAttachmentsApiModel) GetAttachmentIdsOk() (*[]string, bool)`

GetAttachmentIdsOk returns a tuple with the AttachmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentIds

`func (o *UpdateMultipleAttachmentsApiModel) SetAttachmentIds(v []string)`

SetAttachmentIds sets AttachmentIds field to given value.

### HasAttachmentIds

`func (o *UpdateMultipleAttachmentsApiModel) HasAttachmentIds() bool`

HasAttachmentIds returns a boolean if a field has been set.

### SetAttachmentIdsNil

`func (o *UpdateMultipleAttachmentsApiModel) SetAttachmentIdsNil(b bool)`

 SetAttachmentIdsNil sets the value for AttachmentIds to be an explicit nil

### UnsetAttachmentIds
`func (o *UpdateMultipleAttachmentsApiModel) UnsetAttachmentIds()`

UnsetAttachmentIds ensures that no value is present for AttachmentIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


