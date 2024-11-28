# ConfigurationExtractionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to [**NullableGuidExtractionModel**](GuidExtractionModel.md) | Extraction parameters for configurations | [optional] 
**ProjectIds** | Pointer to [**NullableGuidExtractionModel**](GuidExtractionModel.md) | Extraction parameters for projects | [optional] 

## Methods

### NewConfigurationExtractionModel

`func NewConfigurationExtractionModel() *ConfigurationExtractionModel`

NewConfigurationExtractionModel instantiates a new ConfigurationExtractionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationExtractionModelWithDefaults

`func NewConfigurationExtractionModelWithDefaults() *ConfigurationExtractionModel`

NewConfigurationExtractionModelWithDefaults instantiates a new ConfigurationExtractionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *ConfigurationExtractionModel) GetIds() GuidExtractionModel`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ConfigurationExtractionModel) GetIdsOk() (*GuidExtractionModel, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ConfigurationExtractionModel) SetIds(v GuidExtractionModel)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ConfigurationExtractionModel) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *ConfigurationExtractionModel) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *ConfigurationExtractionModel) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetProjectIds

`func (o *ConfigurationExtractionModel) GetProjectIds() GuidExtractionModel`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *ConfigurationExtractionModel) GetProjectIdsOk() (*GuidExtractionModel, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *ConfigurationExtractionModel) SetProjectIds(v GuidExtractionModel)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *ConfigurationExtractionModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *ConfigurationExtractionModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *ConfigurationExtractionModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


