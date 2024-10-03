# ApiV2WebhooksSearchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies a webhook name to search for | [optional] 
**EventTypes** | Pointer to [**[]WebHookEventTypeModel**](WebHookEventTypeModel.md) | Specifies a webhook event types to search for | [optional] 
**Methods** | Pointer to [**[]RequestTypeModel**](RequestTypeModel.md) | Specifies a webhook methods to search for | [optional] 
**ProjectIds** | Pointer to **[]string** | Specifies a webhook project IDs to search for | [optional] 
**IsEnabled** | Pointer to **NullableBool** | Specifies a webhook deleted status to search for | [optional] 

## Methods

### NewApiV2WebhooksSearchPostRequest

`func NewApiV2WebhooksSearchPostRequest() *ApiV2WebhooksSearchPostRequest`

NewApiV2WebhooksSearchPostRequest instantiates a new ApiV2WebhooksSearchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2WebhooksSearchPostRequestWithDefaults

`func NewApiV2WebhooksSearchPostRequestWithDefaults() *ApiV2WebhooksSearchPostRequest`

NewApiV2WebhooksSearchPostRequestWithDefaults instantiates a new ApiV2WebhooksSearchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiV2WebhooksSearchPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2WebhooksSearchPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2WebhooksSearchPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV2WebhooksSearchPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApiV2WebhooksSearchPostRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApiV2WebhooksSearchPostRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEventTypes

`func (o *ApiV2WebhooksSearchPostRequest) GetEventTypes() []WebHookEventTypeModel`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *ApiV2WebhooksSearchPostRequest) GetEventTypesOk() (*[]WebHookEventTypeModel, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *ApiV2WebhooksSearchPostRequest) SetEventTypes(v []WebHookEventTypeModel)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *ApiV2WebhooksSearchPostRequest) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### SetEventTypesNil

`func (o *ApiV2WebhooksSearchPostRequest) SetEventTypesNil(b bool)`

 SetEventTypesNil sets the value for EventTypes to be an explicit nil

### UnsetEventTypes
`func (o *ApiV2WebhooksSearchPostRequest) UnsetEventTypes()`

UnsetEventTypes ensures that no value is present for EventTypes, not even an explicit nil
### GetMethods

`func (o *ApiV2WebhooksSearchPostRequest) GetMethods() []RequestTypeModel`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *ApiV2WebhooksSearchPostRequest) GetMethodsOk() (*[]RequestTypeModel, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *ApiV2WebhooksSearchPostRequest) SetMethods(v []RequestTypeModel)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *ApiV2WebhooksSearchPostRequest) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### SetMethodsNil

`func (o *ApiV2WebhooksSearchPostRequest) SetMethodsNil(b bool)`

 SetMethodsNil sets the value for Methods to be an explicit nil

### UnsetMethods
`func (o *ApiV2WebhooksSearchPostRequest) UnsetMethods()`

UnsetMethods ensures that no value is present for Methods, not even an explicit nil
### GetProjectIds

`func (o *ApiV2WebhooksSearchPostRequest) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *ApiV2WebhooksSearchPostRequest) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *ApiV2WebhooksSearchPostRequest) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *ApiV2WebhooksSearchPostRequest) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *ApiV2WebhooksSearchPostRequest) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *ApiV2WebhooksSearchPostRequest) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetIsEnabled

`func (o *ApiV2WebhooksSearchPostRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ApiV2WebhooksSearchPostRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ApiV2WebhooksSearchPostRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ApiV2WebhooksSearchPostRequest) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *ApiV2WebhooksSearchPostRequest) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *ApiV2WebhooksSearchPostRequest) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


