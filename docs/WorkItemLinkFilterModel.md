# WorkItemLinkFilterModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Types** | Pointer to [**[]LinkType**](LinkType.md) |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Urls** | Pointer to **[]string** |  | [optional] 
**OnlyWithoutLinks** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewWorkItemLinkFilterModel

`func NewWorkItemLinkFilterModel() *WorkItemLinkFilterModel`

NewWorkItemLinkFilterModel instantiates a new WorkItemLinkFilterModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemLinkFilterModelWithDefaults

`func NewWorkItemLinkFilterModelWithDefaults() *WorkItemLinkFilterModel`

NewWorkItemLinkFilterModelWithDefaults instantiates a new WorkItemLinkFilterModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypes

`func (o *WorkItemLinkFilterModel) GetTypes() []LinkType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *WorkItemLinkFilterModel) GetTypesOk() (*[]LinkType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *WorkItemLinkFilterModel) SetTypes(v []LinkType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *WorkItemLinkFilterModel) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *WorkItemLinkFilterModel) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *WorkItemLinkFilterModel) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil
### GetTitle

`func (o *WorkItemLinkFilterModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkItemLinkFilterModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkItemLinkFilterModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkItemLinkFilterModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WorkItemLinkFilterModel) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WorkItemLinkFilterModel) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUrls

`func (o *WorkItemLinkFilterModel) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *WorkItemLinkFilterModel) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *WorkItemLinkFilterModel) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *WorkItemLinkFilterModel) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### SetUrlsNil

`func (o *WorkItemLinkFilterModel) SetUrlsNil(b bool)`

 SetUrlsNil sets the value for Urls to be an explicit nil

### UnsetUrls
`func (o *WorkItemLinkFilterModel) UnsetUrls()`

UnsetUrls ensures that no value is present for Urls, not even an explicit nil
### GetOnlyWithoutLinks

`func (o *WorkItemLinkFilterModel) GetOnlyWithoutLinks() bool`

GetOnlyWithoutLinks returns the OnlyWithoutLinks field if non-nil, zero value otherwise.

### GetOnlyWithoutLinksOk

`func (o *WorkItemLinkFilterModel) GetOnlyWithoutLinksOk() (*bool, bool)`

GetOnlyWithoutLinksOk returns a tuple with the OnlyWithoutLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyWithoutLinks

`func (o *WorkItemLinkFilterModel) SetOnlyWithoutLinks(v bool)`

SetOnlyWithoutLinks sets OnlyWithoutLinks field to given value.

### HasOnlyWithoutLinks

`func (o *WorkItemLinkFilterModel) HasOnlyWithoutLinks() bool`

HasOnlyWithoutLinks returns a boolean if a field has been set.

### SetOnlyWithoutLinksNil

`func (o *WorkItemLinkFilterModel) SetOnlyWithoutLinksNil(b bool)`

 SetOnlyWithoutLinksNil sets the value for OnlyWithoutLinks to be an explicit nil

### UnsetOnlyWithoutLinks
`func (o *WorkItemLinkFilterModel) UnsetOnlyWithoutLinks()`

UnsetOnlyWithoutLinks ensures that no value is present for OnlyWithoutLinks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


