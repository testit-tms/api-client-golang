# SelectTagsApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**NullableTagsFilterApiModel**](TagsFilterApiModel.md) | Filters to select tags | [optional] 
**ExtractionModel** | Pointer to [**NullableTagsExtractionApiModel**](TagsExtractionApiModel.md) | Filters to extract tags | [optional] 

## Methods

### NewSelectTagsApiModel

`func NewSelectTagsApiModel() *SelectTagsApiModel`

NewSelectTagsApiModel instantiates a new SelectTagsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectTagsApiModelWithDefaults

`func NewSelectTagsApiModelWithDefaults() *SelectTagsApiModel`

NewSelectTagsApiModelWithDefaults instantiates a new SelectTagsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *SelectTagsApiModel) GetFilter() TagsFilterApiModel`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SelectTagsApiModel) GetFilterOk() (*TagsFilterApiModel, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SelectTagsApiModel) SetFilter(v TagsFilterApiModel)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *SelectTagsApiModel) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *SelectTagsApiModel) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *SelectTagsApiModel) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetExtractionModel

`func (o *SelectTagsApiModel) GetExtractionModel() TagsExtractionApiModel`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *SelectTagsApiModel) GetExtractionModelOk() (*TagsExtractionApiModel, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *SelectTagsApiModel) SetExtractionModel(v TagsExtractionApiModel)`

SetExtractionModel sets ExtractionModel field to given value.

### HasExtractionModel

`func (o *SelectTagsApiModel) HasExtractionModel() bool`

HasExtractionModel returns a boolean if a field has been set.

### SetExtractionModelNil

`func (o *SelectTagsApiModel) SetExtractionModelNil(b bool)`

 SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil

### UnsetExtractionModel
`func (o *SelectTagsApiModel) UnsetExtractionModel()`

UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


