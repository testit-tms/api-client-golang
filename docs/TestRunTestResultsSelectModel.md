# TestRunTestResultsSelectModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**NullableTestResultsLocalFilterModel**](TestResultsLocalFilterModel.md) | Collection of filters to apply to search | [optional] 
**TestResultIdsExtractionModel** | Pointer to [**NullableGuidExtractionModel**](GuidExtractionModel.md) | Rules to include and exclude certain entities in result | [optional] 

## Methods

### NewTestRunTestResultsSelectModel

`func NewTestRunTestResultsSelectModel() *TestRunTestResultsSelectModel`

NewTestRunTestResultsSelectModel instantiates a new TestRunTestResultsSelectModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunTestResultsSelectModelWithDefaults

`func NewTestRunTestResultsSelectModelWithDefaults() *TestRunTestResultsSelectModel`

NewTestRunTestResultsSelectModelWithDefaults instantiates a new TestRunTestResultsSelectModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *TestRunTestResultsSelectModel) GetFilter() TestResultsLocalFilterModel`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TestRunTestResultsSelectModel) GetFilterOk() (*TestResultsLocalFilterModel, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TestRunTestResultsSelectModel) SetFilter(v TestResultsLocalFilterModel)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TestRunTestResultsSelectModel) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *TestRunTestResultsSelectModel) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *TestRunTestResultsSelectModel) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetTestResultIdsExtractionModel

`func (o *TestRunTestResultsSelectModel) GetTestResultIdsExtractionModel() GuidExtractionModel`

GetTestResultIdsExtractionModel returns the TestResultIdsExtractionModel field if non-nil, zero value otherwise.

### GetTestResultIdsExtractionModelOk

`func (o *TestRunTestResultsSelectModel) GetTestResultIdsExtractionModelOk() (*GuidExtractionModel, bool)`

GetTestResultIdsExtractionModelOk returns a tuple with the TestResultIdsExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResultIdsExtractionModel

`func (o *TestRunTestResultsSelectModel) SetTestResultIdsExtractionModel(v GuidExtractionModel)`

SetTestResultIdsExtractionModel sets TestResultIdsExtractionModel field to given value.

### HasTestResultIdsExtractionModel

`func (o *TestRunTestResultsSelectModel) HasTestResultIdsExtractionModel() bool`

HasTestResultIdsExtractionModel returns a boolean if a field has been set.

### SetTestResultIdsExtractionModelNil

`func (o *TestRunTestResultsSelectModel) SetTestResultIdsExtractionModelNil(b bool)`

 SetTestResultIdsExtractionModelNil sets the value for TestResultIdsExtractionModel to be an explicit nil

### UnsetTestResultIdsExtractionModel
`func (o *TestRunTestResultsSelectModel) UnsetTestResultIdsExtractionModel()`

UnsetTestResultIdsExtractionModel ensures that no value is present for TestResultIdsExtractionModel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


