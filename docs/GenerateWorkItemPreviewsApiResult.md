# GenerateWorkItemPreviewsApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Previews** | [**[]WorkItemPreviewApiModel**](WorkItemPreviewApiModel.md) |  | 
**Link** | Pointer to [**NullablePreviewsIssueLinkApiResult**](PreviewsIssueLinkApiResult.md) |  | [optional] 

## Methods

### NewGenerateWorkItemPreviewsApiResult

`func NewGenerateWorkItemPreviewsApiResult(previews []WorkItemPreviewApiModel, ) *GenerateWorkItemPreviewsApiResult`

NewGenerateWorkItemPreviewsApiResult instantiates a new GenerateWorkItemPreviewsApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateWorkItemPreviewsApiResultWithDefaults

`func NewGenerateWorkItemPreviewsApiResultWithDefaults() *GenerateWorkItemPreviewsApiResult`

NewGenerateWorkItemPreviewsApiResultWithDefaults instantiates a new GenerateWorkItemPreviewsApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreviews

`func (o *GenerateWorkItemPreviewsApiResult) GetPreviews() []WorkItemPreviewApiModel`

GetPreviews returns the Previews field if non-nil, zero value otherwise.

### GetPreviewsOk

`func (o *GenerateWorkItemPreviewsApiResult) GetPreviewsOk() (*[]WorkItemPreviewApiModel, bool)`

GetPreviewsOk returns a tuple with the Previews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviews

`func (o *GenerateWorkItemPreviewsApiResult) SetPreviews(v []WorkItemPreviewApiModel)`

SetPreviews sets Previews field to given value.


### GetLink

`func (o *GenerateWorkItemPreviewsApiResult) GetLink() PreviewsIssueLinkApiResult`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *GenerateWorkItemPreviewsApiResult) GetLinkOk() (*PreviewsIssueLinkApiResult, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *GenerateWorkItemPreviewsApiResult) SetLink(v PreviewsIssueLinkApiResult)`

SetLink sets Link field to given value.

### HasLink

`func (o *GenerateWorkItemPreviewsApiResult) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *GenerateWorkItemPreviewsApiResult) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *GenerateWorkItemPreviewsApiResult) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


