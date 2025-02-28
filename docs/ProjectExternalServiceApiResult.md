# ProjectExternalServiceApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the external service | 
**Name** | **string** | The name of the external service | 
**Metadata** | [**ExternalServiceMetadataApiResult**](ExternalServiceMetadataApiResult.md) | The metadata associated with the external service | 
**Enabled** | **bool** | Indicates whether the external service is enabled or not | 

## Methods

### NewProjectExternalServiceApiResult

`func NewProjectExternalServiceApiResult(id string, name string, metadata ExternalServiceMetadataApiResult, enabled bool, ) *ProjectExternalServiceApiResult`

NewProjectExternalServiceApiResult instantiates a new ProjectExternalServiceApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectExternalServiceApiResultWithDefaults

`func NewProjectExternalServiceApiResultWithDefaults() *ProjectExternalServiceApiResult`

NewProjectExternalServiceApiResultWithDefaults instantiates a new ProjectExternalServiceApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectExternalServiceApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectExternalServiceApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectExternalServiceApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProjectExternalServiceApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectExternalServiceApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectExternalServiceApiResult) SetName(v string)`

SetName sets Name field to given value.


### GetMetadata

`func (o *ProjectExternalServiceApiResult) GetMetadata() ExternalServiceMetadataApiResult`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProjectExternalServiceApiResult) GetMetadataOk() (*ExternalServiceMetadataApiResult, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProjectExternalServiceApiResult) SetMetadata(v ExternalServiceMetadataApiResult)`

SetMetadata sets Metadata field to given value.


### GetEnabled

`func (o *ProjectExternalServiceApiResult) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProjectExternalServiceApiResult) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProjectExternalServiceApiResult) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


