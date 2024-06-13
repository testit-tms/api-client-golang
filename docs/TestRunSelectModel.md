# TestRunSelectModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**TestRunFilterModel**](TestRunFilterModel.md) |  | 
**ExtractionModel** | [**TestRunExtractionModel**](TestRunExtractionModel.md) |  | 

## Methods

### NewTestRunSelectModel

`func NewTestRunSelectModel(filter TestRunFilterModel, extractionModel TestRunExtractionModel, ) *TestRunSelectModel`

NewTestRunSelectModel instantiates a new TestRunSelectModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunSelectModelWithDefaults

`func NewTestRunSelectModelWithDefaults() *TestRunSelectModel`

NewTestRunSelectModelWithDefaults instantiates a new TestRunSelectModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *TestRunSelectModel) GetFilter() TestRunFilterModel`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TestRunSelectModel) GetFilterOk() (*TestRunFilterModel, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TestRunSelectModel) SetFilter(v TestRunFilterModel)`

SetFilter sets Filter field to given value.


### GetExtractionModel

`func (o *TestRunSelectModel) GetExtractionModel() TestRunExtractionModel`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *TestRunSelectModel) GetExtractionModelOk() (*TestRunExtractionModel, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *TestRunSelectModel) SetExtractionModel(v TestRunExtractionModel)`

SetExtractionModel sets ExtractionModel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


