# UpdateParameterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Value** | **string** |  | 
**Name** | **string** |  | 

## Methods

### NewUpdateParameterRequest

`func NewUpdateParameterRequest(value string, name string, ) *UpdateParameterRequest`

NewUpdateParameterRequest instantiates a new UpdateParameterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateParameterRequestWithDefaults

`func NewUpdateParameterRequestWithDefaults() *UpdateParameterRequest`

NewUpdateParameterRequestWithDefaults instantiates a new UpdateParameterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateParameterRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateParameterRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateParameterRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateParameterRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *UpdateParameterRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateParameterRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateParameterRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetName

`func (o *UpdateParameterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateParameterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateParameterRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


