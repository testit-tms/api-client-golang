# TestPlanTestPointsAnalyticsApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountGroupByStatus** | [**[]TestPlanTestPointsStatusGroupApiResult**](TestPlanTestPointsStatusGroupApiResult.md) |  | 
**CountGroupByTesterAndStatus** | [**[]TestPlanTestPointsTesterAndStatusGroupApiResult**](TestPlanTestPointsTesterAndStatusGroupApiResult.md) |  | 
**SumGroupByTester** | [**[]TestPlanTestPointsTesterGroupApiResult**](TestPlanTestPointsTesterGroupApiResult.md) |  | 
**CountGroupByTester** | [**[]TestPlanTestPointsTesterGroupApiResult**](TestPlanTestPointsTesterGroupApiResult.md) |  | 
**CountGroupByStatusType** | [**[]TestPlanTestPointsStatusTypeGroupApiResult**](TestPlanTestPointsStatusTypeGroupApiResult.md) |  | 
**CountGroupByTesterAndStatusType** | [**[]TestPlanTestPointsTesterAndStatusTypeGroupApiResult**](TestPlanTestPointsTesterAndStatusTypeGroupApiResult.md) |  | 

## Methods

### NewTestPlanTestPointsAnalyticsApiResult

`func NewTestPlanTestPointsAnalyticsApiResult(countGroupByStatus []TestPlanTestPointsStatusGroupApiResult, countGroupByTesterAndStatus []TestPlanTestPointsTesterAndStatusGroupApiResult, sumGroupByTester []TestPlanTestPointsTesterGroupApiResult, countGroupByTester []TestPlanTestPointsTesterGroupApiResult, countGroupByStatusType []TestPlanTestPointsStatusTypeGroupApiResult, countGroupByTesterAndStatusType []TestPlanTestPointsTesterAndStatusTypeGroupApiResult, ) *TestPlanTestPointsAnalyticsApiResult`

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


### GetCountGroupByStatusType

`func (o *TestPlanTestPointsAnalyticsApiResult) GetCountGroupByStatusType() []TestPlanTestPointsStatusTypeGroupApiResult`

GetCountGroupByStatusType returns the CountGroupByStatusType field if non-nil, zero value otherwise.

### GetCountGroupByStatusTypeOk

`func (o *TestPlanTestPointsAnalyticsApiResult) GetCountGroupByStatusTypeOk() (*[]TestPlanTestPointsStatusTypeGroupApiResult, bool)`

GetCountGroupByStatusTypeOk returns a tuple with the CountGroupByStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountGroupByStatusType

`func (o *TestPlanTestPointsAnalyticsApiResult) SetCountGroupByStatusType(v []TestPlanTestPointsStatusTypeGroupApiResult)`

SetCountGroupByStatusType sets CountGroupByStatusType field to given value.


### GetCountGroupByTesterAndStatusType

`func (o *TestPlanTestPointsAnalyticsApiResult) GetCountGroupByTesterAndStatusType() []TestPlanTestPointsTesterAndStatusTypeGroupApiResult`

GetCountGroupByTesterAndStatusType returns the CountGroupByTesterAndStatusType field if non-nil, zero value otherwise.

### GetCountGroupByTesterAndStatusTypeOk

`func (o *TestPlanTestPointsAnalyticsApiResult) GetCountGroupByTesterAndStatusTypeOk() (*[]TestPlanTestPointsTesterAndStatusTypeGroupApiResult, bool)`

GetCountGroupByTesterAndStatusTypeOk returns a tuple with the CountGroupByTesterAndStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountGroupByTesterAndStatusType

`func (o *TestPlanTestPointsAnalyticsApiResult) SetCountGroupByTesterAndStatusType(v []TestPlanTestPointsTesterAndStatusTypeGroupApiResult)`

SetCountGroupByTesterAndStatusType sets CountGroupByTesterAndStatusType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


