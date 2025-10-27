# ExternalIssueModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ExternalId** | **string** |  | 
**Url** | **string** |  | 
**Metadata** | [**ExternalIssueMetadataModel**](ExternalIssueMetadataModel.md) |  | 

## Methods

### NewExternalIssueModel

`func NewExternalIssueModel(id string, externalId string, url string, metadata ExternalIssueMetadataModel, ) *ExternalIssueModel`

NewExternalIssueModel instantiates a new ExternalIssueModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalIssueModelWithDefaults

`func NewExternalIssueModelWithDefaults() *ExternalIssueModel`

NewExternalIssueModelWithDefaults instantiates a new ExternalIssueModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalIssueModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalIssueModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalIssueModel) SetId(v string)`

SetId sets Id field to given value.


### GetExternalId

`func (o *ExternalIssueModel) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ExternalIssueModel) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ExternalIssueModel) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetUrl

`func (o *ExternalIssueModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ExternalIssueModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ExternalIssueModel) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMetadata

`func (o *ExternalIssueModel) GetMetadata() ExternalIssueMetadataModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ExternalIssueModel) GetMetadataOk() (*ExternalIssueMetadataModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ExternalIssueModel) SetMetadata(v ExternalIssueMetadataModel)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


