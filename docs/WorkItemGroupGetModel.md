# WorkItemGroupGetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectModel** | Pointer to [**NullableWorkItemGroupGetModelSelectModel**](WorkItemGroupGetModelSelectModel.md) |  | [optional] 
**GroupType** | [**WorkItemGroupType**](WorkItemGroupType.md) |  | 
**CustomAttributeId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWorkItemGroupGetModel

`func NewWorkItemGroupGetModel(groupType WorkItemGroupType, ) *WorkItemGroupGetModel`

NewWorkItemGroupGetModel instantiates a new WorkItemGroupGetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemGroupGetModelWithDefaults

`func NewWorkItemGroupGetModelWithDefaults() *WorkItemGroupGetModel`

NewWorkItemGroupGetModelWithDefaults instantiates a new WorkItemGroupGetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectModel

`func (o *WorkItemGroupGetModel) GetSelectModel() WorkItemGroupGetModelSelectModel`

GetSelectModel returns the SelectModel field if non-nil, zero value otherwise.

### GetSelectModelOk

`func (o *WorkItemGroupGetModel) GetSelectModelOk() (*WorkItemGroupGetModelSelectModel, bool)`

GetSelectModelOk returns a tuple with the SelectModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectModel

`func (o *WorkItemGroupGetModel) SetSelectModel(v WorkItemGroupGetModelSelectModel)`

SetSelectModel sets SelectModel field to given value.

### HasSelectModel

`func (o *WorkItemGroupGetModel) HasSelectModel() bool`

HasSelectModel returns a boolean if a field has been set.

### SetSelectModelNil

`func (o *WorkItemGroupGetModel) SetSelectModelNil(b bool)`

 SetSelectModelNil sets the value for SelectModel to be an explicit nil

### UnsetSelectModel
`func (o *WorkItemGroupGetModel) UnsetSelectModel()`

UnsetSelectModel ensures that no value is present for SelectModel, not even an explicit nil
### GetGroupType

`func (o *WorkItemGroupGetModel) GetGroupType() WorkItemGroupType`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *WorkItemGroupGetModel) GetGroupTypeOk() (*WorkItemGroupType, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *WorkItemGroupGetModel) SetGroupType(v WorkItemGroupType)`

SetGroupType sets GroupType field to given value.


### GetCustomAttributeId

`func (o *WorkItemGroupGetModel) GetCustomAttributeId() string`

GetCustomAttributeId returns the CustomAttributeId field if non-nil, zero value otherwise.

### GetCustomAttributeIdOk

`func (o *WorkItemGroupGetModel) GetCustomAttributeIdOk() (*string, bool)`

GetCustomAttributeIdOk returns a tuple with the CustomAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributeId

`func (o *WorkItemGroupGetModel) SetCustomAttributeId(v string)`

SetCustomAttributeId sets CustomAttributeId field to given value.

### HasCustomAttributeId

`func (o *WorkItemGroupGetModel) HasCustomAttributeId() bool`

HasCustomAttributeId returns a boolean if a field has been set.

### SetCustomAttributeIdNil

`func (o *WorkItemGroupGetModel) SetCustomAttributeIdNil(b bool)`

 SetCustomAttributeIdNil sets the value for CustomAttributeId to be an explicit nil

### UnsetCustomAttributeId
`func (o *WorkItemGroupGetModel) UnsetCustomAttributeId()`

UnsetCustomAttributeId ensures that no value is present for CustomAttributeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


