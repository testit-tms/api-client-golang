# TestPointShortGetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the test point | 
**CreatedDate** | **time.Time** | Creation date of the test point | 
**CreatedById** | **string** | Unique ID of the test point creator | 
**ModifiedDate** | Pointer to **NullableTime** | Last modification date of the test point | [optional] 
**ModifiedById** | Pointer to **NullableString** | Unique ID of the test point last editor | [optional] 
**TesterId** | Pointer to **NullableString** | Unique ID of the test point assigned user | [optional] 
**Parameters** | Pointer to **map[string]string** | Collection of the test point parameters | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | Collection of attributes of work item the test point represents | [optional] 
**Tags** | Pointer to **[]string** | Collection of the test point tags | [optional] 
**Links** | Pointer to **[]string** | Collection of the test point links | [optional] 
**TestSuiteId** | **string** | Unique ID of test suite the test point assigned to | 
**WorkItemId** | **string** | Unique ID of work item the test point represents | 
**WorkItemGlobalId** | **int64** | Global ID of work item the test point represents | 
**WorkItemVersionId** | **string** | Unique ID of work item version the test point represents | 
**WorkItemMedianDuration** | Pointer to **NullableInt64** | Median duration of work item the test point represents | [optional] 
**Status** | [**TestPointStatus**](TestPointStatus.md) |  | 
**Priority** | [**WorkItemPriorityModel**](WorkItemPriorityModel.md) |  | 
**IsAutomated** | **bool** | Indicates if the test point represents an autotest | 
**Name** | Pointer to **NullableString** | Name of the test point | [optional] 
**ConfigurationId** | **string** | Unique ID of the test point configuration | 
**Duration** | **int32** | Duration of the test point | 
**SectionId** | **string** | Unique ID of section where work item the test point represents is located | 
**SectionName** | Pointer to **NullableString** | Name of section where work item the test point represents is located | [optional] 
**ProjectId** | **string** | Unique ID of the test point project | 
**LastTestResult** | [**TestPointShortGetModelLastTestResult**](TestPointShortGetModelLastTestResult.md) |  | 
**IterationId** | **string** | Unique ID of work item iteration the test point represents | 
**WorkItemState** | [**WorkItemState**](WorkItemState.md) |  | 
**WorkItemCreatedById** | **string** | Unique ID of the work item creator | 
**WorkItemCreatedDate** | **time.Time** | Creation date of work item | 
**WorkItemModifiedById** | Pointer to **NullableString** | Unique ID of the work item last editor | [optional] 
**WorkItemModifiedDate** | Pointer to **NullableTime** | Modified date of work item | [optional] 

## Methods

### NewTestPointShortGetModel

`func NewTestPointShortGetModel(id string, createdDate time.Time, createdById string, testSuiteId string, workItemId string, workItemGlobalId int64, workItemVersionId string, status TestPointStatus, priority WorkItemPriorityModel, isAutomated bool, configurationId string, duration int32, sectionId string, projectId string, lastTestResult TestPointShortGetModelLastTestResult, iterationId string, workItemState WorkItemState, workItemCreatedById string, workItemCreatedDate time.Time, ) *TestPointShortGetModel`

NewTestPointShortGetModel instantiates a new TestPointShortGetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPointShortGetModelWithDefaults

`func NewTestPointShortGetModelWithDefaults() *TestPointShortGetModel`

NewTestPointShortGetModelWithDefaults instantiates a new TestPointShortGetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestPointShortGetModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestPointShortGetModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestPointShortGetModel) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedDate

`func (o *TestPointShortGetModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestPointShortGetModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestPointShortGetModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetCreatedById

`func (o *TestPointShortGetModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TestPointShortGetModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TestPointShortGetModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedDate

`func (o *TestPointShortGetModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestPointShortGetModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestPointShortGetModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestPointShortGetModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *TestPointShortGetModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *TestPointShortGetModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedById

`func (o *TestPointShortGetModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *TestPointShortGetModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *TestPointShortGetModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *TestPointShortGetModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *TestPointShortGetModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *TestPointShortGetModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetTesterId

`func (o *TestPointShortGetModel) GetTesterId() string`

GetTesterId returns the TesterId field if non-nil, zero value otherwise.

### GetTesterIdOk

`func (o *TestPointShortGetModel) GetTesterIdOk() (*string, bool)`

GetTesterIdOk returns a tuple with the TesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesterId

`func (o *TestPointShortGetModel) SetTesterId(v string)`

SetTesterId sets TesterId field to given value.

### HasTesterId

`func (o *TestPointShortGetModel) HasTesterId() bool`

HasTesterId returns a boolean if a field has been set.

### SetTesterIdNil

`func (o *TestPointShortGetModel) SetTesterIdNil(b bool)`

 SetTesterIdNil sets the value for TesterId to be an explicit nil

### UnsetTesterId
`func (o *TestPointShortGetModel) UnsetTesterId()`

UnsetTesterId ensures that no value is present for TesterId, not even an explicit nil
### GetParameters

`func (o *TestPointShortGetModel) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TestPointShortGetModel) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TestPointShortGetModel) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TestPointShortGetModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *TestPointShortGetModel) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *TestPointShortGetModel) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetAttributes

`func (o *TestPointShortGetModel) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TestPointShortGetModel) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TestPointShortGetModel) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TestPointShortGetModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *TestPointShortGetModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *TestPointShortGetModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetTags

`func (o *TestPointShortGetModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestPointShortGetModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestPointShortGetModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TestPointShortGetModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *TestPointShortGetModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *TestPointShortGetModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetLinks

`func (o *TestPointShortGetModel) GetLinks() []string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TestPointShortGetModel) GetLinksOk() (*[]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TestPointShortGetModel) SetLinks(v []string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TestPointShortGetModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TestPointShortGetModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TestPointShortGetModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetTestSuiteId

`func (o *TestPointShortGetModel) GetTestSuiteId() string`

GetTestSuiteId returns the TestSuiteId field if non-nil, zero value otherwise.

### GetTestSuiteIdOk

`func (o *TestPointShortGetModel) GetTestSuiteIdOk() (*string, bool)`

GetTestSuiteIdOk returns a tuple with the TestSuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuiteId

`func (o *TestPointShortGetModel) SetTestSuiteId(v string)`

SetTestSuiteId sets TestSuiteId field to given value.


### GetWorkItemId

`func (o *TestPointShortGetModel) GetWorkItemId() string`

GetWorkItemId returns the WorkItemId field if non-nil, zero value otherwise.

### GetWorkItemIdOk

`func (o *TestPointShortGetModel) GetWorkItemIdOk() (*string, bool)`

GetWorkItemIdOk returns a tuple with the WorkItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemId

`func (o *TestPointShortGetModel) SetWorkItemId(v string)`

SetWorkItemId sets WorkItemId field to given value.


### GetWorkItemGlobalId

`func (o *TestPointShortGetModel) GetWorkItemGlobalId() int64`

GetWorkItemGlobalId returns the WorkItemGlobalId field if non-nil, zero value otherwise.

### GetWorkItemGlobalIdOk

`func (o *TestPointShortGetModel) GetWorkItemGlobalIdOk() (*int64, bool)`

GetWorkItemGlobalIdOk returns a tuple with the WorkItemGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemGlobalId

`func (o *TestPointShortGetModel) SetWorkItemGlobalId(v int64)`

SetWorkItemGlobalId sets WorkItemGlobalId field to given value.


### GetWorkItemVersionId

`func (o *TestPointShortGetModel) GetWorkItemVersionId() string`

GetWorkItemVersionId returns the WorkItemVersionId field if non-nil, zero value otherwise.

### GetWorkItemVersionIdOk

`func (o *TestPointShortGetModel) GetWorkItemVersionIdOk() (*string, bool)`

GetWorkItemVersionIdOk returns a tuple with the WorkItemVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionId

`func (o *TestPointShortGetModel) SetWorkItemVersionId(v string)`

SetWorkItemVersionId sets WorkItemVersionId field to given value.


### GetWorkItemMedianDuration

`func (o *TestPointShortGetModel) GetWorkItemMedianDuration() int64`

GetWorkItemMedianDuration returns the WorkItemMedianDuration field if non-nil, zero value otherwise.

### GetWorkItemMedianDurationOk

`func (o *TestPointShortGetModel) GetWorkItemMedianDurationOk() (*int64, bool)`

GetWorkItemMedianDurationOk returns a tuple with the WorkItemMedianDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemMedianDuration

`func (o *TestPointShortGetModel) SetWorkItemMedianDuration(v int64)`

SetWorkItemMedianDuration sets WorkItemMedianDuration field to given value.

### HasWorkItemMedianDuration

`func (o *TestPointShortGetModel) HasWorkItemMedianDuration() bool`

HasWorkItemMedianDuration returns a boolean if a field has been set.

### SetWorkItemMedianDurationNil

`func (o *TestPointShortGetModel) SetWorkItemMedianDurationNil(b bool)`

 SetWorkItemMedianDurationNil sets the value for WorkItemMedianDuration to be an explicit nil

### UnsetWorkItemMedianDuration
`func (o *TestPointShortGetModel) UnsetWorkItemMedianDuration()`

UnsetWorkItemMedianDuration ensures that no value is present for WorkItemMedianDuration, not even an explicit nil
### GetStatus

`func (o *TestPointShortGetModel) GetStatus() TestPointStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestPointShortGetModel) GetStatusOk() (*TestPointStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestPointShortGetModel) SetStatus(v TestPointStatus)`

SetStatus sets Status field to given value.


### GetPriority

`func (o *TestPointShortGetModel) GetPriority() WorkItemPriorityModel`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TestPointShortGetModel) GetPriorityOk() (*WorkItemPriorityModel, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TestPointShortGetModel) SetPriority(v WorkItemPriorityModel)`

SetPriority sets Priority field to given value.


### GetIsAutomated

`func (o *TestPointShortGetModel) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *TestPointShortGetModel) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *TestPointShortGetModel) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.


### GetName

`func (o *TestPointShortGetModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestPointShortGetModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestPointShortGetModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestPointShortGetModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestPointShortGetModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestPointShortGetModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetConfigurationId

`func (o *TestPointShortGetModel) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *TestPointShortGetModel) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *TestPointShortGetModel) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.


### GetDuration

`func (o *TestPointShortGetModel) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TestPointShortGetModel) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TestPointShortGetModel) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetSectionId

`func (o *TestPointShortGetModel) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *TestPointShortGetModel) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *TestPointShortGetModel) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.


### GetSectionName

`func (o *TestPointShortGetModel) GetSectionName() string`

GetSectionName returns the SectionName field if non-nil, zero value otherwise.

### GetSectionNameOk

`func (o *TestPointShortGetModel) GetSectionNameOk() (*string, bool)`

GetSectionNameOk returns a tuple with the SectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionName

`func (o *TestPointShortGetModel) SetSectionName(v string)`

SetSectionName sets SectionName field to given value.

### HasSectionName

`func (o *TestPointShortGetModel) HasSectionName() bool`

HasSectionName returns a boolean if a field has been set.

### SetSectionNameNil

`func (o *TestPointShortGetModel) SetSectionNameNil(b bool)`

 SetSectionNameNil sets the value for SectionName to be an explicit nil

### UnsetSectionName
`func (o *TestPointShortGetModel) UnsetSectionName()`

UnsetSectionName ensures that no value is present for SectionName, not even an explicit nil
### GetProjectId

`func (o *TestPointShortGetModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TestPointShortGetModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TestPointShortGetModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetLastTestResult

`func (o *TestPointShortGetModel) GetLastTestResult() TestPointShortGetModelLastTestResult`

GetLastTestResult returns the LastTestResult field if non-nil, zero value otherwise.

### GetLastTestResultOk

`func (o *TestPointShortGetModel) GetLastTestResultOk() (*TestPointShortGetModelLastTestResult, bool)`

GetLastTestResultOk returns a tuple with the LastTestResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResult

`func (o *TestPointShortGetModel) SetLastTestResult(v TestPointShortGetModelLastTestResult)`

SetLastTestResult sets LastTestResult field to given value.


### GetIterationId

`func (o *TestPointShortGetModel) GetIterationId() string`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *TestPointShortGetModel) GetIterationIdOk() (*string, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *TestPointShortGetModel) SetIterationId(v string)`

SetIterationId sets IterationId field to given value.


### GetWorkItemState

`func (o *TestPointShortGetModel) GetWorkItemState() WorkItemState`

GetWorkItemState returns the WorkItemState field if non-nil, zero value otherwise.

### GetWorkItemStateOk

`func (o *TestPointShortGetModel) GetWorkItemStateOk() (*WorkItemState, bool)`

GetWorkItemStateOk returns a tuple with the WorkItemState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemState

`func (o *TestPointShortGetModel) SetWorkItemState(v WorkItemState)`

SetWorkItemState sets WorkItemState field to given value.


### GetWorkItemCreatedById

`func (o *TestPointShortGetModel) GetWorkItemCreatedById() string`

GetWorkItemCreatedById returns the WorkItemCreatedById field if non-nil, zero value otherwise.

### GetWorkItemCreatedByIdOk

`func (o *TestPointShortGetModel) GetWorkItemCreatedByIdOk() (*string, bool)`

GetWorkItemCreatedByIdOk returns a tuple with the WorkItemCreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemCreatedById

`func (o *TestPointShortGetModel) SetWorkItemCreatedById(v string)`

SetWorkItemCreatedById sets WorkItemCreatedById field to given value.


### GetWorkItemCreatedDate

`func (o *TestPointShortGetModel) GetWorkItemCreatedDate() time.Time`

GetWorkItemCreatedDate returns the WorkItemCreatedDate field if non-nil, zero value otherwise.

### GetWorkItemCreatedDateOk

`func (o *TestPointShortGetModel) GetWorkItemCreatedDateOk() (*time.Time, bool)`

GetWorkItemCreatedDateOk returns a tuple with the WorkItemCreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemCreatedDate

`func (o *TestPointShortGetModel) SetWorkItemCreatedDate(v time.Time)`

SetWorkItemCreatedDate sets WorkItemCreatedDate field to given value.


### GetWorkItemModifiedById

`func (o *TestPointShortGetModel) GetWorkItemModifiedById() string`

GetWorkItemModifiedById returns the WorkItemModifiedById field if non-nil, zero value otherwise.

### GetWorkItemModifiedByIdOk

`func (o *TestPointShortGetModel) GetWorkItemModifiedByIdOk() (*string, bool)`

GetWorkItemModifiedByIdOk returns a tuple with the WorkItemModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemModifiedById

`func (o *TestPointShortGetModel) SetWorkItemModifiedById(v string)`

SetWorkItemModifiedById sets WorkItemModifiedById field to given value.

### HasWorkItemModifiedById

`func (o *TestPointShortGetModel) HasWorkItemModifiedById() bool`

HasWorkItemModifiedById returns a boolean if a field has been set.

### SetWorkItemModifiedByIdNil

`func (o *TestPointShortGetModel) SetWorkItemModifiedByIdNil(b bool)`

 SetWorkItemModifiedByIdNil sets the value for WorkItemModifiedById to be an explicit nil

### UnsetWorkItemModifiedById
`func (o *TestPointShortGetModel) UnsetWorkItemModifiedById()`

UnsetWorkItemModifiedById ensures that no value is present for WorkItemModifiedById, not even an explicit nil
### GetWorkItemModifiedDate

`func (o *TestPointShortGetModel) GetWorkItemModifiedDate() time.Time`

GetWorkItemModifiedDate returns the WorkItemModifiedDate field if non-nil, zero value otherwise.

### GetWorkItemModifiedDateOk

`func (o *TestPointShortGetModel) GetWorkItemModifiedDateOk() (*time.Time, bool)`

GetWorkItemModifiedDateOk returns a tuple with the WorkItemModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemModifiedDate

`func (o *TestPointShortGetModel) SetWorkItemModifiedDate(v time.Time)`

SetWorkItemModifiedDate sets WorkItemModifiedDate field to given value.

### HasWorkItemModifiedDate

`func (o *TestPointShortGetModel) HasWorkItemModifiedDate() bool`

HasWorkItemModifiedDate returns a boolean if a field has been set.

### SetWorkItemModifiedDateNil

`func (o *TestPointShortGetModel) SetWorkItemModifiedDateNil(b bool)`

 SetWorkItemModifiedDateNil sets the value for WorkItemModifiedDate to be an explicit nil

### UnsetWorkItemModifiedDate
`func (o *TestPointShortGetModel) UnsetWorkItemModifiedDate()`

UnsetWorkItemModifiedDate ensures that no value is present for WorkItemModifiedDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


