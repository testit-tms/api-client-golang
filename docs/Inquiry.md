# Inquiry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**NullableCompositeFilter**](CompositeFilter.md) |  | [optional] 
**Order** | [**[]Order**](Order.md) |  | 
**Page** | Pointer to [**NullablePage**](Page.md) |  | [optional] 

## Methods

### NewInquiry

`func NewInquiry(order []Order, ) *Inquiry`

NewInquiry instantiates a new Inquiry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInquiryWithDefaults

`func NewInquiryWithDefaults() *Inquiry`

NewInquiryWithDefaults instantiates a new Inquiry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *Inquiry) GetFilter() CompositeFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *Inquiry) GetFilterOk() (*CompositeFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *Inquiry) SetFilter(v CompositeFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *Inquiry) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *Inquiry) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *Inquiry) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetOrder

`func (o *Inquiry) GetOrder() []Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Inquiry) GetOrderOk() (*[]Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Inquiry) SetOrder(v []Order)`

SetOrder sets Order field to given value.


### GetPage

`func (o *Inquiry) GetPage() Page`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *Inquiry) GetPageOk() (*Page, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *Inquiry) SetPage(v Page)`

SetPage sets Page field to given value.

### HasPage

`func (o *Inquiry) HasPage() bool`

HasPage returns a boolean if a field has been set.

### SetPageNil

`func (o *Inquiry) SetPageNil(b bool)`

 SetPageNil sets the value for Page to be an explicit nil

### UnsetPage
`func (o *Inquiry) UnsetPage()`

UnsetPage ensures that no value is present for Page, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


