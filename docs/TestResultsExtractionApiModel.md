# TestResultsExtractionApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to [**NullableGuidExtractionModel**](GuidExtractionModel.md) | Set of extracted test result IDs | [optional] 

## Methods

### NewTestResultsExtractionApiModel

`func NewTestResultsExtractionApiModel() *TestResultsExtractionApiModel`

NewTestResultsExtractionApiModel instantiates a new TestResultsExtractionApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultsExtractionApiModelWithDefaults

`func NewTestResultsExtractionApiModelWithDefaults() *TestResultsExtractionApiModel`

NewTestResultsExtractionApiModelWithDefaults instantiates a new TestResultsExtractionApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *TestResultsExtractionApiModel) GetIds() GuidExtractionModel`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *TestResultsExtractionApiModel) GetIdsOk() (*GuidExtractionModel, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *TestResultsExtractionApiModel) SetIds(v GuidExtractionModel)`

SetIds sets Ids field to given value.

### HasIds

`func (o *TestResultsExtractionApiModel) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *TestResultsExtractionApiModel) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *TestResultsExtractionApiModel) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


