# ConfigurationPostModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Parameters** | **map[string]string** |  | 
**ProjectId** | **string** | This property is used to link configuration with project | 
**IsDefault** | **bool** |  | 
**Name** | **string** |  | 

## Methods

### NewConfigurationPostModel

`func NewConfigurationPostModel(parameters map[string]string, projectId string, isDefault bool, name string, ) *ConfigurationPostModel`

NewConfigurationPostModel instantiates a new ConfigurationPostModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationPostModelWithDefaults

`func NewConfigurationPostModelWithDefaults() *ConfigurationPostModel`

NewConfigurationPostModelWithDefaults instantiates a new ConfigurationPostModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ConfigurationPostModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigurationPostModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigurationPostModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigurationPostModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ConfigurationPostModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConfigurationPostModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetParameters

`func (o *ConfigurationPostModel) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ConfigurationPostModel) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ConfigurationPostModel) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.


### GetProjectId

`func (o *ConfigurationPostModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ConfigurationPostModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ConfigurationPostModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetIsDefault

`func (o *ConfigurationPostModel) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ConfigurationPostModel) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ConfigurationPostModel) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetName

`func (o *ConfigurationPostModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigurationPostModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigurationPostModel) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


