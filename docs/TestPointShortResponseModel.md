# TestPointShortResponseModel

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
**Attributes** | **map[string]interface{}** | Collection of attributes of work item the test point represents | 
**Tags** | **[]string** | Collection of the test point tags | 
**Links** | **[]string** | Collection of the test point links | 
**TestSuiteId** | **string** | Unique ID of test suite the test point assigned to | 
**TestSuiteName** | **string** | Name of the test suite | 
**WorkItemId** | **string** | Unique ID of work item the test point represents | 
**WorkItemGlobalId** | **int64** | Global ID of work item the test point represents | 
**WorkItemVersionId** | **string** | Unique ID of work item version the test point represents | 
**WorkItemVersionNumber** | **int32** | Number of work item version the test point represents | 
**WorkItemMedianDuration** | Pointer to **NullableInt64** | Median duration of work item the test point represents | [optional] 
**Status** | [**TestPointStatus**](TestPointStatus.md) | Status of the test point | 
**StatusModel** | [**TestStatusApiResult**](TestStatusApiResult.md) | Status of the test point | 
**Priority** | [**WorkItemPriorityModel**](WorkItemPriorityModel.md) | Priority of the test point | 
**SourceType** | [**WorkItemSourceTypeModel**](WorkItemSourceTypeModel.md) | Source type of the test point | 
**IsAutomated** | **bool** | Indicates if the test point represents an autotest | 
**Name** | **string** | Name of the test point | 
**ConfigurationId** | **string** | Unique ID of the test point configuration | 
**Duration** | **int32** | Duration of the test point | 
**SectionId** | **string** | Unique ID of section where work item the test point represents is located | 
**SectionName** | Pointer to **NullableString** | Name of section where work item the test point represents is located | [optional] 
**ProjectId** | **string** | Unique ID of the test point project | 
**LastTestResult** | Pointer to [**NullableLastTestResultModel**](LastTestResultModel.md) | Model of the test point last test result | [optional] 
**IterationId** | **string** | Unique ID of work item iteration the test point represents | 
**WorkItemState** | [**WorkItemState**](WorkItemState.md) | Work item state | 
**WorkItemCreatedById** | **string** | Unique ID of the work item creator | 
**WorkItemCreatedDate** | **time.Time** | Creation date of work item | 
**WorkItemModifiedById** | Pointer to **NullableString** | Unique ID of the work item last editor | [optional] 
**WorkItemModifiedDate** | Pointer to **NullableTime** | Modified date of work item | [optional] 

## Methods

### NewTestPointShortResponseModel

`func NewTestPointShortResponseModel(id string, createdDate time.Time, createdById string, attributes map[string]interface{}, tags []string, links []string, testSuiteId string, testSuiteName string, workItemId string, workItemGlobalId int64, workItemVersionId string, workItemVersionNumber int32, status TestPointStatus, statusModel TestStatusApiResult, priority WorkItemPriorityModel, sourceType WorkItemSourceTypeModel, isAutomated bool, name string, configurationId string, duration int32, sectionId string, projectId string, iterationId string, workItemState WorkItemState, workItemCreatedById string, workItemCreatedDate time.Time, ) *TestPointShortResponseModel`

NewTestPointShortResponseModel instantiates a new TestPointShortResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPointShortResponseModelWithDefaults

`func NewTestPointShortResponseModelWithDefaults() *TestPointShortResponseModel`

NewTestPointShortResponseModelWithDefaults instantiates a new TestPointShortResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestPointShortResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestPointShortResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestPointShortResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedDate

`func (o *TestPointShortResponseModel) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TestPointShortResponseModel) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TestPointShortResponseModel) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetCreatedById

`func (o *TestPointShortResponseModel) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TestPointShortResponseModel) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TestPointShortResponseModel) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedDate

`func (o *TestPointShortResponseModel) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *TestPointShortResponseModel) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *TestPointShortResponseModel) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *TestPointShortResponseModel) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *TestPointShortResponseModel) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *TestPointShortResponseModel) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetModifiedById

`func (o *TestPointShortResponseModel) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *TestPointShortResponseModel) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *TestPointShortResponseModel) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *TestPointShortResponseModel) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *TestPointShortResponseModel) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *TestPointShortResponseModel) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil
### GetTesterId

`func (o *TestPointShortResponseModel) GetTesterId() string`

GetTesterId returns the TesterId field if non-nil, zero value otherwise.

### GetTesterIdOk

`func (o *TestPointShortResponseModel) GetTesterIdOk() (*string, bool)`

GetTesterIdOk returns a tuple with the TesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesterId

`func (o *TestPointShortResponseModel) SetTesterId(v string)`

SetTesterId sets TesterId field to given value.

### HasTesterId

`func (o *TestPointShortResponseModel) HasTesterId() bool`

HasTesterId returns a boolean if a field has been set.

### SetTesterIdNil

`func (o *TestPointShortResponseModel) SetTesterIdNil(b bool)`

 SetTesterIdNil sets the value for TesterId to be an explicit nil

### UnsetTesterId
`func (o *TestPointShortResponseModel) UnsetTesterId()`

UnsetTesterId ensures that no value is present for TesterId, not even an explicit nil
### GetParameters

`func (o *TestPointShortResponseModel) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TestPointShortResponseModel) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TestPointShortResponseModel) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *TestPointShortResponseModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *TestPointShortResponseModel) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *TestPointShortResponseModel) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetAttributes

`func (o *TestPointShortResponseModel) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TestPointShortResponseModel) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TestPointShortResponseModel) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetTags

`func (o *TestPointShortResponseModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestPointShortResponseModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestPointShortResponseModel) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetLinks

`func (o *TestPointShortResponseModel) GetLinks() []string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TestPointShortResponseModel) GetLinksOk() (*[]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TestPointShortResponseModel) SetLinks(v []string)`

SetLinks sets Links field to given value.


### GetTestSuiteId

`func (o *TestPointShortResponseModel) GetTestSuiteId() string`

GetTestSuiteId returns the TestSuiteId field if non-nil, zero value otherwise.

### GetTestSuiteIdOk

`func (o *TestPointShortResponseModel) GetTestSuiteIdOk() (*string, bool)`

GetTestSuiteIdOk returns a tuple with the TestSuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuiteId

`func (o *TestPointShortResponseModel) SetTestSuiteId(v string)`

SetTestSuiteId sets TestSuiteId field to given value.


### GetTestSuiteName

`func (o *TestPointShortResponseModel) GetTestSuiteName() string`

GetTestSuiteName returns the TestSuiteName field if non-nil, zero value otherwise.

### GetTestSuiteNameOk

`func (o *TestPointShortResponseModel) GetTestSuiteNameOk() (*string, bool)`

GetTestSuiteNameOk returns a tuple with the TestSuiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSuiteName

`func (o *TestPointShortResponseModel) SetTestSuiteName(v string)`

SetTestSuiteName sets TestSuiteName field to given value.


### GetWorkItemId

`func (o *TestPointShortResponseModel) GetWorkItemId() string`

GetWorkItemId returns the WorkItemId field if non-nil, zero value otherwise.

### GetWorkItemIdOk

`func (o *TestPointShortResponseModel) GetWorkItemIdOk() (*string, bool)`

GetWorkItemIdOk returns a tuple with the WorkItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemId

`func (o *TestPointShortResponseModel) SetWorkItemId(v string)`

SetWorkItemId sets WorkItemId field to given value.


### GetWorkItemGlobalId

`func (o *TestPointShortResponseModel) GetWorkItemGlobalId() int64`

GetWorkItemGlobalId returns the WorkItemGlobalId field if non-nil, zero value otherwise.

### GetWorkItemGlobalIdOk

`func (o *TestPointShortResponseModel) GetWorkItemGlobalIdOk() (*int64, bool)`

GetWorkItemGlobalIdOk returns a tuple with the WorkItemGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemGlobalId

`func (o *TestPointShortResponseModel) SetWorkItemGlobalId(v int64)`

SetWorkItemGlobalId sets WorkItemGlobalId field to given value.


### GetWorkItemVersionId

`func (o *TestPointShortResponseModel) GetWorkItemVersionId() string`

GetWorkItemVersionId returns the WorkItemVersionId field if non-nil, zero value otherwise.

### GetWorkItemVersionIdOk

`func (o *TestPointShortResponseModel) GetWorkItemVersionIdOk() (*string, bool)`

GetWorkItemVersionIdOk returns a tuple with the WorkItemVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionId

`func (o *TestPointShortResponseModel) SetWorkItemVersionId(v string)`

SetWorkItemVersionId sets WorkItemVersionId field to given value.


### GetWorkItemVersionNumber

`func (o *TestPointShortResponseModel) GetWorkItemVersionNumber() int32`

GetWorkItemVersionNumber returns the WorkItemVersionNumber field if non-nil, zero value otherwise.

### GetWorkItemVersionNumberOk

`func (o *TestPointShortResponseModel) GetWorkItemVersionNumberOk() (*int32, bool)`

GetWorkItemVersionNumberOk returns a tuple with the WorkItemVersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemVersionNumber

`func (o *TestPointShortResponseModel) SetWorkItemVersionNumber(v int32)`

SetWorkItemVersionNumber sets WorkItemVersionNumber field to given value.


### GetWorkItemMedianDuration

`func (o *TestPointShortResponseModel) GetWorkItemMedianDuration() int64`

GetWorkItemMedianDuration returns the WorkItemMedianDuration field if non-nil, zero value otherwise.

### GetWorkItemMedianDurationOk

`func (o *TestPointShortResponseModel) GetWorkItemMedianDurationOk() (*int64, bool)`

GetWorkItemMedianDurationOk returns a tuple with the WorkItemMedianDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemMedianDuration

`func (o *TestPointShortResponseModel) SetWorkItemMedianDuration(v int64)`

SetWorkItemMedianDuration sets WorkItemMedianDuration field to given value.

### HasWorkItemMedianDuration

`func (o *TestPointShortResponseModel) HasWorkItemMedianDuration() bool`

HasWorkItemMedianDuration returns a boolean if a field has been set.

### SetWorkItemMedianDurationNil

`func (o *TestPointShortResponseModel) SetWorkItemMedianDurationNil(b bool)`

 SetWorkItemMedianDurationNil sets the value for WorkItemMedianDuration to be an explicit nil

### UnsetWorkItemMedianDuration
`func (o *TestPointShortResponseModel) UnsetWorkItemMedianDuration()`

UnsetWorkItemMedianDuration ensures that no value is present for WorkItemMedianDuration, not even an explicit nil
### GetStatus

`func (o *TestPointShortResponseModel) GetStatus() TestPointStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestPointShortResponseModel) GetStatusOk() (*TestPointStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestPointShortResponseModel) SetStatus(v TestPointStatus)`

SetStatus sets Status field to given value.


### GetStatusModel

`func (o *TestPointShortResponseModel) GetStatusModel() TestStatusApiResult`

GetStatusModel returns the StatusModel field if non-nil, zero value otherwise.

### GetStatusModelOk

`func (o *TestPointShortResponseModel) GetStatusModelOk() (*TestStatusApiResult, bool)`

GetStatusModelOk returns a tuple with the StatusModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusModel

`func (o *TestPointShortResponseModel) SetStatusModel(v TestStatusApiResult)`

SetStatusModel sets StatusModel field to given value.


### GetPriority

`func (o *TestPointShortResponseModel) GetPriority() WorkItemPriorityModel`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TestPointShortResponseModel) GetPriorityOk() (*WorkItemPriorityModel, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TestPointShortResponseModel) SetPriority(v WorkItemPriorityModel)`

SetPriority sets Priority field to given value.


### GetSourceType

`func (o *TestPointShortResponseModel) GetSourceType() WorkItemSourceTypeModel`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TestPointShortResponseModel) GetSourceTypeOk() (*WorkItemSourceTypeModel, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TestPointShortResponseModel) SetSourceType(v WorkItemSourceTypeModel)`

SetSourceType sets SourceType field to given value.


### GetIsAutomated

`func (o *TestPointShortResponseModel) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *TestPointShortResponseModel) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *TestPointShortResponseModel) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.


### GetName

`func (o *TestPointShortResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestPointShortResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestPointShortResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetConfigurationId

`func (o *TestPointShortResponseModel) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *TestPointShortResponseModel) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *TestPointShortResponseModel) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.


### GetDuration

`func (o *TestPointShortResponseModel) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TestPointShortResponseModel) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TestPointShortResponseModel) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetSectionId

`func (o *TestPointShortResponseModel) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *TestPointShortResponseModel) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *TestPointShortResponseModel) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.


### GetSectionName

`func (o *TestPointShortResponseModel) GetSectionName() string`

GetSectionName returns the SectionName field if non-nil, zero value otherwise.

### GetSectionNameOk

`func (o *TestPointShortResponseModel) GetSectionNameOk() (*string, bool)`

GetSectionNameOk returns a tuple with the SectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionName

`func (o *TestPointShortResponseModel) SetSectionName(v string)`

SetSectionName sets SectionName field to given value.

### HasSectionName

`func (o *TestPointShortResponseModel) HasSectionName() bool`

HasSectionName returns a boolean if a field has been set.

### SetSectionNameNil

`func (o *TestPointShortResponseModel) SetSectionNameNil(b bool)`

 SetSectionNameNil sets the value for SectionName to be an explicit nil

### UnsetSectionName
`func (o *TestPointShortResponseModel) UnsetSectionName()`

UnsetSectionName ensures that no value is present for SectionName, not even an explicit nil
### GetProjectId

`func (o *TestPointShortResponseModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TestPointShortResponseModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TestPointShortResponseModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetLastTestResult

`func (o *TestPointShortResponseModel) GetLastTestResult() LastTestResultModel`

GetLastTestResult returns the LastTestResult field if non-nil, zero value otherwise.

### GetLastTestResultOk

`func (o *TestPointShortResponseModel) GetLastTestResultOk() (*LastTestResultModel, bool)`

GetLastTestResultOk returns a tuple with the LastTestResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestResult

`func (o *TestPointShortResponseModel) SetLastTestResult(v LastTestResultModel)`

SetLastTestResult sets LastTestResult field to given value.

### HasLastTestResult

`func (o *TestPointShortResponseModel) HasLastTestResult() bool`

HasLastTestResult returns a boolean if a field has been set.

### SetLastTestResultNil

`func (o *TestPointShortResponseModel) SetLastTestResultNil(b bool)`

 SetLastTestResultNil sets the value for LastTestResult to be an explicit nil

### UnsetLastTestResult
`func (o *TestPointShortResponseModel) UnsetLastTestResult()`

UnsetLastTestResult ensures that no value is present for LastTestResult, not even an explicit nil
### GetIterationId

`func (o *TestPointShortResponseModel) GetIterationId() string`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *TestPointShortResponseModel) GetIterationIdOk() (*string, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *TestPointShortResponseModel) SetIterationId(v string)`

SetIterationId sets IterationId field to given value.


### GetWorkItemState

`func (o *TestPointShortResponseModel) GetWorkItemState() WorkItemState`

GetWorkItemState returns the WorkItemState field if non-nil, zero value otherwise.

### GetWorkItemStateOk

`func (o *TestPointShortResponseModel) GetWorkItemStateOk() (*WorkItemState, bool)`

GetWorkItemStateOk returns a tuple with the WorkItemState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemState

`func (o *TestPointShortResponseModel) SetWorkItemState(v WorkItemState)`

SetWorkItemState sets WorkItemState field to given value.


### GetWorkItemCreatedById

`func (o *TestPointShortResponseModel) GetWorkItemCreatedById() string`

GetWorkItemCreatedById returns the WorkItemCreatedById field if non-nil, zero value otherwise.

### GetWorkItemCreatedByIdOk

`func (o *TestPointShortResponseModel) GetWorkItemCreatedByIdOk() (*string, bool)`

GetWorkItemCreatedByIdOk returns a tuple with the WorkItemCreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemCreatedById

`func (o *TestPointShortResponseModel) SetWorkItemCreatedById(v string)`

SetWorkItemCreatedById sets WorkItemCreatedById field to given value.


### GetWorkItemCreatedDate

`func (o *TestPointShortResponseModel) GetWorkItemCreatedDate() time.Time`

GetWorkItemCreatedDate returns the WorkItemCreatedDate field if non-nil, zero value otherwise.

### GetWorkItemCreatedDateOk

`func (o *TestPointShortResponseModel) GetWorkItemCreatedDateOk() (*time.Time, bool)`

GetWorkItemCreatedDateOk returns a tuple with the WorkItemCreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemCreatedDate

`func (o *TestPointShortResponseModel) SetWorkItemCreatedDate(v time.Time)`

SetWorkItemCreatedDate sets WorkItemCreatedDate field to given value.


### GetWorkItemModifiedById

`func (o *TestPointShortResponseModel) GetWorkItemModifiedById() string`

GetWorkItemModifiedById returns the WorkItemModifiedById field if non-nil, zero value otherwise.

### GetWorkItemModifiedByIdOk

`func (o *TestPointShortResponseModel) GetWorkItemModifiedByIdOk() (*string, bool)`

GetWorkItemModifiedByIdOk returns a tuple with the WorkItemModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemModifiedById

`func (o *TestPointShortResponseModel) SetWorkItemModifiedById(v string)`

SetWorkItemModifiedById sets WorkItemModifiedById field to given value.

### HasWorkItemModifiedById

`func (o *TestPointShortResponseModel) HasWorkItemModifiedById() bool`

HasWorkItemModifiedById returns a boolean if a field has been set.

### SetWorkItemModifiedByIdNil

`func (o *TestPointShortResponseModel) SetWorkItemModifiedByIdNil(b bool)`

 SetWorkItemModifiedByIdNil sets the value for WorkItemModifiedById to be an explicit nil

### UnsetWorkItemModifiedById
`func (o *TestPointShortResponseModel) UnsetWorkItemModifiedById()`

UnsetWorkItemModifiedById ensures that no value is present for WorkItemModifiedById, not even an explicit nil
### GetWorkItemModifiedDate

`func (o *TestPointShortResponseModel) GetWorkItemModifiedDate() time.Time`

GetWorkItemModifiedDate returns the WorkItemModifiedDate field if non-nil, zero value otherwise.

### GetWorkItemModifiedDateOk

`func (o *TestPointShortResponseModel) GetWorkItemModifiedDateOk() (*time.Time, bool)`

GetWorkItemModifiedDateOk returns a tuple with the WorkItemModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemModifiedDate

`func (o *TestPointShortResponseModel) SetWorkItemModifiedDate(v time.Time)`

SetWorkItemModifiedDate sets WorkItemModifiedDate field to given value.

### HasWorkItemModifiedDate

`func (o *TestPointShortResponseModel) HasWorkItemModifiedDate() bool`

HasWorkItemModifiedDate returns a boolean if a field has been set.

### SetWorkItemModifiedDateNil

`func (o *TestPointShortResponseModel) SetWorkItemModifiedDateNil(b bool)`

 SetWorkItemModifiedDateNil sets the value for WorkItemModifiedDate to be an explicit nil

### UnsetWorkItemModifiedDate
`func (o *TestPointShortResponseModel) UnsetWorkItemModifiedDate()`

UnsetWorkItemModifiedDate ensures that no value is present for WorkItemModifiedDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


