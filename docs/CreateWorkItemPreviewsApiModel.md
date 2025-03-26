# CreateWorkItemPreviewsApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SectionId** | **string** |  | 
**Previews** | [**[]WorkItemPreviewApiModel**](WorkItemPreviewApiModel.md) |  | 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateWorkItemPreviewsApiModel

`func NewCreateWorkItemPreviewsApiModel(sectionId string, previews []WorkItemPreviewApiModel, ) *CreateWorkItemPreviewsApiModel`

NewCreateWorkItemPreviewsApiModel instantiates a new CreateWorkItemPreviewsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkItemPreviewsApiModelWithDefaults

`func NewCreateWorkItemPreviewsApiModelWithDefaults() *CreateWorkItemPreviewsApiModel`

NewCreateWorkItemPreviewsApiModelWithDefaults instantiates a new CreateWorkItemPreviewsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSectionId

`func (o *CreateWorkItemPreviewsApiModel) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *CreateWorkItemPreviewsApiModel) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *CreateWorkItemPreviewsApiModel) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.


### GetPreviews

`func (o *CreateWorkItemPreviewsApiModel) GetPreviews() []WorkItemPreviewApiModel`

GetPreviews returns the Previews field if non-nil, zero value otherwise.

### GetPreviewsOk

`func (o *CreateWorkItemPreviewsApiModel) GetPreviewsOk() (*[]WorkItemPreviewApiModel, bool)`

GetPreviewsOk returns a tuple with the Previews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviews

`func (o *CreateWorkItemPreviewsApiModel) SetPreviews(v []WorkItemPreviewApiModel)`

SetPreviews sets Previews field to given value.


### GetAttributes

`func (o *CreateWorkItemPreviewsApiModel) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CreateWorkItemPreviewsApiModel) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CreateWorkItemPreviewsApiModel) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CreateWorkItemPreviewsApiModel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *CreateWorkItemPreviewsApiModel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *CreateWorkItemPreviewsApiModel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


