# ApiV2WebhooksPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | Unique ID of the webhook project | 
**EventType** | [**WebHookEventTypeModel**](WebHookEventTypeModel.md) |  | 
**Description** | Pointer to **NullableString** | Description of the webhook | [optional] 
**Url** | **string** | Request URL of the webhook | 
**RequestType** | [**RequestTypeModel**](RequestTypeModel.md) |  | 
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

### NewApiV2WebhooksPostRequest

`func NewApiV2WebhooksPostRequest(projectId string, eventType WebHookEventTypeModel, url string, requestType RequestTypeModel, shouldSendBody bool, headers map[string]string, queryParameters map[string]string, isEnabled bool, shouldSendCustomBody bool, shouldReplaceParameters bool, shouldEscapeParameters bool, name string, ) *ApiV2WebhooksPostRequest`

NewApiV2WebhooksPostRequest instantiates a new ApiV2WebhooksPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2WebhooksPostRequestWithDefaults

`func NewApiV2WebhooksPostRequestWithDefaults() *ApiV2WebhooksPostRequest`

NewApiV2WebhooksPostRequestWithDefaults instantiates a new ApiV2WebhooksPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ApiV2WebhooksPostRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ApiV2WebhooksPostRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ApiV2WebhooksPostRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetEventType

`func (o *ApiV2WebhooksPostRequest) GetEventType() WebHookEventTypeModel`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ApiV2WebhooksPostRequest) GetEventTypeOk() (*WebHookEventTypeModel, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ApiV2WebhooksPostRequest) SetEventType(v WebHookEventTypeModel)`

SetEventType sets EventType field to given value.


### GetDescription

`func (o *ApiV2WebhooksPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiV2WebhooksPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiV2WebhooksPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiV2WebhooksPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApiV2WebhooksPostRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApiV2WebhooksPostRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUrl

`func (o *ApiV2WebhooksPostRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApiV2WebhooksPostRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApiV2WebhooksPostRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRequestType

`func (o *ApiV2WebhooksPostRequest) GetRequestType() RequestTypeModel`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *ApiV2WebhooksPostRequest) GetRequestTypeOk() (*RequestTypeModel, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *ApiV2WebhooksPostRequest) SetRequestType(v RequestTypeModel)`

SetRequestType sets RequestType field to given value.


### GetShouldSendBody

`func (o *ApiV2WebhooksPostRequest) GetShouldSendBody() bool`

GetShouldSendBody returns the ShouldSendBody field if non-nil, zero value otherwise.

### GetShouldSendBodyOk

`func (o *ApiV2WebhooksPostRequest) GetShouldSendBodyOk() (*bool, bool)`

GetShouldSendBodyOk returns a tuple with the ShouldSendBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldSendBody

`func (o *ApiV2WebhooksPostRequest) SetShouldSendBody(v bool)`

SetShouldSendBody sets ShouldSendBody field to given value.


### GetHeaders

`func (o *ApiV2WebhooksPostRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ApiV2WebhooksPostRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ApiV2WebhooksPostRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.


### GetQueryParameters

`func (o *ApiV2WebhooksPostRequest) GetQueryParameters() map[string]string`

GetQueryParameters returns the QueryParameters field if non-nil, zero value otherwise.

### GetQueryParametersOk

`func (o *ApiV2WebhooksPostRequest) GetQueryParametersOk() (*map[string]string, bool)`

GetQueryParametersOk returns a tuple with the QueryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParameters

`func (o *ApiV2WebhooksPostRequest) SetQueryParameters(v map[string]string)`

SetQueryParameters sets QueryParameters field to given value.


### GetIsEnabled

`func (o *ApiV2WebhooksPostRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ApiV2WebhooksPostRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ApiV2WebhooksPostRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetShouldSendCustomBody

`func (o *ApiV2WebhooksPostRequest) GetShouldSendCustomBody() bool`

GetShouldSendCustomBody returns the ShouldSendCustomBody field if non-nil, zero value otherwise.

### GetShouldSendCustomBodyOk

`func (o *ApiV2WebhooksPostRequest) GetShouldSendCustomBodyOk() (*bool, bool)`

GetShouldSendCustomBodyOk returns a tuple with the ShouldSendCustomBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldSendCustomBody

`func (o *ApiV2WebhooksPostRequest) SetShouldSendCustomBody(v bool)`

SetShouldSendCustomBody sets ShouldSendCustomBody field to given value.


### GetCustomBody

`func (o *ApiV2WebhooksPostRequest) GetCustomBody() string`

GetCustomBody returns the CustomBody field if non-nil, zero value otherwise.

### GetCustomBodyOk

`func (o *ApiV2WebhooksPostRequest) GetCustomBodyOk() (*string, bool)`

GetCustomBodyOk returns a tuple with the CustomBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBody

`func (o *ApiV2WebhooksPostRequest) SetCustomBody(v string)`

SetCustomBody sets CustomBody field to given value.

### HasCustomBody

`func (o *ApiV2WebhooksPostRequest) HasCustomBody() bool`

HasCustomBody returns a boolean if a field has been set.

### SetCustomBodyNil

`func (o *ApiV2WebhooksPostRequest) SetCustomBodyNil(b bool)`

 SetCustomBodyNil sets the value for CustomBody to be an explicit nil

### UnsetCustomBody
`func (o *ApiV2WebhooksPostRequest) UnsetCustomBody()`

UnsetCustomBody ensures that no value is present for CustomBody, not even an explicit nil
### GetShouldReplaceParameters

`func (o *ApiV2WebhooksPostRequest) GetShouldReplaceParameters() bool`

GetShouldReplaceParameters returns the ShouldReplaceParameters field if non-nil, zero value otherwise.

### GetShouldReplaceParametersOk

`func (o *ApiV2WebhooksPostRequest) GetShouldReplaceParametersOk() (*bool, bool)`

GetShouldReplaceParametersOk returns a tuple with the ShouldReplaceParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldReplaceParameters

`func (o *ApiV2WebhooksPostRequest) SetShouldReplaceParameters(v bool)`

SetShouldReplaceParameters sets ShouldReplaceParameters field to given value.


### GetShouldEscapeParameters

`func (o *ApiV2WebhooksPostRequest) GetShouldEscapeParameters() bool`

GetShouldEscapeParameters returns the ShouldEscapeParameters field if non-nil, zero value otherwise.

### GetShouldEscapeParametersOk

`func (o *ApiV2WebhooksPostRequest) GetShouldEscapeParametersOk() (*bool, bool)`

GetShouldEscapeParametersOk returns a tuple with the ShouldEscapeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldEscapeParameters

`func (o *ApiV2WebhooksPostRequest) SetShouldEscapeParameters(v bool)`

SetShouldEscapeParameters sets ShouldEscapeParameters field to given value.


### GetName

`func (o *ApiV2WebhooksPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2WebhooksPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2WebhooksPostRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


