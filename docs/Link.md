# Link

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Link unique identifier | [optional] 
**Title** | Pointer to **NullableString** | Link name. | [optional] 
**Url** | **string** | Address can be specified without protocol, but necessarily with the domain. | 
**Description** | Pointer to **NullableString** | Link description. | [optional] 
**Type** | Pointer to [**NullableLinkType**](LinkType.md) | Specifies the type of the link. | [optional] 
**HasInfo** | **bool** | Flag defines if link relates to integrated jira service | 

## Methods

### NewLink

`func NewLink(url string, hasInfo bool, ) *Link`

NewLink instantiates a new Link object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkWithDefaults

`func NewLinkWithDefaults() *Link`

NewLinkWithDefaults instantiates a new Link object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Link) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Link) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Link) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Link) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Link) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Link) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTitle

`func (o *Link) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Link) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Link) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Link) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Link) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Link) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUrl

`func (o *Link) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Link) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Link) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDescription

`func (o *Link) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Link) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Link) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Link) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Link) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Link) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *Link) GetType() LinkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Link) GetTypeOk() (*LinkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Link) SetType(v LinkType)`

SetType sets Type field to given value.

### HasType

`func (o *Link) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Link) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Link) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetHasInfo

`func (o *Link) GetHasInfo() bool`

GetHasInfo returns the HasInfo field if non-nil, zero value otherwise.

### GetHasInfoOk

`func (o *Link) GetHasInfoOk() (*bool, bool)`

GetHasInfoOk returns a tuple with the HasInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasInfo

`func (o *Link) SetHasInfo(v bool)`

SetHasInfo sets HasInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


