# ConfigurationSelectApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**NullableConfigurationFilterApiModel**](ConfigurationFilterApiModel.md) | Configuration filters collection | [optional] 
**ExtractionModel** | Pointer to [**NullableConfigurationExtractionApiModel**](ConfigurationExtractionApiModel.md) | Rules for configurations extraction | [optional] 

## Methods

### NewConfigurationSelectApiModel

`func NewConfigurationSelectApiModel() *ConfigurationSelectApiModel`

NewConfigurationSelectApiModel instantiates a new ConfigurationSelectApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationSelectApiModelWithDefaults

`func NewConfigurationSelectApiModelWithDefaults() *ConfigurationSelectApiModel`

NewConfigurationSelectApiModelWithDefaults instantiates a new ConfigurationSelectApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ConfigurationSelectApiModel) GetFilter() ConfigurationFilterApiModel`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ConfigurationSelectApiModel) GetFilterOk() (*ConfigurationFilterApiModel, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ConfigurationSelectApiModel) SetFilter(v ConfigurationFilterApiModel)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ConfigurationSelectApiModel) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *ConfigurationSelectApiModel) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *ConfigurationSelectApiModel) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetExtractionModel

`func (o *ConfigurationSelectApiModel) GetExtractionModel() ConfigurationExtractionApiModel`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *ConfigurationSelectApiModel) GetExtractionModelOk() (*ConfigurationExtractionApiModel, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *ConfigurationSelectApiModel) SetExtractionModel(v ConfigurationExtractionApiModel)`

SetExtractionModel sets ExtractionModel field to given value.

### HasExtractionModel

`func (o *ConfigurationSelectApiModel) HasExtractionModel() bool`

HasExtractionModel returns a boolean if a field has been set.

### SetExtractionModelNil

`func (o *ConfigurationSelectApiModel) SetExtractionModelNil(b bool)`

 SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil

### UnsetExtractionModel
`func (o *ConfigurationSelectApiModel) UnsetExtractionModel()`

UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


