# UpdateMultipleLinksApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**ActionUpdate**](ActionUpdate.md) |  | 
**Links** | Pointer to [**[]CreateLinkApiModel**](CreateLinkApiModel.md) |  | [optional] 

## Methods

### NewUpdateMultipleLinksApiModel

`func NewUpdateMultipleLinksApiModel(action ActionUpdate, ) *UpdateMultipleLinksApiModel`

NewUpdateMultipleLinksApiModel instantiates a new UpdateMultipleLinksApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMultipleLinksApiModelWithDefaults

`func NewUpdateMultipleLinksApiModelWithDefaults() *UpdateMultipleLinksApiModel`

NewUpdateMultipleLinksApiModelWithDefaults instantiates a new UpdateMultipleLinksApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UpdateMultipleLinksApiModel) GetAction() ActionUpdate`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateMultipleLinksApiModel) GetActionOk() (*ActionUpdate, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateMultipleLinksApiModel) SetAction(v ActionUpdate)`

SetAction sets Action field to given value.


### GetLinks

`func (o *UpdateMultipleLinksApiModel) GetLinks() []CreateLinkApiModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UpdateMultipleLinksApiModel) GetLinksOk() (*[]CreateLinkApiModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UpdateMultipleLinksApiModel) SetLinks(v []CreateLinkApiModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UpdateMultipleLinksApiModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *UpdateMultipleLinksApiModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *UpdateMultipleLinksApiModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


