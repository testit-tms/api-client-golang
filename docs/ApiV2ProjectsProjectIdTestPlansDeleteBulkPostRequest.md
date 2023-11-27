# ApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**ApiV2ProjectsProjectIdTestPlansSearchPostRequest**](ApiV2ProjectsProjectIdTestPlansSearchPostRequest.md) |  | 
**ExtractionModel** | Pointer to [**NullableTestPlanExtractionModel**](TestPlanExtractionModel.md) |  | [optional] 

## Methods

### NewApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequest

`func NewApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequest(filter ApiV2ProjectsProjectIdTestPlansSearchPostRequest, ) *ApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequest`

NewApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequest instantiates a new ApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequestWithDefaults

`func NewApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequestWithDefaults() *ApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequest`

NewApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequestWithDefaults instantiates a new ApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequest) GetFilter() ApiV2ProjectsProjectIdTestPlansSearchPostRequest`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequest) GetFilterOk() (*ApiV2ProjectsProjectIdTestPlansSearchPostRequest, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequest) SetFilter(v ApiV2ProjectsProjectIdTestPlansSearchPostRequest)`

SetFilter sets Filter field to given value.


### GetExtractionModel

`func (o *ApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequest) GetExtractionModel() TestPlanExtractionModel`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *ApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequest) GetExtractionModelOk() (*TestPlanExtractionModel, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *ApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequest) SetExtractionModel(v TestPlanExtractionModel)`

SetExtractionModel sets ExtractionModel field to given value.

### HasExtractionModel

`func (o *ApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequest) HasExtractionModel() bool`

HasExtractionModel returns a boolean if a field has been set.

### SetExtractionModelNil

`func (o *ApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequest) SetExtractionModelNil(b bool)`

 SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil

### UnsetExtractionModel
`func (o *ApiV2ProjectsProjectIdTestPlansDeleteBulkPostRequest) UnsetExtractionModel()`

UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


