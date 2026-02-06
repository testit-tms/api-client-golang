# TestPlanTestPointsStatusTypeGroupApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusType** | [**TestStatusApiType**](TestStatusApiType.md) | Collection of possible status types | 
**Statuses** | [**[]TestPlanTestPointsStatusCodeGroupApiResult**](TestPlanTestPointsStatusCodeGroupApiResult.md) |  | 

## Methods

### NewTestPlanTestPointsStatusTypeGroupApiResult

`func NewTestPlanTestPointsStatusTypeGroupApiResult(statusType TestStatusApiType, statuses []TestPlanTestPointsStatusCodeGroupApiResult, ) *TestPlanTestPointsStatusTypeGroupApiResult`

NewTestPlanTestPointsStatusTypeGroupApiResult instantiates a new TestPlanTestPointsStatusTypeGroupApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanTestPointsStatusTypeGroupApiResultWithDefaults

`func NewTestPlanTestPointsStatusTypeGroupApiResultWithDefaults() *TestPlanTestPointsStatusTypeGroupApiResult`

NewTestPlanTestPointsStatusTypeGroupApiResultWithDefaults instantiates a new TestPlanTestPointsStatusTypeGroupApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusType

`func (o *TestPlanTestPointsStatusTypeGroupApiResult) GetStatusType() TestStatusApiType`

GetStatusType returns the StatusType field if non-nil, zero value otherwise.

### GetStatusTypeOk

`func (o *TestPlanTestPointsStatusTypeGroupApiResult) GetStatusTypeOk() (*TestStatusApiType, bool)`

GetStatusTypeOk returns a tuple with the StatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusType

`func (o *TestPlanTestPointsStatusTypeGroupApiResult) SetStatusType(v TestStatusApiType)`

SetStatusType sets StatusType field to given value.


### GetStatuses

`func (o *TestPlanTestPointsStatusTypeGroupApiResult) GetStatuses() []TestPlanTestPointsStatusCodeGroupApiResult`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *TestPlanTestPointsStatusTypeGroupApiResult) GetStatusesOk() (*[]TestPlanTestPointsStatusCodeGroupApiResult, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *TestPlanTestPointsStatusTypeGroupApiResult) SetStatuses(v []TestPlanTestPointsStatusCodeGroupApiResult)`

SetStatuses sets Statuses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


