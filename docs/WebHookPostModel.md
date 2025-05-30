# WebHookPostModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | Unique ID of the webhook project | 
**EventType** | [**WebHookEventTypeModel**](WebHookEventTypeModel.md) | Type of event which triggers the webhook | 
**Description** | Pointer to **NullableString** | Description of the webhook | [optional] 
**Url** | **string** | Request URL of the webhook | 
**RequestType** | [**RequestTypeModel**](RequestTypeModel.md) | Request method of the webhook | 
**ShouldSendBody** | **bool** | Indicates if the webhook sends body | 
**Headers** | **map[string]string** | Collection of the webhook headers | 
**QueryParameters** | **map[string]string** | Collection of the webhook query parameters | 
**IsEnabled** | **bool** | Indicates if the webhook is active | 
**ShouldSendCustomBody** | **bool** | Indicates if the webhook sends custom body | 
**CustomBody** | Pointer to **NullableString** | Custom body of the webhook | [optional] 
**ShouldReplaceParameters** | **bool** | Indicates if the webhook injects parameters | 
**ShouldEscapeParameters** | **bool** | Indicates if the webhook escapes invalid characters in parameters | 
**Name** | **string** | Name of the webhook | 

## Methods

### NewWebHookPostModel

`func NewWebHookPostModel(projectId string, eventType WebHookEventTypeModel, url string, requestType RequestTypeModel, shouldSendBody bool, headers map[string]string, queryParameters map[string]string, isEnabled bool, shouldSendCustomBody bool, shouldReplaceParameters bool, shouldEscapeParameters bool, name string, ) *WebHookPostModel`

NewWebHookPostModel instantiates a new WebHookPostModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebHookPostModelWithDefaults

`func NewWebHookPostModelWithDefaults() *WebHookPostModel`

NewWebHookPostModelWithDefaults instantiates a new WebHookPostModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *WebHookPostModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WebHookPostModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WebHookPostModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetEventType

`func (o *WebHookPostModel) GetEventType() WebHookEventTypeModel`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WebHookPostModel) GetEventTypeOk() (*WebHookEventTypeModel, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WebHookPostModel) SetEventType(v WebHookEventTypeModel)`

SetEventType sets EventType field to given value.


### GetDescription

`func (o *WebHookPostModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebHookPostModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebHookPostModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebHookPostModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WebHookPostModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebHookPostModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUrl

`func (o *WebHookPostModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebHookPostModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebHookPostModel) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRequestType

`func (o *WebHookPostModel) GetRequestType() RequestTypeModel`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *WebHookPostModel) GetRequestTypeOk() (*RequestTypeModel, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *WebHookPostModel) SetRequestType(v RequestTypeModel)`

SetRequestType sets RequestType field to given value.


### GetShouldSendBody

`func (o *WebHookPostModel) GetShouldSendBody() bool`

GetShouldSendBody returns the ShouldSendBody field if non-nil, zero value otherwise.

### GetShouldSendBodyOk

`func (o *WebHookPostModel) GetShouldSendBodyOk() (*bool, bool)`

GetShouldSendBodyOk returns a tuple with the ShouldSendBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldSendBody

`func (o *WebHookPostModel) SetShouldSendBody(v bool)`

SetShouldSendBody sets ShouldSendBody field to given value.


### GetHeaders

`func (o *WebHookPostModel) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WebHookPostModel) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WebHookPostModel) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.


### GetQueryParameters

`func (o *WebHookPostModel) GetQueryParameters() map[string]string`

GetQueryParameters returns the QueryParameters field if non-nil, zero value otherwise.

### GetQueryParametersOk

`func (o *WebHookPostModel) GetQueryParametersOk() (*map[string]string, bool)`

GetQueryParametersOk returns a tuple with the QueryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParameters

`func (o *WebHookPostModel) SetQueryParameters(v map[string]string)`

SetQueryParameters sets QueryParameters field to given value.


### GetIsEnabled

`func (o *WebHookPostModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *WebHookPostModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *WebHookPostModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetShouldSendCustomBody

`func (o *WebHookPostModel) GetShouldSendCustomBody() bool`

GetShouldSendCustomBody returns the ShouldSendCustomBody field if non-nil, zero value otherwise.

### GetShouldSendCustomBodyOk

`func (o *WebHookPostModel) GetShouldSendCustomBodyOk() (*bool, bool)`

GetShouldSendCustomBodyOk returns a tuple with the ShouldSendCustomBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldSendCustomBody

`func (o *WebHookPostModel) SetShouldSendCustomBody(v bool)`

SetShouldSendCustomBody sets ShouldSendCustomBody field to given value.


### GetCustomBody

`func (o *WebHookPostModel) GetCustomBody() string`

GetCustomBody returns the CustomBody field if non-nil, zero value otherwise.

### GetCustomBodyOk

`func (o *WebHookPostModel) GetCustomBodyOk() (*string, bool)`

GetCustomBodyOk returns a tuple with the CustomBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBody

`func (o *WebHookPostModel) SetCustomBody(v string)`

SetCustomBody sets CustomBody field to given value.

### HasCustomBody

`func (o *WebHookPostModel) HasCustomBody() bool`

HasCustomBody returns a boolean if a field has been set.

### SetCustomBodyNil

`func (o *WebHookPostModel) SetCustomBodyNil(b bool)`

 SetCustomBodyNil sets the value for CustomBody to be an explicit nil

### UnsetCustomBody
`func (o *WebHookPostModel) UnsetCustomBody()`

UnsetCustomBody ensures that no value is present for CustomBody, not even an explicit nil
### GetShouldReplaceParameters

`func (o *WebHookPostModel) GetShouldReplaceParameters() bool`

GetShouldReplaceParameters returns the ShouldReplaceParameters field if non-nil, zero value otherwise.

### GetShouldReplaceParametersOk

`func (o *WebHookPostModel) GetShouldReplaceParametersOk() (*bool, bool)`

GetShouldReplaceParametersOk returns a tuple with the ShouldReplaceParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldReplaceParameters

`func (o *WebHookPostModel) SetShouldReplaceParameters(v bool)`

SetShouldReplaceParameters sets ShouldReplaceParameters field to given value.


### GetShouldEscapeParameters

`func (o *WebHookPostModel) GetShouldEscapeParameters() bool`

GetShouldEscapeParameters returns the ShouldEscapeParameters field if non-nil, zero value otherwise.

### GetShouldEscapeParametersOk

`func (o *WebHookPostModel) GetShouldEscapeParametersOk() (*bool, bool)`

GetShouldEscapeParametersOk returns a tuple with the ShouldEscapeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldEscapeParameters

`func (o *WebHookPostModel) SetShouldEscapeParameters(v bool)`

SetShouldEscapeParameters sets ShouldEscapeParameters field to given value.


### GetName

`func (o *WebHookPostModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebHookPostModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebHookPostModel) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


