# TestPlanGroupByTesterAndStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **NullableString** |  | [optional] 
**Status** | **string** |  | 
**Value** | **int64** |  | 

## Methods

### NewTestPlanGroupByTesterAndStatus

`func NewTestPlanGroupByTesterAndStatus(status string, value int64, ) *TestPlanGroupByTesterAndStatus`

NewTestPlanGroupByTesterAndStatus instantiates a new TestPlanGroupByTesterAndStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPlanGroupByTesterAndStatusWithDefaults

`func NewTestPlanGroupByTesterAndStatusWithDefaults() *TestPlanGroupByTesterAndStatus`

NewTestPlanGroupByTesterAndStatusWithDefaults instantiates a new TestPlanGroupByTesterAndStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *TestPlanGroupByTesterAndStatus) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TestPlanGroupByTesterAndStatus) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TestPlanGroupByTesterAndStatus) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *TestPlanGroupByTesterAndStatus) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *TestPlanGroupByTesterAndStatus) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *TestPlanGroupByTesterAndStatus) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetStatus

`func (o *TestPlanGroupByTesterAndStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestPlanGroupByTesterAndStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestPlanGroupByTesterAndStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetValue

`func (o *TestPlanGroupByTesterAndStatus) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TestPlanGroupByTesterAndStatus) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TestPlanGroupByTesterAndStatus) SetValue(v int64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


