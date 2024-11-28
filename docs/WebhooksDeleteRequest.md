# WebhooksDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**WebhooksDeleteFilterRequest**](WebhooksDeleteFilterRequest.md) |  | 
**Extractor** | [**WebhooksExtractionRequest**](WebhooksExtractionRequest.md) |  | 

## Methods

### NewWebhooksDeleteRequest

`func NewWebhooksDeleteRequest(filter WebhooksDeleteFilterRequest, extractor WebhooksExtractionRequest, ) *WebhooksDeleteRequest`

NewWebhooksDeleteRequest instantiates a new WebhooksDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhooksDeleteRequestWithDefaults

`func NewWebhooksDeleteRequestWithDefaults() *WebhooksDeleteRequest`

NewWebhooksDeleteRequestWithDefaults instantiates a new WebhooksDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *WebhooksDeleteRequest) GetFilter() WebhooksDeleteFilterRequest`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *WebhooksDeleteRequest) GetFilterOk() (*WebhooksDeleteFilterRequest, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *WebhooksDeleteRequest) SetFilter(v WebhooksDeleteFilterRequest)`

SetFilter sets Filter field to given value.


### GetExtractor

`func (o *WebhooksDeleteRequest) GetExtractor() WebhooksExtractionRequest`

GetExtractor returns the Extractor field if non-nil, zero value otherwise.

### GetExtractorOk

`func (o *WebhooksDeleteRequest) GetExtractorOk() (*WebhooksExtractionRequest, bool)`

GetExtractorOk returns a tuple with the Extractor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractor

`func (o *WebhooksDeleteRequest) SetExtractor(v WebhooksExtractionRequest)`

SetExtractor sets Extractor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


