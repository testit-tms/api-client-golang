# CreateAndFillByWorkItemsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationIds** | **[]string** | Specifies the configuration GUIDs, from which test points are created. You can specify several GUIDs. | 
**WorkItemIds** | **[]string** | Specifies the work item GUIDs, from which test points are created. You can specify several GUIDs. | 
**ProjectId** | **string** | Specifies the GUID of the project, in which a test run will be created. | 
**TestPlanId** | **string** | Specifies the GUID of the test plan, within which the test run will be created. | 
**Name** | Pointer to **NullableString** | Specifies the name of the test run. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the test run description. | [optional] 
**LaunchSource** | Pointer to **NullableString** | Specifies the test run launch source. | [optional] 
**Attachments** | Pointer to [**[]AttachmentPutModel**](AttachmentPutModel.md) | Collection of attachment ids to relate to the test run | [optional] 
**Links** | Pointer to [**[]LinkPostModel**](LinkPostModel.md) | Collection of links to relate to the test run | [optional] 

## Methods

### NewCreateAndFillByWorkItemsRequest

`func NewCreateAndFillByWorkItemsRequest(configurationIds []string, workItemIds []string, projectId string, testPlanId string, ) *CreateAndFillByWorkItemsRequest`

NewCreateAndFillByWorkItemsRequest instantiates a new CreateAndFillByWorkItemsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAndFillByWorkItemsRequestWithDefaults

`func NewCreateAndFillByWorkItemsRequestWithDefaults() *CreateAndFillByWorkItemsRequest`

NewCreateAndFillByWorkItemsRequestWithDefaults instantiates a new CreateAndFillByWorkItemsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationIds

`func (o *CreateAndFillByWorkItemsRequest) GetConfigurationIds() []string`

GetConfigurationIds returns the ConfigurationIds field if non-nil, zero value otherwise.

### GetConfigurationIdsOk

`func (o *CreateAndFillByWorkItemsRequest) GetConfigurationIdsOk() (*[]string, bool)`

GetConfigurationIdsOk returns a tuple with the ConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationIds

`func (o *CreateAndFillByWorkItemsRequest) SetConfigurationIds(v []string)`

SetConfigurationIds sets ConfigurationIds field to given value.


### GetWorkItemIds

`func (o *CreateAndFillByWorkItemsRequest) GetWorkItemIds() []string`

GetWorkItemIds returns the WorkItemIds field if non-nil, zero value otherwise.

### GetWorkItemIdsOk

`func (o *CreateAndFillByWorkItemsRequest) GetWorkItemIdsOk() (*[]string, bool)`

GetWorkItemIdsOk returns a tuple with the WorkItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemIds

`func (o *CreateAndFillByWorkItemsRequest) SetWorkItemIds(v []string)`

SetWorkItemIds sets WorkItemIds field to given value.


### GetProjectId

`func (o *CreateAndFillByWorkItemsRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateAndFillByWorkItemsRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateAndFillByWorkItemsRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetTestPlanId

`func (o *CreateAndFillByWorkItemsRequest) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *CreateAndFillByWorkItemsRequest) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *CreateAndFillByWorkItemsRequest) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.


### GetName

`func (o *CreateAndFillByWorkItemsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAndFillByWorkItemsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAndFillByWorkItemsRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateAndFillByWorkItemsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateAndFillByWorkItemsRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateAndFillByWorkItemsRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CreateAndFillByWorkItemsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAndFillByWorkItemsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAndFillByWorkItemsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAndFillByWorkItemsRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateAndFillByWorkItemsRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateAndFillByWorkItemsRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLaunchSource

`func (o *CreateAndFillByWorkItemsRequest) GetLaunchSource() string`

GetLaunchSource returns the LaunchSource field if non-nil, zero value otherwise.

### GetLaunchSourceOk

`func (o *CreateAndFillByWorkItemsRequest) GetLaunchSourceOk() (*string, bool)`

GetLaunchSourceOk returns a tuple with the LaunchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSource

`func (o *CreateAndFillByWorkItemsRequest) SetLaunchSource(v string)`

SetLaunchSource sets LaunchSource field to given value.

### HasLaunchSource

`func (o *CreateAndFillByWorkItemsRequest) HasLaunchSource() bool`

HasLaunchSource returns a boolean if a field has been set.

### SetLaunchSourceNil

`func (o *CreateAndFillByWorkItemsRequest) SetLaunchSourceNil(b bool)`

 SetLaunchSourceNil sets the value for LaunchSource to be an explicit nil

### UnsetLaunchSource
`func (o *CreateAndFillByWorkItemsRequest) UnsetLaunchSource()`

UnsetLaunchSource ensures that no value is present for LaunchSource, not even an explicit nil
### GetAttachments

`func (o *CreateAndFillByWorkItemsRequest) GetAttachments() []AttachmentPutModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *CreateAndFillByWorkItemsRequest) GetAttachmentsOk() (*[]AttachmentPutModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *CreateAndFillByWorkItemsRequest) SetAttachments(v []AttachmentPutModel)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *CreateAndFillByWorkItemsRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *CreateAndFillByWorkItemsRequest) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *CreateAndFillByWorkItemsRequest) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetLinks

`func (o *CreateAndFillByWorkItemsRequest) GetLinks() []LinkPostModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateAndFillByWorkItemsRequest) GetLinksOk() (*[]LinkPostModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateAndFillByWorkItemsRequest) SetLinks(v []LinkPostModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateAndFillByWorkItemsRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *CreateAndFillByWorkItemsRequest) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *CreateAndFillByWorkItemsRequest) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


