# ApiV2TagsDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**NullableTagsFilterModel**](TagsFilterModel.md) |  | [optional] 
**ExtractionModel** | Pointer to [**NullableTagExtractionModel**](TagExtractionModel.md) |  | [optional] 

## Methods

### NewApiV2TagsDeleteRequest

`func NewApiV2TagsDeleteRequest() *ApiV2TagsDeleteRequest`

NewApiV2TagsDeleteRequest instantiates a new ApiV2TagsDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2TagsDeleteRequestWithDefaults

`func NewApiV2TagsDeleteRequestWithDefaults() *ApiV2TagsDeleteRequest`

NewApiV2TagsDeleteRequestWithDefaults instantiates a new ApiV2TagsDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ApiV2TagsDeleteRequest) GetFilter() TagsFilterModel`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ApiV2TagsDeleteRequest) GetFilterOk() (*TagsFilterModel, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ApiV2TagsDeleteRequest) SetFilter(v TagsFilterModel)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ApiV2TagsDeleteRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *ApiV2TagsDeleteRequest) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *ApiV2TagsDeleteRequest) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetExtractionModel

`func (o *ApiV2TagsDeleteRequest) GetExtractionModel() TagExtractionModel`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *ApiV2TagsDeleteRequest) GetExtractionModelOk() (*TagExtractionModel, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *ApiV2TagsDeleteRequest) SetExtractionModel(v TagExtractionModel)`

SetExtractionModel sets ExtractionModel field to given value.

### HasExtractionModel

`func (o *ApiV2TagsDeleteRequest) HasExtractionModel() bool`

HasExtractionModel returns a boolean if a field has been set.

### SetExtractionModelNil

`func (o *ApiV2TagsDeleteRequest) SetExtractionModelNil(b bool)`

 SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil

### UnsetExtractionModel
`func (o *ApiV2TagsDeleteRequest) UnsetExtractionModel()`

UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


