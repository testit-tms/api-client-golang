# ApiV2CustomAttributesGlobalPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of attribute | 
**IsEnabled** | Pointer to **NullableBool** | Indicates whether the attribute is available | [optional] 
**IsRequired** | Pointer to **NullableBool** | Indicates whether the attribute value is mandatory to specify | [optional] 
**Options** | Pointer to [**[]CustomAttributeOptionPostModel**](CustomAttributeOptionPostModel.md) | Collection of attribute options  &lt;br /&gt;  Available for attributes of type &#x60;options&#x60; and &#x60;multiple options&#x60; only | [optional] 
**Type** | [**CustomAttributeTypesEnum**](CustomAttributeTypesEnum.md) |  | 

## Methods

### NewApiV2CustomAttributesGlobalPostRequest

`func NewApiV2CustomAttributesGlobalPostRequest(name string, type_ CustomAttributeTypesEnum, ) *ApiV2CustomAttributesGlobalPostRequest`

NewApiV2CustomAttributesGlobalPostRequest instantiates a new ApiV2CustomAttributesGlobalPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2CustomAttributesGlobalPostRequestWithDefaults

`func NewApiV2CustomAttributesGlobalPostRequestWithDefaults() *ApiV2CustomAttributesGlobalPostRequest`

NewApiV2CustomAttributesGlobalPostRequestWithDefaults instantiates a new ApiV2CustomAttributesGlobalPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiV2CustomAttributesGlobalPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2CustomAttributesGlobalPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2CustomAttributesGlobalPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsEnabled

`func (o *ApiV2CustomAttributesGlobalPostRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ApiV2CustomAttributesGlobalPostRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ApiV2CustomAttributesGlobalPostRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ApiV2CustomAttributesGlobalPostRequest) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *ApiV2CustomAttributesGlobalPostRequest) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *ApiV2CustomAttributesGlobalPostRequest) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetIsRequired

`func (o *ApiV2CustomAttributesGlobalPostRequest) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *ApiV2CustomAttributesGlobalPostRequest) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *ApiV2CustomAttributesGlobalPostRequest) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *ApiV2CustomAttributesGlobalPostRequest) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### SetIsRequiredNil

`func (o *ApiV2CustomAttributesGlobalPostRequest) SetIsRequiredNil(b bool)`

 SetIsRequiredNil sets the value for IsRequired to be an explicit nil

### UnsetIsRequired
`func (o *ApiV2CustomAttributesGlobalPostRequest) UnsetIsRequired()`

UnsetIsRequired ensures that no value is present for IsRequired, not even an explicit nil
### GetOptions

`func (o *ApiV2CustomAttributesGlobalPostRequest) GetOptions() []CustomAttributeOptionPostModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ApiV2CustomAttributesGlobalPostRequest) GetOptionsOk() (*[]CustomAttributeOptionPostModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ApiV2CustomAttributesGlobalPostRequest) SetOptions(v []CustomAttributeOptionPostModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ApiV2CustomAttributesGlobalPostRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *ApiV2CustomAttributesGlobalPostRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *ApiV2CustomAttributesGlobalPostRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetType

`func (o *ApiV2CustomAttributesGlobalPostRequest) GetType() CustomAttributeTypesEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiV2CustomAttributesGlobalPostRequest) GetTypeOk() (*CustomAttributeTypesEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiV2CustomAttributesGlobalPostRequest) SetType(v CustomAttributeTypesEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


