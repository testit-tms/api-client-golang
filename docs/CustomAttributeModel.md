# CustomAttributeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the attribute. | 
**Code** | Pointer to **NullableString** | Optional code identifier for the attribute. | [optional] 
**Type** | [**CustomAttributeTypesEnum**](CustomAttributeTypesEnum.md) | Type of the attribute. | 
**Options** | [**[]CustomAttributeOptionModel**](CustomAttributeOptionModel.md) | Collection of the attribute options. | 
**Targets** | **[]string** | Collection of the attribute targets.   Defines where the attribute can be used (e.g., TestCases, AutoTestCases, TestPlans). | 
**IsReadOnly** | **bool** | Indicates if the attribute is read-only. | 
**IsDeleted** | **bool** | Indicates if the attribute is deleted. | 
**IsSystem** | **bool** | Indicates if the attribute is system. | 
**Name** | **string** | Name of the attribute | 
**IsEnabled** | **bool** | Indicates if the attribute is enabled | 
**IsRequired** | **bool** | Indicates if the attribute value is mandatory to specify | 
**IsGlobal** | **bool** | Indicates if the attribute is available across all projects | 

## Methods

### NewCustomAttributeModel

`func NewCustomAttributeModel(id string, type_ CustomAttributeTypesEnum, options []CustomAttributeOptionModel, targets []string, isReadOnly bool, isDeleted bool, isSystem bool, name string, isEnabled bool, isRequired bool, isGlobal bool, ) *CustomAttributeModel`

NewCustomAttributeModel instantiates a new CustomAttributeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAttributeModelWithDefaults

`func NewCustomAttributeModelWithDefaults() *CustomAttributeModel`

NewCustomAttributeModelWithDefaults instantiates a new CustomAttributeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomAttributeModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomAttributeModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomAttributeModel) SetId(v string)`

SetId sets Id field to given value.


### GetCode

`func (o *CustomAttributeModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CustomAttributeModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CustomAttributeModel) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CustomAttributeModel) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *CustomAttributeModel) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *CustomAttributeModel) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetType

`func (o *CustomAttributeModel) GetType() CustomAttributeTypesEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomAttributeModel) GetTypeOk() (*CustomAttributeTypesEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomAttributeModel) SetType(v CustomAttributeTypesEnum)`

SetType sets Type field to given value.


### GetOptions

`func (o *CustomAttributeModel) GetOptions() []CustomAttributeOptionModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CustomAttributeModel) GetOptionsOk() (*[]CustomAttributeOptionModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CustomAttributeModel) SetOptions(v []CustomAttributeOptionModel)`

SetOptions sets Options field to given value.


### GetTargets

`func (o *CustomAttributeModel) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *CustomAttributeModel) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *CustomAttributeModel) SetTargets(v []string)`

SetTargets sets Targets field to given value.


### GetIsReadOnly

`func (o *CustomAttributeModel) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *CustomAttributeModel) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *CustomAttributeModel) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.


### GetIsDeleted

`func (o *CustomAttributeModel) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CustomAttributeModel) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CustomAttributeModel) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetIsSystem

`func (o *CustomAttributeModel) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *CustomAttributeModel) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *CustomAttributeModel) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.


### GetName

`func (o *CustomAttributeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomAttributeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomAttributeModel) SetName(v string)`

SetName sets Name field to given value.


### GetIsEnabled

`func (o *CustomAttributeModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CustomAttributeModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CustomAttributeModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetIsRequired

`func (o *CustomAttributeModel) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *CustomAttributeModel) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *CustomAttributeModel) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.


### GetIsGlobal

`func (o *CustomAttributeModel) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *CustomAttributeModel) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *CustomAttributeModel) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


