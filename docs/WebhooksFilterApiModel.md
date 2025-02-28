# WebhooksFilterApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies a webhook name to search for | [optional] 
**EventTypes** | Pointer to [**[]WebHookEventTypeRequest**](WebHookEventTypeRequest.md) | Specifies a webhook event types to search for | [optional] 
**Methods** | Pointer to [**[]RequestTypeApiModel**](RequestTypeApiModel.md) | Specifies a webhook methods to search for | [optional] 
**ProjectIds** | Pointer to **[]string** | Specifies a webhook project IDs to search for | [optional] 

## Methods

### NewWebhooksFilterApiModel

`func NewWebhooksFilterApiModel() *WebhooksFilterApiModel`

NewWebhooksFilterApiModel instantiates a new WebhooksFilterApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhooksFilterApiModelWithDefaults

`func NewWebhooksFilterApiModelWithDefaults() *WebhooksFilterApiModel`

NewWebhooksFilterApiModelWithDefaults instantiates a new WebhooksFilterApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhooksFilterApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhooksFilterApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhooksFilterApiModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhooksFilterApiModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WebhooksFilterApiModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WebhooksFilterApiModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEventTypes

`func (o *WebhooksFilterApiModel) GetEventTypes() []WebHookEventTypeRequest`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *WebhooksFilterApiModel) GetEventTypesOk() (*[]WebHookEventTypeRequest, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *WebhooksFilterApiModel) SetEventTypes(v []WebHookEventTypeRequest)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *WebhooksFilterApiModel) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### SetEventTypesNil

`func (o *WebhooksFilterApiModel) SetEventTypesNil(b bool)`

 SetEventTypesNil sets the value for EventTypes to be an explicit nil

### UnsetEventTypes
`func (o *WebhooksFilterApiModel) UnsetEventTypes()`

UnsetEventTypes ensures that no value is present for EventTypes, not even an explicit nil
### GetMethods

`func (o *WebhooksFilterApiModel) GetMethods() []RequestTypeApiModel`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *WebhooksFilterApiModel) GetMethodsOk() (*[]RequestTypeApiModel, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *WebhooksFilterApiModel) SetMethods(v []RequestTypeApiModel)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *WebhooksFilterApiModel) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### SetMethodsNil

`func (o *WebhooksFilterApiModel) SetMethodsNil(b bool)`

 SetMethodsNil sets the value for Methods to be an explicit nil

### UnsetMethods
`func (o *WebhooksFilterApiModel) UnsetMethods()`

UnsetMethods ensures that no value is present for Methods, not even an explicit nil
### GetProjectIds

`func (o *WebhooksFilterApiModel) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *WebhooksFilterApiModel) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *WebhooksFilterApiModel) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *WebhooksFilterApiModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *WebhooksFilterApiModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *WebhooksFilterApiModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


