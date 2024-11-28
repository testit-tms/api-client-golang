# WebhooksUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**WebhooksFilterRequest**](WebhooksFilterRequest.md) |  | 
**Model** | [**WebhookBulkUpdateApiModel**](WebhookBulkUpdateApiModel.md) |  | 
**Extractor** | [**WebhooksExtractionRequest**](WebhooksExtractionRequest.md) |  | 

## Methods

### NewWebhooksUpdateRequest

`func NewWebhooksUpdateRequest(filter WebhooksFilterRequest, model WebhookBulkUpdateApiModel, extractor WebhooksExtractionRequest, ) *WebhooksUpdateRequest`

NewWebhooksUpdateRequest instantiates a new WebhooksUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhooksUpdateRequestWithDefaults

`func NewWebhooksUpdateRequestWithDefaults() *WebhooksUpdateRequest`

NewWebhooksUpdateRequestWithDefaults instantiates a new WebhooksUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *WebhooksUpdateRequest) GetFilter() WebhooksFilterRequest`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *WebhooksUpdateRequest) GetFilterOk() (*WebhooksFilterRequest, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *WebhooksUpdateRequest) SetFilter(v WebhooksFilterRequest)`

SetFilter sets Filter field to given value.


### GetModel

`func (o *WebhooksUpdateRequest) GetModel() WebhookBulkUpdateApiModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *WebhooksUpdateRequest) GetModelOk() (*WebhookBulkUpdateApiModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *WebhooksUpdateRequest) SetModel(v WebhookBulkUpdateApiModel)`

SetModel sets Model field to given value.


### GetExtractor

`func (o *WebhooksUpdateRequest) GetExtractor() WebhooksExtractionRequest`

GetExtractor returns the Extractor field if non-nil, zero value otherwise.

### GetExtractorOk

`func (o *WebhooksUpdateRequest) GetExtractorOk() (*WebhooksExtractionRequest, bool)`

GetExtractorOk returns a tuple with the Extractor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractor

`func (o *WebhooksUpdateRequest) SetExtractor(v WebhooksExtractionRequest)`

SetExtractor sets Extractor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


