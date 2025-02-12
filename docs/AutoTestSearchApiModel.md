# AutoTestSearchApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**NullableAutoTestFilterApiModel**](AutoTestFilterApiModel.md) | Object containing different filters to adjust search | [optional] 
**Includes** | Pointer to [**NullableAutoTestSearchIncludeApiModel**](AutoTestSearchIncludeApiModel.md) | Object specifying data to be included | [optional] 

## Methods

### NewAutoTestSearchApiModel

`func NewAutoTestSearchApiModel() *AutoTestSearchApiModel`

NewAutoTestSearchApiModel instantiates a new AutoTestSearchApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoTestSearchApiModelWithDefaults

`func NewAutoTestSearchApiModelWithDefaults() *AutoTestSearchApiModel`

NewAutoTestSearchApiModelWithDefaults instantiates a new AutoTestSearchApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *AutoTestSearchApiModel) GetFilter() AutoTestFilterApiModel`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AutoTestSearchApiModel) GetFilterOk() (*AutoTestFilterApiModel, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AutoTestSearchApiModel) SetFilter(v AutoTestFilterApiModel)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AutoTestSearchApiModel) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *AutoTestSearchApiModel) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *AutoTestSearchApiModel) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetIncludes

`func (o *AutoTestSearchApiModel) GetIncludes() AutoTestSearchIncludeApiModel`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *AutoTestSearchApiModel) GetIncludesOk() (*AutoTestSearchIncludeApiModel, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *AutoTestSearchApiModel) SetIncludes(v AutoTestSearchIncludeApiModel)`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *AutoTestSearchApiModel) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.

### SetIncludesNil

`func (o *AutoTestSearchApiModel) SetIncludesNil(b bool)`

 SetIncludesNil sets the value for Includes to be an explicit nil

### UnsetIncludes
`func (o *AutoTestSearchApiModel) UnsetIncludes()`

UnsetIncludes ensures that no value is present for Includes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


