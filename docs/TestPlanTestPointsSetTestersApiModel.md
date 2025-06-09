# TestPlanTestPointsSetTestersApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**NullableTestPlanTestPointsSearchApiModel**](TestPlanTestPointsSearchApiModel.md) |  | [optional] 
**ExtractionModel** | Pointer to [**NullableTestPlanTestPointsExtractionApiModel**](TestPlanTestPointsExtractionApiModel.md) |  | [optional] 
**TesterIds** | **[]string** |  | 

## Methods

### NewTestPlanTestPointsSetTestersApiModel

`func NewTestPlanTestPointsSetTestersApiModel(testerIds []string, ) *TestPlanTestPointsSetTestersApiModel`

NewTestPlanTestPointsSetTestersApiModel instantiates a new TestPlanTestPointsSetTestersApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanTestPointsSetTestersApiModelWithDefaults

`func NewTestPlanTestPointsSetTestersApiModelWithDefaults() *TestPlanTestPointsSetTestersApiModel`

NewTestPlanTestPointsSetTestersApiModelWithDefaults instantiates a new TestPlanTestPointsSetTestersApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *TestPlanTestPointsSetTestersApiModel) GetFilter() TestPlanTestPointsSearchApiModel`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TestPlanTestPointsSetTestersApiModel) GetFilterOk() (*TestPlanTestPointsSearchApiModel, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TestPlanTestPointsSetTestersApiModel) SetFilter(v TestPlanTestPointsSearchApiModel)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TestPlanTestPointsSetTestersApiModel) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *TestPlanTestPointsSetTestersApiModel) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *TestPlanTestPointsSetTestersApiModel) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetExtractionModel

`func (o *TestPlanTestPointsSetTestersApiModel) GetExtractionModel() TestPlanTestPointsExtractionApiModel`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *TestPlanTestPointsSetTestersApiModel) GetExtractionModelOk() (*TestPlanTestPointsExtractionApiModel, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *TestPlanTestPointsSetTestersApiModel) SetExtractionModel(v TestPlanTestPointsExtractionApiModel)`

SetExtractionModel sets ExtractionModel field to given value.

### HasExtractionModel

`func (o *TestPlanTestPointsSetTestersApiModel) HasExtractionModel() bool`

HasExtractionModel returns a boolean if a field has been set.

### SetExtractionModelNil

`func (o *TestPlanTestPointsSetTestersApiModel) SetExtractionModelNil(b bool)`

 SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil

### UnsetExtractionModel
`func (o *TestPlanTestPointsSetTestersApiModel) UnsetExtractionModel()`

UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil
### GetTesterIds

`func (o *TestPlanTestPointsSetTestersApiModel) GetTesterIds() []string`

GetTesterIds returns the TesterIds field if non-nil, zero value otherwise.

### GetTesterIdsOk

`func (o *TestPlanTestPointsSetTestersApiModel) GetTesterIdsOk() (*[]string, bool)`

GetTesterIdsOk returns a tuple with the TesterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesterIds

`func (o *TestPlanTestPointsSetTestersApiModel) SetTesterIds(v []string)`

SetTesterIds sets TesterIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


