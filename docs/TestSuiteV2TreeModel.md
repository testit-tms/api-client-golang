# TestSuiteV2TreeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]TestSuiteV2TreeModel**](TestSuiteV2TreeModel.md) | nested enumeration of children is allowed | [optional] 
**Id** | **string** | Unique ID of the test suite | 
**RefreshDate** | Pointer to **NullableTime** | Date of the last refresh of the test suite | [optional] 
**ParentId** | Pointer to **NullableString** | Unique ID of the parent test suite in hierarchy | [optional] 
**TestPlanId** | **string** | Unique ID of test plan to which the test suite belongs | 
**Name** | **string** | Name of the test suite | 
**Type** | Pointer to [**NullableTestSuiteType**](TestSuiteType.md) |  | [optional] 
**SaveStructure** | Pointer to **NullableBool** | Indicates if the test suite retains section tree structure | [optional] 
**AutoRefresh** | Pointer to **NullableBool** | Indicates if scheduled auto refresh is enabled for the test suite | [optional] 

## Methods

### NewTestSuiteV2TreeModel

`func NewTestSuiteV2TreeModel(id string, testPlanId string, name string, ) *TestSuiteV2TreeModel`

NewTestSuiteV2TreeModel instantiates a new TestSuiteV2TreeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestSuiteV2TreeModelWithDefaults

`func NewTestSuiteV2TreeModelWithDefaults() *TestSuiteV2TreeModel`

NewTestSuiteV2TreeModelWithDefaults instantiates a new TestSuiteV2TreeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *TestSuiteV2TreeModel) GetChildren() []TestSuiteV2TreeModel`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *TestSuiteV2TreeModel) GetChildrenOk() (*[]TestSuiteV2TreeModel, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *TestSuiteV2TreeModel) SetChildren(v []TestSuiteV2TreeModel)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *TestSuiteV2TreeModel) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### SetChildrenNil

`func (o *TestSuiteV2TreeModel) SetChildrenNil(b bool)`

 SetChildrenNil sets the value for Children to be an explicit nil

### UnsetChildren
`func (o *TestSuiteV2TreeModel) UnsetChildren()`

UnsetChildren ensures that no value is present for Children, not even an explicit nil
### GetId

`func (o *TestSuiteV2TreeModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestSuiteV2TreeModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestSuiteV2TreeModel) SetId(v string)`

SetId sets Id field to given value.


### GetRefreshDate

`func (o *TestSuiteV2TreeModel) GetRefreshDate() time.Time`

GetRefreshDate returns the RefreshDate field if non-nil, zero value otherwise.

### GetRefreshDateOk

`func (o *TestSuiteV2TreeModel) GetRefreshDateOk() (*time.Time, bool)`

GetRefreshDateOk returns a tuple with the RefreshDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshDate

`func (o *TestSuiteV2TreeModel) SetRefreshDate(v time.Time)`

SetRefreshDate sets RefreshDate field to given value.

### HasRefreshDate

`func (o *TestSuiteV2TreeModel) HasRefreshDate() bool`

HasRefreshDate returns a boolean if a field has been set.

### SetRefreshDateNil

`func (o *TestSuiteV2TreeModel) SetRefreshDateNil(b bool)`

 SetRefreshDateNil sets the value for RefreshDate to be an explicit nil

### UnsetRefreshDate
`func (o *TestSuiteV2TreeModel) UnsetRefreshDate()`

UnsetRefreshDate ensures that no value is present for RefreshDate, not even an explicit nil
### GetParentId

`func (o *TestSuiteV2TreeModel) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *TestSuiteV2TreeModel) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *TestSuiteV2TreeModel) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *TestSuiteV2TreeModel) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *TestSuiteV2TreeModel) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *TestSuiteV2TreeModel) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetTestPlanId

`func (o *TestSuiteV2TreeModel) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *TestSuiteV2TreeModel) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *TestSuiteV2TreeModel) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.


### GetName

`func (o *TestSuiteV2TreeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestSuiteV2TreeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestSuiteV2TreeModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *TestSuiteV2TreeModel) GetType() TestSuiteType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestSuiteV2TreeModel) GetTypeOk() (*TestSuiteType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestSuiteV2TreeModel) SetType(v TestSuiteType)`

SetType sets Type field to given value.

### HasType

`func (o *TestSuiteV2TreeModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *TestSuiteV2TreeModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TestSuiteV2TreeModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetSaveStructure

`func (o *TestSuiteV2TreeModel) GetSaveStructure() bool`

GetSaveStructure returns the SaveStructure field if non-nil, zero value otherwise.

### GetSaveStructureOk

`func (o *TestSuiteV2TreeModel) GetSaveStructureOk() (*bool, bool)`

GetSaveStructureOk returns a tuple with the SaveStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveStructure

`func (o *TestSuiteV2TreeModel) SetSaveStructure(v bool)`

SetSaveStructure sets SaveStructure field to given value.

### HasSaveStructure

`func (o *TestSuiteV2TreeModel) HasSaveStructure() bool`

HasSaveStructure returns a boolean if a field has been set.

### SetSaveStructureNil

`func (o *TestSuiteV2TreeModel) SetSaveStructureNil(b bool)`

 SetSaveStructureNil sets the value for SaveStructure to be an explicit nil

### UnsetSaveStructure
`func (o *TestSuiteV2TreeModel) UnsetSaveStructure()`

UnsetSaveStructure ensures that no value is present for SaveStructure, not even an explicit nil
### GetAutoRefresh

`func (o *TestSuiteV2TreeModel) GetAutoRefresh() bool`

GetAutoRefresh returns the AutoRefresh field if non-nil, zero value otherwise.

### GetAutoRefreshOk

`func (o *TestSuiteV2TreeModel) GetAutoRefreshOk() (*bool, bool)`

GetAutoRefreshOk returns a tuple with the AutoRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRefresh

`func (o *TestSuiteV2TreeModel) SetAutoRefresh(v bool)`

SetAutoRefresh sets AutoRefresh field to given value.

### HasAutoRefresh

`func (o *TestSuiteV2TreeModel) HasAutoRefresh() bool`

HasAutoRefresh returns a boolean if a field has been set.

### SetAutoRefreshNil

`func (o *TestSuiteV2TreeModel) SetAutoRefreshNil(b bool)`

 SetAutoRefreshNil sets the value for AutoRefresh to be an explicit nil

### UnsetAutoRefresh
`func (o *TestSuiteV2TreeModel) UnsetAutoRefresh()`

UnsetAutoRefresh ensures that no value is present for AutoRefresh, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


