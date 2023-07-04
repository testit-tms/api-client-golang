# TestRunShortGetModelStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statuses** | Pointer to [**TestResultsStatisticsGetModelStatuses**](TestResultsStatisticsGetModelStatuses.md) |  | [optional] 
**FailureCategories** | Pointer to [**TestResultsStatisticsGetModelFailureCategories**](TestResultsStatisticsGetModelFailureCategories.md) |  | [optional] 

## Methods

### NewTestRunShortGetModelStatistics

`func NewTestRunShortGetModelStatistics() *TestRunShortGetModelStatistics`

NewTestRunShortGetModelStatistics instantiates a new TestRunShortGetModelStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunShortGetModelStatisticsWithDefaults

`func NewTestRunShortGetModelStatisticsWithDefaults() *TestRunShortGetModelStatistics`

NewTestRunShortGetModelStatisticsWithDefaults instantiates a new TestRunShortGetModelStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatuses

`func (o *TestRunShortGetModelStatistics) GetStatuses() TestResultsStatisticsGetModelStatuses`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *TestRunShortGetModelStatistics) GetStatusesOk() (*TestResultsStatisticsGetModelStatuses, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *TestRunShortGetModelStatistics) SetStatuses(v TestResultsStatisticsGetModelStatuses)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *TestRunShortGetModelStatistics) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetFailureCategories

`func (o *TestRunShortGetModelStatistics) GetFailureCategories() TestResultsStatisticsGetModelFailureCategories`

GetFailureCategories returns the FailureCategories field if non-nil, zero value otherwise.

### GetFailureCategoriesOk

`func (o *TestRunShortGetModelStatistics) GetFailureCategoriesOk() (*TestResultsStatisticsGetModelFailureCategories, bool)`

GetFailureCategoriesOk returns a tuple with the FailureCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategories

`func (o *TestRunShortGetModelStatistics) SetFailureCategories(v TestResultsStatisticsGetModelFailureCategories)`

SetFailureCategories sets FailureCategories field to given value.

### HasFailureCategories

`func (o *TestRunShortGetModelStatistics) HasFailureCategories() bool`

HasFailureCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


