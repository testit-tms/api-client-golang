# CreateProjectsAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**[]CustomAttributeOptionPostModel**](CustomAttributeOptionPostModel.md) | Collection of attribute options  &lt;br /&gt;  Available for attributes of type &#x60;options&#x60; and &#x60;multiple options&#x60; only | [optional] 
**Type** | [**CustomAttributeTypesEnum**](CustomAttributeTypesEnum.md) |  | 
**Name** | **string** | Name of the attribute | 
**IsEnabled** | Pointer to **bool** | Indicates if the attribute is enabled | [optional] 
**IsRequired** | Pointer to **bool** | Indicates if the attribute value is mandatory to specify | [optional] 
**IsGlobal** | Pointer to **bool** | Indicates if the attribute is available across all projects | [optional] 

## Methods

### NewCreateProjectsAttributeRequest

`func NewCreateProjectsAttributeRequest(type_ CustomAttributeTypesEnum, name string, ) *CreateProjectsAttributeRequest`

NewCreateProjectsAttributeRequest instantiates a new CreateProjectsAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectsAttributeRequestWithDefaults

`func NewCreateProjectsAttributeRequestWithDefaults() *CreateProjectsAttributeRequest`

NewCreateProjectsAttributeRequestWithDefaults instantiates a new CreateProjectsAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *CreateProjectsAttributeRequest) GetOptions() []CustomAttributeOptionPostModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateProjectsAttributeRequest) GetOptionsOk() (*[]CustomAttributeOptionPostModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateProjectsAttributeRequest) SetOptions(v []CustomAttributeOptionPostModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CreateProjectsAttributeRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *CreateProjectsAttributeRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *CreateProjectsAttributeRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetType

`func (o *CreateProjectsAttributeRequest) GetType() CustomAttributeTypesEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateProjectsAttributeRequest) GetTypeOk() (*CustomAttributeTypesEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateProjectsAttributeRequest) SetType(v CustomAttributeTypesEnum)`

SetType sets Type field to given value.


### GetName

`func (o *CreateProjectsAttributeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProjectsAttributeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProjectsAttributeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsEnabled

`func (o *CreateProjectsAttributeRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CreateProjectsAttributeRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CreateProjectsAttributeRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *CreateProjectsAttributeRequest) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsRequired

`func (o *CreateProjectsAttributeRequest) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *CreateProjectsAttributeRequest) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *CreateProjectsAttributeRequest) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *CreateProjectsAttributeRequest) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetIsGlobal

`func (o *CreateProjectsAttributeRequest) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *CreateProjectsAttributeRequest) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *CreateProjectsAttributeRequest) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.

### HasIsGlobal

`func (o *CreateProjectsAttributeRequest) HasIsGlobal() bool`

HasIsGlobal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


