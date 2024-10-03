# ApiV2TestRunsDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**ApiV2TestRunsSearchPostRequest**](ApiV2TestRunsSearchPostRequest.md) |  | 
**ExtractionModel** | [**TestRunSelectModelExtractionModel**](TestRunSelectModelExtractionModel.md) |  | 

## Methods

### NewApiV2TestRunsDeleteRequest

`func NewApiV2TestRunsDeleteRequest(filter ApiV2TestRunsSearchPostRequest, extractionModel TestRunSelectModelExtractionModel, ) *ApiV2TestRunsDeleteRequest`

NewApiV2TestRunsDeleteRequest instantiates a new ApiV2TestRunsDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2TestRunsDeleteRequestWithDefaults

`func NewApiV2TestRunsDeleteRequestWithDefaults() *ApiV2TestRunsDeleteRequest`

NewApiV2TestRunsDeleteRequestWithDefaults instantiates a new ApiV2TestRunsDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ApiV2TestRunsDeleteRequest) GetFilter() ApiV2TestRunsSearchPostRequest`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ApiV2TestRunsDeleteRequest) GetFilterOk() (*ApiV2TestRunsSearchPostRequest, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ApiV2TestRunsDeleteRequest) SetFilter(v ApiV2TestRunsSearchPostRequest)`

SetFilter sets Filter field to given value.


### GetExtractionModel

`func (o *ApiV2TestRunsDeleteRequest) GetExtractionModel() TestRunSelectModelExtractionModel`

GetExtractionModel returns the ExtractionModel field if non-nil, zero value otherwise.

### GetExtractionModelOk

`func (o *ApiV2TestRunsDeleteRequest) GetExtractionModelOk() (*TestRunSelectModelExtractionModel, bool)`

GetExtractionModelOk returns a tuple with the ExtractionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionModel

`func (o *ApiV2TestRunsDeleteRequest) SetExtractionModel(v TestRunSelectModelExtractionModel)`

SetExtractionModel sets ExtractionModel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


