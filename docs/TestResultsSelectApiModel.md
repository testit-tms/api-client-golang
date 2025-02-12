# TestResultsSelectApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**TestResultsFilterApiModel**](TestResultsFilterApiModel.md) | Test result filters | 
**ExtractionModel** | [**TestResultsExtractionApiModel**](TestResultsExtractionApiModel.md) | Test results extraction model | 

## Methods

### NewTestResultsSelectApiModel

`func NewTestResultsSelectApiModel(filter TestResultsFilterApiModel, extractionModel TestResultsExtractionApiModel, ) *TestResultsSelectApiModel`

NewTestResultsSelectApiModel instantiates a new TestResultsSelectApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultsSelectApiModelWithDefaults

`func NewTestResultsSelectApiModelWithDefaults() *TestResultsSelectApiModel`

NewTestResultsSelectApiModelWithDefaults instantiates a new TestResultsSelectApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *TestResultsSelectApiModel) GetFilter() TestResultsFilterApiModel`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TestResultsSelectApiModel) GetFilterOk() (*TestResultsFilterApiModel, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TestResultsSelectApiModel) SetFilter(v TestResultsFilterApiModel)`

SetFilter sets Filter field to given value.


### GetExtractionModel

`func (o *TestResultsSelectApiModel) GetExtractionModel() TestResultsExtractionApiModel`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *TestResultsSelectApiModel) GetExtractionModelOk() (*TestResultsExtractionApiModel, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *TestResultsSelectApiModel) SetExtractionModel(v TestResultsExtractionApiModel)`

SetExtractionModel sets ExtractionModel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


