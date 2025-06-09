# TestPlanTestPointsGroupSearchApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TestPlanTestPointsGroupSearchItemApiResult**](TestPlanTestPointsGroupSearchItemApiResult.md) |  | 
**TotalCount** | **int32** |  | 
**StatusCounters** | [**TestPlanTestPointsSearchStatusCountersApiResult**](TestPlanTestPointsSearchStatusCountersApiResult.md) |  | 

## Methods

### NewTestPlanTestPointsGroupSearchApiResult

`func NewTestPlanTestPointsGroupSearchApiResult(data []TestPlanTestPointsGroupSearchItemApiResult, totalCount int32, statusCounters TestPlanTestPointsSearchStatusCountersApiResult, ) *TestPlanTestPointsGroupSearchApiResult`

NewTestPlanTestPointsGroupSearchApiResult instantiates a new TestPlanTestPointsGroupSearchApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanTestPointsGroupSearchApiResultWithDefaults

`func NewTestPlanTestPointsGroupSearchApiResultWithDefaults() *TestPlanTestPointsGroupSearchApiResult`

NewTestPlanTestPointsGroupSearchApiResultWithDefaults instantiates a new TestPlanTestPointsGroupSearchApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TestPlanTestPointsGroupSearchApiResult) GetData() []TestPlanTestPointsGroupSearchItemApiResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TestPlanTestPointsGroupSearchApiResult) GetDataOk() (*[]TestPlanTestPointsGroupSearchItemApiResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TestPlanTestPointsGroupSearchApiResult) SetData(v []TestPlanTestPointsGroupSearchItemApiResult)`

SetData sets Data field to given value.


### GetTotalCount

`func (o *TestPlanTestPointsGroupSearchApiResult) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TestPlanTestPointsGroupSearchApiResult) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TestPlanTestPointsGroupSearchApiResult) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetStatusCounters

`func (o *TestPlanTestPointsGroupSearchApiResult) GetStatusCounters() TestPlanTestPointsSearchStatusCountersApiResult`

GetStatusCounters returns the StatusCounters field if non-nil, zero value otherwise.

### GetStatusCountersOk

`func (o *TestPlanTestPointsGroupSearchApiResult) GetStatusCountersOk() (*TestPlanTestPointsSearchStatusCountersApiResult, bool)`

GetStatusCountersOk returns a tuple with the StatusCounters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCounters

`func (o *TestPlanTestPointsGroupSearchApiResult) SetStatusCounters(v TestPlanTestPointsSearchStatusCountersApiResult)`

SetStatusCounters sets StatusCounters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


