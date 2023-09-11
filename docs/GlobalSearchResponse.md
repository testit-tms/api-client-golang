# GlobalSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]GlobalSearchItemResult**](GlobalSearchItemResult.md) |  | 
**MoreResultsAvailable** | **bool** |  | 
**AvailableResourceTypes** | **[]string** |  | 

## Methods

### NewGlobalSearchResponse

`func NewGlobalSearchResponse(results []GlobalSearchItemResult, moreResultsAvailable bool, availableResourceTypes []string, ) *GlobalSearchResponse`

NewGlobalSearchResponse instantiates a new GlobalSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalSearchResponseWithDefaults

`func NewGlobalSearchResponseWithDefaults() *GlobalSearchResponse`

NewGlobalSearchResponseWithDefaults instantiates a new GlobalSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *GlobalSearchResponse) GetResults() []GlobalSearchItemResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GlobalSearchResponse) GetResultsOk() (*[]GlobalSearchItemResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GlobalSearchResponse) SetResults(v []GlobalSearchItemResult)`

SetResults sets Results field to given value.


### GetMoreResultsAvailable

`func (o *GlobalSearchResponse) GetMoreResultsAvailable() bool`

GetMoreResultsAvailable returns the MoreResultsAvailable field if non-nil, zero value otherwise.

### GetMoreResultsAvailableOk

`func (o *GlobalSearchResponse) GetMoreResultsAvailableOk() (*bool, bool)`

GetMoreResultsAvailableOk returns a tuple with the MoreResultsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreResultsAvailable

`func (o *GlobalSearchResponse) SetMoreResultsAvailable(v bool)`

SetMoreResultsAvailable sets MoreResultsAvailable field to given value.


### GetAvailableResourceTypes

`func (o *GlobalSearchResponse) GetAvailableResourceTypes() []string`

GetAvailableResourceTypes returns the AvailableResourceTypes field if non-nil, zero value otherwise.

### GetAvailableResourceTypesOk

`func (o *GlobalSearchResponse) GetAvailableResourceTypesOk() (*[]string, bool)`

GetAvailableResourceTypesOk returns a tuple with the AvailableResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableResourceTypes

`func (o *GlobalSearchResponse) SetAvailableResourceTypes(v []string)`

SetAvailableResourceTypes sets AvailableResourceTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


