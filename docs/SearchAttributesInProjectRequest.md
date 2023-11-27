# SearchAttributesInProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Specifies an attribute name to search for | 
**IsRequired** | Pointer to **NullableBool** | Specifies an attribute mandatory status to search for | [optional] 
**IsGlobal** | Pointer to **NullableBool** | Specifies an attribute global status to search for | [optional] 
**Types** | [**[]CustomAttributeTypesEnum**](CustomAttributeTypesEnum.md) | Specifies an attribute types to search for | 
**IsEnabled** | Pointer to **NullableBool** | Specifies an attribute enabled status to search for | [optional] 

## Methods

### NewSearchAttributesInProjectRequest

`func NewSearchAttributesInProjectRequest(name string, types []CustomAttributeTypesEnum, ) *SearchAttributesInProjectRequest`

NewSearchAttributesInProjectRequest instantiates a new SearchAttributesInProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchAttributesInProjectRequestWithDefaults

`func NewSearchAttributesInProjectRequestWithDefaults() *SearchAttributesInProjectRequest`

NewSearchAttributesInProjectRequestWithDefaults instantiates a new SearchAttributesInProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SearchAttributesInProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchAttributesInProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchAttributesInProjectRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsRequired

`func (o *SearchAttributesInProjectRequest) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *SearchAttributesInProjectRequest) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *SearchAttributesInProjectRequest) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *SearchAttributesInProjectRequest) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### SetIsRequiredNil

`func (o *SearchAttributesInProjectRequest) SetIsRequiredNil(b bool)`

 SetIsRequiredNil sets the value for IsRequired to be an explicit nil

### UnsetIsRequired
`func (o *SearchAttributesInProjectRequest) UnsetIsRequired()`

UnsetIsRequired ensures that no value is present for IsRequired, not even an explicit nil
### GetIsGlobal

`func (o *SearchAttributesInProjectRequest) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *SearchAttributesInProjectRequest) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *SearchAttributesInProjectRequest) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.

### HasIsGlobal

`func (o *SearchAttributesInProjectRequest) HasIsGlobal() bool`

HasIsGlobal returns a boolean if a field has been set.

### SetIsGlobalNil

`func (o *SearchAttributesInProjectRequest) SetIsGlobalNil(b bool)`

 SetIsGlobalNil sets the value for IsGlobal to be an explicit nil

### UnsetIsGlobal
`func (o *SearchAttributesInProjectRequest) UnsetIsGlobal()`

UnsetIsGlobal ensures that no value is present for IsGlobal, not even an explicit nil
### GetTypes

`func (o *SearchAttributesInProjectRequest) GetTypes() []CustomAttributeTypesEnum`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *SearchAttributesInProjectRequest) GetTypesOk() (*[]CustomAttributeTypesEnum, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *SearchAttributesInProjectRequest) SetTypes(v []CustomAttributeTypesEnum)`

SetTypes sets Types field to given value.


### GetIsEnabled

`func (o *SearchAttributesInProjectRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SearchAttributesInProjectRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SearchAttributesInProjectRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *SearchAttributesInProjectRequest) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *SearchAttributesInProjectRequest) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *SearchAttributesInProjectRequest) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


