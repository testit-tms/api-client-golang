# AutoTestResultReasonGroupItemApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**NullableAutoTestResultReasonGroupApiResult**](AutoTestResultReasonGroupApiResult.md) | Group details | [optional] 
**Items** | [**[]AutoTestResultReasonsProjectApiResult**](AutoTestResultReasonsProjectApiResult.md) | Group data | 

## Methods

### NewAutoTestResultReasonGroupItemApiResult

`func NewAutoTestResultReasonGroupItemApiResult(items []AutoTestResultReasonsProjectApiResult, ) *AutoTestResultReasonGroupItemApiResult`

NewAutoTestResultReasonGroupItemApiResult instantiates a new AutoTestResultReasonGroupItemApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestResultReasonGroupItemApiResultWithDefaults

`func NewAutoTestResultReasonGroupItemApiResultWithDefaults() *AutoTestResultReasonGroupItemApiResult`

NewAutoTestResultReasonGroupItemApiResultWithDefaults instantiates a new AutoTestResultReasonGroupItemApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *AutoTestResultReasonGroupItemApiResult) GetGroup() AutoTestResultReasonGroupApiResult`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AutoTestResultReasonGroupItemApiResult) GetGroupOk() (*AutoTestResultReasonGroupApiResult, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AutoTestResultReasonGroupItemApiResult) SetGroup(v AutoTestResultReasonGroupApiResult)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AutoTestResultReasonGroupItemApiResult) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *AutoTestResultReasonGroupItemApiResult) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *AutoTestResultReasonGroupItemApiResult) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetItems

`func (o *AutoTestResultReasonGroupItemApiResult) GetItems() []AutoTestResultReasonsProjectApiResult`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AutoTestResultReasonGroupItemApiResult) GetItemsOk() (*[]AutoTestResultReasonsProjectApiResult, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AutoTestResultReasonGroupItemApiResult) SetItems(v []AutoTestResultReasonsProjectApiResult)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


