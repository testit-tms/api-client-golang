# IFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**[]IFilter**](IFilter.md) |  | 
**Operator** | [**CollectionOperator**](CollectionOperator.md) |  | 
**Value** | Pointer to **NullableString** |  | [optional] 
**Field** | **string** |  | [readonly] 
**Filter** | [**IFilter**](IFilter.md) |  | 

## Methods

### NewIFilter

`func NewIFilter(filters []IFilter, operator CollectionOperator, field string, filter IFilter, ) *IFilter`

NewIFilter instantiates a new IFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIFilterWithDefaults

`func NewIFilterWithDefaults() *IFilter`

NewIFilterWithDefaults instantiates a new IFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *IFilter) GetFilters() []IFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *IFilter) GetFiltersOk() (*[]IFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *IFilter) SetFilters(v []IFilter)`

SetFilters sets Filters field to given value.


### GetOperator

`func (o *IFilter) GetOperator() CollectionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *IFilter) GetOperatorOk() (*CollectionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *IFilter) SetOperator(v CollectionOperator)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *IFilter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IFilter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IFilter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *IFilter) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *IFilter) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetField

`func (o *IFilter) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *IFilter) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *IFilter) SetField(v string)`

SetField sets Field field to given value.


### GetFilter

`func (o *IFilter) GetFilter() IFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *IFilter) GetFilterOk() (*IFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *IFilter) SetFilter(v IFilter)`

SetFilter sets Filter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


