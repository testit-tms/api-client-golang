# GlobalSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** |  | 
**ResourceType** | Pointer to **NullableString** |  | [optional] 
**Take** | **int32** |  | 
**Skip** | **int32** |  | 

## Methods

### NewGlobalSearchRequest

`func NewGlobalSearchRequest(query string, take int32, skip int32, ) *GlobalSearchRequest`

NewGlobalSearchRequest instantiates a new GlobalSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalSearchRequestWithDefaults

`func NewGlobalSearchRequestWithDefaults() *GlobalSearchRequest`

NewGlobalSearchRequestWithDefaults instantiates a new GlobalSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *GlobalSearchRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GlobalSearchRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GlobalSearchRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetResourceType

`func (o *GlobalSearchRequest) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *GlobalSearchRequest) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *GlobalSearchRequest) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *GlobalSearchRequest) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *GlobalSearchRequest) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *GlobalSearchRequest) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetTake

`func (o *GlobalSearchRequest) GetTake() int32`

GetTake returns the Take field if non-nil, zero value otherwise.

### GetTakeOk

`func (o *GlobalSearchRequest) GetTakeOk() (*int32, bool)`

GetTakeOk returns a tuple with the Take field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTake

`func (o *GlobalSearchRequest) SetTake(v int32)`

SetTake sets Take field to given value.


### GetSkip

`func (o *GlobalSearchRequest) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *GlobalSearchRequest) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *GlobalSearchRequest) SetSkip(v int32)`

SetSkip sets Skip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


