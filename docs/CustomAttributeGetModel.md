# CustomAttributeGetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the attribute | 
**Options** | [**[]CustomAttributeOptionModel**](CustomAttributeOptionModel.md) | Collection of the attribute options | 
**Type** | [**CustomAttributeTypesEnum**](CustomAttributeTypesEnum.md) | Type of the attribute | 
**IsDeleted** | **bool** | Indicates if the attribute is deleted | 
**Name** | **string** | Name of the attribute | 
**IsEnabled** | **bool** | Indicates if the attribute is enabled | 
**IsRequired** | **bool** | Indicates if the attribute is mandatory to specify | 
**IsGlobal** | **bool** | Indicates if the attribute is available across all projects | 

## Methods

### NewCustomAttributeGetModel

`func NewCustomAttributeGetModel(id string, options []CustomAttributeOptionModel, type_ CustomAttributeTypesEnum, isDeleted bool, name string, isEnabled bool, isRequired bool, isGlobal bool, ) *CustomAttributeGetModel`

NewCustomAttributeGetModel instantiates a new CustomAttributeGetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAttributeGetModelWithDefaults

`func NewCustomAttributeGetModelWithDefaults() *CustomAttributeGetModel`

NewCustomAttributeGetModelWithDefaults instantiates a new CustomAttributeGetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomAttributeGetModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomAttributeGetModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomAttributeGetModel) SetId(v string)`

SetId sets Id field to given value.


### GetOptions

`func (o *CustomAttributeGetModel) GetOptions() []CustomAttributeOptionModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CustomAttributeGetModel) GetOptionsOk() (*[]CustomAttributeOptionModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CustomAttributeGetModel) SetOptions(v []CustomAttributeOptionModel)`

SetOptions sets Options field to given value.


### GetType

`func (o *CustomAttributeGetModel) GetType() CustomAttributeTypesEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomAttributeGetModel) GetTypeOk() (*CustomAttributeTypesEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomAttributeGetModel) SetType(v CustomAttributeTypesEnum)`

SetType sets Type field to given value.


### GetIsDeleted

`func (o *CustomAttributeGetModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CustomAttributeGetModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CustomAttributeGetModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetName

`func (o *CustomAttributeGetModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomAttributeGetModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomAttributeGetModel) SetName(v string)`

SetName sets Name field to given value.


### GetIsEnabled

`func (o *CustomAttributeGetModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CustomAttributeGetModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CustomAttributeGetModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetIsRequired

`func (o *CustomAttributeGetModel) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *CustomAttributeGetModel) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *CustomAttributeGetModel) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.


### GetIsGlobal

`func (o *CustomAttributeGetModel) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *CustomAttributeGetModel) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *CustomAttributeGetModel) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


