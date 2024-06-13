# TestRunUpdateMultipleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectModel** | [**TestRunSelectModel**](TestRunSelectModel.md) |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**AttachmentUpdateScheme** | [**UpdateAttachmentShortModel**](UpdateAttachmentShortModel.md) |  | 
**LinkUpdateScheme** | [**UpdateLinkShortModel**](UpdateLinkShortModel.md) |  | 

## Methods

### NewTestRunUpdateMultipleModel

`func NewTestRunUpdateMultipleModel(selectModel TestRunSelectModel, attachmentUpdateScheme UpdateAttachmentShortModel, linkUpdateScheme UpdateLinkShortModel, ) *TestRunUpdateMultipleModel`

NewTestRunUpdateMultipleModel instantiates a new TestRunUpdateMultipleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunUpdateMultipleModelWithDefaults

`func NewTestRunUpdateMultipleModelWithDefaults() *TestRunUpdateMultipleModel`

NewTestRunUpdateMultipleModelWithDefaults instantiates a new TestRunUpdateMultipleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectModel

`func (o *TestRunUpdateMultipleModel) GetSelectModel() TestRunSelectModel`

GetSelectModel returns the SelectModel field if non-nil, zero value otherwise.

### GetSelectModelOk

`func (o *TestRunUpdateMultipleModel) GetSelectModelOk() (*TestRunSelectModel, bool)`

GetSelectModelOk returns a tuple with the SelectModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectModel

`func (o *TestRunUpdateMultipleModel) SetSelectModel(v TestRunSelectModel)`

SetSelectModel sets SelectModel field to given value.


### GetDescription

`func (o *TestRunUpdateMultipleModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestRunUpdateMultipleModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestRunUpdateMultipleModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestRunUpdateMultipleModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestRunUpdateMultipleModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestRunUpdateMultipleModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAttachmentUpdateScheme

`func (o *TestRunUpdateMultipleModel) GetAttachmentUpdateScheme() UpdateAttachmentShortModel`

GetAttachmentUpdateScheme returns the AttachmentUpdateScheme field if non-nil, zero value otherwise.

### GetAttachmentUpdateSchemeOk

`func (o *TestRunUpdateMultipleModel) GetAttachmentUpdateSchemeOk() (*UpdateAttachmentShortModel, bool)`

GetAttachmentUpdateSchemeOk returns a tuple with the AttachmentUpdateScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentUpdateScheme

`func (o *TestRunUpdateMultipleModel) SetAttachmentUpdateScheme(v UpdateAttachmentShortModel)`

SetAttachmentUpdateScheme sets AttachmentUpdateScheme field to given value.


### GetLinkUpdateScheme

`func (o *TestRunUpdateMultipleModel) GetLinkUpdateScheme() UpdateLinkShortModel`

GetLinkUpdateScheme returns the LinkUpdateScheme field if non-nil, zero value otherwise.

### GetLinkUpdateSchemeOk

`func (o *TestRunUpdateMultipleModel) GetLinkUpdateSchemeOk() (*UpdateLinkShortModel, bool)`

GetLinkUpdateSchemeOk returns a tuple with the LinkUpdateScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUpdateScheme

`func (o *TestRunUpdateMultipleModel) SetLinkUpdateScheme(v UpdateLinkShortModel)`

SetLinkUpdateScheme sets LinkUpdateScheme field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


