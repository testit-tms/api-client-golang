# CreateTestRunAndFillByConfigurationsApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | Specifies the GUID of the project, in which a test run will be created. | 
**TestPlanId** | **string** | Specifies the GUID of the test plan, within which the test run will be created. | 
**Name** | Pointer to **NullableString** | Specifies the name of the test run. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the test run description. | [optional] 
**LaunchSource** | Pointer to **NullableString** | Specifies the test run launch source. | [optional] 
**Attachments** | Pointer to [**[]AssignAttachmentApiModel**](AssignAttachmentApiModel.md) | Collection of attachment ids to relate to the test run | [optional] 
**Links** | Pointer to [**[]CreateLinkApiModel**](CreateLinkApiModel.md) | Collection of links to relate to the test run | [optional] 
**TestPointSelectors** | [**[]TestPointSelector**](TestPointSelector.md) | Specifies an array of work items and configuration to create a test run for. | 

## Methods

### NewCreateTestRunAndFillByConfigurationsApiModel

`func NewCreateTestRunAndFillByConfigurationsApiModel(projectId string, testPlanId string, testPointSelectors []TestPointSelector, ) *CreateTestRunAndFillByConfigurationsApiModel`

NewCreateTestRunAndFillByConfigurationsApiModel instantiates a new CreateTestRunAndFillByConfigurationsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTestRunAndFillByConfigurationsApiModelWithDefaults

`func NewCreateTestRunAndFillByConfigurationsApiModelWithDefaults() *CreateTestRunAndFillByConfigurationsApiModel`

NewCreateTestRunAndFillByConfigurationsApiModelWithDefaults instantiates a new CreateTestRunAndFillByConfigurationsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *CreateTestRunAndFillByConfigurationsApiModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateTestRunAndFillByConfigurationsApiModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateTestRunAndFillByConfigurationsApiModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetTestPlanId

`func (o *CreateTestRunAndFillByConfigurationsApiModel) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *CreateTestRunAndFillByConfigurationsApiModel) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *CreateTestRunAndFillByConfigurationsApiModel) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.


### GetName

`func (o *CreateTestRunAndFillByConfigurationsApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTestRunAndFillByConfigurationsApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTestRunAndFillByConfigurationsApiModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTestRunAndFillByConfigurationsApiModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateTestRunAndFillByConfigurationsApiModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateTestRunAndFillByConfigurationsApiModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CreateTestRunAndFillByConfigurationsApiModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTestRunAndFillByConfigurationsApiModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTestRunAndFillByConfigurationsApiModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTestRunAndFillByConfigurationsApiModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateTestRunAndFillByConfigurationsApiModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateTestRunAndFillByConfigurationsApiModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLaunchSource

`func (o *CreateTestRunAndFillByConfigurationsApiModel) GetLaunchSource() string`

GetLaunchSource returns the LaunchSource field if non-nil, zero value otherwise.

### GetLaunchSourceOk

`func (o *CreateTestRunAndFillByConfigurationsApiModel) GetLaunchSourceOk() (*string, bool)`

GetLaunchSourceOk returns a tuple with the LaunchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSource

`func (o *CreateTestRunAndFillByConfigurationsApiModel) SetLaunchSource(v string)`

SetLaunchSource sets LaunchSource field to given value.

### HasLaunchSource

`func (o *CreateTestRunAndFillByConfigurationsApiModel) HasLaunchSource() bool`

HasLaunchSource returns a boolean if a field has been set.

### SetLaunchSourceNil

`func (o *CreateTestRunAndFillByConfigurationsApiModel) SetLaunchSourceNil(b bool)`

 SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil

### UnsetLaunchSource
`func (o *CreateTestRunAndFillByConfigurationsApiModel) UnsetLaunchSource()`

UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
### GetAttachments

`func (o *CreateTestRunAndFillByConfigurationsApiModel) GetAttachments() []AssignAttachmentApiModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *CreateTestRunAndFillByConfigurationsApiModel) GetAttachmentsOk() (*[]AssignAttachmentApiModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *CreateTestRunAndFillByConfigurationsApiModel) SetAttachments(v []AssignAttachmentApiModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *CreateTestRunAndFillByConfigurationsApiModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *CreateTestRunAndFillByConfigurationsApiModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *CreateTestRunAndFillByConfigurationsApiModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetLinks

`func (o *CreateTestRunAndFillByConfigurationsApiModel) GetLinks() []CreateLinkApiModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateTestRunAndFillByConfigurationsApiModel) GetLinksOk() (*[]CreateLinkApiModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateTestRunAndFillByConfigurationsApiModel) SetLinks(v []CreateLinkApiModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateTestRunAndFillByConfigurationsApiModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *CreateTestRunAndFillByConfigurationsApiModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *CreateTestRunAndFillByConfigurationsApiModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetTestPointSelectors

`func (o *CreateTestRunAndFillByConfigurationsApiModel) GetTestPointSelectors() []TestPointSelector`

GetTestPointSelectors returns the TestPointSelectors field if non-nil, zero value otherwise.

### GetTestPointSelectorsOk

`func (o *CreateTestRunAndFillByConfigurationsApiModel) GetTestPointSelectorsOk() (*[]TestPointSelector, bool)`

GetTestPointSelectorsOk returns a tuple with the TestPointSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPointSelectors

`func (o *CreateTestRunAndFillByConfigurationsApiModel) SetTestPointSelectors(v []TestPointSelector)`

SetTestPointSelectors sets TestPointSelectors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


