# LinkCreateApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** | Link name. | [optional] 
**Url** | **string** | Address can be specified without protocol, but necessarily with the domain. | 
**Description** | Pointer to **NullableString** | Link description. | [optional] 
**Type** | Pointer to [**NullableLinkType**](LinkType.md) | Specifies the type of the link. | [optional] 
**HasInfo** | **bool** | Flag defines if link relates to integrated external service | 

## Methods

### NewLinkCreateApiModel

`func NewLinkCreateApiModel(url string, hasInfo bool, ) *LinkCreateApiModel`

NewLinkCreateApiModel instantiates a new LinkCreateApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkCreateApiModelWithDefaults

`func NewLinkCreateApiModelWithDefaults() *LinkCreateApiModel`

NewLinkCreateApiModelWithDefaults instantiates a new LinkCreateApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *LinkCreateApiModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LinkCreateApiModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LinkCreateApiModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LinkCreateApiModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *LinkCreateApiModel) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *LinkCreateApiModel) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUrl

`func (o *LinkCreateApiModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LinkCreateApiModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LinkCreateApiModel) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDescription

`func (o *LinkCreateApiModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LinkCreateApiModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LinkCreateApiModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LinkCreateApiModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LinkCreateApiModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LinkCreateApiModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *LinkCreateApiModel) GetType() LinkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkCreateApiModel) GetTypeOk() (*LinkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkCreateApiModel) SetType(v LinkType)`

SetType sets Type field to given value.

### HasType

`func (o *LinkCreateApiModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *LinkCreateApiModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *LinkCreateApiModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetHasInfo

`func (o *LinkCreateApiModel) GetHasInfo() bool`

GetHasInfo returns the HasInfo field if non-nil, zero value otherwise.

### GetHasInfoOk

`func (o *LinkCreateApiModel) GetHasInfoOk() (*bool, bool)`

GetHasInfoOk returns a tuple with the HasInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasInfo

`func (o *LinkCreateApiModel) SetHasInfo(v bool)`

SetHasInfo sets HasInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


