# ConfigurationParameterPreviewApiResultGroupedReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | [**[]ConfigurationParameterPreviewApiResultGroup**](ConfigurationParameterPreviewApiResultGroup.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewConfigurationParameterPreviewApiResultGroupedReply

`func NewConfigurationParameterPreviewApiResultGroupedReply(groups []ConfigurationParameterPreviewApiResultGroup, totalCount int32, ) *ConfigurationParameterPreviewApiResultGroupedReply`

NewConfigurationParameterPreviewApiResultGroupedReply instantiates a new ConfigurationParameterPreviewApiResultGroupedReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationParameterPreviewApiResultGroupedReplyWithDefaults

`func NewConfigurationParameterPreviewApiResultGroupedReplyWithDefaults() *ConfigurationParameterPreviewApiResultGroupedReply`

NewConfigurationParameterPreviewApiResultGroupedReplyWithDefaults instantiates a new ConfigurationParameterPreviewApiResultGroupedReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *ConfigurationParameterPreviewApiResultGroupedReply) GetGroups() []ConfigurationParameterPreviewApiResultGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ConfigurationParameterPreviewApiResultGroupedReply) GetGroupsOk() (*[]ConfigurationParameterPreviewApiResultGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ConfigurationParameterPreviewApiResultGroupedReply) SetGroups(v []ConfigurationParameterPreviewApiResultGroup)`

SetGroups sets Groups field to given value.


### GetTotalCount

`func (o *ConfigurationParameterPreviewApiResultGroupedReply) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ConfigurationParameterPreviewApiResultGroupedReply) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ConfigurationParameterPreviewApiResultGroupedReply) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


