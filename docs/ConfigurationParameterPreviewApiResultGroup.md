# ConfigurationParameterPreviewApiResultGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**GroupKey**](GroupKey.md) |  | 
**Data** | [**[]ConfigurationParameterPreviewApiResult**](ConfigurationParameterPreviewApiResult.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewConfigurationParameterPreviewApiResultGroup

`func NewConfigurationParameterPreviewApiResultGroup(key GroupKey, data []ConfigurationParameterPreviewApiResult, totalCount int32, ) *ConfigurationParameterPreviewApiResultGroup`

NewConfigurationParameterPreviewApiResultGroup instantiates a new ConfigurationParameterPreviewApiResultGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationParameterPreviewApiResultGroupWithDefaults

`func NewConfigurationParameterPreviewApiResultGroupWithDefaults() *ConfigurationParameterPreviewApiResultGroup`

NewConfigurationParameterPreviewApiResultGroupWithDefaults instantiates a new ConfigurationParameterPreviewApiResultGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ConfigurationParameterPreviewApiResultGroup) GetKey() GroupKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ConfigurationParameterPreviewApiResultGroup) GetKeyOk() (*GroupKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ConfigurationParameterPreviewApiResultGroup) SetKey(v GroupKey)`

SetKey sets Key field to given value.


### GetData

`func (o *ConfigurationParameterPreviewApiResultGroup) GetData() []ConfigurationParameterPreviewApiResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConfigurationParameterPreviewApiResultGroup) GetDataOk() (*[]ConfigurationParameterPreviewApiResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConfigurationParameterPreviewApiResultGroup) SetData(v []ConfigurationParameterPreviewApiResult)`

SetData sets Data field to given value.


### GetTotalCount

`func (o *ConfigurationParameterPreviewApiResultGroup) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ConfigurationParameterPreviewApiResultGroup) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ConfigurationParameterPreviewApiResultGroup) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


