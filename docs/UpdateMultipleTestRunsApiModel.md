# UpdateMultipleTestRunsApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectModel** | [**TestRunSelectApiModel**](TestRunSelectApiModel.md) | Test run selection model | 
**Description** | Pointer to **NullableString** | Test run description | [optional] 
**AttachmentUpdateScheme** | Pointer to [**NullableUpdateMultipleAttachmentsApiModel**](UpdateMultipleAttachmentsApiModel.md) | Set of attachment ids | [optional] 
**LinkUpdateScheme** | Pointer to [**NullableUpdateMultipleLinksApiModel**](UpdateMultipleLinksApiModel.md) | Set of links | [optional] 

## Methods

### NewUpdateMultipleTestRunsApiModel

`func NewUpdateMultipleTestRunsApiModel(selectModel TestRunSelectApiModel, ) *UpdateMultipleTestRunsApiModel`

NewUpdateMultipleTestRunsApiModel instantiates a new UpdateMultipleTestRunsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMultipleTestRunsApiModelWithDefaults

`func NewUpdateMultipleTestRunsApiModelWithDefaults() *UpdateMultipleTestRunsApiModel`

NewUpdateMultipleTestRunsApiModelWithDefaults instantiates a new UpdateMultipleTestRunsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectModel

`func (o *UpdateMultipleTestRunsApiModel) GetSelectModel() TestRunSelectApiModel`

GetSelectModel returns the SelectModel field if non-nil, zero value otherwise.

### GetSelectModelOk

`func (o *UpdateMultipleTestRunsApiModel) GetSelectModelOk() (*TestRunSelectApiModel, bool)`

GetSelectModelOk returns a tuple with the SelectModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectModel

`func (o *UpdateMultipleTestRunsApiModel) SetSelectModel(v TestRunSelectApiModel)`

SetSelectModel sets SelectModel field to given value.


### GetDescription

`func (o *UpdateMultipleTestRunsApiModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateMultipleTestRunsApiModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateMultipleTestRunsApiModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateMultipleTestRunsApiModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateMultipleTestRunsApiModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateMultipleTestRunsApiModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAttachmentUpdateScheme

`func (o *UpdateMultipleTestRunsApiModel) GetAttachmentUpdateScheme() UpdateMultipleAttachmentsApiModel`

GetAttachmentUpdateScheme returns the AttachmentUpdateScheme field if non-nil, zero value otherwise.

### GetAttachmentUpdateSchemeOk

`func (o *UpdateMultipleTestRunsApiModel) GetAttachmentUpdateSchemeOk() (*UpdateMultipleAttachmentsApiModel, bool)`

GetAttachmentUpdateSchemeOk returns a tuple with the AttachmentUpdateScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentUpdateScheme

`func (o *UpdateMultipleTestRunsApiModel) SetAttachmentUpdateScheme(v UpdateMultipleAttachmentsApiModel)`

SetAttachmentUpdateScheme sets AttachmentUpdateScheme field to given value.

### HasAttachmentUpdateScheme

`func (o *UpdateMultipleTestRunsApiModel) HasAttachmentUpdateScheme() bool`

HasAttachmentUpdateScheme returns a boolean if a field has been set.

### SetAttachmentUpdateSchemeNil

`func (o *UpdateMultipleTestRunsApiModel) SetAttachmentUpdateSchemeNil(b bool)`

 SetAttachmentUpdateSchemeNil sets the value for AttachmentUpdateScheme to be an explicit nil

### UnsetAttachmentUpdateScheme
`func (o *UpdateMultipleTestRunsApiModel) UnsetAttachmentUpdateScheme()`

UnsetAttachmentUpdateScheme ensures that no value is present for AttachmentUpdateScheme, not even an explicit nil
### GetLinkUpdateScheme

`func (o *UpdateMultipleTestRunsApiModel) GetLinkUpdateScheme() UpdateMultipleLinksApiModel`

GetLinkUpdateScheme returns the LinkUpdateScheme field if non-nil, zero value otherwise.

### GetLinkUpdateSchemeOk

`func (o *UpdateMultipleTestRunsApiModel) GetLinkUpdateSchemeOk() (*UpdateMultipleLinksApiModel, bool)`

GetLinkUpdateSchemeOk returns a tuple with the LinkUpdateScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUpdateScheme

`func (o *UpdateMultipleTestRunsApiModel) SetLinkUpdateScheme(v UpdateMultipleLinksApiModel)`

SetLinkUpdateScheme sets LinkUpdateScheme field to given value.

### HasLinkUpdateScheme

`func (o *UpdateMultipleTestRunsApiModel) HasLinkUpdateScheme() bool`

HasLinkUpdateScheme returns a boolean if a field has been set.

### SetLinkUpdateSchemeNil

`func (o *UpdateMultipleTestRunsApiModel) SetLinkUpdateSchemeNil(b bool)`

 SetLinkUpdateSchemeNil sets the value for LinkUpdateScheme to be an explicit nil

### UnsetLinkUpdateScheme
`func (o *UpdateMultipleTestRunsApiModel) UnsetLinkUpdateScheme()`

UnsetLinkUpdateScheme ensures that no value is present for LinkUpdateScheme, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


