# ApiV2CustomAttributesTemplatesPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the attribute template | 
**CustomAttributeIds** | Pointer to **[]string** | Collection of attribute IDs | [optional] 
**Name** | **string** | Custom attributes template name | 

## Methods

### NewApiV2CustomAttributesTemplatesPutRequest

`func NewApiV2CustomAttributesTemplatesPutRequest(id string, name string, ) *ApiV2CustomAttributesTemplatesPutRequest`

NewApiV2CustomAttributesTemplatesPutRequest instantiates a new ApiV2CustomAttributesTemplatesPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2CustomAttributesTemplatesPutRequestWithDefaults

`func NewApiV2CustomAttributesTemplatesPutRequestWithDefaults() *ApiV2CustomAttributesTemplatesPutRequest`

NewApiV2CustomAttributesTemplatesPutRequestWithDefaults instantiates a new ApiV2CustomAttributesTemplatesPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiV2CustomAttributesTemplatesPutRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiV2CustomAttributesTemplatesPutRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiV2CustomAttributesTemplatesPutRequest) SetId(v string)`

SetId sets Id field to given value.


### GetCustomAttributeIds

`func (o *ApiV2CustomAttributesTemplatesPutRequest) GetCustomAttributeIds() []string`

GetCustomAttributeIds returns the CustomAttributeIds field if non-nil, zero value otherwise.

### GetCustomAttributeIdsOk

`func (o *ApiV2CustomAttributesTemplatesPutRequest) GetCustomAttributeIdsOk() (*[]string, bool)`

GetCustomAttributeIdsOk returns a tuple with the CustomAttributeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributeIds

`func (o *ApiV2CustomAttributesTemplatesPutRequest) SetCustomAttributeIds(v []string)`

SetCustomAttributeIds sets CustomAttributeIds field to given value.

### HasCustomAttributeIds

`func (o *ApiV2CustomAttributesTemplatesPutRequest) HasCustomAttributeIds() bool`

HasCustomAttributeIds returns a boolean if a field has been set.

### SetCustomAttributeIdsNil

`func (o *ApiV2CustomAttributesTemplatesPutRequest) SetCustomAttributeIdsNil(b bool)`

 SetCustomAttributeIdsNil sets the value for CustomAttributeIds to be an explicit nil

### UnsetCustomAttributeIds
`func (o *ApiV2CustomAttributesTemplatesPutRequest) UnsetCustomAttributeIds()`

UnsetCustomAttributeIds ensures that no value is present for CustomAttributeIds, not even an explicit nil
### GetName

`func (o *ApiV2CustomAttributesTemplatesPutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV2CustomAttributesTemplatesPutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV2CustomAttributesTemplatesPutRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


