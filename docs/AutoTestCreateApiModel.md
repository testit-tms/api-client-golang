# AutoTestCreateApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | Unique ID of the autotest project | 
**ExternalId** | **string** | External ID of the autotest | 
**ExternalKey** | Pointer to **NullableString** | External key of the autotest | [optional] 
**Name** | **string** | Name of the autotest | 
**Namespace** | Pointer to **NullableString** | Name of the autotest namespace | [optional] 
**Classname** | Pointer to **NullableString** | Name of the autotest class | [optional] 
**Title** | Pointer to **NullableString** | Name of the autotest in autotest&#39;s card | [optional] 
**Description** | Pointer to **NullableString** | Description of the autotest in autotest&#39;s card | [optional] 
**IsFlaky** | Pointer to **NullableBool** | Indicates if the autotest is marked as flaky | [optional] 
**Steps** | Pointer to [**[]AutoTestStepApiModel**](AutoTestStepApiModel.md) | Collection of the autotest steps | [optional] 
**Setup** | Pointer to [**[]AutoTestStepApiModel**](AutoTestStepApiModel.md) | Collection of the autotest setup steps | [optional] 
**Teardown** | Pointer to [**[]AutoTestStepApiModel**](AutoTestStepApiModel.md) | Collection of the autotest teardown steps | [optional] 
**ShouldCreateWorkItem** | Pointer to **NullableBool** | Creates a test case linked to the autotest. | [optional] 
**WorkItemIds** | Pointer to **[]string** | Specifies the IDs of work items to link your autotest to. You can specify several IDs. | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | Key value pair of custom work item attributes | [optional] 
**WorkItemIdsForLinkWithAutoTest** | Pointer to **[]string** | Specifies the IDs of work items to link your autotest to. You can specify several IDs. | [optional] 
**Labels** | Pointer to [**[]LabelApiModel**](LabelApiModel.md) | Collection of the autotest labels | [optional] 
**Links** | Pointer to [**[]LinkCreateApiModel**](LinkCreateApiModel.md) | Collection of the autotest links | [optional] 
**Tags** | Pointer to **[]string** | Collection of the autotest tags | [optional] 

## Methods

### NewAutoTestCreateApiModel

`func NewAutoTestCreateApiModel(projectId string, externalId string, name string, ) *AutoTestCreateApiModel`

NewAutoTestCreateApiModel instantiates a new AutoTestCreateApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestCreateApiModelWithDefaults

`func NewAutoTestCreateApiModelWithDefaults() *AutoTestCreateApiModel`

NewAutoTestCreateApiModelWithDefaults instantiates a new AutoTestCreateApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *AutoTestCreateApiModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AutoTestCreateApiModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AutoTestCreateApiModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetExternalId

`func (o *AutoTestCreateApiModel) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AutoTestCreateApiModel) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AutoTestCreateApiModel) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetExternalKey

`func (o *AutoTestCreateApiModel) GetExternalKey() string`

GetExternalKey returns the ExternalKey field if non-nil, zero value otherwise.

### GetExternalKeyOk

`func (o *AutoTestCreateApiModel) GetExternalKeyOk() (*string, bool)`

GetExternalKeyOk returns a tuple with the ExternalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalKey

`func (o *AutoTestCreateApiModel) SetExternalKey(v string)`

SetExternalKey sets ExternalKey field to given value.

### HasExternalKey

`func (o *AutoTestCreateApiModel) HasExternalKey() bool`

HasExternalKey returns a boolean if a field has been set.

### SetExternalKeyNil

`func (o *AutoTestCreateApiModel) SetExternalKeyNil(b bool)`

 SetExternalKeyNil sets the value for ExternalKey to be an explicit nil

### UnsetExternalKey
`func (o *AutoTestCreateApiModel) UnsetExternalKey()`

UnsetExternalKey ensures that no value is present for ExternalKey, not even an explicit nil
### GetName

`func (o *AutoTestCreateApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoTestCreateApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoTestCreateApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *AutoTestCreateApiModel) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AutoTestCreateApiModel) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AutoTestCreateApiModel) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AutoTestCreateApiModel) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *AutoTestCreateApiModel) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *AutoTestCreateApiModel) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetClassname

`func (o *AutoTestCreateApiModel) GetClassname() string`

GetClassname returns the Classname field if non-nil, zero value otherwise.

### GetClassnameOk

`func (o *AutoTestCreateApiModel) GetClassnameOk() (*string, bool)`

GetClassnameOk returns a tuple with the Classname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassname

`func (o *AutoTestCreateApiModel) SetClassname(v string)`

SetClassname sets Classname field to given value.

### HasClassname

`func (o *AutoTestCreateApiModel) HasClassname() bool`

HasClassname returns a boolean if a field has been set.

### SetClassnameNil

`func (o *AutoTestCreateApiModel) SetClassnameNil(b bool)`

 SetClassnameNil sets the value for Classname to be an explicit nil

### UnsetClassname
`func (o *AutoTestCreateApiModel) UnsetClassname()`

UnsetClassname ensures that no value is present for Classname, not even an explicit nil
### GetTitle

`func (o *AutoTestCreateApiModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AutoTestCreateApiModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AutoTestCreateApiModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AutoTestCreateApiModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AutoTestCreateApiModel) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AutoTestCreateApiModel) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *AutoTestCreateApiModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutoTestCreateApiModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutoTestCreateApiModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutoTestCreateApiModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AutoTestCreateApiModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AutoTestCreateApiModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsFlaky

`func (o *AutoTestCreateApiModel) GetIsFlaky() bool`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *AutoTestCreateApiModel) GetIsFlakyOk() (*bool, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *AutoTestCreateApiModel) SetIsFlaky(v bool)`

SetIsFlaky sets IsFlaky field to given value.

### HasIsFlaky

`func (o *AutoTestCreateApiModel) HasIsFlaky() bool`

HasIsFlaky returns a boolean if a field has been set.

### SetIsFlakyNil

`func (o *AutoTestCreateApiModel) SetIsFlakyNil(b bool)`

 SetIsFlakyNil sets the value for IsFlaky to be an explicit nil

### UnsetIsFlaky
`func (o *AutoTestCreateApiModel) UnsetIsFlaky()`

UnsetIsFlaky ensures that no value is present for IsFlaky, not even an explicit nil
### GetSteps

`func (o *AutoTestCreateApiModel) GetSteps() []AutoTestStepApiModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *AutoTestCreateApiModel) GetStepsOk() (*[]AutoTestStepApiModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *AutoTestCreateApiModel) SetSteps(v []AutoTestStepApiModel)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *AutoTestCreateApiModel) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *AutoTestCreateApiModel) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *AutoTestCreateApiModel) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil
### GetSetup

`func (o *AutoTestCreateApiModel) GetSetup() []AutoTestStepApiModel`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *AutoTestCreateApiModel) GetSetupOk() (*[]AutoTestStepApiModel, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *AutoTestCreateApiModel) SetSetup(v []AutoTestStepApiModel)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *AutoTestCreateApiModel) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### SetSetupNil

`func (o *AutoTestCreateApiModel) SetSetupNil(b bool)`

 SetSetupNil sets the value for Setup to be an explicit nil

### UnsetSetup
`func (o *AutoTestCreateApiModel) UnsetSetup()`

UnsetSetup ensures that no value is present for Setup, not even an explicit nil
### GetTeardown

`func (o *AutoTestCreateApiModel) GetTeardown() []AutoTestStepApiModel`

GetTeardown returns the Teardown field if non-nil, zero value otherwise.

### GetTeardownOk

`func (o *AutoTestCreateApiModel) GetTeardownOk() (*[]AutoTestStepApiModel, bool)`

GetTeardownOk returns a tuple with the Teardown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardown

`func (o *AutoTestCreateApiModel) SetTeardown(v []AutoTestStepApiModel)`

SetTeardown sets Teardown field to given value.

### HasTeardown

`func (o *AutoTestCreateApiModel) HasTeardown() bool`

HasTeardown returns a boolean if a field has been set.

### SetTeardownNil

`func (o *AutoTestCreateApiModel) SetTeardownNil(b bool)`

 SetTeardownNil sets the value for Teardown to be an explicit nil

### UnsetTeardown
`func (o *AutoTestCreateApiModel) UnsetTeardown()`

UnsetTeardown ensures that no value is present for Teardown, not even an explicit nil
### GetShouldCreateWorkItem

`func (o *AutoTestCreateApiModel) GetShouldCreateWorkItem() bool`

GetShouldCreateWorkItem returns the ShouldCreateWorkItem field if non-nil, zero value otherwise.

### GetShouldCreateWorkItemOk

`func (o *AutoTestCreateApiModel) GetShouldCreateWorkItemOk() (*bool, bool)`

GetShouldCreateWorkItemOk returns a tuple with the ShouldCreateWorkItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCreateWorkItem

`func (o *AutoTestCreateApiModel) SetShouldCreateWorkItem(v bool)`

SetShouldCreateWorkItem sets ShouldCreateWorkItem field to given value.

### HasShouldCreateWorkItem

`func (o *AutoTestCreateApiModel) HasShouldCreateWorkItem() bool`

HasShouldCreateWorkItem returns a boolean if a field has been set.

### SetShouldCreateWorkItemNil

`func (o *AutoTestCreateApiModel) SetShouldCreateWorkItemNil(b bool)`

 SetShouldCreateWorkItemNil sets the value for ShouldCreateWorkItem to be an explicit nil

### UnsetShouldCreateWorkItem
`func (o *AutoTestCreateApiModel) UnsetShouldCreateWorkItem()`

UnsetShouldCreateWorkItem ensures that no value is present for ShouldCreateWorkItem, not even an explicit nil
### GetWorkItemIds

`func (o *AutoTestCreateApiModel) GetWorkItemIds() []string`

GetWorkItemIds returns the WorkItemIds field if non-nil, zero value otherwise.

### GetWorkItemIdsOk

`func (o *AutoTestCreateApiModel) GetWorkItemIdsOk() (*[]string, bool)`

GetWorkItemIdsOk returns a tuple with the WorkItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemIds

`func (o *AutoTestCreateApiModel) SetWorkItemIds(v []string)`

SetWorkItemIds sets WorkItemIds field to given value.

### HasWorkItemIds

`func (o *AutoTestCreateApiModel) HasWorkItemIds() bool`

HasWorkItemIds returns a boolean if a field has been set.

### SetWorkItemIdsNil

`func (o *AutoTestCreateApiModel) SetWorkItemIdsNil(b bool)`

 SetWorkItemIdsNil sets the value for WorkItemIds to be an explicit nil

### UnsetWorkItemIds
`func (o *AutoTestCreateApiModel) UnsetWorkItemIds()`

UnsetWorkItemIds ensures that no value is present for WorkItemIds, not even an explicit nil
### GetAttributes

`func (o *AutoTestCreateApiModel) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AutoTestCreateApiModel) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AutoTestCreateApiModel) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AutoTestCreateApiModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *AutoTestCreateApiModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *AutoTestCreateApiModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetWorkItemIdsForLinkWithAutoTest

`func (o *AutoTestCreateApiModel) GetWorkItemIdsForLinkWithAutoTest() []string`

GetWorkItemIdsForLinkWithAutoTest returns the WorkItemIdsForLinkWithAutoTest field if non-nil, zero value otherwise.

### GetWorkItemIdsForLinkWithAutoTestOk

`func (o *AutoTestCreateApiModel) GetWorkItemIdsForLinkWithAutoTestOk() (*[]string, bool)`

GetWorkItemIdsForLinkWithAutoTestOk returns a tuple with the WorkItemIdsForLinkWithAutoTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemIdsForLinkWithAutoTest

`func (o *AutoTestCreateApiModel) SetWorkItemIdsForLinkWithAutoTest(v []string)`

SetWorkItemIdsForLinkWithAutoTest sets WorkItemIdsForLinkWithAutoTest field to given value.

### HasWorkItemIdsForLinkWithAutoTest

`func (o *AutoTestCreateApiModel) HasWorkItemIdsForLinkWithAutoTest() bool`

HasWorkItemIdsForLinkWithAutoTest returns a boolean if a field has been set.

### SetWorkItemIdsForLinkWithAutoTestNil

`func (o *AutoTestCreateApiModel) SetWorkItemIdsForLinkWithAutoTestNil(b bool)`

 SetWorkItemIdsForLinkWithAutoTestNil sets the value for WorkItemIdsForLinkWithAutoTest to be an explicit nil

### UnsetWorkItemIdsForLinkWithAutoTest
`func (o *AutoTestCreateApiModel) UnsetWorkItemIdsForLinkWithAutoTest()`

UnsetWorkItemIdsForLinkWithAutoTest ensures that no value is present for WorkItemIdsForLinkWithAutoTest, not even an explicit nil
### GetLabels

`func (o *AutoTestCreateApiModel) GetLabels() []LabelApiModel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AutoTestCreateApiModel) GetLabelsOk() (*[]LabelApiModel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AutoTestCreateApiModel) SetLabels(v []LabelApiModel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AutoTestCreateApiModel) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *AutoTestCreateApiModel) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *AutoTestCreateApiModel) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetLinks

`func (o *AutoTestCreateApiModel) GetLinks() []LinkCreateApiModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AutoTestCreateApiModel) GetLinksOk() (*[]LinkCreateApiModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AutoTestCreateApiModel) SetLinks(v []LinkCreateApiModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AutoTestCreateApiModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *AutoTestCreateApiModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *AutoTestCreateApiModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetTags

`func (o *AutoTestCreateApiModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AutoTestCreateApiModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AutoTestCreateApiModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AutoTestCreateApiModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AutoTestCreateApiModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AutoTestCreateApiModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


