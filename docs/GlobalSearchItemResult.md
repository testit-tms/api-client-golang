# GlobalSearchItemResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** |  | 
**ResourceId** | **string** |  | 
**GlobalId** | Pointer to **NullableInt64** |  | [optional] 
**Name** | **string** |  | 
**ProjectGlobalId** | **int64** |  | 

## Methods

### NewGlobalSearchItemResult

`func NewGlobalSearchItemResult(resourceType string, resourceId string, name string, projectGlobalId int64, ) *GlobalSearchItemResult`

NewGlobalSearchItemResult instantiates a new GlobalSearchItemResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalSearchItemResultWithDefaults

`func NewGlobalSearchItemResultWithDefaults() *GlobalSearchItemResult`

NewGlobalSearchItemResultWithDefaults instantiates a new GlobalSearchItemResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *GlobalSearchItemResult) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *GlobalSearchItemResult) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *GlobalSearchItemResult) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetResourceId

`func (o *GlobalSearchItemResult) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *GlobalSearchItemResult) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *GlobalSearchItemResult) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetGlobalId

`func (o *GlobalSearchItemResult) GetGlobalId() int64`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *GlobalSearchItemResult) GetGlobalIdOk() (*int64, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *GlobalSearchItemResult) SetGlobalId(v int64)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *GlobalSearchItemResult) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### SetGlobalIdNil

`func (o *GlobalSearchItemResult) SetGlobalIdNil(b bool)`

 SetGlobalIdNil sets the value for GlobalId to be an explicit nil

### UnsetGlobalId
`func (o *GlobalSearchItemResult) UnsetGlobalId()`

UnsetGlobalId ensures that no value is present for GlobalId, not even an explicit nil
### GetName

`func (o *GlobalSearchItemResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalSearchItemResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalSearchItemResult) SetName(v string)`

SetName sets Name field to given value.


### GetProjectGlobalId

`func (o *GlobalSearchItemResult) GetProjectGlobalId() int64`

GetProjectGlobalId returns the ProjectGlobalId field if non-nil, zero value otherwise.

### GetProjectGlobalIdOk

`func (o *GlobalSearchItemResult) GetProjectGlobalIdOk() (*int64, bool)`

GetProjectGlobalIdOk returns a tuple with the ProjectGlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectGlobalId

`func (o *GlobalSearchItemResult) SetProjectGlobalId(v int64)`

SetProjectGlobalId sets ProjectGlobalId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


