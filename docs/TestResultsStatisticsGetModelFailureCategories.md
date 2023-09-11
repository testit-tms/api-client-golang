# TestResultsStatisticsGetModelFailureCategories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoAnalytics** | **int32** | Number of test results which outcomes were not analyzed | 
**NoDefect** | **int32** | Number of test results which outcomes were not caused by any defect | 
**InfrastructureDefect** | **int32** | Number of test results which outcomes were caused by some infrastructure defect | 
**ProductDefect** | **int32** | Number of test results which outcomes were caused by some tested product defect | 
**TestDefect** | **int32** | Number of test results which outcomes were caused by test itself | 

## Methods

### NewTestResultsStatisticsGetModelFailureCategories

`func NewTestResultsStatisticsGetModelFailureCategories(noAnalytics int32, noDefect int32, infrastructureDefect int32, productDefect int32, testDefect int32, ) *TestResultsStatisticsGetModelFailureCategories`

NewTestResultsStatisticsGetModelFailureCategories instantiates a new TestResultsStatisticsGetModelFailureCategories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultsStatisticsGetModelFailureCategoriesWithDefaults

`func NewTestResultsStatisticsGetModelFailureCategoriesWithDefaults() *TestResultsStatisticsGetModelFailureCategories`

NewTestResultsStatisticsGetModelFailureCategoriesWithDefaults instantiates a new TestResultsStatisticsGetModelFailureCategories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoAnalytics

`func (o *TestResultsStatisticsGetModelFailureCategories) GetNoAnalytics() int32`

GetNoAnalytics returns the NoAnalytics field if non-nil, zero value otherwise.

### GetNoAnalyticsOk

`func (o *TestResultsStatisticsGetModelFailureCategories) GetNoAnalyticsOk() (*int32, bool)`

GetNoAnalyticsOk returns a tuple with the NoAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAnalytics

`func (o *TestResultsStatisticsGetModelFailureCategories) SetNoAnalytics(v int32)`

SetNoAnalytics sets NoAnalytics field to given value.


### GetNoDefect

`func (o *TestResultsStatisticsGetModelFailureCategories) GetNoDefect() int32`

GetNoDefect returns the NoDefect field if non-nil, zero value otherwise.

### GetNoDefectOk

`func (o *TestResultsStatisticsGetModelFailureCategories) GetNoDefectOk() (*int32, bool)`

GetNoDefectOk returns a tuple with the NoDefect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDefect

`func (o *TestResultsStatisticsGetModelFailureCategories) SetNoDefect(v int32)`

SetNoDefect sets NoDefect field to given value.


### GetInfrastructureDefect

`func (o *TestResultsStatisticsGetModelFailureCategories) GetInfrastructureDefect() int32`

GetInfrastructureDefect returns the InfrastructureDefect field if non-nil, zero value otherwise.

### GetInfrastructureDefectOk

`func (o *TestResultsStatisticsGetModelFailureCategories) GetInfrastructureDefectOk() (*int32, bool)`

GetInfrastructureDefectOk returns a tuple with the InfrastructureDefect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDefect

`func (o *TestResultsStatisticsGetModelFailureCategories) SetInfrastructureDefect(v int32)`

SetInfrastructureDefect sets InfrastructureDefect field to given value.


### GetProductDefect

`func (o *TestResultsStatisticsGetModelFailureCategories) GetProductDefect() int32`

GetProductDefect returns the ProductDefect field if non-nil, zero value otherwise.

### GetProductDefectOk

`func (o *TestResultsStatisticsGetModelFailureCategories) GetProductDefectOk() (*int32, bool)`

GetProductDefectOk returns a tuple with the ProductDefect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDefect

`func (o *TestResultsStatisticsGetModelFailureCategories) SetProductDefect(v int32)`

SetProductDefect sets ProductDefect field to given value.


### GetTestDefect

`func (o *TestResultsStatisticsGetModelFailureCategories) GetTestDefect() int32`

GetTestDefect returns the TestDefect field if non-nil, zero value otherwise.

### GetTestDefectOk

`func (o *TestResultsStatisticsGetModelFailureCategories) GetTestDefectOk() (*int32, bool)`

GetTestDefectOk returns a tuple with the TestDefect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestDefect

`func (o *TestResultsStatisticsGetModelFailureCategories) SetTestDefect(v int32)`

SetTestDefect sets TestDefect field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


