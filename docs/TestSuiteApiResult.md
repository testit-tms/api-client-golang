# TestSuiteApiResult

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

## Methods

### NewTestSuiteApiResult

`func NewTestSuiteApiResult(id string, name string, type_ TestSuiteTypeApiResult, saveStructure NullableBool, autoRefresh NullableBool, refreshDate NullableTime, parentId NullableString, testPlanId NullableString, ) *TestSuiteApiResult`

NewTestSuiteApiResult instantiates a new TestSuiteApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestSuiteApiResultWithDefaults

`func NewTestSuiteApiResultWithDefaults() *TestSuiteApiResult`

NewTestSuiteApiResultWithDefaults instantiates a new TestSuiteApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestSuiteApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestSuiteApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestSuiteApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TestSuiteApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestSuiteApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestSuiteApiResult) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *TestSuiteApiResult) GetType() TestSuiteTypeApiResult`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestSuiteApiResult) GetTypeOk() (*TestSuiteTypeApiResult, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestSuiteApiResult) SetType(v TestSuiteTypeApiResult)`

SetType sets Type field to given value.


### GetSaveStructure

`func (o *TestSuiteApiResult) GetSaveStructure() bool`

GetSaveStructure returns the SaveStructure field if non-nil, zero value otherwise.

### GetSaveStructureOk

`func (o *TestSuiteApiResult) GetSaveStructureOk() (*bool, bool)`

GetSaveStructureOk returns a tuple with the SaveStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveStructure

`func (o *TestSuiteApiResult) SetSaveStructure(v bool)`

SetSaveStructure sets SaveStructure field to given value.


### SetSaveStructureNil

`func (o *TestSuiteApiResult) SetSaveStructureNil(b bool)`

 SetSaveStructureNil sets the value for SaveStructure to be an explicit nil

### UnsetSaveStructure
`func (o *TestSuiteApiResult) UnsetSaveStructure()`

UnsetSaveStructure ensures that no value is present for SaveStructure, not even an explicit nil
### GetAutoRefresh

`func (o *TestSuiteApiResult) GetAutoRefresh() bool`

GetAutoRefresh returns the AutoRefresh field if non-nil, zero value otherwise.

### GetAutoRefreshOk

`func (o *TestSuiteApiResult) GetAutoRefreshOk() (*bool, bool)`

GetAutoRefreshOk returns a tuple with the AutoRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRefresh

`func (o *TestSuiteApiResult) SetAutoRefresh(v bool)`

SetAutoRefresh sets AutoRefresh field to given value.


### SetAutoRefreshNil

`func (o *TestSuiteApiResult) SetAutoRefreshNil(b bool)`

 SetAutoRefreshNil sets the value for AutoRefresh to be an explicit nil

### UnsetAutoRefresh
`func (o *TestSuiteApiResult) UnsetAutoRefresh()`

UnsetAutoRefresh ensures that no value is present for AutoRefresh, not even an explicit nil
### GetRefreshDate

`func (o *TestSuiteApiResult) GetRefreshDate() time.Time`

GetRefreshDate returns the RefreshDate field if non-nil, zero value otherwise.

### GetRefreshDateOk

`func (o *TestSuiteApiResult) GetRefreshDateOk() (*time.Time, bool)`

GetRefreshDateOk returns a tuple with the RefreshDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshDate

`func (o *TestSuiteApiResult) SetRefreshDate(v time.Time)`

SetRefreshDate sets RefreshDate field to given value.


### SetRefreshDateNil

`func (o *TestSuiteApiResult) SetRefreshDateNil(b bool)`

 SetRefreshDateNil sets the value for RefreshDate to be an explicit nil

### UnsetRefreshDate
`func (o *TestSuiteApiResult) UnsetRefreshDate()`

UnsetRefreshDate ensures that no value is present for RefreshDate, not even an explicit nil
### GetParentId

`func (o *TestSuiteApiResult) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *TestSuiteApiResult) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *TestSuiteApiResult) SetParentId(v string)`

SetParentId sets ParentId field to given value.


### SetParentIdNil

`func (o *TestSuiteApiResult) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *TestSuiteApiResult) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetTestPlanId

`func (o *TestSuiteApiResult) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *TestSuiteApiResult) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *TestSuiteApiResult) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.


### SetTestPlanIdNil

`func (o *TestSuiteApiResult) SetTestPlanIdNil(b bool)`

 SetTestPlanIdNil sets the value for TestPlanId to be an explicit nil

### UnsetTestPlanId
`func (o *TestSuiteApiResult) UnsetTestPlanId()`

UnsetTestPlanId ensures that no value is present for TestPlanId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


