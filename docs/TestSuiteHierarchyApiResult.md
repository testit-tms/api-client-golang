# TestSuiteHierarchyApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the test suite | 
**Name** | **string** | Name of the test suite | 
**Type** | [**TestSuiteTypeApiResult**](TestSuiteTypeApiResult.md) | Type of the test suite | 
**SaveStructure** | **NullableBool** | Flag indicating whether the structure of the test suite should be saved | 
**AutoRefresh** | **NullableBool** | Flag indicating whether auto-refresh functionality is enabled for the test suite | 
**RefreshDate** | **NullableTime** | The last time the test suite&#39;s results were refreshed | 
**ParentId** | **NullableString** | Unique ID of the parent test suite, if any | 
**TestPlanId** | **NullableString** | Unique ID of the associated test plan | 
**Children** | Pointer to [**[]TestSuiteHierarchyApiResult**](TestSuiteHierarchyApiResult.md) | Collection of child test suites | [optional] 

## Methods

### NewTestSuiteHierarchyApiResult

`func NewTestSuiteHierarchyApiResult(id string, name string, type_ TestSuiteTypeApiResult, saveStructure NullableBool, autoRefresh NullableBool, refreshDate NullableTime, parentId NullableString, testPlanId NullableString, ) *TestSuiteHierarchyApiResult`

NewTestSuiteHierarchyApiResult instantiates a new TestSuiteHierarchyApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestSuiteHierarchyApiResultWithDefaults

`func NewTestSuiteHierarchyApiResultWithDefaults() *TestSuiteHierarchyApiResult`

NewTestSuiteHierarchyApiResultWithDefaults instantiates a new TestSuiteHierarchyApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestSuiteHierarchyApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestSuiteHierarchyApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestSuiteHierarchyApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TestSuiteHierarchyApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestSuiteHierarchyApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestSuiteHierarchyApiResult) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *TestSuiteHierarchyApiResult) GetType() TestSuiteTypeApiResult`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestSuiteHierarchyApiResult) GetTypeOk() (*TestSuiteTypeApiResult, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestSuiteHierarchyApiResult) SetType(v TestSuiteTypeApiResult)`

SetType sets Type field to given value.


### GetSaveStructure

`func (o *TestSuiteHierarchyApiResult) GetSaveStructure() bool`

GetSaveStructure returns the SaveStructure field if non-nil, zero value otherwise.

### GetSaveStructureOk

`func (o *TestSuiteHierarchyApiResult) GetSaveStructureOk() (*bool, bool)`

GetSaveStructureOk returns a tuple with the SaveStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveStructure

`func (o *TestSuiteHierarchyApiResult) SetSaveStructure(v bool)`

SetSaveStructure sets SaveStructure field to given value.


### SetSaveStructureNil

`func (o *TestSuiteHierarchyApiResult) SetSaveStructureNil(b bool)`

 SetSaveStructureNil sets the value for SaveStructure to be an explicit nil

### UnsetSaveStructure
`func (o *TestSuiteHierarchyApiResult) UnsetSaveStructure()`

UnsetSaveStructure ensures that no value is present for SaveStructure, not even an explicit nil
### GetAutoRefresh

`func (o *TestSuiteHierarchyApiResult) GetAutoRefresh() bool`

GetAutoRefresh returns the AutoRefresh field if non-nil, zero value otherwise.

### GetAutoRefreshOk

`func (o *TestSuiteHierarchyApiResult) GetAutoRefreshOk() (*bool, bool)`

GetAutoRefreshOk returns a tuple with the AutoRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRefresh

`func (o *TestSuiteHierarchyApiResult) SetAutoRefresh(v bool)`

SetAutoRefresh sets AutoRefresh field to given value.


### SetAutoRefreshNil

`func (o *TestSuiteHierarchyApiResult) SetAutoRefreshNil(b bool)`

 SetAutoRefreshNil sets the value for AutoRefresh to be an explicit nil

### UnsetAutoRefresh
`func (o *TestSuiteHierarchyApiResult) UnsetAutoRefresh()`

UnsetAutoRefresh ensures that no value is present for AutoRefresh, not even an explicit nil
### GetRefreshDate

`func (o *TestSuiteHierarchyApiResult) GetRefreshDate() time.Time`

GetRefreshDate returns the RefreshDate field if non-nil, zero value otherwise.

### GetRefreshDateOk

`func (o *TestSuiteHierarchyApiResult) GetRefreshDateOk() (*time.Time, bool)`

GetRefreshDateOk returns a tuple with the RefreshDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshDate

`func (o *TestSuiteHierarchyApiResult) SetRefreshDate(v time.Time)`

SetRefreshDate sets RefreshDate field to given value.


### SetRefreshDateNil

`func (o *TestSuiteHierarchyApiResult) SetRefreshDateNil(b bool)`

 SetRefreshDateNil sets the value for RefreshDate to be an explicit nil

### UnsetRefreshDate
`func (o *TestSuiteHierarchyApiResult) UnsetRefreshDate()`

UnsetRefreshDate ensures that no value is present for RefreshDate, not even an explicit nil
### GetParentId

`func (o *TestSuiteHierarchyApiResult) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *TestSuiteHierarchyApiResult) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *TestSuiteHierarchyApiResult) SetParentId(v string)`

SetParentId sets ParentId field to given value.


### SetParentIdNil

`func (o *TestSuiteHierarchyApiResult) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *TestSuiteHierarchyApiResult) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetTestPlanId

`func (o *TestSuiteHierarchyApiResult) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *TestSuiteHierarchyApiResult) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *TestSuiteHierarchyApiResult) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.


### SetTestPlanIdNil

`func (o *TestSuiteHierarchyApiResult) SetTestPlanIdNil(b bool)`

 SetTestPlanIdNil sets the value for TestPlanId to be an explicit nil

### UnsetTestPlanId
`func (o *TestSuiteHierarchyApiResult) UnsetTestPlanId()`

UnsetTestPlanId ensures that no value is present for TestPlanId, not even an explicit nil
### GetChildren

`func (o *TestSuiteHierarchyApiResult) GetChildren() []TestSuiteHierarchyApiResult`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *TestSuiteHierarchyApiResult) GetChildrenOk() (*[]TestSuiteHierarchyApiResult, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *TestSuiteHierarchyApiResult) SetChildren(v []TestSuiteHierarchyApiResult)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *TestSuiteHierarchyApiResult) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### SetChildrenNil

`func (o *TestSuiteHierarchyApiResult) SetChildrenNil(b bool)`

 SetChildrenNil sets the value for Children to be an explicit nil

### UnsetChildren
`func (o *TestSuiteHierarchyApiResult) UnsetChildren()`

UnsetChildren ensures that no value is present for Children, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


