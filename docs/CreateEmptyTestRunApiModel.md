# CreateEmptyTestRunApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | Project unique identifier              This property is to link test run with a project | 
**Name** | Pointer to **NullableString** | Test run name | [optional] 
**Description** | Pointer to **NullableString** | Test run description | [optional] 
**LaunchSource** | Pointer to **NullableString** | Test run launch source | [optional] 
**Attachments** | Pointer to [**[]AssignAttachmentApiModel**](AssignAttachmentApiModel.md) | Collection of attachments to relate to the test run | [optional] 
**Links** | Pointer to [**[]CreateLinkApiModel**](CreateLinkApiModel.md) | Collection of links to relate to the test run | [optional] 

## Methods

### NewCreateEmptyTestRunApiModel

`func NewCreateEmptyTestRunApiModel(projectId string, ) *CreateEmptyTestRunApiModel`

NewCreateEmptyTestRunApiModel instantiates a new CreateEmptyTestRunApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEmptyTestRunApiModelWithDefaults

`func NewCreateEmptyTestRunApiModelWithDefaults() *CreateEmptyTestRunApiModel`

NewCreateEmptyTestRunApiModelWithDefaults instantiates a new CreateEmptyTestRunApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *CreateEmptyTestRunApiModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateEmptyTestRunApiModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateEmptyTestRunApiModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *CreateEmptyTestRunApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEmptyTestRunApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEmptyTestRunApiModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateEmptyTestRunApiModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateEmptyTestRunApiModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateEmptyTestRunApiModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CreateEmptyTestRunApiModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateEmptyTestRunApiModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateEmptyTestRunApiModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateEmptyTestRunApiModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateEmptyTestRunApiModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateEmptyTestRunApiModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLaunchSource

`func (o *CreateEmptyTestRunApiModel) GetLaunchSource() string`

GetLaunchSource returns the LaunchSource field if non-nil, zero value otherwise.

### GetLaunchSourceOk

`func (o *CreateEmptyTestRunApiModel) GetLaunchSourceOk() (*string, bool)`

GetLaunchSourceOk returns a tuple with the LaunchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSource

`func (o *CreateEmptyTestRunApiModel) SetLaunchSource(v string)`

SetLaunchSource sets LaunchSource field to given value.

### HasLaunchSource

`func (o *CreateEmptyTestRunApiModel) HasLaunchSource() bool`

HasLaunchSource returns a boolean if a field has been set.

### SetLaunchSourceNil

`func (o *CreateEmptyTestRunApiModel) SetLaunchSourceNil(b bool)`

 SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil

### UnsetLaunchSource
`func (o *CreateEmptyTestRunApiModel) UnsetLaunchSource()`

UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
### GetAttachments

`func (o *CreateEmptyTestRunApiModel) GetAttachments() []AssignAttachmentApiModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *CreateEmptyTestRunApiModel) GetAttachmentsOk() (*[]AssignAttachmentApiModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *CreateEmptyTestRunApiModel) SetAttachments(v []AssignAttachmentApiModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *CreateEmptyTestRunApiModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *CreateEmptyTestRunApiModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *CreateEmptyTestRunApiModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetLinks

`func (o *CreateEmptyTestRunApiModel) GetLinks() []CreateLinkApiModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateEmptyTestRunApiModel) GetLinksOk() (*[]CreateLinkApiModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateEmptyTestRunApiModel) SetLinks(v []CreateLinkApiModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateEmptyTestRunApiModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *CreateEmptyTestRunApiModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *CreateEmptyTestRunApiModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


