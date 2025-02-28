# IterationApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Parameters** | Pointer to [**[]ParameterShortApiResult**](ParameterShortApiResult.md) |  | [optional] 

## Methods

### NewIterationApiResult

`func NewIterationApiResult(id string, ) *IterationApiResult`

NewIterationApiResult instantiates a new IterationApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIterationApiResultWithDefaults

`func NewIterationApiResultWithDefaults() *IterationApiResult`

NewIterationApiResultWithDefaults instantiates a new IterationApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IterationApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IterationApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IterationApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetParameters

`func (o *IterationApiResult) GetParameters() []ParameterShortApiResult`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IterationApiResult) GetParametersOk() (*[]ParameterShortApiResult, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IterationApiResult) SetParameters(v []ParameterShortApiResult)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IterationApiResult) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *IterationApiResult) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *IterationApiResult) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


