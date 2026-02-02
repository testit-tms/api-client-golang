# AutoTestUpdateApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Autotest unique internal identifier | [optional] 
**ExternalId** | **string** | External ID of the autotest | 
**ExternalKey** | Pointer to **NullableString** | External key of the autotest | [optional] 
**ProjectId** | **string** | Unique ID of the autotest project | 
**Name** | **string** | Name of the autotest | 
**Namespace** | Pointer to **NullableString** | Name of the autotest namespace | [optional] 
**Classname** | Pointer to **NullableString** | Name of the autotest class | [optional] 
**Steps** | Pointer to [**[]AutoTestStepApiModel**](AutoTestStepApiModel.md) | Collection of the autotest steps | [optional] 
**Setup** | Pointer to [**[]AutoTestStepApiModel**](AutoTestStepApiModel.md) | Collection of the autotest setup steps | [optional] 
**Teardown** | Pointer to [**[]AutoTestStepApiModel**](AutoTestStepApiModel.md) | Collection of the autotest teardown steps | [optional] 
**Title** | Pointer to **NullableString** | Name of the autotest in autotest&#39;s card | [optional] 
**Description** | Pointer to **NullableString** | Description of the autotest in autotest&#39;s card | [optional] 
**Labels** | Pointer to [**[]LabelApiModel**](LabelApiModel.md) | Collection of the autotest labels | [optional] 
**Links** | Pointer to [**[]LinkUpdateApiModel**](LinkUpdateApiModel.md) | Collection of the autotest links | [optional] 
**IsFlaky** | Pointer to **NullableBool** | Indicates if the autotest is marked as flaky | [optional] 
**WorkItemIdsForLinkWithAutoTest** | Pointer to **[]string** | Specifies the IDs of work items to link your autotest to. You can specify several IDs. | [optional] 
**WorkItemIds** | Pointer to **[]string** | Specifies the IDs of work items to link your autotest to. You can specify several IDs. | [optional] 

## Methods

### NewAutoTestUpdateApiModel

`func NewAutoTestUpdateApiModel(externalId string, projectId string, name string, ) *AutoTestUpdateApiModel`

NewAutoTestUpdateApiModel instantiates a new AutoTestUpdateApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestUpdateApiModelWithDefaults

`func NewAutoTestUpdateApiModelWithDefaults() *AutoTestUpdateApiModel`

NewAutoTestUpdateApiModelWithDefaults instantiates a new AutoTestUpdateApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoTestUpdateApiModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoTestUpdateApiModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoTestUpdateApiModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutoTestUpdateApiModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AutoTestUpdateApiModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AutoTestUpdateApiModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetExternalId

`func (o *AutoTestUpdateApiModel) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AutoTestUpdateApiModel) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AutoTestUpdateApiModel) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetExternalKey

`func (o *AutoTestUpdateApiModel) GetExternalKey() string`

GetExternalKey returns the ExternalKey field if non-nil, zero value otherwise.

### GetExternalKeyOk

`func (o *AutoTestUpdateApiModel) GetExternalKeyOk() (*string, bool)`

GetExternalKeyOk returns a tuple with the ExternalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalKey

`func (o *AutoTestUpdateApiModel) SetExternalKey(v string)`

SetExternalKey sets ExternalKey field to given value.

### HasExternalKey

`func (o *AutoTestUpdateApiModel) HasExternalKey() bool`

HasExternalKey returns a boolean if a field has been set.

### SetExternalKeyNil

`func (o *AutoTestUpdateApiModel) SetExternalKeyNil(b bool)`

 SetExternalKeyNil sets the value for ExternalKey to be an explicit nil

### UnsetExternalKey
`func (o *AutoTestUpdateApiModel) UnsetExternalKey()`

UnsetExternalKey ensures that no value is present for ExternalKey, not even an explicit nil
### GetProjectId

`func (o *AutoTestUpdateApiModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AutoTestUpdateApiModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AutoTestUpdateApiModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *AutoTestUpdateApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoTestUpdateApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoTestUpdateApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *AutoTestUpdateApiModel) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AutoTestUpdateApiModel) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AutoTestUpdateApiModel) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AutoTestUpdateApiModel) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *AutoTestUpdateApiModel) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *AutoTestUpdateApiModel) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetClassname

`func (o *AutoTestUpdateApiModel) GetClassname() string`

GetClassname returns the Classname field if non-nil, zero value otherwise.

### GetClassnameOk

`func (o *AutoTestUpdateApiModel) GetClassnameOk() (*string, bool)`

GetClassnameOk returns a tuple with the Classname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassname

`func (o *AutoTestUpdateApiModel) SetClassname(v string)`

SetClassname sets Classname field to given value.

### HasClassname

`func (o *AutoTestUpdateApiModel) HasClassname() bool`

HasClassname returns a boolean if a field has been set.

### SetClassnameNil

`func (o *AutoTestUpdateApiModel) SetClassnameNil(b bool)`

 SetClassnameNil sets the value for Classname to be an explicit nil

### UnsetClassname
`func (o *AutoTestUpdateApiModel) UnsetClassname()`

UnsetClassname ensures that no value is present for Classname, not even an explicit nil
### GetSteps

`func (o *AutoTestUpdateApiModel) GetSteps() []AutoTestStepApiModel`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *AutoTestUpdateApiModel) GetStepsOk() (*[]AutoTestStepApiModel, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *AutoTestUpdateApiModel) SetSteps(v []AutoTestStepApiModel)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *AutoTestUpdateApiModel) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *AutoTestUpdateApiModel) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *AutoTestUpdateApiModel) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil
### GetSetup

`func (o *AutoTestUpdateApiModel) GetSetup() []AutoTestStepApiModel`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *AutoTestUpdateApiModel) GetSetupOk() (*[]AutoTestStepApiModel, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *AutoTestUpdateApiModel) SetSetup(v []AutoTestStepApiModel)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *AutoTestUpdateApiModel) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### SetSetupNil

`func (o *AutoTestUpdateApiModel) SetSetupNil(b bool)`

 SetSetupNil sets the value for Setup to be an explicit nil

### UnsetSetup
`func (o *AutoTestUpdateApiModel) UnsetSetup()`

UnsetSetup ensures that no value is present for Setup, not even an explicit nil
### GetTeardown

`func (o *AutoTestUpdateApiModel) GetTeardown() []AutoTestStepApiModel`

GetTeardown returns the Teardown field if non-nil, zero value otherwise.

### GetTeardownOk

`func (o *AutoTestUpdateApiModel) GetTeardownOk() (*[]AutoTestStepApiModel, bool)`

GetTeardownOk returns a tuple with the Teardown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardown

`func (o *AutoTestUpdateApiModel) SetTeardown(v []AutoTestStepApiModel)`

SetTeardown sets Teardown field to given value.

### HasTeardown

`func (o *AutoTestUpdateApiModel) HasTeardown() bool`

HasTeardown returns a boolean if a field has been set.

### SetTeardownNil

`func (o *AutoTestUpdateApiModel) SetTeardownNil(b bool)`

 SetTeardownNil sets the value for Teardown to be an explicit nil

### UnsetTeardown
`func (o *AutoTestUpdateApiModel) UnsetTeardown()`

UnsetTeardown ensures that no value is present for Teardown, not even an explicit nil
### GetTitle

`func (o *AutoTestUpdateApiModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AutoTestUpdateApiModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AutoTestUpdateApiModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AutoTestUpdateApiModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AutoTestUpdateApiModel) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AutoTestUpdateApiModel) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *AutoTestUpdateApiModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutoTestUpdateApiModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutoTestUpdateApiModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutoTestUpdateApiModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AutoTestUpdateApiModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AutoTestUpdateApiModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *AutoTestUpdateApiModel) GetLabels() []LabelApiModel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AutoTestUpdateApiModel) GetLabelsOk() (*[]LabelApiModel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AutoTestUpdateApiModel) SetLabels(v []LabelApiModel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AutoTestUpdateApiModel) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *AutoTestUpdateApiModel) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *AutoTestUpdateApiModel) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetLinks

`func (o *AutoTestUpdateApiModel) GetLinks() []LinkUpdateApiModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AutoTestUpdateApiModel) GetLinksOk() (*[]LinkUpdateApiModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AutoTestUpdateApiModel) SetLinks(v []LinkUpdateApiModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AutoTestUpdateApiModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *AutoTestUpdateApiModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *AutoTestUpdateApiModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetIsFlaky

`func (o *AutoTestUpdateApiModel) GetIsFlaky() bool`

GetIsFlaky returns the IsFlaky field if non-nil, zero value otherwise.

### GetIsFlakyOk

`func (o *AutoTestUpdateApiModel) GetIsFlakyOk() (*bool, bool)`

GetIsFlakyOk returns a tuple with the IsFlaky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlaky

`func (o *AutoTestUpdateApiModel) SetIsFlaky(v bool)`

SetIsFlaky sets IsFlaky field to given value.

### HasIsFlaky

`func (o *AutoTestUpdateApiModel) HasIsFlaky() bool`

HasIsFlaky returns a boolean if a field has been set.

### SetIsFlakyNil

`func (o *AutoTestUpdateApiModel) SetIsFlakyNil(b bool)`

 SetIsFlakyNil sets the value for IsFlaky to be an explicit nil

### UnsetIsFlaky
`func (o *AutoTestUpdateApiModel) UnsetIsFlaky()`

UnsetIsFlaky ensures that no value is present for IsFlaky, not even an explicit nil
### GetWorkItemIdsForLinkWithAutoTest

`func (o *AutoTestUpdateApiModel) GetWorkItemIdsForLinkWithAutoTest() []string`

GetWorkItemIdsForLinkWithAutoTest returns the WorkItemIdsForLinkWithAutoTest field if non-nil, zero value otherwise.

### GetWorkItemIdsForLinkWithAutoTestOk

`func (o *AutoTestUpdateApiModel) GetWorkItemIdsForLinkWithAutoTestOk() (*[]string, bool)`

GetWorkItemIdsForLinkWithAutoTestOk returns a tuple with the WorkItemIdsForLinkWithAutoTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemIdsForLinkWithAutoTest

`func (o *AutoTestUpdateApiModel) SetWorkItemIdsForLinkWithAutoTest(v []string)`

SetWorkItemIdsForLinkWithAutoTest sets WorkItemIdsForLinkWithAutoTest field to given value.

### HasWorkItemIdsForLinkWithAutoTest

`func (o *AutoTestUpdateApiModel) HasWorkItemIdsForLinkWithAutoTest() bool`

HasWorkItemIdsForLinkWithAutoTest returns a boolean if a field has been set.

### SetWorkItemIdsForLinkWithAutoTestNil

`func (o *AutoTestUpdateApiModel) SetWorkItemIdsForLinkWithAutoTestNil(b bool)`

 SetWorkItemIdsForLinkWithAutoTestNil sets the value for WorkItemIdsForLinkWithAutoTest to be an explicit nil

### UnsetWorkItemIdsForLinkWithAutoTest
`func (o *AutoTestUpdateApiModel) UnsetWorkItemIdsForLinkWithAutoTest()`

UnsetWorkItemIdsForLinkWithAutoTest ensures that no value is present for WorkItemIdsForLinkWithAutoTest, not even an explicit nil
### GetWorkItemIds

`func (o *AutoTestUpdateApiModel) GetWorkItemIds() []string`

GetWorkItemIds returns the WorkItemIds field if non-nil, zero value otherwise.

### GetWorkItemIdsOk

`func (o *AutoTestUpdateApiModel) GetWorkItemIdsOk() (*[]string, bool)`

GetWorkItemIdsOk returns a tuple with the WorkItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemIds

`func (o *AutoTestUpdateApiModel) SetWorkItemIds(v []string)`

SetWorkItemIds sets WorkItemIds field to given value.

### HasWorkItemIds

`func (o *AutoTestUpdateApiModel) HasWorkItemIds() bool`

HasWorkItemIds returns a boolean if a field has been set.

### SetWorkItemIdsNil

`func (o *AutoTestUpdateApiModel) SetWorkItemIdsNil(b bool)`

 SetWorkItemIdsNil sets the value for WorkItemIds to be an explicit nil

### UnsetWorkItemIds
`func (o *AutoTestUpdateApiModel) UnsetWorkItemIds()`

UnsetWorkItemIds ensures that no value is present for WorkItemIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


