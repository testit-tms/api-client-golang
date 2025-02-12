# TestPointWithLastResultResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**WorkItemName** | Pointer to **NullableString** |  | [optional] 
**IsAutomated** | **bool** |  | 
**TesterId** | Pointer to **NullableString** |  | [optional] 
**WorkItemId** | **string** |  | 
**ConfigurationId** | Pointer to **NullableString** |  | [optional] 
**TestSuiteId** | **string** |  | 
**LastTestResult** | Pointer to [**NullableLastTestResultModel**](LastTestResultModel.md) |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**StatusModel** | Pointer to [**NullableTestStatusApiResult**](TestStatusApiResult.md) |  | [optional] 
**WorkItemGlobalId** | Pointer to **NullableInt64** |  | [optional] 
**WorkItemEntityTypeName** | Pointer to **NullableString** |  | [optional] 
**SectionId** | **string** |  | 
**SectionName** | Pointer to **NullableString** |  | [optional] 
**CreatedDate** | Pointer to **NullableTime** |  | [optional] 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**CreatedById** | **string** |  | 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**TagNames** | Pointer to **[]string** |  | [optional] 
**Duration** | **int32** |  | 
**Priority** | [**WorkItemPriorityModel**](WorkItemPriorityModel.md) |  | 
**TestSuiteNameBreadCrumbs** | Pointer to **[]string** |  | [optional] 
**GroupCount** | Pointer to **NullableInt32** |  | [optional] 
**Iteration** | Pointer to [**NullableIterationModel**](IterationModel.md) |  | [optional] 

## Methods

### NewTestPointWithLastResultResponseModel

`func NewTestPointWithLastResultResponseModel(id string, isAutomated bool, workItemId string, testSuiteId string, sectionId string, createdById string, duration int32, priority WorkItemPriorityModel, ) *TestPointWithLastResultResponseModel`

NewTestPointWithLastResultResponseModel instantiates a new TestPointWithLastResultResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPointWithLastResultResponseModelWithDefaults

`func NewTestPointWithLastResultResponseModelWithDefaults() *TestPointWithLastResultResponseModel`

NewTestPointWithLastResultResponseModelWithDefaults instantiates a new TestPointWithLastResultResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestPointWithLastResultResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestPointWithLastResultResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestPointWithLastResultResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetWorkItemName

`func (o *TestPointWithLastResultResponseModel) GetWorkItemName() string`

GetWorkItemName returns the WorkItemName field if non-nil, zero value otherwise.

### GetWorkItemNameOk

`func (o *TestPointWithLastResultResponseModel) GetWorkItemNameOk() (*string, bool)`

GetWorkItemNameOk returns a tuple with the WorkItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemName

`func (o *TestPointWithLastResultResponseModel) SetWorkItemName(v string)`

SetWorkItemName sets WorkItemName field to given value.

### HasWorkItemName

`func (o *TestPointWithLastResultResponseModel) HasWorkItemName() bool`

HasWorkItemName returns a boolean if a field has been set.

### SetWorkItemNameNil

`func (o *TestPointWithLastResultResponseModel) SetWorkItemNameNil(b bool)`

 SetWorkItemNameNil sets the value for WorkItemName to be an explicit nil

### UnsetWorkItemName
`func (o *TestPointWithLastResultResponseModel) UnsetWorkItemName()`

UnsetWorkItemName ensures that no value is present for WorkItemName, not even an explicit nil
### GetIsAutomated

`func (o *TestPointWithLastResultResponseModel) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *TestPointWithLastResultResponseModel) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *TestPointWithLastResultResponseModel) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.


### GetTesterId

`func (o *TestPointWithLastResultResponseModel) GetTesterId() string`

GetTesterId returns the TesterId field if non-nil, zero value otherwise.

### GetTesterIdOk

`func (o *TestPointWithLastResultResponseModel) GetTesterIdOk() (*string, bool)`

GetTesterIdOk returns a tuple with the TesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesterId

`func (o *TestPointWithLastResultResponseModel) SetTesterId(v string)`

SetTesterId sets TesterId field to given value.

### HasTesterId

`func (o *TestPointWithLastResultResponseModel) HasTesterId() bool`

HasTesterId returns a boolean if a field has been set.

### SetTesterIdNil

`func (o *TestPointWithLastResultResponseModel) SetTesterIdNil(b bool)`

 SetTesterIdNil sets the value for TesterId to be an explicit nil

### UnsetTesterId
`func (o *TestPointWithLastResultResponseModel) UnsetTesterId()`

UnsetTesterId ensures that no value is present for TesterId, not even an explicit nil
### GetWorkItemId

`func (o *TestPointWithLastResultResponseModel) GetWorkItemId() string`

GetWorkItemId returns the WorkItemId field if non-nil, zero value otherwise.

### GetWorkItemIdOk

`func (o *TestPointWithLastResultResponseModel) GetWorkItemIdOk() (*string, bool)`

GetWorkItemIdOk returns a tuple with the WorkItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemId

`func (o *TestPointWithLastResultResponseModel) SetWorkItemId(v string)`

SetWorkItemId sets WorkItemId field to given value.


### GetConfigurationId

`func (o *TestPointWithLastResultResponseModel) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *TestPointWithLastResultResponseModel) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *TestPointWithLastResultResponseModel) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *TestPointWithLastResultResponseModel) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### SetConfigurationIdNil

`func (o *TestPointWithLastResultResponseModel) SetConfigurationIdNil(b bool)`

 SetConfigurationIdNil sets the value for ConfigurationId to be an explicit nil

### UnsetConfigurationId
`func (o *TestPointWithLastResultResponseModel) UnsetConfigurationId()`

UnsetConfigurationId ensures that no value is present for ConfigurationId, not even an explicit nil
### GetTestSuiteId

`func (o *TestPointWithLastResultResponseModel) GetTestSuiteId() string`

GetTestSuiteId returns the TestSuiteId field if non-nil, zero value otherwise.

### GetTestSuiteIdOk

`func (o *TestPointWithLastResultResponseModel) GetTestSuiteIdOk() (*string, bool)`

GetTestSuiteIdOk returns a tuple with the TestSuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuiteId

`func (o *TestPointWithLastResultResponseModel) SetTestSuiteId(v string)`

SetTestSuiteId sets TestSuiteId field to given value.


### GetLastTestResult

`func (o *TestPointWithLastResultResponseModel) GetLastTestResult() LastTestResultModel`

GetLastTestResult returns the LastTestResult field if non-nil, zero value otherwise.

### GetLastTestResultOk

`func (o *TestPointWithLastResultResponseModel) GetLastTestResultOk() (*LastTestResultModel, bool)`

GetLastTestResultOk returns a tuple with the LastTestResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResult

`func (o *TestPointWithLastResultResponseModel) SetLastTestResult(v LastTestResultModel)`

SetLastTestResult sets LastTestResult field to given value.

### HasLastTestResult

`func (o *TestPointWithLastResultResponseModel) HasLastTestResult() bool`

HasLastTestResult returns a boolean if a field has been set.

### SetLastTestResultNil

`func (o *TestPointWithLastResultResponseModel) SetLastTestResultNil(b bool)`

 SetLastTestResultNil sets the value for LastTestResult to be an explicit nil

### UnsetLastTestResult
`func (o *TestPointWithLastResultResponseModel) UnsetLastTestResult()`

UnsetLastTestResult ensures that no value is present for LastTestResult, not even an explicit nil
### GetStatus

`func (o *TestPointWithLastResultResponseModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestPointWithLastResultResponseModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestPointWithLastResultResponseModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestPointWithLastResultResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TestPointWithLastResultResponseModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TestPointWithLastResultResponseModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusModel

`func (o *TestPointWithLastResultResponseModel) GetStatusModel() TestStatusApiResult`

GetStatusModel returns the StatusModel field if non-nil, zero value otherwise.

### GetStatusModelOk

`func (o *TestPointWithLastResultResponseModel) GetStatusModelOk() (*TestStatusApiResult, bool)`

GetStatusModelOk returns a tuple with the StatusModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusModel

`func (o *TestPointWithLastResultResponseModel) SetStatusModel(v TestStatusApiResult)`

SetStatusModel sets StatusModel field to given value.

### HasStatusModel

`func (o *TestPointWithLastResultResponseModel) HasStatusModel() bool`

HasStatusModel returns a boolean if a field has been set.

### SetStatusModelNil

`func (o *TestPointWithLastResultResponseModel) SetStatusModelNil(b bool)`

 SetStatusModelNil sets the value for StatusModel to be an explicit nil

### UnsetStatusModel
`func (o *TestPointWithLastResultResponseModel) UnsetStatusModel()`

UnsetStatusModel ensures that no value is present for StatusModel, not even an explicit nil
### GetWorkItemGlobalId

`func (o *TestPointWithLastResultResponseModel) GetWorkItemGlobalId() int64`

GetWorkItemGlobalId returns the WorkItemGlobalId field if non-nil, zero value otherwise.

### GetWorkItemGlobalIdOk

`func (o *TestPointWithLastResultResponseModel) GetWorkItemGlobalIdOk() (*int64, bool)`

GetWorkItemGlobalIdOk returns a tuple with the WorkItemGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemGlobalId

`func (o *TestPointWithLastResultResponseModel) SetWorkItemGlobalId(v int64)`

SetWorkItemGlobalId sets WorkItemGlobalId field to given value.

### HasWorkItemGlobalId

`func (o *TestPointWithLastResultResponseModel) HasWorkItemGlobalId() bool`

HasWorkItemGlobalId returns a boolean if a field has been set.

### SetWorkItemGlobalIdNil

`func (o *TestPointWithLastResultResponseModel) SetWorkItemGlobalIdNil(b bool)`

 SetWorkItemGlobalIdNil sets the value for WorkItemGlobalId to be an explicit nil

### UnsetWorkItemGlobalId
`func (o *TestPointWithLastResultResponseModel) UnsetWorkItemGlobalId()`

UnsetWorkItemGlobalId ensures that no value is present for WorkItemGlobalId, not even an explicit nil
### GetWorkItemEntityTypeName

`func (o *TestPointWithLastResultResponseModel) GetWorkItemEntityTypeName() string`

GetWorkItemEntityTypeName returns the WorkItemEntityTypeName field if non-nil, zero value otherwise.

### GetWorkItemEntityTypeNameOk

`func (o *TestPointWithLastResultResponseModel) GetWorkItemEntityTypeNameOk() (*string, bool)`

GetWorkItemEntityTypeNameOk returns a tuple with the WorkItemEntityTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemEntityTypeName

`func (o *TestPointWithLastResultResponseModel) SetWorkItemEntityTypeName(v string)`

SetWorkItemEntityTypeName sets WorkItemEntityTypeName field to given value.

### HasWorkItemEntityTypeName

`func (o *TestPointWithLastResultResponseModel) HasWorkItemEntityTypeName() bool`

HasWorkItemEntityTypeName returns a boolean if a field has been set.

### SetWorkItemEntityTypeNameNil

`func (o *TestPointWithLastResultResponseModel) SetWorkItemEntityTypeNameNil(b bool)`

 SetWorkItemEntityTypeNameNil sets the value for WorkItemEntityTypeName to be an explicit nil

### UnsetWorkItemEntityTypeName
`func (o *TestPointWithLastResultResponseModel) UnsetWorkItemEntityTypeName()`

UnsetWorkItemEntityTypeName ensures that no value is present for WorkItemEntityTypeName, not even an explicit nil
### GetSectionId

`func (o *TestPointWithLastResultResponseModel) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *TestPointWithLastResultResponseModel) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *TestPointWithLastResultResponseModel) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.


### GetSectionName

`func (o *TestPointWithLastResultResponseModel) GetSectionName() string`

GetSectionName returns the SectionName field if non-nil, zero value otherwise.

### GetSectionNameOk

`func (o *TestPointWithLastResultResponseModel) GetSectionNameOk() (*string, bool)`

GetSectionNameOk returns a tuple with the SectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionName

`func (o *TestPointWithLastResultResponseModel) SetSectionName(v string)`

SetSectionName sets SectionName field to given value.

### HasSectionName

`func (o *TestPointWithLastResultResponseModel) HasSectionName() bool`

HasSectionName returns a boolean if a field has been set.

### SetSectionNameNil

`func (o *TestPointWithLastResultResponseModel) SetSectionNameNil(b bool)`

 SetSectionNameNil sets the value for SectionName to be an explicit nil

### UnsetSectionName
`func (o *TestPointWithLastResultResponseModel) UnsetSectionName()`

UnsetSectionName ensures that no value is present for SectionName, not even an explicit nil
### GetCreatedDate

`func (o *TestPointWithLastResultResponseModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestPointWithLastResultResponseModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestPointWithLastResultResponseModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *TestPointWithLastResultResponseModel) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *TestPointWithLastResultResponseModel) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *TestPointWithLastResultResponseModel) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetModifiedDate

`func (o *TestPointWithLastResultResponseModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestPointWithLastResultResponseModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestPointWithLastResultResponseModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestPointWithLastResultResponseModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *TestPointWithLastResultResponseModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *TestPointWithLastResultResponseModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *TestPointWithLastResultResponseModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TestPointWithLastResultResponseModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TestPointWithLastResultResponseModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedById

`func (o *TestPointWithLastResultResponseModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *TestPointWithLastResultResponseModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *TestPointWithLastResultResponseModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *TestPointWithLastResultResponseModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *TestPointWithLastResultResponseModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *TestPointWithLastResultResponseModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetAttributes

`func (o *TestPointWithLastResultResponseModel) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TestPointWithLastResultResponseModel) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TestPointWithLastResultResponseModel) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TestPointWithLastResultResponseModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *TestPointWithLastResultResponseModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *TestPointWithLastResultResponseModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetTagNames

`func (o *TestPointWithLastResultResponseModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *TestPointWithLastResultResponseModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *TestPointWithLastResultResponseModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *TestPointWithLastResultResponseModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### SetTagNamesNil

`func (o *TestPointWithLastResultResponseModel) SetTagNamesNil(b bool)`

 SetTagNamesNil sets the value for TagNames to be an explicit nil

### UnsetTagNames
`func (o *TestPointWithLastResultResponseModel) UnsetTagNames()`

UnsetTagNames ensures that no value is present for TagNames, not even an explicit nil
### GetDuration

`func (o *TestPointWithLastResultResponseModel) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TestPointWithLastResultResponseModel) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TestPointWithLastResultResponseModel) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetPriority

`func (o *TestPointWithLastResultResponseModel) GetPriority() WorkItemPriorityModel`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TestPointWithLastResultResponseModel) GetPriorityOk() (*WorkItemPriorityModel, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TestPointWithLastResultResponseModel) SetPriority(v WorkItemPriorityModel)`

SetPriority sets Priority field to given value.


### GetTestSuiteNameBreadCrumbs

`func (o *TestPointWithLastResultResponseModel) GetTestSuiteNameBreadCrumbs() []string`

GetTestSuiteNameBreadCrumbs returns the TestSuiteNameBreadCrumbs field if non-nil, zero value otherwise.

### GetTestSuiteNameBreadCrumbsOk

`func (o *TestPointWithLastResultResponseModel) GetTestSuiteNameBreadCrumbsOk() (*[]string, bool)`

GetTestSuiteNameBreadCrumbsOk returns a tuple with the TestSuiteNameBreadCrumbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuiteNameBreadCrumbs

`func (o *TestPointWithLastResultResponseModel) SetTestSuiteNameBreadCrumbs(v []string)`

SetTestSuiteNameBreadCrumbs sets TestSuiteNameBreadCrumbs field to given value.

### HasTestSuiteNameBreadCrumbs

`func (o *TestPointWithLastResultResponseModel) HasTestSuiteNameBreadCrumbs() bool`

HasTestSuiteNameBreadCrumbs returns a boolean if a field has been set.

### SetTestSuiteNameBreadCrumbsNil

`func (o *TestPointWithLastResultResponseModel) SetTestSuiteNameBreadCrumbsNil(b bool)`

 SetTestSuiteNameBreadCrumbsNil sets the value for TestSuiteNameBreadCrumbs to be an explicit nil

### UnsetTestSuiteNameBreadCrumbs
`func (o *TestPointWithLastResultResponseModel) UnsetTestSuiteNameBreadCrumbs()`

UnsetTestSuiteNameBreadCrumbs ensures that no value is present for TestSuiteNameBreadCrumbs, not even an explicit nil
### GetGroupCount

`func (o *TestPointWithLastResultResponseModel) GetGroupCount() int32`

GetGroupCount returns the GroupCount field if non-nil, zero value otherwise.

### GetGroupCountOk

`func (o *TestPointWithLastResultResponseModel) GetGroupCountOk() (*int32, bool)`

GetGroupCountOk returns a tuple with the GroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCount

`func (o *TestPointWithLastResultResponseModel) SetGroupCount(v int32)`

SetGroupCount sets GroupCount field to given value.

### HasGroupCount

`func (o *TestPointWithLastResultResponseModel) HasGroupCount() bool`

HasGroupCount returns a boolean if a field has been set.

### SetGroupCountNil

`func (o *TestPointWithLastResultResponseModel) SetGroupCountNil(b bool)`

 SetGroupCountNil sets the value for GroupCount to be an explicit nil

### UnsetGroupCount
`func (o *TestPointWithLastResultResponseModel) UnsetGroupCount()`

UnsetGroupCount ensures that no value is present for GroupCount, not even an explicit nil
### GetIteration

`func (o *TestPointWithLastResultResponseModel) GetIteration() IterationModel`

GetIteration returns the Iteration field if non-nil, zero value otherwise.

### GetIterationOk

`func (o *TestPointWithLastResultResponseModel) GetIterationOk() (*IterationModel, bool)`

GetIterationOk returns a tuple with the Iteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIteration

`func (o *TestPointWithLastResultResponseModel) SetIteration(v IterationModel)`

SetIteration sets Iteration field to given value.

### HasIteration

`func (o *TestPointWithLastResultResponseModel) HasIteration() bool`

HasIteration returns a boolean if a field has been set.

### SetIterationNil

`func (o *TestPointWithLastResultResponseModel) SetIterationNil(b bool)`

 SetIterationNil sets the value for Iteration to be an explicit nil

### UnsetIteration
`func (o *TestPointWithLastResultResponseModel) UnsetIteration()`

UnsetIteration ensures that no value is present for Iteration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


