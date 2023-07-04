# ApiV2TestSuitesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentId** | Pointer to **NullableString** | Unique ID of the parent test suite in hierarchy | [optional] 
**TestPlanId** | **string** | Unique ID of test plan to which the test suite belongs | 
**Name** | **string** | Name of the test suite | 
**Type** | Pointer to [**NullableTestSuiteType**](TestSuiteType.md) |  | [optional] 
**SaveStructure** | Pointer to **NullableBool** | Indicates if the test suite retains section tree structure | [optional] 
**AutoRefresh** | Pointer to **NullableBool** | Indicates if scheduled auto refresh is enabled for the test suite | [optional] 

## Methods

### NewApiV2TestSuitesPostRequest

`func NewApiV2TestSuitesPostRequest(testPlanId string, name string, ) *ApiV2TestSuitesPostRequest`

NewApiV2TestSuitesPostRequest instantiates a new ApiV2TestSuitesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2TestSuitesPostRequestWithDefaults

`func NewApiV2TestSuitesPostRequestWithDefaults() *ApiV2TestSuitesPostRequest`

NewApiV2TestSuitesPostRequestWithDefaults instantiates a new ApiV2TestSuitesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentId

`func (o *ApiV2TestSuitesPostRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ApiV2TestSuitesPostRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ApiV2TestSuitesPostRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ApiV2TestSuitesPostRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *ApiV2TestSuitesPostRequest) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *ApiV2TestSuitesPostRequest) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetTestPlanId

`func (o *ApiV2TestSuitesPostRequest) GetTestPlanId() string`

GetTestPlanId returns the TestPlanId field if non-nil, zero value otherwise.

### GetTestPlanIdOk

`func (o *ApiV2TestSuitesPostRequest) GetTestPlanIdOk() (*string, bool)`

GetTestPlanIdOk returns a tuple with the TestPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPlanId

`func (o *ApiV2TestSuitesPostRequest) SetTestPlanId(v string)`

SetTestPlanId sets TestPlanId field to given value.


### GetName

`func (o *ApiV2TestSuitesPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2TestSuitesPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2TestSuitesPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ApiV2TestSuitesPostRequest) GetType() TestSuiteType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiV2TestSuitesPostRequest) GetTypeOk() (*TestSuiteType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiV2TestSuitesPostRequest) SetType(v TestSuiteType)`

SetType sets Type field to given value.

### HasType

`func (o *ApiV2TestSuitesPostRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ApiV2TestSuitesPostRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ApiV2TestSuitesPostRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetSaveStructure

`func (o *ApiV2TestSuitesPostRequest) GetSaveStructure() bool`

GetSaveStructure returns the SaveStructure field if non-nil, zero value otherwise.

### GetSaveStructureOk

`func (o *ApiV2TestSuitesPostRequest) GetSaveStructureOk() (*bool, bool)`

GetSaveStructureOk returns a tuple with the SaveStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveStructure

`func (o *ApiV2TestSuitesPostRequest) SetSaveStructure(v bool)`

SetSaveStructure sets SaveStructure field to given value.

### HasSaveStructure

`func (o *ApiV2TestSuitesPostRequest) HasSaveStructure() bool`

HasSaveStructure returns a boolean if a field has been set.

### SetSaveStructureNil

`func (o *ApiV2TestSuitesPostRequest) SetSaveStructureNil(b bool)`

 SetSaveStructureNil sets the value for SaveStructure to be an explicit nil

### UnsetSaveStructure
`func (o *ApiV2TestSuitesPostRequest) UnsetSaveStructure()`

UnsetSaveStructure ensures that no value is present for SaveStructure, not even an explicit nil
### GetAutoRefresh

`func (o *ApiV2TestSuitesPostRequest) GetAutoRefresh() bool`

GetAutoRefresh returns the AutoRefresh field if non-nil, zero value otherwise.

### GetAutoRefreshOk

`func (o *ApiV2TestSuitesPostRequest) GetAutoRefreshOk() (*bool, bool)`

GetAutoRefreshOk returns a tuple with the AutoRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRefresh

`func (o *ApiV2TestSuitesPostRequest) SetAutoRefresh(v bool)`

SetAutoRefresh sets AutoRefresh field to given value.

### HasAutoRefresh

`func (o *ApiV2TestSuitesPostRequest) HasAutoRefresh() bool`

HasAutoRefresh returns a boolean if a field has been set.

### SetAutoRefreshNil

`func (o *ApiV2TestSuitesPostRequest) SetAutoRefreshNil(b bool)`

 SetAutoRefreshNil sets the value for AutoRefresh to be an explicit nil

### UnsetAutoRefresh
`func (o *ApiV2TestSuitesPostRequest) UnsetAutoRefresh()`

UnsetAutoRefresh ensures that no value is present for AutoRefresh, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


