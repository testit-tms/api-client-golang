# UpdateProjectsAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the attribute | 
**Options** | Pointer to [**[]CustomAttributeOptionModel**](CustomAttributeOptionModel.md) | Collection of the attribute options  &lt;br /&gt;  Available for attributes of type &#x60;options&#x60; and &#x60;multiple options&#x60; only | [optional] 
**Type** | [**CustomAttributeTypesEnum**](CustomAttributeTypesEnum.md) |  | 
**IsDeleted** | **bool** | Indicates if the entity is deleted | 
**Name** | **string** | Name of the attribute | 
**IsEnabled** | **bool** | Indicates if the attribute is enabled | 
**IsRequired** | **bool** | Indicates if the attribute value is mandatory to specify | 
**IsGlobal** | **bool** | Indicates if the attribute is available across all projects | 

## Methods

### NewUpdateProjectsAttributeRequest

`func NewUpdateProjectsAttributeRequest(id string, type_ CustomAttributeTypesEnum, isDeleted bool, name string, isEnabled bool, isRequired bool, isGlobal bool, ) *UpdateProjectsAttributeRequest`

NewUpdateProjectsAttributeRequest instantiates a new UpdateProjectsAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProjectsAttributeRequestWithDefaults

`func NewUpdateProjectsAttributeRequestWithDefaults() *UpdateProjectsAttributeRequest`

NewUpdateProjectsAttributeRequestWithDefaults instantiates a new UpdateProjectsAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateProjectsAttributeRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateProjectsAttributeRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateProjectsAttributeRequest) SetId(v string)`

SetId sets Id field to given value.


### GetOptions

`func (o *UpdateProjectsAttributeRequest) GetOptions() []CustomAttributeOptionModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateProjectsAttributeRequest) GetOptionsOk() (*[]CustomAttributeOptionModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateProjectsAttributeRequest) SetOptions(v []CustomAttributeOptionModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpdateProjectsAttributeRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *UpdateProjectsAttributeRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *UpdateProjectsAttributeRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetType

`func (o *UpdateProjectsAttributeRequest) GetType() CustomAttributeTypesEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateProjectsAttributeRequest) GetTypeOk() (*CustomAttributeTypesEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateProjectsAttributeRequest) SetType(v CustomAttributeTypesEnum)`

SetType sets Type field to given value.


### GetIsDeleted

`func (o *UpdateProjectsAttributeRequest) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UpdateProjectsAttributeRequest) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UpdateProjectsAttributeRequest) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetName

`func (o *UpdateProjectsAttributeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProjectsAttributeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProjectsAttributeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsEnabled

`func (o *UpdateProjectsAttributeRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UpdateProjectsAttributeRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UpdateProjectsAttributeRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetIsRequired

`func (o *UpdateProjectsAttributeRequest) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *UpdateProjectsAttributeRequest) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *UpdateProjectsAttributeRequest) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.


### GetIsGlobal

`func (o *UpdateProjectsAttributeRequest) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *UpdateProjectsAttributeRequest) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *UpdateProjectsAttributeRequest) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


