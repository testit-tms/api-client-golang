# CustomAttributeOptionApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the attribute option | 
**IsDeleted** | **bool** | Indicates if the attributes option is deleted | 
**Value** | Pointer to **NullableString** | Value of the attribute option | [optional] 
**IsDefault** | **bool** | Indicates if the attribute option is used by default | 

## Methods

### NewCustomAttributeOptionApiResult

`func NewCustomAttributeOptionApiResult(id string, isDeleted bool, isDefault bool, ) *CustomAttributeOptionApiResult`

NewCustomAttributeOptionApiResult instantiates a new CustomAttributeOptionApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAttributeOptionApiResultWithDefaults

`func NewCustomAttributeOptionApiResultWithDefaults() *CustomAttributeOptionApiResult`

NewCustomAttributeOptionApiResultWithDefaults instantiates a new CustomAttributeOptionApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomAttributeOptionApiResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomAttributeOptionApiResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomAttributeOptionApiResult) SetId(v string)`

SetId sets Id field to given value.


### GetIsDeleted

`func (o *CustomAttributeOptionApiResult) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CustomAttributeOptionApiResult) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CustomAttributeOptionApiResult) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetValue

`func (o *CustomAttributeOptionApiResult) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomAttributeOptionApiResult) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomAttributeOptionApiResult) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomAttributeOptionApiResult) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CustomAttributeOptionApiResult) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CustomAttributeOptionApiResult) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsDefault

`func (o *CustomAttributeOptionApiResult) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CustomAttributeOptionApiResult) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CustomAttributeOptionApiResult) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


