# TestResultLinkApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Link unique identifier | [optional] 
**Title** | Pointer to **NullableString** | Link name. | [optional] 
**Url** | **string** | Address can be specified without protocol, but necessarily with the domain. | 
**Description** | Pointer to **NullableString** | Link description. | [optional] 
**Type** | Pointer to [**NullableLinkType**](LinkType.md) | Specifies the type of the link. | [optional] 
**HasInfo** | **bool** | Flag defines if link relates to integrated jira service | 
**Name** | Pointer to **NullableString** | Link name. Backward compatibility. | [optional] [readonly] 

## Methods

### NewTestResultLinkApiResult

`func NewTestResultLinkApiResult(url string, hasInfo bool, ) *TestResultLinkApiResult`

NewTestResultLinkApiResult instantiates a new TestResultLinkApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultLinkApiResultWithDefaults

`func NewTestResultLinkApiResultWithDefaults() *TestResultLinkApiResult`

NewTestResultLinkApiResultWithDefaults instantiates a new TestResultLinkApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestResultLinkApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestResultLinkApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestResultLinkApiResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestResultLinkApiResult) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TestResultLinkApiResult) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TestResultLinkApiResult) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTitle

`func (o *TestResultLinkApiResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TestResultLinkApiResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TestResultLinkApiResult) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TestResultLinkApiResult) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TestResultLinkApiResult) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TestResultLinkApiResult) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUrl

`func (o *TestResultLinkApiResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TestResultLinkApiResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TestResultLinkApiResult) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDescription

`func (o *TestResultLinkApiResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestResultLinkApiResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestResultLinkApiResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestResultLinkApiResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestResultLinkApiResult) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestResultLinkApiResult) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *TestResultLinkApiResult) GetType() LinkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestResultLinkApiResult) GetTypeOk() (*LinkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestResultLinkApiResult) SetType(v LinkType)`

SetType sets Type field to given value.

### HasType

`func (o *TestResultLinkApiResult) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *TestResultLinkApiResult) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TestResultLinkApiResult) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetHasInfo

`func (o *TestResultLinkApiResult) GetHasInfo() bool`

GetHasInfo returns the HasInfo field if non-nil, zero value otherwise.

### GetHasInfoOk

`func (o *TestResultLinkApiResult) GetHasInfoOk() (*bool, bool)`

GetHasInfoOk returns a tuple with the HasInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasInfo

`func (o *TestResultLinkApiResult) SetHasInfo(v bool)`

SetHasInfo sets HasInfo field to given value.


### GetName

`func (o *TestResultLinkApiResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestResultLinkApiResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestResultLinkApiResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestResultLinkApiResult) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestResultLinkApiResult) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestResultLinkApiResult) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


