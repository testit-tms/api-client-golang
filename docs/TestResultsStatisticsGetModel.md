# TestResultsStatisticsGetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statuses** | [**TestRunStatisticsStatusesGetModel**](TestRunStatisticsStatusesGetModel.md) | Test results counts aggregated by outcome | [readonly] 
**FailureCategories** | [**TestRunStatisticsErrorCategoriesGetModel**](TestRunStatisticsErrorCategoriesGetModel.md) | Test results counts aggregated by result failure categories | [readonly] 

## Methods

### NewTestResultsStatisticsGetModel

`func NewTestResultsStatisticsGetModel(statuses TestRunStatisticsStatusesGetModel, failureCategories TestRunStatisticsErrorCategoriesGetModel, ) *TestResultsStatisticsGetModel`

NewTestResultsStatisticsGetModel instantiates a new TestResultsStatisticsGetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultsStatisticsGetModelWithDefaults

`func NewTestResultsStatisticsGetModelWithDefaults() *TestResultsStatisticsGetModel`

NewTestResultsStatisticsGetModelWithDefaults instantiates a new TestResultsStatisticsGetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatuses

`func (o *TestResultsStatisticsGetModel) GetStatuses() TestRunStatisticsStatusesGetModel`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *TestResultsStatisticsGetModel) GetStatusesOk() (*TestRunStatisticsStatusesGetModel, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *TestResultsStatisticsGetModel) SetStatuses(v TestRunStatisticsStatusesGetModel)`

SetStatuses sets Statuses field to given value.


### GetFailureCategories

`func (o *TestResultsStatisticsGetModel) GetFailureCategories() TestRunStatisticsErrorCategoriesGetModel`

GetFailureCategories returns the FailureCategories field if non-nil, zero value otherwise.

### GetFailureCategoriesOk

`func (o *TestResultsStatisticsGetModel) GetFailureCategoriesOk() (*TestRunStatisticsErrorCategoriesGetModel, bool)`

GetFailureCategoriesOk returns a tuple with the FailureCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCategories

`func (o *TestResultsStatisticsGetModel) SetFailureCategories(v TestRunStatisticsErrorCategoriesGetModel)`

SetFailureCategories sets FailureCategories field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


