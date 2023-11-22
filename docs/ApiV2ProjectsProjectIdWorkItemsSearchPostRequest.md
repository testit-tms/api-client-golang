# ApiV2ProjectsProjectIdWorkItemsSearchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**WorkItemSelectModelFilter**](WorkItemSelectModelFilter.md) |  | 
**ExtractionModel** | Pointer to [**NullableWorkItemLocalSelectModelExtractionModel**](WorkItemLocalSelectModelExtractionModel.md) |  | [optional] 

## Methods

### NewApiV2ProjectsProjectIdWorkItemsSearchPostRequest

`func NewApiV2ProjectsProjectIdWorkItemsSearchPostRequest(filter WorkItemSelectModelFilter, ) *ApiV2ProjectsProjectIdWorkItemsSearchPostRequest`

NewApiV2ProjectsProjectIdWorkItemsSearchPostRequest instantiates a new ApiV2ProjectsProjectIdWorkItemsSearchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2ProjectsProjectIdWorkItemsSearchPostRequestWithDefaults

`func NewApiV2ProjectsProjectIdWorkItemsSearchPostRequestWithDefaults() *ApiV2ProjectsProjectIdWorkItemsSearchPostRequest`

NewApiV2ProjectsProjectIdWorkItemsSearchPostRequestWithDefaults instantiates a new ApiV2ProjectsProjectIdWorkItemsSearchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchPostRequest) GetFilter() WorkItemSelectModelFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchPostRequest) GetFilterOk() (*WorkItemSelectModelFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchPostRequest) SetFilter(v WorkItemSelectModelFilter)`

SetFilter sets Filter field to given value.


### GetExtractionModel

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchPostRequest) GetExtractionModel() WorkItemLocalSelectModelExtractionModel`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchPostRequest) GetExtractionModelOk() (*WorkItemLocalSelectModelExtractionModel, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchPostRequest) SetExtractionModel(v WorkItemLocalSelectModelExtractionModel)`

SetExtractionModel sets ExtractionModel field to given value.

### HasExtractionModel

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchPostRequest) HasExtractionModel() bool`

HasExtractionModel returns a boolean if a field has been set.

### SetExtractionModelNil

`func (o *ApiV2ProjectsProjectIdWorkItemsSearchPostRequest) SetExtractionModelNil(b bool)`

 SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil

### UnsetExtractionModel
`func (o *ApiV2ProjectsProjectIdWorkItemsSearchPostRequest) UnsetExtractionModel()`

UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


