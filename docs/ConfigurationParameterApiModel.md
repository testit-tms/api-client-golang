# ConfigurationParameterApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the configuration parameter, must be unique | 
**Values** | [**[]ConfigurationParameterValueApiModel**](ConfigurationParameterValueApiModel.md) | List of possible configuration parameter values | 
**Projects** | Pointer to [**[]ConfigurationParameterProjectApiModel**](ConfigurationParameterProjectApiModel.md) | List of projects to which configuration parameter should be assigned | [optional] 

## Methods

### NewConfigurationParameterApiModel

`func NewConfigurationParameterApiModel(name string, values []ConfigurationParameterValueApiModel, ) *ConfigurationParameterApiModel`

NewConfigurationParameterApiModel instantiates a new ConfigurationParameterApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationParameterApiModelWithDefaults

`func NewConfigurationParameterApiModelWithDefaults() *ConfigurationParameterApiModel`

NewConfigurationParameterApiModelWithDefaults instantiates a new ConfigurationParameterApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConfigurationParameterApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigurationParameterApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigurationParameterApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetValues

`func (o *ConfigurationParameterApiModel) GetValues() []ConfigurationParameterValueApiModel`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ConfigurationParameterApiModel) GetValuesOk() (*[]ConfigurationParameterValueApiModel, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ConfigurationParameterApiModel) SetValues(v []ConfigurationParameterValueApiModel)`

SetValues sets Values field to given value.


### GetProjects

`func (o *ConfigurationParameterApiModel) GetProjects() []ConfigurationParameterProjectApiModel`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ConfigurationParameterApiModel) GetProjectsOk() (*[]ConfigurationParameterProjectApiModel, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ConfigurationParameterApiModel) SetProjects(v []ConfigurationParameterProjectApiModel)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *ConfigurationParameterApiModel) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### SetProjectsNil

`func (o *ConfigurationParameterApiModel) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *ConfigurationParameterApiModel) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


