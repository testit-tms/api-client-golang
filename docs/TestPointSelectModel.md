# TestPointSelectModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**NullableTestPointFilterModel**](TestPointFilterModel.md) |  | [optional] 
**ExtractionModel** | Pointer to [**NullableTestPointsExtractionModel**](TestPointsExtractionModel.md) |  | [optional] 

## Methods

### NewTestPointSelectModel

`func NewTestPointSelectModel() *TestPointSelectModel`

NewTestPointSelectModel instantiates a new TestPointSelectModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPointSelectModelWithDefaults

`func NewTestPointSelectModelWithDefaults() *TestPointSelectModel`

NewTestPointSelectModelWithDefaults instantiates a new TestPointSelectModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *TestPointSelectModel) GetFilter() TestPointFilterModel`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TestPointSelectModel) GetFilterOk() (*TestPointFilterModel, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TestPointSelectModel) SetFilter(v TestPointFilterModel)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TestPointSelectModel) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *TestPointSelectModel) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *TestPointSelectModel) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetExtractionModel

`func (o *TestPointSelectModel) GetExtractionModel() TestPointsExtractionModel`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *TestPointSelectModel) GetExtractionModelOk() (*TestPointsExtractionModel, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *TestPointSelectModel) SetExtractionModel(v TestPointsExtractionModel)`

SetExtractionModel sets ExtractionModel field to given value.

### HasExtractionModel

`func (o *TestPointSelectModel) HasExtractionModel() bool`

HasExtractionModel returns a boolean if a field has been set.

### SetExtractionModelNil

`func (o *TestPointSelectModel) SetExtractionModelNil(b bool)`

 SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil

### UnsetExtractionModel
`func (o *TestPointSelectModel) UnsetExtractionModel()`

UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


