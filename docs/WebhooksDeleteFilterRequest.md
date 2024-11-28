# WebhooksDeleteFilterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies a webhook name to search for | [optional] 
**EventTypes** | Pointer to [**[]WebHookEventTypeRequest**](WebHookEventTypeRequest.md) | Specifies a webhook event types to search for | [optional] 
**Methods** | Pointer to [**[]RequestTypeRequest**](RequestTypeRequest.md) | Specifies a webhook methods to search for | [optional] 
**ProjectIds** | Pointer to **[]string** | Specifies a webhook project IDs to search for | [optional] 
**IsEnabled** | Pointer to **NullableBool** | Specifies a webhook deleted status to search for | [optional] 

## Methods

### NewWebhooksDeleteFilterRequest

`func NewWebhooksDeleteFilterRequest() *WebhooksDeleteFilterRequest`

NewWebhooksDeleteFilterRequest instantiates a new WebhooksDeleteFilterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhooksDeleteFilterRequestWithDefaults

`func NewWebhooksDeleteFilterRequestWithDefaults() *WebhooksDeleteFilterRequest`

NewWebhooksDeleteFilterRequestWithDefaults instantiates a new WebhooksDeleteFilterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhooksDeleteFilterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhooksDeleteFilterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhooksDeleteFilterRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhooksDeleteFilterRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WebhooksDeleteFilterRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WebhooksDeleteFilterRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEventTypes

`func (o *WebhooksDeleteFilterRequest) GetEventTypes() []WebHookEventTypeRequest`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *WebhooksDeleteFilterRequest) GetEventTypesOk() (*[]WebHookEventTypeRequest, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *WebhooksDeleteFilterRequest) SetEventTypes(v []WebHookEventTypeRequest)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *WebhooksDeleteFilterRequest) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### SetEventTypesNil

`func (o *WebhooksDeleteFilterRequest) SetEventTypesNil(b bool)`

 SetEventTypesNil sets the value for EventTypes to be an explicit nil

### UnsetEventTypes
`func (o *WebhooksDeleteFilterRequest) UnsetEventTypes()`

UnsetEventTypes ensures that no value is present for EventTypes, not even an explicit nil
### GetMethods

`func (o *WebhooksDeleteFilterRequest) GetMethods() []RequestTypeRequest`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *WebhooksDeleteFilterRequest) GetMethodsOk() (*[]RequestTypeRequest, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *WebhooksDeleteFilterRequest) SetMethods(v []RequestTypeRequest)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *WebhooksDeleteFilterRequest) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### SetMethodsNil

`func (o *WebhooksDeleteFilterRequest) SetMethodsNil(b bool)`

 SetMethodsNil sets the value for Methods to be an explicit nil

### UnsetMethods
`func (o *WebhooksDeleteFilterRequest) UnsetMethods()`

UnsetMethods ensures that no value is present for Methods, not even an explicit nil
### GetProjectIds

`func (o *WebhooksDeleteFilterRequest) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *WebhooksDeleteFilterRequest) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *WebhooksDeleteFilterRequest) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *WebhooksDeleteFilterRequest) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *WebhooksDeleteFilterRequest) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *WebhooksDeleteFilterRequest) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetIsEnabled

`func (o *WebhooksDeleteFilterRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *WebhooksDeleteFilterRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *WebhooksDeleteFilterRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *WebhooksDeleteFilterRequest) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *WebhooksDeleteFilterRequest) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *WebhooksDeleteFilterRequest) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


