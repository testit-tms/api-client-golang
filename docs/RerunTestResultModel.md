# RerunTestResultModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Outcome** | Pointer to **NullableString** |  | [optional] 
**RunNumber** | **int32** |  | 

## Methods

### NewRerunTestResultModel

`func NewRerunTestResultModel(id string, runNumber int32, ) *RerunTestResultModel`

NewRerunTestResultModel instantiates a new RerunTestResultModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRerunTestResultModelWithDefaults

`func NewRerunTestResultModelWithDefaults() *RerunTestResultModel`

NewRerunTestResultModelWithDefaults instantiates a new RerunTestResultModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RerunTestResultModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RerunTestResultModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RerunTestResultModel) SetId(v string)`

SetId sets Id field to given value.


### GetOutcome

`func (o *RerunTestResultModel) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *RerunTestResultModel) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *RerunTestResultModel) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *RerunTestResultModel) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *RerunTestResultModel) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *RerunTestResultModel) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetRunNumber

`func (o *RerunTestResultModel) GetRunNumber() int32`

GetRunNumber returns the RunNumber field if non-nil, zero value otherwise.

### GetRunNumberOk

`func (o *RerunTestResultModel) GetRunNumberOk() (*int32, bool)`

GetRunNumberOk returns a tuple with the RunNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunNumber

`func (o *RerunTestResultModel) SetRunNumber(v int32)`

SetRunNumber sets RunNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


