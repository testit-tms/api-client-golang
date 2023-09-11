# ApiV2ProjectsRestoreBulkPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**NullableProjectsFilterModel**](ProjectsFilterModel.md) |  | [optional] 
**ExtractionModel** | Pointer to [**NullableProjectExtractionModel**](ProjectExtractionModel.md) |  | [optional] 

## Methods

### NewApiV2ProjectsRestoreBulkPostRequest

`func NewApiV2ProjectsRestoreBulkPostRequest() *ApiV2ProjectsRestoreBulkPostRequest`

NewApiV2ProjectsRestoreBulkPostRequest instantiates a new ApiV2ProjectsRestoreBulkPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2ProjectsRestoreBulkPostRequestWithDefaults

`func NewApiV2ProjectsRestoreBulkPostRequestWithDefaults() *ApiV2ProjectsRestoreBulkPostRequest`

NewApiV2ProjectsRestoreBulkPostRequestWithDefaults instantiates a new ApiV2ProjectsRestoreBulkPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ApiV2ProjectsRestoreBulkPostRequest) GetFilter() ProjectsFilterModel`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ApiV2ProjectsRestoreBulkPostRequest) GetFilterOk() (*ProjectsFilterModel, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ApiV2ProjectsRestoreBulkPostRequest) SetFilter(v ProjectsFilterModel)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ApiV2ProjectsRestoreBulkPostRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *ApiV2ProjectsRestoreBulkPostRequest) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *ApiV2ProjectsRestoreBulkPostRequest) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetExtractionModel

`func (o *ApiV2ProjectsRestoreBulkPostRequest) GetExtractionModel() ProjectExtractionModel`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *ApiV2ProjectsRestoreBulkPostRequest) GetExtractionModelOk() (*ProjectExtractionModel, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *ApiV2ProjectsRestoreBulkPostRequest) SetExtractionModel(v ProjectExtractionModel)`

SetExtractionModel sets ExtractionModel field to given value.

### HasExtractionModel

`func (o *ApiV2ProjectsRestoreBulkPostRequest) HasExtractionModel() bool`

HasExtractionModel returns a boolean if a field has been set.

### SetExtractionModelNil

`func (o *ApiV2ProjectsRestoreBulkPostRequest) SetExtractionModelNil(b bool)`

 SetExtractionModelNil sets the value for ExtractionModel to be an explicit nil

### UnsetExtractionModel
`func (o *ApiV2ProjectsRestoreBulkPostRequest) UnsetExtractionModel()`

UnsetExtractionModel ensures that no value is present for ExtractionModel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


