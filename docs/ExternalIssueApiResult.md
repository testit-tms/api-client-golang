# ExternalIssueApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identiief of external issue | 
**ExternalId** | **string** | Identifier of external issue in external service | 
**Url** | **string** | Url of external issue | 
**Metadata** | [**ExternalIssueApiMetadata**](ExternalIssueApiMetadata.md) | Metadata of external issue from external service | 

## Methods

### NewExternalIssueApiResult

`func NewExternalIssueApiResult(id string, externalId string, url string, metadata ExternalIssueApiMetadata, ) *ExternalIssueApiResult`

NewExternalIssueApiResult instantiates a new ExternalIssueApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalIssueApiResultWithDefaults

`func NewExternalIssueApiResultWithDefaults() *ExternalIssueApiResult`

NewExternalIssueApiResultWithDefaults instantiates a new ExternalIssueApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalIssueApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalIssueApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalIssueApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetExternalId

`func (o *ExternalIssueApiResult) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ExternalIssueApiResult) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ExternalIssueApiResult) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetUrl

`func (o *ExternalIssueApiResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ExternalIssueApiResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ExternalIssueApiResult) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMetadata

`func (o *ExternalIssueApiResult) GetMetadata() ExternalIssueApiMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ExternalIssueApiResult) GetMetadataOk() (*ExternalIssueApiMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ExternalIssueApiResult) SetMetadata(v ExternalIssueApiMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


