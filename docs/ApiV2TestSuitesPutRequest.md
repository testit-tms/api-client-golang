# ApiV2TestSuitesPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**IsDeleted** | **bool** |  | 
**AutoRefresh** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewApiV2TestSuitesPutRequest

`func NewApiV2TestSuitesPutRequest(id string, name string, isDeleted bool, ) *ApiV2TestSuitesPutRequest`

NewApiV2TestSuitesPutRequest instantiates a new ApiV2TestSuitesPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2TestSuitesPutRequestWithDefaults

`func NewApiV2TestSuitesPutRequestWithDefaults() *ApiV2TestSuitesPutRequest`

NewApiV2TestSuitesPutRequestWithDefaults instantiates a new ApiV2TestSuitesPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiV2TestSuitesPutRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiV2TestSuitesPutRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiV2TestSuitesPutRequest) SetId(v string)`

SetId sets Id field to given value.


### GetParentId

`func (o *ApiV2TestSuitesPutRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ApiV2TestSuitesPutRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ApiV2TestSuitesPutRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ApiV2TestSuitesPutRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *ApiV2TestSuitesPutRequest) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *ApiV2TestSuitesPutRequest) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetName

`func (o *ApiV2TestSuitesPutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2TestSuitesPutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2TestSuitesPutRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsDeleted

`func (o *ApiV2TestSuitesPutRequest) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ApiV2TestSuitesPutRequest) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ApiV2TestSuitesPutRequest) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetAutoRefresh

`func (o *ApiV2TestSuitesPutRequest) GetAutoRefresh() bool`

GetAutoRefresh returns the AutoRefresh field if non-nil, zero value otherwise.

### GetAutoRefreshOk

`func (o *ApiV2TestSuitesPutRequest) GetAutoRefreshOk() (*bool, bool)`

GetAutoRefreshOk returns a tuple with the AutoRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRefresh

`func (o *ApiV2TestSuitesPutRequest) SetAutoRefresh(v bool)`

SetAutoRefresh sets AutoRefresh field to given value.

### HasAutoRefresh

`func (o *ApiV2TestSuitesPutRequest) HasAutoRefresh() bool`

HasAutoRefresh returns a boolean if a field has been set.

### SetAutoRefreshNil

`func (o *ApiV2TestSuitesPutRequest) SetAutoRefreshNil(b bool)`

 SetAutoRefreshNil sets the value for AutoRefresh to be an explicit nil

### UnsetAutoRefresh
`func (o *ApiV2TestSuitesPutRequest) UnsetAutoRefresh()`

UnsetAutoRefresh ensures that no value is present for AutoRefresh, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


