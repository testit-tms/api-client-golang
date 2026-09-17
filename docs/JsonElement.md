# JsonElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValueKind** | [**JsonValueKind**](JsonValueKind.md) |  | [readonly] 

## Methods

### NewJsonElement

`func NewJsonElement(valueKind JsonValueKind, ) *JsonElement`

NewJsonElement instantiates a new JsonElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonElementWithDefaults

`func NewJsonElementWithDefaults() *JsonElement`

NewJsonElementWithDefaults instantiates a new JsonElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValueKind

`func (o *JsonElement) GetValueKind() JsonValueKind`

GetValueKind returns the ValueKind field if non-nil, zero value otherwise.

### GetValueKindOk

`func (o *JsonElement) GetValueKindOk() (*JsonValueKind, bool)`

GetValueKindOk returns a tuple with the ValueKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueKind

`func (o *JsonElement) SetValueKind(v JsonValueKind)`

SetValueKind sets ValueKind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


