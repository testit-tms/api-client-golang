# AutoTestSelectModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**AutoTestFilterModel**](AutoTestFilterModel.md) |  | 
**ExtractionModel** | Pointer to [**NullableAutoTestsExtractionModel**](AutoTestsExtractionModel.md) |  | [optional] 

## Methods

### NewAutoTestSelectModel

`func NewAutoTestSelectModel(filter AutoTestFilterModel, ) *AutoTestSelectModel`

NewAutoTestSelectModel instantiates a new AutoTestSelectModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestSelectModelWithDefaults

`func NewAutoTestSelectModelWithDefaults() *AutoTestSelectModel`

NewAutoTestSelectModelWithDefaults instantiates a new AutoTestSelectModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *AutoTestSelectModel) GetFilter() AutoTestFilterModel`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AutoTestSelectModel) GetFilterOk() (*AutoTestFilterModel, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AutoTestSelectModel) SetFilter(v AutoTestFilterModel)`

SetFilter sets Filter field to given value.


### GetExtractionModel

`func (o *AutoTestSelectModel) GetExtractionModel() AutoTestsExtractionModel`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *AutoTestSelectModel) GetExtractionModelOk() (*AutoTestsExtractionModel, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *AutoTestSelectModel) SetExtractionModel(v AutoTestsExtractionModel)`

SetExtractionModel sets ExtractionModel field to given value.

### HasExtractionModel

`func (o *AutoTestSelectModel) HasExtractionModel() bool`

HasExtractionModel returns a boolean if a field has been set.

### SetExtractionModelNil

`func (o *AutoTestSelectModel) SetExtractionModelNil(b bool)`

 SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil

### UnsetExtractionModel
`func (o *AutoTestSelectModel) UnsetExtractionModel()`

UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


