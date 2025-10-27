# FailureCategoryGroupItemApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**NullableFailureCategoryGroupApiResult**](FailureCategoryGroupApiResult.md) | Group details | [optional] 
**Items** | [**[]FailureCategoryItemApiResult**](FailureCategoryItemApiResult.md) | Group data | 

## Methods

### NewFailureCategoryGroupItemApiResult

`func NewFailureCategoryGroupItemApiResult(items []FailureCategoryItemApiResult, ) *FailureCategoryGroupItemApiResult`

NewFailureCategoryGroupItemApiResult instantiates a new FailureCategoryGroupItemApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureCategoryGroupItemApiResultWithDefaults

`func NewFailureCategoryGroupItemApiResultWithDefaults() *FailureCategoryGroupItemApiResult`

NewFailureCategoryGroupItemApiResultWithDefaults instantiates a new FailureCategoryGroupItemApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *FailureCategoryGroupItemApiResult) GetGroup() FailureCategoryGroupApiResult`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *FailureCategoryGroupItemApiResult) GetGroupOk() (*FailureCategoryGroupApiResult, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *FailureCategoryGroupItemApiResult) SetGroup(v FailureCategoryGroupApiResult)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *FailureCategoryGroupItemApiResult) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *FailureCategoryGroupItemApiResult) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *FailureCategoryGroupItemApiResult) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetItems

`func (o *FailureCategoryGroupItemApiResult) GetItems() []FailureCategoryItemApiResult`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FailureCategoryGroupItemApiResult) GetItemsOk() (*[]FailureCategoryItemApiResult, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FailureCategoryGroupItemApiResult) SetItems(v []FailureCategoryItemApiResult)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


