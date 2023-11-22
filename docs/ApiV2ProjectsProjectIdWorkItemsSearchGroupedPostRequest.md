# ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectModel** | Pointer to [**NullableWorkItemGroupGetModelSelectModel**](WorkItemGroupGetModelSelectModel.md) |  | [optional] 
**GroupType** | [**WorkItemGroupType**](WorkItemGroupType.md) |  | 
**CustomAttributeId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest

`func NewApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest(groupType WorkItemGroupType, ) *ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest`

NewApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest instantiates a new ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequestWithDefaults

`func NewApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequestWithDefaults() *ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest`

NewApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequestWithDefaults instantiates a new ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectModel

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) GetSelectModel() WorkItemGroupGetModelSelectModel`

GetSelectModel returns the SelectModel field if non-nil, zero value otherwise.

### GetSelectModelOk

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) GetSelectModelOk() (*WorkItemGroupGetModelSelectModel, bool)`

GetSelectModelOk returns a tuple with the SelectModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectModel

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) SetSelectModel(v WorkItemGroupGetModelSelectModel)`

SetSelectModel sets SelectModel field to given value.

### HasSelectModel

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) HasSelectModel() bool`

HasSelectModel returns a boolean if a field has been set.

### SetSelectModelNil

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) SetSelectModelNil(b bool)`

 SetSelectModelNil sets the value for SelectModel to be an explicit nil

### UnsetSelectModel
`func (o *ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) UnsetSelectModel()`

UnsetSelectModel ensures that no value is present for SelectModel, not even an explicit nil
### GetGroupType

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) GetGroupType() WorkItemGroupType`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) GetGroupTypeOk() (*WorkItemGroupType, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) SetGroupType(v WorkItemGroupType)`

SetGroupType sets GroupType field to given value.


### GetCustomAttributeId

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) GetCustomAttributeId() string`

GetCustomAttributeId returns the CustomAttributeId field if non-nil, zero value otherwise.

### GetCustomAttributeIdOk

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) GetCustomAttributeIdOk() (*string, bool)`

GetCustomAttributeIdOk returns a tuple with the CustomAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributeId

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) SetCustomAttributeId(v string)`

SetCustomAttributeId sets CustomAttributeId field to given value.

### HasCustomAttributeId

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) HasCustomAttributeId() bool`

HasCustomAttributeId returns a boolean if a field has been set.

### SetCustomAttributeIdNil

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) SetCustomAttributeIdNil(b bool)`

 SetCustomAttributeIdNil sets the value for CustomAttributeId to be an explicit nil

### UnsetCustomAttributeId
`func (o *ApiV2ProjectsProjectIdWorkItemsSearchGroupedPostRequest) UnsetCustomAttributeId()`

UnsetCustomAttributeId ensures that no value is present for CustomAttributeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


