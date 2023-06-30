# ApiV2AutoTestsFlakyBulkPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutotestSelect** | Pointer to [**FlakyBulkModelAutotestSelect**](FlakyBulkModelAutotestSelect.md) |  | [optional] 
**Value** | **bool** | Are autotests flaky | 

## Methods

### NewApiV2AutoTestsFlakyBulkPostRequest

`func NewApiV2AutoTestsFlakyBulkPostRequest(value bool, ) *ApiV2AutoTestsFlakyBulkPostRequest`

NewApiV2AutoTestsFlakyBulkPostRequest instantiates a new ApiV2AutoTestsFlakyBulkPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2AutoTestsFlakyBulkPostRequestWithDefaults

`func NewApiV2AutoTestsFlakyBulkPostRequestWithDefaults() *ApiV2AutoTestsFlakyBulkPostRequest`

NewApiV2AutoTestsFlakyBulkPostRequestWithDefaults instantiates a new ApiV2AutoTestsFlakyBulkPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutotestSelect

`func (o *ApiV2AutoTestsFlakyBulkPostRequest) GetAutotestSelect() FlakyBulkModelAutotestSelect`

GetAutotestSelect returns the AutotestSelect field if non-nil, zero value otherwise.

### GetAutotestSelectOk

`func (o *ApiV2AutoTestsFlakyBulkPostRequest) GetAutotestSelectOk() (*FlakyBulkModelAutotestSelect, bool)`

GetAutotestSelectOk returns a tuple with the AutotestSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutotestSelect

`func (o *ApiV2AutoTestsFlakyBulkPostRequest) SetAutotestSelect(v FlakyBulkModelAutotestSelect)`

SetAutotestSelect sets AutotestSelect field to given value.

### HasAutotestSelect

`func (o *ApiV2AutoTestsFlakyBulkPostRequest) HasAutotestSelect() bool`

HasAutotestSelect returns a boolean if a field has been set.

### GetValue

`func (o *ApiV2AutoTestsFlakyBulkPostRequest) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApiV2AutoTestsFlakyBulkPostRequest) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApiV2AutoTestsFlakyBulkPostRequest) SetValue(v bool)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


