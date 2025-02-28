# ManualRerunTestResultApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestResultIds** | Pointer to [**NullableGuidExtractionModel**](GuidExtractionModel.md) | Set of extracted test result IDs | [optional] 

## Methods

### NewManualRerunTestResultApiModel

`func NewManualRerunTestResultApiModel() *ManualRerunTestResultApiModel`

NewManualRerunTestResultApiModel instantiates a new ManualRerunTestResultApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualRerunTestResultApiModelWithDefaults

`func NewManualRerunTestResultApiModelWithDefaults() *ManualRerunTestResultApiModel`

NewManualRerunTestResultApiModelWithDefaults instantiates a new ManualRerunTestResultApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestResultIds

`func (o *ManualRerunTestResultApiModel) GetTestResultIds() GuidExtractionModel`

GetTestResultIds returns the TestResultIds field if non-nil, zero value otherwise.

### GetTestResultIdsOk

`func (o *ManualRerunTestResultApiModel) GetTestResultIdsOk() (*GuidExtractionModel, bool)`

GetTestResultIdsOk returns a tuple with the TestResultIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResultIds

`func (o *ManualRerunTestResultApiModel) SetTestResultIds(v GuidExtractionModel)`

SetTestResultIds sets TestResultIds field to given value.

### HasTestResultIds

`func (o *ManualRerunTestResultApiModel) HasTestResultIds() bool`

HasTestResultIds returns a boolean if a field has been set.

### SetTestResultIdsNil

`func (o *ManualRerunTestResultApiModel) SetTestResultIdsNil(b bool)`

 SetTestResultIdsNil sets the value for TestResultIds to be an explicit nil

### UnsetTestResultIds
`func (o *ManualRerunTestResultApiModel) UnsetTestResultIds()`

UnsetTestResultIds ensures that no value is present for TestResultIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


