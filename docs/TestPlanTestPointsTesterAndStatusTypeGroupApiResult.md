# TestPlanTestPointsTesterAndStatusTypeGroupApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **NullableString** |  | 
**StatusType** | [**TestStatusApiType**](TestStatusApiType.md) | Collection of possible status types | 
**Value** | **int64** |  | 

## Methods

### NewTestPlanTestPointsTesterAndStatusTypeGroupApiResult

`func NewTestPlanTestPointsTesterAndStatusTypeGroupApiResult(userId NullableString, statusType TestStatusApiType, value int64, ) *TestPlanTestPointsTesterAndStatusTypeGroupApiResult`

NewTestPlanTestPointsTesterAndStatusTypeGroupApiResult instantiates a new TestPlanTestPointsTesterAndStatusTypeGroupApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanTestPointsTesterAndStatusTypeGroupApiResultWithDefaults

`func NewTestPlanTestPointsTesterAndStatusTypeGroupApiResultWithDefaults() *TestPlanTestPointsTesterAndStatusTypeGroupApiResult`

NewTestPlanTestPointsTesterAndStatusTypeGroupApiResultWithDefaults instantiates a new TestPlanTestPointsTesterAndStatusTypeGroupApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *TestPlanTestPointsTesterAndStatusTypeGroupApiResult) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TestPlanTestPointsTesterAndStatusTypeGroupApiResult) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TestPlanTestPointsTesterAndStatusTypeGroupApiResult) SetUserId(v string)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *TestPlanTestPointsTesterAndStatusTypeGroupApiResult) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *TestPlanTestPointsTesterAndStatusTypeGroupApiResult) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetStatusType

`func (o *TestPlanTestPointsTesterAndStatusTypeGroupApiResult) GetStatusType() TestStatusApiType`

GetStatusType returns the StatusType field if non-nil, zero value otherwise.

### GetStatusTypeOk

`func (o *TestPlanTestPointsTesterAndStatusTypeGroupApiResult) GetStatusTypeOk() (*TestStatusApiType, bool)`

GetStatusTypeOk returns a tuple with the StatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusType

`func (o *TestPlanTestPointsTesterAndStatusTypeGroupApiResult) SetStatusType(v TestStatusApiType)`

SetStatusType sets StatusType field to given value.


### GetValue

`func (o *TestPlanTestPointsTesterAndStatusTypeGroupApiResult) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TestPlanTestPointsTesterAndStatusTypeGroupApiResult) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TestPlanTestPointsTesterAndStatusTypeGroupApiResult) SetValue(v int64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


