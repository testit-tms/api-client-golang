# WebhookLogApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**IsDeleted** | **bool** |  | 
**WebHookName** | **string** |  | 
**EventType** | [**WebHookEventType**](WebHookEventType.md) |  | 
**WebHookId** | **string** |  | 
**RequestBody** | Pointer to **NullableString** |  | [optional] 
**RequestMeta** | Pointer to **NullableString** |  | [optional] 
**ResponseStatusCode** | **int32** |  | 
**ResponseBody** | Pointer to **NullableString** |  | [optional] 
**ResponseMeta** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | **string** |  | 
**Url** | **string** |  | 
**RequestType** | [**RequestType**](RequestType.md) |  | 
**CreatedDate** | Pointer to **NullableTime** |  | [optional] 
**ModifiedDate** | Pointer to **NullableTime** |  | [optional] 
**CreatedById** | **string** |  | 
**ModifiedById** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWebhookLogApiResult

`func NewWebhookLogApiResult(id string, isDeleted bool, webHookName string, eventType WebHookEventType, webHookId string, responseStatusCode int32, projectId string, url string, requestType RequestType, createdById string, ) *WebhookLogApiResult`

NewWebhookLogApiResult instantiates a new WebhookLogApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookLogApiResultWithDefaults

`func NewWebhookLogApiResultWithDefaults() *WebhookLogApiResult`

NewWebhookLogApiResultWithDefaults instantiates a new WebhookLogApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookLogApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookLogApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookLogApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleted

`func (o *WebhookLogApiResult) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *WebhookLogApiResult) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *WebhookLogApiResult) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetWebHookName

`func (o *WebhookLogApiResult) GetWebHookName() string`

GetWebHookName returns the WebHookName field if non-nil, zero value otherwise.

### GetWebHookNameOk

`func (o *WebhookLogApiResult) GetWebHookNameOk() (*string, bool)`

GetWebHookNameOk returns a tuple with the WebHookName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHookName

`func (o *WebhookLogApiResult) SetWebHookName(v string)`

SetWebHookName sets WebHookName field to given value.


### GetEventType

`func (o *WebhookLogApiResult) GetEventType() WebHookEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WebhookLogApiResult) GetEventTypeOk() (*WebHookEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WebhookLogApiResult) SetEventType(v WebHookEventType)`

SetEventType sets EventType field to given value.


### GetWebHookId

`func (o *WebhookLogApiResult) GetWebHookId() string`

GetWebHookId returns the WebHookId field if non-nil, zero value otherwise.

### GetWebHookIdOk

`func (o *WebhookLogApiResult) GetWebHookIdOk() (*string, bool)`

GetWebHookIdOk returns a tuple with the WebHookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHookId

`func (o *WebhookLogApiResult) SetWebHookId(v string)`

SetWebHookId sets WebHookId field to given value.


### GetRequestBody

`func (o *WebhookLogApiResult) GetRequestBody() string`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *WebhookLogApiResult) GetRequestBodyOk() (*string, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *WebhookLogApiResult) SetRequestBody(v string)`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *WebhookLogApiResult) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### SetRequestBodyNil

`func (o *WebhookLogApiResult) SetRequestBodyNil(b bool)`

 SetRequestBodyNil sets the value for RequestBody to be an explicit nil

### UnsetRequestBody
`func (o *WebhookLogApiResult) UnsetRequestBody()`

UnsetRequestBody ensures that no value is present for RequestBody, not even an explicit nil
### GetRequestMeta

`func (o *WebhookLogApiResult) GetRequestMeta() string`

GetRequestMeta returns the RequestMeta field if non-nil, zero value otherwise.

### GetRequestMetaOk

`func (o *WebhookLogApiResult) GetRequestMetaOk() (*string, bool)`

GetRequestMetaOk returns a tuple with the RequestMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMeta

`func (o *WebhookLogApiResult) SetRequestMeta(v string)`

SetRequestMeta sets RequestMeta field to given value.

### HasRequestMeta

`func (o *WebhookLogApiResult) HasRequestMeta() bool`

HasRequestMeta returns a boolean if a field has been set.

### SetRequestMetaNil

`func (o *WebhookLogApiResult) SetRequestMetaNil(b bool)`

 SetRequestMetaNil sets the value for RequestMeta to be an explicit nil

### UnsetRequestMeta
`func (o *WebhookLogApiResult) UnsetRequestMeta()`

UnsetRequestMeta ensures that no value is present for RequestMeta, not even an explicit nil
### GetResponseStatusCode

`func (o *WebhookLogApiResult) GetResponseStatusCode() int32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *WebhookLogApiResult) GetResponseStatusCodeOk() (*int32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *WebhookLogApiResult) SetResponseStatusCode(v int32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.


### GetResponseBody

`func (o *WebhookLogApiResult) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *WebhookLogApiResult) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *WebhookLogApiResult) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *WebhookLogApiResult) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### SetResponseBodyNil

`func (o *WebhookLogApiResult) SetResponseBodyNil(b bool)`

 SetResponseBodyNil sets the value for ResponseBody to be an explicit nil

### UnsetResponseBody
`func (o *WebhookLogApiResult) UnsetResponseBody()`

UnsetResponseBody ensures that no value is present for ResponseBody, not even an explicit nil
### GetResponseMeta

`func (o *WebhookLogApiResult) GetResponseMeta() string`

GetResponseMeta returns the ResponseMeta field if non-nil, zero value otherwise.

### GetResponseMetaOk

`func (o *WebhookLogApiResult) GetResponseMetaOk() (*string, bool)`

GetResponseMetaOk returns a tuple with the ResponseMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMeta

`func (o *WebhookLogApiResult) SetResponseMeta(v string)`

SetResponseMeta sets ResponseMeta field to given value.

### HasResponseMeta

`func (o *WebhookLogApiResult) HasResponseMeta() bool`

HasResponseMeta returns a boolean if a field has been set.

### SetResponseMetaNil

`func (o *WebhookLogApiResult) SetResponseMetaNil(b bool)`

 SetResponseMetaNil sets the value for ResponseMeta to be an explicit nil

### UnsetResponseMeta
`func (o *WebhookLogApiResult) UnsetResponseMeta()`

UnsetResponseMeta ensures that no value is present for ResponseMeta, not even an explicit nil
### GetProjectId

`func (o *WebhookLogApiResult) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WebhookLogApiResult) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WebhookLogApiResult) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetUrl

`func (o *WebhookLogApiResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookLogApiResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookLogApiResult) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRequestType

`func (o *WebhookLogApiResult) GetRequestType() RequestType`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *WebhookLogApiResult) GetRequestTypeOk() (*RequestType, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *WebhookLogApiResult) SetRequestType(v RequestType)`

SetRequestType sets RequestType field to given value.


### GetCreatedDate

`func (o *WebhookLogApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *WebhookLogApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *WebhookLogApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *WebhookLogApiResult) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### SetCreatedDateNil

`func (o *WebhookLogApiResult) SetCreatedDateNil(b bool)`

 SetCreatedDateNil sets the value for CreatedDate to be an explicit nil

### UnsetCreatedDate
`func (o *WebhookLogApiResult) UnsetCreatedDate()`

UnsetCreatedDate ensures that no value is present for CreatedDate, not even an explicit nil
### GetModifiedDate

`func (o *WebhookLogApiResult) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *WebhookLogApiResult) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *WebhookLogApiResult) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *WebhookLogApiResult) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### SetModifiedDateNil

`func (o *WebhookLogApiResult) SetModifiedDateNil(b bool)`

 SetModifiedDateNil sets the value for ModifiedDate to be an explicit nil

### UnsetModifiedDate
`func (o *WebhookLogApiResult) UnsetModifiedDate()`

UnsetModifiedDate ensures that no value is present for ModifiedDate, not even an explicit nil
### GetCreatedById

`func (o *WebhookLogApiResult) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WebhookLogApiResult) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WebhookLogApiResult) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedById

`func (o *WebhookLogApiResult) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *WebhookLogApiResult) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *WebhookLogApiResult) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *WebhookLogApiResult) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### SetModifiedByIdNil

`func (o *WebhookLogApiResult) SetModifiedByIdNil(b bool)`

 SetModifiedByIdNil sets the value for ModifiedById to be an explicit nil

### UnsetModifiedById
`func (o *WebhookLogApiResult) UnsetModifiedById()`

UnsetModifiedById ensures that no value is present for ModifiedById, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


