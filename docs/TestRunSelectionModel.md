# TestRunSelectionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**NullableTestRunFilterModel**](TestRunFilterModel.md) |  | [optional] 
**ExtractionModel** | Pointer to [**NullableTestRunSelectModelExtractionModel**](TestRunSelectModelExtractionModel.md) |  | [optional] 

## Methods

### NewTestRunSelectionModel

`func NewTestRunSelectionModel() *TestRunSelectionModel`

NewTestRunSelectionModel instantiates a new TestRunSelectionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunSelectionModelWithDefaults

`func NewTestRunSelectionModelWithDefaults() *TestRunSelectionModel`

NewTestRunSelectionModelWithDefaults instantiates a new TestRunSelectionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *TestRunSelectionModel) GetFilter() TestRunFilterModel`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TestRunSelectionModel) GetFilterOk() (*TestRunFilterModel, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TestRunSelectionModel) SetFilter(v TestRunFilterModel)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TestRunSelectionModel) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *TestRunSelectionModel) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *TestRunSelectionModel) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetExtractionModel

`func (o *TestRunSelectionModel) GetExtractionModel() TestRunSelectModelExtractionModel`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *TestRunSelectionModel) GetExtractionModelOk() (*TestRunSelectModelExtractionModel, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *TestRunSelectionModel) SetExtractionModel(v TestRunSelectModelExtractionModel)`

SetExtractionModel sets ExtractionModel field to given value.

### HasExtractionModel

`func (o *TestRunSelectionModel) HasExtractionModel() bool`

HasExtractionModel returns a boolean if a field has been set.

### SetExtractionModelNil

`func (o *TestRunSelectionModel) SetExtractionModelNil(b bool)`

 SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil

### UnsetExtractionModel
`func (o *TestRunSelectionModel) UnsetExtractionModel()`

UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


