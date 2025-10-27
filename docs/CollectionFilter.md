# CollectionFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**CollectionOperator**](CollectionOperator.md) |  | 
**Filter** | [**IFilter**](IFilter.md) |  | 
**Field** | **string** |  | [readonly] 

## Methods

### NewCollectionFilter

`func NewCollectionFilter(operator CollectionOperator, filter IFilter, field string, ) *CollectionFilter`

NewCollectionFilter instantiates a new CollectionFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionFilterWithDefaults

`func NewCollectionFilterWithDefaults() *CollectionFilter`

NewCollectionFilterWithDefaults instantiates a new CollectionFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *CollectionFilter) GetOperator() CollectionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CollectionFilter) GetOperatorOk() (*CollectionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CollectionFilter) SetOperator(v CollectionOperator)`

SetOperator sets Operator field to given value.


### GetFilter

`func (o *CollectionFilter) GetFilter() IFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CollectionFilter) GetFilterOk() (*IFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CollectionFilter) SetFilter(v IFilter)`

SetFilter sets Filter field to given value.


### GetField

`func (o *CollectionFilter) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *CollectionFilter) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *CollectionFilter) SetField(v string)`

SetField sets Field field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


