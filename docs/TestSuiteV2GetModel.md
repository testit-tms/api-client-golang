# TestSuiteV2GetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the test suite | 
**RefreshDate** | Pointer to **NullableTime** | Date of the last refresh of the test suite | [optional] 
**ParentId** | Pointer to **NullableString** | Unique ID of the parent test suite in hierarchy | [optional] 
**TestPlanId** | **string** | Unique ID of test plan to which the test suite belongs | 
**Name** | **string** | Name of the test suite | 
**Type** | Pointer to [**NullableTestSuiteType**](TestSuiteType.md) | Type of the test suite | [optional] 
**SaveStructure** | Pointer to **NullableBool** | Indicates if the test suite retains section tree structure | [optional] 
**AutoRefresh** | Pointer to **NullableBool** | Indicates if scheduled auto refresh is enabled for the test suite | [optional] 

## Methods

### NewTestSuiteV2GetModel

`func NewTestSuiteV2GetModel(id string, testPlanId string, name string, ) *TestSuiteV2GetModel`

NewTestSuiteV2GetModel instantiates a new TestSuiteV2GetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestSuiteV2GetModelWithDefaults

`func NewTestSuiteV2GetModelWithDefaults() *TestSuiteV2GetModel`

NewTestSuiteV2GetModelWithDefaults instantiates a new TestSuiteV2GetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestSuiteV2GetModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestSuiteV2GetModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestSuiteV2GetModel) SetId(v string)`

SetId sets Id field to given value.


### GetRefreshDate

`func (o *TestSuiteV2GetModel) GetRefreshDate() time.Time`

GetRefreshDate returns the RefreshDate field if non-nil, zero value otherwise.

### GetRefreshDateOk

`func (o *TestSuiteV2GetModel) GetRefreshDateOk() (*time.Time, bool)`

GetRefreshDateOk returns a tuple with the RefreshDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshDate

`func (o *TestSuiteV2GetModel) SetRefreshDate(v time.Time)`

SetRefreshDate sets RefreshDate field to given value.

### HasRefreshDate

`func (o *TestSuiteV2GetModel) HasRefreshDate() bool`

HasRefreshDate returns a boolean if a field has been set.

### SetRefreshDateNil

`func (o *TestSuiteV2GetModel) SetRefreshDateNil(b bool)`

 SetRefreshDateNil sets the value for RefreshDate to be an explicit nil

### UnsetRefreshDate
`func (o *TestSuiteV2GetModel) UnsetRefreshDate()`

UnsetRefreshDate ensures that no value is present for RefreshDate, not even an explicit nil
### GetParentId

`func (o *TestSuiteV2GetModel) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *TestSuiteV2GetModel) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *TestSuiteV2GetModel) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *TestSuiteV2GetModel) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *TestSuiteV2GetModel) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *TestSuiteV2GetModel) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetTestPlanId

`func (o *TestSuiteV2GetModel) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *TestSuiteV2GetModel) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *TestSuiteV2GetModel) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.


### GetName

`func (o *TestSuiteV2GetModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestSuiteV2GetModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestSuiteV2GetModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *TestSuiteV2GetModel) GetType() TestSuiteType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestSuiteV2GetModel) GetTypeOk() (*TestSuiteType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestSuiteV2GetModel) SetType(v TestSuiteType)`

SetType sets Type field to given value.

### HasType

`func (o *TestSuiteV2GetModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *TestSuiteV2GetModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TestSuiteV2GetModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetSaveStructure

`func (o *TestSuiteV2GetModel) GetSaveStructure() bool`

GetSaveStructure returns the SaveStructure field if non-nil, zero value otherwise.

### GetSaveStructureOk

`func (o *TestSuiteV2GetModel) GetSaveStructureOk() (*bool, bool)`

GetSaveStructureOk returns a tuple with the SaveStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveStructure

`func (o *TestSuiteV2GetModel) SetSaveStructure(v bool)`

SetSaveStructure sets SaveStructure field to given value.

### HasSaveStructure

`func (o *TestSuiteV2GetModel) HasSaveStructure() bool`

HasSaveStructure returns a boolean if a field has been set.

### SetSaveStructureNil

`func (o *TestSuiteV2GetModel) SetSaveStructureNil(b bool)`

 SetSaveStructureNil sets the value for SaveStructure to be an explicit nil

### UnsetSaveStructure
`func (o *TestSuiteV2GetModel) UnsetSaveStructure()`

UnsetSaveStructure ensures that no value is present for SaveStructure, not even an explicit nil
### GetAutoRefresh

`func (o *TestSuiteV2GetModel) GetAutoRefresh() bool`

GetAutoRefresh returns the AutoRefresh field if non-nil, zero value otherwise.

### GetAutoRefreshOk

`func (o *TestSuiteV2GetModel) GetAutoRefreshOk() (*bool, bool)`

GetAutoRefreshOk returns a tuple with the AutoRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRefresh

`func (o *TestSuiteV2GetModel) SetAutoRefresh(v bool)`

SetAutoRefresh sets AutoRefresh field to given value.

### HasAutoRefresh

`func (o *TestSuiteV2GetModel) HasAutoRefresh() bool`

HasAutoRefresh returns a boolean if a field has been set.

### SetAutoRefreshNil

`func (o *TestSuiteV2GetModel) SetAutoRefreshNil(b bool)`

 SetAutoRefreshNil sets the value for AutoRefresh to be an explicit nil

### UnsetAutoRefresh
`func (o *TestSuiteV2GetModel) UnsetAutoRefresh()`

UnsetAutoRefresh ensures that no value is present for AutoRefresh, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


