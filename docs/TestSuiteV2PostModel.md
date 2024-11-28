# TestSuiteV2PostModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentId** | Pointer to **NullableString** | Unique ID of the parent test suite in hierarchy | [optional] 
**TestPlanId** | **string** | Unique ID of test plan to which the test suite belongs | 
**Name** | **string** | Name of the test suite | 
**Type** | Pointer to [**NullableTestSuiteType**](TestSuiteType.md) | Type of the test suite | [optional] 
**SaveStructure** | Pointer to **NullableBool** | Indicates if the test suite retains section tree structure | [optional] 
**AutoRefresh** | Pointer to **NullableBool** | Indicates if scheduled auto refresh is enabled for the test suite | [optional] 

## Methods

### NewTestSuiteV2PostModel

`func NewTestSuiteV2PostModel(testPlanId string, name string, ) *TestSuiteV2PostModel`

NewTestSuiteV2PostModel instantiates a new TestSuiteV2PostModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestSuiteV2PostModelWithDefaults

`func NewTestSuiteV2PostModelWithDefaults() *TestSuiteV2PostModel`

NewTestSuiteV2PostModelWithDefaults instantiates a new TestSuiteV2PostModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentId

`func (o *TestSuiteV2PostModel) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *TestSuiteV2PostModel) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *TestSuiteV2PostModel) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *TestSuiteV2PostModel) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *TestSuiteV2PostModel) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *TestSuiteV2PostModel) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetTestPlanId

`func (o *TestSuiteV2PostModel) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *TestSuiteV2PostModel) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *TestSuiteV2PostModel) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.


### GetName

`func (o *TestSuiteV2PostModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestSuiteV2PostModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestSuiteV2PostModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *TestSuiteV2PostModel) GetType() TestSuiteType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestSuiteV2PostModel) GetTypeOk() (*TestSuiteType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestSuiteV2PostModel) SetType(v TestSuiteType)`

SetType sets Type field to given value.

### HasType

`func (o *TestSuiteV2PostModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *TestSuiteV2PostModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TestSuiteV2PostModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetSaveStructure

`func (o *TestSuiteV2PostModel) GetSaveStructure() bool`

GetSaveStructure returns the SaveStructure field if non-nil, zero value otherwise.

### GetSaveStructureOk

`func (o *TestSuiteV2PostModel) GetSaveStructureOk() (*bool, bool)`

GetSaveStructureOk returns a tuple with the SaveStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveStructure

`func (o *TestSuiteV2PostModel) SetSaveStructure(v bool)`

SetSaveStructure sets SaveStructure field to given value.

### HasSaveStructure

`func (o *TestSuiteV2PostModel) HasSaveStructure() bool`

HasSaveStructure returns a boolean if a field has been set.

### SetSaveStructureNil

`func (o *TestSuiteV2PostModel) SetSaveStructureNil(b bool)`

 SetSaveStructureNil sets the value for SaveStructure to be an explicit nil

### UnsetSaveStructure
`func (o *TestSuiteV2PostModel) UnsetSaveStructure()`

UnsetSaveStructure ensures that no value is present for SaveStructure, not even an explicit nil
### GetAutoRefresh

`func (o *TestSuiteV2PostModel) GetAutoRefresh() bool`

GetAutoRefresh returns the AutoRefresh field if non-nil, zero value otherwise.

### GetAutoRefreshOk

`func (o *TestSuiteV2PostModel) GetAutoRefreshOk() (*bool, bool)`

GetAutoRefreshOk returns a tuple with the AutoRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRefresh

`func (o *TestSuiteV2PostModel) SetAutoRefresh(v bool)`

SetAutoRefresh sets AutoRefresh field to given value.

### HasAutoRefresh

`func (o *TestSuiteV2PostModel) HasAutoRefresh() bool`

HasAutoRefresh returns a boolean if a field has been set.

### SetAutoRefreshNil

`func (o *TestSuiteV2PostModel) SetAutoRefreshNil(b bool)`

 SetAutoRefreshNil sets the value for AutoRefresh to be an explicit nil

### UnsetAutoRefresh
`func (o *TestSuiteV2PostModel) UnsetAutoRefresh()`

UnsetAutoRefresh ensures that no value is present for AutoRefresh, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


