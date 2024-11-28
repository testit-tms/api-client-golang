# TestResultsStatisticsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statuses** | [**TestResultsStatisticsStatuses**](TestResultsStatisticsStatuses.md) | Test results counts aggregated by outcome | [readonly] 
**FailureCategories** | [**TestResultsStatisticsFailureCategories**](TestResultsStatisticsFailureCategories.md) | Test results counts aggregated by result failure categories | [readonly] 

## Methods

### NewTestResultsStatisticsResponse

`func NewTestResultsStatisticsResponse(statuses TestResultsStatisticsStatuses, failureCategories TestResultsStatisticsFailureCategories, ) *TestResultsStatisticsResponse`

NewTestResultsStatisticsResponse instantiates a new TestResultsStatisticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultsStatisticsResponseWithDefaults

`func NewTestResultsStatisticsResponseWithDefaults() *TestResultsStatisticsResponse`

NewTestResultsStatisticsResponseWithDefaults instantiates a new TestResultsStatisticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatuses

`func (o *TestResultsStatisticsResponse) GetStatuses() TestResultsStatisticsStatuses`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *TestResultsStatisticsResponse) GetStatusesOk() (*TestResultsStatisticsStatuses, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *TestResultsStatisticsResponse) SetStatuses(v TestResultsStatisticsStatuses)`

SetStatuses sets Statuses field to given value.


### GetFailureCategories

`func (o *TestResultsStatisticsResponse) GetFailureCategories() TestResultsStatisticsFailureCategories`

GetFailureCategories returns the FailureCategories field if non-nil, zero value otherwise.

### GetFailureCategoriesOk

`func (o *TestResultsStatisticsResponse) GetFailureCategoriesOk() (*TestResultsStatisticsFailureCategories, bool)`

GetFailureCategoriesOk returns a tuple with the FailureCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategories

`func (o *TestResultsStatisticsResponse) SetFailureCategories(v TestResultsStatisticsFailureCategories)`

SetFailureCategories sets FailureCategories field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


