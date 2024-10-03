# ApiV2CustomAttributesTemplatesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomAttributeIds** | Pointer to **[]string** | Collection of attribute IDs | [optional] 
**Name** | **string** | Custom attributes template name | 

## Methods

### NewApiV2CustomAttributesTemplatesPostRequest

`func NewApiV2CustomAttributesTemplatesPostRequest(name string, ) *ApiV2CustomAttributesTemplatesPostRequest`

NewApiV2CustomAttributesTemplatesPostRequest instantiates a new ApiV2CustomAttributesTemplatesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2CustomAttributesTemplatesPostRequestWithDefaults

`func NewApiV2CustomAttributesTemplatesPostRequestWithDefaults() *ApiV2CustomAttributesTemplatesPostRequest`

NewApiV2CustomAttributesTemplatesPostRequestWithDefaults instantiates a new ApiV2CustomAttributesTemplatesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomAttributeIds

`func (o *ApiV2CustomAttributesTemplatesPostRequest) GetCustomAttributeIds() []string`

GetCustomAttributeIds returns the CustomAttributeIds field if non-nil, zero value otherwise.

### GetCustomAttributeIdsOk

`func (o *ApiV2CustomAttributesTemplatesPostRequest) GetCustomAttributeIdsOk() (*[]string, bool)`

GetCustomAttributeIdsOk returns a tuple with the CustomAttributeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributeIds

`func (o *ApiV2CustomAttributesTemplatesPostRequest) SetCustomAttributeIds(v []string)`

SetCustomAttributeIds sets CustomAttributeIds field to given value.

### HasCustomAttributeIds

`func (o *ApiV2CustomAttributesTemplatesPostRequest) HasCustomAttributeIds() bool`

HasCustomAttributeIds returns a boolean if a field has been set.

### SetCustomAttributeIdsNil

`func (o *ApiV2CustomAttributesTemplatesPostRequest) SetCustomAttributeIdsNil(b bool)`

 SetCustomAttributeIdsNil sets the value for CustomAttributeIds to be an explicit nil

### UnsetCustomAttributeIds
`func (o *ApiV2CustomAttributesTemplatesPostRequest) UnsetCustomAttributeIds()`

UnsetCustomAttributeIds ensures that no value is present for CustomAttributeIds, not even an explicit nil
### GetName

`func (o *ApiV2CustomAttributesTemplatesPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2CustomAttributesTemplatesPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2CustomAttributesTemplatesPostRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


