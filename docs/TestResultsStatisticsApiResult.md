# TestResultsStatisticsApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statuses** | [**TestResultsStatisticsStatusesApiResult**](TestResultsStatisticsStatusesApiResult.md) | Test results counts aggregated by outcome | [readonly] 
**FailureCategories** | [**TestResultsStatisticsFailureCategoriesApiResult**](TestResultsStatisticsFailureCategoriesApiResult.md) | Test results counts aggregated by result failure categories | [readonly] 

## Methods

### NewTestResultsStatisticsApiResult

`func NewTestResultsStatisticsApiResult(statuses TestResultsStatisticsStatusesApiResult, failureCategories TestResultsStatisticsFailureCategoriesApiResult, ) *TestResultsStatisticsApiResult`

NewTestResultsStatisticsApiResult instantiates a new TestResultsStatisticsApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultsStatisticsApiResultWithDefaults

`func NewTestResultsStatisticsApiResultWithDefaults() *TestResultsStatisticsApiResult`

NewTestResultsStatisticsApiResultWithDefaults instantiates a new TestResultsStatisticsApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatuses

`func (o *TestResultsStatisticsApiResult) GetStatuses() TestResultsStatisticsStatusesApiResult`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *TestResultsStatisticsApiResult) GetStatusesOk() (*TestResultsStatisticsStatusesApiResult, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *TestResultsStatisticsApiResult) SetStatuses(v TestResultsStatisticsStatusesApiResult)`

SetStatuses sets Statuses field to given value.


### GetFailureCategories

`func (o *TestResultsStatisticsApiResult) GetFailureCategories() TestResultsStatisticsFailureCategoriesApiResult`

GetFailureCategories returns the FailureCategories field if non-nil, zero value otherwise.

### GetFailureCategoriesOk

`func (o *TestResultsStatisticsApiResult) GetFailureCategoriesOk() (*TestResultsStatisticsFailureCategoriesApiResult, bool)`

GetFailureCategoriesOk returns a tuple with the FailureCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategories

`func (o *TestResultsStatisticsApiResult) SetFailureCategories(v TestResultsStatisticsFailureCategoriesApiResult)`

SetFailureCategories sets FailureCategories field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


