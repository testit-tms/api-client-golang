# WorkItemSelectApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**WorkItemFilterApiModel**](WorkItemFilterApiModel.md) |  | 
**ExtractionModel** | Pointer to [**NullableWorkItemExtractionApiModel**](WorkItemExtractionApiModel.md) |  | [optional] 

## Methods

### NewWorkItemSelectApiModel

`func NewWorkItemSelectApiModel(filter WorkItemFilterApiModel, ) *WorkItemSelectApiModel`

NewWorkItemSelectApiModel instantiates a new WorkItemSelectApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemSelectApiModelWithDefaults

`func NewWorkItemSelectApiModelWithDefaults() *WorkItemSelectApiModel`

NewWorkItemSelectApiModelWithDefaults instantiates a new WorkItemSelectApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *WorkItemSelectApiModel) GetFilter() WorkItemFilterApiModel`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *WorkItemSelectApiModel) GetFilterOk() (*WorkItemFilterApiModel, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *WorkItemSelectApiModel) SetFilter(v WorkItemFilterApiModel)`

SetFilter sets Filter field to given value.


### GetExtractionModel

`func (o *WorkItemSelectApiModel) GetExtractionModel() WorkItemExtractionApiModel`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *WorkItemSelectApiModel) GetExtractionModelOk() (*WorkItemExtractionApiModel, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *WorkItemSelectApiModel) SetExtractionModel(v WorkItemExtractionApiModel)`

SetExtractionModel sets ExtractionModel field to given value.

### HasExtractionModel

`func (o *WorkItemSelectApiModel) HasExtractionModel() bool`

HasExtractionModel returns a boolean if a field has been set.

### SetExtractionModelNil

`func (o *WorkItemSelectApiModel) SetExtractionModelNil(b bool)`

 SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil

### UnsetExtractionModel
`func (o *WorkItemSelectApiModel) UnsetExtractionModel()`

UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


