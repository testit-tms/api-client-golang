# CompositeFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | **[]map[string]interface{}** |  | 
**Operator** | [**LogicalOperator**](LogicalOperator.md) |  | 

## Methods

### NewCompositeFilter

`func NewCompositeFilter(filters []map[string]interface{}, operator LogicalOperator, ) *CompositeFilter`

NewCompositeFilter instantiates a new CompositeFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompositeFilterWithDefaults

`func NewCompositeFilterWithDefaults() *CompositeFilter`

NewCompositeFilterWithDefaults instantiates a new CompositeFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *CompositeFilter) GetFilters() []map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CompositeFilter) GetFiltersOk() (*[]map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CompositeFilter) SetFilters(v []map[string]interface{})`

SetFilters sets Filters field to given value.


### GetOperator

`func (o *CompositeFilter) GetOperator() LogicalOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CompositeFilter) GetOperatorOk() (*LogicalOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CompositeFilter) SetOperator(v LogicalOperator)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


