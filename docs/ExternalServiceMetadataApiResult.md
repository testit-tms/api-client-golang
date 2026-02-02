# ExternalServiceMetadataApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the external service | 
**Code** | **string** | The code of the external service | 
**IconUrl** | **string** | The icon URL of the external service | 
**Category** | [**ApiExternalServiceCategory**](ApiExternalServiceCategory.md) | The category of the external service | 

## Methods

### NewExternalServiceMetadataApiResult

`func NewExternalServiceMetadataApiResult(name string, code string, iconUrl string, category ApiExternalServiceCategory, ) *ExternalServiceMetadataApiResult`

NewExternalServiceMetadataApiResult instantiates a new ExternalServiceMetadataApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalServiceMetadataApiResultWithDefaults

`func NewExternalServiceMetadataApiResultWithDefaults() *ExternalServiceMetadataApiResult`

NewExternalServiceMetadataApiResultWithDefaults instantiates a new ExternalServiceMetadataApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExternalServiceMetadataApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalServiceMetadataApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalServiceMetadataApiResult) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *ExternalServiceMetadataApiResult) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ExternalServiceMetadataApiResult) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ExternalServiceMetadataApiResult) SetCode(v string)`

SetCode sets Code field to given value.


### GetIconUrl

`func (o *ExternalServiceMetadataApiResult) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *ExternalServiceMetadataApiResult) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *ExternalServiceMetadataApiResult) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.


### GetCategory

`func (o *ExternalServiceMetadataApiResult) GetCategory() ApiExternalServiceCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ExternalServiceMetadataApiResult) GetCategoryOk() (*ApiExternalServiceCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ExternalServiceMetadataApiResult) SetCategory(v ApiExternalServiceCategory)`

SetCategory sets Category field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


