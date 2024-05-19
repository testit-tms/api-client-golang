# WebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | Pointer to **NullableString** |  | [optional] 
**StatusCode** | **int32** |  | 
**RequestBody** | Pointer to **NullableString** |  | [optional] 
**RequestMeta** | **string** |  | 
**ResponseBody** | **string** |  | 
**ResponseMeta** | **string** |  | 

## Methods

### NewWebhookResponse

`func NewWebhookResponse(statusCode int32, requestMeta string, responseBody string, responseMeta string, ) *WebhookResponse`

NewWebhookResponse instantiates a new WebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookResponseWithDefaults

`func NewWebhookResponseWithDefaults() *WebhookResponse`

NewWebhookResponseWithDefaults instantiates a new WebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *WebhookResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *WebhookResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *WebhookResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *WebhookResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *WebhookResponse) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *WebhookResponse) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetStatusCode

`func (o *WebhookResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *WebhookResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *WebhookResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetRequestBody

`func (o *WebhookResponse) GetRequestBody() string`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *WebhookResponse) GetRequestBodyOk() (*string, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *WebhookResponse) SetRequestBody(v string)`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *WebhookResponse) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### SetRequestBodyNil

`func (o *WebhookResponse) SetRequestBodyNil(b bool)`

 SetRequestBodyNil sets the value for RequestBody to be an explicit nil

### UnsetRequestBody
`func (o *WebhookResponse) UnsetRequestBody()`

UnsetRequestBody ensures that no value is present for RequestBody, not even an explicit nil
### GetRequestMeta

`func (o *WebhookResponse) GetRequestMeta() string`

GetRequestMeta returns the RequestMeta field if non-nil, zero value otherwise.

### GetRequestMetaOk

`func (o *WebhookResponse) GetRequestMetaOk() (*string, bool)`

GetRequestMetaOk returns a tuple with the RequestMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMeta

`func (o *WebhookResponse) SetRequestMeta(v string)`

SetRequestMeta sets RequestMeta field to given value.


### GetResponseBody

`func (o *WebhookResponse) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *WebhookResponse) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *WebhookResponse) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.


### GetResponseMeta

`func (o *WebhookResponse) GetResponseMeta() string`

GetResponseMeta returns the ResponseMeta field if non-nil, zero value otherwise.

### GetResponseMetaOk

`func (o *WebhookResponse) GetResponseMetaOk() (*string, bool)`

GetResponseMetaOk returns a tuple with the ResponseMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMeta

`func (o *WebhookResponse) SetResponseMeta(v string)`

SetResponseMeta sets ResponseMeta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


