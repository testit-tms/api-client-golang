# ConfigurationParameterPreviewApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the configuration parameter | 
**Name** | **string** | Name of the configuration parameter | 
**IsDeleted** | **bool** | Is configuration parameter deleted? | 
**CreatedDate** | **time.Time** | Date of configuration parameter creation | 
**CreatedById** | **string** | Identifier of user who created configuration parameter | 
**ModifiedDate** | **time.Time** | Date of configuration parameter modification | 
**ModifiedById** | **string** | Identifier of user who modified configuration parameter | 
**Values** | [**ConfigurationParameterValueApiResultApiCollectionPreview**](ConfigurationParameterValueApiResultApiCollectionPreview.md) | Preview of configuration parameter values | 
**Projects** | [**ProjectNameApiResultApiCollectionPreview**](ProjectNameApiResultApiCollectionPreview.md) | Preview of assigned projects | 

## Methods

### NewConfigurationParameterPreviewApiResult

`func NewConfigurationParameterPreviewApiResult(id string, name string, isDeleted bool, createdDate time.Time, createdById string, modifiedDate time.Time, modifiedById string, values ConfigurationParameterValueApiResultApiCollectionPreview, projects ProjectNameApiResultApiCollectionPreview, ) *ConfigurationParameterPreviewApiResult`

NewConfigurationParameterPreviewApiResult instantiates a new ConfigurationParameterPreviewApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationParameterPreviewApiResultWithDefaults

`func NewConfigurationParameterPreviewApiResultWithDefaults() *ConfigurationParameterPreviewApiResult`

NewConfigurationParameterPreviewApiResultWithDefaults instantiates a new ConfigurationParameterPreviewApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigurationParameterPreviewApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigurationParameterPreviewApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigurationParameterPreviewApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ConfigurationParameterPreviewApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigurationParameterPreviewApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigurationParameterPreviewApiResult) SetName(v string)`

SetName sets Name field to given value.


### GetIsDeleted

`func (o *ConfigurationParameterPreviewApiResult) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ConfigurationParameterPreviewApiResult) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ConfigurationParameterPreviewApiResult) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetCreatedDate

`func (o *ConfigurationParameterPreviewApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ConfigurationParameterPreviewApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ConfigurationParameterPreviewApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetCreatedById

`func (o *ConfigurationParameterPreviewApiResult) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *ConfigurationParameterPreviewApiResult) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *ConfigurationParameterPreviewApiResult) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedDate

`func (o *ConfigurationParameterPreviewApiResult) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ConfigurationParameterPreviewApiResult) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ConfigurationParameterPreviewApiResult) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.


### GetModifiedById

`func (o *ConfigurationParameterPreviewApiResult) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *ConfigurationParameterPreviewApiResult) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *ConfigurationParameterPreviewApiResult) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.


### GetValues

`func (o *ConfigurationParameterPreviewApiResult) GetValues() ConfigurationParameterValueApiResultApiCollectionPreview`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ConfigurationParameterPreviewApiResult) GetValuesOk() (*ConfigurationParameterValueApiResultApiCollectionPreview, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ConfigurationParameterPreviewApiResult) SetValues(v ConfigurationParameterValueApiResultApiCollectionPreview)`

SetValues sets Values field to given value.


### GetProjects

`func (o *ConfigurationParameterPreviewApiResult) GetProjects() ProjectNameApiResultApiCollectionPreview`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ConfigurationParameterPreviewApiResult) GetProjectsOk() (*ProjectNameApiResultApiCollectionPreview, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ConfigurationParameterPreviewApiResult) SetProjects(v ProjectNameApiResultApiCollectionPreview)`

SetProjects sets Projects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


