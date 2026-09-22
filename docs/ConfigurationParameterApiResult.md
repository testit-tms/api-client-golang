# ConfigurationParameterApiResult

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
**Values** | [**[]ConfigurationParameterValueApiResult**](ConfigurationParameterValueApiResult.md) | List of configuration parameter values | 
**Projects** | [**[]ProjectNameApiResult**](ProjectNameApiResult.md) | List of assigned projects | 

## Methods

### NewConfigurationParameterApiResult

`func NewConfigurationParameterApiResult(id string, name string, isDeleted bool, createdDate time.Time, createdById string, modifiedDate time.Time, modifiedById string, values []ConfigurationParameterValueApiResult, projects []ProjectNameApiResult, ) *ConfigurationParameterApiResult`

NewConfigurationParameterApiResult instantiates a new ConfigurationParameterApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationParameterApiResultWithDefaults

`func NewConfigurationParameterApiResultWithDefaults() *ConfigurationParameterApiResult`

NewConfigurationParameterApiResultWithDefaults instantiates a new ConfigurationParameterApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigurationParameterApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigurationParameterApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigurationParameterApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ConfigurationParameterApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigurationParameterApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigurationParameterApiResult) SetName(v string)`

SetName sets Name field to given value.


### GetIsDeleted

`func (o *ConfigurationParameterApiResult) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ConfigurationParameterApiResult) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ConfigurationParameterApiResult) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetCreatedDate

`func (o *ConfigurationParameterApiResult) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ConfigurationParameterApiResult) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ConfigurationParameterApiResult) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetCreatedById

`func (o *ConfigurationParameterApiResult) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *ConfigurationParameterApiResult) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *ConfigurationParameterApiResult) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetModifiedDate

`func (o *ConfigurationParameterApiResult) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ConfigurationParameterApiResult) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ConfigurationParameterApiResult) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.


### GetModifiedById

`func (o *ConfigurationParameterApiResult) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *ConfigurationParameterApiResult) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *ConfigurationParameterApiResult) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.


### GetValues

`func (o *ConfigurationParameterApiResult) GetValues() []ConfigurationParameterValueApiResult`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ConfigurationParameterApiResult) GetValuesOk() (*[]ConfigurationParameterValueApiResult, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ConfigurationParameterApiResult) SetValues(v []ConfigurationParameterValueApiResult)`

SetValues sets Values field to given value.


### GetProjects

`func (o *ConfigurationParameterApiResult) GetProjects() []ProjectNameApiResult`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ConfigurationParameterApiResult) GetProjectsOk() (*[]ProjectNameApiResult, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ConfigurationParameterApiResult) SetProjects(v []ProjectNameApiResult)`

SetProjects sets Projects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


