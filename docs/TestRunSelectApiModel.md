# TestRunSelectApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**TestRunFilterApiModel**](TestRunFilterApiModel.md) |  | 
**ExtractionModel** | [**TestRunExtractionApiModel**](TestRunExtractionApiModel.md) | Rules for different level entities inclusion/exclusion | 

## Methods

### NewTestRunSelectApiModel

`func NewTestRunSelectApiModel(filter TestRunFilterApiModel, extractionModel TestRunExtractionApiModel, ) *TestRunSelectApiModel`

NewTestRunSelectApiModel instantiates a new TestRunSelectApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunSelectApiModelWithDefaults

`func NewTestRunSelectApiModelWithDefaults() *TestRunSelectApiModel`

NewTestRunSelectApiModelWithDefaults instantiates a new TestRunSelectApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *TestRunSelectApiModel) GetFilter() TestRunFilterApiModel`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TestRunSelectApiModel) GetFilterOk() (*TestRunFilterApiModel, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TestRunSelectApiModel) SetFilter(v TestRunFilterApiModel)`

SetFilter sets Filter field to given value.


### GetExtractionModel

`func (o *TestRunSelectApiModel) GetExtractionModel() TestRunExtractionApiModel`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *TestRunSelectApiModel) GetExtractionModelOk() (*TestRunExtractionApiModel, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *TestRunSelectApiModel) SetExtractionModel(v TestRunExtractionApiModel)`

SetExtractionModel sets ExtractionModel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


