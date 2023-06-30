# ApiV2CustomAttributesGlobalIdPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of attribute | 
**Options** | Pointer to [**[]CustomAttributeOptionModel**](CustomAttributeOptionModel.md) | Collection of attribute options  &lt;br /&gt;  Available for attributes of type &#x60;options&#x60; and &#x60;multiple options&#x60; only | [optional] 
**IsEnabled** | Pointer to **NullableBool** | Indicates whether the attribute is available | [optional] 
**IsRequired** | Pointer to **NullableBool** | Indicates whether the attribute value is mandatory to specify | [optional] 

## Methods

### NewApiV2CustomAttributesGlobalIdPutRequest

`func NewApiV2CustomAttributesGlobalIdPutRequest(name string, ) *ApiV2CustomAttributesGlobalIdPutRequest`

NewApiV2CustomAttributesGlobalIdPutRequest instantiates a new ApiV2CustomAttributesGlobalIdPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2CustomAttributesGlobalIdPutRequestWithDefaults

`func NewApiV2CustomAttributesGlobalIdPutRequestWithDefaults() *ApiV2CustomAttributesGlobalIdPutRequest`

NewApiV2CustomAttributesGlobalIdPutRequestWithDefaults instantiates a new ApiV2CustomAttributesGlobalIdPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiV2CustomAttributesGlobalIdPutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2CustomAttributesGlobalIdPutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2CustomAttributesGlobalIdPutRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *ApiV2CustomAttributesGlobalIdPutRequest) GetOptions() []CustomAttributeOptionModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ApiV2CustomAttributesGlobalIdPutRequest) GetOptionsOk() (*[]CustomAttributeOptionModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ApiV2CustomAttributesGlobalIdPutRequest) SetOptions(v []CustomAttributeOptionModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ApiV2CustomAttributesGlobalIdPutRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *ApiV2CustomAttributesGlobalIdPutRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *ApiV2CustomAttributesGlobalIdPutRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetIsEnabled

`func (o *ApiV2CustomAttributesGlobalIdPutRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ApiV2CustomAttributesGlobalIdPutRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ApiV2CustomAttributesGlobalIdPutRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ApiV2CustomAttributesGlobalIdPutRequest) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *ApiV2CustomAttributesGlobalIdPutRequest) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *ApiV2CustomAttributesGlobalIdPutRequest) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetIsRequired

`func (o *ApiV2CustomAttributesGlobalIdPutRequest) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *ApiV2CustomAttributesGlobalIdPutRequest) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *ApiV2CustomAttributesGlobalIdPutRequest) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *ApiV2CustomAttributesGlobalIdPutRequest) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### SetIsRequiredNil

`func (o *ApiV2CustomAttributesGlobalIdPutRequest) SetIsRequiredNil(b bool)`

 SetIsRequiredNil sets the value for IsRequired to be an explicit nil

### UnsetIsRequired
`func (o *ApiV2CustomAttributesGlobalIdPutRequest) UnsetIsRequired()`

UnsetIsRequired ensures that no value is present for IsRequired, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


