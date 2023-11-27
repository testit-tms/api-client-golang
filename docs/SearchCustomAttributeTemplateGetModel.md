# SearchCustomAttributeTemplateGetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**IsDeleted** | **bool** |  | 
**Name** | **string** |  | 
**ProjectShortestModels** | [**[]ProjectShortestModel**](ProjectShortestModel.md) |  | 
**CustomAttributeModels** | [**[]CustomAttributeModel**](CustomAttributeModel.md) |  | 

## Methods

### NewSearchCustomAttributeTemplateGetModel

`func NewSearchCustomAttributeTemplateGetModel(id string, isDeleted bool, name string, projectShortestModels []ProjectShortestModel, customAttributeModels []CustomAttributeModel, ) *SearchCustomAttributeTemplateGetModel`

NewSearchCustomAttributeTemplateGetModel instantiates a new SearchCustomAttributeTemplateGetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchCustomAttributeTemplateGetModelWithDefaults

`func NewSearchCustomAttributeTemplateGetModelWithDefaults() *SearchCustomAttributeTemplateGetModel`

NewSearchCustomAttributeTemplateGetModelWithDefaults instantiates a new SearchCustomAttributeTemplateGetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchCustomAttributeTemplateGetModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchCustomAttributeTemplateGetModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchCustomAttributeTemplateGetModel) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleted

`func (o *SearchCustomAttributeTemplateGetModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *SearchCustomAttributeTemplateGetModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *SearchCustomAttributeTemplateGetModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetName

`func (o *SearchCustomAttributeTemplateGetModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchCustomAttributeTemplateGetModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchCustomAttributeTemplateGetModel) SetName(v string)`

SetName sets Name field to given value.


### GetProjectShortestModels

`func (o *SearchCustomAttributeTemplateGetModel) GetProjectShortestModels() []ProjectShortestModel`

GetProjectShortestModels returns the ProjectShortestModels field if non-nil, zero value otherwise.

### GetProjectShortestModelsOk

`func (o *SearchCustomAttributeTemplateGetModel) GetProjectShortestModelsOk() (*[]ProjectShortestModel, bool)`

GetProjectShortestModelsOk returns a tuple with the ProjectShortestModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectShortestModels

`func (o *SearchCustomAttributeTemplateGetModel) SetProjectShortestModels(v []ProjectShortestModel)`

SetProjectShortestModels sets ProjectShortestModels field to given value.


### GetCustomAttributeModels

`func (o *SearchCustomAttributeTemplateGetModel) GetCustomAttributeModels() []CustomAttributeModel`

GetCustomAttributeModels returns the CustomAttributeModels field if non-nil, zero value otherwise.

### GetCustomAttributeModelsOk

`func (o *SearchCustomAttributeTemplateGetModel) GetCustomAttributeModelsOk() (*[]CustomAttributeModel, bool)`

GetCustomAttributeModelsOk returns a tuple with the CustomAttributeModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributeModels

`func (o *SearchCustomAttributeTemplateGetModel) SetCustomAttributeModels(v []CustomAttributeModel)`

SetCustomAttributeModels sets CustomAttributeModels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


