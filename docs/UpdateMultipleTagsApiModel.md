# UpdateMultipleTagsApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**ActionUpdate**](ActionUpdate.md) | The action that specifies which operation should be performed on the supplied tags. | 
**Tags** | Pointer to [**[]TagApiModel**](TagApiModel.md) | The collection of tag names to be processed. | [optional] 

## Methods

### NewUpdateMultipleTagsApiModel

`func NewUpdateMultipleTagsApiModel(action ActionUpdate, ) *UpdateMultipleTagsApiModel`

NewUpdateMultipleTagsApiModel instantiates a new UpdateMultipleTagsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMultipleTagsApiModelWithDefaults

`func NewUpdateMultipleTagsApiModelWithDefaults() *UpdateMultipleTagsApiModel`

NewUpdateMultipleTagsApiModelWithDefaults instantiates a new UpdateMultipleTagsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UpdateMultipleTagsApiModel) GetAction() ActionUpdate`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateMultipleTagsApiModel) GetActionOk() (*ActionUpdate, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateMultipleTagsApiModel) SetAction(v ActionUpdate)`

SetAction sets Action field to given value.


### GetTags

`func (o *UpdateMultipleTagsApiModel) GetTags() []TagApiModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateMultipleTagsApiModel) GetTagsOk() (*[]TagApiModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateMultipleTagsApiModel) SetTags(v []TagApiModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateMultipleTagsApiModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *UpdateMultipleTagsApiModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *UpdateMultipleTagsApiModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


