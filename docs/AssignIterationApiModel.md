# AssignIterationApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | [**[]ParameterIterationModel**](ParameterIterationModel.md) |  | 
**Id** | **string** | Iteration identifier, must be empty for new or changed iteration | 

## Methods

### NewAssignIterationApiModel

`func NewAssignIterationApiModel(parameters []ParameterIterationModel, id string, ) *AssignIterationApiModel`

NewAssignIterationApiModel instantiates a new AssignIterationApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignIterationApiModelWithDefaults

`func NewAssignIterationApiModelWithDefaults() *AssignIterationApiModel`

NewAssignIterationApiModelWithDefaults instantiates a new AssignIterationApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *AssignIterationApiModel) GetParameters() []ParameterIterationModel`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AssignIterationApiModel) GetParametersOk() (*[]ParameterIterationModel, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AssignIterationApiModel) SetParameters(v []ParameterIterationModel)`

SetParameters sets Parameters field to given value.


### GetId

`func (o *AssignIterationApiModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssignIterationApiModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssignIterationApiModel) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


