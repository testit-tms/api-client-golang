# OpenIdConnectionClientShortModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**IsEnabled** | **bool** |  | 
**Settings** | Pointer to [**NullableOpenIdConnectionSettingsShortClientModel**](OpenIdConnectionSettingsShortClientModel.md) |  | [optional] 

## Methods

### NewOpenIdConnectionClientShortModel

`func NewOpenIdConnectionClientShortModel(id string, isEnabled bool, ) *OpenIdConnectionClientShortModel`

NewOpenIdConnectionClientShortModel instantiates a new OpenIdConnectionClientShortModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIdConnectionClientShortModelWithDefaults

`func NewOpenIdConnectionClientShortModelWithDefaults() *OpenIdConnectionClientShortModel`

NewOpenIdConnectionClientShortModelWithDefaults instantiates a new OpenIdConnectionClientShortModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpenIdConnectionClientShortModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenIdConnectionClientShortModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenIdConnectionClientShortModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *OpenIdConnectionClientShortModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenIdConnectionClientShortModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenIdConnectionClientShortModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenIdConnectionClientShortModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OpenIdConnectionClientShortModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OpenIdConnectionClientShortModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsEnabled

`func (o *OpenIdConnectionClientShortModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *OpenIdConnectionClientShortModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *OpenIdConnectionClientShortModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetSettings

`func (o *OpenIdConnectionClientShortModel) GetSettings() OpenIdConnectionSettingsShortClientModel`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *OpenIdConnectionClientShortModel) GetSettingsOk() (*OpenIdConnectionSettingsShortClientModel, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *OpenIdConnectionClientShortModel) SetSettings(v OpenIdConnectionSettingsShortClientModel)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *OpenIdConnectionClientShortModel) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *OpenIdConnectionClientShortModel) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *OpenIdConnectionClientShortModel) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


