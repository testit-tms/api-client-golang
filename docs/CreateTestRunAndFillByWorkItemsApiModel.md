# CreateTestRunAndFillByWorkItemsApiModel

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
**ConfigurationIds** | **[]string** | Specifies the configuration GUIDs, from which test points are created. You can specify several GUIDs. | 
**WorkItemIds** | **[]string** | Specifies the work item GUIDs, from which test points are created. You can specify several GUIDs. | 

## Methods

### NewCreateTestRunAndFillByWorkItemsApiModel

`func NewCreateTestRunAndFillByWorkItemsApiModel(projectId string, testPlanId string, configurationIds []string, workItemIds []string, ) *CreateTestRunAndFillByWorkItemsApiModel`

NewCreateTestRunAndFillByWorkItemsApiModel instantiates a new CreateTestRunAndFillByWorkItemsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTestRunAndFillByWorkItemsApiModelWithDefaults

`func NewCreateTestRunAndFillByWorkItemsApiModelWithDefaults() *CreateTestRunAndFillByWorkItemsApiModel`

NewCreateTestRunAndFillByWorkItemsApiModelWithDefaults instantiates a new CreateTestRunAndFillByWorkItemsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *CreateTestRunAndFillByWorkItemsApiModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateTestRunAndFillByWorkItemsApiModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateTestRunAndFillByWorkItemsApiModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetTestPlanId

`func (o *CreateTestRunAndFillByWorkItemsApiModel) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *CreateTestRunAndFillByWorkItemsApiModel) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *CreateTestRunAndFillByWorkItemsApiModel) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.


### GetName

`func (o *CreateTestRunAndFillByWorkItemsApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTestRunAndFillByWorkItemsApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTestRunAndFillByWorkItemsApiModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTestRunAndFillByWorkItemsApiModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateTestRunAndFillByWorkItemsApiModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateTestRunAndFillByWorkItemsApiModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CreateTestRunAndFillByWorkItemsApiModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTestRunAndFillByWorkItemsApiModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTestRunAndFillByWorkItemsApiModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTestRunAndFillByWorkItemsApiModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateTestRunAndFillByWorkItemsApiModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateTestRunAndFillByWorkItemsApiModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLaunchSource

`func (o *CreateTestRunAndFillByWorkItemsApiModel) GetLaunchSource() string`

GetLaunchSource returns the LaunchSource field if non-nil, zero value otherwise.

### GetLaunchSourceOk

`func (o *CreateTestRunAndFillByWorkItemsApiModel) GetLaunchSourceOk() (*string, bool)`

GetLaunchSourceOk returns a tuple with the LaunchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSource

`func (o *CreateTestRunAndFillByWorkItemsApiModel) SetLaunchSource(v string)`

SetLaunchSource sets LaunchSource field to given value.

### HasLaunchSource

`func (o *CreateTestRunAndFillByWorkItemsApiModel) HasLaunchSource() bool`

HasLaunchSource returns a boolean if a field has been set.

### SetLaunchSourceNil

`func (o *CreateTestRunAndFillByWorkItemsApiModel) SetLaunchSourceNil(b bool)`

 SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil

### UnsetLaunchSource
`func (o *CreateTestRunAndFillByWorkItemsApiModel) UnsetLaunchSource()`

UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
### GetAttachments

`func (o *CreateTestRunAndFillByWorkItemsApiModel) GetAttachments() []AssignAttachmentApiModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *CreateTestRunAndFillByWorkItemsApiModel) GetAttachmentsOk() (*[]AssignAttachmentApiModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *CreateTestRunAndFillByWorkItemsApiModel) SetAttachments(v []AssignAttachmentApiModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *CreateTestRunAndFillByWorkItemsApiModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *CreateTestRunAndFillByWorkItemsApiModel) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *CreateTestRunAndFillByWorkItemsApiModel) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetLinks

`func (o *CreateTestRunAndFillByWorkItemsApiModel) GetLinks() []CreateLinkApiModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateTestRunAndFillByWorkItemsApiModel) GetLinksOk() (*[]CreateLinkApiModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateTestRunAndFillByWorkItemsApiModel) SetLinks(v []CreateLinkApiModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateTestRunAndFillByWorkItemsApiModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *CreateTestRunAndFillByWorkItemsApiModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *CreateTestRunAndFillByWorkItemsApiModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetConfigurationIds

`func (o *CreateTestRunAndFillByWorkItemsApiModel) GetConfigurationIds() []string`

GetConfigurationIds returns the ConfigurationIds field if non-nil, zero value otherwise.

### GetConfigurationIdsOk

`func (o *CreateTestRunAndFillByWorkItemsApiModel) GetConfigurationIdsOk() (*[]string, bool)`

GetConfigurationIdsOk returns a tuple with the ConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIds

`func (o *CreateTestRunAndFillByWorkItemsApiModel) SetConfigurationIds(v []string)`

SetConfigurationIds sets ConfigurationIds field to given value.


### GetWorkItemIds

`func (o *CreateTestRunAndFillByWorkItemsApiModel) GetWorkItemIds() []string`

GetWorkItemIds returns the WorkItemIds field if non-nil, zero value otherwise.

### GetWorkItemIdsOk

`func (o *CreateTestRunAndFillByWorkItemsApiModel) GetWorkItemIdsOk() (*[]string, bool)`

GetWorkItemIdsOk returns a tuple with the WorkItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemIds

`func (o *CreateTestRunAndFillByWorkItemsApiModel) SetWorkItemIds(v []string)`

SetWorkItemIds sets WorkItemIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


