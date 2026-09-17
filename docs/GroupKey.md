# GroupKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | [**JsonElement**](JsonElement.md) |  | 
**DisplayValue** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGroupKey

`func NewGroupKey(value JsonElement, ) *GroupKey`

NewGroupKey instantiates a new GroupKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupKeyWithDefaults

`func NewGroupKeyWithDefaults() *GroupKey`

NewGroupKeyWithDefaults instantiates a new GroupKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *GroupKey) GetValue() JsonElement`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GroupKey) GetValueOk() (*JsonElement, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GroupKey) SetValue(v JsonElement)`

SetValue sets Value field to given value.


### GetDisplayValue

`func (o *GroupKey) GetDisplayValue() string`

GetDisplayValue returns the DisplayValue field if non-nil, zero value otherwise.

### GetDisplayValueOk

`func (o *GroupKey) GetDisplayValueOk() (*string, bool)`

GetDisplayValueOk returns a tuple with the DisplayValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayValue

`func (o *GroupKey) SetDisplayValue(v string)`

SetDisplayValue sets DisplayValue field to given value.

### HasDisplayValue

`func (o *GroupKey) HasDisplayValue() bool`

HasDisplayValue returns a boolean if a field has been set.

### SetDisplayValueNil

`func (o *GroupKey) SetDisplayValueNil(b bool)`

 SetDisplayValueNil sets the value for DisplayValue to be an explicit nil

### UnsetDisplayValue
`func (o *GroupKey) UnsetDisplayValue()`

UnsetDisplayValue ensures that no value is present for DisplayValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


