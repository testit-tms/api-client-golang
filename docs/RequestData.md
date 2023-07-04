# RequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | Pointer to **NullableString** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**RequestBody** | Pointer to **NullableString** |  | [optional] 
**RequestMeta** | Pointer to **string** |  | [optional] 
**ResponseBody** | Pointer to **string** |  | [optional] 
**ResponseMeta** | Pointer to **string** |  | [optional] 

## Methods

### NewRequestData

`func NewRequestData() *RequestData`

NewRequestData instantiates a new RequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestDataWithDefaults

`func NewRequestDataWithDefaults() *RequestData`

NewRequestDataWithDefaults instantiates a new RequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *RequestData) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *RequestData) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *RequestData) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *RequestData) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *RequestData) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *RequestData) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetStatusCode

`func (o *RequestData) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *RequestData) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *RequestData) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *RequestData) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetRequestBody

`func (o *RequestData) GetRequestBody() string`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *RequestData) GetRequestBodyOk() (*string, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *RequestData) SetRequestBody(v string)`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *RequestData) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### SetRequestBodyNil

`func (o *RequestData) SetRequestBodyNil(b bool)`

 SetRequestBodyNil sets the value for RequestBody to be an explicit nil

### UnsetRequestBody
`func (o *RequestData) UnsetRequestBody()`

UnsetRequestBody ensures that no value is present for RequestBody, not even an explicit nil
### GetRequestMeta

`func (o *RequestData) GetRequestMeta() string`

GetRequestMeta returns the RequestMeta field if non-nil, zero value otherwise.

### GetRequestMetaOk

`func (o *RequestData) GetRequestMetaOk() (*string, bool)`

GetRequestMetaOk returns a tuple with the RequestMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMeta

`func (o *RequestData) SetRequestMeta(v string)`

SetRequestMeta sets RequestMeta field to given value.

### HasRequestMeta

`func (o *RequestData) HasRequestMeta() bool`

HasRequestMeta returns a boolean if a field has been set.

### GetResponseBody

`func (o *RequestData) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *RequestData) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *RequestData) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *RequestData) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetResponseMeta

`func (o *RequestData) GetResponseMeta() string`

GetResponseMeta returns the ResponseMeta field if non-nil, zero value otherwise.

### GetResponseMetaOk

`func (o *RequestData) GetResponseMetaOk() (*string, bool)`

GetResponseMetaOk returns a tuple with the ResponseMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMeta

`func (o *RequestData) SetResponseMeta(v string)`

SetResponseMeta sets ResponseMeta field to given value.

### HasResponseMeta

`func (o *RequestData) HasResponseMeta() bool`

HasResponseMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


