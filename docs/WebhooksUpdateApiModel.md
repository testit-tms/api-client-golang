# WebhooksUpdateApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**WebhooksFilterApiModel**](WebhooksFilterApiModel.md) |  | 
**Model** | [**WebhookBulkUpdateApiModel**](WebhookBulkUpdateApiModel.md) |  | 
**Extractor** | [**WebhooksExtractionApiModel**](WebhooksExtractionApiModel.md) |  | 

## Methods

### NewWebhooksUpdateApiModel

`func NewWebhooksUpdateApiModel(filter WebhooksFilterApiModel, model WebhookBulkUpdateApiModel, extractor WebhooksExtractionApiModel, ) *WebhooksUpdateApiModel`

NewWebhooksUpdateApiModel instantiates a new WebhooksUpdateApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhooksUpdateApiModelWithDefaults

`func NewWebhooksUpdateApiModelWithDefaults() *WebhooksUpdateApiModel`

NewWebhooksUpdateApiModelWithDefaults instantiates a new WebhooksUpdateApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *WebhooksUpdateApiModel) GetFilter() WebhooksFilterApiModel`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *WebhooksUpdateApiModel) GetFilterOk() (*WebhooksFilterApiModel, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *WebhooksUpdateApiModel) SetFilter(v WebhooksFilterApiModel)`

SetFilter sets Filter field to given value.


### GetModel

`func (o *WebhooksUpdateApiModel) GetModel() WebhookBulkUpdateApiModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *WebhooksUpdateApiModel) GetModelOk() (*WebhookBulkUpdateApiModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *WebhooksUpdateApiModel) SetModel(v WebhookBulkUpdateApiModel)`

SetModel sets Model field to given value.


### GetExtractor

`func (o *WebhooksUpdateApiModel) GetExtractor() WebhooksExtractionApiModel`

GetExtractor returns the Extractor field if non-nil, zero value otherwise.

### GetExtractorOk

`func (o *WebhooksUpdateApiModel) GetExtractorOk() (*WebhooksExtractionApiModel, bool)`

GetExtractorOk returns a tuple with the Extractor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractor

`func (o *WebhooksUpdateApiModel) SetExtractor(v WebhooksExtractionApiModel)`

SetExtractor sets Extractor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


