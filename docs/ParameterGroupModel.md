# ParameterGroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Values** | **map[string]string** |  | 
**ParameterKeyId** | **string** |  | 

## Methods

### NewParameterGroupModel

`func NewParameterGroupModel(name string, values map[string]string, parameterKeyId string, ) *ParameterGroupModel`

NewParameterGroupModel instantiates a new ParameterGroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterGroupModelWithDefaults

`func NewParameterGroupModelWithDefaults() *ParameterGroupModel`

NewParameterGroupModelWithDefaults instantiates a new ParameterGroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ParameterGroupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParameterGroupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParameterGroupModel) SetName(v string)`

SetName sets Name field to given value.


### GetValues

`func (o *ParameterGroupModel) GetValues() map[string]string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ParameterGroupModel) GetValuesOk() (*map[string]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ParameterGroupModel) SetValues(v map[string]string)`

SetValues sets Values field to given value.


### GetParameterKeyId

`func (o *ParameterGroupModel) GetParameterKeyId() string`

GetParameterKeyId returns the ParameterKeyId field if non-nil, zero value otherwise.

### GetParameterKeyIdOk

`func (o *ParameterGroupModel) GetParameterKeyIdOk() (*string, bool)`

GetParameterKeyIdOk returns a tuple with the ParameterKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterKeyId

`func (o *ParameterGroupModel) SetParameterKeyId(v string)`

SetParameterKeyId sets ParameterKeyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


