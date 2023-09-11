# TestRunUpdateMultipleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectModel** | [**TestRunSelectionModel**](TestRunSelectionModel.md) |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**AttachmentUpdateScheme** | Pointer to [**NullableSetOfAttachmentIds**](SetOfAttachmentIds.md) |  | [optional] 
**LinkUpdateScheme** | Pointer to [**NullableSetOfLinks**](SetOfLinks.md) |  | [optional] 

## Methods

### NewTestRunUpdateMultipleModel

`func NewTestRunUpdateMultipleModel(selectModel TestRunSelectionModel, ) *TestRunUpdateMultipleModel`

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

`func (o *TestRunUpdateMultipleModel) GetSelectModel() TestRunSelectionModel`

GetSelectModel returns the SelectModel field if non-nil, zero value otherwise.

### GetSelectModelOk

`func (o *TestRunUpdateMultipleModel) GetSelectModelOk() (*TestRunSelectionModel, bool)`

GetSelectModelOk returns a tuple with the SelectModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectModel

`func (o *TestRunUpdateMultipleModel) SetSelectModel(v TestRunSelectionModel)`

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

`func (o *TestRunUpdateMultipleModel) GetAttachmentUpdateScheme() SetOfAttachmentIds`

GetAttachmentUpdateScheme returns the AttachmentUpdateScheme field if non-nil, zero value otherwise.

### GetAttachmentUpdateSchemeOk

`func (o *TestRunUpdateMultipleModel) GetAttachmentUpdateSchemeOk() (*SetOfAttachmentIds, bool)`

GetAttachmentUpdateSchemeOk returns a tuple with the AttachmentUpdateScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentUpdateScheme

`func (o *TestRunUpdateMultipleModel) SetAttachmentUpdateScheme(v SetOfAttachmentIds)`

SetAttachmentUpdateScheme sets AttachmentUpdateScheme field to given value.

### HasAttachmentUpdateScheme

`func (o *TestRunUpdateMultipleModel) HasAttachmentUpdateScheme() bool`

HasAttachmentUpdateScheme returns a boolean if a field has been set.

### SetAttachmentUpdateSchemeNil

`func (o *TestRunUpdateMultipleModel) SetAttachmentUpdateSchemeNil(b bool)`

 SetAttachmentUpdateSchemeNil sets the value for AttachmentUpdateScheme to be an explicit nil

### UnsetAttachmentUpdateScheme
`func (o *TestRunUpdateMultipleModel) UnsetAttachmentUpdateScheme()`

UnsetAttachmentUpdateScheme ensures that no value is present for AttachmentUpdateScheme, not even an explicit nil
### GetLinkUpdateScheme

`func (o *TestRunUpdateMultipleModel) GetLinkUpdateScheme() SetOfLinks`

GetLinkUpdateScheme returns the LinkUpdateScheme field if non-nil, zero value otherwise.

### GetLinkUpdateSchemeOk

`func (o *TestRunUpdateMultipleModel) GetLinkUpdateSchemeOk() (*SetOfLinks, bool)`

GetLinkUpdateSchemeOk returns a tuple with the LinkUpdateScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUpdateScheme

`func (o *TestRunUpdateMultipleModel) SetLinkUpdateScheme(v SetOfLinks)`

SetLinkUpdateScheme sets LinkUpdateScheme field to given value.

### HasLinkUpdateScheme

`func (o *TestRunUpdateMultipleModel) HasLinkUpdateScheme() bool`

HasLinkUpdateScheme returns a boolean if a field has been set.

### SetLinkUpdateSchemeNil

`func (o *TestRunUpdateMultipleModel) SetLinkUpdateSchemeNil(b bool)`

 SetLinkUpdateSchemeNil sets the value for LinkUpdateScheme to be an explicit nil

### UnsetLinkUpdateScheme
`func (o *TestRunUpdateMultipleModel) UnsetLinkUpdateScheme()`

UnsetLinkUpdateScheme ensures that no value is present for LinkUpdateScheme, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


