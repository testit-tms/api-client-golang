# TestPlanTestPointsAnalyticsApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountGroupByStatus** | [**[]TestPlanTestPointsStatusGroupApiResult**](TestPlanTestPointsStatusGroupApiResult.md) |  | 
**SumGroupByTester** | [**[]TestPlanTestPointsTesterGroupApiResult**](TestPlanTestPointsTesterGroupApiResult.md) |  | 
**CountGroupByTester** | [**[]TestPlanTestPointsTesterGroupApiResult**](TestPlanTestPointsTesterGroupApiResult.md) |  | 
**CountGroupByTesterAndStatus** | [**[]TestPlanTestPointsTesterAndStatusGroupApiResult**](TestPlanTestPointsTesterAndStatusGroupApiResult.md) |  | 

## Methods

### NewTestPlanTestPointsAnalyticsApiResult

`func NewTestPlanTestPointsAnalyticsApiResult(countGroupByStatus []TestPlanTestPointsStatusGroupApiResult, sumGroupByTester []TestPlanTestPointsTesterGroupApiResult, countGroupByTester []TestPlanTestPointsTesterGroupApiResult, countGroupByTesterAndStatus []TestPlanTestPointsTesterAndStatusGroupApiResult, ) *TestPlanTestPointsAnalyticsApiResult`

NewTestPlanTestPointsAnalyticsApiResult instantiates a new TestPlanTestPointsAnalyticsApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanTestPointsAnalyticsApiResultWithDefaults

`func NewTestPlanTestPointsAnalyticsApiResultWithDefaults() *TestPlanTestPointsAnalyticsApiResult`

NewTestPlanTestPointsAnalyticsApiResultWithDefaults instantiates a new TestPlanTestPointsAnalyticsApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountGroupByStatus

`func (o *TestPlanTestPointsAnalyticsApiResult) GetCountGroupByStatus() []TestPlanTestPointsStatusGroupApiResult`

GetCountGroupByStatus returns the CountGroupByStatus field if non-nil, zero value otherwise.

### GetCountGroupByStatusOk

`func (o *TestPlanTestPointsAnalyticsApiResult) GetCountGroupByStatusOk() (*[]TestPlanTestPointsStatusGroupApiResult, bool)`

GetCountGroupByStatusOk returns a tuple with the CountGroupByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountGroupByStatus

`func (o *TestPlanTestPointsAnalyticsApiResult) SetCountGroupByStatus(v []TestPlanTestPointsStatusGroupApiResult)`

SetCountGroupByStatus sets CountGroupByStatus field to given value.


### GetSumGroupByTester

`func (o *TestPlanTestPointsAnalyticsApiResult) GetSumGroupByTester() []TestPlanTestPointsTesterGroupApiResult`

GetSumGroupByTester returns the SumGroupByTester field if non-nil, zero value otherwise.

### GetSumGroupByTesterOk

`func (o *TestPlanTestPointsAnalyticsApiResult) GetSumGroupByTesterOk() (*[]TestPlanTestPointsTesterGroupApiResult, bool)`

GetSumGroupByTesterOk returns a tuple with the SumGroupByTester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumGroupByTester

`func (o *TestPlanTestPointsAnalyticsApiResult) SetSumGroupByTester(v []TestPlanTestPointsTesterGroupApiResult)`

SetSumGroupByTester sets SumGroupByTester field to given value.


### GetCountGroupByTester

`func (o *TestPlanTestPointsAnalyticsApiResult) GetCountGroupByTester() []TestPlanTestPointsTesterGroupApiResult`

GetCountGroupByTester returns the CountGroupByTester field if non-nil, zero value otherwise.

### GetCountGroupByTesterOk

`func (o *TestPlanTestPointsAnalyticsApiResult) GetCountGroupByTesterOk() (*[]TestPlanTestPointsTesterGroupApiResult, bool)`

GetCountGroupByTesterOk returns a tuple with the CountGroupByTester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountGroupByTester

`func (o *TestPlanTestPointsAnalyticsApiResult) SetCountGroupByTester(v []TestPlanTestPointsTesterGroupApiResult)`

SetCountGroupByTester sets CountGroupByTester field to given value.


### GetCountGroupByTesterAndStatus

`func (o *TestPlanTestPointsAnalyticsApiResult) GetCountGroupByTesterAndStatus() []TestPlanTestPointsTesterAndStatusGroupApiResult`

GetCountGroupByTesterAndStatus returns the CountGroupByTesterAndStatus field if non-nil, zero value otherwise.

### GetCountGroupByTesterAndStatusOk

`func (o *TestPlanTestPointsAnalyticsApiResult) GetCountGroupByTesterAndStatusOk() (*[]TestPlanTestPointsTesterAndStatusGroupApiResult, bool)`

GetCountGroupByTesterAndStatusOk returns a tuple with the CountGroupByTesterAndStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountGroupByTesterAndStatus

`func (o *TestPlanTestPointsAnalyticsApiResult) SetCountGroupByTesterAndStatus(v []TestPlanTestPointsTesterAndStatusGroupApiResult)`

SetCountGroupByTesterAndStatus sets CountGroupByTesterAndStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


