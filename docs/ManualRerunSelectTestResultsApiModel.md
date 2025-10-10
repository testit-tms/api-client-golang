# ManualRerunSelectTestResultsApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**NullableTestResultsFilterApiModel**](TestResultsFilterApiModel.md) | Test results filter. | [optional] 
**ExtractionModel** | Pointer to [**NullableManualRerunTestResultApiModel**](ManualRerunTestResultApiModel.md) | Test results extraction model. | [optional] 
**WebhookIds** | Pointer to **[]string** | Webhook ids to rerun. | [optional] 

## Methods

### NewManualRerunSelectTestResultsApiModel

`func NewManualRerunSelectTestResultsApiModel() *ManualRerunSelectTestResultsApiModel`

NewManualRerunSelectTestResultsApiModel instantiates a new ManualRerunSelectTestResultsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualRerunSelectTestResultsApiModelWithDefaults

`func NewManualRerunSelectTestResultsApiModelWithDefaults() *ManualRerunSelectTestResultsApiModel`

NewManualRerunSelectTestResultsApiModelWithDefaults instantiates a new ManualRerunSelectTestResultsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ManualRerunSelectTestResultsApiModel) GetFilter() TestResultsFilterApiModel`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ManualRerunSelectTestResultsApiModel) GetFilterOk() (*TestResultsFilterApiModel, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ManualRerunSelectTestResultsApiModel) SetFilter(v TestResultsFilterApiModel)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ManualRerunSelectTestResultsApiModel) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *ManualRerunSelectTestResultsApiModel) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *ManualRerunSelectTestResultsApiModel) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetExtractionModel

`func (o *ManualRerunSelectTestResultsApiModel) GetExtractionModel() ManualRerunTestResultApiModel`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *ManualRerunSelectTestResultsApiModel) GetExtractionModelOk() (*ManualRerunTestResultApiModel, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *ManualRerunSelectTestResultsApiModel) SetExtractionModel(v ManualRerunTestResultApiModel)`

SetExtractionModel sets ExtractionModel field to given value.

### HasExtractionModel

`func (o *ManualRerunSelectTestResultsApiModel) HasExtractionModel() bool`

HasExtractionModel returns a boolean if a field has been set.

### SetExtractionModelNil

`func (o *ManualRerunSelectTestResultsApiModel) SetExtractionModelNil(b bool)`

 SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil

### UnsetExtractionModel
`func (o *ManualRerunSelectTestResultsApiModel) UnsetExtractionModel()`

UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil
### GetWebhookIds

`func (o *ManualRerunSelectTestResultsApiModel) GetWebhookIds() []string`

GetWebhookIds returns the WebhookIds field if non-nil, zero value otherwise.

### GetWebhookIdsOk

`func (o *ManualRerunSelectTestResultsApiModel) GetWebhookIdsOk() (*[]string, bool)`

GetWebhookIdsOk returns a tuple with the WebhookIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookIds

`func (o *ManualRerunSelectTestResultsApiModel) SetWebhookIds(v []string)`

SetWebhookIds sets WebhookIds field to given value.

### HasWebhookIds

`func (o *ManualRerunSelectTestResultsApiModel) HasWebhookIds() bool`

HasWebhookIds returns a boolean if a field has been set.

### SetWebhookIdsNil

`func (o *ManualRerunSelectTestResultsApiModel) SetWebhookIdsNil(b bool)`

 SetWebhookIdsNil sets the value for WebhookIds to be an explicit nil

### UnsetWebhookIds
`func (o *ManualRerunSelectTestResultsApiModel) UnsetWebhookIds()`

UnsetWebhookIds ensures that no value is present for WebhookIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


