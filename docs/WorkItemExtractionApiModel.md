# WorkItemExtractionApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectIds** | Pointer to [**NullableGuidExtractionModel**](GuidExtractionModel.md) | Extraction parameters for projects | [optional] 
**Ids** | Pointer to [**NullableGuidExtractionModel**](GuidExtractionModel.md) | Extraction parameters for work items | [optional] 
**SectionIds** | Pointer to [**NullableGuidExtractionModel**](GuidExtractionModel.md) | Extraction parameters for sections | [optional] 

## Methods

### NewWorkItemExtractionApiModel

`func NewWorkItemExtractionApiModel() *WorkItemExtractionApiModel`

NewWorkItemExtractionApiModel instantiates a new WorkItemExtractionApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkItemExtractionApiModelWithDefaults

`func NewWorkItemExtractionApiModelWithDefaults() *WorkItemExtractionApiModel`

NewWorkItemExtractionApiModelWithDefaults instantiates a new WorkItemExtractionApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectIds

`func (o *WorkItemExtractionApiModel) GetProjectIds() GuidExtractionModel`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *WorkItemExtractionApiModel) GetProjectIdsOk() (*GuidExtractionModel, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *WorkItemExtractionApiModel) SetProjectIds(v GuidExtractionModel)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *WorkItemExtractionApiModel) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *WorkItemExtractionApiModel) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *WorkItemExtractionApiModel) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetIds

`func (o *WorkItemExtractionApiModel) GetIds() GuidExtractionModel`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *WorkItemExtractionApiModel) GetIdsOk() (*GuidExtractionModel, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *WorkItemExtractionApiModel) SetIds(v GuidExtractionModel)`

SetIds sets Ids field to given value.

### HasIds

`func (o *WorkItemExtractionApiModel) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *WorkItemExtractionApiModel) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *WorkItemExtractionApiModel) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetSectionIds

`func (o *WorkItemExtractionApiModel) GetSectionIds() GuidExtractionModel`

GetSectionIds returns the SectionIds field if non-nil, zero value otherwise.

### GetSectionIdsOk

`func (o *WorkItemExtractionApiModel) GetSectionIdsOk() (*GuidExtractionModel, bool)`

GetSectionIdsOk returns a tuple with the SectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionIds

`func (o *WorkItemExtractionApiModel) SetSectionIds(v GuidExtractionModel)`

SetSectionIds sets SectionIds field to given value.

### HasSectionIds

`func (o *WorkItemExtractionApiModel) HasSectionIds() bool`

HasSectionIds returns a boolean if a field has been set.

### SetSectionIdsNil

`func (o *WorkItemExtractionApiModel) SetSectionIdsNil(b bool)`

 SetSectionIdsNil sets the value for SectionIds to be an explicit nil

### UnsetSectionIds
`func (o *WorkItemExtractionApiModel) UnsetSectionIds()`

UnsetSectionIds ensures that no value is present for SectionIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


