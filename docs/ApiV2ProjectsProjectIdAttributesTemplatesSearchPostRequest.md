# ApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of custom attribute template | [optional] 
**CustomAttributeTypes** | Pointer to [**[]CustomAttributeTypesEnum**](CustomAttributeTypesEnum.md) | Collection of custom attributes types | [optional] 

## Methods

### NewApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest

`func NewApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest() *ApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest`

NewApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest instantiates a new ApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequestWithDefaults

`func NewApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequestWithDefaults() *ApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest`

NewApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequestWithDefaults instantiates a new ApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCustomAttributeTypes

`func (o *ApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) GetCustomAttributeTypes() []CustomAttributeTypesEnum`

GetCustomAttributeTypes returns the CustomAttributeTypes field if non-nil, zero value otherwise.

### GetCustomAttributeTypesOk

`func (o *ApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) GetCustomAttributeTypesOk() (*[]CustomAttributeTypesEnum, bool)`

GetCustomAttributeTypesOk returns a tuple with the CustomAttributeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributeTypes

`func (o *ApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) SetCustomAttributeTypes(v []CustomAttributeTypesEnum)`

SetCustomAttributeTypes sets CustomAttributeTypes field to given value.

### HasCustomAttributeTypes

`func (o *ApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) HasCustomAttributeTypes() bool`

HasCustomAttributeTypes returns a boolean if a field has been set.

### SetCustomAttributeTypesNil

`func (o *ApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) SetCustomAttributeTypesNil(b bool)`

 SetCustomAttributeTypesNil sets the value for CustomAttributeTypes to be an explicit nil

### UnsetCustomAttributeTypes
`func (o *ApiV2ProjectsProjectIdAttributesTemplatesSearchPostRequest) UnsetCustomAttributeTypes()`

UnsetCustomAttributeTypes ensures that no value is present for CustomAttributeTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


