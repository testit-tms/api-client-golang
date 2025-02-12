# AutoTestFlakyBulkApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoTestSelect** | [**AutoTestSelectApiModel**](AutoTestSelectApiModel.md) |  | 
**Value** | **bool** | Are autotests flaky | 

## Methods

### NewAutoTestFlakyBulkApiModel

`func NewAutoTestFlakyBulkApiModel(autoTestSelect AutoTestSelectApiModel, value bool, ) *AutoTestFlakyBulkApiModel`

NewAutoTestFlakyBulkApiModel instantiates a new AutoTestFlakyBulkApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestFlakyBulkApiModelWithDefaults

`func NewAutoTestFlakyBulkApiModelWithDefaults() *AutoTestFlakyBulkApiModel`

NewAutoTestFlakyBulkApiModelWithDefaults instantiates a new AutoTestFlakyBulkApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoTestSelect

`func (o *AutoTestFlakyBulkApiModel) GetAutoTestSelect() AutoTestSelectApiModel`

GetAutoTestSelect returns the AutoTestSelect field if non-nil, zero value otherwise.

### GetAutoTestSelectOk

`func (o *AutoTestFlakyBulkApiModel) GetAutoTestSelectOk() (*AutoTestSelectApiModel, bool)`

GetAutoTestSelectOk returns a tuple with the AutoTestSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTestSelect

`func (o *AutoTestFlakyBulkApiModel) SetAutoTestSelect(v AutoTestSelectApiModel)`

SetAutoTestSelect sets AutoTestSelect field to given value.


### GetValue

`func (o *AutoTestFlakyBulkApiModel) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AutoTestFlakyBulkApiModel) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AutoTestFlakyBulkApiModel) SetValue(v bool)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


