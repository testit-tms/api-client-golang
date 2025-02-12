# CreateTestRunAndFillByAutoTestsApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | Specifies the GUID of the project, in which a test run will be created. | 
**Name** | Pointer to **NullableString** | Specifies the name of the test run. | [optional] 
**ConfigurationIds** | **[]string** | Specifies the configuration GUIDs, from which test points are created. You can specify several GUIDs. | 
**AutoTestExternalIds** | **[]string** | Specifies the external ID of the autotest. You can specify several IDs. | 
**Description** | Pointer to **NullableString** | Specifies the test run description. | [optional] 
**LaunchSource** | Pointer to **NullableString** | Specifies the test run launch source. | [optional] 
**Attachments** | Pointer to [**[]AssignAttachmentApiModel**](AssignAttachmentApiModel.md) | Collection of attachment ids to relate to the test run | [optional] 
**Links** | Pointer to [**[]CreateLinkApiModel**](CreateLinkApiModel.md) | Collection of links to relate to the test run | [optional] 

## Methods

### NewCreateTestRunAndFillByAutoTestsApiModel

`func NewCreateTestRunAndFillByAutoTestsApiModel(projectId string, configurationIds []string, autoTestExternalIds []string, ) *CreateTestRunAndFillByAutoTestsApiModel`

NewCreateTestRunAndFillByAutoTestsApiModel instantiates a new CreateTestRunAndFillByAutoTestsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTestRunAndFillByAutoTestsApiModelWithDefaults

`func NewCreateTestRunAndFillByAutoTestsApiModelWithDefaults() *CreateTestRunAndFillByAutoTestsApiModel`

NewCreateTestRunAndFillByAutoTestsApiModelWithDefaults instantiates a new CreateTestRunAndFillByAutoTestsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *CreateTestRunAndFillByAutoTestsApiModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateTestRunAndFillByAutoTestsApiModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateTestRunAndFillByAutoTestsApiModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *CreateTestRunAndFillByAutoTestsApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTestRunAndFillByAutoTestsApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTestRunAndFillByAutoTestsApiModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTestRunAndFillByAutoTestsApiModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateTestRunAndFillByAutoTestsApiModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateTestRunAndFillByAutoTestsApiModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetConfigurationIds

`func (o *CreateTestRunAndFillByAutoTestsApiModel) GetConfigurationIds() []string`

GetConfigurationIds returns the ConfigurationIds field if non-nil, zero value otherwise.

### GetConfigurationIdsOk

`func (o *CreateTestRunAndFillByAutoTestsApiModel) GetConfigurationIdsOk() (*[]string, bool)`

GetConfigurationIdsOk returns a tuple with the ConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIds

`func (o *CreateTestRunAndFillByAutoTestsApiModel) SetConfigurationIds(v []string)`

SetConfigurationIds sets ConfigurationIds field to given value.


### GetAutoTestExternalIds

`func (o *CreateTestRunAndFillByAutoTestsApiModel) GetAutoTestExternalIds() []string`

GetAutoTestExternalIds returns the AutoTestExternalIds field if non-nil, zero value otherwise.

### GetAutoTestExternalIdsOk

`func (o *CreateTestRunAndFillByAutoTestsApiModel) GetAutoTestExternalIdsOk() (*[]string, bool)`

GetAutoTestExternalIdsOk returns a tuple with the AutoTestExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestExternalIds

`func (o *CreateTestRunAndFillByAutoTestsApiModel) SetAutoTestExternalIds(v []string)`

SetAutoTestExternalIds sets AutoTestExternalIds field to given value.


### GetDescription

`func (o *CreateTestRunAndFillByAutoTestsApiModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTestRunAndFillByAutoTestsApiModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTestRunAndFillByAutoTestsApiModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTestRunAndFillByAutoTestsApiModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateTestRunAndFillByAutoTestsApiModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateTestRunAndFillByAutoTestsApiModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLaunchSource

`func (o *CreateTestRunAndFillByAutoTestsApiModel) GetLaunchSource() string`

GetLaunchSource returns the LaunchSource field if non-nil, zero value otherwise.

### GetLaunchSourceOk

`func (o *CreateTestRunAndFillByAutoTestsApiModel) GetLaunchSourceOk() (*string, bool)`

GetLaunchSourceOk returns a tuple with the LaunchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSource

`func (o *CreateTestRunAndFillByAutoTestsApiModel) SetLaunchSource(v string)`

SetLaunchSource sets LaunchSource field to given value.

### HasLaunchSource

`func (o *CreateTestRunAndFillByAutoTestsApiModel) HasLaunchSource() bool`

HasLaunchSource returns a boolean if a field has been set.

### SetLaunchSourceNil

`func (o *CreateTestRunAndFillByAutoTestsApiModel) SetLaunchSourceNil(b bool)`

 SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil

### UnsetLaunchSource
`func (o *CreateTestRunAndFillByAutoTestsApiModel) UnsetLaunchSource()`

UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
### GetAttachments

`func (o *CreateTestRunAndFillByAutoTestsApiModel) GetAttachments() []AssignAttachmentApiModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *CreateTestRunAndFillByAutoTestsApiModel) GetAttachmentsOk() (*[]AssignAttachmentApiModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *CreateTestRunAndFillByAutoTestsApiModel) SetAttachments(v []AssignAttachmentApiModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *CreateTestRunAndFillByAutoTestsApiModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *CreateTestRunAndFillByAutoTestsApiModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *CreateTestRunAndFillByAutoTestsApiModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetLinks

`func (o *CreateTestRunAndFillByAutoTestsApiModel) GetLinks() []CreateLinkApiModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateTestRunAndFillByAutoTestsApiModel) GetLinksOk() (*[]CreateLinkApiModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateTestRunAndFillByAutoTestsApiModel) SetLinks(v []CreateLinkApiModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateTestRunAndFillByAutoTestsApiModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *CreateTestRunAndFillByAutoTestsApiModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *CreateTestRunAndFillByAutoTestsApiModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


