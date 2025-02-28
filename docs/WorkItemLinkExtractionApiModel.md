# WorkItemLinkExtractionApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectIds** | Pointer to [**NullableGuidExtractionModel**](GuidExtractionModel.md) |  | [optional] 
**WorkItemIds** | Pointer to [**NullableGuidExtractionModel**](GuidExtractionModel.md) |  | [optional] 
**LinkUrls** | Pointer to [**NullableStringExtractionModel**](StringExtractionModel.md) |  | [optional] 

## Methods

### NewWorkItemLinkExtractionApiModel

`func NewWorkItemLinkExtractionApiModel() *WorkItemLinkExtractionApiModel`

NewWorkItemLinkExtractionApiModel instantiates a new WorkItemLinkExtractionApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemLinkExtractionApiModelWithDefaults

`func NewWorkItemLinkExtractionApiModelWithDefaults() *WorkItemLinkExtractionApiModel`

NewWorkItemLinkExtractionApiModelWithDefaults instantiates a new WorkItemLinkExtractionApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectIds

`func (o *WorkItemLinkExtractionApiModel) GetProjectIds() GuidExtractionModel`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *WorkItemLinkExtractionApiModel) GetProjectIdsOk() (*GuidExtractionModel, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *WorkItemLinkExtractionApiModel) SetProjectIds(v GuidExtractionModel)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *WorkItemLinkExtractionApiModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *WorkItemLinkExtractionApiModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *WorkItemLinkExtractionApiModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetWorkItemIds

`func (o *WorkItemLinkExtractionApiModel) GetWorkItemIds() GuidExtractionModel`

GetWorkItemIds returns the WorkItemIds field if non-nil, zero value otherwise.

### GetWorkItemIdsOk

`func (o *WorkItemLinkExtractionApiModel) GetWorkItemIdsOk() (*GuidExtractionModel, bool)`

GetWorkItemIdsOk returns a tuple with the WorkItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemIds

`func (o *WorkItemLinkExtractionApiModel) SetWorkItemIds(v GuidExtractionModel)`

SetWorkItemIds sets WorkItemIds field to given value.

### HasWorkItemIds

`func (o *WorkItemLinkExtractionApiModel) HasWorkItemIds() bool`

HasWorkItemIds returns a boolean if a field has been set.

### SetWorkItemIdsNil

`func (o *WorkItemLinkExtractionApiModel) SetWorkItemIdsNil(b bool)`

 SetWorkItemIdsNil sets the value for WorkItemIds to be an explicit nil

### UnsetWorkItemIds
`func (o *WorkItemLinkExtractionApiModel) UnsetWorkItemIds()`

UnsetWorkItemIds ensures that no value is present for WorkItemIds, not even an explicit nil
### GetLinkUrls

`func (o *WorkItemLinkExtractionApiModel) GetLinkUrls() StringExtractionModel`

GetLinkUrls returns the LinkUrls field if non-nil, zero value otherwise.

### GetLinkUrlsOk

`func (o *WorkItemLinkExtractionApiModel) GetLinkUrlsOk() (*StringExtractionModel, bool)`

GetLinkUrlsOk returns a tuple with the LinkUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUrls

`func (o *WorkItemLinkExtractionApiModel) SetLinkUrls(v StringExtractionModel)`

SetLinkUrls sets LinkUrls field to given value.

### HasLinkUrls

`func (o *WorkItemLinkExtractionApiModel) HasLinkUrls() bool`

HasLinkUrls returns a boolean if a field has been set.

### SetLinkUrlsNil

`func (o *WorkItemLinkExtractionApiModel) SetLinkUrlsNil(b bool)`

 SetLinkUrlsNil sets the value for LinkUrls to be an explicit nil

### UnsetLinkUrls
`func (o *WorkItemLinkExtractionApiModel) UnsetLinkUrls()`

UnsetLinkUrls ensures that no value is present for LinkUrls, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


