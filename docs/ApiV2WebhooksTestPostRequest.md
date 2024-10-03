# ApiV2WebhooksTestPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestType** | [**RequestTypeModel**](RequestTypeModel.md) |  | 
**Url** | **string** | Request URL of the webhook | 

## Methods

### NewApiV2WebhooksTestPostRequest

`func NewApiV2WebhooksTestPostRequest(requestType RequestTypeModel, url string, ) *ApiV2WebhooksTestPostRequest`

NewApiV2WebhooksTestPostRequest instantiates a new ApiV2WebhooksTestPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2WebhooksTestPostRequestWithDefaults

`func NewApiV2WebhooksTestPostRequestWithDefaults() *ApiV2WebhooksTestPostRequest`

NewApiV2WebhooksTestPostRequestWithDefaults instantiates a new ApiV2WebhooksTestPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestType

`func (o *ApiV2WebhooksTestPostRequest) GetRequestType() RequestTypeModel`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *ApiV2WebhooksTestPostRequest) GetRequestTypeOk() (*RequestTypeModel, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *ApiV2WebhooksTestPostRequest) SetRequestType(v RequestTypeModel)`

SetRequestType sets RequestType field to given value.


### GetUrl

`func (o *ApiV2WebhooksTestPostRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApiV2WebhooksTestPostRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApiV2WebhooksTestPostRequest) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


