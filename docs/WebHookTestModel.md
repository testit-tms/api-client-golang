# WebHookTestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestType** | [**RequestTypeModel**](RequestTypeModel.md) |  | 
**Url** | **string** | Request URL of the webhook | 

## Methods

### NewWebHookTestModel

`func NewWebHookTestModel(requestType RequestTypeModel, url string, ) *WebHookTestModel`

NewWebHookTestModel instantiates a new WebHookTestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebHookTestModelWithDefaults

`func NewWebHookTestModelWithDefaults() *WebHookTestModel`

NewWebHookTestModelWithDefaults instantiates a new WebHookTestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestType

`func (o *WebHookTestModel) GetRequestType() RequestTypeModel`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *WebHookTestModel) GetRequestTypeOk() (*RequestTypeModel, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *WebHookTestModel) SetRequestType(v RequestTypeModel)`

SetRequestType sets RequestType field to given value.


### GetUrl

`func (o *WebHookTestModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebHookTestModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebHookTestModel) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


