# ReplaceProjectExternalServiceApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewExternalServiceId** | **string** | The unique ID of the new external service that will replace the current one | 
**Settings** | Pointer to **interface{}** | External service settings | [optional] 

## Methods

### NewReplaceProjectExternalServiceApiModel

`func NewReplaceProjectExternalServiceApiModel(newExternalServiceId string, ) *ReplaceProjectExternalServiceApiModel`

NewReplaceProjectExternalServiceApiModel instantiates a new ReplaceProjectExternalServiceApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceProjectExternalServiceApiModelWithDefaults

`func NewReplaceProjectExternalServiceApiModelWithDefaults() *ReplaceProjectExternalServiceApiModel`

NewReplaceProjectExternalServiceApiModelWithDefaults instantiates a new ReplaceProjectExternalServiceApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewExternalServiceId

`func (o *ReplaceProjectExternalServiceApiModel) GetNewExternalServiceId() string`

GetNewExternalServiceId returns the NewExternalServiceId field if non-nil, zero value otherwise.

### GetNewExternalServiceIdOk

`func (o *ReplaceProjectExternalServiceApiModel) GetNewExternalServiceIdOk() (*string, bool)`

GetNewExternalServiceIdOk returns a tuple with the NewExternalServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewExternalServiceId

`func (o *ReplaceProjectExternalServiceApiModel) SetNewExternalServiceId(v string)`

SetNewExternalServiceId sets NewExternalServiceId field to given value.


### GetSettings

`func (o *ReplaceProjectExternalServiceApiModel) GetSettings() interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ReplaceProjectExternalServiceApiModel) GetSettingsOk() (*interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ReplaceProjectExternalServiceApiModel) SetSettings(v interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ReplaceProjectExternalServiceApiModel) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *ReplaceProjectExternalServiceApiModel) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *ReplaceProjectExternalServiceApiModel) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


