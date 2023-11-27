# WorkItemGroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **interface{}** |  | [optional] 
**Size** | **int32** |  | 
**WorkItems** | [**[]WorkItemShortModel**](WorkItemShortModel.md) |  | 

## Methods

### NewWorkItemGroupModel

`func NewWorkItemGroupModel(size int32, workItems []WorkItemShortModel, ) *WorkItemGroupModel`

NewWorkItemGroupModel instantiates a new WorkItemGroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemGroupModelWithDefaults

`func NewWorkItemGroupModelWithDefaults() *WorkItemGroupModel`

NewWorkItemGroupModelWithDefaults instantiates a new WorkItemGroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *WorkItemGroupModel) GetKey() interface{}`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WorkItemGroupModel) GetKeyOk() (*interface{}, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WorkItemGroupModel) SetKey(v interface{})`

SetKey sets Key field to given value.

### HasKey

`func (o *WorkItemGroupModel) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *WorkItemGroupModel) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *WorkItemGroupModel) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetSize

`func (o *WorkItemGroupModel) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *WorkItemGroupModel) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *WorkItemGroupModel) SetSize(v int32)`

SetSize sets Size field to given value.


### GetWorkItems

`func (o *WorkItemGroupModel) GetWorkItems() []WorkItemShortModel`

GetWorkItems returns the WorkItems field if non-nil, zero value otherwise.

### GetWorkItemsOk

`func (o *WorkItemGroupModel) GetWorkItemsOk() (*[]WorkItemShortModel, bool)`

GetWorkItemsOk returns a tuple with the WorkItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItems

`func (o *WorkItemGroupModel) SetWorkItems(v []WorkItemShortModel)`

SetWorkItems sets WorkItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


