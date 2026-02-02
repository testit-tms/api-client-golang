# UpdateEmptyTestRunApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Test run unique identifier | 
**Name** | **string** | Test run name | 
**Description** | Pointer to **NullableString** | Test run description | [optional] 
**LaunchSource** | Pointer to **NullableString** | Test run launch source              Once launch source is specified it cannot be updated | [optional] 
**Attachments** | Pointer to [**[]AssignAttachmentApiModel**](AssignAttachmentApiModel.md) | Collection of attachments related to the test run | [optional] 
**Links** | Pointer to [**[]UpdateLinkApiModel**](UpdateLinkApiModel.md) | Collection of links related to the test run | [optional] 

## Methods

### NewUpdateEmptyTestRunApiModel

`func NewUpdateEmptyTestRunApiModel(id string, name string, ) *UpdateEmptyTestRunApiModel`

NewUpdateEmptyTestRunApiModel instantiates a new UpdateEmptyTestRunApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEmptyTestRunApiModelWithDefaults

`func NewUpdateEmptyTestRunApiModelWithDefaults() *UpdateEmptyTestRunApiModel`

NewUpdateEmptyTestRunApiModelWithDefaults instantiates a new UpdateEmptyTestRunApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateEmptyTestRunApiModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateEmptyTestRunApiModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateEmptyTestRunApiModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateEmptyTestRunApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateEmptyTestRunApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateEmptyTestRunApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateEmptyTestRunApiModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateEmptyTestRunApiModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateEmptyTestRunApiModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateEmptyTestRunApiModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateEmptyTestRunApiModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateEmptyTestRunApiModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLaunchSource

`func (o *UpdateEmptyTestRunApiModel) GetLaunchSource() string`

GetLaunchSource returns the LaunchSource field if non-nil, zero value otherwise.

### GetLaunchSourceOk

`func (o *UpdateEmptyTestRunApiModel) GetLaunchSourceOk() (*string, bool)`

GetLaunchSourceOk returns a tuple with the LaunchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSource

`func (o *UpdateEmptyTestRunApiModel) SetLaunchSource(v string)`

SetLaunchSource sets LaunchSource field to given value.

### HasLaunchSource

`func (o *UpdateEmptyTestRunApiModel) HasLaunchSource() bool`

HasLaunchSource returns a boolean if a field has been set.

### SetLaunchSourceNil

`func (o *UpdateEmptyTestRunApiModel) SetLaunchSourceNil(b bool)`

 SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil

### UnsetLaunchSource
`func (o *UpdateEmptyTestRunApiModel) UnsetLaunchSource()`

UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
### GetAttachments

`func (o *UpdateEmptyTestRunApiModel) GetAttachments() []AssignAttachmentApiModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *UpdateEmptyTestRunApiModel) GetAttachmentsOk() (*[]AssignAttachmentApiModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *UpdateEmptyTestRunApiModel) SetAttachments(v []AssignAttachmentApiModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *UpdateEmptyTestRunApiModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *UpdateEmptyTestRunApiModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *UpdateEmptyTestRunApiModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetLinks

`func (o *UpdateEmptyTestRunApiModel) GetLinks() []UpdateLinkApiModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UpdateEmptyTestRunApiModel) GetLinksOk() (*[]UpdateLinkApiModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UpdateEmptyTestRunApiModel) SetLinks(v []UpdateLinkApiModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UpdateEmptyTestRunApiModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *UpdateEmptyTestRunApiModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *UpdateEmptyTestRunApiModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


